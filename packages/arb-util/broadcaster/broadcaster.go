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
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

type Broadcaster struct {
	clientManager         *ClientManager
	startBroadcasterMutex *sync.Mutex
	broadcasterStarted    bool
	settings              configuration.FeedOutput
	poller                netpoll.Poller
	acceptDesc            *netpoll.Desc
	listener              net.Listener
}

func NewBroadcaster(settings configuration.FeedOutput) *Broadcaster {
	return &Broadcaster{
		startBroadcasterMutex: &sync.Mutex{},
		settings:              settings,
		broadcasterStarted:    false,
	}
}

func (b *Broadcaster) ClientCount() int32 {
	return b.clientManager.ClientCount()
}

func (b *Broadcaster) Start(ctx context.Context) error {
	b.startBroadcasterMutex.Lock()
	defer b.startBroadcasterMutex.Unlock()
	if b.broadcasterStarted {
		return nil
	}

	var err error
	b.poller, err = netpoll.New(nil)
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize netpoll for monitoring client connection events")
		return err
	}

	// Make pool of X size, Y sized work queue and one pre-spawned
	// goroutine.
	var pool = gopool.NewPool(b.settings.Workers, b.settings.Queue, 1)
	var clientManager = NewClientManager(pool, b.poller, b.settings)
	clientManager.Start(ctx)

	b.clientManager = clientManager // maintain the pointer in this instance... used for testing

	// handle incoming connection requests.
	// It upgrades TCP connection to WebSocket, registers netpoll listener on
	// it and stores it as a Client connection in ClientManager instance.
	//
	// Called below in accept() loop.
	handle := func(conn net.Conn) {

		safeConn := deadliner{conn, b.settings.IOTimeout}

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
		err = b.poller.Start(desc, func(ev netpoll.Event) {
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
			pool.Schedule(func() {
				// Ignore any messages sent from client
				if _, _, err := client.Receive(ctx, b.settings.ClientTimeout); err != nil {
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
	ln, err := net.Listen("tcp", b.settings.Addr+":"+b.settings.Port)
	if err != nil {
		logger.Error().Err(err).Msg("error calling net.Listen")
		return err
	}

	b.listener = ln

	logger.Info().Str("address", ln.Addr().String()).Msg("arbitrum websocket broadcast server is listening")

	// Create netpoll descriptor for the listener.
	// We use OneShot here to manually resume events stream when we want to.
	acceptDesc, err := netpoll.HandleListener(ln, netpoll.EventRead|netpoll.EventOneShot)
	if err != nil {
		logger.Error().Err(err).Msg("error calling HandleListener")
		return err
	}
	b.acceptDesc = acceptDesc

	// accept is a channel to signal about next incoming connection Accept()
	// results.
	accept := make(chan error, 1)

	// Subscribe to events about listener.
	err = b.poller.Start(acceptDesc, func(e netpoll.Event) {
		// We do not want to accept incoming connection when goroutine pool is
		// busy. So if there are no free goroutines during 1ms we want to
		// cooldown the server and do not receive connection for some short
		// time.
		err := pool.ScheduleTimeout(time.Millisecond, func() {
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

		err = b.poller.Resume(acceptDesc)
		if err != nil {
			logger.Warn().Err(err).Msg("error in poller.Resume")
		}
	})
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Start")
	}

	b.broadcasterStarted = true

	return nil
}

func (b *Broadcaster) BroadcastSingle(prevAcc common.Hash, batchItem inbox.SequencerBatchItem, signature []byte) error {
	return b.clientManager.Broadcast(prevAcc, batchItem, signature)
}

func (b *Broadcaster) Broadcast(prevAcc common.Hash, batchItems []inbox.SequencerBatchItem, dataSigner func([]byte) ([]byte, error)) error {
	for _, item := range batchItems {
		signature, err := dataSigner(hashing.SoliditySHA3WithPrefix(hashing.Bytes32(item.Accumulator)).Bytes())
		if err != nil {
			return err
		}

		err = b.BroadcastSingle(prevAcc, item, signature)
		if err != nil {
			return err
		}
		prevAcc = item.Accumulator
	}

	return nil
}

func (b *Broadcaster) ConfirmedAccumulator(accumulator common.Hash) {
	b.clientManager.confirmedAccumulator(accumulator)
}

func (b *Broadcaster) MessageCacheCount() int {
	return b.clientManager.MessageCacheCount()
}

func (b *Broadcaster) Stop() {
	err := b.listener.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("error in listener.Close")
	}

	err = b.poller.Stop(b.acceptDesc)
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Stop")
	}

	err = b.acceptDesc.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("error in acceptDesc.Close")
	}

	b.clientManager.Stop()
	b.broadcasterStarted = false
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
