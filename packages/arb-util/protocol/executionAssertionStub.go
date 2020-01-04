/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package protocol

import (
	"bytes"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewExecutionAssertionStub(
	afterHash [32]byte,
	didInboxInsn bool,
	numGas uint64,
	messagesAcc [32]byte,
	logsAcc [32]byte,
) *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		AfterHash:        value.NewHashBuf(afterHash),
		DidInboxInsn:     didInboxInsn,
		NumGas:           numGas,
		FirstMessageHash: value.NewHashBuf([32]byte{}),
		LastMessageHash:  value.NewHashBuf(messagesAcc),
		FirstLogHash:     value.NewHashBuf([32]byte{}),
		LastLogHash:      value.NewHashBuf(logsAcc),
	}
}

func (a *ExecutionAssertionStub) Equals(b *ExecutionAssertionStub) bool {
	if a.AfterHash != b.AfterHash ||
		a.NumGas != b.NumGas ||
		!bytes.Equal(a.FirstMessageHash.Value, b.FirstMessageHash.Value) ||
		!bytes.Equal(a.LastMessageHash.Value, b.LastMessageHash.Value) ||
		!bytes.Equal(a.FirstLogHash.Value, b.FirstLogHash.Value) ||
		!bytes.Equal(a.LastLogHash.Value, b.LastLogHash.Value) {
		return false
	}
	return true
}

func (a *ExecutionAssertionStub) AfterHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], a.AfterHash.Value)
	return ret
}

func (a *ExecutionAssertionStub) FirstMessageHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], a.FirstMessageHash.Value)
	return ret
}

func (a *ExecutionAssertionStub) LastMessageHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], a.LastMessageHash.Value)
	return ret
}

func (a *ExecutionAssertionStub) FirstLogHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], a.FirstLogHash.Value)
	return ret
}

func (a *ExecutionAssertionStub) LastLogHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], a.LastLogHash.Value)
	return ret
}

func (a *ExecutionAssertionStub) Hash() [32]byte {
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(a.AfterHash.Value),
		solsha3.Uint64(a.NumGas),
		solsha3.Bytes32(a.FirstMessageHash.Value),
		solsha3.Bytes32(a.LastMessageHash.Value),
		solsha3.Bytes32(a.FirstLogHash.Value),
		solsha3.Bytes32(a.LastLogHash.Value),
	)
	copy(ret[:], hashVal)
	return ret
}

func (a *ExecutionAssertionStub) GeneratePostcondition(pre *Precondition) *Precondition {
	nextBeforeInbox := pre.BeforeInbox
	if a.DidInboxInsn {
		nextBeforeInbox = value.NewHashBuf(value.NewEmptyTuple().Hash())
	}
	return &Precondition{
		BeforeHash:  a.AfterHash,
		TimeBounds:  pre.TimeBounds,
		BeforeInbox: nextBeforeInbox,
	}
}

func GeneratePreconditions(pre *Precondition, assertions []*ExecutionAssertionStub) []*Precondition {
	preconditions := make([]*Precondition, 0, len(assertions))
	for _, assertion := range assertions {
		preconditions = append(preconditions, pre)
		pre = assertion.GeneratePostcondition(pre)
	}
	return preconditions
}
