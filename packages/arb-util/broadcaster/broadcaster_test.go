package broadcaster

import (
	"context"
	"encoding/json"
	"math/big"
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

	err = b.Broadcast(newBroadcastMessage())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	err = b.Broadcast(newBroadcastMessage())
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go connectAndGetCachedMessages(t, i, &wg)
	}

	wg.Wait()

	b.SyncSequence(big.NewInt(42))  // first message is 42, remove it
	if b.messageCacheCount() != 1 { // should have left the second message
		t.Error("Failed to clear cached inbox message")
	}

	b.SyncSequence(big.NewInt(100)) // second message is 43, remove it, and then some
	if b.messageCacheCount() != 0 { // should have emptied.
		t.Error("Failed to clear cached inbox message")
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
