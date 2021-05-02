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

package inbox

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type SequencerBatchItem struct {
	LastSeqNum        *big.Int
	Accumulator       common.Hash
	TotalDelayedCount *big.Int
	SequencerMessage  []byte
}

func NewSequencerItem(totalDelayedCount *big.Int, msg InboxMessage, prevAcc common.Hash) SequencerBatchItem {
	var data []byte
	data = append(data, "Sequencer message:"...)
	data = append(data, prevAcc.Bytes()...)
	data = append(data, math.U256Bytes(msg.InboxSeqNum)...)
	data = append(data, msg.CommitmentHash().Bytes()...)
	return SequencerBatchItem{
		LastSeqNum:        msg.InboxSeqNum,
		Accumulator:       hashing.SoliditySHA3(data),
		TotalDelayedCount: totalDelayedCount,
		SequencerMessage:  msg.ToBytes(),
	}
}

func NewDelayedItem(lastSeqNum *big.Int, totalDelayedCount *big.Int, prevAcc common.Hash, prevDelayedCount *big.Int, delayedAcc common.Hash) SequencerBatchItem {
	var data []byte
	data = append(data, "Delayed messages:"...)
	data = append(data, prevAcc.Bytes()...)
	firstSeqNum := big.NewInt(1)
	firstSeqNum = firstSeqNum.Add(firstSeqNum, lastSeqNum)
	firstSeqNum = firstSeqNum.Add(firstSeqNum, prevDelayedCount)
	firstSeqNum = firstSeqNum.Sub(firstSeqNum, totalDelayedCount)
	data = append(data, math.U256Bytes(firstSeqNum)...)
	data = append(data, math.U256Bytes(prevDelayedCount)...)
	data = append(data, math.U256Bytes(totalDelayedCount)...)
	data = append(data, delayedAcc.Bytes()...)
	return SequencerBatchItem{
		LastSeqNum:        lastSeqNum,
		Accumulator:       hashing.SoliditySHA3(data),
		TotalDelayedCount: totalDelayedCount,
	}
}

func NewSequencerBatchItemFromData(data []byte) (SequencerBatchItem, error) {
	if len(data) < 32*3 {
		return SequencerBatchItem{}, errors.New("Not enough data for sequencer batch item")
	}
	item := SequencerBatchItem{}

	item.LastSeqNum = new(big.Int).SetBytes(data[:32])
	data = data[32:]

	copy(item.Accumulator[:], data[:32])
	data = data[32:]

	item.TotalDelayedCount = new(big.Int).SetBytes(data[:32])
	data = data[32:]

	item.SequencerMessage = data

	return item, nil
}

func (i SequencerBatchItem) ToBytesWithSeqNum() []byte {
	var data []byte
	data = append(data, math.U256Bytes(i.LastSeqNum)...)
	data = append(data, i.Accumulator.Bytes()...)
	data = append(data, math.U256Bytes(i.TotalDelayedCount)...)
	data = append(data, i.SequencerMessage...)
	return data
}
