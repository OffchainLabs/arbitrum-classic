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
	"errors"
	"flag"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
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
		if err := validateRollupChain(); err != nil {
			log.Fatal(err)
		}
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
		GracePeriod:             common.TimeTicks{big.NewInt(13000 * 2)},
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
	client arbbridge.ArbClient,
	rollupAddress common.Address,
	codeFile string,
) (*rollup.ChainObserver, error) {
	ctx := context.Background()
	checkpointer := rollup.NewDummyCheckpointer(codeFile)
	chainObserver, err := rollupmanager.CreateObserver(ctx, rollupAddress, checkpointer, true, client)
	if err != nil {
		return nil, err
	}
	chainObserver.AddListener(&rollup.AnnouncerListener{})
	return chainObserver, nil
}

func validateRollupChain() error {
	// Check number of args

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	rpcEnable := validateCmd.Bool("rpc", false, "rpc")
	err := validateCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if validateCmd.NArg() != 4 {
		return errors.New("usage: rollupServer validate [--rpc] <contract.ao> <private_key.txt> <ethURL> <bridge_eth_addresses.json>")
	}

	// 2) Private key
	keyFile, err := os.Open(validateCmd.Arg(1))
	if err != nil {
		return err
	}
	byteValue, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return err
	}
	if err := keyFile.Close(); err != nil {
		return err
	}
	rawKey := strings.TrimPrefix(strings.TrimSpace(string(byteValue)), "0x")
	key, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		return fmt.Errorf("HexToECDSA private key error: %v", err)
	}

	// 3) URL
	ethURL := validateCmd.Arg(2)

	// 4) Rollup contract address
	addressString := validateCmd.Arg(3)
	address := common.HexToAddress(addressString)

	// Rollup creation
	auth := bind.NewKeyedTransactor(key)
	client, err := ethbridge.NewEthAuthClient(ethURL, auth)
	if err != nil {
		return err
	}

	chainObserver, err := setupChainObserver(client, address, validateCmd.Arg(0))
	if err != nil {
		return err
	}
	validatorListener := rollup.NewValidatorChainListener(chainObserver)
	err = validatorListener.AddStaker(client)
	if err != nil {
		return err
	}
	chainObserver.AddListener(validatorListener)

	if *rpcEnable {
		rollupvalidator.LaunchRPC(chainObserver, "1235")
	} else {
		wait := make(chan bool)
		<-wait
	}
	return nil
}
