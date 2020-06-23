/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"math/rand"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

func setupRand(t *testing.T) {
	currentTime := time.Now().Unix()
	t.Log("seed:", currentTime)
	rand.Seed(currentTime)
}

func TestTransactionMessage(t *testing.T) {
	setupRand(t)
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

func TestTransactionBatchSingleSender(t *testing.T) {
	setupRand(t)
	privateKey, _ := crypto.HexToECDSA(privHex)
	sender := common.NewAddressFromEth(crypto.PubkeyToAddress(privateKey.PublicKey))
	chain := addr3
	batchTx := message.NewRandomBatchTx(chain, privateKey)
	calculatedSender, err := tester.TransactionMessageBatchSingleSender(
		nil,
		big.NewInt(0),
		chain.ToEthAddress(),
		hashing.SoliditySHA3(batchTx.Data),
		batchTx.ToBytes(),
	)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	if calculatedSender != sender.ToEthAddress() {
		t.Error("Transaction sender not calculated correctly: got", hexutil.Encode(calculatedSender[:]), "instead of", hexutil.Encode(sender[:]))
	}
}

func TestTransactionBatchSingleValid(t *testing.T) {
	setupRand(t)
	privateKey, _ := crypto.HexToECDSA(privHex)
	sender := common.NewAddressFromEth(crypto.PubkeyToAddress(privateKey.PublicKey))
	chain := addr3
	batchTx := message.NewRandomBatchTx(chain, privateKey)

	txMessageHash, txReceiptHash, valid, err := tester.TransactionMessageBatchHashSingle(
		nil,
		big.NewInt(0),
		chain.ToEthAddress(),
		batchTx.ToBytes(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("message should have been valid")
	}

	tx := message.Transaction{
		Chain:       chain,
		To:          batchTx.To,
		From:        sender,
		SequenceNum: batchTx.SeqNum,
		Value:       batchTx.Value,
		Data:        batchTx.Data,
	}

	txValueHash := tx.AsInboxValue().Hash()

	t.Log("tx was", tx)

	t.Log("txMessageHash", hexutil.Encode(txMessageHash[:]))
	t.Log("txValueHash", hexutil.Encode(txValueHash[:]))

	if txMessageHash != tx.AsInboxValue().Hash() {
		t.Error("TransactionMessageBatchHashSingle result didn't match")
	}
	if txReceiptHash != tx.ReceiptHash() {
		t.Error("TransactionMessageBatchHashSingle tx receipt hash didn't match")
	}
}

func TestTransactionBatchSingleInvalid(t *testing.T) {
	setupRand(t)
	currentTime := time.Now().Unix()
	t.Log("seed:", currentTime)
	rand.Seed(currentTime)
	privateKey, _ := crypto.HexToECDSA(privHex)
	chain := addr3
	batchTx := message.NewRandomBatchTx(chain, privateKey)

	batchTx.Sig = [65]byte{1, 2, 3}

	_, _, valid, err := tester.TransactionMessageBatchHashSingle(
		nil,
		big.NewInt(0),
		chain.ToEthAddress(),
		batchTx.ToBytes(),
	)
	if err != nil {
		t.Fatal(err)
	}
	if valid {
		t.Fatal("message should have been valid")
	}
}

func TestTransactionBatchMessage(t *testing.T) {
	setupRand(t)
	privateKey, _ := crypto.HexToECDSA(privHex)
	chain := addr3
	batchTxData := make([]byte, 0)
	for i := 0; i < 10; i++ {
		batchTx := message.NewRandomBatchTx(chain, privateKey)
		if i%3 == 0 {
			batchTx.Sig[0]++
		}
		batchTxData = append(batchTxData, batchTx.ToBytes()...)
	}

	// Append some random junk
	batchTxData = append(batchTxData, []byte{54, 76, 23, 87, 34, 32, 87, 32}...)

	deliveryInfo := message.DeliveryInfo{
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(87962345)),
			Timestamp: big.NewInt(35463245),
		},
		TxId: big.NewInt(0),
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

	tup := value.NewEmptyTuple()
	preImage := tup.GetPreImage()

	bridgeInboxHash, err := tester.TransactionMessageBatchHash(
		nil,
		preImage.GetInnerHash(),
		big.NewInt(preImage.Size()),
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
	setupRand(t)
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
	setupRand(t)
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
	setupRand(t)
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
	setupRand(t)
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
		TxId: big.NewInt(98742),
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
		deliveryInfo.TxId,
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
		inboxPreImage.GetInnerHash(),
		big.NewInt(inboxPreImage.Size()),
		deliveryInfo.BlockNum.AsInt(),
		deliveryInfo.Timestamp,
		deliveryInfo.TxId,
		valPreimage.GetInnerHash(),
		big.NewInt(valPreimage.Size()),
	)
	if err != nil {
		t.Fatal(err)
	}

	if messageBridgeHash != message.AddToPrev(inbox, deliveredMsg).Hash().ToEthHash() {
		t.Error(errMsgHash)
	}
}
