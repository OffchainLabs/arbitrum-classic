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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	errors2 "github.com/pkg/errors"
)

type executionChallenge struct {
	*bisectionChallenge
	challenge *ethbridgecontracts.ExecutionChallenge
}

func newExecutionChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*executionChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	executionContract, err := ethbridgecontracts.NewExecutionChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}
	return &executionChallenge{bisectionChallenge: bisectionChallenge, challenge: executionContract}, nil
}

func (c *executionChallenge) BisectAssertion(
	ctx context.Context,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint64,
) error {
	machineHashes := make([][32]byte, 0, len(assertions)+1)
	inboxHashes := make([][32]byte, 0, len(assertions)+1)
	messageAccs := make([][32]byte, 0, len(assertions)+1)
	logAccs := make([][32]byte, 0, len(assertions)+1)
	gasses := make([]uint64, 0, len(assertions))
	machineHashes = append(machineHashes, assertions[0].BeforeMachineHash)
	inboxHashes = append(inboxHashes, assertions[0].BeforeInboxHash)
	messageAccs = append(messageAccs, assertions[0].FirstMessageHash)
	logAccs = append(logAccs, assertions[0].FirstLogHash)
	outCounts := make([]uint64, len(assertions)*2)
	for i, assertion := range assertions {
		machineHashes = append(machineHashes, assertion.AfterMachineHash)
		inboxHashes = append(inboxHashes, assertion.AfterInboxHash)
		messageAccs = append(messageAccs, assertion.LastMessageHash)
		logAccs = append(logAccs, assertion.LastLogHash)
		gasses = append(gasses, assertion.NumGas)
		outCounts[i] = assertion.MessageCount
		outCounts[i+len(assertions)] = assertion.LogCount
	}
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.challenge.BisectAssertion(
		c.auth.getAuth(ctx),
		machineHashes,
		inboxHashes,
		messageAccs,
		logAccs,
		outCounts,
		gasses,
		totalSteps,
	)
	if err != nil {
		return c.challenge.BisectAssertionCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			machineHashes,
			inboxHashes,
			messageAccs,
			logAccs,
			outCounts,
			gasses,
			totalSteps,
		)
	}
	return c.waitForReceipt(ctx, tx, "BisectAssertion")
}

func (c *executionChallenge) OneStepProof(
	ctx context.Context,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.challenge.OneStepProof(
		c.auth.getAuth(ctx),
		assertion.AfterInboxHash,
		assertion.FirstMessageHash,
		assertion.FirstLogHash,
		proof,
	)
	if err != nil {
		return c.challenge.OneStepProofCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			assertion.AfterInboxHash,
			assertion.FirstMessageHash,
			assertion.FirstLogHash,
			proof,
		)
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *executionChallenge) OneStepProofWithMessage(
	ctx context.Context,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
	msg inbox.InboxMessage,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.challenge.OneStepProofWithMessage(
		c.auth.getAuth(ctx),
		assertion.AfterInboxHash,
		assertion.FirstMessageHash,
		assertion.FirstLogHash,
		proof,
		uint8(msg.Kind),
		msg.ChainTime.BlockNum.AsInt(),
		msg.ChainTime.Timestamp,
		msg.Sender.ToEthAddress(),
		msg.InboxSeqNum,
		msg.Data,
	)
	if err != nil {
		return c.challenge.OneStepProofInboxCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			assertion.AfterInboxHash,
			assertion.FirstMessageHash,
			assertion.FirstLogHash,
			proof,
			uint8(msg.Kind),
			msg.ChainTime.BlockNum.AsInt(),
			msg.ChainTime.Timestamp,
			msg.Sender.ToEthAddress(),
			msg.InboxSeqNum,
			msg.Data,
		)
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *executionChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	assertionHashes []common.Hash,
) error {
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		assertionHashes,
	)
}
