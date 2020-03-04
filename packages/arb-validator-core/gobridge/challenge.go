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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

const (
	noChallenge    = 0
	asserterTurn   = 1
	challengerTurn = 2
)

type challenge struct {
	client          *GoArbAuthClient
	challengeData   *challengeData
	contractAddress common.Address
}

func newChallenge(address common.Address, client *GoArbAuthClient) (*challenge, error) {
	chal := challenge{
		client:          client,
		contractAddress: address,
		challengeData:   client.challenges[address],
	}

	return &chal, nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {

	if common.TicksFromBlockNum(c.client.getCurrentBlock().Height).Cmp(c.challengeData.deadline) < 0 {
		return errors.New("Deadline hasn't expired")
	}
	if c.challengeData.state == asserterTurn {
		c.client.pubMsg(c.challengeData, arbbridge.MaybeEvent{
			Event: arbbridge.AsserterTimeoutEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: c.client.getCurrentBlock(),
				},
			},
		})
	} else {
		c.client.pubMsg(c.challengeData, arbbridge.MaybeEvent{
			Event: arbbridge.ChallengerTimeoutEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: c.client.getCurrentBlock(),
				},
			},
		})
	}
	return nil
}

func (c *challenge) commitToSegment(hashes [][32]byte) {
	tree := valprotocol.NewMerkleTree(hashSliceToHashes(hashes))
	c.challengeData.challengerDataHash = tree.GetRoot()
}

func (c *challenge) asserterResponded() {
	c.challengeData.state = challengerTurn
	currentTicks := common.TicksFromBlockNum(c.client.getCurrentBlock().Height)
	c.challengeData.deadline = currentTicks.Add(c.challengeData.challengePeriodTicks)

}

func (c *challenge) challengerResponded() {
	c.challengeData.state = asserterTurn
	currentTicks := common.TicksFromBlockNum(c.client.getCurrentBlock().Height)
	c.challengeData.deadline = currentTicks.Add(c.challengeData.challengePeriodTicks)

}

type challengeWatcher struct {
	client    *goEthdata
	Challenge common.Address
}

func newChallengeWatcher(address common.Address, client *goEthdata) (*challengeWatcher, error) {
	return &challengeWatcher{client: client, Challenge: address}, nil
}
