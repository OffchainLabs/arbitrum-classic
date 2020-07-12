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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/onestepprooftester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

func setupTestValidateProof(t *testing.T) (*onestepprooftester.OneStepProofTester, error) {
	client, auths := test.SimulatedBackend()
	auth := auths[0]
	_, tx, osp, err := onestepprooftester.DeployOneStepProofTester(auth, client)
	if err != nil {
		return nil, err
	}
	client.Commit()
	if _, err := ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		tx,
		"DeployOneStepProof",
	); err != nil {
		return nil, err
	}
	return osp, nil
}

func runTestValidateProof(t *testing.T, contract string, osp *onestepprooftester.OneStepProofTester) {
	t.Log("proof test contact: ", contract)
	mach, err := loader.LoadMachineFromFile(contract, true, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	maxSteps := uint64(100000)
	inbox := value.NewEmptyTuple()

	for i := uint64(0); i < maxSteps; i++ {
		proof, err := mach.MarshalForProof()
		if err != nil {
			t.Fatal(err)
		}
		beforeHash := mach.Hash()
		beforeMach := mach.Clone()
		a, ranSteps := mach.ExecuteAssertion(1, inbox, 0)
		if ranSteps == 0 {
			break
		}
		if ranSteps != 1 {
			t.Fatal("Executed incorrect step count", ranSteps)
		}
		if mach.CurrentStatus() == machine.ErrorStop {
			beforeMach.PrintState()
			mach.PrintState()
			t.Fatal("machine stopped in error state")
		}

		//precond := valprotocol.NewPrecondition(beforeHash, timeBounds, inbox)
		stub := valprotocol.NewExecutionAssertionStubFromAssertion(a)
		hashPreImage := inbox.GetPreImage()
		res, err := osp.ValidateProof(
			&bind.CallOpts{Context: context.Background()},
			beforeHash,
			hashPreImage.GetInnerHash(),
			big.NewInt(hashPreImage.Size()),
			stub.AfterHash,
			stub.DidInboxInsn,
			stub.FirstMessageHash,
			stub.LastMessageHash,
			stub.FirstLogHash,
			stub.LastLogHash,
			stub.NumGas,
			proof,
		)
		if err != nil {
			beforeMach.PrintState()
			mach.PrintState()
			t.Fatal("Proof invalid with error", err)
		} else if res.Cmp(big.NewInt(0)) != 0 {
			mach.PrintState()
			t.Fatal("Proof invalid")
		}

		if a.DidInboxInsn {
			inbox = value.NewEmptyTuple()
		}
	}
}

func TestValidateProof(t *testing.T) {
	testMachines := gotest.OpCodeTestFiles()
	ethCon, err := setupTestValidateProof(t)
	if err != nil {
		t.Fatal(err)
	}
	for _, machName := range testMachines {
		machName := machName // capture range variable
		t.Run(machName, func(t *testing.T) {
			//t.Parallel()
			runTestValidateProof(t, machName, ethCon)
		})
	}
}
