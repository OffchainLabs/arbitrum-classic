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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup/chainlistener"
	"log"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate protoc -I. -I ../.. --go_out=paths=source_relative:. chainobserver.proto

type ChainObserver struct {
	*sync.RWMutex
	NodeGraph           *nodegraph.StakedNodeGraph
	rollupAddr          common.Address
	Inbox               *structures.Inbox
	KnownValidNode      *structures.Node
	calculatedValidNode *structures.Node
	LatestBlockId       *common.BlockId
	listeners           []chainlistener.ChainListener
	checkpointer        checkpointing.RollupCheckpointer
	isOpinionated       bool
	atHead              bool
	pendingState        machine.Machine
}

func NewChain(
	rollupAddr common.Address,
	checkpointer checkpointing.RollupCheckpointer,
	vmParams valprotocol.ChainParams,
	updateOpinion bool,
	startBlockId *common.BlockId,
	creationTxHash common.Hash,
) (*ChainObserver, error) {
	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	nodeGraph := nodegraph.NewStakedNodeGraph(mach, vmParams, creationTxHash)
	ret := &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		NodeGraph:           nodeGraph,
		rollupAddr:          rollupAddr,
		Inbox:               structures.NewInbox(),
		KnownValidNode:      nodeGraph.LatestConfirmed(),
		calculatedValidNode: nodeGraph.LatestConfirmed(),
		LatestBlockId:       startBlockId,
		listeners:           []chainlistener.ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       false,
		atHead:              false,
	}
	ret.Lock()
	defer ret.Unlock()
	if updateOpinion {
		ret.isOpinionated = true
	}
	return ret, nil
}

func InitializeChainObserver(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ChainTimeGetter,
	watcher arbbridge.ArbRollupWatcher,
	checkpointer checkpointing.RollupCheckpointer,
) (*ChainObserver, error) {
	if checkpointer.HasCheckpointedState() {
		var chain *ChainObserver
		if err := checkpointer.RestoreLatestState(ctx, clnt, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext) error {
			chainObserverBuf := &ChainObserverBuf{}
			if err := proto.Unmarshal(chainObserverBytes, chainObserverBuf); err != nil {
				return err
			}
			var err error
			chain, err = chainObserverBuf.UnmarshalFromCheckpoint(restoreCtx, checkpointer)
			return err
		}); err == nil && chain != nil {
			return chain, nil
		}
	}

	log.Println("No valid checkpoints so starting from fresh state")
	params, err := watcher.GetParams(ctx)
	if err != nil {
		log.Fatal(err)
	}
	creationHash, blockId, _, err := watcher.GetCreationInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return NewChain(rollupAddr, checkpointer, params, updateOpinion, blockId, creationHash)
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
				confOpp, _ := chain.NodeGraph.GenerateNextConfProof(common.TicksFromBlockNum(chain.LatestBlockId.Height))
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
				chain.ExecutionPrecondition(challenge.ConflictNode()),
				challenge)

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
				prunesToDo := chain.NodeGraph.GenerateNodePruneInfo(chain.NodeGraph.Stakers())
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
	listener.AddedToChain(ctx, chain.PendingCorrectNodes())
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
		LatestBlockId:       chain.LatestBlockId.MarshalToBuf(),
		IsOpinionated:       chain.isOpinionated,
	}
}

func (x *ChainObserverBuf) UnmarshalFromCheckpoint(
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
		RWMutex:             &sync.RWMutex{},
		NodeGraph:           nodeGraph,
		rollupAddr:          x.ContractAddress.Unmarshal(),
		Inbox:               &structures.Inbox{MessageStack: inbox},
		KnownValidNode:      knownValidNode,
		calculatedValidNode: calculatedValidNode,
		LatestBlockId:       x.LatestBlockId.Unmarshal(),
		listeners:           []chainlistener.ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       x.IsOpinionated,
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
		chain.messageDelivered(ctx, ev)
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
		return chain.confirmNode(ctx, ev)
	}
	return nil
}

func (chain *ChainObserver) NotifyNewBlock(blockId *common.BlockId) {
	chain.Lock()
	defer chain.Unlock()
	chain.LatestBlockId = blockId
	ckptCtx := ckptcontext.NewCheckpointContext()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockId.Clone(), buf, ckptCtx)
}

func (chain *ChainObserver) CurrentBlockId() *common.BlockId {
	chain.RLock()
	defer chain.RUnlock()
	return chain.LatestBlockId.Clone()
}

func (chain *ChainObserver) ContractAddress() common.Address {
	chain.RLock()
	defer chain.RUnlock()
	return chain.rollupAddr
}

func (chain *ChainObserver) LatestKnownValidMachine() machine.Machine {
	chain.RLock()
	defer chain.RUnlock()
	return chain.calculatedValidNode.Machine().Clone()
}

func (chain *ChainObserver) CurrentPendingMachine() machine.Machine {
	chain.RLock()
	defer chain.RUnlock()
	if chain.pendingState == nil {
		return chain.calculatedValidNode.Machine().Clone()
	}
	return chain.pendingState.Clone()
}

func (chain *ChainObserver) RestartFromLatestValid(ctx context.Context) {
	chain.RLock()
	defer chain.RUnlock()
	for _, lis := range chain.listeners {
		lis.RestartingFromLatestValid(ctx, chain.calculatedValidNode)
	}
}

func (chain *ChainObserver) PendingCorrectNodes() []*structures.Node {
	chain.RLock()
	defer chain.RUnlock()
	var nodes []*structures.Node
	for node := chain.calculatedValidNode; node != chain.NodeGraph.LatestConfirmed(); node = node.Prev() {
		nodes = append(nodes, node)
	}
	return nodes
}

func (chain *ChainObserver) messageDelivered(ctx context.Context, ev arbbridge.MessageDeliveredEvent) {
	chain.Inbox.DeliverMessage(ev.Message)
	for _, lis := range chain.listeners {
		lis.MessageDelivered(ctx, ev)
	}
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
			chain.ExecutionPrecondition(challenge.ConflictNode()),
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
	ckptCtx := ckptcontext.NewCheckpointContext()

	// Note that we marshal the node without including the machine state
	// of that node. This is because saving the machine state for every confirmed
	// node would significantly bloat the database and there is currently
	// no usecase for this data
	nodeData := confirmedNode.MarshalForCheckpoint(ckptCtx, false)
	nodeBytes, err := proto.Marshal(nodeData)
	if err != nil {
		return err
	}
	if err := chain.checkpointer.CheckpointConfirmedNode(
		confirmedNode.Hash(),
		confirmedNode.Depth(),
		nodeBytes,
		ckptCtx,
	); err != nil {
		return err
	}

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
	chain.NodeGraph.CreateNodesOnAssert(
		chain.NodeGraph.NodeFromHash(ev.PrevLeafHash),
		ev.Disputable,
		ev.BlockId.Height,
		ev.TxHash,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ctx, ev)
	}
}

func (chain *ChainObserver) equals(co2 *ChainObserver) bool {
	return chain.NodeGraph.Equals(co2.NodeGraph) &&
		bytes.Compare(chain.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		chain.Inbox.Equals(co2.Inbox)
}

func (chain *ChainObserver) ExecutionPrecondition(node *structures.Node) *valprotocol.Precondition {
	vmProtoData := node.Prev().VMProtoData()
	params := node.Disputable().AssertionParams
	inbox, _ := chain.Inbox.GenerateVMInbox(vmProtoData.InboxTop, params.ImportedMessageCount.Uint64())
	return &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		BeforeInbox: inbox.AsValue(),
	}
}
