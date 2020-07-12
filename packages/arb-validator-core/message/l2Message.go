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

const AddressSize = 32

const TransactionHeaderSize = 32*4 + AddressSize
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
		if len(data) <= 32 {
			break
		}
		data = data[32:]
	}
	return ret
}

type AbstractL2Message interface {
	l2Type() L2SubType
	AsData() []byte
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
	ret = append(ret, l.Msg.AsData()...)
	return ret
}

func NewL2MessageFromData(data []byte) (AbstractL2Message, error) {
	l2Type := L2SubType(data[0])
	data = data[1:]
	switch l2Type {
	case TransactionType:
		return newTransactionFromData(data), nil
	case ContractTransactionType:
		return newContractTransactionFromData(data), nil
	case CallType:
		return NewCallFromData(data), nil
	case TransactionBatchType:
		return newTransactionBatchFromData(data), nil
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

func extractUInt256(data []byte) (*big.Int, []byte) {
	val := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return val, data
}

func extractAddress(data []byte) (common.Address, []byte) {
	data = data[12:] // Skip first 12 bytes of 32 byte address data
	var addr common.Address
	copy(addr[:], data[:])
	data = data[20:]
	return addr, data
}

func addressData(addr common.Address) []byte {
	ret := make([]byte, 0, 32)
	ret = append(ret, make([]byte, 12)...)
	ret = append(ret, addr[:]...)
	return ret
}

func newTransactionFromData(data []byte) Transaction {
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	sequenceNum, data := extractUInt256(data)
	destAddress, data := extractAddress(data)
	payment, data := extractUInt256(data)
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

func (t Transaction) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(t.SequenceNum)...)
	ret = append(ret, addressData(t.DestAddress)...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

func (t Transaction) BatchTxHash(chain common.Address) common.Hash {
	data := make([]byte, 0)
	data = append(data, addressData(chain)...)
	data = append(data, t.AsData()...)
	return marshaledBytesHash(data)
}

func (t Transaction) MessageID(sender common.Address) common.Hash {
	data := make([]byte, 0)
	data = append(data, sender[:]...)
	data = append(data, t.AsData()...)
	return marshaledBytesHash(data)
}

type ContractTransaction struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	DestAddress common.Address
	Payment     *big.Int
	Data        []byte
}

func newContractTransactionFromData(data []byte) ContractTransaction {
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	destAddress, data := extractAddress(data)
	payment, data := extractUInt256(data)
	return ContractTransaction{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		DestAddress: destAddress,
		Payment:     payment,
		Data:        data,
	}
}

func NewRandomContractTransaction() ContractTransaction {
	return ContractTransaction{
		MaxGas:      common.RandBigInt(),
		GasPriceBid: common.RandBigInt(),
		DestAddress: common.RandAddress(),
		Payment:     common.RandBigInt(),
		Data:        common.RandBytes(200),
	}
}

func (t ContractTransaction) l2Type() L2SubType {
	return ContractTransactionType
}

func (t ContractTransaction) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, addressData(t.DestAddress)...)
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

func NewSimpleCall(dest common.Address, data []byte) L2Message {
	return L2Message{Msg: Call{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		DestAddress: dest,
		Data:        data,
	}}
}

func NewCallFromData(data []byte) Call {
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	destAddress, data := extractAddress(data)
	return Call{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		DestAddress: destAddress,
		Data:        data,
	}
}

func NewRandomCall() Call {
	return Call{
		MaxGas:      common.RandBigInt(),
		GasPriceBid: common.RandBigInt(),
		DestAddress: common.RandAddress(),
		Data:        common.RandBytes(200),
	}
}

func (t Call) l2Type() L2SubType {
	return CallType
}

func (c Call) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(c.MaxGas)...)
	ret = append(ret, math.U256Bytes(c.GasPriceBid)...)
	ret = append(ret, addressData(c.DestAddress)...)
	ret = append(ret, c.Data...)
	return ret
}

type BatchTx struct {
	Transaction Transaction
	Signature   [SignatureSize]byte
}

func NewRandomBatchTx(chain common.Address, privKey *ecdsa.PrivateKey) BatchTx {
	tx := NewRandomTransaction()
	hash := tx.BatchTxHash(chain)
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
	ret = append(ret, b.Transaction.AsData()...)
	ret = append(ret, b.Signature[:]...)
	return ret
}

func (b BatchTx) Hash() common.Hash {
	data := make([]byte, 0)
	data = append(data, b.Transaction.AsData()...)
	data = append(data, b.Signature[:]...)
	return marshaledBytesHash(data)
}

type TransactionBatch struct {
	Transactions []BatchTx
}

func newTransactionBatchFromData(data []byte) TransactionBatch {
	txes := make([]BatchTx, 0)
	for len(data) >= 8 {
		calldataLength := binary.BigEndian.Uint64(data[:])
		data = data[8:]
		beginningSize := TransactionHeaderSize + calldataLength
		if uint64(len(data)) < beginningSize+SignatureSize {
			// Not enough data remaining
			break
		}
		tx := newTransactionFromData(data[:beginningSize])
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

func NewRandomTransactionBatch(txCount int, chain common.Address, privKey *ecdsa.PrivateKey) TransactionBatch {
	txes := make([]BatchTx, 0, txCount)
	for i := 0; i < txCount; i++ {
		txes = append(txes, NewRandomBatchTx(chain, privKey))
	}
	return TransactionBatch{Transactions: txes}
}

func (t TransactionBatch) l2Type() L2SubType {
	return TransactionBatchType
}

func (t TransactionBatch) AsData() []byte {
	ret := make([]byte, 0)
	for _, tx := range t.Transactions {
		ret = append(ret, tx.AsData()...)
	}
	return ret
}
