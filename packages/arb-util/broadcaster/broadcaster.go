package broadcaster

import (
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws-examples/src/gopool"
	"github.com/mailru/easygo/netpoll"
	"github.com/rs/zerolog/log"

	_ "net/http/pprof"
)

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

type Settings struct {
	Addr      string
	Debug     string
	Workers   int
	Queue     int
	IoTimeout time.Duration
}

type Broadcaster struct {
	clientManager         *ClientManager
	startBroadcasterMutex *sync.Mutex
	broadcasterStarted    bool
	settings              Settings
	poller                *netpoll.Poller
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
	// Initialize netpoll instance. We will use it to be noticed about incoming
	// events from listener of client connections.
	poller, err := netpoll.New(nil)
	b.poller = &poller
	if err != nil {
		logger.Error().Err(err).Msg("unable to initialize netpoll")
		return err
	}

	var (
		// Make pool of X size, Y sized work queue and one pre-spawned
		// goroutine.
		pool                   = gopool.NewPool(b.settings.Workers, b.settings.Queue, 1)
		clientManager          = NewClientManager(pool)
		testMessageBroadcaster = NewTestMessageBroadcaster()
	)

	b.clientManager = clientManager // maintain the pointer in this instance.

	////// TESTING
	// set the clientManager for testing here...
	testMessageBroadcaster.setClientManager(clientManager)

	// handle is a new incoming connection handler.
	// It upgrades TCP connection to WebSocket, registers netpoll listener on
	// it and stores it as a Client connection in ClientManager instance.
	//
	// We will call it below within accept() loop.
	handle := func(conn net.Conn) {

		// NOTE: we wrap conn here to show that ws could work with any kind of
		// io.ReadWriter.
		safeConn := deadliner{conn, b.settings.IoTimeout}

		// Zero-copy upgrade to WebSocket connection.
		hs, err := ws.Upgrade(safeConn)
		if err != nil {
			logger.Warn().Err(err).Str("connection_name", nameConn(conn)).Msg("upgrade error")
			_ = conn.Close()
			return
		}

		logger.Info().Msgf("%s: established websocket connection: %+v", nameConn(conn), hs)

		// Register incoming client in clientManager.
		client := clientManager.Register(safeConn)

		// Create netpoll event descriptor for conn.
		// We want to handle only read events of it.
		desc := netpoll.Must(netpoll.HandleRead(conn))

		//testMessageBroadcaster.startWorker()

		// Subscribe to events about conn.
		err = poller.Start(desc, func(ev netpoll.Event) {
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// When ReadHup or Hup received, this mean that client has
				// closed at least write end of the connection or connections
				// itself. So we want to stop receive events about such conn
				// and remove it from the clientManager registry.
				err = poller.Stop(desc)
				if err != nil {
					logger.Warn().Err(err).Msg("error stopping poller")
				}
				clientManager.Remove(client)
				return
			}
			// Here we can read some new message from connection.
			// We can not read it right here in callback, because then we will
			// block the poller's inner loop.
			// We do not want to spawn a new goroutine to read single message.
			// But we want to reuse previously spawned goroutine.
			pool.Schedule(func() {
				if err := client.Receive(); err != nil {
					// When receive failed, we can only disconnect broken
					// connection and stop to receive events about it.
					_ = poller.Stop(desc)
					clientManager.Remove(client)
				}
			})
		})
		if err != nil {
			logger.Warn().Err(err).Msg("error starting poller")
		}
	}

	// Create incoming connections listener.
	ln, err := net.Listen("tcp", b.settings.Addr)
	if err != nil {
		logger.Error().Err(err).Msg("error calling net.Listen")
		return err
	}

	logger.Info().Msgf("websocket is listening on %s", ln.Addr().String())

	// Create netpoll descriptor for the listener.
	// We use OneShot here to manually resume events stream when we want to.
	acceptDesc := netpoll.Must(netpoll.HandleListener(
		ln, netpoll.EventRead|netpoll.EventOneShot,
	))
	b.acceptDesc = acceptDesc

	// accept is a channel to signal about next incoming connection Accept()
	// results.
	accept := make(chan error, 1)

	// Subscribe to events about listener.
	err = poller.Start(acceptDesc, func(e netpoll.Event) {
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
			if err != gopool.ErrScheduleTimeout {
				goto cooldown
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				goto cooldown
			}

			logger.Fatal().Err(err).Msg("error in poller.Start")

		cooldown:
			delay := 5 * time.Millisecond
			logger.Info().Msgf("accept error: %v; retrying in %s", err, delay)
			time.Sleep(delay)
		}

		err = poller.Resume(acceptDesc)
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
