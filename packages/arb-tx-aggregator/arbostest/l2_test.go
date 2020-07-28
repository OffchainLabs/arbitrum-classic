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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
	"testing"
)

func TestSignedTx(t *testing.T) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	dest := common.RandAddress()
	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))

	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			simpleInitMessage(),
			chain,
			big.NewInt(0),
			chainTime,
		).AsValue(),
	)
	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			message.Eth{
				Dest:  addr,
				Value: big.NewInt(1000),
			},
			common.RandAddress(),
			big.NewInt(1),
			chainTime,
		).AsValue(),
	)

	tx := types.NewTransaction(0, dest.ToEthAddress(), big.NewInt(0), 100000000000, big.NewInt(0), []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(l2message.ChainAddressToID(chain)), pk)
	if err != nil {
		t.Fatal(err)
	}

	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			message.L2Message{Data: l2message.L2MessageAsData(l2message.NewSignedTransactionFromEth(signedTx))},
			common.RandAddress(),
			big.NewInt(2),
			chainTime,
		).AsValue(),
	)
	assertion, _ := mach.ExecuteAssertion(1000000000, inbox, 0)
	logs := assertion.ParseLogs()
	testCase, err := value.TestVectorJSON(inbox, logs, assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	if len(logs) != 1 {
		t.Fatal("incorrect log output count")
	}
	result, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if result.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected result code", result.ResultCode)
	}
	if result.L1Message.Sender != addr {
		t.Error("l2message had incorrect sender", result.L1Message.Sender, addr)
	}
	if result.L1Message.Kind != message.L2Type {
		t.Error("l2message has incorrect type")
	}
	l2Message, err := l2message.NewL2MessageFromData(result.L1Message.Data)
	if err != nil {
		t.Fatal(err)
	}

	if result.L1Message.MessageID().ToEthHash() != signedTx.Hash() {
		t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.L1Message.MessageID(), signedTx.Hash().Hex())
	}

	_, ok := l2Message.(l2message.SignedTransaction)
	if !ok {
		t.Error("bad transaction format")
	}
}

func TestUnsignedTx(t *testing.T) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			simpleInitMessage(),
			chain,
			big.NewInt(0),
			chainTime,
		).AsValue(),
	)
	sender := common.RandAddress()
	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			message.Eth{
				Dest:  sender,
				Value: big.NewInt(1000),
			},
			common.RandAddress(),
			big.NewInt(1),
			chainTime,
		).AsValue(),
	)

	tx1 := l2message.Transaction{
		MaxGas:      big.NewInt(100000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	tx2 := l2message.Transaction{
		MaxGas:      big.NewInt(100000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			message.L2Message{Data: l2message.L2MessageAsData(tx1)},
			sender,
			big.NewInt(2),
			chainTime,
		).AsValue(),
	)
	inbox = value.NewTuple2(
		inbox,
		message.NewInboxMessage(
			message.L2Message{Data: l2message.L2MessageAsData(tx2)},
			sender,
			big.NewInt(3),
			chainTime,
		).AsValue(),
	)
	assertion, _ := mach.ExecuteAssertion(1000000000, inbox, 0)
	logs := assertion.ParseLogs()
	testCase, err := value.TestVectorJSON(inbox, logs, assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	if len(logs) != 2 {
		t.Fatal("incorrect log output count")
	}
	for i, avmLog := range logs {
		result, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		if result.ResultCode != evm.ReturnCode {
			t.Fatal("unexpected result code", result.ResultCode)
		}
		if result.L1Message.Sender != sender {
			t.Error("l2message had incorrect sender", result.L1Message.Sender, sender)
		}
		if result.L1Message.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := l2message.NewL2MessageFromData(result.L1Message.Data)
		if err != nil {
			t.Fatal(err)
		}

		var correctHash common.Hash
		if i == 0 {
			correctHash = tx1.MessageID(sender, chain)
		} else {
			correctHash = tx2.MessageID(sender, chain)
		}
		if result.L1Message.MessageID() != correctHash {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.L1Message.MessageID(), correctHash)
		}

		_, ok := l2Message.(l2message.Transaction)
		if !ok {
			t.Error("bad transaction format")
		}
	}
}
