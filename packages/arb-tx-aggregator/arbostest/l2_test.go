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
	"crypto/ecdsa"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"log"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var contractCreator = common.HexToAddress("0xba59937520bd4c1067bac24fb774b981b4b8c115")
var connAddress = common.HexToAddress("0x9493d820aa2023afdedfc0eba1f86254a253ecdf")

func testBasicTx(t *testing.T, msg message.SafeAbstractL2Message, msg2 message.SafeAbstractL2Message) ([]message.AbstractL2Message, *snapshot.Snapshot) {
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

	sender := common.RandAddress()

	messages = append(
		messages,
		message.NewInboxMessage(
			initMsg(),
			chain,
			big.NewInt(0),
			chainTime,
		),
	)
	messages = append(
		messages,
		message.NewInboxMessage(
			message.Eth{
				Dest:  sender,
				Value: big.NewInt(100),
			},
			chain,
			big.NewInt(1),
			chainTime,
		),
	)

	createTx := message.Transaction{
		MaxGas:      big.NewInt(10000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.ReceiverBin),
	}

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(createTx),
			contractCreator,
			big.NewInt(2),
			chainTime,
		),
	)

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(msg),
			sender,
			big.NewInt(3),
			chainTime,
		),
	)

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(msg2),
			sender,
			big.NewInt(4),
			chainTime,
		),
	)

	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	if len(logs) != 3 {
		t.Fatal("incorrect log output count", len(logs))
	}

	createRes, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(connAddress.Bytes(), createRes.ReturnData[12:]) {
		t.Fatal("incorrect created contract address")
	}

	msgs := make([]message.AbstractL2Message, 0)
	for i, avmLog := range logs[1:] {
		result, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		if result.ResultCode != evm.ReturnCode {
			t.Fatal("unexpected result code", result.ResultCode)
		}
		if result.IncomingRequest.Sender != sender {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, sender)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		if err != nil {
			t.Fatal(err)
		}

		targetHash := hashing.SoliditySHA3(hashing.Uint256(message.ChainAddressToID(chain)), hashing.Uint256(big.NewInt(int64(3+i))))
		if result.IncomingRequest.MessageID != targetHash {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.IncomingRequest.MessageID, targetHash)
		}

		msgs = append(msgs, l2Message)
	}

	snap := snapshot.NewSnapshot(mach.Clone(), inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}, message.ChainAddressToID(chain), big.NewInt(4))

	return msgs, snap
}

func TestCallTx(t *testing.T) {
	tx := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}

	tx2 := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress,
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}
	msgs, snap := testBasicTx(t, tx, tx2)

	for _, l2Message := range msgs {
		_, ok := l2Message.(message.Call)
		if !ok {
			t.Error("bad transaction format")
		}
	}

	balance, err := snap.GetBalance(tx.DestAddress)
	if err != nil {
		t.Fatal(err)
	}
	if balance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("After call to non-contract, balance should still be 0, but was %v", balance)
	}

	balance2, err := snap.GetBalance(tx2.DestAddress)
	if err != nil {
		t.Fatal(err)
	}
	if balance2.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("After call to contract, balance should still be 0, but was %v", balance2)
	}
}

func TestContractTx(t *testing.T) {
	tx := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}

	tx2 := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress,
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}
	msgs, snap := testBasicTx(t, tx, tx2)

	for _, l2Message := range msgs {
		_, ok := l2Message.(message.ContractTransaction)
		if !ok {
			t.Error("bad transaction format")
		}
	}

	balance, err := snap.GetBalance(tx.DestAddress)
	if err != nil {
		t.Fatal(err)
	}
	if balance.Cmp(tx.Payment) != 0 {
		t.Errorf("After call to non-contract, balance should be updated, but was %v", balance)
	}

	balance2, err := snap.GetBalance(tx2.DestAddress)
	if err != nil {
		t.Fatal(err)
	}
	if balance2.Cmp(tx2.Payment) != 0 {
		t.Errorf("After call to contract, balance should be updated, but was %v", balance)
	}
}

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
			initMsg(),
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

	l2msg, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx})
	if err != nil {
		t.Fatal(err)
	}
	messages = append(
		messages,
		message.NewInboxMessage(
			l2msg,
			common.RandAddress(),
			big.NewInt(2),
			chainTime,
		),
	)

	tx2 := types.NewContractCreation(1, big.NewInt(0), 100000000000, big.NewInt(0), hexutil.MustDecode(arbostestcontracts.FibonacciBin))
	signedTx2, err := types.SignTx(tx2, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
	if err != nil {
		t.Fatal(err)
	}

	l2msg2, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx2})
	if err != nil {
		t.Fatal(err)
	}
	messages = append(
		messages,
		message.NewInboxMessage(
			l2msg2,
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
	if len(logs) != 2 {
		t.Fatal("incorrect log output count", len(logs))
	}
	for i, avmLog := range logs {
		result, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		if result.ResultCode != evm.ReturnCode {
			t.Fatal("unexpected result code", result.ResultCode)
		}
		if result.IncomingRequest.Sender != addr {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, addr)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		if err != nil {
			t.Fatal(err)
		}

		var correctHash ethcommon.Hash
		if i == 0 {
			correctHash = signedTx.Hash()
		} else {
			correctHash = signedTx2.Hash()
		}

		if result.IncomingRequest.MessageID.ToEthHash() != correctHash {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.IncomingRequest.MessageID, correctHash.Hex())
		}

		_, ok := l2Message.(message.SignedTransaction)
		if !ok {
			t.Error("bad transaction format", l2Message)
		}
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
			initMsg(),
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
		t.Fatal("incorrect log output count", len(logs))
	}
	for i, avmLog := range logs {
		result, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		if result.ResultCode != evm.ReturnCode {
			t.Fatal("unexpected result code", result.ResultCode)
		}
		if result.IncomingRequest.Sender != sender {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, sender)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		if err != nil {
			t.Fatal(err)
		}

		var correctHash common.Hash
		if i == 0 {
			correctHash = tx1.MessageID(sender, chain)
		} else {
			correctHash = tx2.MessageID(sender, chain)
		}
		if result.IncomingRequest.MessageID != correctHash {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.IncomingRequest.MessageID, correctHash)
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

	runMessage(t, mach, initMsg(), chain)

	constructorData, err := hexutil.Decode(arbostestcontracts.FibonacciBin)
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
		depositEth(t, mach, addr, big.NewInt(1000))
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
	results, sends := runMessage(t, mach, message.NewSafeL2Message(msg), batchSender)
	if len(results) != len(txes) {
		t.Fatal("incorrect result count", len(results), "instead of", len(txes))
	}
	if len(sends) != 0 {
		t.Fatal("incorrect send count", len(sends), "instead of 0")
	}
	for i, result := range results {
		if result.IncomingRequest.Sender != senders[i] {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, senders[i])
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		if result.IncomingRequest.MessageID != hashes[i] {
			t.Error("l2message had incorrect id", result.IncomingRequest.MessageID, hashes[i])
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
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
