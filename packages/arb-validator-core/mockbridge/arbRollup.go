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

package mockbridge

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type ArbRollup struct {
	EthRollupWatcher
	Client *ethclient.Client
}

func NewRollup(address common.Address, client arbbridge.ArbClient) (*ArbRollup, error) {
	//arbitrumRollupContract, err := rollup.NewArbRollup(address, client.(*ArbClient).client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ArbRollup")
	//}
	//vm := &ArbRollup{Client: client.(*ArbClient).client, ArbRollup: arbitrumRollupContract, auth: auth}
	return &ArbRollup{
		Client: nil,
	}, nil
}

func (vm *ArbRollup) PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) ([]arbbridge.Event, error) {
	//call := &bind.TransactOpts{
	//	From:    vm.auth.From,
	//	Signer:  vm.auth.Signer,
	//	Value:   stakeAmount,
	//	Context: ctx,
	//}
	//tx, err := vm.ArbRollup.PlaceStake(
	//	call,
	//	proof1,
	//	proof2,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "PlaceStake")
	return nil, nil
}

func (vm *ArbRollup) RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.RecoverStakeConfirmed(
	//	vm.auth,
	//	proof,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "RecoverStakeConfirmed")
	return nil, nil
}

func (vm *ArbRollup) RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.RecoverStakeOld(
	//	vm.auth,
	//	staker,
	//	proof,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "RecoverStakeOld")
	return nil, nil
}

func (vm *ArbRollup) RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.RecoverStakeMooted(
	//	vm.auth,
	//	staker,
	//	nodeHash,
	//	latestConfirmedProof,
	//	stakerProof,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "RecoverStakeMooted")
	return nil, nil
}

func (vm *ArbRollup) RecoverStakePassedDeadline(ctx context.Context, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal common.Hash, childType uint64, vmProtoStateHash common.Hash, proof []common.Hash) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.RecoverStakePassedDeadline(
	//	vm.auth,
	//	stakerAddress,
	//	deadlineTicks,
	//	disputableNodeHashVal,
	//	new(big.Int).SetUint64(childType),
	//	vmProtoStateHash,
	//	proof,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "RecoverStakePassedDeadline")
	return nil, nil
}

func (vm *ArbRollup) MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.MoveStake(
	//	vm.auth,
	//	proof1,
	//	proof2,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "MoveStake")
	return nil, nil
}

func (vm *ArbRollup) PruneLeaves(ctx context.Context, params []valprotocol.PruneParams) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.PruneLeaf(
	//	vm.auth,
	//	from,
	//	proof1,
	//	proof2,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "PruneLeaf")
	return nil, nil
}

func (vm *ArbRollup) MakeAssertion(
	ctx context.Context,

	prevPrevLeafHash common.Hash,
	prevDataHash common.Hash,
	prevDeadline common.TimeTicks,
	prevChildType valprotocol.ChildType,

	beforeState *valprotocol.VMProtoData,
	assertionParams *valprotocol.AssertionParams,
	assertionClaim *valprotocol.AssertionClaim,
	params [9][32]byte,
	stakerProof []common.Hash,
) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.MakeAssertion(
	//	vm.auth,
	//	[9][32]byte{
	//		beforeState.MachineHash,
	//		beforeState.InboxTop,
	//		prevPrevLeafHash,
	//		prevDataHash,
	//		assertionClaim.AfterInboxTop,
	//		assertionClaim.ImportedMessagesSlice,
	//		assertionClaim.AssertionStub.AfterHashValue(),
	//		assertionClaim.AssertionStub.LastMessageHashValue(),
	//		assertionClaim.AssertionStub.LastLogHashValue(),
	//	},
	//	beforeState.InboxCount,
	//	prevDeadline.Val,
	//	uint32(prevChildType),
	//	assertionParams.NumSteps,
	//	assertionParams.TimeBounds.AsIntArray(),
	//	assertionParams.ImportedMessageCount,
	//	assertionClaim.AssertionStub.DidInboxInsn,
	//	assertionClaim.AssertionStub.NumGas,
	//	stakerProof,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "MakeAssertion")
	return nil, nil
}

func (vm *ArbRollup) Confirm(ctx context.Context, opp *valprotocol.ConfirmOpportunity) ([]arbbridge.Event, error) {
	return nil, nil
}

func (vm *ArbRollup) StartChallenge(
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
) ([]arbbridge.Event, error) {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.StartChallenge(
	//	vm.auth,
	//	asserterAddress,
	//	challengerAddress,
	//	prevNode,
	//	disputableDeadline,
	//	[2]*big.Int{
	//		new(big.Int).SetUint64(uint64(asserterPosition)),
	//		new(big.Int).SetUint64(uint64(challengerPosition)),
	//	},
	//	[2][32]byte{
	//		asserterVMProtoHash,
	//		challengerVMProtoHash,
	//	},
	//	asserterProof,
	//	challengerProof,
	//	asserterDataHash,
	//	asserterPeriodTicks.Val,
	//	challengerNodeHash,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "StartExecutionChallenge")
	return nil, nil
}
