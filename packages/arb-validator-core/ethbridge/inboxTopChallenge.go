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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type inboxTopChallenge struct {
	*bisectionChallenge
	contract *ethbridgecontracts.InboxTopChallenge
}

func newInboxTopChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*inboxTopChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	inboxTopContract, err := ethbridgecontracts.NewInboxTopChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to InboxTopChallenge")
	}
	return &inboxTopChallenge{bisectionChallenge: bisectionChallenge, contract: inboxTopContract}, nil
}

func (c *inboxTopChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	chainLength *big.Int,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.Bisect(
		c.auth.getAuth(ctx),
		common.HashSliceToRaw(chainHashes),
		chainLength,
	)
	if err != nil {
		return c.contract.BisectCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			common.HashSliceToRaw(chainHashes),
			chainLength,
		)
	}
	return c.waitForReceipt(ctx, tx, "Bisect")
}

func (c *inboxTopChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	value common.Hash,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProof(
		c.auth.getAuth(ctx),
		lowerHashA,
		value,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *inboxTopChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	chainLength uint64,
) error {
	bisectionCount := uint64(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint64(0); i < bisectionCount; i++ {
		stepCount := valprotocol.CalculateBisectionStepCount(i, bisectionCount, chainLength)
		bisectionHashes = append(
			bisectionHashes,
			valprotocol.InboxTopChallengeDataHash(
				chainHashes[i],
				chainHashes[i+1],
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
