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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type OutMessage struct {
	Kind   inbox.Type
	Sender common.Address
	Data   []byte
}

func NewOutMessage(msg Message, sender common.Address) OutMessage {
	return OutMessage{
		Kind:   msg.Type(),
		Sender: sender,
		Data:   msg.AsData(),
	}
}

func NewOutMessageFromValue(val value.Value) (OutMessage, error) {
	failRet := OutMessage{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("val must be a tuple")
	}
	if tup.Len() != 3 {
		return failRet, errors.Errorf("expected tuple of length 3, but recieved %v", tup)
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	sender, _ := tup.GetByInt64(1)
	messageData, _ := tup.GetByInt64(2)

	kindInt, ok := kind.(value.IntValue)
	if !ok {
		return failRet, errors.New("outgoing message kind must be an int")
	}
	senderInt, ok := sender.(value.IntValue)
	if !ok {
		return failRet, errors.New("sender must be an int")
	}
	data, err := inbox.ByteStackToHex(messageData)
	if err != nil {
		return failRet, err
	}

	return OutMessage{
		Kind:   inbox.Type(kindInt.BigInt().Uint64()),
		Sender: inbox.NewAddressFromInt(senderInt),
		Data:   data,
	}, nil
}

func NewRandomOutMessage(msg Message) OutMessage {
	return NewOutMessage(msg, common.RandAddress())
}

func (im OutMessage) AsValue() value.Value {
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(im.Kind)),
		inbox.NewIntFromAddress(im.Sender),
		inbox.BytesToByteStack(im.Data),
	})
	return tup
}
