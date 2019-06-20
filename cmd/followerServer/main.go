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
	"crypto/ecdsa"
	jsonenc "encoding/json"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/offchainlabs/arb-avm/loader"
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-validator/valmessage"

	"github.com/offchainlabs/arb-validator/ethvalidator"
)

type FollowerServer struct {
	follower *ethvalidator.ValidatorFollower
}

func NewFollowerServer(
	machine *vm.Machine,
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

	man, err := ethvalidator.NewValidatorFollower("Bob", machine, key, config, false, connectionInfo, ethURL, coordinatorURL)
	if err != nil {
		log.Fatalf("Failed to create follower %v\n", err)
	}

	_, err = man.DepositEth(escrowRequired)
	if err != nil {
		log.Fatal(err)
	}

	err = man.Run()
	if err != nil {
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
	// Check number of args
	if len(os.Args)-1 != 6 {
		log.Fatalln("Expected six arguments")
	}

	// 1) Compiled Arbitrum bytecode
	machine, err := loader.LoadMachineFromFile(os.Args[1], true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	// 2) Private key
	keyFile, err := os.Open(os.Args[2])
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
	addrFile, err := os.Open(os.Args[3])
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
	jsonFile, err := os.Open(os.Args[4])
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
	ethURL := os.Args[5]
	coordinatorURL := os.Args[6]

	// Validator creation
	rpcInterface := NewFollowerServer(machine, key, validators, connectionInfo, ethURL, coordinatorURL)

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(rpcInterface, "Validator"); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/", s).Methods("GET", "POST", "OPTIONS")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// port := fmt.Sprintf(":%d", 1237 + index)
	err = http.ListenAndServe(":1237", handlers.CORS(headersOk, originsOk, methodsOk)(r))
	if err != nil {
		panic(err)
	}
}
