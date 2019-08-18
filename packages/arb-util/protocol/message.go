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
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Identity [32]byte
type TokenType [21]byte

func TokenTypeFromIntValue(val value.IntValue) TokenType {
	var tokType TokenType
	tokBytes := val.ToBytes()
	copy(tokType[:], tokBytes[:])
	return tokType
}

func (t TokenType) ToIntValue() value.IntValue {
	var bigtok [32]byte
	copy(bigtok[:], t[:])
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
	Destination [32]byte
}

func NewMessage(data value.Value, tokenType [21]byte, currency *big.Int, destination [32]byte) Message {
	return Message{data, tokenType, currency, destination}
}

func NewSimpleMessage(data value.Value, tokenType [21]byte, currency *big.Int, sender common.Address) Message {
	senderArr := [32]byte{}
	copy(senderArr[:], sender.Bytes())
	return Message{data, tokenType, currency, senderArr}
}

func NewMessageFromReader(rd io.Reader) (Message, error) {
	data, err := value.UnmarshalValue(rd)
	if err != nil {
		return Message{}, err
	}

	tokenType := [21]byte{}
	_, err = rd.Read(tokenType[:])
	if err != nil {
		return Message{}, fmt.Errorf("error unmarshalling OutgoingMessage: %v", err)
	}

	currency, err := value.NewIntValueFromReader(rd)
	if err != nil {
		return Message{}, fmt.Errorf("error unmarshalling OutgoingMessage: %v", err)
	}

	dest := [32]byte{}
	_, err = rd.Read(tokenType[:])
	if err != nil {
		return Message{}, fmt.Errorf("error unmarshalling OutgoingMessage: %v", err)
	}

	return NewMessage(data, tokenType, currency.BigInt(), dest), nil
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
		destInt.ToBytes(),
	), nil
}

func (msg Message) Marshal(w io.Writer) error {
	if err := value.MarshalValue(msg.Data, w); err != nil {
		return err
	}

	_, err := w.Write(msg.TokenType[:])
	if err != nil {
		return err
	}

	err = value.NewIntValue(msg.Currency).Marshal(w)
	if err != nil {
		return err
	}
	_, err = w.Write(msg.Destination[:])
	if err != nil {
		return err
	}
	return nil
}

func (msg Message) Hash() [32]byte {
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(msg.Data.Hash),
		tokenTypeEncoded(msg.TokenType),
		solsha3.Uint256(msg.Currency),
		solsha3.Bytes32(msg.Destination),
	)
	copy(ret[:], hashVal)
	return ret
}

func (msg Message) Clone() Message {
	return Message{
		msg.Data.Clone(),
		msg.TokenType,
		msg.Currency,
		msg.Destination,
	}
}

func (msg Message) AsValue() value.Value {
	destination := big.NewInt(0)
	destination.SetBytes(msg.Destination[:])
	tokTypeBytes := [32]byte{}
	copy(tokTypeBytes[:], msg.TokenType[:])
	tokTypeInt := big.NewInt(0)
	tokTypeInt.SetBytes(tokTypeBytes[:])
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
