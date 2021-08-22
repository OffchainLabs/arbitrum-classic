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
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
		Data:        arbos.GetTotalOfEthBalances(),
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        GiveOwnershipData(common.RandAddress()),
	}

	tx3 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        GiveOwnershipData(sender),
	}

	tx4 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.StartArbOSUpgradeData(),
	}

	// Actual upgrade tested in dev/upgrade_test.go
	depositAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)
	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.RandAddress(),
				Payment:     depositAmount,
				Data:        nil,
			},
		}),
	}

	ib := InboxBuilder{}
	ib.AddMessage(initMsg(t, nil), common.Address{}, big.NewInt(0), chainTime)
	ib.AddMessage(deposit, chain, big.NewInt(0), chainTime)
	ib.AddMessage(message.NewSafeL2Message(tx1), owner, big.NewInt(0), chainTime)
	ib.AddMessage(message.NewSafeL2Message(tx2), sender, big.NewInt(0), chainTime)
	ib.AddMessage(message.NewSafeL2Message(tx3), owner, big.NewInt(0), chainTime)
	ib.AddMessage(message.NewSafeL2Message(tx4), sender, big.NewInt(0), chainTime)

	results, _ := runTxAssertion(t, ib.Messages)
	succeededTxCheck(t, results[0])
	succeededTxCheck(t, results[1])
	// Transfer from non-owner fails
	revertedTxCheck(t, results[2])
	succeededTxCheck(t, results[3])
	succeededTxCheck(t, results[4])

	totalBalance := new(big.Int).SetBytes(results[1].ReturnData)
	if totalBalance.Cmp(depositAmount) != 0 {
		t.Error("wrong total balance")
	}
}

func GiveOwnershipData(newOwnerAddr common.Address) []byte {
	return arbos.SetChainParameterData(arbos.ChainOwnerParamId, new(big.Int).SetBytes(newOwnerAddr.Bytes()))
}
