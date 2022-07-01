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
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestOutOfGas(t *testing.T) {
	ctx := context.Background()
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)

	upgraderAuth, upgraderAccount := OptsAddressPair(t, nil)

	backend, _, srv, cancelDevNode := NewSimpleTestDevNode(t, config, upgraderAccount)
	defer cancelDevNode()

	auth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	client := web3.NewEthClient(srv, true)

	// Basic Tx
	_, _, transfer, err := arbostestcontracts.DeployTransfer(auth, client)
	test.FailIfError(t, err)

	if doUpgrade {
		UpgradeTestDevNode(t, ctx, backend, srv, upgraderAuth)
	}

	tx, err := transfer.Receive(auth)
	test.FailIfError(t, err)

	auth.GasLimit = tx.Gas()
	_, err = transfer.Spin(auth)
	if err.Error() != "execution ran out of gas" {
		t.Error("Unexpected error from spinning:", err)
	}
}
