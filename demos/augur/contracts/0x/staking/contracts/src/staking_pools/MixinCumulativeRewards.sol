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

import "../../../../utils/contracts/src/LibFractions.sol";
import "../../../../utils/contracts/src/LibSafeMath.sol";
import "../stake/MixinStakeBalances.sol";
import "../immutable/MixinConstants.sol";


contract MixinCumulativeRewards is
    MixinStakeBalances,
    MixinConstants
{
    using LibSafeMath for uint256;

    /// @dev returns true iff Cumulative Rewards are set
    function _isCumulativeRewardSet(IStructs.Fraction memory cumulativeReward)
        internal
        pure
        returns (bool)
    {
        // We use the denominator as a proxy for whether the cumulative
        // reward is set, as setting the cumulative reward always sets this
        // field to at least 1.
        return cumulativeReward.denominator != 0;
    }

    /// @dev Sets a pool's cumulative delegator rewards for the current epoch,
    ///      given the rewards earned and stake from the last epoch, which will
    ///      be summed with the previous cumulative rewards for this pool.
    ///      If the last cumulative reward epoch is the current epoch, this is a
    ///      no-op.
    /// @param poolId The pool ID.
    /// @param reward The total reward earned by pool delegators from the last epoch.
    /// @param stake The total delegated stake in the pool in the last epoch.
    function _addCumulativeReward(
        bytes32 poolId,
        uint256 reward,
        uint256 stake
    )
        internal
    {
        // Fetch the last epoch at which we stored an entry for this pool;
        // this is the most up-to-date cumulative rewards for this pool.
        uint256 lastStoredEpoch = _cumulativeRewardsByPoolLastStored[poolId];
        uint256 currentEpoch_ = currentEpoch;

        // If we already have a record for this epoch, don't overwrite it.
        if (lastStoredEpoch == currentEpoch_) {
            return;
        }

        IStructs.Fraction memory mostRecentCumulativeReward =
            _cumulativeRewardsByPool[poolId][lastStoredEpoch];

        // Compute new cumulative reward
        IStructs.Fraction memory cumulativeReward;
        if (_isCumulativeRewardSet(mostRecentCumulativeReward)) {
            // If we have a prior cumulative reward entry, we sum them as fractions.
            (cumulativeReward.numerator, cumulativeReward.denominator) = LibFractions.add(
                mostRecentCumulativeReward.numerator,
                mostRecentCumulativeReward.denominator,
                reward,
                stake
            );
            // Normalize to prevent overflows in future operations.
            (cumulativeReward.numerator, cumulativeReward.denominator) = LibFractions.normalize(
                cumulativeReward.numerator,
                cumulativeReward.denominator
            );
        } else {
            (cumulativeReward.numerator, cumulativeReward.denominator) = (reward, stake);
        }

        // Store cumulative rewards for this epoch.
        _cumulativeRewardsByPool[poolId][currentEpoch_] = cumulativeReward;
        _cumulativeRewardsByPoolLastStored[poolId] = currentEpoch_;
    }

    /// @dev Sets a pool's cumulative delegator rewards for the current epoch,
    ///      using the last stored cumulative rewards. If we've already set
    ///      a CR for this epoch, this is a no-op.
    /// @param poolId The pool ID.
    function _updateCumulativeReward(bytes32 poolId)
        internal
    {
        // Just add empty rewards for this epoch, which will be added to
        // the previous CR, so we end up with the previous CR being set for
        // this epoch.
        _addCumulativeReward(poolId, 0, 1);
    }

    /// @dev Computes a member's reward over a given epoch interval.
    /// @param poolId Uniqud Id of pool.
    /// @param memberStakeOverInterval Stake delegated to pool by member over
    ///        the interval.
    /// @param beginEpoch Beginning of interval.
    /// @param endEpoch End of interval.
    /// @return rewards Reward accumulated over interval [beginEpoch, endEpoch]
    function _computeMemberRewardOverInterval(
        bytes32 poolId,
        uint256 memberStakeOverInterval,
        uint256 beginEpoch,
        uint256 endEpoch
    )
        internal
        view
        returns (uint256 reward)
    {
        // Sanity check if we can skip computation, as it will result in zero.
        if (memberStakeOverInterval == 0 || beginEpoch == endEpoch) {
            return 0;
        }

        // Sanity check interval
        require(beginEpoch < endEpoch, "CR_INTERVAL_INVALID");

        // Sanity check begin reward
        IStructs.Fraction memory beginReward = _getCumulativeRewardAtEpoch(poolId, beginEpoch);
        IStructs.Fraction memory endReward = _getCumulativeRewardAtEpoch(poolId, endEpoch);

        // Compute reward
        reward = LibFractions.scaleDifference(
            endReward.numerator,
            endReward.denominator,
            beginReward.numerator,
            beginReward.denominator,
            memberStakeOverInterval
        );
    }

    /// @dev Fetch the most recent cumulative reward entry for a pool.
    /// @param poolId Unique ID of pool.
    /// @return cumulativeReward The most recent cumulative reward `poolId`.
    function _getMostRecentCumulativeReward(bytes32 poolId)
        private
        view
        returns (IStructs.Fraction memory cumulativeReward)
    {
        uint256 lastStoredEpoch = _cumulativeRewardsByPoolLastStored[poolId];
        return _cumulativeRewardsByPool[poolId][lastStoredEpoch];
    }

    /// @dev Fetch the cumulative reward for a given epoch.
    ///      If the corresponding CR does not exist in state, then we backtrack
    ///      to find its value by querying `epoch-1` and then most recent CR.
    /// @param poolId Unique ID of pool.
    /// @param epoch The epoch to find the
    /// @return cumulativeReward The cumulative reward for `poolId` at `epoch`.
    /// @return cumulativeRewardStoredAt Epoch that the `cumulativeReward` is stored at.
    function _getCumulativeRewardAtEpoch(bytes32 poolId, uint256 epoch)
        private
        view
        returns (IStructs.Fraction memory cumulativeReward)
    {
        // Return CR at `epoch`, given it's set.
        cumulativeReward = _cumulativeRewardsByPool[poolId][epoch];
        if (_isCumulativeRewardSet(cumulativeReward)) {
            return cumulativeReward;
        }

        // Return CR at `epoch-1`, given it's set.
        uint256 lastEpoch = epoch.safeSub(1);
        cumulativeReward = _cumulativeRewardsByPool[poolId][lastEpoch];
        if (_isCumulativeRewardSet(cumulativeReward)) {
            return cumulativeReward;
        }

        // Return the most recent CR, given it's less than `epoch`.
        uint256 mostRecentEpoch = _cumulativeRewardsByPoolLastStored[poolId];
        if (mostRecentEpoch < epoch) {
            cumulativeReward = _cumulativeRewardsByPool[poolId][mostRecentEpoch];
            if (_isCumulativeRewardSet(cumulativeReward)) {
                return cumulativeReward;
            }
        }

        // Otherwise return an empty CR.
        return IStructs.Fraction(0, 1);
    }
}
