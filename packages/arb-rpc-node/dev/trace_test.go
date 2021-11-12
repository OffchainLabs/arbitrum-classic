package dev

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestTrace(t *testing.T) {
	skipBelowVersion(t, 25)
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
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

	ethServer := web3.NewServer(srv, web3.DefaultConfig, nil)
	tracer := web3.NewTracer(ethServer)

	client := web3.NewEthClient(srv, true)

	simpleAddr, _, simple, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)

	tx, err := simple.Trace(senderAuth, big.NewInt(42356))
	test.FailIfError(t, err)

	blockNum := rpc.LatestBlockNumber
	data := hexutil.Bytes(tx.Data())
	traceData, err := tracer.Call(web3.CallTxArgs{
		From: &senderAuth.From,
		To:   &simpleAddr,
		Data: &data,
	}, []string{"trace"}, rpc.BlockNumberOrHash{BlockNumber: &blockNum})
	test.FailIfError(t, err)

	jsonData, err := json.MarshalIndent(traceData, "", " ")
	test.FailIfError(t, err)
	fmt.Println("data", string(jsonData))
}
