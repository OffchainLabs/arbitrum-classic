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

package broadcastclient

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
)

func TestReceiveMessages(t *testing.T) {
	ctx := context.Background()

	settings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9742",
		Ping:          5 * time.Second,
		ClientTimeout: 20 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	messageCount := 1000
	messageDelay := 0 * time.Millisecond
	clientCount := 2

	b := broadcaster.NewBroadcaster(settings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	// this will send test messages to the clients at an interval
	tmb := broadcaster.NewRandomMessageGenerator(messageCount, messageDelay)
	tmb.SetBroadcaster(b)

	var wg sync.WaitGroup
	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		startMakeBroadcastClient(ctx, t, i, messageCount, &wg)
	}

	errChan := tmb.Start(ctx)
	wg.Wait()

	select {
	case err := <-errChan:
		t.Fatal(err)
	default:
	}
}

func startMakeBroadcastClient(ctx context.Context, t *testing.T, index int, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9742/", nil, 20*time.Second)
	messageCount := 0

	// connect returns
	messageReceiver, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	accListener := broadcastClient.ConfirmedAccumulatorListener

	go func() {
		defer wg.Done()
		defer broadcastClient.Close()
		for {
			select {
			case <-messageReceiver:
				messageCount++

				if messageCount == expectedCount {
					return
				}
			case <-accListener:
			case <-time.After(60 * time.Second):
				t.Errorf("Client %d expected %d meesages, only got %d messages\n", index, expectedCount, messageCount)
				return
			}
		}
	}()

}

func TestServerClientDisconnect(t *testing.T) {
	ctx := context.Background()

	settings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9743",
		Ping:          1 * time.Second,
		ClientTimeout: 2 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b := broadcaster.NewBroadcaster(settings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", nil, 20*time.Second)

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

	broadcastClient.Close()

	disconnectTimeout := time.After(5 * time.Second)
	for {
		if b.ClientCount() == 0 {
			break
		}

		select {
		case <-disconnectTimeout:
			t.Fatal("Client was not disconnected")
		case <-time.After(100 * time.Millisecond):
		}
	}
}

func TestBroadcastClientReconnectsOnServerDisconnect(t *testing.T) {
	ctx := context.Background()

	settings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9743",
		Ping:          50 * time.Second,
		ClientTimeout: 150 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b1 := broadcaster.NewBroadcaster(settings)

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

func TestBroadcasterSendsCachedMessagesOnClientConnect(t *testing.T) {
	ctx := context.Background()

	settings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9842",
		Ping:          5 * time.Second,
		ClientTimeout: 15 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b := broadcaster.NewBroadcaster(settings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	newBroadcastMessage := broadcaster.SequencedMessages()

	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.BroadcastSingle(hash1, feedItem1.BatchItem, signature1.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	hash2, feedItem2, signature2 := newBroadcastMessage()
	err = b.BroadcastSingle(hash2, feedItem2.BatchItem, signature2.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		connectAndGetCachedMessages(ctx, t, i, &wg)
	}

	wg.Wait()

	// give the above connections time to reconnect
	time.Sleep(4 * time.Second)

	// Confirmed Accumulator will also broadcast to the clients.
	b.ConfirmedAccumulator(feedItem1.BatchItem.Accumulator) // remove the first message we generated

	// Send next accumulator because only previous accumulator is sent to clients
	b.ConfirmedAccumulator(feedItem2.BatchItem.Accumulator) // remove the first message we generated

	updateTimeout := time.After(2 * time.Second)
	for {
		if b.MessageCacheCount() == 1 { // should have left the second message
			break
		}

		select {
		case <-updateTimeout:
			t.Fatal("confirmed accumulator did not get updated")
		case <-time.After(10 * time.Millisecond):
		}
	}

	// Send second accumulator again so that the previously added accumulator is sent
	b.ConfirmedAccumulator(feedItem2.BatchItem.Accumulator)

	updateTimeout = time.After(2 * time.Second)
	for {
		if b.MessageCacheCount() == 0 { // should have left the second message
			break
		}

		select {
		case <-updateTimeout:
			t.Fatal("cache did not get cleared")
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func connectAndGetCachedMessages(ctx context.Context, t *testing.T, clientIndex int, wg *sync.WaitGroup) {
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9842/", nil, 60*time.Second)
	testClient, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		defer wg.Done()
		defer broadcastClient.Close()

		t.Logf("client %d %v connected\n", clientIndex, (*broadcastClient).conn.LocalAddr())

		// Wait for client to receive first item
		select {
		case receivedMsg := <-testClient:
			t.Logf("client %d received first message: %v\n", clientIndex, receivedMsg.FeedItem.BatchItem.SequencerMessage)
		case <-time.After(10 * time.Second):
			t.Errorf("client %d did not receive first batch item\n", clientIndex)
			return
		}

		// Wait for client to receive second item
		select {
		case receivedMsg := <-testClient:
			t.Logf("client %d received second message: %v\n", clientIndex, receivedMsg.FeedItem.BatchItem.SequencerMessage)
		case <-time.After(10 * time.Second):
			t.Errorf("client %d did not receive second batch item\n", clientIndex)
			return
		}
	}()
}
