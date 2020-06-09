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

package message

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func generateBatchTx() BatchTx {
	addr1 := common.Address{}
	addr1[0] = 76
	addr1[19] = 93

	addr2 := common.Address{}
	addr2[0] = 43
	addr2[19] = 12

	return BatchTx{
		To:     addr2,
		SeqNum: big.NewInt(2),
		Value:  big.NewInt(43423),
		Data:   []byte{54, 87, 23},
		Sig:    [65]byte{87, 42, 56, 98},
	}

}

func generateTestBatch() TransactionBatch {
	addr1 := common.Address{}
	addr1[0] = 76
	addr1[19] = 93

	tx := generateBatchTx()
	return TransactionBatch{
		Chain:  addr1,
		TxData: tx.ToBytes(),
	}
}

func TestMarshalBatch(t *testing.T) {
	tx := generateBatchTx()
	data := tx.ToBytes()
	tx2, err := NewBatchTxFromData(data, 0)
	if err != nil {
		t.Error(err)
	}

	if !tx.Equals(tx2) {
		t.Error("unmarshaled tx not equal")
	}
}

func TestCheckpointBatch(t *testing.T) {
	msg := generateTestBatch()

	msg2, err := UnmarshalFromCheckpoint(msg.Type(), msg.CheckpointValue())
	if err != nil {
		t.Error(err)
	}

	if !msg.Equals(msg2) {
		t.Error("Unmarshalling didn't reverse marshalling", msg, msg2)
	}
}

func TestGetTransactions(t *testing.T) {
	addr1 := common.Address{}
	addr1[0] = 76
	addr1[19] = 93

	tx1 := generateBatchTx()
	tx2 := generateBatchTx()
	batch := TransactionBatch{
		Chain:  addr1,
		TxData: append(tx1.ToBytes(), tx2.ToBytes()...),
	}

	txes := batch.getBatchTransactions()
	if len(txes) != 2 {
		t.Fatal("didn't get back 2 txes")
	}

	if !txes[0].Equals(tx1) {
		t.Error("unmarshaled tx not equal")
	}

	if !txes[1].Equals(tx2) {
		t.Error("unmarshaled tx not equal")
	}
}
