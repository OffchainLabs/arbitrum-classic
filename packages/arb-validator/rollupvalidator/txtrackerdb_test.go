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

func TestMain(m *testing.M) {
	code := m.Run()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestAddUnconfirmedNode(t *testing.T) {
	cCheckpointer, err := cmachine.NewCheckpoint(dbPath, contractPath)
	if err != nil {
		t.Fatal(err)
	}

	db, err := newTxDB(cCheckpointer, cCheckpointer.GetNodeStore(), chainAddress)
	if err != nil {
		t.Fatal(err)
	}

	mach, err := db.db.GetInitialMachine()
	if err != nil {
		t.Fatal(err)
	}
	node := structures.NewInitialNode(mach)

	key := nodeRecordKey{
		height: node.Depth(),
		hash:   node.Hash(),
	}

	assertion, _ := mach.ExecuteAssertion(
		0,
		&protocol.TimeBounds{
			LowerBoundBlock:     common.NewTimeBlocksInt(0),
			UpperBoundBlock:     common.NewTimeBlocksInt(0),
			LowerBoundTimestamp: big.NewInt(0),
			UpperBoundTimestamp: big.NewInt(0),
		},
		value.NewEmptyTuple(),
		time.Second*10,
	)

	nodeInfo, transactions := processNodeImpl(
		node.Hash(),
		node.Depth(),
		node.AssertionTxHash(),
		valprotocol.ValidChildType,
		assertion,
		chainAddress,
	)

	t.Run("AddUnconfirmedNode", func(t *testing.T) {
		db.addUnconfirmedNode(nodeInfo, transactions)
	})

	heightTest := func(t *testing.T) {
		height, err := db.lookupNodeHeight(nodeInfo.NodeHash)
		if err != nil {
			t.Error(err)
		}
		if height != nodeInfo.NodeHeight {
			t.Error("wrong height")
		}
	}

	hashTest := func(t *testing.T) {
		hash, err := db.lookupNodeHash(nodeInfo.NodeHeight)
		if err != nil {
			t.Error(err)
		}
		if hash != nodeInfo.NodeHash {
			t.Error("wrong hash")
		}
	}

	t.Run("UnconfirmedHeightLookup", heightTest)
	t.Run("UnconfirmedHashLookup", hashTest)

	checkpointContext := ckptcontext.NewCheckpointContext()
	nodeData := node.MarshalForCheckpoint(checkpointContext, false)
	rawNodeData, err := proto.Marshal(nodeData)
	if err != nil {
		t.Fatal(err)
	}
	if err := ckptcontext.SaveCheckpointContext(db.db, checkpointContext); err != nil {
		t.Fatal(err)
	}
	if err := db.confirmedNodeStore.PutNode(nodeInfo.NodeHeight, nodeInfo.NodeHash, rawNodeData); err != nil {
		t.Fatal(err)
	}

	t.Run("ConfirmNode", func(t *testing.T) {
		if err := db.confirmNode(nodeInfo.NodeHash); err != nil {
			t.Error(err)
		}

		if _, err := db.getInMemoryNodeData(key); err == nil {
			t.Error("node data should be removed from memory after confirmation")
		}
	})

	t.Run("ConfirmedHeightLookup", heightTest)
	t.Run("ConfirmedHashLookup", hashTest)
}
