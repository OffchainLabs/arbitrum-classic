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

package validator

import (
	"github.com/offchainlabs/arb-validator/valmessage"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
)



type unanimousUpdateRequest struct {
	valmessage.UnanimousRequestData

	NewMessages []protocol.Message

	Inbox     *protocol.Inbox
	Machine   *vm.Machine
	Assertion *protocol.Assertion

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

type followUnanimousRequest struct {
	valmessage.UnanimousRequestData
	NewMessages []protocol.Message

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

type initiateUnanimousRequest struct {
	TimeLength  uint64
	NewMessages []protocol.Message
	Final       bool
	RequestChan chan<- valmessage.UnanimousRequest
	ResultChan  chan<- valmessage.UnanimousUpdateResults
	ErrChan     chan<- error
}

type unanimousConfirmRequest struct {
	valmessage.UnanimousRequestData
	Signatures [][]byte

	ResultChan chan<- bool
	ErrChan    chan<- error
}

type closeUnanimousAssertionRequest struct {
	ResultChan chan<- bool
	ErrChan    chan<- error
}

type callRequest struct {
	Message    protocol.Message
	ResultChan chan<- value.Value
	ErrorChan  chan<- error
}

type vmStateData struct {
	MachineState [32]byte
	Config       valmessage.VMConfiguration
}

type pendingMessageCheck struct {
	ResultChan chan<- bool
}

type vmStateRequest struct {
	ResultChan chan<- vmStateData
}

type disputableDefenderRequest struct {
	Length                 uint64
	IncludePendingMessages bool
	ResultChan             chan<- bool
}

type disputableAssertionRequest struct {
	State           *vm.Machine
	Defender        protocol.AssertionDefender
	IncludedPending bool
	ResultChan      chan<- bool
}

func (r disputableAssertionRequest) GetPrecondition() *protocol.Precondition {
	return r.Defender.GetPrecondition()
}

func (r disputableAssertionRequest) IncludedPendingInbox() bool {
	return r.IncludedPending
}

func (r disputableAssertionRequest) NotifyInvalid() {
	go func() {
		r.ResultChan <- false
	}()
}

func (r disputableAssertionRequest) NotifyAccepted() {
	go func() {
		r.ResultChan <- true
	}()
}
