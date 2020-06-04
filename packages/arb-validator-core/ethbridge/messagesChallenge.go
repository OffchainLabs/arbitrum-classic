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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/messageschallenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type messagesChallenge struct {
	*bisectionChallenge
	contract *messageschallenge.MessagesChallenge
}

func newMessagesChallenge(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*messagesChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	messagesContract, err := messageschallenge.NewMessagesChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to messagesChallenge")
	}
	return &messagesChallenge{bisectionChallenge: bisectionChallenge, contract: messagesContract}, nil
}

func (c *messagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	hashes := hashSliceToRaw(segmentHashes)
	tx, err := c.contract.Bisect(
		c.auth.getAuth(ctx),
		hashSliceToRaw(chainHashes),
		hashes,
		chainLength,
	)
	if err != nil {
		return c.contract.BisectCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			hashSliceToRaw(chainHashes),
			hashes,
			chainLength)
	}
	return c.waitForReceipt(ctx, tx, "Bisect")
}

func (c *messagesChallenge) OneStepProofTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredTransaction,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofTransactionMessage(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.Chain.ToEthAddress(),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.SequenceNum,
		msg.Value,
		msg.Data,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofTransactionMessage")
}

func (c *messagesChallenge) OneStepProofTransactionBatchMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredTransactionBatch,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofTransactionBatchMessage(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.Chain.ToEthAddress(),
		msg.TxData,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofTransactionBatchMessage")
}

func (c *messagesChallenge) OneStepProofEthMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredEth,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofEthMessage(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)

	if err != nil {
		return c.contract.OneStepProofEthMessageCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			lowerHashA,
			lowerHashB.HashImage,
			msg.To.ToEthAddress(),
			msg.From.ToEthAddress(),
			msg.Value,
			msg.BlockNum.AsInt(),
			msg.Timestamp,
			msg.MessageNum,
		)
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofEthMessage")
}

func (c *messagesChallenge) OneStepProofERC20Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredERC20,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofERC20Message(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofERC20Message")
}

func (c *messagesChallenge) OneStepProofERC721Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredERC721,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofERC721Message(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Id,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofERC721Message")
}

func (c *messagesChallenge) OneStepProofContractTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashPreImage,
	msg message.DeliveredContractTransaction,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProofContractTransactionMessage(
		c.auth.getAuth(ctx),
		lowerHashA,
		lowerHashB.HashImage,
		big.NewInt(lowerHashB.Size),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.Value,
		msg.Data,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProofContractTransactionMessage")
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	bisectionCount := uint64(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint64(0); i < bisectionCount; i++ {
		stepCount := valprotocol.CalculateBisectionStepCount(i, bisectionCount, chainLength.Uint64())
		bisectionHashes = append(
			bisectionHashes,
			valprotocol.MessageChallengeDataHash(
				chainHashes[i],
				chainHashes[i+1],
				segmentHashes[i],
				segmentHashes[i+1],
				new(big.Int).SetUint64(stepCount),
			),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}
