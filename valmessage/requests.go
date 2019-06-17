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
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
)

type UnanimousRequest struct {
	UnanimousRequestData
	NewMessages []protocol.Message
}

type UnanimousUpdateRequest struct {
	UnanimousRequestData

	NewMessages []protocol.Message

	Inbox     *protocol.Inbox
	Machine   *vm.Machine
	Assertion *protocol.Assertion

	ResultChan chan<- UnanimousUpdateResults
	ErrChan    chan<- error
}

type UnanimousRequestData struct {
	BeforeHash  [32]byte
	BeforeInbox [32]byte
	SequenceNum uint64
	TimeBounds  protocol.TimeBounds
}

func (r UnanimousRequestData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(r.BeforeHash),
		solsha3.Bytes32(r.BeforeInbox),
		solsha3.Uint64(r.SequenceNum),
		solsha3.Uint64(r.TimeBounds[0]),
		solsha3.Uint64(r.TimeBounds[1]),
	))
	return ret
}

type FollowUnanimousRequest struct {
	UnanimousRequestData
	NewMessages []protocol.Message

	ResultChan chan<- UnanimousUpdateResults
	ErrChan    chan<- error
}

type InitiateUnanimousRequest struct {
	TimeLength  uint64
	NewMessages []protocol.Message
	Final       bool
	RequestChan chan<- UnanimousRequest
	ResultChan  chan<- UnanimousUpdateResults
	ErrChan     chan<- error
}

type UnanimousConfirmRequest struct {
	UnanimousRequestData
	Signatures []Signature

	ResultChan chan<- bool
	ErrChan    chan<- error
}

type CloseUnanimousAssertionRequest struct {
	ResultChan chan<- bool
	ErrChan    chan<- error
}

type UnanimousUpdateResults struct {
	SequenceNum       uint64
	BeforeHash        [32]byte
	TimeBounds        protocol.TimeBounds
	NewInboxHash      [32]byte
	OriginalInboxHash [32]byte
	Assertion         *protocol.Assertion
}

type CallRequest struct {
	Message    protocol.Message
	ResultChan chan<- value.Value
	ErrorChan  chan<- error
}

type VMStateData struct {
	MachineState [32]byte
	Config       VMConfiguration
}

type PendingMessageCheck struct {
	ResultChan chan<- bool
}

type VMStateRequest struct {
	ResultChan chan<- VMStateData
}

type DisputableDefenderRequest struct {
	Length                 uint64
	IncludePendingMessages bool
	ResultChan             chan<- bool
}

type DisputableAssertionRequest struct {
	State           *vm.Machine
	Defender        protocol.AssertionDefender
	IncludedPending bool
	ResultChan      chan<- bool
}

func (r DisputableAssertionRequest) GetPrecondition() *protocol.Precondition {
	return r.Defender.GetPrecondition()
}

func (r DisputableAssertionRequest) IncludedPendingInbox() bool {
	return r.IncludedPending
}

func (r DisputableAssertionRequest) NotifyInvalid() {
	go func() {
		r.ResultChan <- false
	}()
}

func (r DisputableAssertionRequest) NotifyAccepted() {
	go func() {
		r.ResultChan <- true
	}()
}
