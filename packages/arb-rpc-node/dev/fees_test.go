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

package dev

import (
	"context"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func TestFees(t *testing.T) {
	skipBelowVersion(t, 3)
	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	test.FailIfError(t, err)
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	privkey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	auth := bind.NewKeyedTransactor(privkey)

	privkey2, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	aggAuth := bind.NewKeyedTransactor(privkey2)

	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	netFeeRecipient := common.RandAddress()
	congestionFeeRecipient := common.RandAddress()
	feeConfigInit := message.FeeConfig{
		SpeedLimitPerSecond:    new(big.Int).SetUint64(config.ArbGasSpeedLimitPerSecond),
		L1GasPerL2Tx:           big.NewInt(3700),
		L1GasPerL2Calldata:     big.NewInt(1),
		L1GasPerStorage:        big.NewInt(2000),
		ArbGasDivisor:          big.NewInt(10000),
		NetFeeRecipient:        netFeeRecipient,
		CongestionFeeRecipient: congestionFeeRecipient,
	}

	aggInit := message.DefaultAggConfig{Aggregator: common.NewAddressFromEth(aggAuth.From)}
	monitor, backend, db, rollupAddress := NewDevNode(tmpDir, *arbosfile, config, common.NewAddressFromEth(auth.From), []message.ChainConfigOption{feeConfigInit, aggInit})
	defer monitor.Close()
	defer db.Close()

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(auth.From),
				Payment:     new(big.Int).Exp(big.NewInt(10), big.NewInt(22), nil),
				Data:        nil,
			},
		}),
	}
	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}

	srv := aggregator.NewServer(backend, rollupAddress, db)
	client := web3.NewEthClient(srv, true)

	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)
	arbGasInfo, err := arboscontracts.NewArbGasInfo(arbos.ARB_GAS_INFO_ADDRESS, client)
	test.FailIfError(t, err)
	arbAggregator, err := arboscontracts.NewArbAggregator(arbos.ARB_AGGREGATOR_ADDRESS, client)
	test.FailIfError(t, err)

	feeCollector := common.RandAddress()

	_, feeCollectorErr := arbAggregator.SetFeeCollector(aggAuth, aggInit.Aggregator.ToEthAddress(), feeCollector.ToEthAddress())

	_, err = arbOwner.SetFairGasPriceSender(auth, aggInit.Aggregator.ToEthAddress())
	test.FailIfError(t, err)

	_, err = arbOwner.SetFeesEnabled(auth, true)
	test.FailIfError(t, err)

	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}

	totalPaid := big.NewInt(0)
	for i := 0; i < 5; i++ {
		tx, err := arbOwner.GiveOwnership(auth, auth.From)
		test.FailIfError(t, err)
		paid := checkFees(t, backend, tx)
		totalPaid = totalPaid.Add(totalPaid, paid)
	}

	networkDest, congestionDest, err := arbOwner.GetFeeRecipients(&bind.CallOpts{})
	test.FailIfError(t, err)
	if networkDest != netFeeRecipient.ToEthAddress() {
		t.Error("wrong network dest", networkDest)
	}
	if congestionDest != congestionFeeRecipient.ToEthAddress() {
		t.Error("wrong congestion dest", congestionDest)
	}

	speedLimitPerSecond, gasPoolMax, maxTxGasLimit, err := arbGasInfo.GetGasAccountingParams(&bind.CallOpts{})
	test.FailIfError(t, err)
	if speedLimitPerSecond.Cmp(new(big.Int).SetUint64(config.ArbGasSpeedLimitPerSecond)) != 0 {
		t.Error("wrong speed limit")
	}
	t.Log("gasPoolMax", gasPoolMax)
	t.Log("maxTxGasLimit", maxTxGasLimit)

	perL2TxWei, perL1CalldataByteWei, perStorageWei, perArgGasBaseWei, perArbGasCongestionWei, perArbGasTotalWei, err := arbGasInfo.GetPricesInWei(&bind.CallOpts{})
	test.FailIfError(t, err)
	t.Log("perL2TxWei", perL2TxWei)
	t.Log("perL1CalldataByteWei", perL1CalldataByteWei)
	t.Log("perStorageWei", perStorageWei)
	t.Log("perArgGasBaseWei", perArgGasBaseWei)
	t.Log("perArbGasCongestionWei", perArbGasCongestionWei)
	t.Log("perArbGasTotalWei", perArbGasTotalWei)

	perL2Tx, perL1CalldataByte, perStorage, err := arbGasInfo.GetPricesInArbGas(&bind.CallOpts{})
	test.FailIfError(t, err)
	t.Log("perL2Tx", perL2Tx)
	t.Log("perL1CalldataByte", perL1CalldataByte)
	t.Log("perStorage", perStorage)

	_, tx, _, err := arbostestcontracts.DeploySimple(auth, client)
	test.FailIfError(t, err)

	paid := checkFees(t, backend, tx)
	totalPaid = totalPaid.Add(totalPaid, paid)

	netFeeBal, err := client.BalanceAt(context.Background(), netFeeRecipient.ToEthAddress(), nil)
	test.FailIfError(t, err)

	aggBal, err := client.BalanceAt(context.Background(), aggInit.Aggregator.ToEthAddress(), nil)
	test.FailIfError(t, err)

	feeCollectorBal, err := client.BalanceAt(context.Background(), feeCollector.ToEthAddress(), nil)
	test.FailIfError(t, err)

	totalReceived := new(big.Int).Add(netFeeBal, aggBal)
	totalReceived = totalReceived.Add(totalReceived, feeCollectorBal)
	if totalReceived.Cmp(totalPaid) != 0 {
		t.Error("amount paid different than amount received")
	}

	if arbosVersion <= 4 {
		if aggBal.Cmp(big.NewInt(0)) <= 0 {
			t.Error("aggregator should have nonzero balance")
		}
		if feeCollectorBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("fee collector should have 0 balance")
		}
	} else {
		test.FailIfError(t, feeCollectorErr)
		if aggBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("aggregator should have 0 balance")
		}
		if feeCollectorBal.Cmp(big.NewInt(0)) <= 0 {
			t.Error("fee collector should have nonzero balance")
		}
	}
	t.Log("Paid", totalPaid)
	t.Log("Net bal", netFeeBal)
	t.Log("Agg bal", aggBal)
	t.Log("Fee col bal", feeCollectorBal)
}

func checkFees(t *testing.T, backend *Backend, tx *types.Transaction) *big.Int {
	arbRes, err := backend.db.GetRequest(common.NewHashFromEth(tx.Hash()))
	test.FailIfError(t, err)
	return arbRes.FeeStats.Paid.Total()
}
