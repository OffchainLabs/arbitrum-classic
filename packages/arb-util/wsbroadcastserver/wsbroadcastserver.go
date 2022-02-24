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
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

var logger = arblog.Logger.With().Str("component", "wsbroadcastserver").Logger()

type WSBroadcastServer struct {
	startMutex    *sync.Mutex
	poller        netpoll.Poller
	acceptDesc    *netpoll.Desc
	listener      net.Listener
	settings      configuration.FeedOutput
	started       bool
	clientManager *ClientManager
	catchupBuffer CatchupBuffer
}

func NewWSBroadcastServer(settings configuration.FeedOutput, catchupBuffer CatchupBuffer) *WSBroadcastServer {
	return &WSBroadcastServer{
		startMutex:    &sync.Mutex{},
		settings:      settings,
		started:       false,
		catchupBuffer: catchupBuffer,
	}
}

func (s *WSBroadcastServer) Start(ctx context.Context) error {
	s.startMutex.Lock()
	defer s.startMutex.Unlock()
	if s.started {
		return nil
	}

	var err error
	s.poller, err = netpoll.New(nil)
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize netpoll for monitoring client connection events")
		return err
	}

	// Make pool of X size, Y sized work queue and one pre-spawned
	// goroutine.
	var clientManager = NewClientManager(s.poller, s.settings, s.catchupBuffer)
	clientManager.Start(ctx)

	s.clientManager = clientManager // maintain the pointer in this instance... used for testing

	// handle incoming connection requests.
	// It upgrades TCP connection to WebSocket, registers netpoll listener on
	// it and stores it as a Client connection in ClientManager instance.
	//
	// Called below in accept() loop.
	handle := func(conn net.Conn) {

		safeConn := deadliner{conn, s.settings.IOTimeout}

		// Zero-copy upgrade to WebSocket connection.
		hs, err := ws.Upgrade(safeConn)
		if err != nil {
			logger.Warn().Err(err).Str("connection_name", nameConn(safeConn)).Msg("upgrade error")
			_ = safeConn.Close()
			return
		}

		logger.
			Info().
			Str("connection-name", nameConn(safeConn)).
			Msgf("established websocket connection: %+v", hs)

		// Create netpoll event descriptor to handle only read events.
		desc, err := netpoll.HandleRead(conn)
		if err != nil {
			logger.Warn().Err(err).Str("connection_name", nameConn(conn)).Msg("error in HandleRead")
			_ = conn.Close()
			return
		}

		// Register incoming client in clientManager.
		client := clientManager.Register(safeConn, desc)

		// Subscribe to events about conn.
		err = s.poller.Start(desc, func(ev netpoll.Event) {
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// ReadHup or Hup received, means the client has close the connection
				// remove it from the clientManager registry.
				logger.Info().Str("connection_name", nameConn(safeConn)).Msg("Hup received")
				clientManager.Remove(client)
				return
			}

			if ev > 1 {
				logger.
					Info().
					Str("connection_name", nameConn(safeConn)).
					Int("event", int(ev)).
					Msg("event greater than 1 received")
			}

			// receive client messages, close on error
			clientManager.pool.Schedule(func() {
				// Ignore any messages sent from client
				if _, _, err := client.Receive(ctx, s.settings.ClientTimeout); err != nil {
					logger.Warn().Err(err).Str("connection_name", nameConn(safeConn)).Msg("receive error")
					clientManager.Remove(client)
					return
				}
			})
		})

		if err != nil {
			logger.Warn().Err(err).Msg("error starting client connection poller")
		}
	}

	// Create tcp server for relay connections
	ln, err := net.Listen("tcp", s.settings.Addr+":"+s.settings.Port)
	if err != nil {
		logger.Error().Err(err).Msg("error calling net.Listen")
		return err
	}

	s.listener = ln

	logger.Info().Str("address", ln.Addr().String()).Msg("arbitrum websocket broadcast server is listening")

	// Create netpoll descriptor for the listener.
	// We use OneShot here to manually resume events stream when we want to.
	acceptDesc, err := netpoll.HandleListener(ln, netpoll.EventRead|netpoll.EventOneShot)
	if err != nil {
		logger.Error().Err(err).Msg("error calling HandleListener")
		return err
	}
	s.acceptDesc = acceptDesc

	// accept is a channel to signal about next incoming connection Accept()
	// results.
	accept := make(chan error, 1)

	// Subscribe to events about listener.
	err = s.poller.Start(acceptDesc, func(e netpoll.Event) {
		// We do not want to accept incoming connection when goroutine pool is
		// busy. So if there are no free goroutines during 1ms we want to
		// cooldown the server and do not receive connection for some short
		// time.
		err := clientManager.pool.ScheduleTimeout(time.Millisecond, func() {
			conn, err := ln.Accept()
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
				netError, ok := err.(net.Error)
				if !ok || !netError.Temporary() {
					logger.Error().Err(err).Msg("error in poller.Start")
					return
				}
				if strings.Contains(err.Error(), "file descriptor was not registered") {
					logger.Info().Err(err).Msg("poller exiting")
					return
				}
			}

			// cooldown
			delay := 5 * time.Millisecond
			logger.Info().Err(err).Str("delay", delay.String()).Msg("accept error")
			time.Sleep(delay)
		}

		err = s.poller.Resume(acceptDesc)
		if err != nil {
			logger.Warn().Err(err).Msg("error in poller.Resume")
		}
	})
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Start")
	}

	s.started = true

	return nil
}

func (s *WSBroadcastServer) Stop() {
	err := s.listener.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("error in listener.Close")
	}

	err = s.poller.Stop(s.acceptDesc)
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Stop")
	}

	err = s.acceptDesc.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("error in acceptDesc.Close")
	}

	s.clientManager.Stop()
	s.started = false
}

// Broadcast sends batch item to all clients.
func (s *WSBroadcastServer) Broadcast(bm interface{}) {
	s.clientManager.Broadcast(bm)
}

func (s *WSBroadcastServer) ClientCount() int32 {
	return s.clientManager.ClientCount()
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

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " > " + conn.RemoteAddr().String()
}
