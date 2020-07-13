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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"testing"
)

func testExecutionChallenge(
	t *testing.T,
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
) {
	t.Parallel()

	mach := getTestMachine(t)
	challengeHash, precondition, numSteps := getExecutionChallengeData(mach)

	if err := testChallengerCatchUp(
		client,
		asserter,
		challenger,
		valprotocol.InvalidExecutionChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendExecutionClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				precondition,
				mach.Clone(),
				numSteps,
				4,
				StandardExecutionChallenge(),
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return DefendExecutionClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				precondition,
				mach.Clone(),
				numSteps,
				4,
				ExecutionChallengeInfo{
					true,
					2,
					0,
				},
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return ChallengeExecutionClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				precondition,
				mach.Clone(),
				true,
				StandardExecutionChallenge(),
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient, blockId *common.BlockId) (ChallengeState, error) {
			return ChallengeExecutionClaim(
				context.Background(),
				client,
				challengeAddress,
				blockId,
				0,
				precondition,
				mach.Clone(),
				true,
				ExecutionChallengeInfo{
					true,
					2,
					0,
				},
			)
		},
		testerAddress,
	); err != nil {
		t.Fatal(err)
	}
}

func getExecutionChallengeData(mach machine.Machine) (common.Hash, *valprotocol.Precondition, uint64) {
	afterMachine := mach.Clone()
	precondition := valprotocol.NewPrecondition(mach.Hash(), value.NewEmptyTuple())
	assertion, numSteps := afterMachine.ExecuteAssertion(1000, value.NewEmptyTuple(), 0)

	challengeHash := valprotocol.ExecutionDataHash(
		numSteps,
		precondition.BeforeHash,
		precondition.BeforeInbox.Hash(),
		valprotocol.NewExecutionAssertionStubFromAssertion(assertion).Hash(),
	)

	return challengeHash, precondition, numSteps
}
