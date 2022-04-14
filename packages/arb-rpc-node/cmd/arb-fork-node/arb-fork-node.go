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
	"encoding/json"
	"flag"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/cmd/internal"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/dev"
)

var pprofMux *http.ServeMux

var logger = arblog.Logger.With().Str("component", "arb-fork-node").Logger()

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running node")
	}
}

func startup() error {
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	fs := flag.NewFlagSet("", flag.ContinueOnError)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	walletcount := fs.Int("walletcount", 10, "number of wallets to fund")
	walletbalance := fs.Int64("walletbalance", 100, "amount of funds in each wallet (Eth)")
	dbDir := fs.String("dbdir", "", "directory to load dev node on. Use temporary if empty")
	aggStr := fs.String("aggregator", "", "aggregator to use as the sender from this node")
	chainId64 := fs.Uint64("chainId", 68799, "chain id of chain")
	mnemonic := fs.String(
		"mnemonic",
		"jar deny prosper gasp flush glass core corn alarm treat leg smart",
		"mnemonic to generate accounts from",
	)
	fs.Bool("prettyprint", true, "pretty log output")
	persistState := fs.Bool("persist-state", false, "chain id of chain")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))); err != nil {
		return err
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	chainId := new(big.Int).SetUint64(*chainId64)

	wallet, accounts, err := internal.InitializeWallet(*mnemonic, *walletcount)
	if err != nil {
		return err
	}

	var agg common.Address
	if *aggStr != "" {
		agg = common.HexToAddress(*aggStr)
	} else {
		if len(accounts) < 2 {
			return errors.New("must have at least 2 accounts")
		}
		agg = common.NewAddressFromEth(accounts[1].Address)

	}
	type forkInfo struct {
		LastMessage int64 `json:"last_block"`
	}
	var fork forkInfo

	forkFile := filepath.Join(*dbDir, "fork.json")
	forkData, err := os.ReadFile(forkFile)
	if err == nil {
		if err := json.Unmarshal(forkData, &fork); err != nil {
			return err
		}
	} else {
		msgCount, err := dev.GetMessageCount(*dbDir)
		if err != nil {
			return err
		}
		fork.LastMessage = int64(msgCount)
		forkData, err := json.Marshal(fork)
		if err != nil {
			return err
		}
		if err := os.WriteFile(forkFile, forkData, 0644); err != nil {
			return err
		}
	}
	logger.Info().Int64("message", fork.LastMessage).Msg("Forking chain")

	backend, db, mon, cancel, txDBErrChan, err := dev.NewForkNode(
		ctx,
		*dbDir,
		chainId,
		agg,
		fork.LastMessage,
		*persistState,
	)
	if err != nil {
		return err
	}

	defer cancel()

	srv := aggregator.NewServer(backend, chainId, db)

	depositSize := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	depositSize = depositSize.Mul(depositSize, big.NewInt(*walletbalance))
	for _, account := range accounts {
		dest := common.NewAddressFromEth(account.Address)
		sender := message.L2RemapAccount(dest)
		deposit := message.RetryableTx{
			Destination:       dest,
			Value:             big.NewInt(0),
			Deposit:           depositSize,
			MaxSubmissionCost: depositSize,
			CreditBack:        dest,
			Beneficiary:       dest,
			MaxGas:            big.NewInt(0),
			GasPriceBid:       big.NewInt(0),
		}
		if _, err := backend.AddInboxMessage(ctx, deposit, sender); err != nil {
			return err
		}
	}

	if err := internal.PrintAccountInfo(wallet, accounts); err != nil {
		return err
	}

	plugins := make(map[string]interface{})
	plugins["evm"] = dev.NewEVM(backend)

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	for _, account := range accounts {
		privKey, err := wallet.PrivateKey(account)
		if err != nil {
			return err
		}
		privateKeys = append(privateKeys, privKey)
	}

	web3Server, err := web3.GenerateWeb3Server(srv, privateKeys, web3.DefaultConfig, mon.CoreConfig, plugins, nil)
	if err != nil {
		return err
	}

	{
		owner := common.NewAddressFromEth(accounts[0].Address)
		ownerAdd := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(1000000),
					GasPriceBid: big.NewInt(2000000000),
					DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
					Payment:     big.NewInt(0),
					Data:        arbos.AddChainOwnerData(owner),
				},
			}),
		}
		if _, err := backend.AddInboxMessage(ctx, ownerAdd, common.Address{}); err != nil {
			return err
		}

		setChainId := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(1000000),
					GasPriceBid: big.NewInt(2000000000),
					DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
					Payment:     big.NewInt(0),
					Data:        arbos.SetChainParameterData(arbos.ChainIDId, chainId),
				},
			}),
		}
		if _, err := backend.AddInboxMessage(ctx, setChainId, common.Address{}); err != nil {
			return err
		}
	}

	errChan := make(chan error, 1)
	go func() {
		rpcConfig := configuration.RPC{
			Addr: "0.0.0.0",
			Port: "8547",
			Path: "/",
		}
		wsConfig := configuration.WS{
			Addr: "0.0.0.0",
			Port: "8548",
			Path: "/",
		}
		errChan <- rpc.LaunchPublicServer(ctx, web3Server, rpcConfig, wsConfig)
	}()

	select {
	case err := <-txDBErrChan:
		return err
	case err := <-errChan:
		return err
	case <-cancelChan:
		return nil
	}
}
