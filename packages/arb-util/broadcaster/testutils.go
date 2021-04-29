package broadcaster

import (
	"errors"
	"math/big"
	"reflect"
	"sync"
	"time"

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

func setSequenceNumberInData(data []byte, sequenceNumber *big.Int) []byte {
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

// returns a function that when called returns the next random message in the sequence
func SequencedMessages() func() (common.Hash, SequencerFeedItem, *big.Int) {
	sequenceNumber := big.NewInt(41)
	accumulator := common.RandHash()

	return func() (common.Hash, SequencerFeedItem, *big.Int) {
		prevAccumulator := accumulator
		sequenceNumber = sequenceNumber.Add(sequenceNumber, big.NewInt(1))
		batchItem := inbox.SequencerBatchItem{}
		batchItem.LastSeqNum = sequenceNumber.Add(sequenceNumber, big.NewInt(1))
		batchItem.Accumulator = common.RandHash()
		accumulator = batchItem.Accumulator
		batchItem.TotalDelayedCount = big.NewInt(0)
		batchItem.SequencerMessage = setSequenceNumberInData(common.RandBytes(200), sequenceNumber)

		signature := common.RandBigInt()

		feedItem := SequencerFeedItem{
			BatchItem: batchItem,
			PrevAcc:   prevAccumulator,
		}

		return prevAccumulator, feedItem, signature
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

// SetBroadcaster sets a client manager to broadcast on.
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
	go func() {
		for range ticker.C {
			_ = mg.broadcaster.Broadcast(
				common.HexToHash("0x0001"),
				inbox.SequencerBatchItem{
					LastSeqNum:        big.NewInt(0),
					Accumulator:       common.HexToHash("0x01"),
					TotalDelayedCount: big.NewInt(0),
					SequencerMessage:  big.NewInt(42).Bytes(),
				},
				big.NewInt(0).Bytes(),
			)
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
