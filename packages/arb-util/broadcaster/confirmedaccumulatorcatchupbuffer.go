/*
 * Copyright 2021, Offchain Labs, Inc.
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
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/wsbroadcastserver"
)

type ConfirmedAccumulatorCatchupBuffer struct {
	broadcastMessages []*BroadcastFeedMessage
	cacheSize         int32
}

func NewConfirmedAccumulatorCatchupBuffer() *ConfirmedAccumulatorCatchupBuffer {
	return &ConfirmedAccumulatorCatchupBuffer{}
}

func (q *ConfirmedAccumulatorCatchupBuffer) getCacheMessages(requestedSeqNum *big.Int) *BroadcastMessage {
	if len(q.broadcastMessages) == 0 {
		return nil
	}
	startingIndex := 0
	// Ignore messages older than requested sequence number
	if requestedSeqNum.Cmp(big.NewInt(0)) > 0 {
		requestedLastSeqNum := new(big.Int).Sub(requestedSeqNum, big.NewInt(1))
		if q.broadcastMessages[0].FeedItem.BatchItem.LastSeqNum.Cmp(requestedLastSeqNum) < 0 {
			startingIndex = int(new(big.Int).Sub(requestedLastSeqNum, q.broadcastMessages[0].FeedItem.BatchItem.LastSeqNum).Uint64())
			if startingIndex >= len(q.broadcastMessages) {
				startingIndex = len(q.broadcastMessages) - 1
			}
			comparison := q.broadcastMessages[startingIndex].FeedItem.BatchItem.LastSeqNum.Cmp(requestedLastSeqNum)
			if comparison > 0 {
				for startingIndex > 1 {
					comparison2 := q.broadcastMessages[startingIndex-1].FeedItem.BatchItem.LastSeqNum.Cmp(requestedLastSeqNum)
					if comparison2 < 0 {
						// Found messages to broadcast
						break
					}
					startingIndex--
				}
			} else if comparison < 0 {
				for {
					startingIndex++
					if startingIndex >= len(q.broadcastMessages) {
						// End of array with nothing found
						return nil
					}
					if q.broadcastMessages[startingIndex].FeedItem.BatchItem.LastSeqNum.Cmp(requestedLastSeqNum) >= 0 {
						// Found messages to broadcast
						break
					}
				}
			}
		}
	}

	messagesToSend := q.broadcastMessages[startingIndex:]
	if len(messagesToSend) > 0 {
		bm := BroadcastMessage{
			Version:  1,
			Messages: messagesToSend,
		}

		return &bm
	}

	return nil
}

func (q *ConfirmedAccumulatorCatchupBuffer) OnRegisterClient(ctx context.Context, clientConnection *wsbroadcastserver.ClientConnection) error {
	start := time.Now()
	// send the newly connected client any messages starting with requested sequence number
	bm := q.getCacheMessages(clientConnection.RequestedSeqNum())
	if bm != nil {
		err := clientConnection.Write(bm)
		if err != nil {
			logger.Error().Err(err).Str("client", clientConnection.Name).Str("elapsed", time.Since(start).String()).Msg("error sending client cached messages")
			return err
		}
	}

	logger.Info().Str("client", clientConnection.Name).Str("elapsed", time.Since(start).String()).Msg("client registered")

	return nil
}

func (q *ConfirmedAccumulatorCatchupBuffer) OnDoBroadcast(bmi interface{}) error {
	bm := bmi.(BroadcastMessage)
	if bm.ConfirmedAccumulator.IsConfirmed {
		for i, msg := range q.broadcastMessages {
			if msg.FeedItem.BatchItem.Accumulator == bm.ConfirmedAccumulator.Accumulator {
				// This entry was confirmed, so this and all previous messages should be removed from cache
				unconfirmedIndex := i + 1
				if unconfirmedIndex >= len(q.broadcastMessages) {
					//  Nothing newer, so clear entire cache
					q.broadcastMessages = q.broadcastMessages[:0]
				} else {
					q.broadcastMessages = q.broadcastMessages[unconfirmedIndex:]
				}
				break
			}
		}
	} else if len(bm.Messages) > 0 {
		// Add to cache to send to new clients
		if len(q.broadcastMessages) == 0 {
			// Current list is empty
			q.broadcastMessages = append(q.broadcastMessages, bm.Messages...)
		} else if q.broadcastMessages[len(q.broadcastMessages)-1].FeedItem.BatchItem.Accumulator == bm.Messages[0].FeedItem.PrevAcc {
			q.broadcastMessages = append(q.broadcastMessages, bm.Messages...)
		} else {
			// We need to do a re-org
			logger.Debug().Hex("acc", bm.Messages[0].FeedItem.BatchItem.Accumulator.Bytes()).Msg("broadcaster reorg")
			i := len(q.broadcastMessages) - 1
			for ; i >= 0; i-- {
				if q.broadcastMessages[i].FeedItem.BatchItem.Accumulator == bm.Messages[0].FeedItem.PrevAcc {
					q.broadcastMessages = append(q.broadcastMessages[:i+1], bm.Messages...)
					break
				}
			}

			if i == -1 {
				// All existing messages are out of date
				q.broadcastMessages = append(q.broadcastMessages[:0], bm.Messages...)
			}
		}
	}

	atomic.StoreInt32(&q.cacheSize, int32(len(q.broadcastMessages)))

	return nil
}

func (q *ConfirmedAccumulatorCatchupBuffer) GetMessageCount() int {
	return int(atomic.LoadInt32(&q.cacheSize))
}
