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
	"math/big"

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

func NewChallengeTester(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*ChallengeTester, error) {
	vmCreatorContract, err := challengetester.NewChallengeTester(address, client)
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
) (common.Address, *common.BlockId, error) {
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
