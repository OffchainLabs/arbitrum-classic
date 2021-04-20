package monitor

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type SequencerFeedItem struct {
	BatchItem inbox.SequencerBatchItem
	PrevAcc   common.Hash
}

type InboxReader struct {
	// Only in run thread
	delayedBridge      *ethbridge.DelayedBridgeWatcher
	sequencerInbox     *ethbridge.SequencerInboxWatcher
	db                 core.ArbCore
	firstMessageBlock  *big.Int
	caughtUp           bool
	caughtUpTarget     *big.Int
	healthChan         chan nodehealth.Log
	lastAcc            common.Hash
	sequencerFeedQueue []SequencerFeedItem

	// Only in main thread
	running    bool
	cancelFunc context.CancelFunc
	completed  chan bool

	// Thread safe
	caughtUpChan         chan bool
	MessageDeliveryMutex sync.Mutex
	SequencerFeed        chan SequencerFeedItem
}

func NewInboxReader(ctx context.Context, bridge *ethbridge.DelayedBridgeWatcher, sequencerInbox *ethbridge.SequencerInboxWatcher, db core.ArbCore, healthChan chan nodehealth.Log) (*InboxReader, error) {
	firstMessageBlock, err := bridge.LookupMessageBlock(ctx, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	return &InboxReader{
		delayedBridge:     bridge,
		sequencerInbox:    sequencerInbox,
		db:                db,
		firstMessageBlock: firstMessageBlock.Height.AsInt(),
		completed:         make(chan bool, 1),
		caughtUpChan:      make(chan bool, 1),
		healthChan:        healthChan,
		SequencerFeed:     make(chan SequencerFeedItem, 128),
	}, nil
}

func (ir *InboxReader) Start(parentCtx context.Context) {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	go func() {
		defer func() {
			ir.completed <- true
		}()
		for {
			err := ir.getMessages(ctx)
			if err == nil {
				break
			}
			logger.Warn().Stack().Err(err).Msg("Failed to read inbox messages")
			<-time.After(time.Second * 2)
		}
	}()
	ir.cancelFunc = cancelFunc
	ir.running = true
}

func (ir *InboxReader) Stop() {
	ir.cancelFunc()
	<-ir.completed
	ir.running = false
}

func (ir *InboxReader) IsRunning() bool {
	return ir.running
}

// May only be called once
func (ir *InboxReader) WaitToCatchUp() {
	<-ir.caughtUpChan
}

func (ir *InboxReader) getMessages(ctx context.Context) error {
	from, err := ir.getNextBlockToRead()
	if err != nil {
		return err
	}
	if ir.healthChan != nil && from != nil {
		ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "getNextBlockToRead", ValBigInt: new(big.Int).Set(from)}
	}
	reorging := false
	blocksToFetch := uint64(100)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		currentHeight, err := ir.delayedBridge.CurrentBlockHeight(ctx)
		if err != nil {
			return err
		}

		if ir.healthChan != nil && currentHeight != nil {
			ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "currentHeight", ValBigInt: new(big.Int).Set(currentHeight)}
		}

		for {
			if !ir.caughtUp && ir.caughtUpTarget != nil {
				ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(ir.caughtUpTarget)}
				ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: true}
				arbCorePosition := ir.db.MachineMessagesRead()
				ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: false}
				if ir.healthChan != nil && arbCorePosition != nil {
					ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "arbCorePosition", ValBigInt: new(big.Int).Set(arbCorePosition)}
				}
				if arbCorePosition.Cmp(ir.caughtUpTarget) >= 0 {
					ir.caughtUp = true
					ir.caughtUpChan <- true
				}
			}
			if from.Cmp(currentHeight) >= 0 {
				break
			}
			to := new(big.Int).Add(from, new(big.Int).SetUint64(blocksToFetch))
			if to.Cmp(currentHeight) > 0 {
				to = currentHeight
			}
			delayedMessages, err := ir.delayedBridge.LookupMessagesInRange(ctx, from, to)
			if err != nil {
				return err
			}
			sequencerBatches, err := ir.sequencerInbox.LookupBatchesInRange(ctx, from, to)
			if err != nil {
				return err
			}
			if ir.caughtUpTarget == nil && to.Cmp(currentHeight) == 0 {
				if len(sequencerBatches) > 0 {
					ir.caughtUpTarget = sequencerBatches[len(sequencerBatches)-1].GetAfterCount()
				} else {
					dbMessageCount, err := ir.db.GetMessageCount()
					if err != nil {
						return err
					}
					ir.caughtUpTarget = dbMessageCount
				}
			}
			if len(sequencerBatches) > 0 {
				batchAccs := make([]common.Hash, 0, len(sequencerBatches)+1)
				lastSeqNums := make([]*big.Int, 0, len(sequencerBatches)+1)
				firstBeforeCount := sequencerBatches[0].GetBeforeCount()
				checkingStart := firstBeforeCount.Cmp(big.NewInt(0)) > 0
				if checkingStart {
					lastSeqNums = append(lastSeqNums, new(big.Int).Sub(firstBeforeCount, big.NewInt(1)))
					batchAccs = append(batchAccs, sequencerBatches[0].GetBeforeAcc())
				}
				for _, batch := range sequencerBatches {
					if len(batchAccs) > 0 && batch.GetBeforeAcc() != batchAccs[len(batchAccs)-1] {
						return errors.New("Mismatching batch accumulators; reorg?")
					}
					afterCount := batch.GetAfterCount()
					if afterCount.Cmp(big.NewInt(0)) > 0 {
						lastSeqNums = append(lastSeqNums, new(big.Int).Sub(afterCount, big.NewInt(1)))
						batchAccs = append(batchAccs, batch.GetAfterAcc())
					}
				}
				matching, err := ir.db.CountMatchingBatchAccs(lastSeqNums, batchAccs)
				if err != nil {
					return err
				}
				reorging = false
				if checkingStart {
					if matching == 0 {
						reorging = true
					} else {
						matching--
					}
				}
				sequencerBatches = sequencerBatches[matching:]
			}
			if !reorging && len(delayedMessages) > 0 {
				firstMsg := delayedMessages[0]
				beforeAcc := firstMsg.BeforeInboxAcc
				beforeSeqNum := new(big.Int).Sub(firstMsg.Message.InboxSeqNum, big.NewInt(1))
				if beforeSeqNum.Cmp(big.NewInt(0)) >= 0 {
					haveAcc, err := ir.db.GetDelayedInboxAcc(beforeSeqNum)
					if err != nil || haveAcc != beforeAcc {
						reorging = true
					}
				}
			}
			if ir.healthChan != nil && ir.caughtUpTarget != nil {
				ir.healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(ir.caughtUpTarget)}
			}
			if len(sequencerBatches) < 5 {
				blocksToFetch += 20
			} else if len(sequencerBatches) > 10 {
				blocksToFetch /= 2
			}
			if blocksToFetch == 0 {
				blocksToFetch++
			}
			logger.Debug().
				Str("from", from.String()).
				Str("to", to.String()).
				Int("delayedCount", len(delayedMessages)).
				Int("batchCount", len(sequencerBatches)).
				Msg("Looking up messages")
			if reorging {
				from, err = ir.getPrevBlockForReorg(from)
				if err != nil {
					return err
				}
			} else {
				if len(delayedMessages) != 0 || len(sequencerBatches) != 0 {
					err := ir.addMessages(ctx, sequencerBatches, delayedMessages)
					if err != nil {
						return err
					}
				}
				from = from.Add(to, big.NewInt(1))
			}
		}
		readFromFeed := false
	FeedReadLoop:
		for {
			select {
			case feedItem := <-ir.SequencerFeed:
				readFromFeed = true
				feedReorg := len(ir.sequencerFeedQueue) != 0 && ir.sequencerFeedQueue[len(ir.sequencerFeedQueue)-1].BatchItem.Accumulator != feedItem.PrevAcc
				feedCaughtUp := feedItem.PrevAcc == ir.lastAcc
				if feedReorg || feedCaughtUp {
					ir.sequencerFeedQueue = []SequencerFeedItem{}
				}
				ir.sequencerFeedQueue = append(ir.sequencerFeedQueue, feedItem)
			default:
				break FeedReadLoop
			}
		}
		if !readFromFeed {
			time.Sleep(time.Second)
		}
		if len(ir.sequencerFeedQueue) > 0 && ir.sequencerFeedQueue[0].PrevAcc == ir.lastAcc {
			queueItems := make([]inbox.SequencerBatchItem, 0, len(ir.sequencerFeedQueue))
			for _, item := range ir.sequencerFeedQueue {
				queueItems = append(queueItems, item.BatchItem)
			}
			prevAcc := ir.sequencerFeedQueue[0].PrevAcc
			ir.sequencerFeedQueue = []SequencerFeedItem{}
			ok, err := core.DeliverMessagesAndWait(ir.db, prevAcc, queueItems, []inbox.DelayedMessage{}, nil)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("Failed to deliver sequencer feed messages to ArbCore")
			}
			ir.lastAcc = queueItems[len(queueItems)-1].Accumulator
		}
	}
}

func (ir *InboxReader) getNextBlockToRead() (*big.Int, error) {
	messageCount, err := ir.db.GetMessageCount()
	if err != nil {
		return nil, err
	}
	if messageCount.Cmp(big.NewInt(0)) == 0 {
		return ir.firstMessageBlock, nil
	}
	var acc common.Hash
	if messageCount.Cmp(big.NewInt(0)) > 0 {
		acc, err = ir.db.GetInboxAcc(new(big.Int).Sub(messageCount, big.NewInt(1)))
		if err != nil {
			return nil, err
		}
		ir.lastAcc = acc
	}
	for i, queueItem := range ir.sequencerFeedQueue {
		if queueItem.BatchItem.LastSeqNum.Cmp(messageCount) >= 0 {
			break
		}
		if queueItem.BatchItem.Accumulator.Equals(acc) {
			ir.sequencerFeedQueue = ir.sequencerFeedQueue[(i + 1):]
			break
		}
	}
	seqNum := messageCount
	zeroTime := common.NewTimeBlocksInt(0)
	for {
		seqNum.Sub(seqNum, big.NewInt(1))
		msg, err := core.GetSingleMessage(ir.db, seqNum)
		if err != nil {
			return nil, err
		}
		if msg.ChainTime.BlockNum.Cmp(zeroTime) != 0 {
			return msg.ChainTime.BlockNum.AsInt(), nil
		}
	}
}

func (ir *InboxReader) getPrevBlockForReorg(from *big.Int) (*big.Int, error) {
	if from.Cmp(ir.firstMessageBlock) == 0 {
		return nil, errors.New("can't get older messages")
	}
	newFrom := new(big.Int).Sub(from, big.NewInt(10))
	if newFrom.Cmp(ir.firstMessageBlock) < 0 {
		newFrom = ir.firstMessageBlock
	}
	return newFrom, nil
}

func (ir *InboxReader) addMessages(ctx context.Context, sequencerBatchRefs []ethbridge.SequencerBatchRef, deliveredDelayedMessages []*ethbridge.DeliveredInboxMessage) error {
	var seqBatchItems []inbox.SequencerBatchItem
	for _, ref := range sequencerBatchRefs {
		batch, err := ir.sequencerInbox.ResolveBatchRef(ctx, ref)
		if err != nil {
			return err
		}
		items, err := batch.GetItems()
		if err != nil {
			return err
		}
		seqBatchItems = append(seqBatchItems, items...)
	}
	delayedMessages := make([]inbox.DelayedMessage, 0, len(deliveredDelayedMessages))
	for _, deliveredMsg := range deliveredDelayedMessages {
		msg := inbox.DelayedMessage{
			DelayedSequenceNumber: deliveredMsg.Message.InboxSeqNum,
			DelayedAccumulator:    deliveredMsg.AfterInboxAcc(),
			Message:               deliveredMsg.Message.ToBytes(),
		}
		delayedMessages = append(delayedMessages, msg)
	}
	ir.MessageDeliveryMutex.Lock()
	defer ir.MessageDeliveryMutex.Unlock()
	var beforeAcc common.Hash
	if len(sequencerBatchRefs) > 0 {
		beforeAcc = sequencerBatchRefs[0].GetBeforeAcc()
	}
	ok, err := core.DeliverMessagesAndWait(ir.db, beforeAcc, seqBatchItems, delayedMessages, nil)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Failed to deliver messages to ArbCore")
	}
	for _, item := range seqBatchItems {
		if len(ir.sequencerFeedQueue) == 0 {
			break
		}
		firstQueueItem := ir.sequencerFeedQueue[0].BatchItem
		if item.LastSeqNum.Cmp(firstQueueItem.LastSeqNum) > 0 {
			break
		}
		if item.Accumulator.Equals(firstQueueItem.Accumulator) {
			ir.sequencerFeedQueue = ir.sequencerFeedQueue[1:]
		}
	}
	if len(seqBatchItems) > 0 {
		ir.lastAcc = seqBatchItems[len(seqBatchItems)-1].Accumulator
	}
	return nil
}
