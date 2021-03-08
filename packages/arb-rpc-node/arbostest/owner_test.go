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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestOwner(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.GiveOwnershipData(common.RandAddress()),
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.GiveOwnershipData(sender),
	}

	tx3 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.StartArbOSUpgradeData(),
	}

	tx4 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.ContinueArbOSUpgradeData([]byte{}),
	}

	tx5 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.FinishArbOSUpgradeData(),
	}

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx1), sender, big.NewInt(0), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx2), owner, big.NewInt(1), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx3), sender, big.NewInt(2), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx4), sender, big.NewInt(3), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx5), sender, big.NewInt(4), big.NewInt(0), chainTime),
	}

	logs, _, _, _ := runAssertion(t, messages, 5, 0)
	results := processTxResults(t, logs)
	// Transfer from non-owner fails
	revertedTxCheck(t, results[0])
	succeededTxCheck(t, results[1])
	succeededTxCheck(t, results[2])
	succeededTxCheck(t, results[3])
	revertedTxCheck(t, results[4])
}
