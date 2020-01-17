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
	latestBlockId       *structures.BlockId
	listeners           []ChainListener
	checkpointer        RollupCheckpointer
	isOpinionated       bool
	assertionMadeChan   chan bool
}

func NewChain(
	rollupAddr common.Address,
	checkpointer RollupCheckpointer,
	vmParams structures.ChainParams,
	updateOpinion bool,
	startBlockId *structures.BlockId,
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
		latestBlockId:       startBlockId,
		listeners:           []ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       false,
		assertionMadeChan:   nil,
	}
	ret.Lock()
	defer ret.Unlock()
	if updateOpinion {
		ret.isOpinionated = true
		ret.assertionMadeChan = make(chan bool, 20)
	}
	return ret, nil
}

func (chain *ChainObserver) Start(ctx context.Context) {
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

func MakeInitialChainObserverBuf(
	contractAddress common.Address,
	machineHash common.Hash,
	params *structures.ChainParams,
	opinionated bool,
) *ChainObserverBuf {
	initStakedNodeGraphBuf, initNodeHashBuf := MakeInitialStakedNodeGraphBuf(machineHash, params)
	return &ChainObserverBuf{
		StakedNodeGraph:     initStakedNodeGraphBuf,
		ContractAddress:     contractAddress.MarshallToBuf(),
		PendingInbox:        structures.MakeInitialPendingInboxBuf(),
		KnownValidNode:      initNodeHashBuf,
		CalculatedValidNode: initNodeHashBuf,
		IsOpinionated:       opinionated,
	}
}

func (chain *ChainObserver) marshalForCheckpoint(ctx structures.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph:     chain.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress:     chain.rollupAddr.MarshallToBuf(),
		PendingInbox:        chain.pendingInbox.MarshalForCheckpoint(ctx),
		KnownValidNode:      chain.knownValidNode.hash.MarshalToBuf(),
		CalculatedValidNode: chain.calculatedValidNode.hash.MarshalToBuf(),
		IsOpinionated:       chain.isOpinionated,
	}
}

func (chain *ChainObserver) marshalToBytes(ctx structures.CheckpointContext) ([]byte, error) {
	cob := chain.marshalForCheckpoint(ctx)
	return proto.Marshal(cob)
}

func (m *ChainObserverBuf) UnmarshalFromCheckpoint(
	ctx context.Context,
	restoreCtx structures.RestoreContext,
	latestBlockId *structures.BlockId,
	client arbbridge.ArbRollupWatcher,
	checkpointer RollupCheckpointer,
) *ChainObserver {
	nodeGraph := m.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx)
	chain := &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          m.ContractAddress.Unmarshal(),
		pendingInbox:        &structures.PendingInbox{m.PendingInbox.UnmarshalFromCheckpoint(restoreCtx)},
		knownValidNode:      nodeGraph.nodeFromHash[m.KnownValidNode.Unmarshal()],
		calculatedValidNode: nodeGraph.nodeFromHash[m.CalculatedValidNode.Unmarshal()],
		latestBlockId:       latestBlockId,
		listeners:           []ChainListener{},
		checkpointer:        checkpointer,
		isOpinionated:       m.IsOpinionated,
		assertionMadeChan:   make(chan bool, 10),
	}
	chain.Lock()
	defer chain.Unlock()
	if client != nil {
		log.Println("Starting confirm thread")
		chain.startConfirmThread(ctx)
		log.Println("Starting cleanup thread")
		chain.startCleanupThread(ctx)
	}
	if m.IsOpinionated {
		log.Println("Starting opinion thread")
		chain.startOpinionUpdateThread(ctx)
	}
	return chain
}

func UnmarshalChainObserverFromBytes(ctx context.Context, buf []byte, restoreCtx structures.RestoreContext, latestBlockId *structures.BlockId, client arbbridge.ArbRollupWatcher, checkpointer RollupCheckpointer) (*ChainObserver, error) {
	cob := &ChainObserverBuf{}
	if err := proto.Unmarshal(buf, cob); err != nil {
		return nil, err
	}
	return cob.UnmarshalFromCheckpoint(ctx, restoreCtx, latestBlockId, client, checkpointer), nil
}

func (chain *ChainObserver) MessageDelivered(ctx context.Context, ev arbbridge.MessageDeliveredEvent) {
	chain.pendingInbox.DeliverMessage(ev.Msg.AsValue())
	for _, lis := range chain.listeners {
		lis.MessageDelivered(ctx, chain, ev)
	}
}

func (chain *ChainObserver) PruneLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
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

func (chain *ChainObserver) CreateStake(ctx context.Context, ev arbbridge.StakeCreatedEvent) {
	chain.nodeGraph.CreateStake(ev)
	for _, lis := range chain.listeners {
		lis.StakeCreated(ctx, chain, ev)
	}
}

func (chain *ChainObserver) RemoveStake(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	chain.nodeGraph.RemoveStake(ev.Staker)
	for _, lis := range chain.listeners {
		lis.StakeRemoved(ctx, chain, ev)
	}
}

func (chain *ChainObserver) MoveStake(ctx context.Context, ev arbbridge.StakeMovedEvent) {
	chain.nodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range chain.listeners {
		lis.StakeMoved(ctx, chain, ev)
	}
}

func (chain *ChainObserver) NewChallenge(ctx context.Context, ev arbbridge.ChallengeStartedEvent) {
	asserter := chain.nodeGraph.stakers.Get(ev.Asserter)
	challenger := chain.nodeGraph.stakers.Get(ev.Challenger)
	asserterAncestor, challengerAncestor, err := GetConflictAncestor(asserter.location, challenger.location)
	if err != nil {
		panic("No conflict ancestor for conflict")
	}

	chain.nodeGraph.NewChallenge(
		ev.ChallengeContract,
		ev.Asserter,
		ev.Challenger,
		ev.ChallengeType,
	)
	for _, lis := range chain.listeners {
		lis.StartedChallenge(ctx, chain, ev, asserterAncestor, challengerAncestor)
	}
}

func (chain *ChainObserver) ChallengeResolved(ctx context.Context, ev arbbridge.ChallengeCompletedEvent) {
	chain.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ctx, chain, ev)
	}
}

func (chain *ChainObserver) ConfirmNode(ctx context.Context, ev arbbridge.ConfirmedEvent) {
	newNode := chain.nodeGraph.nodeFromHash[ev.NodeHash]
	if newNode.depth > chain.knownValidNode.depth {
		chain.knownValidNode = newNode
	}
	chain.nodeGraph.latestConfirmed = newNode
	chain.nodeGraph.considerPruningNode(newNode.prev)
	chain.updateOldest(newNode)
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ctx, chain, ev)
	}
}

func (chain *ChainObserver) updateOldest(node *Node) {
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
		chain.nodeGraph.pruneNode(chain.nodeGraph.oldestNode)
		chain.nodeGraph.oldestNode = successor
	}
}

func (chain *ChainObserver) NotifyAssert(ctx context.Context, ev arbbridge.AssertedEvent) error {
	disputableNode := structures.NewDisputableNode(
		ev.Params,
		ev.Claim,
		ev.MaxPendingTop,
		ev.MaxPendingCount,
	)
	chain.nodeGraph.CreateNodesOnAssert(
		chain.nodeGraph.nodeFromHash[ev.PrevLeafHash],
		disputableNode,
		nil,
		ev.BlockId.Height,
		ev.TxHash,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ctx, chain, ev)
	}
	if chain.assertionMadeChan != nil {
		chain.assertionMadeChan <- true
	}
	return nil
}

func (chain *ChainObserver) NotifyNewBlock(blockId *structures.BlockId) {
	chain.Lock()
	defer chain.Unlock()
	chain.latestBlockId = blockId.Clone()
	ckptCtx := structures.NewCheckpointContextImpl()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockId.Clone(), buf, ckptCtx, nil)
}

func (co *ChainObserver) equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.pendingInbox.Equals(co2.pendingInbox)
}

func (chain *ChainObserver) executionPrecondition(node *Node) *valprotocol.Precondition {
	vmProtoData := node.prev.vmProtoData
	messages := chain.pendingInbox.ValueForSubseq(node.prev.vmProtoData.PendingTop, node.disputable.AssertionClaim.AfterPendingTop)
	return &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: messages,
	}
}

func (chain *ChainObserver) currentTimeBounds() *protocol.TimeBoundsBlocks {
	latestTime := chain.latestBlockId.Height
	return &protocol.TimeBoundsBlocks{
		latestTime,
		common.NewTimeBlocks(new(big.Int).Add(latestTime.AsInt(), big.NewInt(10))),
	}
}

func (chain *ChainObserver) CurrentBlockId() *structures.BlockId {
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
	mach := chain.knownValidNode.machine.Clone()
	chain.RUnlock()
	return mach
}
