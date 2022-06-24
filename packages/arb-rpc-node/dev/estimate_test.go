package dev

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestZeroPriceGasEstimationEmpty(t *testing.T) {
	ctx := context.Background()
	_, _, client, _, aggAuth, _, _, _, cancel := setupFeeChain(t, ctx)
	defer cancel()

	gas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  common.RandAddress().ToEthAddress(),
		To:    &aggAuth.From,
		Value: big.NewInt(0),
		Data:  []byte{},
	})
	test.FailIfError(t, err)

	if gas == 0 {
		t.Fatal("EstimateGas returned 0")
	}
}

func TestZeroPriceGasEstimationCall(t *testing.T) {
	ctx := context.Background()
	backend, _, client, auth, _, _, _, _, cancel := setupFeeChain(t, ctx)
	defer cancel()

	simpleAddr, _, _, err := arbostestcontracts.DeploySimple(auth, client)
	test.FailIfError(t, err)

	fundedAddr := addSomeBalance(t, ctx, common.RandAddress(), backend, client)
	unfundedAddr := common.RandAddress()

	fundedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  fundedAddr.ToEthAddress(),
		To:    &simpleAddr,
		Value: big.NewInt(0),
		Data:  []byte{0x26, 0x7c, 0x4a, 0xe4}, // call exists()
	})
	test.FailIfError(t, err)

	unfundedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  unfundedAddr.ToEthAddress(),
		To:    &simpleAddr,
		Value: big.NewInt(0),
		Data:  []byte{0x26, 0x7c, 0x4a, 0xe4}, // call exists()
	})
	test.FailIfError(t, err)

	if fundedGas == 0 || unfundedGas == 0 {
		t.Fatal("EstimateGas returned 0")
	}

	if math.Abs(float64(fundedGas)-float64(unfundedGas)) > 100 {
		t.Fatal("EstimateGas depends on balance", fundedGas, unfundedGas)
	}
}

func TestZeroPriceGasEstimationDeploy(t *testing.T) {
	ctx := context.Background()
	_, _, client, auth, _, _, _, _, cancel := setupFeeChain(t, ctx)
	defer cancel()

	fundedAddr := auth.From
	unfundedAddr := common.RandAddress().ToEthAddress()

	fundedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From: fundedAddr,
		Data: hexutil.MustDecode(conData),
	})
	test.FailIfError(t, err)

	unfundedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From: unfundedAddr,
		Data: hexutil.MustDecode(conData),
	})
	test.FailIfError(t, err)

	if fundedGas == 0 || unfundedGas == 0 {
		t.Fatal("EstimateGas returned 0")
	}

	if math.Abs(float64(fundedGas)-float64(unfundedGas)) > 100 {
		t.Fatal("EstimateGas depends on balance", fundedGas, unfundedGas)
	}
}
