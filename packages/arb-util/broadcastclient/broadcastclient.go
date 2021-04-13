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

	logger.Info().Msgf("Connecting to arbitrum inbox message broadcaster: %v", bc.websocketUrl)
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), bc.websocketUrl)
	if err != nil {
		logger.Error().Msgf("Broadcast Client unable to connect: %v\n", err)
		return nil, err
	} else {
		logger.Info().Msg("Connected")
	}

	poller, err := netpoll.New(nil)
	if err != nil {
		logger.Error().Msgf("Error starting net poller %v\n", err)
	}

	// Get netpoll descriptor with EventRead|EventEdgeTriggered.
	desc := netpoll.Must(netpoll.HandleRead(conn))

	bc.desc = *desc
	bc.poller = poller
	bc.conn = conn

	poller.Start(desc, func(ev netpoll.Event) {
		if ev&netpoll.EventReadHup != 0 {
			logger.Info().Msgf("received hang up")
			poller.Stop(desc)
			conn.Close()
			return
		}

		msg, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			logger.Error().Msgf("can not receive: %v, %v", op, err)
			poller.Stop(desc)
			conn.Close()
			return
		} else {
			res := BroadcastMessages{}
			json.Unmarshal([]byte(msg), &res)
			messageReceiver <- res
		}
	})

	return messageReceiver, nil
}

func (bc *BroadcastClient) Close() {
	bc.poller.Stop(&bc.desc)
	bc.conn.Close()
}
