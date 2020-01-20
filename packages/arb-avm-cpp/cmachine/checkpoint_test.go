/*
* Copyright 2019, Offchain Labs, Inc.
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

package cmachine

import (
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func TestCheckpoint(t *testing.T) {
	codeFile := "contract.ao"
	dePath := "dbPath"

	checkpointStorage, err := NewCheckpoint(dePath, codeFile)
	if err != nil {
		t.Error(err)
	}
	defer checkpointStorage.CloseCheckpointStorage()

	val := checkpointStorage.GetData([]byte("key"))

	if len(val) != 0 {
		t.Error("should have empty value")
	}

	if err := os.RemoveAll(dePath); err != nil {
		log.Fatal(err)
	}
}

func TestCheckpointMachine(t *testing.T) {
	codeFile := "contract.ao"
	dePath := "dbPath2"

	checkpointStorage, err := NewCheckpoint(dePath, codeFile)
	if err != nil {
		t.Fatal(err)
	}
	defer checkpointStorage.CloseCheckpointStorage()

	mach, err := checkpointStorage.GetInitialMachine()
	if err != nil {
		t.Error(err)
	}

	_, numSteps := mach.ExecuteAssertion(1000, &protocol.TimeBoundsBlocks{
		Start: common.NewTimeBlocks(big.NewInt(100)),
		End:   common.NewTimeBlocks(big.NewInt(120)),
	}, value.NewEmptyTuple())

	t.Log("Ran machine for", numSteps, "steps")

	if !mach.Checkpoint(checkpointStorage) {
		t.Error("Failed to checkpoint machine")
	}

	mach2, err := checkpointStorage.GetInitialMachine()
	if err != nil {
		t.Error(err)
	}

	if !mach2.RestoreCheckpoint(checkpointStorage, mach.Hash()) {
		t.Error("Failed to restore machine")
	}

	if err := os.RemoveAll(dePath); err != nil {
		log.Fatal(err)
	}
}
