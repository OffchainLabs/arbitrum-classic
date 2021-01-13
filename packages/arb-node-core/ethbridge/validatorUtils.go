package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type ConfirmType uint8

const (
	CONFIRM_TYPE_NONE ConfirmType = iota
	CONFIRM_TYPE_VALID
	CONFIRM_TYPE_INVALID
)

type ConflictType uint8

const (
	CONFLICT_TYPE_NONE ConflictType = iota
	CONFLICT_TYPE_FOUND
	CONFLICT_TYPE_INDETERMINATE
	CONFLICT_TYPE_INCOMPLETE
)

type ValidatorUtilsWatcher struct {
	con           *ethbridgecontracts.ValidatorUtils
	rollupAddress ethcommon.Address
}

func NewValidatorUtilsWatcher(address, rollupAddress ethcommon.Address, client ethutils.EthClient) (*ValidatorUtilsWatcher, error) {
	con, err := ethbridgecontracts.NewValidatorUtils(address, client)
	if err != nil {
		return nil, err
	}

	return &ValidatorUtilsWatcher{
		con:           con,
		rollupAddress: rollupAddress,
	}, nil
}

func (v *ValidatorUtilsWatcher) RefundableStakers(ctx context.Context) ([]common.Address, error) {
	addresses, err := v.con.RefundableStakers(&bind.CallOpts{Context: ctx}, v.rollupAddress)
	if err != nil {
		return nil, err
	}
	return common.AddressArrayFromEth(addresses), nil
}

type RollupConfig struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}

func (v *ValidatorUtilsWatcher) GetConfig(ctx context.Context) (*RollupConfig, error) {
	config, err := v.con.GetConfig(&bind.CallOpts{Context: ctx}, v.rollupAddress)
	if err != nil {
		return nil, err
	}
	return &RollupConfig{
		ChallengePeriodBlocks:    config.ChallengePeriodBlocks,
		ArbGasSpeedLimitPerBlock: config.ArbGasSpeedLimitPerBlock,
		BaseStake:                config.BaseStake,
		StakeToken:               common.NewAddressFromEth(config.StakeToken),
	}, nil
}

func (v *ValidatorUtilsWatcher) SuccessorNodes(ctx context.Context, node core.NodeID) ([]*big.Int, error) {
	return v.con.SuccessorNodes(&bind.CallOpts{Context: ctx}, v.rollupAddress, node)
}

func (v *ValidatorUtilsWatcher) StakedNodes(ctx context.Context, staker common.Address) ([]*big.Int, error) {
	return v.con.StakedNodes(&bind.CallOpts{Context: ctx}, v.rollupAddress, staker.ToEthAddress())
}

func (v *ValidatorUtilsWatcher) CheckDecidableNextNode(ctx context.Context) (ConfirmType, core.NodeID, common.Address, error) {
	confirmType, successorWithStake, stakerAddress, err := v.con.CheckDecidableNextNode(
		&bind.CallOpts{Context: ctx},
		v.rollupAddress,
		big.NewInt(0),
		math.MaxBig256,
		big.NewInt(0),
		math.MaxBig256,
	)
	if err != nil {
		return CONFIRM_TYPE_NONE, nil, common.Address{}, err
	}
	return ConfirmType(confirmType), successorWithStake, common.NewAddressFromEth(stakerAddress), nil
}

func (v *ValidatorUtilsWatcher) FindStakerConflict(ctx context.Context, staker1, staker2 common.Address) (ConflictType, *big.Int, *big.Int, error) {
	conflictType, staker1Node, staker2Node, err := v.con.FindStakerConflict(
		&bind.CallOpts{Context: ctx},
		v.rollupAddress,
		staker1.ToEthAddress(),
		staker2.ToEthAddress(),
		math.MaxBig256,
	)
	if err != nil {
		return CONFLICT_TYPE_NONE, nil, nil, err
	}
	for ConflictType(conflictType) == CONFLICT_TYPE_INCOMPLETE {
		conflictType, staker1Node, staker2Node, err = v.con.FindNodeConflict(
			&bind.CallOpts{Context: ctx},
			v.rollupAddress,
			staker1Node,
			staker2Node,
			math.MaxBig256,
		)
		if err != nil {
			return CONFLICT_TYPE_NONE, nil, nil, err
		}
	}
	return ConflictType(conflictType), staker1Node, staker2Node, nil
}

type ValidatorUtils struct {
	*ValidatorUtilsWatcher
	auth *TransactAuth
}

func NewValidatorUtils(address, rollupAddress ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*ValidatorUtils, error) {
	watcher, err := NewValidatorUtilsWatcher(address, rollupAddress, client)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtils{
		ValidatorUtilsWatcher: watcher,
		auth:                  auth,
	}, nil
}

func (v *ValidatorUtils) RefundStakers(ctx context.Context, stakers []common.Address) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.RefundStakers(auth, v.rollupAddress, common.AddressArrayToEth(stakers))
	})
}
