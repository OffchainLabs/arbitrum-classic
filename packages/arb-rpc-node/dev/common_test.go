/*
* Copyright 2020, Offchain Labs, Inc.
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
	"encoding/json"
	"flag"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"os"
	"testing"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

var arbosfile *string
var arbosVersion int
var doUpgrade bool
var didUpgrade bool

type ArbOSExec struct {
	Version *int `json:"arbos_version"`
}

func TestMain(m *testing.M) {

	parseDoUpgrade := flag.Bool("upgrade", true, "Test against an upgraded ArbOS. Overrides 'arbos' flag.")
	
	arbosPath, err := arbos.Path(doUpgrade)
	if err != nil {
		panic(err)
	}

	if doUpgrade {
		arbosfile = &arbosPath
	} else {
		arbosfile = flag.String("arbos", arbosPath, "version of arbos to run tests against")
	}
	flag.Parse()

	doUpgrade = *parseDoUpgrade
	didUpgrade = false

	fileData, err := ioutil.ReadFile(*arbosfile)
	if err != nil {
		panic(err)
	}
	var arbosExec ArbOSExec
	if err := json.Unmarshal(fileData, &arbosExec); err != nil {
		panic(err)
	}
	if arbosExec.Version != nil {
		arbosVersion = *arbosExec.Version
	} else {
		arbosVersion = 1
	}
	exitcode := m.Run()
	os.Exit(exitcode)
}

func skipBelowVersion(t *testing.T, ver int) {
	t.Helper()
	if arbosVersion < ver {
		t.Skipf("Skipping test because version %v too below supported version %v", arbosVersion, ver)
	}
}

func NewTestDevNode(
	t *testing.T,
	arbosPath string,
	params protocol.ChainParams,
	owner common.Address,
	config []message.ChainConfigOption,
) (*Backend, *txdb.TxDB, *aggregator.Server, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	agg := common.RandAddress()
	chainId := big.NewInt(42161)
	for i := range config {
		opt := config[len(config)-1-i]
		if aggConfig, ok := opt.(message.DefaultAggConfig); ok {
			agg = aggConfig.Aggregator
			break
		}
		if chainIdConfig, ok := opt.(message.ChainIDConfig); ok {
			chainId = chainIdConfig.ChainId
			break
		}
	}
	backend, db, cancelDevNode, txDBErrChan, err := NewDevNode(
		ctx,
		t.TempDir(),
		arbosPath,
		chainId,
		agg,
		0,
	)
	test.FailIfError(t, err)
	initMsg, err := message.NewInitMessage(params, owner, config)
		
	test.FailIfError(t, err)
	_, err = backend.AddInboxMessage(initMsg, common.Address{})
	test.FailIfError(t, err)

	go func() {
		if err := <-txDBErrChan; err != nil {
			t.Error(err)
			cancel()
		}
	}()

	closeFunc := func() {
		if doUpgrade && !didUpgrade {
			t.Fatal("Test was supposed to perform an upgrade but never did")
		}
		cancelDevNode()
		cancel()
	}
	srv := aggregator.NewServer(backend, common.Address{}, chainId, db)
	return backend, db, srv, closeFunc
}

func UpgradeTestDevNode(t *testing.T, backend *Backend, srv *aggregator.Server, auth *bind.TransactOpts) {
	arbosDir, err := arbos.Dir()
	test.FailIfError(t, err)

	upgradedMach, err := cmachine.New(filepath.Join(arbosDir, "arbos-upgrade.mexe"))
	test.FailIfError(t, err)
	targetHash := upgradedMach.CodePointHash()

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

	updateBytes, err := ioutil.ReadFile(filepath.Join(arbosDir, "upgrade.json"))
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

	t.Log("Finishing code upload")
	_, err = arbOwner.FinishCodeUploadAsArbosUpgrade(auth, codeHash, common.Hash{})
	test.FailIfError(t, err)
	auth.GasLimit = 0

	t.Log("Upgraded! Testing new version")
	_, err = simpleCon.Exists(auth)
	test.FailIfError(t, err)

	newVersion, err := arbSys.ArbOSVersion(&bind.CallOpts{})
	test.FailIfError(t, err)
	t.Log("New Version:", newVersion)
	didUpgrade = true
}

func OwnerAuthPair(t *testing.T, key *ecdsa.PrivateKey) (*bind.TransactOpts, common.Address) {
	
	if key == nil {
		random, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		key = random
	}
	
	auth := bind.NewKeyedTransactor(key)
	address := common.NewAddressFromEth(auth.From)
	return auth, address
}
