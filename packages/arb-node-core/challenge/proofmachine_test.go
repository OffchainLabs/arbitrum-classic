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

package challenge

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
)

func runTestValidateProof(t *testing.T, contract string, osp *ethbridgetestcontracts.OneStepProof, osp2 *ethbridgetestcontracts.OneStepProof2) {
	t.Log("proof test contact: ", contract)
	ctx := context.Background()

	proofs, err := GenerateProofCases(contract, 100000)
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
		opcode := proof.Proof[0]
		t.Run(strconv.FormatUint(uint64(opcode), 10), func(t *testing.T) {
			var err error
			var machineData struct {
				Gas    uint64
				Fields [5][32]byte
			}
			machineFields := [3][32]byte{
				proof.BeforeCut.InboxDelta,
				proof.BeforeCut.SendAcc,
				proof.BeforeCut.LogAcc,
			}
			if len(proof.BufferProof) == 0 {
				machineData, err = osp.ExecuteStep(
					&bind.CallOpts{Context: ctx},
					machineFields,
					proof.Proof,
				)
			} else {
				machineData, err = osp2.ExecuteStep(
					&bind.CallOpts{Context: ctx},
					machineFields,
					proof.Proof,
					proof.BufferProof,
				)
			}
			test.FailIfError(t, err)

			t.Log("Opcode", opcode)
			if err != nil {
				t.Fatal("proof invalid with error", err)
			}
			correctGasUsed := proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
			if machineData.Gas != correctGasUsed {
				t.Fatalf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
			}
			if machineData.Fields[0] != proof.BeforeCut.MachineState {
				t.Fatal("wrong before machine")
			}
			if machineData.Fields[2] != proof.AfterCut.InboxDelta {
				t.Fatal("wrong DidInboxInsn")
			}
			if machineData.Fields[3] != proof.AfterCut.SendAcc {
				t.Fatal("wrong log")
			}
			if machineData.Fields[4] != proof.AfterCut.LogAcc {
				t.Fatal("wrong message")
			}
			if machineData.Fields[1] != proof.AfterCut.MachineState {
				t.Fatal("wrong after machine")
			}
		})
	}
}

func TestValidateProof(t *testing.T) {
	testMachines := gotest.OpCodeTestFiles()
	backend, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])

	_, _, osp, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	_, _, osp2, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			runTestValidateProof(t, machName, osp, osp2)
		})
	}
}
