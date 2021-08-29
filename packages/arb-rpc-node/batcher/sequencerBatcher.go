/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type SequencerLockoutManager interface {
	ShouldSequence() bool
}

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
	auth                            transactauth.TransactAuth
	fromAddress                     common.Address
	chainTimeCheckInterval          time.Duration
	logBatchGasCosts                bool
	feedBroadcaster                 *broadcaster.Broadcaster
	dataSigner                      func([]byte) ([]byte, error)
	maxDelayBlocks                  *big.Int
	maxDelaySeconds                 *big.Int
	updateTimestampInterval         *big.Int
	sequenceDelayedMessagesInterval *big.Int
	createBatchBlockInterval        *big.Int
	LockoutManager                  SequencerLockoutManager
	config                          *configuration.Config
	fb                              *fireblocks.Fireblocks
	consecutiveShouldReorgGaps      int

	signer    types.Signer
	txQueue   chan txQueueItem
	newTxFeed event.Feed

	latestChainTime        inbox.ChainTime
	lastCreatedBatchAt     *big.Int
	lastSequencedDelayedAt *big.Int
	// 1 if we've published a batch to the L1 mempool,
	// but it hasn't been included in an L1 block yet.
	publishingBatchAtomic int32
	// The total estimate of unpublished transactions' gas usage.
	// Added to every time something is sequenced, zeroed when batch posted.
	pendingBatchGasEstimateAtomic int64
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
	sequencerInbox *ethbridgecontracts.SequencerInbox,
	auth *bind.TransactOpts,
	dataSigner func([]byte) ([]byte, error),
	broadcaster *broadcaster.Broadcaster,
	config *configuration.Config,
	walletConfig *configuration.Wallet,
) (*SequencerBatcher, error) {
	chainTime, err := getChainTime(ctx, client)
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: ctx}
	isSequencer, err := sequencerInbox.IsSequencer(callOpts, auth.From)
	if err != nil {
		return nil, err
	}
	if !isSequencer {
		logger.
			Error().
			Hex("current-from", auth.From.Bytes()).
			Msg("Transaction auth isn't for sequencer")
		return nil, errors.New("Transaction auth address isn't for sequencer")
	}

	var transactAuth transactauth.TransactAuth
	var fb *fireblocks.Fireblocks
	if len(walletConfig.Fireblocks.SSLKey) > 0 {
		transactAuth, fb, err = transactauth.NewFireblocksTransactAuth(ctx, client, auth, walletConfig)
	} else {
		transactAuth, err = transactauth.NewTransactAuth(ctx, client, auth)

	}
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

	if config.Node.Sequencer.CreateBatchBlockInterval <= 0 || config.Node.Sequencer.CreateBatchBlockInterval > maxDelayBlocks.Int64() {
		return nil, errors.New("invalid batch creation block interval")
	}

	batcher := &SequencerBatcher{
		db:                         db,
		inboxReader:                inboxReader,
		client:                     client,
		delayedMessagesTargetDelay: big.NewInt(config.Node.Sequencer.DelayedMessagesTargetDelay),
		sequencerInbox:             sequencerInbox,
		auth:                       transactAuth,
		fromAddress:                common.NewAddressFromEth(auth.From),
		chainTimeCheckInterval:     time.Second,
		feedBroadcaster:            broadcaster,
		dataSigner:                 dataSigner,
		maxDelayBlocks:             maxDelayBlocks,
		maxDelaySeconds:            maxDelaySeconds,
		config:                     config,

		// TODO make these configurable
		updateTimestampInterval:         big.NewInt(4),
		sequenceDelayedMessagesInterval: big.NewInt(20),
		createBatchBlockInterval:        big.NewInt(config.Node.Sequencer.CreateBatchBlockInterval),

		signer:                        types.NewEIP155Signer(chainId),
		txQueue:                       make(chan txQueueItem, 10),
		newTxFeed:                     event.Feed{},
		latestChainTime:               chainTime,
		lastSequencedDelayedAt:        chainTime.BlockNum.AsInt(),
		lastCreatedBatchAt:            chainTime.BlockNum.AsInt(),
		publishingBatchAtomic:         0,
		pendingBatchGasEstimateAtomic: int64(gasCostBase),
		fb:                            fb,
	}

	return batcher, nil
}

func (b *SequencerBatcher) PendingTransactionCount(_ context.Context, _ common.Address) (*uint64, error) {
	return nil, nil
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

const maxTxDataSize int = 100_000

func (b *SequencerBatcher) SendTransaction(ctx context.Context, startTx *types.Transaction) error {
	_, err := types.Sender(b.signer, startTx)
	if err != nil {
		logger.Warn().Err(err).Msg("error processing user transaction")
		return err
	}
	if len(startTx.Data()) > maxTxDataSize {
		return errors.New("oversized data")
	}
	logger.Info().Str("hash", startTx.Hash().String()).Msg("got user tx")

	startResultChan := make(chan error, 1)
	b.txQueue <- txQueueItem{tx: startTx, resultChan: startResultChan}
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	if b.LockoutManager != nil && !b.LockoutManager.ShouldSequence() {
		return errors.New("sequencer missing lockout")
	}

	if len(startResultChan) > 0 {
		// startTx was already picked up by another thread
		err = <-startResultChan
		if err != nil {
			core.WaitForMachineIdle(b.db)
		}
		return err
	}

	for {
		var batchTxs []*types.Transaction
		var resultChans []chan error
		var l2BatchContents []message.AbstractL2Message
		var batchDataSize int
		seenOwnTx := false
		emptiedQueue := true
		// This pattern is safe as we acquired a lock so we are the exclusive reader
		for len(b.txQueue) > 0 {
			queueItem := <-b.txQueue
			if batchDataSize+len(queueItem.tx.Data()) > maxTxDataSize {
				// This batch would be too large to publish with this tx added.
				// Put the tx back in the queue so it can be included later.
				b.txQueue <- queueItem
				emptiedQueue = false
				break
			}
			if queueItem.tx == startTx {
				seenOwnTx = true
			}
			batchTxs = append(batchTxs, queueItem.tx)
			resultChans = append(resultChans, queueItem.resultChan)
			l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(queueItem.tx))
			batchDataSize += len(queueItem.tx.Data())
		}
		if !seenOwnTx && emptiedQueue && batchDataSize+len(startTx.Data()) <= maxTxDataSize {
			// Another thread must have encountered an internal error attempting to process startTx
			// Let's try again ourselves (if we fail this time we won't try again)
			batchTxs = append(batchTxs, startTx)
			resultChans = append(resultChans, startResultChan)
			l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(startTx))
			batchDataSize += len(startTx.Data())
			seenOwnTx = true
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
		seqMsg := message.NewInboxMessage(l2Message, b.fromAddress, new(big.Int).Set(msgCount), big.NewInt(0), b.latestChainTime.Clone())

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
			postingCostEstimate := gasCostPerMessage + gasCostPerMessageByte*len(seqMsg.Data)
			atomic.AddInt64(&b.pendingBatchGasEstimateAtomic, int64(postingCostEstimate))
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
				seqMsg := message.NewInboxMessage(l2Message, b.fromAddress, new(big.Int).Set(msgCount), big.NewInt(0), b.latestChainTime.Clone())
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
				postingCostEstimate := gasCostPerMessage + gasCostPerMessageByte*len(seqMsg.Data)
				atomic.AddInt64(&b.pendingBatchGasEstimateAtomic, int64(postingCostEstimate))
				logCount = newLogCount
				resultChans[i] <- nil
			}
		}

		newBlockMessage := message.NewInboxMessage(
			message.EndBlockMessage{},
			b.fromAddress,
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
		atomic.AddInt64(&b.pendingBatchGasEstimateAtomic, int64(gasCostPerMessage))

		if b.feedBroadcaster != nil {
			err = b.feedBroadcaster.Broadcast(originalAcc, sequencedBatchItems, b.dataSigner)
			if err != nil {
				return err
			}
		}

		core.WaitForMachineIdle(b.db)

		b.newTxFeed.Send(ethcore.NewTxsEvent{Txs: sequencedTxs})

		if seenOwnTx {
			break
		}
	}

	return <-startResultChan
}

func (b *SequencerBatcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	// TODO: return latest machine state?
	return nil, nil
}

func (b *SequencerBatcher) Aggregator() *common.Address {
	return &b.fromAddress
}

func (b *SequencerBatcher) deliverDelayedMessages(ctx context.Context, chainTime inbox.ChainTime, bypassLockout bool) (bool, error) {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()
	if !bypassLockout && b.LockoutManager != nil && !b.LockoutManager.ShouldSequence() {
		return false, errors.New("sequencer lockout missing")
	}
	msgCount, err := b.db.GetMessageCount()
	if err != nil {
		return false, err
	}
	oldDelayedCount, err := b.db.GetTotalDelayedMessagesSequenced()
	if err != nil {
		return false, err
	}
	lastConfirmedL1Block := new(big.Int).Sub(chainTime.BlockNum.AsInt(), b.delayedMessagesTargetDelay)
	newDelayedCount, err := b.db.GetDelayedMessagesToSequence(lastConfirmedL1Block)
	if err != nil {
		return false, err
	}
	if newDelayedCount.Cmp(oldDelayedCount) <= 0 {
		logger.Debug().Str("delayedCount", oldDelayedCount.String()).Msg("no delayed messages to sequence")
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
	lastDelayedSeqNum := new(big.Int).Sub(newDelayedCount, big.NewInt(1))
	delayedAcc, err := b.db.GetDelayedInboxAcc(lastDelayedSeqNum)
	if err != nil {
		return false, err
	}
	_, isSimulatedBackend := b.client.(*ethutils.SimulatedEthClient)
	var getDelayedAccTarget *big.Int
	if isSimulatedBackend {
		// The simulated backend doesn't support querying against old blocks
		getDelayedAccTarget = chainTime.BlockNum.AsInt()
	} else {
		// Confirm that the message wasn't reorganized forwards to a further reorganizable block
		getDelayedAccTarget = lastConfirmedL1Block
	}
	l1DelayedAcc, err := b.inboxReader.GetDelayedAccumulator(ctx, lastDelayedSeqNum, getDelayedAccTarget)
	if err != nil {
		return false, err
	}
	if delayedAcc != l1DelayedAcc {
		return false, errors.New("inbox reader missed delayed inbox reorg")
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
	postingCostEstimate := gasCostPerMessage + gasCostDelayedMessages
	atomic.AddInt64(&b.pendingBatchGasEstimateAtomic, int64(postingCostEstimate))

	if b.feedBroadcaster != nil {
		err = b.feedBroadcaster.Broadcast(prevAcc, seqBatchItems, b.dataSigner)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// Warning: bypassLockout should only be used if the lockout manager itself is calling this
func (b *SequencerBatcher) SequenceDelayedMessages(ctx context.Context, bypassLockout bool) error {
	chainTime, err := getChainTime(ctx, b.client)
	if err != nil {
		return err
	}
	_, err = b.deliverDelayedMessages(ctx, chainTime, bypassLockout)
	return err
}

const gasCostBase int = 70292
const gasCostDelayedMessages int = 63505
const gasCostPerMessage int = 1431
const gasCostPerMessageByte int = 16
const gasCostMaximum int = 2_000_000

// Wait this long after batch confirmation before publishing a new batch
const l1RacePrevention time.Duration = time.Second * 10
const maxL1BackwardsReorg int64 = 12

// Updates both prevMsgCount and nonce on success
func (b *SequencerBatcher) publishBatch(ctx context.Context, dontPublishBlockNum *big.Int, prevMsgCount *big.Int, nonce *big.Int) (bool, error) {
	b.inboxReader.MessageDeliveryMutex.Lock()
	batchItems, err := b.db.GetSequencerBatchItems(prevMsgCount)
	origEstimate := atomic.LoadInt64(&b.pendingBatchGasEstimateAtomic)
	b.inboxReader.MessageDeliveryMutex.Unlock()
	if err != nil {
		return false, err
	}
	if len(batchItems) == 0 {
		return true, nil
	}

	if len(batchItems[0].SequencerMessage) >= 128*1024 {
		logger.Error().Int("size", len(batchItems[0].SequencerMessage)).Msg("Sequencer batch item is too big!")
		if b.config.Node.Sequencer.ReorgOutHugeMessages {
			err = b.reorgOutHugeMsg(ctx, prevMsgCount)
			if err != nil {
				return false, err
			}
			return false, errors.New("reorganized out huge message")
		}
	}

	var transactionsData []byte
	var transactionsLengths []*big.Int
	var metadata []*big.Int
	var startDelayedMessagesRead *big.Int
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

			if seqMsg.Sender != (common.Address{}) && seqMsg.Sender != b.fromAddress {
				msg := "sequencer message in database contains messages from another sequencer"
				if b.config.Node.Sequencer.RewriteSequencerAddress {
					msg += "! Reorganizing to compensate..."
				}

				logger.
					Warn().
					Str("messageAddress", seqMsg.Sender.String()).
					Str("fromAddress", b.fromAddress.String()).
					Str("sequenceNumber", seqMsg.InboxSeqNum.String()).
					Msg(msg)

				if b.config.Node.Sequencer.RewriteSequencerAddress {
					b.reorgToNewSequencerAddress(ctx, prevMsgCount)

					return false, errors.New("reorganized to rewrite sequencer address")
				}

				if len(transactionsLengths) > 0 {
					// Still publish what we can
					break
				} else {
					return false, errors.New("sequencer message in database begins with messages from another sequencer")
				}
			}

			estimatedGasCost += gasCostPerMessage + gasCostPerMessageByte*len(seqMsg.Data)
		} else {
			estimatedGasCost += gasCostDelayedMessages
		}
		if i != 0 && estimatedGasCost >= gasCostMaximum && !skippingImplicitEndOfBlock {
			publishingAllBatchItems = false
			break
		}

		mustEndSectionAfterItem := skippingImplicitEndOfBlock
		if len(item.SequencerMessage) == 0 {
			if skippingImplicitEndOfBlock {
				return false, errors.New("back-to-back delayed messages inserted without end of block")
			}
			skippingImplicitEndOfBlock = true
		} else {
			if dontPublishBlockNum != nil && seqMsg.ChainTime.BlockNum.AsInt().Cmp(dontPublishBlockNum) >= 0 && !skippingImplicitEndOfBlock {
				break
			}
			if l1BlockNumber == nil {
				l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
				l1Timestamp = seqMsg.ChainTime.Timestamp
				startDelayedMessagesRead = item.TotalDelayedCount
			} else if l1BlockNumber.Cmp(seqMsg.ChainTime.BlockNum.AsInt()) != 0 || l1Timestamp.Cmp(seqMsg.ChainTime.Timestamp) != 0 {
				sectionCount := len(transactionsLengths) - lastMetadataEnd
				metadata = append(metadata, big.NewInt(int64(sectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
				lastMetadataEnd = len(transactionsLengths)

				l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
				l1Timestamp = seqMsg.ChainTime.Timestamp
				startDelayedMessagesRead = item.TotalDelayedCount
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

		if mustEndSectionAfterItem {
			delayedAcc, err := b.db.GetDelayedInboxAcc(new(big.Int).Sub(item.TotalDelayedCount, big.NewInt(1)))
			if err != nil {
				return false, err
			}
			delayedAccInt := new(big.Int).SetBytes(delayedAcc.Bytes())
			sectionCount := big.NewInt(int64(len(transactionsLengths) - lastMetadataEnd))
			metadata = append(metadata, sectionCount, l1BlockNumber, l1Timestamp, item.TotalDelayedCount, delayedAccInt)
			lastMetadataEnd = len(transactionsLengths)
			l1BlockNumber = nil
			l1Timestamp = nil
			startDelayedMessagesRead = nil
		}
	}
	if lastSeqNum == nil {
		return true, nil
	}
	if skippingImplicitEndOfBlock {
		return false, errors.New("didn't find implicit end of block after delayed messages")
	}

	lastSectionCount := len(transactionsLengths) - lastMetadataEnd
	if lastSectionCount > 0 {
		metadata = append(metadata, big.NewInt(int64(lastSectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
	}

	var minL1BlockNumber *big.Int
	var maxL1BlockNumber *big.Int
	var minL1Timestamp *big.Int
	var maxL1Timestamp *big.Int
	for i := 0; i < len(metadata); i += 5 {
		blockNum := metadata[i+1]
		timestamp := metadata[i+2]
		if minL1BlockNumber == nil || blockNum.Cmp(minL1BlockNumber) < 0 {
			minL1BlockNumber = blockNum
		}
		if maxL1BlockNumber == nil || blockNum.Cmp(maxL1BlockNumber) > 0 {
			maxL1BlockNumber = blockNum
		}
		if minL1Timestamp == nil || timestamp.Cmp(minL1Timestamp) < 0 {
			minL1Timestamp = timestamp
		}
		if maxL1Timestamp == nil || timestamp.Cmp(maxL1Timestamp) > 0 {
			maxL1Timestamp = timestamp
		}
	}

	newestChainTime, err := getChainTime(ctx, b.client)
	if err != nil {
		return false, err
	}
	// These should never be nil, but just for safety
	if minL1BlockNumber == nil {
		minL1BlockNumber = newestChainTime.BlockNum.AsInt()
	}
	if minL1Timestamp == nil {
		minL1Timestamp = newestChainTime.Timestamp
	}
	delayBlocks := new(big.Int).Sub(newestChainTime.BlockNum.AsInt(), minL1BlockNumber)
	delaySeconds := new(big.Int).Sub(newestChainTime.Timestamp, minL1Timestamp)
	aheadBlocks := new(big.Int).Sub(maxL1BlockNumber, newestChainTime.BlockNum.AsInt())
	aheadSeconds := new(big.Int).Sub(maxL1Timestamp, newestChainTime.Timestamp)
	if delayBlocks.Cmp(b.maxDelayBlocks) > 0 || delaySeconds.Cmp(b.maxDelaySeconds) > 0 || aheadBlocks.Cmp(big.NewInt(maxL1BackwardsReorg)) > 0 || aheadSeconds.Cmp(big.NewInt(maxL1BackwardsReorg*15)) > 0 {
		b.consecutiveShouldReorgGaps += 1
		doingReorg := b.consecutiveShouldReorgGaps >= 10

		var msg string
		var log *zerolog.Event
		if doingReorg {
			msg = "Exceeded max sequencer delay! Reorganizing to compensate..."
			log = logger.Error()
		} else {
			msg = "Exceeded max sequencer delay! Waiting to see if this is consistent..."
			log = logger.Warn()
		}

		log.
			Str("delayBlocks", delayBlocks.String()).
			Str("delaySeconds", delaySeconds.String()).
			Str("aheadBlocks", aheadBlocks.String()).
			Str("aheadSeconds", aheadSeconds.String()).
			Int("consecutiveGaps", b.consecutiveShouldReorgGaps).
			Msg(msg)

		if doingReorg {
			err = b.reorgToNewTimestamp(ctx, prevMsgCount, newestChainTime)
			if err != nil {
				return false, errors.Wrap(err, "error during reorg after exceeded max sequencer delay")
			}
			b.consecutiveShouldReorgGaps = 0
		}

		return false, errors.New("exceeded max sequencer delay")
	} else {
		b.consecutiveShouldReorgGaps = 0
	}

	newMsgCount := new(big.Int).Add(lastSeqNum, big.NewInt(1))
	logger.Info().Str("prevMsgCount", prevMsgCount.String()).Int("items", len(batchItems)).Str("newMsgCount", newMsgCount.String()).Msg("Creating sequencer batch")
	arbTx, err := ethbridge.AddSequencerL2BatchFromOriginCustomNonce(ctx, b.sequencerInbox, b.auth, nonce, transactionsData, transactionsLengths, metadata, lastAcc)
	if err != nil {
		return false, err
	}

	var removedPendingGasEstimate int64
	if publishingAllBatchItems {
		// Reset the pending gas estimate to gasCostBase.
		removedPendingGasEstimate = origEstimate
	} else {
		// Since we didn't publish everything, only subtract gas for what we did publish.
		removedPendingGasEstimate = int64(estimatedGasCost)
	}
	atomic.AddInt64(&b.pendingBatchGasEstimateAtomic, int64(gasCostBase)-removedPendingGasEstimate)

	// Update prevMsgCount for the next iteration if we're not publishingAllBatchItems
	// AddSequencerL2BatchFromOriginCustomNonce will have already updated the nonce
	prevMsgCount.Set(newMsgCount)

	atomic.StoreInt32(&b.publishingBatchAtomic, 1)
	go (func() {
		defer atomic.StoreInt32(&b.publishingBatchAtomic, 0)
		receipt, err := transactauth.WaitForReceiptWithResultsAndReplaceByFee(ctx, b.client, b.fromAddress.ToEthAddress(), arbTx, "addSequencerL2BatchFromOrigin", b.auth, b.auth)
		if err != nil {
			logger.Warn().Err(err).Msg("error waiting for batch receipt")
			return
		}

		if b.feedBroadcaster != nil {
			// Confirm feed messages that are already on chain
			b.feedBroadcaster.ConfirmedAccumulator(lastAcc)
		}

		if b.logBatchGasCosts {
			fmt.Printf("%v,%v,%v\n", len(transactionsLengths), len(transactionsData), receipt.GasUsed)
		}

		// Don't set publishingBatchAtomic to 0 until after this.
		// This prevents us from publishing the next batch too quickly.
		// If we don't have this, the MessageCount query might be out of date.
		time.Sleep(l1RacePrevention)
	})()

	return publishingAllBatchItems, nil
}

func (b *SequencerBatcher) reorgAndModifySequencerMessages(ctx context.Context, prevMsgCount *big.Int, modifier func(*inbox.InboxMessage)) error {
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
		item := &batchItems[i]
		if len(item.SequencerMessage) == 0 {
			item.Accumulator = common.Hash{}
		} else {
			seqMsg, err := inbox.NewInboxMessageFromData(item.SequencerMessage)
			if err != nil {
				return err
			}
			modifier(&seqMsg)
			item.SequencerMessage = seqMsg.ToBytes()
			item.Accumulator = common.Hash{}
		}
	}
	err = core.DeliverMessagesAndWait(b.db, prevMsgCount, previousSeqBatchAcc, batchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (b *SequencerBatcher) reorgToNewTimestamp(ctx context.Context, prevMsgCount *big.Int, newChainTime inbox.ChainTime) error {
	return b.reorgAndModifySequencerMessages(ctx, prevMsgCount, func(msg *inbox.InboxMessage) {
		msg.ChainTime = newChainTime
	})
}

func (b *SequencerBatcher) reorgToNewSequencerAddress(ctx context.Context, prevMsgCount *big.Int) error {
	return b.reorgAndModifySequencerMessages(ctx, prevMsgCount, func(msg *inbox.InboxMessage) {
		if msg.Sender != (common.Address{}) {
			msg.Sender = b.fromAddress
		}
	})
}

func (b *SequencerBatcher) reorgOutHugeMsg(ctx context.Context, prevMsgCount *big.Int) error {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	batchItems, err := b.db.GetSequencerBatchItems(new(big.Int).Add(prevMsgCount, big.NewInt(1)))
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
		item := &batchItems[i]
		item.LastSeqNum.Sub(item.LastSeqNum, big.NewInt(1))
		item.Accumulator = common.Hash{}
	}
	err = core.DeliverMessagesAndWait(b.db, prevMsgCount, previousSeqBatchAcc, batchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (b *SequencerBatcher) getLastSequencedChainTime() (inbox.ChainTime, error) {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()
	msgCount, err := b.db.GetMessageCount()
	if err != nil || msgCount.Sign() == 0 {
		return inbox.ChainTime{}, err
	}
	seqNum := new(big.Int).Sub(msgCount, big.NewInt(1))
	batchItems, err := b.db.GetSequencerBatchItems(seqNum)
	if err != nil {
		return inbox.ChainTime{}, err
	}
	if len(batchItems) > 0 {
		item := batchItems[len(batchItems)-1]
		if len(item.SequencerMessage) > 0 {
			seqMsg, err := inbox.NewInboxMessageFromData(item.SequencerMessage)
			if err != nil {
				return inbox.ChainTime{}, err
			}
			return seqMsg.ChainTime, nil
		} else {
			logger.Warn().Str("sequenceNumber", seqNum.String()).Msg("no sequencer message in last batch item on startup")
		}
	}
	return inbox.ChainTime{}, nil
}

func (b *SequencerBatcher) Start(ctx context.Context) {
	logger.Log().Msg("Starting sequencer batch submission thread")
	firstBatchCreation := true
	if b.feedBroadcaster != nil {
		defer b.feedBroadcaster.Stop()
	}

	var chainTime inbox.ChainTime
	for {
		var err error
		chainTime, err = b.getLastSequencedChainTime()
		if err != nil {
			logger.Warn().Err(err).Msg("error getting last sequenced chain time")
		} else {
			break
		}
		select {
		case <-ctx.Done():
			logger.Log().Msg("exiting sequencer since context was canceled")
			return
		case <-time.After(time.Second):
		}
	}
	for {
		select {
		case <-ctx.Done():
			logger.Log().Msg("exiting sequencer since context was canceled")
			return
		case <-time.After(b.chainTimeCheckInterval):
		}

		// Safely get the current chain time
		newChainTime, err := getChainTime(ctx, b.client)
		if err != nil {
			logger.Warn().Err(err).Msg("error getting chain time")
			continue
		}
		var chainTimeCmp int
		if chainTime.BlockNum != nil {
			chainTimeCmp = newChainTime.BlockNum.Cmp(chainTime.BlockNum)
		} else {
			// If we didn't have a previous chain time, we assume we're moving forwards
			chainTimeCmp = 1
		}
		if chainTimeCmp < 0 {
			logger.
				Warn().
				Str("oldBlockNumber", chainTime.BlockNum.String()).
				Str("newBlockNumber", newChainTime.BlockNum.String()).
				Msg("chain time moved backwards")
			continue
		} else if chainTimeCmp == 0 {
			// Chain time hasn't changed
			continue
		}
		chainTime = newChainTime
		blockNum := chainTime.BlockNum.AsInt()

		// Determine if we should create a batch
		shouldSequence := b.LockoutManager == nil || b.LockoutManager.ShouldSequence()
		targetCreateBatch := new(big.Int).Add(b.lastCreatedBatchAt, b.createBatchBlockInterval)
		creatingBatch := blockNum.Cmp(targetCreateBatch) >= 0 ||
			atomic.LoadInt64(&b.pendingBatchGasEstimateAtomic) >= int64(gasCostMaximum)*9/10 ||
			firstBatchCreation
		if creatingBatch && !shouldSequence && !b.config.Node.Sequencer.PublishBatchesWithoutLockout {
			// We don't have the lockout and publishing batches without the lockout is disabled
			creatingBatch = false
		}
		if creatingBatch && atomic.LoadInt32(&b.publishingBatchAtomic) != 0 {
			// The previous batch is still waiting on confirmation; don't attempt to create another yet
			creatingBatch = false
		}
		if creatingBatch && blockNum.Cmp(new(big.Int).Add(targetCreateBatch, big.NewInt(b.config.Node.Sequencer.L1PostingStrategy.HighGasDelayBlocks))) < 0 {
			// Check if gas price is too high, and if so, hold off on creating a batch
			gasPrice, err := b.client.SuggestGasPrice(ctx)
			if err != nil {
				logger.Warn().Err(err).Msg("error getting gas price")
			} else {
				gasPriceFloat := float64(gasPrice.Int64()) / 1e9
				if gasPriceFloat >= b.config.Node.Sequencer.L1PostingStrategy.HighGasThreshold {
					logger.Info().Float64("gasPrice", gasPriceFloat).Float64("highGasPriceConfig", b.config.Node.Sequencer.L1PostingStrategy.HighGasThreshold).Msg("not posting batch yet as gas price is high")
					creatingBatch = false
				}
			}
		}

		// Maybe sequence delayed messages
		sequencedDelayed := false
		var dontPublishBlockNum *big.Int
		if shouldSequence {
			targetSequenceDelayed := new(big.Int).Add(b.lastSequencedDelayedAt, b.sequenceDelayedMessagesInterval)
			if blockNum.Cmp(targetSequenceDelayed) >= 0 || creatingBatch {
				sequencedDelayed, err = b.deliverDelayedMessages(ctx, chainTime, false)
				if err != nil {
					logger.Error().Err(err).Msg("Error delivering delayed messages")
					continue
				}
				b.lastSequencedDelayedAt = blockNum
			}
			targetUpdateTime := new(big.Int).Add(b.latestChainTime.BlockNum.AsInt(), b.updateTimestampInterval)
			if blockNum.Cmp(targetUpdateTime) >= 0 || creatingBatch || sequencedDelayed {
				b.inboxReader.MessageDeliveryMutex.Lock()
				b.latestChainTime = chainTime
				// Avoid inefficency of publishing something that just got put in this timestamp
				dontPublishBlockNum = b.latestChainTime.BlockNum.AsInt()
				b.inboxReader.MessageDeliveryMutex.Unlock()
			}
		}

		// Maybe create a batch
		if creatingBatch {
			prevMsgCount, err := b.sequencerInbox.MessageCount(&bind.CallOpts{
				Context:     ctx,
				BlockNumber: blockNum,
			})
			if err != nil {
				logger.Error().Err(err).Msg("error getting on-chain message count")
				continue
			}
			// Gets the nonce at the latest block's state, *not* the pending state.
			// We attempt to get this at the same block as prevMsgCount,
			// but it isn't perfectly atomic as there could've been a reorg.
			// That's fine though, as the worst case is that batch creation simply fails and we retry.
			nonceInt, err := b.client.NonceAt(ctx, b.fromAddress.ToEthAddress(), blockNum)
			if err != nil {
				logger.Error().Err(err).Msg("error getting latest sequencer nonce")
				continue
			}
			nonce := new(big.Int).SetUint64(nonceInt)
			// Updates both prevMsgCount and nonce on success
			complete, err := b.publishBatch(ctx, dontPublishBlockNum, prevMsgCount, nonce)
			if err != nil {
				logger.Error().Err(err).Msg("error creating batch")
			} else if complete {
				b.lastCreatedBatchAt = blockNum
				firstBatchCreation = false
			}
		}
	}
}
