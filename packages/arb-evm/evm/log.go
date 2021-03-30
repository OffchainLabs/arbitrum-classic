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
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

var logger = log.With().Stack().Caller().Stack().Str("component", "evm").Logger()

type Log struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

func CompareLogs(log1 Log, log2 Log) []string {
	var differences []string
	if log1.Address != log2.Address {
		differences = append(differences, fmt.Sprintf("different address %v and %v", log1.Address, log2.Address))
	}
	if len(log1.Topics) != len(log2.Topics) {
		differences = append(differences, fmt.Sprintf("different topic count %v and %v", len(log1.Topics), len(log2.Topics)))
	} else {
		for i, topic1 := range log1.Topics {
			topic2 := log2.Topics[i]
			if topic1 != topic2 {
				differences = append(differences, fmt.Sprintf("different topic %v and %v", topic1, topic2))
			}
		}
	}
	if !bytes.Equal(log1.Data, log2.Data) {
		differences = append(differences, fmt.Sprintf("different address 0x%X and 0x%X", log1.Data, log2.Data))
	}
	return differences
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
	if tupVal.Len() < 2 {
		return Log{}, errors.Errorf("log tuple must be at least size 2, but is %v", tupVal)
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
	logData, err := inbox.ByteArrayToBytes(logDataByteVal)
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

func LogStackToLogs(val value.Value) ([]Log, error) {
	logValues, err := inbox.StackValueToList(val)
	if err != nil {
		return nil, errors.Wrap(err, "log stack was not a stack")
	}
	logs := make([]Log, 0, len(logValues))
	for _, logVal := range logValues {
		evmLog, err := NewLogFromValue(logVal)
		if err != nil {
			return nil, err
		}
		logs = append(logs, evmLog)
	}
	return logs, nil
}
