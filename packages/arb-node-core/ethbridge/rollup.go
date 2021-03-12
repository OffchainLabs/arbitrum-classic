package ethbridge

import (
	"context"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type RawTransaction struct {
	Data   []byte
	Dest   ethcommon.Address
	Amount *big.Int
}

type Rollup struct {
	*RollupWatcher
	*BuilderBackend
	builderCon *ethbridgecontracts.Rollup
}

func NewRollup(address ethcommon.Address, client ethutils.EthClient, builder *BuilderBackend) (*Rollup, error) {
	builderCon, err := ethbridgecontracts.NewRollup(address, builder)
	if err != nil {
		return nil, err
	}
	watcher, err := NewRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		RollupWatcher:  watcher,
		BuilderBackend: builder,
		builderCon:     builderCon,
	}, nil
}

func (r *Rollup) RejectNextNode(ctx context.Context, staker common.Address) error {
	_, err := r.builderCon.RejectNextNode(authWithContext(ctx, r.builderAuth), staker.ToEthAddress())
	return err
}

func (r *Rollup) ConfirmNextNode(ctx context.Context, assertion *core.Assertion, sends [][]byte) error {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}

	_, err := r.builderCon.ConfirmNextNode(
		authWithContext(ctx, r.builderAuth),
		assertion.Before.SendAcc,
		sendsData,
		sendLengths,
		assertion.After.TotalSendCount,
		assertion.After.LogAcc,
		assertion.After.TotalLogCount,
	)
	return err
}

func (r *Rollup) NewStake(ctx context.Context, amount *big.Int) error {
	tokenType, err := r.StakeToken(ctx)
	if err != nil {
		return err
	}
	emptyAddress := common.Address{}
	if tokenType != emptyAddress {
		_, err := r.builderCon.NewStake(authWithContext(ctx, r.builderAuth), amount)
		return err
	} else {
		_, err := r.builderCon.NewStake(authWithContextAndAmount(ctx, r.builderAuth, amount), big.NewInt(0))
		return err
	}
}

func (r *Rollup) StakeOnExistingNode(ctx context.Context, nodeNumber core.NodeID, nodeHash [32]byte) error {
	_, err := r.builderCon.StakeOnExistingNode(
		authWithContext(ctx, r.builderAuth),
		nodeNumber,
		nodeHash,
	)
	return err
}

func (r *Rollup) StakeOnNewNode(
	ctx context.Context,
	nodeHash [32]byte,
	assertion *core.Assertion,
	prevProposedBlock *big.Int,
	prevInboxMaxCount *big.Int,
) error {
	_, err := r.builderCon.StakeOnNewNode(
		authWithContext(ctx, r.builderAuth),
		nodeHash,
		assertion.BytesFields(),
		assertion.IntFields(),
		prevProposedBlock,
		prevInboxMaxCount,
	)
	return err
}

func (r *Rollup) ReturnOldDeposit(ctx context.Context, staker common.Address) error {
	_, err := r.builderCon.ReturnOldDeposit(authWithContext(ctx, r.builderAuth), staker.ToEthAddress())
	return err
}

func (r *Rollup) AddToDeposit(ctx context.Context, address common.Address, amount *big.Int) error {
	tokenType, err := r.StakeToken(ctx)
	if err != nil {
		return err
	}
	emptyAddress := common.Address{}
	if tokenType != emptyAddress {
		_, err := r.builderCon.AddToDeposit(
			authWithContext(ctx, r.builderAuth),
			address.ToEthAddress(),
			amount,
		)
		return err
	} else {
		_, err := r.builderCon.AddToDeposit(
			authWithContextAndAmount(ctx, r.builderAuth, amount),
			address.ToEthAddress(),
			big.NewInt(0),
		)
		return err
	}
}

func (r *Rollup) ReduceDeposit(ctx context.Context, amount *big.Int) error {
	_, err := r.builderCon.ReduceDeposit(authWithContext(ctx, r.builderAuth), amount)
	return err
}

func (r *Rollup) CreateChallenge(
	ctx context.Context,
	staker1 common.Address,
	node1 *core.NodeInfo,
	staker2 common.Address,
	node2 *core.NodeInfo,
) error {
	_, err := r.builderCon.CreateChallenge(
		authWithContext(ctx, r.builderAuth),
		[2]ethcommon.Address{staker1.ToEthAddress(), staker2.ToEthAddress()},
		[2]*big.Int{node1.NodeNum, node2.NodeNum},
		[2][32]byte{
			node1.Assertion.ExecutionHash(),
			node2.Assertion.ExecutionHash(),
		},
		[2]*big.Int{
			node1.BlockProposed.Height.AsInt(),
			node2.BlockProposed.Height.AsInt(),
		},
		[2]*big.Int{
			node1.Assertion.After.TotalMessagesRead,
			node2.Assertion.After.TotalMessagesRead,
		},
	)
	return err
}

func (r *Rollup) RemoveZombie(ctx context.Context, zombieNum *big.Int, maxNodes *big.Int) error {
	_, err := r.builderCon.RemoveZombie(authWithContext(ctx, r.builderAuth), zombieNum, maxNodes)
	return err
}

func (r *Rollup) RemoveOldZombies(ctx context.Context, startIndex *big.Int) error {
	_, err := r.builderCon.RemoveOldZombies(authWithContext(ctx, r.builderAuth), startIndex)
	return err
}
