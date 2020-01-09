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

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/golang/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph         *StakedNodeGraph
	rollupAddr        common.Address
	pendingInbox      *structures.PendingInbox
	knownValidNode    *Node
	latestBlockNumber *big.Int
	listeners         []ChainListener
	checkpointer      *structures.RollupCheckpointer
	isOpinionated     bool
	assertionMadeChan chan bool
}

func NewChain(
	ctx context.Context,
	rollupAddr common.Address,
	checkpointer *structures.RollupCheckpointer,
	vmParams structures.ChainParams,
	updateOpinion bool,
	startTime *big.Int,
) (*ChainObserver, error) {
	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	ret := &ChainObserver{
		RWMutex:           &sync.RWMutex{},
		nodeGraph:         NewStakedNodeGraph(mach, vmParams),
		rollupAddr:        rollupAddr,
		pendingInbox:      structures.NewPendingInbox(),
		latestBlockNumber: startTime,
		listeners:         []ChainListener{},
		checkpointer:      checkpointer,
	}
	ret.knownValidNode = ret.nodeGraph.latestConfirmed
	ret.Lock()
	defer ret.Unlock()

	ret.startCleanupThread(ctx)
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

func (chain *ChainObserver) MarshalForCheckpoint(ctx structures.CheckpointContext) *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph: chain.nodeGraph.MarshalForCheckpoint(ctx),
		ContractAddress: chain.rollupAddr.Bytes(),
		PendingInbox:    chain.pendingInbox.MarshalForCheckpoint(ctx),
		IsOpinionated:   chain.isOpinionated,
	}
}

func (chain *ChainObserver) MarshalToBytes(ctx structures.CheckpointContext) ([]byte, error) {
	cob := chain.MarshalForCheckpoint(ctx)
	return proto.Marshal(cob)
}

func (m *ChainObserverBuf) UnmarshalFromCheckpoint(
	ctx context.Context,
	restoreCtx structures.RestoreContext,
	_client *ethbridge.ArbRollup,
) *ChainObserver {
	chain := &ChainObserver{
		RWMutex:      &sync.RWMutex{},
		nodeGraph:    m.StakedNodeGraph.UnmarshalFromCheckpoint(restoreCtx),
		rollupAddr:   common.BytesToAddress(m.ContractAddress),
		pendingInbox: &structures.PendingInbox{m.PendingInbox.UnmarshalFromCheckpoint(restoreCtx)},
		listeners:    []ChainListener{},
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

func UnmarshalChainObserverFromBytes(ctx context.Context, buf []byte, restoreCtx structures.RestoreContext, client *ethbridge.ArbRollup) (*ChainObserver, error) {
	cob := &ChainObserverBuf{}
	if err := proto.Unmarshal(buf, cob); err != nil {
		return nil, err
	}
	return cob.UnmarshalFromCheckpoint(ctx, restoreCtx, client), nil
}

func (chain *ChainObserver) PruneNode(ev ethbridge.PrunedEvent) {
	chain.nodeGraph.PruneNodeByHash(ev.Leaf)
}

func (chain *ChainObserver) CreateStake(ev ethbridge.StakeCreatedEvent, currentTime structures.TimeTicks) {
	chain.nodeGraph.CreateStake(ev, currentTime)
	for _, lis := range chain.listeners {
		lis.StakeCreated(ev)
	}
}

func (chain *ChainObserver) RemoveStake(ev ethbridge.StakeRefundedEvent) {
	chain.nodeGraph.RemoveStake(ev.Staker)
	for _, lis := range chain.listeners {
		lis.StakeRemoved(ev)
	}
}

func (chain *ChainObserver) MoveStake(ev ethbridge.StakeMovedEvent) {
	chain.nodeGraph.MoveStake(ev.Staker, ev.Location)
	for _, lis := range chain.listeners {
		lis.StakeMoved(ev)
	}
}

func (chain *ChainObserver) NewChallenge(ev ethbridge.ChallengeStartedEvent) {
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

func (chain *ChainObserver) ChallengeResolved(ev ethbridge.ChallengeCompletedEvent) {
	chain.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ev)
	}
}

func (chain *ChainObserver) ConfirmNode(ev ethbridge.ConfirmedEvent) {
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
	ev ethbridge.AssertedEvent,
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

func (chain *ChainObserver) notifyNewBlockNumber(blockNum *big.Int) {
	chain.Lock()
	defer chain.Unlock()
	chain.latestBlockNumber = blockNum
	ckptCtx := structures.NewCheckpointContextImpl()
	buf, err := chain.MarshalToBytes(ckptCtx)
	if err != nil {
		log.Fatal(err)
	}
	chain.checkpointer.AsyncSaveCheckpoint(blockNum, buf, ckptCtx, nil)
}

func (co *ChainObserver) Equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.pendingInbox.Equals(co2.pendingInbox)
}

func (chain *ChainObserver) ExecutionPrecondition(node *Node) *protocol.Precondition {
	vmProtoData := node.prev.vmProtoData
	inbox := protocol.NewInbox()
	messages := chain.pendingInbox.ValueForSubseq(node.prev.vmProtoData.PendingTop, node.disputable.AssertionClaim.AfterPendingTop)
	inbox.WithAddedMessages(messages)
	return &protocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: inbox.Receive(),
	}
}
