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
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/observer"
)

const defaultMaxReorgDepth = 100

func ensureInitialized(
	ctx context.Context,
	cp *checkpointing.IndexedCheckpointer,
	db *txdb.TxDB,
	clnt arbbridge.ArbClient,
	rollupAddr common.Address,
) error {
	if err := db.Load(ctx); err != nil {
		return err
	}

	// If we're already initialized, do nothing
	if db.LatestBlockId() != nil {
		return nil
	}

	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return err
	}

	inboxAddr, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return err
	}

	// We're starting from scratch. Process the messages from the partial block
	inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
	if err != nil {
		return err
	}

	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return err
	}
	defer cmachine.DestroyValueCache(valueCache)

	initialMachine, err := cp.GetInitialMachine(valueCache)
	if err != nil {
		return err
	}

	if err := rollupWatcher.VerifyArbChain(ctx, initialMachine.Hash()); err != nil {
		return err
	}

	_, eventCreated, _, creationTimestamp, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	events, err := inboxWatcher.GetDeliveredEventsInBlock(ctx, eventCreated.BlockId, creationTimestamp)
	if err != nil {
		return err
	}

	// filter out events before nextEventId
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

	return nil
}

func RunObserver(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	executablePath string,
	dbPath string,
) (*txdb.TxDB, error) {
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

	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}

	inboxAddr, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return nil, err
	}

	db := txdb.New(clnt, cp, cp.GetAggregatorStore(), rollupAddr)

	if err := ensureInitialized(ctx, cp, db, clnt, rollupAddr); err != nil {
		return nil, err
	}

	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			if err := ensureInitialized(ctx, cp, db, clnt, rollupAddr); err != nil {
				log.Fatal(err)
			}

			err = func() error {
				log.Println("Starting observer after", db.LatestBlockId())

				// If the local chain is significantly behind the L1, catch up
				// more efficiently. We process `MaxReorgHeight` blocks at a
				// time up to `MaxReorgHeight` blocks before the current head
				// and and assume that no reorg will occur affecting the blocks
				// we are processing
				maxReorg := cp.MaxReorgHeight()
				for {
					start := new(big.Int).Add(db.LatestBlockId().Height.AsInt(), big.NewInt(1))
					fetchEnd, err := observer.CalculateCatchupFetch(runCtx, start, clnt, maxReorg)
					if err != nil {
						return errors2.Wrap(err, "error calculating fast catchup")
					}
					if fetchEnd == nil {
						break
					}
					log.Println("Getting events between", start, "and", fetchEnd)
					inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(runCtx, start, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}

					endBlock, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocks(fetchEnd))
					if err != nil {
						return errors2.Wrap(err, "error getting end block in fast catchup")
					}
					if err := db.AddMessages(runCtx, inboxDeliveredEvents, endBlock); err != nil {
						return errors2.Wrap(err, "error adding messages to db")
					}
				}

				latest := db.LatestBlockId()
				headersChan, err := clnt.SubscribeBlockHeadersAfter(runCtx, latest)
				if err != nil {
					return errors2.Wrap(err, "can't restart header subscription")
				}
				for maybeBlockId := range headersChan {
					if maybeBlockId.Err != nil {
						return errors2.Wrap(maybeBlockId.Err, "error getting new header")
					}

					blockId := maybeBlockId.BlockId
					timestamp := maybeBlockId.Timestamp

					inboxEvents, err := inboxWatcher.GetDeliveredEventsInBlock(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "manager hit error getting inbox events with block %v", blockId)
					}

					if err := db.AddMessages(runCtx, inboxEvents, blockId); err != nil {
						return errors2.Wrap(err, "error adding messages to db")
					}
				}
				return nil
			}()

			if err != nil {
				log.Println("Error in observer manager:", err)
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
	}()
	return db, nil
}
