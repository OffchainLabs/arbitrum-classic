package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Rollup struct {
	*RollupWatcher
	auth *TransactAuth
}

func NewRollup(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Rollup, error) {
	watcher, err := NewRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		RollupWatcher: watcher,
		auth:          auth,
	}, nil
}

func (r *Rollup) RejectNextNode(ctx context.Context, node core.NodeID, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNode(auth, node, staker.ToEthAddress())
	})
}

func (r *Rollup) RejectNextNodeOutOfOrder(ctx context.Context) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNodeOutOfOrder(auth)
	})
}

func (r *Rollup) ConfirmNextNode(
	ctx context.Context,
	logAcc common.Hash,
	sends [][]byte,
) (*types.Transaction, error) {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ConfirmNextNode(auth, logAcc, sendsData, sendLengths)
	})
}

func (r *Rollup) NewStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node core.NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) AddStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node core.NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) NewStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node core.NodeID,
	prev core.NodeID,
	assertion *core.Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			prev,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) AddStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node core.NodeID,
	assertion *core.Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) ReturnOldDeposit(ctx context.Context, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReturnOldDeposit(auth, staker.ToEthAddress())
	})
}

func (r *Rollup) AddToDeposit(ctx context.Context, address common.Address, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = amount
		return r.con.AddToDeposit(auth, address.ToEthAddress())
	})
}

func (r *Rollup) ReduceDeposit(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReduceDeposit(auth, amount)
	})
}

func (r *Rollup) CreateChallenge(
	ctx context.Context,
	staker1 common.Address,
	node1 core.NodeID,
	staker2 common.Address,
	node2 core.NodeID,
	assertion *core.Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
) (*types.Transaction, error) {
	speedLimit, err := r.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.CreateChallenge(
			auth,
			staker1.ToEthAddress(),
			node1,
			staker2.ToEthAddress(),
			node2,
			assertion.InboxConsistencyHash(inboxMaxHash, inboxMaxCount),
			assertion.InboxDeltaHash(),
			assertion.ExecutionHash(),
			assertion.CheckTime(speedLimit),
		)
	})
}

func (r *Rollup) RemoveZombie(ctx context.Context, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveZombie(auth, zombieNum, maxNodes)
	})
}

func (r *Rollup) RemoveOldZombies(ctx context.Context, startIndex *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveOldZombies(auth, startIndex)
	})
}
