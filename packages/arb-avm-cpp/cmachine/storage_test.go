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

package cmachine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"os"
	"testing"
)

var codeFile = gotest.OpCodeTestFiles()[0]

func TestCheckpoint(t *testing.T) {
	dePath := "dbPath"

	defer func() {
		if err := os.RemoveAll(dePath); err != nil {
			logger.Error().Stack().Err(err).Send()
			t.Fatal(err)
		}
	}()

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := arbStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()
}

func TestCheckpointMachine(t *testing.T) {
	dePath := "dbPath2"

	defer func() {
		if err := os.RemoveAll(dePath); err != nil {
			logger.Error().Stack().Err(err).Send()
			t.Fatal(err)
		}
	}()

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := arbStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()

	mach, err := arbStorage.GetInitialMachine()
	if err != nil {
		t.Error(err)
	}

	hash, err := mach.Hash()
	if err != nil {
		t.Error(err)
	}
	t.Log("Initial machine hash", hash)

	_, _, numSteps := mach.ExecuteAssertion(
		1000,
		true,
		nil,
		false,
	)

	t.Log("Ran machine for", numSteps, "steps")

	if !mach.Checkpoint(arbStorage) {
		t.Error("Failed to checkpoint machine")
	}

	hash, err = mach.Hash()
	if err != nil {
		t.Error(err)
	}

	loadedMach, err := arbStorage.GetMachine(hash)
	if err != nil {
		t.Error(err)
	}

	hash1, err := mach.Hash()
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	hash2, err := loadedMach.Hash()
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	if hash1 != hash2 {
		t.Error("Restored machine with wrong hash", hash1, hash2)
	}
}
