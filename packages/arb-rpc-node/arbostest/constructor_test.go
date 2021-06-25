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
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var constructorData = hexutil.MustDecode(arbostestcontracts.FibonacciBin)

func TestConstructor(t *testing.T) {
	client, pks := test.SimulatedBackend(t)

	tx := types.NewContractCreation(0, big.NewInt(0), 5000000, big.NewInt(0), constructorData)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, pks[0])
	failIfError(t, err)

	targetAddress := crypto.CreateAddress(crypto.PubkeyToAddress(pks[0].PublicKey), 0)

	ctx := context.Background()
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		t.Fatal(err)
	}
	client.Commit()
	ethReceipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
	failIfError(t, err)

	if ethReceipt.ContractAddress != targetAddress {
		t.Error("ethereum contract address incorrect")
	}

	l2Message, err := message.NewL2Message(message.NewCompressedECDSAFromEth(signedTx))
	failIfError(t, err)

	messages := []message.Message{l2Message}
	results, snap := runSimpleTxAssertion(t, messages)

	res := results[0]

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

	arbAddress := common.NewAddressFromEth(targetAddress)
	checkConstructorResult(t, res, arbAddress)

	ethCode, err := client.CodeAt(ctx, ethReceipt.ContractAddress, nil)
	failIfError(t, err)

	arbCode, err := snap.GetCode(arbAddress)
	failIfError(t, err)

	if !bytes.Equal(arbCode, ethCode) {
		t.Error("deployed code is different")
	}
}

func TestConstructorExistingBalance(t *testing.T) {
	factoryABI, err := abi.JSON(strings.NewReader(arbostestcontracts.CloneFactoryABI))
	failIfError(t, err)

	create2Address := common.HexToAddress("0xa0a7964936862853f101d4da3a338fe56d5e0482")

	tx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress2,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, factoryABI.Methods["create2Clone"], connAddress1, big.NewInt(0)),
	}

	messages := []message.Message{
		makeEthDeposit(connAddress1, big.NewInt(100)),
		makeEthDeposit(create2Address, big.NewInt(100)),
		message.NewSafeL2Message(makeSimpleConstructorTx(constructorData, big.NewInt(0))),
		message.NewSafeL2Message(makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.CloneFactoryBin), big.NewInt(1))),
		message.NewSafeL2Message(tx),
	}

	results, _ := runSimpleTxAssertion(t, messages)

	checkConstructorResult(t, results[2], connAddress1)
	checkConstructorResult(t, results[3], connAddress2)
	succeededTxCheck(t, results[4])
	if !bytes.Equal(results[4].ReturnData[12:], create2Address.Bytes()) {
		t.Fatal("incorrect create2 address which should have been", hexutil.Encode(results[4].ReturnData[12:]))
	}
}

func TestConstructorCallback(t *testing.T) {
	skipBelowVersion(t, 18)

	client, pks := test.SimulatedBackend(t)
	auth, err := bind.NewKeyedTransactorWithChainID(pks[0], big.NewInt(1337))
	test.FailIfError(t, err)
	_, _, con, err := arbostestcontracts.DeployConstructorCallback2(auth, client)
	test.FailIfError(t, err)
	client.Commit()
	ethTx, err := con.Test(auth)
	test.FailIfError(t, err)
	client.Commit()
	receipt, err := client.TransactionReceipt(context.Background(), ethTx.Hash())
	test.FailIfError(t, err)

	conABI, err := abi.JSON(strings.NewReader(arbostestcontracts.ConstructorCallback2ABI))
	failIfError(t, err)

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.ConstructorCallback2Bin),
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(100),
		Data:        makeFuncData(t, conABI.Methods["test"]),
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(10000)),
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	}

	results, snap := runSimpleTxAssertion(t, messages)
	allResultsSucceeded(t, results)
	checkConstructorResult(t, results[1], connAddress1)
	res := results[2]
	if len(res.EVMLogs) != len(receipt.Logs) {
		t.Fatal("unexpected log count")
	}
	for i, evmLog := range receipt.Logs {
		arbosLog := res.EVMLogs[i]
		t.Log(evmLog.Address, evmLog.Topics, hexutil.Encode(evmLog.Data))
		t.Log(arbosLog)
		if len(evmLog.Topics) != len(arbosLog.Topics) {
			t.Error("unexpected topic count")
		}
		for j, evmTopic := range evmLog.Topics {
			if evmTopic != arbosLog.Topics[j].ToEthHash() {
				t.Error("wrong topic")
			}
		}
		if !bytes.Equal(evmLog.Data, arbosLog.Data) {
			t.Error("wrong data")
		}
	}

	bal, err := snap.GetBalance(connAddress1)
	test.FailIfError(t, err)
	if bal.Cmp(tx2.Payment) != 0 {
		t.Error("wrong balance")
	}
}
