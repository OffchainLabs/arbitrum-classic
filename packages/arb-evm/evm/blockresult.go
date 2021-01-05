/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package evm

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type OutputStatistics struct {
	GasUsed      *big.Int
	TxCount      *big.Int
	EVMLogCount  *big.Int
	AVMLogCount  *big.Int
	AVMSendCount *big.Int
}

func (os *OutputStatistics) AsValue() value.Value {
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(os.GasUsed),
		value.NewIntValue(os.TxCount),
		value.NewIntValue(os.EVMLogCount),
		value.NewIntValue(os.AVMLogCount),
		value.NewIntValue(os.AVMSendCount),
	})
	return tup
}

type BlockInfo struct {
	BlockNum       *big.Int
	Timestamp      *big.Int
	GasLimit       *big.Int
	BlockStats     *OutputStatistics
	ChainStats     *OutputStatistics
	PreviousHeight *big.Int
}

func (b *BlockInfo) LastAVMLog() *big.Int {
	return new(big.Int).Sub(b.ChainStats.AVMLogCount, big.NewInt(1))
}

func (b *BlockInfo) FirstAVMLog() *big.Int {
	val := new(big.Int).Sub(b.ChainStats.AVMLogCount, b.BlockStats.AVMLogCount)
	// Move back one further to account for the block log itself
	return val.Sub(val, big.NewInt(1))
}

func (b *BlockInfo) LastAVMSend() *big.Int {
	return new(big.Int).Sub(b.ChainStats.AVMSendCount, big.NewInt(1))
}

func (b *BlockInfo) FirstAVMSend() *big.Int {
	return new(big.Int).Sub(b.ChainStats.AVMSendCount, b.BlockStats.AVMSendCount)
}

func (b *BlockInfo) AsValue() value.Value {
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewIntValue(b.BlockNum),
		value.NewIntValue(b.Timestamp),
		value.NewIntValue(b.GasLimit),
		b.BlockStats.AsValue(),
		b.ChainStats.AsValue(),
	})
	return tup
}

func parseBlockResult(
	blockNum value.Value,
	timestamp value.Value,
	gasLimit value.Value,
	blockStatsRaw value.Value,
	chainStatsRaw value.Value,
	previousHeight value.Value,
) (*BlockInfo, error) {
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return nil, errors.New("blockNum must be an int")
	}
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return nil, errors.New("timestamp must be an int")
	}
	gasLimitInt, ok := gasLimit.(value.IntValue)
	if !ok {
		return nil, errors.New("gasLimit must be an int")
	}
	blockStats, err := parseOutputStatistics(blockStatsRaw)
	if err != nil {
		return nil, err
	}

	chainStats, err := parseOutputStatistics(chainStatsRaw)
	if err != nil {
		return nil, err
	}
	previousHeightInt, ok := previousHeight.(value.IntValue)
	if !ok {
		return nil, errors.New("previousHeight must be an int")
	}

	return &BlockInfo{
		BlockNum:       blockNumInt.BigInt(),
		Timestamp:      timestampInt.BigInt(),
		GasLimit:       gasLimitInt.BigInt(),
		BlockStats:     blockStats,
		ChainStats:     chainStats,
		PreviousHeight: previousHeightInt.BigInt(),
	}, nil
}

func parseOutputStatistics(val value.Value) (*OutputStatistics, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 5 {
		return nil, errors.New("expected result to be tuple of length 5")
	}

	// Tuple size already verified above, so error can be ignored
	gasUsed, _ := tup.GetByInt64(0)
	txCount, _ := tup.GetByInt64(1)
	evmLogCount, _ := tup.GetByInt64(2)
	avmLogCount, _ := tup.GetByInt64(3)
	avmSendCount, _ := tup.GetByInt64(4)

	gasUsedInt, ok := gasUsed.(value.IntValue)
	if !ok {
		return nil, errors.New("gasUsed must be an int")
	}
	txCountInt, ok := txCount.(value.IntValue)
	if !ok {
		return nil, errors.New("txCount must be an int")
	}
	evmLogCountInt, ok := evmLogCount.(value.IntValue)
	if !ok {
		return nil, errors.New("evmLogCount must be an int")
	}
	avmLogCountInt, ok := avmLogCount.(value.IntValue)
	if !ok {
		return nil, errors.New("avmLogCount must be an int")
	}
	avmSendCountInt, ok := avmSendCount.(value.IntValue)
	if !ok {
		return nil, errors.New("avmSendCount must be an int")
	}
	return &OutputStatistics{
		GasUsed:      gasUsedInt.BigInt(),
		TxCount:      txCountInt.BigInt(),
		EVMLogCount:  evmLogCountInt.BigInt(),
		AVMLogCount:  avmLogCountInt.BigInt(),
		AVMSendCount: avmSendCountInt.BigInt(),
	}, nil
}

func NewBlockResultFromValue(val value.Value) (*BlockInfo, error) {
	res, err := NewResultFromValue(val)
	if err != nil {
		return nil, err
	}
	txRes, ok := res.(*BlockInfo)
	if !ok {
		return nil, errors.New("got transaction result but expected block")
	}
	return txRes, nil
}
