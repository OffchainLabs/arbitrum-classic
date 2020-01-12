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
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/challengetester"
)

type EthArbClient struct {
	client *ethclient.Client
}

func NewEthClient(ethURL string) (*EthArbClient, error) {
	client, err := ethclient.Dial(ethURL)
	return &EthArbClient{client}, err
}

func (c *EthArbClient) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return newArbFactoryWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return newRollupWatcher(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return newOneStepProof(address.ToEthAddress(), c.client)
}

func (c *EthArbClient) CurrentBlockTime(ctx context.Context) (*common.TimeBlocks, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return common.NewTimeBlocks(header.Number), nil
}

type EthArbAuthClient struct {
	*EthArbClient
	auth *bind.TransactOpts
}

func NewEthAuthClient(ethURL string, auth *bind.TransactOpts) (*EthArbAuthClient, error) {
	client, err := NewEthClient(ethURL)
	if err != nil {
		return nil, err
	}
	return &EthArbAuthClient{
		EthArbClient: client,
		auth:         auth,
	}, nil
}

func (c *EthArbAuthClient) Address() common.Address {
	return common.NewAddressFromEth(c.auth.From)
}

func (c *EthArbAuthClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return newArbFactory(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return newRollup(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return newPendingInbox(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return newChallengeFactory(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return newExecutionChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return newMessagesChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) NewPendingTopChallenge(address common.Address) (arbbridge.PendingTopChallenge, error) {
	return newPendingTopChallenge(address.ToEthAddress(), c.client, c.auth)
}

func (c *EthArbAuthClient) DeployChallengeTest() (*ChallengeTester, error) {
	testerAddress, tx, _, err := challengetester.DeployChallengeTester(c.auth, c.client)
	if err != nil {
		return nil, err
	}
	if err := waitForReceipt(
		context.Background(),
		c.client,
		c.auth.From,
		tx,
		"DeployChallengeTester",
	); err != nil {
		return nil, err
	}
	tester, err := NewChallengeTester(testerAddress, c.client, c.auth)
	if err != nil {
		return nil, err
	}
	return tester, nil
}
