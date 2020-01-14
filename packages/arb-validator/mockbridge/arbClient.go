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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type MockArbClient struct {
	MockEthClient *mockEthdata
}

func NewEthClient(ethURL string) (*MockArbClient, error) {
	// call to mockEth.go - getMockEth(ethURL)
	return &MockArbClient{getMockEth(ethURL)}, nil
}

func (c *MockArbClient) NewArbFactoryWatcher(address common.Address) (arbbridge.ArbFactoryWatcher, error) {
	return newArbFactoryWatcher(address, c)
}

func (c *MockArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return newRollupWatcher(address, c)
}

func (c *MockArbClient) NewExecutionChallengeWatcher(address common.Address) (arbbridge.ExecutionChallengeWatcher, error) {
	return newExecutionChallengeWatcher(address.ToEthAddress(), c)
}

func (c *MockArbClient) NewMessagesChallengeWatcher(address common.Address) (arbbridge.MessagesChallengeWatcher, error) {
	return newMessagesChallengeWatcher(address.ToEthAddress(), c)
}

func (c *MockArbClient) NewPendingTopChallengeWatcher(address common.Address) (arbbridge.PendingTopChallengeWatcher, error) {
	return newPendingTopChallengeWatcher(address.ToEthAddress(), c)
}

func (c *MockArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return newOneStepProof(address, c)
}

func (c *MockArbClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return newPendingInbox(address, c)
}

func (c *MockArbClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return newRollup(address, c)
}

func (c *MockArbClient) CurrentBlockTime(ctx context.Context) (*common.TimeBlocks, error) {
	return common.NewTimeBlocks(c.MockEthClient.LatestHeight), nil
}

func (c *MockArbClient) CurrentBlockTimeAndHash(ctx context.Context) (*common.TimeBlocks, common.Hash, error) {
	return common.NewTimeBlocks(c.MockEthClient.LatestHeight), c.MockEthClient.headerNumber[c.MockEthClient.LatestHeight], nil
}

type MockArbAuthClient struct {
	*MockArbClient
	auth *bind.TransactOpts
}

func NewEthAuthClient(ethURL string, auth *bind.TransactOpts) (*MockArbAuthClient, error) {
	client, err := NewEthClient(ethURL)
	if err != nil {
		return nil, err
	}
	return &MockArbAuthClient{
		MockArbClient: client,
		auth:          auth,
	}, nil
}

func (c *MockArbAuthClient) Address() common.Address {
	return common.NewAddressFromEth(c.auth.From)
}

func (c *MockArbAuthClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return newArbFactory(address, c)
}

func (c *MockArbAuthClient) NewRollup(address common.Address) (arbbridge.ArbRollup, error) {
	return newRollup(address, c)
}

func (c *MockArbAuthClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return newPendingInbox(address, c)
}

func (c *MockArbAuthClient) NewChallengeFactory(address common.Address) (arbbridge.ChallengeFactory, error) {
	return newChallengeFactory(address, c, c.auth)
}

func (c *MockArbAuthClient) NewExecutionChallenge(address common.Address) (arbbridge.ExecutionChallenge, error) {
	return NewExecutionChallenge(address, c)
}

func (c *MockArbAuthClient) NewMessagesChallenge(address common.Address) (arbbridge.MessagesChallenge, error) {
	return newMessagesChallenge(address, c)
}

func (c *MockArbAuthClient) NewPendingTopChallenge(address common.Address) (arbbridge.PendingTopChallenge, error) {
	return NewPendingTopChallenge(address, c)
}
