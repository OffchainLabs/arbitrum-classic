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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	"github.com/pkg/errors"

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
		return nil, errors.Wrap(err, "Failed to connect to arbFactory")
	}
	return &challengeFactory{vmCreatorContract, client, auth}, nil
}

func DeployChallengeFactory(ctx context.Context, authClient *EthArbAuthClient, client ethutils.EthClient) (ethcommon.Address, *types.Transaction, error) {
	inboxTopAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployInboxTopChallenge(auth, client)
	})
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	executionAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployExecutionChallenge(auth, client)
	})
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	ospAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployOneStepProof(auth, client)
	})
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	factoryAddr, tx, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgecontracts.DeployChallengeFactory(auth, client, inboxTopAddr, executionAddr, ospAddr)
	})
	return factoryAddr, tx, err
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
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.contract.CreateChallenge(
			auth,
			asserter.ToEthAddress(),
			challenger.ToEthAddress(),
			challengePeriod.Val,
			challengeHash,
			challengeType,
		)
	})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to call to challengeFactory.CreateChallenge")
	}

	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "CreateChallenge")
	if err != nil {
		return common.Address{}, err
	}

	if len(receipt.Logs) != 1 {
		return common.Address{}, errors.New("Wrong receipt count")
	}

	return common.NewAddressFromEth(receipt.Logs[0].Address), nil
}
