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

package mockbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"
)

type MockArbClient struct {
	MockEthClient *mockEthdata
}

func NewEthClient(ethURL string) (*MockArbClient, error) {
	// call to mockEth.go - getMockEth(ethURL)
	return &MockArbClient{getMockEth(ethURL)}, nil
}

func (c *MockArbClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return NewArbFactory(address, c)
}

func (c *MockArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return NewRollupWatcher(address, c)
}

func (c *MockArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return NewOneStepProof(address, c)
}

func (c *MockArbClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return NewPendingInbox(address, c)
}

func (c *MockArbClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.MockEthClient.headerNumber[number], nil
}

type ArbAuthClient struct {
	*MockArbClient
	auth *bind.TransactOpts
}

func (c *ArbAuthClient) Address() common.Address {
	return c.auth.From
}

func NewEthAuthClient(ethURL string, auth *bind.TransactOpts) (*ArbAuthClient, error) {
	client, err := NewEthClient(ethURL)
	if err != nil {
		return nil, err
	}
	return &ArbAuthClient{
		MockArbClient: client,
		auth:          auth,
	}, nil
}

func (c *ArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return NewRollup(address, c, c.auth)
}

func (c *ArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return NewExecutionChallenge(address, c, c.auth)
}

func (c *ArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return NewMessagesChallenge(address, c, c.auth)
}

func (c *ArbAuthClient) NewPendingTopChallenge(address common.Address) (arbbridge.PendingTopChallenge, error) {
	return NewPendingTopChallenge(address, c, c.auth)
}
