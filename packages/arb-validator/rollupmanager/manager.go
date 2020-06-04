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
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"log"
	"math/big"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
)

const (
	ValidEthBridgeVersion = "2"
)

type Manager struct {
	// listenersLock is locked when writing listeners or activeChain
	listenersLock sync.Mutex
	// validCallLock is locked when there is not a valid chain caught up to head
	validCallLock sync.Mutex
	RollupAddress common.Address
	listeners     []rollup.ChainListener
	activeChain   *rollup.ChainObserver
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
		checkpointing.NewIndexedCheckpointerFactory(
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
	if err := verifyArbChain(ctx, rollupAddr, clnt, ckpFac); err != nil {
		return nil, err
	}

	man := &Manager{
		RollupAddress: rollupAddr,
	}
	man.validCallLock.Lock()
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			checkpointer := ckpFac.New(runCtx)

			watcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			chain, err := initializeChainObserver(
				runCtx,
				rollupAddr,
				updateOpinion,
				clnt,
				watcher,
				checkpointer,
			)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Starting validator from", man.activeChain.CurrentBlockId())

			man.listenersLock.Lock()
			man.activeChain = chain
			// Add manager's listeners
			for _, listener := range man.listeners {
				man.activeChain.AddListener(listener)
			}
			man.listenersLock.Unlock()

			chain.RestartFromLatestValid(runCtx)

			man.activeChain.Start(runCtx)

			current, err := clnt.CurrentBlockId(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			headersChan, err := clnt.SubscribeBlockHeaders(runCtx, man.activeChain.CurrentBlockId())
			if err != nil {
				log.Println("Error subscribing to block headers", err)
				cancelFunc()
				time.Sleep(2 * time.Second)
				continue
			}

			reachedHead := false
		headerLoop:
			for maybeBlockId := range headersChan {
				if maybeBlockId.Err != nil {
					log.Println("Error getting new header", maybeBlockId.Err)
					break
				}

				blockId := maybeBlockId.BlockId
				timestamp := maybeBlockId.Timestamp

				if !reachedHead && blockId.Height.Cmp(current.Height) >= 0 {
					reachedHead = true
					man.activeChain.NowAtHead()
					log.Println("Now at head")
					man.validCallLock.Unlock()
				}

				man.activeChain.NotifyNewBlock(blockId.Clone())
				log.Print(man.activeChain.DebugString("== "))

				events, err := watcher.GetEvents(runCtx, blockId, timestamp)
				if err != nil {
					log.Println("Manager hit error getting events", err)
					break
				}
				for _, event := range events {
					err := man.activeChain.HandleNotification(runCtx, event)
					if err != nil {
						log.Println("Manager hit error processing event", err)
						break headerLoop
					}
				}
			}

			man.validCallLock.Lock()

			man.listenersLock.Lock()
			man.activeChain = nil
			man.listenersLock.Unlock()

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
	man.listenersLock.Lock()
	man.listeners = append(man.listeners, listener)
	if man.activeChain != nil {
		man.activeChain.AddListener(listener)
	}
	man.listenersLock.Unlock()
}

func (man *Manager) ExecuteCall(messages value.TupleValue, maxTime time.Duration) (*protocol.ExecutionAssertion, uint64) {
	man.validCallLock.Lock()
	mach := man.activeChain.LatestKnownValidMachine()
	latestBlock := man.activeChain.CurrentBlockId().Height
	man.validCallLock.Unlock()
	latestTime := big.NewInt(time.Now().Unix())
	timeBounds := &protocol.TimeBounds{latestBlock, latestBlock, latestTime, latestTime}
	assertion, numSteps := mach.ExecuteAssertion(
		// Call execution is only limited by wall time, so use a massive max steps as an approximation to infinity
		10000000000000000,
		timeBounds,
		messages,
		maxTime,
	)
	return assertion, numSteps
}

func (man *Manager) CurrentBlockId() *common.BlockId {
	man.validCallLock.Lock()
	defer man.validCallLock.Unlock()
	return man.activeChain.CurrentBlockId()
}

func verifyArbChain(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	ckpFac checkpointing.RollupCheckpointerFactory,
) error {
	watcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return err
	}

	ethbridgeVersion, err := watcher.GetVersion(ctx)
	if err != nil {
		return err
	}

	if ethbridgeVersion != ValidEthBridgeVersion {
		return fmt.Errorf("VM has EthBridge version %v, but validator implements version %v."+
			" To find a validator version which supports your EthBridge, visit "+
			"https://offchainlabs.com/ethbridge-version-support",
			ethbridgeVersion, ValidEthBridgeVersion)
	}

	_, initialVMHash, err := watcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	initialMachine, err := ckpFac.New(ctx).GetInitialMachine()
	if err != nil {
		return err
	}

	if initialMachine.Hash() != initialVMHash {
		return errors.New("ArbChain was initialized with different VM")
	}
	return nil
}

func initializeChainObserver(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ChainTimeGetter,
	watcher arbbridge.ArbRollupWatcher,
	checkpointer checkpointing.RollupCheckpointer,
) (*rollup.ChainObserver, error) {
	if checkpointer.HasCheckpointedState() {
		var chain *rollup.ChainObserver
		err := checkpointer.RestoreLatestState(ctx, clnt, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext) error {
			chainObserverBuf := &rollup.ChainObserverBuf{}
			if err := proto.Unmarshal(chainObserverBytes, chainObserverBuf); err != nil {
				return err
			}
			var err error
			chain, err = chainObserverBuf.UnmarshalFromCheckpoint(ctx, restoreCtx, checkpointer)
			return err
		})
		return chain, err
	} else {
		params, err := watcher.GetParams(ctx)
		if err != nil {
			log.Fatal(err)
		}
		blockId, _, err := watcher.GetCreationInfo(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return rollup.NewChain(rollupAddr, checkpointer, params, updateOpinion, blockId)
	}
}
