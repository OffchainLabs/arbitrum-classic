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

package mockbridge

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type messagesChallenge struct {
	*bisectionChallenge
}

func newMessagesChallenge(address common.Address, client arbbridge.ArbClient) (*messagesChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client) //, auth??
	if err != nil {
		return nil, err
	}
	vm := &messagesChallenge{bisectionChallenge: bisectionChallenge}
	err = vm.setupContracts()
	return vm, err
}

func (c *messagesChallenge) setupContracts() error {
	//challengeManagerContract, err := messageschallenge.NewMessagesChallenge(c.address, c.Client)
	//if err != nil {
	//	return errors2.Wrap(err, "Failed to connect to MessagesChallenge")
	//}
	//
	//c.challenge = challengeManagerContract
	return nil
}

func (vm *messagesChallenge) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	return nil, nil
}

func (c *messagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.Bisect(
	//	c.auth,
	//	chainHashes,
	//	segmentHashes,
	//	chainLength,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "Bisect")
	return nil
}

func (c *messagesChallenge) OneStepProofTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredTransaction,
) error {
	return nil
}

func (c *messagesChallenge) OneStepProofEthMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredEth,
) error {
	return nil
}

func (c *messagesChallenge) OneStepProofERC20Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC20,
) error {
	return nil
}

func (c *messagesChallenge) OneStepProofERC721Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC721,
) error {
	return nil
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	return nil
}
