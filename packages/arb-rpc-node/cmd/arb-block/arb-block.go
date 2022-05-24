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
	"fmt"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strings"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

var logger zerolog.Logger

var pprofMux *http.ServeMux

const largeChannelBuffer = 200

const (
	failLimit            = 6
	checkFrequency       = time.Second * 30
	blockCheckCountDelay = 5
)

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
	logger = arblog.Logger.With().Str("component", "arb-node").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running node")
		fmt.Printf("\nNotice: %s\n\n", err.Error())
	}
}

func printSampleUsage() {
	fmt.Printf("\n")
	fmt.Printf("Sample usage: arb-block --conf=<filename> \n")
	fmt.Printf("          or: arb-block --l1.url=<L1 RPC> [optional arguments]\n\n")
}

func startup() error {
	ctx, cancelFunc, _ := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	config, _, l1Client, _, err := configuration.ParseCLI(ctx)
	if err != nil || len(config.Persistent.GlobalConfig) == 0 || len(config.L1.URL) == 0 ||
		len(config.Rollup.Address) == 0 || len(config.BridgeUtilsAddress) == 0 ||
		((config.Node.Type() != configuration.SequencerNodeType) && len(config.Node.Sequencer.Lockout.Redis) != 0) ||
		((len(config.Node.Sequencer.Lockout.Redis) == 0) != (len(config.Node.Sequencer.Lockout.SelfRPCURL) == 0)) {
		printSampleUsage()
		if err != nil && !strings.Contains(err.Error(), "help requested") {
			fmt.Printf("\n%s\n", err.Error())
		}

		return nil
	}

	if config.Core.Database.Metadata {
		return cmdhelp.PrintDatabaseMetadata(config.GetDatabasePath(), &config.Core)
	}

	if config.Persistent.Chain == "" {
		return errors.Errorf("Missing --persistent.chain")
	}
	if config.Rollup.Address == "" {
		return errors.Errorf("Missing --rollup.address")
	}
	if config.Node.ChainID == 0 {
		return errors.Errorf("Missing --node.chain-id")
	}
	if config.Rollup.Machine.Filename == "" {
		return errors.Errorf("Missing --rollup.machine.filename")
	}

	defer logger.Log().Msg("Cleanly shutting down node")

	if err := cmdhelp.ParseLogFlags(&config.Log.RPC, &config.Log.Core, gethlog.StreamHandler(os.Stderr, gethlog.JSONFormat())); err != nil {
		return err
	}

	l2ChainId := new(big.Int).SetUint64(config.Node.ChainID)
	rollupAddress := common.HexToAddress(config.Rollup.Address)
	logger.Info().
		Hex("chainaddress", rollupAddress.Bytes()).
		Hex("chainid", l2ChainId.Bytes()).
		Str("type", config.Node.TypeImpl).
		Int64("fromBlock", config.Rollup.FromBlock).
		Msg("Launching arbitrum node")

	rollup, err := ethbridge.NewRollupWatcher(rollupAddress.ToEthAddress(), config.Rollup.FromBlock, l1Client, bind.CallOpts{})
	if err != nil {
		return err
	}

	if config.Node.Type() == configuration.ValidatorNodeType && config.Core.CheckpointMaxExecutionGas != 0 {
		log.Warn().Msg("allowing for infinite core execution because running as validator")
		config.Core.CheckpointMaxExecutionGas = 0
	}

	mon, err := monitor.NewMonitor(config.GetDatabasePath(), &config.Core)
	if err != nil {
		return err
	}
	if err := mon.Initialize(config.Rollup.Machine.Filename); err != nil {
		return err
	}
	if err := mon.Start(); err != nil {
		return err
	}
	defer mon.Close()

	if err := cmdhelp.UpdatePrunePoint(ctx, rollup, mon.Core); err != nil {
		logger.Error().Err(err).Msg("error pruning database")
	}

	nodeStore := mon.Storage.GetNodeStore()

	txDB, _, err := txdb.New(ctx, mon.Core, nodeStore, &config.Node)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer txDB.Close()

	crossDB, err := New(txDB, path.Join(config.Persistent.Chain, "ethdb"))
	if err != nil {
		return err
	}

	maxBlocks, err := txDB.BlockCount()
	if err != nil {
		return err
	}

	currBlocks, err := crossDB.EthBlockNum()
	if err != nil {
		return err
	}

	nextLimit := currBlocks + 1000
	if nextLimit > maxBlocks {
		nextLimit = maxBlocks
	}

	err = crossDB.FillerUp(ctx, nextLimit)
	if err != nil {
		return err
	}

	return dumpArbState(mon.Core, nextLimit, path.Join(config.Persistent.Chain, fmt.Sprint("state_", nextLimit)))
}
