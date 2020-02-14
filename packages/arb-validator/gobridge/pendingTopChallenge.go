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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type pendingTopChallenge struct {
	*bisectionChallenge
	//contract *pendingtopchallenge.PendingTopChallenge
}

func newPendingTopChallenge(address common.Address, client *GoArbAuthClient) (*pendingTopChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}
	return &pendingTopChallenge{bisectionChallenge: bisectionChallenge}, nil
}

func (c *pendingTopChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	chainLength *big.Int,
) error {

	fmt.Println("in (c *pendingTopChallenge) Bisect")

	bisectionCount := len(chainHashes) - 1

	fmt.Println("c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash", c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash)
	fmt.Println("chainHashes[0]", chainHashes[0])
	fmt.Println("chainHashes[bisectionCount]", chainHashes[bisectionCount])
	fmt.Println("chainLength", chainLength)
	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(structures.PendingTopChallengeDataHash(chainHashes[0], chainHashes[bisectionCount], chainLength)) {
		return errors.New("Incorrect previous state")
	}

	if chainLength.Cmp(big.NewInt(1)) < 1 {
		return errors.New("Can't bisect chain of less than 2")
	}

	hashes := make([][32]byte, 0, bisectionCount)
	hashes = append(hashes, structures.PendingTopChallengeDataHash(
		chainHashes[0],
		chainHashes[1],
		new(big.Int).Add(new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount))), new(big.Int).Mod(chainLength, big.NewInt(int64(bisectionCount)))),
	))
	for i := 1; i < bisectionCount; i++ {
		hashes = append(hashes, structures.PendingTopChallengeDataHash(
			chainHashes[i],
			chainHashes[i+1],
			new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount)))))
	}

	c.commitToSegment(hashes)
	c.asserterResponded()

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.PendingTopBisectionEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
			ChainHashes: chainHashes,
			TotalLength: chainLength,
			Deadline:    c.client.GoEthClient.challenges[c.contractAddress].deadline,
		},
	})
	return nil
}

func (c *pendingTopChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	value common.Hash,
) error {
	fmt.Println("in (c *pendingTopChallenge) OneStepProof")
	//c.auth.Lock()
	//defer c.auth.Unlock()
	//tx, err := c.contract.OneStepProof(
	//	c.auth.getAuth(ctx),
	//	lowerHashA,
	//	value,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "OneStepProof")
	//return keccak256(
	//	abi.encodePacked(
	//		pending,
	//		message
	//)
	matchHash := structures.PendingTopChallengeDataHash(lowerHashA, structures.AddMessageToPending(lowerHashA, value), big.NewInt(1))
	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("Incorrect previous state")
	}

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	//_asserterWin()

	return nil
}

func (c *pendingTopChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	chainLength uint64,
) error {
	bisectionCount := uint64(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint64(0); i < bisectionCount; i++ {
		stepCount := structures.CalculateBisectionStepCount(i, bisectionCount, chainLength)
		fmt.Println("PendingTopChallengeDataHash", structures.PendingTopChallengeDataHash(
			chainHashes[i],
			chainHashes[i+1],
			new(big.Int).SetUint64(uint64(stepCount)),
		))
		fmt.Println("chainHashes[i]", chainHashes[i])
		fmt.Println("chainHashes[i+1]", chainHashes[i+1])
		fmt.Println("stepCount", stepCount)

		bisectionHashes = append(
			bisectionHashes,
			structures.PendingTopChallengeDataHash(
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
