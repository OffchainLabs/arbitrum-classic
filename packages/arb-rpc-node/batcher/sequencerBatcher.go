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
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"

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

type SequencerBatcher struct {
	db                         core.ArbCore
	inboxReader                *monitor.InboxReader
	client                     ethutils.EthClient
	delayedMessagesTargetDelay *big.Int
	sequencerInbox             *ethbridgecontracts.SequencerInbox
	auth                       *ethbridge.TransactAuth
	chainTimeCheckInterval     time.Duration
	logBatchGasCosts           bool
	feedBroadcaster            *broadcaster.Broadcaster
	dataSigner                 func([]byte) ([]byte, error)

	sequencer       common.Address
	signer          types.Signer
	txQueue         chan *types.Transaction
	newTxFeed       event.Feed
	latestChainTime inbox.ChainTime
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
	sequencerInbox *ethbridgecontracts.SequencerInbox,
	auth *bind.TransactOpts,
	dataSigner func([]byte) ([]byte, error),
	broadcasterSettings broadcaster.Settings,
) (*SequencerBatcher, error) {
	chainTime, err := getChainTime(ctx, client)
	if err != nil {
		return nil, err
	}

	sequencer, err := sequencerInbox.Sequencer(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	if sequencer != auth.From {
		return nil, errors.New("Transaction auth isn't for sequencer")
	}

	transactAuth, err := ethbridge.NewTransactAuth(ctx, client, auth)
	if err != nil {
		return nil, err
	}

	feedBroadcaster := broadcaster.NewBroadcaster(broadcasterSettings)
	err = feedBroadcaster.Start(ctx)
	if err != nil {
		logger.Warn().Err(err).Msg("error starting feed broadcaster")
		return nil, err
	}

	return &SequencerBatcher{
		db:                         db,
		inboxReader:                inboxReader,
		client:                     client,
		delayedMessagesTargetDelay: delayedMessagesTargetDelay,
		sequencerInbox:             sequencerInbox,
		auth:                       transactAuth,
		chainTimeCheckInterval:     time.Second,
		feedBroadcaster:            feedBroadcaster,
		dataSigner:                 dataSigner,

		sequencer:       common.NewAddressFromEth(sequencer),
		signer:          types.NewEIP155Signer(chainId),
		txQueue:         make(chan *types.Transaction, 10),
		newTxFeed:       event.Feed{},
		latestChainTime: chainTime,
	}, nil
}

func (b *SequencerBatcher) PendingTransactionCount(_ context.Context, _ common.Address) *uint64 {
	return nil
}

func (b *SequencerBatcher) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}

func (b *SequencerBatcher) SendTransaction(_ context.Context, startTx *types.Transaction) error {
	_, err := types.Sender(b.signer, startTx)
	if err != nil {
		logger.Warn().Err(err).Msg("error processing user transaction")
		return err
	}
	logger.Info().Str("hash", startTx.Hash().String()).Msg("got user tx")
	b.txQueue <- startTx
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	var batchTxs []*types.Transaction
	var l2BatchContents []message.AbstractL2Message
	// This pattern is safe as we acquired a lock so we are the exclusive reader
	for len(b.txQueue) > 0 {
		tx := <-b.txQueue
		batchTxs = append(batchTxs, tx)
		l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(tx))
	}
	logger.Info().Int("count", len(l2BatchContents)).Msg("gather user txes")

	if len(l2BatchContents) == 0 {
		// The queue is empty, so another goroutine processed startTx already
		return nil
	}

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

	newBlockSeqNum := new(big.Int).Add(msgCount, big.NewInt(1))
	newBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		b.sequencer,
		newBlockSeqNum,
		big.NewInt(0),
		b.latestChainTime.Clone(),
	)

	txBatchItem := inbox.NewSequencerItem(totalDelayedCount, seqMsg, prevAcc)
	newBlockBatchItem := inbox.NewSequencerItem(totalDelayedCount, newBlockMessage, txBatchItem.Accumulator)
	seqBatchItems := []inbox.SequencerBatchItem{txBatchItem, newBlockBatchItem}
	success, err := core.DeliverMessagesAndWait(b.db, prevAcc, seqBatchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("failed to deliver messages")
	}

	signature, err := b.dataSigner(hashing.SoliditySHA3WithPrefix(hashing.Bytes32(txBatchItem.Accumulator)).Bytes())
	if err != nil {
		return err
	}
	err = b.feedBroadcaster.Broadcast(prevAcc, seqBatchItems[0], signature)
	if err != nil {
		return err
	}

	// TODO check if transaction was valid and if not roll it back
	b.newTxFeed.Send(ethcore.NewTxsEvent{Txs: batchTxs})
	return nil
}

func (b *SequencerBatcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	// TODO: return latest machine state?
	return nil, nil
}

func (b *SequencerBatcher) Aggregator() *common.Address {
	return &b.sequencer
}

func (b *SequencerBatcher) deliverDelayedMessages(chainTime inbox.ChainTime) (*big.Int, error) {
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()
	msgCount, err := b.db.GetMessageCount()
	if err != nil {
		return nil, err
	}
	oldDelayedCount, err := b.db.GetTotalDelayedMessagesSequenced()
	if err != nil {
		return nil, err
	}
	newDelayedCount, err := b.db.GetDelayedMessagesToSequence(new(big.Int).Sub(chainTime.BlockNum.AsInt(), b.delayedMessagesTargetDelay))
	if err != nil {
		return nil, err
	}
	if newDelayedCount.Cmp(oldDelayedCount) <= 0 {
		b.latestChainTime = chainTime
		return msgCount, nil
	}

	delayedRead := new(big.Int).Sub(newDelayedCount, oldDelayedCount)
	newMsgCount := new(big.Int).Add(msgCount, delayedRead)
	newMsgCount.Add(newMsgCount, big.NewInt(1)) // end of block message
	lastSeqNum := new(big.Int).Sub(newMsgCount, big.NewInt(2))

	var prevAcc common.Hash
	if msgCount.Cmp(big.NewInt(1)) > 0 {
		prevAcc, err = b.db.GetInboxAcc(new(big.Int).Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return nil, err
		}
	}
	delayedAcc, err := b.db.GetDelayedInboxAcc(new(big.Int).Sub(newDelayedCount, big.NewInt(1)))
	if err != nil {
		return nil, err
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
	success, err := core.DeliverMessagesAndWait(b.db, prevAcc, seqBatchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errors.New("Failed to deliver messages")
	}

	signature, err := b.dataSigner(hashing.SoliditySHA3WithPrefix(hashing.Bytes32(batchItem.Accumulator)).Bytes())
	if err != nil {
		return nil, err
	}
	err = b.feedBroadcaster.Broadcast(prevAcc, seqBatchItems[0], signature)
	if err != nil {
		return nil, err
	}

	b.latestChainTime = chainTime
	return newMsgCount, nil
}

const gasCostBase int = 70292
const gasCostDelayedMessages int = 63505
const gasCostPerMessage int = 1431
const gasCostPerMessageByte int = 16
const gasCostMaximum int = 2_000_000

func (b *SequencerBatcher) createBatch(ctx context.Context, newMsgCount *big.Int) (bool, error) {
	prevMsgCount, err := b.sequencerInbox.MessageCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, err
	}
	if newMsgCount.Cmp(prevMsgCount) == 0 {
		return true, nil
	}
	batchItems, err := b.db.GetSequencerBatchItems(prevMsgCount, new(big.Int).Sub(newMsgCount, prevMsgCount))
	if err != nil || len(batchItems) == 0 {
		return false, err
	}

	var transactionsData []byte
	var transactionsLengths []*big.Int
	var startDelayedMessagesRead *big.Int
	var totalDelayedMessagesRead *big.Int
	var l1BlockNumber *big.Int
	var l1Timestamp *big.Int
	var lastAcc common.Hash
	estimatedGasCost := gasCostBase
	skippingImplicitEndOfBlock := false
	publishingAllBatchItems := false
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
		if i != 0 && estimatedGasCost >= gasCostMaximum {
			break
		}

		if startDelayedMessagesRead == nil {
			startDelayedMessagesRead = item.TotalDelayedCount
		} else if totalDelayedMessagesRead == nil {
			if item.TotalDelayedCount.Cmp(startDelayedMessagesRead) > 0 {
				totalDelayedMessagesRead = item.TotalDelayedCount
			}
		} else if totalDelayedMessagesRead.Cmp(item.TotalDelayedCount) != 0 {
			break
		}

		if len(item.SequencerMessage) == 0 {
			if skippingImplicitEndOfBlock {
				return false, errors.New("back-to-back delayed messages inserted without end of block")
			}
			skippingImplicitEndOfBlock = true
		} else {
			if l1BlockNumber == nil {
				l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
			} else if l1BlockNumber.Cmp(seqMsg.ChainTime.BlockNum.AsInt()) != 0 {
				break
			}
			if l1Timestamp == nil {
				l1Timestamp = seqMsg.ChainTime.Timestamp
			} else if l1Timestamp.Cmp(seqMsg.ChainTime.Timestamp) != 0 {
				break
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
			} else if totalDelayedMessagesRead != nil {
				// We're attempting to insert sequencer messages after delayed messages which isn't allowed
				break
			} else {
				transactionsData = append(transactionsData, seqMsg.Data...)
				transactionsLengths = append(transactionsLengths, big.NewInt(int64(len(seqMsg.Data))))
			}
		}
		if i == len(batchItems)-1 {
			publishingAllBatchItems = true
		}
		lastAcc = item.Accumulator
	}

	if skippingImplicitEndOfBlock {
		return false, errors.New("didn't find implicit end of block after delayed messages")
	}
	if totalDelayedMessagesRead == nil {
		totalDelayedMessagesRead = startDelayedMessagesRead
	}

	logger.Info().Str("prevMsgCount", prevMsgCount.String()).Str("newMsgCount", newMsgCount.String()).Msg("Creating sequencer batch")
	tx, err := ethbridge.AddSequencerL2BatchFromOrigin(ctx, b.sequencerInbox, b.auth, transactionsData, transactionsLengths, l1BlockNumber, l1Timestamp, totalDelayedMessagesRead, lastAcc)
	if err != nil {
		return false, err
	}

	receipt, err := ethbridge.WaitForReceiptWithResultsSimple(ctx, b.client, tx.ToEthHash())
	if err != nil {
		return false, err
	}

	if b.logBatchGasCosts {
		fmt.Printf("%v,%v,%v\n", len(transactionsLengths), len(transactionsData), receipt.GasUsed)
	}

	return publishingAllBatchItems, nil
}

func (b *SequencerBatcher) Start(ctx context.Context) {
	logger.Log().Msg("Starting sequencer batch submission thread")
	firstBoot := true
	defer b.feedBroadcaster.Stop()

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
		if chainTime.BlockNum.Cmp(b.latestChainTime.BlockNum) <= 0 && !firstBoot {
			continue
		}
		firstBoot = false
		newMsgCount, err := b.deliverDelayedMessages(chainTime)
		if err != nil {
			logger.Error().Err(err).Msg("Error delivering delayed messages")
			continue
		}
		for {
			complete, err := b.createBatch(ctx, newMsgCount)
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
		}
	}
}
