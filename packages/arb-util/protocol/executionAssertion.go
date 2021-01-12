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
	"encoding/binary"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func NewExecutionAssertion(
	beforeMachineHash common.Hash,
	afterMachineHash common.Hash,
	numGas uint64,
	inboxMessagesConsumed uint64,
	outMsgsData []byte,
	outMsgsCount uint64,
	logsData []byte,
	logsCount uint64,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		NumGas:                numGas,
		BeforeMachineHash:     beforeMachineHash.MarshalToBuf(),
		AfterMachineHash:      afterMachineHash.MarshalToBuf(),
		InboxMessagesConsumed: inboxMessagesConsumed,
		OutMsgsData:           outMsgsData,
		OutMsgsCount:          outMsgsCount,
		LogsData:              logsData,
		LogsCount:             logsCount,
	}
}

func valuesToRaw(values []value.Value) []byte {
	var buf bytes.Buffer
	for _, val := range values {
		// Error can only occur with writes and bytes.Buffer is safe
		_ = value.MarshalValue(val, &buf)
	}
	return buf.Bytes()
}

func NewExecutionAssertionFromValues(
	beforeMachineHash common.Hash,
	afterMachineHash common.Hash,
	numGas uint64,
	inboxMessagesConsumed uint64,
	outMsgs []value.Value,
	logs []value.Value,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		BeforeMachineHash:     beforeMachineHash.MarshalToBuf(),
		AfterMachineHash:      afterMachineHash.MarshalToBuf(),
		NumGas:                numGas,
		InboxMessagesConsumed: inboxMessagesConsumed,
		OutMsgsData:           valuesToRaw(outMsgs),
		OutMsgsCount:          uint64(len(outMsgs)),
		LogsData:              valuesToRaw(logs),
		LogsCount:             uint64(len(logs)),
	}
}

func (x *ExecutionAssertion) Equals(b *ExecutionAssertion) bool {
	return bytes.Equal(x.BeforeMachineHash.Value, b.BeforeMachineHash.Value) &&
		bytes.Equal(x.AfterMachineHash.Value, b.AfterMachineHash.Value) &&
		x.NumGas == b.NumGas &&
		x.InboxMessagesConsumed == b.InboxMessagesConsumed &&
		x.OutMsgsCount == b.OutMsgsCount &&
		bytes.Equal(x.OutMsgsData, b.OutMsgsData) &&
		x.LogsCount == b.LogsCount &&
		bytes.Equal(x.LogsData, b.LogsData)
}

func (x *ExecutionAssertion) ParseOutMessages() [][]byte {
	vals := make([][]byte, 0, x.OutMsgsCount)
	rd := bytes.NewReader(x.OutMsgsData)
	for i := uint64(0); i < x.OutMsgsCount; i++ {
		var size uint64
		if err := binary.Read(rd, binary.BigEndian, &size); err != nil {
			panic(err)
		}
		arr := make([]byte, size)
		_, err := rd.Read(arr)
		if err != nil {
			panic(err)
		}
		vals = append(vals, arr)
	}
	return vals
}

func (x *ExecutionAssertion) ParseLogs() []value.Value {
	return BytesArrayToVals(x.LogsData, x.LogsCount)
}

func BytesArrayToVals(data []byte, valCount uint64) []value.Value {
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
