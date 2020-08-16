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

package batcher

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestPrepareTransactions(t *testing.T) {
	type testCase struct {
		orig   []*types.Transaction
		sorted []*types.Transaction
		label  string
	}

	chain := common.RandAddress()
	keys := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < 10; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal()
		}
		keys = append(keys, pk)
	}

	cases := make([]testCase, 0)
	cases = append(cases, func() testCase {
		origTxes := make([]*types.Transaction, 0)
		sortedTxes := make([]*types.Transaction, 0)
		for i := 0; i < 10; i++ {
			tx := message.NewRandomSignedEthTx(chain, keys[0], uint64(i))
			origTxes = append(origTxes, tx)
			sortedTxes = append(sortedTxes, tx)
		}
		return testCase{
			orig:   origTxes,
			sorted: sortedTxes,
			label:  "inorder",
		}
	}())
	cases = append(cases, func() testCase {
		decodedTxes := make([]*types.Transaction, 0)
		sortedTxes := make([]*types.Transaction, 0)
		for i := 0; i < 10; i++ {
			tx := message.NewRandomSignedEthTx(chain, keys[0], uint64(9-i))
			decodedTxes = append(decodedTxes, tx)
		}
		for i := range decodedTxes {
			sortedTxes = append(sortedTxes, decodedTxes[len(decodedTxes)-1-i])
		}
		return testCase{
			orig:   decodedTxes,
			sorted: sortedTxes,
			label:  "reverse",
		}
	}())

	cases = append(cases, func() testCase {
		origTxes := make([]*types.Transaction, 0)
		sortedTxes := make([]*types.Transaction, 0)
		for i := 0; i < 10; i++ {
			tx := message.NewRandomSignedEthTx(chain, keys[i], uint64(9-i))
			origTxes = append(origTxes, tx)
			sortedTxes = append(sortedTxes, tx)
		}
		return testCase{
			orig:   origTxes,
			sorted: sortedTxes,
			label:  "reverseDifferentKeys",
		}
	}())

	signer := types.NewEIP155Signer(message.ChainAddressToID(chain))

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			sortedTxesCal := prepareTransactions(signer, tc.orig)
			t.Log("correct:", tc.sorted)
			t.Log("calculated:", sortedTxesCal)
			if len(sortedTxesCal) != len(tc.sorted) {
				t.Fatal("sorted is wrong length")
			}
			for i, tx := range tc.sorted {
				data1, err := tx.MarshalJSON()
				if err != nil {
					t.Fatal(err)
				}
				data2, err := sortedTxesCal[i].MarshalJSON()
				if err != nil {
					t.Fatal(err)
				}
				if !bytes.Equal(data1, data2) {
					t.Error("tx in wrong order")
					break
				}
			}
		})
	}
}
