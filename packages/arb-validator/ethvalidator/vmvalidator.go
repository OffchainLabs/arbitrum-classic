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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/channelvalidator"

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
	Validators        map[common.Address]validatorInfo
	Bot               *channelvalidator.Validator
	CompletedCallChan chan valmessage.FinalizedAssertion

	mutex *sync.Mutex
	// private thread only
	validator               *Validator
	vmTracker               *ethconnection.ArbChannel
	unprocessedMessageCount uint64
}

func (val *VMValidator) Address() common.Address {
	return val.validator.Address()
}

func (val *VMValidator) ValidatorCount() int {
	return len(val.Validators)
}

type validatorInfo struct {
	indexNum uint16
}

type VMResponse struct {
	Message protocol.Message
	Result  value.Value
	Proof   [][32]byte
}

func NewVMValidator(
	name string,
	val *Validator,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
) (*VMValidator, error) {
	con, err := ethconnection.NewVMTracker(vmID, val.client)
	if err != nil {
		return nil, errors2.Wrap(err, "VMValidator failed to create NewVMTracker")
	}

	callOpts := &bind.CallOpts{
		Pending: false,
		From:    val.Address(),
		Context: context.Background(),
	}
	err = con.VerifyVM(
		callOpts,
		config,
		machine.Hash(),
	)
	if err != nil {
		return nil, errors2.Wrap(err, "VMValidator failed to verify vm")
	}

	manMap := make(map[common.Address]validatorInfo)
	keys := make([]common.Address, 0, len(config.AssertKeys))
	for _, key := range config.AssertKeys {
		var address common.Address
		copy(address[:], key.Value)
		keys = append(keys, address)
	}
	for i, add := range keys {
		manMap[add] = validatorInfo{uint16(i)}
	}

	_, found := manMap[val.Address()]
	if !found {
		return nil, errors.New("key is not a validator of chosen ArbChannel")
	}

	header, err := val.LatestHeader(context.Background())
	if err != nil {
		return nil, errors2.Wrap(err, "VMValidator couldn't get latest error")
	}

	bot := channelvalidator.NewValidator(
		name,
		val.Address(),
		header,
		protocol.NewBalanceTracker(),
		config,
		machine,
		challengeEverything,
		maxCallSteps,
	)

	completedCallChan := make(chan valmessage.FinalizedAssertion, 1024)

	vmVal := &VMValidator{
		vmID,
		manMap,
		bot,
		completedCallChan,
		&sync.Mutex{},
		val,
		con,
		0,
	}
	if err := vmVal.topOffDeposit(context.Background()); err != nil {
		return nil, errors2.Wrap(err, "VMValidator failed to top off deposit")
	}
	return vmVal, nil
}

func (val *VMValidator) ensureVMActivated() error {
	err := val.waitForActivation(context.Background())
	if err != nil {
		return errors2.Wrap(err, "Error checking for VM activation")
	}
	log.Println("Validator is validating vm", hexutil.Encode(val.VMID[:]))
	return nil
}

func (val *VMValidator) topOffDeposit(ctx context.Context) error {
	callOpts := &bind.CallOpts{
		Pending: true,
		From:    val.Address(),
		Context: context.Background(),
	}
	current, err := val.vmTracker.CurrentDeposit(callOpts, val.Address())
	if err != nil {
		return err
	}
	required, err := val.vmTracker.EscrowRequired(callOpts)
	if current.Cmp(required) >= 0 {
		// Validator already has escrow deposited
		return nil
	}
	depToAdd := new(big.Int).Sub(required, current)
	_, err = val.vmTracker.IncreaseDeposit(val.validator.makeAuth(ctx), depToAdd)
	if err != nil {
		return errors2.Wrap(err, "failed calling IncreaseDeposit")
	}
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
			enabled, err := val.vmTracker.IsEnabled(auth)
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
	return val.validator.Sign(msgHash[:])
}

func (val *VMValidator) restartConnection(ctx context.Context) (chan ethconnection.Notification, chan error, error) {
	vmCon, err := ethconnection.NewVMTracker(val.VMID, val.validator.client)
	if err != nil {
		return nil, nil, err
	}
	val.vmTracker = vmCon
	return val.vmTracker.CreateListeners(ctx)
}

func (val *VMValidator) StartListening(ctx context.Context) error {
	if err := val.ensureVMActivated(); err != nil {
		return err
	}
	outChan, errChan, err := val.vmTracker.CreateListeners(ctx)
	if err != nil {
		return err
	}
	parsedChan := make(chan ethconnection.Notification, 1024)

	go func() {
		val.Bot.Run(parsedChan, val, ctx)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case parse, ok := <-outChan:
				if !ok {
					outChan, errChan, err = val.restartConnection(ctx)
					if err != nil {
						panic(err)
					}
					break
				}
				parsedChan <- parse
			case <-errChan:
				// Ignore error and try to reset connection
				// log.Printf("Validator recieved error: %v", err)
				// fmt.Println("Resetting channels")
				for {
					outChan, errChan, err = val.restartConnection(ctx)
					if err == nil {
						break
					}
					log.Println("Error: Validator can't connect to blockchain")
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()

	return nil
}

func (val *VMValidator) AddedNewMessages(count uint64) {
	go func() {
		val.mutex.Lock()
		val.unprocessedMessageCount += count
		val.mutex.Unlock()
	}()
}

func (val *VMValidator) FinalizedAssertion(
	assertion *protocol.Assertion,
	onChainTxHash []byte,
	signatures [][]byte,
	proposalResults *valmessage.UnanimousUpdateResults,
) {
	go func() {
		val.mutex.Lock()
		finalizedAssertion := valmessage.FinalizedAssertion{
			Assertion:       assertion,
			OnChainTxHash:   onChainTxHash,
			Signatures:      signatures,
			ProposalResults: proposalResults,
		}
		val.unprocessedMessageCount -= uint64(len(finalizedAssertion.NewLogs()))
		val.CompletedCallChan <- finalizedAssertion
		val.mutex.Unlock()
	}()
}

func (val *VMValidator) FinalizedUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	signatures [][]byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.mutex.Lock()
		receipt, err := val.vmTracker.FinalizedUnanimousAssert(
			val.validator.makeAuth(ctx),
			newInboxHash,
			assertion,
			signatures,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed sending finalized unanimous assertion")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) PendingUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures [][]byte,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.mutex.Lock()
		receipt, err := val.vmTracker.PendingUnanimousAssert(
			val.validator.makeAuth(ctx),
			newInboxHash,
			assertion,
			sequenceNum,
			signatures,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed proposing unanimous assertion")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) ConfirmUnanimousAsserted(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.mutex.Lock()
		receipt, err := val.vmTracker.ConfirmUnanimousAsserted(
			val.validator.makeAuth(ctx),
			newInboxHash,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming unanimous assertion")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
	}()
	return receiptChan, errChan
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.PendingDisputableAssert(
			val.validator.makeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating disputable assertion")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.ConfirmDisputableAsserted(
			val.validator.makeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed confirming disputable assertion")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.InitiateChallenge(
			val.validator.makeAuth(ctx),
			precondition,
			assertion,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating challenge")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.BisectAssertion(
			val.validator.makeAuth(ctx),
			precondition,
			assertions,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed initiating bisection")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.ContinueChallenge(
			val.validator.makeAuth(ctx),
			assertionToChallenge,
			preconditions,
			assertions,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed continuing challenge")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.OneStepProof(
			val.validator.makeAuth(ctx),
			precondition,
			assertion,
			proof,
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed one step proof")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.AsserterTimedOutChallenge(val.validator.makeAuth(ctx))
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.vmTracker.ChallengerTimedOutChallenge(
			val.validator.makeAuth(ctx),
		)
		if err != nil {
			errChan <- errors2.Wrap(err, "failed timing out challenge")
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
	}()
	return receiptChan, errChan
}

func (val *VMValidator) DepositFunds(
	ctx context.Context,
	amount *big.Int,
) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errChan := make(chan error, 1)
	go func() {
		defer close(receiptChan)
		defer close(errChan)
		val.mutex.Lock()
		receipt, err := val.validator.DepositFunds(val.validator.makeAuth(ctx), amount, val.Address())
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Unlock()
		receipt, err := val.validator.SendMessage(val.validator.makeAuth(ctx), protocol.NewSimpleMessage(data, tokenType, currency, val.VMID))
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.validator.ForwardMessage(val.validator.makeAuth(ctx), protocol.NewSimpleMessage(data, tokenType, currency, val.VMID), sig)
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
		val.mutex.Lock()
		receipt, err := val.validator.SendEthMessage(val.validator.makeAuth(ctx), data, val.VMID, amount)
		if err != nil {
			errChan <- err
		} else {
			receiptChan <- receipt
		}
		val.mutex.Unlock()
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
