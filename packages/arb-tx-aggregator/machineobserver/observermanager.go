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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/observer"
)

var logger = log.With().Caller().Str("component", "machineobserver").Logger()

const defaultMaxReorgDepth = 100

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
	db *txdb.TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
) error {
	logger.Info().Msg("Loading database")
	freshStart, err := db.Load(ctx)
	if err != nil {
		return err
	}

	logger.Info().Msg("Database loaded")

	// If we're already initialized, do nothing
	if !freshStart {
		return nil
	}

	// We're starting from scratch.  Process the messages from initial block
	events, err := inboxWatcher.GetDeliveredEventsInBlock(ctx, db.EventCreated.BlockId, db.CreationTimestamp)
	if err != nil {
		return err
	}

	// filter out events before contract was created
	if len(events) > 0 {
		startIndex := -1
		for i, ev := range events {
			if ev.ChainInfo.Cmp(db.EventCreated) > 0 {
				startIndex = i
			}
		}
		if startIndex >= 0 {
			events = events[startIndex:]
		} else {
			events = nil
		}
	}

	if _, err := db.AddMessages(ctx, extractMessages(events), db.EventCreated.BlockId); err != nil {
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

	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}

	eventCreated, creationTimestamp, err := verifyRollupInstance(ctx, initialMachine.Hash(), rollupWatcher)
	if err != nil {
		return nil, err
	}

	db, err := txdb.New(clnt, cp, cp.GetAggregatorStore(), rollupAddr, eventCreated, creationTimestamp)
	if err != nil {
		return nil, err
	}

	inboxAddr, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return nil, err
	}

	inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	if err := ExecuteObserverAdvanced(
		ctx,
		clnt,
		cp.MaxReorgHeight(),
		db,
		inboxWatcher,
	); err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteObserverAdvanced(
	ctx context.Context,
	clnt arbbridge.ChainInfoGetter,
	maxReorgHeight *big.Int,
	db *txdb.TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
) error {
	logger.Info().Msg("Initializing database")
	// Make first call to ensureInitialized outside of thread to avoid race conditions
	if err := ensureInitialized(ctx, db, inboxWatcher); err != nil {
		return err
	}

	go observerRunThread(ctx, clnt, db, inboxWatcher, maxReorgHeight)
	return nil
}

func observerRunThread(
	ctx context.Context,
	clnt arbbridge.ChainInfoGetter,
	db *txdb.TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
	maxReorgHeight *big.Int,
) {
	firstLoop := true
	for {
		runCtx, cancelFunc := context.WithCancel(ctx)

		logger.Info().Msg("Observer thread")

		err := func() error {
			if firstLoop {
				firstLoop = false
			} else {
				if err := ensureInitialized(ctx, db, inboxWatcher); err != nil {
					logger.Fatal().Stack().Err(err).Send()
				}
			}

			logger.Info().
				Object("blockId", db.LatestBlockId()).
				Msg("Starting observer")

			if err := fastCatchupEvents(runCtx, clnt, db, inboxWatcher, maxReorgHeight); err != nil {
				return err
			}

			if err := watchEvents(runCtx, clnt, db, inboxWatcher); err != nil {
				return err
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

func fastCatchupEvents(
	ctx context.Context,
	clnt arbbridge.ChainInfoGetter,
	db *txdb.TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
	maxReorgHeight *big.Int,
) error {
	// If the local chain is significantly behind the L1, catch up
	// more efficiently. We process `MaxReorgHeight` blocks at a
	// time up to `MaxReorgHeight` blocks before the current head
	// and and assume that no reorg will occur affecting the blocks
	// we are processing
	for {
		start := new(big.Int).Add(db.LatestBlockId().Height.AsInt(), big.NewInt(1))
		fetchEnd, err := observer.CalculateCatchupFetch(ctx, start, clnt, maxReorgHeight)
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
		inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(ctx, start, fetchEnd)
		if err != nil {
			return errors.Wrap(err, "Manager hit error doing fast catchup")
		}

		endBlock, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocks(fetchEnd))
		if err != nil {
			return errors.Wrap(err, "error getting end block in fast catchup")
		}
		if _, err := db.AddMessages(ctx, extractMessages(inboxDeliveredEvents), endBlock); err != nil {
			return errors.Wrap(err, "error adding messages to db")
		}
	}
	return nil
}

func watchEvents(
	ctx context.Context,
	clnt arbbridge.ChainInfoGetter,
	db *txdb.TxDB,
	inboxWatcher arbbridge.GlobalInboxWatcher,
) error {
	latest := db.LatestBlockId()
	headersChan, err := clnt.SubscribeBlockHeadersAfter(ctx, latest)
	if err != nil {
		return errors.Wrap(err, "can't restart header subscription")
	}
	for maybeBlockId := range headersChan {
		if maybeBlockId.Err != nil {
			return errors.Wrap(maybeBlockId.Err, "error getting new header")
		}

		blockId := maybeBlockId.BlockId
		timestamp := maybeBlockId.Timestamp

		inboxEvents, err := inboxWatcher.GetDeliveredEventsInBlock(ctx, blockId, timestamp)
		if err != nil {
			return errors.Wrapf(err, "manager hit error getting inbox events with block %v", blockId)
		}

		if _, err := db.AddMessages(ctx, extractMessages(inboxEvents), blockId); err != nil {
			return errors.Wrap(err, "error adding messages to db")
		}
	}
	return nil
}

func extractMessages(inboxEvents []arbbridge.MessageDeliveredEvent) []inbox.InboxMessage {
	inboxMessages := make([]inbox.InboxMessage, 0, len(inboxEvents))
	for _, ev := range inboxEvents {
		inboxMessages = append(inboxMessages, ev.Message)
	}
	return inboxMessages
}
