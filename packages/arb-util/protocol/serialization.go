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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

//go:generate protoc -I.. -I. --go_out=paths=source_relative:. protocol.proto

func NewTimeBoundsBuf(tb TimeBounds) *TimeBoundsBuf {
	return &TimeBoundsBuf{
		StartTime: tb[0],
		EndTime:   tb[1],
	}
}

func NewTimeBoundsFromBuf(buf *TimeBoundsBuf) TimeBounds {
	return TimeBounds{buf.StartTime, buf.EndTime}
}

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

func NewBalanceTrackerBuf(bt *BalanceTracker) *BalanceTrackerBuf {
	types := make([]*TokenTypeBuf, 0, len(bt.TokenTypes))
	amounts := make([]*value.BigIntegerBuf, 0, len(bt.TokenAmounts))
	for _, tokenType := range bt.TokenTypes {
		types = append(types, &TokenTypeBuf{
			Value: tokenType[:],
		})
	}
	for _, tokenAmount := range bt.TokenAmounts {
		amounts = append(amounts, &value.BigIntegerBuf{
			Value: tokenAmount.Bytes(),
		})
	}
	return &BalanceTrackerBuf{
		Types:   types,
		Amounts: amounts,
	}
}

func NewBalanceTrackerFromBuf(buf *BalanceTrackerBuf) *BalanceTracker {
	types := make([][21]byte, 0, len(buf.Types))
	amounts := make([]*big.Int, 0, len(buf.Amounts))

	for _, tokenType := range buf.Types {
		var typ [21]byte
		copy(typ[:], tokenType.Value)
		types = append(types, typ)
	}
	for _, tokenAmount := range buf.Amounts {
		amounts = append(amounts, value.NewBigIntFromBuf(tokenAmount))
	}
	return NewBalanceTrackerFromLists(types, amounts)
}

func NewPreconditionBuf(pre *Precondition) *PreconditionBuf {
	return &PreconditionBuf{
		BeforeHash:     value.NewHashBuf(pre.BeforeHash),
		TimeBounds:     NewTimeBoundsBuf(pre.TimeBounds),
		BalanceTracker: NewBalanceTrackerBuf(pre.BeforeBalance),
		BeforeInbox:    value.NewHashBuf(pre.BeforeInbox.Hash()),
	}
}

func NewPreconditionFromBuf(buf *PreconditionBuf) *Precondition {
	return &Precondition{
		value.NewHashFromBuf(buf.BeforeHash),
		NewTimeBoundsFromBuf(buf.TimeBounds),
		NewBalanceTrackerFromBuf(buf.BalanceTracker),
		value.NewHashOnlyValue(value.NewHashFromBuf(buf.BeforeInbox), 1),
	}
}

func NewMessageBuf(val Message) *MessageBuf {
	return &MessageBuf{
		Value:     value.NewValueBuf(val.Data),
		TokenType: NewTokenTypeBuf(val.TokenType),
		Amount:    value.NewBigIntBuf(val.Currency),
		Sender:    value.NewHashBuf(val.Destination),
	}
}

func NewMessageFromBuf(buf *MessageBuf) (Message, error) {
	val, err := value.NewValueFromBuf(buf.Value)
	return NewMessage(
		val,
		NewTokenTypeFromBuf(buf.TokenType),
		value.NewBigIntFromBuf(buf.Amount),
		value.NewHashFromBuf(buf.Sender),
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
