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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var contractCreator = common.HexToAddress("0xba59937520bd4c1067bac24fb774b981b4b8c115")
var connAddress = common.HexToAddress("0x9493d820aa2023afdedfc0eba1f86254a253ecdf")
var connAddress2 = common.HexToAddress("0x9276e6abd1b8cb06e5abd72db5140d216148bed3")

func testBasicTx(t *testing.T, msg message.SafeAbstractL2Message, msg2 message.SafeAbstractL2Message) ([]message.AbstractL2Message, *snapshot.Snapshot) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

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
		Data:        hexutil.MustDecode(arbostestcontracts.Receiver2Bin),
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

	var param common.Hash
	copy(param[12:], connAddress.Bytes())
	createTx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        append(hexutil.MustDecode(arbostestcontracts.ReceiverBin), param.Bytes()...),
	}

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(createTx2),
			contractCreator,
			big.NewInt(3),
			chainTime,
		),
	)

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(msg),
			sender,
			big.NewInt(4),
			chainTime,
		),
	)

	messages = append(
		messages,
		message.NewInboxMessage(
			message.NewSafeL2Message(msg2),
			sender,
			big.NewInt(5),
			chainTime,
		),
	)

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	results := processTxResults(t, assertion.ParseLogs())
	if len(results) != 4 {
		t.Fatal("incorrect log output count", len(results))
	}

	allResultsSucceeded(t, results)
	createRes := results[0]

	if !bytes.Equal(connAddress.Bytes(), createRes.ReturnData[12:]) {
		t.Fatal("incorrect created contract address")
	}

	createRes2 := results[1]
	if !bytes.Equal(connAddress2.Bytes(), createRes2.ReturnData[12:]) {
		t.Fatal("incorrect created contract address", hexutil.Encode(createRes2.ReturnData[12:]))
	}

	msgs := make([]message.AbstractL2Message, 0)
	for i, result := range results[2:] {
		if result.IncomingRequest.Sender != sender {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, sender)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		failIfError(t, err)

		targetHash := hashing.SoliditySHA3(hashing.Uint256(message.ChainAddressToID(chain)), hashing.Uint256(big.NewInt(int64(4+i))))
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
			DestAddress: connAddress2,
			Payment:     big.NewInt(10),
			Data:        hexutil.MustDecode("0x7795b5fc"),
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
	failIfError(t, err)
	if balance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("After call to non-contract, balance should still be 0, but was %v", balance)
	}

	balance2, err := snap.GetBalance(tx2.DestAddress)
	failIfError(t, err)
	if balance2.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("After call to contract, balance should still be 0, but was %v", balance2)
	}

	callRes, err := snap.Call(message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(callRes.ReturnData).Cmp(big.NewInt(7)) != 0 {
		t.Errorf("Storage was updated %X", callRes.ReturnData)
	}

	call2Res, err := snap.Call(message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress2,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(call2Res.ReturnData).Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Storage was updated")
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
			DestAddress: connAddress2,
			Payment:     big.NewInt(10),
			Data:        hexutil.MustDecode("0x7795b5fc"),
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
	failIfError(t, err)
	if balance.Cmp(tx.Payment) != 0 {
		t.Errorf("After call to non-contract, balance should be updated, but was %v", balance)
	}

	balance2, err := snap.GetBalance(tx2.DestAddress)
	failIfError(t, err)
	if balance2.Cmp(tx2.Payment) != 0 {
		t.Errorf("After call to contract, balance should be updated, but was %v", balance)
	}

	callRes, err := snap.Call(message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress2,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(callRes.ReturnData).Cmp(big.NewInt(6)) != 0 {
		t.Errorf("Storage wasn't updated %X", callRes.ReturnData)
	}

	callRes2, err := snap.Call(message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(callRes2.ReturnData).Cmp(big.NewInt(8)) != 0 {
		t.Errorf("Storage wasn't updated")
	}
}

func TestSignedTx(t *testing.T) {
	chain := common.RandAddress()
	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

	dest := common.RandAddress()
	pk, err := crypto.GenerateKey()
	failIfError(t, err)
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
	failIfError(t, err)

	l2msg, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx})
	failIfError(t, err)
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
	failIfError(t, err)

	l2msg2, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx2})
	failIfError(t, err)
	messages = append(
		messages,
		message.NewInboxMessage(
			l2msg2,
			common.RandAddress(),
			big.NewInt(2),
			chainTime,
		),
	)

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	testCase, err := inbox.TestVectorJSON(messages, logs, assertion.ParseOutMessages())
	failIfError(t, err)
	t.Log(string(testCase))
	results := processTxResults(t, assertion.ParseLogs())
	if len(results) != 2 {
		t.Fatal("incorrect log output count", len(results))
	}
	allResultsSucceeded(t, results)
	for i, result := range results {
		if result.IncomingRequest.Sender != addr {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, addr)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		failIfError(t, err)

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
	failIfError(t, err)

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

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	testCase, err := inbox.TestVectorJSON(messages, logs, assertion.ParseOutMessages())
	failIfError(t, err)
	t.Log(string(testCase))
	results := processTxResults(t, assertion.ParseLogs())
	if len(results) != 2 {
		t.Fatal("incorrect log output count", len(results))
	}
	allResultsSucceeded(t, results)
	for i, result := range results {
		if result.IncomingRequest.Sender != sender {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, sender)
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error("l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		failIfError(t, err)

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
	failIfError(t, err)

	runMessage(t, mach, initMsg(), chain)

	constructorData, err := hexutil.Decode(arbostestcontracts.FibonacciBin)
	failIfError(t, err)

	dest, err := deployContract(t, mach, common.RandAddress(), constructorData, big.NewInt(0), nil)
	failIfError(t, err)
	batchSize := 21
	senders := make([]common.Address, 0, batchSize)
	pks := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < batchSize; i++ {
		pk, err := crypto.GenerateKey()
		failIfError(t, err)
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
		failIfError(t, err)
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		senders = append(senders, addr)
		txes = append(txes, message.NewCompressedECDSAFromEth(signedTx))
		hashes = append(hashes, common.NewHashFromEth(signedTx.Hash()))
	}

	msg, err := message.NewTransactionBatchFromMessages(txes)
	failIfError(t, err)
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
		failIfError(t, err)
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
		t.Logf("message: %T", l2Message)
	}
}

func generateTestTransactions(t *testing.T, chain common.Address) []*types.Transaction {
	pk, err := crypto.GenerateKey()
	failIfError(t, err)

	tx := types.NewTransaction(0, common.RandAddress().ToEthAddress(), big.NewInt(1), 100000000000, big.NewInt(0), []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
	failIfError(t, err)

	tx2 := types.NewTransaction(1, common.RandAddress().ToEthAddress(), big.NewInt(0), 100000000000, big.NewInt(0), []byte{})
	signedTx2, err := types.SignTx(tx2, types.HomesteadSigner{}, pk)
	failIfError(t, err)

	tx3 := types.NewContractCreation(2, big.NewInt(0), 100000000000, big.NewInt(0), hexutil.MustDecode(arbostestcontracts.FibonacciBin))
	signedTx3, err := types.SignTx(tx3, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
	failIfError(t, err)
	return []*types.Transaction{signedTx, signedTx2, signedTx3}
}

func verifyTxLogs(t *testing.T, signer types.Signer, txes []*types.Transaction, logs []value.Value) {
	results := processTxResults(t, logs)
	allResultsSucceeded(t, results)
	for i, result := range results {
		sender, err := signer.Sender(txes[i])
		failIfError(t, err)
		if result.IncomingRequest.Sender.ToEthAddress() != sender {
			t.Error(i, "l2message had incorrect sender", result.IncomingRequest.Sender, sender.Hex())
		}
		if result.IncomingRequest.Kind != message.L2Type {
			t.Error(i, "l2message has incorrect type")
		}
		l2Message, err := message.L2Message{Data: result.IncomingRequest.Data}.AbstractMessage()
		failIfError(t, err)

		if result.IncomingRequest.MessageID.ToEthHash() != txes[i].Hash() {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.IncomingRequest.MessageID, txes[i].Hash().Hex())
		}

		_, ok := l2Message.(message.SignedTransaction)
		if !ok {
			t.Error("bad transaction format", l2Message)
		}
	}
}

func TestCompressedECDSATx(t *testing.T) {
	chain := common.RandAddress()
	t.Log("Chain address:", chain)
	t.Log("Chain ID:", message.ChainAddressToID(chain))

	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

	signer := types.NewEIP155Signer(message.ChainAddressToID(chain))

	txes := generateTestTransactions(t, chain)

	sender, err := signer.Sender(txes[0])
	failIfError(t, err)
	addr := common.NewAddressFromEth(sender)

	t.Log("Sender Address:", addr.Hex())

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

	for i, tx := range txes {
		l2msg, err := message.NewL2Message(message.NewCompressedECDSAFromEth(tx))
		failIfError(t, err)

		messages = append(
			messages,
			message.NewInboxMessage(
				l2msg,
				common.RandAddress(),
				big.NewInt(int64(2+i)),
				chainTime,
			),
		)
	}

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, messages, 0)
	logs := assertion.ParseLogs()
	testCase, err := inbox.TestVectorJSON(messages, logs, assertion.ParseOutMessages())
	failIfError(t, err)
	t.Log(string(testCase))
	if len(logs) != len(txes) {
		t.Fatal("incorrect log output count", len(logs))
	}
	verifyTxLogs(t, signer, txes, logs)
}
