package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"math/big"
)

type NodeWatcher struct {
	con *ethbridgecontracts.INode
}

func NewNodeWatcher(address ethcommon.Address, client ethutils.EthClient) (*NodeWatcher, error) {
	con, err := ethbridgecontracts.NewINode(address, client)
	if err != nil {
		return nil, err
	}

	return &NodeWatcher{
		con: con,
	}, nil
}

func (n *NodeWatcher) Prev(ctx context.Context) (*big.Int, error) {
	return n.con.Prev(&bind.CallOpts{Context: ctx})
}

func (n *NodeWatcher) DeadlineBlock(ctx context.Context) (*big.Int, error) {
	return n.con.DeadlineBlock(&bind.CallOpts{Context: ctx})
}

func (n *NodeWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	return n.con.StakerCount(&bind.CallOpts{Context: ctx})
}
