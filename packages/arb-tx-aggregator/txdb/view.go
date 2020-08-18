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

package txdb

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type View struct {
	as *cmachine.AggregatorStore
}

func (txdb *View) GetMessage(index uint64) (value.Value, error) {
	return txdb.as.GetMessage(index)
}

func (txdb *View) GetLog(index uint64) (value.Value, error) {
	return txdb.as.GetLog(index)
}

func (txdb *View) GetRequest(requestId common.Hash) (value.Value, error) {
	requestCandidate, err := txdb.as.GetPossibleRequestInfo(requestId)
	if err != nil {
		return nil, err
	}
	logVal, err := txdb.as.GetLog(requestCandidate)
	if err != nil {
		return nil, err
	}
	res, err := evm.NewTxResultFromValue(logVal)
	if err != nil {
		return nil, err
	}
	if res.IncomingRequest.MessageID != requestId {
		return nil, errors.New("request not found")
	}
	return logVal, nil
}

func (txdb *View) GetBlock(height uint64) (*machine.BlockInfo, error) {
	return txdb.as.GetBlock(height)
}

func (txdb *View) LatestBlockId() (*common.BlockId, error) {
	return txdb.as.LatestBlock()
}

func (txdb *View) FindLogs(
	ctx context.Context,
	fromHeight *uint64,
	toHeight *uint64,
	address []common.Address,
	topics [][]common.Hash,
) ([]evm.FullLog, error) {
	latestBlock, err := txdb.LatestBlockId()
	if err != nil {
		return nil, err
	}
	startHeight := uint64(0)
	endHeight := latestBlock.Height.AsInt().Uint64()
	if fromHeight != nil && *fromHeight > 0 {
		startHeight = *fromHeight
	}
	if toHeight != nil {
		altEndHeight := *toHeight + 1
		if endHeight > altEndHeight {
			endHeight = altEndHeight
		}
	}

	logs := make([]evm.FullLog, 0)
	if startHeight >= endHeight {
		return logs, nil
	}

	for i := startHeight; i <= endHeight; i++ {
		select {
		case <-ctx.Done():
			return nil, errors.New("call timed out")
		default:
		}
		blockInfo, err := txdb.GetBlock(i)
		if err != nil {
			return nil, err
		}
		if blockInfo == nil {
			// No arbitrum txes in this block
			continue
		}
		if !maybeMatchesLogQuery(blockInfo.Bloom, address, topics) {
			continue
		}

		res, err := evm.NewBlockResultFromValue(blockInfo.BlockLog)
		if err != nil {
			return nil, err
		}

		first := res.FirstAVMLog().Uint64()
		for j := uint64(0); j < res.BlockStats.AVMLogCount.Uint64(); j++ {
			logVal, err := txdb.GetLog(first + j)
			if err != nil {
				return nil, err
			}

			res, err := evm.NewTxResultFromValue(logVal)
			if err != nil {
				return nil, err
			}

			logIndex := uint64(0)
			for _, evmLog := range res.EVMLogs {
				if evmLog.MatchesQuery(address, topics) {
					logs = append(logs, evm.FullLog{
						Log:     evmLog,
						TxIndex: j,
						TxHash:  res.IncomingRequest.MessageID,
						Index:   logIndex,
						Block: &common.BlockId{
							Height:     common.NewTimeBlocks(new(big.Int).SetUint64(i)),
							HeaderHash: blockInfo.Hash,
						},
					})
				}
				logIndex++
			}
		}
	}
	return logs, nil
}

func maybeMatchesLogQuery(logFilter types.Bloom, addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		match := false
		for _, addr := range addresses {
			if logFilter.TestBytes(addr[:]) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	for _, topicGroup := range topics {
		if len(topicGroup) == 0 {
			continue
		}
		match := false
		for _, topic := range topicGroup {
			if logFilter.TestBytes(topic[:]) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}
