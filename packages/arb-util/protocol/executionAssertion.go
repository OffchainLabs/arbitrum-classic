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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ExecutionAssertion struct {
	AfterHash    common.Hash
	DidInboxInsn bool
	NumGas       uint64
	OutMsgs      []value.Value
	Logs         []value.Value
}

func NewExecutionAssertion(afterHash common.Hash, didInboxInsn bool, numGas uint64, outMsgs []value.Value, logs []value.Value) *ExecutionAssertion {
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
