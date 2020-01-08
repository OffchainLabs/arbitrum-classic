/*
 * Copyright 2019, Offchain Labs, Inc.
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

package arb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func NewArbFactory(address common.Address, client *ethclient.Client) (arbbridge.ArbFactory, error) {
	factory, err := ethbridge.NewArbFactory(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return factory, err
}

func NewBisectionChallenge(address common.Address, client *ethclient.Client, auth *bind.TransactOpts) (arbbridge.BisectionChallenge, error) {
	return NewBisectionChallenge(address, client, auth)
}

func NewOneStepProof(address common.Address, client *ethclient.Client) (arbbridge.OneStepProof, error) {
	oneStepProof, err := ethbridge.NewOneStepProof(address, client)
	if err != nil {
		return nil, err
	}
	return oneStepProof, err
}

func NewPendingInbox(address common.Address, client *ethclient.Client) (arbbridge.PendingInbox, error) {
	pendingInbox, err := ethbridge.NewPendingInbox(address, client)
	if err != nil {
		return nil, err
	}
	return pendingInbox, err
}

func NewPendingTopChallenge(address common.Address, client *ethclient.Client, auth *bind.TransactOpts) (arbbridge.PendingTopChallenge, error) {
	return ethbridge.NewPendingTopChallenge(address, client, auth)
}
