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
	"math/big"
	"os"
	"testing"
)

func TestMessageBatch(t *testing.T) {
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

	arbStorage, err := NewArbStorage(dePath)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	nodeStore := arbStorage.GetNodeStore()
	testBatchNumber := big.NewInt(42)
	testLogIndex := uint64(0xDEADBEEFA1B2C3D4)
	err = nodeStore.SaveMessageBatch(testBatchNumber, testLogIndex)
	if err != nil {
		logger.Error().Stack().Err(err).Send()
		t.Fatal(err)
	}

	logIndex := nodeStore.GetMessageBatch(testBatchNumber)
	if logIndex == nil {
		t.Fatal("GetMessageBatch should have returned log")
	}
	if *logIndex != testLogIndex {
		logger.Error().Msg("logIndex doesnt match testLogIndex")
	}
}
