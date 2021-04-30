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
	"context"
	"crypto"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func testPrecompile(t *testing.T, precompileNum byte, data []byte, correct []byte) {
	precompileAddress := ethcommon.BytesToAddress([]byte{precompileNum})

	ctx := context.Background()
	backend, _ := test.SimulatedBackend(t)

	ethCall := ethereum.CallMsg{
		From:     ethcommon.Address{},
		To:       &precompileAddress,
		Gas:      1000000,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     data,
	}
	ethRes, err := backend.CallContract(ctx, ethCall, nil)
	failIfError(t, err)

	tx := message.Transaction{
		MaxGas:      big.NewInt(100000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(precompileAddress),
		Payment:     big.NewInt(0),
		Data:        data,
	}

	messages := []message.Message{message.NewSafeL2Message(tx)}
	logs, _, _ := runSimpleAssertion(t, messages)
	results := processTxResults(t, logs)

	res := results[0]
	succeededTxCheck(t, res)
	if res.IncomingRequest.Kind != message.L2Type {
		t.Fatal("wrong request type")
	}
	_, err = message.L2Message{Data: res.IncomingRequest.Data}.AbstractMessage()
	failIfError(t, err)

	if !bytes.Equal(res.ReturnData, correct) {
		t.Logf("Got result 0x%x", res.ReturnData)
		t.Logf("Wanted result 0x%x", correct)
		t.Error("calculated result incorrectly")
	}

	if !bytes.Equal(ethRes, correct) {
		t.Logf("Got result 0x%x", res.ReturnData)
		t.Logf("Wanted result 0x%x", correct)
		t.Error("calculated result incorrectly")
	}

	ethGas, err := backend.EstimateGas(ctx, ethCall)
	failIfError(t, err)

	t.Logf("Arb gas = %v", res.GasUsed)
	t.Logf("Eth gas = %v", ethGas)
}

func TestECRecover(t *testing.T) {
	pk, err := gethcrypto.GenerateKey()
	failIfError(t, err)

	hashedMsg := common.RandHash()
	sig, err := gethcrypto.Sign(hashedMsg[:], pk)
	failIfError(t, err)
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
	failIfError(t, err)
	hashedCorrect := sha256.Sum(nil)

	testPrecompile(t, 2, data, hashedCorrect)
}

func TestIdentityPrecompile(t *testing.T) {
	data := common.RandBytes(100)
	testPrecompile(t, 4, data, data)
}

func TestModExp(t *testing.T) {
	x := common.RandBigInt()
	y := common.RandBigInt()
	m := common.RandBigInt()
	correct := new(big.Int).Exp(x, y, m)

	xLen := new(big.Int).SetInt64(int64(len(x.Bytes())))
	yLen := new(big.Int).SetInt64(int64(len(y.Bytes())))
	mLen := new(big.Int).SetInt64(int64(len(m.Bytes())))

	var data []byte
	data = append(data, math.U256Bytes(xLen)...)
	data = append(data, math.U256Bytes(yLen)...)
	data = append(data, math.U256Bytes(mLen)...)
	data = append(data, x.Bytes()...)
	data = append(data, y.Bytes()...)
	data = append(data, m.Bytes()...)
	testPrecompile(t, 5, data, correct.Bytes())
}

func TestECAdd(t *testing.T) {
	_, g1x, err := bn256.RandomG1(rand.Reader)
	failIfError(t, err)
	_, g1y, err := bn256.RandomG1(rand.Reader)
	failIfError(t, err)
	data := append(g1x.Marshal(), g1y.Marshal()...)
	correct := new(bn256.G1).Add(g1x, g1y).Marshal()

	testPrecompile(t, 6, data, correct)
}

func TestECMul(t *testing.T) {
	_, g1x, err := bn256.RandomG1(rand.Reader)
	failIfError(t, err)
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
	failIfError(t, err)
	_, q, err := bn256.RandomG2(rand.Reader)
	failIfError(t, err)

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
		failIfError(t, err)
		_, g2, err := bn256.RandomG2(rand.Reader)
		failIfError(t, err)
		g1Points = append(g1Points, g1)
		g2Points = append(g2Points, g2)
	}

	testECPairing(t, g1Points, g2Points)
}
