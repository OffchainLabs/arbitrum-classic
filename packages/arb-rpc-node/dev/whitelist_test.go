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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/metrics"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func TestWhitelist(t *testing.T) {
	skipBelowVersion(t, 25)
	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey, err := crypto.GenerateKey()
	ownerKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	owner := crypto.PubkeyToAddress(ownerKey.PublicKey)

	backend, _, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, common.NewAddressFromEth(owner), nil)
	defer cancelDevNode()

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(ownerKey, backend.chainID)
	test.FailIfError(t, err)

	client := web3.NewEthClient(srv, true, metrics.NewMetricsConfig(nil))

	_, _, simple, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)

	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)

	_, err = arbOwner.RemoveAllowedSender(ownerAuth, senderAuth.From)
	test.FailIfError(t, err)

	_, err = simple.Exists(senderAuth)
	if err == nil {
		t.Error()
	}

	_, err = arbOwner.AddAllowedSender(ownerAuth, senderAuth.From)
	test.FailIfError(t, err)

	_, err = simple.Exists(senderAuth)
	test.FailIfError(t, err)

	_, err = arbOwner.AllowOnlyOwnerToSend(ownerAuth)
	test.FailIfError(t, err)

	_, err = simple.Exists(senderAuth)
	if err == nil {
		t.Error()
	}

	_, err = arbOwner.AddAllowedSender(ownerAuth, senderAuth.From)
	test.FailIfError(t, err)

	_, err = simple.Exists(senderAuth)
	test.FailIfError(t, err)
}
