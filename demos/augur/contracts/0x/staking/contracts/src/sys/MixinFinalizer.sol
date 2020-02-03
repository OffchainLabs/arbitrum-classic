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


import "../../../../utils/contracts/src/LibSafeMath.sol";
import "../libs/LibCobbDouglas.sol";
import "../libs/LibStakingRichErrors.sol";
import "../interfaces/IStructs.sol";
import "../staking_pools/MixinStakingPoolRewards.sol";


contract MixinFinalizer is
    MixinStakingPoolRewards
{
    using LibSafeMath for uint256;

    /// @dev Begins a new epoch, preparing the prior one for finalization.
    ///      Throws if not enough time has passed between epochs or if the
    ///      previous epoch was not fully finalized.
    /// @return numPoolsToFinalize The number of unfinalized pools.
    function endEpoch()
        external
        returns (uint256)
    {
        uint256 currentEpoch_ = currentEpoch;
        uint256 prevEpoch = currentEpoch_.safeSub(1);

        // Make sure the previous epoch has been fully finalized.
        uint256 numPoolsToFinalizeFromPrevEpoch = aggregatedStatsByEpoch[prevEpoch].numPoolsToFinalize;
        if (numPoolsToFinalizeFromPrevEpoch != 0) {
            revert();
        }

        // Convert all ETH to WETH; the WETH balance of this contract is the total rewards.
        _wrapEth();

        // Load aggregated stats for the epoch we're ending.
        aggregatedStatsByEpoch[currentEpoch_].rewardsAvailable = _getAvailableWethBalance();
        IStructs.AggregatedStats memory aggregatedStats = aggregatedStatsByEpoch[currentEpoch_];

        // Emit an event.
        emit EpochEnded(
            currentEpoch_,
            aggregatedStats.numPoolsToFinalize,
            aggregatedStats.rewardsAvailable,
            aggregatedStats.totalFeesCollected,
            aggregatedStats.totalWeightedStake
        );

        // Advance the epoch. This will revert if not enough time has passed.
        _goToNextEpoch();

        // If there are no pools to finalize then the epoch is finalized.
        if (aggregatedStats.numPoolsToFinalize == 0) {
            emit EpochFinalized(currentEpoch_, 0, aggregatedStats.rewardsAvailable);
        }

        return aggregatedStats.numPoolsToFinalize;
    }

    /// @dev Instantly finalizes a single pool that earned rewards in the previous
    ///      epoch, crediting it rewards for members and withdrawing operator's
    ///      rewards as WETH. This can be called by internal functions that need
    ///      to finalize a pool immediately. Does nothing if the pool is already
    ///      finalized or did not earn rewards in the previous epoch.
    /// @param poolId The pool ID to finalize.
    function finalizePool(bytes32 poolId)
        external
    {
        // Compute relevant epochs
        uint256 currentEpoch_ = currentEpoch;
        uint256 prevEpoch = currentEpoch_.safeSub(1);

        // Load the aggregated stats into memory; noop if no pools to finalize.
        IStructs.AggregatedStats memory aggregatedStats = aggregatedStatsByEpoch[prevEpoch];
        if (aggregatedStats.numPoolsToFinalize == 0) {
            return;
        }

        // Noop if the pool did not earn rewards or already finalized (has no fees).
        IStructs.PoolStats memory poolStats = poolStatsByEpoch[poolId][prevEpoch];
        if (poolStats.feesCollected == 0) {
            return;
        }

        // Clear the pool stats so we don't finalize it again, and to recoup
        // some gas.
        delete poolStatsByEpoch[poolId][prevEpoch];

        // Compute the rewards.
        uint256 rewards = _getUnfinalizedPoolRewardsFromPoolStats(poolStats, aggregatedStats);

        // Pay the operator and update rewards for the pool.
        // Note that we credit at the CURRENT epoch even though these rewards
        // were earned in the previous epoch.
        (uint256 operatorReward, uint256 membersReward) = _syncPoolRewards(
            poolId,
            rewards,
            poolStats.membersStake
        );

        // Emit an event.
        emit RewardsPaid(
            currentEpoch_,
            poolId,
            operatorReward,
            membersReward
        );

        uint256 totalReward = operatorReward.safeAdd(membersReward);

        // Increase `totalRewardsFinalized`.
        aggregatedStatsByEpoch[prevEpoch].totalRewardsFinalized =
            aggregatedStats.totalRewardsFinalized =
            aggregatedStats.totalRewardsFinalized.safeAdd(totalReward);

        // Decrease the number of unfinalized pools left.
        aggregatedStatsByEpoch[prevEpoch].numPoolsToFinalize =
            aggregatedStats.numPoolsToFinalize =
            aggregatedStats.numPoolsToFinalize.safeSub(1);

        // If there are no more unfinalized pools remaining, the epoch is
        // finalized.
        if (aggregatedStats.numPoolsToFinalize == 0) {
            emit EpochFinalized(
                prevEpoch,
                aggregatedStats.totalRewardsFinalized,
                aggregatedStats.rewardsAvailable.safeSub(aggregatedStats.totalRewardsFinalized)
            );
        }
    }

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
            uint256 reward,
            uint256 membersStake
        )
    {
        uint256 prevEpoch = currentEpoch.safeSub(1);
        IStructs.PoolStats memory poolStats = poolStatsByEpoch[poolId][prevEpoch];
        reward = _getUnfinalizedPoolRewardsFromPoolStats(poolStats, aggregatedStatsByEpoch[prevEpoch]);
        membersStake = poolStats.membersStake;
    }

    /// @dev Converts the entire ETH balance of this contract into WETH.
    function _wrapEth()
        internal
    {
        uint256 ethBalance = address(this).balance;
        if (ethBalance != 0) {
            getWethContract().deposit.value(ethBalance)();
        }
    }

    /// @dev Returns the WETH balance of this contract, minus
    ///      any WETH that has already been reserved for rewards.
    function _getAvailableWethBalance()
        internal
        view
        returns (uint256 wethBalance)
    {
        wethBalance = getWethContract().balanceOf(address(this))
            .safeSub(wethReservedForPoolRewards);

        return wethBalance;
    }

    /// @dev Asserts that a pool has been finalized last epoch.
    /// @param poolId The id of the pool that should have been finalized.
    function _assertPoolFinalizedLastEpoch(bytes32 poolId)
        internal
        view
    {
        uint256 prevEpoch = currentEpoch.safeSub(1);
        IStructs.PoolStats memory poolStats = poolStatsByEpoch[poolId][prevEpoch];

        // A pool that has any fees remaining has not been finalized
        if (poolStats.feesCollected != 0) {
            revert();
        }
    }

    /// @dev Computes the reward owed to a pool during finalization.
    /// @param poolStats Stats for a specific pool.
    /// @param aggregatedStats Stats aggregated across all pools.
    /// @return rewards Unfinalized rewards for the input pool.
    function _getUnfinalizedPoolRewardsFromPoolStats(
        IStructs.PoolStats memory poolStats,
        IStructs.AggregatedStats memory aggregatedStats
    )
        private
        view
        returns (uint256 rewards)
    {
        // There can't be any rewards if the pool did not collect any fees.
        if (poolStats.feesCollected == 0) {
            return rewards;
        }

        // Use the cobb-douglas function to compute the total reward.
        rewards = LibCobbDouglas.cobbDouglas(
            aggregatedStats.rewardsAvailable,
            poolStats.feesCollected,
            aggregatedStats.totalFeesCollected,
            poolStats.weightedStake,
            aggregatedStats.totalWeightedStake,
            cobbDouglasAlphaNumerator,
            cobbDouglasAlphaDenominator
        );

        // Clip the reward to always be under
        // `rewardsAvailable - totalRewardsPaid`,
        // in case cobb-douglas overflows, which should be unlikely.
        uint256 rewardsRemaining = aggregatedStats.rewardsAvailable.safeSub(aggregatedStats.totalRewardsFinalized);
        if (rewardsRemaining < rewards) {
            rewards = rewardsRemaining;
        }
    }
}
