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

package inbox

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Type uint8

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

type InboxMessage struct {
	Kind        Type
	Sender      common.Address
	InboxSeqNum *big.Int
	GasPrice    *big.Int
	Data        []byte
	ChainTime   ChainTime
}

func NewInboxMessageFromData(data []byte) (InboxMessage, error) {
	if len(data) < 129 {
		return InboxMessage{}, errors.New("Not enough data for inbox message")
	}
	kind := Type(data[0])
	data = data[1:]

	var sender common.Address
	copy(sender[:], data[:])
	data = data[20:]

	blockNumber := common.NewTimeBlocks(new(big.Int).SetBytes(data[:32]))
	data = data[32:]

	timestamp := new(big.Int).SetBytes(data[:32])
	data = data[32:]

	inboxSeqNum := new(big.Int).SetBytes(data[:32])
	data = data[32:]

	gasPrice := new(big.Int).SetBytes(data[:32])
	data = data[32:]

	return InboxMessage{
		Kind:        kind,
		Sender:      sender,
		ChainTime:   ChainTime{BlockNum: blockNumber, Timestamp: timestamp},
		InboxSeqNum: inboxSeqNum,
		GasPrice:    gasPrice,
		Data:        data,
	}, nil
}

func NewInboxMessageFromValue(val value.Value) (InboxMessage, error) {
	failRet := InboxMessage{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("val must be a tuple")
	}
	if tup.Len() != 8 {
		return failRet, errors.Errorf("expected tuple of length 8, but recieved tuple of length %v", tup.Len())
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	blockNumber, _ := tup.GetByInt64(1)
	timestamp, _ := tup.GetByInt64(2)
	sender, _ := tup.GetByInt64(3)
	inboxSeqNum, _ := tup.GetByInt64(4)
	gasPriceL1, _ := tup.GetByInt64(5)
	msgSize, _ := tup.GetByInt64(6)
	msgData, _ := tup.GetByInt64(7)

	kindInt, ok := kind.(value.IntValue)
	if !ok {
		return failRet, errors.New("inbox message kind must be an int")
	}

	blockNumberInt, ok := blockNumber.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNumber must be an int")
	}

	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be an int")
	}

	senderInt, ok := sender.(value.IntValue)
	if !ok {
		return failRet, errors.New("sender must be an int")
	}

	inboxSeqNumInt, ok := inboxSeqNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("inboxSeqNum must be an int")
	}

	gasPriceInt, ok := gasPriceL1.(value.IntValue)
	if !ok {
		return failRet, errors.New("gasPriceL1 must be an int")
	}

	msgSizeInt, ok := msgSize.(value.IntValue)
	if !ok {
		return failRet, errors.New("msgSize must be an int")
	}
	msgDataBuffer, ok := msgData.(*value.Buffer)
	if !ok {
		return failRet, errors.New("msgData must be an buffer")
	}

	data, err := BufAndLengthToBytes(msgSizeInt.BigInt(), msgDataBuffer)
	if err != nil {
		return failRet, errors.Wrap(err, "unmarshalling input data")
	}

	return InboxMessage{
		Kind:        Type(kindInt.BigInt().Uint64()),
		Sender:      NewAddressFromInt(senderInt),
		InboxSeqNum: inboxSeqNumInt.BigInt(),
		GasPrice:    gasPriceInt.BigInt(),
		Data:        data,
		ChainTime: ChainTime{
			BlockNum:  common.NewTimeBlocks(blockNumberInt.BigInt()),
			Timestamp: timestampInt.BigInt(),
		},
	}, nil
}

func NewRandomInboxMessage() InboxMessage {
	return InboxMessage{
		Kind:        Type(rand.Uint32()),
		Sender:      common.RandAddress(),
		InboxSeqNum: common.RandBigInt(),
		Data:        common.RandBytes(200),
		ChainTime:   NewRandomChainTime(),
	}
}

func (im InboxMessage) String() string {
	return fmt.Sprintf(
		"InboxMessage(%v, %v, %v, %v, %v)",
		im.Kind,
		im.Sender,
		im.InboxSeqNum,
		hexutil.Encode(im.Data),
		im.ChainTime,
	)
}

func (im InboxMessage) AsValue() value.Value {
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(im.Kind)),
		value.NewIntValue(im.ChainTime.BlockNum.AsInt()),
		value.NewIntValue(im.ChainTime.Timestamp),
		NewIntFromAddress(im.Sender),
		value.NewIntValue(im.InboxSeqNum),
		value.NewIntValue(im.GasPrice),
		BytesToByteStack(im.Data),
	})
	return tup
}

func (im InboxMessage) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(im.Kind)),
		hashing.Address(im.Sender),
		hashing.Uint256(im.ChainTime.BlockNum.AsInt()),
		hashing.Uint256(im.ChainTime.Timestamp),
		hashing.Uint256(im.InboxSeqNum),
		hashing.Uint256(im.GasPrice),
		hashing.Bytes32(hashing.SoliditySHA3(im.Data)),
	)
}

func (im InboxMessage) Equals(o InboxMessage) bool {
	return im.Kind == o.Kind &&
		im.Sender == o.Sender &&
		im.InboxSeqNum.Cmp(o.InboxSeqNum) == 0 &&
		bytes.Equal(im.Data, o.Data) &&
		im.GasPrice.Cmp(o.GasPrice) == 0 &&
		im.ChainTime.BlockNum.AsInt().Cmp(o.ChainTime.BlockNum.AsInt()) == 0 &&
		im.ChainTime.Timestamp.Cmp(o.ChainTime.Timestamp) == 0
}

func NewIntFromAddress(address common.Address) value.IntValue {
	addressBytes := [32]byte{}
	copy(addressBytes[12:], address[:])
	addressVal := big.NewInt(0).SetBytes(addressBytes[:])

	return value.NewIntValue(addressVal)
}

func NewAddressFromInt(val value.IntValue) common.Address {
	address := common.Address{}
	valBytes := val.ToBytes()
	copy(address[:], valBytes[12:])
	return address
}

func (im InboxMessage) ToBytes() []byte {
	var data []byte
	data = append(data, uint8(im.Kind))
	data = append(data, im.Sender[:]...)
	data = append(data, math.U256Bytes(im.ChainTime.BlockNum.AsInt())...)
	data = append(data, math.U256Bytes(im.ChainTime.Timestamp)...)
	data = append(data, math.U256Bytes(im.InboxSeqNum)...)
	data = append(data, math.U256Bytes(im.GasPrice)...)
	data = append(data, im.Data...)
	return data
}
