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
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestFib(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

	fib, err := abi.JSON(strings.NewReader(arbostestcontracts.FibonacciABI))
	failIfError(t, err)

	pk, err := crypto.GenerateKey()
	failIfError(t, err)

	addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))

	runMessage(t, mach, initMsg(), chain)

	constructorData, err := hexutil.Decode(arbostestcontracts.FibonacciBin)
	failIfError(t, err)

	fibAddress, err := deployContract(t, mach, addr, constructorData, big.NewInt(0), nil)
	failIfError(t, err)

	snap := snapshot.NewSnapshot(mach.Clone(), chainTime, message.ChainAddressToID(chain), big.NewInt(1))
	code, err := snap.GetCode(fibAddress)
	failIfError(t, err)
	t.Log("code", len(code))

	depositEth(t, mach, addr, big.NewInt(1000))

	fibData, err := generateFib(big.NewInt(20))
	failIfError(t, err)

	generateTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: fibAddress,
		Payment:     big.NewInt(300),
		Data:        fibData,
	}

	generateResult, err := runValidTransaction(t, mach, generateTx, addr)
	failIfError(t, err)
	if len(generateResult.EVMLogs) != 1 {
		t.Fatal("incorrect log count")
	}
	evmLog := generateResult.EVMLogs[0]
	if evmLog.Address != fibAddress {
		t.Fatal("log came from incorrect address")
	}
	if evmLog.Topics[0].ToEthHash() != fib.Events["TestEvent"].ID {
		t.Fatal("incorrect log topic")
	}
	if hexutil.Encode(evmLog.Data) != "0x0000000000000000000000000000000000000000000000000000000000000014" {
		t.Fatal("incorrect log data")
	}

	getFibABI := fib.Methods["getFib"]
	getFibData, err := getFibABI.Inputs.Pack(big.NewInt(5))
	failIfError(t, err)
	getFibTx := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: fibAddress,
			Payment:     big.NewInt(0),
			Data:        append(getFibABI.ID, getFibData...),
		},
	}

	getFibResult, err := runValidTransaction(t, mach, getFibTx, addr)
	failIfError(t, err)
	if hexutil.Encode(getFibResult.ReturnData) != "0x0000000000000000000000000000000000000000000000000000000000000008" {
		t.Fatal("getFib had incorrect result")
	}
}
func TestDeposit(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)

	pk, err := crypto.GenerateKey()
	failIfError(t, err)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	runMessage(t, mach, initMsg(), chain)

	addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))

	amount := big.NewInt(1000)
	depositEth(t, mach, addr, amount)

	snap := snapshot.NewSnapshot(mach.Clone(), chainTime, message.ChainAddressToID(chain), big.NewInt(1))
	balance, err := snap.GetBalance(addr)
	failIfError(t, err)
	if balance.Cmp(amount) != 0 {
		t.Fatal("incorrect balance")
	}
}

func TestBlocks(t *testing.T) {
	messages := make([]inbox.InboxMessage, 0)
	messages = append(
		messages,
		message.NewInboxMessage(
			initMsg(),
			chain,
			big.NewInt(0),
			inbox.ChainTime{
				BlockNum:  common.NewTimeBlocksInt(0),
				Timestamp: big.NewInt(0),
			},
		),
	)
	for i := int64(0); i < 5; i++ {
		tx := message.Transaction{
			MaxGas:      big.NewInt(100000000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(0),
			DestAddress: common.RandAddress(),
			Payment:     big.NewInt(0),
			Data:        []byte{},
		}
		messages = append(
			messages,
			message.NewInboxMessage(
				message.NewSafeL2Message(tx),
				common.RandAddress(),
				big.NewInt(i+1),
				inbox.ChainTime{
					BlockNum:  common.NewTimeBlocksInt(i + 1),
					Timestamp: big.NewInt(10 + i + 1),
				},
			),
		)
	}

	// Last value returned is not an error type
	avmLogs, _, _ := runAssertion(t, messages, 9, 0)
	t.Log("Got", len(avmLogs), "logs")
	blockGasUsed := big.NewInt(0)
	blockAVMLogCount := big.NewInt(0)
	blockEVMLogCount := big.NewInt(0)
	blockTxCount := big.NewInt(0)

	totalGasUsed := big.NewInt(0)
	totalAVMLogCount := big.NewInt(0)
	totalEVMLogCount := big.NewInt(0)
	totalTxCount := big.NewInt(0)
	for i, avmLog := range avmLogs {
		totalAVMLogCount = totalAVMLogCount.Add(totalAVMLogCount, big.NewInt(1))
		res, err := evm.NewResultFromValue(avmLog)
		failIfError(t, err)

		if i%2 == 0 {
			res, ok := res.(*evm.TxResult)
			if !ok {
				t.Error("incorrect result type")
			}
			succeededTxCheck(t, res)
			blockGasUsed = blockGasUsed.Add(blockGasUsed, res.GasUsed)
			blockEVMLogCount = blockEVMLogCount.Add(blockEVMLogCount, big.NewInt(int64(len(res.EVMLogs))))
			blockTxCount = blockTxCount.Add(blockTxCount, big.NewInt(1))
			blockAVMLogCount = blockAVMLogCount.Add(blockAVMLogCount, big.NewInt(1))

			totalGasUsed = totalGasUsed.Add(totalGasUsed, res.GasUsed)
			totalEVMLogCount = totalEVMLogCount.Add(totalEVMLogCount, big.NewInt(int64(len(res.EVMLogs))))
			totalTxCount = totalTxCount.Add(totalTxCount, big.NewInt(1))
		} else {
			res, ok := res.(*evm.BlockInfo)
			if !ok {
				t.Fatal("incorrect result type")
			}
			if res.BlockNum.Cmp(big.NewInt(int64(i/2+1))) != 0 {
				t.Error("unexpected block height")
			}
			if res.Timestamp.Cmp(big.NewInt(int64(10+i/2+1))) != 0 {
				t.Error("unexpected block height")
			}

			if res.BlockStats.GasUsed.Cmp(blockGasUsed) != 0 {
				t.Error("unexpected chain gas used")
			}
			if res.BlockStats.AVMLogCount.Cmp(blockAVMLogCount) != 0 {
				t.Error("unexpected block log count", res.BlockStats.AVMLogCount, "instead of", blockAVMLogCount)
			}
			if res.BlockStats.AVMSendCount.Cmp(big.NewInt(0)) != 0 {
				t.Error("unexpected block send count")
			}
			if res.BlockStats.EVMLogCount.Cmp(blockEVMLogCount) != 0 {
				t.Error("unexpected block evm log count")
			}
			if res.BlockStats.TxCount.Cmp(blockTxCount) != 0 {
				t.Error("unexpected block tx count", res.BlockStats.TxCount)
			}

			if res.ChainStats.GasUsed.Cmp(totalGasUsed) != 0 {
				t.Error("unexpected chain gas used")
			}
			if res.ChainStats.AVMLogCount.Cmp(totalAVMLogCount) != 0 {
				t.Error("unexpected chain log count", res.ChainStats.AVMLogCount, "instead of", totalAVMLogCount)
			}
			if res.ChainStats.AVMSendCount.Cmp(big.NewInt(0)) != 0 {
				t.Error("unexpected chain send count")
			}
			if res.ChainStats.EVMLogCount.Cmp(totalEVMLogCount) != 0 {
				t.Error("unexpected chain evm log count")
			}
			if res.ChainStats.TxCount.Cmp(totalTxCount) != 0 {
				t.Error("unexpected chain tx count", res.ChainStats.TxCount, "instead of", totalTxCount)
			}

			blockGasUsed = big.NewInt(0)
			blockAVMLogCount = big.NewInt(0)
			blockEVMLogCount = big.NewInt(0)
			blockTxCount = big.NewInt(0)
		}
	}
}
