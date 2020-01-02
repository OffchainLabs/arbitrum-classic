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
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func AddMessageAcc(currentInbox value.Value, msgVal value.Value) value.TupleValue {
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(0),
		currentInbox,
		msgVal,
	})
	return tup
}

func AddMessage(currentInbox value.Value, msgVal value.Value) value.TupleValue {
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		currentInbox,
		msgVal,
	})
	return tup
}

type MessageStack struct {
	msg value.TupleValue
}

func NewMessageStack() *MessageStack {
	return &MessageStack{value.NewEmptyTuple()}
}

func (in *MessageStack) Clone() *MessageStack {
	return &MessageStack{in.msg}
}

func (in *MessageStack) String() string {
	return fmt.Sprintf("MessageStack(%v)", in.msg)
}

func (in *MessageStack) GetValue() value.TupleValue {
	return in.msg
}

func (in *MessageStack) IsEmpty() bool {
	return in.msg.Equal(value.NewEmptyTuple())
}

func (in *MessageStack) AddMessage(msgVal value.Value) {
	in.msg, _ = value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(0),
		in.msg,
		msgVal,
	})
}

type Inbox struct {
	messages value.TupleValue
}

func NewInbox() *Inbox {
	return &Inbox{value.NewEmptyTuple()}
}

func (in *Inbox) Clone() *Inbox {
	return &Inbox{in.messages}
}

func (in *Inbox) String() string {
	return fmt.Sprintf("Inbox(%v)", in.messages)
}

func (in *Inbox) WithAddedMessages(messages value.TupleValue) *Inbox {
	if messages.Len() == 0 {
		return in
	}
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		in.messages,
		messages,
	})
	return &Inbox{tup}
}

func (in *Inbox) Receive() value.TupleValue {
	return in.messages
}
func (in *Inbox) EmptyAccepted() {
	in.Accepted = NewMessageQueues()
}
