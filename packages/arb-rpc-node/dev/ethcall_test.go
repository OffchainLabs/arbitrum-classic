package dev

import (
	// "encoding/hex"
	//	"encoding/json"
	// "fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestEthCall(t *testing.T) {
	//skipBelowVersion(t, 46)
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)

	ownerKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	owner := crypto.PubkeyToAddress(ownerKey.PublicKey)

	backend, _, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, common.NewAddressFromEth(owner), nil)
	defer cancelDevNode()

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	ethServer := web3.NewServer(srv, false)

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
	callRes, err := ethServer.Call(getXTxArgs, block, nil)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x100") {
		t.Fatal("Unexpected return val")
	}

	callRes, err = ethServer.Call(sloadTxArgs, block, nil)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x0") {
		t.Fatal("Unexpected return val")
	}

	callRes, err = ethServer.Call(getBalanceTxArgs, block, nil)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x0") {
		t.Fatal("Unexpected Balance")
	}

	t.Log("Override Balance")
	overrideMap := make(map[ethcommon.Address]web3.EthCallOverride)
	overrideMap[testerAddr] = web3.EthCallOverride{
		Balance: (*hexutil.Big)(hexutil.MustDecodeBig("0x3000")),
	}
	callRes, err = ethServer.Call(sloadTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x0") {
		t.Fatal("Unexpected return val")
	}

	callRes, err = ethServer.Call(getBalanceTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x3000") {
		t.Fatal("Unexpected Balance")
	}

	t.Log("StateDiff")
	stateMap := make(map[ethcommon.Hash]ethcommon.Hash)
	stateMap[ethcommon.HexToHash("0x0")] = ethcommon.HexToHash("0x10")
	stateMap[ethcommon.HexToHash("0x100")] = ethcommon.HexToHash("0x90")
	overrideMap = make(map[ethcommon.Address]web3.EthCallOverride)
	overrideMap[testerAddr] = web3.EthCallOverride{
		StateDiff: &stateMap,
	}

	callRes, err = ethServer.Call(sloadTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x90") {
		t.Fatal("Unexpected return val")
	}

	callRes, err = ethServer.Call(getXTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x10") {
		t.Fatal("Unexpected return val")
	}

	t.Log("State")
	overrideMap = make(map[ethcommon.Address]web3.EthCallOverride)
	overrideMap[testerAddr] = web3.EthCallOverride{
		State: &stateMap,
	}

	callRes, err = ethServer.Call(sloadTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x90") {
		t.Fatal("Unexpected return val")
	}

	callRes, err = ethServer.Call(getXTxArgs, block, &overrideMap)
	test.FailIfError(t, err)
	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x10") {
		t.Fatal("Unexpected return val")
	}

}
