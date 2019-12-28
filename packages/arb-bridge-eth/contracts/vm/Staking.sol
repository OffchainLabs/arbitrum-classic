/*
 * Copyright 2019, Offchain Labs, Inc.
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

import "./RollupUtils.sol";

import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";


contract Staking {

    uint internal constant VALID_CHILD_TYPE = 0;
    uint internal constant INVALID_PENDING_TOP_CHILD_TYPE = 1;
    uint internal constant INVALID_MESSAGES_CHILD_TYPE = 2;
    uint internal constant INVALID_EXECUTION_CHILD_TYPE = 3;
    uint internal constant MAX_CHILD_TYPE = 3;

    IChallengeFactory public challengeFactory;

    struct Staker {
        bytes32 location;
        uint128 creationTime;
        bool inChallenge;
    }

    uint128 private stakeRequirement;
    mapping(address => Staker) private stakers;
    uint private stakerCount;

    event RollupStakeCreated(
        address staker,
        bytes32 nodeHash,
        uint    blockNumber
    );

    event RollupChallengeStarted(
        address asserter,
        address challenger,
        uint    challengeType,
        address challengeContract
    );

    event RollupChallengeCompleted(
        address challengeContract,
        address winner,
        address loser
    );

    struct ChallengeData {
        address payable[2] stakerAddresses;
        bytes32 node;
        uint disputableDeadline;
        uint[2] stakerPositions;
        bytes32[2] vmProtoHashes;
        bytes32[] proof1;
        bytes32[] proof2;
    }

    struct StartExecutionChallengeData {
        bytes32 beforeHash;
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes32 pendingAssertion;
        bytes32 beforePendingTop;
        bytes32 importedMessageSlice;
        uint32 importedMessageCount;
        bytes32 assertionHash;
    }

    struct StartPendingTopChallengeData {
        bytes32 preconditionHash;
        bytes32 afterPendingTop;
        bytes32 currentPending;
        bytes32 importedAssertion;
        bytes32 assertionHash;
    }

    struct StartMessagesChallengeData {
        bytes32 preconditionHash;
        bytes32 afterPendingTop;
        bytes32 currentPending;
        bytes32 beforePendingTop;
        bytes32 importedMessageSlice;
        uint32 importedMessageCount;
        bytes32 assertionHash;
    }

    function resolveChallenge(address payable winner, address loser) external {
        address sender = msg.sender;
        bytes32 codehash;
        assembly { codehash := extcodehash(sender) }
        address challengeContract = challengeFactory.generateCloneAddress(address(winner), loser, codehash);
        require(challengeContract == msg.sender, "Challenge can only be resolved by spawned contract");
        Staker storage winningStaker = getValidStaker(address(winner));
        winner.transfer(stakeRequirement / 2);
        winningStaker.inChallenge = false;
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, address(winner), loser);
    }

    // fields
    //  node
    //  beforeHash
    //  beforeInbox
    //  pendingAssertionHash
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startExecutionChallenge(
        bytes32[7] calldata _fields,
        address payable[2] calldata stakerAddresses,
        uint disputableDeadline,
        uint[2] calldata stakerPositions,
        bytes32[2] calldata vmProtoHashes,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2,
        uint64[2] calldata _timeBounds,
        uint32 _importedMessageCount
    )
        external
    {
        return _startExecutionChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartExecutionChallengeData(
                _fields[1],
                _fields[2],
                _timeBounds,
                _fields[3],
                _fields[4],
                _fields[5],
                _importedMessageCount,
                _fields[6]
            )
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  importedAssertion
    //  assertionHash

    function startPendingTopChallenge(
        bytes32[6] calldata _fields,
        address payable[2] calldata stakerAddresses,
        uint disputableDeadline,
        uint[2] calldata stakerPositions,
        bytes32[2] calldata vmProtoHashes,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
    {
        return _startPendingTopChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartPendingTopChallengeData(
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _fields[5]
            )
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startMessagesChallenge(
        bytes32[7] calldata _fields,
        address payable[2] calldata stakerAddresses,
        uint disputableDeadline,
        uint[2] calldata stakerPositions,
        bytes32[2] calldata vmProtoHashes,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2,
        uint32 _importedMessageCount
    )
        external
    {
        return _startMessagesChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartMessagesChallengeData(
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _fields[5],
                _importedMessageCount,
                _fields[6]
            )
        );
    }

    function init(
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        require(address(challengeFactory) == address(0), "VM already initialized");
        require(_challengeFactoryAddress != address(0), "Challenge factory must be nonzero");

        challengeFactory = IChallengeFactory(_challengeFactoryAddress);

        // VM parameters
        stakeRequirement = _stakeRequirement;
    }

    function getValidStaker(address _stakerAddress) internal view returns (Staker storage) {
        Staker storage staker = stakers[_stakerAddress];
        require(staker.location != 0x00, "Invalid staker");
        return staker;
    }

    function createStake(
        bytes32 location
    )
        internal
    {
        require(msg.value == stakeRequirement, "must supply stake value");
        require(stakers[msg.sender].location != 0x00, "Staker already exists");
        stakers[msg.sender] = Staker(
            location,
            uint128(block.number),
            false
        );
        stakerCount++;

        emit RollupStakeCreated(msg.sender, location, block.number);
    }

    function deleteStakerWithPayout(address payable _stakerAddress) internal {
        deleteStaker(_stakerAddress);
        _stakerAddress.transfer(stakeRequirement);
    }

    function deleteStaker(address _stakerAddress) private {
        delete stakers[_stakerAddress];
        stakerCount--;
    }

    function getStakerCount() internal view returns(uint) {
        return stakerCount;
    }

    function _startPendingTopChallenge(
        ChallengeData memory _challenge,
        StartPendingTopChallengeData memory data
    )
        private
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_PENDING_TOP_CHILD_TYPE, "Stakers must have a conflict over pending top");

        verifyConflict(
            staker1,
            staker2,
            _challenge.disputableDeadline,
            _challenge.node,
            RollupUtils.disputableNodeHash(
                _challenge.disputableDeadline,
                data.preconditionHash,
                RollupUtils.pendingAssertionHash(
                    data.afterPendingTop,
                    data.currentPending
                ),
                data.importedAssertion,
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createPendingTopChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.currentPending,
            data.afterPendingTop
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startMessagesChallenge(
        ChallengeData memory _challenge,
        StartMessagesChallengeData memory data
    )
        private
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_MESSAGES_CHILD_TYPE, "Stakers must have a conflict over pending top");

        verifyConflict(
            staker1,
            staker2,
            _challenge.disputableDeadline,
            _challenge.node,
            RollupUtils.disputableNodeHash(
                _challenge.disputableDeadline,
                data.preconditionHash,
                RollupUtils.pendingAssertionHash(
                    data.afterPendingTop,
                    data.currentPending
                ),
                RollupUtils.importedAssertionHash(
                    data.beforePendingTop,
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createMessagesChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.beforePendingTop,
            data.afterPendingTop,
            data.importedMessageSlice,
            data.importedMessageCount
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startExecutionChallenge(
        ChallengeData memory _challenge,
        StartExecutionChallengeData memory data
    )
        private
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_EXECUTION_CHILD_TYPE, "Stakers must have a conflict over execution");

        verifyConflict(
            staker1,
            staker2,
            _challenge.disputableDeadline,
            _challenge.node,
            RollupUtils.disputableNodeHash(
                _challenge.disputableDeadline,
                Protocol.generatePreconditionHash(
                    data.beforeHash,
                    data.timeBounds,
                    data.beforeInbox
                ),
                data.pendingAssertion,
                RollupUtils.importedAssertionHash(
                    data.beforePendingTop,
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createExecutionChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.beforeHash,
            Protocol.addMessagesToInbox(data.beforeInbox, data.importedMessageSlice),
            data.timeBounds,
            data.assertionHash
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function verifyConflict(
        Staker storage staker1,
        Staker storage staker2,
        uint disputableDeadline,
        bytes32 node,
        bytes32 disputableNodeHash,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2

    )
        private
        view
    {
        require(staker1.creationTime < disputableDeadline, "staker1 staked after deadline");
        require(staker2.creationTime < disputableDeadline, "staker2 staked after deadline");
        require(!staker1.inChallenge, "staker1 already in a challenge");
        require(!staker2.inChallenge, "staker2 already in a challenge");
        require(stakerPositions[0] < stakerPositions[1], "Child types must be ordered");
        require(stakerPositions[1]<=MAX_CHILD_TYPE, "Invalid child type");
        require(
            RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    node,
                    disputableNodeHash,
                    stakerPositions[0],
                    vmProtoHashes[0]
                ),
                staker1.location,
                proof1
            ) && RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    node,
                    disputableNodeHash,
                    stakerPositions[1],
                    vmProtoHashes[1]
                ),
                staker2.location,
                proof2
            ),
            "Invalid conflict proof"
        );
    }
}
