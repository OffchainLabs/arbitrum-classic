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

package valprotocol

import (
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type ExecutionAssertionStub struct {
	AfterHash        common.Hash
	DidInboxInsn     bool
	NumGas           uint64
	FirstMessageHash common.Hash
	LastMessageHash  common.Hash
	FirstLogHash     common.Hash
	LastLogHash      common.Hash
}

func NewExecutionAssertionStubFromAssertion(a *protocol.ExecutionAssertion) *ExecutionAssertionStub {
	var lastMsgHash common.Hash
	for _, msg := range a.OutMsgs {
		lastMsgHash = hashing.SoliditySHA3(hashing.Bytes32(lastMsgHash), hashing.Bytes32(msg.Hash()))
	}
	var lastLogHash common.Hash
	for _, logVal := range a.Logs {
		lastLogHash = hashing.SoliditySHA3(hashing.Bytes32(lastLogHash), hashing.Bytes32(logVal.Hash()))
	}

	return &ExecutionAssertionStub{
		AfterHash:        a.AfterHash,
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: common.Hash{},
		LastMessageHash:  lastMsgHash,
		FirstLogHash:     common.Hash{},
		LastLogHash:      lastLogHash,
	}
}

func (a *ExecutionAssertionStub) MarshalToBuf() *ExecutionAssertionStubBuf {
	return &ExecutionAssertionStubBuf{
		AfterHash:        a.AfterHash.MarshalToBuf(),
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: a.FirstMessageHash.MarshalToBuf(),
		LastMessageHash:  a.LastMessageHash.MarshalToBuf(),
		FirstLogHash:     a.FirstLogHash.MarshalToBuf(),
		LastLogHash:      a.LastLogHash.MarshalToBuf(),
	}
}

func (a *ExecutionAssertionStubBuf) Unmarshal() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		AfterHash:        a.AfterHash.Unmarshal(),
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: a.FirstMessageHash.Unmarshal(),
		LastMessageHash:  a.LastMessageHash.Unmarshal(),
		FirstLogHash:     a.FirstLogHash.Unmarshal(),
		LastLogHash:      a.LastLogHash.Unmarshal(),
	}
}

func (a *ExecutionAssertionStub) Clone() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		AfterHash:        a.AfterHash,
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: a.FirstMessageHash,
		LastMessageHash:  a.LastMessageHash,
		FirstLogHash:     a.FirstLogHash,
		LastLogHash:      a.LastLogHash,
	}
}

func (a *ExecutionAssertionStub) String() string {
	return fmt.Sprintf(
		"Assertion(AfterHash: %v, DidInboxInsn: %v, NumGas: %v, "+
			"FirstMessageHash: %v, LastMessageHash: %v, FirstLogHash: %v LastLogHash: %v)",
		a.AfterHash,
		a.DidInboxInsn,
		a.NumGas,
		a.FirstMessageHash,
		a.LastMessageHash,
		a.FirstLogHash,
		a.LastLogHash,
	)
}

func (a *ExecutionAssertionStub) Equals(b *ExecutionAssertionStub) bool {
	return a.AfterHash == b.AfterHash &&
		a.NumGas == b.NumGas &&
		a.FirstMessageHash == b.FirstMessageHash &&
		a.LastMessageHash == b.LastMessageHash &&
		a.FirstLogHash == b.FirstLogHash &&
		a.LastLogHash == b.LastLogHash
}

func (a *ExecutionAssertionStub) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(a.AfterHash),
		hashing.Bool(a.DidInboxInsn),
		hashing.Uint64(a.NumGas),
		hashing.Bytes32(a.FirstMessageHash),
		hashing.Bytes32(a.LastMessageHash),
		hashing.Bytes32(a.FirstLogHash),
		hashing.Bytes32(a.LastLogHash),
	)
}
