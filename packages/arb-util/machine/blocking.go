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

package machine

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type BlockReason interface {
	IsBlocked(m Machine, currentTime uint64) bool
	Equals(b BlockReason) bool
}

type HaltBlocked struct {
}

func (b HaltBlocked) IsBlocked(m Machine, currentTime uint64) bool {
	return true
}

func (b HaltBlocked) Equals(a BlockReason) bool {
	_, ok := a.(HaltBlocked)
	return ok
}

type ErrorBlocked struct {
}

func (b ErrorBlocked) IsBlocked(m Machine, currentTime uint64) bool {
	return true
}

func (b ErrorBlocked) Equals(a BlockReason) bool {
	_, ok := a.(ErrorBlocked)
	return ok
}

type BreakpointBlocked struct {
}

func (b BreakpointBlocked) IsBlocked(m Machine, currentTime uint64) bool {
	return false
}

func (b BreakpointBlocked) Equals(a BlockReason) bool {
	_, ok := a.(BreakpointBlocked)
	return ok
}

type InboxBlocked struct {
	Inbox value.HashOnlyValue
}

func (b InboxBlocked) IsBlocked(m Machine, currentTime uint64) bool {
	return value.Eq(m.InboxHash(), b.Inbox)
}

func (b InboxBlocked) Equals(a BlockReason) bool {
	aBlock, ok := a.(InboxBlocked)
	if !ok {
		return false
	}
	return value.Eq(aBlock.Inbox, b.Inbox)
}

type SendBlocked struct {
	Currency  *big.Int
	TokenType protocol.TokenType
}

func (b SendBlocked) IsBlocked(m Machine, currentTime uint64) bool {
	return m.CanSpend(b.TokenType, b.Currency)
}

func (b SendBlocked) Equals(a BlockReason) bool {
	aBlock, ok := a.(SendBlocked)
	if !ok {
		return false
	}
	return aBlock.Currency.Cmp(b.Currency) == 0 && aBlock.TokenType == b.TokenType
}
