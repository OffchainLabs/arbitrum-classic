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
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ArbRollup struct {
	*nodeGraph
	Client MockArbAuthClient
}

func newRollup(address common.Address, client MockArbAuthClient) (*ArbRollup, error) {
	//arbitrumRollupContract, err := rollup.NewArbRollup(address, client.(*ArbClient).client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ArbRollup")
	//}
	//vm := &ArbRollup{Client: client.(*MockArbClient).client, ArbRollup: arbitrumRollupContract, auth: auth}

	// arbRollup init()
	// 	NodeGraph init()
	//	staking init()
	//        require(address(challengeFactory) == address(0), INIT_TWICE);
	//        require(_challengeFactoryAddress != address(0), INIT_NONZERO);
	//
	//        challengeFactory = IChallengeFactory(_challengeFactoryAddress);
	//
	//        // VM parameters
	//        stakeRequirement = _stakeRequirement;

	vm := newNodeGraph(client.auth)
	return &ArbRollup{
		nodeGraph: vm,
		Client:    client,
	}, nil
}

func (vm *ArbRollup) PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) error {
	location := calculatePath(vm.lastConfirmed, proof1)
	leaf := calculatePath(location, proof2)
	if !vm.leaves[leaf] {
		errors.New("invalid path proof")
	}
	if err := vm.createStake(stakeAmount, location); err != nil {
		return err
	}

	//emit RollupStakeCreated(msg.sender, location);
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.StakeCreatedEvent{
			Staker:   vm.Client.auth.From,
			NodeHash: vm.Client.MockEthClient.LatestBlock.HeaderHash,
		},
	})
	return nil
}

func (vm *ArbRollup) createStake(stakeAmount *big.Int, location common.Hash) error {
	// require(msg.value == stakeRequirement, STK_AMT);
	if vm.Client.auth.Value != vm.stakeRequirement {
		return errors.New("invalid stake amount")
	}
	if _, ok := vm.stakers[vm.Client.auth.From]; ok {
		return errors.New("staker already exists")
	}
	// require(stakers[msg.sender].location == 0x00, ALRDY_STAKED);
	vm.stakers[vm.Client.auth.From] = &staker{location, vm.Client.MockEthClient.LatestBlock.Height, false, stakeAmount}

	return nil
}

func (vm *ArbRollup) refundStaker(staker common.Address) {
	//refundStaker(stakerAddress);
	delete(vm.stakers, staker)
	//transfer stake requirement
	// ???
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.StakeRefundedEvent{
			Staker: staker,
		},
	})

}
func (vm *ArbRollup) RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) error {
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	//require(RollupUtils.calculatePath(stakerLocation, proof) == latestConfirmed(), RECOV_PATH_PROOF);
	//refundStaker(stakerAddress);

	staker, ok := vm.stakers[vm.Client.auth.From]
	if !ok {
		return errors.New("staker not found")
	}

	if calculatePath(staker.location, proof) != vm.lastConfirmed {
		return errors.New("invalid path proof")
	}

	// refundStaker
	vm.refundStaker(vm.Client.auth.From)

	//emit RollupStakeRefunded(address(_stakerAddress));
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.StakeRefundedEvent{
			Staker: vm.Client.auth.From,
		},
	})

	return nil
}

func (vm *ArbRollup) RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) error {
	//require(proof.length > 0, RECVOLD_LENGTH);
	if len(proof) <= 0 {
		return errors.New("proof must be non-zero length")
	}
	//_recoverStakeConfirmed(stakerAddress, proof);
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	st, ok := vm.stakers[staker]
	if !ok {
		return errors.New("staker not found")
	}
	//require(RollupUtils.calculatePath(stakerLocation, proof) == latestConfirmed(), RECOV_PATH_PROOF);
	if calculatePath(st.location, proof) != vm.lastConfirmed {
		return errors.New("invalid path proof")
	}
	vm.refundStaker(staker)

	return nil
}

func (vm *ArbRollup) RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) error {
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	//require(
	//	latestConfirmedProof[0] != stakerProof[0] &&
	//		RollupUtils.calculatePath(node, latestConfirmedProof) == latestConfirmed() &&
	//		RollupUtils.calculatePath(node, stakerProof) == stakerLocation,
	//	RECOV_CONFLICT_PROOF
	//);
	if latestConfirmedProof[0] == stakerProof[0] ||
		calculatePath(nodeHash, latestConfirmedProof) == vm.lastConfirmed ||
		calculatePath(nodeHash, stakerProof) != vm.stakers[vm.Client.auth.From].location {
		return errors.New("Invalid conflict proof")
	}
	//refundStaker(stakerAddress);
	vm.refundStaker(staker)

	return nil
}

func (vm *ArbRollup) RecoverStakePassedDeadline(
	ctx context.Context,
	stakerAddress common.Address,
	deadlineTicks *big.Int,
	disputableNodeHashVal common.Hash,
	childType uint64,
	vmProtoStateHash common.Hash,
	proof []common.Hash) error {
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	//bytes32 nextNode = RollupUtils.childNodeHash(
	//	stakerLocation,
	//	deadlineTicks,
	//	disputableNodeHashVal,
	//	childType,
	//	vmProtoStateHash
	//);
	//bytes32 leaf = RollupUtils.calculatePath(nextNode, proof);
	//???
	//require(isValidLeaf(leaf), RECOV_DEADLINE_LEAF);
	//require(block.number >= RollupTime.blocksToTicks(deadlineTicks), RECOV_DEADLINE_TIME);
	if vm.Client.MockEthClient.LatestBlock.Height.Cmp(deadlineTicks) < 0 {
		return errors.New("Node is not passed deadline")
	}
	//refundStaker(stakerAddress);
	vm.refundStaker(stakerAddress)

	return nil
}

func (vm *ArbRollup) MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) error {
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	//bytes32 newLocation = RollupUtils.calculatePath(stakerLocation, proof1);
	//bytes32 leaf = RollupUtils.calculatePath(newLocation, proof2);
	//require(isValidLeaf(leaf), MOVE_LEAF);
	//updateStakerLocation(msg.sender, newLocation);
	location := vm.stakers[vm.Client.auth.From].location
	newLocation := calculatePath(location, proof1)
	leaf := calculatePath(newLocation, proof2)
	if !vm.leaves[leaf] {
		return errors.New("MoveStake - invalid leaf")
	}
	vm.stakers[vm.Client.auth.From].location = newLocation
	//emit RollupStakeRefunded(address(_stakerAddress));
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.StakeRefundedEvent{
			Staker: vm.Client.auth.From,
		},
	})

	return nil
}

func (vm *ArbRollup) PruneLeaf(ctx context.Context, from common.Hash, leafProof []common.Hash, ancProof []common.Hash) error {
	//bytes32 leaf = RollupUtils.calculatePath(from, leafProof);
	leaf := calculatePath(from, leafProof)
	//require(isValidLeaf(leaf), PRUNE_LEAF);
	if !vm.leaves[leaf] {
		return errors.New("MoveStake - invalid leaf")
	}
	//require(
	//	leafProof[0] != latestConfirmedProof[0] &&
	//		RollupUtils.calculatePath(from, latestConfirmedProof) == latestConfirmed(),
	//	PRUNE_CONFLICT
	//);
	if leafProof[0] == ancProof[0] ||
		calculatePath(from, ancProof) != vm.lastConfirmed {
		return errors.New("prune conflict")
	}
	//delete leaves[leaf];
	delete(vm.leaves, leaf)
	//
	//emit RollupPruned(leaf);
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.PrunedEvent{
			leaf,
		},
	})

	return nil
}

func (vm *ArbRollup) MakeAssertion(
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
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.MakeAssertion(
	//	vm.auth,
	//	[9][32]byte{
	//		beforeState.MachineHash,
	//		beforeState.PendingTop,
	//		prevPrevLeafHash,
	//		prevDataHash,
	//		assertionClaim.AfterPendingTop,
	//		assertionClaim.ImportedMessagesSlice,
	//		assertionClaim.AssertionStub.AfterHashValue(),
	//		assertionClaim.AssertionStub.LastMessageHashValue(),
	//		assertionClaim.AssertionStub.LastLogHashValue(),
	//	},
	//	beforeState.PendingCount,
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
	//(bytes32 prevLeaf, bytes32 newValid) = makeAssertion(
	//	MakeAssertionData(
	prevLeaf, newValid, err := vm.makeAssertion(assertionData)
	if err != nil {
		return err
	}
	//bytes32 stakerLocation = getStakerLocation(msg.sender);
	//require(RollupUtils.calculatePath(stakerLocation, _stakerProof) == prevLeaf, MAKE_STAKER_PROOF);
	if calculatePath(vm.stakers[vm.Client.auth.From].location, stakerProof) != prevLeaf {
		return errors.New("invalid staker location proof")
	}

	//updateStakerLocation(msg.sender, newValid);
	vm.stakers[vm.Client.auth.From].location = newValid
	//emit RollupStakeRefunded(address(_stakerAddress));
	vm.Client.MockEthClient.pubMsg(arbbridge.Notification{
		BlockId: vm.Client.MockEthClient.LatestBlock,
		Event: arbbridge.StakeRefundedEvent{
			Staker: vm.Client.auth.From,
		},
	})

	return nil
}

func (vm *ArbRollup) makeAssertion(prevPrevLeafHash common.Hash,
	prevDataHash common.Hash,
	prevDeadline common.TimeTicks,
	prevChildType structures.ChildType,

	beforeState *structures.VMProtoData,
	assertionParams *structures.AssertionParams,
	assertionClaim *structures.AssertionClaim,
) (common.Hash, common.Hash, error) {
	//bytes32 vmProtoHashBefore = RollupUtils.protoStateHash(
	//data.beforeVMHash,
	//data.beforePendingTop,
	//data.beforePendingCount
	//);
	protoHashBefore := beforeState.Hash()
	//bytes32 prevLeaf = RollupUtils.childNodeHash(
	//data.prevPrevLeafHash,
	//data.prevDeadlineTicks,
	//data.prevDataHash,
	//data.prevChildType,
	//vmProtoHashBefore
	//);
	prevLeaf := childNodeHash(prevPrevLeafHash,
		prevDeadline,
		prevDataHash,
		prevChildType,
		protoHashBefore,
	)
	//require(isValidLeaf(prevLeaf), MAKE_LEAF);
	if !vm.leaves[prevLeaf] {
		return common.Hash{}, common.Hash{}, errors.New("makeAssertion - invalid leaf")
	}
	//require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
	// if
	//require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
	if assertionParams.NumSteps > vmParams.maxExecutionSteps {
		return common.Hash{}, common.Hash{}, errors.New("makeAssertion - Tried to execute too many steps")
	}
	//require(VM.withinTimeBounds(data.timeBoundsBlocks), MAKE_TIME);
	if !withinTimeBounds(assertionParams.TimeBounds) {
		return common.Hash{}, common.Hash{}, errors.New("makeAssertion - Precondition: not within time bounds")
	}
	//require(data.importedMessageCount == 0 || data.didInboxInsn, MAKE_MESSAGES);
	if assertionParams.ImportedMessageCount.Cmp(big.NewInt(0)) != 0 && !assertionClaim.AssertionStub.DidInboxInsn {
		return common.Hash{}, common.Hash{}, errors.New("makeAssertion - Imported messages without reading them")
	}
	//(bytes32 pendingValue, uint256 pendingCount) = globalInbox.getPending();
	pendingValue, pendingCount = globalInbox.getPending()
	//require(data.importedMessageCount <= pendingCount.sub(data.beforePendingCount), MAKE_MESSAGE_CNT);
	if assertionParams.ImportedMessageCount > pendingCount.sub(beforeState.PendingCount) {
		return errors.New("makeAssertion - Tried to import more messages than exist in pending inbox")
	}
	//
	//uint256 gracePeriodTicks = vmParams.gracePeriodTicks;
	//uint256 checkTimeTicks = data.numArbGas / vmParams.arbGasSpeedLimitPerTick;
	//uint256 deadlineTicks = RollupTime.blocksToTicks(block.number) + gracePeriodTicks;
	//if (deadlineTicks < data.prevDeadlineTicks) {
	//deadlineTicks = data.prevDeadlineTicks;
	//}
	//deadlineTicks += checkTimeTicks;

	return common.Hash{}, common.Hash{}, nil
}

//
//
//bytes32 invalidPending = generateInvalidPendingTopLeaf(
//data,
//prevLeaf,
//deadlineTicks,
//pendingValue,
//pendingCount,
//vmProtoHashBefore,
//gracePeriodTicks
//);
//bytes32 invalidMessages = generateInvalidMessagesLeaf(
//data,
//prevLeaf,
//deadlineTicks,
//vmProtoHashBefore,
//gracePeriodTicks
//);
//bytes32 invalidExec = generateInvalidExecutionLeaf(
//data,
//prevLeaf,
//deadlineTicks,
//vmProtoHashBefore,
//gracePeriodTicks,
//checkTimeTicks
//);
//bytes32 validHash = generateValidLeaf(
//data,
//prevLeaf,
//deadlineTicks
//);
//
//leaves[invalidPending] = true;
//leaves[invalidMessages] = true;
//leaves[invalidExec] = true;
//leaves[validHash] = true;
//delete leaves[prevLeaf];
//
//emitAssertedEvent(data, prevLeaf, pendingValue, pendingCount);
//return (prevLeaf, validHash);

func (vm *ArbRollup) ConfirmValid(
	ctx context.Context,
	deadline common.TimeTicks,
	outMsgs []value.Value,
	logsAccHash common.Hash,
	protoHash common.Hash,
	stakerAddresses []common.Address,
	stakerProofs []common.Hash,
	stakerProofOffsets []*big.Int,
) error {
	//vm.auth.Context = ctx
	//messages := hashing.CombineMessages(outMsgs)
	//tx, err := vm.ArbRollup.ConfirmValid(
	//	vm.auth,
	//	deadline.Val,
	//	messages,
	//	logsAccHash,
	//	protoHash,
	//	stakerAddresses,
	//	stakerProofs,
	//	stakerProofOffsets,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "ConfirmValid")
	return nil
}

func (vm *ArbRollup) ConfirmInvalid(
	ctx context.Context,
	deadline common.TimeTicks,
	challengeNodeData common.Hash,
	branch structures.ChildType,
	protoHash common.Hash,
	stakerAddresses []common.Address,
	stakerProofs []common.Hash,
	stakerProofOffsets []*big.Int,
) error {
	//vm.auth.Context = ctx
	//tx, err := vm.ArbRollup.ConfirmInvalid(
	//	vm.auth,
	//	deadline.Val,
	//	challengeNodeData,
	//	new(big.Int).SetUint64(uint64(branch)),
	//	protoHash,
	//	stakerAddresses,
	//	stakerProofs,
	//	stakerProofOffsets,
	//)
	//if err != nil {
	//	return err
	//}
	//return vm.waitForReceipt(ctx, tx, "ConfirmInvalid")
	return nil
}

func (vm *ArbRollup) StartChallenge(
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
	return nil
}

func (vm *ArbRollup) IsStaked(address common.Address) (bool, error) {
	return false, nil
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

//func (vm *ArbRollup) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
//	return waitForReceipt(ctx, vm.Client, vm.auth.From, tx, methodName)
//}
