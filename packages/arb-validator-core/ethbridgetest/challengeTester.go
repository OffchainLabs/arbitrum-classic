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

package ethbridgetest

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/challengetester"
)

type ChallengeTester struct {
	Address  ethcommon.Address
	contract *challengetester.ChallengeTester
	client   *ethclient.Client
	auth     *bind.TransactOpts
}

func DeployChallengeTest(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, challengeFactory common.Address) (*ChallengeTester, error) {
	testerAddress, tx, _, err := challengetester.DeployChallengeTester(auth, client, challengeFactory.ToEthAddress())
	if err != nil {
		return nil, err
	}
	if _, err := ethbridge.WaitForReceiptWithResults(
		ctx,
		client,
		auth.From,
		tx,
		"DeployChallengeTester",
	); err != nil {
		return nil, err
	}
	tester, err := NewChallengeTester(testerAddress, client, auth)
	if err != nil {
		return nil, err
	}
	return tester, nil
}

func NewChallengeTester(address ethcommon.Address, client *ethclient.Client, auth *bind.TransactOpts) (*ChallengeTester, error) {
	vmCreatorContract, err := challengetester.NewChallengeTester(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeTester")
	}
	return &ChallengeTester{address, vmCreatorContract, client, auth}, nil
}

func (con *ChallengeTester) StartChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, *common.BlockId, error) {
	con.auth.Context = ctx
	tx, err := con.contract.StartChallenge(
		con.auth,
		asserter.ToEthAddress(),
		challenger.ToEthAddress(),
		challengePeriod.Val,
		challengeHash,
		challengeType,
	)
	if err != nil {
		return common.Address{}, nil, errors2.Wrap(err, "Failed to call to ChallengeTester.StartChallenge")
	}

	receipt, err := ethbridge.WaitForReceiptWithResults(ctx, con.client, con.auth.From, tx, "CreateChallenge")
	if err != nil {
		return common.Address{}, nil, err
	}

	if len(receipt.Logs) != 1 {
		return common.Address{}, nil, errors2.New("Wrong receipt count")
	}

	return common.NewAddressFromEth(receipt.Logs[0].Address), ethbridge.GetReceiptBlockID(receipt), nil
}
