/*
 * Copyright 2019, Offchain Labs, Inc.
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

package ethvalidator

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000000
)

type Client struct {
	ToClient   chan []byte
	FromClient chan []byte

	conn    *websocket.Conn
	Address common.Address
}

func NewClient(conn *websocket.Conn, address common.Address) *Client {
	return &Client{
		make(chan []byte, 128),
		make(chan []byte, 128),
		conn,
		address,
	}
}

func (c *Client) run() error {
	var rerr error
	defer func() {
		rerr = c.conn.Close()
		close(c.FromClient)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return err
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	readErrChan := make(chan error, 1)
	writeErrChan := make(chan error, 1)
	go func() {
		if err := c.writePump(readErrChan); err != nil {
			writeErrChan <- err
		}
	}()

	if err := c.readPump(writeErrChan); err != nil {
		readErrChan <- err
		return err
	}
	return rerr
}

func (c *Client) readPump(writeErrChan chan error) error {
	var rerr error
	defer func() {
		rerr = c.conn.Close()
	}()

	for {
		select {
		case err := <-writeErrChan:
			// Error writing to the client so break off this connection
			return err
		default:
		}
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return rerr
		}

		c.FromClient <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump(readErrChan chan error) error {
	var rerr error
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		rerr = c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.ToClient:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return err
			}
			if !ok {
				// The hub closed the channel.
				_ = c.conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				// There's nothing much we can do if we fail send a close message
				// Return an error if we fail to close the connection (occurs in the defer block)
				return rerr
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return err
			}
			if _, err := w.Write(message); err != nil {
				return err
			}
			if err := w.Close(); err != nil {
				return err
			}
		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return err
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return err
			}
		case <-readErrChan:
			// Error reading from the client so break off this connection
			return nil
		}
	}
}
