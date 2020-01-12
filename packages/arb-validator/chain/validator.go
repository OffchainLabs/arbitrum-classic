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

package chain

import (
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Validator struct {
	*ethvalidator.VMValidator
	arbChain *ethbridge.ArbChain
}

func (val *Validator) Address() common.Address {
	return val.Validator.Address()
}

func NewValidator(
	val *ethvalidator.Validator,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
) (*Validator, error) {
	con, err := ethbridge.NewArbChain(vmID, val.Client)
	if err != nil {
		return nil, err
	}

	vmVal, err := ethvalidator.NewVMValidator(
		val,
		vmID,
		machine,
		config,
		challengeEverything,
		con,
	)
	if err != nil {
		return nil, err
	}

	chanVal := &Validator{
		vmVal,
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
	current, err := val.arbChain.CurrentDeposit(callOpts, val.Address())
	if err != nil {
		return err
	}
	required, err := val.arbChain.EscrowRequired(callOpts)
	if current.Cmp(required) >= 0 {
		// Validator already has escrow deposited
		return nil
	}
	depToAdd := new(big.Int).Sub(required, current)
	_, err = val.arbChain.IncreaseDeposit(val.Validator.MakeAuth(ctx), depToAdd)
	if err != nil {
		return errors2.Wrap(err, "failed calling IncreaseDeposit")
	}
	return nil
}
