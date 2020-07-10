/*
 * Copyright 2020, Offchain Labs, Inc.
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

package arbbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type Event interface {
	GetChainInfo() ChainInfo
}

// MergeEventsUnsafe assumes that both sets of events are disjoint and come from
// the same chain state rather than from two different states caused by a reorg
func MergeEventsUnsafe(events1 []Event, events2 []Event) []Event {
	totalLen := len(events1) + len(events2)
	events := make([]Event, 0, totalLen)
	events1Index := 0
	events2Index := 0
	for i := 0; i < totalLen; i++ {
		if events1Index == len(events1) {
			events = append(events, events2[events2Index])
			events2Index++
		} else if events2Index == len(events2) {
			events = append(events, events1[events1Index])
			events1Index++
		} else {
			event1 := events1[events1Index]
			event2 := events2[events2Index]
			if event1.GetChainInfo().BlockId.Height.AsInt().Cmp(event2.GetChainInfo().BlockId.Height.AsInt()) < 0 {
				events = append(events, events1[events1Index])
				events1Index++
			} else {
				events = append(events, events2[events2Index])
				events2Index++
			}
		}
	}
	return events
}

type ChainInfo struct {
	BlockId  *common.BlockId
	LogIndex uint
	TxHash   [32]byte
}

func (c ChainInfo) GetChainInfo() ChainInfo {
	return c
}

type StakeCreatedEvent struct {
	ChainInfo
	Staker   common.Address
	NodeHash common.Hash
}

type ChallengeStartedEvent struct {
	ChainInfo
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     valprotocol.ChildType
	ChallengeContract common.Address
}

type ChallengeCompletedEvent struct {
	ChainInfo
	Winner            common.Address
	Loser             common.Address
	ChallengeContract common.Address
}

type StakeRefundedEvent struct {
	ChainInfo
	Staker common.Address
}

type PrunedEvent struct {
	ChainInfo
	Leaf common.Hash
}

type StakeMovedEvent struct {
	ChainInfo
	Staker   common.Address
	Location common.Hash
}

type AssertedEvent struct {
	ChainInfo
	PrevLeafHash common.Hash
	Disputable   *valprotocol.DisputableNode
}

type ConfirmedEvent struct {
	ChainInfo
	NodeHash common.Hash
}

type ConfirmedAssertionEvent struct {
	ChainInfo
	LogsAccHash []common.Hash
}

type InitiateChallengeEvent struct {
	ChainInfo
	Deadline common.TimeTicks
}

type AsserterTimeoutEvent struct {
	ChainInfo
}

type ChallengerTimeoutEvent struct {
	ChainInfo
}

type ContinueChallengeEvent struct {
	ChainInfo
	SegmentIndex *big.Int
	Deadline     common.TimeTicks
}

type OneStepProofEvent struct {
	ChainInfo
}

type InboxTopBisectionEvent struct {
	ChainInfo
	ChainHashes []common.Hash
	TotalLength *big.Int
	Deadline    common.TimeTicks
}

type MessagesBisectionEvent struct {
	ChainInfo
	ChainHashes   []common.Hash
	SegmentHashes []common.Hash
	TotalLength   *big.Int
	Deadline      common.TimeTicks
}

type ExecutionBisectionEvent struct {
	ChainInfo
	Assertions []*valprotocol.ExecutionAssertionStub
	TotalSteps uint64
	Deadline   common.TimeTicks
}

type MessageDeliveredEvent struct {
	ChainInfo
	Message message.InboxMessage
}

type NewTimeEvent struct {
	ChainInfo
}
