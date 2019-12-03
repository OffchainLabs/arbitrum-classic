/*
 * Copyright 2019, Offchain Labs, Inc.
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

package protocol

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type NFTKey struct {
	tokenType TokenType
	intVal    [32]byte
}

func NewNFTKey(tokenType TokenType, num *big.Int) NFTKey {
	key := NFTKey{}
	key.tokenType = tokenType
	var b bytes.Buffer
	_ = value.NewIntValue(num).Marshal(&b)
	copy(key.intVal[:], b.Bytes())
	return key
}

func (k NFTKey) Value() *big.Int {
	rd := bytes.NewBuffer(k.intVal[:])
	v, _ := value.NewIntValueFromReader(rd)
	return v.BigInt()
}

type tokenEntry struct {
	tokenType TokenType
	amount    *big.Int
}

type TokenTracker struct {
	entries     []tokenEntry
	tokenLookup map[TokenType]int
	nftLookup   map[NFTKey]int
}

func NewTokenTrackerFromMessages(msgs []Message) *TokenTracker {
	entries := make([]tokenEntry, 0, len(msgs))
	for _, msg := range msgs {
		entries = append(entries, tokenEntry{
			tokenType: msg.TokenType,
			amount:    msg.Currency,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		tokDiff := entries[i].tokenType.ToIntValue().BigInt().Cmp(entries[j].tokenType.ToIntValue().BigInt())
		if tokDiff < 0 {
			return true
		} else if tokDiff > 0 {
			return false
		} else {
			return entries[i].amount.Cmp(entries[j].amount) < 0
		}
	})

	tokenLookup := make(map[TokenType]int)
	nftLookup := make(map[NFTKey]int)

	for i, entry := range entries {
		if entry.tokenType.IsToken() {
			tokenLookup[entry.tokenType] = i
		} else {
			nftLookup[NewNFTKey(entry.tokenType, entry.amount)] = i
		}
	}
	return &TokenTracker{
		entries:     entries,
		tokenLookup: tokenLookup,
		nftLookup:   nftLookup,
	}
}

func (b *TokenTracker) GetTypesAndAmounts() ([][21]byte, []*big.Int) {
	tokTypes := make([][21]byte, 0, len(b.entries))
	amounts := make([]*big.Int, 0, len(b.entries))
	for _, entry := range b.entries {
		tokTypes = append(tokTypes, entry.tokenType)
		amounts = append(amounts, entry.amount)
	}
	return tokTypes, amounts
}

func (b *TokenTracker) TokenIndex(tokenType [21]byte, amount *big.Int) int {
	tokType := TokenType{}
	copy(tokType[:], tokenType[:])
	if tokType.IsToken() {
		return b.tokenLookup[tokType]
	}
	return b.nftLookup[NewNFTKey(tokType, amount)]
}
