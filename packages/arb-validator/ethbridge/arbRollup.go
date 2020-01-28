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

package ethbridge

import (
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type arbRollup struct {
	Client          *ethclient.Client
	ArbRollup       *rollup.ArbRollup
	auth            *TransactAuth
	contractAddress ethcommon.Address
}

func newRollup(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*arbRollup, error) {
	arbitrumRollupContract, err := rollup.NewArbRollup(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbRollup")
	}
	vm := &arbRollup{Client: client, ArbRollup: arbitrumRollupContract, auth: auth, contractAddress: address}
	return vm, err
}

func (vm *arbRollup) PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	call := &bind.TransactOpts{
		From:    vm.auth.auth.From,
		Signer:  vm.auth.auth.Signer,
		Value:   stakeAmount,
		Context: ctx,
	}
	tx, err := vm.ArbRollup.PlaceStake(
		call,
		hashSliceToRaw(proof1),
		hashSliceToRaw(proof2),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "PlaceStake")
}

func (vm *arbRollup) RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.RecoverStakeConfirmed(
		vm.auth.getAuth(ctx),
		hashSliceToRaw(proof),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeConfirmed")
}

func (vm *arbRollup) RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.RecoverStakeOld(
		vm.auth.getAuth(ctx),
		staker.ToEthAddress(),
		hashSliceToRaw(proof),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeOld")
}

func (vm *arbRollup) RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.RecoverStakeMooted(
		vm.auth.getAuth(ctx),
		staker.ToEthAddress(),
		nodeHash,
		hashSliceToRaw(latestConfirmedProof),
		hashSliceToRaw(stakerProof),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeMooted")
}

func (vm *arbRollup) RecoverStakePassedDeadline(ctx context.Context, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal common.Hash, childType uint64, vmProtoStateHash common.Hash, proof []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.RecoverStakePassedDeadline(
		vm.auth.getAuth(ctx),
		stakerAddress.ToEthAddress(),
		deadlineTicks,
		disputableNodeHashVal,
		new(big.Int).SetUint64(childType),
		vmProtoStateHash,
		hashSliceToRaw(proof),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakePassedDeadline")
}

func (vm *arbRollup) MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.MoveStake(
		vm.auth.getAuth(ctx),
		hashSliceToRaw(proof1),
		hashSliceToRaw(proof2),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "MoveStake")
}

func (vm *arbRollup) PruneLeaf(ctx context.Context, from common.Hash, proof1 []common.Hash, proof2 []common.Hash) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.PruneLeaf(
		vm.auth.getAuth(ctx),
		from,
		hashSliceToRaw(proof1),
		hashSliceToRaw(proof2),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "PruneLeaf")
}

func (vm *arbRollup) MakeAssertion(
	ctx context.Context,

	prevPrevLeafHash common.Hash,
	prevDataHash common.Hash,
	prevDeadline common.TimeTicks,
	prevChildType structures.ChildType,

	beforeState *structures.VMProtoData,
	assertionParams *structures.AssertionParams,
	assertionClaim *structures.AssertionClaim,
	stakerProof []common.Hash,
) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.MakeAssertion(
		vm.auth.getAuth(ctx),
		[9][32]byte{
			beforeState.MachineHash,
			beforeState.PendingTop,
			prevPrevLeafHash,
			prevDataHash,
			assertionClaim.AfterPendingTop,
			assertionClaim.ImportedMessagesSlice,
			assertionClaim.AssertionStub.AfterHash,
			assertionClaim.AssertionStub.LastMessageHash,
			assertionClaim.AssertionStub.LastLogHash,
		},
		beforeState.PendingCount,
		prevDeadline.Val,
		uint32(prevChildType),
		assertionParams.NumSteps,
		assertionParams.TimeBounds.AsIntArray(),
		assertionParams.ImportedMessageCount,
		assertionClaim.AssertionStub.DidInboxInsn,
		assertionClaim.AssertionStub.NumGas,
		hashSliceToRaw(stakerProof),
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "MakeAssertion")
}

func (vm *arbRollup) ConfirmValid(
	ctx context.Context,
	deadline common.TimeTicks,
	outMsgs []value.Value,
	logsAccHash common.Hash,
	protoHash common.Hash,
	stakerAddresses []common.Address,
	stakerProofs []common.Hash,
	stakerProofOffsets []*big.Int,
) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	messages := combineMessages(outMsgs)
	tx, err := vm.ArbRollup.ConfirmValid(
		vm.auth.getAuth(ctx),
		deadline.Val,
		messages,
		logsAccHash,
		protoHash,
		addressSliceToRaw(stakerAddresses),
		hashSliceToRaw(stakerProofs),
		stakerProofOffsets,
	)
	if err != nil {
		return vm.ArbRollup.ConfirmValidCall(
			ctx,
			vm.Client,
			vm.auth.auth.From,
			vm.contractAddress,
			deadline.Val,
			messages,
			logsAccHash,
			protoHash,
			addressSliceToRaw(stakerAddresses),
			hashSliceToRaw(stakerProofs),
			stakerProofOffsets,
		)
	}
	return vm.waitForReceipt(ctx, tx, "ConfirmValid")
}

func (vm *arbRollup) ConfirmInvalid(
	ctx context.Context,
	deadline common.TimeTicks,
	challengeNodeData common.Hash,
	branch structures.ChildType,
	protoHash common.Hash,
	stakerAddresses []common.Address,
	stakerProofs []common.Hash,
	stakerProofOffsets []*big.Int,
) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.ConfirmInvalid(
		vm.auth.getAuth(ctx),
		deadline.Val,
		challengeNodeData,
		new(big.Int).SetUint64(uint64(branch)),
		protoHash,
		addressSliceToRaw(stakerAddresses),
		hashSliceToRaw(stakerProofs),
		stakerProofOffsets,
	)
	if err != nil {
		return vm.ArbRollup.ConfirmInvalidCall(
			ctx,
			vm.Client,
			vm.auth.auth.From,
			vm.contractAddress,
			deadline.Val,
			challengeNodeData,
			new(big.Int).SetUint64(uint64(branch)),
			protoHash,
			addressSliceToRaw(stakerAddresses),
			hashSliceToRaw(stakerProofs),
			stakerProofOffsets,
		)
	}
	return vm.waitForReceipt(ctx, tx, "ConfirmInvalid")
}

func (vm *arbRollup) StartChallenge(
	ctx context.Context,
	asserterAddress common.Address,
	challengerAddress common.Address,
	prevNode common.Hash,
	disputableDeadline *big.Int,
	asserterPosition structures.ChildType,
	challengerPosition structures.ChildType,
	asserterVMProtoHash common.Hash,
	challengerVMProtoHash common.Hash,
	asserterProof []common.Hash,
	challengerProof []common.Hash,
	asserterNodeHash common.Hash,
	challengerDataHash common.Hash,
	challengerPeriodTicks common.TimeTicks,
) error {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.StartChallenge(
		vm.auth.getAuth(ctx),
		asserterAddress.ToEthAddress(),
		challengerAddress.ToEthAddress(),
		prevNode,
		disputableDeadline,
		[2]*big.Int{
			new(big.Int).SetUint64(uint64(asserterPosition)),
			new(big.Int).SetUint64(uint64(challengerPosition)),
		},
		[2][32]byte{
			asserterVMProtoHash,
			challengerVMProtoHash,
		},
		hashSliceToRaw(asserterProof),
		hashSliceToRaw(challengerProof),
		asserterNodeHash,
		challengerDataHash,
		challengerPeriodTicks.Val,
	)
	if err != nil {
		return err
	}
	return vm.waitForReceipt(ctx, tx, "StartExecutionChallenge")
}

func (vm *arbRollup) IsStaked(address common.Address) (bool, error) {
	return vm.ArbRollup.IsStaked(nil, address.ToEthAddress())
}

//func (vm *arbRollup) VerifyVM(
//	auth *bind.CallOpts,
//	config *valmessage.VMConfiguration,
//	machine common.Hash,
//) error {
//	//code, err := vm.client.CodeAt(auth.Context, vm.address, nil)
//	// Verify that VM has correct code
//	vmInfo, err := vm.ArbRollup.Vm(auth)
//	if err != nil {
//		return err
//	}
//
//	if vmInfo.MachineHash != machine {
//		return errors.New("VM has different machine hash")
//	}
//
//	if config.GracePeriod != uint64(vmInfo.GracePeriod) {
//		return errors.New("VM has different grace period")
//	}
//
//	if value.NewBigIntFromBuf(config.EscrowRequired).Cmp(vmInfo.EscrowRequired) != 0 {
//		return errors.New("VM has different escrow required")
//	}
//
//	if config.MaxExecutionStepCount != vmInfo.MaxExecutionSteps {
//		return errors.New("VM has different mxa steps")
//	}
//
//	owner, err := vm.ArbRollup.Owner(auth)
//	if err != nil {
//		return err
//	}
//	if protocol.NewAddressFromBuf(config.Owner) != owner {
//		return errors.New("VM has different owner")
//	}
//	return nil
//}

func (vm *arbRollup) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, vm.Client, vm.auth.auth.From, tx, methodName)
}
