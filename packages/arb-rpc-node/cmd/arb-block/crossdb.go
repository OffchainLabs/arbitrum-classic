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

package main

import (
	"context"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type CrossDB struct {
	txDB  *txdb.TxDB
	ethDB ethdb.Database
}

func New(
	txdb *txdb.TxDB,
	ethDBPath string,
) (*CrossDB, error) {

	freezer := filepath.Join(ethDBPath, "ancient")
	ethDB, err := rawdb.NewLevelDBDatabaseWithFreezer(ethDBPath, 0, 0, freezer, "", false)

	if err != nil {
		return nil, err
	}
	return &CrossDB{
		txDB:  txdb,
		ethDB: ethDB,
	}, nil
}

func (c *CrossDB) EthBlockNum() (uint64, error) {
	return c.ethDB.Ancients()
}

func (c *CrossDB) FillerUp(ctx context.Context, limit uint64) error {
	blockCount, err := c.ethDB.Ancients()
	if err != nil {
		return err
	}
	if blockCount > limit {
		logger.Info().Uint64("exists", blockCount).Uint64("requested", limit).Msg("block translation done")
		return nil
	}
	prevBlockHash := common.Hash{}
	if blockCount > 0 {
		prevBlockHash = rawdb.ReadCanonicalHash(c.ethDB, blockCount-1)
	}
	for blockCount < limit {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		logger.Info().Uint64("block", blockCount).Msg("importing block")
		machineBlockInfo, err := c.txDB.GetBlock(blockCount)
		if err != nil {
			return err
		}
		blockInfo, txResults, err := c.txDB.GetBlockResults(machineBlockInfo)
		if err != nil {
			return err
		}
		outputTxs := make([]*types.Transaction, 0)
		outputReceipts := make([]*types.Receipt, 0)
		for _, txRes := range txResults {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			if (txRes.ResultCode == evm.SequenceNumberTooLow) ||
				(txRes.ResultCode == evm.SequenceNumberTooHigh) {
				continue
			}

			processedTx, err := evm.GetTransaction(txRes)
			if err != nil {
				return err
			}
			tx := processedTx.Tx
			res := processedTx.Result
			vVal, rVal, sVal := tx.RawSignatureValues()

			arblegacy := types.ArbitrumLegacyTxData{
				Gas:      tx.Gas(),
				GasPrice: tx.GasPrice(),
				Hash:     res.IncomingRequest.MessageID.ToEthHash(),
				Data:     tx.Data(),
				Nonce:    tx.Nonce(),
				To:       tx.To(),
				Value:    tx.Value(),
				V:        vVal,
				R:        rVal,
				S:        sVal,
			}

			outputTxs = append(outputTxs, types.NewTx(&arblegacy))
			outputReceipts = append(outputReceipts, txRes.ToEthReceipt(arbcommon.Hash{}))
		}
		//header := types.CopyHeader(machineBlockInfo.Header)
		header := &types.Header{
			ParentHash: prevBlockHash,
			Difficulty: big.NewInt(1), //TODO -
			Number:     new(big.Int).SetUint64(blockCount),
			GasLimit:   blockInfo.GasLimit().Uint64(),
			GasUsed:    blockInfo.BlockStats.GasUsed.Uint64(),
			Time:       blockInfo.Timestamp.Uint64(),
			Root:       machineBlockInfo.Header.Root,
			Extra:      nil,
		}

		header.ParentHash = prevBlockHash
		block := types.NewBlock(header, outputTxs, nil, outputReceipts, trie.NewStackTrie(nil))
		_, err = rawdb.WriteAncientBlocks(c.ethDB, []*types.Block{block}, []types.Receipts{outputReceipts}, new(big.Int).SetUint64(blockCount))
		if err != nil {
			return err
		}
		blockCount++
	}
	return nil
}
