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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
	Hash() (common.Hash, error)
	Clone() Machine

	CurrentStatus() Status
	IsBlocked(newMessages bool) BlockReason

	ExecuteAssertion(maxGas uint64, goOverGas bool, messages []inbox.InboxMessage, finalMessageOfBlock bool) (*protocol.ExecutionAssertion, []value.Value, uint64)
	ExecuteAssertionAdvanced(maxGas uint64, goOverGas bool, messages []inbox.InboxMessage, finalMessageOfBlock bool, sideloads []inbox.InboxMessage, stopOnSideload bool, beforeSendAcc common.Hash, beforeLogAcc common.Hash) (*protocol.ExecutionAssertion, []value.Value, uint64)

	MarshalForProof() ([]byte, []byte, error)

	MarshalState() ([]byte, error)
}
