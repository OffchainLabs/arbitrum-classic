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
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"

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

func (l L2Message) String() string {
	msg, err := l.AbstractMessage()
	if err != nil {
		return "invalid message"
	} else {
		return fmt.Sprintf("%v", msg)
	}
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
	ContractTransactionType L2SubType = 1
	CallType                L2SubType = 2
	TransactionBatchType    L2SubType = 3
	SignedTransactionType   L2SubType = 4
	HeartbeatType           L2SubType = 6
	CompressedECDSA         L2SubType = 7
)

type AbstractL2Message interface {
	L2Type() L2SubType
	AsData() ([]byte, error)
}

type SafeAbstractL2Message interface {
	AbstractL2Message
	AsDataSafe() []byte
}

type AbstractTransaction interface {
	Destination() common.Address
}

type EthConvertable interface {
	AsEthTx() *types.Transaction
}

func NewL2Message(msg AbstractL2Message) (L2Message, error) {
	msgData, err := msg.AsData()
	if err != nil {
		return L2Message{}, err
	}
	data := make([]byte, 0)
	data = append(data, byte(msg.L2Type()))
	data = append(data, msgData...)
	return L2Message{Data: data}, nil
}

func NewSafeL2Message(msg SafeAbstractL2Message) L2Message {
	data := make([]byte, 0)
	data = append(data, byte(msg.L2Type()))
	data = append(data, msg.AsDataSafe()...)
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
		return newSignedTransactionFromData(data)
	case CompressedECDSA:
		return newCompressedECDSATxFromData(data)
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

func (t Transaction) AsEthTx() *types.Transaction {
	emptyAddress := common.Address{}
	if t.DestAddress == emptyAddress {
		return types.NewContractCreation(t.SequenceNum.Uint64(), t.Payment, t.MaxGas.Uint64(), t.GasPriceBid, t.Data)
	} else {
		return types.NewTransaction(t.SequenceNum.Uint64(), t.DestAddress.ToEthAddress(), t.Payment, t.MaxGas.Uint64(), t.GasPriceBid, t.Data)
	}
}

func (t Transaction) Destination() common.Address {
	return t.DestAddress
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"Transaction(gas=%v, gasprice=%v, seq=%v, dest=%v, payment=%v, data=%v)",
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

func (t Transaction) AsData() ([]byte, error) {
	return t.AsDataSafe(), nil
}

func (t Transaction) AsDataSafe() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(t.SequenceNum)...)
	ret = append(ret, AddressData(t.DestAddress)...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

func (t Transaction) MessageID(sender common.Address, chainId *big.Int) common.Hash {
	l2 := NewSafeL2Message(t)
	dataHash := hashing.SoliditySHA3(l2.AsData())
	inner := hashing.SoliditySHA3(hashing.Uint256(chainId), hashing.Bytes32(dataHash))
	return hashing.SoliditySHA3(AddressData(sender), hashing.Bytes32(inner))
}

type BasicTx struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	DestAddress common.Address
	Payment     *big.Int
	Data        []byte
}

func newBasicTxFromData(data []byte) BasicTx {
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	destAddress, data := extractAddress(data)
	payment, data := extractUInt256(data)
	return BasicTx{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		DestAddress: destAddress,
		Payment:     payment,
		Data:        data,
	}
}

func newRandomBasicTx() BasicTx {
	return BasicTx{
		MaxGas:      common.RandBigInt(),
		GasPriceBid: common.RandBigInt(),
		DestAddress: common.RandAddress(),
		Payment:     common.RandBigInt(),
		Data:        common.RandBytes(200),
	}
}

func (t BasicTx) Destination() common.Address {
	return t.DestAddress
}

func (t BasicTx) AsData() ([]byte, error) {
	return t.AsDataSafe(), nil
}

func (t BasicTx) AsDataSafe() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(new(big.Int).Set(t.MaxGas))...)
	ret = append(ret, math.U256Bytes(new(big.Int).Set(t.GasPriceBid))...)
	ret = append(ret, AddressData(t.DestAddress)...)
	ret = append(ret, math.U256Bytes(new(big.Int).Set(t.Payment))...)
	ret = append(ret, t.Data...)
	return ret
}

type ContractTransaction struct {
	BasicTx
}

func NewContractTransactionFromData(data []byte) ContractTransaction {
	return ContractTransaction{BasicTx: newBasicTxFromData(data)}
}

func NewRandomContractTransaction() ContractTransaction {
	return ContractTransaction{BasicTx: newRandomBasicTx()}
}

func (t ContractTransaction) L2Type() L2SubType {
	return ContractTransactionType
}

func (t ContractTransaction) AsEthTx() *types.Transaction {
	emptyAddress := common.Address{}
	if t.DestAddress == emptyAddress {
		return types.NewContractCreation(0, t.Payment, t.MaxGas.Uint64(), t.GasPriceBid, t.Data)
	} else {
		return types.NewTransaction(0, t.DestAddress.ToEthAddress(), t.Payment, t.MaxGas.Uint64(), t.GasPriceBid, t.Data)
	}
}

func (t ContractTransaction) AsNonConstructorTx() *types.Transaction {
	return types.NewTransaction(0, t.DestAddress.ToEthAddress(), t.Payment, t.MaxGas.Uint64(), t.GasPriceBid, t.Data)
}

type Call struct {
	BasicTx
}

func NewCallFromData(data []byte) Call {
	return Call{BasicTx: newBasicTxFromData(data)}
}

func NewRandomCall() Call {
	return Call{BasicTx: newRandomBasicTx()}
}

func (c Call) L2Type() L2SubType {
	return CallType
}

type SignedTransaction struct {
	Tx *types.Transaction
}

func newSignedTransactionFromData(data []byte) (SignedTransaction, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(data, tx); err != nil {
		return SignedTransaction{}, err
	}
	return SignedTransaction{Tx: tx}, nil
}

func (t SignedTransaction) Destination() common.Address {
	dest := t.Tx.To()
	if dest != nil {
		return common.NewAddressFromEth(*dest)
	}
	return common.Address{}
}

func NewRandomSignedEthTx(privKey *ecdsa.PrivateKey, nonce uint64, chainId *big.Int) (*types.Transaction, error) {
	tx := NewRandomTransaction()
	tx.SequenceNum = new(big.Int).SetUint64(nonce)
	ethTx := tx.AsEthTx()
	return types.SignTx(ethTx, types.NewEIP155Signer(chainId), privKey)
}

func NewRandomSignedTx(privKey *ecdsa.PrivateKey, nonce uint64, chainId *big.Int) (SignedTransaction, error) {
	signedTx, err := NewRandomSignedEthTx(privKey, nonce, chainId)
	if err != nil {
		return SignedTransaction{}, err
	}
	return SignedTransaction{Tx: signedTx}, nil
}

func (t SignedTransaction) String() string {
	j, err := t.Tx.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("SignedTransaction(%v)", err)
	} else {
		return string(j)
	}
}

func (t SignedTransaction) AsEthTx() *types.Transaction {
	return t.Tx
}

func (t SignedTransaction) L2Type() L2SubType {
	return SignedTransactionType
}

func (t SignedTransaction) Equals(o SignedTransaction) bool {
	tJson, err := t.Tx.MarshalJSON()
	if err != nil {
		return false
	}
	oJson, err := o.Tx.MarshalJSON()
	if err != nil {
		return false
	}
	return bytes.Equal(tJson, oJson)
}

func (t SignedTransaction) AsData() ([]byte, error) {
	return rlp.EncodeToBytes(t.Tx)
}

type CompressedTx struct {
	SequenceNum *big.Int
	GasPrice    *big.Int
	GasLimit    *big.Int
	To          CompressedAddress
	Payment     *big.Int
	Calldata    []byte
}

func newCompressedTxFromEth(tx *types.Transaction) CompressedTx {
	var to CompressedAddress
	if tx.To() != nil {
		to = CompressedAddressFull{common.NewAddressFromEth(*tx.To())}
	}
	return CompressedTx{
		SequenceNum: new(big.Int).SetUint64(tx.Nonce()),
		GasPrice:    tx.GasPrice(),
		GasLimit:    new(big.Int).SetUint64(tx.Gas()),
		To:          to,
		Payment:     tx.Value(),
		Calldata:    tx.Data(),
	}
}

type CompressedECDSATransaction struct {
	CompressedTx

	V byte
	R *big.Int
	S *big.Int
}

func NewCompressedECDSAFromEth(tx *types.Transaction) CompressedECDSATransaction {
	v, r, s := tx.RawSignatureValues()
	vByte := byte(0)
	if !tx.Protected() {
		// None EIP-155 tx
		vByte = byte(v.Uint64())
	} else {
		vByte = byte(v.Uint64() % 2)
	}
	return CompressedECDSATransaction{
		CompressedTx: newCompressedTxFromEth(tx),
		V:            vByte,
		R:            r,
		S:            s,
	}
}

func newCompressedECDSATxFromData(data []byte) (CompressedECDSATransaction, error) {
	if len(data) < 66 {
		return CompressedECDSATransaction{}, errors.New("data is too short")
	}

	if data[0] != 0xff {
		return CompressedECDSATransaction{}, errors.New("parsing compressed tx using function table not supported")
	}

	compressedTx, err := decodeCompressedTx(bytes.NewReader(data[1 : len(data)-65]))
	if err != nil {
		return CompressedECDSATransaction{}, err
	}
	v, r, s, err := decodeECDSASig(bytes.NewReader(data[len(data)-65:]))
	if err != nil {
		return CompressedECDSATransaction{}, err
	}
	return CompressedECDSATransaction{
		CompressedTx: compressedTx,
		V:            v,
		R:            r,
		S:            s,
	}, nil
}

func (t CompressedECDSATransaction) String() string {
	dest := "ContractCreation"
	if t.To != nil {
		dest = t.To.String()
	}
	return fmt.Sprintf("CompressedECDSATransaction(%v, %v, %v, %v, %v, 0x%X)", t.SequenceNum, t.GasPrice, t.GasLimit, dest, t.Payment, t.Calldata)
}

func (t CompressedECDSATransaction) L2Type() L2SubType {
	return CompressedECDSA
}

func (t CompressedECDSATransaction) AsData() ([]byte, error) {
	data := []byte{0xff}
	txData, err := encodeUnsignedTx(t.CompressedTx)
	if err != nil {
		return nil, err
	}
	data = append(data, txData...)
	data = append(data, encodeECDSASig(t.V, t.R, t.S)...)
	return data, nil
}

func (t CompressedECDSATransaction) IsEIP155() bool {
	// If transaction is an EIP-155 transaction, v will be 0 or 1
	// If transaction is a pre-EIP-155 transaction, it will be 27 or 28
	return t.V == 0 || t.V == 1
}

func (t CompressedECDSATransaction) AsEthTx(chainId *big.Int) (*types.Transaction, error) {
	to, ok := t.To.(CompressedAddressFull)
	if !ok {
		return nil, errors.New("can only convert to tx if address is full")
	}
	var dest []byte
	emptyAddress := common.Address{}
	if to.Address != emptyAddress {
		dest = to.Address[:]
	}
	var v *big.Int
	if !t.IsEIP155() {
		v = big.NewInt(int64(t.V))
	} else {
		v = new(big.Int).Mul(chainId, big.NewInt(2))
		v = v.Add(v, big.NewInt(35+int64(1-t.V)))
	}
	txData := []interface{}{
		t.SequenceNum,
		t.GasPrice,
		t.GasLimit,
		dest,
		t.Payment,
		t.Calldata,
		v,
		t.R,
		t.S,
	}
	rlpTxData, err := rlp.EncodeToBytes(txData)
	if err != nil {
		return nil, errors.Wrap(err, "error encoding transaction")
	}
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(rlpTxData, tx); err != nil {
		return nil, errors.Wrap(err, "error decoding transaction")
	}
	return tx, nil
}

type TransactionBatch struct {
	Transactions [][]byte
}

func NewTransactionBatchFromMessages(messages []AbstractL2Message) (TransactionBatch, error) {
	txes := make([][]byte, 0)
	for _, msg := range messages {
		msg, err := NewL2Message(msg)
		if err != nil {
			return TransactionBatch{}, err
		}
		txes = append(txes, msg.AsData())
	}
	return TransactionBatch{Transactions: txes}, nil
}

func newTransactionBatchFromData(data []byte) TransactionBatch {
	txes := make([][]byte, 0)

	r := bytes.NewReader(data)
	for {
		msgLength := new(big.Int)
		if err := rlp.Decode(r, msgLength); err != nil {
			break
		}
		if big.NewInt(int64(r.Len())).Cmp(msgLength) < 0 {
			// Not enough data remaining
			logger.Warn().Msg("Received batch containing invalid data at end")
			break
		}
		txData := make([]byte, msgLength.Uint64())
		// Read wont error since we've already checked for remaining length
		_, _ = r.Read(txData)
		txes = append(txes, txData)
	}
	return TransactionBatch{Transactions: txes}
}

func NewRandomTransactionBatch(txCount int, privKey *ecdsa.PrivateKey, initialNonce uint64, chainId *big.Int) (TransactionBatch, error) {
	messages := make([]AbstractL2Message, 0, txCount)
	for i := 0; i < txCount; i++ {
		tx, err := NewRandomSignedTx(privKey, initialNonce, chainId)
		if err != nil {
			return TransactionBatch{}, err
		}
		messages = append(messages, tx)
		initialNonce++
	}
	return NewTransactionBatchFromMessages(messages)
}

func (t TransactionBatch) String() string {
	var sb strings.Builder
	sb.WriteString("TransactionBatch(")
	for i, txData := range t.Transactions {
		msg, err := L2Message{Data: txData}.AbstractMessage()
		if err != nil {
			sb.WriteString("invalid tx")
		} else {
			sb.WriteString(fmt.Sprintf("%v", msg))
		}
		if i < len(t.Transactions)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

func (t TransactionBatch) L2Type() L2SubType {
	return TransactionBatchType
}

func (t TransactionBatch) AsData() ([]byte, error) {
	return t.AsDataSafe(), nil
}

func (t TransactionBatch) AsDataSafe() []byte {
	ret := make([]byte, 0)
	for _, tx := range t.Transactions {
		encodedLength, err := rlp.EncodeToBytes(big.NewInt(int64(len(tx))))
		if err != nil {
			// This should never occur
			panic(err)
		}
		ret = append(ret, encodedLength[:]...)
		ret = append(ret, tx...)
	}
	return ret
}

type HeartbeatMessage struct {
}

func (t HeartbeatMessage) L2Type() L2SubType {
	return HeartbeatType
}

func (t HeartbeatMessage) AsData() ([]byte, error) {
	return t.AsDataSafe(), nil
}

func (t HeartbeatMessage) AsDataSafe() []byte {
	return nil
}
