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
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type L2SubType uint8

const (
	TransactionType         L2SubType = 0
	ContractTransactionType           = 1
	CallType                          = 2
	TransactionBatchType              = 3
)

const TransactionHeaderSize = 32*4 + 20
const SignatureSize = 65

func marshaledBytesHash(data []byte) common.Hash {
	var ret common.Hash
	copy(ret[:], math.U256Bytes(big.NewInt(int64(len(data)))))
	for len(data) > 0 {
		var nextVal common.Hash
		copy(nextVal[:], data[:])
		ret = hashing.SoliditySHA3(
			hashing.Bytes32(ret),
			hashing.Bytes32(nextVal),
		)
		data = data[32:]
	}
	return ret
}

type AbstractL2Message interface {
	l2Type() L2SubType
	asData() []byte
}

type L2Message struct {
	Msg AbstractL2Message
}

func (l L2Message) Type() Type {
	return L2Type
}

func (l L2Message) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, byte(l.Msg.l2Type()))
	ret = append(ret, l.Msg.asData()...)
	return ret
}

func NewL2MessageFromData(data []byte) (AbstractL2Message, error) {
	l2Type := L2SubType(data[0])
	data = data[1:]
	switch l2Type {
	case TransactionType:
		return NewTransactionFromData(data), nil
	case ContractTransactionType:
		return NewContractTransactionFromData(data), nil
	case CallType:
		return NewCallFromData(data), nil
	case TransactionBatchType:
		return NewTransactionBatchFromData(data), nil
	default:
		return nil, errors.New("invalid l2 message type")
	}
}

type Transaction struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	SequenceNum *big.Int
	DestAddress common.Address
	Payment     *big.Int
	Data        []byte
}

func NewTransactionFromData(data []byte) Transaction {
	maxGas := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	gasPriceBid := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	sequenceNum := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	var destAddress common.Address
	copy(destAddress[:], data[:])
	data = data[:20]
	payment := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return Transaction{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		SequenceNum: sequenceNum,
		DestAddress: destAddress,
		Payment:     payment,
		Data:        data,
	}
}

func NewRandomTransaction() Transaction {
	return Transaction{
		MaxGas:      common.RandBigInt(),
		GasPriceBid: common.RandBigInt(),
		SequenceNum: common.RandBigInt(),
		DestAddress: common.RandAddress(),
		Payment:     common.RandBigInt(),
		Data:        common.RandBytes(200),
	}
}

func (b Transaction) Equals(o Transaction) bool {
	return b.MaxGas.Cmp(o.MaxGas) == 0 &&
		b.GasPriceBid.Cmp(o.GasPriceBid) == 0 &&
		b.SequenceNum.Cmp(o.SequenceNum) == 0 &&
		b.DestAddress == o.DestAddress &&
		b.Payment.Cmp(o.Payment) == 0 &&
		bytes.Equal(b.Data, o.Data)
}

func (t Transaction) l2Type() L2SubType {
	return TransactionType
}

func (t Transaction) asData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(t.SequenceNum)...)
	ret = append(ret, t.DestAddress[:]...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

func (t Transaction) BatchTxHash(chain common.Address) common.Hash {
	data := make([]byte, 0)
	data = append(data, chain[:]...)
	data = append(data, t.asData()...)
	return marshaledBytesHash(data)
}

func (t Transaction) MessageID(sender common.Address) common.Hash {
	data := make([]byte, 0)
	data = append(data, sender[:]...)
	data = append(data, t.asData()...)
	return marshaledBytesHash(data)
}

type ContractTransaction struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	DestAddress common.Address
	Payment     *big.Int
	Data        []byte
}

func NewContractTransactionFromData(data []byte) ContractTransaction {
	maxGas := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	gasPriceBid := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	var destAddress common.Address
	copy(destAddress[:], data[:])
	data = data[:20]
	payment := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return ContractTransaction{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		DestAddress: destAddress,
		Payment:     payment,
		Data:        data,
	}
}

func (t ContractTransaction) l2Type() L2SubType {
	return ContractTransactionType
}

func (t ContractTransaction) asData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, t.DestAddress[:]...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

type Call struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	DestAddress common.Address
	Data        []byte
}

func NewSimpleCall(dest common.Address, data []byte) Call {
	return Call{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		DestAddress: dest,
		Data:        data,
	}
}

func NewCallFromData(data []byte) Call {
	maxGas := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	gasPriceBid := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	var destAddress common.Address
	copy(destAddress[:], data[:])
	data = data[:20]
	return Call{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		DestAddress: destAddress,
		Data:        data,
	}
}

func (t Call) l2Type() L2SubType {
	return CallType
}

func (c Call) asData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(c.MaxGas)...)
	ret = append(ret, math.U256Bytes(c.GasPriceBid)...)
	ret = append(ret, c.DestAddress[:]...)
	ret = append(ret, c.Data...)
	return ret
}

type BatchTx struct {
	Transaction Transaction
	Signature   [SignatureSize]byte
}

func NewRandomBatchTx(chain common.Address, privKey *ecdsa.PrivateKey) BatchTx {
	tx := NewRandomTransaction()
	data := make([]byte, 0)
	data = append(data, chain[:]...)
	data = append(data, tx.asData()...)
	hash := marshaledBytesHash(data)
	messageHash := hashing.SoliditySHA3WithPrefix(hash[:])
	sigBytes, _ := crypto.Sign(messageHash.Bytes(), privKey)

	var sig [65]byte
	copy(sig[:], sigBytes)

	return BatchTx{
		Transaction: tx,
		Signature:   sig,
	}
}

func (b BatchTx) Equals(o BatchTx) bool {
	return b.Transaction.Equals(o.Transaction) &&
		b.Signature == o.Signature
}

func (b BatchTx) AsData() []byte {
	ret := make([]byte, 0)
	encodedLength := make([]byte, 8)
	binary.BigEndian.PutUint64(encodedLength[:], uint64(len(b.Transaction.Data)))
	ret = append(ret, encodedLength[:]...)
	ret = append(ret, b.Transaction.asData()...)
	ret = append(ret, b.Signature[:]...)
	return ret
}

func (b BatchTx) Hash() common.Hash {
	data := make([]byte, 0)
	data = append(data, b.Transaction.asData()...)
	data = append(data, b.Signature[:]...)
	return marshaledBytesHash(data)
}

type TransactionBatch struct {
	Transactions []BatchTx
}

func NewTransactionBatchFromData(data []byte) TransactionBatch {
	txes := make([]BatchTx, 0)
	for len(data) >= 8 {
		calldataLength := binary.BigEndian.Uint64(data[:])
		data = data[8:]
		beginningSize := TransactionHeaderSize + calldataLength
		if uint64(len(data)) < beginningSize+SignatureSize {
			// Not enough data remaining
			break
		}
		tx := NewTransactionFromData(data[:beginningSize])
		data = data[beginningSize:]
		var sig [SignatureSize]byte
		copy(sig[:], data[:])
		data = data[SignatureSize:]
		txes = append(txes, BatchTx{
			Transaction: tx,
			Signature:   sig,
		})
	}
	return TransactionBatch{Transactions: txes}
}

func (t TransactionBatch) l2Type() L2SubType {
	return TransactionBatchType
}

func (t TransactionBatch) asData() []byte {
	ret := make([]byte, 0)
	for _, tx := range t.Transactions {
		ret = append(ret, tx.AsData()...)
	}
	return ret
}
