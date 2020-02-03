pragma solidity 0.5.15;


interface IStakingEvents {

    /// @dev Emitted by MixinStake when ZRX is staked.
    /// @param staker of ZRX.
    /// @param amount of ZRX staked.
    event Stake(
        address indexed staker,
        uint256 amount
    );

    /// @dev Emitted by MixinStake when ZRX is unstaked.
    /// @param staker of ZRX.
    /// @param amount of ZRX unstaked.
    event Unstake(
        address indexed staker,
        uint256 amount
    );

    /// @dev Emitted by MixinStake when ZRX is unstaked.
    /// @param staker of ZRX.
    /// @param amount of ZRX unstaked.
    event MoveStake(
        address indexed staker,
        uint256 amount,
        uint8 fromStatus,
        bytes32 indexed fromPool,
        uint8 toStatus,
        bytes32 indexed toPool
    );

    /// @dev Emitted by MixinExchangeManager when an exchange is added.
    /// @param exchangeAddress Address of new exchange.
    event ExchangeAdded(
        address exchangeAddress
    );

    /// @dev Emitted by MixinExchangeManager when an exchange is removed.
    /// @param exchangeAddress Address of removed exchange.
    event ExchangeRemoved(
        address exchangeAddress
    );

    /// @dev Emitted by MixinExchangeFees when a pool starts earning rewards in an epoch.
    /// @param epoch The epoch in which the pool earned rewards.
    /// @param poolId The ID of the pool.
    event StakingPoolEarnedRewardsInEpoch(
        uint256 indexed epoch,
        bytes32 indexed poolId
    );

    /// @dev Emitted by MixinFinalizer when an epoch has ended.
    /// @param epoch The epoch that ended.
    /// @param numPoolsToFinalize Number of pools that earned rewards during `epoch` and must be finalized.
    /// @param rewardsAvailable Rewards available to all pools that earned rewards during `epoch`.
    /// @param totalWeightedStake Total weighted stake across all pools that earned rewards during `epoch`.
    /// @param totalFeesCollected Total fees collected across all pools that earned rewards during `epoch`.
    event EpochEnded(
        uint256 indexed epoch,
        uint256 numPoolsToFinalize,
        uint256 rewardsAvailable,
        uint256 totalFeesCollected,
        uint256 totalWeightedStake
    );

    /// @dev Emitted by MixinFinalizer when an epoch is fully finalized.
    /// @param epoch The epoch being finalized.
    /// @param rewardsPaid Total amount of rewards paid out.
    /// @param rewardsRemaining Rewards left over.
    event EpochFinalized(
        uint256 indexed epoch,
        uint256 rewardsPaid,
        uint256 rewardsRemaining
    );

    /// @dev Emitted by MixinFinalizer when rewards are paid out to a pool.
    /// @param epoch The epoch when the rewards were paid out.
    /// @param poolId The pool's ID.
    /// @param operatorReward Amount of reward paid to pool operator.
    /// @param membersReward Amount of reward paid to pool members.
    event RewardsPaid(
        uint256 indexed epoch,
        bytes32 indexed poolId,
        uint256 operatorReward,
        uint256 membersReward
    );

    /// @dev Emitted whenever staking parameters are changed via the `setParams()` function.
    /// @param epochDurationInSeconds Minimum seconds between epochs.
    /// @param rewardDelegatedStakeWeight How much delegated stake is weighted vs operator stake, in ppm.
    /// @param minimumPoolStake Minimum amount of stake required in a pool to collect rewards.
    /// @param cobbDouglasAlphaNumerator Numerator for cobb douglas alpha factor.
    /// @param cobbDouglasAlphaDenominator Denominator for cobb douglas alpha factor.
    event ParamsSet(
        uint256 epochDurationInSeconds,
        uint32 rewardDelegatedStakeWeight,
        uint256 minimumPoolStake,
        uint256 cobbDouglasAlphaNumerator,
        uint256 cobbDouglasAlphaDenominator
    );

    /// @dev Emitted by MixinStakingPool when a new pool is created.
    /// @param poolId Unique id generated for pool.
    /// @param operator The operator (creator) of pool.
    /// @param operatorShare The share of rewards given to the operator, in ppm.
    event StakingPoolCreated(
        bytes32 poolId,
        address operator,
        uint32 operatorShare
    );

    /// @dev Emitted by MixinStakingPool when a maker sets their pool.
    /// @param makerAddress Adress of maker added to pool.
    /// @param poolId Unique id of pool.
    event MakerStakingPoolSet(
        address indexed makerAddress,
        bytes32 indexed poolId
    );

    /// @dev Emitted when a staking pool's operator share is decreased.
    /// @param poolId Unique Id of pool.
    /// @param oldOperatorShare Previous share of rewards owned by operator.
    /// @param newOperatorShare Newly decreased share of rewards owned by operator.
    event OperatorShareDecreased(
        bytes32 indexed poolId,
        uint32 oldOperatorShare,
        uint32 newOperatorShare
    );
}
