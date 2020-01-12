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
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ExecutionAssertion struct {
	AfterHash    [32]byte
	DidInboxInsn bool
	NumGas       uint64
	OutMsgs      []value.Value
	Logs         []value.Value
}

func NewExecutionAssertion(afterHash [32]byte, didInboxInsn bool, numGas uint64, outMsgs []value.Value, logs []value.Value) *ExecutionAssertion {
	return &ExecutionAssertion{afterHash, didInboxInsn, numGas, outMsgs, logs}
}

func (a *ExecutionAssertion) Equals(b *ExecutionAssertion) bool {
	if a.AfterHash != b.AfterHash ||
		a.DidInboxInsn != b.DidInboxInsn ||
		a.NumGas != b.NumGas ||
		len(a.OutMsgs) != len(b.OutMsgs) ||
		len(a.Logs) != len(b.Logs) {
		return false
	}
	for i, ao := range a.OutMsgs {
		if !value.Eq(ao, b.OutMsgs[i]) {
			return false
		}
	}
	for i, ao := range a.Logs {
		if !value.Eq(ao, b.Logs[i]) {
			return false
		}
	}
	return true
}

func (a *ExecutionAssertion) LogsHash() [32]byte {
	var logHash [32]byte
	for _, logVal := range a.Logs {
		next := solsha3.SoliditySHA3(solsha3.Bytes32(logHash), solsha3.Bytes32(logVal.Hash()))
		copy(logHash[:], next)
	}
	return logHash
}

func (a *ExecutionAssertion) Stub() *ExecutionAssertionStub {
	var lastHash [32]byte
	for _, msg := range a.OutMsgs {
		next := solsha3.SoliditySHA3(solsha3.Bytes32(lastHash), solsha3.Bytes32(msg.Hash()))
		copy(lastHash[:], next)
	}

	return NewExecutionAssertionStub(
		a.AfterHash,
		a.DidInboxInsn,
		a.NumGas,
		lastHash,
		a.LogsHash(),
	)
}
