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
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func expectHex(t *testing.T, result []byte, resultErr error, expected string) {
	t.Helper()
	test.FailIfError(t, resultErr)
	if (ethcommon.BytesToHash(result) != ethcommon.HexToHash(expected)) || (len(result) != 32) {
		t.Fatal("Expected ", expected, "Got", ethcommon.BytesToHash(result), " [", len(result), "]")
	}
}

func TestEthCall(t *testing.T) {
	//skipBelowVersion(t, 46)
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	auth, owner := OptsAddressPair(t, nil)
	senderKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	ctx := context.Background()

	backend, _, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, owner, nil)
	defer cancelDevNode()

	if doUpgrade {
		UpgradeTestDevNode(t, backend, srv, auth)
		enableRewrites(t, backend, srv, auth)
	}

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	ethServer := web3.NewServer(srv, web3.DefaultConfig, nil)

	client := web3.NewEthClient(srv, true)

	testerAddr, _, _, err := arbostestcontracts.DeployEthCallTester(senderAuth, client)
	test.FailIfError(t, err)

	testerABI, err := arbostestcontracts.EthCallTesterMetaData.GetAbi()
	test.FailIfError(t, err)

	rpcLatest := rpc.LatestBlockNumber
	block := rpc.BlockNumberOrHash{BlockNumber: &rpcLatest}

	getXdata := testerABI.Methods["getX"].ID
	getXTxArgs := web3.CallTxArgs{
		To:   &testerAddr,
		Data: (*hexutil.Bytes)(&getXdata),
	}

	getBalancedata := testerABI.Methods["getBalance"].ID
	getBalanceTxArgs := web3.CallTxArgs{
		To:   &testerAddr,
		Data: (*hexutil.Bytes)(&getBalancedata),
	}

	sloadArgs, err := testerABI.Methods["sLoad"].Inputs.Pack(big.NewInt(0x100))
	test.FailIfError(t, err)
	sloadData := append(testerABI.Methods["sLoad"].ID, sloadArgs...)
	sloadTxArgs := web3.CallTxArgs{
		To:   &testerAddr,
		Data: (*hexutil.Bytes)(&sloadData),
	}

	t.Log("No Overrides")
	callRes, err := ethServer.Call(ctx, getXTxArgs, block, nil)
	expectHex(t, callRes, err, "0x100")

	callRes, err = ethServer.Call(ctx, sloadTxArgs, block, nil)
	expectHex(t, callRes, err, "0x0")

	callRes, err = ethServer.Call(ctx, getBalanceTxArgs, block, nil)
	expectHex(t, callRes, err, "0x0")

	t.Log("Override Balance")
	overrideMap := make(map[ethcommon.Address]snapshot.EthCallOverride)
	overrideMap[testerAddr] = snapshot.EthCallOverride{
		Balance: (*hexutil.Big)(hexutil.MustDecodeBig("0x3000")),
	}
	callRes, err = ethServer.Call(ctx, sloadTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x0")

	callRes, err = ethServer.Call(ctx, getBalanceTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x3000")

	t.Log("StateDiff")
	stateMap := make(map[ethcommon.Hash]ethcommon.Hash)
	stateMap[ethcommon.HexToHash("0x0")] = ethcommon.HexToHash("0x10")
	stateMap[ethcommon.HexToHash("0x100")] = ethcommon.HexToHash("0x90")
	overrideMap = make(map[ethcommon.Address]snapshot.EthCallOverride)
	overrideMap[testerAddr] = snapshot.EthCallOverride{
		StateDiff: &stateMap,
	}

	callRes, err = ethServer.Call(ctx, sloadTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x90")

	callRes, err = ethServer.Call(ctx, getXTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x10")

	t.Log("State")
	overrideMap = make(map[ethcommon.Address]snapshot.EthCallOverride)
	overrideMap[testerAddr] = snapshot.EthCallOverride{
		State: &stateMap,
	}

	callRes, err = ethServer.Call(ctx, sloadTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x90")

	callRes, err = ethServer.Call(ctx, getXTxArgs, block, &overrideMap)
	expectHex(t, callRes, err, "0x10")

	t.Log("Code over empty")
	newContractAddr := ethcommon.HexToAddress("0x3000000000")
	// code translates to: storage[1] = storage[0] + 0x10, return (storage[1])
	code, err := hex.DecodeString("6000546010018060015560005260206000f3")
	test.FailIfError(t, err)
	noData := make(hexutil.Bytes, 0)
	newContractTxArgs := web3.CallTxArgs{
		To:   &newContractAddr,
		Data: &noData,
	}
	codeOverride := make(map[ethcommon.Address]snapshot.EthCallOverride)
	codeOverride[newContractAddr] = snapshot.EthCallOverride{
		Code: (*hexutil.Bytes)(&code),
	}
	callRes, err = ethServer.Call(ctx, newContractTxArgs, block, &codeOverride)
	expectHex(t, callRes, err, "0x10")

	t.Log("Code + state over existing")
	codeOverride[testerAddr] = snapshot.EthCallOverride{
		Code:  (*hexutil.Bytes)(&code),
		State: &stateMap,
	}
	overrideTxArgs := web3.CallTxArgs{
		To:   &testerAddr,
		Data: &noData,
	}
	callRes, err = ethServer.Call(ctx, overrideTxArgs, block, &codeOverride)
	expectHex(t, callRes, err, "0x20")

	t.Log("Nonce")
	arbostest, err := abi.JSON(strings.NewReader(arboscontracts.ArbosTestABI))
	test.FailIfError(t, err)
	arbosGetAccountInfo := arbostest.Methods["getAccountInfo"]
	getInfoArgs, err := arbosGetAccountInfo.Inputs.Pack(newContractAddr)
	test.FailIfError(t, err)
	getInfoData := append(arbosGetAccountInfo.ID, getInfoArgs...)
	getInfoTxArgs := web3.CallTxArgs{
		To:   &arbos.ARB_TEST_ADDRESS,
		Data: (*hexutil.Bytes)(&getInfoData),
	}
	nonceToTest := uint64(0x60)
	codeOverride[newContractAddr] = snapshot.EthCallOverride{
		Nonce: (*hexutil.Uint64)(&nonceToTest),
	}
	callRes, err = ethServer.Call(ctx, getInfoTxArgs, block, &codeOverride)
	//returns some information.. 2nd uint256 is nonce for the request
	expectHex(t, callRes[32:64], err, "0x60")

}
