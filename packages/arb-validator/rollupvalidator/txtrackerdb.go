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
	"encoding/binary"
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/golang-lru"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var txRecordPrefix = []byte{53}
var nodeRecordPrefix = []byte{54}

type nodeRecordKey struct {
	height uint64
	hash   common.Hash
}

func txRecordKey(txHash common.Hash) []byte {
	return append(append([]byte{}, txRecordPrefix...), txHash.Bytes()...)
}

func nodeMetadataKey(key nodeRecordKey) []byte {
	encodedHeight := make([]byte, 8)
	binary.BigEndian.PutUint64(encodedHeight, key.height)

	rawKey := append([]byte{}, nodeRecordPrefix...)
	rawKey = append(rawKey, encodedHeight...)
	return append(rawKey, key.hash.Bytes()...)
}

func newNodeMetadata(node *nodeInfo) *NodeMetadata {
	return &NodeMetadata{LogBloom: node.calculateBloomFilter().Bytes()}
}

type txDB struct {
	db                 machine.CheckpointStorage
	confirmedNodeStore machine.NodeStore

	confirmedNodeCache *lru.Cache

	transactions     map[common.Hash]*TxRecord
	nodeInfo         map[nodeRecordKey]*nodeInfo
	nodeMetadata     map[nodeRecordKey]*NodeMetadata
	nodeHashLookup   map[common.Hash]uint64
	nodeHeightLookup map[uint64]common.Hash
	chainAddress     common.Address
}

func newTxDB(db machine.CheckpointStorage, ns machine.NodeStore, chainAddress common.Address) (*txDB, error) {
	lruCache, err := lru.New(500)
	if err != nil {
		return nil, err
	}
	return &txDB{
		db:                 db,
		confirmedNodeStore: ns,
		confirmedNodeCache: lruCache,
		transactions:       make(map[common.Hash]*TxRecord),
		nodeInfo:           make(map[nodeRecordKey]*nodeInfo),
		nodeMetadata:       make(map[nodeRecordKey]*NodeMetadata),
		nodeHashLookup:     make(map[common.Hash]uint64),
		nodeHeightLookup:   make(map[uint64]common.Hash),
		chainAddress:       chainAddress,
	}, nil
}

func (txdb *txDB) removeUnconfirmedNode(nodeHeight uint64) error {
	nodeHash, err := txdb.lookupNodeHash(nodeHeight)
	if err != nil {
		return err
	}
	node, err := txdb.lookupNodeRecord(nodeHeight, nodeHash)
	if err != nil {
		return err
	}
	txdb.deleteNode(node)
	return nil
}

func (txdb *txDB) deleteNode(node *nodeInfo) {
	for _, txHash := range node.TransactionHashes {
		delete(txdb.transactions, txHash)
	}
	delete(txdb.nodeHeightLookup, node.NodeHeight)
	delete(txdb.nodeHashLookup, node.NodeHash)
	key := nodeRecordKey{
		height: node.NodeHeight,
		hash:   node.NodeHash,
	}
	delete(txdb.nodeInfo, key)
	delete(txdb.nodeMetadata, key)
}

func (txdb *txDB) addUnconfirmedNode(info *nodeInfo, txes []txRecordInfo) {
	for _, tx := range txes {
		txdb.transactions[tx.txHash] = tx.record
	}
	txdb.nodeHeightLookup[info.NodeHeight] = info.NodeHash
	txdb.nodeHashLookup[info.NodeHash] = info.NodeHeight
	key := nodeRecordKey{
		height: info.NodeHeight,
		hash:   info.NodeHash,
	}
	txdb.nodeInfo[key] = info
	txdb.nodeMetadata[key] = newNodeMetadata(info)
}

func (txdb *txDB) confirmNode(nodeHash common.Hash) error {
	height, err := txdb.lookupNodeHeight(nodeHash)
	if err != nil {
		return err
	}
	metadata, err := txdb.lookupNodeMetadata(height, nodeHash)
	if err != nil {
		return err
	}
	data, err := proto.Marshal(metadata)
	if err != nil {
		return err
	}
	key := nodeRecordKey{
		height: height,
		hash:   nodeHash,
	}
	if !txdb.db.SaveData(nodeMetadataKey(key), data) {
		return errors.New("failed to save node metadata record")
	}
	node, err := txdb.lookupNodeRecord(height, nodeHash)
	if err != nil {
		return err
	}
	for _, txHash := range node.TransactionHashes {
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
	}

	txdb.deleteNode(node)
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

func (txdb *txDB) lookupNodeHash(nodeHeight uint64) (common.Hash, error) {
	nodeHash, found := txdb.nodeHeightLookup[nodeHeight]
	if found {
		return nodeHash, nil
	}

	nodeHash, err := txdb.confirmedNodeStore.GetNodeHash(nodeHeight)
	if err != nil {
		return common.Hash{}, err
	}
	return nodeHash, nil
}

func (txdb *txDB) lookupNodeHeight(nodeHash common.Hash) (uint64, error) {
	nodeHeight, found := txdb.nodeHashLookup[nodeHash]
	if found {
		return nodeHeight, nil
	}

	nodeHeight, err := txdb.confirmedNodeStore.GetNodeHeight(nodeHash)
	if err != nil {
		return 0, err
	}
	return nodeHeight, nil
}

func (txdb *txDB) getInMemoryNodeData(key nodeRecordKey) (*nodeInfo, error) {
	infoData, ok := txdb.nodeInfo[key]
	if ok {
		return infoData, nil
	}

	nodeInfoCache, ok := txdb.confirmedNodeCache.Get(key)
	if ok {
		return nodeInfoCache.(*nodeInfo), nil
	}
	return nil, errors.New("not found")
}

func (txdb *txDB) lookupNodeRecord(nodeHeight uint64, nodeHash common.Hash) (*nodeInfo, error) {
	key := nodeRecordKey{height: nodeHeight, hash: nodeHash}
	info, err := txdb.getInMemoryNodeData(key)
	if err == nil {
		return info, nil
	}

	nodeData, err := txdb.confirmedNodeStore.GetNode(nodeHeight, nodeHash)
	if err != nil {
		return nil, err
	}
	nodeBuf := &structures.NodeBuf{}
	if err := proto.Unmarshal(nodeData, nodeBuf); err != nil {
		return nil, err
	}
	restoreContext := ckptcontext.NewSimpleRestore(txdb.db)
	node := nodeBuf.UnmarshalFromCheckpoint(restoreContext)

	info, _ = processNode(node, txdb.chainAddress)
	txdb.confirmedNodeCache.Add(key, info)
	return info, nil
}

func (txdb *txDB) lookupNodeMetadata(nodeHeight uint64, nodeHash common.Hash) (*NodeMetadata, error) {
	key := nodeRecordKey{height: nodeHeight, hash: nodeHash}
	metadata, ok := txdb.nodeMetadata[key]
	if ok {
		return metadata, nil
	}

	metadataRaw := txdb.db.GetData(nodeMetadataKey(key))
	if len(metadataRaw) == 0 {
		return nil, errors.New("not found")
	}

	metadata = &NodeMetadata{}
	if err := proto.Unmarshal(metadataRaw, metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}
