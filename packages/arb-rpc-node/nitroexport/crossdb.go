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
	"encoding/binary"
	"fmt"
	"math"
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

var batchNumKey = hashing.SoliditySHA3([]byte("NextMsgBatch"))

type CrossDB struct {
	txDB  *txdb.TxDB
	ethDB ethdb.Database
	msgDB ethdb.Database

	err error

	// atomic reads / writes:
	targetBlockPlusOne    uint64
	targetMsgBatchPlusOne uint64

	targetChangedChan chan struct{}
}

func NewCrossDB(
	txDB *txdb.TxDB,
	exportPath string,
) (*CrossDB, error) {
	ethDBPath := filepath.Join(exportPath, "l2chaindata")
	freezer := filepath.Join(ethDBPath, "ancient")
	ethDB, err := rawdb.NewLevelDBDatabaseWithFreezer(ethDBPath, 0, 0, freezer, "", false)
	if err != nil {
		return nil, err
	}
	msgDBPath := filepath.Join(exportPath, "classic-msg")
	msgDB, err := rawdb.NewLevelDBDatabase(msgDBPath, 0, 0, "", false)
	if err != nil {
		return nil, err
	}

	return &CrossDB{
		err:               nil,
		txDB:              txDB,
		ethDB:             ethDB,
		msgDB:             msgDB,
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
	blockInfo, txResults, err := c.txDB.GetBlockResults(machineBlockInfo)
	if err != nil {
		return err
	}
	outputTxs := make([]*types.Transaction, 0)
	outputReceipts := make([]*types.Receipt, 0)
	processedTxes := evm.FilterEthTxResults(txResults)

	for i, processedTx := range processedTxes {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		txRes := processedTx.Result

		txHash := txRes.IncomingRequest.MessageID.ToEthHash()
		effectiveGasPrice := txRes.FeeStats.Price.L2Computation.Uint64()
		tx, err := types.NewArbitrumLegacyTx(processedTx.Tx, txHash, effectiveGasPrice, blockInfo.L1BlockNum.Uint64())
		if err != nil {
			return err
		}

		outputTxs = append(outputTxs, tx)
		receipt := txRes.ToEthReceipt(arbcommon.NewHashFromEth(machineBlockInfo.Header.Hash()))
		receipt.TransactionIndex = uint(i)
		for _, log := range receipt.Logs {
			log.TxIndex = uint(i)
		}
		outputReceipts = append(outputReceipts, receipt)
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

func (c *CrossDB) storeMerkle(ctx context.Context, node evm.MerkleNode, batch ethdb.Batch) error {
	key := node.Hash().Bytes()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	switch n := node.(type) {
	case *evm.MerkleInteriorNode:
		if node.Lowest() == node.Highest() {
			return errors.New("one index on internal merkle")
		}
		if err := c.storeMerkle(ctx, n.Left, batch); err != nil {
			return err
		}
		if err := c.storeMerkle(ctx, n.Right, batch); err != nil {
			return err
		}
		data := append(n.Left.Hash().Bytes(), n.Right.Hash().Bytes()...)
		return batch.Put(key, data)
	case *evm.MerkleLeaf:
		return batch.Put(key, n.Data)
	}
	return errors.New("unexpected MerkleNode type")
}

func msgBatchKey(batchNum *big.Int) []byte {
	return hashing.SoliditySHA3(append([]byte("msgBatch"), batchNum.Bytes()...)).Bytes()
}

func (c *CrossDB) BatchesExported() uint64 {
	batchNumBytes, err := c.msgDB.Get(batchNumKey[:])
	if err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(batchNumBytes)
}

func (c *CrossDB) importMsgBatch(ctx context.Context, batchNum uint64) (bool, error) {
	batchNumBig := new(big.Int).SetUint64(batchNum)
	merkle, err := c.txDB.GetMessageBatch(batchNumBig)
	if err != nil {
		return false, err
	}
	if merkle == nil {
		return false, nil
	}
	batch := c.msgDB.NewBatch()
	if err := c.storeMerkle(ctx, merkle.Tree, batch); err != nil {
		return false, err
	}
	batchEntry := make([]byte, 8)
	binary.BigEndian.PutUint64(batchEntry[0:8], merkle.NumInBatch.Uint64())
	batchEntry = append(batchEntry, merkle.Tree.Hash().Bytes()...)
	batch.Put(msgBatchKey(batchNumBig), batchEntry)
	return true, batch.Write()
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
	batchNum := c.BatchesExported()
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
		flag := true
		for flag {
			flag = false
			if batchNum < atomic.LoadUint64(&c.targetMsgBatchPlusOne) {
				batchFound, err := c.importMsgBatch(ctx, batchNum)
				if err != nil {
					c.err = err
					return
				}
				if batchFound {
					batchNum++
					batchNumBytes := make([]byte, 8)
					binary.BigEndian.PutUint64(batchNumBytes[:], batchNum)
					err = c.msgDB.Put(batchNumKey[:], batchNumBytes)
					if err != nil {
						c.err = err
						return
					}
					flag = true
				}
			}
			if blockCount < atomic.LoadUint64(&c.targetBlockPlusOne) {
				err := c.importBlock(ctx, blockCount)
				if err != nil {
					c.err = err
					return
				}
				blockCount++
				flag = true
			}
		}
		select {
		case <-c.targetChangedChan:
		case <-ctx.Done():
			c.err = ctx.Err()
			return
		}
	}
}

func (c *CrossDB) UpdateTargetBatch(target uint64) {
	if target != math.MaxUint64 {
		target += 1
	}
	atomic.StoreUint64(&c.targetMsgBatchPlusOne, target)
	select {
	case c.targetChangedChan <- struct{}{}:
	default:
	}
}

func (c *CrossDB) UpdateTargetBlock(target uint64) {
	atomic.StoreUint64(&c.targetBlockPlusOne, target+1)
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
