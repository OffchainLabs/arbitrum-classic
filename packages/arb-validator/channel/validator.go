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

package channel

import (
	"context"
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Validator struct {
	*ethvalidator.VMValidator
	Validators map[common.Address]validatorInfo
	arbChannel *ethbridge.ArbChannel
}

func (val *Validator) ValidatorCount() int {
	return len(val.Validators)
}

type validatorInfo struct {
	indexNum uint16
}

func NewValidator(
	val *ethvalidator.Validator,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
) (*Validator, error) {
	con, err := ethbridge.NewArbChannel(vmID, val.Client)
	if err != nil {
		return nil, err
	}

	vmVal, err := ethvalidator.NewVMValidator(
		val,
		vmID,
		machine,
		config,
		con,
	)
	if err != nil {
		return nil, err
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

	chanVal := &Validator{
		vmVal,
		manMap,
		con,
	}
	if err := chanVal.topOffDeposit(context.Background()); err != nil {
		return nil, errors2.Wrap(err, "Validator failed to top off deposit")
	}
	return chanVal, nil
}

func (val *Validator) topOffDeposit(ctx context.Context) error {
	callOpts := &bind.CallOpts{
		Pending: true,
		From:    val.Address(),
		Context: context.Background(),
	}
	current, err := val.arbChannel.CurrentDeposit(callOpts, val.Address())
	if err != nil {
		return err
	}
	required, err := val.arbChannel.EscrowRequired(callOpts)
	if current.Cmp(required) >= 0 {
		// Validator already has escrow deposited
		return nil
	}
	depToAdd := new(big.Int).Sub(required, current)
	_, err = val.arbChannel.IncreaseDeposit(val.Validator.MakeAuth(ctx), depToAdd)
	if err != nil {
		return errors2.Wrap(err, "failed calling IncreaseDeposit")
	}
	return nil
}

func (val *Validator) FinalizedUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	signatures [][]byte,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	//log.Println(" ********** FinalizedUnanimousAssert")
	//newInboxHash[0]=5
	receipt, err := val.arbChannel.FinalizedUnanimousAssert(
		val.Validator.MakeAuth(ctx),
		newInboxHash,
		assertion,
		signatures,
	)
	val.Mutex.Unlock()
	return receipt, err
}

func (val *Validator) PendingUnanimousAssert(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures [][]byte,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	//log.Println(" ********** PendingUnanimousAssert")
	//newInboxHash[0]=5
	receipt, err := val.arbChannel.PendingUnanimousAssert(
		val.Validator.MakeAuth(ctx),
		newInboxHash,
		assertion,
		sequenceNum,
		signatures,
	)
	val.Mutex.Unlock()
	return receipt, err
}

func (val *Validator) ConfirmUnanimousAsserted(
	ctx context.Context,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	val.Mutex.Lock()
	//log.Println(" ********** ConfirmUnanimousAsserted")
	//newInboxHash[0]=5
	receipt, err := val.arbChannel.ConfirmUnanimousAsserted(
		val.Validator.MakeAuth(ctx),
		newInboxHash,
		assertion,
	)
	val.Mutex.Unlock()
	return receipt, err
}
