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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
)

var sequencerBridgeABI abi.ABI
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
	sequencerBridgeABI = parsedBridgeABI
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

func (r *SequencerInboxWatcher) CurrentBlockHeight(ctx context.Context) (*big.Int, error) {
	latestHeader, err := r.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return latestHeader.Number, nil
}

type SequencerBatchRef interface {
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
	transactionsData         []byte
	transactionLengths       []*big.Int
	L1BlockNumber            *big.Int
	Timestamp                *big.Int
	TotalDelayedMessagesRead *big.Int
	BeforeCount              *big.Int
	BeforeAcc                common.Hash
	AfterCount               *big.Int
	AfterAcc                 common.Hash
	DelayedAcc               common.Hash
	ChainTime                inbox.ChainTime
	Sequencer                common.Address
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

func newEndOfBlockMessage(seqNum *big.Int) inbox.InboxMessage {
	return inbox.InboxMessage{
		Kind:        message.EndOfBlockType,
		Sender:      common.Address{},
		InboxSeqNum: seqNum,
		GasPrice:    big.NewInt(0),
		Data:        []byte{},
		ChainTime: inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	}
}

func (b SequencerBatch) GetItems() []inbox.SequencerBatchItem {
	count := new(big.Int).Sub(b.AfterCount, b.BeforeCount)
	delayedCount := new(big.Int).Sub(count, big.NewInt(int64(len(b.transactionLengths))))
	hasDelayed := delayedCount.Cmp(big.NewInt(0)) > 0
	startDelayedCount := b.TotalDelayedMessagesRead
	if hasDelayed {
		// Subtract out the end of block message, which isn't really delayed
		delayedCount.Sub(delayedCount, big.NewInt(1))
		startDelayedCount = new(big.Int).Sub(b.TotalDelayedMessagesRead, delayedCount)
	}

	ret := make([]inbox.SequencerBatchItem, 0, len(b.transactionLengths)+2)
	lastAcc := b.BeforeAcc
	nextSeqNum := new(big.Int).Set(b.BeforeCount)
	dataOffset := 0
	for i := 0; i < len(b.transactionLengths); i++ {
		// Sequencer batch items
		var message inbox.InboxMessage
		length := int(b.transactionLengths[i].Int64())
		if length == 0 {
			message = newEndOfBlockMessage(nextSeqNum)
		} else {
			message = inbox.InboxMessage{
				Kind:        3,
				Sender:      b.Sequencer,
				InboxSeqNum: nextSeqNum,
				GasPrice:    big.NewInt(0),
				Data:        b.transactionsData[dataOffset:(dataOffset + length)],
				ChainTime:   b.ChainTime,
			}
		}
		dataOffset += length
		item := inbox.SequencerBatchItem{
			LastSeqNum:        nextSeqNum,
			Accumulator:       common.Hash{},
			TotalDelayedCount: startDelayedCount,
			SequencerMessage:  message.ToBytes(),
		}
		item.RecomputeAccumulator(lastAcc, startDelayedCount, common.Hash{})
		lastAcc = item.Accumulator
		nextSeqNum = new(big.Int).Add(nextSeqNum, big.NewInt(1))
		ret = append(ret, item)
	}

	if hasDelayed {
		// Create batch item to read delayed messages
		lastSeqNum := new(big.Int).Sub(b.AfterCount, big.NewInt(2))
		item := inbox.SequencerBatchItem{
			LastSeqNum:        lastSeqNum,
			Accumulator:       common.Hash{},
			TotalDelayedCount: b.TotalDelayedMessagesRead,
			SequencerMessage:  []byte{},
		}
		item.RecomputeAccumulator(lastAcc, startDelayedCount, b.DelayedAcc)
		ret = append(ret, item)

		endSeqNum := new(big.Int).Add(lastSeqNum, big.NewInt(1))
		item2 := inbox.SequencerBatchItem{
			LastSeqNum:        endSeqNum,
			Accumulator:       common.Hash{},
			TotalDelayedCount: b.TotalDelayedMessagesRead,
			SequencerMessage:  newEndOfBlockMessage(endSeqNum).ToBytes(),
		}
		item2.RecomputeAccumulator(item.Accumulator, b.TotalDelayedMessagesRead, b.DelayedAcc)
		ret = append(ret, item2)
	}

	return ret
}

type sequencerBatchOriginRef struct {
	blockHash   ethcommon.Hash
	txIndex     uint
	beforeCount *big.Int
	beforeAcc   common.Hash
	afterCount  *big.Int
	afterAcc    common.Hash
	delayedAcc  common.Hash
	chainTime   inbox.ChainTime
	sequencer   common.Address
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
	sequencerEthAddr, err := r.con.Sequencer(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	sequencer := common.NewAddressFromEth(sequencerEthAddr)
	blockTimes := make(map[ethcommon.Hash]*big.Int)
	refs := make([]SequencerBatchRef, 0, len(logs))
	for _, log := range logs {
		blockTime, ok := blockTimes[log.BlockHash]
		if !ok {
			header, err := r.client.HeaderByHash(ctx, log.BlockHash)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			blockTime = new(big.Int).SetUint64(header.Time)
			blockTimes[log.BlockHash] = blockTime
		}
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(new(big.Int).SetUint64(log.BlockNumber)),
			Timestamp: blockTime,
		}

		if log.Topics[0] == sequencerBatchDeliveredID {
			parsed, err := r.con.ParseSequencerBatchDelivered(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			refs = append(refs, SequencerBatch{
				transactionsData:         parsed.Transactions,
				transactionLengths:       parsed.Lengths,
				L1BlockNumber:            parsed.L1BlockNumber,
				Timestamp:                parsed.Timestamp,
				TotalDelayedMessagesRead: parsed.TotalDelayedMessagesRead,
				BeforeCount:              parsed.FirstMessageNum,
				BeforeAcc:                parsed.BeforeAcc,
				AfterCount:               parsed.NewMessageCount,
				AfterAcc:                 parsed.AfterAcc,
				DelayedAcc:               parsed.DelayedAcc,
				Sequencer:                sequencer,
				ChainTime:                chainTime,
			})
		} else if log.Topics[0] == sequencerBatchDeliveredFromOriginID {
			parsed, err := r.con.ParseSequencerBatchDeliveredFromOrigin(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			refs = append(refs, sequencerBatchOriginRef{
				blockHash:   log.BlockHash,
				txIndex:     log.TxIndex,
				beforeCount: parsed.FirstMessageNum,
				beforeAcc:   parsed.BeforeAcc,
				afterCount:  parsed.NewMessageCount,
				afterAcc:    parsed.AfterAcc,
				delayedAcc:  parsed.DelayedAcc,
				sequencer:   sequencer,
				chainTime:   chainTime,
			})
		} else if log.Topics[0] == delayedInboxForcedID {
			parsed, err := r.con.ParseDelayedInboxForced(log)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			refs = append(refs, SequencerBatch{
				TotalDelayedMessagesRead: parsed.TotalDelayedMessagesRead,
				BeforeCount:              parsed.FirstMessageNum,
				BeforeAcc:                parsed.BeforeAcc,
				AfterCount:               parsed.NewMessageCount,
				AfterAcc:                 parsed.AfterAccAndDelayed[0],
				DelayedAcc:               parsed.AfterAccAndDelayed[1],
				Sequencer:                sequencer,
				ChainTime:                chainTime,
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

	return SequencerBatch{
		transactionsData:         args["transactions"].([]byte),
		transactionLengths:       args["lengths"].([]*big.Int),
		L1BlockNumber:            args["l1BlockNumber"].(*big.Int),
		Timestamp:                args["timestamp"].(*big.Int),
		TotalDelayedMessagesRead: args["_totalDelayedMessagesRead"].(*big.Int),
		BeforeCount:              ref.beforeCount,
		BeforeAcc:                ref.beforeAcc,
		AfterCount:               ref.afterCount,
		AfterAcc:                 ref.afterAcc,
		DelayedAcc:               ref.delayedAcc,
		Sequencer:                ref.sequencer,
		ChainTime:                ref.chainTime,
	}, nil
}
