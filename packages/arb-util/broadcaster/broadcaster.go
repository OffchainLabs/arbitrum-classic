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

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/wsbroadcastserver"

	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

type Broadcaster struct {
	acceptor      *wsbroadcastserver.Acceptor
	catchupBuffer wsbroadcastserver.CatchupBuffer
}

func NewBroadcaster(settings configuration.FeedOutput) *Broadcaster {
	catchupBuffer := NewConfirmedAccumulatorCatchupBuffer()
	return &Broadcaster{
		acceptor:      wsbroadcastserver.NewAcceptor(settings, catchupBuffer),
		catchupBuffer: catchupBuffer,
	}
}

func (b *Broadcaster) ClientCount() int32 {
	return b.acceptor.ClientCount()
}

func (b *Broadcaster) Start(ctx context.Context) error {
	return b.acceptor.Start(ctx)
}

func (b *Broadcaster) BroadcastSingle(prevAcc common.Hash, batchItem inbox.SequencerBatchItem, signature []byte) error {
	var broadcastMessages []*BroadcastFeedMessage

	logger.Debug().Hex("acc", batchItem.Accumulator.Bytes()).Msg("sending batch Item")

	msg := BroadcastFeedMessage{
		FeedItem: SequencerFeedItem{
			BatchItem: batchItem,
			PrevAcc:   prevAcc,
		},
		Signature: signature,
	}

	broadcastMessages = append(broadcastMessages, &msg)

	bm := BroadcastMessage{
		Version:  1,
		Messages: broadcastMessages,
	}

	b.acceptor.Broadcast(bm)
	return nil
}

func (b *Broadcaster) Broadcast(prevAcc common.Hash, batchItems []inbox.SequencerBatchItem, dataSigner func([]byte) ([]byte, error)) error {
	for _, item := range batchItems {
		signature, err := dataSigner(hashing.SoliditySHA3WithPrefix(hashing.Bytes32(item.Accumulator)).Bytes())
		if err != nil {
			return err
		}

		err = b.BroadcastSingle(prevAcc, item, signature)
		if err != nil {
			return err
		}
		prevAcc = item.Accumulator
	}

	return nil
}

func (b *Broadcaster) ConfirmedAccumulator(accumulator common.Hash) {
	logger.Debug().Hex("acc", accumulator.Bytes()).Msg("confirming accumulator")

	bm := BroadcastMessage{
		Version: 1,
		ConfirmedAccumulator: ConfirmedAccumulator{
			IsConfirmed: true,
			Accumulator: accumulator,
		},
	}

	b.acceptor.Broadcast(bm)
}

func (b *Broadcaster) MessageCacheCount() int {
	return b.catchupBuffer.GetMessageCount()
}

func (b *Broadcaster) Stop() {
	b.acceptor.Stop()
}
