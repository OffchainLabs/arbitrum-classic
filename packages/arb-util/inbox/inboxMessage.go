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
	Data        []byte
	ChainTime   ChainTime
}

func NewInboxMessageFromValue(val value.Value) (InboxMessage, error) {
	failRet := InboxMessage{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("val must be a tuple")
	}
	if tup.Len() != 6 {
		return failRet, errors.Errorf("expected tuple of length 6, but recieved tuple of length %v", tup.Len())
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	blockNumber, _ := tup.GetByInt64(1)
	timestamp, _ := tup.GetByInt64(2)
	sender, _ := tup.GetByInt64(3)
	inboxSeqNum, _ := tup.GetByInt64(4)
	messageData, _ := tup.GetByInt64(5)

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

	data, err := ByteStackToHex(messageData)
	if err != nil {
		return failRet, errors.Wrap(err, "unmarshalling input data")
	}

	return InboxMessage{
		Kind:        Type(kindInt.BigInt().Uint64()),
		Sender:      NewAddressFromInt(senderInt),
		InboxSeqNum: inboxSeqNumInt.BigInt(),
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
		hashing.Bytes32(hashing.SoliditySHA3(im.Data)),
	)
}

func (im InboxMessage) Equals(o InboxMessage) bool {
	return im.Kind == o.Kind &&
		im.Sender == o.Sender &&
		im.InboxSeqNum.Cmp(o.InboxSeqNum) == 0 &&
		bytes.Equal(im.Data, o.Data) &&
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
