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

package wsbroadcastserver

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

/* Protocol-specific client catch-up logic can be injected using this interface. */
type CatchupBuffer interface {
	OnRegisterClient(context.Context, *ClientConnection) error
	OnDoBroadcast(interface{}) error
	GetMessageCount() int
}

// ClientManager manages client connections
type ClientManager struct {
	cancelFunc    context.CancelFunc
	clientPtrMap  map[*ClientConnection]bool
	clientCount   int32
	pool          *gopool.Pool
	poller        netpoll.Poller
	broadcastChan chan interface{}
	clientAction  chan ClientConnectionAction
	settings      configuration.FeedOutput
	catchupBuffer CatchupBuffer
	maxSendQueue  int
}

type ClientConnectionAction struct {
	cc     *ClientConnection
	create bool
}

func NewClientManager(poller netpoll.Poller, settings configuration.FeedOutput, catchupBuffer CatchupBuffer) *ClientManager {
	return &ClientManager{
		poller:        poller,
		pool:          gopool.NewPool(settings.Workers, settings.Queue, 1),
		clientPtrMap:  make(map[*ClientConnection]bool),
		broadcastChan: make(chan interface{}, 1),
		clientAction:  make(chan ClientConnectionAction, 128),
		settings:      settings,
		catchupBuffer: catchupBuffer,
		maxSendQueue:  settings.MaxSendQueue,
	}
}

func (cm *ClientManager) registerClient(ctx context.Context, clientConnection *ClientConnection) error {
	if err := cm.catchupBuffer.OnRegisterClient(ctx, clientConnection); err != nil {
		return err
	}

	clientConnection.Start(ctx)
	cm.clientPtrMap[clientConnection] = true
	atomic.AddInt32(&cm.clientCount, 1)

	return nil
}

// Register registers new connection as a Client.
func (cm *ClientManager) Register(conn net.Conn, desc *netpoll.Desc) *ClientConnection {
	createClient := ClientConnectionAction{
		NewClientConnection(conn, desc, cm),
		true,
	}

	cm.clientAction <- createClient

	return createClient.cc
}

// removeAll removes all clients after main ClientManager thread exits
func (cm *ClientManager) removeAll() {
	// Only called after main ClientManager thread exits, so remove client directly
	for client := range cm.clientPtrMap {
		cm.removeClientImpl(client)
	}
}

func (cm *ClientManager) removeClientImpl(clientConnection *ClientConnection) {
	clientConnection.Stop()

	err := cm.poller.Stop(clientConnection.desc)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to stop poller")
	}

	err = clientConnection.conn.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to close client connection")
	}

	atomic.AddInt32(&cm.clientCount, -1)
}

func (cm *ClientManager) removeClient(clientConnection *ClientConnection) {
	if !cm.clientPtrMap[clientConnection] {
		return
	}

	cm.removeClientImpl(clientConnection)

	delete(cm.clientPtrMap, clientConnection)
}

func (cm *ClientManager) Remove(clientConnection *ClientConnection) {
	cm.clientAction <- ClientConnectionAction{
		clientConnection,
		false,
	}
}

func (cm *ClientManager) ClientCount() int32 {
	return atomic.LoadInt32(&cm.clientCount)
}

// Broadcast sends batch item to all clients.
func (cm *ClientManager) Broadcast(bm interface{}) {
	cm.broadcastChan <- bm
}

func (cm *ClientManager) doBroadcast(bm interface{}) ([]*ClientConnection, error) {
	if err := cm.catchupBuffer.OnDoBroadcast(bm); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	writer := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(bm); err != nil {
		return nil, errors.Wrap(err, "unable to encode message")
	}
	if err := writer.Flush(); err != nil {
		return nil, errors.Wrap(err, "unable to flush message")
	}

	clientDeleteList := make([]*ClientConnection, 0, len(cm.clientPtrMap))
	for client := range cm.clientPtrMap {
		if len(client.out) >= cm.maxSendQueue {
			// Queue for client too backed up, disconnect instead of blocking on channel send
			logger.Error().Str("client", client.Name).Int("sendQueue", len(client.out)).Msg("disconnecting because sendQueue too large")
			clientDeleteList = append(clientDeleteList, client)
		} else {
			client.out <- buf.Bytes()
		}
	}

	return clientDeleteList, nil
}

// verifyClients should be called every cm.settings.ClientPingInterval
func (cm *ClientManager) verifyClients() []*ClientConnection {
	clientConnectionCount := len(cm.clientPtrMap)

	// Create list of clients to remove
	clientDeleteList := make([]*ClientConnection, 0, clientConnectionCount)

	// Send ping to all connected clients
	logger.Debug().Int("count", len(cm.clientPtrMap)).Msg("pinging clients")
	for client := range cm.clientPtrMap {
		diff := time.Since(client.GetLastHeard())
		if diff > cm.settings.ClientTimeout {
			logger.Debug().Str("client", client.Name).Msg("disconnecting because connection timed out")
			clientDeleteList = append(clientDeleteList, client)
		} else {
			err := client.Ping()
			if err != nil {
				logger.Warn().Str("client", client.Name).Msg("disconnecting because error pinging client")
				clientDeleteList = append(clientDeleteList, client)
			}
		}
	}

	return clientDeleteList
}

func (cm *ClientManager) Stop() {
	cm.cancelFunc()
}

func (cm *ClientManager) Start(parentCtx context.Context) {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	cm.cancelFunc = cancelFunc

	go func() {
		defer cancelFunc()
		defer cm.removeAll()

		pingInterval := time.NewTicker(cm.settings.Ping)
		defer pingInterval.Stop()
		var clientDeleteList []*ClientConnection
		for {
			select {
			case <-ctx.Done():
				return
			case clientAction := <-cm.clientAction:
				if clientAction.create {
					err := cm.registerClient(ctx, clientAction.cc)
					if err != nil {
						// Log message already output in registerClient
						cm.removeClientImpl(clientAction.cc)
					}
				} else {
					cm.removeClient(clientAction.cc)
				}
			case bm := <-cm.broadcastChan:
				var err error
				clientDeleteList, err = cm.doBroadcast(bm)
				if err != nil {
					logError(err, "failed to do a broadcast")
				}
			case <-pingInterval.C:
				clientDeleteList = cm.verifyClients()
			}

			if len(clientDeleteList) > 0 {
				for _, client := range clientDeleteList {
					cm.removeClient(client)
				}
				clientDeleteList = nil
			}
		}
	}()
}
