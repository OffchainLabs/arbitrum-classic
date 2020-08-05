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
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type messageStackItem struct {
	message inbox.InboxMessage
	prev    *messageStackItem
	next    *messageStackItem
	hash    common.Hash
	count   *big.Int
}

func (msi *messageStackItem) skipNext(n uint64) *messageStackItem {
	ret := msi
	for i := uint64(0); i < n && ret != nil; i++ {
		ret = ret.next
	}
	return ret
}

func (msi *messageStackItem) skipPrev(n uint64) *messageStackItem {
	ret := msi
	for i := uint64(0); i < n && ret != nil; i++ {
		ret = ret.prev
	}
	return ret
}

func (msi *messageStackItem) Equals(msi2 *messageStackItem) bool {
	return msi.hash == msi2.hash &&
		msi.count.Cmp(msi2.count) == 0 &&
		msi.message.Equals(msi2.message) &&
		(msi.prev == nil) == (msi2.prev == nil) &&
		(msi.next == nil) == (msi2.next == nil)
}

type MessageStack struct {
	newest     *messageStackItem
	oldest     *messageStackItem
	index      map[common.Hash]*messageStackItem
	hashOfRest common.Hash
}

func NewMessageStack() *MessageStack {
	return &MessageStack{
		newest: nil,
		oldest: nil,
		index:  make(map[common.Hash]*messageStackItem),
	}
}

func NewRandomMessageStack(count int) *MessageStack {
	ms := NewMessageStack()
	for i := 0; i < count; i++ {
		ms.DeliverMessage(inbox.NewRandomInboxMessage())
	}
	return ms
}

func (ms *MessageStack) String() string {
	hashes := make([]common.Hash, 0)
	hashes = append(hashes, ms.hashOfRest)
	for item := ms.oldest; item != nil; item = item.next {
		hashes = append(hashes, item.hash)
	}
	return fmt.Sprintf("%v", hashes)
}

func (ms *MessageStack) GetTopHash() common.Hash {
	if ms.newest == nil {
		return common.Hash{}
	} else {
		return ms.newest.hash
	}
}

func (ms *MessageStack) GetMaxAtHeight(maxHeight *common.TimeBlocks) (bool, *big.Int) {
	msg := ms.newest
	for msg != nil && msg.message.ChainTime.BlockNum.Cmp(maxHeight) > 0 {
		msg = msg.prev
	}
	if msg == nil {
		return false, nil
	}
	return true, msg.count
}

func (ms *MessageStack) TopCount() *big.Int {
	if ms.newest == nil {
		return big.NewInt(0)
	} else {
		return ms.newest.count
	}
}

func (ms *MessageStack) bottomIndex() *big.Int {
	if ms.oldest == nil {
		return big.NewInt(0)
	} else {
		return ms.oldest.count
	}
}

func (ms *MessageStack) DeliverMessage(msg inbox.InboxMessage) {
	newTopCount := new(big.Int).Add(ms.TopCount(), big.NewInt(1))
	if ms.newest == nil {
		item := &messageStackItem{
			message: msg,
			prev:    nil,
			next:    nil,
			hash:    hash2(ms.hashOfRest, msg.CommitmentHash()),
			count:   newTopCount,
		}
		ms.newest = item
		ms.oldest = item
		ms.index[item.hash] = item
	} else {
		item := &messageStackItem{
			message: msg,
			prev:    ms.newest,
			next:    nil,
			hash:    hash2(ms.newest.hash, msg.CommitmentHash()),
			count:   newTopCount,
		}
		ms.newest = item
		item.prev.next = item
		ms.index[item.hash] = item
	}
}

func (ms *MessageStack) GetHashAtIndex(height *big.Int) (common.Hash, error) {
	if height.Cmp(big.NewInt(0)) == 0 {
		return common.Hash{}, nil
	}
	if height.Cmp(ms.bottomIndex()) < 0 {
		return common.Hash{}, errors.New("height is below bottom of l2message stack")
	}
	if height.Cmp(ms.TopCount()) > 0 {
		return common.Hash{}, errors.New("height is above top of l2message stack")
	}
	offset := new(big.Int).Sub(height, ms.bottomIndex())
	return ms.oldest.skipNext(offset.Uint64()).hash, nil
}

func hash2(h1, h2 common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(h1), hashing.Bytes32(h2))
}

func (ms *MessageStack) MarshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *InboxBuf {
	var items []*common.HashBuf
	for item := ms.newest; item != nil; item = item.prev {
		checkpointVal := item.message.AsValue()
		ctx.AddValue(checkpointVal)
		items = append(items, checkpointVal.Hash().MarshalToBuf())
	}
	var topCount *big.Int
	if ms.newest == nil {
		topCount = big.NewInt(0)
	} else {
		topCount = ms.newest.count
	}
	return &InboxBuf{
		TopCount:   common.MarshalBigInt(topCount),
		Items:      items,
		HashOfRest: ms.hashOfRest.MarshalToBuf(),
	}
}

func (x *InboxBuf) UnmarshalFromCheckpoint(ctx ckptcontext.RestoreContext) (*MessageStack, error) {
	ret := NewMessageStack()
	ret.hashOfRest = x.HashOfRest.Unmarshal()
	for i := len(x.Items) - 1; i >= 0; i = i - 1 {
		val := ctx.GetValue(x.Items[i].Unmarshal())
		msg, err := inbox.NewInboxMessageFromValue(val)
		if err != nil {
			return nil, err
		}
		ret.DeliverMessage(msg)
	}
	return ret, nil
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

func (ms *MessageStack) itemAfterHash(acc common.Hash) (*messageStackItem, bool) {
	if acc == ms.hashOfRest {
		return ms.oldest, true
	}
	item, found := ms.index[acc]
	if !found {
		return nil, false
	}
	return item.next, true
}

func (ms *MessageStack) itemSkippedAfterHash(acc common.Hash, count uint64) (common.Hash, bool) {
	if count == 0 {
		return acc, true
	}
	next, ok := ms.itemAfterHash(acc)
	if !ok {
		return common.Hash{}, false
	}
	node := next.skipNext(count - 1)
	if node == nil {
		return common.Hash{}, false
	}
	return node.hash, true
}

func (ms *MessageStack) itemAtHash(acc common.Hash) (*messageStackItem, bool) {
	item, found := ms.index[acc]
	return item, found
}

func segmentSizes(segments, count uint64) (uint64, uint64, uint64) {
	if count < segments {
		segments = count
	}
	firstSegmentSize := count/segments + count%segments
	otherSegmentSize := count / segments
	return segments, firstSegmentSize, otherSegmentSize
}

func (ms *MessageStack) GenerateBisection(startItemHash common.Hash, segments, count uint64) ([]common.Hash, error) {
	startItem, ok := ms.itemAfterHash(startItemHash)
	if !ok {
		return nil, errors.New("bisection startItemHash not found")
	}

	segments, firstSegmentSize, otherSegmentSize := segmentSizes(segments, count)

	cuts := make([]common.Hash, 0, segments+1)
	cuts = append(cuts, startItemHash)
	item := startItem.skipNext(firstSegmentSize - 1)
	if item == nil {
		return nil, errors.New("inbox too short start")
	}
	cuts = append(cuts, item.hash)
	for i := uint64(1); i < segments; i++ {
		item = item.skipNext(otherSegmentSize)
		if item == nil {
			return nil, errors.New("inbox too short rest")
		}
		cuts = append(cuts, item.hash)
	}
	return cuts, nil
}

func (ms *MessageStack) GenerateBisectionReverse(afterItemHash common.Hash, segments, count uint64) ([]common.Hash, error) {
	startItem, ok := ms.itemAtHash(afterItemHash)
	if !ok {
		return nil, errors.New("bisection startItemHash not found")
	}

	segments, firstSegmentSize, otherSegmentSize := segmentSizes(segments, count)

	cuts := make([]common.Hash, 0, segments+1)
	cuts = append(cuts, afterItemHash)
	item := startItem.skipPrev(firstSegmentSize)
	if item == nil {
		return nil, errors.New("inbox too short start")
	}
	cuts = append(cuts, item.hash)
	for i := uint64(1); i < segments; i++ {
		item = item.skipPrev(otherSegmentSize)
		if item == nil {
			return nil, errors.New("inbox too short rest")
		}
		cuts = append(cuts, item.hash)
	}
	return cuts, nil
}

func (ms *MessageStack) InboxMessageAt(afterGlobalInbox common.Hash) (inbox.InboxMessage, error) {
	item, ok := ms.itemAtHash(afterGlobalInbox)
	if !ok {
		return inbox.InboxMessage{}, errors.New("one step proof startItemHash not found")
	}
	return item.message, nil
}

func (ms *MessageStack) InboxMessageBefore(afterGlobalInbox common.Hash) (inbox.InboxMessage, error) {
	item, ok := ms.itemAtHash(afterGlobalInbox)
	if !ok {
		return inbox.InboxMessage{}, errors.New("one step proof startItemHash not found")
	}
	if item.prev == nil {
		return inbox.InboxMessage{}, errors.New("no previous message")
	}
	return item.prev.message, nil
}

func (ms *MessageStack) InboxMessageAfter(startItemHash common.Hash) (inbox.InboxMessage, error) {
	item, ok := ms.itemAfterHash(startItemHash)
	if !ok {
		return inbox.InboxMessage{}, errors.New("one step proof startItemHash not found")
	}
	return item.message, nil
}

func (ms *MessageStack) GenerateVMInbox(olderAcc common.Hash, count uint64) (*VMInbox, error) {
	if count == 0 {
		return NewVMInbox(nil), nil
	}
	oldItem, ok := ms.itemAfterHash(olderAcc)
	if !ok {
		return nil, errors.New("olderAcc not found")
	}

	item := oldItem
	messages := make([]inbox.InboxMessage, 0, count)
	for i := uint64(0); i < count; i++ {
		if item == nil {
			return nil, errors.New("not enough Messages in inbox")
		}
		messages = append(messages, item.message)
		item = item.next
	}
	return NewVMInbox(messages), nil
}

func (ms *MessageStack) GetMessages(olderAcc common.Hash, count uint64) ([]inbox.InboxMessage, error) {
	if count == 0 {
		return nil, nil
	}
	oldItem, ok := ms.itemAfterHash(olderAcc)
	if !ok {
		return nil, errors.New("olderAcc not found")
	}

	item := oldItem
	messages := make([]inbox.InboxMessage, 0, count)
	for i := uint64(0); i < count; i++ {
		if item == nil {
			return nil, errors.New("not enough Messages in inbox")
		}
		messages = append(messages, item.message)
		item = item.next
	}
	return messages, nil
}

func (ms *MessageStack) GetAssertionMessages(beforeInboxHash common.Hash, afterInboxHash common.Hash) ([]inbox.InboxMessage, error) {
	if beforeInboxHash == afterInboxHash {
		return nil, nil
	}
	item, ok := ms.itemAfterHash(beforeInboxHash)
	if !ok || item == nil {
		return nil, errors.New("beforeInboxHash not found")
	}

	messages := make([]inbox.InboxMessage, 0)
	for item.hash != afterInboxHash {
		messages = append(messages, item.message)
		item = item.next
		if item == nil {
			return nil, errors.New("not enough Messages in inbox")
		}
	}
	return messages, nil
}

func (ms *MessageStack) GetAllMessagesAfter(olderAcc common.Hash) ([]inbox.InboxMessage, error) {
	item, ok := ms.itemAfterHash(olderAcc)
	if !ok {
		return nil, errors.New("olderAcc not found")
	}

	messages := make([]inbox.InboxMessage, 0)
	for item != nil {
		messages = append(messages, item.message)
		item = item.next
	}
	return messages, nil
}

func (ms *MessageStack) GetAllMessages() []inbox.InboxMessage {
	msgs := make([]inbox.InboxMessage, 0)
	for item := ms.oldest; item != nil; item = item.next {
		msgs = append(msgs, item.message)
	}
	return msgs
}

func (ms *MessageStack) GetAllHashes() []common.Hash {
	hashes := make([]common.Hash, 0)
	hashes = append(hashes, ms.hashOfRest)
	for item := ms.oldest; item != nil; item = item.next {
		hashes = append(hashes, item.hash)
	}
	return hashes
}

type Inbox struct {
	*MessageStack
}

func NewInbox() *Inbox {
	return &Inbox{
		MessageStack: NewMessageStack(),
	}
}

func (pi *Inbox) DiscardUpToCount(count *big.Int) {
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

func (pi *Inbox) Equals(pi2 *Inbox) bool {
	return pi.MessageStack.Equals(pi2.MessageStack)
}
