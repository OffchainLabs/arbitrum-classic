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

package structures

import (
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type messageStackItem struct {
	message value.Value
	prev    *messageStackItem
	next    *messageStackItem
	hash    [32]byte
	count   *big.Int
}

func (pii *messageStackItem) skipNext(n uint64) *messageStackItem {
	ret := pii
	for i := uint64(0); i < n && ret != nil; i++ {
		ret = ret.next
	}
	return ret
}

func (msi *messageStackItem) Equals(msi2 *messageStackItem) bool {
	return msi.hash == msi2.hash &&
		msi.count.Cmp(msi2.count) == 0 &&
		value.Eq(msi.message, msi2.message) &&
		(msi.prev == nil) == (msi2.prev == nil) &&
		(msi.next == nil) == (msi2.next == nil)
}

type MessageStack struct {
	newest     *messageStackItem
	oldest     *messageStackItem
	index      map[[32]byte]*messageStackItem
	hashOfRest [32]byte
}

func NewMessageStack() *MessageStack {
	return &MessageStack{
		newest:     nil,
		oldest:     nil,
		index:      make(map[[32]byte]*messageStackItem),
		hashOfRest: value.NewEmptyTuple().Hash(),
	}
}

func (ms *MessageStack) GetTopHash() [32]byte {
	if ms.newest == nil {
		return value.NewEmptyTuple().Hash()
	} else {
		return ms.newest.hash
	}
}

func (ms *MessageStack) TopCount() *big.Int {
	if ms.newest == nil {
		return big.NewInt(0)
	} else {
		return ms.newest.count
	}
}

func (ms *MessageStack) BottomCount() *big.Int {
	if ms.oldest == nil {
		return big.NewInt(0)
	} else {
		return ms.oldest.count
	}
}

func (pi *MessageStack) DeliverMessage(msg value.Value) {
	newTopCount := new(big.Int).Add(pi.TopCount(), big.NewInt(1))
	if pi.newest == nil {
		item := &messageStackItem{
			message: msg,
			prev:    nil,
			next:    nil,
			hash:    hash2(pi.hashOfRest, msg.Hash()),
			count:   newTopCount,
		}
		pi.newest = item
		pi.oldest = item
		pi.index[item.hash] = item
	} else {
		item := &messageStackItem{
			message: msg,
			prev:    pi.newest,
			next:    nil,
			hash:    hash2(pi.newest.hash, msg.Hash()),
			count:   newTopCount,
		}
		pi.newest = item
		item.prev.next = item
		pi.index[item.hash] = item
	}
}

func (pi *MessageStack) GetHeight(acc [32]byte) (*big.Int, bool) {
	item, ok := pi.index[acc]
	if !ok {
		if pi.hashOfRest == acc {
			if pi.BottomCount().Cmp(big.NewInt(0)) == 0 {
				return big.NewInt(0), true
			} else {
				return new(big.Int).Sub(pi.BottomCount(), big.NewInt(1)), true
			}
		} else {
			return nil, false
		}
	}
	return item.count, true
}

func (pi *MessageStack) SegmentSize(olderAcc [32]byte, newerAcc [32]byte) (uint64, error) {
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

func hash2(h1, h2 [32]byte) [32]byte {
	return value.NewTuple2(
		value.NewHashOnlyValue(h1, 1),
		value.NewHashOnlyValue(h2, 1),
	).Hash()
}

func (pi *MessageStack) MarshalForCheckpoint(ctx CheckpointContext) *PendingInboxBuf {
	var msgHashes []*value.HashBuf
	for item := pi.newest; item != nil; item = item.prev {
		ctx.AddValue(item.message)
		msgHashes = append(msgHashes, utils.MarshalHash(item.message.Hash()))
	}
	var topCount *big.Int
	if pi.newest == nil {
		topCount = big.NewInt(0)
	} else {
		topCount = pi.newest.count
	}
	return &PendingInboxBuf{
		TopCount:   utils.MarshalBigInt(topCount),
		ItemHashes: msgHashes,
		HashOfRest: utils.MarshalHash(pi.hashOfRest),
	}
}

func (buf *PendingInboxBuf) UnmarshalFromCheckpoint(ctx RestoreContext) *MessageStack {
	ret := NewMessageStack()
	ret.hashOfRest = utils.UnmarshalHash(buf.HashOfRest)
	for i := len(buf.ItemHashes) - 1; i >= 0; i = i - 1 {
		val := ctx.GetValue(utils.UnmarshalHash(buf.ItemHashes[i]))
		ret.DeliverMessage(val)
	}
	return ret
}

func (ms *MessageStack) Equals(ms2 *MessageStack) bool {
	if ms.hashOfRest != ms2.hashOfRest || len(ms.index) != len(ms2.index) {
		return false
	}
	for h, m := range ms.index {
		m2 := ms2.index[h]
		if m2 == nil || !m.Equals(m2) {
			return false
		}
	}
	return true
}

func (pi *MessageStack) GenerateBisection(startItemHash [32]byte, endItemHash [32]byte, segments uint64) ([][32]byte, error) {
	startItem, ok := pi.index[startItemHash]
	if !ok {
		return nil, errors.New("startItemHash not found")
	}

	endItem, ok := pi.index[endItemHash]
	if !ok {
		return nil, errors.New("endItemHash not found")
	}

	count := new(big.Int).Sub(pi.newest.count, endItem.count).Uint64()
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

func (pi *MessageStack) CheckBisection(segments [][32]byte) (uint64, error) {
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

func (pi *MessageStack) GenerateOneStepProof(startItemHash [32]byte) ([32]byte, [32]byte, error) {
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

func (pi *MessageStack) Substack(olderAcc, newerAcc [32]byte) (*MessageStack, error) {
	if olderAcc == newerAcc {
		return NewMessageStack(), nil
	}
	var oldItem *messageStackItem
	if olderAcc == pi.hashOfRest {
		oldItem = pi.oldest
	} else {
		old, ok := pi.index[olderAcc]
		if !ok {
			return nil, errors.New("olderAcc not found")
		}
		oldItem = old.next
	}

	newItem, ok := pi.index[newerAcc]
	if !ok {
		return nil, errors.New("newerAcc not found")
	}

	if oldItem.count.Cmp(newItem.count) > 0 {
		return nil, errors.New("olderAcc is not before newerAcc")
	}
	item := oldItem
	stack := NewMessageStack()
	for item != newItem {
		stack.DeliverMessage(item.message)
		item = item.next
	}
	stack.DeliverMessage(item.message)
	return stack, nil
}

func (pi *MessageStack) ValueForSubseq(olderAcc, newerAcc [32]byte) value.TupleValue {
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

func valueForSubseq2(oldItem, newItem *messageStackItem) value.TupleValue {
	if newItem == oldItem {
		return value.NewEmptyTuple()
	} else {
		return value.NewTuple2(
			valueForSubseq2(oldItem, newItem.prev),
			newItem.message,
		)
	}
}

type PendingInbox struct {
	*MessageStack
}

func NewPendingInbox() *PendingInbox {
	return &PendingInbox{
		MessageStack: NewMessageStack(),
	}
}

func (pi *PendingInbox) DiscardUpToCount(count *big.Int) {
	for pi.oldest != nil && pi.oldest.count.Cmp(count) < 0 {
		victim := pi.oldest
		if victim == pi.newest {
			pi.oldest = nil
			pi.newest = nil
		} else {
			pi.oldest = victim.next
			pi.oldest.prev = nil
		}
		delete(pi.index, victim.hash)
	}
}

func (pi *PendingInbox) Equals(pi2 *PendingInbox) bool {
	return pi.MessageStack.Equals(pi2.MessageStack)
}
