package broadcaster

import (
	"context"
	"encoding/json"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func TestBroadcasterSendsCachedMessagesOnClientConnect(t *testing.T) {
	broadcasterSettings := Settings{
		Addr:      ":9642",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	newBroadcastMessage := SequencedMessages()

	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.Broadcast(hash1, feedItem1, signature1)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	hash2, feedItem2, signature2 := newBroadcastMessage()
	err = b.Broadcast(hash2, feedItem2, signature2)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go connectAndGetCachedMessages(t, i, &wg)
	}

	wg.Wait()

	b.SyncMessages(feedItem1.BatchItem.Accumulator) // remove the first message we generated
	if b.messageCacheCount() != 1 {                 // should have left the second message
		t.Errorf("1. Failed to clear cached inbox message. MessageCacheCount: %v", b.messageCacheCount())
	}

	b.SyncMessages(feedItem2.BatchItem.Accumulator) // remove the second message we generated
	if b.messageCacheCount() != 0 {                 // should have emptied.
		t.Errorf("2. Failed to clear cached inbox message. MessageCacheCount: %v", b.messageCacheCount())
	}
}

func connectAndGetCachedMessages(t *testing.T, i int, wg *sync.WaitGroup) {
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
			t.Logf("%d closed\n", i)
		}
	}(conn)

	t.Logf("%d connected\n", i)

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
		messagesReceived = len(res.Messages)
		t.Logf("%d receive: %v，type: %v\n", i, res, op)
	}

	if messagesReceived != 2 {
		t.Errorf("%d Should have received two cached messages: %s\n", i, err)
	}
}

func TestBroadcasterRespondsToPing(t *testing.T) {
	broadcasterSettings := Settings{
		Addr:      ":9643",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}
	b := NewBroadcaster(broadcasterSettings)

	err := b.Start()
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

	conn.Write(ws.CompiledPing)

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
	broadcasterSettings := Settings{
		Addr:      ":9642",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	b := NewBroadcaster(broadcasterSettings)

	err := b.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	newBroadcastMessage := SequencedMessages()

	hash1, feedItem1, signature1 := newBroadcastMessage()
	err = b.Broadcast(hash1, feedItem1, signature1)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	hash2, feedItem2, signature2 := newBroadcastMessage()
	err = b.Broadcast(hash2, feedItem2, signature2)
	if err != nil {
		t.Fatal(err)
	}

	hash3, feedItem3, signature3 := newBroadcastMessage()
	err = b.Broadcast(hash3, feedItem3, signature3)
	if err != nil {
		t.Fatal(err)
	}

	_, feedItem4, signature4 := newBroadcastMessage()
	feedItem4.PrevAcc = feedItem1.BatchItem.Accumulator
	err = b.Broadcast(feedItem4.PrevAcc, feedItem4, signature4)
	if err != nil {
		t.Fatal(err)
	}

	if b.messageCacheCount() != 2 {
		t.Errorf("1. Failed to reorganized cached inbox message. MessageCacheCount: %v", b.messageCacheCount())
	}

	//TODO: Add some more assertions about the state of the cache
}
