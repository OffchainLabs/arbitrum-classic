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

package ethbridgetest

import (
	"github.com/pkg/errors"
	"math/rand"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
		Kind:        0,
		Sender:      common.RandAddress(),
		InboxSeqNum: common.RandBigInt(),
		Data:        common.RandBytes(200),
		ChainTime:   inbox.NewRandomChainTime(),
	}

	bridgeHash, err := messageTester.MessageHash(
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

	messageBridgeHash, err := messageTester.MessageValueHash(
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

	beforeInbox := common.RandHash()
	msgHash := common.RandHash()
	bridgeInboxAcc, err := messageTester.AddMessageToInbox(
		nil,
		beforeInbox,
		msgHash,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeInboxAcc != hashing.SoliditySHA3(
		hashing.Bytes32(beforeInbox),
		hashing.Bytes32(msgHash),
	) {
		t.Error("incorrect AddMessageToInbox")
	}
}
