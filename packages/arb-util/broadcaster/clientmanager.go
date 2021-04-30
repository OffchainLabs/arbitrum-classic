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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/rand"
	"net"
	"strconv"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

// ClientManager manages client connections
type ClientManager struct {
	mu                sync.RWMutex
	seq               uint
	clientPtrMap      map[*ClientConnection]bool
	broadcastMessages []*BroadcastFeedMessage
	pool              *gopool.Pool
	poller            netpoll.Poller
	out               chan []byte
}

func NewClientManager(pool *gopool.Pool, poller netpoll.Poller) *ClientManager {
	return &ClientManager{
		poller:       poller,
		pool:         pool,
		clientPtrMap: make(map[*ClientConnection]bool),
		out:          make(chan []byte, 1),
	}
}

// Register registers new connection as a Client.
func (cm *ClientManager) Register(conn net.Conn, desc *netpoll.Desc) *ClientConnection {
	clientConnection := &ClientConnection{
		clientManager: cm,
		conn:          conn,
		desc:          desc,
	}

	{
		cm.mu.Lock()
		defer cm.mu.Unlock()

		clientConnection.id = cm.seq
		clientConnection.name = conn.RemoteAddr().String() + strconv.Itoa(rand.Intn(10))

		cm.clientPtrMap[clientConnection] = true

		if len(cm.broadcastMessages) > 0 {
			// send the newly connected client all the messages we've got...
			bm := BroadcastMessage{Version: 1, Messages: cm.broadcastMessages}

			_ = clientConnection.write(bm)
		}

		cm.seq++
	}

	return clientConnection
}

// RemoveAll removes all clients.
func (cm *ClientManager) RemoveAll() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Make copy of list because the remove() affects the client list held by the instance
	clientList := make([]*ClientConnection, len(cm.clientPtrMap))
	var i uint64
	for client := range cm.clientPtrMap {
		clientList[i] = client
		i++
	}

	// Only called by destructor, so keep mutex while looping through client list
	for i := range clientList {
		cm.remove(clientList[i])
	}
}

// Remove removes client from stream.
func (cm *ClientManager) Remove(clientConnection *ClientConnection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.remove(clientConnection)
}

// ConfirmedAccumulator clears out entry that matches accumulator and all older entries
func (cm *ClientManager) confirmedAccumulator(accumulator common.Hash) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for i, msg := range cm.broadcastMessages {
		if msg.FeedItem.BatchItem.Accumulator == accumulator {
			// This entry was confirmed, so this and all previous messages should be removed from cache
			unconfirmedIndex := i + 1
			if unconfirmedIndex >= len(cm.broadcastMessages) {
				//  Nothing newer, so clear entire cache
				cm.broadcastMessages = nil
			} else {
				cm.broadcastMessages = cm.broadcastMessages[unconfirmedIndex:]
			}
			break
		}
	}

	bm := BroadcastMessage{Version: 1}
	bm.ConfirmedAccumulator = ConfirmedAccumulator{
		IsConfirmed: true,
		Accumulator: accumulator,
	}

	var buf bytes.Buffer
	w := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(bm); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	cm.out <- buf.Bytes()

	return nil
}

// Broadcast sends message to all clients.
func (cm *ClientManager) Broadcast(prevAcc common.Hash, batchItem inbox.SequencerBatchItem, signature []byte) error {
	var buf bytes.Buffer

	w := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	var broadcastMessages []*BroadcastFeedMessage

	msg := BroadcastFeedMessage{
		FeedItem:  SequencerFeedItem{BatchItem: batchItem, PrevAcc: prevAcc},
		Signature: signature,
	}

	broadcastMessages = append(broadcastMessages, &msg)

	// also add this to our global list for broadcasting to clients when connecting
	{
		cm.mu.Lock()
		defer cm.mu.Unlock()

		if len(cm.broadcastMessages) == 0 {
			cm.broadcastMessages = append(cm.broadcastMessages, &msg)
		} else if cm.broadcastMessages[len(cm.broadcastMessages)-1].FeedItem.BatchItem.Accumulator == prevAcc {
			cm.broadcastMessages = append(cm.broadcastMessages, &msg)
		} else {
			// We need to do a re-org
			i := len(cm.broadcastMessages) - 1
			for ; i >= 0; i-- {
				if cm.broadcastMessages[i].FeedItem.BatchItem.Accumulator == prevAcc {
					broadcastMessages := cm.broadcastMessages[:i+1]
					cm.broadcastMessages = append(broadcastMessages, &msg)
					break
				}
			}

			if i == -1 { // didn't even find the previous accumulator... start from here.
				cm.broadcastMessages = append(cm.broadcastMessages, &msg)
			}
		}

	}

	bm := BroadcastMessage{Version: 1, Messages: broadcastMessages}

	if err := encoder.Encode(bm); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	cm.out <- buf.Bytes()

	return nil
}

// startWriter starts thread to write broadcast messages from cm.out channel.
func (cm *ClientManager) startWriter(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-cm.out:
				cm.mu.RLock()
				// Copy list so data can be written to each client without lock held
				clientList := make([]*ClientConnection, len(cm.clientPtrMap))
				var i uint64
				for client := range cm.clientPtrMap {
					clientList[i] = client
					i++
				}
				cm.mu.RUnlock()

				for _, c := range clientList {
					c := c // For closure.
					cm.pool.Schedule(func() {
						err := c.writeRaw(data)
						if err != nil {
							logger.Warn().Err(err).Msg("error with writeRaw")
						}
					})
				}
			}
		}
	}()
}

// mutex must be held before calling
func (cm *ClientManager) remove(clientConnection *ClientConnection) bool {
	if !cm.clientPtrMap[clientConnection] {
		return false
	}

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
