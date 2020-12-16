/*
 * Copyright 2020, Offchain Labs, Inc.
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

package chainobserver

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"math/big"
	"testing"
)

func TestMoveStake(t *testing.T) {
	rollup := getRollup(t)
	ctx := context.Background()
	_, err := rollup.MoveStake(
		ctx,
		[]common.Hash{},
		[]common.Hash{},
	)
	if err == nil {
		t.Fatal("Should not be able to move stake")
	}

	_, err = rollup.PlaceStake(
		ctx,
		big.NewInt(0),
		[]common.Hash{},
		[]common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = rollup.MoveStake(
		ctx,
		[]common.Hash{},
		[]common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecoverStakeOld(t *testing.T) {
	rollup := getRollup(t)

	ctx := context.Background()
	_, err := rollup.RecoverStakeOld(
		ctx,
		common.Address{},
		[]common.Hash{},
	)
	if err == nil {
		t.Fatal(err)
	}

	_, err = rollup.PlaceStake(
		ctx,
		big.NewInt(0),
		[]common.Hash{},
		[]common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = rollup.RecoverStakeOld(
		ctx,
		common.Address{},
		[]common.Hash{},
	)
	if err == nil {
		t.Fatal(err)
	}
}

func TestRecoverStake(t *testing.T) {
	rollup := getRollup(t)

	ctx := context.Background()
	_, err := rollup.RecoverStakeConfirmed(
		ctx,
		[]common.Hash{},
	)
	if err == nil {
		t.Fatal(err)
	}

	_, err = rollup.PlaceStake(
		ctx,
		big.NewInt(0),
		[]common.Hash{},
		[]common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = rollup.RecoverStakeConfirmed(
		ctx,
		[]common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func getRollup(t *testing.T) arbbridge.ArbRollup {
	ctx := context.Background()

	chainParams := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(0),
		GracePeriod:             common.TicksFromSeconds(1),
		MaxExecutionSteps:       100000,
		ArbGasSpeedLimitPerTick: 100000,
	}

	arbFactoryAddress, err := ethbridge.DeployRollupFactory(ctx, authClient)
	if err != nil {
		t.Fatal(err)
	}

	factory, err := authClient.NewArbFactory(common.NewAddressFromEth(arbFactoryAddress))
	if err != nil {
		t.Fatal(err)
	}

	mach, err := loader.LoadMachineFromFile(arbos.Path(), false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	rollupAddress, _, err := factory.CreateRollup(
		ctx,
		mach.Hash(),
		chainParams,
		common.Address{},
	)
	if err != nil {
		t.Fatal(err)
	}

	rollupContract, err := authClient.NewRollup(rollupAddress)
	if err != nil {
		t.Fatal(err)
	}

	return rollupContract
}
