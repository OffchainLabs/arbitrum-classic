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
	"math/big"
	"math/rand"
	"testing"
)

func generateSampleTxInfo() (TxInfo, error) {
	var txHash common.Hash
	var nodeHash common.Hash

	slicesToFill := [][]byte{txHash[:], nodeHash[:]}
	for _, sl := range slicesToFill {
		_, err := rand.Read(sl)
		if err != nil {
			return TxInfo{}, err
		}
	}

	delivered, _ := message.NewSingleDelivered(message.Delivered{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: new(big.Int).SetUint64(rand.Uint64()),
		},
		DeliveryInfo: message.DeliveryInfo{
			ChainTime: message.ChainTime{
				BlockNum:  common.NewTimeBlocks(new(big.Int).SetUint64(rand.Uint64())),
				Timestamp: new(big.Int).SetUint64(rand.Uint64()),
			},
			TxId: new(big.Int).SetUint64(rand.Uint64()),
		},
	})

	val, err := NewVMResultValue(
		delivered,
		StopCode,
		[]byte{},
		[]Log{},
	)
	if err != nil {
		return TxInfo{}, err
	}

	return TxInfo{
		Found:            true,
		NodeHeight:       rand.Uint64(),
		NodeHash:         nodeHash,
		TransactionIndex: rand.Uint64(),
		TransactionHash:  txHash,
		RawVal:           val,
		LogsPreHash:      "",
		LogsPostHash:     "",
		LogsValHashes:    nil,
		OnChainTxHash:    common.Hash{},
	}, nil
}

func TestTxInfoMarshal(t *testing.T) {
	rand.Seed(43242)
	tx, err := generateSampleTxInfo()
	if err != nil {
		t.Fatal(err)
	}

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
	l, err := generateSampleTxInfo()
	if err != nil {
		t.Fatal(err)
	}

	_, err = l.ToEthReceipt(common.Address{65, 43, 65})
	if err != nil {
		t.Fatal(err)
	}
}
