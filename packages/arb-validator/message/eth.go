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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Eth struct {
	To    common.Address
	From  common.Address
	Value *big.Int
}

func (m Eth) String() string {
	return fmt.Sprintf("Eth(to: %v, from: %v, value: %v)", m.To, m.From, m.Value)
}

func (m Eth) Equals(o Eth) bool {
	return m.To == o.To &&
		m.From == o.From &&
		m.Value.Cmp(o.Value) == 0
}

func (m Eth) Type() MessageType {
	return EthType
}

func (m Eth) GetFuncName() string {
	return "EthTransfer"
}

func (m Eth) AsValue() value.Value {
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

type DeliveredEth struct {
	Eth
	BlockNum   *common.TimeBlocks
	MessageNum *big.Int
}

func (m DeliveredEth) Equals(other DeliveredMessage) bool {
	o, ok := other.(DeliveredEth)
	if !ok {
		return false
	}
	return m.Eth.Equals(o.Eth) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.MessageNum.Cmp(o.MessageNum) == 0
}

func (m DeliveredEth) DeliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredEth) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		hashing.Address(m.To),
		hashing.Address(m.From),
		hashing.Uint256(m.Value),
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.MessageNum),
	)
}

func (m DeliveredEth) ReceiptHash() common.Hash {
	return value.NewIntValue(m.MessageNum).ToBytes()
}

func (m DeliveredEth) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.To),
		addressToIntValue(m.From),
		value.NewIntValue(m.Value),
		value.NewIntValue(m.BlockNum.AsInt()),
		value.NewIntValue(m.MessageNum),
	})
	return val
}

func UnmarshalEthFromCheckpoint(v value.Value) (DeliveredEth, error) {
	tup, ok := v.(value.TupleValue)
	if !ok || tup.Len() != 5 {
		return DeliveredEth{}, errors.New("tx val must be 5-tuple")
	}
	to, _ := tup.GetByInt64(0)
	toInt, ok := to.(value.IntValue)
	if !ok {
		return DeliveredEth{}, errors.New("to must be int")
	}
	from, _ := tup.GetByInt64(1)
	fromInt, ok := from.(value.IntValue)
	if !ok {
		return DeliveredEth{}, errors.New("from must be int")
	}
	val, _ := tup.GetByInt64(2)
	valInt, ok := val.(value.IntValue)
	if !ok {
		return DeliveredEth{}, errors.New("chain must be int")
	}
	blockNum, _ := tup.GetByInt64(3)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return DeliveredEth{}, errors.New("blockNum must be int")
	}
	messageNum, _ := tup.GetByInt64(4)
	messageNumInt, ok := messageNum.(value.IntValue)
	if !ok {
		return DeliveredEth{}, errors.New("messageNum must be int")
	}

	return DeliveredEth{
		Eth: Eth{
			To:    intValueToAddress(toInt),
			From:  intValueToAddress(fromInt),
			Value: valInt.BigInt(),
		},
		BlockNum:   common.NewTimeBlocks(blockNumInt.BigInt()),
		MessageNum: messageNumInt.BigInt(),
	}, nil
}
