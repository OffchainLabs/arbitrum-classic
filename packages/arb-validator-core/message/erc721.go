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

type ERC721 struct {
	To           common.Address
	From         common.Address
	TokenAddress common.Address
	Id           *big.Int
}

func (m ERC721) String() string {
	return fmt.Sprintf("ERC721(to: %v, from: %v, token: %v, id: %v)",
		m.To,
		m.From,
		m.TokenAddress,
		m.Id,
	)
}

func (m ERC721) Equals(other Message) bool {
	o, ok := other.(ERC721)
	if !ok {
		return false
	}
	return m.To == o.To &&
		m.From == o.From &&
		m.TokenAddress == o.TokenAddress &&
		m.Id.Cmp(o.Id) == 0
}

func (m ERC721) Type() Type {
	return ERC721Type
}

func (m ERC721) GetFuncName() string {
	return "ERC721Transfer"
}

func (m ERC721) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Address(m.TokenAddress),
		hashing.Uint256(m.Id),
	)
}

func (m ERC721) AsInboxValue() value.TupleValue {
	val1, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.TokenAddress),
		addressToIntValue(m.To),
		value.NewIntValue(new(big.Int).Set(m.Id)),
	})
	val2, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(big.NewInt(int64(m.Type()))),
		addressToIntValue(m.From),
		val1,
	})
	return val2
}

func UnmarshalERC721(val value.Value) (ERC721, error) {
	from, token, to, amount, err := unmarshalToken(val, ERC721Type)
	if err != nil {
		return ERC721{}, err
	}

	return ERC721{
		To:           to,
		From:         from,
		TokenAddress: token,
		Id:           amount,
	}, nil
}

func (m ERC721) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		addressToIntValue(m.TokenAddress),
		value.NewIntValue(new(big.Int).Set(m.Id)),
	})
	return val
}

func UnmarshalERC721FromCheckpoint(v value.Value) (ERC721, error) {
	tup, ok := v.(value.TupleValue)
	failRet := ERC721{}
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

	return ERC721{
		To:           intValueToAddress(toInt),
		From:         intValueToAddress(fromInt),
		TokenAddress: intValueToAddress(tokenAddressInt),
		Id:           valInt.BigInt(),
	}, nil
}
