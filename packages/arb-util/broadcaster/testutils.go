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
	"errors"
	"math/big"
	"reflect"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

// This is used to reverse the slice for the sequence number field
// so that it will be in correct byte order when creating a new big.Int out of it
func reverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= (valueLen-1)/2; i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

func setSequenceNumberInData(data []byte, sequenceNumber *big.Int) []byte {
	seqNumOffset := 85
	seqNumEnd := seqNumOffset + 32
	prefixData := data[:seqNumOffset]
	postfixData := data[seqNumEnd:]
	sequenceNumberBytes := sequenceNumber.Bytes()
	sequenceNumberByteField := make([]byte, 32)
	copy(sequenceNumberByteField, sequenceNumberBytes)
	reverseSlice(sequenceNumberByteField)
	sequenceNumberWithPrefix := append(prefixData, sequenceNumberByteField...)
	completeDataWithSequenceNumberSet := append(sequenceNumberWithPrefix, postfixData...)
	return completeDataWithSequenceNumberSet
}

// SequencedMessages returns a function that when called returns the next random message in the sequence
func SequencedMessages() func() (common.Hash, SequencerFeedItem, *big.Int) {
	sequenceNumber := big.NewInt(41)
	accumulator := common.RandHash()

	return func() (common.Hash, SequencerFeedItem, *big.Int) {
		prevAccumulator := accumulator

		batchItem := inbox.SequencerBatchItem{
			LastSeqNum:  big.NewInt(0).Set(sequenceNumber),
			Accumulator: common.RandHash(),
		}

		sequenceNumber = sequenceNumber.Add(sequenceNumber, big.NewInt(1))
		accumulator = batchItem.Accumulator
		batchItem.TotalDelayedCount = big.NewInt(0)
		batchItem.SequencerMessage = setSequenceNumberInData(common.RandBytes(200), sequenceNumber)

		signature := common.RandBigInt()

		feedItem := SequencerFeedItem{
			BatchItem: batchItem,
			PrevAcc:   prevAccumulator,
		}

		return prevAccumulator, feedItem, signature
	}
}

// RandomMessageGenerator sends out generated test broadcast messages
type RandomMessageGenerator struct {
	broadcaster      *Broadcaster
	cancelFunc       context.CancelFunc
	count            int
	intervalDuration time.Duration
	workerStarted    bool
}

// NewRandomMessageGenerator creates a new test message generator
func NewRandomMessageGenerator(count int, interval time.Duration) *RandomMessageGenerator {
	gm := &RandomMessageGenerator{}
	gm.intervalDuration = interval
	gm.count = count
	return gm
}

// SetBroadcaster sets a client manager to broadcast on.
func (mg *RandomMessageGenerator) SetBroadcaster(broadcaster *Broadcaster) {
	mg.broadcaster = broadcaster
}

func (mg *RandomMessageGenerator) Start(parentCtx context.Context) <-chan error {
	errChan := make(chan error, 1)
	ctx, cancelFunc := context.WithCancel(parentCtx)

	go func() {
		defer cancelFunc()
		var ticker *time.Ticker
		if mg.intervalDuration > 0 {
			ticker = time.NewTicker(mg.intervalDuration)
		}

		prevAcc := common.RandHash()
		currAcc := common.RandHash()

		var lastSeq int64

		messageCount := 0
		if ticker != nil {
			for {
				currSeq := lastSeq + 1
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					err := mg.broadcaster.BroadcastSingle(
						prevAcc,
						inbox.SequencerBatchItem{
							LastSeqNum:        big.NewInt(lastSeq),
							Accumulator:       currAcc,
							TotalDelayedCount: big.NewInt(0),
							SequencerMessage:  big.NewInt(currSeq).Bytes(),
						},
						make([]byte, 0),
					)
					if err != nil {
						errChan <- errors.New("error broadcasting message")
						return
					}
					messageCount++
					prevAcc = currAcc
					lastSeq = currSeq
					if messageCount == mg.count {
						return
					}
				}
			}
		} else {
			for {
				currSeq := lastSeq + 1
				select {
				case <-ctx.Done():
					return
				default:
				}

				err := mg.broadcaster.BroadcastSingle(
					prevAcc,
					inbox.SequencerBatchItem{
						LastSeqNum:        big.NewInt(lastSeq),
						Accumulator:       currAcc,
						TotalDelayedCount: big.NewInt(0),
						SequencerMessage:  big.NewInt(currSeq).Bytes(),
					},
					make([]byte, 0),
				)
				if err != nil {
					errChan <- errors.New("error broadcasting message")
					return
				}
				messageCount++
				prevAcc = currAcc
				lastSeq = currSeq
				if messageCount == mg.count {
					return
				}
			}
		}
	}()

	return errChan
}

func (mg *RandomMessageGenerator) Stop() {
	mg.cancelFunc()
}
