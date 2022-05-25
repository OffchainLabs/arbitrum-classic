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
	golog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"

	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
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

func startup() error {
	ctx, cancelFunc, _ := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	chainDir := fs.String("chaindir", "", "chain directory to dump ArbOS state of")
	arbosPath := fs.String("arbospath", "", "ArbOS mexe file path")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}
	if len(*chainDir) == 0 {
		flag.Usage()
		return errors.New("no chain directory specified")
	}
	if len(*arbosPath) == 0 {
		flag.Usage()
		return errors.New("no ArbOS mexe path")
	}

	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))); err != nil {
		return err
	}

	nodeConfig := configuration.DefaultNodeSettings()
	coreConfig := configuration.DefaultCoreSettingsMaxExecution()
	coreConfig.LazyLoadCoreMachine = true
	coreConfig.LazyLoadArchiveQueries = true
	mon, err := monitor.NewMonitor(filepath.Join(*chainDir, "db"), coreConfig)
	if err != nil {
		return err
	}
	if err := mon.Initialize(*arbosPath); err != nil {
		return err
	}
	if err := mon.Start(); err != nil {
		return err
	}
	defer mon.Close()

	nodeStore := mon.Storage.GetNodeStore()

	txDB, _, err := txdb.New(ctx, mon.Core, nodeStore, nodeConfig)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer txDB.Close()

	crossDB, err := New(txDB, path.Join(*chainDir, "ethdb"))
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

	return dumpArbState(mon.Core, nextLimit, path.Join(*chainDir, fmt.Sprint("state_", nextLimit)))
}
