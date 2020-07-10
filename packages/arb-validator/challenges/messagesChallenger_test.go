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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
)

func testMessagesChallenge(t *testing.T) {
	t.Parallel()
	messageStack := getMsgStack()
	messageCount := uint64(4)
	startIndex := big.NewInt(2)

	beforeInbox, challengeHash := getMsgChallengeData(
		t,
		messageStack,
		startIndex,
		messageCount)

	if err := testChallenge(
		valprotocol.InvalidMessagesChildType,
		challengeHash,
		"d26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e",
		"af5c2984cb1e2f668ae3fd5bbfe0471f68417efd012493538dcd42692299155b",
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendMessagesClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				beforeInbox,
				new(big.Int).SetUint64(messageCount),
				2,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return ChallengeMessagesClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				beforeInbox,
				new(big.Int).SetUint64(messageCount),
				true,
			)
		},
		testerAddress,
	); err != nil {
		t.Fatal(err)
	}
}

func getMsgChallengeData(
	t *testing.T,
	messageStack *structures.MessageStack,
	startIndex *big.Int,
	msgCount uint64,
) (common.Hash, common.Hash) {

	beforeInbox, err := messageStack.GetHashAtIndex(startIndex)
	if err != nil {
		t.Fatal(err)
	}

	startIndex = startIndex.Add(startIndex, new(big.Int).SetUint64(msgCount))
	afterInbox, err := messageStack.GetHashAtIndex(startIndex)
	if err != nil {
		t.Fatal(err)
	}

	inbox, err := messageStack.GenerateVMInbox(beforeInbox, msgCount)
	if err != nil {
		t.Fatal(err)
	}

	importedMessages := inbox.Hash().Hash()
	challengeHash := valprotocol.MessageChallengeDataHash(
		beforeInbox,
		afterInbox,
		value.NewEmptyTuple().Hash(),
		importedMessages,
		big.NewInt(4),
	)
	return beforeInbox, challengeHash
}

func getMsgStack() *structures.MessageStack {
	messageStack := structures.NewMessageStack()
	for i := int64(0); i < 8; i++ {
		msg := message.NewRandomInboxMessage(message.NewRandomEth())
		messageStack.DeliverMessage(msg)
	}
	return messageStack
}
