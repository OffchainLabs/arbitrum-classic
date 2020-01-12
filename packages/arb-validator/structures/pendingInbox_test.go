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
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func getStack() *MessageStack {
	messageStack := NewMessageStack()
	messageStack.DeliverMessage(value.NewInt64Value(0))
	messageStack.DeliverMessage(value.NewInt64Value(1))
	messageStack.DeliverMessage(value.NewInt64Value(2))
	messageStack.DeliverMessage(value.NewInt64Value(3))
	return messageStack
}

func TestBisection(t *testing.T) {
	messageStack := getStack()

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	topHash, err := messageStack.GetHashAtIndex(big.NewInt(4))
	if err != nil {
		t.Fatal(err)
	}

	_, err = messageStack.GenerateBisection(bottomHash, topHash, 50)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSegmentSize(t *testing.T) {
	messageStack := getStack()

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	topHash, err := messageStack.GetHashAtIndex(big.NewInt(4))
	if err != nil {
		t.Fatal(err)
	}
	size, err := messageStack.SegmentSize(bottomHash, topHash)
	if err != nil {
		t.Fatal(err)
	}
	if size != 4 {
		t.Errorf("Segment starting a beginning has wrong size %v", size)
	}
}

func TestPendingInboxInsert(t *testing.T) {
	pi := NewPendingInbox()
	if pi.newest != nil {
		t.Error("newest of new PendingInbox should be nil")
	}
	pi2 := marshalUnmarshal(pi)
	if pi.hashOfRest != pi2.hashOfRest {
		t.Error("marshal/unmarshal changes hash of empty pending inbox")
	}

	val1 := value.NewEmptyTuple()
	val2 := value.NewTuple2(val1, value.NewTuple2(val1, val1))

	pi.DeliverMessage(val1)
	if !pi.newest.message.Equal(val1) {
		t.Error("newest of PendingInbox wrong at val1")
	}
	pi2 = marshalUnmarshal(pi)
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item pending inbox")
	}

	pi.DeliverMessage(val2)
	if !pi.newest.message.Equal(val2) {
		t.Error("newest of PendingInbox wrong at val2")
	}
	pi2 = marshalUnmarshal(pi)
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of two-item pending inbox")
	}

	val3 := pi.ValueForSubseq(pi.hashOfRest, pi.newest.hash)
	if val3.Hash() != pi.newest.hash {
		t.Error("unexpected hash for extracted inbox")
	}

	pi.DiscardUpToCount(big.NewInt(0))
	pi2 = marshalUnmarshal(pi)
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item pending inbox")
	}

	pi.DiscardUpToCount(big.NewInt(1))
	pi2 = marshalUnmarshal(pi)
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item pending inbox")
	}
}

func marshalUnmarshal(pi *PendingInbox) *MessageStack {
	ctx := NewCheckpointContextImpl()
	return pi.MarshalForCheckpoint(ctx).UnmarshalFromCheckpoint(ctx)
}
