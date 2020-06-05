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
	"bytes"
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
	msg := message.Transaction{
		Chain:       addr3,
		To:          addr1,
		From:        addr2,
		SequenceNum: big.NewInt(74563),
		Value:       big.NewInt(89735406),
		Data:        []byte{65, 23, 68, 87, 12},
	}
	bridgeHash, err := tester.TransactionHash(
		nil,
		msg.Chain.ToEthAddress(),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.SequenceNum,
		msg.Value,
		msg.Data,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	messageBridgeHash, txReceiptHash, err := tester.TransactionMessageHash(
		nil,
		msg.Chain.ToEthAddress(),
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.SequenceNum,
		msg.Value,
		msg.Data,
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != msg.AsInboxValue().Hash().ToEthHash() {
		t.Error(errMsgHash)
	}

	if msg.ReceiptHash() != txReceiptHash {
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

	deliveryInfo := message.DeliveryInfo{
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(87962345)),
			Timestamp: big.NewInt(35463245),
		},
		MessageNum: big.NewInt(0),
	}

	msg := message.TransactionBatch{
		Chain:  addr3,
		TxData: batchTxData,
	}

	deliveredBatchMsg := message.Delivered{
		Message:      msg,
		DeliveryInfo: deliveryInfo,
	}

	bridgeHash, err := tester.TransactionBatchHash(
		nil,
		batchTxData,
	)
	if err != nil {
		t.Fatal(err)
	}
	if bridgeHash != msg.CommitmentHash().ToEthHash() {
		t.Error(errHash)
	}

	txMessageHash, txReceiptHash, err := tester.TransactionMessageBatchHashSingle(
		nil,
		big.NewInt(0),
		msg.Chain.ToEthAddress(),
		batchTxData,
	)
	if err != nil {
		t.Fatal(err)
	}
	if txMessageHash != tx.AsInboxValue().Hash() {
		t.Error("TransactionMessageBatchHashSingle result didn't match")
	}
	if txReceiptHash != tx.ReceiptHash() {
		t.Error("TransactionMessageBatchHashSingle tx receipt hash didn't match")
	}

	tup := value.NewEmptyTuple()
	preImage := tup.GetPreImage()

	bridgeInboxHash, err := tester.TransactionMessageBatchHash(
		nil,
		preImage.HashImage,
		big.NewInt(preImage.Size),
		msg.Chain.ToEthAddress(),
		batchTxData,
		deliveryInfo.BlockNum.AsInt(),
		deliveryInfo.Timestamp,
	)
	if err != nil {
		t.Fatal(err)
	}
	valInboxHash := message.AddToPrev(value.NewEmptyTuple(), deliveredBatchMsg).Hash()
	if bridgeInboxHash != valInboxHash.ToEthHash() {
		t.Error("TransactionMessageBatchHash result didn't match")
	}
}

func TestEthMessage(t *testing.T) {
	msg := message.Eth{
		To:    addr1,
		From:  addr2,
		Value: big.NewInt(89735406),
	}
	bridgeHash, err := tester.EthHash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.Value,
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
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != msg.AsInboxValue().Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestERC20Message(t *testing.T) {
	msg := message.ERC20{
		To:           addr1,
		From:         addr2,
		TokenAddress: addr3,
		Value:        big.NewInt(89735406),
	}
	bridgeHash, err := tester.Erc20Hash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Value,
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
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != msg.AsInboxValue().Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestERC721Message(t *testing.T) {
	msg := message.ERC721{
		To:           addr1,
		From:         addr2,
		TokenAddress: addr3,
		Id:           big.NewInt(89735406),
	}
	bridgeHash, err := tester.Erc721Hash(
		nil,
		msg.To.ToEthAddress(),
		msg.From.ToEthAddress(),
		msg.TokenAddress.ToEthAddress(),
		msg.Id,
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
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != msg.AsInboxValue().Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}

func TestDeliveredMessage(t *testing.T) {
	msg := message.ERC721{
		To:           addr1,
		From:         addr2,
		TokenAddress: addr3,
		Id:           big.NewInt(89735406),
	}
	deliveryInfo := message.DeliveryInfo{
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(87962345)),
			Timestamp: big.NewInt(35463245),
		},
		MessageNum: big.NewInt(98742),
	}
	deliveredMsg := message.Delivered{
		Message:      msg,
		DeliveryInfo: deliveryInfo,
	}
	inboxHash, err := tester.AddMessageToInbox(
		nil,
		common.Hash{},
		msg.CommitmentHash().ToEthHash(),
		deliveryInfo.BlockNum.AsInt(),
		deliveryInfo.Timestamp,
		deliveryInfo.MessageNum,
	)
	if err != nil {
		t.Fatal(err)
	}
	inboxHash2 := hashing.SoliditySHA3(hashing.Bytes32(common.Hash{}), hashing.Bytes32(deliveredMsg.CommitmentHash()))
	if inboxHash != inboxHash2 {
		t.Error(errHash)
	}

	var msgDataBuf bytes.Buffer
	if err := msg.AsInboxValue().MarshalForProof(&msgDataBuf); err != nil {
		t.Fatal(err)
	}

	inbox := value.NewEmptyTuple()
	inboxPreImage := inbox.GetPreImage()
	valPreimage := msg.AsInboxValue().GetPreImage()

	messageBridgeHash, err := tester.AddMessageToVMInboxHash(
		nil,
		inboxPreImage.HashImage,
		big.NewInt(inboxPreImage.Size),
		deliveryInfo.BlockNum.AsInt(),
		deliveryInfo.Timestamp,
		deliveryInfo.MessageNum,
		valPreimage.HashImage,
		big.NewInt(valPreimage.Size),
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.AddToPrev(inbox, deliveredMsg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}
