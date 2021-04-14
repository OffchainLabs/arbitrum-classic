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
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
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

type upgrade struct {
	Instructions []string `json:"instructions"`
}

func TestUpgrade(t *testing.T) {
	skipBelowVersion(t, 4)

	upgradedMach, err := cmachine.New(filepath.Join(arbos.Dir(), "arbos-upgrade.mexe"))
	test.FailIfError(t, err)
	targetHash := upgradedMach.CodePointHash()

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

	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	arbosFile := filepath.Join(arbos.Dir(), "arbos_before.mexe")
	monitor, backend, db, rollupAddress := NewDevNode(tmpDir, arbosFile, config, common.NewAddressFromEth(auth.From), nil)
	defer monitor.Close()
	defer db.Close()

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

	srv := aggregator.NewServer(backend, rollupAddress, db)
	client := web3.NewEthClient(srv, true)
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	test.FailIfError(t, err)

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

	updateBytes, err := ioutil.ReadFile(filepath.Join(arbos.Dir(), "upgrade.json"))
	test.FailIfError(t, err)

	upgrade := upgrade{}
	err = json.Unmarshal(updateBytes, &upgrade)
	test.FailIfError(t, err)
	chunkSize := 100000
	chunks := []string{"0x"}
	for _, insn := range upgrade.Instructions {
		if len(chunks[len(chunks)-1])+len(insn) > chunkSize {
			chunks = append(chunks, "0x")
		}
		chunks[len(chunks)-1] += insn
	}

	auth.GasLimit = 10000000000
	_, err = arbOwner.StartCodeUpload(auth)
	test.FailIfError(t, err)

	for i, upgradeChunk := range chunks {
		t.Log("Upgrade chunk", i)
		_, err = arbOwner.ContinueCodeUpload(auth, hexutil.MustDecode(upgradeChunk))
		test.FailIfError(t, err)
	}

	codeHash, err := arbOwner.GetUploadedCodeHash(&bind.CallOpts{})
	test.FailIfError(t, err)

	if codeHash != targetHash {
		t.Fatal("uploaded codehash was incorrect after 1st upgrade")
	}

	_, err = arbOwner.FinishCodeUploadAsArbosUpgrade(auth, codeHash)
	test.FailIfError(t, err)
	auth.GasLimit = 0

	_, err = simpleCon.Exists(auth)
	test.FailIfError(t, err)

	newVersion, err := arbSys.ArbOSVersion(&bind.CallOpts{})
	test.FailIfError(t, err)

	t.Log("New Version:", newVersion)
	//if newVersion.Cmp(oldVersion) <= 0 {
	//	t.Error("didn't change to new version")
	//}
}
