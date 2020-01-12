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
	"fmt"
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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
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
	switch os.Args[1] {
	case "create":
		createRollupChain()
	case "validate":
		validateRollupChain()
	default:
	}
}

func createRollupChain() {
	// Check number of args
	flag.Parse()
	if flag.NArg() != 5 {
		log.Fatalln("usage: rollupServer create <contract.ao> <private_key.txt> <ethURL> <bridge_eth_addresses.json>")
	}

	// 1) Compiled Arbitrum bytecode
	mach, err := loader.LoadMachineFromFile(flag.Arg(1), true, "cpp")
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	// 2) Private key
	keyFile, err := os.Open(flag.Arg(2))
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

	// 3) URL
	ethURL := flag.Arg(3)

	// 4) Global EthBridge addresses json
	jsonFile, err := os.Open(flag.Arg(4))
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

	config := structures.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             structures.TimeTicks{big.NewInt(13000 * 2)},
		MaxExecutionSteps:       250000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	// Rollup creation
	auth := bind.NewKeyedTransactor(key)
	client, err := ethbridge.NewEthAuthClient(ethURL, auth)
	if err != nil {
		log.Fatal(err)
	}

	factory, err := client.NewArbFactory(connectionInfo.ArbFactoryAddress())
	if err != nil {
		log.Fatal(err)
	}

	address, err := factory.CreateRollup(
		context.Background(),
		mach.Hash(),
		config,
		common.Address{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
}

func setupChainObserver(
	client arbbridge.ArbAuthClient,
	rollupAddress common.Address,
	codeFile string,
) (*rollup.ChainObserver, error) {
	ctx := context.Background()
	currentTime, err := client.CurrentBlockTime(ctx)
	if err != nil {
		return nil, err
	}
	checkpointer := rollup.NewDummyCheckpointer(codeFile)
	chainObserver, err := rollup.CreateObserver(ctx, rollupAddress, checkpointer, true, currentTime, client)
	if err != nil {
		return nil, err
	}
	validatorListener := rollup.NewValidatorChainListener(chainObserver)
	err = validatorListener.AddStaker(client)
	if err != nil {
		return nil, err
	}
	chainObserver.AddListener(&rollup.AnnouncerListener{})
	chainObserver.AddListener(validatorListener)
	return chainObserver, nil
}

func validateRollupChain() {
	// Check number of args

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	rpcEnable := validateCmd.Bool("rpc", false, "rpc")
	err := validateCmd.Parse(os.Args[2:])
	if err != nil {
		log.Fatalln(err)
	}

	if validateCmd.NArg() != 4 {
		log.Fatalln("usage: rollupServer validate [--rpc] <contract.ao> <private_key.txt> <ethURL> <bridge_eth_addresses.json>")
	}

	// 2) Private key
	keyFile, err := os.Open(validateCmd.Arg(1))
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

	// 3) URL
	ethURL := validateCmd.Arg(2)

	// 4) Rollup contract address
	addressString := validateCmd.Arg(3)
	address := common.NewAddressFromEth(ethcommon.HexToAddress(addressString))

	// Rollup creation
	auth := bind.NewKeyedTransactor(key)
	client, err := ethbridge.NewEthAuthClient(ethURL, auth)
	if err != nil {
		log.Fatal(err)
	}

	chainObserver, err := setupChainObserver(client, address, validateCmd.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if *rpcEnable {
		server, err := rollupvalidator.NewRPCServer(chainObserver, 200000)
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
	} else {
		wait := make(chan bool)
		<-wait
	}
}
