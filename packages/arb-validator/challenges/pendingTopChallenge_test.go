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

package challenges

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func TestPendingTopChallenge(t *testing.T) {
	messageStack := structures.NewMessageStack()
	messageStack.DeliverMessage(value.NewInt64Value(0))
	messageStack.DeliverMessage(value.NewInt64Value(1))
	messageStack.DeliverMessage(value.NewInt64Value(2))
	messageStack.DeliverMessage(value.NewInt64Value(3))

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	topHash, err := messageStack.GetHashAtIndex(big.NewInt(3))
	if err != nil {
		t.Fatal(err)
	}
	challengeHash := structures.PendingTopChallengeDataHash(bottomHash, topHash, big.NewInt(3))

	if err := testChallenge(
		structures.InvalidPendingChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return DefendPendingTopClaim(
				client,
				challengeAddress,
				common.NewTimeBlocks(big.NewInt(0)),
				0,
				messageStack,
				bottomHash,
				topHash,
				2,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return ChallengePendingTopClaim(
				client,
				challengeAddress,
				common.NewTimeBlocks(big.NewInt(0)),
				0,
				messageStack,
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}
