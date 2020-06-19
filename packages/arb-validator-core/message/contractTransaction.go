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
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ContractTransaction struct {
	To    common.Address
	From  common.Address
	Value *big.Int
	Data  []byte
}

func (m ContractTransaction) String() string {
	return fmt.Sprintf("ContractTransaction(to: %v, from: %v, value: %v, data: %v)",
		m.To,
		m.From,
		m.Value,
		m.Data,
	)
}

func (m ContractTransaction) GetFuncName() string {
	return hexutil.Encode(m.Data[:4])
}

func (m ContractTransaction) DestAddress() common.Address {
	return m.To
}

func (m ContractTransaction) SenderAddress() common.Address {
	return m.From
}

func (m ContractTransaction) Equals(other Message) bool {
	o, ok := other.(ContractTransaction)
	if !ok {
		return false
	}
	return m.To == o.To &&
		m.From == o.From &&
		m.Value.Cmp(o.Value) == 0 &&
		bytes.Equal(m.Data, o.Data)
}

func (m ContractTransaction) Type() Type {
	return ContractTransactionType
}

func (m ContractTransaction) VMInboxMessages() []SingleMessage {
	return []SingleMessage{m}
}

func (m ContractTransaction) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
		m.Data,
	)
}

func (m ContractTransaction) AsInboxValue() value.TupleValue {
	val1, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		value.NewIntValue(new(big.Int).Set(m.Value)),
		BytesToByteStack(m.Data),
	})
	val2, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(big.NewInt(int64(m.Type()))),
		addressToIntValue(m.From),
		val1,
	})
	return val2
}

func UnmarshalContractTransaction(val value.Value) (ContractTransaction, error) {
	from, tup, err := unmarshalTxWrapped(val, ContractTransactionType)
	if err != nil {
		return ContractTransaction{}, err
	}

	if tup.Len() != 3 {
		return ContractTransaction{}, fmt.Errorf("expected tuple of length 2, but recieved %v", tup)
	}
	destVal, _ := tup.GetByInt64(0)
	amountVal, _ := tup.GetByInt64(1)
	dataVal, _ := tup.GetByInt64(2)

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return ContractTransaction{}, errors.New("dest must be an int")
	}

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return ContractTransaction{}, errors.New("amount must be an int")
	}

	data, err := ByteStackToHex(dataVal)
	if err != nil {
		return ContractTransaction{}, err
	}

	return ContractTransaction{
		To:    intValueToAddress(destInt),
		From:  from,
		Value: amountInt.BigInt(),
		Data:  data,
	}, nil
}
func (m ContractTransaction) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(new(big.Int).Set(m.Value)),
		BytesToByteStack(m.Data),
	})
	return val
}

func UnmarshalContractTransactionFromCheckpoint(v value.Value) (ContractTransaction, error) {
	tup, ok := v.(value.TupleValue)
	failRet := ContractTransaction{}
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
	val, _ := tup.GetByInt64(2)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	data, _ := tup.GetByInt64(3)
	dataBytes, err := ByteStackToHex(data)
	if err != nil {
		return failRet, err
	}

	return ContractTransaction{
		To:    intValueToAddress(toInt),
		From:  intValueToAddress(fromInt),
		Value: valInt.BigInt(),
		Data:  dataBytes,
	}, nil
}
