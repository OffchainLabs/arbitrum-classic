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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"log"
	"math/big"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
)

const (
	ValidEthBridgeVersion = "4"
)

type Manager struct {
	// listenersLock is should be held whenever listeners is accessed or set
	// and whenever activeChain is going to be set (but not when it is accessed)
	listenersLock sync.Mutex

	// validCallLock is held by the main loop whenever there is not a non-nil
	// activeChain running which has caught up to the current L1 block. It is
	// held by all other functions when they wish to access an update to date
	// chain state. This means that these functions will block when there isn't
	// an appropriate way to resolve them
	validCallLock sync.Mutex

	listeners   []rollup.ChainListener
	activeChain *rollup.ChainObserver

	// These variables are only written by the constructor
	RollupAddress common.Address
	checkpointer  checkpointing.RollupCheckpointer
}

const defaultMaxReorgDepth = 100

func CreateManager(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	aoFilePath string,
	dbPath string,
) (*Manager, error) {
	return CreateManagerAdvanced(
		ctx,
		rollupAddr,
		true,
		clnt,
		checkpointing.NewIndexedCheckpointer(
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
	checkpointer checkpointing.RollupCheckpointer,
) (*Manager, error) {
	if err := verifyArbChain(ctx, rollupAddr, clnt, checkpointer); err != nil {
		return nil, err
	}

	man := &Manager{
		RollupAddress: rollupAddr,
		checkpointer:  checkpointer,
	}
	man.validCallLock.Lock()
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

			chain, err := initializeChainObserver(
				runCtx,
				rollupAddr,
				updateOpinion,
				clnt,
				rollupWatcher,
				checkpointer,
			)
			if err != nil {
				log.Fatal(err)
			}

			man.listenersLock.Lock()
			man.activeChain = chain
			// Add manager's listeners
			for _, listener := range man.listeners {
				man.activeChain.AddListener(runCtx, listener)
			}
			man.listenersLock.Unlock()

			currentProcessedBlockId := man.activeChain.CurrentBlockId()
			time.Sleep(time.Second) // give time for things to settle, post-reorg, before restarting stuff

			log.Println("Starting validator from", currentProcessedBlockId)

			man.activeChain.RestartFromLatestValid(runCtx)

			man.activeChain.Start(runCtx)

			headersChan, err := clnt.SubscribeBlockHeaders(runCtx, currentProcessedBlockId)
			if err != nil {
				log.Println("Error subscribing to block headers", err)
				cancelFunc()
				time.Sleep(2 * time.Second)
				continue
			}

			caughtUpToL1 := false
		headerLoop:
			for maybeBlockId := range headersChan {
				if maybeBlockId.Err != nil {
					log.Println("Error getting new header", maybeBlockId.Err)
					break
				}

				blockId := maybeBlockId.BlockId
				timestamp := maybeBlockId.Timestamp

				currentOnChain, err := clnt.CurrentBlockId(runCtx)
				if err != nil {
					log.Fatal(err)
				}

				if !caughtUpToL1 && blockId.Height.Cmp(currentOnChain.Height) >= 0 {
					caughtUpToL1 = true
					man.activeChain.NowAtHead()
					log.Println("Now at head")
					man.validCallLock.Unlock()
				}

				man.activeChain.NotifyNewBlock(blockId.Clone())
				log.Print(man.activeChain.DebugString("== "))

				inboxEvents, err := inboxWatcher.GetEvents(runCtx, blockId, timestamp)
				if err != nil {
					log.Println("Manager hit error getting events", err)
					break
				}

				events, err := rollupWatcher.GetEvents(runCtx, blockId, timestamp)
				if err != nil {
					log.Println("Manager hit error getting events", err)
					break
				}

				for _, event := range inboxEvents {
					err := man.activeChain.HandleNotification(runCtx, event)
					if err != nil {
						log.Println("Manager hit error processing event", err)
						break headerLoop
					}
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

			}
		}
	}()

	return man, nil
}

func (man *Manager) AddListener(listener rollup.ChainListener) {
	man.listenersLock.Lock()
	man.listeners = append(man.listeners, listener)
	if man.activeChain != nil {
		man.activeChain.AddListener(context.Background(), listener)
	}
	man.listenersLock.Unlock()
}

func (man *Manager) GetLatestMachine() machine.Machine {
	man.validCallLock.Lock()
	defer man.validCallLock.Unlock()
	return man.activeChain.LatestKnownValidMachine()
}

func (man *Manager) GetLatestBlock() *common.BlockId {
	man.validCallLock.Lock()
	defer man.validCallLock.Unlock()
	return man.activeChain.CurrentBlockId()
}

func (man *Manager) GetPendingMachine() machine.Machine {
	man.validCallLock.Lock()
	defer man.validCallLock.Unlock()
	return man.activeChain.CurrentPendingMachine()
}

func (man *Manager) CurrentBlockId() *common.BlockId {
	man.validCallLock.Lock()
	defer man.validCallLock.Unlock()
	return man.activeChain.CurrentBlockId()
}

func (man *Manager) GetCheckpointer() checkpointing.RollupCheckpointer {
	return man.checkpointer
}

func verifyArbChain(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	checkpointer checkpointing.RollupCheckpointer,
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

	_, _, initialVMHash, err := watcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	initialMachine, err := checkpointer.GetInitialMachine()
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
		if err := checkpointer.RestoreLatestState(ctx, clnt, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext) error {
			chainObserverBuf := &rollup.ChainObserverBuf{}
			if err := proto.Unmarshal(chainObserverBytes, chainObserverBuf); err != nil {
				return err
			}
			var err error
			chain, err = chainObserverBuf.UnmarshalFromCheckpoint(restoreCtx, checkpointer)
			return err
		}); err == nil && chain != nil {
			return chain, nil
		}
	}

	log.Println("No valid checkpoints so starting from fresh state")
	params, err := watcher.GetParams(ctx)
	if err != nil {
		log.Fatal(err)
	}
	txHash, blockId, _, err := watcher.GetCreationInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return rollup.NewChain(rollupAddr, checkpointer, params, updateOpinion, blockId, txHash)
}
