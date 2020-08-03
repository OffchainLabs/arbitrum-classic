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

package message

import (
	"bytes"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var errHash = errors.New("ethbridge calculated wrong hash")
var errMsgHash = errors.New("ethbridge calculated wrong l2message hash")

func setupRand(t *testing.T) {
	currentTime := time.Now().Unix()
	t.Log("seed:", currentTime)
	rand.Seed(currentTime)
}

func TestMessage(t *testing.T) {
	setupRand(t)

	msg := inbox.InboxMessage{
		Kind:        L2Type,
		Sender:      common.RandAddress(),
		InboxSeqNum: common.RandBigInt(),
		Data:        common.RandBytes(200),
		ChainTime:   inbox.NewRandomChainTime(),
	}

	bridgeHash, err := tester.MessageHash(
		nil,
		uint8(msg.Kind),
		msg.Sender.ToEthAddress(),
		msg.ChainTime.BlockNum.AsInt(),
		msg.ChainTime.Timestamp,
		msg.InboxSeqNum,
		hashing.SoliditySHA3(msg.Data),
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, err := tester.MessageValueHash(
		nil,
		uint8(msg.Kind),
		msg.ChainTime.BlockNum.AsInt(),
		msg.ChainTime.Timestamp,
		msg.Sender.ToEthAddress(),
		msg.InboxSeqNum,
		msg.Data,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != msg.AsValue().Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestDeliveredMessage(t *testing.T) {
	setupRand(t)

	startInbox := value.NewEmptyTuple()
	startInboxPre := startInbox.GetPreImage()
	msg := value.NewTuple2(value.NewInt64Value(2), value.NewInt64Value(6))
	msgPre := msg.GetPreImage()
	bridgeVMInboxHash, err := tester.AddMessageToVMInboxHash(
		nil,
		startInboxPre.GetInnerHash(),
		big.NewInt(startInboxPre.Size()),
		msgPre.GetInnerHash(),
		big.NewInt(msgPre.Size()),
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeVMInboxHash != value.NewTuple2(startInbox, msg).Hash() {
		t.Error("incorrect AddMessageToVMInboxHash")
	}

	beforeInbox := common.RandHash()
	msgHash := common.RandHash()
	bridgeInboxHash, err := tester.AddMessageToInbox(
		nil,
		beforeInbox,
		msgHash,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeInboxHash != hashing.SoliditySHA3(
		hashing.Bytes32(beforeInbox),
		hashing.Bytes32(msgHash),
	) {
		t.Error("incorrect AddMessageToInbox")
	}
}

func TestUnmarshalOutgoing(t *testing.T) {
	msg := NewRandomOutMessage(NewRandomEth())
	var valData bytes.Buffer
	if err := value.MarshalValue(msg.AsValue(), &valData); err != nil {
		t.Fatal(err)
	}
	valid, offset, kind, sender, data, err := tester.UnmarshalOutgoingMessage(nil, valData.Bytes(), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("invalid l2message")
	}
	if offset.Uint64() != uint64(len(valData.Bytes())) {
		t.Error("incorrect offset")
	}
	if inbox.Type(kind) != msg.Kind {
		t.Error("incorrect l2message type")
	}
	if sender != msg.Sender.ToEthAddress() {
		t.Error("incorrect sender")
	}
	if !bytes.Equal(data, msg.Data) {
		t.Error("incorrect data")
	}
}

func TestParseEthMessage(t *testing.T) {
	msg := NewRandomEth()
	ret, err := tester.ParseEthMessage(nil, msg.AsData())
	if err != nil {
		t.Fatal(err)
	}
	if !ret.Valid {
		t.Error("invalid l2message")
	}
	parsedMsg := ret.Message
	if parsedMsg.Value.Cmp(msg.Value) != 0 {
		t.Error("incorrect value")
	}
	if parsedMsg.Dest != msg.Dest.ToEthAddress() {
		t.Error("incorrect address")
	}
}

func TestParseERC20Message(t *testing.T) {
	msg := NewRandomERC20()
	ret, err := tester.ParseERC20Message(nil, msg.AsData())
	if err != nil {
		t.Fatal(err)
	}
	if !ret.Valid {
		t.Error("invalid l2message")
	}
	parsedMsg := ret.Message
	if parsedMsg.Token != msg.Token.ToEthAddress() {
		t.Error("incorrect token")
	}
	if parsedMsg.Value.Cmp(msg.Value) != 0 {
		t.Error("incorrect value")
	}
	if parsedMsg.Dest != msg.Dest.ToEthAddress() {
		t.Error("incorrect address")
	}
}

func TestParseERC721Message(t *testing.T) {
	msg := NewRandomERC721()
	ret, err := tester.ParseERC721Message(nil, msg.AsData())
	if err != nil {
		t.Error(err)
	}
	if !ret.Valid {
		t.Error("invalid l2message")
	}
	parsedMsg := ret.Message
	if parsedMsg.Token != msg.Token.ToEthAddress() {
		t.Error("incorrect token")
	}
	if parsedMsg.Id.Cmp(msg.ID) != 0 {
		t.Error("incorrect value")
	}
	if parsedMsg.Dest != msg.Dest.ToEthAddress() {
		t.Error("incorrect address")
	}
}
