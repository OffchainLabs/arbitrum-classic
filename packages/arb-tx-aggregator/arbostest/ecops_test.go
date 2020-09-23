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
	"crypto/rand"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func testPrecompile(t *testing.T, precompileNum byte, data []byte, correct []byte) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), common.RandAddress(), big.NewInt(0), chainTime))
	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(0),
			DestAddress: common.NewAddressFromEth(ethcommon.BytesToAddress([]byte{precompileNum})),
			Payment:     big.NewInt(0),
			Data:        data,
		}),
		common.RandAddress(),
		big.NewInt(1),
		chainTime,
	))

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	assertion, _ := mach.ExecuteAssertion(1000000000, inboxMessages, 0)
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

	if !bytes.Equal(res.ReturnData, correct) {
		t.Logf("Got result 0x%x", res.ReturnData)
		t.Logf("Wanted result 0x%x", correct)
		t.Error("calculated result incorrectly")
	}

	t.Logf("Used %v gas", res.GasUsed)
}

func TestECAdd(t *testing.T) {
	_, g1x, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	_, g1y, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	data := append(g1x.Marshal(), g1y.Marshal()...)
	correct := new(bn256.G1).Add(g1x, g1y).Marshal()

	testPrecompile(t, 6, data, correct)
}

func TestECMul(t *testing.T) {
	_, g1x, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	k := common.RandBigInt()
	data := append(g1x.Marshal(), math.U256Bytes(k)...)
	correct := new(bn256.G1).ScalarMult(g1x, k).Marshal()

	testPrecompile(t, 7, data, correct)
}
