package dev

import (
	// "encoding/hex"
	//	"encoding/json"
	//	"fmt"
	//	"math/big"
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

	t.Log("Onestep")
	callRes, err := ethServer.Call(getXTxArgs, block, nil)
	test.FailIfError(t, err)

	if ethcommon.BytesToHash(callRes) != ethcommon.HexToHash("0x100") {
		t.Fatal("Unexpected return val")
	}

}
