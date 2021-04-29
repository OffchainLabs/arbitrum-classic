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
	"math/big"
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog/log"
)

type BroadcastClient struct {
	websocketUrl                 string
	lastInboxSeqNum              *big.Int
	conn                         net.Conn
	desc                         netpoll.Desc
	poller                       netpoll.Poller
	startingBroadcastClientMutex *sync.Mutex
	RetryCount                   int
	retrying                     bool
	ConfirmedAccumulatorListener chan common.Hash
}

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

func NewBroadcastClient(websocketUrl string, lastInboxSeqNum *big.Int) *BroadcastClient {
	bc := &BroadcastClient{}
	bc.startingBroadcastClientMutex = &sync.Mutex{}
	bc.websocketUrl = websocketUrl
	if lastInboxSeqNum == nil {
		bc.lastInboxSeqNum = big.NewInt(0)
	} else {
		bc.lastInboxSeqNum = lastInboxSeqNum
	}

	return bc
}

func (bc *BroadcastClient) Connect() (chan broadcaster.BroadcastFeedMessage, error) {
	messageReceiver := make(chan broadcaster.BroadcastFeedMessage)
	bc.ConfirmedAccumulatorListener = make(chan common.Hash)
	return bc.connect(messageReceiver)
}

func (bc *BroadcastClient) GetConfirmedAccumulatorListner() chan common.Hash {
	return bc.ConfirmedAccumulatorListener
}

func (bc *BroadcastClient) connect(messageReceiver chan broadcaster.BroadcastFeedMessage) (chan broadcaster.BroadcastFeedMessage, error) {
	if len(bc.websocketUrl) == 0 {
		// Nothing to do
		return nil, nil
	}

	logger.Info().Str("url", bc.websocketUrl).Msg("connecting to arbitrum inbox message broadcaster")
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), bc.websocketUrl)
	if err != nil {
		logger.Error().Err(err).Msg("broadcast client unable to connect")
		return nil, err
	} else {
		logger.Info().Msg("Connected")
	}

	poller, err := netpoll.New(nil)
	if err != nil {
		logger.Error().Err(err).Msg("error starting net poller")
		return nil, err
	}

	// Get netpoll descriptor with EventRead|EventEdgeTriggered.
	desc, err := netpoll.HandleRead(conn)
	if err != nil {
		logger.Error().Err(err).Msg("error getting netpoll descriptor")
		return nil, err
	}

	bc.desc = *desc
	bc.poller = poller
	bc.conn = conn

	err = poller.Start(desc, func(ev netpoll.Event) {
		if ev&netpoll.EventReadHup != 0 {
			logger.Info().Msg("received hang up")
			_ = poller.Stop(desc)
			_ = conn.Close()
			bc.RetryConnect(messageReceiver)
			return
		}

		if ev != 0 {
			logger.Info().Int("event", int(ev)).Msg("non-zero netpoll event")
		}

		msg, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			logger.Error().Err(err).Int("opcode", int(op)).Msgf("error calling ReadServerData")
			_ = poller.Stop(desc)
			_ = conn.Close()
			return
		}

		res := broadcaster.BroadcastMessage{}
		err = json.Unmarshal(msg, &res)
		if err != nil {
			logger.Error().Err(err).Msg("error unmarshalling message")
			return
		}

		if res.Version == 1 {
			for _, message := range res.Messages {
				messageReceiver <- *message
			}

			if res.ConfirmedAccumulator.IsConfirmed {
				bc.ConfirmedAccumulatorListener <- res.ConfirmedAccumulator.Accumulator
			}
		}
	})

	return messageReceiver, err
}

func (bc *BroadcastClient) Ping() (<-chan string, error) {

	_, err := bc.conn.Write(ws.CompiledPing)
	if err != nil {
		return nil, err
	}

	out := make(chan string)
	go func() {
		h, _, err := wsutil.NextReader(bc.conn, ws.StateClientSide)
		if err != nil {
			out <- err.Error()
			return
		}

		switch h.OpCode {
		case ws.OpPong:
			logger.Info().Msg("pong")
			out <- "pong"
		default:
			str := fmt.Sprintf("Received unknown JSON OpCode from server after ping: %v", h.OpCode)
			logger.Error().Msg(str)
			out <- str
		}
	}()

	return out, nil
}

func (bc *BroadcastClient) RetryConnect(messageReceiver chan broadcaster.BroadcastFeedMessage) {
	MaxWaitMs := 15000
	waitMs := 500
	bc.retrying = true
	for {
		bc.RetryCount++
		_, err := bc.connect(messageReceiver)
		if err == nil {
			bc.retrying = false
			return
		}
		time.Sleep(time.Duration(waitMs) * time.Millisecond)
		if waitMs < MaxWaitMs {
			waitMs += 500
		}
	}
}

func (bc *BroadcastClient) Close() {
	_ = bc.poller.Stop(&bc.desc)
	_ = bc.conn.Close()
}
