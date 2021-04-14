package broadcaster

import (
	"context"
	"encoding/json"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func TestBroadcaster(t *testing.T) {
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

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go broadcastWait(t, i, &wg)
	}
	ib := inbox.InboxMessage{
		Kind: 1,
		Sender: common.HexToAddress("0x12345678123456781234567812345678"),
		InboxSeqNum: big.NewInt(42),
		GasPrice: big.NewInt(43),
		Data: []byte{4, 2},
		ChainTime: inbox.NewRandomChainTime(),
	}

	messages := []*inbox.InboxMessage{
		&ib,
	}
	err = b.Broadcast(messages)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	err = b.Broadcast(messages)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
}

func broadcastWait(t *testing.T, i int, wg *sync.WaitGroup) {
	defer wg.Done()

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
		res := Response{}
		err = json.Unmarshal([]byte(msg), &res)
		if err != nil {
			t.Errorf("Unable to marshal message: %v\n", err)
			return
		}
		// println(res.ID)
		t.Logf("%d receive: %v，type: %v\n", i, res, op)
	}

	time.Sleep(3 * time.Second)
}
