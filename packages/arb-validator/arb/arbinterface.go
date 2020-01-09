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

package arb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/mockbridge"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

var mock bool = false

//var mock bool = true

func NewArbFactory(address common.Address, client arbbridge.ArbClient) (arbbridge.ArbFactory, error) {
	var factory arbbridge.ArbFactory
	var err error
	if mock {
		factory, err = mockbridge.NewArbFactory(address, client)
	} else {
		factory, err = ethbridge.NewArbFactory(address, client)
		if err != nil {
			log.Fatal(err)
		}
	}
	return factory, err
}

func NewBisectionChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.BisectionChallenge, error) {
	return NewBisectionChallenge(address, client, auth)
}

func NewOneStepProof(address common.Address, client arbbridge.ArbClient) (arbbridge.OneStepProof, error) {
	oneStepProof, err := ethbridge.NewOneStepProof(address, client)
	if err != nil {
		return nil, err
	}
	return oneStepProof, err
}

func NewPendingInbox(address common.Address, client arbbridge.ArbClient) (arbbridge.PendingInbox, error) {
	pendingInbox, err := ethbridge.NewPendingInbox(address, client)
	if err != nil {
		return nil, err
	}
	return pendingInbox, err
}

func NewPendingTopChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.PendingTopChallenge, error) {
	return ethbridge.NewPendingTopChallenge(address, client, auth)
}

func NewRollup(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.ArbRollup, error) {
	return ethbridge.NewRollup(address, client, auth)
}

func NewExecutionChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.ExecutionChallenge, error) {
	return ethbridge.NewExecutionChallenge(address, client, auth)
}

func NewMessagesChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.MessagesChallenge, error) {
	return ethbridge.NewMessagesChallenge(address, client, auth)
}

func NewRollupWatcher(address common.Address, client arbbridge.ArbClient) (arbbridge.ArbRollupWatcher, error) {
	return ethbridge.NewRollupWatcher(address, client)
}

func NewArbClient() arbbridge.ArbClient {
	var client arbbridge.ArbClient
	if mock {
		//client = mockbridge.NewArbClient()
	} else {
		client = ethbridge.NewEthClient()
	}
	return client
}
