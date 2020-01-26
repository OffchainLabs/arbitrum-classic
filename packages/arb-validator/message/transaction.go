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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/evm"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Transaction struct {
	Chain       common.Address
	To          common.Address
	From        common.Address
	SequenceNum *big.Int
	Value       *big.Int
	Data        []byte
}

func (m Transaction) Equals(o Transaction) bool {
	return m.Chain == o.Chain &&
		m.To == o.To &&
		m.From == o.From &&
		m.SequenceNum.Cmp(o.SequenceNum) == 0 &&
		m.Value.Cmp(o.Value) == 0 &&
		bytes.Equal(m.Data, o.Data)
}

func (m Transaction) Type() MessageType {
	return TransactionType
}

func (m Transaction) AsValue() value.Value {
	val1, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		value.NewIntValue(m.Value),
	})
	val2, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(big.NewInt(int64(m.Type()))),
		addressToIntValue(m.From),
		val1,
	})
	return val2
}

func (m Transaction) ReceiptHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
		hashing.Uint256(m.SequenceNum),
		m.Data,
	)
}

type DeliveredTransaction struct {
	Transaction
	BlockNum *common.TimeBlocks
}

func (m DeliveredTransaction) Equals(other DeliveredMessage) bool {
	o, ok := other.(DeliveredTransaction)
	if !ok {
		return false
	}
	return m.Transaction.Equals(o.Transaction) &&
		m.BlockNum.Cmp(o.BlockNum) == 0
}

func (m DeliveredTransaction) DeliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredTransaction) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
		hashing.Uint256(m.SequenceNum),
		hashing.Uint256(m.BlockNum.AsInt()),
		m.Data,
	)
}

func (m DeliveredTransaction) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.Chain),
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(m.SequenceNum),
		value.NewIntValue(m.Value),
		evm.BytesToByteStack(m.Data),
		value.NewIntValue(m.BlockNum.AsInt()),
	})
	return val
}

func UnmarshalTransaction(v value.Value) (DeliveredTransaction, error) {
	tup, ok := v.(value.TupleValue)
	if !ok || tup.Len() != 7 {
		return DeliveredTransaction{}, errors.New("tx val must be 7-tuple")
	}
	chain, _ := tup.GetByInt64(0)
	chainInt, ok := chain.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("chain must be int")
	}
	to, _ := tup.GetByInt64(1)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(2)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("from must be int")
	}
	seq, _ := tup.GetByInt64(3)
	seqInt, ok := seq.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("seq must be int")
	}
	val, _ := tup.GetByInt64(4)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("chain must be int")
	}
	data, _ := tup.GetByInt64(5)
	dataBytes, err := evm.ByteStackToHex(data)
	if err != nil {
		return DeliveredTransaction{}, err
	}
	blockNum, _ := tup.GetByInt64(6)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return DeliveredTransaction{}, errors.New("blockNum must be int")
	}

	return DeliveredTransaction{
		Transaction: Transaction{
			Chain:       intValueToAddress(chainInt),
			To:          intValueToAddress(toInt),
			From:        intValueToAddress(fromInt),
			SequenceNum: seqInt.BigInt(),
			Value:       valInt.BigInt(),
			Data:        dataBytes,
		},
		BlockNum: common.NewTimeBlocks(blockNumInt.BigInt()),
	}, nil
}
