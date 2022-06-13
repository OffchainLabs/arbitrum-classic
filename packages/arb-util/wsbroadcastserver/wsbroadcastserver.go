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
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
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

func NewWSBroadcastServer(settings *configuration.FeedOutput, catchupBuffer CatchupBuffer) *WSBroadcastServer {
	return &WSBroadcastServer{
		startMutex:    &sync.Mutex{},
		settings:      *settings,
		started:       false,
		catchupBuffer: catchupBuffer,
	}
}

func (s *WSBroadcastServer) Start(ctx context.Context) (chan error, error) {
	s.startMutex.Lock()
	defer s.startMutex.Unlock()
	if s.started {
		return nil, errors.New("broadcast server already started")
	}

	var err error
	s.poller, err = netpoll.New(nil)
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize netpoll for monitoring client connection events")
		return nil, err
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
		return nil, err
	}

	s.listener = ln

	logger.Info().Str("address", ln.Addr().String()).Msg("arbitrum websocket broadcast server is listening")

	// Create netpoll descriptor for the listener.
	// We use OneShot here to synchronously manage the rate that new connections are accepted
	acceptDesc, err := netpoll.HandleListener(ln, netpoll.EventRead|netpoll.EventOneShot)
	if err != nil {
		logger.Error().Err(err).Msg("error calling HandleListener")
		return nil, err
	}
	s.acceptDesc = acceptDesc

	broadcasterErrChan := make(chan error, 10)
	acceptErrChan := make(chan error, 10)

	// Subscribe to events about listener.
	err = s.poller.Start(acceptDesc, func(e netpoll.Event) {
		select {
		case <-ctx.Done():
			broadcasterErrChan <- errors.New("broadcaster poller context done")
			return
		default:
		}
		// We do not want to accept incoming connection when goroutine pool is
		// busy. So if there are no free goroutines during 1ms we want to
		// cooldown the server and do not receive connection for some short
		// time.
		err := clientManager.pool.ScheduleTimeout(time.Millisecond, func() {
			conn, err := ln.Accept()
			if err != nil {
				acceptErrChan <- err
				return
			}

			acceptErrChan <- nil
			handle(conn)
		})
		if err == nil {
			err = <-acceptErrChan
		}
		if err != nil {
			if err == gopool.ErrScheduleTimeout {
				var netError net.Error
				isNetError := errors.As(err, &netError)
				if strings.Contains(err.Error(), "file descriptor was not registered") {
					logger.Error().Err(err).Msg("broadcast poller unable to register file descriptor")
				} else if !isNetError || !netError.Timeout() {
					logger.Error().Err(err).Msg("broadcast poller error")
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
			broadcasterErrChan <- err
			return
		}
	})
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Start")
		return nil, err
	}

	s.started = true

	return broadcasterErrChan, nil
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
