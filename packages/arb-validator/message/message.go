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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type MessageType uint8

const (
	TransactionType MessageType = iota
	EthType
	ERC20Type
	ERC721Type
)

type Message interface {
	Type() MessageType
	AsValue() value.Value
}

type DeliveredMessage interface {
	Message
	Equals(o DeliveredMessage) bool
	CommitmentHash() common.Hash
	ReceiptHash() common.Hash
	DeliveredHeight() *common.TimeBlocks
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

func Unmarshal(msgType MessageType, v value.Value) (DeliveredMessage, error) {
	switch msgType {
	case TransactionType:
		return UnmarshalTransaction(v)
	case EthType:
		return UnmarshalEth(v)
	case ERC20Type:
		return UnmarshalERC20(v)
	case ERC721Type:
		return UnmarshalERC721(v)
	default:
		return nil, errors.New("bad message type")
	}
}

func DeliveredValue(m DeliveredMessage) value.Value {
	receiptHash := m.ReceiptHash()
	receiptVal := big.NewInt(0).SetBytes(receiptHash[:])
	msg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(m.DeliveredHeight().AsInt()),
		value.NewIntValue(receiptVal),
		m.AsValue(),
	})
	return msg
}
