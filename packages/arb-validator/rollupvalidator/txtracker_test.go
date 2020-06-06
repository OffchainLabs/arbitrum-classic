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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
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

	for _, node := range nodes {
		t.Run("AdvancedKnownNode", func(t *testing.T) {
			txTracker.AdvancedKnownNode(context.Background(), nil, node)
		})

		t.Run("AssertionCountBeforeConfirm", countTest(node))

		if err := saveNode(checkpointer, ns, node); err != nil {
			t.Fatal(err)
		}

		t.Run("ConfirmNode", func(t *testing.T) {
			txTracker.ConfirmedNode(context.Background(), nil, arbbridge.ConfirmedEvent{
				ChainInfo: arbbridge.ChainInfo{},
				NodeHash:  node.Hash(),
			})
		})

		t.Run("AssertionCountAfterConfirm", countTest(node))
	}

	txTracker.RestartingFromLatestValid(context.Background(), nil, nodes[0])

	t.Run("AssertionCountAfterReorg", countTest(nodes[0]))

	checkpointer.CloseCheckpointStorage()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
}
