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

type ExecutionAssertion struct {
	NumGas                uint64
	InboxMessagesConsumed uint64
	Sends                 [][]byte
	SendAcc               common.Hash
	Logs                  []value.Value
	LogAcc                common.Hash
}

func NewExecutionAssertion(
	numGas uint64,
	inboxMessagesConsumed uint64,
	sendsData []byte,
	sendsCount uint64,
	sendAcc common.Hash,
	logsData []byte,
	logsCount uint64,
	logAcc common.Hash,
) *ExecutionAssertion {
	return &ExecutionAssertion{
		NumGas:                numGas,
		InboxMessagesConsumed: inboxMessagesConsumed,
		Sends:                 parseSends(sendsData, sendsCount),
		SendAcc:               sendAcc,
		Logs:                  BytesArrayToVals(logsData, logsCount),
		LogAcc:                logAcc,
	}
}

func parseSends(sendData []byte, sendCount uint64) [][]byte {
	vals := make([][]byte, 0, sendCount)
	rd := bytes.NewReader(sendData)
	for i := uint64(0); i < sendCount; i++ {
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
