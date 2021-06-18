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
	"encoding/json"
	"fmt"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/metrics"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

type ChainState struct {
	ValidatorWallet string `json:"validatorWallet"`
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-validator").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running validator")
	}
}

func startup() error {
	config, wallet, err := configuration.Parse()
	if err != nil || len(config.Database.Path) == 0 || len(config.L1.URL) == 0 ||
		len(config.Rollup.Address) == 0 || len(config.Bridge.Utils.Address) == 0 ||
		len(config.Validator.Utils.Address) == 0 || len(config.Validator.WalletFactory.Address) == 0 ||
		len(config.Validator.Strategy) != 0 {
		fmt.Printf("\n")
		fmt.Printf("usage arb-validator --conf=<filename> \n")
		fmt.Printf("   or arb-validator --l1.url=<url> --database.path=<path> --mainnet.arb1 \n")
		fmt.Printf("   or arb-validator --l1.url=<url> --database.path=<path> --testnet.rinkeby \n")
		if err != nil {
			return err
		}

		return nil
	}

	defer logger.Log().Msg("Cleanly shutting down validator")
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	if config.PProf.Enabled {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	// Dummy sequencerFeed since validator doesn't use it
	dummySequencerFeed := make(chan broadcaster.BroadcastFeedMessage)

	metricsConfig := metrics.NewMetricsConfig(&config.Healthcheck.Metrics.Prefix)
	metricsConfig.RegisterSystemMetrics()
	metricsConfig.RegisterStaticMetrics()

	const largeChannelBuffer = 200
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)

	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, metricsConfig.Registry, metricsConfig.Registerer)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	folder := os.Args[1]
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: config.Healthcheck.Metrics.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: config.Healthcheck.Sequencer.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: config.Healthcheck.L1Node.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: config.Healthcheck.Addr + ":" + config.Healthcheck.Port}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: config.L1.URL}
	nodehealth.Init(healthChan)

	client, err := ethutils.NewRPCEthClient(os.Args[2])
	if err != nil {
		return errors.Wrap(err, "error creating Ethereum RPC client")
	}
	var l1ChainId *big.Int
	for {
		l1ChainId, err = client.ChainID(ctx)
		if err == nil {
			break
		}
		logger.Warn().Err(err).Msg("Error getting chain ID")
		time.Sleep(time.Second * 5)
	}

	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	rollupAddr := ethcommon.HexToAddress(os.Args[3])
	bridgeUtilsAddr := ethcommon.HexToAddress(os.Args[4])
	validatorUtilsAddr := ethcommon.HexToAddress(os.Args[5])
	validatorWalletFactoryAddr := ethcommon.HexToAddress(os.Args[6])
	auth, _, err := cmdhelp.GetKeystore(config.Database.Path, wallet, l1ChainId)
	if err != nil {
		return errors.Wrap(err, "error loading wallet keystore")
	}
	logger.Info().Str("address", auth.From.String()).Msg("Loaded wallet")

	strategyString := os.Args[7]
	var strategy staker.Strategy
	if strategyString == "MakeNodes" {
		strategy = staker.MakeNodesStrategy
	} else if strategyString == "StakeLatest" {
		strategy = staker.StakeLatestStrategy
	} else if strategyString == "Defensive" {
		strategy = staker.DefensiveStrategy
	} else {
		return errors.New("unsupported strategy specified. Currently supported: MakeNodes, StakeLatest")
	}

	chainState := ChainState{}
	chainStatePath := path.Join(folder, "chainState.json")
	chainStateFile, err := os.Open(chainStatePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "failed to open chainState.json")
		}
	} else {
		chainStateData, err := ioutil.ReadAll(chainStateFile)
		if err != nil {
			return errors.Wrap(err, "failed to read chain state")
		}
		err = json.Unmarshal(chainStateData, &chainState)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal chain state")
		}
	}

	valAuth, err := ethbridge.NewTransactAuth(ctx, client, auth, config.GasPriceUrl)
	if err != nil {
		return errors.Wrap(err, "error creating connecting to chain")
	}
	validatorAddress := ethcommon.Address{}
	if chainState.ValidatorWallet == "" {
		for {
			validatorAddress, err = ethbridge.CreateValidatorWallet(ctx, validatorWalletFactoryAddr, config.Rollup.FromBlock, valAuth, client)
			if err == nil {
				break
			}
			logger.Warn().Err(err).
				Str("sender", auth.From.Hex()).
				Msg("Failed to deploy validator wallet")
			time.Sleep(time.Second * 5)
		}
		chainState.ValidatorWallet = validatorAddress.String()

		newChainStateData, err := json.Marshal(chainState)
		if err != nil {
			return errors.Wrap(err, "failed to marshal chain state")
		}
		if err := ioutil.WriteFile(chainStatePath, newChainStateData, 0644); err != nil {
			return errors.Wrap(err, "failed to write chain state config")
		}
	} else {
		validatorAddress = ethcommon.HexToAddress(chainState.ValidatorWallet)
	}

	dbPath := path.Join(folder, "arbStorage")
	arbosPath := path.Join(folder, "arbos.mexe")
	mon, err := monitor.NewMonitor(dbPath, arbosPath)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, valAuth)
	if err != nil {
		return errors.Wrap(err, "error creating validator wallet")
	}

	stakerManager, _, err := staker.NewStaker(ctx, mon.Core, client, val, config.Rollup.FromBlock, common.NewAddressFromEth(validatorUtilsAddr), strategy)
	if err != nil {
		return errors.Wrap(err, "error setting up staker")
	}

	chainMachineHash, err := stakerManager.GetInitialMachineHash(ctx)
	if err != nil {
		return errors.Wrap(err, "error checking initial chain state")
	}
	initialExecutionCursor, err := mon.Core.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		return errors.Wrap(err, "error loading initial ArbCore machine")
	}
	initialMachineHash := initialExecutionCursor.MachineHash()
	if initialMachineHash != chainMachineHash {
		return errors.Errorf("Initial machine hash loaded from arbos.mexe doesn't match chain's initial machine hash: chain %v, arbCore %v", hexutil.Encode(chainMachineHash[:]), initialMachineHash)
	}

	_, err = mon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), config.Rollup.FromBlock, common.NewAddressFromEth(bridgeUtilsAddr), healthChan, dummySequencerFeed)
	if err != nil {
		return errors.Wrap(err, "failed to create inbox reader")
	}

	logger.Info().Int("strategy", int(strategy)).Msg("Initialized validator")
	select {
	case <-cancelChan:
		return nil
	case <-stakerManager.RunInBackground(ctx):
		return nil
	}
}
