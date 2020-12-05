/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"strconv"

	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"

	"path/filepath"
	"runtime"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
)

func runTestValidateProof(t *testing.T, contract string, osp *ethbridgecontracts.OneStepProof) {
	t.Log("proof test contact: ", contract)

	ctx := context.Background()

	proofs, err := generateProofCases(contract)
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(proofs)
	if err != nil {
		t.Fatal(err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err := errors.New("failed to get filename")
		t.Fatal(err)
	}

	file := filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/test/proofs", filepath.Base(contract)+"-proofs.json")

	if err := ioutil.WriteFile(file, data, 0644); err != nil {
		t.Fatal(err)
	}

	for _, proof := range proofs {
		opcode := proof.Proof[len(proof.Proof)-1]
		t.Run(strconv.FormatUint(uint64(opcode), 10), func(t *testing.T) {
			var err error
			var machineData struct {
				Gas    uint64
				Fields [5][32]byte
			}

			if proof.Message != nil {
				machineData, err = osp.ExecuteStepWithMessage(
					&bind.CallOpts{Context: ctx},
					proof.Assertion.AfterInboxHash,
					proof.Assertion.FirstMessageHash,
					proof.Assertion.FirstLogHash,
					proof.Proof,
					uint8(proof.Message.Kind),
					proof.Message.ChainTime.BlockNum.AsInt(),
					proof.Message.ChainTime.Timestamp,
					proof.Message.Sender.ToEthAddress(),
					proof.Message.InboxSeqNum,
					proof.Message.Data,
				)
			} else {
				machineData, err = osp.ExecuteStep(
					&bind.CallOpts{Context: ctx},
					proof.Assertion.AfterInboxHash,
					proof.Assertion.FirstMessageHash,
					proof.Assertion.FirstLogHash,
					proof.Proof,
				)
			}
			t.Log("Opcode", opcode)
			if err != nil {
				t.Fatal("proof invalid with error", err)
			}
			if machineData.Gas != proof.Assertion.NumGas {
				t.Fatalf("wrong gas %v instead of %v", machineData.Gas, proof.Assertion.NumGas)
			}
			if machineData.Fields[0] != proof.Assertion.BeforeMachineHash {
				t.Fatal("wrong before machine")
			}
			if machineData.Fields[2] != proof.Assertion.AfterInboxHash {
				t.Fatal("wrong DidInboxInsn")
			}
			if machineData.Fields[3] != proof.Assertion.LastLogHash {
				t.Fatal("wrong log")
			}
			if machineData.Fields[4] != proof.Assertion.LastMessageHash {
				t.Fatal("wrong message")
			}
			if machineData.Fields[1] != proof.Assertion.AfterMachineHash {
				t.Fatal("wrong after machine")
			}
		})
	}
}

func TestValidateProof(t *testing.T) {
	ctx := context.Background()
	testMachines := gotest.OpCodeTestFiles()
	backend, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])
	authClient, err := ethbridge.NewEthAuthClient(ctx, client, auth)
	if err != nil {
		t.Fatal(err)
	}
	ospAddr, tx, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployOneStepProof(auth, client)
	})
	if err != nil {
		t.Fatal(err)
	}
	osp, err := ethbridgecontracts.NewOneStepProof(ospAddr, backend)
	if err != nil {
		t.Fatal(err)
	}
	client.Commit()
	if _, err := ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		tx,
		"DeployOneStepProof",
	); err != nil {
		t.Fatal(err)
	}

	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			runTestValidateProof(t, machName, osp)
		})
	}
}
