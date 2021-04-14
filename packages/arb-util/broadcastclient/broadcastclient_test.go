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

// This sends out generated test broadcast messages
type MessageGenerator struct {
	broadcaster              *broadcaster.Broadcaster
	startWorkerMutex         *sync.Mutex
	messageBroadcasterWorker *time.Ticker
	count                    int
	intervalDuration         time.Duration
	workerStarted            bool
}

// create a new test message generator
func NewMessageGenerator(count int, ms int) *MessageGenerator {
	gm := &MessageGenerator{}
	gm.startWorkerMutex = &sync.Mutex{}
	gm.intervalDuration = time.Duration(ms) * time.Millisecond
	gm.workerStarted = false
	gm.count = count
	return gm
}

// give it a client manager to broadcast on.
func (gm *MessageGenerator) setBroadcaster(broadcaster *broadcaster.Broadcaster) {
	gm.broadcaster = broadcaster
}

func (gm *MessageGenerator) startWorker() {
	gm.startWorkerMutex.Lock()
	defer gm.startWorkerMutex.Unlock()
	if gm.workerStarted {
		return
	}

	ticker := time.NewTicker(gm.intervalDuration)
	messageCount := 0
	go func() {
		for t := range ticker.C {
			ib := inbox.InboxMessage{}
			ib.InboxSeqNum = big.NewInt(t.UnixNano())

			messages := []*inbox.InboxMessage{
				&ib,
			}
			gm.broadcaster.Broadcast(messages)
			messageCount++
			if messageCount == gm.count {
				ticker.Stop()
				return
			}
		}
	}()

	gm.messageBroadcasterWorker = ticker
	gm.workerStarted = true
}

func (mg *MessageGenerator) stopWorker() {
	if mg.messageBroadcasterWorker != nil {
		mg.messageBroadcasterWorker.Stop()
		mg.workerStarted = false
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
	defer b.Stop()

	// this will send test messages to the clients at an interval
	tmb := NewMessageGenerator(10, 100)
	tmb.setBroadcaster(b)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go makeBroadcastClient(t, 10, &wg)
	}

	tmb.startWorker()
	wg.Wait()
	tmb.stopWorker()
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
