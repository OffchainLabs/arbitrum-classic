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

package ethvalidator

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChallengeValidator struct {
	Mutex *sync.Mutex
	// private thread only
	Validator      *Validator
	challenge      ethbridge.ChallengeConnection
	MessageMonChan chan bridge.BridgeMessage
	ErrorMonChan   chan bridge.Error
}

func (val *ChallengeValidator) Address() common.Address {
	return val.Validator.Address()
}

func (val *ChallengeValidator) SendMonitorMsg(msg bridge.BridgeMessage) {
	val.MessageMonChan <- msg
}

func (val *ChallengeValidator) SendMonitorErr(msg bridge.Error) {
	val.ErrorMonChan <- msg
}

func NewChallengeValidator(
	val *Validator,
	challengeContract common.Address,
) (*ChallengeValidator, error) {
	con, err := ethbridge.NewChallenge(challengeContract, val.Client)
	if err != nil {
		return nil, err
	}

	msgmon := make(chan bridge.BridgeMessage, 100)
	errmon := make(chan bridge.Error, 100)
	vmVal := &ChallengeValidator{
		&sync.Mutex{},
		val,
		con,
		msgmon,
		errmon,
	}
	return vmVal, nil
}

func (val *ChallengeValidator) StartListening(ctx context.Context) (chan ethbridge.Notification, error) {
	parsedChan := make(chan ethbridge.Notification, 1024)

	if err := val.challenge.StartConnection(ctx); err != nil {
		return nil, err
	}

	outChan, errChan := val.challenge.GetChans()
	go func() {
		for {
			hitError := false
			select {
			case <-ctx.Done():
				break
			case parse, ok := <-outChan:
				if !ok {
					hitError = true
					break
				}
				parsedChan <- parse
			case <-errChan:
				// log.Printf("Validator recieved error: %v\n", err)
				// fmt.Println("Resetting channels")
				hitError = true

			}

			if hitError {
				// Ignore error and try to reset connection
				for {
					if err := val.challenge.StartConnection(ctx); err == nil {
						break
					}
					log.Println("Error: Validator can't connect to blockchain")
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()

	return parsedChan, nil
}

func (val *ChallengeValidator) BisectAssertion(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertions []*protocol.ExecutionAssertionStub,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	receipt, err := val.challenge.BisectAssertion(
		val.Validator.MakeAuth(ctx),
		precondition,
		assertions,
	)
	val.Mutex.Unlock()
	return receipt, err
}

func (val *ChallengeValidator) ContinueChallenge(
	ctx context.Context,
	assertionToChallenge uint16,
	precondition *protocol.Precondition,
	assertions []*protocol.ExecutionAssertionStub,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	receipt, err := val.challenge.ContinueChallenge(
		val.Validator.MakeAuth(ctx),
		assertionToChallenge,
		precondition,
		assertions,
	)
	val.Mutex.Unlock()
	return receipt, err
}

func (val *ChallengeValidator) OneStepProof(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.ExecutionAssertionStub,
	proof []byte,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	receipt, err := val.challenge.OneStepProof(
		val.Validator.MakeAuth(ctx),
		precondition,
		assertion,
		proof,
	)
	val.Mutex.Unlock()
	return receipt, err
}

func (val *ChallengeValidator) AsserterTimedOut(
	ctx context.Context,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	receipt, err := val.challenge.AsserterTimedOutChallenge(val.Validator.MakeAuth(ctx))
	val.Mutex.Unlock()
	return receipt, err
}

func (val *ChallengeValidator) ChallengerTimedOut(
	ctx context.Context,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	receipt, err := val.challenge.ChallengerTimedOutChallenge(
		val.Validator.MakeAuth(ctx),
	)
	val.Mutex.Unlock()
	return receipt, err
}
