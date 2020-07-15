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
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var contractPath = arbos.Path()
var dbPath = "./testdb"

func generateResults() []*evm.Result {
	results := make([]*evm.Result, 0, 5)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}
	return results
}

func setupNodes() ([]*structures.Node, error) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		return nil, err
	}
	node := structures.NewInitialNode(mach.Clone(), common.RandHash())
	nextNode := structures.NewRandomNodeFromValidPrev(node, generateResults())
	return []*structures.Node{node, nextNode}, nil
}

func saveNode(checkpointer *cmachine.CheckpointStorage, ns machine.ConfirmedNodeStore, node *structures.Node) error {
	checkpointContext := ckptcontext.NewCheckpointContext()
	nodeData := node.MarshalForCheckpoint(checkpointContext, false)
	rawNodeData, err := proto.Marshal(nodeData)
	if err != nil {
		return err
	}
	if err := ckptcontext.SaveCheckpointContext(checkpointer, checkpointContext); err != nil {
		return err
	}
	if err := ns.PutNode(node.Depth(), node.Hash(), rawNodeData); err != nil {
		return err
	}
	return nil
}

func TestTrackerDB(t *testing.T) {
	checkpointer, err := cmachine.NewCheckpoint(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := checkpointer.Initialize(contractPath); err != nil {
		t.Fatal(err)
	}

	nodes, err := setupNodes()
	if err != nil {
		t.Fatal(err)
	}

	db, err := newTxDB(checkpointer, checkpointer.GetConfirmedNodeStore())

	heightTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			height, err := db.lookupNodeHeight(node.Hash())
			if err != nil {
				t.Error(err)
			}
			if height != node.Depth() {
				t.Error("wrong height")
			}
		}
	}

	hashTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			hash, err := db.lookupNodeHash(node.Depth())
			if err != nil {
				t.Error(err)
			}
			if hash != node.Hash() {
				t.Error("wrong hash")
			}
		}
	}

	nodeRecordTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			nodeInfo, err := processNode(node)
			if err != nil {
				t.Fatal(err)
			}

			info, err := db.lookupNodeRecord(node.Depth(), node.Hash())
			if err != nil {
				t.Fatal(err)
			}

			if info == nil {
				t.Fatal("node info nil")
			}

			if !info.Equals(nodeInfo) {
				t.Error("nodeInfo not equal")
			}
		}
	}

	nodeMetadataTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			nodeInfo, err := processNode(node)
			if err != nil {
				t.Fatal(err)
			}

			metadata, err := db.lookupNodeMetadata(node.Depth(), node.Hash())
			if err != nil {
				t.Fatal(err)
			}

			if metadata == nil {
				t.Fatal("node info nil")
			}

			if !bytes.Equal(metadata.LogBloom, newNodeMetadata(nodeInfo).LogBloom) {
				t.Error("nodeInfo not equal")
			}
		}
	}

	txRecordTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			nodeInfo, err := processNode(node)
			if err != nil {
				t.Fatal(err)
			}

			for i, txHash := range nodeInfo.EVMTransactionHashes {
				record, err := db.lookupTxRecord(txHash)
				if err != nil {
					t.Fatal(err)
				}

				if record == nil {
					t.Fatal("txrecord nil")
				}

				if !record.Equals(&TxRecord{
					NodeHeight:       node.Depth(),
					NodeHash:         node.Hash().MarshalToBuf(),
					TransactionIndex: uint64(i),
				}) {
					t.Error("Got wrong record")
				}
			}
		}
	}

	for _, node := range nodes {
		t.Run("AddUnconfirmedNode", func(t *testing.T) {
			nodeInfo, err := processNode(node)
			if err != nil {
				t.Fatal(err)
			}
			if err := db.addUnconfirmedNode(nodeInfo); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("UnconfirmedHeightLookup", heightTest(node))
		t.Run("UnconfirmedHashLookup", hashTest(node))
		t.Run("UnconfirmedNodeRecord", nodeRecordTest(node))
		t.Run("UnconfirmedMetadata", nodeMetadataTest(node))
		t.Run("UnconfirmedTxRecord", txRecordTest(node))

		if err := saveNode(checkpointer, db.confirmedNodeStore, node); err != nil {
			t.Fatal(err)
		}

		t.Run("ConfirmNode", func(t *testing.T) {
			if err := db.confirmNode(node.Hash()); err != nil {
				t.Error(err)
			}

			key := nodeRecordKey{
				height: node.Depth(),
				hash:   node.Hash(),
			}

			if _, err := db.getInMemoryNodeData(key); err == nil {
				t.Error("node data should be removed from memory after confirmation")
			}
		})

		t.Run("ConfirmedHeightLookup", heightTest(node))
		t.Run("ConfirmedHashLookup", hashTest(node))
		t.Run("ConfirmedNodeRecord", nodeRecordTest(node))
		t.Run("ConfirmedMetadata", nodeMetadataTest(node))
		t.Run("ConfirmedTxRecord", txRecordTest(node))

		t.Run("CachedConfirmedNodeRecord", nodeRecordTest(node))
	}

	checkpointer.CloseCheckpointStorage()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
}

func TestMetadataLogMatch(t *testing.T) {
	nodes, err := setupNodes()
	if err != nil {
		t.Fatal(err)
	}

	nodeInfo, err := processNode(nodes[1])
	if err != nil {
		t.Fatal(err)
	}

	metadata := newNodeMetadata(nodeInfo)
	flatLogs := nodeInfo.fullLogs()

	evm.LogMatchTest(
		t,
		func(addresses []common.Address, topics [][]common.Hash) bool {
			return metadata.MaybeMatchesLogQuery(addresses, topics)
		},
		flatLogs[0].Log,
	)
}

func TestUnconfirmedDB(t *testing.T) {
	checkpointer, err := cmachine.NewCheckpoint(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := checkpointer.Initialize(contractPath); err != nil {
		t.Fatal(err)
	}

	nodes, err := setupNodes()
	if err != nil {
		t.Fatal(err)
	}

	results := generateResults()

	nodeA := structures.NewRandomNodeFromValidPrev(nodes[0], results)

	nodeB := structures.NewRandomNodeFromValidPrev(nodes[0], results)

	nodeInfoA, err := processNode(nodeA)
	if err != nil {
		t.Fatal(err)
	}

	nodeInfoB, err := processNode(nodeB)
	if err != nil {
		t.Fatal(err)
	}

	db, err := newTxDB(checkpointer, checkpointer.GetConfirmedNodeStore())
	if err != nil {
		t.Fatal(err)
	}

	// Clear location to emulate a pending node
	nodeInfoA.Location = nil

	db.addPendingNode(nodeInfoA)

	txInfoA, err := db.lookupTxInfo(nodeInfoA.EVMTransactionHashes[1])
	if err != nil {
		t.Fatal(err)
	}

	if !nodeInfoA.getTxInfo(1).Equals(txInfoA) {
		t.Error("wrong tx")
	}

	if err := db.addUnconfirmedNode(nodeInfoB); err != nil {
		t.Fatal(err)
	}

	txInfoB, err := db.lookupTxInfo(nodeInfoA.EVMTransactionHashes[1])
	if err != nil {
		t.Fatal(err)
	}

	if !nodeInfoB.getTxInfo(1).Equals(txInfoB) {
		t.Error("wrong tx")
	}

	checkpointer.CloseCheckpointStorage()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
}
