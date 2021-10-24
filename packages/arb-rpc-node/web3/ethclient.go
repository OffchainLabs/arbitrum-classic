package web3

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func NewEthClient(srv *aggregator.Server, ganacheMode bool) (*ethclient.Client, error) {
	var mode RpcMode
	if ganacheMode {
		mode = GanacheMode
	} else {
		mode = NormalMode
	}
	rpcServer, err := GenerateWeb3Server(srv, nil, mode, nil)
	if err != nil {
		panic(err)
	}
	return ethclient.NewClient(rpc.DialInProc(rpcServer)), nil
}

func NewTestEthClient(t *testing.T, srv *aggregator.Server, ganacheMode bool) *ethclient.Client {
	clnt, err := NewEthClient(srv, ganacheMode)
	test.FailIfError(t, err)
	return clnt
}
