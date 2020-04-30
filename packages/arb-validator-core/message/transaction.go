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

type Transaction struct {
	Chain       common.Address
	To          common.Address
	From        common.Address
	SequenceNum *big.Int
	Value       *big.Int
	Data        []byte
}

func (m Transaction) String() string {
	return fmt.Sprintf("Transaction(chain: %v, to: %v, from: %v, seq: %v, value: %v, data: %v)",
		m.Chain,
		m.To,
		m.From,
		m.SequenceNum,
		m.Value,
		m.Data,
	)
}

func (m Transaction) GetFuncName() string {
	return hexutil.Encode(m.Data[:4])
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
		value.NewIntValue(new(big.Int).Set(m.SequenceNum)),
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

func UnmarshalTransaction(val value.Value, chain common.Address) (Transaction, error) {
	from, tup, err := unmarshalTxWrapped(val, TransactionType)
	if err != nil {
		return Transaction{}, err
	}

	if tup.Len() != 4 {
		return Transaction{}, fmt.Errorf("expected tuple of length 2, but recieved %v", tup)
	}
	destVal, _ := tup.GetByInt64(0)
	seqVal, _ := tup.GetByInt64(1)
	amountVal, _ := tup.GetByInt64(2)
	dataVal, _ := tup.GetByInt64(3)

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return Transaction{}, errors.New("dest must be an int")
	}

	seqInt, ok := seqVal.(value.IntValue)
	if !ok {
		return Transaction{}, errors.New("seq must be an int")
	}

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return Transaction{}, errors.New("amount must be an int")
	}

	data, err := ByteStackToHex(dataVal)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		Chain:       chain,
		To:          intValueToAddress(destInt),
		From:        from,
		SequenceNum: seqInt.BigInt(),
		Value:       amountInt.BigInt(),
		Data:        data,
	}, nil
}

func (m Transaction) ReceiptHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.Chain),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.SequenceNum),
		hashing.Uint256(m.Value),
		m.Data,
	)
}

type DeliveredTransaction struct {
	Transaction
	BlockNum  *common.TimeBlocks
	Timestamp *big.Int
}

func (m DeliveredTransaction) Equals(other Message) bool {
	o, ok := other.(DeliveredTransaction)
	if !ok {
		return false
	}
	return m.Transaction.Equals(o.Transaction) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0
}

func (m DeliveredTransaction) DeliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredTransaction) DeliveredTimestamp() *big.Int {
	return m.Timestamp
}

func (m DeliveredTransaction) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.Chain),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.SequenceNum),
		hashing.Uint256(m.Value),
		m.Data,
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.Timestamp),
	)
}

func (m DeliveredTransaction) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.Chain),
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(new(big.Int).Set(m.SequenceNum)),
		value.NewIntValue(new(big.Int).Set(m.Value)),
		BytesToByteStack(m.Data),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
	})
	return val
}

func UnmarshalTransactionFromCheckpoint(v value.Value) (DeliveredTransaction, error) {
	tup, ok := v.(value.TupleValue)
	failRet := DeliveredTransaction{}
	if !ok || tup.Len() != 8 {
		return failRet, errors.New("tx val must be 7-tuple")
	}
	chain, _ := tup.GetByInt64(0)
	chainInt, ok := chain.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	to, _ := tup.GetByInt64(1)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return failRet, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(2)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return failRet, errors.New("from must be int")
	}
	seq, _ := tup.GetByInt64(3)
	seqInt, ok := seq.(value.IntValue)
	if !ok {
		return failRet, errors.New("seq must be int")
	}
	val, _ := tup.GetByInt64(4)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	data, _ := tup.GetByInt64(5)
	dataBytes, err := ByteStackToHex(data)
	if err != nil {
		return failRet, err
	}
	blockNum, _ := tup.GetByInt64(6)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNum must be int")
	}
	timestamp, _ := tup.GetByInt64(7)
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be int")
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
		BlockNum:  common.NewTimeBlocks(blockNumInt.BigInt()),
		Timestamp: timestampInt.BigInt(),
	}, nil
}
