/*
 * Copyright 2020, Offchain Labs, Inc.
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
	"flag"

	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := utils.AddWalletFlags(fs)
	rpcVars := utils2.AddRPCFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if fs.NArg() != 3 {
		log.Fatalf(
			"usage: arb-tx-aggregator %v %v",
			utils.WalletArgsString,
			utils.RollupArgsString,
		)
	}

	rollupArgs := utils.ParseRollupCommand(fs, 0)

	auth, err := utils.GetKeystore(rollupArgs.ValidatorFolder, walletArgs, fs)
	if err != nil {
		log.Fatal(err)
	}

	ethclint, err := ethclient.Dial(rollupArgs.EthURL)
	if err != nil {
		log.Fatal(err)
	}
	client := ethbridge.NewEthAuthClient(ethclint, auth)

	if err := arbbridge.WaitForNonZeroBalance(
		context.Background(),
		client,
		common.NewAddressFromEth(auth.From),
	); err != nil {
		log.Fatal(err)
	}

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, "contract.mexe")
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	if err := aggregator.LaunchAggregator(
		context.Background(),
		client,
		rollupArgs.Address,
		contractFile,
		dbPath,
		"1235",
		rpcVars,
	); err != nil {
		log.Fatal(err)
	}
}
