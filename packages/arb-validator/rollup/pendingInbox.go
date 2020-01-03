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

package rollup

import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"log"
)

type pendingInboxItem struct {
	message value.Value
	prev    *pendingInboxItem
	next    *pendingInboxItem
	hash    [32]byte
}

type PendingInbox struct {
	head       *pendingInboxItem
	index      map[[32]byte]*pendingInboxItem
	hashOfRest [32]byte
}

func NewPendingInbox() *PendingInbox {
	return &PendingInbox{
		head:       nil,
		index:      make(map[[32]byte]*pendingInboxItem),
		hashOfRest: value.NewEmptyTuple().Hash(),
	}
}

func (pi *PendingInbox) DeliverMessage(msg value.Value) {
	if pi.head == nil {
		item := &pendingInboxItem{
			message: msg,
			prev:    nil,
			next:    nil,
			hash:    hash2(pi.hashOfRest, msg.Hash()),
		}
		pi.head = item
		pi.index[item.hash] = item
	} else {
		item := &pendingInboxItem{
			message: msg,
			prev:    pi.head,
			next:    nil,
			hash:    hash2(pi.head.hash, msg.Hash()),
		}
		pi.head = item
		item.prev.next = item
		pi.index[item.hash] = item
	}
}

func (pi *PendingInbox) DiscardUpTo(hash [32]byte) (discardedSomething bool) {
	item, ok := pi.index[hash]
	if !ok {
		return false
	}
	pi.hashOfRest = item.hash
	if item.next != nil {
		item.next.prev = nil
	}
	pi.discardItems(item)
	return true
}

func (pi *PendingInbox) discardItems(item *pendingInboxItem) {
	for item != nil {
		delete(pi.index, item.hash)
		item = item.prev
	}
}

func (pi *PendingInbox) ValueForSubseq(olderAcc, newerAcc [32]byte) value.Value {
	oldItem, ok := pi.index[olderAcc]
	if !ok {
		oldItem = nil
	}
	newItem, ok := pi.index[newerAcc]
	if !ok {
		newItem = nil
	}
	return valueForSubseq2(oldItem, newItem)
}

func valueForSubseq2(oldItem, newItem *pendingInboxItem) value.Value {
	if newItem == oldItem {
		return value.NewEmptyTuple()
	} else {
		return value.NewTuple2(
			valueForSubseq2(oldItem, newItem.prev),
			newItem.message,
		)
	}
}

func hash2(h1, h2 [32]byte) [32]byte {
	return value.NewTuple2(
		value.NewHashOnlyValue(h1, 1),
		value.NewHashOnlyValue(h2, 1),
	).Hash()
}

func (pi *PendingInbox) MarshalToBuf() *PendingInboxBuf {
	var msgs [][]byte
	for item := pi.head; item != nil; item = item.prev {
		bb := bytes.NewBuffer(nil)
		err := value.MarshalValue(item.message, bb)
		if err != nil {
			log.Fatal(err)
		}
		msgs = append(msgs, bb.Bytes())
	}
	return &PendingInboxBuf{
		Items:      msgs,
		HashOfRest: marshalHash(pi.hashOfRest),
	}
}

func (buf *PendingInboxBuf) Unmarshal() *PendingInbox {
	ret := NewPendingInbox()
	ret.hashOfRest = unmarshalHash(buf.HashOfRest)
	for i := len(buf.Items) - 1; i >= 0; i = i - 1 {
		val, err := value.UnmarshalValue(bytes.NewBuffer([]byte(buf.Items[i])))
		if err != nil {
			log.Fatal(err)
		}
		ret.DeliverMessage(val)
	}
	return ret
}
