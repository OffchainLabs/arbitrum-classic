package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestRelayRebroadcasts(t *testing.T) {
	ctx := context.Background()

	// Start up an arbitrum sequencer broadcaster
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9742",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	bc := broadcaster.NewBroadcaster(broadcasterSettings)

	err := bc.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer bc.Stop()

	relaySettings := broadcaster.Settings{
		Addr:      ":7429",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	// Start up an arbitrum sequencer relay
	arbRelay := NewArbRelay("ws://127.0.0.1:9742/", relaySettings)
	arbRelay.Start(ctx, false)
	defer arbRelay.Stop()

	// Create RandomMessageGenerator
	tmb := broadcaster.NewRandomMessageGenerator(10, 100)
	tmb.SetBroadcaster(bc)

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go makeRelayClient(t, 10, &wg)
	}

	tmb.StartWorker()
	wg.Wait()
	tmb.StopWorker()
}

func makeRelayClient(t *testing.T, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := broadcastclient.NewBroadcastClient("ws://127.0.0.1:7429/", nil)
	defer wg.Done()
	messageCount := 0

	// connect returns
	messageReceiver, err := broadcastClient.Connect()
	if err != nil {
		t.Errorf("Can not connect: %v\n", err)
	}

	accList := make(chan common.Hash)

	broadcastClient.SetConfirmedAccumulatorListner(accList)

	for {
		select {
		case receivedMsg := <-messageReceiver:
			t.Logf("Received Message, Sequence Message: %v\n", receivedMsg.FeedItem.BatchItem.SequencerMessage)
			messageCount++

			if messageCount == expectedCount {
				broadcastClient.Close()
				return
			}
		case confirmedAccumulator := <-accList:
			t.Logf("Received confirmedAccumulator, Sequence Message: %v\n", confirmedAccumulator.ShortString())
		}
	}
}
