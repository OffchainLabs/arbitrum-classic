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
	"encoding/json"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/proofmachine"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

func generateProofCases(t *testing.T, arbCore *monitor.Monitor) ([]*proofmachine.ProofData, []string) {
	var cursors []core.ExecutionCursor
	cursor, err := arbCore.Core.GetExecutionCursor(big.NewInt(0))
	test.FailIfError(t, err)
	cursors = append(cursors, cursor.Clone())
	for {
		err = arbCore.Core.AdvanceExecutionCursor(cursor, big.NewInt(1), true)
		test.FailIfError(t, err)
		if cursor.TotalGasConsumed().Cmp(cursors[len(cursors)-1].TotalGasConsumed()) == 0 {
			break
		}
		cursors = append(cursors, cursor.Clone())
	}

	proofs := make([]*proofmachine.ProofData, 0)
	machineStates := make([]string, 0)

	for i := 0; i < len(cursors)-1; i++ {
		mach, err := arbCore.Core.TakeMachine(cursors[i].Clone())
		test.FailIfError(t, err)
		machineStates = append(machineStates, mach.String())

		proof, bproof, err := mach.MarshalForProof()
		test.FailIfError(t, err)

		beforeState, err := core.NewExecutionState(cursors[i])
		test.FailIfError(t, err)
		afterState, err := core.NewExecutionState(cursors[i+1])
		test.FailIfError(t, err)
		proofs = append(proofs, &proofmachine.ProofData{
			Assertion: &core.Assertion{
				Before: beforeState,
				After:  afterState,
			},
			Proof:       proof,
			BufferProof: bproof,
		})
	}
	mach, err := arbCore.Core.TakeMachine(cursors[len(cursors)-1].Clone())
	test.FailIfError(t, err)
	machineStates = append(machineStates, mach.String())
	return proofs, machineStates
}

func TestValidateProof(t *testing.T) {
	testMachines, err := gotest.OpCodeTestFiles()
	test.FailIfError(t, err)
	backend, auths := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	proofChecker, err := proofmachine.NewProofChecker(auth, client)
	test.FailIfError(t, err)
	t.Log(testMachines)

	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			arbCore, cancel := monitor.PrepareArbCoreWithMexe(t, machName)
			proofs, _ := generateProofCases(t, arbCore)
			cancel()

			data, err := json.Marshal(proofs)
			test.FailIfError(t, err)

			_, filename, _, ok := runtime.Caller(0)
			if !ok {
				t.Fatal("failed to get filename")
			}

			file := filepath.Join(filepath.Dir(filename), "../../arb-bridge-eth/test/proofs", filepath.Base(machName)+"-proofs.json")
			err = ioutil.WriteFile(file, data, 0644)
			test.FailIfError(t, err)

			for _, proof := range proofs {
				errors := proofChecker.CheckProof(proof)
				if len(errors) > 0 {
					t.Logf("error checking proof for opcode 0x%x", proof.Proof[0])
				}
				for _, err := range proofChecker.CheckProof(proof) {
					t.Error(err)
				}
			}
		})
	}
}
