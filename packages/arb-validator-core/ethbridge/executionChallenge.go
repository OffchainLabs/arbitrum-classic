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

package ethbridge

import (
	"context"
	errors2 "github.com/pkg/errors"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/executionchallenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type executionChallenge struct {
	*bisectionChallenge
	challenge *executionchallenge.ExecutionChallenge
}

func newExecutionChallenge(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*executionChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	executionContract, err := executionchallenge.NewExecutionChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}
	return &executionChallenge{bisectionChallenge: bisectionChallenge, challenge: executionContract}, nil
}

func (c *executionChallenge) BisectAssertion(
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
	for _, assertion := range assertions {
		machineHashes = append(machineHashes, assertion.AfterHash)
		didInboxInsns = append(didInboxInsns, assertion.DidInboxInsn)
		messageAccs = append(messageAccs, assertion.LastMessageHash)
		logAccs = append(logAccs, assertion.LastLogHash)
		gasses = append(gasses, assertion.NumGas)
	}
	c.auth.Lock()
	defer c.auth.Unlock()
	beforeInboxHash := precondition.BeforeInbox.Hash()
	tx, err := c.challenge.BisectAssertion(
		c.auth.getAuth(ctx),
		beforeInboxHash,
		precondition.TimeBounds.AsIntArray(),
		machineHashes,
		didInboxInsns,
		messageAccs,
		logAccs,
		gasses,
		totalSteps,
	)
	if err != nil {
		return c.challenge.BisectAssertionCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			beforeInboxHash,
			precondition.TimeBounds.AsIntArray(),
			machineHashes,
			didInboxInsns,
			messageAccs,
			logAccs,
			gasses,
			totalSteps,
		)
	}
	return c.waitForReceipt(ctx, tx, "BisectAssertion")
}

func (c *executionChallenge) OneStepProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	hashPreImage := precondition.BeforeInbox.GetPreImage()
	tx, err := c.challenge.OneStepProof(
		c.auth.getAuth(ctx),
		precondition.BeforeHash,
		hashPreImage.GetPreImageHash(),
		big.NewInt(hashPreImage.Size()),
		precondition.TimeBounds.AsIntArray(),
		assertion.AfterHash,
		assertion.DidInboxInsn,
		assertion.FirstMessageHash,
		assertion.LastMessageHash,
		assertion.FirstLogHash,
		assertion.LastLogHash,
		assertion.NumGas,
		proof,
	)
	if err != nil {
		return c.challenge.OneStepProofCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			precondition.BeforeHash,
			hashPreImage.GetPreImageHash(),
			big.NewInt(hashPreImage.Size()),
			precondition.TimeBounds.AsIntArray(),
			assertion.AfterHash,
			assertion.DidInboxInsn,
			assertion.FirstMessageHash,
			assertion.LastMessageHash,
			assertion.FirstLogHash,
			assertion.LastLogHash,
			assertion.NumGas,
			proof,
		)
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *executionChallenge) ChooseSegment(
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
