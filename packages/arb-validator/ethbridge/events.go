/*
 * Copyright 2019, Offchain Labs, Inc.
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

package ethbridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type IncomingMessageType int

const (
	CommonMessage IncomingMessageType = iota
	ChallengeMessage
)

type Event interface {
}

type VMEvent interface {
	GetIncomingMessageType() IncomingMessageType
}

type Notification struct {
	Header *types.Header
	VMID   [32]byte
	Event  Event
	TxHash [32]byte
}

type FinalizedUnanimousAssertEvent struct {
	UnanHash [32]byte
}

func (FinalizedUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type PendingUnanimousAssertEvent struct {
	UnanHash    [32]byte
	SequenceNum uint64
}

func (PendingUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedUnanimousAssertEvent struct {
	SequenceNum uint64
}

func (ConfirmedUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type PendingDisputableAssertionEvent struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
	Asserter     common.Address
}

func (PendingDisputableAssertionEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedDisputableAssertEvent struct {
	TxHash   [32]byte
	LogsHash [32]byte
}

func (ConfirmedDisputableAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type InitiateChallengeEvent struct {
	Challenger common.Address
}

func (InitiateChallengeEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type BisectionEvent struct {
	Assertions []*protocol.AssertionStub
}

func (BisectionEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type ContinueChallengeEvent struct {
	ChallengedAssertion uint16
}

func (ContinueChallengeEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type ChallengerTimeoutEvent struct{}

func (ChallengerTimeoutEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type AsserterTimeoutEvent struct{}

func (AsserterTimeoutEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type OneStepProofEvent struct{}

func (OneStepProofEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type VMCreatedEvent struct {
	GracePeriod         uint32
	EscrowRequired      *big.Int
	EscrowCurrency      common.Address
	MaxExecutionSteps   uint32
	VMState             [32]byte
	ChallengeManagerNum uint16
	Owner               common.Address
	Validators          []common.Address
}

func (VMCreatedEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type MessageDeliveredEvent struct {
	Msg protocol.Message
}

type NewTimeEvent struct{}
