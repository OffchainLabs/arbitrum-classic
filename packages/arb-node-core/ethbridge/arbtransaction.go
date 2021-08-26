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

package ethbridge

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ArbTransaction struct {
	tx   *types.Transaction
	hash ethcommon.Hash
}

func NewArbTransaction(tx *types.Transaction) *ArbTransaction {
	v, r, s := tx.RawSignatureValues()
	if v.Cmp(big.NewInt(0)) != 0 && r.Cmp(big.NewInt(0)) == 0 && s.Cmp(big.NewInt(0)) == 0 {
		// When r and s are empty and v is not empty, v contains transaction hash from fireblocks
		return NewFireblocksArbTransaction(tx, ethcommon.BytesToHash(v.Bytes()))
	}

	return &ArbTransaction{
		tx:   tx,
		hash: tx.Hash(),
	}
}

func NewMockArbTx(hash ethcommon.Hash) *ArbTransaction {
	return &ArbTransaction{hash: hash}
}

func NewFireblocksArbTransaction(tx *types.Transaction, hash ethcommon.Hash) *ArbTransaction {
	return &ArbTransaction{
		tx:   tx,
		hash: hash,
	}
}

func (t *ArbTransaction) Hash() ethcommon.Hash {
	return t.hash
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
