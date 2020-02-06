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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

func testExecutionChallenge(t *testing.T) {
	t.Parallel()
	contract := "../contract.ao"

	mach, err := loader.LoadMachineFromFile(contract, true, "test")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := &protocol.TimeBoundsBlocks{
		common.NewTimeBlocks(big.NewInt(100)),
		common.NewTimeBlocks(big.NewInt(120)),
	}
	afterMachine := mach.Clone()
	precondition := valprotocol.NewPrecondition(mach.Hash(), timeBounds, value.NewEmptyTuple())
	assertion, numSteps := afterMachine.ExecuteAssertion(1000, timeBounds, value.NewEmptyTuple(), 0)

	challengeHash := structures.ExecutionDataHash(
		numSteps,
		precondition.Hash(),
		valprotocol.NewExecutionAssertionStubFromAssertion(assertion).Hash(),
	)

	if err := testChallenge(
		valprotocol.InvalidExecutionChildType,
		challengeHash,
		"9af1e691e3db692cc9cad4e87b6490e099eb291e3b434a0d3f014dfd2bb747cc",
		"27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa",
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
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}
