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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

var initiatedChallengeID ethcommon.Hash
var timedOutAsserterID ethcommon.Hash
var timedOutChallengerID ethcommon.Hash

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
		challengeData:   client.GoEthClient.challenges[address],
	}

	return &chal, nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {

	if common.TicksFromBlockNum(c.client.GoEthClient.getCurrentBlock().Height).Cmp(c.challengeData.deadline) < 0 {
		return errors.New("Deadline hasn't expired")
	}
	if c.challengeData.state == asserterTurn {
		c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
			Event: arbbridge.AsserterTimeoutEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: c.client.GoEthClient.getCurrentBlock(),
				},
			},
		})
	} else {
		c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
			Event: arbbridge.ChallengerTimeoutEvent{
				ChainInfo: arbbridge.ChainInfo{
					BlockId: c.client.GoEthClient.getCurrentBlock(),
				},
			},
		})
	}
	return nil
}

type challengeWatcher struct {
	client    *GoArbClient
	Challenge common.Address
}

func newChallengeWatcher(address common.Address, client *GoArbClient) (*challengeWatcher, error) {

	return &challengeWatcher{client: client, Challenge: address}, nil
}

func (c *challengeWatcher) topics() []ethcommon.Hash {
	return []ethcommon.Hash{
		initiatedChallengeID,
		timedOutAsserterID,
		timedOutChallengerID,
	}
}
