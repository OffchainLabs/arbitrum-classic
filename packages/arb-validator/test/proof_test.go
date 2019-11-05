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
	jsonenc "encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/proofmachine"
	"io/ioutil"
	"math/big"
	brand "math/rand"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestValidateProof(t *testing.T) {
	//var mach machine.Machine
	var connectionInfo ethbridge.ArbAddresses

	bridge_eth_addresses := "bridge_eth_addresses.json"
	contract := "contract.ao"
	ethURL := "ws://127.0.0.1:7545"

	//seed := time.Now().UnixNano()
	seed := int64(1571337692091150000)
	fmt.Println("seed", seed)
	brand.Seed(seed)
	jsonFile, err := os.Open(bridge_eth_addresses)
	if err != nil {
		t.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		t.Fatal(err)
	}

	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Fatal(err)
	}

	basemach, err := loader.LoadMachineFromFile(contract, true, "test")
	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		t.Fatal(err)
	}

	mach, err := proofmachine.New(contract, basemach, true, common.HexToAddress(connectionInfo.OneStepProof), key1, ethURL)
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	var timeBounds [2]uint64

	keyAddr := crypto.PubkeyToAddress(key1.PublicKey)

	dataBytes, _ := hexutil.Decode("0x2ddec39b0000000000000000000000000000000000000000000000000000000000000028")
	data, _ := evm.BytesToSizedByteArray(dataBytes)

	msg := protocol.NewSimpleMessage(data, [21]byte{}, big.NewInt(0), keyAddr)
	callingMessage := protocol.Message{
		Data:        data,
		TokenType:   msg.TokenType,
		Currency:    msg.Currency,
		Destination: msg.Destination,
	}
	mach.SendOffchainMessages([]protocol.Message{callingMessage})

	stepIncrease := int32(1)
	maxSteps := int32(1000)
	for i := int32(0); i < maxSteps; i += stepIncrease {
		timeBounds[0] = uint64(i)
		timeBounds[1] = uint64(i + stepIncrease)
		steps := int32(stepIncrease)

		a := mach.ExecuteAssertion(steps, timeBounds)
		if a.NumSteps == 0 {
			fmt.Println(" machine halted ")
			break
		}
		if a.NumSteps != 1 {
			t.Log("Num steps = ", a.NumSteps)
		}
		fmt.Println("executed ", i, " steps")

	}

	t.Log("called ValidateProof")
	time.Sleep(5 * time.Second)
	t.Log("done")
}
