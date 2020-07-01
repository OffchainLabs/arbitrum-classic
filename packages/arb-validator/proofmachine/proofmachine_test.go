/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"context"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/onestepprooftester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

func setupTestValidateProof(t *testing.T) (*Connection, error) {
	ethURL := test.GetEthUrl()

	seed := time.Now().UnixNano()
	//seed := int64(1571337692091150000)
	fmt.Println("seed", seed)
	rand.Seed(seed)

	auth, err := test.SetupAuth("9af1e691e3db692cc9cad4e87b6490e099eb291e3b434a0d3f014dfd2bb747cc")
	if err != nil {
		t.Fatal(err)
	}
	ethclint, err := ethclient.Dial(ethURL)
	if err != nil {
		t.Fatal(err)
	}
	_, tx, osp, err := onestepprooftester.DeployOneStepProofTester(auth, ethclint)
	if err != nil {
		return nil, err
	}
	if _, err := ethbridge.WaitForReceiptWithResults(
		context.Background(),
		ethclint,
		auth.From,
		tx,
		"DeployOneStepProof",
	); err != nil {
		return nil, err
	}
	proofbounds := [2]uint64{0, 10000}
	return NewEthConnection(osp, proofbounds), nil
}

func runTestValidateProof(t *testing.T, contract string, ethCon *Connection) {
	t.Log("proof test contact: ", contract)
	basemach, err := loader.LoadMachineFromFile(contract, true, "cpp")

	if err != nil {
		t.Fatal(err)
	}

	mach, err := New(basemach, ethCon)
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := &protocol.TimeBounds{
		LowerBoundBlock:     common.NewTimeBlocks(big.NewInt(0)),
		UpperBoundBlock:     common.NewTimeBlocks(big.NewInt(10000)),
		LowerBoundTimestamp: big.NewInt(100),
		UpperBoundTimestamp: big.NewInt(120),
	}
	steps := uint64(100000)
	cont := true

	for cont {
		_, stepsExecuted := mach.ExecuteAssertion(steps, timeBounds, value.NewEmptyTuple(), 0)
		if mach.CurrentStatus() == machine.ErrorStop {
			t.Fatal("Machine in error state")
		}
		if stepsExecuted == 0 {
			blocked := mach.IsBlocked(common.NewTimeBlocks(big.NewInt(0)), false)
			if blocked != nil {
				cont = false
			}
		}
		if stepsExecuted != 1 {
			t.Log("Num steps = ", stepsExecuted)
		}
	}
	t.Log("called ValidateProof")

	time.Sleep(5 * time.Second)

	t.Log("done")
}

func TestValidateProof(t *testing.T) {
	testMachines := []string{
		"opcodetesttuple.mexe",
		"opcodetestlogic.mexe",
		"opcodetestmath.mexe",
		"opcodetesthash.mexe",
		"opcodetestethhash2.mexe",
		"opcodeteststack.mexe",
		"opcodetestdup.mexe",
		"opcodetestarbgas.mexe",
		"opcodetestecrecover.mexe",
		gotest.TestMachinePath(),
	}
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
