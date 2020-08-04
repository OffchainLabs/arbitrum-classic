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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type messagesChallenge struct {
	*bisectionChallenge
	contract *ethbridgecontracts.MessagesChallenge
}

func newMessagesChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*messagesChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	messagesContract, err := ethbridgecontracts.NewMessagesChallenge(address, client)
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

	hashes := common.HashSliceToRaw(segmentHashes)
	tx, err := c.contract.Bisect(
		c.auth.getAuth(ctx),
		common.HashSliceToRaw(chainHashes),
		hashes,
		chainLength,
	)
	if err != nil {
		return c.contract.BisectCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			common.HashSliceToRaw(chainHashes),
			hashes,
			chainLength)
	}
	return c.waitForReceipt(ctx, tx, "Bisect")
}

func (c *messagesChallenge) OneStepProof(
	ctx context.Context,
	afterGlobalInbox common.Hash,
	beforeVmInbox value.HashPreImage,
	msg inbox.InboxMessage,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.contract.OneStepProof(
		c.auth.getAuth(ctx),
		afterGlobalInbox,
		beforeVmInbox.GetInnerHash(),
		big.NewInt(beforeVmInbox.Size()),
		uint8(msg.Kind),
		msg.ChainTime.BlockNum.AsInt(),
		msg.ChainTime.Timestamp,
		msg.Sender.ToEthAddress(),
		msg.InboxSeqNum,
		msg.Data,
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
