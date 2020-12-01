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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var continuedChallengeID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(ethbridgecontracts.BisectionChallengeABI))
	if err != nil {
		panic(err)
	}
	continuedChallengeID = parsed.Events["Continued"].ID
}

type bisectionChallenge struct {
	*challenge
	BisectionChallenge *ethbridgecontracts.BisectionChallenge
}

func newBisectionChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*bisectionChallenge, error) {
	challenge, err := newChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	bisectionContract, err := ethbridgecontracts.NewBisectionChallenge(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ChallengeManager")
	}
	vm := &bisectionChallenge{
		challenge:          challenge,
		BisectionChallenge: bisectionContract,
	}
	return vm, err
}

func (c *bisectionChallenge) chooseSegment(
	ctx context.Context,
	segmentToChallenge uint16,
	segments []common.Hash,
) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	if int(segmentToChallenge) >= len(segments) {
		return errors.New("invalid assertionToChallenge")
	}

	tree := NewMerkleTree(segments)
	tx, err := c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.BisectionChallenge.ChooseSegment(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			tree.GetProofFlat(int(segmentToChallenge)),
			tree.GetRoot(),
			tree.GetNode(int(segmentToChallenge)),
		)
	})
	if err != nil {
		return c.BisectionChallenge.ChooseSegmentCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
			big.NewInt(int64(segmentToChallenge)),
			tree.GetProofFlat(int(segmentToChallenge)),
			tree.GetRoot(),
			tree.GetNode(int(segmentToChallenge)),
		)
	}
	return c.waitForReceipt(ctx, tx, "ChooseSegment")
}

type bisectionChallengeWatcher struct {
	*challengeWatcher
	BisectionChallenge *ethbridgecontracts.BisectionChallenge
}

func newBisectionChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*bisectionChallengeWatcher, error) {
	challenge, err := newChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	bisectionContract, err := ethbridgecontracts.NewBisectionChallenge(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ChallengeManager")
	}
	vm := &bisectionChallengeWatcher{
		challengeWatcher:   challenge,
		BisectionChallenge: bisectionContract,
	}
	return vm, err
}

func (c *bisectionChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		continuedChallengeID,
	}
	return append(tops, c.challengeWatcher.topics()...)
}

func (c *bisectionChallengeWatcher) parseBisectionEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == continuedChallengeID {
		contChal, err := c.BisectionChallenge.ParseContinued(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.ContinueChallengeEvent{
			ChainInfo:    chainInfo,
			SegmentIndex: contChal.SegmentIndex,
			Deadline:     common.TimeTicks{Val: contChal.DeadlineTicks},
		}, nil
	} else {
		return c.challengeWatcher.parseChallengeEvent(chainInfo, log)
	}
}
