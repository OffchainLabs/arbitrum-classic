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

package gobridge

import (
	"context"
	"errors"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"
)

var continuedChallengeID ethcommon.Hash

type bisectionChallenge struct {
	*challenge
}

func newBisectionChallenge(address common.Address, client *GoArbAuthClient) (*bisectionChallenge, error) {
	challenge, err := newChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &bisectionChallenge{
		challenge: challenge,
	}
	return vm, nil
}

func (c *bisectionChallenge) chooseSegment(
	ctx context.Context,
	segmentToChallenge uint16,
	segments []common.Hash,
) error {
	fmt.Println("in bisectionChallenge - chooseSegment")
	tree := valprotocol.NewMerkleTree(segments)

	if !tree.GetRoot().Equals(c.challengeData.challengerDataHash) {
		return errors.New("chooseSegment Incorrect previous state")
	}

	// TODO: figure out merkle verify proof
	//require(
	//	MerkleLib.verifyProof(
	//		_proof,
	//		_bisectionRoot,
	//		_bisectionHash,
	//		_segmentToChallenge + 1
	//),
	//CON_PROOF
	//);
	//
	c.challengerResponded()
	c.challengeData.challengerDataHash = segments[segmentToChallenge]
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.ContinueChallengeEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
			SegmentIndex: big.NewInt(int64(segmentToChallenge)),
			Deadline:     c.challengeData.deadline,
		},
	})

	return nil
}

type bisectionChallengeWatcher struct {
	*challengeWatcher
}

func newBisectionChallengeWatcher(address common.Address, client *GoArbClient) (*bisectionChallengeWatcher, error) {
	challenge, err := newChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	vm := &bisectionChallengeWatcher{
		challengeWatcher: challenge,
	}
	return vm, err
}

func (c *bisectionChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		continuedChallengeID,
	}
	return append(tops, c.challengeWatcher.topics()...)
}

func (c *challenge) commitToSegment(hashes [][32]byte) {
	tree := valprotocol.NewMerkleTree(hashSliceToHashes(hashes))
	c.challengeData.challengerDataHash = tree.GetRoot()
}

func (c *challenge) asserterResponded() {
	c.challengeData.state = challengerTurn
	currentTicks := common.TicksFromBlockNum(c.client.GoEthClient.getCurrentBlock().Height)
	c.challengeData.deadline = currentTicks.Add(c.challengeData.challengePeriodTicks)

}

func (c *challenge) challengerResponded() {
	c.challengeData.state = asserterTurn
	currentTicks := common.TicksFromBlockNum(c.client.GoEthClient.getCurrentBlock().Height)
	c.challengeData.deadline = currentTicks.Add(c.challengeData.challengePeriodTicks)

}
