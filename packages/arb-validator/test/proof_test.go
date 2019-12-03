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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/proofmachine"
)

func setupTestValidateProof(t *testing.T) (*proofmachine.Connection, error) {
	var connectionInfo ethbridge.ArbAddresses

	bridge_eth_addresses := "bridge_eth_addresses.json"
	ethURL := "ws://127.0.0.1:7546"

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
	key1, err := crypto.HexToECDSA("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	if err != nil {
		t.Fatal(err)
	}
	proofbounds := [2]uint64{0, 10000}
	return proofmachine.NewEthConnection(common.HexToAddress(connectionInfo.OneStepProof), key1, ethURL, proofbounds)
}

func runTestValidateProof(t *testing.T, contract string, ethCon *proofmachine.Connection) {
	basemach, err := loader.LoadMachineFromFile(contract, true, "test")

	if err != nil {
		t.Fatal(err)
	}

	mach, err := proofmachine.New(basemach, ethCon)
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := protocol.NewTimeBounds(0, 10000)
	steps := int32(100000)
	cont := true

	for cont {
		a := mach.ExecuteAssertion(steps, timeBounds)
		lastReason := mach.LastBlockReason()
		if lastReason != nil {
			if lastReason.IsBlocked(mach, 0) && lastReason.Equals(machine.ErrorBlocked{}) {
				t.Fatal("Machine in error state")
				break
			}
		}
		if a.NumSteps == 0 {
			if lastReason.IsBlocked(mach, 0) && !lastReason.Equals(machine.BreakpointBlocked{}) {
				cont = false
			}
			fmt.Println(" machine halted ")
			//break
		}
		if a.NumSteps != 1 {
			t.Log("Num steps = ", a.NumSteps)
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
