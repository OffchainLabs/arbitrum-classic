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

package valmessage

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arb-avm/protocol"
)

type IncomingMessageType int

const (
	CommonMessage IncomingMessageType = iota
	ChallengeMessage
)

type IncomingMessage interface {
	GetIncomingMessageType() IncomingMessageType
}

type TimeUpdate struct {
	NewTime uint64
}

type ProposedUnanimousAssertMessage struct {
	UnanHash    [32]byte
	SequenceNum uint64
}

func (ProposedUnanimousAssertMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type FinalUnanimousAssertMessage struct {
	UnanHash [32]byte
}

func (FinalUnanimousAssertMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedUnanimousAssertMessage struct {
	SequenceNum uint64
}

func (ConfirmedUnanimousAssertMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type AssertMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
	Asserter     common.Address
}

func (AssertMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedAssertMessage struct {
}

func (ConfirmedAssertMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type InitiateChallengeMessage struct {
	Challenger common.Address
}

func (InitiateChallengeMessage) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ContinueChallengeMessage struct {
	ChallengedAssertion uint16
}

func (ContinueChallengeMessage) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type BisectMessage struct {
	Assertions []*protocol.AssertionStub
}

func (BisectMessage) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type ChallengerTimeoutMessage struct{}

func (ChallengerTimeoutMessage) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type AsserterTimeoutMessage struct{}

func (AsserterTimeoutMessage) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type OneStepProofMessage struct{}

func (OneStepProofMessage) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}
