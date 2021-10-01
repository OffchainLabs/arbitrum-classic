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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

type upgrade struct {
	Instructions []string `json:"instructions"`
}

func TestUpgrade(t *testing.T) {
	skipBelowVersion(t, 4)

	arbosFile, _ := arbos.Path(true)

	privkey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	auth, owner := OwnerAuthPair(t, privkey)

	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	backend, _, srv, cancelDevNode := NewTestDevNode(t, arbosFile, config, owner, nil)
	defer cancelDevNode()

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(auth.From),
				Payment:     big.NewInt(100),
				Data:        nil,
			},
		}),
	}
	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}

	client := web3.NewEthClient(srv, true)

	arbSys, err := arboscontracts.NewArbSys(arbos.ARB_SYS_ADDRESS, client)
	test.FailIfError(t, err)

	oldVersion, err := arbSys.ArbOSVersion(&bind.CallOpts{})
	test.FailIfError(t, err)

	t.Log("Old Version:", oldVersion)

	_, _, simpleCon, err := arbostestcontracts.DeploySimple(auth, client)
	test.FailIfError(t, err)

	_, err = simpleCon.Exists(auth)
	test.FailIfError(t, err)

	auth.Value = big.NewInt(1)
	_, err = simpleCon.RejectPayment(auth)
	if err == nil {
		t.Fatal("tx should have failed")
	}
	auth.Value = big.NewInt(0)

	UpgradeTestDevNode(t, backend, srv, auth)

	_, err = simpleCon.Exists(auth)
	test.FailIfError(t, err)

	// Try to start a new upgrade to make sure the owner auth still works
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)
	auth.GasLimit = 10000000000
	_, err = arbOwner.StartCodeUpload(auth)
	test.FailIfError(t, err)

	newVersion, err := arbSys.ArbOSVersion(&bind.CallOpts{})
	test.FailIfError(t, err)

	t.Log("New Version:", newVersion)
	if newVersion.Cmp(oldVersion) <= 0 {
		t.Error("didn't change to new version")
	}
}
