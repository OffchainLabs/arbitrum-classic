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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChallengeTester struct {
	contract common.Address
	client   *GoArbAuthClient
}

func NewChallengeTester(client *GoArbAuthClient) (*ChallengeTester, error) {
	return &ChallengeTester{client.GoEthClient.getNextAddress(), client}, nil
}

func (con *ChallengeTester) StartChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, *common.BlockId, error) {
	eth := con.client.GoEthClient

	// create clone
	newAddr := eth.getNextAddress()
	eth.challenges[newAddr] = &challengeData{deadline: challengePeriod, challengerDataHash: challengeHash, challengePeriodTicks: challengePeriod}

	//initializeBisection
	eth.challenges[newAddr].deadline = common.TicksFromBlockNum(eth.LastMinedBlock.Height).Add(challengePeriod)
	eth.challenges[newAddr].state = asserterTurn
	// emit InitiatedChallenge
	InitiateChallengeEvent := arbbridge.InitiateChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: eth.getCurrentBlock(),
		},
		Deadline: eth.challenges[newAddr].deadline,
	}
	eth.pubMsg(eth.challenges[newAddr], arbbridge.MaybeEvent{
		Event: InitiateChallengeEvent,
	})
	//return clone address
	return newAddr, eth.getLastBlock(), nil
}
