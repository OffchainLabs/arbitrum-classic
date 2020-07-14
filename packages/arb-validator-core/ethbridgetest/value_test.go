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
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func TestTupleHashing(t *testing.T) {

	intVal := value.NewInt64Value(111)
	emptyTup := value.NewEmptyTuple()

	tup := value.NewTuple2(intVal, emptyTup)
	preImage := tup.GetPreImage()

	testTupleBridgeHash, err := valueTester.HashTestTuple(nil)
	if err != nil {
		t.Fatal(err)
	}

	preImageBridgeHash, err := valueTester.HashTuplePreImage(nil, preImage.GetInnerHash(), big.NewInt(preImage.Size()))
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

func TestBytesStack(t *testing.T) {
	data := common.RandBytes(200)
	bytestack := message.BytesToByteStack(data)
	t.Log("bytestack", bytestack)

	bridgeStackHash, err := valueTester.BytesToBytestackHash(nil, data, big.NewInt(0), big.NewInt(int64(len(data))))
	if err != nil {
		t.Fatal(err)
	}

	if bytestack.Hash().ToEthHash() != bridgeStackHash {
		t.Error("calculated wrong byte stack hash")
	}

	var bytestackValBytes bytes.Buffer
	if err := value.MarshalValue(bytestack, &bytestackValBytes); err != nil {
		t.Fatal(err)
	}
	valid, offset, parsedData, err := valueTester.BytestackToBytes(nil, bytestackValBytes.Bytes(), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("failed to parse bytestack")
	}
	if offset.Cmp(big.NewInt(int64(len(bytestackValBytes.Bytes())))) != 0 {
		t.Error("incorrect offset")
	}
	if !bytes.Equal(parsedData, data) {
		t.Error("incorrect data")
	}
}

func TestBytesToBytestackHash(t *testing.T) {
	datas := [][]byte{
		common.RandBytes(5),
		common.RandBytes(32),
		common.RandBytes(33),
		common.RandBytes(64),
		common.RandBytes(200),
	}
	for _, data := range datas {
		valueHash, err := valueTester.BytesToBytestackHash(nil, data, big.NewInt(0), big.NewInt(int64(len(data))))
		if err != nil {
			t.Fatal(err)
		}
		calcDataValue := message.BytesToByteStack(data)
		if calcDataValue.Hash() != valueHash {
			t.Error("hash not equal with data length", len(data))
		}
	}
}

type TestCase struct {
	Value string `json:"value"`
	Hash  string `json:"hash"`
	Name  string `json:"name"`
}

func TestDeserialize(t *testing.T) {
	jsonFile, err := os.Open("../../arb-util/value/test_cases.json")
	if err != nil {
		t.Error(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testCases []TestCase
	err = json.Unmarshal(byteValue, &testCases)
	if err != nil {
		t.Error(err)
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Name == "simple codepoint" || testCase.Name == "immediate codepoint" {
				return
			}
			valBytes, err := hexutil.Decode("0x" + testCase.Value)
			if err != nil {
				t.Error(err)
			}
			val, err := value.UnmarshalValue(bytes.NewReader(valBytes))
			if err != nil {
				t.Error(err)
			}

			offset, valHash, err := valueTester.DeserializeHash(nil, valBytes, big.NewInt(0))
			if err != nil {
				t.Error(err)
			}
			if offset.Cmp(big.NewInt(int64(len(valBytes)))) != 0 {
				t.Errorf("offset was incorrect, was %v, should have been %v", offset, len(valBytes))
			}
			if valHash != val.Hash() {
				t.Error("Incorrect hash")
			}
		})
	}
}
