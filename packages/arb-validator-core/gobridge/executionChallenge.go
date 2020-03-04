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

package gobridge

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type ExecutionChallenge struct {
	*bisectionChallenge
}

func NewExecutionChallenge(address common.Address, client *GoArbAuthClient) (*ExecutionChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}

	vm := &ExecutionChallenge{bisectionChallenge: bisectionChallenge}
	return vm, err
}

func (c *ExecutionChallenge) BisectAssertion(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint64,
) error {
	machineHashes := make([][32]byte, 0, len(assertions)+1)
	didInboxInsns := make([]bool, 0, len(assertions))
	messageAccs := make([][32]byte, 0, len(assertions)+1)
	logAccs := make([][32]byte, 0, len(assertions)+1)
	gasses := make([]uint64, 0, len(assertions))
	machineHashes = append(machineHashes, precondition.BeforeHash)
	messageAccs = append(messageAccs, assertions[0].FirstMessageHash)
	logAccs = append(logAccs, assertions[0].FirstLogHash)

	var totalGas uint64
	everDidInboxInsn := false
	for _, assertion := range assertions {
		totalGas += assertion.NumGas
		everDidInboxInsn = everDidInboxInsn || assertion.DidInboxInsn

		machineHashes = append(machineHashes, assertion.AfterHash)
		didInboxInsns = append(didInboxInsns, assertion.DidInboxInsn)
		messageAccs = append(messageAccs, assertion.LastMessageHash)
		logAccs = append(logAccs, assertion.LastLogHash)
		gasses = append(gasses, assertion.NumGas)
	}

	bisectionCount := len(machineHashes) - 1

	preconditionHash := valprotocol.ExecutionPreconditionHash(machineHashes[0], precondition.TimeBounds, precondition.BeforeInbox.Hash())

	assertionHash := generateAssertionHash(
		machineHashes[bisectionCount],
		everDidInboxInsn,
		totalGas,
		messageAccs[0],
		messageAccs[bisectionCount],
		logAccs[0],
		logAccs[bisectionCount],
	)

	if !c.client.challenges[c.contractAddress].challengerDataHash.Equals(valprotocol.ExecutionDataHash(totalSteps, preconditionHash, assertionHash)) {
		return errors.New("BisectAssertion Incorrect previous state")
	}

	assertionHash = generateAssertionHash(
		machineHashes[1],
		didInboxInsns[0],
		gasses[0],
		messageAccs[0],
		messageAccs[1],
		logAccs[0],
		logAccs[1],
	)

	hashes := make([][32]byte, 0, bisectionCount)
	hashes = append(hashes, valprotocol.ExecutionDataHash(
		totalSteps/uint64(bisectionCount)+totalSteps%uint64(bisectionCount),
		preconditionHash,
		assertionHash,
	))

	for i := 1; i < bisectionCount; i++ {
		if didInboxInsns[i-1] {
			precondition.BeforeInbox = value.NewEmptyTuple()
		}
		assertionHash = generateAssertionHash(
			machineHashes[i+1],
			didInboxInsns[i],
			gasses[i],
			messageAccs[i],
			messageAccs[i+1],
			logAccs[i],
			logAccs[i+1],
		)
		hashes = append(hashes, valprotocol.ExecutionDataHash(
			totalSteps/uint64(bisectionCount),
			valprotocol.ExecutionPreconditionHash(machineHashes[i], precondition.TimeBounds, precondition.BeforeInbox.Hash()),
			assertionHash))
	}

	c.commitToSegment(hashes)
	c.asserterResponded()

	c.client.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.ExecutionBisectionEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			},
			Assertions: assertions,
			TotalSteps: totalSteps,
			Deadline:   c.client.challenges[c.contractAddress].deadline,
		},
	})

	return nil
}

func (c *ExecutionChallenge) OneStepProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) error {
	valprotocol.ExecutionPreconditionHash(precondition.BeforeHash, precondition.TimeBounds, precondition.BeforeInbox.Hash())
	precondition.Hash()

	matchHash := valprotocol.ExecutionDataHash(1, precondition.Hash(), assertion.Hash())
	if !c.client.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("OneStepProof Incorrect previous state")
	}

	// TODO: executionChallenge one step proof validation
	//	uint256 correctProof = OneStepProof.validateProof(
	//		_beforeHash,
	//		_timeBoundsBlocks,
	//		_beforeInbox,
	//		_afterHash,
	//		_didInboxInsns,
	//		_firstMessage,
	//		_lastMessage,
	//		_firstLog,
	//		_lastLog,
	//		_gas,
	//		_proof
	//	);
	//
	// for now make OSP always valid

	//	require(correctProof == 0, OSP_PROOF);
	//	emit OneStepProofCompleted();
	c.client.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			},
		},
	})

	//c.client.GoEthClient.pubMsg(arbbridge.MaybeEvent{
	//	Event: arbbridge.ChallengeCompletedEvent{
	//		ChainInfo: arbbridge.ChainInfo{
	//			BlockId: c.client.GoEthClient.getCurrentBlock(),
	//		},
	//		Winner:            c.client.GoEthClient.challenges[c.contractAddress].asserter,
	//		Loser:             c.client.GoEthClient.challenges[c.contractAddress].challenger,
	//		ChallengeContract: c.contractAddress,
	//	},
	//})

	//	_asserterWin();

	return nil
}

func (c *ExecutionChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	preconditions []*valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint64,
) error {
	bisectionHashes := make([]common.Hash, 0, len(assertions))
	for i := range assertions {
		stepCount := valprotocol.CalculateBisectionStepCount(uint64(i), uint64(len(assertions)), totalSteps)
		bisectionHashes = append(
			bisectionHashes,
			valprotocol.ExecutionDataHash(stepCount, preconditions[i].Hash(), assertions[i].Hash()),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}

func generateAssertionHash(
	machineHash [32]byte,
	everDidInboxInsn bool,
	numGas uint64,
	firstMsgHash [32]byte,
	lastMsgHash [32]byte,
	firstLogHash [32]byte,
	lastLogHash [32]byte,
) common.Hash {
	stub := valprotocol.ExecutionAssertionStub{
		machineHash,
		everDidInboxInsn,
		numGas,
		firstMsgHash,
		lastMsgHash,
		firstLogHash,
		lastLogHash,
	}
	return stub.Hash()
}
