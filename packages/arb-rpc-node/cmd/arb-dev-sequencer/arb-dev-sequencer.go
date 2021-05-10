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
	"flag"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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
	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)
	privKeyString := fs.String("privkey", "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d", "funded private key")

	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if fs.NArg() != 2 {
		fmt.Println("usage: arb-dev-sequencer <ethURL> <rollup_creator_address>")
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
	rollupCreator := common.HexToAddress(rollupCreatorAddresssString)

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

	privKey, err := crypto.HexToECDSA(*privKeyString)
	deployer := bind.NewKeyedTransactor(privKey)

	seqPrivKey, err := crypto.GenerateKey()
	if err != nil {
		return errors.Wrap(err, "error generating key")
	}
	seqAuth := bind.NewKeyedTransactor(seqPrivKey)

	owner := ethcommon.Address{}
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
	l1ChainId, err := ethclint.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting chain ID")
	}
	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	logger.Info().Hex("chainaddress", rollupAddress.Bytes()).Str("chainid", message.ChainAddressToID(rollupAddress).String()).Msg("Launching arbitrum node")

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
		inboxReader, err = mon.StartInboxReader(ctx, ethclint, rollupAddress, nil, dummySequencerFeed)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", ethURL).
			Str("rollup", rollupAddress.Hex()).
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
		Auth:                       seqAuth,
		Core:                       mon.Core,
		InboxReader:                inboxReader,
		DelayedMessagesTargetDelay: big.NewInt(*delayedMessagesTargetDelay),
	}
	broadcasterSettings := broadcaster.Settings{
		Addr:                    "",
		Workers:                 128,
		Queue:                   1,
		IoReadWriteTimeout:      2 * time.Second,
		ClientPingInterval:      5 * time.Second,
		ClientNoResponseTimeout: 15 * time.Second,
	}

	db, txDBErrChan, err := txdb.New(ctx, mon.Core, mon.Storage.GetNodeStore(), rollupAddress, 100*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer db.Close()

	var dummyDataSigner = func([]byte) ([]byte, error) {
		return common.HexToHash("0x0").Bytes(), nil
	}

	batch, err := rpc.SetupBatcher(
		ctx,
		ethclint,
		rollupAddress,
		db,
		time.Duration(5)*time.Second,
		batcherMode,
		dummyDataSigner,
		broadcasterSettings,
	)
	if err != nil {
		return err
	}

	srv := aggregator.NewServer(batch, rollupAddress, db)
	web3Server, err := web3.GenerateWeb3Server(srv, nil, false, nil)
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	defer close(errChan)
	go func() {
		err := rpc.LaunchPublicServer(ctx, web3Server, "8547", "8548")
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
