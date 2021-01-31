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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestGetStorageAt(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	constructorTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.StorageBin), big.NewInt(0))

	getStorageAtTx := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        snapshot.StorageAtData(connAddress1, big.NewInt(1)),
		},
	}

	failGetStorageAtTx := message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress1,
			Payment:     big.NewInt(0),
			Data:        hexutil.MustDecode("0x188f9139"),
		},
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx), sender, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(getStorageAtTx), common.Address{}, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(failGetStorageAtTx), sender, big.NewInt(3), chainTime),
	}

	logs, _, _, _ := runAssertion(t, inboxMessages, 3, 0)
	results := processTxResults(t, logs)

	constructorRes := results[0]
	getStorageAtRes := results[1]
	failGetStorageAtRes := results[2]

	checkConstructorResult(t, constructorRes, connAddress1)
	succeededTxCheck(t, getStorageAtRes)
	storageVal := new(big.Int).SetBytes(getStorageAtRes.ReturnData)
	if storageVal.Cmp(big.NewInt(12345)) != 0 {
		t.Fatal("expected storage to be 12345 but got", storageVal)
	}

	revertedTxCheck(t, failGetStorageAtRes)
}
