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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type arbRollup struct {
	*ethRollupWatcher
	auth *TransactAuth
}

func newRollup(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*arbRollup, error) {
	watcher, err := newRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	vm := &arbRollup{ethRollupWatcher: watcher, auth: auth}
	return vm, err
}

func (vm *arbRollup) PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()

	blankAddress := ethcommon.Address{}
	st, err := vm.ArbRollup.GetStakeToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		if st == blankAddress {
			auth.Value = stakeAmount
		}
		return vm.ArbRollup.PlaceStake(
			auth,
			common.HashSliceToRaw(proof1),
			common.HashSliceToRaw(proof2),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "PlaceStake")
}

func (vm *arbRollup) RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.RecoverStakeConfirmed(
			auth,
			common.HashSliceToRaw(proof),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeConfirmed")
}

func (vm *arbRollup) RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.RecoverStakeOld(
			auth,
			staker.ToEthAddress(),
			common.HashSliceToRaw(proof),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeOld")
}

func (vm *arbRollup) RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.RecoverStakeMooted(
			auth,
			staker.ToEthAddress(),
			nodeHash,
			common.HashSliceToRaw(latestConfirmedProof),
			common.HashSliceToRaw(stakerProof),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakeMooted")
}

func (vm *arbRollup) RecoverStakePassedDeadline(ctx context.Context, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal common.Hash, childType uint64, vmProtoStateHash common.Hash, proof []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.RecoverStakePassedDeadline(
			auth,
			stakerAddress.ToEthAddress(),
			deadlineTicks,
			disputableNodeHashVal,
			new(big.Int).SetUint64(childType),
			vmProtoStateHash,
			common.HashSliceToRaw(proof),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "RecoverStakePassedDeadline")
}

func (vm *arbRollup) MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.MoveStake(
			auth,
			common.HashSliceToRaw(proof1),
			common.HashSliceToRaw(proof2),
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "MoveStake")
}

func (vm *arbRollup) PruneLeaves(ctx context.Context, opps []valprotocol.PruneParams) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()

	fromNodes := make([]common.Hash, 0, len(opps))
	leafProofs := make([]common.Hash, 0, len(opps))
	leafProofLengths := make([]*big.Int, 0, len(opps))
	confProofs := make([]common.Hash, 0, len(opps))
	confProofLengths := make([]*big.Int, 0, len(opps))
	for _, opp := range opps {
		fromNodes = append(fromNodes, opp.AncestorHash)
		leafProofs = append(leafProofs, opp.LeafProof...)
		leafProofLengths = append(leafProofLengths, big.NewInt(int64(len(opp.LeafProof))))
		confProofs = append(confProofs, opp.AncProof...)
		confProofLengths = append(confProofLengths, big.NewInt(int64(len(opp.AncProof))))
	}

	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.PruneLeaves(
			auth,
			common.HashSliceToRaw(fromNodes),
			common.HashSliceToRaw(leafProofs),
			leafProofLengths,
			common.HashSliceToRaw(confProofs),
			confProofLengths,
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "PruneLeaf")
}

func (vm *arbRollup) MakeAssertion(
	ctx context.Context,
	prevPrevLeafHash common.Hash,
	prevDataHash common.Hash,
	prevDeadline common.TimeTicks,
	prevChildType valprotocol.ChildType,
	beforeState *valprotocol.VMProtoData,
	assertionParams *valprotocol.AssertionParams,
	assertion *valprotocol.ExecutionAssertionStub,
	stakerProof []common.Hash,
	validBlock *common.BlockId,
) ([]arbbridge.Event, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	fields := [8][32]byte{
		beforeState.MachineHash,
		assertion.AfterMachineHash,
		assertion.BeforeInboxAcc,
		assertion.AfterInboxAcc,
		assertion.LastMessageHash,
		assertion.LastLogHash,
		prevPrevLeafHash,
		prevDataHash,
	}
	fields2 := [5]*big.Int{
		beforeState.InboxCount,
		prevDeadline.Val,
		assertionParams.ImportedMessageCount,
		beforeState.MessageCount,
		beforeState.LogCount,
	}
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.MakeAssertion(
			auth,
			fields,
			fields2,
			validBlock.HeaderHash,
			validBlock.Height.AsInt(),
			assertion.MessageCount,
			assertion.LogCount,
			uint32(prevChildType),
			assertionParams.NumSteps,
			assertion.NumGas,
			common.HashSliceToRaw(stakerProof),
		)
	})
	if err != nil {
		callErr := vm.ArbRollup.MakeAssertionCall(
			ctx,
			vm.client,
			vm.auth.auth.From,
			vm.rollupAddress,
			fields,
			fields2,
			validBlock.HeaderHash,
			validBlock.Height.AsInt(),
			assertion.MessageCount,
			assertion.LogCount,
			uint32(prevChildType),
			assertionParams.NumSteps,
			assertion.NumGas,
			common.HashSliceToRaw(stakerProof),
		)
		return nil, errors.WithStack(callErr)
	}
	return vm.waitForReceipt(ctx, tx, "MakeAssertion")
}

func (vm *arbRollup) Confirm(ctx context.Context, opp *valprotocol.ConfirmOpportunity) ([]arbbridge.Event, error) {
	proof := opp.PrepareProof()
	vm.auth.Lock()
	defer vm.auth.Unlock()

	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.Confirm(
			auth,
			proof.InitalProtoStateHash,
			proof.BeforeSendCount,
			proof.BranchesNums,
			proof.DeadlineTicks,
			proof.ChallengeNodeData,
			proof.LogsAcc,
			proof.VMProtoStateHashes,
			proof.MessageCounts,
			proof.Messages,
			addressSliceToRaw(opp.StakerAddresses),
			proof.CombinedProofs,
			proof.StakerProofOffsets,
		)
	})
	if err != nil {
		return nil, vm.ArbRollup.ConfirmCall(
			ctx,
			vm.client,
			vm.auth.auth.From,
			vm.rollupAddress,
			proof.InitalProtoStateHash,
			proof.BeforeSendCount,
			proof.BranchesNums,
			proof.DeadlineTicks,
			proof.ChallengeNodeData,
			proof.LogsAcc,
			proof.VMProtoStateHashes,
			proof.MessageCounts,
			proof.Messages,
			addressSliceToRaw(opp.StakerAddresses),
			proof.CombinedProofs,
			proof.StakerProofOffsets,
		)
	}
	return vm.waitForReceipt(ctx, tx, "Confirm")
}

func (vm *arbRollup) StartChallenge(
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
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return vm.ArbRollup.StartChallenge(
			auth,
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
			common.HashSliceToRaw(asserterProof),
			common.HashSliceToRaw(challengerProof),
			asserterNodeHash,
			challengerDataHash,
			challengerPeriodTicks.Val,
		)
	})
	if err != nil {
		return nil, err
	}
	return vm.waitForReceipt(ctx, tx, "StartExecutionChallenge")
}

func (vm *arbRollup) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) ([]arbbridge.Event, error) {
	receipt, err := WaitForReceiptWithResults(ctx, vm.client, vm.auth.auth.From, tx, methodName)
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(receipt.Logs))

	for _, log := range receipt.Logs {
		chainInfo := getLogChainInfo(*log)
		ev, err := vm.processEvents(chainInfo, *log)
		if err != nil {
			continue
		}
		events = append(events, ev)
	}

	return events, nil
}
