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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func testMessagesChallenge(
	t *testing.T,
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
) {
	t.Parallel()
	testMessages := getTestMessages()
	messageStack := makeMessageStack(testMessages)
	messageCount := uint64(4)
	startIndex := big.NewInt(2)

	beforeInboxTop, afterInboxTop, challengeHash := getMsgChallengeData(
		t,
		testMessages,
		startIndex,
		messageCount,
	)

	testChallenge(
		t,
		client,
		asserter,
		challenger,
		valprotocol.InvalidMessagesChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendMessagesClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				beforeInboxTop,
				afterInboxTop,
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
				beforeInboxTop,
				new(big.Int).SetUint64(messageCount),
				true,
			)
		},
		testerAddress,
	)
}

func getMsgChallengeData(
	t *testing.T,
	messages []inbox.InboxMessage,
	startIndex *big.Int,
	msgCount uint64,
) (common.Hash, common.Hash, common.Hash) {
	messageStack := makeMessageStack(messages)

	bottomInboxHash, err := messageStack.GetHashAtIndex(startIndex)
	if err != nil {
		t.Fatal(err)
	}

	endIndex := new(big.Int).Add(startIndex, new(big.Int).SetUint64(msgCount))
	topInboxHash, err := messageStack.GetHashAtIndex(endIndex)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Message stack:")
	stackHashes := messageStack.GetAllHashes()
	for i := range messages {
		t.Log("stack:", stackHashes[len(stackHashes)-1-i])
		t.Log("message:", messages[len(messages)-1-i].AsValue().Hash())
	}
	t.Log(stackHashes[0])

	t.Log("topHash", topInboxHash)
	t.Log("bottomHash", bottomInboxHash)

	vmInbox, err := messageStack.GenerateVMInbox(bottomInboxHash, msgCount)
	if err != nil {
		t.Fatal(err)
	}

	calculatedInboxMessages := vmInbox.Messages()
	if uint64(len(calculatedInboxMessages)) != msgCount {
		t.Fatal("unexpected vm inbox message count", len(calculatedInboxMessages), "instead of", msgCount)
	}
	start := startIndex.Uint64()
	for i, msg := range messages[start : start+msgCount] {
		if !msg.Equals(calculatedInboxMessages[i]) {
			t.Fatal("generated vm inbox had bad contents")
		}

	}

	challengeHash := valprotocol.MessageChallengeDataHash(
		topInboxHash,
		bottomInboxHash,
		value.NewEmptyTuple().Hash(),
		vmInbox.Hash().Hash(),
		big.NewInt(4),
	)
	return bottomInboxHash, topInboxHash, challengeHash
}

func makeMessageStack(messages []inbox.InboxMessage) *structures.MessageStack {
	messageStack := structures.NewMessageStack()
	for _, msg := range messages {
		messageStack.DeliverMessage(msg)
	}
	return messageStack
}

func getTestMessages() []inbox.InboxMessage {
	messages := make([]inbox.InboxMessage, 0, 8)
	for i := int64(0); i < 8; i++ {
		messages = append(messages, inbox.NewRandomInboxMessage())
	}
	return messages
}
