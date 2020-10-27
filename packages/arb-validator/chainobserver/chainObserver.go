/*
* Copyright 2019-2020, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package chainobserver

import (
	"bytes"
	"context"
	"fmt"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate protoc -I. -I ../.. --go_out=paths=source_relative:. chainobserver.proto

type ChainObserver struct {
	sync.RWMutex
	NodeGraph           *nodegraph.StakedNodeGraph
	rollupAddr          common.Address
	Inbox               *structures.Inbox
	KnownValidNode      *structures.Node
	calculatedValidNode *structures.Node
	currentEventId      arbbridge.ChainInfo
	// assumedValidBlock is used when making assertions so that we don't include messages past the threshold
	// If there is a reorg past that depth, our assertion could be invalidated. Invalidation isn't a problem
	// and would only lose the validator a small amount of gas money, so this assumption being violated is safe.
	assumedValidBlock *common.BlockId
	listeners         []chainlistener.ChainListener
	checkpointer      checkpointing.RollupCheckpointer
	isOpinionated     bool
	atHead            bool
}

func tryRestoreFromCheckpoint(
	ctx context.Context,
	clnt arbbridge.ChainTimeGetter,
	checkpointer checkpointing.RollupCheckpointer,
) (*ChainObserver, *common.BlockId) {
	var chain *ChainObserver
	var id *common.BlockId
	err := checkpointer.RestoreLatestState(ctx, clnt, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext, blockId *common.BlockId) error {
		chainObserverBuf := &ChainObserverBuf{}
		if err := proto.Unmarshal(chainObserverBytes, chainObserverBuf); err != nil {
			return err
		}
		var err error
		chain, err = chainObserverBuf.unmarshalFromCheckpoint(restoreCtx, checkpointer)
		if err != nil {
			return err
		}
		id = blockId
		return nil
	})
	if err != nil || chain == nil {
		return nil, nil
	}
	return chain, id
}

func NewChainObserver(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ChainTimeGetter,
	watcher arbbridge.ArbRollupWatcher,
	checkpointer checkpointing.RollupCheckpointer,
	assumedValidDepth int64,
) (*ChainObserver, error) {
	var chain *ChainObserver
	var currentEventId arbbridge.ChainInfo
	if checkpointer.HasCheckpointedState() {
		restoredChain, restoredBlockId := tryRestoreFromCheckpoint(ctx, clnt, checkpointer)
		currentEventId = arbbridge.ChainInfo{
			BlockId:  restoredBlockId,
			LogIndex: 0,
		}
		chain = restoredChain
	}
	if chain == nil {
		params, err := watcher.GetParams(ctx)
		if err != nil {
			return nil, err
		}

		_, createdEventId, _, _, err := watcher.GetCreationInfo(ctx)
		if err != nil {
			return nil, err
		}
		currentEventId = createdEventId
		currentEventId.LogIndex++

		chain, err = newChain(
			rollupAddr,
			checkpointer,
			params,
		)
		if err != nil {
			return nil, err
		}
	}

	chain.isOpinionated = updateOpinion
	chain.currentEventId = currentEventId

	if err := chain.UpdateAssumedValidBlock(ctx, clnt, assumedValidDepth); err != nil {
		return nil, err
	}

	return chain, nil
}

func (chain *ChainObserver) startConfirmThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				chain.RLock()
				if !chain.atHead {
					chain.RUnlock()
					break
				}

				confOpp, _ := chain.NodeGraph.GenerateNextConfProof(common.TicksFromBlockNum(chain.currentEventId.BlockId.Height))
				if confOpp != nil {
					for _, listener := range chain.listeners {
						listener.ConfirmableNodes(ctx, confOpp)
					}
				}
				chain.RUnlock()
			}
		}
	}()
}

func (chain *ChainObserver) Start(ctx context.Context) {
	chain.NodeGraph.Challenges.Forall(func(challenge *nodegraph.Challenge) {
		for _, listener := range chain.listeners {
			listener.ResumedChallenge(
				ctx,
				chain.Inbox.MessageStack,
				challenge,
			)
		}
	})
	chain.startCleanupThread(ctx)
	chain.startConfirmThread(ctx)

	if chain.isOpinionated {
		chain.startOpinionUpdateThread(ctx)
	}
}

func (chain *ChainObserver) startCleanupThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				chain.RLock()
				if !chain.atHead {
					chain.RUnlock()
					break
				}
				prunesToDo := chain.NodeGraph.GenerateNodePruneInfo()
				mootedToDo, oldToDo := chain.NodeGraph.GenerateStakerPruneInfo()
				chain.RUnlock()

				if len(prunesToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.PrunableLeafs(ctx, prunesToDo)
					}
				}
				if len(mootedToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.MootableStakes(ctx, mootedToDo)
					}
				}
				if len(oldToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.OldStakes(ctx, oldToDo)
					}
				}

			}
		}
	}()
}

func (chain *ChainObserver) AddListener(
	ctx context.Context,
	listener chainlistener.ChainListener,
) {
	chain.Lock()
	chain.listeners = append(chain.listeners, listener)
	chain.Unlock()
	chain.RLock()
	listener.AddedToChain(ctx, chain.pendingCorrectNodes())
	chain.RUnlock()
}

func (chain *ChainObserver) NowAtHead() {
	chain.Lock()
	defer chain.Unlock()
	chain.atHead = true
}

func (chain *ChainObserver) IsAtHead() bool {
	chain.RLock()
	defer chain.RUnlock()
	return chain.atHead
}

func (chain *ChainObserver) GetChainParams() valprotocol.ChainParams {
	chain.RLock()
	defer chain.RUnlock()
	return chain.NodeGraph.Params()
}

func (chain *ChainObserver) marshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph:     chain.NodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress:     chain.rollupAddr.MarshallToBuf(),
		Inbox:               chain.Inbox.MarshalForCheckpoint(ctx),
		KnownValidNode:      chain.KnownValidNode.Hash().MarshalToBuf(),
		CalculatedValidNode: chain.calculatedValidNode.Hash().MarshalToBuf(),
	}
}

func (x *ChainObserverBuf) unmarshalFromCheckpoint(
	restoreCtx ckptcontext.RestoreContext,
	checkpointer checkpointing.RollupCheckpointer,
) (*ChainObserver, error) {
	nodeGraph, err := x.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx)
	if err != nil {
		return nil, err
	}
	inbox, err := x.Inbox.UnmarshalFromCheckpoint(restoreCtx)
	if err != nil {
		return nil, err
	}
	knownValidNode := nodeGraph.NodeFromHash(x.KnownValidNode.Unmarshal())
	if knownValidNode == nil {
		return nil, fmt.Errorf("knownValidNode %v was nil", x.KnownValidNode.Unmarshal())
	}

	calculatedValidNode := nodeGraph.NodeFromHash(x.CalculatedValidNode.Unmarshal())
	if calculatedValidNode == nil {
		return nil, fmt.Errorf("calculatedValidNode %v was nil", x.CalculatedValidNode.Unmarshal())
	}

	return &ChainObserver{
		NodeGraph:           nodeGraph,
		rollupAddr:          x.ContractAddress.Unmarshal(),
		Inbox:               &structures.Inbox{MessageStack: inbox},
		KnownValidNode:      knownValidNode,
		calculatedValidNode: calculatedValidNode,
		listeners:           []chainlistener.ChainListener{},
		checkpointer:        checkpointer,
		atHead:              false,
	}, nil
}

func newChain(
	rollupAddr common.Address,
	checkpointer checkpointing.RollupCheckpointer,
	vmParams valprotocol.ChainParams,
) (*ChainObserver, error) {
	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return nil, err
	}

	mach, err := checkpointer.GetInitialMachine(valueCache)
	if err != nil {
		return nil, err
	}
	nodeGraph := nodegraph.NewStakedNodeGraph(mach, vmParams)
	return &ChainObserver{
		NodeGraph:           nodeGraph,
		rollupAddr:          rollupAddr,
		Inbox:               structures.NewInbox(),
		KnownValidNode:      nodeGraph.LatestConfirmed(),
		calculatedValidNode: nodeGraph.LatestConfirmed(),
		listeners:           []chainlistener.ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       true,
		atHead:              false,
	}, nil
}

func (chain *ChainObserver) marshalToBytes(ctx *ckptcontext.CheckpointContext) ([]byte, error) {
	cob := chain.marshalForCheckpoint(ctx)
	return proto.Marshal(cob)
}

func (chain *ChainObserver) DebugString(prefix string) string {
	chain.Lock()
	defer chain.Unlock()
	labels := make(map[*structures.Node][]string)
	labels[chain.calculatedValidNode] = append(labels[chain.calculatedValidNode], "calculatedValidNode")
	labels[chain.KnownValidNode] = append(labels[chain.KnownValidNode], "knownValidNode")
	return chain.NodeGraph.DebugString(prefix, labels)
}

func (chain *ChainObserver) HandleNotification(ctx context.Context, event arbbridge.Event) error {
	chain.Lock()
	defer chain.Unlock()
	switch ev := event.(type) {
	case arbbridge.MessageDeliveredEvent:
		if err := chain.messageDelivered(ctx, ev); err != nil {
			return err
		}
	case arbbridge.StakeCreatedEvent:
		chain.createStake(ctx, ev)
	case arbbridge.ChallengeStartedEvent:
		chain.newChallenge(ctx, ev)
	case arbbridge.ChallengeCompletedEvent:
		chain.challengeResolved(ctx, ev)
	case arbbridge.StakeRefundedEvent:
		chain.removeStake(ctx, ev)
	case arbbridge.PrunedEvent:
		chain.pruneLeaf(ctx, ev)
	case arbbridge.StakeMovedEvent:
		chain.moveStake(ctx, ev)
	case arbbridge.AssertedEvent:
		chain.notifyAssert(ctx, ev)
	case arbbridge.ConfirmedEvent:
		if err := chain.confirmNode(ctx, ev); err != nil {
			return err
		}
	}
	chain.currentEventId = arbbridge.ChainInfo{
		BlockId:  event.GetChainInfo().BlockId.Clone(),
		LogIndex: event.GetChainInfo().LogIndex + 1,
	}
	return nil
}

func (chain *ChainObserver) NotifyNextEvent(blockId *common.BlockId) {
	chain.Lock()
	defer chain.Unlock()
	chain.currentEventId = arbbridge.ChainInfo{
		BlockId:  blockId.Clone(),
		LogIndex: 0,
	}
}

func (chain *ChainObserver) UpdateAssumedValidBlock(ctx context.Context, clnt arbbridge.ChainTimeGetter, assumedValidDepth int64) error {
	latestL1BlockId, err := clnt.BlockIdForHeight(ctx, nil)
	if err != nil {
		return errors2.Wrap(err, "Getting current block header")
	}

	validHeight := new(big.Int).Sub(latestL1BlockId.Height.AsInt(), big.NewInt(assumedValidDepth))
	if validHeight.Cmp(big.NewInt(0)) < 0 {
		validHeight = big.NewInt(0)
	}
	assumedValidBlock, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocks(validHeight))
	if err != nil {
		return errors2.Wrapf(err, "Getting assumed valid block header at height %v", validHeight)
	}

	chain.Lock()
	defer chain.Unlock()
	chain.assumedValidBlock = assumedValidBlock
	return nil
}

func (chain *ChainObserver) NotifyNewBlock(blockId *common.BlockId) {
	chain.Lock()
	defer chain.Unlock()
	ckptCtx := ckptcontext.NewCheckpointContext()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockId.Clone(), buf, ckptCtx)
}

func (chain *ChainObserver) CurrentEventId() arbbridge.ChainInfo {
	chain.RLock()
	defer chain.RUnlock()
	return arbbridge.ChainInfo{
		BlockId:  chain.currentEventId.BlockId.Clone(),
		LogIndex: chain.currentEventId.LogIndex,
	}
}

func (chain *ChainObserver) ContractAddress() common.Address {
	chain.RLock()
	defer chain.RUnlock()
	return chain.rollupAddr
}

func (chain *ChainObserver) RestartFromLatestValid(ctx context.Context) {
	chain.RLock()
	defer chain.RUnlock()
	for _, lis := range chain.listeners {
		lis.RestartingFromLatestValid(ctx, chain.calculatedValidNode)
	}
}

func (chain *ChainObserver) pendingCorrectNodes() []*structures.Node {
	var nodes []*structures.Node
	for node := chain.calculatedValidNode; node != chain.NodeGraph.LatestConfirmed(); node = node.Prev() {
		nodes = append(nodes, node)
	}
	return nodes
}

func (chain *ChainObserver) messageDelivered(ctx context.Context, ev arbbridge.MessageDeliveredEvent) error {
	if err := chain.Inbox.DeliverMessage(ev.Message); err != nil {
		return err
	}
	for _, lis := range chain.listeners {
		lis.MessageDelivered(ctx, ev)
	}
	return nil
}

func (chain *ChainObserver) pruneLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
	leaf := chain.NodeGraph.NodeFromHash(ev.Leaf)
	if leaf == nil {
		panic("Tried to prune nonexistant leaf")
	}
	chain.NodeGraph.DeleteLeaf(leaf)
	chain.NodeGraph.PruneNodeByHash(ev.Leaf)
	for _, lis := range chain.listeners {
		lis.PrunedLeaf(ctx, ev)
	}
	chain.updateOldest()
}

func (chain *ChainObserver) createStake(ctx context.Context, ev arbbridge.StakeCreatedEvent) {
	chain.NodeGraph.CreateStake(ev)
	for _, listener := range chain.listeners {
		listener.StakeCreated(ctx, chain.NodeGraph, ev)
	}
}

func (chain *ChainObserver) removeStake(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	chain.NodeGraph.RemoveStake(ev.Staker)
	for _, lis := range chain.listeners {
		lis.StakeRemoved(ctx, ev)
	}
}

func (chain *ChainObserver) moveStake(ctx context.Context, ev arbbridge.StakeMovedEvent) {
	chain.NodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range chain.listeners {
		lis.StakeMoved(ctx, chain.NodeGraph, ev)
	}
}

func (chain *ChainObserver) newChallenge(ctx context.Context, ev arbbridge.ChallengeStartedEvent) {
	asserter := chain.NodeGraph.Stakers().Get(ev.Asserter)
	challenger := chain.NodeGraph.Stakers().Get(ev.Challenger)
	_, challengerAncestor, err := structures.GetConflictAncestor(asserter.Location(), challenger.Location())
	if err != nil {
		panic("No conflict ancestor for conflict")
	}
	challenge := nodegraph.NewChallengeFromEvent(ev, challengerAncestor)

	chain.NodeGraph.NewChallenge(challenge)
	for _, lis := range chain.listeners {
		lis.StartedChallenge(
			ctx,
			chain.Inbox.MessageStack,
			challenge)
	}
}

func (chain *ChainObserver) challengeResolved(ctx context.Context, ev arbbridge.ChallengeCompletedEvent) {
	chain.NodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ctx, chain.NodeGraph, ev)
	}
}

func (chain *ChainObserver) confirmNode(ctx context.Context, ev arbbridge.ConfirmedEvent) error {
	confirmedNode := chain.NodeGraph.NodeFromHash(ev.NodeHash)
	if confirmedNode.Depth() > chain.KnownValidNode.Depth() {
		chain.KnownValidNode = confirmedNode
	}
	chain.NodeGraph.UpdateLatestConfirmed(confirmedNode)
	chain.NodeGraph.ConsiderPruningNode(confirmedNode.Prev())

	chain.updateOldest()
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ctx, ev)
	}
	return nil
}

func (chain *ChainObserver) updateOldest() {
	// Don't update the oldest more than 1 block behind the current latest confirmed node
	for chain.NodeGraph.OldestNode().Depth()+1 < chain.NodeGraph.LatestConfirmed().Depth() {
		if chain.NodeGraph.OldestNode().NumStakers() > 0 {
			return
		}
		if chain.calculatedValidNode == chain.NodeGraph.OldestNode() {
			return
		}
		var successor *structures.Node
		for _, successorHash := range chain.NodeGraph.OldestNode().SuccessorHashes() {
			if successorHash != nodegraph.ZeroBytes32 {
				if successor != nil {
					return
				}
				successor = chain.NodeGraph.NodeFromHash(successorHash)
			}
		}
		chain.NodeGraph.PruneOldestNode(chain.NodeGraph.OldestNode())
		chain.NodeGraph.UpdateOldestNode(successor)
	}
}

func (chain *ChainObserver) notifyAssert(ctx context.Context, ev arbbridge.AssertedEvent) {
	prevNode := chain.NodeGraph.NodeFromHash(ev.PrevLeafHash)
	disNode := valprotocol.NewDisputableNode(
		ev.AssertionParams,
		&valprotocol.ExecutionAssertionStub{
			NumGas:            ev.NumGas,
			BeforeMachineHash: prevNode.VMProtoData().MachineHash,
			AfterMachineHash:  ev.AfterMachineHash,
			BeforeInboxHash:   prevNode.VMProtoData().InboxTop,
			AfterInboxHash:    ev.AfterInboxHash,
			FirstMessageHash:  common.Hash{},
			LastMessageHash:   ev.LastMessageHash,
			MessageCount:      ev.MessageCount,
			FirstLogHash:      common.Hash{},
			LastLogHash:       ev.LastLogHash,
			LogCount:          ev.LogCount,
		},
		ev.MaxInboxTop,
		ev.MaxInboxCount,
	)
	chain.NodeGraph.CreateNodesOnAssert(
		chain.NodeGraph.NodeFromHash(ev.PrevLeafHash),
		disNode,
		ev.BlockId.Height,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ctx, ev)
	}
}

func (chain *ChainObserver) equals(co2 *ChainObserver) bool {
	return chain.NodeGraph.Equals(co2.NodeGraph) &&
		bytes.Equal(chain.rollupAddr[:], co2.rollupAddr[:]) &&
		chain.Inbox.Equals(co2.Inbox)
}
