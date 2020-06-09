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
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/keycard-go/hexutils"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type TransactionBatch struct {
	Chain  common.Address
	TxData []byte
}

// BatchTxHash hashes the transaction data. This hash is signed by users
// who submit transactions as part of a batch
func BatchTxHash(
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

type BatchTx struct {
	To     common.Address
	SeqNum *big.Int
	Value  *big.Int
	Data   []byte
	Sig    [65]byte
}

func (b BatchTx) Equals(o BatchTx) bool {
	return b.To == o.To &&
		b.SeqNum.Cmp(o.SeqNum) == 0 &&
		b.Value.Cmp(o.Value) == 0 &&
		bytes.Equal(b.Data, o.Data) &&
		b.Sig == o.Sig
}

func NewBatchTxFromData(data []byte, offset int) (BatchTx, error) {
	dataLength := int(binary.BigEndian.Uint16(data[offset : offset+2]))
	if offset+DataOffset+dataLength > len(data) {
		return BatchTx{}, fmt.Errorf("not enough data remaining (offset: %v, DataOffset: %v, dataLength: %v, totalLength: %v)", offset, DataOffset, dataLength, len(data))
	}
	offset += 2
	toRaw := data[offset : offset+20]
	var to common.Address
	copy(to[:], toRaw)
	offset += 20
	seqRaw := data[offset : offset+32]
	seq := new(big.Int).SetBytes(seqRaw)
	offset += 32
	valueRaw := data[offset : offset+32]
	val := new(big.Int).SetBytes(valueRaw)
	offset += 32
	sigRaw := data[offset : offset+65]
	var sig [65]byte
	copy(sig[:], sigRaw)
	offset += 65
	txData := data[offset : offset+dataLength]

	return BatchTx{
		To:     to,
		SeqNum: seq,
		Value:  val,
		Data:   txData,
		Sig:    sig,
	}, nil
}

func (b BatchTx) String() string {
	return fmt.Sprintf("BatchTx(to: %v, seq: %v, value: %v, data: %v)",
		b.To,
		b.SeqNum,
		b.Value,
		hexutils.BytesToHex(b.Data),
	)
}

func (b BatchTx) encodedLength() int {
	return DataOffset + len(b.Data)
}

func (b BatchTx) ToBytes() []byte {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data[:], uint16(len(b.Data)))
	data = append(data, b.To[:]...)
	data = append(data, abi.U256(b.SeqNum)...)
	data = append(data, abi.U256(b.Value)...)
	data = append(data, b.Sig[:]...)
	data = append(data, b.Data...)
	return data
}

var DataOffset = 151

func (m TransactionBatch) getBatchTransactions() []BatchTx {
	txes := make([]BatchTx, 0)
	offset := 0
	data := m.TxData
	for offset+DataOffset < len(data) {
		batch, err := NewBatchTxFromData(data, offset)
		if err != nil {
			log.Println("Transaction batch", hexutil.Encode(data), "at offset", offset, "included invalid tx", err)
			break
		}
		txes = append(txes, batch)
		offset += batch.encodedLength()
	}
	return txes
}

func (m TransactionBatch) getTransactions() []Transaction {
	txes := make([]Transaction, 0)
	for _, tx := range m.getBatchTransactions() {
		batchTxHash := BatchTxHash(
			m.Chain,
			tx.To,
			tx.SeqNum,
			tx.Value,
			tx.Data,
		)
		messageHash := hashing.SoliditySHA3WithPrefix(batchTxHash[:])
		pubkey, err := crypto.SigToPub(messageHash.Bytes(), tx.Sig[:])
		if err != nil {
			// TODO: Is this possible? If so we need to handle it
			// What are the possible failure conditions and how do they relate
			// to ecrecover's behavior
			log.Fatalln("Invalid sig", err)
		}

		from := common.NewAddressFromEth(crypto.PubkeyToAddress(*pubkey))
		fullTx := Transaction{
			Chain:       m.Chain,
			To:          tx.To,
			From:        from,
			SequenceNum: tx.SeqNum,
			Value:       tx.Value,
			Data:        tx.Data,
		}
		txes = append(txes, fullTx)
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
	return m.Chain != o.Chain || bytes.Equal(m.TxData, o.TxData)
}

func (m TransactionBatch) Type() Type {
	return TransactionBatchType
}

func (m TransactionBatch) CommitmentHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint8(uint8(m.Type())),
		m.TxData,
	)
}

func (m TransactionBatch) CheckpointValue() value.Value {
	val, _ := value.NewTupleFromSlice([]value.Value{
		addressToIntValue(m.Chain),
		BytesToByteStack(m.TxData),
	})
	return val
}

func UnmarshalTransactionBatchFromCheckpoint(v value.Value) (TransactionBatch, error) {
	tup, ok := v.(value.TupleValue)
	failRet := TransactionBatch{}
	if !ok || tup.Len() != 2 {
		return failRet, errors.New("tx val must be 2-tuple")
	}
	chain, _ := tup.GetByInt64(0)
	chainInt, ok := chain.(value.IntValue)
	if !ok {
		return failRet, errors.New("chain must be int")
	}
	data, _ := tup.GetByInt64(1)
	dataBytes, err := ByteStackToHex(data)
	if err != nil {
		return failRet, err
	}

	return TransactionBatch{
		Chain:  intValueToAddress(chainInt),
		TxData: dataBytes,
	}, nil
}
