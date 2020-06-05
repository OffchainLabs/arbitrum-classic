/*
* Copyright 2020, Offchain Labs, Inc.
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

package message

import (
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ChainTime struct {
	BlockNum  *common.TimeBlocks
	Timestamp *big.Int
}

type DeliveryInfo struct {
	ChainTime
	MessageNum *big.Int
}

type Received struct {
	Message Message
	ChainTime
}

func (m Received) Equals(o Received) bool {
	return m.Message.Equals(o.Message) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0
}

type Delivered struct {
	Message Message
	DeliveryInfo
}

func (m Delivered) GetReceived() Received {
	return Received{
		Message:   m.Message,
		ChainTime: m.ChainTime,
	}
}

func (m Delivered) Equals(o Delivered) bool {
	return m.Message.Equals(o.Message) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0 &&
		m.MessageNum.Cmp(o.MessageNum) == 0
}

func (m Delivered) deliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m Delivered) deliveredTimestamp() *big.Int {
	return m.Timestamp
}

func (m Delivered) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(m.Message.CommitmentHash()),
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.Timestamp),
		hashing.Uint256(m.MessageNum),
	)
}

func (m Delivered) ReceiptHash() common.Hash {
	if msg, ok := m.Message.(ReceiptMessage); ok {
		return msg.ReceiptHash()
	}
	return value.NewIntValue(m.MessageNum).ToBytes()
}

func (m Delivered) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(m.Message.Type())),
		m.Message.CheckpointValue(),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
		value.NewIntValue(new(big.Int).Set(m.MessageNum)),
	})
	return val
}

func UnmarshalDeliveredFromCheckpoint(v value.Value) (Delivered, error) {
	tup, ok := v.(value.TupleValue)
	failRet := Delivered{}
	if !ok || tup.Len() != 5 {
		return failRet, errors.New("delivered val must be 5-tuple")
	}
	typecode, _ := tup.GetByInt64(0)
	typecodeInt, ok := typecode.(value.IntValue)
	msgType := Type(typecodeInt.BigInt().Uint64())

	msg, _ := tup.GetByInt64(1)
	message, err := UnmarshalFromCheckpoint(msgType, msg)
	if err != nil {
		return failRet, err
	}
	blockNum, _ := tup.GetByInt64(2)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNum must be int")
	}
	timestamp, _ := tup.GetByInt64(3)
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be int")
	}
	messageNum, _ := tup.GetByInt64(4)
	messageNumInt, ok := messageNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("messageNum must be int")
	}

	return Delivered{
		Message: message,
		DeliveryInfo: DeliveryInfo{
			ChainTime: ChainTime{
				BlockNum:  common.NewTimeBlocks(blockNumInt.BigInt()),
				Timestamp: timestampInt.BigInt(),
			},
			MessageNum: messageNumInt.BigInt(),
		},
	}, nil
}
