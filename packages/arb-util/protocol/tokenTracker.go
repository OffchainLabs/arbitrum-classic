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

type BalanceTracker struct {
	TokenTypes   [][21]byte
	TokenAmounts []*big.Int
	TokenLookup  map[TokenType]int
	NFTLookup    map[nftKey]int
}

func NewBalanceTracker() *BalanceTracker {
	return &BalanceTracker{
		make([][21]byte, 0),
		make([]*big.Int, 0),
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

func (b *BalanceTracker) Equals(o *BalanceTracker) bool {
	if len(b.TokenTypes) != len(o.TokenTypes) {
		return false
	}

	for i := 0; i < len(b.TokenTypes); i++ {
		if b.TokenTypes[i] != o.TokenTypes[i] || b.TokenAmounts[i].Cmp(o.TokenAmounts[i]) != 0 {
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
	count := int64(len(b.TokenAmounts))
	err := binary.Write(wr, binary.LittleEndian, &count)
	if err != nil {
		return err
	}
	for _, tokenType := range b.TokenTypes {
		_, err := wr.Write(tokenType[:])
		if err != nil {
			return err
		}
	}
	for _, amount := range b.TokenAmounts {
		err := value.NewIntValue(amount).Marshal(wr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BalanceTracker) Clone() *BalanceTracker {
	tokenTypes := make([][21]byte, 0, len(b.TokenTypes))
	tokenAmounts := make([]*big.Int, 0, len(b.TokenTypes))
	for i := range b.TokenTypes {
		tokenTypes = append(tokenTypes, b.TokenTypes[i])
		newAmount := big.NewInt(0)
		newAmount.Set(b.TokenAmounts[i])
		tokenAmounts = append(tokenAmounts, newAmount)
	}
	return NewBalanceTrackerFromLists(tokenTypes, tokenAmounts)
}

func (b *BalanceTracker) TokenIndex(tokenType [21]byte, amount *big.Int) int {
	tokType := TokenType{}
	copy(tokType[:], tokenType[:])
	if tokType.IsToken() {
		return b.TokenLookup[tokType]
	}
	return b.NFTLookup[newNFTKey(tokType, amount)]
}

func (b *BalanceTracker) CanSpend(tokenType TokenType, amount *big.Int) bool {
	if tokenType.IsToken() {
		return amount.Cmp(b.TokenAmounts[b.TokenLookup[tokenType]]) <= 0
	}
	index, ok := b.NFTLookup[newNFTKey(tokenType, amount)]
	if !ok {
		return false
	}
	return b.TokenAmounts[index].Cmp(amount) == 0
}

func (b *BalanceTracker) Spend(tokenType TokenType, amount *big.Int) error {
	if !b.CanSpend(tokenType, amount) {
		return errors.New("not enough balance to spend")
	}

	if tokenType.IsToken() {
		b.TokenAmounts[b.TokenLookup[tokenType]].Sub(b.TokenAmounts[b.TokenLookup[tokenType]], amount)
	} else {
		b.TokenAmounts[b.NFTLookup[newNFTKey(tokenType, amount)]] = big.NewInt(0)
	}
	return nil
}

func (b *BalanceTracker) Add(tokenType TokenType, amount *big.Int) {
	if tokenType.IsToken() {
		if index, ok := b.TokenLookup[tokenType]; ok {
			b.TokenAmounts[index].Add(b.TokenAmounts[index], amount)
		} else {
			b.TokenTypes = append(b.TokenTypes, tokenType)
			b.TokenAmounts = append(b.TokenAmounts, amount)
			b.TokenLookup[tokenType] = len(b.TokenTypes) - 1
		}
	} else {
		if index, ok := b.NFTLookup[newNFTKey(tokenType, amount)]; ok {
			b.TokenAmounts[index] = amount
		} else {
			b.TokenTypes = append(b.TokenTypes, tokenType)
			b.TokenAmounts = append(b.TokenAmounts, amount)
			b.NFTLookup[newNFTKey(tokenType, amount)] = len(b.TokenTypes) - 1
		}
	}
}

func (b *BalanceTracker) SpendAll(o *BalanceTracker) error {
	for i := range o.TokenTypes {
		err := b.Spend(o.TokenTypes[i], o.TokenAmounts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BalanceTracker) AddAll(o *BalanceTracker) {
	for i := range o.TokenTypes {
		b.Add(o.TokenTypes[i], o.TokenAmounts[i])
	}
}

func (b *BalanceTracker) CanSpendAll(o *BalanceTracker) bool {
	c := b.Clone()
	return c.SpendAll(o) == nil
}

func (b *BalanceTracker) ValidAssertionStub(a *AssertionStub) bool {
	c := b.Clone()
	o := NewBalanceTrackerFromLists(b.TokenTypes, a.TotalVals)
	return c.SpendAll(o) == nil
}
