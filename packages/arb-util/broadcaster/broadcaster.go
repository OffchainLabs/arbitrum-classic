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
	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
	"github.com/rs/zerolog/log"
	"net"
	"time"
)

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

type Broadcaster struct {
	// Parameters
	addr string
	maxWorkers int
	queueSize int
	ioTimeout time.Duration

	// State
	poller *netpoll.Poller
	listener *net.Listener
	acceptDesc *netpoll.Desc
}

func NewBroadcaster(addr string) (*Broadcaster, error) {
	return &Broadcaster {
		addr: addr,
		maxWorkers: 128,
		queueSize: 1,
		ioTimeout: 1 * time.Second,
	}, nil
}

func (b *Broadcaster) Start() error {
	// To get notified about incoming events from user connections.
	poller, err := netpoll.New(nil)
	if err != nil {
		return err
	}
	b.poller = &poller

	var (
		// Make pool of X size, Y sized work queue and one pre-spawned
		// goroutine.
		pool = gopool.NewPool(b.maxWorkers, b.queueSize, 1)
		chat = NewChat(pool)
	)
	handle := func(conn net.Conn) {
		// NOTE: we wrap conn here to show that ws could work with any kind of
		// io.ReadWriter.
		safeConn := deadliner{conn, b.ioTimeout}

		// Zero-copy upgrade to WebSocket connection.
		hs, err := ws.Upgrade(safeConn)
		if err != nil {
			logger.Error().Err(err).Str("connection_name", nameConn(conn)).Msg("upgrade error")
			_ = conn.Close()
			return
		}

		logger.Printf("%s: established websocket connection: %+v", nameConn(conn), hs)

		// Register incoming user in chat.
		user := chat.Register(safeConn)

		// Create netpoll event descriptor for conn.
		// We want to handle only read events of it.
		desc := netpoll.Must(netpoll.HandleRead(conn))

		// Subscribe to events about conn.
		err = poller.Start(desc, func(ev netpoll.Event) {
			// User should never send anything, so simply disconnect
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// When ReadHup or Hup received, this mean that client has
				// closed at least write end of the connection or connections
				// itself. So we want to stop receive events about such conn
				// and remove it from the chat registry.
				_ = poller.Stop(desc)
				chat.Remove(user)
				return
			}
			// Here we can read some new message from connection.
			// We can not read it right here in callback, because then we will
			// block the poller's inner loop.
			// We do not want to spawn a new goroutine to read single message.
			// But we want to reuse previously spawned goroutine.
			pool.Schedule(func() {
				if err := user.Receive(); err != nil {
					// When receive failed, we can only disconnect broken
					// connection and stop to receive events about it.
					_ = poller.Stop(desc)
					chat.Remove(user)
				}
			})
		})
	}

	listener, err := net.Listen("tcp", b.addr)
	if err != nil {
		return err
	}

	logger.Printf("websocket is listening on %s", listener.Addr().String())

	// Create netpoll descriptor for the listener.
	// We use OneShot here to manually resume events stream when we want to.
	b.acceptDesc = netpoll.Must(netpoll.HandleListener(
		listener, netpoll.EventRead|netpoll.EventOneShot,
	))

	// Subscribe to events about listener.
	err = poller.Start(b.acceptDesc, func(e netpoll.Event) {
		// accept is a channel to signal about next incoming connection Accept()
		// results.
		accept := make(chan error, 1)

		// We do not want to accept incoming connection when goroutine pool is
		// busy. So if there are no free goroutines during 1ms we want to
		// cooldown the server and do not receive connection for some short
		// time.
		err := pool.ScheduleTimeout(time.Millisecond, func() {
			conn, err := listener.Accept()
			if err != nil {
				accept <- err
				return
			}

			accept <- nil
			handle(conn)
		})
		if err == nil {
			err = <-accept
		}
		if err != nil {
			if err == gopool.ErrScheduleTimeout {
				ne, ok := err.(net.Error)
				if !ok || !ne.Temporary() {
					logger.Fatal().Err(err).Msg("accept error")
				}
			}

			delay := 5 * time.Millisecond
			logger.Err(err).Dur("delay", delay).Msg("accept error")
			time.Sleep(delay)
		}

		err = poller.Resume(b.acceptDesc)
	})

	return nil
}

func (b *Broadcaster) Stop() {
	_ = (*b.poller).Stop(b.acceptDesc)
	_ = b.acceptDesc.Close()
}

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " > " + conn.RemoteAddr().String()
}

// deadliner is a wrapper around net.Conn that sets read/write deadlines before
// every Read() or Write() call.
type deadliner struct {
	net.Conn
	t time.Duration
}

func (d deadliner) Write(p []byte) (int, error) {
	if err := d.Conn.SetWriteDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Write(p)
}

func (d deadliner) Read(p []byte) (int, error) {
	if err := d.Conn.SetReadDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Read(p)
}
