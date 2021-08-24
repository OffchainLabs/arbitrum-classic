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
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
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
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	config, walletConfig, feedSignerConfig, l1Client, l1ChainId, err := configuration.ParseValidator(ctx)
	if err != nil || len(config.Persistent.GlobalConfig) == 0 || len(config.L1.URL) == 0 ||
		len(config.Rollup.Address) == 0 || len(config.BridgeUtilsAddress) == 0 ||
		len(config.Validator.UtilsAddress) == 0 || len(config.Validator.WalletFactoryAddress) == 0 ||
		len(config.Validator.Strategy) == 0 {
		fmt.Printf("\n")
		fmt.Printf("Sample usage: arb-validator --conf=<filename> \n")
		fmt.Printf("          or: arb-validator --persistent.storage.path=<path> --l1.url=<L1 RPC> --feed.input.url=<feed websocket>\n\n")
		if err != nil && !strings.Contains(err.Error(), "help requested") {
			fmt.Printf("%s\n", err.Error())
		}

		return nil
	}

	defer logger.Log().Msg("Cleanly shutting down validator")

	if config.PProfEnable {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	// Dummy sequencerFeed since validator doesn't use it
	dummySequencerFeed := make(chan broadcaster.BroadcastFeedMessage)

	metricsConfig := metrics.NewMetricsConfig(config.MetricsServer, &config.Healthcheck.MetricsPrefix)

	const largeChannelBuffer = 200
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)

	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, metricsConfig.Registry)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: config.Healthcheck.Metrics}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: !config.Healthcheck.Sequencer}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: !config.Healthcheck.L1Node}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: config.Healthcheck.Addr + ":" + config.Healthcheck.Port}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: config.L1.URL}
	nodehealth.Init(healthChan)

	rollupAddr := ethcommon.HexToAddress(config.Rollup.Address)
	bridgeUtilsAddr := ethcommon.HexToAddress(config.BridgeUtilsAddress)
	validatorUtilsAddr := ethcommon.HexToAddress(config.Validator.UtilsAddress)
	validatorWalletFactoryAddr := ethcommon.HexToAddress(config.Validator.WalletFactoryAddress)
	auth, _, err := cmdhelp.GetKeystore(config, walletConfig, feedSignerConfig, l1ChainId, false)
	if err != nil {
		return errors.Wrap(err, "error loading wallet keystore")
	}
	logger.Info().Str("address", auth.From.String()).Msg("Loaded wallet")

	strategyString := config.Validator.Strategy
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
	chainStatePath := path.Join(config.Persistent.Chain, "chainState.json")
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

	var valAuth *ethbridge.TransactAuth
	if len(config.Wallet.Fireblocks.SSLKey) > 0 {
		valAuth, _, err = ethbridge.NewFireblocksTransactAuthAdvanced(ctx, l1Client, auth, walletConfig, false)
	} else {
		valAuth, err = ethbridge.NewTransactAuthAdvanced(ctx, l1Client, auth, false)

	}
	if err != nil {
		return errors.Wrap(err, "error creating connecting to chain")
	}
	validatorAddress := ethcommon.Address{}
	if chainState.ValidatorWallet == "" {
		for {
			validatorAddress, err = ethbridge.CreateValidatorWallet(ctx, validatorWalletFactoryAddr, config.Rollup.FromBlock, valAuth, l1Client)
			if err == nil {
				break
			}
			logger.Warn().Err(err).
				Str("sender", auth.From.Hex()).
				Msg("Failed to deploy validator wallet")

			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Second * 5):
			}
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

	mon, err := monitor.NewMonitor(config.GetValidatorDatabasePath(), config.Rollup.Machine.Filename, &config.Core)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, l1Client, valAuth)
	if err != nil {
		return errors.Wrap(err, "error creating validator wallet")
	}

	stakerManager, _, err := staker.NewStaker(ctx, mon.Core, l1Client, val, config.Rollup.FromBlock, common.NewAddressFromEth(validatorUtilsAddr), strategy, bind.CallOpts{}, valAuth, config.Validator)
	if err != nil {
		return errors.Wrap(err, "error setting up staker")
	}

	_, err = mon.StartInboxReader(ctx, l1Client, common.NewAddressFromEth(rollupAddr), config.Rollup.FromBlock, common.NewAddressFromEth(bridgeUtilsAddr), healthChan, dummySequencerFeed)
	if err != nil {
		return errors.Wrap(err, "failed to create inbox reader")
	}

	logger.Info().Int("strategy", int(strategy)).Msg("Initialized validator")
	select {
	case <-cancelChan:
		return nil
	case <-stakerManager.RunInBackground(ctx, config.Validator.StakerDelay):
		return nil
	}
}
