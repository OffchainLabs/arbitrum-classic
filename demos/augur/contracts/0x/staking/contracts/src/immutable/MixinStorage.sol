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


import "../../../../utils/contracts/src/Authorizable.sol";
import "./MixinConstants.sol";
import "../interfaces/IZrxVault.sol";
import "../interfaces/IStructs.sol";
import "../libs/LibStakingRichErrors.sol";


// solhint-disable max-states-count, no-empty-blocks
contract MixinStorage is
    Authorizable
{
    // address of staking contract
    address public stakingContract;

    // mapping from StakeStatus to global stored balance
    // NOTE: only Status.DELEGATED is used to access this mapping, but this format
    // is used for extensibility
    mapping (uint8 => IStructs.StoredBalance) internal _globalStakeByStatus;

    // mapping from StakeStatus to address of staker to stored balance
    mapping (uint8 => mapping (address => IStructs.StoredBalance)) internal _ownerStakeByStatus;

    // Mapping from Owner to Pool Id to Amount Delegated
    mapping (address => mapping (bytes32 => IStructs.StoredBalance)) internal _delegatedStakeToPoolByOwner;

    // Mapping from Pool Id to Amount Delegated
    mapping (bytes32 => IStructs.StoredBalance) internal _delegatedStakeByPoolId;

    // tracking Pool Id, a unique identifier for each staking pool.
    bytes32 public lastPoolId;

    /// @dev Mapping from Maker Address to pool Id of maker
    /// @param 0 Maker address.
    /// @return 0 The pool ID.
    mapping (address => bytes32) public poolIdByMaker;

    // mapping from Pool Id to Pool
    mapping (bytes32 => IStructs.Pool) internal _poolById;

    /// @dev mapping from pool ID to reward balance of members
    /// @param 0 Pool ID.
    /// @return 0 The total reward balance of members in this pool.
    mapping (bytes32 => uint256) public rewardsByPoolId;

    // The current epoch.
    uint256 public currentEpoch;

    // The current epoch start time.
    uint256 public currentEpochStartTimeInSeconds;

    // mapping from Pool Id to Epoch to Reward Ratio
    mapping (bytes32 => mapping (uint256 => IStructs.Fraction)) internal _cumulativeRewardsByPool;

    // mapping from Pool Id to Epoch
    mapping (bytes32 => uint256) internal _cumulativeRewardsByPoolLastStored;

    /// @dev Registered 0x Exchange contracts, capable of paying protocol fees.
    /// @param 0 The address to check.
    /// @return 0 Whether the address is a registered exchange.
    mapping (address => bool) public validExchanges;

    /* Tweakable parameters */

    // Minimum seconds between epochs.
    uint256 public epochDurationInSeconds;

    // How much delegated stake is weighted vs operator stake, in ppm.
    uint32 public rewardDelegatedStakeWeight;

    // Minimum amount of stake required in a pool to collect rewards.
    uint256 public minimumPoolStake;

    // Numerator for cobb douglas alpha factor.
    uint32 public cobbDouglasAlphaNumerator;

    // Denominator for cobb douglas alpha factor.
    uint32 public cobbDouglasAlphaDenominator;

    /* State for finalization */

    /// @dev Stats for each pool that generated fees with sufficient stake to earn rewards.
    ///      See `_minimumPoolStake` in `MixinParams`.
    /// @param 0 Pool ID.
    /// @param 1 Epoch number.
    /// @return 0 Pool fee stats.
    mapping (bytes32 => mapping (uint256 => IStructs.PoolStats)) public poolStatsByEpoch;

    /// @dev Aggregated stats across all pools that generated fees with sufficient stake to earn rewards.
    ///      See `_minimumPoolStake` in MixinParams.
    /// @param 0 Epoch number.
    /// @return 0 Reward computation stats.
    mapping (uint256 => IStructs.AggregatedStats) public aggregatedStatsByEpoch;

    /// @dev The WETH balance of this contract that is reserved for pool reward payouts.
    uint256 public wethReservedForPoolRewards;
}
