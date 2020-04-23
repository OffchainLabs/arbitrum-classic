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

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/txaggregator"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := utils.AddFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if fs.NArg() != 3 {
		log.Fatalf("usage: arb-tx-aggregator [--password=pass] [--gasprice==FloatInGwei] %v", utils.RollupArgsString)
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

	if err := arbbridge.WaitForNonZeroBalance(context.Background(), client, common.NewAddressFromEth(auth.From)); err != nil {
		log.Fatal(err)
	}

	rollupContract, err := client.NewRollupWatcher(rollupArgs.Address)
	if err != nil {
		log.Fatal(err)
	}
	inboxAddress, err := rollupContract.InboxAddress(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	globalInbox, err := client.NewGlobalInbox(inboxAddress)
	if err != nil {
		log.Fatal(err)
	}

	server := txaggregator.NewRPCServer(context.Background(), globalInbox, rollupArgs.Address)

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(server, "TxAggregator"); err != nil {
		log.Fatal(err)
	}

	log.Fatal(utils.LaunchRPC(s, "1237"))
}
