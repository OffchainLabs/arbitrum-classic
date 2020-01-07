/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./NodeGraph.sol";
import "./Staking.sol";


contract ArbRollup is NodeGraph, Staking {

    // invalid path proof
    string constant PLACE_LEAF = "PLACE_LEAF";

    // invalid leaf
    string constant MOVE_LEAF = "MOVE_LEAF";

    // invalid path proof
    string constant RECOV_PATH_PROOF = "RECOV_PATH_PROOF";
    // Invalid conflict proof
    string constant RECOV_CONFLICT_PROOF = "RECOV_CONFLICT_PROOF";
    // Proof must be of nonzero length
    string constant RECVOLD_LENGTH = "RECVOLD_LENGTH";
    // invalid leaf
    string constant RECOV_DEADLINE_LEAF = "RECOV_DEADLINE_LEAF";
    // Node is not passed deadline
    string constant RECOV_DEADLINE_TIME = "RECOV_DEADLINE_TIME";

    // invalid staker location proof
    string constant MAKE_STAKER_PROOF = "MAKE_STAKER_PROOF";

    // Type is not invalid
    string constant CONF_INV_TYPE = "CONF_INV_TYPE";
    // Node is not passed deadline
    string constant CONF_TIME = "CONF_TIME";
    // There must be at least one staker
    string constant CONF_HAS_STAKER = "CONF_HAS_STAKER";

    // Only callable by owner
    string constant ONLY_OWNER = "ONLY_OWNER";

    address owner;

    event ConfirmedAssertion(
        bytes32 logsAccHash
    );

    function init(
        bytes32 _vmState,
        uint128 _gracePeriodTicks,
        uint128 _arbGasSpeedLimitPerTick,
        uint32 _maxExecutionSteps,
        uint128 _stakeRequirement,
        address payable _owner,
        address _challengeFactoryAddress,
        address _globalInboxAddress
    )
        external
    {
        NodeGraph.init(
            _vmState,
            _gracePeriodTicks,
            _arbGasSpeedLimitPerTick,
            _maxExecutionSteps,
            _globalInboxAddress
        );
        Staking.init(
            _stakeRequirement,
            _challengeFactoryAddress
        );
        owner = _owner;
    }

    function placeStake(
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
        payable
    {
        bytes32 location = RollupUtils.calculatePath(latestConfirmed(), proof1);
        bytes32 leaf = RollupUtils.calculatePath(location, proof2);
        require(isValidLeaf(leaf), PLACE_LEAF);
        createStake(location);
    }

    function moveStake(
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
    {
        Staker storage staker = getValidStaker(msg.sender);
        bytes32 newLocation = RollupUtils.calculatePath(staker.location, proof1);
        bytes32 leaf = RollupUtils.calculatePath(newLocation, proof2);
        require(isValidLeaf(leaf), MOVE_LEAF);
        updateStakerLocation(msg.sender, newLocation);
    }

    function recoverStakeConfirmed(bytes32[] calldata proof) external {
        _recoverStakeConfirmed(msg.sender, proof);
    }

    function recoverStakeOld(address payable stakerAddress, bytes32[] calldata proof) external {
        require(proof.length > 0, RECVOLD_LENGTH);
        _recoverStakeConfirmed(stakerAddress, proof);
    }

    function recoverStakeMooted(
        address payable stakerAddress,
        bytes32 node,
        bytes32[] calldata latestConfirmedProof,
        bytes32[] calldata stakerProof
    )
        external
    {
        Staker storage staker = getValidStaker(stakerAddress);
        require(
            latestConfirmedProof[0] != stakerProof[0] &&
            RollupUtils.calculatePath(node, latestConfirmedProof) == latestConfirmed() &&
            RollupUtils.calculatePath(node, stakerProof) == staker.location,
            RECOV_CONFLICT_PROOF
        );
        refundStaker(stakerAddress);
    }

    // Kick off if successor node whose deadline has passed
    function recoverStakePassedDeadline(
        address payable stakerAddress,
        uint256 deadlineTicks,
        bytes32 disputableNodeHashVal,
        uint256 childType,
        bytes32 vmProtoStateHash,
        bytes32[] calldata proof
    )
        external
    {
        Staker storage staker = getValidStaker(stakerAddress);
        bytes32 nextNode = RollupUtils.childNodeHash(
            staker.location,
            deadlineTicks,
            disputableNodeHashVal,
            childType,
            vmProtoStateHash
        );
        bytes32 leaf = RollupUtils.calculatePath(nextNode, proof);
        require(isValidLeaf(leaf), RECOV_DEADLINE_LEAF);
        require(block.number >= RollupTime.blocksToTicks(deadlineTicks), RECOV_DEADLINE_TIME);

        refundStaker(stakerAddress);
    }

    // fields
     // beforeVMHash
     // beforePendingTop
     // prevPrevLeafHash
     // prevDisputableNodeHash
     // afterPendingTop
     // importedMessagesSlice
     // afterVMHash
     // messagesAccHash
     // logsAccHash

    function makeAssertion(
        bytes32[9] calldata _fields,
        uint256 _beforePendingCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint32 _numSteps,
        uint128[2] calldata _timeBoundsBlocks,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas,
        bytes32[] calldata _stakerProof
    )
        external
    {
        (bytes32 prevLeaf, bytes32 newValid) = makeAssertion(
            MakeAssertionData(
                _fields[0],
                _fields[1],
                _beforePendingCount,

                _fields[2],
                _prevDeadlineTicks,
                _fields[3],
                _prevChildType,

                _numSteps,
                _timeBoundsBlocks,
                _importedMessageCount,

                _fields[4],

                _fields[5],

                _fields[6],
                _didInboxInsn,
                _numArbGas,
                _fields[7],
                _fields[8]
            )
        );
        Staker storage staker = getValidStaker(msg.sender);
        require(RollupUtils.calculatePath(staker.location, _stakerProof) == prevLeaf, MAKE_STAKER_PROOF);
        updateStakerLocation(msg.sender, newValid);
    }

    function confirmValid(
        uint256 deadlineTicks,
        bytes calldata _messages,
        bytes32 logsAcc,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint256[] calldata stakerProofOffsets
    )
        external
    {
        _confirmNode(
            deadlineTicks,
            RollupUtils.validNodeHash(
                Protocol.generateLastMessageHash(_messages),
                logsAcc
            ),
            VALID_CHILD_TYPE,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );

        globalInbox.sendMessages(_messages);

        emit ConfirmedAssertion(
            logsAcc
        );
    }

    function confirmInvalid(
        uint256 deadlineTicks,
        bytes32 challengeNodeData,
        uint256 branch,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint256[] calldata stakerProofOffsets
    )
        external
    {
        require(branch < VALID_CHILD_TYPE, CONF_INV_TYPE);
        _confirmNode(
            deadlineTicks,
            challengeNodeData,
            branch,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );
    }

    modifier onlyOwner() {
        require(msg.sender == owner, ONLY_OWNER);
        _;
    }

/*    function activateVM() external onlyOwner {
        if (vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function ownerShutdown() external onlyOwner {
        _shutdown();
    }
    */

    function _recoverStakeConfirmed(address payable stakerAddress, bytes32[] memory proof) private {
        Staker storage staker = getValidStaker(stakerAddress);
        require(RollupUtils.calculatePath(staker.location, proof) == latestConfirmed(), RECOV_PATH_PROOF);
        refundStaker(stakerAddress);
    }

    function _confirmNode(
        uint256 deadlineTicks,
        bytes32 nodeDataHash,
        uint256 branch,
        bytes32 vmProtoStateHash,
        address[] memory stakerAddresses,
        bytes32[] memory stakerProofs,
        uint256[] memory stakerProofOffsets
    )
        private
    {
        bytes32 to = RollupUtils.childNodeHash(
            latestConfirmed(),
            deadlineTicks,
            nodeDataHash,
            branch,
            vmProtoStateHash
        );
        require(RollupTime.blocksToTicks(block.number) >= deadlineTicks, CONF_TIME);
        uint activeCount = checkAlignedStakers(
            to,
            deadlineTicks,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );
        require(activeCount > 0, CONF_HAS_STAKER);

        confirmNode(to);
    }

}
