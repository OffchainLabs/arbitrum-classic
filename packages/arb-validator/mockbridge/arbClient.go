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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type ArbClient struct {
	client arbbridge.ArbClient
}

//func (c *ArbClient) GetClient() *ethclient.Client {
//	return c.client
//}

func NewEthClient(ethURL string) (*ArbClient, error) {
	// call to mockEth.go - initMock(ethURL)
	//client, err := ethclient.Dial(ethURL)
	return &ArbClient{nil}, nil
}

func (c *ArbClient) NewArbFactory(address common.Address) (arbbridge.ArbFactory, error) {
	return NewArbFactory(address, c.client)
}

func (c *ArbClient) NewRollup(address common.Address, auth *bind.TransactOpts) (arbbridge.ArbRollup, error) {
	return NewRollup(address, c.client, auth)
}

func (c *ArbClient) NewRollupWatcher(address common.Address) (arbbridge.ArbRollupWatcher, error) {
	return NewRollupWatcher(address, c.client)
}

func (c *ArbClient) NewExecutionChallenge(address common.Address, auth *bind.TransactOpts) (arbbridge.ExecutionChallenge, error) {
	return NewExecutionChallenge(address, c.client, auth)
}

func (c *ArbClient) NewMessagesChallenge(address common.Address, auth *bind.TransactOpts) (arbbridge.MessagesChallenge, error) {
	return NewMessagesChallenge(address, c.client, auth)
}

func (c *ArbClient) NewOneStepProof(address common.Address) (arbbridge.OneStepProof, error) {
	return NewOneStepProof(address, c.client)
}

func (c *ArbClient) NewPendingInbox(address common.Address) (arbbridge.PendingInbox, error) {
	return NewPendingInbox(address, c.client)
}

func (c *ArbClient) NewPendingTopChallenge(address common.Address, auth *bind.TransactOpts) (arbbridge.PendingTopChallenge, error) {
	return NewPendingTopChallenge(address, c.client, auth)
}

func (c *ArbClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.client.HeaderByNumber(ctx, number)
}
