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

	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type messageStackItem struct {
	message message.PendingMessage
	prev    *messageStackItem
	next    *messageStackItem
	hash    common.Hash
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
		newest:     nil,
		oldest:     nil,
		index:      make(map[common.Hash]*messageStackItem),
		hashOfRest: value.NewEmptyTuple().Hash(),
	}
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

func (ms *MessageStack) bottomIndex() *big.Int {
	if ms.oldest == nil {
		return big.NewInt(0)
	} else {
		return ms.oldest.count
	}
}

func (ms *MessageStack) DeliverMessage(msg message.PendingMessage) {
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
		return value.NewEmptyTuple().Hash(), nil
	}
	if height.Cmp(ms.bottomIndex()) < 0 {
		return common.Hash{}, errors.New("Height is below bottom of message stack")
	}
	if height.Cmp(ms.TopCount()) > 0 {
		return common.Hash{}, errors.New("Height is above top of message stack")
	}
	offset := new(big.Int).Sub(height, ms.bottomIndex())
	return ms.oldest.skipNext(offset.Uint64()).hash, nil
}

func hash2(h1, h2 common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(h1), hashing.Bytes32(h2))
}

func (ms *MessageStack) MarshalForCheckpoint(ctx CheckpointContext) *PendingInboxBuf {
	var items []*InboxItemBuf
	for item := ms.newest; item != nil; item = item.prev {
		checkpointVal := item.message.CheckpointValue()
		ctx.AddValue(checkpointVal)
		items = append(items, &InboxItemBuf{
			ValType: uint32(item.message.Type()),
			ValHash: checkpointVal.Hash().MarshalToBuf(),
		})
	}
	var topCount *big.Int
	if ms.newest == nil {
		topCount = big.NewInt(0)
	} else {
		topCount = ms.newest.count
	}
	return &PendingInboxBuf{
		TopCount:   common.MarshalBigInt(topCount),
		Items:      items,
		HashOfRest: ms.hashOfRest.MarshalToBuf(),
	}
}

func (buf *PendingInboxBuf) UnmarshalFromCheckpoint(ctx RestoreContext) (*MessageStack, error) {
	ret := NewMessageStack()
	ret.hashOfRest = buf.HashOfRest.Unmarshal()
	for i := len(buf.Items) - 1; i >= 0; i = i - 1 {
		val := ctx.GetValue(buf.Items[i].ValHash.Unmarshal())
		msg, err := message.UnmarshalFromCheckpoint(message.MessageType(buf.Items[i].ValType), val)
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

func (ms *MessageStack) GenerateBisection(startItemHash common.Hash, segments, count uint64) ([]common.Hash, error) {
	startItem, ok := ms.itemAfterHash(startItemHash)
	if !ok {
		return nil, errors.New("bisection startItemHash not found")
	}

	if count < segments {
		segments = count
	}

	cuts := make([]common.Hash, 0, segments+1)
	cuts = append(cuts, startItemHash)
	firstSegmentSize := count/segments + count%segments
	otherSegmentSize := count / segments
	item := startItem.skipNext(firstSegmentSize - 1)
	if item == nil {
		return nil, errors.New("pending inbox too short start")
	}
	cuts = append(cuts, item.hash)
	for i := uint64(1); i < segments; i++ {
		item = item.skipNext(otherSegmentSize)
		if item == nil {
			return nil, errors.New("pending inbox too short rest")
		}
		cuts = append(cuts, item.hash)
	}
	return cuts, nil
}

func (ms *MessageStack) GenerateOneStepProof(startItemHash common.Hash) (message.PendingMessage, error) {
	item, ok := ms.itemAfterHash(startItemHash)
	if !ok {
		return nil, errors.New("one step proof startItemHash not found")
	}
	return item.message, nil
}

func (ms *MessageStack) GenerateInbox(olderAcc common.Hash, count uint64) (*Inbox, error) {
	if count == 0 {
		return NewInbox(), nil
	}
	oldItem, ok := ms.itemAfterHash(olderAcc)
	if !ok {
		return nil, errors.New("olderAcc not found")
	}

	item := oldItem
	inbox := NewInbox()
	for i := uint64(0); i < count; i++ {
		if item == nil {
			return nil, errors.New("Not enough Messages in pending inbox")
		}
		inbox.DeliverMessage(item.message)
		item = item.next
	}
	return inbox, nil
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
