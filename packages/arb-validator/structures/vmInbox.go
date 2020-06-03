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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type VMInbox struct {
	preImageHashes []value.HashPreImage
	value          value.TupleValue
}

func NewVMInbox() *VMInbox {
	tuple := value.NewEmptyTuple()
	hashPreImage := tuple.GetPreImage()

	preImageHashes := make([]value.HashPreImage, 0)
	preImageHashes = append(preImageHashes, hashPreImage)

	return &VMInbox{
		value:          value.NewEmptyTuple(),
		preImageHashes: preImageHashes,
	}
}

func (b *VMInbox) DeliverMessage(msg message.Message) {
	b.value = message.AddToPrev(b.value, msg)

	hashPreImage := b.value.GetPreImage()
	b.preImageHashes = append(b.preImageHashes, hashPreImage)
}

func (b *VMInbox) GenerateBisection(startIndex, segments, count uint64) ([]value.HashPreImage, error) {
	if count > uint64(len(b.preImageHashes)) {
		return nil, fmt.Errorf("can't generate bisection of %v with only %v items", count, len(b.preImageHashes))
	}
	if count < segments {
		segments = count
	}
	item := startIndex
	inboxCuts := make([]value.HashPreImage, 0, segments+1)
	inboxCuts = append(inboxCuts, b.preImageHashes[item])

	otherSegmentSize := count / segments
	item += count/segments + count%segments
	inboxCuts = append(inboxCuts, b.preImageHashes[item])

	for i := uint64(1); i < segments; i++ {
		item += otherSegmentSize
		inboxCuts = append(inboxCuts, b.preImageHashes[item])
	}
	return inboxCuts, nil
}

func (b *VMInbox) String() string {
	return fmt.Sprintf("%v", b.preImageHashes)
}

func (b *VMInbox) AsValue() value.TupleValue {
	return b.value
}

func (b *VMInbox) Hash() value.HashPreImage {
	return b.preImageHashes[len(b.preImageHashes)-1]
}
