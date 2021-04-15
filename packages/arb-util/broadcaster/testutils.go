package broadcaster

import (
	"errors"
	"math/big"
	"reflect"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

func SequencedMessages() func() (*big.Int, []byte, *big.Int) {
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
type RandomMessageGenerator struct {
	broadcaster              *Broadcaster
	startWorkerMutex         *sync.Mutex
	messageBroadcasterWorker *time.Ticker
	count                    int
	intervalDuration         time.Duration
	workerStarted            bool
}

// create a new test message generator
func NewRandomMessageGenerator(count int, ms int) *RandomMessageGenerator {
	gm := &RandomMessageGenerator{}
	gm.startWorkerMutex = &sync.Mutex{}
	gm.intervalDuration = time.Duration(ms) * time.Millisecond
	gm.workerStarted = false
	gm.count = count
	return gm
}

// give it a client manager to broadcast on.
func (mg *RandomMessageGenerator) SetBroadcaster(broadcaster *Broadcaster) {
	mg.broadcaster = broadcaster
}

func (mg *RandomMessageGenerator) StartWorker() {
	mg.startWorkerMutex.Lock()
	defer mg.startWorkerMutex.Unlock()
	if mg.workerStarted {
		return
	}

	ticker := time.NewTicker(mg.intervalDuration)
	messageCount := 0
	newBroadcastMessage := SequencedMessages()
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

func (mg *RandomMessageGenerator) StopWorker() {
	if mg.messageBroadcasterWorker != nil {
		mg.messageBroadcasterWorker.Stop()
		mg.workerStarted = false
	}
}
