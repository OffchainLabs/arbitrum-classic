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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var MessageCount = 10
var ClientCount = 100

func TestBroadcasterLoad(t *testing.T) {
	ctx := context.Background()

	broadcasterSettings := Settings{
		Addr:                    ":9642",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      5 * time.Second,
		ClientNoResponseTimeout: 15 * time.Second,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	var wg sync.WaitGroup
	for i := 0; i < ClientCount; i++ {
		wg.Add(1)
		go receiveMessages(t, i, &wg)
	}

	// probably should wait on connections being established
	time.Sleep(10 * time.Millisecond)

	go broadcastTonsOfMessages(b, t)

	wg.Wait()

}

func receiveMessages(t *testing.T, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	messagesReceived := 0
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9642/")
	if err != nil {
		t.Errorf("%d can not connect: %v\n", i, err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			t.Errorf("%d can not close: %v\n", i, err)
		} else {
			//t.Logf("%d closed\n", i)
		}
	}(conn)

	var prevAcc common.Hash

	//t.Logf("%d connected\n", i)

	for {
		msg, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			t.Errorf("%d can not receive: %v\n", i, err)
			return
		} else {
			res := BroadcastMessage{}
			err = json.Unmarshal(msg, &res)
			if err != nil {
				t.Errorf("%d error unmarshalling message: %s\n", i, err)
				return
			}
			messagesReceived += len(res.Messages)
			for i := range res.Messages {
				msg := res.Messages[i]
				if prevAcc == common.HexToHash("0x0") || prevAcc == msg.FeedItem.PrevAcc {
					prevAcc = msg.FeedItem.BatchItem.Accumulator
				} else {

					t.Errorf("Message received out of order: previous: %v, expected previous: %v, current %v, client: %v\n", prevAcc, msg.FeedItem.PrevAcc, msg.FeedItem.BatchItem.Accumulator, conn.LocalAddr().String())
				}
			}
			_ = op
			//t.Logf("%d receive: %v，type: %v\n", i, res, op)
			if messagesReceived == MessageCount {
				break
			}
		}
	}

	if messagesReceived != MessageCount {
		t.Errorf("%d Should have received %d cached messages: %s\n", i, MessageCount, err)
	}
}

func broadcastTonsOfMessages(b *Broadcaster, t *testing.T) {
	newBroadcastMessage := SequencedMessages()
	for i := 0; i < MessageCount; i++ {
		hash1, feedItem1, signature1 := newBroadcastMessage()
		t.Logf("sending accumulator: %s", feedItem1.BatchItem.Accumulator.String())
		err := b.BroadcastSingle(hash1, feedItem1.BatchItem, signature1.Bytes())
		if err != nil {
			t.Error(err)
		}
		//t.Logf("sent %d messages", i+1)
		time.Sleep(10 * time.Millisecond)
	}
}
