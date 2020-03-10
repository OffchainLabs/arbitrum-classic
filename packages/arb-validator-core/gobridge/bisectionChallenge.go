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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"
)

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

	if !tree.GetRoot().Equals(c.challenge.challengeData.challengerDataHash) {
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
	c.challenge.challengeData.challengerDataHash = segments[segmentToChallenge]
	c.client.pubMsg(c.contractAddress, arbbridge.ContinueChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: c.client.getCurrentBlock(),
		},
		SegmentIndex: big.NewInt(int64(segmentToChallenge)),
		Deadline:     c.challenge.challengeData.deadline,
	})

	return nil
}

func hashSliceToHashes(slice [][32]byte) []common.Hash {
	ret := make([]common.Hash, 0, len(slice))
	for _, a := range slice {
		ret = append(ret, a)
	}
	return ret
}
