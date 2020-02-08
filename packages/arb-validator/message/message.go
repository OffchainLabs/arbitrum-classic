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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type MessageType uint8

const (
	TransactionType MessageType = iota
	EthType
	ERC20Type
	ERC721Type
	CallType
)

type UnsentMessage interface {
	Type() MessageType
	AsValue() value.Value
	GetFuncName() string
}

type Message interface {
	UnsentMessage
	Equals(o Message) bool
	ReceiptHash() common.Hash
	DeliveredHeight() *common.TimeBlocks
}

type InboxMessage interface {
	Message
	CommitmentHash() common.Hash
	CheckpointValue() value.Value
}

func addressToIntValue(address common.Address) value.IntValue {
	addressBytes := [32]byte{}
	copy(addressBytes[12:], address[:])
	addressVal := big.NewInt(0).SetBytes(addressBytes[:])

	return value.NewIntValue(addressVal)
}

func intValueToAddress(val value.IntValue) common.Address {
	address := common.Address{}
	valBytes := val.ToBytes()
	copy(address[:], valBytes[12:])
	return address
}

func UnmarshalFromCheckpoint(msgType MessageType, v value.Value) (InboxMessage, error) {
	switch msgType {
	case TransactionType:
		return UnmarshalTransactionFromCheckpoint(v)
	case EthType:
		return UnmarshalEthFromCheckpoint(v)
	case ERC20Type:
		return UnmarshalERC20FromCheckpoint(v)
	case ERC721Type:
		return UnmarshalERC721FromCheckpoint(v)
	default:
		return nil, errors.New("bad message type")
	}
}

func DeliveredValue(m Message) value.Value {
	receiptHash := m.ReceiptHash()
	receiptVal := big.NewInt(0).SetBytes(receiptHash[:])
	msg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(m.DeliveredHeight().AsInt()),
		value.NewIntValue(receiptVal),
		m.AsValue(),
	})
	return msg
}

func unmarshalTxWrapped(val value.Value, msgType MessageType) (common.Address, value.TupleValue, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return common.Address{}, value.TupleValue{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 3 {
		return common.Address{}, value.TupleValue{}, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	msgTypeVal, _ := tup.GetByInt64(0)
	msgTypeInt, ok := msgTypeVal.(value.IntValue)
	if !ok {
		return common.Address{}, value.TupleValue{}, errors.New("msg type must be an int")
	}

	if MessageType(msgTypeInt.BigInt().Uint64()) != msgType {
		return common.Address{}, value.TupleValue{}, errors.New("wrong msg type")
	}

	fromVal, _ := tup.GetByInt64(1)
	fromInt, ok := fromVal.(value.IntValue)
	if !ok {
		return common.Address{}, value.TupleValue{}, errors.New("from must be an int")
	}
	val2, _ := tup.GetByInt64(2)

	tup2, ok := val2.(value.TupleValue)
	if !ok {
		return common.Address{}, value.TupleValue{}, fmt.Errorf("expected tuple, but recieved %v", tup2)
	}
	return intValueToAddress(fromInt), tup2, nil
}

func unmarshalToken(val value.Value, msgType MessageType) (common.Address, common.Address, common.Address, *big.Int, error) {
	from, tup, err := unmarshalTxWrapped(val, msgType)
	if err != nil {
		return common.Address{}, common.Address{}, common.Address{}, nil, err
	}

	if tup.Len() != 3 {
		return common.Address{}, common.Address{}, common.Address{}, nil, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	tokenVal, _ := tup.GetByInt64(0)
	destVal, _ := tup.GetByInt64(1)
	amountVal, _ := tup.GetByInt64(2)

	tokenInt, ok := tokenVal.(value.IntValue)
	if !ok {
		return common.Address{}, common.Address{}, common.Address{}, nil, errors.New("token must be an int")
	}

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return common.Address{}, common.Address{}, common.Address{}, nil, errors.New("dest must be an int")
	}

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return common.Address{}, common.Address{}, common.Address{}, nil, errors.New("amount must be an int")
	}

	return from, intValueToAddress(tokenInt), intValueToAddress(destInt), amountInt.BigInt(), nil
}
