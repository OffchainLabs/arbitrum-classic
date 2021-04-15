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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/pkg/errors"
)

type SequencerBatchItem struct {
	LastSeqNum        *big.Int
	Accumulator       common.Hash
	TotalDelayedCount *big.Int
	SequencerMessage  []byte
}

func (i *SequencerBatchItem) RecomputeAccumulator(prevAcc common.Hash, prevDelayedCount *big.Int, delayedAcc common.Hash) error {
	var data []byte
	delayedCmp := i.TotalDelayedCount.Cmp(prevDelayedCount)
	if delayedCmp > 0 {
		if len(i.SequencerMessage) > 0 {
			return errors.New("Sequencer batch item has both sequencer message and delayed messages")
		}
		data = append(data, "Delayed messages:"...)
		data = append(data, prevAcc.Bytes()...)
		firstSeqNum := big.NewInt(1)
		firstSeqNum = firstSeqNum.Add(firstSeqNum, i.LastSeqNum)
		firstSeqNum = firstSeqNum.Add(firstSeqNum, prevDelayedCount)
		firstSeqNum = firstSeqNum.Sub(firstSeqNum, i.TotalDelayedCount)
		data = append(data, math.U256Bytes(firstSeqNum)...)
		data = append(data, math.U256Bytes(prevDelayedCount)...)
		data = append(data, math.U256Bytes(i.TotalDelayedCount)...)
		data = append(data, delayedAcc.Bytes()...)
	} else if delayedCmp == 0 {
		data = append(data, "Sequencer message:"...)
		data = append(data, prevAcc.Bytes()...)
		data = append(data, math.U256Bytes(i.LastSeqNum)...)
		msg, err := NewInboxMessageFromData(i.SequencerMessage)
		if err != nil {
			return err
		}
		data = append(data, msg.CommitmentHash().Bytes()...)
	} else {
		return errors.New("Sequencer batch item delayed count went backwards")
	}
	i.Accumulator = hashing.SoliditySHA3(data)
	return nil
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
