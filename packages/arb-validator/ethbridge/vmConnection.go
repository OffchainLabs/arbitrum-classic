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
	"log"
	"time"

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
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	ConfirmDisputableAsserted(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	InitiateChallenge(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)
}

type ChallengeConnection interface {
	StartConnection(ctx context.Context) error

	GetChans() (chan Notification, chan error)

	BisectAssertion(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	ContinueChallenge(
		auth *bind.TransactOpts,
		assertionToChallenge uint16,
		precondition *protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	OneStepProof(
		auth *bind.TransactOpts,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
		proof []byte,
	) (*types.Receipt, error)

	AsserterTimedOutChallenge(
		auth *bind.TransactOpts,
	) (*types.Receipt, error)

	ChallengerTimedOutChallenge(
		auth *bind.TransactOpts,
	) (*types.Receipt, error)
}

type ContractConnection interface {
	StartConnection(context.Context, chan Notification, chan error) error
}

type ChallengeContract interface {
	TimeoutChallenge(
		ctx context.Context,
	) (*types.Receipt, error)
}

func HandleBlockchainNotifications(ctx context.Context, noteChan chan Notification, contract ContractConnection) {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)
	defer close(outChan)
	defer close(errChan)
	if err := contract.StartConnection(ctx, outChan, errChan); err != nil {
		return
	}
	for {
		hitError := false
		select {
		case <-ctx.Done():
			break
		case notification, ok := <-outChan:
			if !ok {
				hitError = true
				break
			}
			noteChan <- notification
		case <-errChan:
			hitError = true
		}

		if hitError {
			// Ignore error and try to reset connection
			for {
				if err := contract.StartConnection(ctx, outChan, errChan); err == nil {
					break
				}
				log.Println("Error: Can't connect to blockchain")
				time.Sleep(5 * time.Second)
			}
		}
	}
}
