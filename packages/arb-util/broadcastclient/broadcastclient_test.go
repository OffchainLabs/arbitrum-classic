package broadcastclient

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
	seqNumEnd := 117
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

// This sends out generated test broadcast messages
type MessageGenerator struct {
	broadcaster              *broadcaster.Broadcaster
	startWorkerMutex         *sync.Mutex
	messageBroadcasterWorker *time.Ticker
	count                    int
	intervalDuration         time.Duration
	workerStarted            bool
}

// create a new test message generator
func NewMessageGenerator(count int, ms int) *MessageGenerator {
	gm := &MessageGenerator{}
	gm.startWorkerMutex = &sync.Mutex{}
	gm.intervalDuration = time.Duration(ms) * time.Millisecond
	gm.workerStarted = false
	gm.count = count
	return gm
}

// give it a client manager to broadcast on.
func (mg *MessageGenerator) setBroadcaster(broadcaster *broadcaster.Broadcaster) {
	mg.broadcaster = broadcaster
}

func (mg *MessageGenerator) startWorker() {
	mg.startWorkerMutex.Lock()
	defer mg.startWorkerMutex.Unlock()
	if mg.workerStarted {
		return
	}

	ticker := time.NewTicker(mg.intervalDuration)
	messageCount := 0
	newBroadcastMessage := sequencedMessages()
	go func() {
		for range ticker.C {
			_ = mg.broadcaster.Broadcast(newBroadcastMessage())
			messageCount++
			if messageCount == mg.count {
				ticker.Stop()
				return
			}
		}
	}()

	mg.messageBroadcasterWorker = ticker
	mg.workerStarted = true
}

func (mg *MessageGenerator) stopWorker() {
	if mg.messageBroadcasterWorker != nil {
		mg.messageBroadcasterWorker.Stop()
		mg.workerStarted = false
	}
}

func TestBroadCastClient(t *testing.T) {
	broadcasterSettings := broadcaster.Settings{
		Addr:      ":9643",
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}

	b := broadcaster.NewBroadcaster(broadcasterSettings)

	err := b.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer b.Stop()

	// this will send test messages to the clients at an interval
	tmb := NewMessageGenerator(10, 100)
	tmb.setBroadcaster(b)

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go makeBroadcastClient(t, 10, &wg)
	}

	tmb.startWorker()
	wg.Wait()
	tmb.stopWorker()
}

func makeBroadcastClient(t *testing.T, expectedCount int, wg *sync.WaitGroup) {
	broadcastClient := NewBroadcastClient("ws://127.0.0.1:9643/", nil)
	defer wg.Done()
	messageCount := 0

	// connect returns
	messages, err := broadcastClient.Connect()
	if err != nil {
		t.Errorf("Can not connect: %v\n", err)
	}

	for {
		select {
		case receivedMsgs := <-messages:
			for i := range receivedMsgs.Messages {
				fmt.Printf("Received Message, Sequence Number: %v\n", inbox.GetSequenceNumber(receivedMsgs.Messages[i].InboxMessage))
				messageCount++
				if messageCount == expectedCount {
					broadcastClient.Close()
					return
				}
			}
		}
	}

}
