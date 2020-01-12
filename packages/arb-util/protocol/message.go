/*
 * Copyright 2019, Offchain Labs, Inc.
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

package protocol

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Identity [32]byte
type TokenType [21]byte

func TokenTypeFromIntValue(val value.IntValue) TokenType {
	var tokType TokenType
	tokBytes := val.ToBytes()
	copy(tokType[:], tokBytes[11:])
	return tokType
}

func (t TokenType) ToIntValue() value.IntValue {
	var bigtok [32]byte
	copy(bigtok[11:], t[:])
	return value.NewIntValue(new(big.Int).SetBytes(bigtok[:]))
}

func (t TokenType) IsToken() bool {
	return t[20] == 0
}

func tokenTypeEncoded(input [21]byte) []byte {
	return common.RightPadBytes(input[:], 21)
}

func TokenTypeArrayEncoded(input [][21]byte) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, common.RightPadBytes(val[:], 32)...)
	}
	return values
}

type Message struct {
	Data        value.Value
	TokenType   [21]byte
	Currency    *big.Int
	Destination common.Address
}

func NewMessage(data value.Value, tokenType [21]byte, currency *big.Int, destination *big.Int) Message {
	var dest common.Address
	destBytes := value.NewIntValue(destination).ToBytes()
	copy(dest[:], destBytes[12:])
	return Message{data, tokenType, new(big.Int).Set(currency), dest}
}

func NewSimpleMessage(data value.Value, tokenType [21]byte, currency *big.Int, sender common.Address) Message {
	return Message{data, tokenType, currency, sender}
}

func NewMessageFromValue(val value.Value) (Message, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return Message{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 4 {
		return Message{}, fmt.Errorf("advise expected tuple of length 5, but recieved %v", tup)
	}
	data, _ := tup.GetByInt64(0)
	destVal, _ := tup.GetByInt64(1)
	amountVal, _ := tup.GetByInt64(2)
	typeVal, _ := tup.GetByInt64(3)

	typeInt, ok := typeVal.(value.IntValue)
	if !ok {
		return Message{}, errors.New("type must be an int")
	}

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return Message{}, errors.New("value must be an int")
	}

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return Message{}, errors.New("value must be an int")
	}

	typeBytes := typeInt.ToBytes()
	var tokenType [21]byte
	copy(tokenType[:], typeBytes[:21])

	return NewMessage(
		data,
		tokenType,
		amountInt.BigInt(),
		destInt.BigInt(),
	), nil
}

func (msg Message) Clone() Message {
	// Message shouldn't require cloning currency, but something is mutating that variable elsewhere in the code
	return Message{
		msg.Data.Clone(),
		msg.TokenType,
		new(big.Int).Set(msg.Currency),
		msg.Destination,
	}
}

func (msg Message) AsValue() value.Value {
	destinationBytes := [32]byte{}
	copy(destinationBytes[12:], msg.Destination[:])
	destination := big.NewInt(0).SetBytes(destinationBytes[:])
	tokTypeBytes := [32]byte{}
	copy(tokTypeBytes[11:], msg.TokenType[:])
	tokTypeInt := big.NewInt(0).SetBytes(tokTypeBytes[:])
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		msg.Data,
		value.NewIntValue(destination),
		value.NewIntValue(msg.Currency),
		value.NewIntValue(tokTypeInt),
	})
	return newTup
}

func (msg Message) Equals(b Message) bool {
	if msg.TokenType != b.TokenType {
		return false
	}
	if !value.Eq(msg.Data, b.Data) {
		return false
	}
	if msg.Currency.Cmp(b.Currency) != 0 {
		return false
	}
	if msg.Destination != b.Destination {
		return false
	}
	return true
}
