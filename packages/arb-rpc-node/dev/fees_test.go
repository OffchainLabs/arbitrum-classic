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
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func setupFeeChain(t *testing.T) (*Backend, *web3.Server, *web3.EthClient, *bind.TransactOpts, *bind.TransactOpts, message.FeeConfig, protocol.ChainParams, common.Address, func()) {
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
		ArbGasSpeedLimitPerSecond: 2000000000,
	}

	netFeeRecipient := common.RandAddress()
	congestionFeeRecipient := common.RandAddress()
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

	aggInit := message.DefaultAggConfig{Aggregator: common.NewAddressFromEth(aggAuth.From)}
	backend, _, srv, cancelDevNode := NewTestDevNode(
		t,
		*arbosfile,
		config,
		common.NewAddressFromEth(auth.From),
		[]message.ChainConfigOption{feeConfigInit, aggInit},
	)

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(auth.From),
				Payment:     new(big.Int).Exp(big.NewInt(10), big.NewInt(17), nil),
				Data:        nil,
			},
		}),
	}
	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}

	web3Server := web3.NewServer(srv, true)

	client := web3.NewEthClient(srv, true)

	arbAggregator, err := arboscontracts.NewArbAggregator(arbos.ARB_AGGREGATOR_ADDRESS, client)
	test.FailIfError(t, err)

	feeCollector := common.RandAddress()
	aggAuth.GasLimit = 100000000
	_, feeCollectorErr := arbAggregator.SetFeeCollector(aggAuth, aggAuth.From, feeCollector.ToEthAddress())
	aggAuth.GasLimit = 0
	if arbosVersion >= 5 {
		test.FailIfError(t, feeCollectorErr)
	}

	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)

	auth.GasLimit = 100000000
	_, err = arbOwner.SetFairGasPriceSender(auth, aggInit.Aggregator.ToEthAddress())
	test.FailIfError(t, err)

	_, err = arbOwner.SetFeesEnabled(auth, true)
	test.FailIfError(t, err)
	auth.GasLimit = 0

	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}
	return backend, web3Server, client, auth, aggAuth, feeConfigInit, config, feeCollector, cancelDevNode
}

func TestFees(t *testing.T) {
	skipBelowVersion(t, 3)
	backend, _, client, auth, aggAuth, feeConfig, config, feeCollector, cancel := setupFeeChain(t)
	defer cancel()

	agg := common.NewAddressFromEth(aggAuth.From)

	arbGasInfo, err := arboscontracts.NewArbGasInfo(arbos.ARB_GAS_INFO_ADDRESS, client)
	test.FailIfError(t, err)

	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)

	totalPaid := big.NewInt(0)
	for i := 0; i < 5; i++ {
		tx, err := arbOwner.GiveOwnership(auth, auth.From)
		test.FailIfError(t, err)
		paid := checkFees(t, backend, tx)
		totalPaid = totalPaid.Add(totalPaid, paid)
	}

	networkDest, congestionDest, err := arbOwner.GetFeeRecipients(&bind.CallOpts{})
	test.FailIfError(t, err)
	if networkDest != feeConfig.NetFeeRecipient.ToEthAddress() {
		t.Error("wrong network dest", networkDest)
	}
	if congestionDest != feeConfig.CongestionFeeRecipient.ToEthAddress() {
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

	netFeeBal, err := client.BalanceAt(context.Background(), feeConfig.NetFeeRecipient.ToEthAddress(), nil)
	test.FailIfError(t, err)

	aggBal, err := client.BalanceAt(context.Background(), agg.ToEthAddress(), nil)
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
			t.Error("currentAggregator should have nonzero balance")
		}
		if feeCollectorBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("fee collector should have 0 balance")
		}
	} else {
		if aggBal.Cmp(big.NewInt(0)) != 0 {
			t.Error("currentAggregator should have 0 balance")
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
	t.Helper()
	arbRes, err := backend.db.GetRequest(common.NewHashFromEth(tx.Hash()))
	test.FailIfError(t, err)
	t.Log("Gas used:", arbRes.CalcGasUsed().Uint64())
	extra := tx.Gas() - arbRes.CalcGasUsed().Uint64()
	t.Log("gas remaining", extra)
	if extra > 5000000 {
		t.Error("too much extra gas estimated")
	}
	return arbRes.FeeStats.Paid.Total()
}

func TestNonAggregatorFee(t *testing.T) {
	skipBelowVersion(t, 3)
	backend, web3SServer, client, auth, _, _, _, _, cancel := setupFeeChain(t)
	defer cancel()

	simpleAddr, _, simple, err := arbostestcontracts.DeploySimple(auth, client)
	test.FailIfError(t, err)
	backend.currentAggregator = common.Address{}

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	test.FailIfError(t, err)
	data := simpleABI.Methods["exists"].ID
	emptyAgg := ethcommon.Address{}

	estimatedGas, err := web3SServer.EstimateGas(web3.CallTxArgs{
		From:       &auth.From,
		To:         &simpleAddr,
		Data:       (*hexutil.Bytes)(&data),
		Aggregator: &emptyAgg,
	})
	test.FailIfError(t, err)
	auth.GasLimit = uint64(estimatedGas)
	tx, err := simple.Exists(auth)
	test.FailIfError(t, err)
	checkFees(t, backend, tx)
}

func TestDeposit(t *testing.T) {
	skipBelowVersion(t, 3)
	backend, _, client, _, _, _, _, _, cancel := setupFeeChain(t)
	defer cancel()

	tx := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(5000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.RandAddress(),
				Payment:     big.NewInt(0),
				Data:        nil,
			},
		}),
	}
	txHash, err := backend.AddInboxMessage(tx, common.RandAddress())
	test.FailIfError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), txHash.ToEthHash())
	test.FailIfError(t, err)

	if receipt == nil {
		t.Fatal("expected receipt")
	}
	block, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	test.FailIfError(t, err)
	if len(block.Transactions()) != 1 {
		t.Fatal("expected 1 tx in block")
	}

	arbRes, err := backend.db.GetRequest(txHash)
	test.FailIfError(t, err)

	t.Log("arbRes", arbRes.IncomingRequest.Kind)
}
