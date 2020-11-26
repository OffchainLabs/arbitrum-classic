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
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"math/big"
	"testing"
)

var constructorData = hexutil.MustDecode(arbostestcontracts.FibonacciBin)

func TestContructor(t *testing.T) {
	client, pks := test.SimulatedBackend()

	tx := types.NewContractCreation(0, big.NewInt(0), 1000000, big.NewInt(0), constructorData)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, pks[0])
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		t.Fatal(err)
	}
	client.Commit()
	ethReceipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	l2Message, err := message.NewL2Message(message.NewCompressedECDSAFromEth(signedTx))
	if err != nil {
		t.Fatal(err)
	}

	chain := common.RandAddress()
	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime))
	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		l2Message,
		common.RandAddress(),
		big.NewInt(1),
		chainTime,
	))

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(1000000000, inboxMessages, 0)
	logs := assertion.ParseLogs()

	if len(logs) != 1 {
		t.Fatal("unexpected log count", len(logs))
	}

	res, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == evm.ReturnCode {
		if ethReceipt.Status != 1 {
			t.Fatal("arb tx succeeded but eth tx failed")
		}
		t.Log("constructors succeeded")
	} else {
		if ethReceipt.Status != 0 {
			t.Fatal("arb tx failed but eth tx succeeded")
		}
		t.Log("constructors failed")
	}

	if res.ResultCode != evm.ReturnCode {
		// Nothing else to check if the tx failed
		return
	}

	arbAddress, err := getConstructorResult(res)
	if err != nil {
		t.Fatal(err)
	}
	if arbAddress.ToEthAddress() != ethReceipt.ContractAddress {
		t.Error("contracts deployed at different addresses")
	}

	ethCode, err := client.CodeAt(ctx, ethReceipt.ContractAddress, nil)
	if err != nil {
		t.Fatal(err)
	}

	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(9999999))
	arbCode, err := snap.GetCode(arbAddress)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(arbCode, ethCode) {
		t.Error("deployed code is different")
	}
}
