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
	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

//go:generate protoc -I.. -I. --go_out=paths=source_relative:. protocol.proto

func NewTokenTypeBuf(tok [21]byte) *TokenTypeBuf {
	return &TokenTypeBuf{
		Value: tok[:],
	}
}

func NewTokenTypeFromBuf(buf *TokenTypeBuf) [21]byte {
	var ret [21]byte
	copy(ret[:], buf.Value)
	return ret
}

func NewAddressBuf(tok common.Address) *AddressBuf {
	return &AddressBuf{
		Value: tok.Bytes(),
	}
}

func NewAddressFromBuf(buf *AddressBuf) common.Address {
	var ret common.Address
	copy(ret[:], buf.Value)
	return ret
}

func NewMessageBuf(val Message) *MessageBuf {
	return &MessageBuf{
		Value:     value.NewValueBuf(val.Data),
		TokenType: NewTokenTypeBuf(val.TokenType),
		Amount:    value.NewBigIntBuf(val.Currency),
		Sender:    NewAddressBuf(val.Destination),
	}
}

func NewMessageFromBuf(buf *MessageBuf) (Message, error) {
	val, err := value.NewValueFromBuf(buf.Value)
	return NewSimpleMessage(
		val,
		NewTokenTypeFromBuf(buf.TokenType),
		value.NewBigIntFromBuf(buf.Amount),
		NewAddressFromBuf(buf.Sender),
	), err
}

func NewAssertionBuf(a *Assertion) *AssertionBuf {
	messages := make([]*MessageBuf, 0, len(a.OutMsgs))
	for _, msg := range a.OutMsgs {
		messages = append(messages, NewMessageBuf(msg))
	}
	return &AssertionBuf{
		AfterHash: value.NewHashBuf(a.AfterHash),
		NumSteps:  a.NumSteps,
		Messages:  messages,
	}
}

func NewAssertionFromBuf(buf *AssertionBuf) (*Assertion, error) {
	messages := make([]Message, 0, len(buf.Messages))
	for _, msg := range buf.Messages {
		m, err := NewMessageFromBuf(msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	logs := make([]value.Value, 0, len(buf.Logs))
	for _, valLog := range buf.Logs {
		v, err := value.NewValueFromBuf(valLog)
		if err != nil {
			return nil, err
		}
		logs = append(logs, v)
	}
	return NewAssertion(
		value.NewHashFromBuf(buf.AfterHash),
		buf.NumSteps,
		messages,
		logs,
	), nil
}
