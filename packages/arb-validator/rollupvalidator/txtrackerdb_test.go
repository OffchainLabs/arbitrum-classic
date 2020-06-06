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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var contractPath = "../contract.ao"
var dbPath = "./testdb"
var chainAddress = common.Address{65, 87, 32}

func setupContext() (*cmachine.CheckpointStorage, []*structures.Node, error) {
	cCheckpointer, err := cmachine.NewCheckpoint(dbPath, contractPath)
	if err != nil {
		return nil, nil, err
	}

	mach, err := cCheckpointer.GetInitialMachine()
	if err != nil {
		return nil, nil, err
	}
	node := structures.NewInitialNode(mach.Clone())

	tb := &protocol.TimeBounds{
		LowerBoundBlock:     common.NewTimeBlocksInt(0),
		UpperBoundBlock:     common.NewTimeBlocksInt(0),
		LowerBoundTimestamp: big.NewInt(0),
		UpperBoundTimestamp: big.NewInt(0),
	}
	assertion, numSteps := mach.ExecuteAssertion(
		0,
		tb,
		value.NewEmptyTuple(),
		time.Second*10,
	)

	nextNode := structures.NewNodeFromValidPrev(
		node,
		&valprotocol.DisputableNode{
			AssertionParams: &valprotocol.AssertionParams{
				NumSteps:             numSteps,
				TimeBounds:           tb,
				ImportedMessageCount: big.NewInt(0),
			},
			AssertionClaim: &valprotocol.AssertionClaim{
				AfterInboxTop:         common.Hash{},
				ImportedMessagesSlice: common.Hash{},
				AssertionStub:         valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
			},
			MaxInboxTop:   common.Hash{},
			MaxInboxCount: big.NewInt(0),
		},
		valprotocol.ChainParams{
			StakeRequirement:        nil,
			GracePeriod:             common.TimeTicks{Val: big.NewInt(100000)},
			MaxExecutionSteps:       10000,
			MaxBlockBoundsWidth:     10,
			MaxTimestampBoundsWidth: 10,
			ArbGasSpeedLimitPerTick: 10000000,
		},
		common.NewTimeBlocksInt(0),
		common.Hash{54, 87, 23, 65},
	)

	if err := nextNode.UpdateValidOpinion(mach.Clone(), assertion); err != nil {
		return nil, nil, err
	}
	return cCheckpointer, []*structures.Node{node, nextNode}, nil
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
	checkpointer, nodes, err := setupContext()
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

	for _, node := range nodes {
		t.Run("AddUnconfirmedNode", func(t *testing.T) {
			nodeInfo, transactions := processNode(node, chainAddress)
			db.addUnconfirmedNode(nodeInfo, transactions)
		})

		t.Run("UnconfirmedHeightLookup", heightTest(node))
		t.Run("UnconfirmedHashLookup", hashTest(node))

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
	}
	checkpointer.CloseCheckpointStorage()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
}
