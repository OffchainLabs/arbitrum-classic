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
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
)

type SequencerBatcher struct {
	db                         core.ArbCore
	inboxReader                *staker.InboxReader
	client                     ethutils.EthClient
	delayedMessagesTargetDelay *big.Int
	sequencerInbox             *ethbridgecontracts.SequencerInbox
	auth                       *ethbridge.TransactAuth

	sequencer       common.Address
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

func NewSequencerBatcher(ctx context.Context, db core.ArbCore, inboxReader *staker.InboxReader, client ethutils.EthClient, delayedMessagesTargetDelay *big.Int, sequencerInbox *ethbridgecontracts.SequencerInbox, auth *bind.TransactOpts) (*SequencerBatcher, error) {
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

	batcher := &SequencerBatcher{
		db:                         db,
		inboxReader:                inboxReader,
		client:                     client,
		delayedMessagesTargetDelay: delayedMessagesTargetDelay,
		sequencerInbox:             sequencerInbox,
		auth:                       transactAuth,

		sequencer:       common.NewAddressFromEth(sequencer),
		txQueue:         make(chan *types.Transaction, 10),
		newTxFeed:       event.Feed{},
		latestChainTime: chainTime,
	}
	// TODO: publish any unpublished messages on startup
	go batcher.chainManager(ctx)
	return batcher, nil
}

func (b *SequencerBatcher) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	var x uint64
	return &x
}

func (b *SequencerBatcher) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}

func (b *SequencerBatcher) SendTransaction(ctx context.Context, startTx *types.Transaction) error {
	b.txQueue <- startTx
	b.inboxReader.MessageDeliveryMutex.Lock()
	defer b.inboxReader.MessageDeliveryMutex.Unlock()

	var batchTxs []*types.Transaction
	var l2BatchContents []message.AbstractL2Message
	// This pattern is safe as we acquired a lock so we are the exclusive reader
	for len(b.txQueue) > 0 {
		tx := <-b.txQueue
		batchTxs = append(batchTxs, tx)
		l2BatchContents = append(l2BatchContents, message.NewTransactionFromEthTx(tx))
	}

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
	if totalDelayedCount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("chain not yet initialized")
	}

	batch, err := message.NewTransactionBatchFromMessages(l2BatchContents)
	if err != nil {
		return err
	}
	l2Message := message.NewSafeL2Message(batch)
	seqMsg := message.NewInboxMessage(l2Message, b.sequencer, new(big.Int).Set(msgCount), big.NewInt(0), b.latestChainTime)

	newBlockSeqNum := new(big.Int).Add(msgCount, big.NewInt(1))
	newBlockMessage := inbox.InboxMessage{
		Kind:        message.EndOfBlockType,
		Sender:      common.Address{},
		InboxSeqNum: newBlockSeqNum,
		GasPrice:    big.NewInt(0),
		Data:        []byte{},
		ChainTime: inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	}

	txBatchItem := inbox.SequencerBatchItem{
		LastSeqNum:        msgCount,
		Accumulator:       common.Hash{},
		TotalDelayedCount: totalDelayedCount,
		SequencerMessage:  seqMsg.ToBytes(),
	}
	txBatchItem.RecomputeAccumulator(prevAcc, totalDelayedCount, common.Hash{})
	newBlockBatchItem := inbox.SequencerBatchItem{
		LastSeqNum:        newBlockSeqNum,
		Accumulator:       common.Hash{},
		TotalDelayedCount: totalDelayedCount,
		SequencerMessage:  newBlockMessage.ToBytes(),
	}
	newBlockBatchItem.RecomputeAccumulator(txBatchItem.Accumulator, totalDelayedCount, common.Hash{})

	seqBatchItems := []inbox.SequencerBatchItem{txBatchItem, newBlockBatchItem}
	success, err := core.DeliverMessagesAndWait(b.db, prevAcc, seqBatchItems, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("failed to deliver messages")
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

func (b *SequencerBatcher) deliverDelayedMessages(ctx context.Context, chainTime inbox.ChainTime) (*big.Int, error) {
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
	if newDelayedCount.Cmp(big.NewInt(0)) == 0 {
		// Immediately sequence the init message, even if it technically isn't confirmed yet
		newDelayedCount = big.NewInt(1)
	}
	if newDelayedCount.Cmp(oldDelayedCount) <= 0 {
		return msgCount, nil
	}

	delayedRead := new(big.Int).Sub(newDelayedCount, oldDelayedCount)
	newMsgCount := new(big.Int).Add(msgCount, delayedRead)
	newMsgCount.Add(newMsgCount, big.NewInt(1)) // end of block message
	lastSeqNum := new(big.Int).Sub(newMsgCount, big.NewInt(2))
	batchItem := inbox.SequencerBatchItem{
		LastSeqNum:        lastSeqNum,
		Accumulator:       common.Hash{},
		TotalDelayedCount: newDelayedCount,
		SequencerMessage:  []byte{},
	}
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
	err = batchItem.RecomputeAccumulator(prevAcc, oldDelayedCount, delayedAcc)
	if err != nil {
		return nil, err
	}

	endOfBlockSeqNum := new(big.Int).Add(lastSeqNum, big.NewInt(1))
	endOfBlockMessage := inbox.InboxMessage{
		Kind:        message.EndOfBlockType,
		Sender:      common.Address{},
		InboxSeqNum: endOfBlockSeqNum,
		GasPrice:    big.NewInt(0),
		Data:        []byte{},
		ChainTime: inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	}
	endBlockBatchItem := inbox.SequencerBatchItem{
		LastSeqNum:        endOfBlockSeqNum,
		Accumulator:       common.Hash{},
		TotalDelayedCount: newDelayedCount,
		SequencerMessage:  endOfBlockMessage.ToBytes(),
	}
	err = endBlockBatchItem.RecomputeAccumulator(batchItem.Accumulator, newDelayedCount, delayedAcc)
	if err != nil {
		return nil, err
	}

	success, err := core.DeliverMessagesAndWait(b.db, prevAcc, []inbox.SequencerBatchItem{batchItem, endBlockBatchItem}, []inbox.DelayedMessage{}, nil)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errors.New("Failed to deliver messages")
	}

	b.latestChainTime = chainTime
	return newMsgCount, nil
}

func (b *SequencerBatcher) createBatch(ctx context.Context, newMsgCount *big.Int) error {
	prevMsgCount, err := b.sequencerInbox.MessageCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	if newMsgCount.Cmp(prevMsgCount) == 0 {
		return nil
	}
	batchItems, err := b.db.GetSequencerBatchItems(prevMsgCount, new(big.Int).Sub(newMsgCount, prevMsgCount))
	if err != nil || len(batchItems) == 0 {
		return err
	}

	var transactionsData []byte
	var transactionsLengths []*big.Int
	var startDelayedMessagesRead *big.Int
	var totalDelayedMessagesRead *big.Int
	var l1BlockNumber *big.Int
	var l1Timestamp *big.Int
	var lastAcc common.Hash
	skippingImplicitEndOfBlock := false
	for _, item := range batchItems {
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
				return errors.New("back-to-back delayed messages inserted without end of block")
			}
			skippingImplicitEndOfBlock = true
		} else {
			seqMsg, err := inbox.NewInboxMessageFromData(item.SequencerMessage)
			if err != nil {
				return err
			}

			if seqMsg.Kind == message.L2Type {
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
			} else if seqMsg.Kind == message.EndOfBlockType {
				if len(seqMsg.Data) != 0 {
					return errors.New("end of block message has data")
				}
			} else {
				return errors.Errorf("unexpected sequencer message kind %v", seqMsg.Kind)
			}

			if skippingImplicitEndOfBlock {
				if seqMsg.Kind != message.EndOfBlockType {
					return errors.New("found non-end-of-block sequencer message after delayed messages")
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
		lastAcc = item.Accumulator
	}

	if skippingImplicitEndOfBlock {
		return errors.New("didn't find implicit end of block after delayed messages")
	}
	if totalDelayedMessagesRead == nil {
		totalDelayedMessagesRead = startDelayedMessagesRead
	}
	if l1BlockNumber == nil {
		// This batch consists exclusively of delayed messages,
		// so the l1 block number is largely irrelevant
		chainTime, err := getChainTime(ctx, b.client)
		if err != nil {
			return err
		}
		l1BlockNumber = chainTime.BlockNum.AsInt()
		l1Timestamp = chainTime.Timestamp
	}

	logger.Info().Str("prevMsgCount", prevMsgCount.String()).Str("newMsgCount", newMsgCount.String()).Msg("Creating sequencer batch")
	_, err = ethbridge.AddSequencerL2BatchFromOrigin(ctx, b.sequencerInbox, b.auth, transactionsData, transactionsLengths, l1BlockNumber, l1Timestamp, totalDelayedMessagesRead, lastAcc)

	return err
}

func (b *SequencerBatcher) chainManager(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		chainTime, err := getChainTime(ctx, b.client)
		if err != nil {
			logger.Warn().Err(err).Msg("Error getting chain time")
			continue
		}
		if chainTime.BlockNum.Cmp(b.latestChainTime.BlockNum) <= 0 {
			continue
		}
		newMsgCount, err := b.deliverDelayedMessages(ctx, chainTime)
		if err != nil {
			logger.Error().Err(err).Msg("Error delivering delayed messages")
			continue
		}
		for {
			err = b.createBatch(ctx, newMsgCount)
			if err == nil {
				time.Sleep(10 * time.Second)
				break
			} else {
				logger.Error().Err(err).Msg("Error creating batch")
			}
			time.Sleep(time.Second)
		}
	}
}
