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

package state

import (
	"github.com/offchainlabs/arb-util/machine"
	"github.com/offchainlabs/arb-util/protocol"

	"github.com/offchainlabs/arb-validator/valmessage"
)

type UnanimousUpdateRequest struct {
	valmessage.UnanimousRequestData

	NewMessages []protocol.Message

	Machine   machine.Machine
	Assertion *protocol.Assertion

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

type DisputableAssertionRequest struct {
	AfterState      machine.Machine
	Precondition    *protocol.Precondition
	Assertion       *protocol.Assertion
	ResultChan      chan<- bool
}

func (r DisputableAssertionRequest) GetPrecondition() *protocol.Precondition {
	return r.Precondition
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
