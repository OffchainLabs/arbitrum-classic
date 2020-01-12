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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type EthArbClient struct {
	Client *ethclient.Client
}

func (c *EthArbClient) GetClient() *ethclient.Client {
	return c.Client
}

func NewEthClient(ethURL string) (*EthArbClient, error) {
	client, err := ethclient.Dial(ethURL)
	return &EthArbClient{client}, err
}

func (c *EthArbClient) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return NewArbFactoryWatcher(address.ToEthAddress(), c.Client)
}

func (c *EthArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return NewRollupWatcher(address.ToEthAddress(), c.Client)
}

func (c *EthArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return NewOneStepProof(address.ToEthAddress(), c.Client)
}

func (c *EthArbClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return NewPendingInbox(address.ToEthAddress(), c.Client)
}

func (c *EthArbClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.Client.HeaderByNumber(ctx, number)
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
	return NewArbFactory(address.ToEthAddress(), c.Client, c.auth)
}

func (c *EthArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return NewRollup(address.ToEthAddress(), c.Client, c.auth)
}

func (c *EthArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return NewChallengeFactory(address.ToEthAddress(), c.Client, c.auth)
}

func (c *EthArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return NewExecutionChallenge(address.ToEthAddress(), c.Client, c.auth)
}

func (c *EthArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return NewMessagesChallenge(address.ToEthAddress(), c.Client, c.auth)
}

func (c *EthArbAuthClient) NewPendingTopChallenge(address common.Address) (arbbridge.PendingTopChallenge, error) {
	return NewPendingTopChallenge(address.ToEthAddress(), c.Client, c.auth)
}
