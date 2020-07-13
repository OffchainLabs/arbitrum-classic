// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

import "./RollupUtils.sol";
import "../libraries/RollupTime.sol";

import "../challenge/ChallengeUtils.sol";
import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";

contract Staking {
    // VM already initialized"
    string private constant INIT_TWICE = "INIT_TWICE";
    // Challenge factory must be nonzero
    string private constant INIT_NONZERO = "INIT_NONZERO";

    // Invalid staker
    string private constant INV_STAKER = "INV_STAKER";

    // must supply stake value
    string private constant STK_AMT = "STK_AMT";
    // Staker already exists
    string private constant ALRDY_STAKED = "ALRDY_STAKED";

    // Challenge can only be resolved by spawned contract
    string private constant RES_CHAL_SENDER = "RES_CHAL_SENDER";

    // staker1 staked after deadline
    string private constant STK1_DEADLINE = "STK1_DEADLINE";
    // staker2 staked after deadline
    string private constant STK2_DEADLINE = "STK2_DEADLINE";
    // staker1 already in a challenge
    string private constant STK1_IN_CHAL = "STK1_IN_CHAL";
    // staker2 already in a challenge
    string private constant STK2_IN_CHAL = "STK1_IN_CHAL";
    // Child types must be ordered
    string private constant TYPE_ORDER = "TYPE_ORDER";
    // Invalid child type
    string private constant INVLD_CHLD_TYPE = "INVLD_CHLD_TYPE";
    // Challenge asserter proof
    string private constant ASSERT_PROOF = "ASSERT_PROOF";
    // Challenge challenger proof
    string private constant CHAL_PROOF = "CHAL_PROOF";

    // must include proof for all stakers
    string private constant CHCK_COUNT = "CHCK_COUNT";
    // Stakers must be ordered
    string private constant CHCK_ORDER = "CHCK_ORDER";
    // at least one active staker disagrees
    string private constant CHCK_STAKER_PROOF = "CHCK_STAKER_PROOF";
    string private constant CHCK_OFFSETS = "CHCK_OFFSETS";

    uint256 private constant MAX_CHILD_TYPE = 3;

    IChallengeFactory public challengeFactory;

    struct Staker {
        bytes32 location;
        uint128 creationTimeBlocks;
        bool inChallenge;
    }

    uint128 private stakeRequirement;
    mapping(address => Staker) private stakers;
    uint256 private stakerCount;
    mapping(address => bool) private challenges;

    event RollupStakeCreated(address staker, bytes32 nodeHash);

    event RollupStakeRefunded(address staker);

    event RollupStakeMoved(address staker, bytes32 toNodeHash);

    event RollupChallengeStarted(
        address asserter,
        address challenger,
        uint256 challengeType,
        address challengeContract
    );

    event RollupChallengeCompleted(
        address challengeContract,
        address winner,
        address loser
    );

    function getStakeRequired() external view returns (uint128) {
        return stakeRequirement;
    }

    function isStaked(address _stakerAddress) external view returns (bool) {
        return stakers[_stakerAddress].location != 0x00;
    }

    function resolveChallenge(address payable winner, address loser) external {
        require(challenges[msg.sender], RES_CHAL_SENDER);
        delete challenges[msg.sender];

        Staker storage winningStaker = getValidStaker(address(winner));
        winner.transfer(stakeRequirement / 2);
        winningStaker.inChallenge = false;
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, address(winner), loser);
    }

    function startChallenge(
        address payable asserterAddress,
        address payable challengerAddress,
        bytes32 prevNode,
        uint256 deadlineTicks,
        uint256[2] memory stakerNodeTypes, // [asserterNodeType, challengerNodeType]
        bytes32[2] memory vmProtoHashes, // [asserterVMProtoHash, challengerVMProtoHash]
        bytes32[] memory asserterProof,
        bytes32[] memory challengerProof,
        bytes32 asserterNodeHash,
        bytes32 challengerDataHash,
        uint128 challengerPeriodTicks
    ) public {
        Staker storage asserter = getValidStaker(asserterAddress);
        Staker storage challenger = getValidStaker(challengerAddress);

        require(
            RollupTime.blocksToTicks(asserter.creationTimeBlocks) <
                deadlineTicks,
            STK1_DEADLINE
        );
        require(
            RollupTime.blocksToTicks(challenger.creationTimeBlocks) <
                deadlineTicks,
            STK2_DEADLINE
        );
        require(!asserter.inChallenge, STK1_IN_CHAL);
        require(!challenger.inChallenge, STK2_IN_CHAL);
        require(stakerNodeTypes[0] > stakerNodeTypes[1], TYPE_ORDER);
        require(
            RollupUtils.calculateLeafFromPath(
                RollupUtils.childNodeHash(
                    prevNode,
                    deadlineTicks,
                    asserterNodeHash,
                    stakerNodeTypes[0],
                    vmProtoHashes[0]
                ),
                asserterProof
            ) == asserter.location,
            ASSERT_PROOF
        );
        require(
            RollupUtils.calculateLeafFromPath(
                RollupUtils.childNodeHash(
                    prevNode,
                    deadlineTicks,
                    RollupUtils.challengeDataHash(
                        challengerDataHash,
                        challengerPeriodTicks
                    ),
                    stakerNodeTypes[1],
                    vmProtoHashes[1]
                ),
                challengerProof
            ) == challenger.location,
            CHAL_PROOF
        );

        asserter.inChallenge = true;
        challenger.inChallenge = true;

        createChallenge(
            asserterAddress,
            challengerAddress,
            challengerPeriodTicks,
            challengerDataHash,
            stakerNodeTypes[1]
        );
    }

    function createChallenge(
        address payable asserterAddress,
        address payable challengerAddress,
        uint128 challengerPeriodTicks,
        bytes32 challengerDataHash,
        uint256 stakerNodeType
    ) internal {
        address newChallengeAddr = challengeFactory.createChallenge(
            asserterAddress,
            challengerAddress,
            challengerPeriodTicks,
            challengerDataHash,
            stakerNodeType
        );

        challenges[newChallengeAddr] = true;

        emit RollupChallengeStarted(
            asserterAddress,
            challengerAddress,
            stakerNodeType,
            newChallengeAddr
        );
    }

    function init(uint128 _stakeRequirement, address _challengeFactoryAddress)
        internal
    {
        require(address(challengeFactory) == address(0), INIT_TWICE);
        require(_challengeFactoryAddress != address(0), INIT_NONZERO);

        challengeFactory = IChallengeFactory(_challengeFactoryAddress);

        // VM parameters
        stakeRequirement = _stakeRequirement;
    }

    function getStakerLocation(address _stakerAddress)
        internal
        view
        returns (bytes32)
    {
        bytes32 location = stakers[_stakerAddress].location;
        require(location != 0x00, INV_STAKER);
        return location;
    }

    function createStake(bytes32 location) internal {
        require(msg.value == stakeRequirement, STK_AMT);
        require(stakers[msg.sender].location == 0x00, ALRDY_STAKED);
        stakers[msg.sender] = Staker(location, uint128(block.number), false);
        stakerCount++;

        emit RollupStakeCreated(msg.sender, location);
    }

    function updateStakerLocation(address _stakerAddress, bytes32 _location)
        internal
    {
        stakers[_stakerAddress].location = _location;
        emit RollupStakeMoved(_stakerAddress, _location);
    }

    function refundStaker(address payable _stakerAddress) internal {
        deleteStaker(_stakerAddress);
        _stakerAddress.transfer(stakeRequirement);

        emit RollupStakeRefunded(address(_stakerAddress));
    }

    function getValidStaker(address _stakerAddress)
        private
        view
        returns (Staker storage)
    {
        Staker storage staker = stakers[_stakerAddress];
        require(staker.location != 0x00, INV_STAKER);
        return staker;
    }

    function deleteStaker(address _stakerAddress) private {
        delete stakers[_stakerAddress];
        stakerCount--;
    }

    function checkAlignedStakers(
        bytes32 node,
        uint256 deadlineTicks,
        address[] memory stakerAddresses,
        bytes32[] memory stakerProofs,
        uint256[] memory stakerProofOffsets
    ) internal view returns (uint256) {
        uint256 _stakerCount = stakerAddresses.length;
        require(_stakerCount == stakerCount, CHCK_COUNT);
        require(_stakerCount + 1 == stakerProofOffsets.length, CHCK_OFFSETS);

        bytes20 prevStaker = 0x00;
        uint256 activeCount = 0;
        bool isActive = false;

        for (uint256 index = 0; index < _stakerCount; index++) {
            address currentStaker = stakerAddresses[index];

            isActive = _verifyAlignedStaker(
                node,
                stakerProofs,
                deadlineTicks,
                currentStaker,
                prevStaker,
                stakerProofOffsets[index],
                stakerProofOffsets[index + 1]
            );

            if (isActive) {
                activeCount++;
            }

            prevStaker = bytes20(currentStaker);
        }
        return activeCount;
    }

    function _verifyAlignedStaker(
        bytes32 node,
        bytes32[] memory stakerProofs,
        uint256 deadlineTicks,
        address stakerAddress,
        bytes20 prevStaker,
        uint256 proofStart,
        uint256 proofEnd
    ) private view returns (bool) {
        require(bytes20(stakerAddress) > prevStaker, CHCK_ORDER);
        Staker storage staker = getValidStaker(stakerAddress);
        bool isActive = RollupTime.blocksToTicks(staker.creationTimeBlocks) <
            deadlineTicks;

        if (isActive) {
            require(
                RollupUtils.calculateLeafFromPath(
                    node,
                    stakerProofs,
                    proofStart,
                    proofEnd
                ) == staker.location,
                CHCK_STAKER_PROOF
            );
        }

        return isActive;
    }
}
