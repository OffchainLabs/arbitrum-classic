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

package gobridge

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

const VALID_CHILD_TYPE = 3

type staker struct {
	location           common.Hash
	creationTimeBlocks *common.TimeBlocks
	inChallenge        bool
	stakeAmount        *big.Int
}

type rollupData struct {
	initVMHash    common.Hash
	VMstate       machine.Status
	state         EthState
	chainParams   valprotocol.ChainParams
	owner         common.Address
	events        map[*common.BlockId][]arbbridge.Event
	creation      *common.BlockId
	stakers       map[common.Address]*staker
	leaves        map[common.Hash]bool
	lastConfirmed common.Hash
	nextConfirmed common.Hash
}

type arbRollup struct {
	*rollupData
	client          *GoArbAuthClient
	contractAddress common.Address
}

func getRollupContract(contractAddress common.Address, client *GoArbAuthClient) (*arbRollup, error) {
	return client.rollups[contractAddress], nil
}

func newRollup(con *arbFactory,
	address common.Address,
	vmState common.Hash,
	params valprotocol.ChainParams,
	owner common.Address,
) {
	events := make(map[*common.BlockId][]arbbridge.Event)
	vmProto := hashing.SoliditySHA3(
		hashing.Bytes32(vmState),
		hashing.Bytes32(value.NewEmptyTuple().Hash()),
		hashing.Uint256(big.NewInt(0)),
	)
	innerHash := hashing.SoliditySHA3(
		hashing.Bytes32(vmProto),
		hashing.Uint256(big.NewInt(0)),
		hashing.Uint256(big.NewInt(0)),
		hashing.Uint256(big.NewInt(0)),
	)
	initialNode := hashing.SoliditySHA3(
		hashing.Uint256(big.NewInt(0)),
		hashing.Bytes32(innerHash),
	)
	ruData := &rollupData{
		initVMHash:    vmState,
		VMstate:       machine.Extensive,
		state:         Uninitialized,
		chainParams:   params,
		owner:         owner,
		events:        events,
		creation:      con.client.getCurrentBlock(),
		stakers:       make(map[common.Address]*staker),
		leaves:        make(map[common.Hash]bool),
		lastConfirmed: initialNode,
	}
	con.client.rollups[address] = &arbRollup{
		client:          con.client,
		rollupData:      ruData,
		contractAddress: address,
	}
	con.client.rollups[address].leaves[initialNode] = true

}

func (vm *arbRollup) PlaceStake(ctx context.Context, stakeAmount *big.Int, proof1 []common.Hash, proof2 []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()
	location := calculatePath(vm.lastConfirmed, proof1)
	leaf := calculatePath(location, proof2)
	if !vm.leaves[leaf] {
		return errors.New("invalid path proof")
	}
	if err := createStake(vm, stakeAmount, location); err != nil {
		return err
	}

	event := arbbridge.StakeCreatedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: vm.client.getCurrentBlock(),
		},
		Staker:   vm.client.fromAddr,
		NodeHash: location,
	}
	vm.client.pubMsg(vm.contractAddress, event)
	return nil
}

func createStake(vm *arbRollup, stakeAmount *big.Int, location common.Hash) error {
	if stakeAmount != vm.chainParams.StakeRequirement {
		return errors.New("invalid stake amount")
	}
	if _, ok := vm.stakers[vm.client.fromAddr]; ok {
		return errors.New("staker already exists")
	}
	err := transferEth(vm.client.goEthdata, vm.contractAddress, vm.client.fromAddr, stakeAmount)
	if err != nil {
		return err
	}
	vm.stakers[vm.client.fromAddr] = &staker{location, vm.client.getCurrentBlock().Height, false, stakeAmount}

	return nil
}

func refundStaker(vm *arbRollup, staker common.Address) error {
	if err := transferEth(vm.client.goEthdata, staker, vm.contractAddress, vm.stakers[staker].stakeAmount); err != nil {
		return err
	}
	delete(vm.stakers, staker)

	vm.client.pubMsg(vm.contractAddress,
		arbbridge.StakeRefundedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: vm.client.getCurrentBlock(),
			},
			Staker: staker,
		})
	return nil
}

func (vm *arbRollup) RecoverStakeConfirmed(ctx context.Context, proof []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	staker, ok := vm.stakers[vm.client.fromAddr]
	if !ok {
		return errors.New("staker not found")
	}

	if calculatePath(staker.location, proof) != vm.lastConfirmed {
		return errors.New("invalid path proof")
	}

	// refundStaker
	if err := refundStaker(vm, vm.client.fromAddr); err != nil {
		return err
	}

	//emit RollupStakeRefunded(contractAddress(_stakerAddress));
	vm.client.pubMsg(vm.contractAddress,
		arbbridge.StakeRefundedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: vm.client.getCurrentBlock(),
			},
			Staker: vm.client.fromAddr,
		},
	)

	return nil
}

func (vm *arbRollup) RecoverStakeOld(ctx context.Context, staker common.Address, proof []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()
	if len(proof) <= 0 {
		return errors.New("proof must be non-zero length")
	}

	st, ok := vm.stakers[staker]
	if !ok {
		return errors.New("staker not found")
	}
	if calculatePath(st.location, proof) != vm.lastConfirmed {
		return errors.New("invalid path proof")
	}
	if err := refundStaker(vm, staker); err != nil {
		return err
	}

	return nil
}

func (vm *arbRollup) RecoverStakeMooted(ctx context.Context, nodeHash common.Hash, staker common.Address, latestConfirmedProof []common.Hash, stakerProof []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	if latestConfirmedProof[0] == stakerProof[0] ||
		calculatePath(nodeHash, latestConfirmedProof) == vm.lastConfirmed ||
		calculatePath(nodeHash, stakerProof) != vm.stakers[vm.client.fromAddr].location {
		return errors.New("Invalid conflict proof")
	}
	if err := refundStaker(vm, staker); err != nil {
		return err
	}

	return nil
}

func (vm *arbRollup) RecoverStakePassedDeadline(
	ctx context.Context,
	stakerAddress common.Address,
	deadlineTicks *big.Int,
	disputableNodeHashVal common.Hash,
	childType uint64,
	vmProtoStateHash common.Hash,
	proof []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	nextNode, _ := valprotocol.NodeHash(
		vm.stakers[stakerAddress].location,
		vmProtoStateHash,
		common.TimeTicks{deadlineTicks},
		disputableNodeHashVal,
		valprotocol.ChildType(childType),
	)

	leaf := calculatePath(nextNode, proof)
	if !vm.leaves[leaf] {
		return errors.New("invalid leaf")
	}

	if common.TicksFromBlockNum(vm.client.getCurrentBlock().Height).Val.Cmp(deadlineTicks) < 0 {
		return errors.New("Node is not passed deadline")
	}
	if err := refundStaker(vm, stakerAddress); err != nil {
		return err
	}

	return nil
}

func (vm *arbRollup) MoveStake(ctx context.Context, proof1 []common.Hash, proof2 []common.Hash) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	staker, ok := vm.stakers[vm.client.fromAddr]
	if !ok {
		return errors.New("invalid staker address")
	}
	location := staker.location
	newLocation := calculatePath(location, proof1)
	leaf := calculatePath(newLocation, proof2)
	if !vm.leaves[leaf] {
		return errors.New("MoveStake - invalid leaf")
	}
	vm.stakers[vm.client.fromAddr].location = newLocation
	//emit RollupStakeRefunded(contractAddress(_stakerAddress));
	vm.client.pubMsg(vm.contractAddress,
		arbbridge.StakeRefundedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: vm.client.getCurrentBlock(),
			},
			Staker: vm.client.fromAddr,
		},
	)

	return nil
}

func (vm *arbRollup) pruneLeaf(ctx context.Context, from common.Hash, leafProof []common.Hash, ancProof []common.Hash) error {
	leaf := calculatePath(from, leafProof)
	if !vm.leaves[leaf] {
		return errors.New("pruneLeaf - invalid leaf")
	}
	if leafProof[0] == ancProof[0] || calculatePath(from, ancProof) != vm.lastConfirmed {
		return errors.New("prune conflict")
	}
	delete(vm.leaves, leaf)

	vm.client.pubMsg(vm.contractAddress,
		arbbridge.PrunedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: vm.client.getCurrentBlock(),
			},
			Leaf: leaf,
		},
	)

	return nil
}

func (vm *arbRollup) PruneLeaves(ctx context.Context, opps []valprotocol.PruneParams) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()
	for _, opp := range opps {
		err := vm.pruneLeaf(ctx, opp.AncestorHash, opp.LeafProof, opp.AncProof)
		if err != nil {
			return err
		}
	}

	return nil
}

func (vm *arbRollup) MakeAssertion(
	ctx context.Context,

	prevPrevLeafHash common.Hash,
	prevDataHash common.Hash,
	prevDeadline common.TimeTicks,
	prevChildType valprotocol.ChildType,

	beforeState *valprotocol.VMProtoData,
	assertionParams *valprotocol.AssertionParams,
	assertionClaim *valprotocol.AssertionClaim,
	stakerProof []common.Hash,
) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	protoHashBefore := beforeState.Hash()
	prevLeaf, _ := valprotocol.NodeHash(prevPrevLeafHash,
		protoHashBefore,
		prevDeadline,
		prevDataHash,
		prevChildType,
	)
	if !vm.leaves[prevLeaf] {
		return errors.New("makeAssertion - invalid leaf")
	}
	if vm.VMstate == machine.ErrorStop || vm.VMstate == machine.Halt {
		return errors.New("Can only disputable assert if machine is not errored or halted")
	}
	if assertionParams.NumSteps > vm.chainParams.MaxExecutionSteps {
		return errors.New("makeAssertion - Tried to execute too many steps")
	}
	if assertionParams.TimeBounds.IsValidTime(vm.client.getCurrentBlock().Height) != nil {
		return errors.New("makeAssertion - Precondition: not within time bounds")
	}
	if assertionParams.ImportedMessageCount.Cmp(big.NewInt(0)) != 0 && !assertionClaim.AssertionStub.DidInboxInsn {
		return errors.New("makeAssertion - Imported messages without reading them")
	}
	inboxPending, inboxExists := vm.client.globalInbox.inbox[vm.contractAddress]
	if inboxExists { // valid inbox - check size
		if assertionParams.ImportedMessageCount.Cmp(inboxPending.count.Sub(inboxPending.count, beforeState.InboxCount)) > 0 {
			return errors.New("makeAssertion - Tried to import more messages than exist in pending inbox")
		}
	} else { //no inbox - verify assertion has no messages
		if assertionParams.ImportedMessageCount.Cmp(big.NewInt(0)) != 0 {
			return errors.New("makeAssertion - Tried to import more messages when none exist in pending inbox")
		}
	}

	currentTicks := common.TicksFromBlockNum(vm.client.getCurrentBlock().Height)
	deadlineTicks := currentTicks.Add(vm.chainParams.GracePeriod)
	if deadlineTicks.Cmp(prevDeadline) < 0 {
		return errors.New("Node is not passed deadline")
	}
	checkTimeTicks := assertionClaim.AssertionStub.NumGas / vm.chainParams.ArbGasSpeedLimitPerTick
	deadlineTicks = deadlineTicks.Add(common.TicksFromSeconds(int64(checkTimeTicks)))

	protoStateHash := hashing.SoliditySHA3(
		hashing.Bytes32(assertionClaim.AssertionStub.AfterHash),
		hashing.Bytes32(assertionClaim.AfterInboxTop),
		hashing.Uint256(beforeState.InboxCount.Add(beforeState.InboxCount, assertionParams.ImportedMessageCount)),
	)
	protoData := hashing.SoliditySHA3(
		hashing.Bytes32(assertionClaim.AssertionStub.LastMessageHash),
		hashing.Bytes32(assertionClaim.AssertionStub.LastLogHash),
	)

	//build invalidPendingInbox node
	var pendingTopCount *big.Int
	var pendingTopHash common.Hash
	if !inboxExists {
		pendingTopCount = big.NewInt(0)
		pendingTopHash = value.NewEmptyTuple().Hash()
	} else {
		pendingTopCount = inboxPending.count
		pendingTopHash = inboxPending.value
	}
	left := new(big.Int).Add(beforeState.InboxCount, assertionParams.ImportedMessageCount)
	left = left.Sub(pendingTopCount, left)
	invPendingChallengeDataHash := valprotocol.InboxTopChallengeDataHash(
		assertionClaim.AfterInboxTop,
		pendingTopHash,
		left,
	)
	ticks := vm.chainParams.GracePeriod.Add(common.TicksFromBlockNum(common.NewTimeBlocks(big.NewInt(1))))
	invPendingProtoData := hashing.SoliditySHA3(
		hashing.Bytes32(invPendingChallengeDataHash),
		hashing.TimeTicks(ticks),
	)
	invalidPending, _ := valprotocol.NodeHash(prevLeaf,
		protoHashBefore,
		deadlineTicks,
		invPendingProtoData,
		valprotocol.InvalidInboxTopChildType)

	// build invalid messages node
	invMsgsChallengeDataHash := valprotocol.MessageChallengeDataHash(
		beforeState.InboxTop,
		assertionClaim.AfterInboxTop,
		value.NewEmptyTuple().Hash(),
		assertionClaim.ImportedMessagesSlice,
		assertionParams.ImportedMessageCount,
	)
	invMsgsProtoData := hashing.SoliditySHA3(
		hashing.Bytes32(invMsgsChallengeDataHash),
		hashing.TimeTicks(vm.chainParams.GracePeriod.Add(common.TicksFromBlockNum(common.NewTimeBlocks(big.NewInt(1))))),
	)
	invalidMessages, _ := valprotocol.NodeHash(prevLeaf,
		protoHashBefore,
		deadlineTicks,
		invMsgsProtoData,
		valprotocol.InvalidMessagesChildType)

	// build invalidExecutions node
	invExecChallengeDataHash := valprotocol.ExecutionDataHash(
		assertionParams.NumSteps,
		valprotocol.ExecutionPreconditionHash(beforeState.MachineHash, assertionParams.TimeBounds, assertionClaim.ImportedMessagesSlice),
		assertionClaim.AssertionStub.Hash(),
	)
	invExecProtoData := hashing.SoliditySHA3(
		hashing.Bytes32(invExecChallengeDataHash),
		hashing.TimeTicks(vm.chainParams.GracePeriod.Add(common.TimeTicks{new(big.Int).SetUint64(assertionClaim.AssertionStub.NumGas / vm.chainParams.ArbGasSpeedLimitPerTick)})),
	)
	invalidExecution, _ := valprotocol.NodeHash(prevLeaf,
		protoHashBefore,
		deadlineTicks,
		invExecProtoData,
		valprotocol.InvalidExecutionChildType,
	)

	// build valid node
	valid, _ := valprotocol.NodeHash(prevLeaf,
		protoStateHash,
		deadlineTicks,
		protoData,
		valprotocol.ValidChildType,
	)

	vm.leaves[invalidPending] = true
	vm.leaves[invalidMessages] = true
	vm.leaves[invalidExecution] = true
	vm.leaves[valid] = true
	delete(vm.leaves, prevLeaf)

	event := arbbridge.AssertedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: vm.client.getCurrentBlock(),
		},
		PrevLeafHash:  prevLeaf,
		Params:        assertionParams,
		Claim:         assertionClaim,
		MaxInboxTop:   beforeState.InboxTop,
		MaxInboxCount: beforeState.InboxCount,
	}
	vm.client.pubMsg(vm.contractAddress, event)

	if calculatePath(vm.stakers[vm.client.fromAddr].location, stakerProof) != prevLeaf {
		return errors.New("invalid staker location proof")
	}
	vm.stakers[vm.client.fromAddr].location = valid
	vm.nextConfirmed = valid
	stakeMovedEvent := arbbridge.StakeMovedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: vm.client.getCurrentBlock(),
		},
		Staker:   vm.client.fromAddr,
		Location: valid,
	}
	vm.client.pubMsg(vm.contractAddress, stakeMovedEvent)

	return nil
}

func (vm *arbRollup) Confirm(ctx context.Context, opp *valprotocol.ConfirmOpportunity) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()

	nodeOpps := opp.Nodes
	nodeCount := len(nodeOpps)
	lastMsg := common.Hash{}
	initalProtoStateHash := nodeOpps[0].StateHash()
	var lastLogHash common.Hash
	var nodeDataHash common.Hash
	var messages []value.Value
	validNum := 0
	invalidNum := 0
	confNode := opp.CurrentLatestConfirmed
	for _, opp := range nodeOpps {
		var linkType valprotocol.ChildType
		switch opp := opp.(type) {
		case valprotocol.ConfirmValidOpportunity:
			{
				linkType = VALID_CHILD_TYPE
				lastLogHash = opp.LogsAcc
				messages = opp.Messages
				if len(opp.Messages) > 0 {
					lastMsg = opp.Messages[len(opp.Messages)-1].Hash()
				}
				nodeDataHash = hashing.SoliditySHA3(
					hashing.Bytes32(lastMsg),
					hashing.Bytes32(lastLogHash),
				)
				validNum++
			}
		case valprotocol.ConfirmInvalidOpportunity:
			linkType = valprotocol.MaxInvalidChildType
			nodeDataHash = opp.ChallengeNodeData
			invalidNum++
		}
		confNode, _ = valprotocol.NodeHash(
			confNode,
			initalProtoStateHash,
			opp.Deadline(),
			nodeDataHash,
			linkType,
		)
	}

	if common.TicksFromBlockNum(vm.client.LastMinedBlock.Height).Cmp(opp.Nodes[nodeCount-1].Deadline()) == -1 {
		//panic("Node is not passed deadline")
		return errors.New("node is not passed deadline")
	}

	activeCount := 0
	for i, staker := range opp.StakerAddresses {
		if !calculatePath(confNode, opp.StakerProofs[i]).Equals(vm.stakers[staker].location) {
			return errors.New("at least one active staker disagrees")
		}
		activeCount++
	}

	if activeCount == 0 {
		return errors.New("There must be at least one staker")
	}
	vm.lastConfirmed = confNode

	confirmedEvent := arbbridge.ConfirmedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: vm.client.getCurrentBlock(),
		},
		NodeHash: confNode,
	}
	vm.client.pubMsg(vm.contractAddress, confirmedEvent)

	for _, msg := range messages {
		msgTypeVal, _ := msg.(value.TupleValue).GetByInt64(0)
		msgTypeInt, ok := msgTypeVal.(value.IntValue)
		if !ok {
			return errors.New("msg type must be an int")
		}
		switch message.MessageType(msgTypeInt.BigInt().Uint64()) {
		case message.EthType:
			ethMsg, err := message.UnmarshalEth(msg)
			if err != nil {
				return errors.New("invalid message")
			}
			err = transferEth(vm.client.goEthdata, ethMsg.To, ethMsg.From, ethMsg.Value)
			if err != nil {
				return err
			}
		case message.ERC20Type:
			//transfer ERC20(msg.sender, to, erc20, value)
			erc20Msg, err := message.UnmarshalERC20(msg)
			if err != nil {
				return errors.New("invalid message")
			}
			if !transferERC20(vm.client.goEthdata, erc20Msg) {
				return errors.New("token transfer error")
			}

		case message.ERC721Type:
			erc721Msg, err := message.UnmarshalERC721(msg)
			if err != nil {
				return errors.New("invalid message")
			}
			if !transferNFTToken(vm.client.goEthdata, erc721Msg) {
				return errors.New("token transfer error")
			}

		}
	}

	if validNum > 0 {
		confirmedAssertionEvent := arbbridge.ConfirmedAssertionEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: vm.client.getCurrentBlock(),
			},
		}
		vm.client.pubMsg(vm.contractAddress, confirmedAssertionEvent)
	}
	return nil
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
) error {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()
	eth := vm.client
	asserter, ok := vm.stakers[asserterAddress]
	if !ok {
		return errors.New("unknown asserter")
	}
	challenger, ok := vm.stakers[challengerAddress]
	if !ok {
		return errors.New("unknown challenger")
	}

	if asserter.creationTimeBlocks.AsInt().Cmp(disputableDeadline) == 1 {
		return errors.New("asserter staked after deadline")
	}
	if challenger.creationTimeBlocks.AsInt().Cmp(disputableDeadline) == 1 {
		return errors.New("challenger staked after deadline")
	}
	if asserter.inChallenge {
		return errors.New("asserter already in challenge")
	}
	if challenger.inChallenge {
		return errors.New("challenger already in challenge")
	}
	if asserterPosition <= challengerPosition {
		return errors.New("Child types must be ordere")
	}

	assnodedata, _ := valprotocol.NodeHash(
		prevNode,
		asserterVMProtoHash,
		common.TimeTicks{disputableDeadline},
		asserterNodeHash,
		asserterPosition,
	)
	if calculatePath(assnodedata, asserterProof) != asserter.location {
		return errors.New("Challenge asserter proof error")
	}

	chalnodedata, _ := valprotocol.NodeHash(
		prevNode,
		challengerVMProtoHash,
		common.TimeTicks{disputableDeadline},
		challengerDataHash,
		challengerPosition,
	)
	if calculatePath(chalnodedata, challengerProof) != challenger.location {
		return errors.New("Challenge challenger proof error")
	}

	asserter.inChallenge = true
	challenger.inChallenge = true

	// generate contractAddress
	challengeAddress, _ := eth.challengeFactory.CreateChallenge(
		ctx,
		asserterAddress,
		challengerAddress,
		challengerPeriodTicks,
		challengerDataHash,
		new(big.Int).SetUint64(uint64(challengerPosition)),
	)
	eth.challenges[challengeAddress].rollupAddr = vm.contractAddress

	eth.pubMsg(vm.contractAddress, arbbridge.InitiateChallengeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: eth.getCurrentBlock(),
		},
		Deadline: eth.challenges[challengeAddress].deadline,
	})

	eth.pubMsg(vm.contractAddress, arbbridge.ChallengeStartedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: eth.getCurrentBlock(),
		},
		Asserter:          asserterAddress,
		Challenger:        challengerAddress,
		ChallengeType:     asserterPosition,
		ChallengeContract: challengeAddress,
	})
	return nil
}

func (vm *arbRollup) IsStaked(address common.Address) (bool, error) {
	vm.client.goEthMutex.Lock()
	defer vm.client.goEthMutex.Unlock()
	_, ok := vm.stakers[address]
	if !ok {
		return false, nil
	}

	return true, nil
}
