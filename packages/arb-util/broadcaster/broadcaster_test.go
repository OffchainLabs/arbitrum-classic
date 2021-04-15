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
}

func connectAndGetCachedMessages(t *testing.T, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	messagesReceived := 0
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9642/")
	if err != nil {
		t.Fatalf("%d can not connect: %v\n", i, err)
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

	req := Request{}
	req.ID = 1
	req.Method = "ping"
	msg, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Can not Marshal ping request: %v\n", err)
	}

	err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
	if err != nil {
		t.Errorf("Can not send: %v\n", err)
		return
	}

	msg, op, err := wsutil.ReadServerData(conn)
	if err != nil {
		t.Errorf("Can not receive: %v\n", err)
		return
	} else {
		res := BroadcastMessage{}
		err = json.Unmarshal(msg, &res)
		if err != nil {
			t.Errorf("Error unmarshalling message: %s\n", err)
			return
		}

		if len(res.PongResponse) == 0 {
			t.Errorf("Should have received a ping response: %s\n", err)
		}
		t.Logf("Receive: %v，type: %v\n", res, op)
	}

	time.Sleep(1 * time.Second)
}

func TestBroadcasterRemovesOldSequences(t *testing.T) {

}
