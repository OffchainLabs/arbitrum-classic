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
	"math/big"
	"net"
	"sync"
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
	mu                sync.RWMutex
	clientPtrMap      map[*ClientConnection]bool
	delayedMessages   []inbox.DelayedMessage
	sequencerItems    []inbox.SequencerBatchItem
	sequencedMetadata SequencedMetadata
	pool              *gopool.Pool
	poller            netpoll.Poller
	out               chan []byte
	settings          ClientManagerSettings
}

func NewClientManager(pool *gopool.Pool, poller netpoll.Poller, settings ClientManagerSettings) *ClientManager {
	return &ClientManager{
		poller:       poller,
		pool:         pool,
		clientPtrMap: make(map[*ClientConnection]bool),
		out:          make(chan []byte, 1),
		settings:     settings,
	}
}

// Register registers new connection as a Client.
func (cm *ClientManager) Register(ctx context.Context, conn net.Conn, desc *netpoll.Desc) *ClientConnection {
	clientConnection := NewClientConnection(conn, desc, cm)

	{
		cm.mu.Lock()

		cm.clientPtrMap[clientConnection] = true

		if len(cm.sequencerItems) > 0 {
			var seqMetadataPtr *SequencedMetadata
			if cm.sequencedMetadata.PrevAcc != (common.Hash{}) {
				seqMetadataPtr = &cm.sequencedMetadata
			}
			// send the newly connected client all the messages we've got...
			bm := BroadcastMessage{
				Version:           1,
				DelayedMessages:   cm.delayedMessages,
				Messages:          cm.sequencerItems,
				SequencedMetadata: seqMetadataPtr,
			}

			_ = clientConnection.write(bm)
		}

		clientConnection.Start(ctx)
		cm.mu.Unlock()
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
func (cm *ClientManager) confirmedAccumulators(accumulators [2]common.Hash) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	seqAcc := accumulators[0]
	if seqAcc != (common.Hash{}) {
		for i, msg := range cm.sequencerItems {
			if msg.Accumulator == seqAcc {
				// This entry was confirmed, so this and all previous messages should be removed from cache
				unconfirmedIndex := i + 1
				if unconfirmedIndex >= len(cm.sequencerItems) {
					//  Nothing newer, so clear entire cache
					cm.sequencerItems = cm.sequencerItems[:0]
				} else {
					cm.sequencedMetadata.PrevAcc = seqAcc
					cm.sequencerItems = cm.sequencerItems[unconfirmedIndex:]
				}
				break
			}
		}
	}

	delayedAcc := accumulators[1]
	if delayedAcc != (common.Hash{}) {
		for i, msg := range cm.delayedMessages {
			if msg.DelayedAccumulator == delayedAcc {
				// This entry was confirmed, so this and all previous messages should be removed from cache
				unconfirmedIndex := i + 1
				if unconfirmedIndex >= len(cm.delayedMessages) {
					//  Nothing newer, so clear entire cache
					cm.delayedMessages = cm.delayedMessages[:0]
				} else {
					cm.delayedMessages = cm.delayedMessages[unconfirmedIndex:]
				}
				break
			}
		}
	}

	bm := BroadcastMessage{Version: 2}
	bm.ConfirmedAccumulators = &accumulators

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
func (cm *ClientManager) Broadcast(prevAcc common.Hash, delayedMessages []inbox.DelayedMessage, batchItems []inbox.SequencerBatchItem, signature []byte) error {
	var buf bytes.Buffer

	if len(batchItems) > 0 {
		logger.Debug().Hex("acc", batchItems[len(batchItems)-1].Accumulator.Bytes()).Int("count", len(batchItems)).Msg("sending batch items")
	}

	writer := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(writer)

	// also add this to our global list for broadcasting to clients when connecting
	{
		cm.mu.Lock()

		if len(delayedMessages) > 0 {
			var cmDelayedCount *big.Int
			if len(cm.delayedMessages) == 0 {
				cmDelayedCount = big.NewInt(0)
			} else {
				cmDelayedCount = new(big.Int).Add(cm.delayedMessages[len(cm.delayedMessages)-1].DelayedSequenceNumber, big.NewInt(1))
			}
			diff := new(big.Int).Sub(cmDelayedCount, delayedMessages[0].DelayedSequenceNumber)
			if diff.Sign() != 0 {
				logger.Debug().Str("prevCount", cmDelayedCount.String()).Str("newSeqNum", delayedMessages[0].DelayedSequenceNumber.String()).Int("cacheCount", len(cm.delayedMessages)).Msg("broadcaster delayed reorg")
				if diff.Sign() > 0 && diff.IsInt64() && diff.Int64() <= int64(len(cm.delayedMessages)) {
					// Go back to where the sequence numbers match up
					newPos := len(cm.delayedMessages) - int(diff.Int64())
					cm.delayedMessages = cm.delayedMessages[:newPos]
				} else {
					// The sequence numbers don't match up anywhere, clear our cache
					cm.delayedMessages = cm.delayedMessages[:0]
				}
			}
			cm.delayedMessages = append(cm.delayedMessages, delayedMessages...)
		}

		if len(cm.sequencerItems) > 0 {
			if len(batchItems) > 0 && cm.sequencerItems[len(cm.sequencerItems)-1].Accumulator != prevAcc {
				// We need to do a re-org
				logger.Debug().Hex("prevAcc", prevAcc.Bytes()).Hex("acc", batchItems[0].Accumulator.Bytes()).Msg("broadcaster reorg")
				i := len(cm.sequencerItems) - 1
				for ; i >= 0; i-- {
					if cm.sequencerItems[i].Accumulator == prevAcc {
						cm.sequencerItems = cm.sequencerItems[:i+1]
						break
					}
				}

				if i == -1 {
					// Don't use anything in existing slice
					cm.sequencerItems = cm.sequencerItems[:0]
				}
			}
			cm.sequencedMetadata.Signature = signature
			cm.sequencedMetadata.PrevAcc = prevAcc
		}
		cm.sequencerItems = append(cm.sequencerItems, batchItems...)

		cm.mu.Unlock()
	}

	bm := BroadcastMessage{Version: 2, DelayedMessages: delayedMessages, Messages: batchItems}

	if err := encoder.Encode(bm); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	cm.out <- buf.Bytes()

	if len(batchItems) > 0 {
		logger.Debug().Int("count", len(batchItems)).Hex("acc", batchItems[len(batchItems)-1].Accumulator.Bytes()).Msg("batch items queued")
	}

	return nil
}

func (cm *ClientManager) ClientConnectionCount() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.clientPtrMap)
}

// verifyClients should be called every cm.settings.ClientPingInterval
func (cm *ClientManager) verifyClients() {
	cm.mu.RLock()
	clientConnectionCount := len(cm.clientPtrMap)
	logger.Debug().Int("feed_client_count", clientConnectionCount).Send()

	// Create list of clients to ping and clients to remove
	clientList := make([]*ClientConnection, 0, clientConnectionCount)
	deadClientList := make([]*ClientConnection, 0, clientConnectionCount)
	var deadClientCount uint64
	var aliveClientCount uint64
	for client := range cm.clientPtrMap {
		diff := time.Since(client.GetLastHeard())
		if diff > cm.settings.ClientNoResponseTimeout {
			deadClientList = append(deadClientList, client)
			deadClientCount++
		} else {
			clientList = append(clientList, client)
			aliveClientCount++
		}
	}
	cm.mu.RUnlock()

	logger.Debug().Uint64("disconnecting clients", deadClientCount).Send()
	for _, deadClient := range deadClientList {
		cm.Remove(deadClient)
	}

	logger.Debug().Uint64("pinging clients", aliveClientCount).Send()
	for _, client := range clientList {
		err := client.Ping()
		if err != nil {
			logger.Debug().Err(err).Str("name", client.name).Uint64("error pinging client", aliveClientCount).Send()
		}
	}
}

// startWriter starts thread to write broadcast messages from cm.out channel.
func (cm *ClientManager) startWriter(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-cm.out:
				cm.mu.Lock()
				clientDeleteList := make([]*ClientConnection, 0, len(cm.clientPtrMap))
				// Lock mutex while writing to channels to ensure items delivered in order
				for i := 0; i < MaxSendCount; i++ {
					for client := range cm.clientPtrMap {
						if len(client.out) == MaxSendQueue {
							// Queue for client too backed up, so delete after going through all other clients
							clientDeleteList = append(clientDeleteList, client)
						} else {
							client.out <- data
						}
					}

					select {
					case <-ctx.Done():
						return
					case data = <-cm.out:
						continue
					default:
					}
					break
				}
				cm.mu.Unlock()

				for _, client := range clientDeleteList {
					logger.Warn().Str("client", client.name).Msg("disconnecting client, queue too large")
					cm.Remove(client)
				}
			}
		}
	}()
}

// startVerifier starts thread to ping active connections and remove expired connections
func (cm *ClientManager) startVerifier(ctx context.Context) {
	go func() {
		pingInterval := time.NewTicker(cm.settings.ClientPingInterval)
		defer pingInterval.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-pingInterval.C:
				cm.verifyClients()
			}
		}
	}()
}

// mutex must be held before calling
func (cm *ClientManager) remove(clientConnection *ClientConnection) bool {
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
