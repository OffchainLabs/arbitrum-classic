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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type Observer struct {
}

func NewObserver(chain *ChainObserver, clnt *ethclient.Client) (*Observer, error) {
	rollup, err := ethbridge.NewRollupWatcher(chain.rollupAddr, clnt)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	outChan := make(chan ethbridge.Notification, 1024)
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
					chain.notifyNewBlockNumber(lastBlockNumberSeen)

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
	return &Observer{}, nil
}

func handleNotification(notification ethbridge.Notification, chain *ChainObserver) {
	chain.Lock()
	defer chain.Unlock()
	switch ev := notification.Event.(type) {
	case ethbridge.StakeCreatedEvent:
		chain.CreateStake(ev.Staker, ev.NodeHash, structures.RollupTimeFromBlockNum(notification.Header.Number))
		if chain.listener != nil && chain.listenForAddress == ev.Staker {
			chain.listener.Notify(
				&StakeCreatedChainEvent{
					ev.Staker,
					ev.NodeHash,
					structures.RollupTimeFromBlockNum(notification.Header.Number),
				},
			)
		}
	case ethbridge.ChallengeStartedEvent:
		_ = chain.NewChallenge(ev.ChallengeContract, ev.Asserter, ev.Challenger, ChallengeType(ev.ChallengeType))
		if chain.listener != nil && (chain.listenForAddress == ev.Asserter || chain.listenForAddress == ev.Challenger) {
			chain.listener.Notify(
				&ChallengeStartedChainEvent{
					ev.ChallengeContract,
					ev.Asserter,
					ev.Challenger,
					ChallengeType(ev.ChallengeType),
				},
			)
		}
	case ethbridge.ChallengeCompletedEvent:
		chain.ChallengeResolved(ev.ChallengeContract, ev.Winner, ev.Loser)
		if chain.listener != nil && (chain.listenForAddress == ev.Winner || chain.listenForAddress == ev.Loser) {
			chain.listener.Notify(
				&ChallengeCompletedChainEvent{
					ev.ChallengeContract,
					ev.Winner,
					ev.Loser,
				},
			)
		}
	case ethbridge.StakeRefundedEvent:
		chain.RemoveStake(ev.Staker)
		if chain.listener != nil && chain.listenForAddress == ev.Staker {
			chain.listener.Notify(&StakeRefundedChainEvent{ev.Staker})
		}
	case ethbridge.PrunedEvent:
		chain.PruneNodeByHash(ev.Leaf)
	case ethbridge.StakeMovedEvent:
		chain.MoveStake(ev.Staker, ev.Location)
		if chain.listener != nil && chain.listenForAddress == ev.Staker {
			chain.listener.Notify(&StakeMovedChainEvent{ev.Staker, ev.Location})
		}
	case ethbridge.AssertedEvent:
		chain.notifyAssert(
			ev.PrevLeafHash,
			ev.Params,
			ev.Claim,
		)
	case ethbridge.ConfirmedEvent:
		chain.ConfirmNode(ev.NodeHash)
	}
}

func calcSigHash(sig string) common.Hash {
	return crypto.Keccak256Hash([]byte(sig))
}
