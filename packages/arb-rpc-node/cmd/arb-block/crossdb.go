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
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pkg/errors"

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
	if blockCount > 0 {
		lastBlock := blockCount - 1
		storedHash := rawdb.ReadCanonicalHash(c.ethDB, lastBlock)
		txBlockInfo, err := c.txDB.GetBlock(lastBlock)
		if err != nil {
			return err
		}
		if txBlockInfo.Header.Hash() != storedHash {
			return errors.Errorf("ethDB holds bad last block")
		}
	}
	if blockCount > limit {
		logger.Info().Uint64("exists", blockCount).Uint64("requested", limit).Msg("block translation done")
		return nil
	}
	for blockCount < limit {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if blockCount%1000 == 0 {
			logger.Info().Uint64("block", blockCount).Msg("importing block")
		}
		machineBlockInfo, err := c.txDB.GetBlock(blockCount)
		if err != nil {
			return err
		}
		_, txResults, err := c.txDB.GetBlockResults(machineBlockInfo)
		if err != nil {
			return err
		}
		outputTxs := make([]*types.Transaction, 0)
		outputReceipts := make([]*types.Receipt, 0)
		for _, txRes := range txResults {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			processedTx, err := evm.GetTransaction(txRes)
			if err != nil {
				return err
			}

			txHash := txRes.IncomingRequest.MessageID.ToEthHash()
			effectiveGasPrice := txRes.FeeStats.Price.L2Computation.Uint64()
			tx, err := types.NewArbitrumLegacyTx(processedTx.Tx, txHash, effectiveGasPrice, blockCount)
			if err != nil {
				return err
			}

			outputTxs = append(outputTxs, tx)
			outputReceipts = append(outputReceipts, txRes.ToEthReceipt(arbcommon.NewHashFromEth(machineBlockInfo.Header.Hash())))
		}
		header := types.CopyHeader(machineBlockInfo.Header)

		block := types.NewBlock(header, outputTxs, nil, outputReceipts, trie.NewStackTrie(nil))
		blockHash := block.Header().Hash()
		if blockHash != machineBlockInfo.Header.Hash() {
			errStr := ""
			for i, tx := range outputTxs {
				errStr += fmt.Sprint(i, ": ", tx.Hash(), "\n")
			}
			return errors.Errorf("bad block ", blockCount, "\n", machineBlockInfo.Header.Hash(), machineBlockInfo.Header, "\n", block.Header().Hash(), block.Header(), "\n", errStr)
		}
		_, err = rawdb.WriteAncientBlocks(c.ethDB, []*types.Block{block}, []types.Receipts{outputReceipts}, big.NewInt(0))
		if err != nil {
			return err
		}
		blockCount++
	}
	return nil
}
