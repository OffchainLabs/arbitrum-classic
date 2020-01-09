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

func NewArbClient() arbbridge.ArbClient {
	var client arbbridge.ArbClient
	if mock {
		//client = mockbridge.NewArbClient()
	} else {
		client = ethbridge.NewEthClient()
	}
	return client
}

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

func NewRollup(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.ArbRollup, error) {
	var rollup arbbridge.ArbRollup
	var err error
	if mock {
		rollup, err = mockbridge.NewRollup(address, client, auth)
	} else {
		rollup, err = ethbridge.NewRollup(address, client, auth)
		if err != nil {
			log.Fatal(err)
		}
	}
	return rollup, err
}

func NewRollupWatcher(address common.Address, client arbbridge.ArbClient) (arbbridge.ArbRollupWatcher, error) {
	var rollupWatcher arbbridge.ArbRollupWatcher
	var err error
	if mock {
		rollupWatcher, err = mockbridge.NewRollupWatcher(address, client)
	} else {
		rollupWatcher, err = ethbridge.NewRollupWatcher(address, client)
		if err != nil {
			log.Fatal(err)
		}
	}
	return rollupWatcher, err
}

func NewExecutionChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.ExecutionChallenge, error) {
	var executionChallenge arbbridge.ExecutionChallenge
	var err error
	if mock {
		executionChallenge, err = mockbridge.NewExecutionChallenge(address, client, auth)
	} else {
		executionChallenge, err = ethbridge.NewExecutionChallenge(address, client, auth)
		if err != nil {
			log.Fatal(err)
		}
	}
	return executionChallenge, err
}

func NewMessagesChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.MessagesChallenge, error) {
	var messagesChallenge arbbridge.MessagesChallenge
	var err error
	if mock {
		messagesChallenge, err = mockbridge.NewMessagesChallenge(address, client, auth)
	} else {
		messagesChallenge, err = ethbridge.NewMessagesChallenge(address, client, auth)
		if err != nil {
			log.Fatal(err)
		}
	}
	return messagesChallenge, err
}

func NewOneStepProof(address common.Address, client arbbridge.ArbClient) (arbbridge.OneStepProof, error) {
	var oneStepProof arbbridge.OneStepProof
	var err error
	if mock {
		oneStepProof, err = mockbridge.NewOneStepProof(address, client)
	} else {
		oneStepProof, err = ethbridge.NewOneStepProof(address, client)
		if err != nil {
			return nil, err
		}
	}
	return oneStepProof, err
}

func NewPendingInbox(address common.Address, client arbbridge.ArbClient) (arbbridge.PendingInbox, error) {
	var pendingInbox arbbridge.PendingInbox
	var err error
	if mock {
		pendingInbox, err = mockbridge.NewPendingInbox(address, client)
	} else {
		pendingInbox, err = ethbridge.NewPendingInbox(address, client)
		if err != nil {
			return nil, err
		}
	}
	return pendingInbox, err
}

func NewPendingTopChallenge(address common.Address, client arbbridge.ArbClient, auth *bind.TransactOpts) (arbbridge.PendingTopChallenge, error) {
	var pendingTopChallenge arbbridge.PendingTopChallenge
	var err error
	if mock {
		pendingTopChallenge, err = mockbridge.NewPendingTopChallenge(address, client, auth)
	} else {
		pendingTopChallenge, err = ethbridge.NewPendingTopChallenge(address, client, auth)
		if err != nil {
			log.Fatal(err)
		}
	}
	return pendingTopChallenge, err
}
