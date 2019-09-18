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
	"log"
	"math"
	"math/big"
	brand "math/rand"
	"os"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/channel"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

func main() {
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	brand.Seed(seed)
	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		log.Fatalln(err)
	}

	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		log.Fatalln(err)
	}

	machine, err := loader.LoadMachineFromFile(os.Args[2], true, "go")
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		log.Fatal(err)
	}
	key2, err := crypto.HexToECDSA("979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76")
	if err != nil {
		log.Fatal(err)
	}

	auth1 := bind.NewKeyedTransactor(key1)
	auth2 := bind.NewKeyedTransactor(key2)

	validators := []common.Address{auth1.From, auth2.From}
	escrowRequired := big.NewInt(10)
	config := valmessage.NewVMConfiguration(
		10,
		escrowRequired,
		common.Address{}, // Address 0 is eth
		validators,
		200000,
		common.Address{}, // Address 0 means no owner
	)

	ethURL := os.Args[3]

	val1, err := ethvalidator.NewValidator(
		key1,
		connectionInfo,
		ethURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	val2, err := ethvalidator.NewValidator(
		key2,
		connectionInfo,
		ethURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	address, err := val1.LaunchChannel(context.Background(), config, machine.Hash())
	if err != nil {
		log.Fatal(err)
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
	)

	if err != nil {
		log.Fatal(err)
	}

	coordinator.StartServer(context.Background())

	challenger, err := channel.NewValidatorFollower(
		"Bob",
		val2,
		machine.Clone(),
		config,
		false,
		math.MaxInt32, // maxCallSteps,
		math.MaxInt32, // maxUnanSteps
		"wss://127.0.0.1:1236/ws",
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := coordinator.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := challenger.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

	dataBytes, _ := hexutil.Decode("0x2ddec39b0000000000000000000000000000000000000000000000000000000000000028")
	data, _ := evm.BytesToSizedByteArray(dataBytes)
	addressInt, _ := new(big.Int).SetString("784030224795475933405737832577560929931042096197", 10)
	seq := value.NewInt64Value(100)

	tup, _ := value.NewTupleFromSlice([]value.Value{
		data,
		value.NewIntValue(addressInt),
		seq,
	})

	receipt, err := coordinator.Val.SendEthMessage(
		context.Background(),
		tup,
		big.NewInt(0),
	)
	if err != nil {
		log.Fatalln("Send error", err)
	}
	if receipt.Status == 0 {
		log.Fatalln("Follower could not deposit funds")
	}

	time.Sleep(2000 * time.Millisecond)
	successChan, errChan := coordinator.InitiateUnanimousAssertion(true)
	select {
	case result := <-successChan:
		fmt.Println("ChallengeTest: Unanimous assertion successful", result)
	case err := <-errChan:
		panic(fmt.Sprintf("Error Running unanimous assertion: %v", err))
	}

	successChan = coordinator.InitiateDisputableAssertion()
	result := <-successChan
	fmt.Println("Result", result)

	time.Sleep(10 * time.Second)
}
