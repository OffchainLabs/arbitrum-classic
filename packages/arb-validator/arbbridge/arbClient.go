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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type MaybeBlockID struct {
	BlockID *structures.BlockID
	Err     error
}

type ArbClient interface {
	SubscribeBlockHeaders(ctx context.Context, startBlockID *structures.BlockID) (<-chan MaybeBlockID, error)

	NewArbFactoryWatcher(address common.Address) (ArbFactoryWatcher, error)
	NewRollupWatcher(address common.Address) (ArbRollupWatcher, error)
	NewPendingInboxWatcher(address common.Address, rollupAddress common.Address) (PendingInboxWatcher, error)
	NewExecutionChallengeWatcher(address common.Address) (ExecutionChallengeWatcher, error)
	NewMessagesChallengeWatcher(address common.Address) (MessagesChallengeWatcher, error)
	NewPendingTopChallengeWatcher(address common.Address) (PendingTopChallengeWatcher, error)
	NewOneStepProof(address common.Address) (OneStepProof, error)
	CurrentBlockID(ctx context.Context) (*structures.BlockID, error)
	BlockIDForHeight(ctx context.Context, height *common.TimeBlocks) (*structures.BlockID, error)
}

type ArbAuthClient interface {
	ArbClient
	Address() common.Address
	NewArbFactory(address common.Address) (ArbFactory, error)
	NewRollup(address common.Address) (ArbRollup, error)
	NewPendingInbox(address common.Address) (PendingInbox, error)
	NewChallengeFactory(address common.Address) (ChallengeFactory, error)
	NewExecutionChallenge(address common.Address) (ExecutionChallenge, error)
	NewMessagesChallenge(address common.Address) (MessagesChallenge, error)
	NewPendingTopChallenge(address common.Address) (PendingTopChallenge, error)
}
