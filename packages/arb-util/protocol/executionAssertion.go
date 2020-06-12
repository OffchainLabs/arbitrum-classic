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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func NewExecutionAssertion(
	afterHash common.Hash,
	didInboxInsn bool,
	numGas uint64,
	outMsgsData []byte,
	outMsgsCount uint64,
	logsData []byte,
	logsCount uint64,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		AfterHash:    afterHash.MarshalToBuf(),
		DidInboxInsn: didInboxInsn,
		NumGas:       numGas,
		OutMsgsData:  outMsgsData,
		OutMsgsCount: outMsgsCount,
		LogsData:     logsData,
		LogsCount:    logsCount,
	}
}

func valuesToRaw(values []value.Value) []byte {
	var buf bytes.Buffer
	for _, val := range values {
		_ = value.MarshalValue(val, &buf)
	}
	return buf.Bytes()
}

func NewExecutionAssertionFromValues(
	afterHash common.Hash,
	didInboxInsn bool,
	numGas uint64,
	outMsgs []value.Value,
	logs []value.Value,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		AfterHash:    afterHash.MarshalToBuf(),
		DidInboxInsn: didInboxInsn,
		NumGas:       numGas,
		OutMsgsData:  valuesToRaw(outMsgs),
		OutMsgsCount: uint64(len(outMsgs)),
		LogsData:     valuesToRaw(logs),
		LogsCount:    uint64(len(logs)),
	}
}

func (a *ExecutionAssertion) Equals(b *ExecutionAssertion) bool {
	return a.AfterHash == b.AfterHash &&
		a.DidInboxInsn != b.DidInboxInsn &&
		a.NumGas != b.NumGas &&
		a.OutMsgsCount != b.OutMsgsCount &&
		bytes.Equal(a.OutMsgsData, b.OutMsgsData) &&
		a.LogsCount != b.LogsCount &&
		bytes.Equal(a.LogsData, b.LogsData)

}

func (a *ExecutionAssertion) ParseOutMessages() []value.Value {
	return bytesArrayToVals(a.OutMsgsData, a.OutMsgsCount)
}

func (a *ExecutionAssertion) ParseLogs() []value.Value {
	return bytesArrayToVals(a.LogsData, a.LogsCount)
}

func bytesArrayToVals(data []byte, valCount uint64) []value.Value {
	rd := bytes.NewReader(data)
	vals := make([]value.Value, 0, valCount)
	for i := uint64(0); i < valCount; i++ {
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return vals
}
