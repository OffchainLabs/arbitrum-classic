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
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/golang/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph           *StakedNodeGraph
	rollupAddr          common.Address
	inbox               *structures.Inbox
	knownValidNode      *Node
	calculatedValidNode *Node
	latestBlockId       *common.BlockId
	listeners           []ChainListener
	checkpointer        checkpointing.RollupCheckpointer
	isOpinionated       bool
	atHead              bool
}

func NewChain(
	rollupAddr common.Address,
	checkpointer checkpointing.RollupCheckpointer,
	vmParams valprotocol.ChainParams,
	updateOpinion bool,
	startBlockId *common.BlockId,
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

func (chain *ChainObserver) AddListener(listener ChainListener) {
	chain.Lock()
	chain.listeners = append(chain.listeners, listener)
	chain.Unlock()
}

func (chain *ChainObserver) NowAtHead() {
	chain.Lock()
	chain.atHead = true
	chain.Unlock()
}

func (chain *ChainObserver) marshalForCheckpoint(ctx checkpointing.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph:     chain.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress:     chain.rollupAddr.MarshallToBuf(),
		Inbox:               chain.inbox.MarshalForCheckpoint(ctx),
		KnownValidNode:      chain.knownValidNode.hash.MarshalToBuf(),
		CalculatedValidNode: chain.calculatedValidNode.hash.MarshalToBuf(),
		LatestBlockId:       chain.latestBlockId.MarshalToBuf(),
		IsOpinionated:       chain.isOpinionated,
	}
}

func (chain *ChainObserver) marshalToBytes(ctx checkpointing.CheckpointContext) ([]byte, error) {
	cob := chain.marshalForCheckpoint(ctx)
	return proto.Marshal(cob)
}

func (m *ChainObserverBuf) UnmarshalFromCheckpoint(
	ctx context.Context,
	restoreCtx checkpointing.RestoreContext,
	checkpointer checkpointing.RollupCheckpointer,
) (*ChainObserver, error) {
	nodeGraph := m.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx)
	inbox, err := m.Inbox.UnmarshalFromCheckpoint(restoreCtx)
	if err != nil {
		return nil, err
	}
	return &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          m.ContractAddress.Unmarshal(),
		inbox:               &structures.Inbox{inbox},
		knownValidNode:      nodeGraph.nodeFromHash[m.KnownValidNode.Unmarshal()],
		calculatedValidNode: nodeGraph.nodeFromHash[m.CalculatedValidNode.Unmarshal()],
		latestBlockId:       m.LatestBlockId.Unmarshal(),
		listeners:           []ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       m.IsOpinionated,
		atHead:              false,
	}, nil
}

func (chain *ChainObserver) DebugString(prefix string) string {
	chain.Lock()
	defer chain.Unlock()
	return chain.nodeGraph.DebugString(prefix)
}

func (chain *ChainObserver) HandleNotification(ctx context.Context, event arbbridge.Event) {
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
		err := chain.notifyAssert(ctx, ev)
		if err != nil {
			panic(err)
		}
	case arbbridge.ConfirmedEvent:
		chain.confirmNode(ctx, ev)
	}
}

func (chain *ChainObserver) NotifyNewBlock(blockId *common.BlockId) {
	chain.Lock()
	defer chain.Unlock()
	chain.latestBlockId = blockId
	ckptCtx := checkpointing.NewCheckpointContextImpl()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockId.Clone(), buf, ckptCtx, nil)
}

func (chain *ChainObserver) CurrentBlockId() *common.BlockId {
	chain.RLock()
	blockId := chain.latestBlockId
	chain.RUnlock()
	return blockId
}

func (chain *ChainObserver) ContractAddress() common.Address {
	chain.RLock()
	address := chain.rollupAddr
	chain.RUnlock()
	return address
}

func (chain *ChainObserver) LatestKnownValidMachine() machine.Machine {
	chain.RLock()
	mach := chain.calculatedValidNode.machine.Clone()
	chain.RUnlock()
	return mach
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
	_, challengerAncestor, err := GetConflictAncestor(asserter.location, challenger.location)
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

func (chain *ChainObserver) confirmNode(ctx context.Context, ev arbbridge.ConfirmedEvent) {
	newNode := chain.nodeGraph.nodeFromHash[ev.NodeHash]
	if newNode.depth > chain.knownValidNode.depth {
		chain.knownValidNode = newNode
	}
	chain.nodeGraph.latestConfirmed = newNode
	chain.nodeGraph.considerPruningNode(newNode.prev)
	chain.updateOldest()
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ctx, chain, ev)
	}
}

func (chain *ChainObserver) updateOldest() {
	for chain.nodeGraph.oldestNode != chain.nodeGraph.latestConfirmed {
		if chain.nodeGraph.oldestNode.numStakers > 0 {
			return
		}
		if chain.calculatedValidNode == chain.nodeGraph.oldestNode {
			return
		}
		var successor *Node
		for _, successorHash := range chain.nodeGraph.oldestNode.successorHashes {
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

func (chain *ChainObserver) notifyAssert(ctx context.Context, ev arbbridge.AssertedEvent) error {
	disputableNode := valprotocol.NewDisputableNode(
		ev.Params,
		ev.Claim,
		ev.MaxInboxTop,
		ev.MaxInboxCount,
	)
	chain.nodeGraph.CreateNodesOnAssert(
		chain.nodeGraph.nodeFromHash[ev.PrevLeafHash],
		disputableNode,
		ev.BlockId.Height,
		ev.TxHash,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ctx, chain, ev)
	}
	return nil
}

func (co *ChainObserver) equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.inbox.Equals(co2.inbox)
}

func (chain *ChainObserver) executionPrecondition(node *Node) *valprotocol.Precondition {
	vmProtoData := node.prev.vmProtoData
	inbox, _ := chain.inbox.GenerateVMInbox(node.prev.vmProtoData.InboxTop, node.disputable.AssertionParams.ImportedMessageCount.Uint64())
	return &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: inbox.AsValue(),
	}
}

func (chain *ChainObserver) currentTimeBounds() *protocol.TimeBounds {
	latestBlock := chain.latestBlockId.Height
	// Start timestamp slightly in the past to avoid it being invalid
	latestTimestamp := time.Now().Unix() - 60
	return &protocol.TimeBounds{
		StartBlock: latestBlock,
		EndBlock:   common.NewTimeBlocks(new(big.Int).Add(latestBlock.AsInt(), big.NewInt(int64(chain.nodeGraph.params.MaxBlockBoundsWidth)))),
		StartTime:  big.NewInt(latestTimestamp),
		EndTime:    big.NewInt(latestTimestamp + int64(chain.nodeGraph.params.MaxTimestampBoundsWidth)),
	}
}
