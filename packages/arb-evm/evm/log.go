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
	"bytes"
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	errors2 "github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Log struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

func NewRandomLog(topicCount int32) Log {
	topics := make([]common.Hash, 0, topicCount)
	for i := int32(0); i < topicCount; i++ {
		topics = append(topics, common.RandHash())
	}
	return Log{
		Address: common.RandAddress(),
		Topics:  topics,
		Data:    common.RandBytes(200),
	}
}

func (l Log) MatchesQuery(addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		match := false
		for _, addr := range addresses {
			if l.Address == addr {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	if len(topics) > len(l.Topics) {
		return false
	}

	for i, topicGroup := range topics {
		if len(topicGroup) == 0 {
			continue
		}
		match := false
		for _, topic := range topicGroup {
			if l.Topics[i] == topic {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func (l Log) Equals(o Log) bool {
	if len(l.Topics) != len(o.Topics) {
		return false
	}
	for i, topic := range l.Topics {
		if topic != o.Topics[i] {
			return false
		}
	}
	return l.Address == o.Address &&
		bytes.Equal(l.Data, o.Data)
}

func (l Log) String() string {
	var sb strings.Builder
	sb.WriteString("Log(contract: ")
	sb.WriteString(l.Address.String())
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
	return sb.String()
}

func NewLogFromValue(val value.Value) (Log, error) {
	tupVal, ok := val.(*value.TupleValue)
	if !ok {
		return Log{}, errors.New("log must be a tuple")
	}
	if tupVal.Len() < 3 {
		return Log{}, fmt.Errorf("log tuple must be at least size 3, but is %v", tupVal)
	}

	// Tuple size already verified above, so error can be ignored
	contractIDVal, _ := tupVal.GetByInt64(0)
	contractIDInt, ok := contractIDVal.(value.IntValue)
	if !ok {
		return Log{}, errors.New("log contract id must be an int")
	}
	contractIDBytes := contractIDInt.ToBytes()
	var address common.Address
	copy(address[:], contractIDBytes[12:])
	logDataByteVal, _ := tupVal.GetByInt64(1)
	logData, err := inbox.ByteStackToHex(logDataByteVal)
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

	return Log{address, topics, logData}, nil
}

func (l Log) AsValue() (*value.TupleValue, error) {
	data := []value.Value{
		value.NewValueFromAddress(l.Address),
		inbox.BytesToByteStack(l.Data),
	}
	for _, topic := range l.Topics {
		data = append(data, value.NewIntValue(new(big.Int).SetBytes(topic.Bytes())))
	}
	return value.NewTupleFromSlice(data)
}

func LogStackToLogs(val value.Value) ([]Log, error) {
	logValues, err := inbox.StackValueToList(val)
	if err != nil {
		return nil, errors2.Wrap(err, "log stack was not a stack")
	}
	logs := make([]Log, 0, len(logValues))
	for i := range logValues {
		// Flip the order of the logs
		log, err := NewLogFromValue(logValues[len(logValues)-1-i])
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func LogsToLogStack(logs []Log) *value.TupleValue {
	logValues := make([]value.Value, 0, len(logs))
	for i := range logs {
		logValue, err := logs[len(logs)-1-i].AsValue()
		if err == nil {
			logValues = append(logValues, logValue)
		}
	}
	return inbox.ListToStackValue(logValues)
}
