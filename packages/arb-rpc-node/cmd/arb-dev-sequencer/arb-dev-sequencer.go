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
	"flag"
	"fmt"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/dev"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

var logger zerolog.Logger

var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-node").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running node")
	}
}

func startup() error {
	defer logger.Log().Msg("Cleanly shutting down node")
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	delayedMessagesTargetDelay := fs.Int64("delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	createBatchBlockInterval := fs.Int64("create-batch-block-interval", 1, "block interval at which to create new batches")
	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)
	privKeyString := fs.String("privkey", "979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76", "funded private key")
	//fundedAccount := fs.String("account", "0x9a6C04fBf4108E2c1a1306534A126381F99644cf", "account to fund")
	chainId64 := fs.Uint64("chainId", 68799, "chain id of chain")
	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if fs.NArg() != 3 {
		fmt.Println("usage: arb-dev-sequencer <ethURL> <rollup_creator_address> <bridge_utils_adddress>")
		return errors.New("invalid arguments")
	}

	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel); err != nil {
		return err
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	ethURL := fs.Arg(0)
	rollupCreatorAddresssString := fs.Arg(1)
	bridgeUtilsAddressString := fs.Arg(2)
	rollupCreator := common.HexToAddress(rollupCreatorAddresssString)
	bridgeUtilsAddress := common.HexToAddress(bridgeUtilsAddressString)

	ethclint, err := ethutils.NewRPCEthClient(ethURL)
	if err != nil {
		return errors.Wrap(err, "error running NewRPcEthClient")
	}

	arbosPath, err := arbos.Path()
	if err != nil {
		return err
	}
	initialMachine, err := cmachine.New(arbosPath)
	if err != nil {
		return err
	}

	creator, err := ethbridgecontracts.NewRollupCreator(rollupCreator.ToEthAddress(), ethclint)
	if err != nil {
		return errors.Wrap(err, "error getting chain creator")
	}

	l1ChainId, err := ethclint.ChainID(ctx)
	if err != nil {
		return err
	}

	privKey, err := crypto.HexToECDSA(*privKeyString)
	if err != nil {
		return err
	}
	deployer, err := bind.NewKeyedTransactorWithChainID(privKey, l1ChainId)
	if err != nil {
		return err
	}

	seqPrivKey, err := crypto.GenerateKey()
	if err != nil {
		return errors.Wrap(err, "error generating key")
	}
	seqAuth, err := bind.NewKeyedTransactorWithChainID(seqPrivKey, l1ChainId)
	if err != nil {
		return err
	}

	ownerPrivKey, err := crypto.GenerateKey()
	l1OwnerAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivKey, l1ChainId)
	if err != nil {
		return err
	}

	owner := l1OwnerAuth.From
	sequencer := crypto.PubkeyToAddress(seqPrivKey.PublicKey)

	tx, err := creator.CreateRollup(
		deployer,
		initialMachine.Hash(),
		big.NewInt(5),
		big.NewInt(0),
		big.NewInt(2000000000000),
		big.NewInt(10),
		ethcommon.Address{},
		owner,
		sequencer,
		big.NewInt(15),
		big.NewInt(900),
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "error creating rollup")
	}
	receipt, err := ethbridge.WaitForReceiptWithResults(ctx, ethclint, deployer.From, tx, "CreateRollup")
	if err != nil {
		return errors.Wrap(err, "error getting transaction receipt")
	}
	createdEvent, err := creator.ParseRollupCreated(*receipt.Logs[len(receipt.Logs)-1])
	if err != nil {
		return err
	}
	rollupAddress := common.NewAddressFromEth(createdEvent.RollupAddress)
	//inboxAddress := createdEvent.InboxAddress

	l2ChainId := new(big.Int).SetUint64(*chainId64)
	l2OwnerAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivKey, l2ChainId)
	if err != nil {
		return err
	}

	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")
	logger.Info().Hex("chainaddress", rollupAddress.Bytes()).Str("chainid", l2ChainId.String()).Msg("Launching arbitrum node")

	dbPath, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		return errors.Wrap(err, "error generating temporary directory")
	}

	defer func() {
		if err := os.RemoveAll(dbPath); err != nil {
			panic(err)
		}
	}()

	mon, err := monitor.NewMonitor(dbPath, arbosPath)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	dummySequencerFeed := make(chan broadcaster.BroadcastFeedMessage)
	var inboxReader *monitor.InboxReader
	for {
		inboxReader, err = mon.StartInboxReader(ctx, ethclint, rollupAddress, 0, bridgeUtilsAddress, nil, dummySequencerFeed)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", ethURL).
			Str("rollup", rollupAddress.Hex()).
			Str("bridgeUtils", bridgeUtilsAddress.Hex()).
			Msg("failed to start inbox reader, waiting and retrying")
		time.Sleep(time.Second * 5)
	}

	logger.Info().Hex("from", seqAuth.From.Bytes()).Msg("Arbitrum node submitting batches")

	nonce, err := ethclint.PendingNonceAt(ctx, deployer.From)
	if err != nil {
		return err
	}
	transferSize, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		return errors.New("invalid value for deposit amount")
	}
	transferTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(0),
		Gas:      21000,
		To:       &seqAuth.From,
		Value:    transferSize,
		Data:     nil,
	})
	transferTx, err = types.SignTx(transferTx, types.HomesteadSigner{}, privKey)
	if err != nil {
		return err
	}
	if err := ethclint.SendTransaction(ctx, transferTx); err != nil {
		return err
	}

	if err := ethbridge.WaitForBalance(
		ctx,
		ethclint,
		common.Address{},
		common.NewAddressFromEth(seqAuth.From),
	); err != nil {
		return errors.Wrap(err, "error waiting for balance")
	}

	batcherMode := rpc.SequencerBatcherMode{
		Auth:        seqAuth,
		Core:        mon.Core,
		InboxReader: inboxReader,
	}
	config := configuration.Config{
		Feed: configuration.Feed{
			Output: configuration.FeedOutput{
				Addr:          "127.0.0.1",
				IOTimeout:     2 * time.Second,
				Port:          "9642",
				Ping:          5 * time.Second,
				ClientTimeout: 15 * time.Second,
				Queue:         1,
				Workers:       2,
			},
		},
		Node: configuration.Node{
			Sequencer: configuration.Sequencer{
				CreateBatchBlockInterval:   *createBatchBlockInterval,
				DelayedMessagesTargetDelay: *delayedMessagesTargetDelay,
			},
		},
	}

	db, txDBErrChan, err := txdb.New(ctx, mon.Core, mon.Storage.GetNodeStore(), 100*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer db.Close()

	signer := func(hash []byte) ([]byte, error) {
		return crypto.Sign(hash, seqPrivKey)
	}

	batch, err := rpc.SetupBatcher(
		ctx,
		ethclint,
		rollupAddress,
		l2ChainId,
		db,
		time.Duration(5)*time.Second,
		batcherMode,
		signer,
		&config,
		&config.Wallet,
	)
	if err != nil {
		return err
	}

	srv := aggregator.NewServer(batch, rollupAddress, l2ChainId, db)

	// TODO: Add back in funding of fundedAccount
	// Note: The dev sequencer isn't being used anywhere currently
	//inboxCon, err := ethbridgecontracts.NewInbox(inboxAddress, ethclint)
	//if err != nil {
	//	return err
	//}
	//deployer.Value = transferSize
	//_, err = inboxCon.DepositEth(deployer, ethcommon.HexToAddress(*fundedAccount))
	//if err != nil {
	//	return err
	//}
	//
	//_, err = inboxCon.DepositEth(deployer, ethcommon.HexToAddress(*fundedAccount))
	//if err != nil {
	//	return err
	//}
	//deployer.Value = nil

	time.Sleep(time.Second * 40)

	if err := dev.EnableFees(srv, l2OwnerAuth, sequencer); err != nil {
		return err
	}

	web3Server, err := web3.GenerateWeb3Server(srv, nil, false, nil)
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	defer close(errChan)
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
		err := rpc.LaunchPublicServer(ctx, web3Server, rpcConfig, wsConfig)
		if err != nil {
			errChan <- err
		}
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
