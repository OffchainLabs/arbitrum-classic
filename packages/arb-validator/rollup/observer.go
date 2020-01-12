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

package rollup

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func CreateObserver(
	ctx context.Context,
	rollupAddr common.Address,
	checkpointer RollupCheckpointer,
	updateOpinion bool,
	startTime *common.TimeBlocks,
	clnt arbbridge.ArbClient,
) (*ChainObserver, error) {
	rollup, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}
	vmParams, err := rollup.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	chain, err := NewChain(
		ctx,
		rollupAddr,
		checkpointer,
		vmParams,
		updateOpinion,
		startTime,
	)
	if err != nil {
		return nil, err
	}

	outChan := make(chan arbbridge.Notification, 1024)
	errChan := make(chan error, 1024)
	if err := rollup.StartConnection(ctx, outChan, errChan); err != nil {
		return nil, err
	}

	go func() {
		lastBlockNumberSeen := big.NewInt(0)
		for {
			hitError := false
			select {
			case <-ctx.Done():
				break
			case notification, ok := <-outChan:
				if !ok {
					hitError = true
					break
				}
				if notification.Header.Number.Cmp(lastBlockNumberSeen) > 0 {
					lastBlockNumberSeen = notification.Header.Number
					blockHeaderHash := notification.Header.Root
					chain.notifyNewBlock(common.NewTimeBlocks(lastBlockNumberSeen), blockHeaderHash)
				}
				handleNotification(notification, chain)
			case <-errChan:
				hitError = true
			}

			if hitError {
				// Ignore error and try to reset connection
				for {
					if err := rollup.StartConnection(ctx, outChan, errChan); err == nil {
						break
					}
					log.Println("Error: Can't connect to blockchain")
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()
	return chain, nil
}

func handleNotification(notification arbbridge.Notification, chain *ChainObserver) {
	chain.Lock()
	defer chain.Unlock()
	switch ev := notification.Event.(type) {
	case arbbridge.MessageDeliveredEvent:
		chain.messageDelivered(ev)
	case arbbridge.StakeCreatedEvent:
		currentTime := structures.TimeFromBlockNum(common.NewTimeBlocks(notification.Header.Number))
		chain.createStake(ev, currentTime)
	case arbbridge.ChallengeStartedEvent:
		chain.newChallenge(ev)
	case arbbridge.ChallengeCompletedEvent:
		chain.challengeResolved(ev)
	case arbbridge.StakeRefundedEvent:
		chain.removeStake(ev)
	case arbbridge.PrunedEvent:
		chain.pruneLeaf(ev)
	case arbbridge.StakeMovedEvent:
		chain.moveStake(ev)
	case arbbridge.AssertedEvent:
		currentTime := common.NewTimeBlocks(notification.Header.Number)
		err := chain.notifyAssert(ev, currentTime, notification.TxHash)
		if err != nil {
			panic(err)
		}
	case arbbridge.ConfirmedEvent:
		chain.confirmNode(ev)
	}
}
