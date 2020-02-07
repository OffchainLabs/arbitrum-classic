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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func testInboxTopChallenge(t *testing.T) {
	t.Parallel()
	msg1 := message.DeliveredEth{
		Eth: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(532)),
		MessageNum: big.NewInt(1),
	}
	msg2 := message.DeliveredEth{
		Eth: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(532)),
		MessageNum: big.NewInt(2),
	}
	msg3 := message.DeliveredEth{
		Eth: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(532)),
		MessageNum: big.NewInt(3),
	}
	msg4 := message.DeliveredEth{
		Eth: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(532)),
		MessageNum: big.NewInt(4),
	}
	messageStack := structures.NewMessageStack()
	messageStack.DeliverMessage(msg1)
	messageStack.DeliverMessage(msg2)
	messageStack.DeliverMessage(msg3)
	messageStack.DeliverMessage(msg4)

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	messageCount := big.NewInt(3)
	topHash, err := messageStack.GetHashAtIndex(messageCount)
	if err != nil {
		t.Fatal(err)
	}
	challengeHash := structures.InboxTopChallengeDataHash(bottomHash, topHash, big.NewInt(3))

	if err := testChallenge(
		valprotocol.InvalidInboxTopChildType,
		challengeHash,
		"ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39",
		"979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76",
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendInboxTopClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				messageStack,
				bottomHash,
				messageCount,
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
	); err != nil {
		t.Fatal(err)
	}
}
