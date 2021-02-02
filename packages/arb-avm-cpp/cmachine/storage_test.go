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
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"os"
	"testing"
)

var codeFile = arbos.Path()

func TestCheckpoint(t *testing.T) {
	dePath := "dbPath"

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := arbStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()

	val, err := arbStorage.GetData([]byte("key"))
	if err == nil {
		t.Error("should have failed")
	}

	if len(val) != 0 {
		t.Error("should have empty value")
	}

	if err := os.RemoveAll(dePath); err != nil {
		t.Fatal(err)
	}
}

func TestCheckpointMachine(t *testing.T) {
	dePath := "dbPath2"

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

	t.Log("Initial machine hash", mach.Hash())

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

	loadedMach, err := arbStorage.GetMachine(mach.Hash())
	if err != nil {
		t.Error(err)
	}

	if mach.Hash() != loadedMach.Hash() {
		t.Error("Restored machine with wrong hash", mach.Hash(), loadedMach.Hash())
	}

	if err := os.RemoveAll(dePath); err != nil {
		t.Fatal(err)
	}
}
