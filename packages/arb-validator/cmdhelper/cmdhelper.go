/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package cmdhelper

import (
	"context"
	"flag"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"os"
	"path/filepath"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
)

var ContractName = "contract.mexe"

// ValidateRollupChain creates a validator given the managerCreationFunc.
// This allows for the abstraction of the manager setup away from command line
// parsing and initialization of common structures and behavior
func ValidateRollupChain(
	execName string,
	managerCreationFunc func(
		ctx context.Context,
		rollupAddress common.Address,
		client arbbridge.ArbClient,
		contractFile string, dbPath string,
	) (*rollupmanager.Manager, error),
) error {
	ctx := context.Background()
	// Check number of args

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	walletVars := utils.AddWalletFlags(validateCmd)
	blocktime := validateCmd.Int64(
		"blocktime",
		2,
		"blocktime=NumSeconds",
	)
	err := validateCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if validateCmd.NArg() != 3 {
		return fmt.Errorf(
			"usage: %v validate %v [--blocktime=NumSeconds] %v",
			execName,
			utils.WalletArgsString,
			utils.RollupArgsString,
		)
	}

	common.SetDurationPerBlock(time.Duration(*blocktime) * time.Second)

	rollupArgs := utils.ParseRollupCommand(validateCmd, 0)

	auth, err := utils.GetKeystore(
		rollupArgs.ValidatorFolder,
		walletVars,
		validateCmd,
	)
	if err != nil {
		return err
	}

	// Rollup creation
	ethclint, err := ethutils.NewRPCEthClient(rollupArgs.EthURL)
	if err != nil {
		return err
	}
	client, err := ethbridge.NewEthAuthClient(ctx, ethclint, auth)
	if err != nil {
		return err
	}

	rollup, err := client.NewRollup(rollupArgs.Address)
	if err != nil {
		return err
	}

	params, err := rollup.GetParams(ctx)
	if err != nil {
		return err
	}

	if err := arbbridge.WaitForBalance(ctx, client, params.StakeToken, common.NewAddressFromEth(auth.From)); err != nil {
		return err
	}

	validatorListener := chainlistener.NewValidatorChainListener(
		ctx,
		rollupArgs.Address,
		rollup,
	)
	err = validatorListener.AddStaker(client)
	if err != nil {
		return err
	}

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, ContractName)
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	manager, err := managerCreationFunc(
		ctx,
		rollupArgs.Address,
		client,
		contractFile,
		dbPath,
	)

	if err != nil {
		return err
	}
	manager.AddListener(ctx, &chainlistener.AnnouncerListener{})
	manager.AddListener(ctx, validatorListener)

	wait := make(chan bool)
	<-wait
	return nil
}

// ValidateRollupChain creates a validator given the managerCreationFunc.
// This allows for the abstraction of the manager setup away from command line
// parsing and initialization of common structures and behavior
func ObserveRollupChain(
	execName string,
	managerCreationFunc func(
		ctx context.Context,
		rollupAddress common.Address,
		client arbbridge.ArbClient,
		contractFile string, dbPath string,
	) (*rollupmanager.Manager, error),
) error {
	ctx := context.Background()
	// Check number of args
	validateCmd := flag.NewFlagSet("observe", flag.ExitOnError)
	quietFlag := validateCmd.Bool(
		"q",
		false,
		"quiet validator output",
	)
	err := validateCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if validateCmd.NArg() != 3 {
		return fmt.Errorf(
			"usage: %v validate %v",
			execName,
			utils.RollupArgsString,
		)
	}

	rollupArgs := utils.ParseRollupCommand(validateCmd, 0)

	// Rollup creation
	ethclint, err := ethutils.NewRPCEthClient(rollupArgs.EthURL)
	if err != nil {
		return err
	}
	client := ethbridge.NewEthClient(ethclint)

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, ContractName)
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	manager, err := managerCreationFunc(
		ctx,
		rollupArgs.Address,
		client,
		contractFile,
		dbPath,
	)

	if err != nil {
		return err
	}
	if !*quietFlag {
		manager.AddListener(ctx, &chainlistener.AnnouncerListener{})
	}

	wait := make(chan bool)
	<-wait
	return nil
}
