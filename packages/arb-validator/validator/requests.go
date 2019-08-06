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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type callRequest struct {
	Message    protocol.Message
	ResultChan chan<- value.Value
	ErrorChan  chan<- error
}

type pendingMessageCheck struct {
	ResultChan chan<- bool
}

type openAssertionCheck struct {
	ResultChan chan<- bool
}

type vmStateRequest struct {
	ResultChan chan<- valmessage.VMStateData
}

type disputableDefenderRequest struct {
	Length     uint64
	ResultChan chan<- bool
}

type initiateUnanimousRequest struct {
	TimeLength    uint64
	NewMessages   []protocol.Message
	MessageHashes [][]byte
	Final         bool
	MaxSteps      int32
	RequestChan   chan<- valmessage.UnanimousRequest
	ResultChan    chan<- valmessage.UnanimousUpdateResults
	ErrChan       chan<- error
}

type followUnanimousRequest struct {
	valmessage.UnanimousRequestData
	NewMessages []protocol.Message
	MaxSteps    int32

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

type closeUnanimousAssertionRequest struct {
	ResultChan chan<- bool
	ErrChan    chan<- error
}

type unanimousConfirmRequest struct {
	valmessage.UnanimousRequestData
	Signatures [][]byte

	ResultChan chan<- bool
	ErrChan    chan<- error
}
