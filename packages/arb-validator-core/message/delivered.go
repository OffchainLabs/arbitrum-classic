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
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ChainTime struct {
	BlockNum  *common.TimeBlocks
	Timestamp *big.Int
}

func NewRandomChainTime() ChainTime {
	return ChainTime{
		BlockNum:  common.NewTimeBlocks(common.RandBigInt()),
		Timestamp: common.RandBigInt(),
	}
}

type DeliveryInfo struct {
	ChainTime
	TxId *big.Int
}

func (d DeliveryInfo) Equals(o DeliveryInfo) bool {
	return d.BlockNum.Cmp(o.BlockNum) == 0 &&
		d.Timestamp.Cmp(o.Timestamp) == 0 &&
		d.TxId.Cmp(o.TxId) == 0
}

func (d DeliveryInfo) deliveredHeight() *common.TimeBlocks {
	return d.BlockNum
}

func (d DeliveryInfo) deliveredTimestamp() *big.Int {
	return d.Timestamp
}

func (d DeliveryInfo) TxHash() common.Hash {
	var hash common.Hash
	copy(hash[:], math.U256Bytes(d.TxId))
	return hash
}

func NewRandomDeliveryInfo() DeliveryInfo {
	return DeliveryInfo{
		ChainTime: NewRandomChainTime(),
		TxId:      common.RandBigInt(),
	}
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

func (m Received) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(m.Message.Type())),
		m.Message.CheckpointValue(),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
	})
	return val
}

func UnmarshalReceivedFromCheckpoint(v value.Value) (Received, error) {
	tup, ok := v.(value.TupleValue)
	failRet := Received{}
	if !ok || tup.Len() != 4 {
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

	return Received{
		Message: message,
		ChainTime: ChainTime{
			BlockNum:  common.NewTimeBlocks(blockNumInt.BigInt()),
			Timestamp: timestampInt.BigInt(),
		},
	}, nil
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

func (m Delivered) VMInboxMessages() []SingleDelivered {
	messages := m.Message.VMInboxMessages()
	ret := make([]SingleDelivered, 0, len(messages))
	for _, msg := range messages {
		ret = append(ret, SingleDelivered{
			Message:      msg,
			DeliveryInfo: m.DeliveryInfo,
		})
	}
	return ret
}

func (m Delivered) Equals(o Delivered) bool {
	return m.Message.Equals(o.Message) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0 &&
		m.TxId.Cmp(o.TxId) == 0
}

func (m Delivered) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(m.Message.CommitmentHash()),
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.Timestamp),
		hashing.Uint256(m.TxId),
	)
}

type RawDelivered struct {
	Type    Type
	Message value.Value
	DeliveryInfo
}

func UnmarshalRawDelivered(val value.Value) (RawDelivered, error) {
	tup, ok := val.(value.TupleValue)
	invalid := RawDelivered{}
	if !ok {
		return invalid, errors.New("msg must be tuple value")
	}
	if tup.Len() != 4 {
		return invalid, fmt.Errorf("expected tuple of length 4, but recieved %v", tup)
	}
	blockNumberVal, _ := tup.GetByInt64(0)
	timestampVal, _ := tup.GetByInt64(1)
	txIdVal, _ := tup.GetByInt64(2)
	restVal, _ := tup.GetByInt64(3)

	blockNumberInt, ok := blockNumberVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("block number must be an int")
	}

	timestampInt, ok := timestampVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("timestamp must be an int")
	}

	txId, ok := txIdVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("tx hash must be an int")
	}

	restValTup, ok := restVal.(value.TupleValue)
	if !ok {
		return invalid, errors.New("message must be a tup")
	}

	typeVal, _ := restValTup.GetByInt64(0)
	typeInt, ok := typeVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("type must be an int")
	}
	typecode := uint8(typeInt.BigInt().Uint64())

	return RawDelivered{
		Type:    Type(typecode),
		Message: restValTup,
		DeliveryInfo: DeliveryInfo{
			ChainTime: ChainTime{
				BlockNum:  common.NewTimeBlocks(blockNumberInt.BigInt()),
				Timestamp: timestampInt.BigInt(),
			},
			TxId: txId.BigInt(),
		},
	}, nil
}

func (m Delivered) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(m.Message.Type())),
		m.Message.CheckpointValue(),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
		value.NewIntValue(new(big.Int).Set(m.TxId)),
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
			TxId: messageNumInt.BigInt(),
		},
	}, nil
}

type SingleDelivered struct {
	Message SingleMessage
	DeliveryInfo
}

func NewSingleDelivered(d Delivered) (SingleDelivered, error) {
	msg, ok := d.Message.(ExecutionMessage)
	if !ok {
		return SingleDelivered{}, errors.New("must construct from execution message")
	}
	return SingleDelivered{Message: msg}, nil
}

func (m SingleDelivered) Equals(o SingleDelivered) bool {
	return m.Message.Equals(o.Message) &&
		m.DeliveryInfo.Equals(o.DeliveryInfo)
}

func UnmarshalSingleDelivered(val value.Value, chain common.Address) (SingleDelivered, error) {
	invalid := SingleDelivered{}
	rawDelivered, err := UnmarshalRawDelivered(val)
	if err != nil {
		return invalid, err
	}
	arbMessage, err := UnmarshalExecuted(rawDelivered.Type, rawDelivered.Message, chain)
	if err != nil {
		return invalid, err
	}

	return SingleDelivered{
		Message:      arbMessage,
		DeliveryInfo: rawDelivered.DeliveryInfo,
	}, nil
}

func NewRandomSingleDelivered(msg ExecutionMessage) SingleDelivered {
	return SingleDelivered{
		Message:      msg,
		DeliveryInfo: NewRandomDeliveryInfo(),
	}
}

func (s SingleDelivered) ExectutedMessage() ExecutionMessage {
	return s.Message.(ExecutionMessage)
}

func (s SingleDelivered) ReceiptHash() common.Hash {
	if msg, ok := s.Message.(ReceiptMessage); ok {
		return msg.ReceiptHash()
	}
	return value.NewIntValue(s.TxId).ToBytes()
}

func (s SingleDelivered) AsInboxValue() value.Value {
	receiptHash := s.ReceiptHash()
	receiptVal := big.NewInt(0).SetBytes(receiptHash[:])
	msg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(s.deliveredHeight().AsInt()),
		value.NewIntValue(s.deliveredTimestamp()),
		value.NewIntValue(receiptVal),
		s.Message.AsInboxValue(),
	})
	return msg
}
