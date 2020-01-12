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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/golang/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
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
	latestBlockNumber   *common.TimeBlocks
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
	startTime *common.TimeBlocks,
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
		latestBlockNumber:   startTime,
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
	machineHash [32]byte,
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
	_client arbbridge.ArbRollup,
) *ChainObserver {
	nodeGraph := m.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx)
	chain := &ChainObserver{
		RWMutex:             &sync.RWMutex{},
		nodeGraph:           nodeGraph,
		rollupAddr:          m.ContractAddress.Unmarshal(),
		pendingInbox:        &structures.PendingInbox{m.PendingInbox.UnmarshalFromCheckpoint(restoreCtx)},
		knownValidNode:      nodeGraph.nodeFromHash[m.KnownValidNode.Unmarshal()],
		calculatedValidNode: nodeGraph.nodeFromHash[m.CalculatedValidNode.Unmarshal()],
		latestBlockNumber:   nil,
		listeners:           []ChainListener{},
		checkpointer:        nil,
		isOpinionated:       false,
		assertionMadeChan:   nil,
	}
	chain.Lock()
	defer chain.Unlock()
	if _client != nil {
		chain.startCleanupThread(ctx)
	}
	if m.IsOpinionated {
		chain.isOpinionated = true
		chain.assertionMadeChan = make(chan bool)
		chain.startOpinionUpdateThread(ctx)
	}
	return chain
}

func UnmarshalChainObserverFromBytes(ctx context.Context, buf []byte, restoreCtx structures.RestoreContext, client arbbridge.ArbRollup) (*ChainObserver, error) {
	cob := &ChainObserverBuf{}
	if err := proto.Unmarshal(buf, cob); err != nil {
		return nil, err
	}
	return cob.UnmarshalFromCheckpoint(ctx, restoreCtx, client), nil
}

func (chain *ChainObserver) messageDelivered(ev arbbridge.MessageDeliveredEvent) {
	chain.pendingInbox.DeliverMessage(ev.Msg.AsValue())
	for _, lis := range chain.listeners {
		lis.MessageDelivered(ev)
	}
}

func (chain *ChainObserver) pruneLeaf(ev arbbridge.PrunedEvent) {
	leaf, found := chain.nodeGraph.nodeFromHash[ev.Leaf]
	if !found {
		panic("Tried to prune nonexistant leaf")
	}
	chain.nodeGraph.leaves.Delete(leaf)
	chain.nodeGraph.PruneNodeByHash(ev.Leaf)
	for _, lis := range chain.listeners {
		lis.PrunedLeaf(ev)
	}
}

func (chain *ChainObserver) createStake(ev arbbridge.StakeCreatedEvent, currentTime common.TimeTicks) {
	chain.nodeGraph.CreateStake(ev, currentTime)
	for _, lis := range chain.listeners {
		lis.StakeCreated(ev)
	}
}

func (chain *ChainObserver) removeStake(ev arbbridge.StakeRefundedEvent) {
	chain.nodeGraph.RemoveStake(ev.Staker)
	for _, lis := range chain.listeners {
		lis.StakeRemoved(ev)
	}
}

func (chain *ChainObserver) moveStake(ev arbbridge.StakeMovedEvent) {
	chain.nodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range chain.listeners {
		lis.StakeMoved(ev)
	}
}

func (chain *ChainObserver) newChallenge(ev arbbridge.ChallengeStartedEvent) {
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
		lis.StartedChallenge(ev, asserterAncestor, challengerAncestor)
	}
}

func (chain *ChainObserver) challengeResolved(ev arbbridge.ChallengeCompletedEvent) {
	chain.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ev)
	}
}

func (chain *ChainObserver) confirmNode(ev arbbridge.ConfirmedEvent) {
	newNode := chain.nodeGraph.nodeFromHash[ev.NodeHash]
	if newNode.depth > chain.knownValidNode.depth {
		chain.knownValidNode = newNode
	}
	chain.nodeGraph.latestConfirmed = newNode
	chain.nodeGraph.considerPruningNode(newNode.prev)
	chain.updateOldest(newNode)
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ev)
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

func (chain *ChainObserver) notifyAssert(
	ev arbbridge.AssertedEvent,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) error {
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
		currentTime,
		assertionTxHash,
	)
	for _, listener := range chain.listeners {
		listener.SawAssertion(ev, currentTime, assertionTxHash)
	}
	if chain.assertionMadeChan != nil {
		chain.assertionMadeChan <- true
	}
	return nil
}

func (chain *ChainObserver) notifyNewBlock(blockNum *common.TimeBlocks, blockHash [32]byte) {
	chain.Lock()
	defer chain.Unlock()
	chain.latestBlockNumber = blockNum
	ckptCtx := structures.NewCheckpointContextImpl()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockNum, blockHash, buf, ckptCtx, nil)
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
	return &protocol.TimeBoundsBlocks{
		chain.latestBlockNumber,
		common.NewTimeBlocks(new(big.Int).Add(chain.latestBlockNumber.AsInt(), big.NewInt(10))),
	}
}

func (chain *ChainObserver) CurrentTime() *common.TimeBlocks {
	chain.RLock()
	time := chain.latestBlockNumber
	chain.RUnlock()
	return time
}

func (chain *ChainObserver) ContractAddress() common.Address {
	chain.RLock()
	address := chain.rollupAddr
	chain.RUnlock()
	return address
}

func (chain *ChainObserver) ExecuteCall(messages value.TupleValue, maxSteps uint32) (*protocol.ExecutionAssertion, uint32) {
	chain.RLock()
	mach := chain.knownValidNode.machine.Clone()
	chain.RUnlock()
	assertion, steps := mach.ExecuteAssertion(maxSteps, chain.currentTimeBounds(), messages)
	return assertion, steps
}
