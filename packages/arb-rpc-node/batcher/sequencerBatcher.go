/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package batcher

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type txQueueItem struct {
	tx         *types.Transaction
	resultChan chan error
}

type SequencerBatcher struct {
	db                              core.ArbCore
	inboxReader                     *monitor.InboxReader
	client                          ethutils.EthClient
	delayedMessagesTargetDelay      *big.Int
	sequencerInbox                  *ethbridgecontracts.SequencerInbox
	auth                            *ethbridge.TransactAuth
	chainTimeCheckInterval          time.Duration
	logBatchGasCosts                bool
	feedBroadcaster                 *broadcaster.Broadcaster
	dataSigner                      func([]byte) ([]byte, error)
	maxDelayBlocks                  *big.Int
	maxDelaySeconds                 *big.Int
	updateTimestampInterval         *big.Int
	sequenceDelayedMessagesInterval *big.Int
	createBatchBlockInterval        *big.Int

	sequencer common.Address
	signer    types.Signer
	txQueue   chan txQueueItem
	newTxFeed event.Feed

	latestChainTime        inbox.ChainTime
	lastCreatedBatchAt     *big.Int
	lastSequencedDelayedAt *big.Int
}

func getChainTime(ctx context.Context, client ethutils.EthClient) (inbox.ChainTime, error) {
	latestL1BlockInfo, err := client.BlockInfoByNumber(ctx, nil)
	if err != nil {
		return inbox.ChainTime{}, err
	}
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks((*big.Int)(latestL1BlockInfo.Number)),
		Timestamp: big.NewInt(int64(latestL1BlockInfo.Time)),
	}
	return chainTime, nil
}

func NewSequencerBatcher(
	ctx context.Context,
	db core.ArbCore,
	chainId *big.Int,
	inboxReader *monitor.InboxReader,
	client ethutils.EthClient,
	delayedMessagesTargetDelay *big.Int,
	createBatchBlockInterval *big.Int,
	sequencerInbox *ethbridgecontracts.SequencerInbox,
	auth *bind.TransactOpts,
	dataSigner func([]byte) ([]byte, error),
	broadcaster *broadcaster.Broadcaster,
	gasPriceUrl string,
) (*SequencerBatcher, error) {
	chainTime, err := getChainTime(ctx, client)
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: ctx}
	sequencer, err := sequencerInbox.Sequencer(callOpts)
	if err != nil {
		return nil, err
	}
	if sequencer != auth.From {
		return nil, errors.New("Transaction auth isn't for sequencer")
	}

	transactAuth, err := ethbridge.NewTransactAuth(ctx, client, auth, gasPriceUrl)
	if err != nil {
		return nil, err
	}

	maxDelayBlocks, err := sequencerInbox.MaxDelayBlocks(callOpts)
	if err != nil {
		return nil, err
	}
	maxDelaySeconds, err := sequencerInbox.MaxDelaySeconds(callOpts)
	if err != nil {
		return nil, err
	}

	if createBatchBlockInterval.Cmp(big.NewInt(0)) <= 0 || createBatchBlockInterval.Cmp(maxDelayBlocks) >= 0 {
		return nil, errors.New("invalid batch creation block interval")
	}

	return &SequencerBatcher{
		db:                         db,
		inboxReader:                inboxReader,
		client:                     client,
		delayedMessagesTargetDelay: delayedMessagesTargetDelay,
		sequencerInbox:             sequencerInbox,
		auth:                       transactAuth,
		chainTimeCheckInterval:     time.Second,
		feedBroadcaster:            broadcaster,
		dataSigner:                 dataSigner,
		maxDelayBlocks:             maxDelayBlocks,
		maxDelaySeconds:            maxDelaySeconds,

		// TODO make these configurable
		updateTimestampInterval:         big.NewInt(4),
		sequenceDelayedMessagesInterval: big.NewInt(20),
		createBatchBlockInterval:        createBatchBlockInterval,

		sequencer:              common.NewAddressFromEth(sequencer),
		signer:                 types.NewEIP155Signer(chainId),
		txQueue:                make(chan txQueueItem, 10),
		newTxFeed:              event.Feed{},
		latestChainTime:        chainTime,
		lastSequencedDelayedAt: chainTime.BlockNum.AsInt(),
		lastCreatedBatchAt:     chainTime.BlockNum.AsInt(),
	}, nil
}

func (b *SequencerBatcher) PendingTransactionCount(_ context.Context, _ common.Address) *uint64 {
	return nil
}

func (b *SequencerBatcher) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}

const maxExcludeComputation int64 = 10_000

func shouldIncludeTxResult(txRes *evm.TxResult) bool {
	if txRes == nil {
		// Tx receipt not found
		return false
	}
	if txRes.ResultCode == evm.ReturnCode {
		return true
	}
	if txRes.ResultCode == evm.RevertCode {
		// Still include computations taking up a lot of gas to avoid DoS
		return txRes.FeeStats.Paid.L2Computation.Cmp(big.NewInt(maxExcludeComputation)) > 0
	}
	// Other failure (probably not enough ETH balance)
	return false
}

func txLogsToResults(logs []value.Value) (map[common.Hash]*evm.TxResult, error) {
	resMap := make(map[common.Hash]*evm.TxResult)
	for _, log := range logs {
		res, err := evm.NewResultFromValue(log)
		if err != nil {
			return nil, err
		}
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			continue
		}
		resMap[txRes.IncomingRequest.MessageID] = txRes
	}
	return resMap, nil
}

func (b *SequencerBatcher) SendTransaction(_ context.Context, startTx *types.Transaction) error {
	_, err := types.Sender(b.signer, startTx)
	if err != nil {
		logger.Warn().Err(err).Msg("error processing user transaction")
		return err
	}
	logger.Info().Str("hash", startTx.Hash().String()).Msg("got user tx")
	startResultChan := make(chan error, 1)
	b.txQueue <- txQueueItem{tx: startTx, resultChan: startResultChan}
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	if len(startResultChan) > 0 {
		// startTx was already picked up by another thread
		err = <-startResultChan
		if err != nil {
			core.WaitForMachineIdle(b.db)
		}
		return err
	}

	var batchTxs []*types.Transaction
	var resultChans []chan error
	var l2BatchContents []message.AbstractL2Message
	seenOwnTx := false
	// This pattern is safe as we acquired a lock so we are the exclusive reader
	for len(b.txQueue) > 0 {
		queueItem := <-b.txQueue
		if queueItem.tx == startTx {
			seenOwnTx = true
		}
		batchTxs = append(batchTxs, queueItem.tx)
		resultChans = append(resultChans, queueItem.resultChan)
		l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(queueItem.tx))
	}
	if !seenOwnTx {
		// Another thread must have encountered an internal error attempting to process startTx
		// Let's try again ourselves (if we fail this time we won't try again)
		batchTxs = append(batchTxs, startTx)
		resultChans = append(resultChans, startResultChan)
		l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(startTx))
	}
	logger.Info().Int("count", len(l2BatchContents)).Msg("gather user txes")

	msgCount, err := b.db.GetMessageCount()
	if err != nil {
		return err
	}
	var prevAcc common.Hash
	if msgCount.Cmp(big.NewInt(0)) > 0 {
		prevAcc, err = b.db.GetInboxAcc(new(big.Int).Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return err
		}
	}
	originalAcc := prevAcc
	totalDelayedCount, err := b.db.GetTotalDelayedMessagesSequenced()
	if err != nil {
		return err
	}
	if totalDelayedCount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("chain not yet initialized")
	}

	batch, err := message.NewTransactionBatchFromMessages(l2BatchContents)
	if err != nil {
		return err
	}
	l2Message := message.NewSafeL2Message(batch)
	seqMsg := message.NewInboxMessage(l2Message, b.sequencer, new(big.Int).Set(msgCount), big.NewInt(0), b.latestChainTime.Clone())

	logCount, err := b.db.GetLogCount()
	if err != nil {
		return err
	}

	txBatchItem := inbox.NewSequencerItem(totalDelayedCount, seqMsg, prevAcc)
	err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, []inbox.SequencerBatchItem{txBatchItem}, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}
	core.WaitForMachineIdle(b.db)

	var sequencedTxs []*types.Transaction
	var sequencedBatchItems []inbox.SequencerBatchItem

	newLogCount, err := b.db.GetLogCount()
	if err != nil {
		return err
	}
	txLogs, err := b.db.GetLogs(logCount, new(big.Int).Sub(newLogCount, logCount))
	if err != nil {
		return err
	}
	txResults, err := txLogsToResults(txLogs)
	if err != nil {
		return err
	}

	txHashes := make([]common.Hash, 0, len(batchTxs))
	for _, tx := range batchTxs {
		txHashes = append(txHashes, common.NewHashFromEth(tx.Hash()))
	}

	successCount := 0
	for _, hash := range txHashes {
		if shouldIncludeTxResult(txResults[hash]) {
			successCount++
		}
	}
	if successCount == len(batchTxs) {
		sequencedTxs = batchTxs
		msgCount = new(big.Int).Add(msgCount, big.NewInt(1))
		prevAcc = txBatchItem.Accumulator
		sequencedBatchItems = append(sequencedBatchItems, txBatchItem)
		for _, c := range resultChans {
			c <- nil
		}
	} else {
		// Reorg to before we processed the batch and re-process the messages individually
		err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, nil, nil, msgCount)
		if err != nil {
			return err
		}
		core.WaitForMachineIdle(b.db)
		if successCount == 0 {
			// All of the transactions failed
			for i, c := range resultChans {
				c <- evm.HandleCallError(txResults[txHashes[i]], false)
			}
			return <-startResultChan
		}
		// At least one of the transactions failed and one of the transactions succeeded
		for i, tx := range batchTxs {
			txHash := txHashes[i]
			if !shouldIncludeTxResult(txResults[txHash]) {
				resultChans[i] <- evm.HandleCallError(txResults[txHash], false)
				continue
			}
			l2Msg := message.NewCompressedECDSAFromEth(tx)
			batch, err = message.NewTransactionBatchFromMessages([]message.AbstractL2Message{l2Msg})
			if err != nil {
				return err
			}
			l2Message := message.NewSafeL2Message(batch)
			seqMsg := message.NewInboxMessage(l2Message, b.sequencer, new(big.Int).Set(msgCount), big.NewInt(0), b.latestChainTime.Clone())
			txBatchItem := inbox.NewSequencerItem(totalDelayedCount, seqMsg, prevAcc)
			err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, []inbox.SequencerBatchItem{txBatchItem}, []inbox.DelayedMessage{}, nil)
			if err != nil {
				return err
			}
			core.WaitForMachineIdle(b.db)
			newLogCount, err = b.db.GetLogCount()
			if err != nil {
				return err
			}
			txLogs, err = b.db.GetLogs(logCount, new(big.Int).Sub(newLogCount, logCount))
			if err != nil {
				return err
			}
			newTxResults, err := txLogsToResults(txLogs)
			if err != nil {
				return err
			}
			txResult := newTxResults[txHash]
			if !shouldIncludeTxResult(txResult) {
				err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, nil, nil, msgCount)
				if err != nil {
					return err
				}
				resultChans[i] <- evm.HandleCallError(txResult, false)
				continue
			}
			msgCount = new(big.Int).Add(msgCount, big.NewInt(1))
			prevAcc = txBatchItem.Accumulator
			sequencedBatchItems = append(sequencedBatchItems, txBatchItem)
			sequencedTxs = append(sequencedTxs, tx)
			logCount = newLogCount
			resultChans[i] <- nil
		}
	}

	newBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		b.sequencer,
		new(big.Int).Set(msgCount),
		big.NewInt(0),
		b.latestChainTime.Clone(),
	)

	newBlockBatchItem := inbox.NewSequencerItem(totalDelayedCount, newBlockMessage, prevAcc)
	sequencedBatchItems = append(sequencedBatchItems, newBlockBatchItem)
	err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, []inbox.SequencerBatchItem{newBlockBatchItem}, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}

	if b.feedBroadcaster != nil {
		err = b.feedBroadcaster.Broadcast(originalAcc, sequencedBatchItems, b.dataSigner)
		if err != nil {
			return err
		}
	}

	core.WaitForMachineIdle(b.db)

	b.newTxFeed.Send(ethcore.NewTxsEvent{Txs: sequencedTxs})
	return <-startResultChan
}

func (b *SequencerBatcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	// TODO: return latest machine state?
	return nil, nil
}

func (b *SequencerBatcher) Aggregator() *common.Address {
	return &b.sequencer
}

func (b *SequencerBatcher) deliverDelayedMessages(chainTime inbox.ChainTime) (bool, error) {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()
	msgCount, err := b.db.GetMessageCount()
	if err != nil {
		return false, err
	}
	oldDelayedCount, err := b.db.GetTotalDelayedMessagesSequenced()
	if err != nil {
		return false, err
	}
	newDelayedCount, err := b.db.GetDelayedMessagesToSequence(new(big.Int).Sub(chainTime.BlockNum.AsInt(), b.delayedMessagesTargetDelay))
	if err != nil {
		return false, err
	}
	if newDelayedCount.Cmp(oldDelayedCount) <= 0 {
		return false, nil
	}

	delayedRead := new(big.Int).Sub(newDelayedCount, oldDelayedCount)
	newMsgCount := new(big.Int).Add(msgCount, delayedRead)
	newMsgCount.Add(newMsgCount, big.NewInt(1)) // end of block message
	lastSeqNum := new(big.Int).Sub(newMsgCount, big.NewInt(2))

	var prevAcc common.Hash
	if msgCount.Cmp(big.NewInt(0)) > 0 {
		prevAcc, err = b.db.GetInboxAcc(new(big.Int).Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return false, err
		}
	}
	delayedAcc, err := b.db.GetDelayedInboxAcc(new(big.Int).Sub(newDelayedCount, big.NewInt(1)))
	if err != nil {
		return false, err
	}
	batchItem := inbox.NewDelayedItem(lastSeqNum, newDelayedCount, prevAcc, oldDelayedCount, delayedAcc)
	logger.Info().
		Str("old", oldDelayedCount.String()).
		Str("new", newDelayedCount.String()).
		Msg("Adding messages from delayed queue")

	endOfBlockSeqNum := new(big.Int).Add(lastSeqNum, big.NewInt(1))
	endOfBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		endOfBlockSeqNum,
		big.NewInt(0),
		b.latestChainTime.Clone(),
	)
	endBlockBatchItem := inbox.NewSequencerItem(newDelayedCount, endOfBlockMessage, batchItem.Accumulator)
	seqBatchItems := []inbox.SequencerBatchItem{batchItem, endBlockBatchItem}
	err = core.DeliverMessagesAndWait(b.db, msgCount, prevAcc, seqBatchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return false, err
	}

	if b.feedBroadcaster != nil {
		err = b.feedBroadcaster.Broadcast(prevAcc, seqBatchItems, b.dataSigner)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

const gasCostBase int = 70292
const gasCostDelayedMessages int = 63505
const gasCostPerMessage int = 1431
const gasCostPerMessageByte int = 16
const gasCostMaximum int = 2_000_000

func (b *SequencerBatcher) createBatch(ctx context.Context, dontPublishBlockNum *big.Int) (bool, error) {
	prevMsgCount, err := b.sequencerInbox.MessageCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, err
	}
	batchItems, err := b.db.GetSequencerBatchItems(prevMsgCount)
	if err != nil {
		return false, err
	}
	if len(batchItems) == 0 {
		return true, nil
	}

	// Check if we need to reorg because we've exceeded the window
	lastBatchItem := batchItems[len(batchItems)-1]
	lastSeqMsg, err := inbox.NewInboxMessageFromData(lastBatchItem.SequencerMessage)
	if err != nil {
		return false, err
	}
	newestChainTime, err := getChainTime(ctx, b.client)
	if err != nil {
		return false, err
	}
	delayBlocks := new(big.Int).Sub(newestChainTime.BlockNum.AsInt(), lastSeqMsg.ChainTime.BlockNum.AsInt())
	delaySeconds := new(big.Int).Sub(newestChainTime.Timestamp, lastSeqMsg.ChainTime.Timestamp)
	if delayBlocks.Cmp(b.maxDelayBlocks) > 0 || delaySeconds.Cmp(b.maxDelaySeconds) > 0 {
		logger.Error().Str("delayBlocks", delayBlocks.String()).Str("delaySeconds", delaySeconds.String()).Msg("Exceeded max sequencer delay! Reorganizing to compensate...")

		b.reorgToNewTimestamp(ctx, prevMsgCount, newestChainTime)

		return false, errors.New("exceeded max sequencer delay, reorganized to compensate")
	}

	var transactionsData []byte
	var transactionsLengths []*big.Int
	var metadata []*big.Int
	var startDelayedMessagesRead *big.Int
	var totalDelayedMessagesRead *big.Int
	var l1BlockNumber *big.Int
	var l1Timestamp *big.Int
	var lastAcc common.Hash
	var lastSeqNum *big.Int
	estimatedGasCost := gasCostBase
	skippingImplicitEndOfBlock := false
	publishingAllBatchItems := true
	lastMetadataEnd := 0
	for i, item := range batchItems {
		var seqMsg inbox.InboxMessage
		if len(item.SequencerMessage) > 0 {
			seqMsg, err = inbox.NewInboxMessageFromData(item.SequencerMessage)
			if err != nil {
				return false, err
			}

			estimatedGasCost += gasCostPerMessage + gasCostPerMessageByte*len(seqMsg.Data)
		} else {
			estimatedGasCost += gasCostDelayedMessages
		}
		if i != 0 && estimatedGasCost >= gasCostMaximum && !skippingImplicitEndOfBlock {
			publishingAllBatchItems = false
			break
		}

		if startDelayedMessagesRead == nil {
			startDelayedMessagesRead = item.TotalDelayedCount
		} else if totalDelayedMessagesRead != nil && !skippingImplicitEndOfBlock {
			delayedAcc, err := b.db.GetDelayedInboxAcc(new(big.Int).Sub(totalDelayedMessagesRead, big.NewInt(1)))
			if err != nil {
				return false, err
			}
			delayedAccInt := new(big.Int).SetBytes(delayedAcc.Bytes())
			sectionCount := big.NewInt(int64(len(transactionsLengths) - lastMetadataEnd))
			metadata = append(metadata, sectionCount, l1BlockNumber, l1Timestamp, totalDelayedMessagesRead, delayedAccInt)
			lastMetadataEnd = len(transactionsLengths)
			l1BlockNumber = nil
			l1Timestamp = nil
			startDelayedMessagesRead = nil
			totalDelayedMessagesRead = nil
		}

		if len(item.SequencerMessage) == 0 {
			if skippingImplicitEndOfBlock {
				return false, errors.New("back-to-back delayed messages inserted without end of block")
			}
			totalDelayedMessagesRead = item.TotalDelayedCount
			skippingImplicitEndOfBlock = true
		} else {
			if dontPublishBlockNum != nil && seqMsg.ChainTime.BlockNum.AsInt().Cmp(dontPublishBlockNum) >= 0 && !skippingImplicitEndOfBlock {
				break
			}
			if l1BlockNumber == nil || l1BlockNumber.Cmp(seqMsg.ChainTime.BlockNum.AsInt()) != 0 || l1Timestamp.Cmp(seqMsg.ChainTime.Timestamp) != 0 {
				sectionCount := len(transactionsLengths) - lastMetadataEnd
				if sectionCount > 0 {
					metadata = append(metadata, big.NewInt(int64(sectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
					lastMetadataEnd = len(transactionsLengths)
				}
				l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
				l1Timestamp = seqMsg.ChainTime.Timestamp
			}

			// Do some basic validation of the message
			if seqMsg.Kind == message.EndOfBlockType {
				if len(seqMsg.Data) != 0 {
					return false, errors.New("end of block message has data")
				}
			} else if seqMsg.Kind != message.L2Type {
				return false, errors.Errorf("unexpected sequencer message kind %v", seqMsg.Kind)
			}

			if skippingImplicitEndOfBlock {
				if seqMsg.Kind != message.EndOfBlockType {
					return false, errors.New("found non-end-of-block sequencer message after delayed messages")
				}
				skippingImplicitEndOfBlock = false
			} else {
				transactionsData = append(transactionsData, seqMsg.Data...)
				transactionsLengths = append(transactionsLengths, big.NewInt(int64(len(seqMsg.Data))))
			}
		}
		lastAcc = item.Accumulator
		lastSeqNum = item.LastSeqNum
	}
	if lastSeqNum == nil {
		return true, nil
	}
	if skippingImplicitEndOfBlock {
		return false, errors.New("didn't find implicit end of block after delayed messages")
	}

	lastSectionCount := len(transactionsLengths) - lastMetadataEnd
	if lastSectionCount > 0 || totalDelayedMessagesRead != nil {
		var delayedAcc common.Hash
		if totalDelayedMessagesRead != nil {
			delayedAcc, err = b.db.GetDelayedInboxAcc(new(big.Int).Sub(totalDelayedMessagesRead, big.NewInt(1)))
			if err != nil {
				return false, err
			}
		} else {
			totalDelayedMessagesRead = startDelayedMessagesRead
		}
		delayedAccInt := new(big.Int).SetBytes(delayedAcc.Bytes())
		metadata = append(metadata, big.NewInt(int64(lastSectionCount)), l1BlockNumber, l1Timestamp, totalDelayedMessagesRead, delayedAccInt)
	}

	newMsgCount := new(big.Int).Add(lastSeqNum, big.NewInt(1))
	logger.Info().Str("prevMsgCount", prevMsgCount.String()).Int("items", len(batchItems)).Str("newMsgCount", newMsgCount.String()).Msg("Creating sequencer batch")
	tx, err := ethbridge.AddSequencerL2BatchFromOrigin(ctx, b.sequencerInbox, b.auth, transactionsData, transactionsLengths, metadata, lastAcc)
	if err != nil {
		return false, err
	}

	receipt, err := ethbridge.WaitForReceiptWithResults(ctx, b.client, b.sequencer.ToEthAddress(), tx, "addSequencerL2BatchFromOrigin")
	if err != nil {
		return false, err
	}

	if b.feedBroadcaster != nil {
		// Confirm feed messages that are already on chain
		err = b.feedBroadcaster.ConfirmedAccumulator(lastAcc)
		if err != nil {
			return false, err
		}
	}

	if b.logBatchGasCosts {
		fmt.Printf("%v,%v,%v\n", len(transactionsLengths), len(transactionsData), receipt.GasUsed)
	}

	return publishingAllBatchItems, nil
}

func (b *SequencerBatcher) reorgToNewTimestamp(ctx context.Context, prevMsgCount *big.Int, newChainTime inbox.ChainTime) error {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	batchItems, err := b.db.GetSequencerBatchItems(prevMsgCount)
	if err != nil {
		return err
	}

	var previousSeqBatchAcc common.Hash
	if prevMsgCount.Cmp(big.NewInt(0)) > 0 {
		previousSeqBatchAcc, err = b.db.GetInboxAcc(new(big.Int).Sub(prevMsgCount, big.NewInt(1)))
		if err != nil {
			return err
		}
	}
	for i := range batchItems {
		item := batchItems[i]
		if len(item.SequencerMessage) == 0 {
			item.Accumulator = common.Hash{}
		} else {
			seqMsg, err := inbox.NewInboxMessageFromData(item.SequencerMessage)
			if err != nil {
				return err
			}
			seqMsg.ChainTime = newChainTime
			item.SequencerMessage = seqMsg.ToBytes()
			item.Accumulator = common.Hash{}
		}
		batchItems[i] = item
	}
	err = core.DeliverMessagesAndWait(b.db, prevMsgCount, previousSeqBatchAcc, batchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (b *SequencerBatcher) Start(ctx context.Context) {
	logger.Log().Msg("Starting sequencer batch submission thread")
	firstBoot := true
	if b.feedBroadcaster != nil {
		defer b.feedBroadcaster.Stop()
	}

	for {
		select {
		case <-ctx.Done():
			logger.Log().Msg("exiting sequencer since context was canceled")
			return
		default:
		}
		time.Sleep(b.chainTimeCheckInterval)
		chainTime, err := getChainTime(ctx, b.client)
		if err != nil {
			logger.Warn().Err(err).Msg("Error getting chain time")
			continue
		}
		blockNum := chainTime.BlockNum.AsInt()
		targetCreateBatch := new(big.Int).Add(b.lastCreatedBatchAt, b.createBatchBlockInterval)
		creatingBatch := blockNum.Cmp(targetCreateBatch) >= 0
		targetSequenceDelayed := new(big.Int).Add(b.lastSequencedDelayedAt, b.sequenceDelayedMessagesInterval)
		sequencedDelayed := false
		if blockNum.Cmp(targetSequenceDelayed) >= 0 || creatingBatch || firstBoot {
			sequencedDelayed, err = b.deliverDelayedMessages(chainTime)
			if err != nil {
				logger.Error().Err(err).Msg("Error delivering delayed messages")
				continue
			}
			b.lastSequencedDelayedAt = blockNum
		}
		targetUpdateTime := new(big.Int).Add(b.latestChainTime.BlockNum.AsInt(), b.updateTimestampInterval)
		var dontPublishBlockNum *big.Int
		if blockNum.Cmp(targetUpdateTime) >= 0 || creatingBatch || sequencedDelayed || firstBoot {
			b.inboxReader.MessageDeliveryMutex.Lock()
			b.latestChainTime = chainTime
			// Avoid inefficency of publishing something that just got put in this timestamp
			dontPublishBlockNum = b.latestChainTime.BlockNum.AsInt()
			b.inboxReader.MessageDeliveryMutex.Unlock()
		}
		if creatingBatch || firstBoot {
			for {
				complete, err := b.createBatch(ctx, dontPublishBlockNum)
				if err == nil {
					if complete {
						time.Sleep(10 * b.chainTimeCheckInterval)
						break
					} else {
						time.Sleep(time.Second)
					}
				} else {
					logger.Error().Err(err).Msg("Error creating batch")
					time.Sleep(5 * time.Second)
				}
				dontPublishBlockNum = nil
			}
			b.lastCreatedBatchAt = blockNum
		}
		firstBoot = false
	}
}
