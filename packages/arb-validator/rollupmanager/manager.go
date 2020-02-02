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

package rollupmanager

import (
	"context"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type Manager struct {
	sync.Mutex
	RollupAddress common.Address
	client        arbbridge.ArbClient
	listeners     []rollup.ChainListener

	listenerAddChan chan rollup.ChainListener
	actionChan      chan func(*rollup.ChainObserver)
	ckpFac          checkpointing.RollupCheckpointerFactory
}

const defaultMaxReorgDepth = 100

func CreateManager(
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	aoFilePath string,
	dbPath string,
) (*Manager, error) {
	return CreateManagerAdvanced(
		context.Background(),
		rollupAddr,
		true,
		clnt,
		checkpointing.NewRollupCheckpointerImplFactory(
			rollupAddr,
			aoFilePath,
			dbPath,
			big.NewInt(defaultMaxReorgDepth),
			false,
		),
	)
}

func CreateManagerAdvanced(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ArbClient,
	ckpFac checkpointing.RollupCheckpointerFactory,
) (*Manager, error) {
	man := &Manager{
		RollupAddress:   rollupAddr,
		client:          clnt,
		listenerAddChan: make(chan rollup.ChainListener, 10),
		actionChan:      make(chan func(*rollup.ChainObserver), 10),
		ckpFac:          ckpFac,
	}
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			checkpointer := man.ckpFac.New(runCtx)

			var chain *rollup.ChainObserver

			watcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			pendingInboxAddress, err := watcher.InboxAddress(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			pendingInboxWatcher, err := clnt.NewPendingInboxWatcher(pendingInboxAddress, rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			if checkpointer.HasCheckpointedState() {
				chainObserverBytes, restoreCtx, err := checkpointer.RestoreLatestState(runCtx, clnt, rollupAddr, updateOpinion)
				if err != nil {
					log.Fatal(err)
				}
				chainObserverBuf := &rollup.ChainObserverBuf{}
				if err := proto.Unmarshal(chainObserverBytes, chainObserverBuf); err != nil {
					log.Fatal(err)
				}
				chain, err = chainObserverBuf.UnmarshalFromCheckpoint(runCtx, restoreCtx, checkpointer)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				params, err := watcher.GetParams(ctx)
				if err != nil {
					log.Fatal(err)
				}
				blockID, err := watcher.GetCreationHeight(ctx)
				if err != nil {
					log.Fatal(err)
				}
				chain, err = rollup.NewChain(rollupAddr, checkpointer, params, updateOpinion, blockID)
				if err != nil {
					log.Fatal(err)
				}
			}

			log.Println("Starting validator from", chain.CurrentBlockID())

			man.Lock()
			// Clear pending listeners
			for len(man.listenerAddChan) > 0 {
				<-man.listenerAddChan
			}
			// Add manager's listeners
			for _, listener := range man.listeners {
				chain.AddListener(listener)
			}
			man.Unlock()

			chain.Start(runCtx)

			current, err := clnt.CurrentBlockID(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			headersChan, err := clnt.SubscribeBlockHeaders(runCtx, chain.CurrentBlockID())
			if err != nil {
				blockID, err := clnt.BlockIDForHeight(ctx, common.NewTimeBlocks(big.NewInt(0)))
				if err != nil {
					panic(err)
				}
				log.Println("Error subscribing to block headers", chain.CurrentBlockID().HeaderHash, chain.CurrentBlockID().Height.AsInt(), blockID.HeaderHash, blockID.Height.AsInt(), err)

				cancelFunc()
				time.Sleep(2 * time.Second)
				continue
			}
			reachedHead := false
		runLoop:
			for {
				select {
				case maybeBlockID, ok := <-headersChan:
					if !ok {
						log.Println("Manager stopped receiving headers")
						break runLoop
					}
					if maybeBlockID.Err != nil {
						log.Println("Error getting new header", maybeBlockID.Err)
						break runLoop
					}

					blockID := maybeBlockID.BlockID

					if !reachedHead && blockID.Height.Cmp(current.Height) >= 0 {
						log.Println("Reached head")
						reachedHead = true
						chain.NowAtHead()
						log.Println("Now at head")
					}

					chain.NotifyNewBlock(blockID.Clone())
					log.Print(chain.DebugString("== "))

					inboxEvents, err := pendingInboxWatcher.GetEvents(runCtx, blockID)
					if err != nil {
						log.Println("Manager hit error getting events", err)
						break runLoop
					}
					for _, event := range inboxEvents {
						chain.HandleNotification(runCtx, event)
					}

					events, err := watcher.GetEvents(runCtx, blockID)
					if err != nil {
						log.Println("Manager hit error getting events", err)
						break runLoop
					}
					for _, event := range events {
						chain.HandleNotification(runCtx, event)
					}
				case action := <-man.actionChan:
					action(chain)
				}
			}

			cancelFunc()

			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Second) // give time for things to settle, post-reorg, before restarting stuff
			}
		}
	}()

	return man, nil
}

func (man *Manager) AddListener(listener rollup.ChainListener) {
	man.Lock()
	man.listeners = append(man.listeners, listener)
	man.listenerAddChan <- listener
	man.Unlock()
}

func (man *Manager) ExecuteCall(messages value.TupleValue, maxTime time.Duration) (*protocol.ExecutionAssertion, uint64) {
	retChan := make(chan struct {
		*protocol.ExecutionAssertion
		uint64
	}, 1)
	man.actionChan <- func(chain *rollup.ChainObserver) {
		mach := chain.LatestKnownValidMachine()
		latestTime := chain.CurrentBlockID().Height
		timeBounds := &protocol.TimeBoundsBlocks{Start: latestTime, End: latestTime}
		go func() {
			assertion, numSteps := mach.ExecuteAssertion(
				// Call execution is only limited by wall time, so use a massive max steps as an approximation to infinity
				10000000000000000,
				timeBounds,
				messages,
				maxTime,
			)
			retChan <- struct {
				*protocol.ExecutionAssertion
				uint64
			}{assertion, numSteps}
		}()
	}
	ret := <-retChan
	return ret.ExecutionAssertion, ret.uint64
}

func (man *Manager) CurrentBlockID() *structures.BlockID {
	retChan := make(chan *structures.BlockID, 1)
	man.actionChan <- func(chain *rollup.ChainObserver) {
		retChan <- chain.CurrentBlockID()
	}
	return <-retChan
}
