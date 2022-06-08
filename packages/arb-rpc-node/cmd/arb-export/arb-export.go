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
	"os"
	"path"
	"path/filepath"
	"time"

	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/nitroexport"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

var logger zerolog.Logger

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = arblog.Logger.With().Str("component", "arb-export").Logger()

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
	blockNum := fs.Uint64("blocknum", 0, "block number to import")
	skipBlocks := fs.Bool("skipblocks", false, "don't import blocks")
	skipState := fs.Bool("skipstate", false, "don't import state")
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

	crossDB, err := nitroexport.NewCrossDB(txDB, path.Join(*chainDir, "ethdb"))
	if err != nil {
		return err
	}

	maxBlocks, err := txDB.BlockCount()
	if err != nil {
		return err
	}

	if *blockNum > maxBlocks {
		return errors.Errorf("requested {} blocks but only {} exist", *blockNum, maxBlocks)
	}

	if *blockNum == 0 {
		*blockNum = maxBlocks
		logger.Info().Uint64("blocks", *blockNum)
	}

	if !*skipBlocks {
		crossDB.Start(ctx)
		crossDB.UpdateTarget(*blockNum + 1)
		logger.Info().Msg("starting history export")
	}

	if !*skipState {
		logger.Info().Msg("starting state export")
		err = nitroexport.ExportState(mon.Core, *blockNum, path.Join(*chainDir, fmt.Sprint("state_", *blockNum)))
		if err != nil {
			return err
		}
		logger.Info().Msg("state export done")
	}

	if !*skipBlocks {
		for {
			blocksDone, err := crossDB.BlocksExported()
			if err != nil {
				return err
			}
			if blocksDone > *blockNum {
				break
			}
			logger.Info().Uint64("imported", blocksDone).Uint64("out of", *blockNum+1).Msg(".. importing blocks")
			err = crossDB.CurrentError()
			if err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				break
			case <-time.After(time.Minute):
			}
		}
		logger.Info().Msg("history export done")
	}

	return nil
}
