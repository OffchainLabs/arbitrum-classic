package broadcastclient

import (
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

// This sends out test broadcast messages
type TestMessageBroadcaster struct {
	broadcaster              *broadcaster.Broadcaster
	startWorkerMutex         *sync.Mutex
	messageBroadcasterWorker *time.Ticker
	count                    int
	intervalDuration         time.Duration
	workerStarted            bool
}

// create a new test message broadcaster
func NewTestMessageBroadcaster(count int, ms int) *TestMessageBroadcaster {
	tmb := &TestMessageBroadcaster{}
	tmb.startWorkerMutex = &sync.Mutex{}
	tmb.intervalDuration = time.Duration(ms) * time.Millisecond
	tmb.workerStarted = false
	tmb.count = count
	return tmb
}

// give it a client manager to broadcast on.
func (tmb *TestMessageBroadcaster) setBroadcaster(broadcaster *broadcaster.Broadcaster) {
	tmb.broadcaster = broadcaster
}

func (tmb *TestMessageBroadcaster) startWorker() {
	tmb.startWorkerMutex.Lock()
	defer tmb.startWorkerMutex.Unlock()
	if tmb.workerStarted {
		return
	}

	ticker := time.NewTicker(tmb.intervalDuration)
	messageCount := 0
	go func() {
		for t := range ticker.C {
			ib := inbox.InboxMessage{}
			ib.InboxSeqNum = big.NewInt(t.UnixNano())

			messages := []*inbox.InboxMessage{
				&ib,
			}
			tmb.broadcaster.Broadcast(messages)
			messageCount++
			if messageCount == tmb.count {
				ticker.Stop()
				return
			}
		}
	}()

	tmb.messageBroadcasterWorker = ticker
	tmb.workerStarted = true
}

func (tmb *TestMessageBroadcaster) stopWorker() {
	if tmb.messageBroadcasterWorker != nil {
		tmb.messageBroadcasterWorker.Stop()
		tmb.workerStarted = false
	}
}

func TestBroadcaster(t *testing.T) {
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9642",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	b := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b.Start()
	if err != nil {
		t.Fatal(err)
	}

	// this will send test messages to the clients at an interval
	tmb := NewTestMessageBroadcaster(10, 100)
	tmb.setBroadcaster(b)

	var wg sync.WaitGroup
	// for i := 0; i < 2; i++ {
	wg.Add(1)
	go makeBroadcastClient(t, 10, &wg)
	// }

	tmb.startWorker()
	wg.Wait()
	tmb.stopWorker()
	b.Stop()
}

func makeBroadcastClient(t *testing.T, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9642/", nil)
	defer wg.Done()
	messageCount := 0

	// connect returns
	messages, err := broadcastClient.Connect()
	if err != nil {
		t.Errorf("Can not connect: %v\n", err)
	}

	for {
		select {
		case receivedMsgs := <-messages:
			for i := range receivedMsgs.Messages {
				fmt.Printf("Received Message, Sequence Number: %v\n", &receivedMsgs.Messages[i].InboxSeqNum)
				messageCount++
				if messageCount == expectedCount {
					broadcastClient.Close()
					return
				}
			}
		}
	}

}
