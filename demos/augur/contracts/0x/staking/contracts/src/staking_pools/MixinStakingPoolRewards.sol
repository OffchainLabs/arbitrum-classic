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

import "../../../../exchange-libs/contracts/src/LibMath.sol";
import "../../../../utils/contracts/src/LibSafeMath.sol";
import "./MixinCumulativeRewards.sol";
import "../sys/MixinAbstract.sol";


contract MixinStakingPoolRewards is
    MixinAbstract,
    MixinCumulativeRewards
{
    using LibSafeMath for uint256;

    /// @dev Withdraws the caller's WETH rewards that have accumulated
    ///      until the last epoch.
    /// @param poolId Unique id of pool.
    function withdrawDelegatorRewards(bytes32 poolId)
        external
    {
        _withdrawAndSyncDelegatorRewards(poolId, msg.sender);
    }

    /// @dev Computes the reward balance in ETH of the operator of a pool.
    /// @param poolId Unique id of pool.
    /// @return totalReward Balance in ETH.
    function computeRewardBalanceOfOperator(bytes32 poolId)
        external
        view
        returns (uint256 reward)
    {
        // Because operator rewards are immediately withdrawn as WETH
        // on finalization, the only factor in this function are unfinalized
        // rewards.
        IStructs.Pool memory pool = _poolById[poolId];
        // Get any unfinalized rewards.
        (uint256 unfinalizedTotalRewards, uint256 unfinalizedMembersStake) =
            _getUnfinalizedPoolRewards(poolId);

        // Get the operators' portion.
        (reward,) = _computePoolRewardsSplit(
            pool.operatorShare,
            unfinalizedTotalRewards,
            unfinalizedMembersStake
        );
        return reward;
    }

    /// @dev Computes the reward balance in ETH of a specific member of a pool.
    /// @param poolId Unique id of pool.
    /// @param member The member of the pool.
    /// @return totalReward Balance in ETH.
    function computeRewardBalanceOfDelegator(bytes32 poolId, address member)
        external
        view
        returns (uint256 reward)
    {
        IStructs.Pool memory pool = _poolById[poolId];
        // Get any unfinalized rewards.
        (uint256 unfinalizedTotalRewards, uint256 unfinalizedMembersStake) =
            _getUnfinalizedPoolRewards(poolId);

        // Get the members' portion.
        (, uint256 unfinalizedMembersReward) = _computePoolRewardsSplit(
            pool.operatorShare,
            unfinalizedTotalRewards,
            unfinalizedMembersStake
        );
        return _computeDelegatorReward(
            poolId,
            member,
            unfinalizedMembersReward,
            unfinalizedMembersStake
        );
    }

    /// @dev Syncs rewards for a delegator. This includes withdrawing rewards
    ///      rewards and adding/removing dependencies on cumulative rewards.
    /// @param poolId Unique id of pool.
    /// @param member of the pool.
    function _withdrawAndSyncDelegatorRewards(
        bytes32 poolId,
        address member
    )
        internal
    {
        // Ensure the pool is finalized.
        _assertPoolFinalizedLastEpoch(poolId);

        // Compute balance owed to delegator
        uint256 balance = _computeDelegatorReward(
            poolId,
            member,
            // No unfinalized values because we ensured the pool is already
            // finalized.
            0,
            0
        );

        // Sync the delegated stake balance. This will ensure future calls of
        // `_computeDelegatorReward` during this epoch will return 0, 
        // preventing a delegator from withdrawing more than once an epoch.
        _delegatedStakeToPoolByOwner[member][poolId] =
            _loadCurrentBalance(_delegatedStakeToPoolByOwner[member][poolId]);

        // Withdraw non-0 balance
        if (balance != 0) {
            // Decrease the balance of the pool
            _decreasePoolRewards(poolId, balance);

            // Withdraw the member's WETH balance
            getWethContract().transfer(member, balance);
        }

        // Ensure a cumulative reward entry exists for this epoch,
        // copying the previous epoch's CR if one doesn't exist already.
        _updateCumulativeReward(poolId);
    }

    /// @dev Handles a pool's reward at the current epoch.
    ///      This will split the reward between the operator and members,
    ///      depositing them into their respective vaults, and update the
    ///      accounting needed to allow members to withdraw their individual
    ///      rewards.
    /// @param poolId Unique Id of pool.
    /// @param reward received by the pool.
    /// @param membersStake the amount of non-operator delegated stake that
    ///        will split the  reward.
    /// @return operatorReward Portion of `reward` given to the pool operator.
    /// @return membersReward Portion of `reward` given to the pool members.
    function _syncPoolRewards(
        bytes32 poolId,
        uint256 reward,
        uint256 membersStake
    )
        internal
        returns (uint256 operatorReward, uint256 membersReward)
    {
        IStructs.Pool memory pool = _poolById[poolId];

        // Split the reward between operator and members
        (operatorReward, membersReward) = _computePoolRewardsSplit(
            pool.operatorShare,
            reward,
            membersStake
        );

        if (operatorReward > 0) {
            // Transfer the operator's weth reward to the operator
            getWethContract().transfer(pool.operator, operatorReward);
        }

        if (membersReward > 0) {
            // Increase the balance of the pool
            _increasePoolRewards(poolId, membersReward);
            // Create a cumulative reward entry at the current epoch.
            _addCumulativeReward(poolId, membersReward, membersStake);
        }

        return (operatorReward, membersReward);
    }

    /// @dev Compute the split of a pool reward between the operator and members
    ///      based on the `operatorShare` and `membersStake`.
    /// @param operatorShare The fraction of rewards owed to the operator,
    ///        in PPM.
    /// @param totalReward The pool reward.
    /// @param membersStake The amount of member (non-operator) stake delegated
    ///        to the pool in the epoch the rewards were earned.
    /// @return operatorReward Portion of `totalReward` given to the pool operator.
    /// @return membersReward Portion of `totalReward` given to the pool members.
    function _computePoolRewardsSplit(
        uint32 operatorShare,
        uint256 totalReward,
        uint256 membersStake
    )
        internal
        pure
        returns (uint256 operatorReward, uint256 membersReward)
    {
        if (membersStake == 0) {
            operatorReward = totalReward;
        } else {
            operatorReward = LibMath.getPartialAmountCeil(
                uint256(operatorShare),
                PPM_DENOMINATOR,
                totalReward
            );
            membersReward = totalReward.safeSub(operatorReward);
        }
        return (operatorReward, membersReward);
    }

    /// @dev Computes the reward balance in ETH of a specific member of a pool.
    /// @param poolId Unique id of pool.
    /// @param member of the pool.
    /// @param unfinalizedMembersReward Unfinalized total members reward (if any).
    /// @param unfinalizedMembersStake Unfinalized total members stake (if any).
    /// @return reward Balance in WETH.
    function _computeDelegatorReward(
        bytes32 poolId,
        address member,
        uint256 unfinalizedMembersReward,
        uint256 unfinalizedMembersStake
    )
        private
        view
        returns (uint256 reward)
    {
        uint256 currentEpoch_ = currentEpoch;
        IStructs.StoredBalance memory delegatedStake = _delegatedStakeToPoolByOwner[member][poolId];

        // There can be no rewards if the last epoch when stake was stored is
        // equal to the current epoch, because all prior rewards, including
        // rewards finalized this epoch have been claimed.
        if (delegatedStake.currentEpoch == currentEpoch_) {
            return 0;
        }

        // We account for rewards over 3 intervals, below.

        // 1/3 Unfinalized rewards earned in `currentEpoch - 1`.
        reward = _computeUnfinalizedDelegatorReward(
            delegatedStake,
            currentEpoch_,
            unfinalizedMembersReward,
            unfinalizedMembersStake
        );

        // 2/3 Finalized rewards earned in epochs [`delegatedStake.currentEpoch + 1` .. `currentEpoch - 1`]
        uint256 delegatedStakeNextEpoch = uint256(delegatedStake.currentEpoch).safeAdd(1);
        reward = reward.safeAdd(
            _computeMemberRewardOverInterval(
                poolId,
                delegatedStake.currentEpochBalance,
                delegatedStake.currentEpoch,
                delegatedStakeNextEpoch
            )
        );

        // 3/3 Finalized rewards earned in epoch `delegatedStake.currentEpoch`.
        reward = reward.safeAdd(
            _computeMemberRewardOverInterval(
                poolId,
                delegatedStake.nextEpochBalance,
                delegatedStakeNextEpoch,
                currentEpoch_
            )
        );

        return reward;
    }

    /// @dev Computes the unfinalized rewards earned by a delegator in the last epoch.
    /// @param delegatedStake Amount of stake delegated to pool by a specific staker
    /// @param currentEpoch_ The epoch in which this call is executing
    /// @param unfinalizedMembersReward Unfinalized total members reward (if any).
    /// @param unfinalizedMembersStake Unfinalized total members stake (if any).
    /// @return reward Balance in WETH.
    function _computeUnfinalizedDelegatorReward(
        IStructs.StoredBalance memory delegatedStake,
        uint256 currentEpoch_,
        uint256 unfinalizedMembersReward,
        uint256 unfinalizedMembersStake
    )
        private
        pure
        returns (uint256)
    {
        // If there are unfinalized rewards this epoch, compute the member's
        // share.
        if (unfinalizedMembersReward == 0 || unfinalizedMembersStake == 0) {
            return 0;
        }

        // Unfinalized rewards are always earned from stake in
        // the prior epoch so we want the stake at `currentEpoch_-1`.
        uint256 unfinalizedStakeBalance = delegatedStake.currentEpoch >= currentEpoch_.safeSub(1) ?
            delegatedStake.currentEpochBalance :
            delegatedStake.nextEpochBalance;

        // Sanity check to save gas on computation
        if (unfinalizedStakeBalance == 0) {
            return 0;
        }

        // Compute unfinalized reward
        return LibMath.getPartialAmountFloor(
            unfinalizedMembersReward,
            unfinalizedMembersStake,
            unfinalizedStakeBalance
        );
    }

    /// @dev Increases rewards for a pool.
    /// @param poolId Unique id of pool.
    /// @param amount Amount to increment rewards by.
    function _increasePoolRewards(bytes32 poolId, uint256 amount)
        private
    {
        rewardsByPoolId[poolId] = rewardsByPoolId[poolId].safeAdd(amount);
        wethReservedForPoolRewards = wethReservedForPoolRewards.safeAdd(amount);
    }

    /// @dev Decreases rewards for a pool.
    /// @param poolId Unique id of pool.
    /// @param amount Amount to decrement rewards by.
    function _decreasePoolRewards(bytes32 poolId, uint256 amount)
        private
    {
        rewardsByPoolId[poolId] = rewardsByPoolId[poolId].safeSub(amount);
        wethReservedForPoolRewards = wethReservedForPoolRewards.safeSub(amount);
    }
}
