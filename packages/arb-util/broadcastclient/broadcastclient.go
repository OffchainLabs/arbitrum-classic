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
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/wsbroadcastserver"
)

type BroadcastClient struct {
	websocketUrl    string
	lastInboxSeqNum *big.Int

	connMutex *sync.Mutex
	conn      net.Conn

	retryMutex *sync.Mutex
	retryCount int

	retrying                     bool
	shuttingDown                 bool
	ConfirmedAccumulatorListener chan common.Hash
	idleTimeout                  time.Duration
}

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

func NewBroadcastClient(websocketUrl string, lastInboxSeqNum *big.Int, idleTimeout time.Duration) *BroadcastClient {
	var seqNum *big.Int
	if lastInboxSeqNum == nil {
		seqNum = big.NewInt(0)
	} else {
		seqNum = lastInboxSeqNum
	}

	return &BroadcastClient{
		websocketUrl:    websocketUrl,
		lastInboxSeqNum: seqNum,
		connMutex:       &sync.Mutex{},
		retryMutex:      &sync.Mutex{},
		idleTimeout:     idleTimeout,
	}
}

func (bc *BroadcastClient) Connect(ctx context.Context) (chan broadcaster.BroadcastFeedMessage, error) {
	messageReceiver := make(chan broadcaster.BroadcastFeedMessage)

	return messageReceiver, bc.ConnectWithChannel(ctx, messageReceiver)
}

func (bc *BroadcastClient) ConnectWithChannel(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) error {
	_, err := bc.connect(ctx, messageReceiver)
	if err != nil {
		return err
	}

	bc.startBackgroundReader(ctx, messageReceiver)

	return nil
}

func (bc *BroadcastClient) ConnectInBackground(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) {
	go (func() {
		for {
			err := bc.ConnectWithChannel(ctx, messageReceiver)
			if err == nil {
				break
			}
			logger.Warn().Str("url", bc.websocketUrl).Err(err).
				Msg("failed connect to sequencer broadcast, waiting and retrying")
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	})()
}

func (bc *BroadcastClient) connect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) (chan broadcaster.BroadcastFeedMessage, error) {

	if len(bc.websocketUrl) == 0 {
		// Nothing to do
		return nil, nil
	}

	logger.Info().Str("url", bc.websocketUrl).Msg("connecting to arbitrum inbox message broadcaster")
	timeoutDialer := ws.Dialer{
		Timeout: 10 * time.Second,
	}

	conn, _, _, err := timeoutDialer.Dial(ctx, bc.websocketUrl)
	if err != nil {
		logger.Warn().Err(err).Msg("broadcast client unable to connect")
		return nil, errors.Wrap(err, "broadcast client unable to connect")
	}

	bc.connMutex.Lock()
	bc.conn = conn
	bc.connMutex.Unlock()

	logger.Info().Msg("Connected")

	return messageReceiver, nil
}

func (bc *BroadcastClient) startBackgroundReader(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			msg, op, err := wsbroadcastserver.ReadData(ctx, bc.conn, bc.idleTimeout, ws.StateClientSide)
			if err != nil {
				if bc.shuttingDown {
					return
				}
				if strings.Contains(err.Error(), "i/o timeout") {
					logger.Error().Str("feed", bc.websocketUrl).Msg("Server connection timed out without receiving data")
				} else {
					logger.Error().Err(err).Str("feed", bc.websocketUrl).Int("opcode", int(op)).Msgf("error calling readData")
				}
				_ = bc.conn.Close()
				bc.RetryConnect(ctx, messageReceiver)
				continue
			}

			if msg != nil {
				res := broadcaster.BroadcastMessage{}
				err = json.Unmarshal(msg, &res)
				if err != nil {
					logger.Error().Err(err).Str("message", string(msg)).Msg("error unmarshalling message")
					continue
				}

				if len(res.Messages) > 0 {
					logger.Debug().Int("count", len(res.Messages)).Hex("acc", res.Messages[0].FeedItem.BatchItem.Accumulator.Bytes()).Msg("received batch item")
				} else if res.ConfirmedAccumulator.IsConfirmed {
					logger.Debug().Hex("acc", res.ConfirmedAccumulator.Accumulator.Bytes()).Msg("confirmed accumulator")
				} else {
					logger.Debug().Int("length", len(msg)).Msg("received broadcast without any messages or confirmations")
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
		}
	}()
}

func (bc *BroadcastClient) GetRetryCount() int {
	bc.retryMutex.Lock()
	defer bc.retryMutex.Unlock()

	return bc.retryCount
}

func (bc *BroadcastClient) RetryConnect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) {
	bc.retryMutex.Lock()
	defer bc.retryMutex.Unlock()

	maxWaitDuration := 15 * time.Second
	waitDuration := 500 * time.Millisecond
	bc.retrying = true
	for !bc.shuttingDown {
		select {
		case <-ctx.Done():
			return
		case <-time.After(waitDuration):
		}

		bc.retryCount++
		_, err := bc.connect(ctx, messageReceiver)
		if err == nil {
			bc.retrying = false
			return
		}

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
