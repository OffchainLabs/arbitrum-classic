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
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func TestReceiveMessages(t *testing.T) {
	ctx := context.Background()

	settings := configuration.DefaultFeedOutput()
	settings.Port = "9742"

	messageCount := 1000
	messageDelay := 0 * time.Millisecond
	clientCount := 10

	b := broadcaster.NewBroadcaster(settings, 9742)

	broadcasterErrChan, err := b.Start(ctx)
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
	case err := <-broadcasterErrChan:
		t.Fatal(err)
	case err := <-errChan:
		t.Fatal(err)
	default:
	}
}

func startMakeBroadcastClient(ctx context.Context, t *testing.T, index int, expectedCount int, wg *sync.WaitGroup) {
	broadcastClientErrChan := make(chan error)
	broadcastClient := NewBroadcastClient(
		"ws://127.0.0.1:9742/",
		9742,
		nil,
		20*time.Second,
		broadcastClientErrChan,
	)
	messageCount := 0

	// connect returns
	messageReceiver, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	accListener := broadcastClient.ConfirmedAccumulatorListener

	if broadcastClient.chainId != 9742 {
		t.Fatalf("Incorrect chain id: %d", broadcastClient.chainId)
	}

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
			case err := <-broadcastClientErrChan:
				t.Errorf("broadcast client error: %s", err.Error())
				return
			case <-time.After(60 * time.Second):
				t.Errorf("Client %d expected %d meesages, only got %d messages\n", index, expectedCount, messageCount)
				return
			}
		}
	}()

}

func TestServerClientDisconnect(t *testing.T) {
	ctx := context.Background()

	settings := configuration.DefaultFeedOutput()
	settings.Port = "9743"
	settings.Ping = 1 * time.Second
	settings.ClientTimeout = 2 * time.Second

	b := broadcaster.NewBroadcaster(settings, 9743)

	_, err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	broadcastClientErrChan := make(chan error)
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9743/", 9743, nil, 20*time.Second, broadcastClientErrChan)

	client, err := broadcastClient.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	newBroadcastMessage := broadcaster.SequencedMessages()
	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.BroadcastSingle(hash1, feedItem1.BatchItem, signature1.Bytes())

	// Wait for client to receive batch to ensure it is connected
	select {
	case err := <-broadcastClientErrChan:
		t.Fatalf("broadcast client error: %s", err)
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

	settings := configuration.DefaultFeedOutput()
	settings.Port = "9744"
	settings.Ping = 50 * time.Second
	settings.ClientTimeout = 150 * time.Second

	b1 := broadcaster.NewBroadcaster(settings, 9744)

	_, err := b1.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b1.Stop()

	broadcastClientErrChan := make(chan error)
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9744/", 9744, nil, 2*time.Second, broadcastClientErrChan)

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

func implementBroadcasterSendsCachedMessagesOnClientConnect(t *testing.T, clientChainId uint64, serverChainId uint64) {
	ctx := context.Background()

	settings := configuration.DefaultFeedOutput()
	settings.Port = "9842"

	b := broadcaster.NewBroadcaster(settings, serverChainId)

	_, err := b.Start(ctx)
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

	hash3, feedItem3, signature3 := newBroadcastMessage()
	err = b.BroadcastSingle(hash3, feedItem3.BatchItem, signature3.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		err := connectAndGetCachedMessages(ctx, t, i, &wg, clientChainId, clientChainId != serverChainId)
		if err != nil {
			if clientChainId == serverChainId {
				t.Log("error occurred when chain id is correct")
				t.Fatal(err)
			}
			// Abort test without error
			wg.Done()
			break
		} else if clientChainId != serverChainId {
			t.Error("no error occurred when chain id incorrect")
		}
	}

	wg.Wait()

	// give the above connections time to reconnect
	time.Sleep(4 * time.Second)

	// Confirmed Accumulator will also broadcast to the clients.
	b.ConfirmedAccumulator(feedItem1.BatchItem.Accumulator) // remove the first message we generated

	// Send next accumulator because only previous accumulator is sent to clients
	b.ConfirmedAccumulator(feedItem2.BatchItem.Accumulator) // remove the first message we generated

	// Send next accumulator because only previous accumulator is sent to clients
	b.ConfirmedAccumulator(feedItem3.BatchItem.Accumulator) // remove the first message we generated

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
	b.ConfirmedAccumulator(feedItem3.BatchItem.Accumulator)

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

func TestBroadcasterSendsCachedMessagesOnClientConnect(t *testing.T) {
	chainId := uint64(9842)
	implementBroadcasterSendsCachedMessagesOnClientConnect(t, chainId, chainId)
}

func TestBadBroadcasterSendsCachedMessagesOnClientConnect(t *testing.T) {
	chainId := uint64(9842)
	implementBroadcasterSendsCachedMessagesOnClientConnect(t, chainId, chainId+1)
}

func connectAndGetCachedMessages(ctx context.Context, t *testing.T, clientIndex int, wg *sync.WaitGroup, chainId uint64, connectShouldFail bool) error {
	broadcastClientErrChan := make(chan error)
	requestedSeqNum := big.NewInt(42)
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9842/", chainId, requestedSeqNum, 60*time.Second, broadcastClientErrChan)
	testClient, err := broadcastClient.Connect(ctx)
	if err != nil {
		if !connectShouldFail {
			t.Fatal(err)
		}
		return err
	} else if connectShouldFail {
		t.Fatal("broadcast client didn't fail when it should have")
	}

	go func() {
		defer wg.Done()
		defer broadcastClient.Close()

		t.Logf("client %d %v connected\n", clientIndex, (*broadcastClient).conn.LocalAddr())

		// Wait for client to receive first item
		select {
		case err := <-broadcastClientErrChan:
			t.Errorf("broadcast client error: %s", err.Error())
		case receivedMsg := <-testClient:
			if receivedMsg.FeedItem.BatchItem.LastSeqNum.Cmp(requestedSeqNum) != 0 {
				t.Errorf("expected seqnum %d but got %d instead", requestedSeqNum, receivedMsg.FeedItem.BatchItem.LastSeqNum)
			}
			t.Logf("client %d received first message: (%v) %v\n", clientIndex, receivedMsg.FeedItem.BatchItem.LastSeqNum, receivedMsg.FeedItem.BatchItem.SequencerMessage)
		case <-time.After(10 * time.Second):
			t.Errorf("client %d did not receive first batch item\n", clientIndex)
			return
		}

		// Wait for client to receive second item
		select {
		case err := <-broadcastClientErrChan:
			t.Errorf("broadcast client error: %s", err.Error())
		case receivedMsg := <-testClient:
			t.Logf("client %d received second message: (%v) %v\n", clientIndex, receivedMsg.FeedItem.BatchItem.LastSeqNum, receivedMsg.FeedItem.BatchItem.SequencerMessage)
		case <-time.After(10 * time.Second):
			t.Errorf("client %d did not receive second batch item\n", clientIndex)
			return
		}
	}()

	return nil
}
