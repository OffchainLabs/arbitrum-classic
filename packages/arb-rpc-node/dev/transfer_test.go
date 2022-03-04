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
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func transferTx(t *testing.T, ctx context.Context, nonce uint64, client bind.ContractBackend, dest ethcommon.Address) *types.Transaction {
	head, err := client.HeaderByNumber(ctx, nil)
	test.FailIfError(t, err)
	if head.BaseFee != nil {
		gasTipCap, err := client.SuggestGasTipCap(ctx)
		test.FailIfError(t, err)
		return types.NewTx(&types.DynamicFeeTx{
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: head.BaseFee,
			Gas:       100000,
			To:        &dest,
			Value:     big.NewInt(100),
		})
	} else {
		gasPrice, err := client.SuggestGasPrice(ctx)
		test.FailIfError(t, err)
		return types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      100000,
			To:       &dest,
			Value:    big.NewInt(100),
		})
	}

}

func setupTransferTest(t *testing.T, auth *bind.TransactOpts, client bind.ContractBackend) (*arbostestcontracts.Transfer, ethcommon.Address) {
	ctx := context.Background()
	transferConLogicAddr, _, _, err := arbostestcontracts.DeployTransfer(auth, client)
	test.FailIfError(t, err)
	transferConAddr, _, _, err := ethbridgetestcontracts.DeployTransparentUpgradeableProxy(auth, client, transferConLogicAddr, ethcommon.Address{}, nil)
	test.FailIfError(t, err)
	nonce, err := client.PendingNonceAt(ctx, auth.From)
	test.FailIfError(t, err)
	tx := transferTx(t, ctx, nonce, client, transferConAddr)
	tx, err = auth.Signer(auth.From, tx)
	test.FailIfError(t, err)
	err = client.SendTransaction(ctx, tx)
	test.FailIfError(t, err)
	transferCon, err := arbostestcontracts.NewTransfer(transferConAddr, client)
	test.FailIfError(t, err)
	return transferCon, transferConAddr
}

func getEVMTrace(debugPrints []value.Value) (*evm.EVMTrace, error) {
	var trace *evm.EVMTrace
	for _, debugPrint := range debugPrints {
		ll, err := evm.NewLogLineFromValue(debugPrint)
		if err != nil {
			return nil, err
		}
		if foundTrace, ok := ll.(*evm.EVMTrace); ok {
			if trace != nil {
				return nil, errors.New("found multiple traces")
			}
			trace = foundTrace
		}
	}
	if trace == nil {
		return nil, errors.New("found no trace")
	}
	return trace, nil
}

func TestTransfer(t *testing.T) {
	ctx := context.Background()
	skipBelowVersion(t, 42)
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey, err := crypto.GenerateKey()
	ownerKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)

	auth, owner := OwnerAuthPair(t, ownerKey)

	ethBackend, ethAuths := test.SimulatedBackend(t)

	backend, _, srv, cancelDevNode := NewSimpleTestDevNode(t, config, owner)
	defer cancelDevNode()

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	if doUpgrade {
		UpgradeTestDevNode(t, backend, srv, auth)
	}

	deposit := makeDepositMessage(common.NewAddressFromEth(senderAuth.From))
	_, err = backend.AddInboxMessage(deposit, common.RandAddress())
	test.FailIfError(t, err)

	client := web3.NewEthClient(srv, true)

	ethTransferCon, _ := setupTransferTest(t, ethAuths[0], ethBackend)
	arbTransferCon, arbTransferConAddr := setupTransferTest(t, senderAuth, client)

	// Figure out gas required in stipend
	arbTx, err := arbTransferCon.Send4(senderAuth, big.NewInt(10000))
	test.FailIfError(t, err)
	snap, err := srv.PendingSnapshot(ctx)
	test.FailIfError(t, err)
	_, debugPrints, err := snap.EstimateGas(ctx, arbTx, common.Address{}, common.NewAddressFromEth(senderAuth.From), 100000000)
	test.FailIfError(t, err)
	trace, err := getEVMTrace(debugPrints)
	test.FailIfError(t, err)
	if len(trace.Items) != 8 {
		t.Fatal("unexpected number of items in trace")
	}
	returnItem := trace.Items[5]
	returnTrace, ok := returnItem.(*evm.ReturnTrace)
	if !ok {
		t.Fatal("expected return trace")
	}
	_, err = ethTransferCon.Send3(ethAuths[0])
	test.FailIfError(t, err)
	_, err = arbTransferCon.Send3(senderAuth)
	if err != nil {
		transferABI, err := arbostestcontracts.TransferMetaData.GetAbi()
		test.FailIfError(t, err)
		_, debugPrints, err := snap.Call(ctx, message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(arbTransferConAddr),
				Payment:     big.NewInt(0),
				Data:        transferABI.Methods["send3"].ID,
			},
		}, common.NewAddressFromEth(senderAuth.From), math.MaxUint64)
		test.FailIfError(t, err)
		trace, err := getEVMTrace(debugPrints)
		test.FailIfError(t, err)
		t.Log("call trace", trace)

		if len(trace.Items) < 4 {
			t.Fatal("too few items in trace")
		}
		callItem := trace.Items[2]
		callTrace, ok := callItem.(*evm.CallTrace)
		if !ok {
			t.Fatal("expected call trace")
		}
		t.Log("Gas stipend", callTrace.Gas)
		t.Log("Gas required in stipend", returnTrace.GasUsed)
		t.Fatal("insufficient stipend")
	}
	test.FailIfError(t, err)
}
