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
		c.client.pubMsg(c, arbbridge.MaybeEvent{
			Event: arbbridge.AsserterTimeoutEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: c.client.getCurrentBlock(),
				},
			},
		})
	} else {
		c.client.pubMsg(c, arbbridge.MaybeEvent{
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
	client *goEthdata
	*challenge
	Challenge common.Address
}

func newChallengeWatcher(address common.Address, client *goEthdata) (*challengeWatcher, error) {
	chalData := client.challenges[address]
	if _, ok := client.challengeWatcherEvents[chalData]; !ok {
		client.challengeWatcherEvents[chalData] = make(map[*common.BlockId][]arbbridge.Event)
	}

	return &challengeWatcher{challenge: chalData, client: client}, nil
}

func (c *challengeWatcher) GetEvents(ctx context.Context, blockID *common.BlockId) ([]arbbridge.Event, error) {
	c.client.goEthMutex.Lock()
	defer c.client.goEthMutex.Unlock()
	cw := c.client.challengeWatcherEvents[c.challenge][blockID]
	return cw, nil
}
