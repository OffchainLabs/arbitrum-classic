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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainobserver"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
)

const (
	ValidEthBridgeVersion = "develop"
)

var errNoActiveChain = errors.New("validator has no active chain")
var notAtHead = errors.New("validator is catching up to latest")

type reorgCache struct {
	latestValidMachine    machine.Machine
	currentPendingMachine machine.Machine
	currentBlockId        *common.BlockId
}

type Manager struct {
	// The mutex should be held whenever listeres or reorgCache are accessed or
	// set and whenever activeChain is going to be set (but not when it is accessed)
	sync.Mutex

	listeners   []chainlistener.ChainListener
	activeChain *chainobserver.ChainObserver
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
}

const defaultMaxReorgDepth = 100

const assumedValidThreshold = 2

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
			dbPath,
			big.NewInt(defaultMaxReorgDepth),
			false,
		),
		aoFilePath,
	)
}

func CreateManagerAdvanced(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ArbClient,
	checkpointer checkpointing.RollupCheckpointer,
	aoFilePath string,
) (*Manager, error) {
	if !checkpointer.Initialized() {
		if err := checkpointer.Initialize(aoFilePath); err != nil {
			return nil, err
		}
	}
	if err := verifyArbChain(ctx, rollupAddr, clnt, checkpointer); err != nil {
		return nil, err
	}

	man := &Manager{
		RollupAddress: rollupAddr,
		checkpointer:  checkpointer,
	}
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

			chain, err := chainobserver.NewChainObserver(
				runCtx,
				rollupAddr,
				updateOpinion,
				clnt,
				rollupWatcher,
				checkpointer,
				assumedValidThreshold,
			)
			if err != nil {
				log.Fatal(err)
			}

			man.Lock()
			man.activeChain = chain
			// Add manager's listeners
			for _, listener := range man.listeners {
				man.activeChain.AddListener(runCtx, listener)
			}
			man.Unlock()

			currentProcessedBlockId := man.activeChain.CurrentBlockId()
			time.Sleep(time.Second) // give time for things to settle, post-reorg, before restarting stuff

			log.Println("Starting validator from", currentProcessedBlockId)

			man.activeChain.RestartFromLatestValid(runCtx)

			man.activeChain.Start(runCtx)

			caughtUpToL1 := false

			err = func() error {
				// If the local chain is significantly behind the L1, catch up
				// more efficiently. We process `MaxReorgHeight` blocks at a
				// time up to `MaxReorgHeight` blocks before the current head
				// and and assume that no reorg will occur affecting the blocks
				// we are processing
				maxReorg := checkpointer.MaxReorgHeight()
				for {
					if err := man.activeChain.UpdateAssumedValidBlock(runCtx, clnt, assumedValidThreshold); err != nil {
						return err
					}
					currentProcessedBlockId := man.activeChain.CurrentBlockId()
					currentLocalHeight := currentProcessedBlockId.Height.AsInt()

					currentOnChain, err := clnt.CurrentBlockId(runCtx)
					if err != nil {
						return err
					}
					currentL1Height := currentOnChain.Height.AsInt()

					fastCatchupEndHeight := new(big.Int).Sub(currentL1Height, maxReorg)
					if currentLocalHeight.Cmp(fastCatchupEndHeight) >= 0 {
						break
					}

					fetchSize := new(big.Int).Sub(fastCatchupEndHeight, currentLocalHeight)
					if fetchSize.Cmp(big.NewInt(1)) <= 0 {
						break
					}
					if fetchSize.Cmp(maxReorg) >= 0 {
						fetchSize = maxReorg
					}
					fetchEnd := new(big.Int).Add(currentLocalHeight, fetchSize)
					fetchEnd = fetchEnd.Sub(fetchEnd, big.NewInt(1))

					log.Println("Getting events between", currentLocalHeight, "and", fetchEnd)
					inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(runCtx, currentLocalHeight, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}
					inboxEvents := make([]arbbridge.Event, 0, len(inboxDeliveredEvents))
					for _, ev := range inboxDeliveredEvents {
						inboxEvents = append(inboxEvents, ev)
					}

					events, err := rollupWatcher.GetAllEvents(runCtx, currentLocalHeight, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}
					allEvents := arbbridge.MergeEventsUnsafe(inboxEvents, events)
					for _, ev := range allEvents {
						blockId := ev.GetChainInfo().BlockId
						if blockId.Height.AsInt().Cmp(currentLocalHeight) > 0 {
							man.activeChain.NotifyNewBlock(blockId.Clone())
							currentLocalHeight = blockId.Height.AsInt()
						}
						err := man.activeChain.HandleNotification(runCtx, ev)
						if err != nil {
							return errors2.Wrap(err, "Manager hit error processing event during fast catchup")
						}
					}
					if fetchEnd.Cmp(currentLocalHeight) > 0 {
						endBlockId, err := clnt.BlockIdForHeight(runCtx, common.NewTimeBlocks(fetchEnd))
						if err != nil {
							return err
						}
						man.activeChain.NotifyNewBlock(endBlockId)
					}
				}

				headersChan, err := clnt.SubscribeBlockHeaders(runCtx, man.activeChain.CurrentBlockId())
				if err != nil {
					return errors2.Wrap(err, "Error subscribing to block headers")
				}

				lastDebugPrint := time.Now()
				for maybeBlockId := range headersChan {
					if maybeBlockId.Err != nil {
						return errors2.Wrap(maybeBlockId.Err, "Error getting new header")
					}

					if err := man.activeChain.UpdateAssumedValidBlock(runCtx, clnt, assumedValidThreshold); err != nil {
						return err
					}

					blockId := maybeBlockId.BlockId
					timestamp := maybeBlockId.Timestamp

					currentOnChain, err := clnt.CurrentBlockId(runCtx)
					if err != nil {
						return err
					}

					if !caughtUpToL1 && blockId.Height.Cmp(currentOnChain.Height) >= 0 {
						caughtUpToL1 = true
						man.activeChain.NowAtHead()
						log.Println("Now at head")
						man.Lock()
						man.reorgCache = nil
						man.Unlock()
					}

					man.activeChain.NotifyNewBlock(blockId.Clone())

					if caughtUpToL1 || time.Since(lastDebugPrint).Seconds() > 5 {
						log.Print(man.activeChain.DebugString("== "))
						lastDebugPrint = time.Now()
					}

					inboxEvents, err := inboxWatcher.GetEvents(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "Manager hit error getting inbox events with block %v", blockId)
					}

					events, err := rollupWatcher.GetEvents(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "Manager hit error getting rollup events with block %v", blockId)
					}

					for _, event := range arbbridge.MergeEventsUnsafe(inboxEvents, events) {
						err := man.activeChain.HandleNotification(runCtx, event)
						if err != nil {
							return errors2.Wrap(err, "Manager hit error processing event")
						}
					}
				}
				return nil
			}()

			if err != nil {
				log.Println(err)
			}

			man.Lock()
			man.reorgCache = &reorgCache{
				latestValidMachine:    man.activeChain.LatestKnownValidMachine(),
				currentPendingMachine: man.activeChain.CurrentPendingMachine(),
				currentBlockId:        man.activeChain.CurrentBlockId(),
			}
			man.activeChain = nil
			man.Unlock()

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

func (man *Manager) AddListener(listener chainlistener.ChainListener) {
	man.Lock()
	defer man.Unlock()
	man.listeners = append(man.listeners, listener)
	if man.activeChain != nil {
		man.activeChain.AddListener(context.Background(), listener)
	}
}

func (man *Manager) GetLatestMachine() (machine.Machine, error) {
	man.Lock()
	defer man.Unlock()
	if man.reorgCache != nil {
		return man.reorgCache.latestValidMachine.Clone(), nil
	}
	if man.activeChain == nil {
		return nil, errNoActiveChain
	}
	if man.activeChain.IsAtHead() {
		return man.activeChain.LatestKnownValidMachine(), nil
	}
	return nil, notAtHead
}

func (man *Manager) GetPendingMachine() (machine.Machine, error) {
	man.Lock()
	defer man.Unlock()
	if man.reorgCache != nil {
		return man.reorgCache.currentPendingMachine.Clone(), nil
	}
	if man.activeChain == nil {
		return nil, errNoActiveChain
	}
	if man.activeChain.IsAtHead() {
		return man.activeChain.CurrentPendingMachine(), nil
	}
	return nil, notAtHead
}

func (man *Manager) CurrentBlockId() (*common.BlockId, error) {
	man.Lock()
	defer man.Unlock()
	if man.reorgCache != nil {
		return man.reorgCache.currentBlockId.Clone(), nil
	}
	if man.activeChain == nil {
		return nil, errNoActiveChain
	}
	if man.activeChain.IsAtHead() {
		return man.activeChain.CurrentBlockId(), nil
	}
	return nil, notAtHead
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

	initialMachineHash := initialMachine.Hash()
	if initialMachineHash != initialVMHash {
		return fmt.Errorf("ArbChain was initialized with VM with hash %v, but local validator has VM with hash %v", initialVMHash, initialMachineHash)
	}
	return nil
}
