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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"

	"testing"
)

//func TestGenerateLastMessageHash(t *testing.T){
//	msgs := make([]value.Value, 0)
//	for i := 0; i < 5; i++ {
//		//intVal := value.NewHashOnlyValueFromValue(value.NewInt64Value(1))
//		intVal := value.NewEmptyTuple()
//		msgs = append(msgs, intVal)
//	}
//	confirmOpp := valprotocol.ConfirmValidOpportunity{}
//	confirmOpp.Messages = msgs
//
//	msgBytes := confirmOpp.MarshalMsgsForConfirmation()
//	expectedCount := 5
//	expectedHash := hashMsgs(msgs)
//
//	ethbridgeHash, msgCounts, err := protocolTester.GenerateLastMessageHash(
//		nil,
//		msgBytes,
//		big.NewInt(0),
//		big.NewInt(int64(len(msgBytes))))
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if expectedHash != ethbridgeHash {
//		t.Error(errors.New("calculated wrong last message hash"))
//		fmt.Println(expectedHash)
//		fmt.Println(ethbridgeHash)
//	}
//
//	if big.NewInt(int64(expectedCount)) != msgCounts {
//		t.Error(errors.New("calculated wrong message count"))
//		fmt.Println(expectedCount)
//		fmt.Println(msgCounts)
//	}
//}
//
//func hashMsgs(msgs []value.Value) common.Hash {
//	currentHash := value.NewEmptyTuple().Hash()
//
//	for _, val := range msgs {
//		currentHash = hash2(currentHash, val.Hash())
//	}
//
//	return currentHash
//}

//func hash2(h1, h2 common.Hash) common.Hash {
//	return hashing.SoliditySHA3(hashing.Bytes32(h1), hashing.Bytes32(h2))
//}

func TestGeneratePreconditionHash(t *testing.T) {
	intVal := value.NewInt64Value(1)
	tuple := value.NewTuple2(intVal, intVal)
	timeBounds := &protocol.TimeBounds{}
	timeBounds.LowerBoundBlock = common.NewTimeBlocks(big.NewInt(0))
	timeBounds.UpperBoundBlock = common.NewTimeBlocks(big.NewInt(1))
	timeBounds.LowerBoundTimestamp = big.NewInt(0)
	timeBounds.UpperBoundTimestamp = big.NewInt(1)

	precondition := valprotocol.NewPrecondition(intVal.Hash(), timeBounds, tuple)
	expectedHash := precondition.Hash().ToEthHash()

	ethbridgeHash, err := protocolTester.GeneratePreconditionHash(
		nil,
		intVal.Hash(),
		timeBounds.AsIntArray(),
		tuple.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if expectedHash != ethbridgeHash {
		t.Error(errors.New("calculated wrong precondition hash"))
		fmt.Println(expectedHash)
		fmt.Println(ethbridgeHash)
	}
}

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
		fmt.Println(expectedHash)
		fmt.Println(ethbridgeHash)
	}
}
