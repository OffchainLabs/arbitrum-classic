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

package nitroexport

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"
	"reflect"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type CrossDB struct {
	txDB  *txdb.TxDB
	ethDB ethdb.Database

	err error

	// atomic reads / writes:
	targetBlock       uint64
	latestStoredBlock uint64

	targetChangedChan chan struct{}
}

func NewCrossDB(
	txDB *txdb.TxDB,
	ethDBPath string,
) (*CrossDB, error) {
	freezer := filepath.Join(ethDBPath, "ancient")
	ethDB, err := rawdb.NewLevelDBDatabaseWithFreezer(ethDBPath, 0, 0, freezer, "", false)

	if err != nil {
		return nil, err
	}
	return &CrossDB{
		err:               nil,
		txDB:              txDB,
		ethDB:             ethDB,
		targetChangedChan: make(chan struct{}),
	}, nil
}

func (c *CrossDB) BlocksExported() (uint64, error) {
	return c.ethDB.Ancients()
}

func (c *CrossDB) importBlock(ctx context.Context, blockNumber uint64) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	machineBlockInfo, err := c.txDB.GetBlock(blockNumber)
	if err != nil {
		return err
	}
	_, txResults, err := c.txDB.GetBlockResults(machineBlockInfo)
	if err != nil {
		return err
	}
	outputTxs := make([]*types.Transaction, 0)
	outputReceipts := make([]*types.Receipt, 0)
	processedTxes := evm.FilterEthTxResults(txResults)

	for _, processedTx := range processedTxes {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		txRes := processedTx.Result

		txHash := txRes.IncomingRequest.MessageID.ToEthHash()
		effectiveGasPrice := txRes.FeeStats.Price.L2Computation.Uint64()
		tx, err := types.NewArbitrumLegacyTx(processedTx.Tx, txHash, effectiveGasPrice, blockNumber)
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
		return errors.Errorf("bad block %v", blockNumber)
	}
	_, err = rawdb.WriteAncientBlocks(c.ethDB, []*types.Block{block}, []types.Receipts{outputReceipts}, big.NewInt(0))
	receiptsRead := rawdb.ReadReceipts(c.ethDB, blockHash, blockNumber, params.ArbitrumRinkebyDevTestChainConfig())
	for i := range receiptsRead {
		if !reflect.DeepEqual(receiptsRead[i], outputReceipts[i]) {
			return errors.New(fmt.Sprintf("stored %v != original %v", receiptsRead[i], outputReceipts[i]))
		}
	}
	return err
}

func (c *CrossDB) mainThread(ctx context.Context) {
	if c.err != nil {
		return
	}
	blockCount, err := c.ethDB.Ancients()
	if err != nil {
		c.err = err
		return
	}
	if blockCount > 0 {
		lastBlock := blockCount - 1
		storedHash := rawdb.ReadCanonicalHash(c.ethDB, lastBlock)
		txBlockInfo, err := c.txDB.GetBlock(lastBlock)
		if err != nil {
			c.err = err
		}
		if txBlockInfo.Header.Hash() != storedHash {
			c.err = errors.Errorf("ethDB holds bad last block")
			return
		}
	}
	for {
		for blockCount < atomic.LoadUint64(&c.targetBlock) {
			err := c.importBlock(ctx, blockCount)
			if err != nil {
				c.err = err
				return
			}
			blockCount++
		}
		select {
		case <-c.targetChangedChan:
		case <-ctx.Done():
			c.err = ctx.Err()
			return
		}
	}
}

func (c *CrossDB) UpdateTarget(target uint64) {
	atomic.StoreUint64(&c.targetBlock, target)
	select {
	case c.targetChangedChan <- struct{}{}:
	default:
	}
}

func (c *CrossDB) CurrentError() error {
	return c.err
}

func (c *CrossDB) Start(ctx context.Context) {
	go c.mainThread(ctx)
}
