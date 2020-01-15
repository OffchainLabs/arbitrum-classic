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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/messageschallenge"
	errors2 "github.com/pkg/errors"
	"math/big"

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
	//messagesContract, err := messageschallenge.NewMessagesChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to messagesChallenge")
	//}
	return &messagesChallenge{bisectionChallenge: bisectionChallenge}, nil
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

func (c *messagesChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	topHashA common.Hash,
	lowerHashB common.Hash,
	topHashB common.Hash,
	value common.Hash,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.OneStepProof(
	//	c.auth,
	//	lowerHashA,
	//	topHashA,
	//	lowerHashB,
	//	topHashB,
	//	value,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "OneStepProof")
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
