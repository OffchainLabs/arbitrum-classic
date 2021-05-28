/*
 * Copyright 2021, Offchain Labs, Inc.
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

package ethbridge

import (
	"context"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/monitor"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var sequencerBatchDeliveredID ethcommon.Hash
var sequencerBatchDeliveredFromOriginID ethcommon.Hash
var delayedInboxForcedID ethcommon.Hash
var addSequencerL2BatchFromOriginABI abi.Method

func init() {
	parsedBridgeABI, err := abi.JSON(strings.NewReader(ethbridgecontracts.SequencerInboxABI))
	if err != nil {
		panic(err)
	}
	sequencerBatchDeliveredID = parsedBridgeABI.Events["SequencerBatchDelivered"].ID
	sequencerBatchDeliveredFromOriginID = parsedBridgeABI.Events["SequencerBatchDeliveredFromOrigin"].ID
	delayedInboxForcedID = parsedBridgeABI.Events["DelayedInboxForced"].ID
	addSequencerL2BatchFromOriginABI = parsedBridgeABI.Methods["addSequencerL2BatchFromOrigin"]
}

type SequencerInboxWatcher struct {
	con     *ethbridgecontracts.SequencerInbox
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewSequencerInboxWatcher(address ethcommon.Address, client ethutils.EthClient) (*SequencerInboxWatcher, error) {
	con, err := ethbridgecontracts.NewSequencerInbox(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &SequencerInboxWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (r *SequencerInboxWatcher) Address() ethcommon.Address {
	return r.address
}

func (r *SequencerInboxWatcher) CurrentBlockHeight(ctx context.Context) (*big.Int, error) {
	latestHeader, err := r.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return latestHeader.Number, nil
}

type SequencerBatchRef interface {
	GetBatchIndex() *big.Int
	GetBeforeCount() *big.Int
	GetBeforeAcc() common.Hash
	GetAfterCount() *big.Int
	GetAfterAcc() common.Hash
}

func (r *SequencerInboxWatcher) LookupBatchesInRange(ctx context.Context, from, to *big.Int) ([]SequencerBatchRef, error) {
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: from,
		ToBlock:   to,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{sequencerBatchDeliveredID, sequencerBatchDeliveredFromOriginID, delayedInboxForcedID}},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, evmLog := range logs {
		monitor.GlobalMonitor.ReaderGotBatch(common.NewHashFromEth(evmLog.TxHash))
	}
	return r.logsToBatchRefs(ctx, logs)
}

type SequencerBatch struct {
	transactionsData   []byte
	transactionLengths []*big.Int
	sectionsMetadata   []*big.Int
	BatchIndex         *big.Int
	BeforeCount        *big.Int
	BeforeAcc          common.Hash
	AfterCount         *big.Int
	AfterAcc           common.Hash
	DelayedAcc         common.Hash
	Sequencer          common.Address
}

func (b SequencerBatch) GetBatchIndex() *big.Int {
	return b.BatchIndex
}

func (b SequencerBatch) GetBeforeCount() *big.Int {
	return b.BeforeCount
}

func (b SequencerBatch) GetBeforeAcc() common.Hash {
	return b.BeforeAcc
}

func (b SequencerBatch) GetAfterCount() *big.Int {
	return b.AfterCount
}

func (b SequencerBatch) GetAfterAcc() common.Hash {
	return b.AfterAcc
}

type sectionMetadata struct {
	numItems                *big.Int
	chainTime               inbox.ChainTime
	newTotalDelayedMessages *big.Int
	newDelayedAcc           common.Hash
}

func (b SequencerBatch) GetItems() ([]inbox.SequencerBatchItem, error) {
	sectionsMetadata := make([]sectionMetadata, 0, len(b.sectionsMetadata)/5)
	for i := 0; i+5 <= len(b.sectionsMetadata); i += 5 {
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(b.sectionsMetadata[i+1]),
			Timestamp: b.sectionsMetadata[i+2],
		}
		var newDelayedAcc common.Hash
		copy(newDelayedAcc[:], math.U256Bytes(b.sectionsMetadata[i+4]))
		sectionsMetadata = append(sectionsMetadata, sectionMetadata{
			numItems:                b.sectionsMetadata[i],
			chainTime:               chainTime,
			newTotalDelayedMessages: b.sectionsMetadata[i+3],
			newDelayedAcc:           newDelayedAcc,
		})
	}
	if len(sectionsMetadata) == 0 {
		logger.Warn().Msg("encountered sequencer batch with no batch items")
		return []inbox.SequencerBatchItem{}, nil
	}
	unaccountedTransactions := new(big.Int).Sub(b.AfterCount, b.BeforeCount)
	// Iterate backwards through all but the first section metadata
	for i := len(sectionsMetadata) - 1; i >= 1; i-- {
		meta := sectionsMetadata[i]
		txCount := meta.numItems
		unaccountedTransactions.Sub(unaccountedTransactions, txCount)
		delayedCount := meta.newTotalDelayedMessages
		prevDelayedCount := sectionsMetadata[i-1].newTotalDelayedMessages
		unaccountedTransactions.Sub(unaccountedTransactions, delayedCount)
		unaccountedTransactions.Add(unaccountedTransactions, prevDelayedCount)
		if delayedCount.Cmp(prevDelayedCount) > 0 {
			// Account for the end-of-block message
			unaccountedTransactions.Sub(unaccountedTransactions, big.NewInt(1))
		}
	}
	firstSectionMeta := sectionsMetadata[0]
	unaccountedTransactions.Sub(unaccountedTransactions, firstSectionMeta.numItems)
	if unaccountedTransactions.Sign() > 0 {
		// Account for the end-of-block message
		unaccountedTransactions.Sub(unaccountedTransactions, big.NewInt(1))
	} else if unaccountedTransactions.Sign() < 0 {
		return nil, errors.New("found a negative amount of unaccounted transactions")
	}
	// Any remaining unaccounted transactions are delayed messages in the first batch
	runningTotalDelayedMessages := new(big.Int).Sub(firstSectionMeta.newTotalDelayedMessages, unaccountedTransactions)

	ret := make([]inbox.SequencerBatchItem, 0, len(b.transactionLengths)+2)
	lastAcc := b.BeforeAcc
	nextSeqNum := new(big.Int).Set(b.BeforeCount)
	dataOffset := 0
	lengthsOffset := 0
	for _, meta := range sectionsMetadata {
		for j := 0; int64(j) < meta.numItems.Int64(); j++ {
			// Sequencer batch items
			length := int(b.transactionLengths[lengthsOffset].Int64())
			lengthsOffset += 1
			messageKind := message.L2Type
			if length == 0 {
				messageKind = message.EndOfBlockType
			}
			seqMsg := inbox.InboxMessage{
				Kind:        messageKind,
				Sender:      b.Sequencer,
				InboxSeqNum: nextSeqNum,
				GasPrice:    big.NewInt(0),
				Data:        b.transactionsData[dataOffset:(dataOffset + length)],
				ChainTime:   meta.chainTime,
			}
			dataOffset += length
			item := inbox.NewSequencerItem(runningTotalDelayedMessages, seqMsg, lastAcc)
			lastAcc = item.Accumulator
			nextSeqNum = new(big.Int).Add(nextSeqNum, big.NewInt(1))
			ret = append(ret, item)
		}

		if meta.newTotalDelayedMessages.Cmp(runningTotalDelayedMessages) > 0 {
			// Create batch item to read delayed messages
			lastSeqNum := new(big.Int).Add(nextSeqNum, meta.newTotalDelayedMessages)
			lastSeqNum.Sub(lastSeqNum, runningTotalDelayedMessages)
			nextSeqNum = new(big.Int).Set(lastSeqNum)
			lastSeqNum.Sub(lastSeqNum, big.NewInt(1))
			item := inbox.NewDelayedItem(lastSeqNum, meta.newTotalDelayedMessages, lastAcc, runningTotalDelayedMessages, meta.newDelayedAcc)
			lastAcc = item.Accumulator
			runningTotalDelayedMessages = meta.newTotalDelayedMessages
			ret = append(ret, item)

			endBlockMessage := inbox.InboxMessage{
				Kind:        message.EndOfBlockType,
				Sender:      common.Address{},
				InboxSeqNum: nextSeqNum,
				GasPrice:    big.NewInt(0),
				Data:        []byte{},
				ChainTime:   meta.chainTime,
			}
			item2 := inbox.NewSequencerItem(meta.newTotalDelayedMessages, endBlockMessage, lastAcc)
			lastAcc = item2.Accumulator
			nextSeqNum = new(big.Int).Add(nextSeqNum, big.NewInt(1))
			ret = append(ret, item2)
		}
	}

	if nextSeqNum.Cmp(b.AfterCount) != 0 {
		return nil, errors.New("computed unexpected batch end count")
	}

	if !lastAcc.Equals(b.AfterAcc) {
		return nil, errors.New("computed unexpected batch end accumulator")
	}

	return ret, nil
}

type sequencerBatchOriginRef struct {
	blockHash   ethcommon.Hash
	txIndex     uint
	batchIndex  *big.Int
	beforeCount *big.Int
	beforeAcc   common.Hash
	afterCount  *big.Int
	afterAcc    common.Hash
	delayedAcc  common.Hash
}

func (b sequencerBatchOriginRef) GetBatchIndex() *big.Int {
	return b.batchIndex
}

func (b sequencerBatchOriginRef) GetBeforeCount() *big.Int {
	return b.beforeCount
}

func (b sequencerBatchOriginRef) GetBeforeAcc() common.Hash {
	return b.beforeAcc
}

func (b sequencerBatchOriginRef) GetAfterCount() *big.Int {
	return b.afterCount
}

func (b sequencerBatchOriginRef) GetAfterAcc() common.Hash {
	return b.afterAcc
}

func (r *SequencerInboxWatcher) logsToBatchRefs(ctx context.Context, logs []types.Log) ([]SequencerBatchRef, error) {
	if len(logs) == 0 {
		return nil, nil
	}
	refs := make([]SequencerBatchRef, 0, len(logs))
	for _, log := range logs {
		if log.Topics[0] == sequencerBatchDeliveredID {
			parsed, err := r.con.ParseSequencerBatchDelivered(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			refs = append(refs, SequencerBatch{
				transactionsData:   parsed.Transactions,
				transactionLengths: parsed.Lengths,
				sectionsMetadata:   parsed.SectionsMetadata,
				BatchIndex:         parsed.SeqBatchIndex,
				BeforeCount:        parsed.FirstMessageNum,
				BeforeAcc:          parsed.BeforeAcc,
				AfterCount:         parsed.NewMessageCount,
				AfterAcc:           parsed.AfterAcc,
				Sequencer:          common.NewAddressFromEth(parsed.Sequencer),
			})
		} else if log.Topics[0] == sequencerBatchDeliveredFromOriginID {
			parsed, err := r.con.ParseSequencerBatchDeliveredFromOrigin(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			refs = append(refs, sequencerBatchOriginRef{
				blockHash:   log.BlockHash,
				txIndex:     log.TxIndex,
				batchIndex:  parsed.SeqBatchIndex,
				beforeCount: parsed.FirstMessageNum,
				beforeAcc:   parsed.BeforeAcc,
				afterCount:  parsed.NewMessageCount,
				afterAcc:    parsed.AfterAcc,
			})
		} else if log.Topics[0] == delayedInboxForcedID {
			parsed, err := r.con.ParseDelayedInboxForced(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			header, err := r.client.HeaderByHash(ctx, log.BlockHash)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			blockNum := new(big.Int).SetUint64(log.BlockNumber)
			blockTime := new(big.Int).SetUint64(header.Time)
			delayedAccInt := new(big.Int).SetBytes(parsed.AfterAccAndDelayed[1][:])
			sectionsMetadata := []*big.Int{big.NewInt(0), blockNum, blockTime, parsed.TotalDelayedMessagesRead, delayedAccInt}
			refs = append(refs, SequencerBatch{
				sectionsMetadata: sectionsMetadata,
				BatchIndex:       parsed.SeqBatchIndex,
				BeforeCount:      parsed.FirstMessageNum,
				BeforeAcc:        parsed.BeforeAcc,
				AfterCount:       parsed.NewMessageCount,
				AfterAcc:         parsed.AfterAccAndDelayed[0],
			})
		} else {
			return nil, errors.Errorf("Unexpected log topic %v", log.Topics[0].String())
		}
	}
	return refs, nil
}

func (r *SequencerInboxWatcher) ResolveBatchRef(ctx context.Context, genericRef SequencerBatchRef) (SequencerBatch, error) {
	if batch, ok := genericRef.(SequencerBatch); ok {
		return batch, nil
	}
	ref := genericRef.(sequencerBatchOriginRef)

	tx, err := r.client.TransactionInBlock(ctx, ref.blockHash, ref.txIndex)
	if err != nil {
		return SequencerBatch{}, errors.WithStack(err)
	}

	args := make(map[string]interface{})
	err = addSequencerL2BatchFromOriginABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
	if err != nil {
		return SequencerBatch{}, err
	}

	sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return SequencerBatch{}, err
	}
	return SequencerBatch{
		transactionsData:   args["transactions"].([]byte),
		transactionLengths: args["lengths"].([]*big.Int),
		sectionsMetadata:   args["sectionsMetadata"].([]*big.Int),
		BeforeCount:        ref.beforeCount,
		BeforeAcc:          ref.beforeAcc,
		AfterCount:         ref.afterCount,
		AfterAcc:           ref.afterAcc,
		DelayedAcc:         ref.delayedAcc,
		Sequencer:          common.NewAddressFromEth(sender),
	}, nil
}

func (r *SequencerInboxWatcher) GetMaxDelayBlocks(ctx context.Context) (*big.Int, error) {
	return r.con.MaxDelayBlocks(&bind.CallOpts{Context: ctx})
}
