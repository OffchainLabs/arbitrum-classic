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

package rollupvalidator

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"google.golang.org/protobuf/proto"
)

var txRecordPrefix = []byte{53}

type nodeRecordKey struct {
	height uint64
	hash   common.Hash
}

type txDB struct {
	db                 machine.CheckpointStorage
	confirmedNodeStore machine.NodeStore

	transactions     map[common.Hash]*TxRecord
	nodeInfo         map[nodeRecordKey]*nodeInfo
	nodeHashLookup   map[common.Hash]uint64
	nodeHeightLookup map[uint64]common.Hash
	chainAddress     common.Address
}

func newTxDB(db machine.CheckpointStorage, ns machine.NodeStore, chainAddress common.Address) *txDB {
	return &txDB{
		db:                 db,
		confirmedNodeStore: ns,
		transactions:       make(map[common.Hash]*TxRecord),
		nodeInfo:           make(map[nodeRecordKey]*nodeInfo),
		nodeHashLookup:     make(map[common.Hash]uint64),
		nodeHeightLookup:   make(map[uint64]common.Hash),
	}
}

func txRecordKey(txHash common.Hash) []byte {
	return append(append([]byte{}, txRecordPrefix...), txHash.Bytes()...)
}

func (txdb *txDB) removeUnconfirmedNode(nodeHeight uint64) error {
	node, err := txdb.lookupNodeWithHeight(nodeHeight)
	if err != nil {
		return err
	}
	for _, txHash := range node.TransactionHashes {
		delete(txdb.transactions, txHash)
	}
	delete(txdb.nodeHeightLookup, nodeHeight)
	delete(txdb.nodeHashLookup, node.NodeHash)
	delete(txdb.nodeInfo, nodeRecordKey{
		height: nodeHeight,
		hash:   node.NodeHash,
	})
	return nil
}

func (txdb *txDB) addUnconfirmedNode(info *nodeInfo, txes []txRecordInfo) {
	for _, tx := range txes {
		txdb.transactions[tx.txHash] = tx.record
	}
	txdb.nodeHeightLookup[info.NodeHeight] = info.NodeHash
	txdb.nodeHashLookup[info.NodeHash] = info.NodeHeight
	txdb.nodeInfo[nodeRecordKey{
		height: info.NodeHeight,
		hash:   info.NodeHash,
	}] = info
}

func (txdb *txDB) confirmNode(nodeHash common.Hash) error {
	node, err := txdb.lookupNodeWithHash(nodeHash)
	if err != nil {
		return err
	}
	for _, txHash := range node.TransactionHashes {
		if err := txdb.confirmTx(txHash); err != nil {
			return err
		}
	}
	return nil
}

func (txdb *txDB) confirmTx(txHash common.Hash) error {
	txRecord, ok := txdb.transactions[txHash]
	if !ok {
		return errors.New("failed to find transaction while confirming")
	}
	data, err := proto.Marshal(txRecord)
	if err != nil {
		return err
	}
	if !txdb.db.SaveData(txRecordKey(txHash), data) {
		return errors.New("failed to save tx record")
	}

	delete(txdb.transactions, txHash)
	return nil
}

func (txdb *txDB) lookupTxRecord(txHash common.Hash) (*TxRecord, error) {
	txRecord, ok := txdb.transactions[txHash]
	if ok {
		return txRecord, nil
	}

	txData := txdb.db.GetData(txRecordKey(txHash))
	if len(txData) == 0 {
		return nil, errors.New("tx not found")
	}

	if err := proto.Unmarshal(txData, txRecord); err != nil {
		return nil, err
	}
	return txRecord, nil
}

func (txdb *txDB) lookupNodeWithHeight(nodeHeight uint64) (*nodeInfo, error) {
	nodeHash, found := txdb.nodeHeightLookup[nodeHeight]
	if found {
		return txdb.lookupNodeRecord(nodeHeight, nodeHash)
	}

	nodeHash, err := txdb.confirmedNodeStore.GetNodeHash(nodeHeight)
	if err != nil {
		return nil, err
	}
	return txdb.lookupNodeRecord(nodeHeight, nodeHash)
}

func (txdb *txDB) lookupNodeWithHash(nodeHash common.Hash) (*nodeInfo, error) {
	nodeHeight, found := txdb.nodeHashLookup[nodeHash]
	if found {
		return txdb.lookupNodeRecord(nodeHeight, nodeHash)
	}

	nodeHeight, err := txdb.confirmedNodeStore.GetNodeHeight(nodeHash)
	if err != nil {
		return nil, err
	}
	return txdb.lookupNodeRecord(nodeHeight, nodeHash)
}

func (txdb *txDB) lookupNodeRecord(nodeHeight uint64, nodeHash common.Hash) (*nodeInfo, error) {
	key := nodeRecordKey{height: nodeHeight, hash: nodeHash}
	nodeInfo, ok := txdb.nodeInfo[key]
	if ok {
		return nodeInfo, nil
	}

	nodeData, err := txdb.confirmedNodeStore.GetNode(nodeHeight, nodeHash)
	if err != nil {
		return nil, err
	}
	var nodeBuf *structures.NodeBuf
	if err := proto.Unmarshal(nodeData, nodeBuf); err != nil {
		return nil, err
	}
	restoreContext := ckptcontext.NewSimpleRestore(txdb.db)
	node := nodeBuf.UnmarshalFromCheckpoint(restoreContext)

	nodeInfo, _ = processNode(node, txdb.chainAddress)
	txdb.nodeInfo[key] = nodeInfo
	return nodeInfo, nil
}
