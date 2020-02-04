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
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
	"math/big"
)

var continuedChallengeID ethcommon.Hash

type bisectionChallenge struct {
	*challenge
	//challengeState
}

func newBisectionChallenge(address common.Address, client *MockArbAuthClient) (*bisectionChallenge, error) {
	challenge, err := newChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &bisectionChallenge{
		challenge: challenge,
	}
	//err = vm.setupContracts()
	return vm, nil
}

func (c *bisectionChallenge) chooseSegment(
	ctx context.Context,
	segmentToChallenge uint16,
	segments []common.Hash,
) error {
	fmt.Println("in bisectionChallenge - chooseSegment")
	//tree := NewMerkleTree(segments)
	//c.auth.Context = ctx
	//tx, err := c.bisectionChallenge.ChooseSegment(
	//	c.auth,
	//	big.NewInt(int64(segmentToChallenge)),
	//	tree.GetProofFlat(int(segmentToChallenge)),
	//	tree.GetRoot(),
	//	tree.GetNode(int(segmentToChallenge)),
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "ChooseSegment")

	//require(_bisectionRoot == challengeState, CON_PREV);
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
	//challengeState = _bisectionHash;
	//
	//challengerResponded();
	//emit Continued(_segmentToChallenge, deadlineTicks);
	c.client.MockEthClient.pubMsg(arbbridge.MaybeEvent{
		Event: arbbridge.ContinueChallengeEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.MockEthClient.getCurrentBlock(),
			},
			SegmentIndex: big.NewInt(int64(segmentToChallenge)),
			Deadline:     c.client.MockEthClient.challenges[c.contractAddress].deadline,
		},
	})

	return nil
}

type bisectionChallengeWatcher struct {
	*challengeWatcher
	BisectionChallenge *executionchallenge.BisectionChallenge
}

func newBisectionChallengeWatcher(address common.Address, client *MockArbClient) (*bisectionChallengeWatcher, error) {
	challenge, err := newChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	//bisectionContract, err := executionchallenge.newBisectionChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//}
	vm := &bisectionChallengeWatcher{
		challengeWatcher:   challenge,
		BisectionChallenge: nil,
	}
	return vm, err
}

func (c *bisectionChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		continuedChallengeID,
	}
	return append(tops, c.challengeWatcher.topics()...)
}

//
//func (c *bisectionChallengeWatcher) parseBisectionEvent(log types.Log) (arbbridge.Event, error) {
//	if log.Topics[0] == continuedChallengeID {
//		contChal, err := c.BisectionChallenge.ParseContinued(log)
//		if err != nil {
//			return nil, err
//		}
//		return arbbridge.ContinueChallengeEvent{
//			SegmentIndex: contChal.SegmentIndex,
//			Deadline:     common.TimeTicks{Val: contChal.DeadlineTicks},
//		}, nil
//	} else {
//		return c.challengeWatcher.parseChallengeEvent(log)
//	}
//}
