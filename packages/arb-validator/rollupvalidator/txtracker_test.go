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
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"log"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func TestTxTracker(t *testing.T) {
	checkpointer, nodes, err := setupContext()
	if err != nil {
		t.Fatal(err)
	}

	ns := checkpointer.GetNodeStore()
	txTracker, err := newTxTracker(checkpointer, ns, chainAddress)
	if err != nil {
		t.Fatal(err)
	}

	countTest := func(node *structures.Node) func(*testing.T) {
		return func(t *testing.T) {
			count, err := txTracker.AssertionCount(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			if count != uint64(node.Depth()) {
				t.Error("wrong assertion count")
			}
		}
	}

	findLogsTest := func(t *testing.T) {
		logs, err := txTracker.FindLogs(context.Background(), nil, nil, nil, []common.Hash{{}})
		if err != nil {
			t.Fatal(err)
		}
		if len(logs) != 0 {
			t.Error("wrong logs count")
		}
	}

	txInfoTest := func(t *testing.T) {
		info, err := txTracker.TxInfo(context.Background(), common.Hash{})
		if err != nil {
			t.Fatal(err)
		}
		if info.Found != false {
			t.Error("found non-existant tx")
		}
	}

	for _, node := range nodes {
		t.Run("AdvancedKnownNode", func(t *testing.T) {
			txTracker.AdvancedKnownNode(context.Background(), nil, node)
		})

		t.Run("AssertionCount", countTest(node))
		t.Run("FindLogs", findLogsTest)
		t.Run("TxInfo", txInfoTest)
	}

	txTracker.ConfirmedNode(context.Background(), nil, arbbridge.ConfirmedEvent{
		ChainInfo: arbbridge.ChainInfo{},
		NodeHash:  nodes[0].Hash(),
	})

	txTracker.RestartingFromLatestValid(context.Background(), nil, nodes[0])

	t.Run("AssertionCountAfterReorg", countTest(nodes[0]))
	t.Run("FindLogsAfterReorg", findLogsTest)
	t.Run("TxInfoAfterReorg", txInfoTest)

	checkpointer.CloseCheckpointStorage()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
}
