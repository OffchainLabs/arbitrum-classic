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

	"github.com/ethereum/go-ethereum/common"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/rs/zerolog/log"
)

type Request struct {
	ID     int                    `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}
type BroadcastMessage struct {
	// Kind        Type           `json:"kind"`
	Sender      common.Address `json:"sender"`
	InboxSeqNum *big.Int       `json:"seqnum"`
	GasPrice    *big.Int       `json:"gasprice"`
	Data        []byte         `json:"data"`
}

type BroadcastMessages struct {
	Messages []BroadcastMessage `json:"messages"`
}

type BroadcastClient struct {
	websocketUrl    string
	lastInboxSeqNum *big.Int
	conn            net.Conn
	desc            netpoll.Desc
	poller          netpoll.Poller
}

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

func NewBroadcastClient(websocketUrl string, lastInboxSeqNum *big.Int) *BroadcastClient {
	bc := &BroadcastClient{}
	bc.websocketUrl = websocketUrl
	if lastInboxSeqNum == nil {
		bc.lastInboxSeqNum = big.NewInt(0)
	} else {
		bc.lastInboxSeqNum = lastInboxSeqNum
	}

	return bc
}

func (bc *BroadcastClient) Connect() (<-chan BroadcastMessages, error) {
	messageReceiver := make(chan BroadcastMessages)

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
	}

	// Get netpoll descriptor with EventRead|EventEdgeTriggered.
	desc := netpoll.Must(netpoll.HandleRead(conn))

	bc.desc = *desc
	bc.poller = poller
	bc.conn = conn

	err = poller.Start(desc, func(ev netpoll.Event) {
		if ev & netpoll.EventReadHup != 0 {
			logger.Info().Msg("received hang up")
			_ = poller.Stop(desc)
			_ = conn.Close()
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

		res := BroadcastMessages{}
		err = json.Unmarshal(msg, &res)
		if err != nil {
			logger.Error().Err(err).Msg("error unmarshalling message")
			return
		}

		messageReceiver <- res
	})

	return messageReceiver, nil
}

func (bc *BroadcastClient) Close() {
	_ = bc.poller.Stop(&bc.desc)
	_ = bc.conn.Close()
}
