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

import "./Confirming.sol";


contract Leaves is Confirming {
    // invalid leaf
    string constant MOVE_LEAF = "MOVE_LEAF";
    // new stake location is not in path
    string constant MOVE_LOC = "MOVE_LOC";

    // Node is not passed deadline
    string constant RECOV_DEADLINE_TIME = "RECOV_DEADLINE_TIME";
    // Node proof invalid
    string constant RECOV_DEADLINE_PROOF = "RECOV_DEADLINE_PROOF";

    mapping (bytes32 => bool) private leaves;

    event RollupPruned(bytes32 leaf);

    event RollupStakeMoved(
        address staker,
        bytes32 toNodeHash
    );

    function placeStake(
        bytes32 location,
        bytes32 _leaf,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
        payable
    {
        require(isValidLeaf(_leaf), "invalid leaf");
        require(
            RollupUtils.isInPath(
                latestConfirmed(),
                location,
                _leaf,
                proof1,
                proof2
            ),
            PLACE_PATH_PROOF
        );
        createStake(location);
    }

    function moveStake(
        bytes32 newLocation,
        bytes32    _leaf,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(isValidLeaf(_leaf), MOVE_LEAF);
        require(
            RollupUtils.isInPath(
                staker.location,
                newLocation,
                _leaf,
                proof1,
                proof2
            ),
            MOVE_LOC
        );

        staker.location = newLocation;

        emit RollupStakeMoved(msg.sender, newLocation);
    }

    function pruneLeaf(
        bytes32 _leaf,
        bytes32 from,
        bytes32[] calldata leafProof,
        bytes32[] calldata latestConfirmedProof
    )
        external
    {
        require(isValidLeaf(_leaf), "invalid leaf");
        require(
            RollupUtils.isConflict(
                from,
                _leaf,
                latestConfirmed(),
                leafProof,
                latestConfirmedProof
            ),
            "Invalid conflict proof"
        );
        delete leaves[_leaf];

        emit RollupPruned(_leaf);
    }

    // Kick off if successor node whose deadline has passed
    function recoverStakePassedDeadline(
        address payable stakerAddress,
        uint deadlineTicks,
        bytes32 disputableNodeHashVal,
        uint    childType,
        bytes32 vmProtoStateHash,
        bytes32 leaf,
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
        require(block.number >= RollupTime.blocksToTicks(deadlineTicks), RECOV_DEADLINE_TIME);

        require(RollupUtils.isPath(nextNode, leaf, proof), RECOV_DEADLINE_PROOF);
        deleteStakerWithPayout(stakerAddress);

        emit RollupStakeRefunded(stakerAddress);
    }

    function init(
        bytes32 _vmState,
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        Confirming.init(_vmState, _stakeRequirement, _challengeFactoryAddress);
        leaves[latestConfirmed()] = true;
    }

    function isValidLeaf(bytes32 leaf) internal view returns(bool) {
        return leaves[leaf];
    }

    function splitLeaf(bytes32 currentLeaf, bytes32[] memory children) internal {
        delete leaves[currentLeaf];
        for (uint i = 0; i < MAX_CHILD_TYPE; i++) {
            leaves[children[i]] = true;
        }
    }
}
