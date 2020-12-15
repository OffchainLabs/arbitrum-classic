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

package machineobserver

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/observer"
)

var logger = log.With().Caller().Str("component", "machineobserver").Logger()

const defaultMaxReorgDepth = 100

type TxDB interface {
	Load(ctx context.Context) error
	LatestBlockId() *common.BlockId
	AddInitialBlock(ctx context.Context, initialBlockHeight *big.Int) error
	AddMessages(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent, finishedBlock *common.BlockId) error
}

func verifyRollupInstance(
	ctx context.Context,
	initialMachineHash common.Hash,
	rollupWatcher arbbridge.ArbRollupWatcher,
) (arbbridge.ChainInfo, *big.Int, error) {
	if err := rollupWatcher.VerifyArbChain(ctx, initialMachineHash); err != nil {
		return arbbridge.ChainInfo{}, nil, err
	}

	logger.Info().Msg("L2 chain verified")

	_, eventCreated, _, creationTimestamp, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return arbbridge.ChainInfo{}, nil, err
	}
	logger.Info().Msg("Creation info retrieved")
	return eventCreated, creationTimestamp, nil
}

func ensureInitialized(
	ctx context.Context,
	db TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
	eventCreated arbbridge.ChainInfo,
	creationTimestamp *big.Int,
) error {
	logger.Info().Msg("Loading database")
	if err := db.Load(ctx); err != nil {
		return err
	}

	logger.Info().Msg("Database loaded")

	// If we're already initialized, do nothing
	if db.LatestBlockId() != nil {
		return nil
	}

	// We're starting from scratch.  Process the messages from initial block
	if err := db.AddInitialBlock(ctx, new(big.Int).Sub(eventCreated.BlockId.Height.AsInt(), big.NewInt(1))); err != nil {
		return err
	}

	events, err := inboxWatcher.GetDeliveredEventsInBlock(ctx, eventCreated.BlockId, creationTimestamp)
	if err != nil {
		return err
	}

	// filter out events before contract was created
	if len(events) > 0 {
		startIndex := -1
		for i, ev := range events {
			if ev.ChainInfo.Cmp(eventCreated) > 0 {
				startIndex = i
			}
		}
		if startIndex >= 0 {
			events = events[startIndex:]
		} else {
			events = nil
		}
	}

	if err := db.AddMessages(ctx, events, eventCreated.BlockId); err != nil {
		return err
	}

	logger.Info().Msg("Initial messages from first block have been added")

	return nil
}

func initializeCheckpointDB(
	rollupAddr common.Address,
	executablePath string,
	dbPath string,
) (*checkpointing.IndexedCheckpointer, error) {
	cp, err := checkpointing.NewIndexedCheckpointer(
		rollupAddr,
		dbPath,
		big.NewInt(defaultMaxReorgDepth),
		false,
	)
	if err != nil {
		return nil, err
	}

	if !cp.Initialized() {
		if err := cp.Initialize(executablePath); err != nil {
			return nil, err
		}
	}
	return cp, nil
}

func RunObserver(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	executablePath string,
	dbPath string,
) (*txdb.TxDB, error) {
	logger.Info().Msg("Creating indexed checkpointer")
	cp, err := initializeCheckpointDB(rollupAddr, executablePath, dbPath)
	if err != nil {
		return nil, err
	}

	// Verify contract version
	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return nil, err
	}

	initialMachine, err := cp.GetInitialMachine(valueCache)
	if err != nil {
		return nil, err
	}

	db := txdb.New(clnt, cp, cp.GetAggregatorStore(), rollupAddr)

	if err := ExecuteObserver(
		ctx,
		rollupAddr,
		clnt,
		initialMachine.Hash(),
		cp.MaxReorgHeight(),
		db,
	); err != nil {
		return nil, err
	}
	return db, nil
}

func ExecuteObserver(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	initialMachineHash common.Hash,
	maxReorgHeight *big.Int,
	db TxDB,
) error {
	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return err
	}

	eventCreated, creationTimestamp, err := verifyRollupInstance(ctx, initialMachineHash, rollupWatcher)
	if err != nil {
		return err
	}

	inboxAddr, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return err
	}

	inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	logger.Info().Msg("Initializing database")
	// Make first call to ensureInitialized outside of thread to avoid race conditions
	if err := ensureInitialized(ctx, db, inboxWatcher, eventCreated, creationTimestamp); err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	go observerRunThread(ctx, clnt, db, inboxWatcher, maxReorgHeight, eventCreated)
	return nil
}

func observerRunThread(
	ctx context.Context,
	clnt arbbridge.ChainInfoGetter,
	db TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
	maxReorgHeight *big.Int,
	eventCreated arbbridge.ChainInfo,
) {
	firstLoop := true
	for {
		runCtx, cancelFunc := context.WithCancel(ctx)

		logger.Info().Msg("Observer thread")

		err := func() error {
			if firstLoop {
				firstLoop = false
			} else {
				creationTimestamp, err := clnt.TimestampForBlockHash(runCtx, eventCreated.BlockId.HeaderHash)
				if err != nil {
					return err
				}
				if err := ensureInitialized(ctx, db, inboxWatcher, eventCreated, creationTimestamp); err != nil {
					logger.Fatal().Stack().Err(err).Send()
				}
			}

			logger.Info().
				Object("blockId", db.LatestBlockId()).
				Msg("Starting observer")

			// If the local chain is significantly behind the L1, catch up
			// more efficiently. We process `MaxReorgHeight` blocks at a
			// time up to `MaxReorgHeight` blocks before the current head
			// and and assume that no reorg will occur affecting the blocks
			// we are processing
			for {
				start := new(big.Int).Add(db.LatestBlockId().Height.AsInt(), big.NewInt(1))
				fetchEnd, err := observer.CalculateCatchupFetch(runCtx, start, clnt, maxReorgHeight)
				if err != nil {
					return errors.Wrap(err, "error calculating fast catchup")
				}
				if fetchEnd == nil {
					break
				}
				currentOnChain, err := clnt.BlockIdForHeight(ctx, nil)
				if err != nil {
					return err
				}
				logger.Info().
					Str("startEvent", start.String()).
					Str("endEvent", fetchEnd.String()).
					Str("blocksRemaining", new(big.Int).Sub(currentOnChain.Height.AsInt(), start).String()).
					Msg("Getting events")
				inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(runCtx, start, fetchEnd)
				if err != nil {
					return errors.Wrap(err, "Manager hit error doing fast catchup")
				}

				endBlock, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocks(fetchEnd))
				if err != nil {
					return errors.Wrap(err, "error getting end block in fast catchup")
				}
				if err := db.AddMessages(runCtx, inboxDeliveredEvents, endBlock); err != nil {
					return errors.Wrap(err, "error adding messages to db")
				}
			}

			latest := db.LatestBlockId()
			headersChan, err := clnt.SubscribeBlockHeadersAfter(runCtx, latest)
			if err != nil {
				return errors.Wrap(err, "can't restart header subscription")
			}
			for maybeBlockId := range headersChan {
				if maybeBlockId.Err != nil {
					return errors.Wrap(maybeBlockId.Err, "error getting new header")
				}

				blockId := maybeBlockId.BlockId
				timestamp := maybeBlockId.Timestamp

				inboxEvents, err := inboxWatcher.GetDeliveredEventsInBlock(runCtx, blockId, timestamp)
				if err != nil {
					return errors.Wrapf(err, "manager hit error getting inbox events with block %v", blockId)
				}

				if err := db.AddMessages(runCtx, inboxEvents, blockId); err != nil {
					return errors.Wrap(err, "error adding messages to db")
				}
			}
			return nil
		}()

		if err != nil {
			logger.Error().Stack().Err(err).Msg("Error in observer manager")
		}

		cancelFunc()

		select {
		case <-ctx.Done():
			return
		default:

		}
		// Wait for things to settle
		time.Sleep(time.Second)
	}
}
