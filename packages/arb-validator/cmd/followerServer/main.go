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
	"crypto/ecdsa"
	jsonenc "encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type FollowerServer struct {
	follower *ethvalidator.ValidatorFollower
}

func NewFollowerServer(
	machine machine.Machine,
	key *ecdsa.PrivateKey,
	validators []common.Address,
	connectionInfo ethbridge.ArbAddresses,
	ethURL string,
	coordinatorURL string,
) *FollowerServer {
	escrowRequired := big.NewInt(10)
	config := valmessage.NewVMConfiguration(
		10,
		escrowRequired,
		common.Address{}, // Address 0 is eth
		validators,
		200000,
		common.Address{}, // Address 0 means no owner
	)

	man, err := ethvalidator.NewValidatorFollower(
		"Bob",
		machine,
		key,
		config,
		false,
		math.MaxInt32, // maxCallSteps
		connectionInfo,
		ethURL,
		coordinatorURL,
		math.MaxInt32, // maxUnanSteps
	)
	if err != nil {
		log.Fatalf("Failed to create follower %v\n", err)
	}

	if err = man.Run(); err != nil {
		log.Fatal(err)
	}

	receiptChan, errChan := man.DepositFunds(context.Background(), escrowRequired)
	select {
	case receipt := <-receiptChan:
		if receipt.Status == 0 {
			log.Fatalln("Follower could not deposit funds")
		}
	case err := <-errChan:
		log.Fatal(err)
	}

	return &FollowerServer{man}
}

func (m *FollowerServer) SendMessage(r *http.Request, args *bool, reply *bool) error {
	return nil
}

// Launches one validator with the following command line arguments:
// 1) compiled Arbitrum bytecode file
// 2) private key
// 3) public key addresses (newline separated)
// 4) global EthBridge addresses in json
// 5) ethURL
// 6) coordinatorURL
func main() {
	vmType := flag.String("avm", "cpp", "Select the AVM implementation")
	flag.Parse()

	// Check number of args
	if len(flag.Args()) != 6 {
		log.Fatalln("usage: followerServer <contract.ao> <private_key.txt> <validator_addresses.txt> <bridge_eth_addresses.json> <ethURL> <coordinatorURL>")
	}

	// 1) Compiled Arbitrum bytecode
	m, err := loader.LoadMachineFromFile(flag.Arg(0), true, *vmType)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	// 2) Private key
	keyFile, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err := ioutil.ReadAll(keyFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := keyFile.Close(); err != nil {
		log.Fatalln(err)
	}
	rawKey := strings.TrimPrefix(strings.TrimSpace(string(byteValue)), "0x")
	key, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		log.Fatal("HexToECDSA private key error: ", err)
	}

	// 3) All public key addresses
	addrFile, err := os.Open(flag.Arg(2))
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err = ioutil.ReadAll(addrFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := addrFile.Close(); err != nil {
		log.Fatalln(err)
	}
	validatorHexAddrs := strings.Split(
		strings.TrimPrefix(strings.TrimSpace(string(byteValue)), "0x"), "\n")
	validators := make([]common.Address, len(validatorHexAddrs))
	for i, v := range validatorHexAddrs {
		validators[i] = common.HexToAddress(v)
	}

	// 4) Global EthBridge addresses json
	jsonFile, err := os.Open(flag.Arg(3))
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		log.Fatalln(err)
	}
	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		log.Fatalln(err)
	}

	// 5) ethURL 6) coordinatorURL
	ethURL := flag.Arg(4)
	coordinatorURL := flag.Arg(5)

	// Validator creation
	NewFollowerServer(
		m,
		key,
		validators,
		connectionInfo,
		ethURL,
		coordinatorURL,
	)

	blockChan := make(chan struct{})
	<-blockChan
}
