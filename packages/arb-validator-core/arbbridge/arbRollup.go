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

package arbbridge

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type ArbRollup interface {
	ArbRollupWatcher
	PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) ([]Event, error)
	RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) ([]Event, error)
	RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) ([]Event, error)
	RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) ([]Event, error)
	RecoverStakePassedDeadline(ctx context.Context, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal common.Hash, childType uint64, vmProtoStateHash common.Hash, proof []common.Hash) ([]Event, error)
	MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) ([]Event, error)
	PruneLeaves(ctx context.Context, params []valprotocol.PruneParams) ([]Event, error)
	MakeAssertion(
		ctx context.Context,
		prevPrevLeafHash common.Hash,
		prevDataHash common.Hash,
		prevDeadline common.TimeTicks,
		prevChildType valprotocol.ChildType,
		beforeState *valprotocol.VMProtoData,
		assertionParams *valprotocol.AssertionParams,
		assertionClaim *valprotocol.AssertionClaim,
		stakerProof []common.Hash) ([]Event, error)
	Confirm(ctx context.Context, opp *valprotocol.ConfirmOpportunity) ([]Event, error)
	StartChallenge(
		ctx context.Context,
		asserterAddress common.Address,
		challengerAddress common.Address,
		prevNode common.Hash,
		disputableDeadline *big.Int,
		asserterPosition valprotocol.ChildType,
		challengerPosition valprotocol.ChildType,
		asserterVMProtoHash common.Hash,
		challengerVMProtoHash common.Hash,
		asserterProof []common.Hash,
		challengerProof []common.Hash,
		asserterNodeHash common.Hash,
		challengerDataHash common.Hash,
		challengerPeriodTicks common.TimeTicks,
	) ([]Event, error)
}
