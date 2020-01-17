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
	dbPrefix string,
	stressTest bool, // if true, generate artificial chaos to stress-test the implementation
) (*Manager, error) {
	man := &Manager{
		RollupAddress:   rollupAddr,
		client:          clnt,
		listenerAddChan: make(chan rollup.ChainListener, 10),
		actionChan:      make(chan func(*rollup.ChainObserver), 10),
	}
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			checkpointer := rollup.NewProductionCheckpointer(
				runCtx,
				rollupAddr,
				arbitrumCodeFilePath,
				big.NewInt(maxReorgDepth),
				dbPrefix,
				false,
			)

			latestBlockId, chainObserverBuf, restoreCtx := checkpointer.RestoreLatestState(clnt, rollupAddr, updateOpinion)
			watcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}
			if stressTest {
				watcher = NewStressTestWatcher(watcher, 30*time.Second)
			}
			chain := chainObserverBuf.UnmarshalFromCheckpoint(runCtx, restoreCtx, latestBlockId, watcher, checkpointer)

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

			reorgCtx, eventChan, err := arbbridge.HandleBlockchainNotifications(runCtx, latestBlockId, 0, watcher)
			if err != nil {
				log.Fatal(err)
			}

		runLoop:
			for {
				select {
				case <-reorgCtx.Done():
					log.Println("Reorg context done")
					break runLoop
				case listener := <-man.listenerAddChan:
					chain.AddListener(listener)
				case action := <-man.actionChan:
					action(chain)
				case event, ok := <-eventChan:
					if !ok {
						break runLoop
					}
					chainInfo := event.GetChainInfo()

					if chainInfo.BlockId.Height.Cmp(latestBlockId.Height) > 0 {
						chain.NotifyNewBlock(chainInfo.BlockId)
					}

					handleNotification(event, chain)
					latestBlockId = chainInfo.BlockId
				}
			}
			cancelFunc()

			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(15 * time.Second) // give time for things to settle, post-reorg, before restarting stuff
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

func handleNotification(event arbbridge.Event, chain *rollup.ChainObserver) {
	log.Printf("Handling event %T\n", event)
	chain.Lock()
	defer chain.Unlock()
	switch ev := event.(type) {
	case arbbridge.MessageDeliveredEvent:
		chain.MessageDelivered(ev)
	case arbbridge.StakeCreatedEvent:
		currentTime := common.TimeFromBlockNum(ev.BlockId.Height)
		chain.CreateStake(ev, currentTime)
	case arbbridge.ChallengeStartedEvent:
		chain.NewChallenge(ev)
	case arbbridge.ChallengeCompletedEvent:
		chain.ChallengeResolved(ev)
	case arbbridge.StakeRefundedEvent:
		chain.RemoveStake(ev)
	case arbbridge.PrunedEvent:
		chain.PruneLeaf(ev)
	case arbbridge.StakeMovedEvent:
		chain.MoveStake(ev)
	case arbbridge.AssertedEvent:
		err := chain.NotifyAssert(ev, ev.BlockId.Height, ev.TxHash)
		if err != nil {
			panic(err)
		}
	case arbbridge.ConfirmedEvent:
		chain.ConfirmNode(ev)
	}
}
