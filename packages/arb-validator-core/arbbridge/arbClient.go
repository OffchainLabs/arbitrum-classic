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

package arbbridge

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type MaybeBlockId struct {
	BlockId   *common.BlockId
	Timestamp *big.Int
	Err       error
}

type ChainTimeGetter interface {
	CurrentBlockId(ctx context.Context) (*common.BlockId, error)
	BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error)
}

type ArbClient interface {
	ChainTimeGetter
	SubscribeBlockHeaders(ctx context.Context, startBlockId *common.BlockId) (<-chan MaybeBlockId, error)
	SubscribeBlockHeadersAfter(ctx context.Context, prevBlockId *common.BlockId) <-chan MaybeBlockId

	NewArbFactoryWatcher(address common.Address) (ArbFactoryWatcher, error)
	NewRollupWatcher(address common.Address) (ArbRollupWatcher, error)
	NewGlobalInboxWatcher(address common.Address, rollupAddress common.Address) (GlobalInboxWatcher, error)
	NewExecutionChallengeWatcher(address common.Address) (ExecutionChallengeWatcher, error)
	NewInboxTopChallengeWatcher(address common.Address) (InboxTopChallengeWatcher, error)

	GetBalance(ctx context.Context, account common.Address) (*big.Int, error)
}

type ArbAuthClient interface {
	ArbClient
	Address() common.Address
	NewArbFactory(address common.Address) (ArbFactory, error)
	NewRollup(address common.Address) (ArbRollup, error)
	NewGlobalInbox(address common.Address, rollupAddress common.Address) (GlobalInbox, error)
	NewChallengeFactory(address common.Address) (ChallengeFactory, error)
	NewExecutionChallenge(address common.Address) (ExecutionChallenge, error)
	NewInboxTopChallenge(address common.Address) (InboxTopChallenge, error)
}

func WaitForNonZeroBalance(ctx context.Context, client ArbClient, account common.Address) error {
	balance, err := client.GetBalance(ctx, account)
	if err != nil {
		return err
	}
	if balance.Cmp(big.NewInt(0)) > 0 {
		return nil
	}
	log.Println("Waiting for account", account, "to receive funds")
	timer := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ctx.Done():
			return errors.New("timed out waiting for balance")
		case <-timer.C:
			balance, err := client.GetBalance(ctx, account)
			if err != nil {
				return err
			}
			if balance.Cmp(big.NewInt(0)) > 0 {
				return nil
			}
		}
	}
}
