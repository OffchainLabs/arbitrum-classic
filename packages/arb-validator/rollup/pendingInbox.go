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
	"errors"
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type pendingInboxItem struct {
	message value.Value
	prev    *pendingInboxItem
	next    *pendingInboxItem
	hash    [32]byte
	count   *big.Int
}

func (pii *pendingInboxItem) skipNext(n uint64) *pendingInboxItem {
	ret := pii
	for i := uint64(0); i < n && ret != nil; i++ {
		ret = ret.next
	}
	return ret
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

func (pi *PendingInbox) SegmentSize(olderAcc [32]byte, newerAcc [32]byte) (uint64, error) {
	oldItem, ok := pi.index[olderAcc]
	if !ok {
		return 0, errors.New("olderAcc not found")
	}
	newItem, ok := pi.index[newerAcc]
	if !ok {
		return 0, errors.New("newerAcc not found")
	}
	return new(big.Int).Sub(newItem.count, oldItem.count).Uint64(), nil
}

func (pi *PendingInbox) GenerateBisection(startItemHash [32]byte, endItemHash [32]byte, segments uint64) ([][32]byte, error) {
	startItem, ok := pi.index[startItemHash]
	if !ok {
		return nil, errors.New("startItemHash not found")
	}

	endItem, ok := pi.index[endItemHash]
	if !ok {
		return nil, errors.New("endItemHash not found")
	}

	count := new(big.Int).Sub(pi.head.count, endItem.count).Uint64()
	if count < segments {
		segments = count
	}

	cuts := make([][32]byte, 0, segments+1)
	cuts = append(cuts, startItemHash)
	firstSegmentSize := count/segments + count%segments
	otherSegmentSize := count / segments
	item := startItem.skipNext(firstSegmentSize)
	if item == nil {
		return nil, errors.New("pending inbox too short")
	}
	cuts = append(cuts, item.hash)
	for i := uint64(0); i < segments; i++ {
		item = item.skipNext(otherSegmentSize)
		if item == nil {
			return nil, errors.New("pending inbox too short")
		}
		cuts = append(cuts, item.hash)
	}
	return cuts, nil
}

func (pi *PendingInbox) CheckBisection(segments [][32]byte) (uint64, error) {
	segmentCount := uint64(len(segments))
	totalLength, err := pi.SegmentSize(segments[0], segments[segmentCount-1])
	if err != nil {
		return 0, errors.New("first and last segment must exist")
	}
	for i := uint64(1); i < segmentCount; i++ {
		targetLength := uint64(0)
		if i == segmentCount-1 {
			targetLength = totalLength/segmentCount + totalLength%segmentCount
		} else {
			targetLength = totalLength / segmentCount
		}
		length, err := pi.SegmentSize(segments[segmentCount-i-1], segments[segmentCount-i])
		if err != nil || length != targetLength {
			return uint64(segmentCount - i - 1), nil
		}
	}
	return 0, errors.New("all segments were correct")
}

func (pi *PendingInbox) GenerateOneStepProof(startItemHash [32]byte) ([32]byte, [32]byte, error) {
	item, ok := pi.index[startItemHash]
	if !ok {
		return [32]byte{}, [32]byte{}, errors.New("startItemHash not found")
	}
	next := item.next
	if next == nil {
		return [32]byte{}, [32]byte{}, errors.New("startItemHash is last item")
	}
	return next.hash, next.message.Hash(), nil
}

func (pi *PendingInbox) DeliverMessage(msg value.Value) {
	if pi.head == nil {
		item := &pendingInboxItem{
			message: msg,
			prev:    nil,
			next:    nil,
			hash:    hash2(pi.hashOfRest, msg.Hash()),
			count:   big.NewInt(1),
		}
		pi.head = item
		pi.index[item.hash] = item
	} else {
		item := &pendingInboxItem{
			message: msg,
			prev:    pi.head,
			next:    nil,
			hash:    hash2(pi.head.hash, msg.Hash()),
			count:   new(big.Int).Add(pi.head.count, big.NewInt(1)),
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
