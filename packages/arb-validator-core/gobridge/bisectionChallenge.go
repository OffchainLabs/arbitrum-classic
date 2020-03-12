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
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
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

	root := tree.GetRoot()
	if !root.Equals(c.challenge.challengerDataHash) {
		return errors.New("chooseSegment Incorrect previous state")
	}
	proof := tree.GetProofFlat(int(segmentToChallenge))
	hash := tree.GetNode(int(segmentToChallenge))
	index := uint(segmentToChallenge) + 1
	h := hash

	for j := 32; j <= len(proof); j += 32 {
		el := proof[j-32 : j]

		// calculate remaining elements in proof
		remaining := uint((len(proof) - j + 32) / 32)

		// we don't assume that the tree is padded to a power of 2
		// if the index is odd then the proof will start with a hash at a higher
		// layer, so we have to adjust the index to be the index at that layer
		for remaining > 0 && index%2 == 1 && index > 1<<remaining {
			index = uint(index)/2 + 1
		}

		if index%2 == 0 {
			h = hashing.SoliditySHA3(el, hashing.Bytes32(h))
			index = index / 2
		} else {
			h = hashing.SoliditySHA3(hashing.Bytes32(h), el)
			index = uint(index)/2 + 1
		}
	}

	if !h.Equals(root) {
		return errors.New("Invalid assertion selected")
	}

	c.challengerResponded()
	c.challenge.challengerDataHash = segments[segmentToChallenge]
	c.client.pubMsg(c.contractAddress, arbbridge.ContinueChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: c.client.getCurrentBlock(),
		},
		SegmentIndex: big.NewInt(int64(segmentToChallenge)),
		Deadline:     c.challenge.deadline,
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
