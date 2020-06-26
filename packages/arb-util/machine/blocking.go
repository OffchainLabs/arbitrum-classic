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
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type BlockReason interface {
	IsBlocked(m Machine, currentTime *common.TimeBlocks, newMessages bool) bool
	Equals(b BlockReason) bool
}

type HaltBlocked struct {
}

func (b HaltBlocked) String() string {
	return "HaltBlocked"
}

func (b HaltBlocked) IsBlocked(Machine, *common.TimeBlocks, bool) bool {
	return true
}

func (b HaltBlocked) Equals(a BlockReason) bool {
	_, ok := a.(HaltBlocked)
	return ok
}

type ErrorBlocked struct {
}

func (b ErrorBlocked) String() string {
	return "ErrorBlocked"
}

func (b ErrorBlocked) IsBlocked(Machine, *common.TimeBlocks, bool) bool {
	return true
}

func (b ErrorBlocked) Equals(a BlockReason) bool {
	_, ok := a.(ErrorBlocked)
	return ok
}

type BreakpointBlocked struct {
}

func (b BreakpointBlocked) String() string {
	return "BreakpointBlocked"
}

func (b BreakpointBlocked) IsBlocked(Machine, *common.TimeBlocks, bool) bool {
	return false
}

func (b BreakpointBlocked) Equals(a BlockReason) bool {
	_, ok := a.(BreakpointBlocked)
	return ok
}

type InboxBlocked struct {
	Timeout value.IntValue
}

func (b InboxBlocked) String() string {
	return fmt.Sprintf("InboxBlocked(%v)", b.Timeout)
}

func (b InboxBlocked) IsBlocked(_ Machine, currentTime *common.TimeBlocks, newMessages bool) bool {
	return b.Timeout.BigInt().Cmp(currentTime.AsInt()) > 0 && !newMessages
}

func (b InboxBlocked) Equals(a BlockReason) bool {
	aBlock, ok := a.(InboxBlocked)
	if !ok {
		return false
	}
	return value.Eq(aBlock.Timeout, b.Timeout)
}
