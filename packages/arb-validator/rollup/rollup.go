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
	"log"
	"math/big"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/golang/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserverConfig struct {
	MaxCallSteps uint32
}

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph         *StakedNodeGraph
	rollupAddr        common.Address
	config            ChainObserverConfig
	pendingInbox      *structures.PendingInbox
	knownValidNode    *Node
	latestBlockNumber *protocol.TimeBlocks
	listeners         []ChainListener
	checkpointer      *structures.RollupCheckpointer
	isOpinionated     bool
	assertionMadeChan chan bool
}

func NewChain(
	ctx context.Context,
	rollupAddr common.Address,
	config ChainObserverConfig,
	checkpointer *structures.RollupCheckpointer,
	vmParams structures.ChainParams,
	updateOpinion bool,
	startTime *protocol.TimeBlocks,
) (*ChainObserver, error) {
	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	nodeGraph := NewStakedNodeGraph(mach, vmParams)
	ret := &ChainObserver{
		RWMutex:           &sync.RWMutex{},
		nodeGraph:         nodeGraph,
		rollupAddr:        rollupAddr,
		config:            config,
		pendingInbox:      structures.NewPendingInbox(),
		knownValidNode:    nodeGraph.latestConfirmed,
		latestBlockNumber: startTime,
		listeners:         []ChainListener{},
		checkpointer:      checkpointer,
		isOpinionated:     false,
		assertionMadeChan: nil,
	}
	ret.Lock()
	defer ret.Unlock()

	ret.startCleanupThread(ctx)
	ret.startConfirmThread(ctx)
	if updateOpinion {
		ret.isOpinionated = true
		ret.assertionMadeChan = make(chan bool, 20)
		ret.startOpinionUpdateThread(ctx)
	}
	return ret, nil
}

func (chain *ChainObserver) AddListener(listener ChainListener) {
	chain.Lock()
	chain.listeners = append(chain.listeners, listener)
	chain.Unlock()
}

func (chain *ChainObserver) marshalForCheckpoint(ctx structures.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph: chain.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress: chain.rollupAddr.Bytes(),
		PendingInbox:    chain.pendingInbox.MarshalForCheckpoint(ctx),
		KnownValidNode:  utils.MarshalHash(chain.knownValidNode.hash),
		IsOpinionated:   chain.isOpinionated,
		MaxCallSteps:    chain.config.MaxCallSteps,
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
		RWMutex:           &sync.RWMutex{},
		nodeGraph:         nodeGraph,
		rollupAddr:        common.BytesToAddress(m.ContractAddress),
		config:            ChainObserverConfig{MaxCallSteps: m.MaxCallSteps},
		pendingInbox:      &structures.PendingInbox{m.PendingInbox.UnmarshalFromCheckpoint(restoreCtx)},
		knownValidNode:    nodeGraph.nodeFromHash[utils.UnmarshalHash(m.KnownValidNode)],
		latestBlockNumber: nil,
		listeners:         []ChainListener{},
		checkpointer:      nil,
		isOpinionated:     false,
		assertionMadeChan: nil,
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
}

func (chain *ChainObserver) pruneLeaf(ev arbbridge.PrunedEvent) {
	chain.nodeGraph.leaves.Delete(chain.nodeGraph.nodeFromHash[ev.Leaf])
	chain.nodeGraph.PruneNodeByHash(ev.Leaf)
	for _, lis := range chain.listeners {
		lis.PrunedLeaf(ev)
	}
}

func (chain *ChainObserver) createStake(ev arbbridge.StakeCreatedEvent, currentTime structures.TimeTicks) {
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
	chain.nodeGraph.ConfirmNode(ev.NodeHash)
	for _, listener := range chain.listeners {
		listener.ConfirmedNode(ev)
	}
}

func (chain *ChainObserver) notifyAssert(
	ev arbbridge.AssertedEvent,
	currentTime *protocol.TimeBlocks,
	assertionTxHash [32]byte,
) error {
	topPendingCount, ok := chain.pendingInbox.GetHeight(ev.MaxPendingTop)
	if !ok {
		return fmt.Errorf("Couldn't find top message in inbox: %v", hexutil.Encode(ev.MaxPendingTop[:]))
	}
	disputableNode := structures.NewDisputableNode(
		ev.Params,
		ev.Claim,
		ev.MaxPendingTop,
		topPendingCount,
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

func (chain *ChainObserver) notifyNewBlockNumber(blockNum *protocol.TimeBlocks) {
	chain.Lock()
	defer chain.Unlock()
	chain.latestBlockNumber = blockNum
	ckptCtx := structures.NewCheckpointContextImpl()
	buf, err := chain.marshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockNum.AsInt(), buf, ckptCtx, nil)
}

func (co *ChainObserver) equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.pendingInbox.Equals(co2.pendingInbox)
}

func (chain *ChainObserver) executionPrecondition(node *Node) *protocol.Precondition {
	vmProtoData := node.prev.vmProtoData
	messages := chain.pendingInbox.ValueForSubseq(node.prev.vmProtoData.PendingTop, node.disputable.AssertionClaim.AfterPendingTop)
	return &protocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: messages,
	}
}

func (chain *ChainObserver) currentTimeBounds() *protocol.TimeBoundsBlocks {
	return protocol.NewTimeBoundsBlocks(
		chain.latestBlockNumber,
		protocol.NewTimeBlocks(new(big.Int).Add(chain.latestBlockNumber.AsInt(), big.NewInt(10))),
	)
}

func (chain *ChainObserver) CurrentTime() *protocol.TimeBlocks {
	chain.RLock()
	time := chain.latestBlockNumber
	chain.RUnlock()
	return time
}

func (chain *ChainObserver) ExecuteCall(messages value.TupleValue) (*protocol.ExecutionAssertion, uint32) {
	chain.RLock()
	mach := chain.knownValidNode.machine.Clone()
	chain.RUnlock()
	assertion, steps := mach.ExecuteAssertion(chain.config.MaxCallSteps, chain.currentTimeBounds(), messages)
	return assertion, steps
}
