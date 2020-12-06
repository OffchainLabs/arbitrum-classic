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
	"crypto/rand"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
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

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, inboxMessages, 0)

	results := processTxResults(t, assertion.ParseLogs())
	if len(results) != 1 {
		t.Fatal("unexpected log count", len(results))
	}

	res := results[0]
	succeededTxCheck(t, res)
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

func TestECRecover(t *testing.T) {
	pk, err := gethcrypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	hashedMsg := common.RandHash()
	sig, err := gethcrypto.Sign(hashedMsg[:], pk)
	if err != nil {
		t.Fatal(err)
	}
	sig[64] += 27

	data := hashedMsg.Bytes()
	data = append(data, math.U256Bytes(big.NewInt(int64(sig[64])))...)
	data = append(data, sig[:64]...)

	signer := gethcrypto.PubkeyToAddress(pk.PublicKey)
	ret := make([]byte, 32)
	copy(ret[12:], signer.Bytes())

	testPrecompile(t, 1, data, ret[:])
}

func TestSha256(t *testing.T) {
	data := common.RandBytes(100)
	sha256 := crypto.SHA256.New()
	_, err := sha256.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	hashedCorrect := sha256.Sum(nil)

	testPrecompile(t, 2, data, hashedCorrect)
}

func TestIdentityPrecompile(t *testing.T) {
	data := common.RandBytes(100)
	testPrecompile(t, 4, data, data)
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

func testECPairing(t *testing.T, g1Points []*bn256.G1, g2Points []*bn256.G2) {
	var data []byte

	if len(g1Points) != len(g2Points) {
		t.Fatal("input slices have different lengths")
	}

	for i, g1 := range g1Points {
		data = append(data, g1.Marshal()...)
		data = append(data, g2Points[i].Marshal()...)
	}

	isPair := bn256.PairingCheck(g1Points, g2Points)
	var correct []byte
	if isPair {
		correct = math.U256Bytes(big.NewInt(1))
	} else {
		correct = math.U256Bytes(big.NewInt(0))
	}

	testPrecompile(t, 8, data, correct)
}

func TestECPairing(t *testing.T) {
	_, p, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	_, q, err := bn256.RandomG2(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	s := common.RandBigInt()
	negOne := big.NewInt(-1)

	sP := new(bn256.G1).ScalarMult(p, s)
	sQ := new(bn256.G2).ScalarMult(q, s)

	g1Points := []*bn256.G1{
		new(bn256.G1).ScalarMult(p, negOne),
		sP,
	}

	g2Points := []*bn256.G2{
		sQ,
		q,
	}

	testECPairing(t, g1Points, g2Points)
}

func TestECPairingRandom(t *testing.T) {
	var g1Points []*bn256.G1
	var g2Points []*bn256.G2

	for i := 0; i < 5; i++ {
		_, g1, err := bn256.RandomG1(rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		_, g2, err := bn256.RandomG2(rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		g1Points = append(g1Points, g1)
		g2Points = append(g2Points, g2)

	}

	testECPairing(t, g1Points, g2Points)
}
