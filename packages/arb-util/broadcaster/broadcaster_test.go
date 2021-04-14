package broadcaster

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// This is used to reverse the slice for the sequence number field
// so that it will be in correct byte order when creating a new big.Int out of it
func reverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= (valueLen-1)/2; i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

func setSequenceNumber(data []byte, sequenceNumber *big.Int) []byte {
	seqNumOffset := 85
	seqNumEnd := seqNumOffset + 32
	prefixData := data[:seqNumOffset]
	postfixData := data[seqNumEnd:]
	sequenceNumberBytes := sequenceNumber.Bytes()
	sequenceNumberByteField := make([]byte, 32)
	copy(sequenceNumberByteField, sequenceNumberBytes)
	reverseSlice(sequenceNumberByteField)
	sequenceNumberWithPrefix := append(prefixData, sequenceNumberByteField...)
	completeDataWithSequenceNumberSet := append(sequenceNumberWithPrefix, postfixData...)
	return completeDataWithSequenceNumberSet
}

func sequencedMessages() func() (*big.Int, []byte, *big.Int) {
	sequenceNumber := big.NewInt(41)
	return func() (*big.Int, []byte, *big.Int) {
		sequenceNumber = sequenceNumber.Add(sequenceNumber, big.NewInt(1))
		inboxMessage := setSequenceNumber(common.RandBytes(200), sequenceNumber)
		beforeAccumulator := common.RandBigInt()
		signature := common.RandBigInt()
		return beforeAccumulator, inboxMessage, signature
	}
}

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
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go broadcastWait(t, i, &wg)
	}

	newBroadcastMessage := sequencedMessages()

	err = b.Broadcast(newBroadcastMessage())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	err = b.Broadcast(newBroadcastMessage())
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
		err = json.Unmarshal(msg, &res)
		if err != nil {
			t.Errorf("Unable to marshal message: %v\n", err)
			return
		}
		t.Logf("%d receive: %v，type: %v\n", i, res, op)
	}

	time.Sleep(3 * time.Second)
}
