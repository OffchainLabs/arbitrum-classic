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
	"log"
	"math/big"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/golang/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph           *StakedNodeGraph
	rollupAddr          common.Address
	pendingInbox        *structures.PendingInbox
	knownValidNode      *Node
	calculatedValidNode *Node
	latestBlockID       *structures.BlockID
	listeners           []ChainListener
	checkpointer        checkpointing.RollupCheckpointer
	isOpinionated       bool
	atHead              bool
}

func NewChain(
	rollupAddr common.Address,
	checkpointer checkpointing.RollupCheckpointer,
	vmParams structures.ChainParams,
	updateOpinion bool,
	startBlockID *structures.BlockID,
) (*ChainObserver, error) {
	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	nodeGraph := NewStakedNodeGraph(mach, vmParams)
	ret := &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          rollupAddr,
		pendingInbox:        structures.NewPendingInbox(),
		knownValidNode:      nodeGraph.latestConfirmed,
		calculatedValidNode: nodeGraph.latestConfirmed,
		latestBlockID:       startBlockID,
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

func (co *ChainObserver) Start(ctx context.Context) {
	co.startCleanupThread(ctx)
	co.startConfirmThread(ctx)

	if co.isOpinionated {
		co.startOpinionUpdateThread(ctx)
	}
}

func (co *ChainObserver) AddListener(listener ChainListener) {
	co.Lock()
	co.listeners = append(co.listeners, listener)
	co.Unlock()
}

func (co *ChainObserver) NowAtHead() {
	co.Lock()
	co.atHead = true
	co.Unlock()
}

func (co *ChainObserver) marshalForCheckpoint(ctx structures.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph:     co.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress:     co.rollupAddr.MarshallToBuf(),
		PendingInbox:        co.pendingInbox.MarshalForCheckpoint(ctx),
		KnownValidNode:      co.knownValidNode.hash.MarshalToBuf(),
		CalculatedValidNode: co.calculatedValidNode.hash.MarshalToBuf(),
		LatestBlockID:       co.latestBlockID.MarshalToBuf(),
		IsOpinionated:       co.isOpinionated,
	}
}

func (co *ChainObserver) marshalToBytes(ctx structures.CheckpointContext) ([]byte, error) {
	cob := co.marshalForCheckpoint(ctx)
	return proto.Marshal(cob)
}

func (m *ChainObserverBuf) UnmarshalFromCheckpoint(
	ctx context.Context,
	restoreCtx structures.RestoreContext,
	checkpointer checkpointing.RollupCheckpointer,
) (*ChainObserver, error) {
	nodeGraph := m.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx)
	pendingInbox, err := m.PendingInbox.UnmarshalFromCheckpoint(restoreCtx)
	if err != nil {
		return nil, err
	}
	return &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          m.ContractAddress.Unmarshal(),
		pendingInbox:        &structures.PendingInbox{MessageStack: pendingInbox},
		knownValidNode:      nodeGraph.nodeFromHash[m.KnownValidNode.Unmarshal()],
		calculatedValidNode: nodeGraph.nodeFromHash[m.CalculatedValidNode.Unmarshal()],
		latestBlockID:       m.LatestBlockID.Unmarshal(),
		listeners:           []ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       m.IsOpinionated,
		atHead:              false,
	}, nil
}

func (co *ChainObserver) DebugString(prefix string) string {
	co.Lock()
	defer co.Unlock()
	return co.nodeGraph.DebugString(prefix)
}

func (co *ChainObserver) HandleNotification(ctx context.Context, event arbbridge.Event) {
	co.Lock()
	defer co.Unlock()
	switch ev := event.(type) {
	case arbbridge.MessageDeliveredEvent:
		co.messageDelivered(ctx, ev)
	case arbbridge.StakeCreatedEvent:
		co.createStake(ctx, ev)
	case arbbridge.ChallengeStartedEvent:
		co.newChallenge(ctx, ev)
	case arbbridge.ChallengeCompletedEvent:
		co.challengeResolved(ctx, ev)
	case arbbridge.StakeRefundedEvent:
		co.removeStake(ctx, ev)
	case arbbridge.PrunedEvent:
		co.pruneLeaf(ctx, ev)
	case arbbridge.StakeMovedEvent:
		co.moveStake(ctx, ev)
	case arbbridge.AssertedEvent:
		err := co.notifyAssert(ctx, ev)
		if err != nil {
			panic(err)
		}
	case arbbridge.ConfirmedEvent:
		co.confirmNode(ctx, ev)
	}
}

func (co *ChainObserver) NotifyNewBlock(blockID *structures.BlockID) {
	co.Lock()
	defer co.Unlock()
	co.latestBlockID = blockID
	ckptCtx := structures.NewCheckpointContextImpl()
	buf, err := co.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	co.checkpointer.AsyncSaveCheckpoint(blockID.Clone(), buf, ckptCtx, nil)
}

func (co *ChainObserver) CurrentBlockID() *structures.BlockID {
	co.RLock()
	blockID := co.latestBlockID
	co.RUnlock()
	return blockID
}

func (co *ChainObserver) ContractAddress() common.Address {
	co.RLock()
	address := co.rollupAddr
	co.RUnlock()
	return address
}

func (co *ChainObserver) LatestKnownValidMachine() machine.Machine {
	co.RLock()
	mach := co.calculatedValidNode.machine.Clone()
	co.RUnlock()
	return mach
}

func (co *ChainObserver) messageDelivered(ctx context.Context, ev arbbridge.MessageDeliveredEvent) {
	co.pendingInbox.DeliverMessage(ev.Message)
	for _, lis := range co.listeners {
		lis.MessageDelivered(ctx, co, ev)
	}
}

func (co *ChainObserver) pruneLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
	leaf, found := co.nodeGraph.nodeFromHash[ev.Leaf]
	if !found {
		panic("Tried to prune nonexistant leaf")
	}
	co.nodeGraph.leaves.Delete(leaf)
	co.nodeGraph.PruneNodeByHash(ev.Leaf)
	for _, lis := range co.listeners {
		lis.PrunedLeaf(ctx, co, ev)
	}
}

func (co *ChainObserver) createStake(ctx context.Context, ev arbbridge.StakeCreatedEvent) {
	co.nodeGraph.CreateStake(ev)
	for _, lis := range co.listeners {
		lis.StakeCreated(ctx, co, ev)
	}
}

func (co *ChainObserver) removeStake(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	co.nodeGraph.RemoveStake(ev.Staker)
	for _, lis := range co.listeners {
		lis.StakeRemoved(ctx, co, ev)
	}
}

func (co *ChainObserver) moveStake(ctx context.Context, ev arbbridge.StakeMovedEvent) {
	co.nodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range co.listeners {
		lis.StakeMoved(ctx, co, ev)
	}
}

func (co *ChainObserver) newChallenge(ctx context.Context, ev arbbridge.ChallengeStartedEvent) {
	asserter := co.nodeGraph.stakers.Get(ev.Asserter)
	challenger := co.nodeGraph.stakers.Get(ev.Challenger)
	asserterAncestor, challengerAncestor, _, err := GetConflictAncestor(asserter.location, challenger.location)
	if err != nil {
		panic("No conflict ancestor for conflict")
	}

	co.nodeGraph.NewChallenge(
		ev.ChallengeContract,
		ev.Asserter,
		ev.Challenger,
		ev.ChallengeType,
	)
	for _, lis := range co.listeners {
		lis.StartedChallenge(ctx, co, ev, challengerAncestor, asserterAncestor)
	}
}

func (co *ChainObserver) challengeResolved(ctx context.Context, ev arbbridge.ChallengeCompletedEvent) {
	co.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range co.listeners {
		lis.CompletedChallenge(ctx, co, ev)
	}
}

func (co *ChainObserver) confirmNode(ctx context.Context, ev arbbridge.ConfirmedEvent) {
	newNode := co.nodeGraph.nodeFromHash[ev.NodeHash]
	if newNode.depth > co.knownValidNode.depth {
		co.knownValidNode = newNode
	}
	co.nodeGraph.latestConfirmed = newNode
	co.nodeGraph.considerPruningNode(newNode.prev)
	co.updateOldest()
	for _, listener := range co.listeners {
		listener.ConfirmedNode(ctx, co, ev)
	}
}

func (co *ChainObserver) updateOldest() {
	for co.nodeGraph.oldestNode != co.nodeGraph.latestConfirmed {
		if co.nodeGraph.oldestNode.numStakers > 0 {
			return
		}
		if co.calculatedValidNode == co.nodeGraph.oldestNode {
			return
		}
		var successor *Node
		for _, successorHash := range co.nodeGraph.oldestNode.successorHashes {
			if successorHash != zeroBytes32 {
				if successor != nil {
					return
				}
				successor = co.nodeGraph.nodeFromHash[successorHash]
			}
		}
		co.nodeGraph.pruneOldestNode(co.nodeGraph.oldestNode)
		co.nodeGraph.oldestNode = successor
	}
}

func (co *ChainObserver) notifyAssert(ctx context.Context, ev arbbridge.AssertedEvent) error {
	disputableNode := structures.NewDisputableNode(
		ev.Params,
		ev.Claim,
		ev.MaxPendingTop,
		ev.MaxPendingCount,
	)
	if err := co.nodeGraph.createNodesOnAssert(
		co.nodeGraph.nodeFromHash[ev.PrevLeafHash],
		disputableNode,
		ev.BlockID.Height,
		ev.TxHash,
	); err != nil {
		return err
	}
	for _, listener := range co.listeners {
		listener.SawAssertion(ctx, co, ev)
	}
	return nil
}

func (co *ChainObserver) equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Equal(co.rollupAddr[:], co2.rollupAddr[:]) &&
		co.pendingInbox.Equals(co2.pendingInbox)
}

func (co *ChainObserver) executionPrecondition(node *Node) *valprotocol.Precondition {
	vmProtoData := node.prev.vmProtoData
	inbox, _ := co.pendingInbox.GenerateInbox(node.prev.vmProtoData.PendingTop, node.disputable.AssertionParams.ImportedMessageCount.Uint64())
	return &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: inbox.AsValue(),
	}
}

func (co *ChainObserver) currentTimeBounds() *protocol.TimeBoundsBlocks {
	latestTime := co.latestBlockID.Height
	return &protocol.TimeBoundsBlocks{
		Start: latestTime,
		End:   common.NewTimeBlocks(new(big.Int).Add(latestTime.AsInt(), big.NewInt(int64(co.nodeGraph.params.MaxTimeBoundsWidth)))),
	}
}
