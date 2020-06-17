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

package rollup

import (
	"bytes"
	"context"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"log"
	"math/big"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate protoc -I. -I ../.. --go_out=paths=source_relative:. rollup.proto

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph           *StakedNodeGraph
	rollupAddr          common.Address
	inbox               *structures.Inbox
	knownValidNode      *structures.Node
	calculatedValidNode *structures.Node
	latestBlockId       *common.BlockId
	listeners           []ChainListener
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
	nodeGraph := NewStakedNodeGraph(mach, vmParams, creationTxHash)
	ret := &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          rollupAddr,
		inbox:               structures.NewInbox(),
		knownValidNode:      nodeGraph.latestConfirmed,
		calculatedValidNode: nodeGraph.latestConfirmed,
		latestBlockId:       startBlockId,
		listeners:           []ChainListener{},
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

func (chain *ChainObserver) Start(ctx context.Context) {
	chain.nodeGraph.challenges.forall(func(c *Challenge) {
		for _, listener := range chain.listeners {
			listener.ResumedChallenge(ctx, chain, c)
		}
	})
	chain.startCleanupThread(ctx)
	chain.startConfirmThread(ctx)

	if chain.isOpinionated {
		chain.startOpinionUpdateThread(ctx)
	}
}

func (chain *ChainObserver) AddListener(ctx context.Context, listener ChainListener) {
	chain.Lock()
	chain.listeners = append(chain.listeners, listener)
	chain.Unlock()
	chain.RLock()
	listener.AddedToChain(ctx, chain)
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
	return chain.nodeGraph.params
}

func (chain *ChainObserver) marshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph:     chain.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress:     chain.rollupAddr.MarshallToBuf(),
		Inbox:               chain.inbox.MarshalForCheckpoint(ctx),
		KnownValidNode:      chain.knownValidNode.Hash().MarshalToBuf(),
		CalculatedValidNode: chain.calculatedValidNode.Hash().MarshalToBuf(),
		LatestBlockId:       chain.latestBlockId.MarshalToBuf(),
		IsOpinionated:       chain.isOpinionated,
	}
}

func (chain *ChainObserver) marshalToBytes(ctx *ckptcontext.CheckpointContext) ([]byte, error) {
	cob := chain.marshalForCheckpoint(ctx)
	return proto.Marshal(cob)
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
	knownValidNode := nodeGraph.nodeFromHash[x.KnownValidNode.Unmarshal()]
	if knownValidNode == nil {
		return nil, fmt.Errorf("knownValidNode %v was nil", x.KnownValidNode.Unmarshal())
	}

	calculatedValidNode := nodeGraph.nodeFromHash[x.CalculatedValidNode.Unmarshal()]
	if calculatedValidNode == nil {
		return nil, fmt.Errorf("calculatedValidNode %v was nil", x.CalculatedValidNode.Unmarshal())
	}

	return &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          x.ContractAddress.Unmarshal(),
		inbox:               &structures.Inbox{MessageStack: inbox},
		knownValidNode:      knownValidNode,
		calculatedValidNode: calculatedValidNode,
		latestBlockId:       x.LatestBlockId.Unmarshal(),
		listeners:           []ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       x.IsOpinionated,
		atHead:              false,
	}, nil
}

func (chain *ChainObserver) DebugString(prefix string) string {
	chain.Lock()
	defer chain.Unlock()
	labels := make(map[*structures.Node][]string)
	labels[chain.calculatedValidNode] = append(labels[chain.calculatedValidNode], "calculatedValidNode")
	labels[chain.knownValidNode] = append(labels[chain.knownValidNode], "knownValidNode")
	return chain.nodeGraph.DebugString(prefix, labels)
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
	chain.latestBlockId = blockId
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
	return chain.latestBlockId.Clone()
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
		lis.RestartingFromLatestValid(ctx, chain, chain.calculatedValidNode)
	}
}

func (chain *ChainObserver) PendingCorrectNodes() []*structures.Node {
	chain.RLock()
	defer chain.RUnlock()
	var nodes []*structures.Node
	for node := chain.calculatedValidNode; node != chain.nodeGraph.latestConfirmed; node = node.Prev() {
		nodes = append(nodes, node)
	}
	return nodes
}

func (chain *ChainObserver) messageDelivered(ctx context.Context, ev arbbridge.MessageDeliveredEvent) {
	chain.inbox.DeliverMessage(ev.Message)
	for _, lis := range chain.listeners {
		lis.MessageDelivered(ctx, chain, ev)
	}
}

func (chain *ChainObserver) pruneLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
	leaf, found := chain.nodeGraph.nodeFromHash[ev.Leaf]
	if !found {
		panic("Tried to prune nonexistant leaf")
	}
	chain.nodeGraph.leaves.Delete(leaf)
	chain.nodeGraph.PruneNodeByHash(ev.Leaf)
	for _, lis := range chain.listeners {
		lis.PrunedLeaf(ctx, chain, ev)
	}
	chain.updateOldest()
}

func (chain *ChainObserver) createStake(ctx context.Context, ev arbbridge.StakeCreatedEvent) {
	chain.nodeGraph.CreateStake(ev)
	for _, lis := range chain.listeners {
		lis.StakeCreated(ctx, chain, ev)
	}
}

func (chain *ChainObserver) removeStake(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	chain.nodeGraph.RemoveStake(ev.Staker)
	for _, lis := range chain.listeners {
		lis.StakeRemoved(ctx, chain, ev)
	}
}

func (chain *ChainObserver) moveStake(ctx context.Context, ev arbbridge.StakeMovedEvent) {
	chain.nodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range chain.listeners {
		lis.StakeMoved(ctx, chain, ev)
	}
}

func (chain *ChainObserver) newChallenge(ctx context.Context, ev arbbridge.ChallengeStartedEvent) {
	asserter := chain.nodeGraph.stakers.Get(ev.Asserter)
	challenger := chain.nodeGraph.stakers.Get(ev.Challenger)
	_, challengerAncestor, err := structures.GetConflictAncestor(asserter.location, challenger.location)
	if err != nil {
		panic("No conflict ancestor for conflict")
	}
	challenge := &Challenge{
		blockId:      ev.BlockId,
		logIndex:     ev.LogIndex,
		asserter:     ev.Asserter,
		challenger:   ev.Challenger,
		contract:     ev.ChallengeContract,
		conflictNode: challengerAncestor,
	}

	chain.nodeGraph.NewChallenge(challenge)
	for _, lis := range chain.listeners {
		lis.StartedChallenge(ctx, chain, challenge)
	}
}

func (chain *ChainObserver) challengeResolved(ctx context.Context, ev arbbridge.ChallengeCompletedEvent) {
	chain.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ctx, chain, ev)
	}
}

func (chain *ChainObserver) confirmNode(ctx context.Context, ev arbbridge.ConfirmedEvent) error {
	confirmedNode := chain.nodeGraph.nodeFromHash[ev.NodeHash]
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

	if confirmedNode.Depth() > chain.knownValidNode.Depth() {
		chain.knownValidNode = confirmedNode
	}
	chain.nodeGraph.latestConfirmed = confirmedNode
	chain.nodeGraph.considerPruningNode(confirmedNode.Prev())

	chain.updateOldest()
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ctx, chain, ev)
	}
	return nil
}

func (chain *ChainObserver) updateOldest() {
	for chain.nodeGraph.oldestNode != chain.nodeGraph.latestConfirmed {
		if chain.nodeGraph.oldestNode.NumStakers() > 0 {
			return
		}
		if chain.calculatedValidNode == chain.nodeGraph.oldestNode {
			return
		}
		var successor *structures.Node
		for _, successorHash := range chain.nodeGraph.oldestNode.SuccessorHashes() {
			if successorHash != zeroBytes32 {
				if successor != nil {
					return
				}
				successor = chain.nodeGraph.nodeFromHash[successorHash]
			}
		}
		chain.nodeGraph.pruneOldestNode(chain.nodeGraph.oldestNode)
		chain.nodeGraph.oldestNode = successor
	}
}

func (chain *ChainObserver) notifyAssert(ctx context.Context, ev arbbridge.AssertedEvent) {
	chain.nodeGraph.CreateNodesOnAssert(
		chain.nodeGraph.nodeFromHash[ev.PrevLeafHash],
		ev.Disputable,
		ev.BlockId.Height,
		ev.TxHash,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ctx, chain, ev)
	}
}

func (chain *ChainObserver) equals(co2 *ChainObserver) bool {
	return chain.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(chain.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		chain.inbox.Equals(co2.inbox)
}

func (chain *ChainObserver) executionPrecondition(node *structures.Node) *valprotocol.Precondition {
	vmProtoData := node.Prev().VMProtoData()
	params := node.Disputable().AssertionParams
	inbox, _ := chain.inbox.GenerateVMInbox(vmProtoData.InboxTop, params.ImportedMessageCount.Uint64())
	return &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  params.TimeBounds,
		BeforeInbox: inbox.AsValue(),
	}
}

func (chain *ChainObserver) currentTimeBounds() *protocol.TimeBounds {
	latestBlock := chain.latestBlockId.Height
	// Start timestamp slightly in the past to avoid it being invalid
	latestTimestamp := time.Now().Unix() - 60
	return &protocol.TimeBounds{
		LowerBoundBlock:     latestBlock,
		UpperBoundBlock:     common.NewTimeBlocks(new(big.Int).Add(latestBlock.AsInt(), big.NewInt(int64(chain.nodeGraph.params.MaxBlockBoundsWidth)))),
		LowerBoundTimestamp: big.NewInt(latestTimestamp),
		UpperBoundTimestamp: big.NewInt(latestTimestamp + int64(chain.nodeGraph.params.MaxTimestampBoundsWidth)),
	}
}
