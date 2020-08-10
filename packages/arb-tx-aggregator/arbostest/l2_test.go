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
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	messages := make([]inbox.InboxMessage, 0)
	messages = append(
		messages,
		message.NewInboxMessage(
			simpleInitMessage(),
			chain,
			big.NewInt(0),
			chainTime,
		),
	)
	messages = append(
		messages,
		message.NewInboxMessage(
			message.Eth{
				Dest:  addr,
				Value: big.NewInt(1000),
			},
			common.RandAddress(),
			big.NewInt(1),
			chainTime,
		),
	)

	tx := types.NewTransaction(0, dest.ToEthAddress(), big.NewInt(0), 100000000000, big.NewInt(0), []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
	if err != nil {
		t.Fatal(err)
	}

	l2, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx})
	if err != nil {
		t.Fatal(err)
	}
	messages = append(
		messages,
		message.NewInboxMessage(
			l2,
			common.RandAddress(),
			big.NewInt(2),
			chainTime,
		),
	)
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	testCase, err := inbox.TestVectorJSON(messages, logs, assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	if len(logs) != 1 {
		t.Fatal("incorrect log output count", len(logs))
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
	l2Message, err := message.L2Message{Data: result.L1Message.Data}.AbstractMessage()
	if err != nil {
		t.Fatal(err)
	}

	if result.L1Message.MessageID().ToEthHash() != signedTx.Hash() {
		t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.L1Message.MessageID(), signedTx.Hash().Hex())
	}

	_, ok := l2Message.(message.SignedTransaction)
	if !ok {
		t.Error("bad transaction format", l2Message)
	}
}

func TestUnsignedTx(t *testing.T) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	messages := make([]inbox.InboxMessage, 0)
	messages = append(
		messages,
		message.NewInboxMessage(
			simpleInitMessage(),
			chain,
			big.NewInt(0),
			chainTime,
		),
	)
	sender := common.RandAddress()
	messages = append(
		messages,
		message.NewInboxMessage(
			message.Eth{
				Dest:  sender,
				Value: big.NewInt(1000),
			},
			common.RandAddress(),
			big.NewInt(1),
			chainTime,
		),
	)

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(100000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(100000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(tx1),
			sender,
			big.NewInt(2),
			chainTime,
		),
	)
	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(tx2),
			sender,
			big.NewInt(3),
			chainTime,
		),
	)
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	testCase, err := inbox.TestVectorJSON(messages, logs, assertion.ParseOutMessages())
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
		l2Message, err := message.L2Message{Data: result.L1Message.Data}.AbstractMessage()
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
		_, ok := l2Message.(message.Transaction)
		if !ok {
			t.Error("bad transaction format")
		}
	}
}

func TestBatch(t *testing.T) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
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
	results := runMessage(t, mach, initMsg, chain)
	log.Println(results)

	constructorData, err := hexutil.Decode(FibonacciBin)
	if err != nil {
		t.Fatal(err)
	}

	dest, err := deployContract(t, mach, common.RandAddress(), constructorData, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	batchSize := 21
	senders := make([]common.Address, 0, batchSize)
	pks := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < batchSize; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		depositResults := runMessage(
			t,
			mach,
			message.Eth{
				Dest:  addr,
				Value: big.NewInt(1000),
			},
			addr,
		)
		if len(depositResults) != 0 {
			t.Fatal("deposit should not have had a result")
		}
		pks = append(pks, pk)
	}
	batchSender := common.NewAddressFromEth(crypto.PubkeyToAddress(pks[0].PublicKey))
	txes := make([]message.AbstractL2Message, 0)
	hashes := make([]common.Hash, 0)
	batchSenderSeq := int64(0)
	for i := 0; i < 10; i++ {
		tx := message.Transaction{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(batchSenderSeq),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        []byte{},
		}
		senders = append(senders, batchSender)
		txes = append(txes, tx)
		hashes = append(hashes, tx.MessageID(batchSender, chain))
		batchSenderSeq++
	}
	for _, pk := range pks[1:] {
		tx := types.NewTransaction(0, dest.ToEthAddress(), big.NewInt(0), 100000000000, big.NewInt(0), []byte{})
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
		if err != nil {
			t.Fatal(err)
		}
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		senders = append(senders, addr)
		txes = append(txes, message.SignedTransaction{Tx: signedTx})
		hashes = append(hashes, common.NewHashFromEth(signedTx.Hash()))
	}

	msg, err := message.NewTransactionBatchFromMessages(txes)
	if err != nil {
		t.Fatal(err)
	}
	results = runMessage(t, mach, message.NewSafeL2Message(msg), batchSender)
	if len(results) != len(txes) {
		t.Fatal("incorrect result count", len(results), "instead of", len(txes))
	}
	for i, result := range results {
		if result.L1Message.Sender != senders[i] {
			t.Error("l2message had incorrect sender", result.L1Message.Sender, senders[i])
		}
		if result.L1Message.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		if result.L1Message.MessageID() != hashes[i] {
			t.Error("l2message had incorrect id", result.L1Message.MessageID(), hashes[i])
		}
		l2Message, err := message.L2Message{Data: result.L1Message.Data}.AbstractMessage()
		if err != nil {
			t.Fatal(err)
		}
		if i < 10 {
			_, ok := l2Message.(message.Transaction)
			if !ok {
				t.Error("bad transaction format")
			}
		} else {
			_, ok := l2Message.(message.SignedTransaction)
			if !ok {
				t.Error("bad transaction format")
			}
		}
		log.Printf("message: %T\n", l2Message)
	}
}
