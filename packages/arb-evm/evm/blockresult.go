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
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type OutputStatistics struct {
	GasUsed      *big.Int
	TxCount      *big.Int
	EVMLogCount  *big.Int
	AVMLogCount  *big.Int
	AVMSendCount *big.Int
}

type GasAccountingSummary struct {
	PricePerL2Tx             *big.Int
	PricePerL1CalldataByte   *big.Int
	PricePerStorageCell      *big.Int
	PricePerArbGasBase       *big.Int
	PricePerArbGasCongestion *big.Int
	PricePerArbGasTotal      *big.Int
	GasPool                  *big.Int
}

type BlockInfo struct {
	BlockNum       *big.Int
	Timestamp      *big.Int
	BlockStats     *OutputStatistics
	ChainStats     *OutputStatistics
	GasSummary     *GasAccountingSummary
	PreviousHeight *big.Int
	L1BlockNum     *big.Int
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

func (b *BlockInfo) GasLimit() *big.Int {
	limit := big.NewInt(0)
	if b.GasSummary.GasPool.Cmp(limit) > 0 {
		limit = limit.Set(b.GasSummary.GasPool)
	}
	if b.BlockStats.GasUsed.Cmp(limit) > 0 {
		limit = limit.Set(b.BlockStats.GasUsed)
	}
	return limit
}

func parseBlockResult(
	blockNum value.Value,
	timestamp value.Value,
	blockStatsRaw value.Value,
	chainStatsRaw value.Value,
	gasStatsRaw value.Value,
	previousHeight value.Value,
	l1BlockNum value.Value,
) (*BlockInfo, error) {
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return nil, errors.New("blockNum must be an int")
	}
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return nil, errors.New("timestamp must be an int")
	}
	blockStats, err := parseOutputStatistics(blockStatsRaw)
	if err != nil {
		return nil, err
	}

	chainStats, err := parseOutputStatistics(chainStatsRaw)
	if err != nil {
		return nil, err
	}
	gasStats, err := parseGasAccountingSummary(gasStatsRaw)
	if err != nil {
		return nil, err
	}
	previousHeightInt, ok := previousHeight.(value.IntValue)
	if !ok {
		return nil, errors.New("previousHeight must be an int")
	}
	l1BlockNumInt, ok := l1BlockNum.(value.IntValue)
	if !ok {
		return nil, errors.New("l1BlockNum must be an int")
	}

	return &BlockInfo{
		BlockNum:       blockNumInt.BigInt(),
		Timestamp:      timestampInt.BigInt(),
		BlockStats:     blockStats,
		ChainStats:     chainStats,
		GasSummary:     gasStats,
		PreviousHeight: previousHeightInt.BigInt(),
		L1BlockNum:     l1BlockNumInt.BigInt(),
	}, nil
}

func parseOutputStatistics(val value.Value) (*OutputStatistics, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 5 {
		return nil, errors.New("expected output statistics to be tuple of length 5")
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

func parseGasAccountingSummary(val value.Value) (*GasAccountingSummary, error) {
	tup, ok := val.(*value.TupleValue)

	offset := tup.Len() - 6;    // we skip the first field for old ArbOS versions
	
	if !ok || tup.Len() < 6 || tup.Len() > 7 {
		return nil, errors.New("could not parse gas accounting summary")
	}

	// Tuple size already verified above, so error can be ignored
	pricePerL1CalldataByte, _   := tup.GetByInt64(0 + offset)
	pricePerStorageCell, _      := tup.GetByInt64(1 + offset)
	pricePerArbGasBase, _       := tup.GetByInt64(2 + offset)
	pricePerArbGasCongestion, _ := tup.GetByInt64(3 + offset)
	pricePerArbGasTotal, _      := tup.GetByInt64(4 + offset)
	gasPool, _                  := tup.GetByInt64(5 + offset)

	pricePerL1CalldataByteInt, ok := pricePerL1CalldataByte.(value.IntValue)
	if !ok {
		return nil, errors.New("pricePerL1CalldataByte must be an int")
	}
	pricePerStorageCellInt, ok := pricePerStorageCell.(value.IntValue)
	if !ok {
		return nil, errors.New("pricePerStorageCell must be an int")
	}
	pricePerArbGasBaseInt, ok := pricePerArbGasBase.(value.IntValue)
	if !ok {
		return nil, errors.New("pricePerArbGasBase must be an int")
	}
	pricePerArbGasCongestionInt, ok := pricePerArbGasCongestion.(value.IntValue)
	if !ok {
		return nil, errors.New("pricePerArbGasCongestion must be an int")
	}
	pricePerArbGasTotalInt, ok := pricePerArbGasTotal.(value.IntValue)
	if !ok {
		return nil, errors.New("pricePerArbGasTotal must be an int")
	}
	gasPoolInt, ok := gasPool.(value.IntValue)
	if !ok {
		return nil, errors.New("gasPool must be an int")
	}
	return &GasAccountingSummary{
		PricePerL1CalldataByte:   pricePerL1CalldataByteInt.BigInt(),
		PricePerStorageCell:      pricePerStorageCellInt.BigInt(),
		PricePerArbGasBase:       pricePerArbGasBaseInt.BigInt(),
		PricePerArbGasCongestion: pricePerArbGasCongestionInt.BigInt(),
		PricePerArbGasTotal:      pricePerArbGasTotalInt.BigInt(),
		GasPool:                  math.S256(gasPoolInt.BigInt()),
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
