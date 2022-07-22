/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

package broadcaster

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestGetEmptyCacheMessages(t *testing.T) {
	buffer := ConfirmedAccumulatorCatchupBuffer{
		broadcastMessages: []*BroadcastFeedMessage{},
		cacheSize:         10,
	}

	// Get everything
	bm := buffer.getCacheMessages(big.NewInt(0))
	if bm != nil {
		t.Error("shouldn't have returned anything")
	}
}

func createDummyBroadcastMessages(lastSeqNums []int) []*BroadcastFeedMessage {
	broadcastMessages := make([]*BroadcastFeedMessage, 0, len(lastSeqNums))
	for _, lastSeqNum := range lastSeqNums {
		broadcastMessage := &BroadcastFeedMessage{
			FeedItem: SequencerFeedItem{
				BatchItem: inbox.SequencerBatchItem{
					LastSeqNum:        big.NewInt(int64(lastSeqNum)),
					Accumulator:       common.Hash{},
					TotalDelayedCount: big.NewInt(0),
					SequencerMessage:  []byte{},
				},
				PrevAcc: common.Hash{},
			},
			Signature: []byte{},
		}
		broadcastMessages = append(broadcastMessages, broadcastMessage)
	}

	return broadcastMessages
}

func TestGetCacheMessages(t *testing.T) {
	buffer := ConfirmedAccumulatorCatchupBuffer{
		broadcastMessages: createDummyBroadcastMessages([]int{40, 40, 41, 45, 46, 47, 48}),
		cacheSize:         10,
	}

	// Get everything
	bm := buffer.getCacheMessages(big.NewInt(0))
	if len(bm.Messages) != 7 {
		t.Error("didn't return all messages")
	}

	// Get everything
	bm = buffer.getCacheMessages(big.NewInt(1))
	if len(bm.Messages) != 7 {
		t.Error("didn't return all messages")
	}

	// Get everything
	bm = buffer.getCacheMessages(big.NewInt(41))
	if len(bm.Messages) != 7 {
		t.Error("didn't return all messages")
	}

	// Get nothing
	bm = buffer.getCacheMessages(big.NewInt(100))
	if bm != nil {
		t.Error("should not have returned anything")
	}

	// Test single
	bm = buffer.getCacheMessages(big.NewInt(49))
	if len(bm.Messages) != 1 {
		t.Errorf("expected 1 message, got %d messages", len(bm.Messages))
	}
	if bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Cmp(big.NewInt(48)) != 0 {
		t.Errorf("expected lastSeqNum 48, got %d", bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Int64())
	}

	// Test when messages missing
	bm = buffer.getCacheMessages(big.NewInt(46))
	if len(bm.Messages) != 4 {
		t.Errorf("expected 4 messages, got %d messages", len(bm.Messages))
	}
	if bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Cmp(big.NewInt(45)) != 0 {
		t.Errorf("expected lastSeqNum 45, got %d", bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Int64())
	}

	// Test when duplicate message
	bm = buffer.getCacheMessages(big.NewInt(42))
	if len(bm.Messages) != 5 {
		t.Errorf("expected only 5 messages, got %d messages", len(bm.Messages))
	}
	if bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Cmp(big.NewInt(41)) != 0 {
		t.Errorf("expected lastSeqNum 41, got %d", bm.Messages[0].FeedItem.BatchItem.LastSeqNum.Int64())
	}

}
