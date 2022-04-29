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
	"context"
	"encoding/json"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
)

// MaxSendQueue is the maximum number of items in a clients out channel before client gets disconnected.
// If set too low, a burst of items will cause all clients to be disconnected
const MaxSendQueue = 1000

// ClientConnection represents client connection.
type ClientConnection struct {
	ioMutex sync.Mutex
	conn    net.Conn

	desc          *netpoll.Desc
	Name          string
	clientManager *ClientManager

	lastHeardUnix int64
	cancelFunc    context.CancelFunc
	out           chan []byte
}

func NewClientConnection(conn net.Conn, desc *netpoll.Desc, clientManager *ClientManager) *ClientConnection {
	return &ClientConnection{
		conn:          conn,
		desc:          desc,
		Name:          conn.RemoteAddr().String() + strconv.Itoa(rand.Intn(10)),
		clientManager: clientManager,
		lastHeardUnix: time.Now().Unix(),
		out:           make(chan []byte, MaxSendQueue),
	}
}

func (cc *ClientConnection) Start(parentCtx context.Context) {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	cc.cancelFunc = cancelFunc

	go func() {
		defer cc.cancelFunc()
		defer close(cc.out)
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-cc.out:
				err := cc.writeRaw(data)
				if err != nil {
					logWarn(err, "error writing data to client")
					cc.clientManager.Remove(cc)
					for {
						// Consume and ignore channel data until client properly stopped to prevent deadlock
						select {
						case <-ctx.Done():
							return
						case <-cc.out:
						}
					}
				}
			}
		}
	}()
}

func (cc *ClientConnection) Stop() {
	if cc.cancelFunc != nil {
		cc.cancelFunc()
	} else {
		// If client connection never started, need to close channel
		close(cc.out)
	}
}

func (cc *ClientConnection) GetLastHeard() time.Time {
	return time.Unix(atomic.LoadInt64(&cc.lastHeardUnix), 0)
}

// Receive reads next message from client's underlying connection.
// It blocks until full message received.
func (cc *ClientConnection) Receive(ctx context.Context, timeout time.Duration) ([]byte, ws.OpCode, error) {
	msg, op, err := cc.readRequest(ctx, timeout)
	if err != nil {
		_ = cc.conn.Close()
		return nil, op, err
	}

	return msg, op, err
}

// readRequests reads json-rpc request from connection.
func (cc *ClientConnection) readRequest(ctx context.Context, timeout time.Duration) ([]byte, ws.OpCode, error) {
	cc.ioMutex.Lock()
	defer cc.ioMutex.Unlock()

	atomic.StoreInt64(&cc.lastHeardUnix, time.Now().Unix())

	return ReadData(ctx, cc.conn, nil, timeout, ws.StateServerSide)
}

func (cc *ClientConnection) Write(x interface{}) error {
	writer := wsutil.NewWriter(cc.conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(writer)

	cc.ioMutex.Lock()
	defer cc.ioMutex.Unlock()

	if err := encoder.Encode(x); err != nil {
		return err
	}

	return writer.Flush()
}

func (cc *ClientConnection) writeRaw(p []byte) error {
	cc.ioMutex.Lock()
	defer cc.ioMutex.Unlock()

	_, err := cc.conn.Write(p)

	return err
}

func (cc *ClientConnection) Ping() error {
	cc.ioMutex.Lock()
	defer cc.ioMutex.Unlock()
	_, err := cc.conn.Write(ws.CompiledPing)
	if err != nil {
		return err
	}

	return nil
}
