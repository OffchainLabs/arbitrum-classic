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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/test"
)

func setupTestValidateProof(t *testing.T) (*Connection, error) {
	var connectionInfo ethbridge.ArbAddresses

	bridge_eth_addresses := "../bridge_eth_addresses.json"
	ethURL := test.GetEthUrl()

	seed := time.Now().UnixNano()
	//seed := int64(1571337692091150000)
	fmt.Println("seed", seed)
	rand.Seed(seed)
	jsonFile, err := os.Open(bridge_eth_addresses)
	if err != nil {
		t.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Fatal(err)
	}
	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		t.Fatal(err)
	}
	proofbounds := [2]uint32{0, 10000}
	return NewEthConnection(common.HexToAddress(connectionInfo.OneStepProof), key1, ethURL, proofbounds)
}

func runTestValidateProof(t *testing.T, contract string, ethCon *Connection) {
	basemach, err := loader.LoadMachineFromFile(contract, true, "test")

	if err != nil {
		t.Fatal(err)
	}

	mach, err := New(basemach, ethCon)
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := &protocol.TimeBoundsBlocks{protocol.NewTimeBlocks(big.NewInt(0)), protocol.NewTimeBlocks(big.NewInt(10000))}
	steps := uint32(100000)
	cont := true

	for cont {
		_, stepsExecuted := mach.ExecuteAssertion(steps, timeBounds, value.NewEmptyTuple())
		lastReason := mach.LastBlockReason()
		if lastReason != nil {
			if lastReason.IsBlocked(mach, protocol.NewTimeBlocks(big.NewInt(0)), false) && lastReason.Equals(machine.ErrorBlocked{}) {
				t.Fatal("Machine in error state")
				break
			}
		}
		if stepsExecuted == 0 {
			if lastReason.IsBlocked(mach, protocol.NewTimeBlocks(big.NewInt(0)), false) && !lastReason.Equals(machine.BreakpointBlocked{}) {
				cont = false
			}
			fmt.Println(" machine halted ")
			//break
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
		"opcodetestmath.ao",
		"opcodetestlogic.ao",
		"opcodetesthash.ao",
		"opcodeteststack.ao",
		"opcodetestdup.ao",
		"opcodetesttuple.ao",
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
