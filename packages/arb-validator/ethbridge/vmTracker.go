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
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
)

type VMTracker struct {
	contract *ethconnection.VMTracker
	address  common.Address
}

func NewVMTracker(address common.Address, client *ethclient.Client) (*VMTracker, error) {
	vm, err := ethconnection.NewVMTracker(address, client)
	return &VMTracker{vm, address}, err
}

func (vm *VMTracker) CreateListeners(ctx context.Context) (chan ethconnection.Notification, chan error, error) {
	return vm.contract.CreateListeners(ctx)
}
func (vm *VMTracker) FinalizedUnanimousAssert(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	signatures [][]byte,
) (*types.Receipt, error) {
	tx, err := vm.contract.FinalizedUnanimousAssert(
		auth,
		newInboxHash,
		assertion,
		signatures,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) PendingUnanimousAssert(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures [][]byte,
) (*types.Receipt, error) {
	tx, err := vm.contract.PendingUnanimousAssert(
		auth,
		newInboxHash,
		assertion,
		sequenceNum,
		signatures,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) ConfirmUnanimousAsserted(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	tx, err := vm.contract.ConfirmUnanimousAsserted(
		auth,
		newInboxHash,
		assertion,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) PendingDisputableAssert(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	tx, err := vm.contract.PendingDisputableAssert(
		auth,
		precondition,
		assertion,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) ConfirmDisputableAsserted(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	tx, err := vm.contract.ConfirmDisputableAsserted(
		auth,
		precondition,
		assertion,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) InitiateChallenge(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (*types.Receipt, error) {
	tx, err := vm.contract.InitiateChallenge(
		auth,
		precondition,
		assertion,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) BisectAssertion(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	tx, err := vm.contract.BisectAssertion(
		auth,
		precondition,
		assertions,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) ContinueChallenge(
	auth *bind.TransactOpts,
	assertionToChallenge uint16,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	tx, err := vm.contract.ContinueChallenge(
		auth,
		assertionToChallenge,
		preconditions,
		assertions,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) OneStepProof(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
) (*types.Receipt, error) {
	tx, err := vm.contract.OneStepProof(
		auth,
		precondition,
		assertion,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) AsserterTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := vm.contract.AsserterTimedOutChallenge(
		auth,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) ChallengerTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := vm.contract.ChallengerTimedOutChallenge(
		auth,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) CurrentDeposit(
	auth *bind.CallOpts,
	address common.Address,
) (*big.Int, error) {
	return vm.contract.Tracker.CurrentDeposit(auth, address)
}

func (vm *VMTracker) EscrowRequired(
	auth *bind.CallOpts,
) (*big.Int, error) {
	return vm.contract.Tracker.EscrowRequired(auth)
}

func (vm *VMTracker) IsEnabled(
	auth *bind.CallOpts,
) (bool, error) {
	status, err := vm.contract.Tracker.GetState(auth)
	return status != 0, err
}

func (vm *VMTracker) IncreaseDeposit(
	auth *bind.TransactOpts,
	amount *big.Int,
) (*types.Receipt, error) {
	tx, err := vm.contract.IncreaseDeposit(auth, amount)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.contract.Client, tx.Hash())
}

func (vm *VMTracker) VerifyVM(
	auth *bind.CallOpts,
	config *valmessage.VMConfiguration,
	machine [32]byte,
) error {
	//code, err := vm.contract.Client.CodeAt(auth.Context, vm.address, nil)
	// Verify that VM has correct code
	vmInfo, err := vm.contract.Tracker.Vm(auth)
	if err != nil {
		return err
	}

	if vmInfo.MachineHash != machine {
		return errors.New("VM has different machine hash")
	}

	if config.GracePeriod != uint64(vmInfo.GracePeriod) {
		return errors.New("VM has different grace period")
	}

	if value.NewBigIntFromBuf(config.EscrowRequired).Cmp(vmInfo.EscrowRequired) != 0 {
		return errors.New("VM has different escrow required")
	}

	escrowCurrency, err := vm.contract.Tracker.EscrowCurrency(auth)
	if err != nil {
		return err
	}
	if protocol.NewAddressFromBuf(config.EscrowCurrency) != escrowCurrency {
		return errors.New("VM has different escrow currency")
	}

	validators := make([]common.Address, 0, len(config.AssertKeys))
	for _, assertKey := range config.AssertKeys {
		validators = append(validators, protocol.NewAddressFromBuf(assertKey))
	}
	correctValidators, err := vm.contract.Tracker.IsValidatorList(auth, validators)
	if err != nil {
		return err
	}
	if !correctValidators {
		return errors.New("VM has different validator list")
	}

	if config.MaxExecutionStepCount != vmInfo.MaxExecutionSteps {
		return errors.New("VM has different mxa steps")
	}

	owner, err := vm.contract.Tracker.Owner(auth)
	if err != nil {
		return err
	}
	if protocol.NewAddressFromBuf(config.Owner) != owner {
		return errors.New("VM has different owner")
	}
	return nil
}
