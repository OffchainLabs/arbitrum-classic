
package dev

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestZeroPriceGasEstimation(t *testing.T) {
	ctx := context.Background()
	_, _, client, _, aggAuth, _, _, _, cancel := setupFeeChain(t)
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
