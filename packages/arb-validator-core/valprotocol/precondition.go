/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package valprotocol

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Precondition struct {
	BeforeHash  common.Hash
	TimeBounds  *protocol.TimeBoundsBlocks
	BeforeInbox value.TupleValue
}

func NewPrecondition(beforeHash common.Hash, timeBounds *protocol.TimeBoundsBlocks, beforeInbox value.TupleValue) *Precondition {
	return &Precondition{BeforeHash: beforeHash, TimeBounds: timeBounds, BeforeInbox: beforeInbox}
}

func (pre *Precondition) String() string {
	inboxHash := pre.BeforeInbox.Hash()
	return fmt.Sprintf(
		"Precondition(beforeHash: %v, timebounds: [%v, %v], BeforeInbox: %v)",
		pre.BeforeHash,
		pre.TimeBounds.Start.AsInt(),
		pre.TimeBounds.End.AsInt(),
		inboxHash,
	)
}

func (pre *Precondition) Equals(b *Precondition) bool {
	return pre.BeforeHash == b.BeforeHash ||
		pre.TimeBounds.Equals(b.TimeBounds) ||
		value.Eq(pre.BeforeInbox, b.BeforeInbox)
}

func (pre *Precondition) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(pre.BeforeHash),
		hashing.TimeBlocks(pre.TimeBounds.Start),
		hashing.TimeBlocks(pre.TimeBounds.End),
		hashing.Bytes32(pre.BeforeInbox.Hash()),
	)
}

func (pre *Precondition) GeneratePostcondition(a *ExecutionAssertionStub) *Precondition {
	nextBeforeInbox := pre.BeforeInbox
	if a.DidInboxInsn {
		nextBeforeInbox = value.NewEmptyTuple()
	}
	return &Precondition{
		BeforeHash:  a.AfterHash,
		TimeBounds:  pre.TimeBounds,
		BeforeInbox: nextBeforeInbox,
	}
}

func GeneratePreconditions(pre *Precondition, assertions []*ExecutionAssertionStub) []*Precondition {
	preconditions := make([]*Precondition, 0, len(assertions))
	for _, assertion := range assertions {
		preconditions = append(preconditions, pre)
		pre = pre.GeneratePostcondition(assertion)
	}
	return preconditions
}
