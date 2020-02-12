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

func (m ContractTransaction) Equals(o ContractTransaction) bool {
	return m.To == o.To &&
		m.From == o.From &&
		m.Value.Cmp(o.Value) == 0 &&
		bytes.Equal(m.Data, o.Data)
}

func (m ContractTransaction) Type() MessageType {
	return ContractTransactionType
}

func (m ContractTransaction) AsValue() value.Value {
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

type DeliveredContractTransaction struct {
	ContractTransaction
	BlockNum   *common.TimeBlocks
	MessageNum *big.Int
}

func (m DeliveredContractTransaction) Equals(other Message) bool {
	o, ok := other.(DeliveredContractTransaction)
	if !ok {
		return false
	}
	return m.ContractTransaction.Equals(o.ContractTransaction) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.MessageNum.Cmp(o.MessageNum) == 0
}

func (m DeliveredContractTransaction) DeliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredContractTransaction) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
		m.Data,
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.MessageNum),
	)
}

func (m DeliveredContractTransaction) ReceiptHash() common.Hash {
	return value.NewIntValue(m.MessageNum).ToBytes()
}

func (m DeliveredContractTransaction) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(new(big.Int).Set(m.Value)),
		BytesToByteStack(m.Data),
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(m.MessageNum),
	})
	return val
}

func UnmarshalContractTransactionFromCheckpoint(v value.Value) (DeliveredContractTransaction, error) {
	tup, ok := v.(value.TupleValue)
	if !ok || tup.Len() != 6 {
		return DeliveredContractTransaction{}, errors.New("tx val must be 7-tuple")
	}
	to, _ := tup.GetByInt64(0)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return DeliveredContractTransaction{}, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(1)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return DeliveredContractTransaction{}, errors.New("from must be int")
	}
	val, _ := tup.GetByInt64(2)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return DeliveredContractTransaction{}, errors.New("chain must be int")
	}
	data, _ := tup.GetByInt64(3)
	dataBytes, err := ByteStackToHex(data)
	if err != nil {
		return DeliveredContractTransaction{}, err
	}
	blockNum, _ := tup.GetByInt64(4)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return DeliveredContractTransaction{}, errors.New("blockNum must be int")
	}
	msgNum, _ := tup.GetByInt64(5)
	msgNumInt, ok := msgNum.(value.IntValue)
	if !ok {
		return DeliveredContractTransaction{}, errors.New("msgNum must be int")
	}

	return DeliveredContractTransaction{
		ContractTransaction: ContractTransaction{
			To:    intValueToAddress(toInt),
			From:  intValueToAddress(fromInt),
			Value: valInt.BigInt(),
			Data:  dataBytes,
		},
		BlockNum:   common.NewTimeBlocks(blockNumInt.BigInt()),
		MessageNum: msgNumInt.BigInt(),
	}, nil
}
