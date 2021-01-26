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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"testing"
)

func testExecutionChallenge(t *testing.T, ctx context.Context, client ethutils.EthClient, asserterClient *ethbridge.EthArbAuthClient, challengerClient *ethbridge.EthArbAuthClient) {
	t.Parallel()

	mach := getTestMachine(t)
	challengeHash, assertion, inboxStack, numSteps := getExecutionChallengeData(mach)

	testChallengerCatchUp(t, ctx, client, asserterClient, challengerClient, valprotocol.InvalidExecutionChildType, challengeHash, func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
		return DefendExecutionClaim(
			ctx,
			client,
			challengeAddress,
			blockId,
			0,
			mach.Clone(),
			assertion,
			inboxStack,
			numSteps,
			4,
			StandardExecutionChallenge(),
		)
	}, func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
		return DefendExecutionClaim(
			ctx,
			client,
			challengeAddress,
			blockId,
			0,
			mach.Clone(),
			assertion,
			inboxStack,
			numSteps,
			4,
			ExecutionChallengeInfo{
				true,
				2,
				0,
			},
		)
	}, func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
		return ChallengeExecutionClaim(
			ctx,
			client,
			challengeAddress,
			blockId,
			0,
			inboxStack,
			numSteps,
			mach.Clone(),
			assertion.BeforeInboxHash,
			true,
			StandardExecutionChallenge(),
		)
	}, func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
		return ChallengeExecutionClaim(
			ctx,
			client,
			challengeAddress,
			blockId,
			0,
			inboxStack,
			numSteps,
			mach.Clone(),
			assertion.BeforeInboxHash,
			true,
			ExecutionChallengeInfo{
				true,
				2,
				0,
			},
		)
	}, testerAddress)
}

func getExecutionChallengeData(mach machine.Machine) (common.Hash, *valprotocol.ExecutionAssertionStub, *structures.MessageStack, uint64) {
	ms := structures.NewRandomMessageStack(1000)
	afterMachine := mach.Clone()
	assertion, _, numSteps := afterMachine.ExecuteAssertion(1000, true, ms.GetAllMessages(), true)
	stub := structures.NewExecutionAssertionStubFromWholeAssertion(assertion, common.Hash{}, ms)
	challengeHash := valprotocol.ExecutionDataHash(numSteps, stub)
	return challengeHash, stub, ms, numSteps
}
