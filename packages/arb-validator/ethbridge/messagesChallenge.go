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
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/messageschallenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
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
	tx, err := c.contract.Bisect(
		c.auth.getAuth(ctx),
		hashSliceToRaw(chainHashes),
		hashSliceToRaw(segmentHashes),
		chainLength,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "Bisect")
}

func (c *messagesChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	topHashA common.Hash,
	lowerHashB common.Hash,
	topHashB common.Hash,
	value common.Hash,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProof(
		c.auth.getAuth(ctx),
		lowerHashA,
		topHashA,
		lowerHashB,
		topHashB,
		value,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	bisectionCount := uint32(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint32(0); i < bisectionCount; i++ {
		stepCount := structures.CalculateBisectionStepCount(i, bisectionCount, uint32(chainLength.Uint64()))
		bisectionHashes = append(
			bisectionHashes,
			structures.MessageChallengeDataHash(
				chainHashes[i],
				chainHashes[i+1],
				segmentHashes[i],
				segmentHashes[i+1],
				new(big.Int).SetUint64(uint64(stepCount)),
			),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}
