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
	"context"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func testMessagesChallenge(t *testing.T) {
	t.Parallel()
	messageStack := structures.NewMessageStack()
	for i := int64(0); i < 8; i++ {
		messageStack.DeliverMessage(value.NewInt64Value(i))
	}
	beforePending, err := messageStack.GetHashAtIndex(big.NewInt(2))
	if err != nil {
		t.Fatal(err)
	}
	afterPending, err := messageStack.GetHashAtIndex(big.NewInt(6))
	if err != nil {
		t.Fatal(err)
	}

	substack, err := messageStack.Substack(beforePending, afterPending)
	if err != nil {
		t.Fatal(err)
	}

	importedMessages := substack.GetTopHash()
	challengeHash := structures.MessageChallengeDataHash(
		beforePending,
		afterPending,
		value.NewEmptyTuple().Hash(),
		importedMessages,
		big.NewInt(4),
	)

	if err := testChallenge(
		structures.InvalidMessagesChildType,
		challengeHash,
		"d26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e",
		"af5c2984cb1e2f668ae3fd5bbfe0471f68417efd012493538dcd42692299155b",
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *structures.BlockId) (ChallengeState, error) {
			return DefendMessagesClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				beforePending,
				afterPending,
				importedMessages,
				2,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *structures.BlockId) (ChallengeState, error) {
			return ChallengeMessagesClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				beforePending,
				afterPending,
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}
