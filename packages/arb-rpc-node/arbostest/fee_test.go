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
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"math/big"
	"strings"
	"testing"
)

func addInitialization(ib *InboxBuilder, aggregator, netFeeRecipient, congestionFeeRecipient common.Address) {
	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(0),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
		MaxExecutionSteps:         0,
		ArbGasSpeedLimitPerSecond: 1000000000,
	}

	feeConfigInit := message.FeeConfig{
		SpeedLimitPerSecond:    new(big.Int).SetUint64(config.ArbGasSpeedLimitPerSecond),
		L1GasPerL2Tx:           big.NewInt(3700),
		L1GasPerL2Calldata:     big.NewInt(16),
		L1GasPerStorage:        big.NewInt(2000),
		ArbGasDivisor:          big.NewInt(10000),
		NetFeeRecipient:        netFeeRecipient,
		CongestionFeeRecipient: congestionFeeRecipient,
	}
	aggInit := message.DefaultAggConfig{Aggregator: aggregator}
	init := message.NewInitMessage(config, owner, []message.ChainConfigOption{feeConfigInit, aggInit})

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	ib.AddMessage(init, chain, big.NewInt(0), chainTime)
}

func addEnableFeesMessages(ib *InboxBuilder) {
	ownerTx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.SetFairGasPriceSender(owner),
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

func addUserTxes(t *testing.T, ib *InboxBuilder, privKey *ecdsa.PrivateKey, userTxes []txTemplate, aggregator common.Address, initialDeposit *big.Int) {
	userAddress := crypto.PubkeyToAddress(privKey.PublicKey)
	signer := types.NewEIP155Signer(message.ChainAddressToID(chain))

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(10000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(userAddress),
				Payment:     initialDeposit,
				Data:        nil,
			},
		}),
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(int64(len(ib.Messages))),
		Timestamp: big.NewInt(0),
	}
	ib.AddMessage(deposit, chain, big.NewInt(0), chainTime)

	for i, rawTx := range userTxes {
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    uint64(i),
			GasPrice: rawTx.GasPrice,
			Gas:      rawTx.Gas,
			To:       rawTx.To,
			Value:    rawTx.Value,
			Data:     rawTx.Data,
		})
		signedTx, err := types.SignTx(tx, signer, privKey)
		failIfError(t, err)

		batch, err := message.NewTransactionBatchFromMessages([]message.AbstractL2Message{message.NewCompressedECDSAFromEth(signedTx)})
		failIfError(t, err)
		ib.AddMessage(message.NewSafeL2Message(batch), aggregator, big.NewInt(0), chainTime)
		chainTime.BlockNum = common.NewTimeBlocksInt(int64(len(ib.Messages)))
	}
}

type txTemplate struct {
	GasPrice *big.Int
	Gas      uint64
	To       *ethcommon.Address
	Value    *big.Int
	Data     []byte

	// Data to verify tx
	nonzeroComputation bool
	correctStorageUsed int
}

func TestFees(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	failIfError(t, err)
	userAddress := common.NewAddressFromEth(crypto.PubkeyToAddress(privKey.PublicKey))

	initialDeposit := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	gasUsedABI, err := abi.JSON(strings.NewReader(arbostestcontracts.GasUsedABI))
	failIfError(t, err)

	contractDest := crypto.CreateAddress(userAddress.ToEthAddress(), 0)
	eoaDest := common.RandAddress().ToEthAddress()
	rawTxes := []txTemplate{
		{
			GasPrice: big.NewInt(0),
			Gas:      1000000000,
			Value:    big.NewInt(0),
			Data:     hexutil.MustDecode(arbostestcontracts.GasUsedBin),

			nonzeroComputation: true,
			correctStorageUsed: (len(hexutil.MustDecode(arbostestcontracts.GasUsedBin)) + 32) / 32,
		},
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     makeFuncData(t, gasUsedABI.Methods["noop"]),

			nonzeroComputation: true,
			correctStorageUsed: 0,
		},
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &contractDest,
			Value:    big.NewInt(0),
			Data:     makeFuncData(t, gasUsedABI.Methods["sstore"]),

			nonzeroComputation: true,
			correctStorageUsed: 1,
		},
		{
			GasPrice: big.NewInt(0),
			Gas:      100000000,
			To:       &eoaDest,
			Value:    big.NewInt(100000),

			nonzeroComputation: false,
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
		addInitialization(ib, aggregator, netFeeRecipient, congestionFeeRecipient)
	}

	addUserTxesLoc := func(ib *InboxBuilder, agg common.Address) {
		t.Helper()
		addUserTxes(t, ib, privKey, rawTxes, agg, initialDeposit)
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

	noFeeResults, noFeeSnap := processMessages(t, noFeeIB, len(rawTxes))
	feeResults, feeSnap := processMessages(t, feeIB, len(rawTxes))
	feeWithAggResults, feeWithAggSnap := processMessages(t, feeWithAggIB, len(rawTxes))

	checkResults := func(results []*evm.TxResult, aggregator common.Address) {
		for i, res := range results {
			checkGas(t, res, rawTxes[i].nonzeroComputation, rawTxes[i].correctStorageUsed, aggregator)
		}
	}

	t.Log("Checking results for no fee")
	checkResults(noFeeResults, aggregator)
	t.Log("Checking results for fee")
	checkResults(feeResults, otherAgg)
	t.Log("Checking results for fee with agg")
	checkResults(feeWithAggResults, aggregator)

	checkPairedUnitsEqual(t, noFeeResults, feeResults)
	checkPairedUnitsEqual(t, noFeeResults, feeWithAggResults)

	calcDiff := func(a, b *big.Rat) *big.Rat {
		diff := new(big.Rat).Sub(a, b)
		return diff.Abs(diff).Quo(diff, b)
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
			t.Error("tx price with agg is wrong")
		}

		l1CalldataDiff := calculateFeeAggDiff(noAggPrice.L1Calldata, aggPrice.L1Calldata)
		if l1CalldataDiff.Cmp(big.NewRat(1, 100)) > 0 {
			t.Error("tx price with agg is wrong")
		}

		if noAggPrice.L2Computation.Cmp(aggPrice.L2Computation) != 0 {
			t.Error("wrong l2 computation price")
		}

		if noAggPrice.L2Storage.Cmp(aggPrice.L2Storage) != 0 {
			t.Error("wrong l2 storage price")
		}
	}

	checkPaid := func(snap *snapshot.Snapshot, results []*evm.TxResult) *big.Int {
		t.Helper()

		txCount, err := snap.GetTransactionCount(userAddress)
		test.FailIfError(t, err)

		if txCount.Cmp(big.NewInt(int64(len(rawTxes)))) != 0 {
			t.Error("wrong tx count", txCount)
		}

		userBal, err := snap.GetBalance(userAddress)
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

	checkTotalReceived := func(snap *snapshot.Snapshot, results []*evm.TxResult) (*big.Int, *big.Int) {
		t.Helper()
		aggBal, err := snap.GetBalance(aggregator)
		test.FailIfError(t, err)

		netFeeRecipientBal, err := snap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)

		totalPaidL1Tx := big.NewInt(0)
		totalPaidL1Calldata := big.NewInt(0)
		totalPaidL2Computation := big.NewInt(0)
		totalPaidL2Storage := big.NewInt(0)
		for _, res := range results {
			totalPaidL1Tx = totalPaidL1Tx.Add(totalPaidL1Tx, res.FeeStats.Paid.L1Transaction)
			totalPaidL1Calldata = totalPaidL1Calldata.Add(totalPaidL1Calldata, res.FeeStats.Paid.L1Calldata)
			totalPaidL2Computation = totalPaidL2Computation.Add(totalPaidL2Computation, res.FeeStats.Paid.L2Computation)
			totalPaidL2Storage = totalPaidL2Storage.Add(totalPaidL2Storage, res.FeeStats.Paid.L2Storage)
		}
		totalL1Paid := new(big.Int).Add(totalPaidL1Tx, totalPaidL1Calldata)
		totalL2Paid := new(big.Int).Add(totalPaidL2Computation, totalPaidL2Storage)
		totalPaid := new(big.Int).Add(totalL1Paid, totalL2Paid)

		totalReceived := new(big.Int).Add(aggBal, netFeeRecipientBal)
		if totalPaid.Cmp(totalReceived) != 0 {
			t.Error("total paid was", totalPaid, "but aggregator + network received", totalReceived)
		}
		return totalL1Paid, totalL2Paid
	}

	checkNoCongestionFee(noFeeSnap)
	checkNoNonPreferredAggFee(noFeeSnap)
	checkNoAggFee(noFeeSnap)

	checkNoCongestionFee(feeSnap)
	checkNoNonPreferredAggFee(feeSnap)
	checkNoAggFee(feeSnap)

	checkNoCongestionFee(feeWithAggSnap)
	checkNoNonPreferredAggFee(feeWithAggSnap)

	noFeePaid := checkPaid(noFeeSnap, noFeeResults)
	if noFeePaid.Cmp(big.NewInt(0)) != 0 {
		t.Error("paid fee with fees disabled")
	}

	noAggPaid := checkPaid(feeSnap, feeResults)
	checkPaid(feeWithAggSnap, feeWithAggResults)

	checkTotalReceived(noFeeSnap, noFeeResults)
	checkTotalReceived(feeSnap, feeResults)
	l1PaidWithAgg, l2PaidWithAgg := checkTotalReceived(feeWithAggSnap, feeWithAggResults)
	{
		netFeeRecipientBal, err := feeWithAggSnap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)

		if netFeeRecipientBal.Cmp(noAggPaid) != 0 {
			t.Error("network fee should be the same")
		}
	}

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

		if calcDiff(l1ToAgg, new(big.Rat).SetInt(aggBal)).Cmp(big.NewRat(1, 100)) > 0 {
			t.Error("unexpected aggregator fee collected")
		}

		if calcDiff(totalToNetworkFee, new(big.Rat).SetInt(netFeeRecipientBal)).Cmp(big.NewRat(1, 100)) > 0 {
			t.Error("unexpected network fee collected")
		}
	}
}

func processMessages(t *testing.T, ib *InboxBuilder, userTxCount int) ([]*evm.TxResult, *snapshot.Snapshot) {
	t.Helper()
	logs, _, snap, _ := runAssertion(t, ib.Messages, math.MaxInt32, 0)
	results := extractTxResults(t, logs)
	allResultsSucceeded(t, results)
	return results[len(results)-userTxCount:], snap
}

func checkPairedUnitsEqual(t *testing.T, res1 []*evm.TxResult, res2 []*evm.TxResult) {
	t.Helper()
	for i, res := range res1 {
		checkUnitsEqual(t, res, res2[i])
	}
}

func checkUnitsEqual(t *testing.T, res1 *evm.TxResult, res2 *evm.TxResult) {
	t.Helper()
	unitsUsed1 := res1.FeeStats.UnitsUsed
	unitsUsed2 := res2.FeeStats.UnitsUsed
	if unitsUsed1.L1Calldata.Cmp(unitsUsed2.L1Calldata) != 0 {
		t.Error("different calldata used", unitsUsed1.L1Calldata, unitsUsed2.L1Calldata)
	}
	if unitsUsed1.L1Transaction.Cmp(unitsUsed2.L1Transaction) != 0 {
		t.Error("different transaction count used")
	}
	if new(big.Int).Sub(unitsUsed1.L2Computation, unitsUsed2.L2Computation).CmpAbs(big.NewInt(2000)) > 0 {
		t.Error("computation used outside of acceptable range", unitsUsed1.L2Computation, unitsUsed2.L2Computation)
	}
	if unitsUsed1.L2Storage.Cmp(unitsUsed2.L2Storage) != 0 {
		t.Error("different storage count used", unitsUsed1.L2Storage, unitsUsed2.L2Storage)
	}
}

func checkGas(t *testing.T, res *evm.TxResult, nonzeroComputation bool, correctStorageUsed int, aggregator common.Address) {
	t.Helper()
	unitsUsed := res.FeeStats.UnitsUsed
	prices := res.FeeStats.Price
	paid := res.FeeStats.Paid
	t.Log("UnitsUsed", res.FeeStats.UnitsUsed)
	t.Log("Price", res.FeeStats.Price)
	t.Log("Paid", res.FeeStats.Paid, "Total", res.FeeStats.Paid.Total())
	if unitsUsed.L1Calldata.Cmp(big.NewInt(0)) <= 0 {
		t.Error("should have nonzero calldata used")
	}
	if unitsUsed.L1Transaction.Cmp(big.NewInt(1)) != 0 {
		t.Error("should have one tx used")
	}
	if nonzeroComputation {
		if unitsUsed.L2Computation.Cmp(big.NewInt(0)) <= 0 {
			t.Error("should have nonzero computation used")
		}
	} else {
		if unitsUsed.L2Computation.Cmp(big.NewInt(0)) != 0 {
			t.Error("should have zero computation used")
		}
	}

	if unitsUsed.L2Storage.Cmp(big.NewInt(int64(correctStorageUsed))) != 0 {
		t.Error("wrong storage count used got", unitsUsed.L2Storage, "but expected", correctStorageUsed)
	}

	if res.IncomingRequest.AggregatorInfo.Aggregator == nil {
		t.Error("should come from aggregator")
	} else if *res.IncomingRequest.AggregatorInfo.Aggregator != aggregator {
		t.Error("wrong aggregator", *res.IncomingRequest.AggregatorInfo.Aggregator)
	}

	l1TxPaid := new(big.Int).Mul(unitsUsed.L1Transaction, prices.L1Transaction)
	if l1TxPaid.Cmp(paid.L1Transaction) != 0 {
		t.Error("wrong paid for l1 transaction got", paid.L1Transaction, "but expected", l1TxPaid)
	}
	l1CalldataPaid := new(big.Int).Mul(unitsUsed.L1Calldata, prices.L1Calldata)
	if l1CalldataPaid.Cmp(paid.L1Calldata) != 0 {
		t.Error("wrong paid for l1 calldata got", paid.L1Calldata, "but expected", l1CalldataPaid)
	}
	if new(big.Int).Mul(unitsUsed.L2Computation, prices.L2Computation).Cmp(paid.L2Computation) != 0 {
		t.Error("wrong paid for l2 computation")
	}
	if new(big.Int).Mul(unitsUsed.L2Storage, prices.L2Storage).Cmp(paid.L2Storage) != 0 {
		t.Error("wrong paid for l2 storage")
	}
}
