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

package ethbridgemachine

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/pkg/errors"
	"math/big"
	"testing"
)

func getTester(t *testing.T) *ethbridgetestcontracts.MachineTester {
	backend, pks := test.SimulatedBackend()
	ctx := context.Background()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])
	authClient, err := ethbridge.NewEthAuthClient(ctx, client, auth)
	if err != nil {
		t.Fatal(err)
	}

	machineTesterAddr, machineTx, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgetestcontracts.DeployMachineTester(auth, client)
	})
	if err != nil {
		t.Fatal(err)
	}

	client.Commit()

	_, err = ethbridge.WaitForReceiptWithResults(
		ctx,
		client,
		auth.From,
		machineTx,
		"deployedMachineTester",
	)
	if err != nil {
		t.Fatal(err)
	}

	deployedMachineTester, err := ethbridgetestcontracts.NewMachineTester(machineTesterAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	return deployedMachineTester
}

func TestDeserializeMachine(t *testing.T) {
	machineTester := getTester(t)
	machine, err := loader.LoadMachineFromFile(arbos.Path(), true, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	stateData, err := machine.MarshalState()
	if err != nil {
		t.Fatal(err)
	}

	expectedHash := machine.Hash()

	offset, bridgeHash, err := machineTester.DeserializeMachine(nil, stateData)
	if err != nil {
		t.Fatal(err)
	}

	if offset.Cmp(big.NewInt(int64(len(stateData)))) != 0 {
		t.Error("incorrect offset")
	}

	if expectedHash.ToEthHash() != bridgeHash {
		t.Log("local hash", expectedHash)
		t.Log("ethbridge hash", hexutil.Encode(bridgeHash[:]))
		t.Error(errors.New("calculated wrong state hash"))
	}
}

func TestAddValueToStack(t *testing.T) {
	machineTester := getTester(t)

	stack := value.NewEmptyTuple()
	intval := value.NewInt64Value(1)

	stack2 := value.NewTuple2(intval, stack)
	expectedHash := stack2.Hash().ToEthHash()

	buf1 := new(bytes.Buffer)
	err := value.MarshalValue(stack, buf1)
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
