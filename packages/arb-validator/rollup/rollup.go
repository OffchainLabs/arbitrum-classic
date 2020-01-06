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
	"errors"
	"math/big"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*sync.RWMutex
	nodeGraph    *StakedNodeGraph
	rollupAddr   common.Address
	pendingInbox *structures.PendingInbox
	listeners    []ChainListener
}

func NewChain(
	_rollupAddr common.Address,
	_machine machine.Machine,
	_vmParams structures.ChainParams,
) *ChainObserver {
	ret := &ChainObserver{
		nodeGraph:    NewStakedNodeGraph(_machine, _vmParams),
		rollupAddr:   _rollupAddr,
		pendingInbox: structures.NewPendingInbox(),
		listeners:    []ChainListener{},
	}
	ret.startCleanupThread(nil)
	return ret
}

func (chain *ChainObserver) AddListener(listener ChainListener) {
	chain.listeners = append(chain.listeners, listener)
}

func (chain *ChainObserver) MarshalToBuf() *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph: chain.nodeGraph.MarshalToBuf(),
		ContractAddress: chain.rollupAddr.Bytes(),
		PendingInbox:    chain.pendingInbox.MarshalToBuf(),
	}
}

func (m *ChainObserverBuf) Unmarshal(_listenForAddress common.Address) *ChainObserver {
	chain := &ChainObserver{
		nodeGraph:    m.StakedNodeGraph.Unmarshal(),
		rollupAddr:   common.BytesToAddress(m.ContractAddress),
		pendingInbox: &structures.PendingInbox{m.PendingInbox.Unmarshal()},
		listeners:    []ChainListener{},
	}
	chain.startCleanupThread(nil)
	return chain
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
	conflictNode, disputeType, err := GetConflictAncestor(asserter.location, challenger.location)
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
		lis.StartedChallenge(ev, conflictNode, disputeType)
	}
}

func (chain *ChainObserver) ChallengeResolved(ev ethbridge.ChallengeCompletedEvent) {
	chain.nodeGraph.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
	for _, lis := range chain.listeners {
		lis.CompletedChallenge(ev)
	}
}

func (chain *ChainObserver) ConfirmNode(ev ethbridge.ConfirmedEvent) {
	chain.nodeGraph.ConfirmNode(ev.NodeHash)
}

func (chain *ChainObserver) notifyAssert(
	ev ethbridge.AssertedEvent,
	currentTime *protocol.TimeBlocks,
) error {
	topPendingCount, ok := chain.pendingInbox.GetHeight(ev.MaxPendingTop)
	if !ok {
		return errors.New("Couldn't find top message in inbox")
	}
	disputableNode := structures.NewDisputableNode(
		ev.Params,
		ev.Claim,
		ev.MaxPendingTop,
		topPendingCount,
	)
	chain.nodeGraph.CreateNodesOnAssert(chain.nodeGraph.nodeFromHash[ev.PrevLeafHash], disputableNode, nil, currentTime)
	return nil
}

func (chain *ChainObserver) notifyNewBlockNumber(blockNum *big.Int) {
	chain.Lock()
	defer chain.Unlock()
	//TODO: checkpoint, and take other appropriate actions for new block
}

func (co *ChainObserver) Equals(co2 *ChainObserver) bool {
	return co.nodeGraph.Equals(co2.nodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.pendingInbox.Equals(co2.pendingInbox)
}
