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

package nodeview

import (
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"google.golang.org/protobuf/proto"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

var contractPath = arbos.Path()
var dbPath = "./testdb"
var maxReorgHeight = big.NewInt(100)

var initialEntryBlockId = &common.BlockId{
	Height:     common.NewTimeBlocksInt(10),
	HeaderHash: common.Hash{20},
}
var checkpointData = []byte{5, 3, 2}

func TestMain(m *testing.M) {
	code := m.Run()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestConfirm(t *testing.T) {
	var rollupAddr common.Address
	cp, err := checkpointing.NewIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}

	if err := cp.Initialize(contractPath); err != nil {
		t.Fatal(err)
	}

	checkpointContext := ckptcontext.NewCheckpointContext()
	errChan := cp.AsyncSaveCheckpoint(initialEntryBlockId, checkpointData, checkpointContext)

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second * 5):
		t.Fatal("Timed out saving checkpoint")
	}

	mach, err := cp.GetInitialMachine()
	if err != nil {
		t.Error()
	}

	nd := structures.NewInitialNode(mach, common.Hash{})

	buf := nd.MarshalForCheckpoint(checkpointContext, false)
	data, err := proto.Marshal(buf)
	if err != nil {
		t.Error()
	}

	if err := cp.CheckpointConfirmedNode(
		common.Hash{75},
		3,
		data,
		checkpointContext,
	); err != nil {
		t.Error(err)
	}

	view := New(cp.GetConfirmedNodeStore(), cp.GetCheckpointDB())

	loadedNd, err := view.GetNode(3, common.Hash{75})
	if err != nil {
		t.Fatal(err)
	}

	if !loadedNd.EqualsFull(nd) {
		t.Error("Loaded node not equal to original")
	}
}
