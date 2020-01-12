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

	"github.com/ethereum/go-ethereum/common/hexutil"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Precondition struct {
	BeforeHash  [32]byte
	TimeBounds  *protocol.TimeBoundsBlocks
	BeforeInbox value.Value
}

func NewPrecondition(beforeHash [32]byte, timeBounds *protocol.TimeBoundsBlocks, beforeInbox value.Value) *Precondition {
	return &Precondition{BeforeHash: beforeHash, TimeBounds: timeBounds, BeforeInbox: beforeInbox}
}

func (pre *Precondition) String() string {
	inboxHash := pre.BeforeInbox.Hash()
	return fmt.Sprintf(
		"Precondition(beforeHash: %v, timebounds: [%v, %v], BeforeInbox: %v)",
		hexutil.Encode(pre.BeforeHash[:]),
		pre.TimeBounds.Start.AsInt(),
		pre.TimeBounds.End.AsInt(),
		hexutil.Encode(inboxHash[:]),
	)
}

func (pre *Precondition) Equals(b *Precondition) bool {
	return pre.BeforeHash == b.BeforeHash ||
		pre.TimeBounds.Equals(b.TimeBounds) ||
		value.Eq(pre.BeforeInbox, b.BeforeInbox)
}

func (pre *Precondition) Hash() [32]byte {
	var ret [32]byte
	bounds := pre.TimeBounds.AsIntArray()
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(pre.BeforeHash),
		solsha3.Uint128(bounds[0]),
		solsha3.Uint128(bounds[1]),
		solsha3.Bytes32(pre.BeforeInbox.Hash()),
	))
	return ret
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
