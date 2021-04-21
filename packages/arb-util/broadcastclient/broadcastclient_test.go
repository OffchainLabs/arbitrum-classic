package broadcastclient

import (
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestBroadcastClientConnectsAndReceivesSequences(t *testing.T) {
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
				t.Logf("Received Message, Sequence Number: %v\n", inbox.GetSequenceNumber(receivedMsgs.Messages[i].InboxMessage))
				messageCount++
				if messageCount == expectedCount {
					broadcastClient.Close()
					return
				}
			}
		}
	}

}

func TestBroadcastClientPings(t *testing.T) {
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9743",
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
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil)

	// connect returns
	_, err = broadcastClient.Connect()
	if err != nil {
		t.Errorf("Can not connect: %v\n", err)
	}
	pong, err := broadcastClient.Ping()
	if err != nil {
		t.Errorf("Cannot send ping: %v\n", err)
	}
	p := <-pong
	if p != "pong" {
		t.Error("No response from ping")
	}
}

func TestBroadcastClientReconnectsOnServerDisconnect(t *testing.T) {
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9743",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	b1 := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b1.Start()
	if err != nil {
		t.Fatal(err)
	}

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil)

	// connect returns
	_, err = broadcastClient.Connect()
	if err != nil {
		t.Errorf("Can not connect: %v\n", err)
	}

	b1.Stop()

	time.Sleep(1000 * time.Millisecond)

	pong, err := broadcastClient.Ping()
	if err != nil {
		t.Fatal("error sending ping")
	}

	p := <-pong
	if p == "pong" {
		t.Error("Should not have received a response")
	}

	b2 := broadcaster.NewBroadcaster(broadcasterSettings)
	err = b2.Start()
	if err != nil {
		t.Fatal("error restarting broadcaster")
	}

	time.Sleep(1000 * time.Millisecond)

	pong2, err := broadcastClient.Ping()
	if err != nil {
		t.Fatal("error sending ping2")
	}

	p = <-pong2
	if p != "pong" {
		t.Error("No response from ping")
	}

	if broadcastClient.RetryCount <= 0 {
		t.Error("Should have had some retry counts")
	}
}
