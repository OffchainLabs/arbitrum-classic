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
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TransactionBatch struct {
	Chain        common.Address
	Tos          []common.Address
	SequenceNums []*big.Int
	Values       []*big.Int
	DataLengths  []uint32
	Data         []byte
	Signatures   []byte
}

func OffchainTxHash(
	rollupAddress common.Address,
	to common.Address,
	sequenceNum *big.Int,
	val *big.Int,
	data []byte,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Address(rollupAddress),
		hashing.Address(to),
		hashing.Uint256(sequenceNum),
		hashing.Uint256(val),
		hashing.Bytes32(hashing.SoliditySHA3(data)),
	)
}

func (m DeliveredTransactionBatch) getTransactions() []DeliveredTransaction {
	txes := make([]DeliveredTransaction, 0, len(m.Tos))
	dataOffset := uint32(0)
	for i := range m.Tos {
		data := m.Data[dataOffset : dataOffset+m.DataLengths[i]]
		offchainHash := OffchainTxHash(
			m.Chain,
			m.Tos[i],
			m.SequenceNums[i],
			m.Values[i],
			data,
		)
		messageHash := hashing.SoliditySHA3WithPrefix(offchainHash[:])
		sig := m.Signatures[i*65 : (i+1)*65]
		pubkey, err := crypto.SigToPub(messageHash.Bytes(), sig)
		if err != nil {
			log.Fatalln("Invalid sig", err)
		}

		fromAddress := common.NewAddressFromEth(crypto.PubkeyToAddress(*pubkey))

		tx := Transaction{
			Chain:       m.Chain,
			To:          m.Tos[i],
			From:        fromAddress,
			SequenceNum: m.SequenceNums[i],
			Value:       m.Values[i],
			Data:        data,
		}

		txes = append(txes, DeliveredTransaction{
			Transaction: tx,
			BlockNum:    m.BlockNum,
			Timestamp:   m.Timestamp,
		})
		dataOffset += m.DataLengths[i]
	}
	return txes
}

func (m TransactionBatch) String() string {
	return fmt.Sprintf("TransactionBatch()")
}

// Equals check for equality between this object and any other message by
// checking for full equality of all members
func (m TransactionBatch) Equals(other Message) bool {
	o, ok := other.(TransactionBatch)
	if !ok {
		return false
	}
	if m.Chain != o.Chain ||
		len(m.Tos) != len(o.Tos) ||
		len(m.SequenceNums) != len(o.SequenceNums) ||
		len(m.Values) != len(o.Values) ||
		len(m.Data) != len(o.Data) ||
		len(m.Signatures) != len(o.Signatures) {
		return false
	}

	for i := range m.Tos {
		if m.Tos[i] != o.Tos[i] {
			return false
		}
	}
	for i := range m.Tos {
		if m.SequenceNums[i].Cmp(o.SequenceNums[i]) != 0 {
			return false
		}
	}
	for i := range m.Tos {
		if m.Values[i].Cmp(o.Values[i]) != 0 {
			return false
		}
	}
	if !bytes.Equal(m.Data, o.Data) {
		return false
	}
	if !bytes.Equal(m.Signatures, o.Signatures) {
		return false
	}

	return true
}

func (m TransactionBatch) Type() MessageType {
	return TransactionBatchType
}

type DeliveredTransactionBatch struct {
	TransactionBatch
	BlockNum  *common.TimeBlocks
	Timestamp *big.Int
}

func (m DeliveredTransactionBatch) Equals(other Message) bool {
	o, ok := other.(DeliveredTransactionBatch)
	if !ok {
		return false
	}
	return m.TransactionBatch.Equals(o.TransactionBatch) &&
		m.BlockNum.Cmp(o.BlockNum) == 0 &&
		m.Timestamp.Cmp(o.Timestamp) == 0
}

func (m DeliveredTransactionBatch) deliveredHeight() *common.TimeBlocks {
	return m.BlockNum
}

func (m DeliveredTransactionBatch) deliveredTimestamp() *big.Int {
	return m.Timestamp
}

func (m DeliveredTransactionBatch) CommitmentHash() common.Hash {
	addressTy, _ := abi.NewType("address", "", nil)
	addressArrayTy, _ := abi.NewType("address[]", "", nil)
	uint256ArrayTy, _ := abi.NewType("uint256[]", "", nil)
	uint32ArrayTy, _ := abi.NewType("uint32[]", "", nil)
	bytesTy, _ := abi.NewType("bytes", "", nil)
	args := abi.Arguments{
		abi.Argument{
			Type: addressTy,
		},
		abi.Argument{
			Type: addressArrayTy,
		},
		abi.Argument{
			Type: uint256ArrayTy,
		},
		abi.Argument{
			Type: uint256ArrayTy,
		},
		abi.Argument{
			Type: uint32ArrayTy,
		},
		abi.Argument{
			Type: bytesTy,
		},
		abi.Argument{
			Type: bytesTy,
		},
	}
	packedData, err := args.Pack(
		m.Chain,
		common.AddressArrayToEth(m.Tos),
		m.SequenceNums,
		m.Values,
		m.DataLengths,
		m.Data,
		m.Signatures,
	)
	if err != nil {
		log.Fatal(err)
	}

	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		packedData,
		hashing.Uint256(m.BlockNum.AsInt()),
		hashing.Uint256(m.Timestamp),
	)
}

func (m DeliveredTransactionBatch) CheckpointValue() value.Value {
	innerTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(new(big.Int).Set(m.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).Set(m.Timestamp)),
	})
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.Chain),
		addressesToValue(m.Tos),
		intsToValue(m.SequenceNums),
		intsToValue(m.Values),
		uint32sToValue(m.DataLengths),
		BytesToByteStack(m.Data),
		BytesToByteStack(m.Signatures),
		innerTup,
	})
	return val
}

func UnmarshalTransactionBatchFromCheckpoint(v value.Value) (DeliveredTransactionBatch, error) {
	tup, ok := v.(value.TupleValue)
	failRet := DeliveredTransactionBatch{}
	if !ok || tup.Len() != 8 {
		return failRet, errors.New("tx val must be 7-tuple")
	}
	chain, _ := tup.GetByInt64(0)
	chainInt, ok := chain.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	tos, _ := tup.GetByInt64(1)
	tosAddresses, err := valueToAddresses(tos)
	if err != nil {
		return failRet, err
	}

	sequenceNums, _ := tup.GetByInt64(2)
	sequenceNumsInts, err := valueToInts(sequenceNums)
	if err != nil {
		return failRet, err
	}
	values, _ := tup.GetByInt64(3)
	valuesInts, err := valueToInts(values)
	if err != nil {
		return failRet, err
	}
	dataLengths, _ := tup.GetByInt64(4)
	dataLengthsInts, err := valueToUInt32s(dataLengths)
	if err != nil {
		return failRet, err
	}
	data, _ := tup.GetByInt64(5)
	dataBytes, err := ByteStackToHex(data)
	if err != nil {
		return failRet, err
	}
	signatures, _ := tup.GetByInt64(6)
	signaturesBytes, err := ByteStackToHex(signatures)
	if err != nil {
		return failRet, err
	}

	innerVal, _ := tup.GetByInt64(7)
	innerValTup, ok := innerVal.(value.TupleValue)
	if !ok {
		return failRet, errors.New("innerVal must be a tup")
	}

	blockNum, _ := innerValTup.GetByInt64(0)
	blockNumInt, ok := blockNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNum must be int")
	}
	timestamp, _ := innerValTup.GetByInt64(1)
	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be int")
	}

	return DeliveredTransactionBatch{
		TransactionBatch: TransactionBatch{
			Chain:        intValueToAddress(chainInt),
			Tos:          tosAddresses,
			SequenceNums: sequenceNumsInts,
			Values:       valuesInts,
			DataLengths:  dataLengthsInts,
			Data:         dataBytes,
			Signatures:   signaturesBytes,
		},
		BlockNum:  common.NewTimeBlocks(blockNumInt.BigInt()),
		Timestamp: timestampInt.BigInt(),
	}, nil
}
