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

package broadcaster

import (
	"context"
	"encoding/json"
	"github.com/mailru/easygo/netpoll"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func TestBroadcasterSendsConfirmedAccumulatorMessages(t *testing.T) {
	ctx := context.Background()

	broadcasterSettings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9642",
		Ping:          5 * time.Second,
		ClientTimeout: 20 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	newBroadcastMessage := SequencedMessages()

	_, feedItem, _ := newBroadcastMessage()
	time.Sleep(1 * time.Second)

	accumulatorConfirmed := make(chan common.Hash)
	var wg sync.WaitGroup
	wg.Add(1)
	startReceivedConfirmedAccumulator(t, &wg, accumulatorConfirmed)

	time.Sleep(2 * time.Second)

	// Only previous accumulator will also broadcast to the clients, so send twice for test
	b.ConfirmedAccumulator(feedItem.BatchItem.Accumulator) // remove the first message we generated
	b.ConfirmedAccumulator(feedItem.BatchItem.Accumulator) // remove the first message we generated

	acc := <-accumulatorConfirmed
	if acc != feedItem.BatchItem.Accumulator {
		t.Error("Did not receive expected accumulator")
	}

	wg.Wait()
}

func startReceivedConfirmedAccumulator(t *testing.T, wg *sync.WaitGroup, accumulatorConfirmed chan common.Hash) {

	go func() {
		confirmedAccumulatorReceived := 0
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9642/")
		if err != nil {
			t.Errorf("Can not connect: %v\n", err)
			return
		}

		poller, err := netpoll.New(nil)
		if err != nil {
			t.Error("error starting net poller")
			return
		}

		desc, err := netpoll.HandleRead(conn)
		if err != nil {
			t.Error("error getting netpoll descriptor")
			return
		}

		err = poller.Start(desc, func(ev netpoll.Event) {
			if ev&netpoll.EventReadHup != 0 {
				t.Error("received hang up")
				_ = poller.Stop(desc)
				_ = conn.Close()
				wg.Done()
				return
			}

			msg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				t.Error("error calling ReadServerData")
				_ = poller.Stop(desc)
				_ = conn.Close()
				wg.Done()
				return
			}

			res := BroadcastMessage{}
			err = json.Unmarshal(msg, &res)
			if err != nil {
				logger.Error().Err(err).Msg("error unmarshalling message")
				_ = poller.Stop(desc)
				_ = conn.Close()
				wg.Done()

				return
			}

			if res.Version != 1 {
				t.Error("This is not version 1")
			}

			if res.ConfirmedAccumulator.IsConfirmed {
				confirmedAccumulatorReceived++
				accumulatorConfirmed <- res.ConfirmedAccumulator.Accumulator
			}

			if confirmedAccumulatorReceived == 1 { // this gets called twice from the test
				_ = poller.Stop(desc)
				_ = conn.Close()
				wg.Done()
				return
			}

		})

		if err != nil {
			t.Errorf("Problem starting poller: %v\n", err)
		}
	}()
}

func TestBroadcasterRespondsToPing(t *testing.T) {
	t.Skip("Server is not responding to ping anymore")
	ctx := context.Background()

	broadcasterSettings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9643",
		Ping:          5 * time.Second,
		ClientTimeout: 20 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9643/")
	if err != nil {
		t.Fatalf("Can not connect: %v\n", err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			t.Errorf("Can not close: %v\n", err)
		} else {
			t.Log("%Closed\n")
		}
	}(conn)

	t.Logf("Connected")

	_, err = conn.Write(ws.CompiledPing)
	if err != nil {
		t.Fatalf("unable to write: %v\n", err)
	}

	time.Sleep(1 * time.Second)

	h, _, _ := wsutil.NextReader(conn, ws.StateClientSide)
	switch h.OpCode {
	case ws.OpPing:
		t.Errorf("Received ping but should have be a pong")
	case ws.OpPong:
		t.Log("Received pong!")
	default:
		t.Errorf("Received uknown OpCode from server after ping: %v", h.OpCode)
	}

	time.Sleep(1 * time.Second)
}

func TestBroadcasterReorganizesCacheBasedOnAccumulator(t *testing.T) {
	ctx, cancelFunc, _ := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	broadcasterSettings := configuration.FeedOutput{
		Addr:          "0.0.0.0",
		IOTimeout:     2 * time.Second,
		Port:          "9642",
		Ping:          5 * time.Second,
		ClientTimeout: 30 * time.Second,
		Queue:         1,
		Workers:       128,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	newBroadcastMessage := SequencedMessages()

	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.BroadcastSingle(hash1, feedItem1.BatchItem, signature1.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)

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

	_, feedItem4, signature4 := newBroadcastMessage()
	feedItem4.PrevAcc = feedItem1.BatchItem.Accumulator
	err = b.BroadcastSingle(feedItem4.PrevAcc, feedItem4.BatchItem, signature4.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	if b.MessageCacheCount() != 2 {
		t.Errorf("1. Failed to reorganized cached inbox message. MessageCacheCount: %v", b.MessageCacheCount())
	}

	//TODO: Add some more assertions about the state of the cache
}
