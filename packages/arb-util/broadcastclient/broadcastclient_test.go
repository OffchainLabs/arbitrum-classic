package broadcastclient

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
)

func TestBroadcastClientConnectsAndReceivesMessages(t *testing.T) {
	ctx := context.Background()

	broadcasterSettings := broadcaster.Settings{
		Addr:                    ":9742",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      5 * time.Second,
		ClientNoResponseTimeout: 15 * time.Second,
	}

	b := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
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
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9742/", nil, 20*time.Second)
	defer wg.Done()
	messageCount := 0
	ctx := context.Background()

	// connect returns
	messageReceiver, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	accListener := broadcastClient.ConfirmedAccumulatorListener

	for {
		select {
		case receivedMsg := <-messageReceiver:
			t.Logf("Received Message, Sequence Message: %v\n", receivedMsg.FeedItem.BatchItem.SequencerMessage)
			messageCount++

			if messageCount == expectedCount {
				broadcastClient.Close()
				return
			}
		case confirmedAccumulator := <-accListener:
			t.Logf("Received confirmedAccumulator, Sequence Message: %v\n", confirmedAccumulator.ShortString())
		}
	}

}

func TestServerDisconnectsAClientIfItDoesNotRespondToPings(t *testing.T) {
	ctx := context.Background()

	broadcasterSettings := broadcaster.Settings{
		Addr:                    ":9743",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      1 * time.Second,
		ClientNoResponseTimeout: 2 * time.Second,
	}

	b := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil, 20*time.Second)

	// connect returns
	client, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	newBroadcastMessage := broadcaster.SequencedMessages()
	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.BroadcastSingle(hash1, feedItem1.BatchItem, signature1.Bytes())

	// Wait for client to receive batch to ensure it is connected
	select {
	case receivedMsg := <-client:
		t.Logf("Received Message, Sequence Message: %v\n", receivedMsg.FeedItem.BatchItem.SequencerMessage)
	case <-time.After(5 * time.Second):
		t.Fatal("Client did not receive batch item")
	}

	connectionCount := b.ClientConnectionCount()
	if connectionCount != 1 {
		t.Fatalf("Client Connection Count error %v\n", connectionCount)
	}

	broadcastClient.Close()

	// Wait for client to be disconnected from server
	disconnectTimeout := time.After(5 * time.Second)
	for {
		if b.ClientConnectionCount() == 0 {
			break
		}

		select {
		case <-disconnectTimeout:
			t.Fatal("Client did not receive batch item")
		case <-time.After(100 * time.Millisecond):
		}
	}
}

func TestBroadcastClientReconnectsOnServerDisconnect(t *testing.T) {
	ctx := context.Background()

	broadcasterSettings := broadcaster.Settings{
		Addr:                    ":9743",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      50 * time.Second,
		ClientNoResponseTimeout: 150 * time.Second,
	}

	b1 := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b1.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b1.Stop()

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil, 2*time.Second)

	// connect returns
	_, err = broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Client set to timeout connection at 2 seconds, and server set to send ping every 50 seconds,
	// so at least one timeout/reconnect should happen after 4 seconds
	time.Sleep(4 * time.Second)

	if broadcastClient.GetRetryCount() <= 0 {
		t.Error("Should have had some retry counts")
	}
}
