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

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

const defaultMaxReorgDepth = 100

func calculateCatchupFetch(ctx context.Context, start *big.Int, clnt arbbridge.ChainTimeGetter, maxReorg *big.Int) (*big.Int, error) {
	currentLocalHeight := start
	currentOnChain, err := clnt.CurrentBlockId(ctx)
	if err != nil {
		return nil, err
	}
	currentL1Height := currentOnChain.Height.AsInt()

	fastCatchupEndHeight := new(big.Int).Sub(currentL1Height, maxReorg)
	if currentLocalHeight.Cmp(fastCatchupEndHeight) >= 0 {
		return nil, nil
	}

	fetchSize := new(big.Int).Sub(fastCatchupEndHeight, currentLocalHeight)
	if fetchSize.Cmp(big.NewInt(1)) <= 0 {
		return nil, nil
	}
	if fetchSize.Cmp(maxReorg) >= 0 {
		fetchSize = maxReorg
	}
	fetchEnd := new(big.Int).Add(currentLocalHeight, fetchSize)
	fetchEnd = fetchEnd.Sub(fetchEnd, big.NewInt(1))
	return fetchEnd, nil
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
	initialMachine, err := cp.GetInitialMachine()
	if err != nil {
		return nil, err
	}

	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}
	if err := rollupWatcher.VerifyArbChain(ctx, initialMachine.Hash()); err != nil {
		return nil, err
	}

	inboxAddr, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return nil, err
	}

	_, blockCreated, _, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return nil, err
	}

	db, err := txdb.New(ctx, clnt, cp, cp.GetAggregatorStore(), blockCreated)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			if err := db.RestoreFromCheckpoint(ctx); err != nil {
				log.Fatal(err)
			}

			latest, err := db.LatestBlock()
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Starting observer from", latest)

			err = func() error {
				// If the local chain is significantly behind the L1, catch up
				// more efficiently. We process `MaxReorgHeight` blocks at a
				// time up to `MaxReorgHeight` blocks before the current head
				// and and assume that no reorg will occur affecting the blocks
				// we are processing
				maxReorg := cp.MaxReorgHeight()
				for {
					latestBlock, err := db.LatestBlock()
					if err != nil {
						return err
					}
					start := new(big.Int).Add(latestBlock.Height.AsInt(), big.NewInt(1))
					fetchEnd, err := calculateCatchupFetch(runCtx, start, clnt, maxReorg)
					if err != nil {
						return err
					}
					if fetchEnd == nil {
						break
					}
					log.Println("Getting events between", start, "and", fetchEnd)
					inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(runCtx, start, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}
					if err := db.AddMessages(runCtx, inboxDeliveredEvents); err != nil {
						return err
					}
				}

				latest, err := db.LatestBlock()
				if err != nil {
					return err
				}

				headersChan, err := clnt.SubscribeBlockHeadersAfter(runCtx, latest)
				if err != nil {
					return errors2.Wrap(err, "Error subscribing to block headers")
				}

				for maybeBlockId := range headersChan {
					if maybeBlockId.Err != nil {
						return errors2.Wrap(maybeBlockId.Err, "Error getting new header")
					}

					blockId := maybeBlockId.BlockId
					timestamp := maybeBlockId.Timestamp

					inboxEvents, err := inboxWatcher.GetDeliveredEventsInBlock(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "Manager hit error getting inbox events with block %v", blockId)
					}

					if err := db.AddMessages(runCtx, inboxEvents); err != nil {
						return err
					}
				}
				return nil
			}()

			if err != nil {
				log.Println(err)
			}

			cancelFunc()

			select {
			case <-ctx.Done():
				return
			default:

			}
		}
	}()
	return db, nil
}
