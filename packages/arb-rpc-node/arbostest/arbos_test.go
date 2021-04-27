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

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestFib(t *testing.T) {
	fib, err := abi.JSON(strings.NewReader(arbostestcontracts.FibonacciABI))
	failIfError(t, err)

	constructorData, err := hexutil.Decode(arbostestcontracts.FibonacciBin)
	failIfError(t, err)

	constructTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        constructorData,
	}

	generateTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(300),
		Data:        generateFib(t, big.NewInt(20)),
	}

	getFibTx := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: connAddress1,
			Payment:     big.NewInt(0),
			Data:        makeFuncData(t, fib.Methods["getFib"], big.NewInt(5)),
		},
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(1000)),
		message.NewSafeL2Message(constructTx),
		message.NewSafeL2Message(generateTx),
		message.NewSafeL2Message(getFibTx),
	}

	logs, _, snap, _ := runSimpleAssertion(t, messages)
	results := processTxResults(t, logs)
	allResultsSucceeded(t, results)
	checkConstructorResult(t, results[1], connAddress1)

	generateResult := results[2]
	if len(generateResult.EVMLogs) != 1 {
		t.Fatal("incorrect log count")
	}
	evmLog := generateResult.EVMLogs[0]
	if evmLog.Address != connAddress1 {
		t.Fatal("log came from incorrect address")
	}
	if evmLog.Topics[0].ToEthHash() != fib.Events["TestEvent"].ID {
		t.Fatal("incorrect log topic")
	}
	if hexutil.Encode(evmLog.Data) != "0x0000000000000000000000000000000000000000000000000000000000000014" {
		t.Fatal("incorrect log data")
	}

	if hexutil.Encode(results[3].ReturnData) != "0x0000000000000000000000000000000000000000000000000000000000000008" {
		t.Fatal("getFib had incorrect result")
	}

	code, err := snap.GetCode(connAddress1)
	failIfError(t, err)
	t.Log("code", len(code))

}

func TestDeposit(t *testing.T) {
	amount := big.NewInt(1000)
	messages := []message.Message{
		makeEthDeposit(sender, amount),
	}

	_, _, snap, _ := runSimpleAssertion(t, messages)
	checkBalance(t, snap, sender, amount)
}

func TestBlocks(t *testing.T) {
	messages := make([]inbox.InboxMessage, 0)
	startTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(1),
		Timestamp: big.NewInt(1),
	}

	messages = append(
		messages,
		message.NewInboxMessage(initMsg(t, nil), chain, big.NewInt(0), big.NewInt(0), startTime),
	)

	messages = append(
		messages,
		message.NewInboxMessage(makeEthDeposit(sender, big.NewInt(1000)), chain, big.NewInt(0), big.NewInt(0), startTime),
	)

	halfSendCount := int64(5)

	blockTimes := make([]inbox.ChainTime, 0)
	blockTimes = append(blockTimes, inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(1),
		Timestamp: big.NewInt(1),
	})
	for i := int64(0); i < halfSendCount; i++ {
		time := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(2 + i),
			Timestamp: big.NewInt(11 + i),
		}
		blockTimes = append(blockTimes, time)
	}

	for i := int64(0); i < halfSendCount; i++ {
		tx := message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(i * 2),
			DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
			Payment:     big.NewInt(i * 2),
			Data:        arbos.WithdrawEthData(common.RandAddress()),
		}
		tx2 := message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(i*2 + 1),
			DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
			Payment:     big.NewInt(i*2 + 1),
			Data:        arbos.WithdrawEthData(common.RandAddress()),
		}
		messages = append(
			messages,
			message.NewInboxMessage(
				message.NewSafeL2Message(tx),
				sender,
				big.NewInt(i*2+2),
				big.NewInt(0),
				blockTimes[i+1],
			),
		)
		messages = append(
			messages,
			message.NewInboxMessage(
				message.NewSafeL2Message(tx2),
				sender,
				big.NewInt(i*2+2),
				big.NewInt(0),
				blockTimes[i+1],
			),
		)
	}

	type TargetBlockInfo struct {
		txCount       int
		includesBatch bool
	}

	type resType int
	const (
		txRes resType = iota
		sendRes
		merkleRes
		blockRes
	)

	targetBlocks := []TargetBlockInfo{
		{
			txCount:       1,
			includesBatch: false,
		},
		{
			txCount:       2,
			includesBatch: true,
		},
		{
			txCount:       2,
			includesBatch: false,
		},
		{
			txCount:       2,
			includesBatch: true,
		},
		{
			txCount:       2,
			includesBatch: false,
		},
		{
			txCount:       2,
			includesBatch: false,
		},
	}

	resultTypes := make([]resType, 0)
	sendCount := 0
	for i, targetBlock := range targetBlocks {
		for i := 0; i < targetBlock.txCount; i++ {
			resultTypes = append(resultTypes, txRes)
		}
		if i != len(targetBlocks)-1 {
			if i != 0 {
				for i := 0; i < targetBlock.txCount; i++ {
					resultTypes = append(resultTypes, sendRes)
				}
			}
			if targetBlock.includesBatch {
				resultTypes = append(resultTypes, merkleRes)
				sendCount++
			}
			resultTypes = append(resultTypes, blockRes)
		}
	}

	// Last value returned is not an error type
	avmLogs, sends, _, _ := runAssertion(t, messages, len(resultTypes), sendCount)
	results := make([]evm.Result, 0)
	for _, avmLog := range avmLogs {
		res, err := evm.NewResultFromValue(avmLog)
		failIfError(t, err)
		results = append(results, res)
	}

	for i, res := range results {
		switch res := res.(type) {
		case *evm.TxResult:
			t.Logf("%v %T, L1Block=%v L2Block=%v\n", i, res, res.IncomingRequest.L1BlockNumber, res.IncomingRequest.L2BlockNumber)
		case *evm.BlockInfo:
			t.Logf("%v %T, L2Block=%v\n", i, res, res.BlockNum)
		default:
			t.Logf("%v %T\n", i, res)
		}
	}

	arbSendsAccumulated := big.NewInt(0)

	blockGasUsed := big.NewInt(0)
	blockAVMLogCount := big.NewInt(0)
	blockAVMSendCount := big.NewInt(0)
	blockEVMLogCount := big.NewInt(0)
	blockTxCount := big.NewInt(0)

	totalGasUsed := big.NewInt(0)
	totalAVMLogCount := big.NewInt(0)
	totalAVMSendCount := big.NewInt(0)
	totalEVMLogCount := big.NewInt(0)
	totalTxCount := big.NewInt(0)
	blockCount := big.NewInt(0)
	prevBlockNum := abi.MaxUint256

	blocks := make([]*evm.BlockInfo, 0)
	merkleRoots := make([]*evm.MerkleRootResult, 0)

	for i, res := range results {
		totalAVMLogCount = totalAVMLogCount.Add(totalAVMLogCount, big.NewInt(1))

		switch resultTypes[i] {
		case blockRes:
			res, ok := res.(*evm.BlockInfo)
			if !ok {
				t.Fatal("incorrect result type", i)
			}
			blocks = append(blocks, res)

			correctTime := blockTimes[blockCount.Uint64()]
			if res.BlockNum.Cmp(blockCount) != 0 {
				t.Error("unexpected block height", res.BlockNum, i)
			}
			if res.Timestamp.Cmp(correctTime.Timestamp) != 0 {
				t.Error("unexpected timestamp", res.BlockNum, res.Timestamp, correctTime.Timestamp)
			}

			if res.BlockStats.GasUsed.Cmp(blockGasUsed) != 0 {
				t.Error("unexpected chain gas used")
			}
			if res.BlockStats.AVMLogCount.Cmp(blockAVMLogCount) != 0 {
				t.Error("unexpected block log count", res.BlockStats.AVMLogCount, "instead of", blockAVMLogCount)
			}
			if res.BlockStats.AVMSendCount.Cmp(blockAVMSendCount) != 0 {
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
			if res.ChainStats.AVMSendCount.Cmp(totalAVMSendCount) != 0 {
				t.Error("unexpected chain send count")
			}
			if res.ChainStats.EVMLogCount.Cmp(totalEVMLogCount) != 0 {
				t.Error("unexpected chain evm log count")
			}
			if res.ChainStats.TxCount.Cmp(totalTxCount) != 0 {
				t.Error("unexpected chain tx count", res.ChainStats.TxCount, "instead of", totalTxCount)
			}

			if res.LastAVMLog().Uint64() != uint64(i) {
				t.Error("incorrect last log")
			}

			if res.PreviousHeight.Cmp(prevBlockNum) != 0 {
				t.Error("incorrect prev block num")
			}

			blockGasUsed = big.NewInt(0)
			blockAVMLogCount = big.NewInt(0)
			blockAVMSendCount = big.NewInt(0)
			blockEVMLogCount = big.NewInt(0)
			blockTxCount = big.NewInt(0)
			prevBlockNum = res.BlockNum
			blockCount = blockCount.Add(blockCount, big.NewInt(1))
		case txRes:
			res, ok := res.(*evm.TxResult)
			if !ok {
				t.Fatal("incorrect result type")
			}
			succeededTxCheck(t, res)
			blockGasUsed = blockGasUsed.Add(blockGasUsed, res.GasUsed)
			blockEVMLogCount = blockEVMLogCount.Add(blockEVMLogCount, big.NewInt(int64(len(res.EVMLogs))))
			blockTxCount = blockTxCount.Add(blockTxCount, big.NewInt(1))
			blockAVMLogCount = blockAVMLogCount.Add(blockAVMLogCount, big.NewInt(1))

			totalGasUsed = totalGasUsed.Add(totalGasUsed, res.GasUsed)
			totalEVMLogCount = totalEVMLogCount.Add(totalEVMLogCount, big.NewInt(int64(len(res.EVMLogs))))
			totalTxCount = totalTxCount.Add(totalTxCount, big.NewInt(1))
		case sendRes:
			_, ok := res.(*evm.SendResult)
			if !ok {
				t.Fatal("incorrect result type")
			}
			blockAVMLogCount = blockAVMLogCount.Add(blockAVMLogCount, big.NewInt(1))
			arbSendsAccumulated = arbSendsAccumulated.Add(arbSendsAccumulated, big.NewInt(1))
		case merkleRes:
			root, ok := res.(*evm.MerkleRootResult)
			if !ok {
				t.Fatal("incorrect result type", i)
			}
			if root.NumInBatch.Cmp(arbSendsAccumulated) != 0 {
				t.Fatal("unexpected send count in merkle root")
			}
			merkleRoots = append(merkleRoots, root)
			arbSendsAccumulated = big.NewInt(0)
			blockAVMLogCount = blockAVMLogCount.Add(blockAVMLogCount, big.NewInt(1))
			blockAVMSendCount = blockAVMSendCount.Add(blockAVMSendCount, big.NewInt(1))
			totalAVMSendCount = totalAVMSendCount.Add(totalAVMSendCount, big.NewInt(1))
		default:
			t.Fatal("unknown result type")
		}
	}

	if len(sends) != len(merkleRoots) {
		t.Fatal("incorrect send or send log count")
	}
	for i, send := range sends {
		outMsg, err := message.NewOutMessageFromBytes(send)
		failIfError(t, err)

		sendMessageRoot, ok := outMsg.(*message.SendMessageRoot)
		if !ok {
			t.Fatal("send had wrong kind")
		}

		sendMessageLog := merkleRoots[i]
		if sendMessageRoot.BatchNumber.Cmp(big.NewInt(int64(i))) != 0 {
			t.Log("merkle send had wrong batch number")
		}
		if sendMessageLog.BatchNumber.Cmp(big.NewInt(int64(i))) != 0 {
			t.Log("merkle log had wrong batch num")
		}

		if sendMessageRoot.NumInBatch.Cmp(sendMessageLog.NumInBatch) != 0 {
			t.Error("merkle send and log have different num in batch")
		}

		treeHash := sendMessageLog.Tree.Hash()
		if treeHash != sendMessageRoot.OutputRoot {
			t.Error("incorrect send root", treeHash, sendMessageRoot.OutputRoot)
		}
	}

	for i, block := range blocks {
		target := targetBlocks[i]
		txCount := block.BlockStats.TxCount.Uint64()

		if uint64(target.txCount) != txCount {
			t.Fatal("wrong tx count in block, got", txCount, "but expected", target.txCount, "in block", i)
		}

		startLog := block.FirstAVMLog().Uint64()
		for i := uint64(0); i < txCount; i++ {
			txRes, ok := results[startLog+i].(*evm.TxResult)
			if !ok {
				continue
			}
			if txRes.IncomingRequest.L2BlockNumber.Cmp(block.BlockNum) != 0 {
				t.Error("tx in block had wrong block num")
			}
			if txRes.IncomingRequest.L2Timestamp.Cmp(block.Timestamp) != 0 {
				t.Error("tx in block had wrong timestamp")
			}
		}

		sendCount := block.BlockStats.AVMSendCount.Uint64()
		if target.includesBatch {
			if sendCount != 1 {
				t.Fatal("should have had 1 send")
			}
		} else {
			if sendCount != 0 {
				t.Fatal("should have had 0 sends")
			}
		}
	}
}
