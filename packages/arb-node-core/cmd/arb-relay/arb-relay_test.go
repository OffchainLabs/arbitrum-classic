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

package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestRelayRebroadcasts(t *testing.T) {
	ctx := context.Background()

	// Start up an Arbitrum sequencer broadcaster
	broadcasterSettings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9742",
		Ping:          5 * time.Second,
		ClientTimeout: 15 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	bc := broadcaster.NewBroadcaster(broadcasterSettings)

	err := bc.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer bc.Stop()

	relaySettings := configuration.Feed{
		Input: configuration.FeedInput{
			Timeout: 20 * time.Second,
			URLs:    "ws://127.0.0.1:9742",
		},
		Output: configuration.FeedOutput{
			Addr:          "0.0.0.0",
			IOTimeout:     2 * time.Second,
			Port:          "7429",
			Ping:          5 * time.Second,
			ClientTimeout: 15 * time.Second,
			Queue:         1,
			Workers:       128,
		},
	}

	// Start up an arbitrum sequencer relay
	arbRelay := NewArbRelay(relaySettings)
	_, err = arbRelay.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer arbRelay.Stop()

	// Create RandomMessageGenerator
	tmb := broadcaster.NewRandomMessageGenerator(10, 100*time.Millisecond)
	tmb.SetBroadcaster(bc)

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go makeRelayClient(t, 10, &wg)
	}

	errChan := tmb.Start(ctx)
	wg.Wait()

	select {
	case err := <-errChan:
		t.Fatal(err)
	default:
	}
}

func makeRelayClient(t *testing.T, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := broadcastclient.NewBroadcastClient("ws://127.0.0.1:7429/", nil, 20*time.Second)
	broadcastClient.ConfirmedAccumulatorListener = make(chan common.Hash, 1)
	defer wg.Done()
	messageCount := 0
	ctx := context.Background()

	// connect returns
	messageReceiver, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for {
		select {
		case receivedMsg := <-messageReceiver:
			t.Logf("Received Message, Sequence Message: %v\n", receivedMsg.FeedItem.BatchItem.SequencerMessage)
			messageCount++

			if messageCount == expectedCount {
				broadcastClient.Close()
				return
			}
		case confirmedAccumulator := <-broadcastClient.ConfirmedAccumulatorListener:
			t.Logf("Received confirmedAccumulator, Sequence Message: %v\n", confirmedAccumulator.ShortString())
		}
	}
}
