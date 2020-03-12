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
}

type challenge struct {
	client               *goEthdata
	rollupAddr           common.Address
	contractAddress      common.Address
	deadline             common.TimeTicks
	challengerDataHash   common.Hash
	state                int
	challengePeriodTicks common.TimeTicks
	asserter             common.Address
	challenger           common.Address
	challengeType        *big.Int
}

func newChallenge(address common.Address, client *GoArbAuthClient) (*challenge, error) {
	challenge, ok := client.challenges[address]
	if !ok {
		return nil, errors.New("invalid challenge address")
	}
	return challenge, nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {
	c.client.goEthMutex.Lock()
	defer c.client.goEthMutex.Unlock()

	if common.TicksFromBlockNum(c.client.getCurrentBlock().Height).Cmp(c.deadline) < 0 {
		return errors.New("Deadline hasn't expired")
	}
	if c.state == asserterTurn {
		err := c.resolveChallenge(c.challenger, c.asserter)
		if err != nil {
			return err
		}
		c.client.pubMsg(c.contractAddress, arbbridge.AsserterTimeoutEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			}})
	} else {
		err := c.resolveChallenge(c.asserter, c.challenger)
		if err != nil {
			return err
		}
		c.client.pubMsg(c.contractAddress, arbbridge.ChallengerTimeoutEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.getCurrentBlock(),
			}})
	}
	return nil
}

func (c *challenge) commitToSegment(hashes [][32]byte) {
	tree := valprotocol.NewMerkleTree(hashSliceToHashes(hashes))
	c.challengerDataHash = tree.GetRoot()
}

func (c *challenge) asserterResponded() {
	c.state = challengerTurn
	currentTicks := common.TicksFromBlockNum(c.client.getCurrentBlock().Height)
	c.deadline = currentTicks.Add(c.challengePeriodTicks)

}

func (c *challenge) challengerResponded() {
	c.state = asserterTurn
	currentTicks := common.TicksFromBlockNum(c.client.getCurrentBlock().Height)
	c.deadline = currentTicks.Add(c.challengePeriodTicks)

}

func (c *challenge) resolveChallenge(winner common.Address, loser common.Address) error {
	rollup, ok := c.client.rollups[c.rollupAddr]
	if ok {
		winningStaker := rollup.stakers[winner]
		winningStaker.inChallenge = false
		half := new(big.Int).Div(rollup.chainParams.StakeRequirement, big.NewInt(2))
		err := transferEth(c.client, winner, c.contractAddress, half)
		if err != nil {
			return err
		}
		delete(c.client.rollups[c.contractAddress].stakers, loser)
	}
	c.client.pubMsg(c.contractAddress, arbbridge.ChallengeCompletedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: c.client.getCurrentBlock(),
		},
		Winner:            winner,
		Loser:             loser,
		ChallengeContract: c.contractAddress,
	})
	return nil
}

type challengeWatcher struct {
	client *goEthdata
	*challenge
}

func newChallengeWatcher(address common.Address, client *goEthdata) (*challengeWatcher, error) {
	chalData := client.challengeFactory.challenges[address]

	return &challengeWatcher{challenge: chalData, client: client}, nil
}

func (c *challengeWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	c.client.goEthMutex.Lock()
	defer c.client.goEthMutex.Unlock()

	return c.client.blockMsgs[blockId].msgs[c.contractAddress], nil
}
