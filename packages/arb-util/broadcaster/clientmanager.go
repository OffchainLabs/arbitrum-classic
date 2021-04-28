package broadcaster

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

// ClientManager manages client connections
type ClientManager struct {
	mu                sync.RWMutex
	seq               uint
	clientList        []*ClientConnection
	clientMap         map[string]*ClientConnection
	broadcastMessages []*BroadcastInboxMessage
	pool              *gopool.Pool
	poller            netpoll.Poller
	out               chan []byte
}

func NewClientManager(pool *gopool.Pool, poller netpoll.Poller) *ClientManager {
	clientManager := &ClientManager{
		poller:    poller,
		pool:      pool,
		clientMap: make(map[string]*ClientConnection),
		out:       make(chan []byte, 1),
	}

	go clientManager.writer()

	return clientManager
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

		cm.clientList = append(cm.clientList, clientConnection)
		cm.clientMap[clientConnection.name] = clientConnection

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
	// the remove() affects the client list held by the instance
	clientList := make([]*ClientConnection, len(cm.clientList))
	for i := range cm.clientList {
		clientList[i] = cm.clientList[i]
	}
	for i := range clientList {
		cm.remove(clientList[i])
	}

	cm.mu.Unlock()
}

// Remove removes client from stream.
func (cm *ClientManager) Remove(clientConnection *ClientConnection) {
	cm.mu.Lock()

	cm.remove(clientConnection)
	cm.mu.Unlock()
}

// ConfirmedAccumulator clears out everything prior to finding the matching accumulator
func (cm *ClientManager) confirmedAccumulator(accumulator common.Hash) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	broadcastMessages := make([]*BroadcastInboxMessage, 0)
	accumulatorFound := false

	for i := range cm.broadcastMessages {
		accum := cm.broadcastMessages[i].FeedItem.BatchItem.Accumulator
		if accumulatorFound {
			broadcastMessages = append(broadcastMessages, cm.broadcastMessages[i])
		} else if accum.Equals(accumulator) {
			accumulatorFound = true
		}
	}

	cm.broadcastMessages = broadcastMessages

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

	var broadcastMessages []*BroadcastInboxMessage

	msg := BroadcastInboxMessage{
		FeedItem:  monitor.SequencerFeedItem{BatchItem: batchItem, PrevAcc: prevAcc},
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

// writer writes broadcast messages from cm.out channel.
func (cm *ClientManager) writer() {
	for bts := range cm.out {
		// For closure
		bts := bts
		cm.mu.RLock()
		cl := cm.clientList
		cm.mu.RUnlock()

		for _, c := range cl {
			c := c // For closure.
			cm.pool.Schedule(func() {
				err := c.writeRaw(bts)
				if err != nil {
					logger.Warn().Err(err).Msg("error with writeRaw")
				}
			})
		}
	}
}

// mutex must be held before calling
func (cm *ClientManager) remove(clientConnection *ClientConnection) bool {
	if _, has := cm.clientMap[clientConnection.name]; !has {
		return false
	}

	// doing this causes a hang.
	// cm.poller.Stop(clientConnection.desc)

	err := clientConnection.conn.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to close client connection")
	}

	delete(cm.clientMap, clientConnection.name)

	i := sort.Search(len(cm.clientList), func(i int) bool {
		return cm.clientList[i].id >= clientConnection.id
	})

	if i >= len(cm.clientList) {
		panic("stream: inconsistent state")
	}

	without := make([]*ClientConnection, len(cm.clientList)-1)
	copy(without[:i], cm.clientList[:i])
	copy(without[i:], cm.clientList[i+1:])
	cm.clientList = without

	return true
}
