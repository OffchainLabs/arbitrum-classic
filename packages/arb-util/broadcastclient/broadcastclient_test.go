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
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9742/", nil)
	defer wg.Done()
	messageCount := 0
	ctx := context.Background()

	// connect returns
	messageReceiver := broadcastClient.Connect(ctx)

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

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil)

	// connect returns
	_ = broadcastClient.Connect(ctx)
	time.Sleep(500 * time.Millisecond)

	connectionCount := b.ClientConnectionCount()
	if connectionCount != 1 {
		t.Errorf("Client Connection Count error %v\n", connectionCount)
	}

	broadcastClient.Close()
	time.Sleep(3 * time.Second)

	connectionCount = b.ClientConnectionCount()
	if connectionCount != 0 {
		t.Errorf("Client Connection Count error %v\n", connectionCount)
	}

}

// with the current functionality,
// there is no way for the client to know that
// it's connection has been closed by the server
// something needs to be built in the client
// where by if it doesn't hear from the server
// in 10 seconds, it should re-initiate the connection

func TestBroadcastClientReconnectsOnServerDisconnect(t *testing.T) {
	t.Skip("currently broken")
	ctx := context.Background()

	broadcasterSettings := broadcaster.Settings{
		Addr:                    ":9743",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      5 * time.Second,
		ClientNoResponseTimeout: 15 * time.Second,
	}

	b1 := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b1.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil)

	// connect returns
	_ = broadcastClient.Connect(ctx)

	b1.Stop()

	time.Sleep(1000 * time.Millisecond)

	b2 := broadcaster.NewBroadcaster(broadcasterSettings)
	err = b2.Start(ctx)
	if err != nil {
		t.Fatal("error restarting broadcaster")
	}

	for broadcastClient.retrying == true {
		time.Sleep(1000 * time.Millisecond)
	}

	if broadcastClient.RetryCount <= 0 {
		t.Error("Should have had some retry counts")
	}
}
