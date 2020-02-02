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

func NewExecutionAssertion(
	afterHash common.Hash,
	didInboxInsn bool,
	numGas uint64,
	outMsgs []value.Value,
	logs []value.Value,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		AfterHash:    afterHash,
		DidInboxInsn: didInboxInsn,
		NumGas:       numGas,
		OutMsgs:      outMsgs,
		Logs:         logs,
	}
}

func (a *ExecutionAssertion) MarshalToBuf() *ExecutionAssertionBuf {
	messages := make([][]byte, 0, len(a.OutMsgs))
	for _, msg := range a.OutMsgs {
		valBytes := value.MarshalValueToBytes(msg)
		messages = append(messages, valBytes)
	}
	logs := make([][]byte, 0, len(a.Logs))
	for _, msg := range a.OutMsgs {
		valBytes := value.MarshalValueToBytes(msg)
		logs = append(logs, valBytes)
	}

	return &ExecutionAssertionBuf{
		AfterHash:    a.AfterHash.MarshalToBuf(),
		DidInboxInsn: a.DidInboxInsn,
		NumGas:       a.NumGas,
		Messages:     messages,
		Logs:         logs,
	}
}

func (a *ExecutionAssertionBuf) Unmarshal() (*ExecutionAssertion, error) {
	messages := make([]value.Value, 0, len(a.Logs))
	for _, valLog := range a.Messages {
		val, err := value.UnmarshalValueFromBytes(valLog)
		if err != nil {
			return nil, err
		}
		messages = append(messages, val)
	}

	logs := make([]value.Value, 0, len(a.Logs))
	for _, valLog := range a.Logs {
		val, err := value.UnmarshalValueFromBytes(valLog)
		if err != nil {
			return nil, err
		}
		logs = append(logs, val)
	}

	return &ExecutionAssertion{
		AfterHash:    a.AfterHash.Unmarshal(),
		DidInboxInsn: a.DidInboxInsn,
		NumGas:       a.NumGas,
		OutMsgs:      messages,
		Logs:         logs,
	}, nil
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
