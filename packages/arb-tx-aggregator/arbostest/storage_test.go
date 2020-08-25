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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestGetStorageAt(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}
	chain := common.RandAddress()
	sender := common.RandAddress()

	runMessage(t, mach, initMsg(), chain)

	constructorData := hexutil.MustDecode(arbostestcontracts.StorageBin)
	storageAddress, err := deployContract(t, mach, sender, constructorData, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	tx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: storageAddress,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x188f9139"),
	}
	res, err := runTransaction(t, mach, tx, sender)
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode != evm.RevertCode {
		t.Fatal("tx should have reverted")
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(1))
	slot0, err := snap.GetStorageAt(storageAddress, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	if slot0.Cmp(big.NewInt(0)) != 0 {
		t.Error("expected 0 slot to be 0 but got", slot0)
	}
	slot1, err := snap.GetStorageAt(storageAddress, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	if slot1.Cmp(big.NewInt(12345)) != 0 {
		t.Error("expected 1 slot to be 12345 but got", slot1)
	}
}
