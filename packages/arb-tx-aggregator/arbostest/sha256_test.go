/*
* Copyright 2020, Offchain Labs, Inc.
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

package arbostest

import (
	"bytes"
	"crypto"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestSha256(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	addr := common.Address{1, 2, 3, 4, 5}

	data := common.RandBytes(100)

	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), addr, big.NewInt(0), chainTime))
	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(0),
			DestAddress: common.NewAddressFromEth(ethcommon.BytesToAddress([]byte{2})),
			Payment:     big.NewInt(0),
			Data:        data,
		}),
		addr,
		big.NewInt(1),
		chainTime,
	))

	sha256 := crypto.SHA256.New()
	sha256.Write(data)
	hashedCorrect := sha256.Sum(nil)

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	assertion, _ := mach.ExecuteAssertion(1000000000, inboxMessages, 0)
	//data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(data))

	logs := assertion.ParseLogs()

	if len(logs) != 1 {
		t.Fatal("unexpected log count", len(logs))
	}

	res, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode != evm.ReturnCode {
		t.Error("tx failed", res.ResultCode)
	}

	if res.IncomingRequest.Kind != message.L2Type {
		t.Fatal("wrong request type")
	}
	_, err = message.L2Message{Data: res.IncomingRequest.Data}.AbstractMessage()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(res.ReturnData, hashedCorrect) {
		t.Error("calculated hash incorrectly")
	}
}
