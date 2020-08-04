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

package structures

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type VMInbox struct {
	inboxHashes []common.Hash
	messages    []inbox.InboxMessage
}

func NewVMInbox(messages []inbox.InboxMessage) *VMInbox {
	inboxHash := common.Hash{}
	inboxHashes := make([]common.Hash, 0)
	inboxHashes = append(inboxHashes, inboxHash)

	for i := range messages {
		inboxHash = hashing.SoliditySHA3(
			hashing.Bytes32(inboxHash),
			hashing.Bytes32(messages[len(messages)-1-i].AsValue().Hash()),
		)
		inboxHashes = append(inboxHashes, inboxHash)
	}

	return &VMInbox{
		inboxHashes: inboxHashes,
		messages:    messages,
	}
}

func (b *VMInbox) GenerateBisection(startIndex, segments, count uint64) ([]common.Hash, error) {
	if count > uint64(len(b.inboxHashes)) {
		return nil, fmt.Errorf("can't generate bisection of %v with only %v items", count, len(b.inboxHashes))
	}
	if count < segments {
		segments = count
	}
	item := startIndex
	inboxCuts := make([]common.Hash, 0, segments+1)
	inboxCuts = append(inboxCuts, b.inboxHashes[item])

	otherSegmentSize := count / segments
	item += count/segments + count%segments
	inboxCuts = append(inboxCuts, b.inboxHashes[item])

	for i := uint64(1); i < segments; i++ {
		item += otherSegmentSize
		inboxCuts = append(inboxCuts, b.inboxHashes[item])
	}
	return inboxCuts, nil
}

func (b *VMInbox) String() string {
	return fmt.Sprintf("%v", b.inboxHashes)
}

func (b *VMInbox) Messages() []inbox.InboxMessage {
	return b.messages
}

func (b *VMInbox) Hash() common.Hash {
	return b.inboxHashes[len(b.inboxHashes)-1]
}
