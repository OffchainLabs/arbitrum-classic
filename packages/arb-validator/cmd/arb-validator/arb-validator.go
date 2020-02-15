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
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/cmdhelper"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
)

func main() {
	// Check number of args
	flag.Parse()
	switch os.Args[1] {
	case "create":
		if err := createRollupChain(); err != nil {
			log.Fatal(err)
		}
	case "validate":
		if err := cmdhelper.ValidateRollupChain("arb-validator", createManager); err != nil {
			log.Fatal(err)
		}
	default:
	}
}

func createRollupChain() error {
	createCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	walletVars := utils.AddWalletFlags(createCmd)
	tokenAddressString := createCmd.String("staketoken", "", "staketoken=TokenAddress")
	stakeAmountString := createCmd.String("stakeamount", "", "stakeamount=Amount")
	err := createCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if createCmd.NArg() != 3 {
		return fmt.Errorf("usage: arb-validator create %v <validator_folder> <ethURL> <factoryAddress> [staketoken=TokenAddress] [staketoken=TokenAddress]", utils.WalletArgsString)
	}

	validatorFolder := createCmd.Arg(0)
	ethURL := createCmd.Arg(1)
	addressString := createCmd.Arg(2)
	factoryAddress := common.HexToAddress(addressString)
	contractFile := filepath.Join(validatorFolder, cmdhelper.ContractName)

	// 1) Compiled Arbitrum bytecode
	mach, err := loader.LoadMachineFromFile(contractFile, true, "cpp")
	if err != nil {
		return errors2.Wrap(err, "loader error")
	}

	auth, err := utils.GetKeystore(validatorFolder, walletVars, createCmd)
	if err != nil {
		return err
	}

	ethclint, err := ethclient.Dial(ethURL)
	if err != nil {
		return err
	}

	// Rollup creation
	client := ethbridge.NewEthAuthClient(ethclint, auth)

	if err := arbbridge.WaitForNonZeroBalance(context.Background(), client, common.NewAddressFromEth(auth.From)); err != nil {
		return err
	}

	factory, err := client.NewArbFactory(factoryAddress)
	if err != nil {
		return err
	}

	params := rollup.DefaultChainParams()
	if *tokenAddressString != "" {
		params = params.WithStakeToken(common.HexToAddress(*tokenAddressString))
	}

	if *stakeAmountString == "" {
		stakeAmount, success := new(big.Int).SetString(*stakeAmountString, 10)
		if success {
			params = params.WithStakeRequirement(stakeAmount)
		} else {
			return errors.New("invalid stake amount: expected an integer")
		}
	}

	address, _, err := factory.CreateRollup(
		context.Background(),
		mach.Hash(),
		rollup.DefaultChainParams(),
		common.Address{},
	)
	if err != nil {
		return err
	}
	fmt.Println(address.Hex())
	return nil
}

func createManager(rollupAddress common.Address, client arbbridge.ArbAuthClient, contractFile string, dbPath string) (*rollupmanager.Manager, error) {
	return rollupmanager.CreateManager(context.Background(), rollupAddress, client, contractFile, dbPath)
}
