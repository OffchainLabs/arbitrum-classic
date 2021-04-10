package broadcaster

import (
	"context"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"sync"
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	broadcasterSettings := Settings{
		Addr:      ":9642",
		Debug:     "",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}
	b := NewBroadcaster(broadcasterSettings)

	err := b.Start()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go chatWait(t, i, &wg)
	}
	wg.Wait()

	b.Stop()
}

func chatWait(t *testing.T, i int, wg *sync.WaitGroup) {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:9642/")
	defer wg.Done()
	if err != nil {
		t.Errorf("%d can not connect: %v\n", i, err)
	} else {
		t.Logf("%d connected\n", i)
		msg := []byte("OK+OK")
		err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
		if err != nil {
			t.Errorf("%d can not send: %v\n", i, err)
			return
		} else {
			t.Logf("%d send: %s, type: %v\n", i, msg, ws.OpText)
		}

		msg, op, err := wsutil.ReadServerData(conn)
		if err != nil {
			t.Errorf("%d can not receive: %v\n", i, err)
			return
		} else {
			t.Logf("%d receive: %s，type: %v\n", i, msg, op)
		}

		time.Sleep(time.Duration(3) * time.Second)

		err = conn.Close()
		if err != nil {
			t.Errorf("%d can not close: %v\n", i, err)
		} else {
			t.Logf("%d closed\n", i)
		}
	}
}
