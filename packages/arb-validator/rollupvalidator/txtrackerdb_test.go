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

var contractPath = "../contract.ao"
var dbPath = "./testdb"
var chainAddress = common.Address{65, 87, 32}

func setupNodes() ([]*structures.Node, []evm.Result, error) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		return nil, nil, err
	}
	node := structures.NewInitialNode(mach.Clone())

	results := make([]evm.Result, 0, 5)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomStop(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := structures.NewRandomNodeFromValidPrev(node, results)
	return []*structures.Node{node, nextNode}, results, nil
}

func saveNode(checkpointer *cmachine.CheckpointStorage, ns machine.NodeStore, node *structures.Node) error {
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
	checkpointer, err := cmachine.NewCheckpoint(dbPath, contractPath)
	if err != nil {
		t.Fatal(err)
	}

	nodes, _, err := setupNodes()
	if err != nil {
		t.Fatal(err)
	}

	db, err := newTxDB(checkpointer, checkpointer.GetNodeStore(), chainAddress)

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
			nodeInfo, _ := processNode(node, chainAddress)

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
			nodeInfo, _ := processNode(node, chainAddress)

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
			_, txInfos := processNode(node, chainAddress)

			for _, r := range txInfos {
				record, err := db.lookupTxRecord(r.txHash)
				if err != nil {
					t.Fatal(err)
				}

				if record == nil {
					t.Fatal("txrecord nil")
				}

				if !record.Equals(r.record) {
					t.Error("Got wrong record")
				}
			}
		}
	}

	for _, node := range nodes {
		t.Run("AddUnconfirmedNode", func(t *testing.T) {
			nodeInfo, transactions := processNode(node, chainAddress)
			db.addUnconfirmedNode(nodeInfo, transactions)
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
	nodes, results, err := setupNodes()
	if err != nil {
		t.Fatal(err)
	}

	nodeInfo, _ := processNode(nodes[1], chainAddress)
	metadata := newNodeMetadata(nodeInfo)
	flatLogs := extractLogResponses(results)

	evm.LogMatchTest(
		t,
		func(addresses []common.Address, topics [][]common.Hash) bool {
			return metadata.MaybeMatchesLogQuery(addresses, topics)
		},
		flatLogs[0].Log,
	)
}
