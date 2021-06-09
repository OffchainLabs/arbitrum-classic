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
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
)

const MaxSendCount = 10
const MaxSendQueue = 20

// ClientConnection represents client connection.
type ClientConnection struct {
	parentCtx context.Context
	ioMutex   sync.Mutex
	conn      io.ReadWriteCloser

	desc          *netpoll.Desc
	name          string
	clientManager *ClientManager

	timeoutMutex sync.Mutex
	lastHeard    time.Time
	cancelFunc   context.CancelFunc
	out          chan []byte
}

func NewClientConnection(ctx context.Context, conn net.Conn, desc *netpoll.Desc, clientManager *ClientManager) *ClientConnection {
	return &ClientConnection{
		parentCtx:     ctx,
		conn:          conn,
		desc:          desc,
		name:          conn.RemoteAddr().String() + strconv.Itoa(rand.Intn(10)),
		clientManager: clientManager,
		lastHeard:     time.Now(),
		out:           make(chan []byte, MaxSendQueue),
	}
}

func (cc *ClientConnection) Start() {
	ctx, cancelFunc := context.WithCancel(cc.parentCtx)
	cc.cancelFunc = cancelFunc

	go func() {
		defer cancelFunc()
		defer close(cc.out)
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-cc.out:
				err := cc.writeRaw(data)
				if err != nil {
					logger.Error().Err(err).Str("client", cc.name).Msg("error writing data to client")
					cc.clientManager.Remove(cc)
					return
				}
			}
		}
	}()
}

func (cc *ClientConnection) Stop() {
	cc.cancelFunc()
}

func (cc *ClientConnection) GetLastHeard() time.Time {
	cc.timeoutMutex.Lock()
	defer cc.timeoutMutex.Unlock()

	return cc.lastHeard
}

// Receive reads next message from client's underlying connection.
// It blocks until full message received.
func (cc *ClientConnection) Receive() error {
	err := cc.readRequest()
	if err != nil {
		_ = cc.conn.Close()
		return err
	}

	return nil
}

// readRequests reads json-rpc request from connection.
func (cc *ClientConnection) readRequest() error {
	cc.ioMutex.Lock()
	defer cc.ioMutex.Unlock()

	cc.timeoutMutex.Lock()
	cc.lastHeard = time.Now()
	cc.timeoutMutex.Unlock()

	h, r, err := wsutil.NextReader(cc.conn, ws.StateServerSide)
	if err != nil {
		if h.OpCode.IsControl() {
			return wsutil.ControlFrameHandler(cc.conn, ws.StateServerSide)(h, r)
		}

		return err
	}

	return nil
}

func (cc *ClientConnection) write(x interface{}) error {
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
