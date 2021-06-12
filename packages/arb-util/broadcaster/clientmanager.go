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
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net"
	"sync/atomic"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ClientManagerSettings struct {
	ClientPingInterval      time.Duration
	ClientNoResponseTimeout time.Duration
}

// ClientManager manages client connections
type ClientManager struct {
	cancelFunc        context.CancelFunc
	clientPtrMap      map[*ClientConnection]bool
	broadcastMessages []*BroadcastFeedMessage
	cacheSize         int32
	pool              *gopool.Pool
	poller            netpoll.Poller
	broadcastChan     chan *BroadcastMessage
	clientAdd         chan *ClientConnection
	clientRemove      chan *ClientConnection
	accConfirm        chan common.Hash
	settings          ClientManagerSettings
}

func NewClientManager(pool *gopool.Pool, poller netpoll.Poller, settings ClientManagerSettings) *ClientManager {
	return &ClientManager{
		poller:        poller,
		pool:          pool,
		clientPtrMap:  make(map[*ClientConnection]bool),
		broadcastChan: make(chan *BroadcastMessage, 1),
		clientAdd:     make(chan *ClientConnection, 128),
		clientRemove:  make(chan *ClientConnection, 128),
		accConfirm:    make(chan common.Hash, 128),
		settings:      settings,
	}
}

func (cm *ClientManager) registerClient(clientConnection *ClientConnection) error {
	if len(cm.broadcastMessages) > 0 {
		// send the newly connected client all the messages we've got...
		bm := BroadcastMessage{
			Version:  1,
			Messages: cm.broadcastMessages,
		}

		err := clientConnection.write(bm)
		if err != nil {
			return err
		}
	}

	cm.clientPtrMap[clientConnection] = true

	clientConnection.Start()

	return nil
}

// Register registers new connection as a Client.
func (cm *ClientManager) Register(ctx context.Context, conn net.Conn, desc *netpoll.Desc) *ClientConnection {
	clientConnection := NewClientConnection(ctx, conn, desc, cm)

	cm.clientAdd <- clientConnection

	return clientConnection
}

// removeAll removes all clients after main ClientManager thread exits
func (cm *ClientManager) removeAll() {
	// Make copy of list because the remove() affects the client list held by the instance
	clientList := make([]*ClientConnection, len(cm.clientPtrMap))
	var i uint64
	for client := range cm.clientPtrMap {
		clientList[i] = client
		i++
	}

	// Only called after main ClientManager thread exits, so use removeClient directly
	for _, client := range clientList {
		cm.removeClient(client)
	}
}

func (cm *ClientManager) removeClient(clientConnection *ClientConnection) bool {
	if !cm.clientPtrMap[clientConnection] {
		return false
	}

	clientConnection.Stop()

	err := cm.poller.Stop(clientConnection.desc)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to stop poller")
		return false
	}

	err = clientConnection.conn.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to close client connection")
	}

	delete(cm.clientPtrMap, clientConnection)

	// TODO: properly close file descriptor
	//_ = clientConnection.desc.Close()

	return true
}

func (cm *ClientManager) Remove(clientConnection *ClientConnection) {
	cm.clientRemove <- clientConnection
}

// ConfirmedAccumulator clears out entry that matches accumulator and all older entries
func (cm *ClientManager) doConfirmedAccumulator(accumulator common.Hash) {
	for i, msg := range cm.broadcastMessages {
		if msg.FeedItem.BatchItem.Accumulator == accumulator {
			// This entry was confirmed, so this and all previous messages should be removed from cache
			unconfirmedIndex := i + 1
			if unconfirmedIndex >= len(cm.broadcastMessages) {
				//  Nothing newer, so clear entire cache
				cm.broadcastMessages = cm.broadcastMessages[:0]
			} else {
				cm.broadcastMessages = cm.broadcastMessages[unconfirmedIndex:]
			}
			break
		}
	}

	bm := BroadcastMessage{
		Version: 1,
		ConfirmedAccumulator: ConfirmedAccumulator{
			IsConfirmed: true,
			Accumulator: accumulator,
		},
	}

	cm.broadcastChan <- &bm
}

func (cm *ClientManager) confirmedAccumulator(accumulator common.Hash) {
	cm.accConfirm <- accumulator
}

// Broadcast sends batch item to all clients.
func (cm *ClientManager) Broadcast(prevAcc common.Hash, batchItem inbox.SequencerBatchItem, signature []byte) error {
	logger.Debug().Hex("acc", batchItem.Accumulator.Bytes()).Msg("sending batch Item")

	var broadcastMessages []*BroadcastFeedMessage

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

	cm.broadcastChan <- &bm

	return nil
}

func (cm *ClientManager) doBroadcast(bm *BroadcastMessage) error {
	// Don't add confirmed accumulator to cache
	if len(bm.Messages) > 0 {
		// Add to cache to send to new clients
		if len(cm.broadcastMessages) == 0 {
			// Current list is empty
			cm.broadcastMessages = append(cm.broadcastMessages, bm.Messages...)
		} else if cm.broadcastMessages[len(cm.broadcastMessages)-1].FeedItem.BatchItem.Accumulator == bm.Messages[0].FeedItem.PrevAcc {
			cm.broadcastMessages = append(cm.broadcastMessages, bm.Messages...)
		} else {
			// We need to do a re-org
			logger.Debug().Hex("acc", bm.Messages[0].FeedItem.BatchItem.Accumulator.Bytes()).Msg("broadcaster reorg")
			i := len(cm.broadcastMessages) - 1
			for ; i >= 0; i-- {
				if cm.broadcastMessages[i].FeedItem.BatchItem.Accumulator == bm.Messages[0].FeedItem.PrevAcc {
					cm.broadcastMessages = append(cm.broadcastMessages[:i+1], bm.Messages...)
					break
				}
			}

			if i == -1 {
				// All existing messages are out of date
				cm.broadcastMessages = append(cm.broadcastMessages[:0], bm.Messages...)
			}
		}
	}

	var buf bytes.Buffer
	writer := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(bm); err != nil {
		return errors.Wrap(err, "unable to encode message")
	}
	if err := writer.Flush(); err != nil {
		return errors.Wrap(err, "unable to flush message")
	}

	clientDeleteList := make([]*ClientConnection, 0, len(cm.clientPtrMap))
	for client := range cm.clientPtrMap {
		if len(client.out) == MaxSendQueue {
			// Queue for client too backed up, so delete after going through all other clients
			clientDeleteList = append(clientDeleteList, client)
		} else {
			client.out <- buf.Bytes()
		}
	}

	for _, client := range clientDeleteList {
		logger.Warn().Str("client", client.name).Msg("disconnecting client, queue too large")
		cm.Remove(client)
	}

	return nil
}

// verifyClients should be called every cm.settings.ClientPingInterval
func (cm *ClientManager) verifyClients() {
	clientConnectionCount := len(cm.clientPtrMap)
	logger.Debug().Int("feed_client_count", clientConnectionCount).Send()

	// Create list of clients to clients to remove
	deadClientList := make([]*ClientConnection, 0, clientConnectionCount)
	for client := range cm.clientPtrMap {
		diff := time.Since(client.GetLastHeard())
		if diff > cm.settings.ClientNoResponseTimeout {
			deadClientList = append(deadClientList, client)
		}
	}

	if len(deadClientList) > 0 {
		logger.Debug().Int("count", len(deadClientList)).Msg("disconnecting timed out clients")
		for _, deadClient := range deadClientList {
			cm.Remove(deadClient)
		}
	}

	// Send ping to all remaining clients
	logger.Debug().Int("count", len(cm.clientPtrMap)).Msg("pinging clients")
	for client := range cm.clientPtrMap {
		err := client.Ping()
		if err != nil {
			logger.Error().Err(err).Str("name", client.name).Msg("error pinging client")
		}
	}
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

		pingInterval := time.NewTicker(cm.settings.ClientPingInterval)
		defer pingInterval.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case clientConnection := <-cm.clientAdd:
				err := cm.registerClient(clientConnection)
				if err != nil {
					logger.Error().Err(err).Str("client", clientConnection.name).Msg("Failed to register client")
					close(clientConnection.out)
				}
			case clientConnection := <-cm.clientRemove:
				cm.removeClient(clientConnection)
			case accumulator := <-cm.accConfirm:
				cm.doConfirmedAccumulator(accumulator)
				atomic.StoreInt32(&cm.cacheSize, int32(len(cm.broadcastMessages)))
			case bm := <-cm.broadcastChan:
				err := cm.doBroadcast(bm)
				if err != nil {
					logger.Error().Err(err).Msg("Failed to do broadcast")
				}
				atomic.StoreInt32(&cm.cacheSize, int32(len(cm.broadcastMessages)))
			case <-time.After(cm.settings.ClientPingInterval / 2):
			}
			select {
			case <-pingInterval.C:
				cm.verifyClients()
			default:
			}
		}
	}()
}

func (cm *ClientManager) MessageCacheCount() int {
	return int(atomic.LoadInt32(&cm.cacheSize))
}
