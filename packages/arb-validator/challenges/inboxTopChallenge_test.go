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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func testInboxTopChallenge(
	t *testing.T,
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
) {
	t.Parallel()

	messageStack := getInboxMsgStack()
	count := new(big.Int).Sub(messageStack.TopCount(), big.NewInt(1))
	bottomHash, challengeHash := getChallengeData(t, messageStack, count)

	if err := testChallenge(
		client,
		asserter,
		challenger,
		valprotocol.InvalidInboxTopChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendInboxTopClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				bottomHash,
				count,
				2,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return ChallengeInboxTopClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				true,
			)
		},
		testerAddress,
	); err != nil {
		t.Fatal(err)
	}
}

func getChallengeData(t *testing.T, messageStack *structures.MessageStack, messageCount *big.Int) (common.Hash, common.Hash) {
	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	topHash, err := messageStack.GetHashAtIndex(messageCount)
	if err != nil {
		t.Fatal(err)
	}
	challengeHash := valprotocol.InboxTopChallengeDataHash(bottomHash, topHash, messageCount)

	return bottomHash, challengeHash
}

func getInboxMsgStack() *structures.MessageStack {
	messageStack := structures.NewMessageStack()
	messageStack.DeliverMessage(inbox.NewRandomInboxMessage())
	messageStack.DeliverMessage(inbox.NewRandomInboxMessage())
	messageStack.DeliverMessage(inbox.NewRandomInboxMessage())
	messageStack.DeliverMessage(inbox.NewRandomInboxMessage())

	return messageStack
}
