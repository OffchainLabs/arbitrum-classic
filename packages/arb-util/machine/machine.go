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

package machine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Status int

const (
	Extensive Status = iota
	ErrorStop
	Halt
)

type Machine interface {
	Hash() common.Hash
	Clone() Machine
	PrintState()

	CurrentStatus() Status
	IsBlocked(newMessages bool) BlockReason

	ExecuteAssertion(maxSteps uint64, messages []inbox.InboxMessage, maxWallTime time.Duration) (*protocol.ExecutionAssertion, []value.Value, uint64)

	// Supply a value that the inbox peek opcode will return if the inbox
	// runs out of messages. For ArbOS, this can be used to simulate a message
	// from the next block arriving in order to trigger end-of-block processes
	// without waiting for the next block
	ExecuteCallServerAssertion(maxSteps uint64, inboxMessages []inbox.InboxMessage, fakeInboxPeekValue value.Value, maxWallTime time.Duration) (*protocol.ExecutionAssertion, []value.Value, uint64)

	MarshalForProof() ([]byte, error)

	MarshalBufferProof() ([]byte, error)

	MarshalState() ([]byte, error)

	Checkpoint(storage ArbStorage) bool
}
