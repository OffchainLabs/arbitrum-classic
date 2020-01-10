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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ArbRollup interface {
	PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 [][32]byte, proof2 [][32]byte) error
	RecoverStakeConfirmed(ctx context.Context, proof [][32]byte) error
	RecoverStakeOld(ctx context.Context, staker common.Address, proof [][32]byte) error
	RecoverStakeMooted(ctx context.Context, nodeHash [32]byte, staker common.Address, latestConfirmedProof [][32]byte, stakerProof [][32]byte) error
	RecoverStakePassedDeadline(ctx context.Context, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal [32]byte, childType uint64, vmProtoStateHash [32]byte, proof [][32]byte) error
	MoveStake(ctx context.Context, proof1 [][32]byte, proof2 [][32]byte) error
	PruneLeaf(ctx context.Context, from [32]byte, proof1 [][32]byte, proof2 [][32]byte) error
	MakeAssertion(ctx context.Context, prevPrevLeafHash [32]byte, prevDataHash [32]byte, prevDeadline structures.TimeTicks, prevChildType structures.ChildType, beforeState *structures.VMProtoData, assertionParams *structures.AssertionParams, assertionClaim *structures.AssertionClaim, stakerProof [][32]byte) error
	ConfirmValid(
		ctx context.Context,
		deadline structures.TimeTicks,
		outMsgs []value.Value,
		logsAccHash [32]byte,
		protoHash [32]byte,
		stakerAddresses []common.Address,
		stakerProofs [][32]byte,
		stakerProofOffsets []*big.Int,
	) error
	ConfirmInvalid(
		ctx context.Context,
		deadline structures.TimeTicks,
		challengeNodeData [32]byte,
		branch structures.ChildType,
		protoHash [32]byte,
		stakerAddresses []common.Address,
		stakerProofs [][32]byte,
		stakerProofOffsets []*big.Int,
	) error
	StartChallenge(
		ctx context.Context,
		asserterAddress common.Address,
		challengerAddress common.Address,
		prevNode [32]byte,
		disputableDeadline *big.Int,
		asserterPosition structures.ChildType,
		challengerPosition structures.ChildType,
		asserterVMProtoHash [32]byte,
		challengerVMProtoHash [32]byte,
		asserterProof [][32]byte,
		challengerProof [][32]byte,
		asserterNodeHash [32]byte,
		challengerDataHash [32]byte,
		challengerPeriodTicks structures.TimeTicks,
	) error
	IsStaked(address common.Address) (bool, error)
}

//func (vm *ArbRollup) VerifyVM(
//	auth *bind.CallOpts,
//	config *valmessage.VMConfiguration,
//	machine [32]byte,
//) error {
//	//code, err := vm.contract.Client.CodeAt(auth.Context, vm.address, nil)
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

//func (vm *ArbRollup) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) (*types.Receipt, error) {
//	return waitForReceipt(ctx, vm.Client, vm.auth.From, tx, methodName)
//}
