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

package message

import (
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ERC20 struct {
	To           common.Address
	From         common.Address
	TokenAddress common.Address
	Value        *big.Int
}

func (m ERC20) String() string {
	return fmt.Sprintf("ERC20(to: %v, from: %v, token: %v, value: %v)",
		m.To,
		m.From,
		m.TokenAddress,
		m.Value,
	)
}

func (m ERC20) Equals(other Message) bool {
	o, ok := other.(ERC20)
	if !ok {
		return false
	}
	return m.To == o.To &&
		m.From == o.From &&
		m.TokenAddress == o.TokenAddress &&
		m.Value.Cmp(o.Value) == 0
}

func (m ERC20) Type() Type {
	return ERC20Type
}

func (m ERC20) VMInboxMessages() []SingleMessage {
	return []SingleMessage{m}
}

func (m ERC20) GetFuncName() string {
	return "ERC20Transfer"
}

func (m ERC20) DestAddress() common.Address {
	return m.To
}

func (m ERC20) SenderAddress() common.Address {
	return m.From
}

func (m ERC20) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Address(m.TokenAddress),
		hashing.Uint256(m.Value),
	)
}

func (m ERC20) AsInboxValue() value.TupleValue {
	val1, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.TokenAddress),
		addressToIntValue(m.To),
		value.NewIntValue(new(big.Int).Set(m.Value)),
	})
	val2, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(big.NewInt(int64(m.Type()))),
		addressToIntValue(m.From),
		val1,
	})
	return val2
}

func UnmarshalERC20(val value.Value) (ERC20, error) {
	from, token, to, amount, err := unmarshalToken(val, ERC20Type)
	if err != nil {
		return ERC20{}, err
	}

	return ERC20{
		To:           to,
		From:         from,
		TokenAddress: token,
		Value:        amount,
	}, nil
}

func (m ERC20) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		addressToIntValue(m.TokenAddress),
		value.NewIntValue(new(big.Int).Set(m.Value)),
	})
	return val
}

func UnmarshalERC20FromCheckpoint(v value.Value) (ERC20, error) {
	tup, ok := v.(value.TupleValue)
	failRet := ERC20{}
	if !ok || tup.Len() != 4 {
		return failRet, errors.New("tx val must be 4-tuple")
	}
	to, _ := tup.GetByInt64(0)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return failRet, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(1)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return failRet, errors.New("from must be int")
	}
	tokenAddress, _ := tup.GetByInt64(2)
	tokenAddressInt, ok := tokenAddress.(value.IntValue)
	if !ok {
		return failRet, errors.New("tokenAddress must be int")
	}
	val, _ := tup.GetByInt64(3)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}

	return ERC20{
		To:           intValueToAddress(toInt),
		From:         intValueToAddress(fromInt),
		TokenAddress: intValueToAddress(tokenAddressInt),
		Value:        valInt.BigInt(),
	}, nil
}
