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
	"github.com/rs/zerolog/log"
)

type BroadcastClient struct {
	websocketUrl                 string
	lastInboxSeqNum              *big.Int
	conn                         net.Conn
	connMutex                    *sync.Mutex
	RetryCount                   int
	retrying                     bool
	shuttingDown                 bool
	ConfirmedAccumulatorListener chan common.Hash
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
		connMutex:       &sync.Mutex{},
		websocketUrl:    websocketUrl,
		lastInboxSeqNum: seqNum,
	}
}

func (bc *BroadcastClient) Connect(ctx context.Context) chan broadcaster.BroadcastFeedMessage {
	messageReceiver := make(chan broadcaster.BroadcastFeedMessage)
	return bc.connect(ctx, messageReceiver)
}

func (bc *BroadcastClient) connect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) chan broadcaster.BroadcastFeedMessage {
	if len(bc.websocketUrl) == 0 {
		// Nothing to do
		return nil
	}

	logger.Info().Str("url", bc.websocketUrl).Msg("connecting to arbitrum inbox message broadcaster")
	go func() {
		timeoutDialer := ws.Dialer{
			Timeout: 10 * time.Second,
		}

		for {
			conn, _, _, err := timeoutDialer.Dial(ctx, bc.websocketUrl)
			if err != nil {
				logger.Warn().Err(err).Msg("broadcast client unable to connect, retrying")

				select {
				case <-ctx.Done():
					return
				case <-time.After(10 * time.Second):
				}

				continue
			}

			bc.connMutex.Lock()
			bc.conn = conn
			bc.connMutex.Unlock()

			logger.Info().Msg("Connected")
			break
		}

		bc.startBackgroundReader(ctx, messageReceiver)
	}()

	return messageReceiver
}

func (bc *BroadcastClient) startBackgroundReader(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			msg, op, err := wsutil.ReadServerData(bc.conn)
			if err != nil {
				if bc.shuttingDown {
					return
				}
				logger.Error().Err(err).Int("opcode", int(op)).Msgf("error calling ReadServerData")
				_ = bc.conn.Close()
				// Starts up a new backgroundReader
				bc.RetryConnect(ctx, messageReceiver)
				return
			}

			res := broadcaster.BroadcastMessage{}
			err = json.Unmarshal(msg, &res)
			if err != nil {
				logger.Error().Err(err).Msg("error unmarshalling message")
				continue
			}

			if len(res.Messages) > 0 {
				logger.Debug().Int("count", len(res.Messages)).Hex("acc", res.Messages[0].FeedItem.BatchItem.Accumulator.Bytes()).Msg("received batch item")
			} else {
				logger.Debug().Int("length", len(msg)).Msg("received broadcast without any messages")
			}

			if res.Version == 1 {
				for _, message := range res.Messages {
					messageReceiver <- *message
				}

				if res.ConfirmedAccumulator.IsConfirmed && bc.ConfirmedAccumulatorListener != nil {
					bc.ConfirmedAccumulatorListener <- res.ConfirmedAccumulator.Accumulator
				}
			}
		}
	}()
}

func (bc *BroadcastClient) RetryConnect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) {
	maxWaitDuration := 15 * time.Second
	waitDuration := 500 * time.Millisecond
	bc.retrying = true
	for !bc.shuttingDown {
		select {
		case <-ctx.Done():
			return
		case <-time.After(waitDuration):
		}

		bc.RetryCount++
		_ = bc.connect(ctx, messageReceiver)

		if waitDuration < maxWaitDuration {
			waitDuration += 500 * time.Millisecond
		}
	}
}

func (bc *BroadcastClient) Close() {
	logger.Debug().Msg("closing broadcaster client connection")
	bc.shuttingDown = true
	bc.connMutex.Lock()
	if bc.conn != nil {
		_ = bc.conn.Close()
	}
	bc.connMutex.Unlock()
}
