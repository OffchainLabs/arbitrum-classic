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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/golang-lru"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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

func (x *NodeMetadata) MaybeMatchesLogQuery(addresses []common.Address, topics [][]common.Hash) bool {
	logFilter := types.BytesToBloom(x.LogBloom)

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

// txDB is not fully thread safe, but can be accessed safely by multiple
// readers. Each method of this class is labeled with what type
// of lock its caller requires
type txDB struct {
	db                 machine.CheckpointStorage
	confirmedNodeStore machine.ConfirmedNodeStore
	confirmedNodeCache *lru.Cache

	transactions     map[common.Hash]*TxRecord
	nodeInfo         map[nodeRecordKey]*nodeInfo
	nodeMetadata     map[nodeRecordKey]*NodeMetadata
	nodeHashLookup   map[common.Hash]uint64
	nodeHeightLookup map[uint64]common.Hash

	pendingTransactions map[common.Hash]*evm.TxInfo
}

func newTxDB(db machine.CheckpointStorage, ns machine.ConfirmedNodeStore) (*txDB, error) {
	lruCache, err := lru.New(500)
	if err != nil {
		return nil, err
	}
	return &txDB{
		db:                  db,
		confirmedNodeStore:  ns,
		confirmedNodeCache:  lruCache,
		transactions:        make(map[common.Hash]*TxRecord),
		nodeInfo:            make(map[nodeRecordKey]*nodeInfo),
		nodeMetadata:        make(map[nodeRecordKey]*NodeMetadata),
		nodeHashLookup:      make(map[common.Hash]uint64),
		nodeHeightLookup:    make(map[uint64]common.Hash),
		pendingTransactions: make(map[common.Hash]*evm.TxInfo),
	}, nil
}

// removeUnconfirmedNode requires holding a write lock
func (txdb *txDB) removeUnconfirmedNode(nodeHeight uint64) error {
	nodeHash, err := txdb.lookupNodeHash(nodeHeight)
	if err != nil {
		return err
	}
	node, err := txdb.lookupNodeRecord(nodeHeight, nodeHash)
	if err != nil {
		return err
	}
	txdb.deleteTransactions(node.EVMTransactionHashes)
	txdb.deleteNode(node.Location)
	return nil
}

// addUnconfirmedNode requires holding a write lock
func (txdb *txDB) addUnconfirmedNode(info *nodeInfo) error {
	location := info.Location
	if location == nil {
		return errors.New("node has no location")
	}
	nodeHash := location.NodeHashVal()
	for i, txHash := range info.EVMTransactionHashes {
		txdb.transactions[txHash] = &TxRecord{
			NodeHeight:       location.NodeHeight,
			NodeHash:         nodeHash.MarshalToBuf(),
			TransactionIndex: uint64(i),
		}
		// If we had a pending result for this transaction, delete it
		delete(txdb.pendingTransactions, txHash)
	}
	txdb.nodeHeightLookup[location.NodeHeight] = nodeHash
	txdb.nodeHashLookup[nodeHash] = location.NodeHeight
	key := nodeRecordKey{
		height: location.NodeHeight,
		hash:   nodeHash,
	}
	txdb.nodeInfo[key] = info
	txdb.nodeMetadata[key] = newNodeMetadata(info)
	return nil
}

// addUnconfirmedNode requires holding a write lock
func (txdb *txDB) addPendingNode(info *nodeInfo) {
	for i := range info.EVMTransactionHashes {
		info := info.getTxInfo(uint64(i))
		txdb.pendingTransactions[info.TransactionHash] = info
	}
}

// confirmNode requires holding a write lock
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
	for _, txHash := range node.EVMTransactionHashes {
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

	txdb.deleteTransactions(node.EVMTransactionHashes)
	txdb.deleteNode(node.Location)
	return nil
}

func (txdb *txDB) lookupTxInfo(txHash common.Hash) (*evm.TxInfo, error) {
	pendingInfo, found := txdb.pendingTransactions[txHash]
	if found {
		return pendingInfo, nil
	}
	tx, err := txdb.lookupTxRecord(txHash)
	if err != nil || tx == nil {
		return nil, err
	}
	nodeInfo, err := txdb.lookupNodeRecord(tx.NodeHeight, tx.NodeHash.Unmarshal())
	if err != nil {
		return nil, nil
	}
	return nodeInfo.getTxInfo(tx.TransactionIndex), nil
}

// lookupTxRecord requires holding a read lock
// If the record is found it returns it. If the record is not found, it
// returns nil with no error, otherwise it returns an error
func (txdb *txDB) lookupTxRecord(txHash common.Hash) (*TxRecord, error) {
	txRecord, ok := txdb.transactions[txHash]
	if ok {
		return txRecord, nil
	}

	txData := txdb.db.GetData(txRecordKey(txHash))
	if len(txData) == 0 {
		return nil, nil
	}

	txRecord = &TxRecord{}
	if err := proto.Unmarshal(txData, txRecord); err != nil {
		return nil, err
	}
	return txRecord, nil
}

// lookupNodeHash requires holding a read lock
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

// lookupNodeHeight requires holding a read lock
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

// lookupNodeRecord requires holding a read lock
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
	node, err := nodeBuf.UnmarshalFromCheckpoint(restoreContext)
	if err != nil {
		return nil, err
	}

	info, err = processNode(node)
	if err != nil {
		return nil, err
	}
	txdb.confirmedNodeCache.Add(key, info)
	return info, nil
}

// lookupNodeMetadata requires holding a read lock
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

// getInMemoryNodeData is a private method that should not be called externally
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

// deleteNode is a private method that should not be called externally
func (txdb *txDB) deleteNode(location *evm.NodeLocation) {
	nodeHash := location.NodeHashVal()
	delete(txdb.nodeHeightLookup, location.NodeHeight)
	delete(txdb.nodeHashLookup, nodeHash)
	key := nodeRecordKey{
		height: location.NodeHeight,
		hash:   nodeHash,
	}
	delete(txdb.nodeInfo, key)
	delete(txdb.nodeMetadata, key)
}

// deleteTransactions is a private method that should not be called externally
func (txdb *txDB) deleteTransactions(txHashes []common.Hash) {
	for _, txHash := range txHashes {
		delete(txdb.transactions, txHash)
	}
}
