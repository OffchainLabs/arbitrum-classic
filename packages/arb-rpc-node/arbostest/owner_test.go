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

	depositAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	fundAddress := func(address common.Address) message.EthDepositTx {
		return message.EthDepositTx{
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
	}

	addAddress := func(address common.Address, sequenceNum int) message.L2Message {
		return message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(int64(sequenceNum)),
			DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        arbos.AddChainOwnerData(message.L2RemapAccount(address)),
		})
	}

	removeAddress := func(address common.Address, sequenceNum int) message.L2Message {
		return message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(int64(sequenceNum)),
			DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        arbos.RemoveChainOwnerData(message.L2RemapAccount(address)),
		})
	}
	
	var shouldSucceed []bool;
	
	ib := InboxBuilder{}
	ib.AddMessage(initMsg(t, nil), common.Address{}, big.NewInt(0), chainTime)

	// Add and remove a bunch of random owners
	
	priorRandomOwner := owner
	ib.AddMessage(addAddress(priorRandomOwner, 0), priorRandomOwner, big.NewInt(0), chainTime)
	shouldSucceed = append(shouldSucceed, true)
	
	for i := 0; i < 24; i++ {

		println(len(shouldSucceed))

		randomOwner := common.RandAddress()

		ib.AddMessage(fundAddress(randomOwner), randomOwner, big.NewInt(0), chainTime)
		ib.AddMessage(addAddress(randomOwner, i % 2 + 1), priorRandomOwner, big.NewInt(0), chainTime)
		shouldSucceed = append(shouldSucceed, true, true)

		if i % 2 == 1 {
			ib.AddMessage(removeAddress(priorRandomOwner, 0), randomOwner, big.NewInt(0), chainTime)
			ib.AddMessage(removeAddress(randomOwner, 3), priorRandomOwner, big.NewInt(0), chainTime)
			shouldSucceed = append(shouldSucceed, true, false)
			
			priorRandomOwner = randomOwner
		}
	}

	// make sender an owner, then have sender remove prior owner
	ib.AddMessage(addAddress(sender, 4), owner, big.NewInt(0), chainTime)
	ib.AddMessage(addAddress(sender, 1), priorRandomOwner, big.NewInt(0), chainTime)
	ib.AddMessage(removeAddress(priorRandomOwner, 0), sender, big.NewInt(0), chainTime)
	shouldSucceed = append(shouldSucceed, false, true, true)

	// try to start an upgrade with the new owner, which you can only do if you're an owner.
	// this won't actually perform an upgrade (we test this elsewhere)
	ownerOnlyAction := func(sequenceNum int) message.L2Message {
		return message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(int64(sequenceNum)),
			DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        arbos.StartArbOSUpgradeData(),
		})
	}

	ib.AddMessage(ownerOnlyAction(5), owner, big.NewInt(0), chainTime)
	ib.AddMessage(ownerOnlyAction(1), sender, big.NewInt(0), chainTime)
	ib.AddMessage(ownerOnlyAction(2), priorRandomOwner, big.NewInt(0), chainTime)
	shouldSucceed = append(shouldSucceed, false, true, false)

	results, _ := runTxAssertion(t, ib.Messages)

	if len(results) != len(shouldSucceed) {
		t.Log("length of results does not match checks", len(results), len(shouldSucceed))
	}

	for i := 0; i < len(results); i++ {
		t.Log("Checking txn", i, shouldSucceed[i])
		if shouldSucceed[i] {
			succeededTxCheck(t, results[i])
		} else {
			revertedTxCheck(t, results[i])
		}
	}
}

