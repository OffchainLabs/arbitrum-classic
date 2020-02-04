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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChallengeTester struct {
	contract common.Address
	client   *MockArbAuthClient
}

func newChallengeTester(address common.Address, client *MockArbAuthClient) (*ChallengeTester, error) {
	//vmCreatorContract, err := challengetester.DeployChallengeTester(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ChallengeTester")
	//}
	return &ChallengeTester{client.MockEthClient.getNextAddress(), client}, nil
}

func (con *ChallengeTester) StartChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, *structures.BlockId, error) {
	eth := con.client.MockEthClient
	//con.auth.Context = ctx
	//tx, err := con.contract.StartChallenge(
	//	con.auth,
	//	factory.ToEthAddress(),
	//	asserter.ToEthAddress(),
	//	challenger.ToEthAddress(),
	//	challengePeriod.Val,
	//	challengeHash,
	//	challengeType,
	//)
	//if err != nil {
	//	return common.Address{}, errors2.Wrap(err, "Failed to call to ChallengeTester.StartChallenge")
	//}
	//
	//receipt, err := WaitForReceiptWithResults(con.auth.Context, con.client, con.auth.From, tx, "CreateChallenge")
	//if err != nil {
	//	return common.Address{}, err
	//}
	//
	//if len(receipt.Logs) != 1 {
	//	return common.Address{}, errors2.New("Wrong receipt count")
	//}

	// create clone
	newAddr := eth.getNextAddress()
	eth.challenges[newAddr] = &challengeData{deadline: challengePeriod}

	//initializeBisection
	eth.challenges[newAddr].deadline = common.TimeFromBlockNum(eth.LastMinedBlock.Height).Add(challengePeriod)
	// emit InitiatedChallenge
	InitiateChallengeEvent := arbbridge.InitiateChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: eth.getCurrentBlock(),
		},
		Deadline: eth.challenges[newAddr].deadline,
	}
	fmt.Println("challengeTester - publishing InitiateChallengeEvent")
	eth.pubMsg(arbbridge.MaybeEvent{
		Event: InitiateChallengeEvent,
	})
	//return clone address
	return newAddr, eth.getLastBlock(), nil
}
