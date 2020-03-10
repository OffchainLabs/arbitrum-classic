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
	"math/big"
)

const (
	noChallenge    = 0
	asserterTurn   = 1
	challengerTurn = 2
)

type challengeData struct {
	deadline             common.TimeTicks
	challengerDataHash   common.Hash
	state                int
	challengePeriodTicks common.TimeTicks
	asserter             common.Address
	challenger           common.Address
	challengeType        *big.Int
}

type challenge struct {
	client          *GoArbAuthClient
	challengeData   *challengeData
	contractAddress common.Address
}

func newChallenge(address common.Address, client *GoArbAuthClient) (*challenge, error) {
	return client.challenges[address], nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {
	c.client.goEthMutex.Lock()
	defer c.client.goEthMutex.Unlock()

	if common.TicksFromBlockNum(c.client.getCurrentBlock().Height).Cmp(c.challengeData.deadline) < 0 {
		return errors.New("Deadline hasn't expired")
	}
	if c.challengeData.state == asserterTurn {
		c.resolveChallenge(c.challengeData.challenger, c.challengeData.asserter)
		c.client.pubMsg(c.contractAddress, arbbridge.AsserterTimeoutEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			}})
	} else {
		c.resolveChallenge(c.challengeData.asserter, c.challengeData.challenger)
		c.client.pubMsg(c.contractAddress, arbbridge.ChallengerTimeoutEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			}})
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

func (c *challenge) resolveChallenge(winner common.Address, loser common.Address) {
	rollup, ok := c.client.rollups[c.contractAddress]
	if ok {
		winningStaker := rollup.rollup.stakers[winner]
		winningStaker.inChallenge = false
		transferEth(c.client.goEthdata, winner, loser, rollup.rollup.chainParams.StakeRequirement)
		delete(c.client.rollups[c.contractAddress].rollup.stakers, loser)
	}
	c.client.pubMsg(c.contractAddress, arbbridge.ChallengeCompletedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: c.client.getCurrentBlock(),
		},
		Winner:            winner,
		Loser:             loser,
		ChallengeContract: c.contractAddress,
	})

}

type challengeWatcher struct {
	client *goEthdata
	*challenge
}

func newChallengeWatcher(address common.Address, client *goEthdata) (*challengeWatcher, error) {
	chalData := client.challenges[address]

	return &challengeWatcher{challenge: chalData, client: client}, nil
}

func (c *challengeWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	c.client.goEthMutex.Lock()
	defer c.client.goEthMutex.Unlock()

	return c.client.blockMsgs[blockId].msgs[c.contractAddress], nil
}
