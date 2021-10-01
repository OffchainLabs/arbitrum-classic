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

package challenge

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func getTester(t *testing.T) *ethbridgetestcontracts.MachineTester {
	backend, auths := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	_, _, machineTester, err := ethbridgetestcontracts.DeployMachineTester(auth, client)
	test.FailIfError(t, err)
	client.Commit()
	return machineTester
}

func TestDeserializeMachine(t *testing.T) {
	machineTester := getTester(t)
	arbosPath, err := arbos.Path(false)
	test.FailIfError(t, err)
	machine, err := cmachine.New(arbosPath)
	test.FailIfError(t, err)

	stateData, err := machine.MarshalState()
	test.FailIfError(t, err)

	expectedHash := machine.Hash()
	offset, bridgeHash, err := machineTester.DeserializeMachine(nil, stateData)
	test.FailIfError(t, err)

	if offset.Cmp(big.NewInt(int64(len(stateData)))) != 0 {
		t.Error("incorrect offset")
	}

	if expectedHash.ToEthHash() != bridgeHash {
		t.Log("local hash", expectedHash)
		t.Log("ethbridge hash", hexutil.Encode(bridgeHash[:]))
		t.Error(errors.New("calculated wrong state hash"))
	}
}

//func TestAddValueToStack(t *testing.T) {
//	machineTester := getTester(t)
//
//	stack := value.NewEmptyTuple()
//	intval := value.NewInt64Value(1)
//
//	stack2 := value.NewTuple2(intval, stack)
//	expectedHash := stack2.Hash().ToEthHash()
//
//	buf1 := new(bytes.Buffer)
//	err := value.MarshalValue(stack, buf1)
//	if err != nil {
//		t.Fatal(err)
//	}
//	data1 := buf1.Bytes()
//
//	buf2 := new(bytes.Buffer)
//	err = value.MarshalValue(intval, buf2)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	data2 := buf2.Bytes()
//
//	bridgeHash, err := machineTester.AddStackVal(nil, data1, data2)
//	if err != nil {
//		fmt.Println(buf1.Bytes())
//		fmt.Println(buf2.Bytes())
//		t.Fatal(err)
//	}
//
//	if expectedHash != bridgeHash {
//		t.Error(errors.New("calculated wrong state hash"))
//		fmt.Println(expectedHash)
//		fmt.Println(bridgeHash)
//	}
//
//}
