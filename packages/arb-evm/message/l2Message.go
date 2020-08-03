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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type L2Message struct {
	Data []byte
}

func (l L2Message) Type() inbox.Type {
	return L2Type
}

func (l L2Message) AsData() []byte {
	return l.Data
}

type L2SubType uint8

const (
	TransactionType         L2SubType = 0
	ContractTransactionType           = 1
	CallType                          = 2
	TransactionBatchType              = 3
	SignedTransactionType             = 4
)

const SignatureSize = 65

type AbstractL2Message interface {
	L2Type() L2SubType
	AsData() []byte
}

type AbstractTransaction interface {
	Destination() common.Address
}

type EthConvertable interface {
	AsEthTx(chain common.Address) (*types.Transaction, error)
}

func NewL2Message(msg AbstractL2Message) L2Message {
	data := make([]byte, 0)
	data = append(data, byte(msg.L2Type()))
	data = append(data, msg.AsData()...)
	return L2Message{Data: data}
}

func (l L2Message) AbstractMessage() (AbstractL2Message, error) {
	data := l.Data
	l2Type := L2SubType(data[0])
	data = data[1:]
	switch l2Type {
	case TransactionType:
		return newTransactionFromData(data), nil
	case ContractTransactionType:
		return NewContractTransactionFromData(data), nil
	case CallType:
		return NewCallFromData(data), nil
	case TransactionBatchType:
		return newTransactionBatchFromData(data), nil
	case SignedTransactionType:
		return newSignedTransactionFromData(data), nil
	default:
		return nil, errors.New("invalid l2 l2message type")
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

func NewTransactionFromEthTx(tx *types.Transaction) Transaction {
	var dest common.Address
	if tx.To() != nil {
		dest = common.NewAddressFromEth(*tx.To())
	}
	return Transaction{
		MaxGas:      new(big.Int).SetUint64(tx.Gas()),
		GasPriceBid: tx.GasPrice(),
		SequenceNum: new(big.Int).SetUint64(tx.Nonce()),
		DestAddress: dest,
		Payment:     tx.Value(),
		Data:        tx.Data(),
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

func (t Transaction) AsEthTx(_ common.Address) (*types.Transaction, error) {
	emptyAddress := common.Address{}
	if t.DestAddress == emptyAddress {
		return types.NewContractCreation(t.SequenceNum.Uint64(), t.GasPriceBid, t.MaxGas.Uint64(), t.Payment, t.Data), nil
	} else {
		return types.NewTransaction(t.SequenceNum.Uint64(), t.DestAddress.ToEthAddress(), t.GasPriceBid, t.MaxGas.Uint64(), t.Payment, t.Data), nil
	}
}

func (t Transaction) Destination() common.Address {
	return t.DestAddress
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"Transaction(%v, %v, %v, %v, %v, %v)",
		t.MaxGas,
		t.GasPriceBid,
		t.SequenceNum,
		t.DestAddress,
		t.Payment,
		hexutil.Encode(t.Data),
	)
}

func (t Transaction) Equals(o Transaction) bool {
	return t.MaxGas.Cmp(o.MaxGas) == 0 &&
		t.GasPriceBid.Cmp(o.GasPriceBid) == 0 &&
		t.SequenceNum.Cmp(o.SequenceNum) == 0 &&
		t.DestAddress == o.DestAddress &&
		t.Payment.Cmp(o.Payment) == 0 &&
		bytes.Equal(t.Data, o.Data)
}

func (t Transaction) L2Type() L2SubType {
	return TransactionType
}

func (t Transaction) AsData() []byte {
	return t.asData()
}

func (t Transaction) asData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(t.SequenceNum)...)
	ret = append(ret, addressData(t.DestAddress)...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

func (t Transaction) MessageID(sender common.Address, chain common.Address) common.Hash {
	data := NewL2Message(t).AsData()
	inner := hashing.SoliditySHA3(hashing.Uint256(ChainAddressToID(chain)), hashing.Bytes32(marshaledBytesHash(data)))
	return hashing.SoliditySHA3(addressData(sender), hashing.Bytes32(inner))
}

type ContractTransaction struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	DestAddress common.Address
	Payment     *big.Int
	Data        []byte
}

func NewContractTransactionFromData(data []byte) ContractTransaction {
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

func (t ContractTransaction) Destination() common.Address {
	return t.DestAddress
}

func (t ContractTransaction) L2Type() L2SubType {
	return ContractTransactionType
}

func (t ContractTransaction) AsData() []byte {
	return t.asData()
}

func (t ContractTransaction) asData() []byte {
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

func (c Call) Destination() common.Address {
	return c.DestAddress
}

func (c Call) L2Type() L2SubType {
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

type SignedTransaction struct {
	Transaction Transaction
	Signature   [SignatureSize]byte
}

func newSignedTransactionFromData(data []byte) SignedTransaction {
	var sig [SignatureSize]byte
	tx := newTransactionFromData(data[:len(data)-SignatureSize])
	data = data[len(data)-SignatureSize:]
	copy(sig[:], data[:])
	return SignedTransaction{
		Transaction: tx,
		Signature:   sig,
	}
}

func NewSignedTransactionFromEth(tx *types.Transaction) SignedTransaction {
	v, r, s := tx.RawSignatureValues()
	var sig [65]byte
	copy(sig[:], math.U256Bytes(r))
	copy(sig[32:], math.U256Bytes(s))
	sig[64] = byte(1 - v.Uint64())

	return SignedTransaction{
		Transaction: NewTransactionFromEthTx(tx),
		Signature:   sig,
	}
}

func (t SignedTransaction) Destination() common.Address {
	return t.Transaction.Destination()
}

func ChainAddressToID(chain common.Address) *big.Int {
	return new(big.Int).SetBytes(chain[14:])
}

func NewRandomBatchTx(chain common.Address, privKey *ecdsa.PrivateKey) SignedTransaction {
	tx := NewRandomTransaction()
	ethTx, err := tx.AsEthTx(chain)
	if err != nil {
		panic(err)
	}
	signedTx, err := types.SignTx(ethTx, types.NewEIP155Signer(ChainAddressToID(chain)), privKey)
	if err != nil {
		panic(err)
	}
	v, r, s := signedTx.RawSignatureValues()
	var sig [65]byte
	copy(sig[:], math.U256Bytes(r))
	copy(sig[32:], math.U256Bytes(s))
	sig[64] = byte(v.Uint64() % 2)

	return SignedTransaction{
		Transaction: tx,
		Signature:   sig,
	}
}

func (t SignedTransaction) AsEthTx(chain common.Address) (*types.Transaction, error) {
	tx, err := t.Transaction.AsEthTx(chain)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(types.NewEIP155Signer(ChainAddressToID(chain)), t.Signature[:])
}

func (t SignedTransaction) L2Type() L2SubType {
	return SignedTransactionType
}

func (t SignedTransaction) Equals(o SignedTransaction) bool {
	return t.Transaction.Equals(o.Transaction) &&
		t.Signature == o.Signature
}

func (t SignedTransaction) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, t.Transaction.asData()...)
	ret = append(ret, t.Signature[:]...)
	return ret
}

type TransactionBatch struct {
	Transactions [][]byte
}

func NewTransactionBatchFromMessages(messages []AbstractL2Message) TransactionBatch {
	txes := make([][]byte, 0)
	for _, msg := range messages {
		txes = append(txes, NewL2Message(msg).AsData())
	}
	return TransactionBatch{Transactions: txes}
}

func newTransactionBatchFromData(data []byte) TransactionBatch {
	txes := make([][]byte, 0)
	for len(data) >= 8 {
		msgLength := binary.BigEndian.Uint64(data[:])
		data = data[8:]
		if uint64(len(data)) < msgLength {
			// Not enough data remaining
			break
		}
		txes = append(txes, data[:msgLength])
		data = data[msgLength:]
	}
	return TransactionBatch{Transactions: txes}
}

func NewRandomTransactionBatch(txCount int, chain common.Address, privKey *ecdsa.PrivateKey) TransactionBatch {
	txes := make([][]byte, 0, txCount)
	for i := 0; i < txCount; i++ {
		txes = append(txes, NewRandomBatchTx(chain, privKey).AsData())
	}
	return TransactionBatch{Transactions: txes}
}

func (t TransactionBatch) L2Type() L2SubType {
	return TransactionBatchType
}

func (t TransactionBatch) AsData() []byte {
	ret := make([]byte, 0)
	for _, tx := range t.Transactions {
		encodedLength := make([]byte, 8)
		binary.BigEndian.PutUint64(encodedLength[:], uint64(len(tx)))
		ret = append(ret, encodedLength[:]...)
		ret = append(ret, tx...)
	}
	return ret
}
