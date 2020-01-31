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

package ethbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/challengetester"
)

type ChallengeTester struct {
	contract *challengetester.ChallengeTester
	client   *ethclient.Client
	auth     *TransactAuth
}

func deployChallengeTester(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*ChallengeTester, error) {
	auth.Lock()
	defer auth.Unlock()
	testerAddress, tx, _, err := challengetester.DeployChallengeTester(auth.auth, client, address)
	if err != nil {
		return nil, err
	}
	if err := waitForReceipt(
		context.Background(),
		client,
		auth.auth.From,
		tx,
		"DeployChallengeTester",
	); err != nil {
		return nil, err
	}
	vmCreatorContract, err := challengetester.NewChallengeTester(testerAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeTester")
	}
	return &ChallengeTester{vmCreatorContract, client, auth}, nil
}

func (con *ChallengeTester) StartChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, *structures.BlockId, error) {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.contract.StartChallenge(
		con.auth.getAuth(ctx),
		asserter.ToEthAddress(),
		challenger.ToEthAddress(),
		challengePeriod.Val,
		challengeHash,
		challengeType,
	)
	if err != nil {
		return common.Address{}, nil, errors2.Wrap(err, "Failed to call to ChallengeTester.StartChallenge")
	}

	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "CreateChallenge")
	if err != nil {
		return common.Address{}, nil, err
	}

	if len(receipt.Logs) != 1 {
		return common.Address{}, nil, errors2.New("Wrong receipt count")
	}

	return common.NewAddressFromEth(receipt.Logs[0].Address), getTxBlockID(receipt), nil
}

//func (con *ChallengeTester) DeployChallengeTest(ctx context.Context, challengeFactory common.Address) (arbbridge.ChallengeTester, error) {
//	con.auth.Lock()
//	defer con.auth.Unlock()
//	testerAddress, tx, _, err := challengetester.DeployChallengeTester(con.auth.auth, con.client, challengeFactory.ToEthAddress())
//	if err != nil {
//		return nil, err
//	}
//	if err := waitForReceipt(
//		ctx,
//		con.client,
//		con.auth.auth.From,
//		tx,
//		"DeployChallengeTester",
//	); err != nil {
//		return nil, err
//	}
//	tester, err := c.DeployChallengeTester(common.NewAddressFromEth(testerAddress))
//	if err != nil {
//		return nil, err
//	}
//	return tester, nil
//}

func (c *EthArbAuthClient) DeployOneStepProof(ctx context.Context) (arbbridge.OneStepProof, error) {
	c.auth.Lock()
	defer c.auth.Unlock()
	ospAddress, tx, _, err := executionchallenge.DeployOneStepProof(c.auth.auth, c.client)
	if err != nil {
		return nil, err
	}
	if err := waitForReceipt(
		ctx,
		c.client,
		c.auth.auth.From,
		tx,
		"DeployOneStepProof",
	); err != nil {
		return nil, err
	}
	osp, err := c.NewOneStepProof(common.NewAddressFromEth(ospAddress))
	if err != nil {
		return nil, err
	}
	return osp, nil
}
