package ethbridge

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Assertion struct {
}

func (a *Assertion) BytesFields() [7][32]byte {
	return [7][32]byte{}
}

func (a *Assertion) IntFields() [11]*big.Int {
	return [11]*big.Int{}
}

func (a *Assertion) InboxConsistencyHash(inboxMaxHash common.Hash, inboxMaxCount *big.Int) common.Hash {
	return common.Hash{}
}

func (a *Assertion) InboxDeltaHash() common.Hash {
	return common.Hash{}
}

func (a *Assertion) ExecutionHash() common.Hash {
	return common.Hash{}
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return big.NewInt(0)
}

type NodeID *big.Int

type RollupWatcher struct {
	con *ethbridgecontracts.Rollup

	arbGasSpeedLimitPerBlock *big.Int
}

func NewRollupWatcher(address ethcommon.Address, client ethutils.EthClient) (*RollupWatcher, error) {
	con, err := ethbridgecontracts.NewRollup(address, client)
	if err != nil {
		return nil, err
	}
	arbGasSpeedLimitPerBlock, err := con.ArbGasSpeedLimitPerBlock(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return &RollupWatcher{
		con:                      con,
		arbGasSpeedLimitPerBlock: arbGasSpeedLimitPerBlock,
	}, nil
}

func (r *RollupWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	return r.con.StakerCount(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) CurrentRequiredStake(ctx context.Context) (*big.Int, error) {
	return r.con.CurrentRequiredStake(&bind.CallOpts{Context: ctx})
}

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

func (r *Rollup) RejectNextNode(ctx context.Context, node NodeID, staker common.Address) (*types.Transaction, error) {
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
	messages [][]byte,
) (*types.Transaction, error) {
	var messageData []byte
	messageLengths := make([]*big.Int, 0, len(messages))
	for _, msg := range messages {
		messageData = append(messageData, msg...)
		messageLengths = append(messageLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ConfirmNextNode(auth, logAcc, messageData, messageLengths)
	})
}

func (r *Rollup) NewStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
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
	node NodeID,
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
	node NodeID,
	prev NodeID,
	assertion *Assertion,
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
	node NodeID,
	assertion *Assertion,
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
	node1 NodeID,
	staker2 common.Address,
	node2 NodeID,
	assertion *Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
) (*types.Transaction, error) {
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
			assertion.CheckTime(r.arbGasSpeedLimitPerBlock),
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
