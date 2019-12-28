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

import "./Staking.sol";


contract Confirming is Staking {
    bytes32 private latestConfirmedPriv;

    event RollupStakeRefunded(address staker);

    function init(
        bytes32 _vmState,
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        Staking.init(_stakeRequirement, _challengeFactoryAddress);

        // VM protocol state
        bytes32 vmProtoStateHash = RollupUtils.protoStateHash(_vmState, Value.hashEmptyTuple(), Value.hashEmptyTuple());
        updateLatestConfirmed(RollupUtils.childNodeHash(
            0,
            0,
            0,
            vmProtoStateHash
        ));
    }

    function placeStake(
        bytes32 location,
        bytes32[] calldata proof
    )
        external
        payable
    {
        require(RollupUtils.isPath(latestConfirmed(), location, proof), "invalid path proof");
        // TODO: Also check if location is on path to leaf?
        createStake(location);
    }

    function recoverStakeConfirmed(
        bytes32[] calldata proof
    )
        external
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(RollupUtils.isPath(staker.location, latestConfirmed(), proof), "invalid path proof");
        deleteStakerWithPayout(msg.sender);

        emit RollupStakeRefunded(msg.sender);
    }

    function recoverStakeMooted(
        bytes32 disputableHash,
        bytes32[] calldata latestConfirmedProof,
        bytes32[] calldata nodeProof
    )
        external
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(
            RollupUtils.isConflict(
                staker.location,
                disputableHash,
                latestConfirmed(),
                latestConfirmedProof,
                nodeProof
            ),
            "Invalid conflict proof"
        );
        deleteStakerWithPayout(msg.sender);

        emit RollupStakeRefunded(msg.sender);
    }

    function latestConfirmed() internal view returns (bytes32) {
        return latestConfirmedPriv;
    }

    function updateLatestConfirmed(bytes32 node) internal {
        latestConfirmedPriv = node;
    }
}
