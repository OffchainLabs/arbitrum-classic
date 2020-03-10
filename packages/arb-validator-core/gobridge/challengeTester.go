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
	fact, _ := newChallengeFactory(client.getNextAddress(), client)
	return &ChallengeTester{fact.challengeFactoryContract, client}, nil
}

func (con *ChallengeTester) StartChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, *common.BlockId, error) {
	eth := con.client
	eth.goEthMutex.Lock()
	defer eth.goEthMutex.Unlock()

	// create clone
	newAddr, _ := eth.challengeFactoryContract.createChallenge(
		ctx,
		asserter,
		challenger,
		challengePeriod,
		challengeHash,
		challengeType)

	// emit InitiatedChallenge
	InitiateChallengeEvent := arbbridge.InitiateChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: eth.getCurrentBlock(),
		},
		Deadline: eth.challenges[newAddr].challengeData.deadline,
	}
	eth.pubMsg(newAddr, InitiateChallengeEvent)
	return newAddr, eth.getLastBlock(), nil
}
