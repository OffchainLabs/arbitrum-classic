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
	"os"
	"testing"
)

func TestMachineCreation(t *testing.T) {
	dePath := "dbPath"

	if err := os.RemoveAll(dePath); err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	defer func() {
		if err := os.RemoveAll(dePath); err != nil {
			logger.Error().Stack().Err(err).Send()
			t.Fatal(err)
		}
	}()

	mach1, err := New(codeFile)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	if err := arbStorage.Initialize(codeFile); err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	defer arbStorage.CloseArbStorage()
	mach2, err := arbStorage.GetInitialMachine()
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	hash1, err := mach1.Hash()
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	hash2, err := mach2.Hash()
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
	if hash1 != hash2 {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}
}
