//go:build openethereumtest
// +build openethereumtest

/*
* Copyright 2022, Offchain Labs, Inc.
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

// This test expects to be run against an open ethereum dev node. To boot the node, run
// openethereum --config dev-insecure --tracing=on

package dev

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
	"math/big"
	"os"
	"reflect"
	"testing"
)

func TestTraceOpenEthereum(t *testing.T) {
	skipBelowVersion(t, 25)
	oeUrl := os.Getenv("TEST_OPEN_ETHEREUM")
	if oeUrl == "" {
		oeUrl = "http://127.0.0.1:8545"
	}

	ctx := context.Background()
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	backend, _, srv, cancelDevNode := NewSimpleTestDevNode(t, config, common.RandAddress())
	defer cancelDevNode()
	ethServer := web3.NewServer(srv, web3.DefaultConfig, nil)
	arbClient := web3.NewEthClient(srv, true)
	arbSigner := types.NewEIP155Signer(backend.chainID)
	tracer := web3.NewTracer(ethServer, configuration.DefaultCoreSettingsMaxExecution())

	oeRpc, err := rpc.Dial(oeUrl)
	test.FailIfError(t, err)
	oeClient := ethclient.NewClient(oeRpc)
	oeSigner := types.NewEIP155Signer(big.NewInt(17))

	senderKey := test.MustGenerateKey(t)
	senderAddr := crypto.PubkeyToAddress(senderKey.PublicKey)

	// Fund arb
	deposit := makeDepositMessage(common.NewAddressFromEth(senderAddr))
	_, err = backend.AddInboxMessage(ctx, deposit, common.RandAddress())
	test.FailIfError(t, err)

	// Fund oe
	oeFundingKey, err := crypto.ToECDSA(hexutil.MustDecode("0x4d5db4107d237df6a3d58ee5f70ae63d73d7658d4026f2eefd2f204c81682cb7"))
	test.FailIfError(t, err)
	{
		oeFundingAddr := crypto.PubkeyToAddress(oeFundingKey.PublicKey)
		startNonce, err := oeClient.NonceAt(ctx, oeFundingAddr, nil)
		test.FailIfError(t, err)
		gasPrice, err := oeClient.SuggestGasPrice(ctx)
		test.FailIfError(t, err)
		tx, err := types.SignNewTx(oeFundingKey, oeSigner, &types.LegacyTx{
			Nonce:    startNonce,
			GasPrice: gasPrice,
			Gas:      30000,
			To:       &senderAddr,
			Value:    big.NewInt(100000000),
		})
		test.FailIfError(t, err)
		err = oeClient.SendTransaction(ctx, tx)
	}

	_, arbUserTx := makeTxes(t, arbClient, arbSigner, senderKey)
	arbTraceData, err := tracer.Transaction(ctx, arbUserTx.Bytes())
	test.FailIfError(t, err)

	_, oeUserTx := makeTxes(t, oeClient, oeSigner, senderKey)
	var oeTraceData []web3.TraceFrame
	err = oeRpc.Call(&oeTraceData, "trace_transaction", hexutil.Bytes(oeUserTx.Bytes()))
	test.FailIfError(t, err)

	arbTraceData = clearTxTraceData(arbTraceData)
	oeTraceData = clearTxTraceData(oeTraceData)
	assertTraceEqual(t, arbTraceData, oeTraceData)
}

func clearTxTraceData(traces []web3.TraceFrame) []web3.TraceFrame {
	newTraces := make([]web3.TraceFrame, 0, len(traces))
	var removedTraces [][]int
	for _, trace := range traces {
		// Remove selfdestruct frames
		if trace.Type == "suicide" {
			removedTraces = append(removedTraces, trace.TraceAddress)
			continue
		}
		trace.TransactionHash = nil
		trace.TransactionPosition = nil
		trace.BlockNumber = nil
		trace.BlockHash = nil
		if trace.Result != nil {
			trace.Result.GasUsed = 0
		}
		// Fixup different revert error text
		if trace.Error != nil && *trace.Error == "Reverted" {
			tmp := "Revert"
			trace.Error = &tmp
		}
		trace.Action.Gas = 0
		newTraces = append(newTraces, trace)
	}

	// Adjust subtrace count for removed frames
	removedIndex := 0
	for traceIndex := 0; traceIndex < len(newTraces); traceIndex++ {
		if removedIndex >= len(removedTraces) {
			break
		}
		traceTarget := removedTraces[removedIndex]
		traceTarget = traceTarget[:len(traceTarget)-1]
		if reflect.DeepEqual(newTraces[traceIndex].TraceAddress, traceTarget) {
			newTraces[traceIndex].Subtraces--
			removedIndex++
		}
	}
	return newTraces
}

func makeTxes(t *testing.T, client ethutils.BasicEthClient, signer types.Signer, privKey *ecdsa.PrivateKey) (common.Hash, common.Hash) {
	meta := arbostestcontracts.SimpleMetaData
	constructorData := hexutil.MustDecode(meta.Bin)
	simpleABI, err := meta.GetAbi()
	test.FailIfError(t, err)
	ctx := context.Background()

	addr := crypto.PubkeyToAddress(privKey.PublicKey)

	startNonce, err := client.NonceAt(ctx, addr, nil)
	test.FailIfError(t, err)
	gasPrice, err := client.SuggestGasPrice(ctx)
	test.FailIfError(t, err)

	makeTx := func(data []byte, to *ethcommon.Address, value *big.Int) (common.Hash, uint64) {
		gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
			From:  addr,
			To:    to,
			Value: value,
			Data:  data,
		})
		test.FailIfError(t, err)

		tx, err := types.SignNewTx(privKey, signer, &types.LegacyTx{
			Nonce:    startNonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       to,
			Value:    value,
			Data:     data,
		})
		test.FailIfError(t, err)
		err = client.SendTransaction(ctx, tx)
		test.FailIfError(t, err)
		startNonce++
		return common.NewHashFromEth(tx.Hash()), startNonce - 1
	}

	txHash1, nonce := makeTx(
		constructorData,
		nil,
		big.NewInt(0),
	)

	connAddress := crypto.CreateAddress(addr, nonce)

	txHash2, _ := makeTx(
		makeFuncData(t, simpleABI.Methods["trace"], big.NewInt(42356)),
		&connAddress,
		big.NewInt(200),
	)
	return txHash1, txHash2
}

func makeFuncData(t *testing.T, funcABI abi.Method, params ...interface{}) []byte {
	txData, err := funcABI.Inputs.Pack(params...)
	test.FailIfError(t, err)
	return append(funcABI.ID, txData...)
}
