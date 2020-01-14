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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

const (
	maxReorgDepth = 100
)

type Manager struct {
	rollupAddr common.Address
	client     arbbridge.ArbClient
}

func CreateManager(
	ctx context.Context,
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	updateOpinion bool,
	clnt arbbridge.ArbClient,
) (*Manager, error) {
	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			reorgCtx, cancelFunc := context.WithCancel(ctx)

			checkpointer := rollup.NewProductionCheckpointer(
				reorgCtx,
				rollupAddr,
				arbitrumCodeFilePath,
				big.NewInt(maxReorgDepth),
				false,
			)

			latestBlockId, chainObserverBuf, restoreCtx := checkpointer.RestoreLatestState(clnt, rollupAddr, true)
			watcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}
			chain := chainObserverBuf.UnmarshalFromCheckpoint(reorgCtx, restoreCtx, watcher)

			fmt.Println("Starting connection")

			outChan := make(chan arbbridge.Notification, 1024)
			errChan := make(chan error, 1024)
			if err := rollupWatcher.StartConnection(ctx, latestBlockId.Height, 0, errChan, outChan); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Started connection")

			chain.Start(reorgCtx)

			go func() {
				latestLogIndex := uint(0)
				for {
					hitError := false
					select {
					case <-ctx.Done():
						return
					case notification, ok := <-outChan:
						if !ok {
							hitError = true
							break
						}
						switch notification.BlockId.Height.Cmp(latestBlockId.Height) {
						case -1:
							// reorg
							cancelFunc()
							return
						case 0:
							if !notification.BlockId.HeaderHash.Equals(latestBlockId.HeaderHash) {
								// reorg
								cancelFunc()
								return
							}
							if notification.LogIndex > latestLogIndex {
								latestLogIndex = notification.LogIndex
								handleNotification(notification, chain)
							}
						case 1:
							latestBlockId = notification.BlockId
							latestLogIndex = notification.LogIndex
							chain.NotifyNewBlock(notification.BlockId)
							handleNotification(notification, chain)
						}
					case <-errChan:
						hitError = true
					}

					if hitError {
						// Ignore error and try to reset connection
						for {
							if err := rollupWatcher.StartConnection(ctx, latestBlockId.Height, latestLogIndex+1, errChan, outChan); err == nil {
								break
							}
							log.Println("Error: Can't connect to blockchain")
							time.Sleep(5 * time.Second)
						}
					}
				}
			}()

			<-reorgCtx.Done()
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(15 * time.Second) // give time for things to settle, post-reorg, before restarting stuff
			}
		}
	}()

	return &Manager{rollupAddr, clnt}, nil
}

func handleNotification(notification arbbridge.Notification, chain *rollup.ChainObserver) {
	chain.Lock()
	defer chain.Unlock()
	switch ev := notification.Event.(type) {
	case arbbridge.MessageDeliveredEvent:
		chain.MessageDelivered(ev)
	case arbbridge.StakeCreatedEvent:
		currentTime := common.TimeFromBlockNum(notification.BlockId.Height)
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
		err := chain.NotifyAssert(ev, notification.BlockId.Height, notification.TxHash)
		if err != nil {
			panic(err)
		}
	case arbbridge.ConfirmedEvent:
		chain.ConfirmNode(ev)
	}
}
