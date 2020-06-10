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

type Type uint8

const (
	TransactionType Type = iota
	EthType
	ERC20Type
	ERC721Type
	ContractTransactionType
	CallType
	TransactionBatchType
)

type SingleMessage interface {
	Message
	AsInboxValue() value.TupleValue
}

type ReceiptMessage interface {
	ReceiptHash() common.Hash
}

type Message interface {
	fmt.Stringer
	Type() Type
	Equals(o Message) bool

	CommitmentHash() common.Hash
	CheckpointValue() value.Value

	VMInboxMessages() []SingleMessage
}

type ExecutionMessage interface {
	SingleMessage
	GetFuncName() string
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

// UnmarshalExecuted converts the given Arbitrum message value which is the
// encoding of one of our value types for a VM, back into the original message
// type
func UnmarshalExecuted(typecode Type, messageVal value.Value, chain common.Address) (ExecutionMessage, error) {
	switch typecode {
	case TransactionType:
		return UnmarshalTransaction(messageVal, chain)
	case EthType:
		return UnmarshalEth(messageVal)
	case ERC20Type:
		return UnmarshalERC20(messageVal)
	case ERC721Type:
		return UnmarshalERC721(messageVal)
	case ContractTransactionType:
		return UnmarshalContractTransaction(messageVal)
	case CallType:
		return UnmarshalCall(messageVal)
	default:
		return nil, errors.New("invalid message type")
	}
}

func UnmarshalFromCheckpoint(msgType Type, v value.Value) (Message, error) {
	switch msgType {
	case TransactionType:
		return UnmarshalTransactionFromCheckpoint(v)
	case EthType:
		return UnmarshalEthFromCheckpoint(v)
	case ERC20Type:
		return UnmarshalERC20FromCheckpoint(v)
	case ERC721Type:
		return UnmarshalERC721FromCheckpoint(v)
	case ContractTransactionType:
		return UnmarshalContractTransactionFromCheckpoint(v)
	case TransactionBatchType:
		return UnmarshalTransactionBatchFromCheckpoint(v)
	case CallType:
		return UnmarshalCallFromCheckpoint(v)
	default:
		return nil, errors.New("bad message type")
	}
}

func AddToPrev(prev value.TupleValue, delivered Delivered) value.TupleValue {
	ret := prev
	for _, msg := range delivered.VMInboxMessages() {
		ret = value.NewTuple2(ret, msg.AsInboxValue())
	}
	return ret
}

func unmarshalTxWrapped(val value.Value, msgType Type) (common.Address, value.TupleValue, error) {
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

	if Type(msgTypeInt.BigInt().Uint64()) != msgType {
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

func unmarshalToken(val value.Value, msgType Type) (common.Address, common.Address, common.Address, *big.Int, error) {
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
