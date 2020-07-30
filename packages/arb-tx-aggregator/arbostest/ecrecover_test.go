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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func TestECRecover(t *testing.T) {
	constructorData, err := hexutil.Decode(ECRecoverTestBin)
	if err != nil {
		t.Fatal(err)
	}

	ecTestABI, err := abi.JSON(strings.NewReader(ECRecoverTestABI))
	if err != nil {
		t.Fatal(err)
	}

	recoverSignerABI := ecTestABI.Methods["recoverSigner"]
	recoverSignerSignature, err := hexutil.Decode("0x2e295ec9")
	if err != nil {
		t.Fatal(err)
	}

	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	addr := common.Address{1, 2, 3, 4, 5}

	contractAddress := common.HexToAddress("0xba59937520bd4c1067bac24fb774b981b4b8c115")

	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	msg := common.RandBytes(200)

	hashedMsg := hashing.SoliditySHA3(msg)
	hashedPrefixMsg := hashing.SoliditySHA3WithPrefix(hashedMsg[:])
	sig, err := crypto.Sign(hashedPrefixMsg[:], pk)
	if err != nil {
		t.Fatal(err)
	}
	sig[64] += 27
	signer := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
	t.Log("Signer is", signer)
	t.Log("Sig is", hexutil.Encode(sig))

	recoverSignerData, err := recoverSignerABI.Inputs.Pack(msg, sig)
	if err != nil {
		t.Fatal(err)
	}

	initMsg := message.Init{
		ChainParams: valprotocol.ChainParams{
			StakeRequirement:        big.NewInt(0),
			GracePeriod:             common.TimeTicks{Val: big.NewInt(0)},
			MaxExecutionSteps:       0,
			ArbGasSpeedLimitPerTick: 0,
		},
		Owner:       common.Address{},
		ExtraConfig: []byte{},
	}
	inbox := value.NewEmptyTuple()

	inbox = value.NewTuple2(inbox, message.NewInboxMessage(initMsg, addr, big.NewInt(0), chainTime).AsValue())

	inbox = value.NewTuple2(inbox, message.NewInboxMessage(
		message.L2Message{Data: l2message.L2MessageAsData(makeConstructorTx(constructorData, big.NewInt(0)))},
		addr,
		big.NewInt(1),
		chainTime,
	).AsValue())

	inbox = value.NewTuple2(inbox, message.NewInboxMessage(
		message.L2Message{Data: l2message.L2MessageAsData(l2message.Transaction{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(1),
			DestAddress: contractAddress,
			Payment:     big.NewInt(0),
			Data:        append(recoverSignerSignature, recoverSignerData...),
		})},
		addr,
		big.NewInt(3),
		chainTime,
	).AsValue())

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	assertion, _ := mach.ExecuteAssertion(1000000000, inbox, 0)
	//data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(data))

	logs := assertion.ParseLogs()

	if len(logs) != 2 {
		t.Fatal("unexpected log count", len(logs))
	}

	for i, logVal := range assertion.ParseLogs() {
		res, err := evm.NewResultFromValue(logVal)
		if err != nil {
			t.Fatal(err)
		}

		// Recover log
		if i == 1 {
			t.Log("Recover data", string(res.ReturnData))

			vals, err := recoverSignerABI.Outputs.UnpackValues(res.ReturnData)
			if err != nil {
				t.Fatal(err)
			}

			if len(vals) != 1 {
				t.Fatal("unexpected tx result")
			}
			calculatedSigner := vals[0].(ethcommon.Address)
			if calculatedSigner != signer.ToEthAddress() {
				t.Fatal("recovered incorrect signer", calculatedSigner.Hex())
			}
		}

		if res.ResultCode != evm.ReturnCode {
			t.Error("tx failed", res.ResultCode)
		}
		log.Println("ReturnData", hexutil.Encode(res.ReturnData))
		if res.L1Message.Kind == message.L2Type {
			l2, err := l2message.NewL2MessageFromData(res.L1Message.Data)
			if err != nil {
				t.Fatal(err)
			}
			log.Println(l2)
		}

	}
}
