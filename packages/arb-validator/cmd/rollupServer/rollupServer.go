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
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

// Launches the rollup validator with the following command line arguments:
// 1) Compiled Arbitrum bytecode file
// 2) private key file
// 3) Global EthBridge addresses json file
// 4) ethURL
func main() {
	// Check number of args
	flag.Parse()
	if flag.NArg() != 4 {
		log.Fatalln("usage: rollupServer <contract.ao> <private_key.txt> <bridge_eth_addresses.json> <ethURL>")
	}

	// 1) Compiled Arbitrum bytecode
	mach, err := loader.LoadMachineFromFile(flag.Arg(0), true, "cpp")
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

	// 3) Global EthBridge addresses json
	jsonFile, err := os.Open(flag.Arg(2))
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

	// 5) URL
	ethURL := flag.Arg(3)

	config := structures.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             structures.TimeTicks{big.NewInt(13000 * 2)},
		MaxExecutionSteps:       200000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	// Rollup creation
	auth := bind.NewKeyedTransactor(key)
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		log.Fatal(err)
	}

	factory, err := ethbridge.NewArbFactory(common.HexToAddress(connectionInfo.ArbFactory), client)
	if err != nil {
		log.Fatal(err)
	}
	auth.Context = context.Background()
	address, err := factory.CreateRollup(
		auth,
		mach.Hash(),
		config,
		common.Address{},
	)

	server, err := rollupvalidator.NewRPCServer(auth, client, address, mach, config)
	if err != nil {
		log.Fatal(err)
	}

	// Run server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(server, "Validator"); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/", s).Methods("GET", "POST", "OPTIONS")
	//attachProfiler(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(":1235", handlers.CORS(headersOk, originsOk, methodsOk)(r))
	if err != nil {
		panic(err)
	}
}
