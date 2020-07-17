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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"log"
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
	SignedTransactionType             = 4
)

const AddressSize = 32

const TransactionHeaderSize = 32*4 + AddressSize
const SignatureSize = 65

type AbstractL2Message interface {
	L2Type() L2SubType
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
	ret = append(ret, byte(l.Msg.L2Type()))
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
	case SignedTransactionType:
		log.Println("GOT SIGNED TX")
		return newSignedTransactionFromData(data), nil
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

func (b Transaction) AsEthTx() *types.Transaction {
	return types.NewTransaction(b.SequenceNum.Uint64(), b.DestAddress.ToEthAddress(), b.GasPriceBid, b.MaxGas.Uint64(), b.Payment, b.Data)
}

func (b Transaction) String() string {
	return fmt.Sprintf(
		"Transaction(%v, %v, %v, %v, %v, %v)",
		b.MaxGas,
		b.GasPriceBid,
		b.SequenceNum,
		b.DestAddress,
		b.Payment,
		hexutil.Encode(b.Data),
	)
}

func (b Transaction) Equals(o Transaction) bool {
	return b.MaxGas.Cmp(o.MaxGas) == 0 &&
		b.GasPriceBid.Cmp(o.GasPriceBid) == 0 &&
		b.SequenceNum.Cmp(o.SequenceNum) == 0 &&
		b.DestAddress == o.DestAddress &&
		b.Payment.Cmp(o.Payment) == 0 &&
		bytes.Equal(b.Data, o.Data)
}

func (t Transaction) Type() Type {
	return L2Type
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
	return hashing.SoliditySHA3(hashing.Address(sender), hashing.Address(chain), t.AsData())
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

func (t ContractTransaction) Type() Type {
	return L2Type
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

func (t Call) L2Type() L2SubType {
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
	sig[64] = byte(v.Uint64() % 2)
	return SignedTransaction{
		Transaction: NewTransactionFromEthTx(tx),
		Signature:   sig,
	}
}

func ChainAddressToID(chain common.Address) *big.Int {
	return new(big.Int).SetBytes(chain[14:])
}

func NewRandomBatchTx(chain common.Address, privKey *ecdsa.PrivateKey) SignedTransaction {
	tx := NewRandomTransaction()

	signedTx, err := types.SignTx(tx.AsEthTx(), types.NewEIP155Signer(ChainAddressToID(chain)), privKey)
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
	return t.Transaction.AsEthTx().WithSignature(types.NewEIP155Signer(ChainAddressToID(chain)), t.Signature[:])
}

func (t SignedTransaction) L2Type() L2SubType {
	return SignedTransactionType
}

func (b SignedTransaction) Equals(o SignedTransaction) bool {
	return b.Transaction.Equals(o.Transaction) &&
		b.Signature == o.Signature
}

func (b SignedTransaction) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, b.Transaction.AsData()...)
	ret = append(ret, b.Signature[:]...)
	return ret
}

func (b SignedTransaction) Hash() common.Hash {
	data := make([]byte, 0)
	data = append(data, b.Transaction.AsData()...)
	data = append(data, b.Signature[:]...)
	return marshaledBytesHash(data)
}

type TransactionBatch struct {
	Transactions [][]byte
}

func NewTransactionBatchFromMessages(messages []L2Message) TransactionBatch {
	txes := make([][]byte, 0)
	for _, msg := range messages {
		txes = append(txes, msg.AsData())
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
