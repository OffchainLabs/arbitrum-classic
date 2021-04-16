package broadcastclient

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestBroadCastClientConnectsAndReceivesSequences(t *testing.T) {
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9742",
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
	tmb := broadcaster.NewRandomMessageGenerator(10, 100)
	tmb.SetBroadcaster(b)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go makeBroadcastClient(t, 10, &wg)
	}

	tmb.StartWorker()
	wg.Wait()
	tmb.StopWorker()
}

func makeBroadcastClient(t *testing.T, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9742/", nil)
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
				fmt.Printf("Received Message, Sequence Number: %v\n", inbox.GetSequenceNumber(receivedMsgs.Messages[i].InboxMessage))
				messageCount++
				if messageCount == expectedCount {
					broadcastClient.Close()
					return
				}
			}
		}
	}

}
