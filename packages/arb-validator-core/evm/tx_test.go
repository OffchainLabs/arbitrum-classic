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

package evm

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/rand"
	"testing"
)

func newRandomTxInfo(r *Result) *TxInfo {
	return &TxInfo{
		TransactionIndex: rand.Uint64(),
		TransactionHash:  common.RandHash(),
		RawVal:           r.AsValue(),
		StartLogIndex:    rand.Uint64(),
		Block: &common.BlockId{
			Height:     common.NewTimeBlocks(common.RandBigInt()),
			HeaderHash: common.RandHash(),
		},
		Proof: &AVMLogProof{
			LogPreHash:   "",
			LogPostHash:  "",
			LogValHashes: nil,
		},
	}
}

func TestTxInfoMarshal(t *testing.T) {
	rand.Seed(43242)
	tx := newRandomTxInfo(NewRandomResult(message.NewRandomEth(), rand.Int31n(5)))

	txBuf := tx.Marshal()

	tx2, err := txBuf.Unmarshal()
	if err != nil {
		t.Fatal(err)
	}

	if !tx.Equals(tx2) {
		t.Fatal("not equal after unmarshal")
	}
}

func TestTxInfoToEthReceipt(t *testing.T) {
	rand.Seed(43242)
	l := newRandomTxInfo(NewRandomResult(message.NewRandomEth(), rand.Int31n(5)))
	_, err := l.ToEthReceipt()
	if err != nil {
		t.Fatal(err)
	}
}
