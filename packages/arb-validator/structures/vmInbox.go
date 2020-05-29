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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type VMInbox struct {
	hashes         []value.HashOnlyValue
	preImageHashes []value.HashPreImage
	value          value.TupleValue
}

func NewVMInbox() *VMInbox {
	tuple := value.NewEmptyTuple()
	hash := tuple.Hash()
	hashPreImage := tuple.GetPreImage()

	hashes := make([]value.HashOnlyValue, 0)
	hashes = append(hashes, value.NewHashOnlyValue(hash, hashPreImage.Size))

	preImageHashes := make([]value.HashPreImage, 0)
	preImageHashes = append(preImageHashes, hashPreImage)

	return &VMInbox{
		hashes:         hashes,
		value:          value.NewEmptyTuple(),
		preImageHashes: preImageHashes,
	}
}

func (b *VMInbox) DeliverMessage(msg message.Message) {
	tuple := value.NewTuple2(b.value, message.DeliveredValue(msg))
	b.value = tuple

	hash := b.value.Hash()
	hashPreImage := b.value.GetPreImage()

	b.hashes = append(b.hashes, value.NewHashOnlyValue(hash, hashPreImage.Size))
	b.preImageHashes = append(b.preImageHashes, hashPreImage)
}

func (b *VMInbox) GenerateBisection(startIndex, segments, count uint64) ([]value.HashOnlyValue, []common.Hash, error) {
	if count > uint64(len(b.hashes)) {
		return nil, nil, fmt.Errorf("can't generate bisection of %v with only %v items", count, len(b.hashes))
	}
	if count < segments {
		segments = count
	}
	item := startIndex
	inboxCuts := make([]value.HashOnlyValue, 0, segments+1)
	inboxCuts = append(inboxCuts, b.hashes[item])

	imageCuts := make([]common.Hash, 0, segments+1)
	imageCuts = append(imageCuts, b.preImageHashes[item].HashImage)

	otherSegmentSize := count / segments
	item += count/segments + count%segments
	inboxCuts = append(inboxCuts, b.hashes[item])
	imageCuts = append(imageCuts, b.preImageHashes[item].HashImage)
	for i := uint64(1); i < segments; i++ {
		item += otherSegmentSize
		inboxCuts = append(inboxCuts, b.hashes[item])
		imageCuts = append(imageCuts, b.preImageHashes[item].HashImage)
	}
	return inboxCuts, imageCuts, nil
}

func (b *VMInbox) String() string {
	return fmt.Sprintf("%v", b.hashes)
}

func (b *VMInbox) AsValue() value.TupleValue {
	return b.value
}

func (b *VMInbox) Hash() value.HashOnlyValue {
	return b.hashes[len(b.hashes)-1]
}
