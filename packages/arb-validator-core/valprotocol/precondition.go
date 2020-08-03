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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Precondition struct {
	BeforeHash    common.Hash
	InboxMessages []inbox.InboxMessage
}

func NewPrecondition(beforeHash common.Hash, inboxMessages []inbox.InboxMessage) *Precondition {
	return &Precondition{BeforeHash: beforeHash, InboxMessages: inboxMessages}
}

func (pre *Precondition) String() string {
	return fmt.Sprintf(
		"Precondition(beforeHash: %v, InboxMessages: %v)",
		pre.BeforeHash,
		pre.InboxMessages,
	)
}

func (pre *Precondition) Equals(b *Precondition) bool {
	if pre.BeforeHash != b.BeforeHash ||
		len(pre.InboxMessages) != len(b.InboxMessages) {
		return false
	}
	for i, msg := range pre.InboxMessages {
		if !value.Eq(msg.AsValue(), b.InboxMessages[i].AsValue()) {
			return false
		}
	}
	return false
}

func (pre *Precondition) GeneratePostcondition(a *ExecutionAssertionStub) *Precondition {
	nextInboxMessages := pre.InboxMessages
	if a.DidInboxInsn {
		nextInboxMessages = nil
	}
	return &Precondition{
		BeforeHash:    a.AfterHash,
		InboxMessages: nextInboxMessages,
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
