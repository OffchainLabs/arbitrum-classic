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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

func testBasicTx(t *testing.T, msg message.AbstractL2Message, msg2 message.AbstractL2Message) ([]message.AbstractL2Message, *snapshot.Snapshot) {
	ethDeposit := message.Eth{
		Dest:  sender,
		Value: big.NewInt(100),
	}

	createTx := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.Receiver2Bin),
	}

	var param common.Hash
	copy(param[12:], connAddress1.Bytes())
	createTx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        append(hexutil.MustDecode(arbostestcontracts.ReceiverBin), param.Bytes()...),
	}

	l2Message, err := message.NewL2Message(msg)
	failIfError(t, err)

	l2Message2, err := message.NewL2Message(msg2)
	failIfError(t, err)

	messages := makeSimpleInbox([]message.Message{
		ethDeposit,
		message.NewSafeL2Message(createTx),
		message.NewSafeL2Message(createTx2),
		l2Message,
		l2Message2,
	})

	logs, _, snap, _ := runAssertion(t, messages, 4, 0)
	results := processTxResults(t, logs)

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[0], connAddress1)
	checkConstructorResult(t, results[1], connAddress2)

	msgs := make([]message.AbstractL2Message, 0)
	for i, result := range results[2:] {
		if result.IncomingRequest.Sender != sender {
			t.Error("l2message had incorrect sender", result.IncomingRequest.Sender, sender)
		}
		msg, err := message.NestedMessage(result.IncomingRequest.Data, result.IncomingRequest.Kind)
		failIfError(t, err)

		l2Msg, ok := msg.(message.L2Message)
		if !ok {
			t.Fatal("expected l2 message")
		}

		l2Message, err := l2Msg.AbstractMessage()
		failIfError(t, err)

		targetHash := hashing.SoliditySHA3(hashing.Uint256(message.ChainAddressToID(chain)), hashing.Uint256(big.NewInt(int64(4+i))))
		if result.IncomingRequest.MessageID != targetHash {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", l2Message, result.IncomingRequest.MessageID, targetHash)
		}

		msgs = append(msgs, l2Message)
	}
	return msgs, snap
}

func TestCallTx(t *testing.T) {
	tx := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}

	tx2 := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
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

	// After call to non-contract, balance should still be 0
	checkBalance(t, snap, tx.DestAddress, big.NewInt(0))
	// After call to contract, balance should still be 0
	checkBalance(t, snap, tx2.DestAddress, big.NewInt(0))

	callRes, err := snap.Call(message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress1,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(callRes.ReturnData).Cmp(big.NewInt(7)) != 0 {
		t.Errorf("Storage was updated %X", callRes.ReturnData)
	}

	call2Res, err := snap.Call(message.ContractTransaction{
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

	_, err = snap.Call(message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.Address{},
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode(arbostestcontracts.SimpleBin),
		},
	}, sender)
	failIfError(t, err)
}

func TestContractTx(t *testing.T) {
	tx := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(10),
			Data:        []byte{},
		},
	}

	tx2 := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
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

	checkBalance(t, snap, tx.DestAddress, tx.Payment)
	checkBalance(t, snap, tx2.DestAddress, tx2.Payment)

	callRes, err := snap.Call(message.ContractTransaction{
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

	callRes2, err := snap.Call(message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress1,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0xf8a8fd6d"),
		},
	}, common.Address{})
	failIfError(t, err)
	if new(big.Int).SetBytes(callRes2.ReturnData).Cmp(big.NewInt(8)) != 0 {
		t.Errorf("Storage wasn't updated")
	}
}

func TestUnsignedTx(t *testing.T) {
	ethDeposit := message.Eth{
		Dest:  sender,
		Value: big.NewInt(1000),
	}

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	messages := makeSimpleInbox([]message.Message{
		ethDeposit,
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	})

	logs, _, _, _ := runAssertion(t, messages, 2, 0)
	results := processTxResults(t, logs)
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
	batchSize := 21
	senders := make([]common.Address, 0, batchSize)
	pks := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < batchSize; i++ {
		pk, err := crypto.GenerateKey()
		failIfError(t, err)
		pks = append(pks, pk)
	}

	dest := common.RandAddress()
	txes := make([]message.AbstractL2Message, 0)
	hashes := make([]common.Hash, 0)
	batchSenderSeq := int64(0)
	for i := 0; i < 10; i++ {
		tx := message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(batchSenderSeq),
			DestAddress: dest,
			Payment:     big.NewInt(0),
			Data:        []byte{},
		}
		senders = append(senders, sender)
		txes = append(txes, tx)
		hashes = append(hashes, tx.MessageID(sender, chain))
		batchSenderSeq++
	}
	for _, pk := range pks[1:] {
		tx := types.NewTransaction(0, dest.ToEthAddress(), big.NewInt(0), 10000000, big.NewInt(0), []byte{})
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
		failIfError(t, err)
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		senders = append(senders, addr)
		txes = append(txes, message.NewCompressedECDSAFromEth(signedTx))
		hashes = append(hashes, common.NewHashFromEth(signedTx.Hash()))
	}

	msg, err := message.NewTransactionBatchFromMessages(txes)
	failIfError(t, err)

	var messages []message.Message
	for _, pk := range pks {
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		messages = append(messages, message.Eth{
			Dest:  addr,
			Value: big.NewInt(1000),
		})
	}
	messages = append(messages, message.NewSafeL2Message(msg))

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(txes), 0)
	results := processTxResults(t, logs)

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

	tx := types.NewTransaction(0, common.RandAddress().ToEthAddress(), big.NewInt(1), 10000000, big.NewInt(0), []byte{})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(chain)), pk)
	failIfError(t, err)

	tx2 := types.NewTransaction(1, common.RandAddress().ToEthAddress(), big.NewInt(0), 1000000, big.NewInt(0), []byte{})
	signedTx2, err := types.SignTx(tx2, types.HomesteadSigner{}, pk)
	failIfError(t, err)

	tx3 := types.NewContractCreation(2, big.NewInt(0), 1000000, big.NewInt(0), hexutil.MustDecode(arbostestcontracts.FibonacciBin))
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
	t.Log("Chain address:", chain)
	t.Log("Chain ID:", message.ChainAddressToID(chain))

	signer := types.NewEIP155Signer(message.ChainAddressToID(chain))

	txes := generateTestTransactions(t, chain)

	sender, err := signer.Sender(txes[0])
	failIfError(t, err)
	addr := common.NewAddressFromEth(sender)

	t.Log("Sender Address:", addr.Hex())

	messages := make([]message.Message, 0)
	messages = append(
		messages,
		message.Eth{
			Dest:  addr,
			Value: big.NewInt(1000),
		},
	)

	for _, tx := range txes {
		l2msg, err := message.NewL2Message(message.NewCompressedECDSAFromEth(tx))
		failIfError(t, err)
		messages = append(
			messages,
			l2msg,
		)
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(txes), 0)
	verifyTxLogs(t, signer, txes, logs)
}

func TestCall(t *testing.T) {
	tx1 := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0))
	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	tx2 := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress1,
			Payment:     big.NewInt(0),
			Data:        makeFuncData(t, simpleABI.Methods["exists"]),
		},
	}

	messages := []message.Message{
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	}
	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(messages), 0)
	results := processTxResults(t, logs)
	allResultsSucceeded(t, results)
	checkConstructorResult(t, results[0], connAddress1)
}
