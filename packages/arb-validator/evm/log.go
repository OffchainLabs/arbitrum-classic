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

package evm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Log struct {
	ContractID value.IntValue
	Data       []byte
	Topics     []common.Hash
}

func (l Log) String() string {
	var sb strings.Builder
	sb.WriteString("Log(contract: ")
	sb.WriteString(l.ContractID.String())
	sb.WriteString(", topics: [")
	for i, topic := range l.Topics {
		sb.WriteString(hexutil.Encode(topic[:]))
		if i != len(l.Topics)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("], data:")
	sb.WriteString(hexutil.Encode(l.Data))
	sb.WriteString(")")
	return ""
}

func LogValToLog(val value.Value) (Log, error) {
	tupVal, ok := val.(value.TupleValue)
	if !ok {
		return Log{}, errors.New("log must be a tuple")
	}
	if tupVal.Len() < 3 {
		return Log{}, fmt.Errorf("log tuple must be at least size 3, but is %v", tupVal)
	}
	contractIDVal, _ := tupVal.GetByInt64(0)
	contractIDInt, ok := contractIDVal.(value.IntValue)
	if !ok {
		return Log{}, errors.New("log contract id must be an int")
	}
	logDataByteVal, _ := tupVal.GetByInt64(1)
	logData, err := message.ByteStackToHex(logDataByteVal)
	if err != nil {
		return Log{}, err
	}
	topics := make([]common.Hash, 0, tupVal.Len()-2)
	for _, topicVal := range tupVal.Contents()[2:] {
		topicValInt, ok := topicVal.(value.IntValue)
		if !ok {
			return Log{}, errors.New("log topic must be an int")
		}
		topics = append(topics, topicValInt.ToBytes())
	}

	return Log{contractIDInt, logData, topics}, nil
}

func LogStackToLogs(val value.Value) ([]Log, error) {
	logValues, err := message.StackValueToList(val)
	if err != nil {
		return nil, err
	}
	logs := make([]Log, 0, len(logValues))
	for _, logVal := range logValues {
		log, err := LogValToLog(logVal)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
