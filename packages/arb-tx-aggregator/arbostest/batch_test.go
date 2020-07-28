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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

func TestBatch(t *testing.T) {
	mach, chain := initArbOS(t)

	pks := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < 20; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		depositResults := runMessage(
			t,
			mach,
			message.Eth{
				Dest:  common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey)),
				Value: big.NewInt(1000),
			},
			common.RandAddress(),
		)
		if len(depositResults) != 0 {
			t.Fatal("deposit should not have had a result")
		}
		pks = append(pks, pk)
	}

	batchSender := common.RandAddress()

	depositResults := runMessage(
		t,
		mach,
		message.Eth{
			Dest:  batchSender,
			Value: big.NewInt(1000),
		},
		common.RandAddress(),
	)
	if len(depositResults) != 0 {
		t.Fatal("deposit should not have had a result")
	}

	senders := make([]common.Address, 0)
	txes := make([]l2message.AbstractL2Message, 0)
	hashes := make([]common.Hash, 0)

	for _, pk := range pks {
		tx := types.NewTransaction(0, common.RandAddress().ToEthAddress(), big.NewInt(0), 100000000000, big.NewInt(0), []byte{})
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(l2message.ChainAddressToID(chain)), pk)
		if err != nil {
			t.Fatal(err)
		}
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		senders = append(senders, addr)
		txes = append(txes, l2message.NewSignedTransactionFromEth(signedTx))
		hashes = append(hashes, common.NewHashFromEth(signedTx.Hash()))
	}

	batchSenderSeq := int64(0)
	for i := 0; i < 10; i++ {
		tx := l2message.Transaction{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(batchSenderSeq),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(0),
			Data:        []byte{},
		}
		hashes = append(hashes, tx.MessageID(batchSender, chain))
		senders = append(senders, batchSender)
		txes = append(txes, tx)
		batchSenderSeq++
	}

	msg := l2message.NewTransactionBatchFromMessages(txes)
	results := runMessage(t, mach, message.L2Message{Data: l2message.L2MessageAsData(msg)}, batchSender)
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
		msg, err := l2message.NewL2MessageFromData(result.L1Message.Data)
		if err != nil {
			t.Fatal(err)
		}

		if result.L1Message.MessageID() != hashes[i] {
			t.Errorf("l2message of type %T had incorrect id %v instead of %v", msg, result.L1Message.MessageID(), hashes[i])
		} else {
			t.Log("correct request id")
		}

		switch msg := msg.(type) {
		case l2message.Transaction:
			match, ok := txes[i].(l2message.Transaction)
			if !ok {
				t.Error("l2 type didn't match")
			}
			if !msg.Equals(match) {
				t.Error("transaction didn't match input")
			}
		case l2message.SignedTransaction:
			match, ok := txes[i].(l2message.SignedTransaction)
			if !ok {
				t.Error("l2 type didn't match")
			}
			if !msg.Equals(match) {
				t.Error("signed transaction didn't match input")
			}
		default:
			t.Error("unexpected output type")
		}
	}
}
