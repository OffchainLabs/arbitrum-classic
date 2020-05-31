/*
 * Copyright 2019, Offchain Labs, Inc.
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

package ethbridgetest

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

func TestTransactionMessage(t *testing.T) {
	msg := message.DeliveredTransaction{
		Transaction: message.Transaction{
			Chain:       addr3,
			To:          addr1,
			From:        addr2,
			SequenceNum: big.NewInt(74563),
			Value:       big.NewInt(89735406),
			Data:        []byte{65, 23, 68, 87, 12},
		},
		BlockNum:  common.NewTimeBlocks(big.NewInt(87962345)),
		Timestamp: big.NewInt(35463245),
	}
	bridgeHash, err := tester.TransactionHash(
		nil,
		msg.Chain.ToEthAddress(),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.SequenceNum,
		msg.Value,
		msg.Data,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, err := tester.TransactionMessageHash(
		nil,
		msg.Chain.ToEthAddress(),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.SequenceNum,
		msg.Value,
		msg.Data,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestTransactionBatchMessage(t *testing.T) {
	tx := message.Transaction{
		Chain:       addr3,
		To:          addr1,
		From:        common.NewAddressFromEth(auth.From),
		SequenceNum: big.NewInt(74563),
		Value:       big.NewInt(89735406),
		Data:        []byte{65, 23, 68, 87, 12},
	}

	offchainHash := message.BatchTxHash(
		tx.Chain,
		tx.To,
		tx.SequenceNum,
		tx.Value,
		tx.Data,
	)
	messageHash := hashing.SoliditySHA3WithPrefix(offchainHash[:])

	privateKey, err := crypto.HexToECDSA(privHex)
	if err != nil {
		t.Fatal(err)
	}

	sigBytes, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}
	var sig [65]byte
	copy(sig[:], sigBytes)

	batchTx := message.BatchTx{
		To:     tx.To,
		SeqNum: tx.SequenceNum,
		Value:  tx.Value,
		Data:   tx.Data,
		Sig:    sig,
	}

	batchTxData := batchTx.ToBytes()

	sender, err := tester.TransactionMessageBatchSingleSender(
		nil,
		big.NewInt(0),
		tx.Chain.ToEthAddress(),
		hashing.SoliditySHA3(tx.Data),
		batchTxData,
	)
	if err != nil {
		t.Error(err)
	}

	if sender != auth.From {
		t.Error("Transaction sender not calculated correctly: got", hexutil.Encode(sender[:]), "instead of", hexutil.Encode(auth.From[:]))
	}

	deliveredTx := message.DeliveredTransaction{
		Transaction: tx,
		BlockNum:    common.NewTimeBlocks(big.NewInt(87962345)),
		Timestamp:   big.NewInt(35463245),
	}

	msg := message.DeliveredTransactionBatch{
		TransactionBatch: message.TransactionBatch{
			Chain:  addr3,
			TxData: batchTxData,
		},
		BlockNum:  deliveredTx.BlockNum,
		Timestamp: deliveredTx.Timestamp,
	}

	bridgeHash, err := tester.TransactionBatchHash(
		nil,
		batchTxData,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	txMessageHash, err := tester.TransactionMessageBatchHashSingle(
		nil,
		big.NewInt(0),
		msg.Chain.ToEthAddress(),
		batchTxData,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}
	if txMessageHash != message.DeliveredValue(deliveredTx).Hash() {
		t.Error("TransactionMessageBatchHashSingle result didn't match")
	}

	tup := value.NewEmptyTuple()
	preImage := tup.GetPreImage()

	bridgeInboxHash, err := tester.TransactionMessageBatchHash(
		nil,
		preImage.HashImage,
		big.NewInt(preImage.Size),
		msg.Chain.ToEthAddress(),
		batchTxData,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}
	valInboxHash := message.AddToPrev(value.NewEmptyTuple(), msg).Hash()
	if bridgeInboxHash != valInboxHash.ToEthHash() {
		t.Error("TransactionMessageBatchHash result didn't match")
	}
}

func TestEthMessage(t *testing.T) {
	msg := message.DeliveredEth{
		Eth: message.Eth{
			To:    addr1,
			From:  addr2,
			Value: big.NewInt(89735406),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
		Timestamp:  big.NewInt(35463245),
		MessageNum: big.NewInt(98742),
	}
	bridgeHash, err := tester.EthHash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, err := tester.EthMessageHash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestERC20Message(t *testing.T) {
	msg := message.DeliveredERC20{
		ERC20: message.ERC20{
			To:           addr1,
			From:         addr2,
			TokenAddress: addr3,
			Value:        big.NewInt(89735406),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
		Timestamp:  big.NewInt(35463245),
		MessageNum: big.NewInt(98742),
	}
	bridgeHash, err := tester.Erc20Hash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, err := tester.Erc20MessageHash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Value,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestERC721Message(t *testing.T) {
	msg := message.DeliveredERC721{
		ERC721: message.ERC721{
			To:           addr1,
			From:         addr2,
			TokenAddress: addr3,
			Id:           big.NewInt(89735406),
		},
		BlockNum:   common.NewTimeBlocks(big.NewInt(87962345)),
		Timestamp:  big.NewInt(35463245),
		MessageNum: big.NewInt(98742),
	}
	bridgeHash, err := tester.Erc721Hash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Id,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, err := tester.Erc721MessageHash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Id,
		msg.BlockNum.AsInt(),
		msg.Timestamp,
		msg.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.DeliveredValue(msg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}
