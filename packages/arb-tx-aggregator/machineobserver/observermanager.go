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
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/txdb"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"sync"
)

type reorgCache struct {
	latestValidMachine    machine.Machine
	currentPendingMachine machine.Machine
	currentBlockId        *common.BlockId
}

type Manager struct {
	// The mutex should be held whenever listeres or reorgCache are accessed or
	// set and whenever activeChain is going to be set (but not when it is accessed)
	sync.Mutex

	// reorgCache is nil when the validator is functioning normally. When the
	// validator experiences a reorg it stores the current state in the reorg
	// cache. It uses this cache to respond to non-mutating queries from users
	// until it is caught back up with the latest state at which point it clears
	// the cache and starts answering queries based on the current state.
	// This approach let's us provide a best effort response to users quickly
	// rather than blocking until the validator fully recovers from the reorg.
	reorgCache *reorgCache

	// These variables are only written by the constructor
	RollupAddress common.Address
	checkpointer  checkpointing.RollupCheckpointer
	txdb          *txdb.TxDB
}

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

func CreateManager(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	aoFilePath string,
	dbPath string,
) (*Manager, error) {
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
		if err := cp.Initialize(aoFilePath); err != nil {
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

	man := &Manager{
		RollupAddress: rollupAddr,
		checkpointer:  cp,
	}
	var initialBlockId *common.BlockId
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			inboxAddr, err := rollupWatcher.InboxAddress(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			db, err := txdb.New(runCtx, clnt, cp, nil, initialBlockId)
			if err != nil {
				log.Fatal(err)
			}

			blockCount, err := db.BlockCount()
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Starting observer from", blockCount)

			err = func() error {
				// If the local chain is significantly behind the L1, catch up
				// more efficiently. We process `MaxReorgHeight` blocks at a
				// time up to `MaxReorgHeight` blocks before the current head
				// and and assume that no reorg will occur affecting the blocks
				// we are processing
				maxReorg := cp.MaxReorgHeight()
				for {
					blockCount, err := db.BlockCount()
					if err != nil {
						return err
					}
					start := new(big.Int).SetUint64(blockCount)
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
					if err := man.txdb.AddMessages(runCtx, inboxDeliveredEvents); err != nil {
						return err
					}
				}

				//startBlock
				//
				//headersChan, err := clnt.SubscribeBlockHeaders(runCtx, observer.currentBlockId)
				//if err != nil {
				//	return errors2.Wrap(err, "Error subscribing to block headers")
				//}
				//
				//for maybeBlockId := range headersChan {
				//	if maybeBlockId.Err != nil {
				//		return errors2.Wrap(maybeBlockId.Err, "Error getting new header")
				//	}
				//
				//	blockId := maybeBlockId.BlockId
				//	timestamp := maybeBlockId.Timestamp
				//
				//	inboxEvents, err := inboxWatcher.GetDeliveredEventsInBlock(runCtx, blockId, timestamp)
				//	if err != nil {
				//		return errors2.Wrapf(err, "Manager hit error getting inbox events with block %v", blockId)
				//	}
				//
				//	for _, ev := range inboxEvents {
				//		observer.processNextMessage(ev)
				//	}
				//}
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

	return man, nil
}
