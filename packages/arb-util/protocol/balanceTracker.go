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
	"encoding/binary"
	"errors"
	"io"
	"math/big"
)

type BalanceTracker struct {
	tokenLookup map[TokenType]*big.Int
	nftLookup   map[NFTKey]bool
}

func NewBalanceTracker() *BalanceTracker {
	return &BalanceTracker{
		make(map[TokenType]*big.Int),
		make(map[NFTKey]bool),
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

func (b *BalanceTracker) Clone() *BalanceTracker {
	tokenLookup := make(map[TokenType]*big.Int, len(b.tokenLookup))
	for key, val := range b.tokenLookup {
		tokenLookup[key] = val
	}

	nftLookup := make(map[NFTKey]bool, len(b.nftLookup))
	for key, val := range b.nftLookup {
		nftLookup[key] = val
	}
	return &BalanceTracker{
		tokenLookup: tokenLookup,
		nftLookup:   nftLookup,
	}
}

func NewBalanceTrackerFromReader(rd io.Reader) (*BalanceTracker, error) {
	var tokenLookupCount uint64
	if err := binary.Read(rd, binary.LittleEndian, &tokenLookupCount); err != nil {
		return nil, err
	}
	var nftLookupCount uint64
	if err := binary.Read(rd, binary.LittleEndian, &nftLookupCount); err != nil {
		return nil, err
	}
	tokenLookup := make(map[TokenType]*big.Int, tokenLookupCount)
	nftLookup := make(map[NFTKey]bool, nftLookupCount)

	for i := uint64(0); i < tokenLookupCount; i++ {
		var tokenType TokenType
		_, err := io.ReadFull(rd, tokenType[:])
		if err != nil {
			return nil, err
		}
		var data [32]byte
		if _, err = rd.Read(data[:]); err != nil {
			return nil, err
		}
		tokenLookup[tokenType] = new(big.Int).SetBytes(data[:])
	}

	for i := uint64(0); i < nftLookupCount; i++ {
		var tokenType TokenType
		_, err := io.ReadFull(rd, tokenType[:])
		if err != nil {
			return nil, err
		}
		var data [32]byte
		if _, err = rd.Read(data[:]); err != nil {
			return nil, err
		}
		nftLookup[NewNFTKey(tokenType, new(big.Int).SetBytes(data[:]))] = true
	}
	return &BalanceTracker{
		tokenLookup: tokenLookup,
		nftLookup:   nftLookup,
	}, nil
}

func (b *BalanceTracker) Marshal(wr io.Writer) error {
	tokenLookupCount := uint64(len(b.tokenLookup))
	if err := binary.Write(wr, binary.LittleEndian, &tokenLookupCount); err != nil {
		return err
	}
	nftLookupCount := uint64(len(b.nftLookup))
	if err := binary.Write(wr, binary.LittleEndian, &nftLookupCount); err != nil {
		return err
	}

	for tokenType, amount := range b.tokenLookup {
		_, err := wr.Write(tokenType[:])
		if err != nil {
			return err
		}
		_, err = wr.Write(amount.Bytes())
		if err != nil {
			return err
		}
	}
	for nftkey, _ := range b.nftLookup {
		_, err := wr.Write(nftkey.tokenType[:])
		if err != nil {
			return err
		}
		_, err = wr.Write(nftkey.intVal[:])
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BalanceTracker) CanSpend(tokenType TokenType, amount *big.Int) bool {
	if tokenType.IsToken() {
		if tokenVal, ok := b.tokenLookup[tokenType]; ok {
			return amount.Cmp(tokenVal) <= 0
		} else if amount.Cmp(big.NewInt(0)) == 0 {
			return true
		} else {
			return false
		}
	}
	_, ok := b.nftLookup[NewNFTKey(tokenType, amount)]
	return ok
}

func (b *BalanceTracker) Spend(tokenType TokenType, amount *big.Int) error {
	if !b.CanSpend(tokenType, amount) {
		return errors.New("not enough balance to spend")
	}

	if tokenType.IsToken() {
		b.tokenLookup[tokenType].Sub(b.tokenLookup[tokenType], amount)
	} else {
		delete(b.nftLookup, NewNFTKey(tokenType, amount))
	}
	return nil
}

func (b *BalanceTracker) Add(tokenType TokenType, amount *big.Int) {
	if tokenType.IsToken() {
		if tokenVal, ok := b.tokenLookup[tokenType]; ok {
			tokenVal.Add(tokenVal, amount)
		} else {
			b.tokenLookup[tokenType] = amount
		}
	} else {
		b.nftLookup[NewNFTKey(tokenType, amount)] = true
	}
}

func (b *BalanceTracker) SpendAll(o *BalanceTracker) error {
	for key, val := range o.tokenLookup {
		if err := b.Spend(key, val); err != nil {
			return err
		}
	}
	for key, _ := range o.nftLookup {
		_, ok := b.nftLookup[key]
		if !ok {
			return errors.New("Balance tracker tried to spend unowned NFT")
		}
		delete(b.nftLookup, key)
	}
	return nil
}

func (b *BalanceTracker) SpendAllTokens(o *TokenTracker) error {
	for key, index := range o.tokenLookup {
		entry := o.entries[index]
		if err := b.Spend(key, entry.amount); err != nil {
			return err
		}
	}
	for key, _ := range o.nftLookup {
		_, ok := b.nftLookup[key]
		if !ok {
			return errors.New("Balance tracker tried to spend unowned NFT")
		}
		delete(b.nftLookup, key)
	}
	return nil
}

func (b *BalanceTracker) CanSpendAll(o *BalanceTracker) bool {
	c := b.Clone()
	return c.SpendAll(o) == nil
}

func (b *BalanceTracker) CanSpendAllTokens(o *TokenTracker) bool {
	c := b.Clone()
	return c.SpendAllTokens(o) == nil
}
