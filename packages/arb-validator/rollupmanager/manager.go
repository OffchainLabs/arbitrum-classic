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

const (
	maxReorgDepth = 100
)

type Manager struct {
	sync.Mutex
	RollupAddress common.Address
	client        arbbridge.ArbClient
	listeners     []rollup.ChainListener

	listenerAddChan chan rollup.ChainListener
	actionChan      chan func(*rollup.ChainObserver)
}

func CreateManager(
	ctx context.Context,
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	updateOpinion bool,
	clnt arbbridge.ArbClient,
	forceFreshStart bool,
	dbPrefix string,
	stressTest bool, // if true, generate artificial chaos to stress-test the implementation
) (*Manager, error) {
	if stressTest {
		clnt = NewStressTestClient(clnt, 10*time.Second)
	}
	man := &Manager{
		RollupAddress:   rollupAddr,
		client:          clnt,
		listenerAddChan: make(chan rollup.ChainListener, 10),
		actionChan:      make(chan func(*rollup.ChainObserver), 10),
	}
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			checkpointer := checkpointing.NewProductionCheckpointer(
				runCtx,
				rollupAddr,
				arbitrumCodeFilePath,
				big.NewInt(maxReorgDepth),
				dbPrefix,
				forceFreshStart,
			)

			var chain *rollup.ChainObserver

			watcher, err := clnt.NewRollupWatcher(rollupAddr)
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
				chain = chainObserverBuf.UnmarshalFromCheckpoint(runCtx, restoreCtx, checkpointer)
			} else {
				params, err := watcher.GetParams(ctx)
				if err != nil {
					log.Fatal(err)
				}
				blockId, err := watcher.GetCreationHeight(ctx)
				if err != nil {
					log.Fatal(err)
				}
				chain, err = rollup.NewChain(rollupAddr, checkpointer, params, updateOpinion, blockId)
				if err != nil {
					log.Fatal(err)
				}
			}

			log.Println("Starting validator from", chain.CurrentBlockId())

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

			current, err := clnt.CurrentBlockId(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			headersChan, err := clnt.SubscribeBlockHeaders(runCtx, chain.CurrentBlockId())
			if err != nil {
				blockId, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocks(big.NewInt(0)))
				if err != nil {
					panic(err)
				}
				log.Println("Error subscribing to block headers", chain.CurrentBlockId().HeaderHash, chain.CurrentBlockId().Height.AsInt(), blockId.HeaderHash, blockId.Height.AsInt(), err)

				cancelFunc()
				time.Sleep(2 * time.Second)
				continue
			}
			reachedHead := false
		runLoop:
			for {
				select {
				case maybeBlockId, ok := <-headersChan:
					if !ok {
						log.Println("Manager stopped receiving headers")
						break runLoop
					}
					if maybeBlockId.Err != nil {
						log.Println("Error getting new header", maybeBlockId.Err)
						break runLoop
					}

					blockId := maybeBlockId.BlockId

					if !reachedHead && blockId.Height.Cmp(current.Height) >= 0 {
						log.Println("Reached head")
						reachedHead = true
						chain.NowAtHead()
						log.Println("Now at head")
					}

					chain.NotifyNewBlock(blockId.Clone())

					events, err := watcher.GetEvents(runCtx, blockId)
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

func (man *Manager) ExecuteCall(messages value.TupleValue, maxSteps uint32) (*protocol.ExecutionAssertion, uint32) {
	retChan := make(chan struct {
		*protocol.ExecutionAssertion
		uint32
	}, 1)
	man.actionChan <- func(chain *rollup.ChainObserver) {
		mach := chain.LatestKnownValidMachine()
		latestTime := chain.CurrentBlockId().Height
		timeBounds := &protocol.TimeBoundsBlocks{latestTime, latestTime}
		go func() {
			assertion, numSteps := mach.ExecuteAssertion(maxSteps, timeBounds, messages)
			retChan <- struct {
				*protocol.ExecutionAssertion
				uint32
			}{assertion, numSteps}
		}()
	}
	ret := <-retChan
	return ret.ExecutionAssertion, ret.uint32
}

func (man *Manager) CurrentBlockId() *structures.BlockId {
	retChan := make(chan *structures.BlockId, 1)
	man.actionChan <- func(chain *rollup.ChainObserver) {
		retChan <- chain.CurrentBlockId()
	}
	return <-retChan
}
