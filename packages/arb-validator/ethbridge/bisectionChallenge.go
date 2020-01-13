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
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
)

var continuedChallengeID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.BisectionChallengeABI))
	if err != nil {
		panic(err)
	}
	continuedChallengeID = parsed.Events["Continued"].ID()
}

type bisectionChallenge struct {
	*challenge
	BisectionChallenge *executionchallenge.BisectionChallenge
}

func newBisectionChallenge(address ethcommon.Address, client *ethclient.Client, auth *bind.TransactOpts) (*bisectionChallenge, error) {
	challenge, err := newChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	vm := &bisectionChallenge{
		challenge:          challenge,
		BisectionChallenge: nil,
	}
	err = vm.setupContracts()
	return vm, err
}

func (c *bisectionChallenge) setupContracts() error {
	challengeManagerContract, err := executionchallenge.NewBisectionChallenge(c.address, c.client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	c.BisectionChallenge = challengeManagerContract
	return nil
}

func (c *bisectionChallenge) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		continuedChallengeID,
	}
	return append(tops, c.challenge.topics()...)
}

func (c *bisectionChallenge) parseBisectionEvent(log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == continuedChallengeID {
		contChal, err := c.BisectionChallenge.ParseContinued(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.ContinueChallengeEvent{
			SegmentIndex: contChal.SegmentIndex,
			Deadline:     common.TimeTicks{Val: contChal.DeadlineTicks},
		}, nil
	} else {
		return c.challenge.parseChallengeEvent(log)
	}
}

func (c *bisectionChallenge) chooseSegment(
	ctx context.Context,
	segmentToChallenge uint16,
	segments []common.Hash,
) error {
	tree := NewMerkleTree(segments)
	c.auth.Context = ctx
	tx, err := c.BisectionChallenge.ChooseSegment(
		c.auth,
		big.NewInt(int64(segmentToChallenge)),
		tree.GetProofFlat(int(segmentToChallenge)),
		tree.GetRoot(),
		tree.GetNode(int(segmentToChallenge)),
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "ChooseSegment")
}
