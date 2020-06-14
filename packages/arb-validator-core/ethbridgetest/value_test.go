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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
	"testing"
)

func TestEmptyTupleHashing(t *testing.T) {

	tup := value.NewEmptyTuple()
	preImage := tup.GetPreImage()

	emptyBridgeHash, err := valueTester.HashEmptyTuple(nil)
	if err != nil {
		t.Fatal(err)
	}

	preImageBridgeHash, err := valueTester.HashTuplePreImage(nil, preImage.HashImage, big.NewInt(preImage.Size))
	if err != nil {
		t.Fatal(err)
	}

	if preImage.Hash().ToEthHash() != preImageBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}

	if tup.Hash().ToEthHash() != emptyBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}

	if preImage.Hash().ToEthHash() != emptyBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}
}

func TestTupleHashing(t *testing.T) {

	intVal := value.NewInt64Value(111)
	emptyTup := value.NewEmptyTuple()

	tup := value.NewTuple2(intVal, emptyTup)
	preImage := tup.GetPreImage()

	testTupleBridgeHash, err := valueTester.HashTestTuple(nil)
	if err != nil {
		t.Fatal(err)
	}

	preImageBridgeHash, err := valueTester.HashTuplePreImage(nil, preImage.HashImage, big.NewInt(preImage.Size))
	if err != nil {
		t.Fatal(err)
	}

	if preImage.Hash().ToEthHash() != preImageBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}

	if tup.Hash().ToEthHash() != testTupleBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}

	if preImage.Hash().ToEthHash() != testTupleBridgeHash {
		t.Error(errors.New("calculated wrong empty tuple hash"))
	}
}

func TestBytesStackHash(t *testing.T) {
	data := []byte{65, 23, 68, 87, 12}
	stackHash := message.BytesToByteStack(data).Hash().ToEthHash()

	bridgeStackHash, err := valueTester.BytesToBytestackHash(nil, data)
	if err != nil {
		t.Fatal(err)
	}

	if stackHash != bridgeStackHash {
		t.Error(errors.New("calculated wrong byte stack hash: "))
		fmt.Println(stackHash)
		fmt.Println(bridgeStackHash)
	}
}
