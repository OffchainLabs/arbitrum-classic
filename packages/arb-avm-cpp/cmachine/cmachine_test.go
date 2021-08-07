/*
* Copyright 2019-2021, Offchain Labs, Inc.
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
	"math/big"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func TestMachineCreation(t *testing.T) {
	dePath := "dbPath"

	if err := os.RemoveAll(dePath); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.RemoveAll(dePath); err != nil {
			t.Fatal(err)
		}
	}()

	mach1, err := New(codeFile)
	if err != nil {
		t.Fatal(err)
	}

	coreConfig := configuration.DefaultCoreSettings()
	arbStorage, err := NewArbStorage(dePath, coreConfig)
	if err != nil {
		t.Fatal(err)
	}
	if err := arbStorage.Initialize(codeFile); err != nil {
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()
	core := arbStorage.GetArbCore()
	cursor, err := core.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	mach2, err := core.TakeMachine(cursor)
	if err != nil {
		t.Fatal(err)
	}

	hash1 := mach1.Hash()
	hash2 := mach2.Hash()
	if hash1 != hash2 {
		t.Fatal(err)
	}
}
