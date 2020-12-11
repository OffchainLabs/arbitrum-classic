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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestDepositEthTx(t *testing.T) {
	depositDest := common.RandAddress()

	deployTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0))

	tx := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(1),
			DestAddress: depositDest,
			Payment:     big.NewInt(100),
			Data:        nil,
		}),
	}

	tx2 := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(2),
			DestAddress: connAddress1,
			Payment:     big.NewInt(100),
			Data:        nil,
		}),
	}

	messages := []message.Message{message.NewSafeL2Message(deployTx), tx, tx2}
	logs, _, mach := runAssertion(t, makeSimpleInbox(messages), 3, 0)
	results := processTxResults(t, logs)

	checkConstructorResult(t, results[0], connAddress1)
	succeededTxCheck(t, results[1])
	succeededTxCheck(t, results[2])

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(2))
	balance1, err := snap.GetBalance(depositDest)
	failIfError(t, err)
	if balance1.Cmp(big.NewInt(100)) != 00 {
		t.Fatal("wrong clone code length")
	}
}
