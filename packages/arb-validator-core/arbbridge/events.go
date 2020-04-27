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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type Event interface {
	GetChainInfo() ChainInfo
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
	PrevLeafHash  common.Hash
	Params        *valprotocol.AssertionParams
	Claim         *valprotocol.AssertionClaim
	MaxInboxTop   common.Hash
	MaxInboxCount *big.Int
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

type MessageBatchDeliveredEvent struct {
	ChainInfo
	Messages []message.InboxMessage
}

type NewTimeEvent struct {
	ChainInfo
}
