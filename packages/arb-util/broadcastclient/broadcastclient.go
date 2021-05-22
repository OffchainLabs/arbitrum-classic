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

package broadcastclient

import (
	"context"
	"encoding/json"
	"math/big"
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/rs/zerolog/log"
)

type BroadcastClient struct {
	websocketUrl                  string
	lastInboxSeqNum               *big.Int
	conn                          net.Conn
	startingBroadcastClientMutex  *sync.Mutex
	RetryCount                    int
	retrying                      bool
	shuttingDown                  bool
	ConfirmedAccumulatorsListener chan [2]common.Hash
}

type BatchItemAndPrevAcc struct {
	BatchItem inbox.SequencerBatchItem
	PrevAcc   common.Hash
}

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

func NewBroadcastClient(websocketUrl string, lastInboxSeqNum *big.Int) *BroadcastClient {
	var seqNum *big.Int
	if lastInboxSeqNum == nil {
		seqNum = big.NewInt(0)
	} else {
		seqNum = lastInboxSeqNum
	}

	return &BroadcastClient{
		startingBroadcastClientMutex: &sync.Mutex{},
		websocketUrl:                 websocketUrl,
		lastInboxSeqNum:              seqNum,
	}
}

func (bc *BroadcastClient) Connect() (chan BatchItemAndPrevAcc, chan inbox.DelayedMessage, error) {
	messageReceiver := make(chan BatchItemAndPrevAcc, 128)
	delayedMessageReceiver := make(chan inbox.DelayedMessage, 1024)
	return messageReceiver, delayedMessageReceiver, bc.connect(messageReceiver, delayedMessageReceiver)
}

func (bc *BroadcastClient) connect(messageReceiver chan BatchItemAndPrevAcc, delayedMessageReceiver chan inbox.DelayedMessage) error {
	if len(bc.websocketUrl) == 0 {
		// Nothing to do
		return nil
	}

	logger.Info().Str("url", bc.websocketUrl).Msg("connecting to arbitrum inbox message broadcaster")
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), bc.websocketUrl)
	if err != nil {
		logger.Error().Err(err).Msg("broadcast client unable to connect")
		return err
	} else {
		logger.Info().Msg("Connected")
	}

	bc.conn = conn

	go bc.backgroundReader(messageReceiver, delayedMessageReceiver)

	return err
}

func (bc *BroadcastClient) backgroundReader(messageReceiver chan BatchItemAndPrevAcc, delayedMessageReceiver chan inbox.DelayedMessage) {
	for {
		msg, op, err := wsutil.ReadServerData(bc.conn)
		if err != nil {
			if bc.shuttingDown {
				return
			}
			logger.Error().Err(err).Int("opcode", int(op)).Msgf("error calling ReadServerData")
			_ = bc.conn.Close()
			// Starts up a new backgroundReader
			bc.RetryConnect(messageReceiver, delayedMessageReceiver)
			return
		}

		res := broadcaster.BroadcastMessage{}
		err = json.Unmarshal(msg, &res)
		if err != nil {
			logger.Error().Err(err).Msg("error unmarshalling message")
			continue
		}

		if len(res.Messages) > 0 {
			logger.Debug().Int("count", len(res.Messages)).Hex("acc", res.Messages[len(res.Messages)-1].Accumulator.Bytes()).Msg("received batch items")
		} else {
			logger.Debug().Int("length", len(msg)).Msg("received broadcast without any messages")
		}

		if res.Version == 2 {
			// TODO: check signature
			var prevAcc common.Hash
			if res.SequencedMetadata != nil {
				prevAcc = res.SequencedMetadata.PrevAcc
			}
			for _, message := range res.DelayedMessages {
				delayedMessageReceiver <- message
			}
			for _, message := range res.Messages {
				messageReceiver <- BatchItemAndPrevAcc{
					BatchItem: message,
					PrevAcc:   prevAcc,
				}
			}

			if res.ConfirmedAccumulators != nil && bc.ConfirmedAccumulatorsListener != nil {
				bc.ConfirmedAccumulatorsListener <- *res.ConfirmedAccumulators
			}
		}
	}
}

func (bc *BroadcastClient) RetryConnect(messageReceiver chan BatchItemAndPrevAcc, delayedMessageReceiver chan inbox.DelayedMessage) {
	MaxWaitMs := 15000
	waitMs := 500
	bc.retrying = true
	for !bc.shuttingDown {
		time.Sleep(time.Duration(waitMs) * time.Millisecond)

		bc.RetryCount++
		err := bc.connect(messageReceiver, delayedMessageReceiver)
		if err == nil {
			bc.retrying = false
			return
		}

		if waitMs < MaxWaitMs {
			waitMs += 500
		}
	}
}

func (bc *BroadcastClient) Close() {
	logger.Debug().Msg("closing broadcaster client connection")
	bc.shuttingDown = true
	_ = bc.conn.Close()
}
