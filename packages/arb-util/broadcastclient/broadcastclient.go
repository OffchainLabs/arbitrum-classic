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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"io"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/gobwas/ws"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/wsbroadcastserver"
)

type BroadcastClient struct {
	websocketUrl string

	// sequence number of the previous message
	lastInboxSeqNum *big.Int

	chainId uint64

	connMutex *sync.Mutex
	conn      net.Conn
	errChan   chan error

	retryMutex *sync.Mutex
	retryCount int

	retrying                     bool
	shuttingDown                 bool
	ConfirmedAccumulatorListener chan common.Hash
	idleTimeout                  time.Duration
}

var logger = arblog.Logger.With().Str("component", "broadcaster").Logger()

func NewBroadcastClient(
	websocketUrl string,
	chainId uint64,
	requestedInboxSeqNum *big.Int,
	idleTimeout time.Duration,
	broadcastClientErrChan chan error,
) *BroadcastClient {
	var lastSeqNum *big.Int
	if requestedInboxSeqNum == nil || requestedInboxSeqNum.Cmp(big.NewInt(0)) <= 0 {
		lastSeqNum = big.NewInt(0)
	} else {
		lastSeqNum = new(big.Int).Sub(requestedInboxSeqNum, big.NewInt(1))
	}

	return &BroadcastClient{
		websocketUrl:    websocketUrl,
		chainId:         chainId,
		lastInboxSeqNum: lastSeqNum,
		connMutex:       &sync.Mutex{},
		errChan:         broadcastClientErrChan,
		retryMutex:      &sync.Mutex{},
		idleTimeout:     idleTimeout,
	}
}

func (bc *BroadcastClient) Connect(ctx context.Context) (chan broadcaster.BroadcastFeedMessage, error) {
	messageReceiver := make(chan broadcaster.BroadcastFeedMessage)

	err := bc.ConnectWithChannel(ctx, messageReceiver)
	if err != nil {
		return nil, err
	}
	return messageReceiver, nil
}

func (bc *BroadcastClient) ConnectWithChannel(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) error {
	earlyFrameData, _, err := bc.connect(ctx, messageReceiver, bc.lastInboxSeqNum)
	if err != nil {
		return err
	}

	bc.startBackgroundReader(ctx, messageReceiver, earlyFrameData)

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

var ErrIncorrectFeedServerVersion = errors.New("incorrect feed server version")
var ErrIncorrectChainId = errors.New("incorrect chain id")

func (bc *BroadcastClient) connect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage, previousSequenceNumber *big.Int) (io.Reader, chan broadcaster.BroadcastFeedMessage, error) {

	if len(bc.websocketUrl) == 0 {
		// Nothing to do
		return nil, nil, nil
	}

	var requestedSequenceNumber string
	if previousSequenceNumber.Cmp(big.NewInt(0)) > 0 {
		// previousSequenceNumber is 1 before current, and we want 1 after current, so add 2.
		requestedSequenceNumber = new(big.Int).Add(previousSequenceNumber, big.NewInt(2)).String()
	} else {
		requestedSequenceNumber = "0"
	}
	header := ws.HandshakeHeaderHTTP(http.Header{
		wsbroadcastserver.HTTPHeaderFeedClientVersion:       []string{strconv.Itoa(wsbroadcastserver.FeedClientVersion)},
		wsbroadcastserver.HTTPHeaderRequestedSequenceNumber: []string{requestedSequenceNumber},
	})

	logger.Info().Str("url", bc.websocketUrl).Msg("connecting to arbitrum inbox message broadcaster")
	var feedServerVersion uint64
	timeoutDialer := ws.Dialer{
		Header: header,
		OnHeader: func(key, value []byte) (err error) {
			headerName := string(key)
			headerValue := string(value)
			if headerName == wsbroadcastserver.HTTPHeaderFeedServerVersion {
				feedServerVersion, err = strconv.ParseUint(headerValue, 0, 64)
				if err != nil {
					return err
				}
				if feedServerVersion != wsbroadcastserver.FeedServerVersion {
					logger.
						Error().
						Uint64("expectedFeedServerVersion", wsbroadcastserver.FeedServerVersion).
						Uint64("actualFeedServerVersion", feedServerVersion).
						Msg("incorrect feed server version")
					return ErrIncorrectFeedServerVersion
				}
			} else if headerName == wsbroadcastserver.HTTPHeaderChainId {
				chainId, err := strconv.ParseUint(headerValue, 0, 64)
				if err != nil {
					return err
				}
				if chainId != bc.chainId {
					logger.
						Error().
						Uint64("expectedChainId", bc.chainId).
						Uint64("actualChainId", chainId).
						Msg("incorrect chain id when connecting to server feed")
					return ErrIncorrectChainId
				}
			}
			return nil
		},
		Timeout: 10 * time.Second,
	}

	conn, br, _, err := timeoutDialer.Dial(ctx, bc.websocketUrl)
	if err != nil {
		logger.Warn().Err(err).Msg("broadcast client unable to connect")
		return nil, nil, errors.Wrap(err, "broadcast client unable to connect")
	}

	var earlyFrameData io.Reader
	if br != nil {
		// Depending on how long the client takes to read the response, there may be
		// data after the WebSocket upgrade response in a single read from the socket,
		// ie WebSocket frames sent by the server. If this happens, Dial returns
		// a non-nil bufio.Reader so that data isn't lost. But beware, this buffered
		// reader is still hooked up to the socket; trying to read past what had already
		// been buffered will do a blocking read on the socket, so we have to wrap it
		// in a LimitedReader.
		earlyFrameData = io.LimitReader(br, int64(br.Buffered()))
	}

	bc.connMutex.Lock()
	bc.conn = conn
	bc.connMutex.Unlock()

	logger.Info().Uint64("chainId", bc.chainId).Uint64("feedServerVersion", feedServerVersion).Msg("Connected")

	return earlyFrameData, messageReceiver, nil
}

func (bc *BroadcastClient) startBackgroundReader(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage, earlyFrameData io.Reader) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			msg, op, err := wsbroadcastserver.ReadData(ctx, bc.conn, earlyFrameData, bc.idleTimeout, ws.StateClientSide)
			if err != nil {
				if bc.shuttingDown {
					return
				}
				if strings.Contains(err.Error(), "i/o timeout") {
					logger.Error().Str("feed", bc.websocketUrl).Msg("Server connection timed out without receiving data")
				} else if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
					logger.Warn().Err(err).Str("feed", bc.websocketUrl).Int("opcode", int(op)).Msgf("readData returned EOF")
				} else {
					logger.Error().Err(err).Str("feed", bc.websocketUrl).Int("opcode", int(op)).Msgf("error calling readData")
				}
				_ = bc.conn.Close()
				earlyFrameData = bc.RetryConnect(ctx, messageReceiver)
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
					var currentLastSeqNum *big.Int
					for _, message := range res.Messages {
						currentLastSeqNum = message.FeedItem.BatchItem.LastSeqNum
						messageReceiver <- *message
					}
					bc.lastInboxSeqNum = new(big.Int).Add(currentLastSeqNum, big.NewInt(1))

					if res.ConfirmedAccumulator.IsConfirmed && bc.ConfirmedAccumulatorListener != nil {
						bc.ConfirmedAccumulatorListener <- res.ConfirmedAccumulator.Accumulator
					}
				} else if res.Version >= 2 {
					bc.errChan <- fmt.Errorf("connected to nitro feed with classic client, server version: %d", res.Version)
					break
				} else {
					bc.errChan <- fmt.Errorf("unrecognized feed version, server version: %d", res.Version)
					break
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

func (bc *BroadcastClient) RetryConnect(ctx context.Context, messageReceiver chan broadcaster.BroadcastFeedMessage) io.Reader {
	bc.retryMutex.Lock()
	defer bc.retryMutex.Unlock()

	maxWaitDuration := 15 * time.Second
	waitDuration := 500 * time.Millisecond
	bc.retrying = true
	for !bc.shuttingDown {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(waitDuration):
		}

		bc.retryCount++
		earlyFrameData, _, err := bc.connect(ctx, messageReceiver, bc.lastInboxSeqNum)
		if err == nil {
			bc.retrying = false
			return earlyFrameData
		}

		if waitDuration < maxWaitDuration {
			waitDuration += 500 * time.Millisecond
		}
	}

	return nil
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
