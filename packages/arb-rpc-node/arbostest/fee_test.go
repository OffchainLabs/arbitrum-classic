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

func addUserTxes(t *testing.T, ib *InboxBuilder, privKey *ecdsa.PrivateKey, aggregator common.Address, initialDeposit *big.Int) {
	userAddress := crypto.PubkeyToAddress(privKey.PublicKey)
	signer := types.NewEIP155Signer(message.ChainAddressToID(chain))
	gasUsedABI, err := abi.JSON(strings.NewReader(arbostestcontracts.GasUsedABI))
	failIfError(t, err)

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

	rawTx1 := types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(0),
		Gas:      1000000000,
		Value:    big.NewInt(0),
		Data:     hexutil.MustDecode(arbostestcontracts.GasUsedBin),
	})

	dest := crypto.CreateAddress(userAddress, 0)
	rawTx2 := types.NewTx(&types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(0),
		Gas:      100000000,
		To:       &dest,
		Value:    big.NewInt(0),
		Data:     makeFuncData(t, gasUsedABI.Methods["sstore"]),
	})

	userTxes := []*types.Transaction{rawTx1, rawTx2}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(int64(len(ib.Messages))),
		Timestamp: big.NewInt(0),
	}
	ib.AddMessage(deposit, chain, big.NewInt(0), chainTime)

	for _, rawTx := range userTxes {
		tx, err := types.SignTx(rawTx, signer, privKey)
		failIfError(t, err)

		batch, err := message.NewTransactionBatchFromMessages([]message.AbstractL2Message{message.NewCompressedECDSAFromEth(tx)})
		failIfError(t, err)
		ib.AddMessage(message.NewSafeL2Message(batch), aggregator, big.NewInt(0), chainTime)
		chainTime.BlockNum = common.NewTimeBlocksInt(int64(len(ib.Messages)))
	}
}

func TestFees(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	failIfError(t, err)
	userAddress := common.NewAddressFromEth(crypto.PubkeyToAddress(privKey.PublicKey))
	initialDeposit := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

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
		addUserTxes(t, ib, privKey, agg, initialDeposit)
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

	noFeeResults, noFeeSnap := processMessages(t, noFeeIB)
	feeResults, feeSnap := processMessages(t, feeIB)
	feeWithAggResults, feeWithAggSnap := processMessages(t, feeWithAggIB)

	noFeeRes1, noFeeRes2 := checkResults(t, "no fee", noFeeResults, aggregator)
	feeRes1, feeRes2 := checkResults(t, "fee", feeResults, otherAgg)
	feeWithAggRes1, feeWithAggRes2 := checkResults(t, "fee with agg", feeWithAggResults, aggregator)

	checkUnitsEqual(t, noFeeRes1, feeWithAggRes1)
	checkUnitsEqual(t, noFeeRes2, feeWithAggRes2)
	checkUnitsEqual(t, noFeeRes1, feeRes1)
	checkUnitsEqual(t, noFeeRes2, feeRes2)

	checkBalance := func(snap *snapshot.Snapshot, results []*evm.TxResult, preferredAgg, usedAgg common.Address) {
		t.Helper()
		userBal, err := snap.GetBalance(userAddress)
		test.FailIfError(t, err)
		preferredAggBal, err := snap.GetBalance(preferredAgg)
		test.FailIfError(t, err)
		usedAggBal, err := snap.GetBalance(usedAgg)
		test.FailIfError(t, err)
		netFeeRecipientBal, err := snap.GetBalance(netFeeRecipient)
		test.FailIfError(t, err)
		congestionFeeRecipientBal, err := snap.GetBalance(congestionFeeRecipient)
		test.FailIfError(t, err)

		res1 := results[len(results)-2]
		res2 := results[len(results)-1]
		reportedPaid := new(big.Int).Add(res1.FeeStats.Paid.Total(), res2.FeeStats.Paid.Total())
		amountPaid := new(big.Int).Sub(initialDeposit, userBal)
		if amountPaid.Cmp(reportedPaid) != 0 {
			t.Error("wrong amount deducted from user expected", reportedPaid, "but got", amountPaid)
		}
		t.Log("Preferred agg bal", preferredAggBal)
		t.Log("Used agg bal", usedAggBal)
		t.Log("Net fee bal", netFeeRecipientBal)
		t.Log("Congestion fee bal", congestionFeeRecipientBal)
	}

	checkBalance(noFeeSnap, noFeeResults, aggregator, aggregator)
	checkBalance(feeSnap, feeResults, aggregator, otherAgg)
	checkBalance(feeWithAggSnap, feeWithAggResults, aggregator, aggregator)
}

func processMessages(t *testing.T, ib *InboxBuilder) ([]*evm.TxResult, *snapshot.Snapshot) {
	logs, _, snap, _ := runAssertion(t, ib.Messages, math.MaxInt32, 0)
	results := extractTxResults(t, logs)
	allResultsSucceeded(t, results)
	return results, snap
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

func checkResults(t *testing.T, label string, results []*evm.TxResult, aggregator common.Address) (*evm.TxResult, *evm.TxResult) {
	t.Log("Checking results for", label)
	storageUsed0 := (len(hexutil.MustDecode(arbostestcontracts.GasUsedBin)) + 32) / 32
	storageUsed1 := 1

	res1 := results[len(results)-2]
	res2 := results[len(results)-1]

	checkGas(t, res1, storageUsed0, aggregator)
	checkGas(t, res2, storageUsed1, aggregator)
	return res1, res2
}

func checkGas(t *testing.T, res *evm.TxResult, correctStorageUsed int, aggregator common.Address) {
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
	if unitsUsed.L2Computation.Cmp(big.NewInt(0)) <= 0 {
		t.Error("should have nonzero computation used")
	}
	if unitsUsed.L2Storage.Cmp(big.NewInt(int64(correctStorageUsed))) != 0 {
		t.Error("wrong storage count used")
	}

	if res.IncomingRequest.AggregatorInfo.Aggregator == nil {
		t.Error("should come from aggregator")
	} else if *res.IncomingRequest.AggregatorInfo.Aggregator != aggregator {
		t.Error("wrong aggregator", *res.IncomingRequest.AggregatorInfo.Aggregator)
	}

	l1TxPaid := new(big.Int).Mul(unitsUsed.L1Transaction, prices.L1Transaction)
	if l1TxPaid.Cmp(paid.L1Transaction) != 0 {
		t.Error("wrong paid for l1 transaction got", paid.L1Transaction, "expected", l1TxPaid)
	}
	l1CalldataPaid := new(big.Int).Mul(unitsUsed.L1Calldata, prices.L1Calldata)
	if l1CalldataPaid.Cmp(paid.L1Calldata) != 0 {
		t.Error("wrong paid for l1 calldata got", paid.L1Calldata, "expected", l1CalldataPaid)
	}
	if new(big.Int).Mul(unitsUsed.L2Computation, prices.L2Computation).Cmp(paid.L2Computation) != 0 {
		t.Error("wrong paid for l2 computation")
	}
	if new(big.Int).Mul(unitsUsed.L2Storage, prices.L2Storage).Cmp(paid.L2Storage) != 0 {
		t.Error("wrong paid for l2 storage")
	}
}
