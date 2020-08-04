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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type challengeFactory struct {
	contract *ethbridgecontracts.ChallengeFactory
	client   ethutils.EthClient
	auth     *TransactAuth
}

func newChallengeFactory(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*challengeFactory, error) {
	vmCreatorContract, err := ethbridgecontracts.NewChallengeFactory(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	}
	return &challengeFactory{vmCreatorContract, client, auth}, nil
}

func DeployChallengeFactory(auth *bind.TransactOpts, client ethutils.EthClient) (ethcommon.Address, error) {
	inboxTopAddr, _, _, err := ethbridgecontracts.DeployInboxTopChallenge(auth, client)
	if err != nil {
		return ethcommon.Address{}, err
	}
	executionAddr, _, _, err := ethbridgecontracts.DeployExecutionChallenge(auth, client)
	if err != nil {
		return ethcommon.Address{}, err
	}
	factoryAddr, _, _, err := ethbridgecontracts.DeployChallengeFactory(auth, client, inboxTopAddr, executionAddr)
	return factoryAddr, err
}

func (con *challengeFactory) CreateChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, error) {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.contract.CreateChallenge(
		con.auth.getAuth(ctx),
		asserter.ToEthAddress(),
		challenger.ToEthAddress(),
		challengePeriod.Val,
		challengeHash,
		challengeType,
	)
	if err != nil {
		return common.Address{}, errors2.Wrap(err, "Failed to call to challengeFactory.CreateChallenge")
	}

	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "CreateChallenge")
	if err != nil {
		return common.Address{}, err
	}

	if len(receipt.Logs) != 1 {
		return common.Address{}, errors2.New("Wrong receipt count")
	}

	return common.NewAddressFromEth(receipt.Logs[0].Address), nil
}
