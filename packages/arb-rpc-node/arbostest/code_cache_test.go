/*
* Copyright 2021, Offchain Labs, Inc.
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
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestCodeCache(t *testing.T) {
	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(0),
		Data:        common.RandBytes(1000),
	}
	//tx2 := message.Transaction{
	//	MaxGas:      big.NewInt(10000000),
	//	GasPriceBid: big.NewInt(0),
	//	SequenceNum: big.NewInt(1),
	//	DestAddress: common.Address{},
	//	Payment:     big.NewInt(0),
	//	Data:        hexutil.MustDecode(arbostestcontracts.FailedSendBin),
	//}

	messages := []message.Message{
		message.NewSafeL2Message(tx1),
		//message.NewSafeL2Message(tx2),
	}

	logs, _, _, assertion := runAssertionWithoutPrint(t, makeSimpleInbox(messages), len(messages), 0)
	results := processTxResults(t, logs)
	//checkConstructorResult(t, results[0], connAddress1)
	//checkConstructorResult(t, results[1], connAddress2)
	//if arbosVersion < 8 {
	//
	//}
	t.Log(results[0].FeeStats.UnitsUsed)
	//t.Log(results[1].FeeStats.UnitsUsed)
	t.Log("assertion gas", assertion.NumGas)
	//if results[1].FeeStats.UnitsUsed.L2Storage.Cmp(big.NewInt(0)) != 0 {
	//	t.Error("l2 storage used should be zero if contract is in cache")
	//}
	t.Fatal("error")
}
