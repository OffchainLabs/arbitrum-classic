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
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type nftKey struct {
	tokenType TokenType
	intVal    [32]byte
}

func newNFTKey(tokenType TokenType, num *big.Int) nftKey {
	key := nftKey{}
	key.tokenType = tokenType
	var b bytes.Buffer
	_ = value.NewIntValue(num).Marshal(&b)
	copy(key.intVal[:], b.Bytes())
	return key
}

func (k nftKey) Value() *big.Int {
	rd := bytes.NewBuffer(k.intVal[:])
	v, _ := value.NewIntValueFromReader(rd)
	return v.BigInt()
}

type tokenEntry struct {
	tokenType TokenType
	amount    *big.Int
}

type BalanceTracker struct {
	entries     []tokenEntry
	tokenLookup map[TokenType]int
	nftLookup   map[nftKey]int
}

func tokenTypeComparator(a, b interface{}) int {
	object1 := a.(tokenEntry)
	object2 := b.(tokenEntry)
	tokDiff := object1.tokenType.ToIntValue().BigInt().Cmp(object2.tokenType.ToIntValue().BigInt())
	if tokDiff < 0 {
		return -1
	} else if tokDiff > 0 {
		return 1
	} else {
		return object1.amount.Cmp(object2.amount)
	}
}

func NewBalanceTracker() *BalanceTracker {
	return &BalanceTracker{
		make([]tokenEntry, 0),
		make(map[TokenType]int),
		make(map[nftKey]int),
	}
}

func NewBalanceTrackerFromMessages(msgs []Message) *BalanceTracker {
	tracker := NewBalanceTracker()
	for _, msg := range msgs {
		tracker.Add(msg.TokenType, msg.Currency)
	}
	return tracker
}

func NewBalanceTrackerFromLists(types [][21]byte, amounts []*big.Int) *BalanceTracker {
	tracker := NewBalanceTracker()
	for i := 0; i < len(types); i++ {
		tracker.Add(types[i], amounts[i])
	}
	return tracker
}

func (b *BalanceTracker) GetTypesAndAmounts() ([][21]byte, []*big.Int) {
	entries := b.Clone().entries
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
	tokTypes := make([][21]byte, 0, len(b.entries))
	amounts := make([]*big.Int, 0, len(b.entries))
	for _, entry := range entries {
		tokTypes = append(tokTypes, entry.tokenType)
		amounts = append(amounts, entry.amount)
	}
	return tokTypes, amounts
}

func (b *BalanceTracker) RemoveAssertionValues(totalVals []*big.Int) {
	for i, val := range totalVals {
		b.entries[i].amount.Sub(b.entries[i].amount, val)
	}
}

func (b *BalanceTracker) Equals(o *BalanceTracker) bool {
	if len(b.entries) != len(o.entries) {
		return false
	}

	for i := 0; i < len(b.entries); i++ {
		if b.entries[i].tokenType != o.entries[i].tokenType || b.entries[i].amount.Cmp(o.entries[i].amount) != 0 {
			return false
		}
	}
	return true
}

func NewBalanceTrackerFromReader(rd io.Reader) (*BalanceTracker, error) {
	types := make([][21]byte, 0)
	amounts := make([]*big.Int, 0)

	var count int64
	err := binary.Read(rd, binary.LittleEndian, &count)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < count; i++ {
		var tokenType TokenType
		_, err := io.ReadFull(rd, tokenType[:])
		if err != nil {
			return nil, err
		}
		types = append(types, tokenType)
	}

	for i := int64(0); i < count; i++ {
		val, err := value.NewIntValueFromReader(rd)
		if err != nil {
			return nil, err
		}
		amounts = append(amounts, val.BigInt())
	}
	return NewBalanceTrackerFromLists(types, amounts), nil
}

func (b *BalanceTracker) Marshal(wr io.Writer) error {
	count := int64(len(b.entries))
	err := binary.Write(wr, binary.LittleEndian, &count)
	if err != nil {
		return err
	}
	for _, entry := range b.entries {
		_, err := wr.Write(entry.tokenType[:])
		if err != nil {
			return err
		}
	}
	for _, entry := range b.entries {
		err := value.NewIntValue(entry.amount).Marshal(wr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BalanceTracker) Clone() *BalanceTracker {
	tokenTypes := make([][21]byte, 0, len(b.entries))
	tokenAmounts := make([]*big.Int, 0, len(b.entries))
	for _, entry := range b.entries {
		tokenTypes = append(tokenTypes, entry.tokenType)
		newAmount := big.NewInt(0).Set(entry.amount)
		tokenAmounts = append(tokenAmounts, newAmount)
	}
	return NewBalanceTrackerFromLists(tokenTypes, tokenAmounts)
}

func (b *BalanceTracker) TokenIndex(tokenType [21]byte, amount *big.Int) int {
	tokType := TokenType{}
	copy(tokType[:], tokenType[:])
	if tokType.IsToken() {
		return b.tokenLookup[tokType]
	}
	return b.nftLookup[newNFTKey(tokType, amount)]
}

func (b *BalanceTracker) CanSpend(tokenType TokenType, amount *big.Int) bool {
	if tokenType.IsToken() {
		return amount.Cmp(b.entries[b.tokenLookup[tokenType]].amount) <= 0
	}
	index, ok := b.nftLookup[newNFTKey(tokenType, amount)]
	if !ok {
		return false
	}
	return b.entries[index].amount.Cmp(amount) == 0
}

func (b *BalanceTracker) Spend(tokenType TokenType, amount *big.Int) error {
	if !b.CanSpend(tokenType, amount) {
		return errors.New("not enough balance to spend")
	}

	if tokenType.IsToken() {
		b.entries[b.tokenLookup[tokenType]].amount.Sub(b.entries[b.tokenLookup[tokenType]].amount, amount)
	} else {
		b.entries[b.nftLookup[newNFTKey(tokenType, amount)]].amount.SetUint64(0)
	}
	return nil
}

func (b *BalanceTracker) Add(tokenType TokenType, amount *big.Int) {
	if tokenType.IsToken() {
		if index, ok := b.tokenLookup[tokenType]; ok {
			b.entries[index].amount.Add(b.entries[index].amount, amount)
		} else {
			b.entries = append(b.entries, tokenEntry{
				tokenType: tokenType,
				amount:    amount,
			})
			b.tokenLookup[tokenType] = len(b.entries) - 1
		}
	} else {
		if index, ok := b.nftLookup[newNFTKey(tokenType, amount)]; ok {
			b.entries[index].amount.Set(amount)
		} else {
			b.entries = append(b.entries, tokenEntry{
				tokenType: tokenType,
				amount:    amount,
			})
			b.nftLookup[newNFTKey(tokenType, amount)] = len(b.entries) - 1
		}
	}
}

func (b *BalanceTracker) SpendAll(o *BalanceTracker) error {
	for _, entry := range o.entries {
		err := b.Spend(entry.tokenType, entry.amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BalanceTracker) CanSpendAll(o *BalanceTracker) bool {
	c := b.Clone()
	return c.SpendAll(o) == nil
}
