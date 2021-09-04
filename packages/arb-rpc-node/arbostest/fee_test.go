/*
* Copyright 2021, Offchain Labs, Inc.
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
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func addEnableFeesMessages(ib *InboxBuilder) {
	ownerTx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.SetFairGasPriceSender(owner, true),
	}

	ownerTx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.SetFeesEnabled(true),
	}

	ownerMessages := []message.Transaction{ownerTx1, ownerTx2}
	for _, msg := range ownerMessages {
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(int64(len(ib.Messages))),
			Timestamp: big.NewInt(0),
		}
		ib.AddMessage(message.NewSafeL2Message(msg), owner, big.NewInt(0), chainTime)
	}
}

type txTemplate struct {
	GasPrice *big.Int
	Gas      uint64
	To       *ethcommon.Address
	Value    *big.Int
	Data     []byte

	// Data to verify tx
	resultType         []evm.ResultType
	nonzeroComputation []bool
	correctStorageUsed int
	calldata           int
	ranOutOfFunds      bool
}

func TestArbOSFees(t *testing.T) {
	skipBelowVersion(t, 35)

	countCalldataUnits := func(data []byte) int {
		units := 0
		for _, val := range data {
			if val == 0 {
				units += 4
			} else {
				units += 16
			}
		}
		return units
	}

	privKey, err := crypto.GenerateKey()
	failIfError(t, err)
	signer := types.NewEIP155Signer(chainId)
	userAddress := common.NewAddressFromEth(crypto.PubkeyToAddress(privKey.PublicKey))

	initialDeposit := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)
	initialDeposit = initialDeposit.Mul(initialDeposit, big.NewInt(4))

	gasUsedABI, err := abi.JSON(strings.NewReader(arbostestcontracts.GasUsedABI))
	failIfError(t, err)

	contractDest := crypto.CreateAddress(userAddress.ToEthAddress(), 0)
	eoaDest := common.RandAddress().ToEthAddress()

	var conDeployedLength int
	{
		client, auths := test.SimulatedBackend(t)
		addr, _, _, err := arbostestcontracts.DeployGasUsed(auths[0], client, false)
		failIfError(t, err)
		client.Commit()
		deployedCode, err := client.CodeAt(context.Background(), addr, nil)
		failIfError(t, err)
		conDeployedLength = len(deployedCode)
	}

	conDataSuccess := hexutil.MustDecode(arbostestcontracts.GasUsedBin)
	conDataSuccess = append(conDataSuccess, math.U256Bytes(big.NewInt(0))...)

	conDataFailure := hexutil.MustDecode(arbostestcontracts.GasUsedBin)
	conDataFailure = append(conDataFailure, math.U256Bytes(big.NewInt(1))...)

	rawTxes := []txTemplate{
		// Successful call to constructor
		{
			GasPrice: big.NewInt(0),
			Gas:      300000000,
			Value:    big.NewInt(0),
			Data:     conDataSuccess,

			resultType:         []evm.ResultType{evm.ReturnCode},
			nonzeroComputation: []bool{true},
			correctStorageUsed: (conDeployedLength + 31) / 32,
		},
		// Successful call to method without storage
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     makeFuncData(t, gasUsedABI.Methods["noop"]),

			resultType:         []evm.ResultType{evm.ReturnCode},
			nonzeroComputation: []bool{true},
			correctStorageUsed: 0,
		},
		// Successful call to storage allocating method
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     makeFuncData(t, gasUsedABI.Methods["sstore"]),

			resultType:         []evm.ResultType{evm.ReturnCode},
			nonzeroComputation: []bool{true},
			correctStorageUsed: 1,
		},
		// Successful eth transfer to EOA
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &eoaDest,
			Value:    big.NewInt(100000),

			resultType:         []evm.ResultType{evm.ReturnCode},
			nonzeroComputation: []bool{false},
			correctStorageUsed: 0,
		},
		// Reverted constructor
		{
			GasPrice: big.NewInt(0),
			Gas:      1000000000,
			Value:    big.NewInt(0),
			Data:     conDataFailure,

			resultType:         []evm.ResultType{evm.RevertCode},
			nonzeroComputation: []bool{true},
			correctStorageUsed: 0,
		},
		// Reverted storage allocating function call
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     makeFuncData(t, gasUsedABI.Methods["fail"]),

			resultType:         []evm.ResultType{evm.RevertCode},
			nonzeroComputation: []bool{true},
			correctStorageUsed: 0,
		},
		// Reverted since insufficient funds
		{
			GasPrice: big.NewInt(0),
			Gas:      1000000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     common.RandBytes(100000),

			resultType:         []evm.ResultType{evm.RevertCode, evm.InsufficientGasFundsCode, evm.InsufficientGasFundsCode, evm.InsufficientGasFundsCode},
			nonzeroComputation: []bool{true, false, false, false},
			correctStorageUsed: 0,
		},
	}
	valueTransfered := big.NewInt(0)
	for _, tx := range rawTxes {
		valueTransfered = valueTransfered.Add(valueTransfered, tx.Value)
	}

	aggregator := common.RandAddress()
	netFeeRecipient := common.RandAddress()
	congestionFeeRecipient := common.RandAddress()

	t.Log("User", userAddress)
	t.Log("Net fee recipient", netFeeRecipient)
	t.Log("Congestion recipient", congestionFeeRecipient)

	addInitializationLoc := func(ib *InboxBuilder) {
		config := protocol.ChainParams{
			GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
			ArbGasSpeedLimitPerSecond: 1000000000,
		}

		feeConfigInit := message.FeeConfig{
			SpeedLimitPerSecond:    new(big.Int).SetUint64(config.ArbGasSpeedLimitPerSecond),
			L1GasPerL2Tx:           big.NewInt(3700),
			ArbGasPerL2Tx:          big.NewInt(0),
			L1GasPerL2Calldata:     big.NewInt(1),
			ArbGasPerL2Calldata:    big.NewInt(0),
			L1GasPerStorage:        big.NewInt(2000),
			ArbGasPerStorage:       big.NewInt(0),
			ArbGasDivisor:          big.NewInt(10000),
			NetFeeRecipient:        netFeeRecipient,
			CongestionFeeRecipient: congestionFeeRecipient,
		}
		aggInit := message.DefaultAggConfig{Aggregator: aggregator}
		chainIdInit := message.ChainIDConfig{ChainId: chainId}
		init, err := message.NewInitMessage(config, owner, []message.ChainConfigOption{chainIdInit, feeConfigInit, aggInit})
		test.FailIfError(t, err)

		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		}
		ib.AddMessage(init, common.Address{}, big.NewInt(0), chainTime)

		deposit := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(10000000),
					GasPriceBid: big.NewInt(0),
					DestAddress: userAddress,
					Payment:     initialDeposit,
					Data:        nil,
				},
			}),
		}
		ib.AddMessage(deposit, chain, big.NewInt(0), chainTime)
	}

	ethTxes := make([]*types.Transaction, 0)
	for i, rawTx := range rawTxes {
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    uint64(i),
			GasPrice: new(big.Int).Set(rawTx.GasPrice),
			Gas:      rawTx.Gas,
			To:       rawTx.To,
			Value:    new(big.Int).Set(rawTx.Value),
			Data:     append([]byte{}, rawTx.Data...),
		})
		ethTxes = append(ethTxes, tx)
	}

	buildCompressedTxes := func() []message.CompressedECDSATransaction {
		compressedTxes := make([]message.CompressedECDSATransaction, 0)
		for _, tx := range ethTxes {
			signedTx, err := types.SignTx(tx, signer, privKey)
			failIfError(t, err)
			compressedTxes = append(compressedTxes, message.NewCompressedECDSAFromEth(signedTx))
		}
		return compressedTxes
	}
	for i, compressedTx := range buildCompressedTxes() {
		l2, err := message.NewL2Message(compressedTx)
		failIfError(t, err)
		rawTxes[i].calldata = countCalldataUnits(l2.Data)
	}

	addUserTxesLoc := func(ib *InboxBuilder, agg common.Address) {
		t.Helper()
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(int64(len(ib.Messages))),
			Timestamp: big.NewInt(0),
		}
		for _, tx := range buildCompressedTxes() {
			batch, err := message.NewTransactionBatchFromMessages([]message.AbstractL2Message{tx})
			failIfError(t, err)
			ib.AddMessage(message.NewSafeL2Message(batch), agg, big.NewInt(0), chainTime)
			chainTime.BlockNum = common.NewTimeBlocksInt(int64(len(ib.Messages)))
		}
	}

	noFeeIB := &InboxBuilder{}
	addInitializationLoc(noFeeIB)
	addUserTxesLoc(noFeeIB, aggregator)

	otherAgg := common.RandAddress()
	feeIB := &InboxBuilder{}
	addInitializationLoc(feeIB)
	addEnableFeesMessages(feeIB)
	addUserTxesLoc(feeIB, otherAgg)

	feeWithAggIB := &InboxBuilder{}
	addInitializationLoc(feeWithAggIB)
	addEnableFeesMessages(feeWithAggIB)
	addUserTxesLoc(feeWithAggIB, aggregator)

	feeWithContractTxIB := &InboxBuilder{}
	addInitializationLoc(feeWithContractTxIB)
	addEnableFeesMessages(feeWithContractTxIB)
	contractTxData := make([]int, 0, len(ethTxes))
	{
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(int64(len(feeWithContractTxIB.Messages))),
			Timestamp: big.NewInt(0),
		}
		for _, tx := range ethTxes {
			var dest common.Address
			if tx.To() != nil {
				dest = common.NewAddressFromEth(*tx.To())
			}
			l2msg := message.ContractTransaction{BasicTx: message.BasicTx{
				MaxGas:      new(big.Int).SetUint64(tx.Gas()),
				GasPriceBid: tx.GasPrice(),
				DestAddress: dest,
				Payment:     tx.Value(),
				Data:        tx.Data(),
			}}
			msg := message.NewSafeL2Message(l2msg)
			feeWithContractTxIB.AddMessage(msg, userAddress, big.NewInt(0), chainTime)
			chainTime.BlockNum = common.NewTimeBlocksInt(int64(len(feeWithContractTxIB.Messages)))
			contractTxData = append(contractTxData, countCalldataUnits(msg.Data))
		}
	}

	estimateFeeIB := &InboxBuilder{}
	addInitializationLoc(estimateFeeIB)
	addEnableFeesMessages(estimateFeeIB)
	{
		chainTime := inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(int64(len(estimateFeeIB.Messages))),
			Timestamp: big.NewInt(0),
		}
		for _, tx := range ethTxes {
			compressed := message.NewCompressedECDSAFromEth(tx)
			compressed.GasLimit = big.NewInt(0)
			compressed.GasPrice = big.NewInt(0)
			msg, err := message.NewGasEstimationMessage(aggregator, big.NewInt(100000000), compressed)
			test.FailIfError(t, err)
			estimateFeeIB.AddMessage(msg, userAddress, big.NewInt(0), chainTime)
			chainTime.BlockNum = common.NewTimeBlocksInt(int64(len(estimateFeeIB.Messages)))
		}
	}

	processMessages := func(ib *InboxBuilder, index int, aggregator *common.Address) ([]*evm.TxResult, *snapshot.Snapshot, *big.Int) {
		t.Helper()
		results, _, snap := runAssertion(t, ib.Messages, math.MaxInt32, 0)
		rawResults := extractTxResults(t, results)
		allResultsSucceeded(t, rawResults[:len(rawResults)-len(rawTxes)])
		extractedResults := rawResults[len(rawResults)-len(rawTxes):]
		amountUnpaid := big.NewInt(0)
		for i, res := range extractedResults {
			resType := rawTxes[i].resultType[0]
			if len(rawTxes[i].resultType) > 1 {
				resType = rawTxes[i].resultType[index]
			}
			if res.ResultCode != resType {
				t.Fatal("unexpected result got", res.ResultCode, "but expected", resType, "for", index, i)
			}
			if res.ResultCode != evm.InsufficientGasFundsCode {
				// we don't include intentionally unpaid fees as in test 6
				unpaid := checkGas(t, res, aggregator, index == 3)
				amountUnpaid = amountUnpaid.Add(amountUnpaid, unpaid)
			}
		}
		return extractedResults, snap, amountUnpaid
	}

	checkAllUnits := func(results []*evm.TxResult, index int, calldataExact bool) {
		t.Helper()
		for i, res := range results {
			checkUnits(t, res, rawTxes[i], index, rawTxes[i].calldata, calldataExact)
		}
	}

	t.Log("Checking results for no fee")
	noFeeResults, noFeeSnap, unpaidNoFee := processMessages(noFeeIB, 0, &aggregator)
	t.Log("Checking results for fee")
	feeResults, feeSnap, unpaidFee := processMessages(feeIB, 1, &otherAgg)
	t.Log("Checking results for fee with agg")
	feeWithAggResults, feeWithAggSnap, _ := processMessages(feeWithAggIB, 2, &aggregator)
	t.Log("Checking results for fee with contract")
	feeWithContractResults, feeWithContractSnap, unpaidFeeWithContract := processMessages(feeWithContractTxIB, 1, nil)
	t.Log("Checking results for estimate")
	estimateFeeResults, estimateFeeSnap, _ := processMessages(estimateFeeIB, 3, &aggregator)

	if unpaidNoFee.Cmp(big.NewInt(0)) != 0 {
		t.Error("shouldn't have unpaid")
	}
	if unpaidFee.Cmp(big.NewInt(0)) != 0 {
		t.Error("shouldn't have unpaid")
	}
	if unpaidFeeWithContract.Cmp(big.NewInt(0)) != 0 {
		t.Error("shouldn't have unpaid")
	}

	checkAllUnits(noFeeResults, 0, true)
	checkAllUnits(feeResults, 1, true)
	checkAllUnits(feeWithAggResults, 2, true)
	for i, res := range feeWithContractResults {
		checkUnits(t, res, rawTxes[i], 1, contractTxData[i], true)
	}
	checkAllUnits(estimateFeeResults, 3, false)

	checkSameL2ComputationUnits(t, noFeeResults, feeResults)
	checkSameL2ComputationUnits(t, noFeeResults, feeWithAggResults)
	checkSameL2ComputationUnits(t, noFeeResults, feeWithContractResults)
	checkSameL2ComputationUnits(t, noFeeResults, estimateFeeResults)

	calcDiff := func(a, b *big.Rat) *big.Rat {
		diff := new(big.Rat).Sub(a, b)
		return diff.Abs(diff).Quo(diff, b)
	}
	calcDiffSigned := func(a, b *big.Rat) *big.Rat {
		diff := new(big.Rat).Sub(a, b)
		return diff.Quo(diff, b)
	}

	estimateFeeWithAgg := func(withoutAgg *big.Int) *big.Rat {
		calcAggTxPrice := new(big.Rat).SetInt(withoutAgg)
		return calcAggTxPrice.Mul(calcAggTxPrice, big.NewRat(115, 15))
	}

	calculateFeeAggDiff := func(withoutAgg, withAgg *big.Int) *big.Rat {
		calcAggTxPrice := estimateFeeWithAgg(withoutAgg)
		correctAggTxPrice := new(big.Rat).SetInt(withAgg)
		return calcDiff(calcAggTxPrice, correctAggTxPrice)
	}

	for i, res := range feeResults {
		noAggPrice := res.FeeStats.Price
		aggPrice := feeWithAggResults[i].FeeStats.Price

		l1TxDiff := calculateFeeAggDiff(noAggPrice.L1Transaction, aggPrice.L1Transaction)
		if l1TxDiff.Cmp(big.NewRat(1, 100)) > 0 {
			t.Error("tx price with agg is wrong ", i, " ", l1TxDiff)
		}

		l1CalldataDiff := calculateFeeAggDiff(noAggPrice.L1Calldata, aggPrice.L1Calldata)
		if l1CalldataDiff.Cmp(big.NewRat(1, 100)) > 0 {
			t.Error("tx price with agg is wrong ", i, l1CalldataDiff)
		}

		if noAggPrice.L2Computation.Cmp(aggPrice.L2Computation) != 0 {
			t.Error("wrong l2 computation price")
		}

		if noAggPrice.L2Storage.Cmp(aggPrice.L2Storage) != 0 {
			t.Error("wrong l2 storage price")
		}
	}

	t.Log("Checking estimate")
	for i, res := range feeWithAggResults {
		aggPrice := res.FeeStats.Price
		estimatePrice := estimateFeeResults[i].FeeStats.Price
		checkL1UnitsEqual(t, aggPrice, estimatePrice)
		checkL2UnitsEqual(t, aggPrice, estimatePrice)
	}
	t.Log("Finished checking estimate")

	checkTxCount := func(snap *snapshot.Snapshot, results []*evm.TxResult) {
		t.Helper()

		correctCount := int64(0)
		for _, res := range results {
			if arbosVersion < 16 {
				if res.ResultCode == evm.ReturnCode || res.ResultCode == evm.RevertCode {
					correctCount++
				}
			} else {
				if res.ResultCode != evm.BadSequenceCode {
					correctCount++
				}
			}
		}

		txCount, err := snap.GetTransactionCount(userAddress)
		test.FailIfError(t, err)

		if txCount.Cmp(big.NewInt(correctCount)) != 0 {
			t.Error("wrong tx count", txCount)
		}
	}

	checkPaid := func(snap *snapshot.Snapshot, results []*evm.TxResult) *big.Int {
		t.Helper()

		userBal, err := snap.GetBalance(userAddress)
		test.FailIfError(t, err)

		aggBal, err := snap.GetBalance(aggregator)
		test.FailIfError(t, err)

		t.Log("netFeeRecipient", netFeeRecipient)

		netFeeRecipientBal, err := snap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)

		reportedPaid := big.NewInt(0)
		for _, res := range results {
			reportedPaid = reportedPaid.Add(reportedPaid, res.FeeStats.Paid.Total())
		}
		amountPaid := new(big.Int).Sub(initialDeposit, userBal)
		amountPaid = amountPaid.Sub(amountPaid, valueTransfered)
		if amountPaid.Cmp(reportedPaid) != 0 {
			t.Error("wrong amount deducted from user got", amountPaid, "but expected", reportedPaid)
		}
		t.Log("Total paid", amountPaid)
		t.Log("Remaining balance", userBal)

		amountReceived := new(big.Int).Add(aggBal, netFeeRecipientBal)
		if amountReceived.Cmp(amountPaid) != 0 {
			t.Error("payment was not equal to amount received", amountReceived, amountPaid)
		}
		return amountPaid
	}

	checkNoCongestionFee := func(snap *snapshot.Snapshot) {
		t.Helper()
		congestionFeeRecipientBal, err := snap.GetBalance(congestionFeeRecipient)
		test.FailIfError(t, err)
		if congestionFeeRecipientBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("wrong congestion fee balance got", congestionFeeRecipientBal, "but expected 0")
		}
	}

	checkNoNonPreferredAggFee := func(snap *snapshot.Snapshot) {
		t.Helper()
		otherAggBal, err := snap.GetBalance(otherAgg)
		test.FailIfError(t, err)
		if otherAggBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("wrong other agg balance", otherAggBal, "but expected 0")
		}
	}

	checkNoAggFee := func(snap *snapshot.Snapshot) {
		t.Helper()
		aggBal, err := snap.GetBalance(aggregator)
		test.FailIfError(t, err)
		if aggBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("wrong other agg balance", aggBal, "but expected 0")
		}
	}

	checkTotalReceived := func(snap *snapshot.Snapshot, results []*evm.TxResult) (*big.Int, *big.Int, *big.Int) {
		t.Helper()
		aggBal, err := snap.GetBalance(aggregator)
		test.FailIfError(t, err)

		netFeeRecipientBal, err := snap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)

		totalPaidL1Tx := big.NewInt(0)
		totalPaidL1Calldata := big.NewInt(0)
		totalPaidL2Computation := big.NewInt(0)
		totalPaidL2Storage := big.NewInt(0)
		totalPaidUnderfunded := big.NewInt(0)
		for _, res := range results {
			if res.ResultCode == evm.InsufficientGasFundsCode {
				totalPaidUnderfunded.Add(totalPaidUnderfunded, res.FeeStats.Paid.L1Transaction)
				totalPaidUnderfunded.Add(totalPaidUnderfunded, res.FeeStats.Paid.L1Calldata)
			} else {
				totalPaidL1Tx.Add(totalPaidL1Tx, res.FeeStats.Paid.L1Transaction)
				totalPaidL1Calldata.Add(totalPaidL1Calldata, res.FeeStats.Paid.L1Calldata)
			}
			totalPaidL2Computation.Add(totalPaidL2Computation, res.FeeStats.Paid.L2Computation)
			totalPaidL2Storage.Add(totalPaidL2Storage, res.FeeStats.Paid.L2Storage)
		}
		totalL1Paid := new(big.Int).Add(totalPaidL1Tx, totalPaidL1Calldata)
		totalL2Paid := new(big.Int).Add(totalPaidL2Computation, totalPaidL2Storage)
		totalPaid := new(big.Int).Add(totalL1Paid, totalL2Paid)
		totalPaid.Add(totalPaid, totalPaidUnderfunded)

		totalReceived := new(big.Int).Add(aggBal, netFeeRecipientBal)
		if totalPaid.Cmp(totalReceived) != 0 {
			t.Error("total paid was", totalPaid, "but aggregator + network received", totalReceived)
		}
		return totalL1Paid, totalL2Paid, totalPaidUnderfunded
	}

	checkNoCongestionFee(noFeeSnap)
	checkNoNonPreferredAggFee(noFeeSnap)
	checkNoAggFee(noFeeSnap)

	checkNoCongestionFee(feeSnap)
	checkNoNonPreferredAggFee(feeSnap)
	checkNoAggFee(feeSnap)

	checkNoCongestionFee(feeWithAggSnap)
	checkNoNonPreferredAggFee(feeWithAggSnap)

	checkNoCongestionFee(feeWithContractSnap)
	checkNoNonPreferredAggFee(feeWithContractSnap)
	checkNoAggFee(feeWithContractSnap)

	noFeePaid := checkPaid(noFeeSnap, noFeeResults)
	if noFeePaid.Cmp(big.NewInt(0)) != 0 {
		t.Error("paid fee with fees disabled")
	}

	checkPaid(feeSnap, feeResults)
	checkPaid(feeWithAggSnap, feeWithAggResults)
	checkPaid(feeWithContractSnap, feeWithContractResults)

	checkTxCount(noFeeSnap, noFeeResults)
	checkTxCount(feeSnap, noFeeResults)
	checkTxCount(feeWithAggSnap, feeWithAggResults)

	estimateUserBal, err := estimateFeeSnap.GetBalance(userAddress)
	test.FailIfError(t, err)
	t.Log("Remaining balance", estimateUserBal)

	checkTotalReceived(noFeeSnap, noFeeResults)
	checkTotalReceived(feeSnap, feeResults)
	l1PaidWithAgg, l2PaidWithAgg, l1PaidUnderfunded := checkTotalReceived(feeWithAggSnap, feeWithAggResults)
	{
		t.Helper()
		aggBal, err := feeWithAggSnap.GetBalance(aggregator)
		test.FailIfError(t, err)

		netFeeRecipientBal, err := feeWithAggSnap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)

		l1RatioToAgg := big.NewRat(100, 115)

		l1ToAgg := new(big.Rat).Mul(new(big.Rat).SetInt(l1PaidWithAgg), l1RatioToAgg)
		l1ToNetwork := new(big.Rat).Sub(new(big.Rat).SetInt(l1PaidWithAgg), l1ToAgg)

		totalToNetworkFee := new(big.Rat).Add(l1ToNetwork, new(big.Rat).SetInt(l2PaidWithAgg))

		// add back in the underpaid amounts, since they break the ratio
		totalToNetworkFee.Add(totalToNetworkFee, new(big.Rat).SetInt(l1PaidUnderfunded))

		aggregatorDiff := calcDiffSigned(l1ToAgg, new(big.Rat).SetInt(aggBal))
		networkDiff := calcDiffSigned(totalToNetworkFee, new(big.Rat).SetInt(netFeeRecipientBal))

		if new(big.Rat).Abs(aggregatorDiff).Cmp(big.NewRat(1, 100)) > 0 {
			as_float, _ := aggregatorDiff.Float64()
			t.Error("unexpected aggregator fee collected", 100.0*as_float,
				"\naggregatorDiff", aggregatorDiff, "\naggBal", aggBal, "\nl1ToAgg",
				l1ToAgg, "\nnetFeeRecipientBal", netFeeRecipientBal, "\nl1ToNetwork",
				l1ToNetwork, "\ntotalToNetworkFee", totalToNetworkFee)
		}
		if new(big.Rat).Abs(networkDiff).Cmp(big.NewRat(1, 100)) > 0 {
			as_float, _ := networkDiff.Float64()
			t.Error("unexpected network fee collected", 100.0*as_float, networkDiff)
		}
	}

	_, err = feeSnap.GetPricesInWei()
	test.FailIfError(t, err)
	_, err = feeWithAggSnap.GetPricesInWei()
	test.FailIfError(t, err)
}

func checkSameL2ComputationUnits(t *testing.T, res1 []*evm.TxResult, res2 []*evm.TxResult) {
	for i, res := range res1 {
		unitsUsed1 := res.FeeStats.UnitsUsed
		unitsUsed2 := res2[i].FeeStats.UnitsUsed
		if res.ResultCode == res2[i].ResultCode {
			if new(big.Int).Sub(unitsUsed1.L2Computation, unitsUsed2.L2Computation).CmpAbs(big.NewInt(2000)) > 0 {
				t.Error("computation used outside of acceptable range", unitsUsed1.L2Computation, unitsUsed2.L2Computation)
			}
		}
	}
}

func checkL1UnitsEqual(t *testing.T, unitsUsed1 *evm.FeeSet, unitsUsed2 *evm.FeeSet) {
	t.Helper()
	if unitsUsed1.L1Calldata.Cmp(unitsUsed2.L1Calldata) != 0 {
		t.Error("different calldata used", unitsUsed1.L1Calldata, unitsUsed2.L1Calldata)
	}
	if unitsUsed1.L1Transaction.Cmp(unitsUsed2.L1Transaction) != 0 {
		t.Error("different transaction count used")
	}
}

func checkL2UnitsEqual(t *testing.T, unitsUsed1 *evm.FeeSet, unitsUsed2 *evm.FeeSet) {
	t.Helper()
	if unitsUsed1.L2Computation.Cmp(unitsUsed2.L2Computation) != 0 {
		t.Error("different computation used", unitsUsed1.L2Computation, unitsUsed2.L2Computation)
	}
	if unitsUsed1.L2Storage.Cmp(unitsUsed2.L2Storage) != 0 {
		t.Error("different storage count used", unitsUsed1.L2Storage, unitsUsed2.L2Storage)
	}
}
func checkUnits(t *testing.T, res *evm.TxResult, correct txTemplate, index, calldataUnits int, calldataExact bool) {
	t.Helper()
	unitsUsed := res.FeeStats.UnitsUsed
	t.Log("UnitsUsed", res.FeeStats.UnitsUsed)
	if calldataExact {
		if unitsUsed.L1Calldata.Cmp(big.NewInt(int64(calldataUnits))) != 0 {
			t.Error("wrong calldata used, got", unitsUsed.L1Calldata, "but expected", calldataUnits)
		}
	} else {
		// Adjust units used for gas used
		gasUsed := len(res.CalcGasUsed().Bytes()) * 16
		adjustedCalldata := int(unitsUsed.L1Calldata.Int64()) + gasUsed

		if adjustedCalldata < calldataUnits {
			t.Error("calldata used should be upper bound, got", adjustedCalldata, "but expected", calldataUnits)
		}
		unitsDifference := adjustedCalldata - calldataUnits
		if unitsDifference > 200 {
			t.Error("calldata difference too large", unitsDifference)
		} else {
			t.Log("calldata estimate was over by", unitsDifference)
		}
	}

	if unitsUsed.L1Transaction.Cmp(big.NewInt(1)) != 0 {
		t.Error("should have one tx used")
	}
	nonZeroComp := correct.nonzeroComputation[0]
	if len(correct.nonzeroComputation) > 1 {
		nonZeroComp = correct.nonzeroComputation[index]
	}
	if nonZeroComp {
		if unitsUsed.L2Computation.Cmp(big.NewInt(0)) <= 0 {
			t.Error("should have nonzero computation used")
		}
	} else {
		if unitsUsed.L2Computation.Cmp(big.NewInt(0)) != 0 {
			t.Error("should have zero computation used")
		}
	}
	if unitsUsed.L2Storage.Cmp(big.NewInt(int64(correct.correctStorageUsed))) != 0 {
		t.Error("wrong storage count used got", unitsUsed.L2Storage, "but expected", correct.correctStorageUsed)
	}
}

func checkGas(t *testing.T, res *evm.TxResult, aggregator *common.Address, l2Unpaid bool) *big.Int {
	t.Helper()
	unitsUsed := res.FeeStats.UnitsUsed
	prices := res.FeeStats.Price
	paid := res.FeeStats.Paid
	t.Log("Units", res.FeeStats.UnitsUsed)
	t.Log("Price", res.FeeStats.Price)
	t.Log("Paid", res.FeeStats.Paid, "Total", res.FeeStats.Paid.Total())

	if res.IncomingRequest.AggregatorInfo == nil {
		t.Error("expected aggregator info")
	} else if aggregator != nil && res.IncomingRequest.AggregatorInfo.Aggregator == nil {
		t.Error("should come from aggregator")
	} else if aggregator == nil && res.IncomingRequest.AggregatorInfo.Aggregator != nil {
		t.Error("should not come from aggregator")
	} else if aggregator != nil && *res.IncomingRequest.AggregatorInfo.Aggregator != *aggregator {
		t.Error("wrong aggregator", *res.IncomingRequest.AggregatorInfo.Aggregator)
	}

	l1TxPaidGoal := new(big.Int).Mul(unitsUsed.L1Transaction, prices.L1Transaction)
	l1CalldataGoal := new(big.Int).Mul(unitsUsed.L1Calldata, prices.L1Calldata)
	l2ComputationGoal := new(big.Int).Mul(unitsUsed.L2Computation, prices.L2Computation)
	l2StorageGoal := new(big.Int).Mul(unitsUsed.L2Storage, prices.L2Storage)
	if l1TxPaidGoal.Cmp(paid.L1Transaction) < 0 {
		t.Error("overpaid for l1 tx, paid", paid.L1Transaction, "but wanted", l1TxPaidGoal)
	}
	if l1CalldataGoal.Cmp(paid.L1Calldata) < 0 {
		t.Error("overpaid for l1 calldata got", paid.L1Calldata, "but expected", l1CalldataGoal)
	}
	if l2ComputationGoal.Cmp(paid.L2Computation) < 0 {
		t.Error("overpaid for l2 computation")
	}
	if l2StorageGoal.Cmp(paid.L2Storage) < 0 {
		t.Error("overpaid for l2 storage")
	}

	l1TxUnpaid := new(big.Int).Sub(l1TxPaidGoal, paid.L1Transaction)
	l1CalldataUnpaid := new(big.Int).Sub(l1CalldataGoal, paid.L1Calldata)
	l2ComputationUnpaid := new(big.Int).Sub(l2ComputationGoal, paid.L2Computation)
	l2StorageUnpaid := new(big.Int).Sub(l2StorageGoal, paid.L2Storage)

	if l2ComputationUnpaid.Cmp(big.NewInt(0)) != 0 {
		t.Error("incorrectly unpaid computation amount", l2ComputationUnpaid)
	}

	if l2StorageUnpaid.Cmp(big.NewInt(0)) != 0 {
		t.Error("incorrectly unpaid storage amount", l2StorageUnpaid)
	}

	totalUnpaid := new(big.Int).Add(l1TxUnpaid, l1CalldataUnpaid)
	if totalUnpaid.Cmp(big.NewInt(0)) != 0 && res.ResultCode != evm.InsufficientGasForBaseFee && res.ResultCode != evm.InsufficientGasFundsCode {
		t.Log("totalUnpaid: ", l1TxUnpaid, "+", l1CalldataUnpaid)
		t.Error("gas left unpaid, but got wrong result", res.ResultCode)
	}
	return totalUnpaid
}
