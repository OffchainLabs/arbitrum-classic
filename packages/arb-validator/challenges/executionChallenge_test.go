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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

func TestExecution(t *testing.T) {
	contract := "../contract.ao"

	mach, err := loader.LoadMachineFromFile(contract, true, "test")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := &protocol.TimeBoundsBlocks{
		common.NewTimeBlocks(big.NewInt(100)),
		common.NewTimeBlocks(big.NewInt(200)),
	}
	afterMachine := mach.Clone()
	precondition := valprotocol.NewPrecondition(mach.Hash(), timeBounds, value.NewEmptyTuple())
	assertion, numSteps := afterMachine.ExecuteAssertion(1000, timeBounds, value.NewEmptyTuple())

	challengeHash := structures.ExecutionDataHash(
		numSteps,
		precondition.Hash(),
		valprotocol.NewExecutionAssertionStubFromAssertion(assertion).Hash(),
	)

	if err := testChallenge(
		structures.InvalidExecutionChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return DefendExecutionClaim(
				client,
				challengeAddress,
				precondition,
				mach.Clone(),
				numSteps,
				2,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return ChallengeExecutionClaim(
				client,
				challengeAddress,
				precondition,
				mach.Clone(),
				true,
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}
