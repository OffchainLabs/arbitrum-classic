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

package proofmachine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	stack2 "github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm/stack"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/proofmachine/machinetester"
	"log"
	"testing"
)

var privHex = "27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa"

func getTester(m *testing.T) *machinetester.MachineTester {
	auth, err := test.SetupAuth(privHex)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(test.GetEthUrl())
	if err != nil {
		log.Fatal(err)
	}

	_, machineTx, deployedMachineTester, err := machinetester.DeployMachineTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		machineTx,
		"deployedMachineTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	return deployedMachineTester
}

func TestDeserializeMachine(t *testing.T) {
	machineTester := getTester(t)
	machine, err := loader.LoadMachineFromFile("../contract.ao", true, "test")
	if err != nil {
		t.Fatal(err)
	}

	stateData, err := machine.MarshalState()
	if err != nil {
		t.Fatal(err)
	}

	expectedHash := machine.Hash().ToEthHash()

	bridgeHash, err := machineTester.DeserializeMachine(nil, stateData)
	if err != nil {
		t.Fatal(err)
	}

	if expectedHash != bridgeHash {
		t.Error(errors.New("calculated wrong state hash"))
		fmt.Println(expectedHash)
		fmt.Println(bridgeHash)
	}
}

func TestAddValueToStack(t *testing.T) {
	machineTester := getTester(t)

	stack := stack2.NewEmptyFlat()
	bridgeStack := stack.StateValue()
	intval := value.NewInt64Value(1)

	stack.Push(intval)
	expectedHash := stack.StateValue().Hash().ToEthHash()

	buf1 := new(bytes.Buffer)
	err := value.MarshalValue(bridgeStack, buf1)
	if err != nil {
		t.Fatal(err)
	}
	data1 := buf1.Bytes()

	buf2 := new(bytes.Buffer)
	err = value.MarshalValue(intval, buf2)
	if err != nil {
		t.Fatal(err)
	}

	data2 := buf2.Bytes()

	bridgeHash, err := machineTester.AddStackVal(nil, data1, data2)
	if err != nil {
		fmt.Println(buf1.Bytes())
		fmt.Println(buf2.Bytes())
		t.Fatal(err)
	}

	if expectedHash != bridgeHash {
		t.Error(errors.New("calculated wrong state hash"))
		fmt.Println(expectedHash)
		fmt.Println(bridgeHash)
	}

}
