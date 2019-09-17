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
	"errors"
	"log"
	"math/big"
	"sync"
	"time"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type VMValidator struct {
	// Safe public interface
	VMID              common.Address
	CompletedCallChan chan valmessage.FinalizedAssertion

	Mutex *sync.Mutex
	// private thread only
	Validator               *Validator
	arbitrumVM              ethconnection.VMConnection
	unprocessedMessageCount uint64
}

func (val *VMValidator) Address() common.Address {
	return val.Validator.Address()
}

func NewVMValidator(
	val *Validator,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	con ethconnection.VMConnection,
) (*VMValidator, error) {
	callOpts := &bind.CallOpts{
		Pending: false,
		From:    val.Address(),
		Context: context.Background(),
	}

	err := con.VerifyVM(
		callOpts,
		config,
		machine.Hash(),
	)
	if err != nil {
		return nil, errors2.Wrap(err, "ChannelValidator failed to verify vm")
	}

	completedCallChan := make(chan valmessage.FinalizedAssertion, 1024)

	vmVal := &VMValidator{
		vmID,
		completedCallChan,
		&sync.Mutex{},
		val,
		con,
		0,
	}
	return vmVal, nil
}

func (val *VMValidator) ensureVMActivated() error {
	err := val.waitForActivation(context.Background())
	if err != nil {
		return errors2.Wrap(err, "Error checking for VM activation")
	}
	log.Println("ChannelValidator is validating vm", hexutil.Encode(val.VMID[:]))
	return nil
}

func (val *VMValidator) waitForActivation(
	ctx context.Context,
) error {
	auth := &bind.CallOpts{
		Pending: false,
		From:    val.Address(),
		Context: ctx,
	}

	for {
		select {
		case _ = <-time.After(time.Second):
			enabled, err := val.arbitrumVM.IsEnabled(auth)
			if err != nil {
				return err
			}
			if enabled {
				return nil
			}
		case _ = <-ctx.Done():
			return errors.New("VM never enabled")
		}
	}
}

func (val *VMValidator) Sign(msgHash [32]byte) ([]byte, error) {
	return val.Validator.Sign(msgHash[:])
}

func (val *VMValidator) StartListening(ctx context.Context) (chan ethconnection.Notification, error) {
	if err := val.ensureVMActivated(); err != nil {
		return nil, err
	}
	parsedChan := make(chan ethconnection.Notification, 1024)

	if err := val.arbitrumVM.StartConnection(ctx); err != nil {
		return nil, err
	}

	outChan, errChan := val.arbitrumVM.GetChans()
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
				// log.Printf("ChannelValidator recieved error: %v", err)
				// fmt.Println("Resetting channels")
				hitError = true

			}

			if hitError {
				// Ignore error and try to reset connection
				for {
					if err := val.arbitrumVM.StartConnection(ctx); err == nil {
						break
					}
					log.Println("Error: ChannelValidator can't connect to blockchain")
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()

	return parsedChan, nil
}

func (val *VMValidator) AddedNewMessages(count uint64) {
	go func() {
		val.Mutex.Lock()
		val.unprocessedMessageCount += count
		val.Mutex.Unlock()
	}()
}

func (val *VMValidator) FinalizedAssertion(
	assertion *protocol.Assertion,
	onChainTxHash []byte,
	signatures [][]byte,
	proposalResults *valmessage.UnanimousUpdateResults,
) {
	go func() {
		val.Mutex.Lock()
		finalizedAssertion := valmessage.FinalizedAssertion{
			Assertion:       assertion,
			OnChainTxHash:   onChainTxHash,
			Signatures:      signatures,
			ProposalResults: proposalResults,
		}
		val.unprocessedMessageCount -= uint64(len(finalizedAssertion.NewLogs()))
		val.CompletedCallChan <- finalizedAssertion
		val.Mutex.Unlock()
	}()
}

func (val *VMValidator) PendingDisputableAssert(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.PendingDisputableAssert(
			val.Validator.MakeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating disputable assertion")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) ConfirmDisputableAsserted(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.ConfirmDisputableAsserted(
			val.Validator.MakeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming disputable assertion")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) InitiateChallenge(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.InitiateChallenge(
			val.Validator.MakeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating challenge")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) BisectAssertion(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.BisectAssertion(
			val.Validator.MakeAuth(ctx),
			precondition,
			assertions,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating bisection")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) ContinueChallenge(
	ctx context.Context,
	assertionToChallenge uint16,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.ContinueChallenge(
			val.Validator.MakeAuth(ctx),
			assertionToChallenge,
			preconditions,
			assertions,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed continuing challenge")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) OneStepProof(
	ctx context.Context,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.OneStepProof(
			val.Validator.MakeAuth(ctx),
			precondition,
			assertion,
			proof,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed one step proof")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) AsserterTimedOut(
	ctx context.Context,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.AsserterTimedOutChallenge(val.Validator.MakeAuth(ctx))
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) ChallengerTimedOut(
	ctx context.Context,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.arbitrumVM.ChallengerTimedOutChallenge(
			val.Validator.MakeAuth(ctx),
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) SendMessage(
	ctx context.Context,
	data value.Value,
	tokenType [21]byte,
	currency *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Unlock()
		receipt, err := val.Validator.SendMessage(val.Validator.MakeAuth(ctx), protocol.NewSimpleMessage(data, tokenType, currency, val.VMID))
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) ForwardMessage(
	ctx context.Context,
	data value.Value,
	tokenType [21]byte,
	currency *big.Int,
	sig []byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.Validator.ForwardMessage(val.Validator.MakeAuth(ctx), protocol.NewSimpleMessage(data, tokenType, currency, val.VMID), sig)
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) SendEthMessage(
	ctx context.Context,
	data value.Value,
	amount *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.Mutex.Lock()
		receipt, err := val.Validator.SendEthMessage(val.Validator.MakeAuth(ctx), data, val.VMID, amount)
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.Mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) UnanimousAssertHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	return hashing.UnanimousAssertHash(
		val.VMID,
		sequenceNum,
		beforeHash,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
}
