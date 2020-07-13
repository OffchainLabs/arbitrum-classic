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

package ethbridgetest

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"testing"
)

func TestGenerateAssertionHash(t *testing.T) {
	intVal := value.NewInt64Value(1)
	intVal2 := value.NewInt64Value(2)
	tupVal := value.NewTuple2(intVal, intVal)

	afterHash := intVal.Hash()
	firstMessageHash := tupVal.Hash()
	lastMessageHash := value.NewTuple2(tupVal, intVal).Hash()
	firstLogHash := value.NewTuple2(tupVal, intVal2).Hash()
	lastLogHash := value.NewTuple2(intVal2, tupVal).Hash()

	assertStub := valprotocol.ExecutionAssertionStub{}
	assertStub.AfterHash = afterHash
	assertStub.FirstMessageHash = firstMessageHash
	assertStub.LastMessageHash = lastMessageHash
	assertStub.FirstLogHash = firstLogHash
	assertStub.LastLogHash = lastLogHash
	assertStub.DidInboxInsn = true
	assertStub.NumGas = 10

	expectedHash := assertStub.Hash()

	ethbridgeHash, err := protocolTester.GenerateAssertionHash(
		nil,
		afterHash,
		true,
		10,
		firstMessageHash,
		lastMessageHash,
		firstLogHash,
		lastLogHash)
	if err != nil {
		t.Fatal(err)
	}

	if expectedHash != ethbridgeHash {
		t.Error(errors.New("calculated wrong assertion hash"))
	}
}
