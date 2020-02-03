/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;


/// @dev Exposes some internal functions from various contracts to avoid
///      cyclical dependencies.
contract MixinAbstract {

    /// @dev Computes the reward owed to a pool during finalization.
    ///      Does nothing if the pool is already finalized.
    /// @param poolId The pool's ID.
    /// @return totalReward The total reward owed to a pool.
    /// @return membersStake The total stake for all non-operator members in
    ///         this pool.
    function _getUnfinalizedPoolRewards(bytes32 poolId)
        internal
        view
        returns (
            uint256 totalReward,
            uint256 membersStake
        );

    /// @dev Asserts that a pool has been finalized last epoch.
    /// @param poolId The id of the pool that should have been finalized.
    function _assertPoolFinalizedLastEpoch(bytes32 poolId)
        internal
        view;
}
