/*
 * Copyright 2021, Offchain Labs, Inc.
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

package arbtransaction

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ArbTransaction struct {
	tx         *types.Transaction
	hash       ethcommon.Hash
	hashForRbf ethcommon.Hash
	id         string
}

func NewArbTransaction(tx *types.Transaction) *ArbTransaction {
	hash := tx.Hash()
	return &ArbTransaction{
		tx:         tx,
		hash:       hash,
		hashForRbf: hash,
	}
}

func NewMockArbTx(hash ethcommon.Hash) *ArbTransaction {
	return &ArbTransaction{hash: hash}
}

func NewFireblocksArbTransaction(tx *types.Transaction, details *fireblocks.TransactionDetails) (*ArbTransaction, error) {
	if len(details.TxHash) == 0 {
		return nil, errors.New("missing txHash")
	}
	hashString := details.TxHash
	if strings.HasPrefix(hashString, "0x") {
		hashString = hashString[2:]
	}
	if len(hashString) != 64 {
		return nil, errors.New("txHash wrong size")
	}
	txHashBytes, err := hex.DecodeString(hashString)
	if err != nil {
		return nil, err
	}

	txHash := ethcommon.BytesToHash(txHashBytes)
	return &ArbTransaction{
		tx:         tx,
		hash:       txHash,
		hashForRbf: txHash,
		id:         details.Id,
	}, nil
}

func (t *ArbTransaction) OverrideHash(hash ethcommon.Hash) {
	t.hash = hash
}

func (t *ArbTransaction) InheritFireblocksFieldsFrom(other *ArbTransaction) {
	t.hashForRbf = other.hash
	t.id = other.id
}

func (t *ArbTransaction) Id() string {
	return t.id
}

func (t *ArbTransaction) Hash() ethcommon.Hash {
	return t.hash
}

func (t *ArbTransaction) HashForRbf() ethcommon.Hash {
	return t.hashForRbf
}

func (t *ArbTransaction) To() *ethcommon.Address {
	return t.tx.To()
}

func (t *ArbTransaction) Value() *big.Int {
	return t.tx.Value()
}

func (t *ArbTransaction) GasPrice() *big.Int {
	return t.tx.GasPrice()
}

func (t *ArbTransaction) GasTipCap() *big.Int {
	return t.tx.GasTipCap()
}

func (t *ArbTransaction) GasFeeCap() *big.Int {
	return t.tx.GasFeeCap()
}

func (t *ArbTransaction) Gas() uint64 {
	return t.tx.Gas()
}

func (t *ArbTransaction) Nonce() uint64 {
	return t.tx.Nonce()
}

func (t *ArbTransaction) Data() []byte {
	return t.tx.Data()
}

func (t *ArbTransaction) AccessList() types.AccessList {
	return t.tx.AccessList()
}

func (t *ArbTransaction) Type() uint8 {
	return t.tx.Type()
}

func (t *ArbTransaction) ChainId() *big.Int {
	return t.tx.ChainId()
}
