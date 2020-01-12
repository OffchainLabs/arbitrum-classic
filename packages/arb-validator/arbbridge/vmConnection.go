/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package arbbridge

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type VMConnection interface {
	StartConnection(ctx context.Context) error

	GetChans() (chan Notification, chan error)

	VerifyVM(
		auth *bind.CallOpts,
		config *valmessage.VMConfiguration,
		machine common.Hash,
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
		precondition *valprotocol.Precondition,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	ConfirmDisputableAsserted(
		auth *bind.TransactOpts,
		precondition *valprotocol.Precondition,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	InitiateChallenge(
		auth *bind.TransactOpts,
		precondition *valprotocol.Precondition,
		assertion *valprotocol.ExecutionAssertionStub,
	) (*types.Receipt, error)
}

type ChallengeConnection interface {
	StartConnection(ctx context.Context) error

	GetChans() (chan Notification, chan error)

	BisectAssertion(
		auth *bind.TransactOpts,
		precondition *valprotocol.Precondition,
		assertions []*valprotocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	ContinueChallenge(
		auth *bind.TransactOpts,
		assertionToChallenge uint16,
		precondition *valprotocol.Precondition,
		assertions []*valprotocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	OneStepProof(
		auth *bind.TransactOpts,
		precondition *valprotocol.Precondition,
		assertion *valprotocol.ExecutionAssertionStub,
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

type ChainContract interface {
	CurrentBlockTime(ctx context.Context) (*common.TimeBlocks, error)
}

type ChallengeContract interface {
	ChainContract

	TimeoutChallenge(ctx context.Context) error
}

func HandleBlockchainNotifications(ctx context.Context, contract ContractConnection) chan Notification {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)
	if err := contract.StartConnection(ctx, outChan, errChan); err != nil {
		close(outChan)
		close(errChan)
		return nil
	}

	noteChan := make(chan Notification, 1024)
	go func() {
		defer close(outChan)
		defer close(errChan)
		defer close(noteChan)
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
	}()
	return noteChan
}
