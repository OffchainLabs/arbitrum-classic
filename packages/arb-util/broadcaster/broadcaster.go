package broadcaster

import (
	"math/big"
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

type Settings struct {
	Addr      string
	Workers   int
	Queue     int
	IoTimeout time.Duration
}

type Broadcaster struct {
	clientManager         *ClientManager
	startBroadcasterMutex *sync.Mutex
	broadcasterStarted    bool
	settings              Settings
	poller                netpoll.Poller
	acceptDesc            *netpoll.Desc
}

func NewBroadcaster(settings Settings) *Broadcaster {
	return &Broadcaster{
		startBroadcasterMutex: &sync.Mutex{},
		settings:              settings,
		broadcasterStarted:    false,
	}
}

func (b *Broadcaster) Start() error {
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
	var clientManager = NewClientManager(pool)

	b.clientManager = clientManager // maintain the pointer in this instance... used for testing

	// handle incoming connection requests.
	// It upgrades TCP connection to WebSocket, registers netpoll listener on
	// it and stores it as a Client connection in ClientManager instance.
	//
	// Called below in accept() loop.
	handle := func(conn net.Conn) {

		safeConn := deadliner{conn, b.settings.IoTimeout}

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

		// Register incoming client in clientManager.
		client := clientManager.Register(safeConn)

		// Create netpoll event descriptor to handle only read events.
		desc, err := netpoll.HandleRead(conn)
		if err != nil {
			logger.Warn().Err(err).Str("connection_name", nameConn(conn)).Msg("error in HandleRead")
			_ = conn.Close()
			return
		}

		// Subscribe to events about conn.
		err = b.poller.Start(desc, func(ev netpoll.Event) {
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// ReadHup or Hup received, means the client has close the connection
				// remove it from the clientManager registry.
				logger.Info().Str("connection_name", nameConn(safeConn)).Msg("Hup received")
				_ = b.poller.Stop(desc)
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
				if err := client.Receive(); err != nil {
					logger.Warn().Err(err).Str("connection_name", nameConn(safeConn)).Msg("receive error")
					_ = b.poller.Stop(desc)
					_ = desc.Close()
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
	ln, err := net.Listen("tcp", b.settings.Addr)
	if err != nil {
		logger.Error().Err(err).Msg("error calling net.Listen")
		return err
	}

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
					logger.Fatal().Err(err).Msg("error in poller.Start")
				}
			}

			// cooldown
			delay := 5 * time.Millisecond
			logger.Info().Msgf("accept error: %v; retrying in %s", err, delay)
			time.Sleep(delay)
		}

		err = b.poller.Resume(acceptDesc)
		if err != nil {
			logger.Warn().Err(err).Msg("error in poller.Start")
		}
	})
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Resume")
	}

	b.broadcasterStarted = true

	return nil
}

func (b *Broadcaster) Broadcast(beforeAccumulator *big.Int, inboxMessage []byte, signature *big.Int) error {
	return b.clientManager.Broadcast(beforeAccumulator, inboxMessage, signature)
}

func (b *Broadcaster) SyncSequence(fromSequenceNumber *big.Int) {
	b.clientManager.syncSequence(fromSequenceNumber)
}

func (b *Broadcaster) messageCacheCount() int {
	count := len(b.clientManager.broadcastMessages)
	return count
}

func (b *Broadcaster) Stop() {
	err := b.poller.Stop(b.acceptDesc)
	if err != nil {
		logger.Warn().Err(err).Msg("error in poller.Stop")
	}

	err = b.acceptDesc.Close()
	if err != nil {
		logger.Warn().Err(err).Msg("error in acceptDesc.Stop")
	}
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
