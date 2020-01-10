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
	in.msg = value.NewTuple2(in.msg, msgVal)
}
