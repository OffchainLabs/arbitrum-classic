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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
)

func getStack() *MessageStack {
	msg1 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(2868),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(64521)),
			Timestamp: big.NewInt(5435254),
		},
	}

	msg2 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(2868),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(64521)),
			Timestamp: big.NewInt(5435254),
		},
	}

	msg3 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(2868),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(64521)),
			Timestamp: big.NewInt(5435254),
		},
	}

	msg4 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(2868),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(64521)),
			Timestamp: big.NewInt(5435254),
		},
	}

	messageStack := NewMessageStack()
	messageStack.DeliverMessage(msg1)
	messageStack.DeliverMessage(msg2)
	messageStack.DeliverMessage(msg3)
	messageStack.DeliverMessage(msg4)
	return messageStack
}

func TestBisection(t *testing.T) {
	messageStack := getStack()

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	sections, err := messageStack.GenerateBisection(bottomHash, 50, 4)
	if err != nil {
		t.Fatal(err)
	}

	for i, section := range sections {
		if section != messageStack.GetTopHash() {
			nextVal, err := messageStack.GetHashAtIndex(big.NewInt(int64(i + 1)))
			if err != nil {
				t.Fatal(err)
			}
			msg, err := messageStack.GenerateOneStepProof(section)
			if err != nil {
				t.Fatal(err)
			}
			if hash2(section, msg.CommitmentHash()) != nextVal {
				t.Error("Hashes not equal")
			}
		}
	}
}

func TestInboxInsert(t *testing.T) {
	pi := NewInbox()
	if pi.newest != nil {
		t.Error("newest of new Inbox should be nil")
	}
	pi2, err := marshalUnmarshal(pi)
	if err != nil {
		t.Error(err)
	}
	if pi.hashOfRest != pi2.hashOfRest {
		t.Error("marshal/unmarshal changes hash of empty inbox")
	}

	msg1 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(2868),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(64521)),
			Timestamp: big.NewInt(5435254),
		},
	}

	msg2 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(8741),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(1735)),
			Timestamp: big.NewInt(5435254),
		},
	}

	pi.DeliverMessage(msg1)
	msg1Delivered := pi.newest.message
	if !msg1Delivered.GetReceived().Equals(msg1) {
		t.Error("newest of Inbox wrong at val1")
	}
	if msg1Delivered.MessageNum.Cmp(big.NewInt(1)) != 0 {
		t.Error("msg 1 messageNum should have been 1, but was", msg1Delivered.MessageNum)
	}
	pi2, err = marshalUnmarshal(pi)
	if err != nil {
		t.Error(err)
	}
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item inbox")
	}

	pi.DeliverMessage(msg2)
	msg2Delivered := pi.newest.message
	if !msg2Delivered.GetReceived().Equals(msg2) {
		t.Error("newest of Inbox wrong at val2")
	}
	if msg2Delivered.MessageNum.Cmp(big.NewInt(2)) != 0 {
		t.Error("msg 2 messageNum should have been 2, but was", msg2Delivered.MessageNum)
	}
	pi2, err = marshalUnmarshal(pi)
	if err != nil {
		t.Error(err)
	}
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of two-item inbox")
	}

	pi.DiscardUpToCount(big.NewInt(0))
	pi2, err = marshalUnmarshal(pi)
	if err != nil {
		t.Error(err)
	}
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item inbox")
	}

	pi.DiscardUpToCount(big.NewInt(1))
	pi2, err = marshalUnmarshal(pi)
	if err != nil {
		t.Error(err)
	}
	if pi.newest.hash != pi2.newest.hash {
		t.Error("marshal/unmarshal changes hash of one-item inbox")
	}
}

func marshalUnmarshal(pi *Inbox) (*MessageStack, error) {
	ctx := ckptcontext.NewCheckpointContext()
	return pi.MarshalForCheckpoint(ctx).UnmarshalFromCheckpoint(ctx)
}
