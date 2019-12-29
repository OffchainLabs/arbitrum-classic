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

    mapping (bytes32 => bool) private leaves;

    event RollupPruned(bytes32 nodeHash);

    event RollupStakeMoved(
        address staker,
        bytes32 toNodeHash
    );

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
