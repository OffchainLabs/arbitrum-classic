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

type Eth struct {
	To    common.Address
	From  common.Address
	Value *big.Int
}

func NewRandomEth() Eth {
	return Eth{
		To:    common.RandAddress(),
		From:  common.RandAddress(),
		Value: common.RandBigInt(),
	}
}

func (m Eth) String() string {
	return fmt.Sprintf("Eth(to: %v, from: %v, value: %v)", m.To, m.From, m.Value)
}

func (m Eth) Equals(other Message) bool {
	o, ok := other.(Eth)
	if !ok {
		return false
	}
	return m.To == o.To &&
		m.From == o.From &&
		m.Value.Cmp(o.Value) == 0
}

func (m Eth) Type() Type {
	return EthType
}

func (m Eth) GetFuncName() string {
	return "EthTransfer"
}

func (m Eth) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
	)
}

func (m Eth) AsInboxValue() value.TupleValue {
	val1, _ := value.NewTupleFromSlice([]value.Value{
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

func UnmarshalEth(val value.Value) (Eth, error) {
	from, tup, err := unmarshalTxWrapped(val, EthType)
	if err != nil {
		return Eth{}, err
	}

	if tup.Len() != 2 {
		return Eth{}, fmt.Errorf("expected tuple of length 2, but recieved %v", tup)
	}
	destVal, _ := tup.GetByInt64(0)
	amountVal, _ := tup.GetByInt64(1)

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return Eth{}, errors.New("dest must be an int")
	}

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return Eth{}, errors.New("amount must be an int")
	}

	return Eth{
		To:    intValueToAddress(destInt),
		From:  from,
		Value: amountInt.BigInt(),
	}, nil
}

func (m Eth) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(new(big.Int).Set(m.Value)),
	})
	return val
}

func UnmarshalEthFromCheckpoint(v value.Value) (Eth, error) {
	tup, ok := v.(value.TupleValue)
	failRet := Eth{}
	if !ok || tup.Len() != 3 {
		return failRet, errors.New("tx val must be 3-tuple")
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
	val, _ := tup.GetByInt64(2)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}

	return Eth{
		To:    intValueToAddress(toInt),
		From:  intValueToAddress(fromInt),
		Value: valInt.BigInt(),
	}, nil
}
