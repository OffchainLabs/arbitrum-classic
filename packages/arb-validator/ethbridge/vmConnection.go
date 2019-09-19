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

package ethbridge

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type VMConnection interface {
	StartConnection(ctx context.Context) error

	GetChans() (chan Notification, chan error)

	VerifyVM(
		auth *bind.CallOpts,
		config *valmessage.VMConfiguration,
		machine [32]byte,
	) error

	IsEnabled(
		auth *bind.CallOpts,
	) (bool, error)

	IsPendingUnanimous(
		auth *bind.CallOpts,
	) (bool, error)

	IsInChallenge(
		auth *bind.CallOpts,
	) (bool, error)

	PendingDisputableAssert(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	) (*types.Receipt, error)

	ConfirmDisputableAsserted(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	) (*types.Receipt, error)

	InitiateChallenge(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.AssertionStub,
	) (*types.Receipt, error)

	BisectAssertion(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertions []*protocol.AssertionStub,
	) (*types.Receipt, error)

	ContinueChallenge(
		auth *bind.TransactOpts,
		assertionToChallenge uint16,
		precondition *protocol.Precondition,
		assertions []*protocol.AssertionStub,
	) (*types.Receipt, error)

	OneStepProof(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.AssertionStub,
		proof []byte,
	) (*types.Receipt, error)

	AsserterTimedOutChallenge(
		auth *bind.TransactOpts,
	) (*types.Receipt, error)

	ChallengerTimedOutChallenge(
		auth *bind.TransactOpts,
	) (*types.Receipt, error)
}
