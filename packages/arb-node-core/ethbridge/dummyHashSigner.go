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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type dummyHashSigner struct {
	chainId *big.Int
}

func NewDummyHashSigner(chainId *big.Int) *dummyHashSigner {
	return &dummyHashSigner{
		chainId: chainId,
	}
}

func (s *dummyHashSigner) Sender(tx *types.Transaction) (ethcommon.Address, error) {
	panic("dummyHashSigner.Sender not implemented")
}

func (s *dummyHashSigner) SignatureValues(tx *types.Transaction, sig []byte) (*big.Int, *big.Int, *big.Int, error) {
	R := big.NewInt(0)
	R.SetBytes(sig[0:32])
	return R, big.NewInt(0), big.NewInt(0), nil
}

func (s *dummyHashSigner) ChainID() *big.Int {
	return s.chainId
}

func (s *dummyHashSigner) Hash(tx *types.Transaction) ethcommon.Hash {
	panic("dummyHashSigner.Hash not implemented")
}

func (s *dummyHashSigner) Equal(types.Signer) bool {
	panic("dummyHashSigner.Equal not implemented")
}
