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
	"context"
	jsonenc "encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	brand "math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/channel"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

func TestChallenge(t *testing.T) {

	bridge_eth_addresses := "bridge_eth_addresses.json"
	contract := "contract.ao"
	ethURL := "ws://127.0.0.1:7546"

	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
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

	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Fatal(err)
	}

	key1, err := crypto.HexToECDSA("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	if err != nil {
		t.Fatal(err)
	}
	key2, err := crypto.HexToECDSA("6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1")
	if err != nil {
		t.Fatal(err)
	}

	basemachine, err := loader.LoadMachineFromFile(contract, true, "test")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}
	machine := basemachine

	auth1 := bind.NewKeyedTransactor(key1)
	auth2 := bind.NewKeyedTransactor(key2)

	validators := []common.Address{auth1.From, auth2.From}
	escrowRequired := big.NewInt(10)
	config := valmessage.NewVMConfiguration(
		2,
		escrowRequired,
		common.Address{}, // Address 0 is eth
		validators,
		200000,
		common.Address{}, // Address 0 means no owner
	)

	val1, err := ethvalidator.NewValidator(
		key1,
		connectionInfo,
		ethURL,
	)
	if err != nil {
		t.Fatal(err)
	}

	val2, err := ethvalidator.NewValidator(
		key2,
		connectionInfo,
		ethURL,
	)
	if err != nil {
		t.Fatal(err)
	}

	address, err := val1.CreateChannel(context.Background(), config, machine.Hash())
	if err != nil {
		t.Fatal(err)
	}

	coordinator, err := channel.NewCoordinator(
		"Alice",
		val1,
		address,
		machine.Clone(),
		config,
		false,
		math.MaxInt32, // maxCallSteps,
		math.MaxInt32, // maxUnanSteps
		time.Second,
	)

	if err != nil {
		t.Fatal(err)
	}

	coordinator.StartServer(context.Background())

	time.Sleep(1 * time.Second)

	challenger, err := channel.NewValidatorFollower(
		"Bob",
		val2,
		machine.Clone(),
		config,
		true,
		math.MaxInt32, // maxCallSteps,
		math.MaxInt32, // maxUnanSteps
		"wss://127.0.0.1:1236/ws",
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := coordinator.Run(context.Background()); err != nil {
		t.Fatal(err)
	}

	if err := challenger.Run(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Log("Everyone is running")
	coordinatorMsgMonitorChan := coordinator.Val.MessageMonChan
	coordinatorErrMonitorChan := coordinator.Val.ErrorMonChan
	challengerMsgMonitorChan := challenger.Validator.MessageMonChan
	challengerErrMonitorChan := challenger.Validator.ErrorMonChan
	time.Sleep(2 * time.Second)

	challenger.IgnoreCoordinator()

	dataBytes, _ := hexutil.Decode("0x2ddec39b0000000000000000000000000000000000000000000000000000000000000028")
	data, _ := evm.BytesToSizedByteArray(dataBytes)
	addressInt, _ := new(big.Int).SetString("784030224795475933405737832577560929931042096197", 10)
	seq := value.NewInt64Value(100)

	tup, _ := value.NewTupleFromSlice([]value.Value{
		data,
		value.NewIntValue(addressInt),
		seq,
	})
	//coordinator.Val.Validator.
	receipt, err := coordinator.Val.SendEthMessage(
		context.Background(),
		tup,
		big.NewInt(0),
	)
	if err != nil {
		t.Fatal("Send error", err)
	}
	if receipt.Status == 0 {
		t.Fatal("Follower could not send message")
	}
	for {
		select {
		case message := <-challengerMsgMonitorChan:
			if message == bridge.ProofAccepted {
				t.Log("***************************")
				t.Log("Challenger received ProofAccepted message = ", message)
				t.Log("***************************")
			}
		case message := <-coordinatorMsgMonitorChan:
			if message == bridge.ProofAccepted {
				t.Log("***************************")
				t.Log("Coordinator received ProofAccepted message = ", message)
				t.Log("***************************")
				return
			}
		case message := <-coordinatorErrMonitorChan:
			if !message.Recoverable {
				t.Error("coordinator unexpected unrecoverable error")
				t.Log("***************************")
				t.Log("unrecoverable error exiting")
				t.Log(message.Message)
				t.Log(message.Err)
				t.Log("***************************")
				return
			} else {
				t.Log("****************************")
				t.Log("recoverable error continuing")
				t.Log(message)
				t.Log("****************************")
			}
		case message := <-challengerErrMonitorChan:
			if !message.Recoverable {
				t.Error("challenger unexpected unrecoverable error")
				t.Log("***************************")
				t.Log("unrecoverable challenger error exiting")
				t.Log(message.Message)
				t.Log(message.Err)
				t.Log("***************************")
				return
			} else {
				t.Error("challenger unexpected recoverable error")
				t.Log("****************************")
				t.Log("recoverable challenger error continuing")
				t.Log(message.Message)
				t.Log("****************************")
			}
		case <-time.After(100 * time.Second):
			t.Error("Never received proof accepted message")
			fmt.Println("test complete")
			return
		}
	}
	//time.Sleep(60 * time.Second)
}
