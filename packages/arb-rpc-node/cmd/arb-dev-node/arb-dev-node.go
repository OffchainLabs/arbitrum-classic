/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
	"flag"
	"fmt"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/dev"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	accounts2 "github.com/ethereum/go-ethereum/accounts"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

var canceled = false

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()

	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-dev-node").Logger()
}

func main() {
	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running dev node")
	}
}

func startup() error {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	rpcVars := utils2.AddRPCFlags(fs)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	saveMessages := fs.String("save", "", "save messages")
	walletcount := fs.Int("walletcount", 10, "number of wallets to fund")
	walletbalance := fs.Int64("walletbalance", 100, "amount of funds in each wallet (Eth)")
	arbosPath := fs.String("arbos", "", "ArbOS version")
	mnemonic := fs.String(
		"mnemonic",
		"jar deny prosper gasp flush glass core corn alarm treat leg smart",
		"mnemonic to generate accounts from",
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		return errors.Wrap(err, "error generating temporary directory")
	}

	wallet, err := hdwallet.NewFromMnemonic(*mnemonic)
	if err != nil {
		return err
	}

	depositSize, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		return errors.New("invalid value for deposit amount")
	}
	depositSize = depositSize.Mul(depositSize, big.NewInt(*walletbalance))

	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	accounts := make([]accounts2.Account, 0)
	for i := 0; i < *walletcount; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			return err
		}
		accounts = append(accounts, account)
	}

	if *arbosPath == "" {
		arbosPathStr, err := arbos.Path()
		if err != nil {
			return err
		}
		arbosPath = &arbosPathStr
	}

	backend, db, rollupAddress, cancelDevNode, err := dev.NewDevNode(tmpDir, *arbosPath, config, common.NewAddressFromEth(accounts[0].Address), nil)
	if err != nil {
		return err
	}

	cancel := func() {
		if !canceled {
			cancelDevNode()
			if err := os.RemoveAll(tmpDir); err != nil {
				panic(err)
			}
			canceled = true
		}
	}
	defer cancel()

	for _, account := range accounts {
		deposit := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(1000000),
					GasPriceBid: big.NewInt(0),
					DestAddress: common.NewAddressFromEth(account.Address),
					Payment:     depositSize,
					Data:        nil,
				},
			}),
		}
		if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
			logger.Fatal().Err(err).Send()
		}
	}

	fmt.Println("Arbitrum Dev Chain")
	fmt.Println("")
	fmt.Println("Available Accounts")
	fmt.Println("==================")
	for i, account := range accounts {
		fmt.Printf("(%v) %v (100 ETH)\n", i, account.Address.Hex())
	}

	fmt.Println("\nPrivate Keys")
	fmt.Println("==================")
	for i, account := range accounts {
		privKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			return err
		}
		fmt.Printf("(%v) 0x%v\n", i, privKey)
	}
	fmt.Println("")

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	for _, account := range accounts {
		privKey, err := wallet.PrivateKey(account)
		if err != nil {
			return err
		}
		privateKeys = append(privateKeys, privKey)
	}

	errChan := make(chan error, 10)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if *saveMessages != "" {
			data, err := backend.ExportData()
			if err != nil {
				errChan <- errors.Wrap(err, "error exporting data from backend")
				return
			}

			if err := ioutil.WriteFile(*saveMessages, data, 777); err != nil {
				errChan <- errors.Wrap(err, "error saving exported data")
				return
			}
		}
		errChan <- nil
	}()

	plugins := make(map[string]interface{})
	plugins["evm"] = dev.NewEVM(backend)

	go func() {
		errChan <- rpc.LaunchNodeAdvanced(
			db,
			rollupAddress,
			"8547",
			"8548",
			rpcVars,
			backend,
			privateKeys,
			true,
			plugins,
		)
	}()

	err = <-errChan
	return err
}
