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

package ethbridge

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
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

type ValidatorUtils struct {
	con           *ethbridgecontracts.ValidatorUtils
	client        ethutils.EthClient
	address       ethcommon.Address
	rollupAddress ethcommon.Address
	baseCallOpts  bind.CallOpts
}

func NewValidatorUtils(address, rollupAddress ethcommon.Address, client ethutils.EthClient, callOpts bind.CallOpts) (*ValidatorUtils, error) {
	con, err := ethbridgecontracts.NewValidatorUtils(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ValidatorUtils{
		con:           con,
		client:        client,
		address:       address,
		rollupAddress: rollupAddress,
		baseCallOpts:  callOpts,
	}, nil
}

func (v *ValidatorUtils) getCallOpts(ctx context.Context) *bind.CallOpts {
	opts := v.baseCallOpts
	opts.Context = ctx
	return &opts
}

func (v *ValidatorUtils) RefundableStakers(ctx context.Context) ([]common.Address, error) {
	addresses, err := v.con.RefundableStakers(v.getCallOpts(ctx), v.rollupAddress)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return common.AddressArrayFromEth(addresses), nil
}

func (v *ValidatorUtils) TimedOutChallenges(ctx context.Context, max int) ([]common.Address, error) {
	i := big.NewInt(0)
	count := big.NewInt(1024)
	addresses := make([]ethcommon.Address, 0)
	for {
		newAddrs, hasMore, err := v.con.TimedOutChallenges(v.getCallOpts(ctx), v.rollupAddress, i, count)
		addresses = append(addresses, newAddrs...)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !hasMore {
			break
		}
		if len(addresses) >= max {
			break
		}
		i = i.Add(i, count)
	}
	if len(addresses) > max {
		addresses = addresses[:max]
	}
	return common.AddressArrayFromEth(addresses), nil
}

type RollupConfig struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}

func (v *ValidatorUtils) GetConfig(ctx context.Context) (*RollupConfig, error) {
	config, err := v.con.GetConfig(v.getCallOpts(ctx), v.rollupAddress)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &RollupConfig{
		ConfirmPeriodBlocks:      config.ConfirmPeriodBlocks,
		ExtraChallengeTimeBlocks: config.ExtraChallengeTimeBlocks,
		ArbGasSpeedLimitPerBlock: config.ArbGasSpeedLimitPerBlock,
		BaseStake:                config.BaseStake,
		StakeToken:               common.Address{},
	}, nil
}

func (v *ValidatorUtils) GetStakers(ctx context.Context) ([]common.Address, error) {
	addresses, _, err := v.con.GetStakers(v.getCallOpts(ctx), v.rollupAddress, big.NewInt(0), math.MaxBig256)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return common.AddressArrayFromEth(addresses), nil
}

func (v *ValidatorUtils) LatestStaked(ctx context.Context, staker common.Address) (*big.Int, [32]byte, error) {
	amount, hash, err := v.con.LatestStaked(v.getCallOpts(ctx), v.rollupAddress, staker.ToEthAddress())
	return amount, hash, errors.WithStack(err)
}

func (v *ValidatorUtils) StakedNodes(ctx context.Context, staker common.Address) ([]*big.Int, error) {
	nodes, err := v.con.StakedNodes(v.getCallOpts(ctx), v.rollupAddress, staker.ToEthAddress())
	return nodes, errors.WithStack(err)
}

func (v *ValidatorUtils) AreUnresolvedNodesLinear(ctx context.Context) (bool, error) {
	linear, err := v.con.AreUnresolvedNodesLinear(v.getCallOpts(ctx), v.rollupAddress)
	return linear, errors.WithStack(err)
}

func (v *ValidatorUtils) CheckDecidableNextNode(ctx context.Context) (ConfirmType, error) {
	confirmType, err := v.con.CheckDecidableNextNode(
		v.getCallOpts(ctx),
		v.rollupAddress,
	)
	if err != nil {
		return CONFIRM_TYPE_NONE, errors.WithStack(err)
	}
	return ConfirmType(confirmType), nil
}

func (v *ValidatorUtils) FindStakerConflict(ctx context.Context, staker1, staker2 common.Address) (ConflictType, *big.Int, *big.Int, error) {
	conflictType, staker1Node, staker2Node, err := v.con.FindStakerConflict(
		v.getCallOpts(ctx),
		v.rollupAddress,
		staker1.ToEthAddress(),
		staker2.ToEthAddress(),
		math.MaxBig256,
	)
	if err != nil {
		return CONFLICT_TYPE_NONE, nil, nil, errors.WithStack(err)
	}
	for ConflictType(conflictType) == CONFLICT_TYPE_INCOMPLETE {
		conflictType, staker1Node, staker2Node, err = v.con.FindNodeConflict(
			v.getCallOpts(ctx),
			v.rollupAddress,
			staker1Node,
			staker2Node,
			math.MaxBig256,
		)
		if err != nil {
			return CONFLICT_TYPE_NONE, nil, nil, errors.WithStack(err)
		}
	}
	return ConflictType(conflictType), staker1Node, staker2Node, nil
}
