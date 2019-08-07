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

type MessageQueue struct {
	msg      value.TupleValue
	msgCount uint64
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{value.NewEmptyTuple(), 0}
}

func (in *MessageQueue) Clone() *MessageQueue {
	return &MessageQueue{in.msg, in.msgCount}
}

func (in *MessageQueue) String() string {
	return fmt.Sprintf("MessageQueue(%v)", in.msg)
}

func (in *MessageQueue) MessageCount() uint64 {
	return in.msgCount
}

func (in *MessageQueue) IsEmpty() bool {
	return in.msg.Equal(value.NewEmptyTuple())
}

func (in *MessageQueue) AddMessage(msg Message) {
	in.AddRawMessage(msg.AsValue())
}

func (in *MessageQueue) AddRawMessage(msgVal value.Value) {
	in.msg, _ = value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(0),
		in.msg,
		msgVal,
	})
	in.msgCount++
}

type MessageQueues struct {
	queues value.TupleValue
}

func NewMessageQueues() *MessageQueues {
	return &MessageQueues{value.NewEmptyTuple()}
}

func (in *MessageQueues) Clone() *MessageQueues {
	return &MessageQueues{in.queues}
}

func (in *MessageQueues) String() string {
	return fmt.Sprintf("MessageQueues(%v)", in.queues)
}

func (in *MessageQueues) WithAddedQueue(queue *MessageQueue) *MessageQueues {
	if queue.IsEmpty() {
		return in
	}
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		in.queues,
		queue.msg,
	})
	return &MessageQueues{tup}
}

type Inbox struct {
	Accepted     *MessageQueues
	PendingQueue *MessageQueue
}

func NewInbox(inbox *MessageQueues, pending *MessageQueue) *Inbox {
	return &Inbox{inbox, pending}
}

func NewEmptyInbox() *Inbox {
	return &Inbox{NewMessageQueues(), NewMessageQueue()}
}

func (in *Inbox) Clone() *Inbox {
	return &Inbox{in.Accepted, in.PendingQueue}
}

func (in *Inbox) SendMessage(msg Message) {
	in.PendingQueue.AddMessage(msg)
}

func (in *Inbox) SendRawMessage(msg value.Value) {
	in.PendingQueue.AddRawMessage(msg)
}

func (in *Inbox) InsertMessageGroup(msgs []Message) {
	if len(msgs) > 0 {
		q := NewMessageQueue()
		for _, msg := range msgs {
			q.AddMessage(msg)
		}
		in.InsertMessageQueue(q)
	}
}

func (in *Inbox) InsertMessageQueue(mq *MessageQueue) {
	in.Accepted = in.Accepted.WithAddedQueue(mq)
}

func (in *Inbox) DeliverMessages() {
	in.InsertMessageQueue(in.PendingQueue)
	in.PendingQueue = NewMessageQueue()
}

func (in *Inbox) Receive() value.TupleValue {
	return in.Accepted.queues
}

func (in *Inbox) ReceivePending() value.TupleValue {
	return in.Accepted.WithAddedQueue(in.PendingQueue).queues
}

func (in *Inbox) Pending() value.TupleValue {
	return in.PendingQueue.msg
}
