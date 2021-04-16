package broadcaster

import (
	"bytes"
	"encoding/json"
	"math/big"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/gobwas/ws/wsutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

// ClientManager manages client connections
type ClientManager struct {
	mu                sync.RWMutex
	seq               uint
	clientList        []*ClientConnection
	clientMap         map[string]*ClientConnection
	broadcastMessages []*BroadcastInboxMessage
	pool              *gopool.Pool
	out               chan []byte
}

func NewClientManager(pool *gopool.Pool) *ClientManager {
	clientManager := &ClientManager{
		pool:      pool,
		clientMap: make(map[string]*ClientConnection),
		out:       make(chan []byte, 1),
	}

	go clientManager.writer()

	return clientManager
}

// Register registers new connection as a Client.
func (cm *ClientManager) Register(conn net.Conn) *ClientConnection {
	clientConnection := &ClientConnection{
		clientManager: cm,
		conn:          conn,
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
			bm := BroadcastMessage{}
			bm.Messages = cm.broadcastMessages

			_ = clientConnection.write(bm)
		}

		cm.seq++
	}

	return clientConnection
}

// Remove removes client from stream.
func (cm *ClientManager) Remove(clientConnection *ClientConnection) {
	cm.mu.Lock()
	cm.remove(clientConnection)
	cm.mu.Unlock()
}

// ** Need to add a Mutex in here
// SyncSequence clears out everything prior
func (cm *ClientManager) syncSequence(fromSequenceNumber *big.Int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	broadcastMessages := make([]*BroadcastInboxMessage, 0)
	for i := range cm.broadcastMessages {
		if cm.broadcastMessages[i].SeqNum.CmpAbs(fromSequenceNumber) > 0 {
			broadcastMessages = append(broadcastMessages, cm.broadcastMessages[i])
		}
	}
	cm.broadcastMessages = broadcastMessages
}

// Broadcast sends message to all clients.
func (cm *ClientManager) Broadcast(beforeAccumulator *big.Int, inboxMessage []byte, signature *big.Int) error {
	var buf bytes.Buffer

	w := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	var broadcastMessages []*BroadcastInboxMessage

	ibMsg := BroadcastInboxMessage{}
	ibMsg.BeforeAccumulator = beforeAccumulator
	ibMsg.InboxMessage = inboxMessage
	ibMsg.Signature = signature
	ibMsg.SeqNum = inbox.GetSequenceNumber(inboxMessage)

	broadcastMessages = append(broadcastMessages, &ibMsg)

	// also add this to our global list for broadcasting to clients when connecting
	{
		cm.mu.Lock()
		defer cm.mu.Unlock()
		cm.broadcastMessages = append(cm.broadcastMessages, &ibMsg)
	}

	bm := BroadcastMessage{}
	bm.Messages = broadcastMessages

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

// mutex must be held.
func (cm *ClientManager) remove(clientConnection *ClientConnection) bool {
	if _, has := cm.clientMap[clientConnection.name]; !has {
		return false
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
