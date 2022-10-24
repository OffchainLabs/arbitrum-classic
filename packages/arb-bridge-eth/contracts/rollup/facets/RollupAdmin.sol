// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../Rollup.sol";
import "../INode.sol";
import "./IRollupFacets.sol";
import { NitroReadyMagicNums } from "../../bridge/NitroMigratorUtil.sol";
import "../../bridge/interfaces/IOutbox.sol";
import "../../bridge/interfaces/ISequencerInbox.sol";
import "../../libraries/Whitelist.sol";

import "@openzeppelin/contracts/proxy/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @notice DEPRECATED - only for classic version, see new repo (https://github.com/OffchainLabs/nitro/tree/master/contracts)
 * for new updates
 */
contract RollupAdminFacet is RollupBase, IRollupAdmin {
    event NodeDestroyedInMigration(uint256 nodeNum);
    event ChallengeDestroyedInMigration(address challenge);
    event StakerWithdrawnInMigration(address staker);

    /**
     * Functions are only to reach this facet if the caller is the owner
     * so there is no need for a redundant onlyOwner check
     */

    /**
     * @notice Add a contract authorized to put messages into this rollup's inbox
     * @param _outbox Outbox contract to add
     */
    function setOutbox(IOutbox _outbox) external override {
        outbox = _outbox;
        delayedBridge.setOutbox(address(_outbox), true);
        emit OwnerFunctionCalled(0);
    }

    /**
     * @notice Disable an old outbox from interacting with the bridge
     * @param _outbox Outbox contract to remove
     */
    function removeOldOutbox(address _outbox) external override {
        require(_outbox != address(outbox), "CUR_OUTBOX");
        delayedBridge.setOutbox(_outbox, false);
        emit OwnerFunctionCalled(1);
    }

    /**
     * @notice Enable or disable an inbox contract
     * @param _inbox Inbox contract to add or remove
     * @param _enabled New status of inbox
     */
    function setInbox(address _inbox, bool _enabled) external override {
        delayedBridge.setInbox(address(_inbox), _enabled);
        emit OwnerFunctionCalled(2);
    }

    /**
     * @notice Pause interaction with the rollup contract.
     * The time spent paused is not incremented in the rollup's timing for node validation.
     */
    function pause() external override {
        _pause();
        emit OwnerFunctionCalled(3);
    }

    /**
     * @notice Resume interaction with the rollup contract
     */
    function resume() external override {
        _unpause();
        emit OwnerFunctionCalled(4);
    }

    /**
     * @notice Set the addresses of rollup logic facets called
     * @param newAdminFacet address of logic that owner of rollup calls
     * @param newUserFacet address of logic that user of rollup calls
     */
    function setFacets(address newAdminFacet, address newUserFacet) external override {
        facets[0] = newAdminFacet;
        facets[1] = newUserFacet;
        emit OwnerFunctionCalled(5);
    }

    /**
     * @notice Set the addresses of the validator whitelist
     * @dev It is expected that both arrays are same length, and validator at
     * position i corresponds to the value at position i
     * @param _validator addresses to set in the whitelist
     * @param _val value to set in the whitelist for corresponding address
     */
    function setValidator(address[] memory _validator, bool[] memory _val) external override {
        require(_validator.length == _val.length, "WRONG_LENGTH");

        for (uint256 i = 0; i < _validator.length; i++) {
            isValidator[_validator[i]] = _val[i];
        }
        emit OwnerFunctionCalled(6);
    }

    /**
     * @notice Set a new owner address for the rollup
     * @param newOwner address of new rollup owner
     */
    function setOwner(address newOwner) external override {
        owner = newOwner;
        emit OwnerFunctionCalled(7);
    }

    /**
     * @notice Set minimum assertion period for the rollup
     * @param newPeriod new minimum period for assertions
     */
    function setMinimumAssertionPeriod(uint256 newPeriod) external override {
        minimumAssertionPeriod = newPeriod;
        emit OwnerFunctionCalled(8);
    }

    /**
     * @notice Set number of blocks until a node is considered confirmed
     * @param newConfirmPeriod new number of blocks
     */
    function setConfirmPeriodBlocks(uint256 newConfirmPeriod) external override {
        confirmPeriodBlocks = newConfirmPeriod;
        emit OwnerFunctionCalled(9);
    }

    /**
     * @notice Set number of extra blocks after a challenge
     * @param newExtraTimeBlocks new number of blocks
     */
    function setExtraChallengeTimeBlocks(uint256 newExtraTimeBlocks) external override {
        extraChallengeTimeBlocks = newExtraTimeBlocks;
        emit OwnerFunctionCalled(10);
    }

    /**
     * @notice Set speed limit per block
     * @param newAvmGasSpeedLimitPerBlock maximum avmgas to be used per block
     */
    function setAvmGasSpeedLimitPerBlock(uint256 newAvmGasSpeedLimitPerBlock) external override {
        avmGasSpeedLimitPerBlock = newAvmGasSpeedLimitPerBlock;
        emit OwnerFunctionCalled(11);
    }

    /**
     * @notice Set base stake required for an assertion
     * @param newBaseStake minimum amount of stake required
     */
    function setBaseStake(uint256 newBaseStake) external override {
        baseStake = newBaseStake;
        emit OwnerFunctionCalled(12);
    }

    /**
     * @notice Set the token used for stake, where address(0) == eth
     * @dev Before changing the base stake token, you might need to change the
     * implementation of the Rollup User facet!
     * @param newStakeToken address of token used for staking
     */
    function setStakeToken(address newStakeToken) external override {
        stakeToken = newStakeToken;
        emit OwnerFunctionCalled(13);
    }

    /**
     * @notice Set max delay for sequencer inbox
     * @param newSequencerInboxMaxDelayBlocks max number of blocks
     * @param newSequencerInboxMaxDelaySeconds max number of seconds
     */
    function setSequencerInboxMaxDelay(
        uint256 newSequencerInboxMaxDelayBlocks,
        uint256 newSequencerInboxMaxDelaySeconds
    ) external override {
        ISequencerInbox(sequencerBridge).setMaxDelay(
            newSequencerInboxMaxDelayBlocks,
            newSequencerInboxMaxDelaySeconds
        );
        emit OwnerFunctionCalled(14);
    }

    /**
     * @notice Set execution bisection degree
     * @param newChallengeExecutionBisectionDegree execution bisection degree
     */
    function setChallengeExecutionBisectionDegree(uint256 newChallengeExecutionBisectionDegree)
        external
        override
    {
        challengeExecutionBisectionDegree = newChallengeExecutionBisectionDegree;
        emit OwnerFunctionCalled(16);
    }

    /**
     * @notice Updates a whitelist address for its consumers
     * @dev setting the newWhitelist to address(0) disables it for consumers
     * @param whitelist old whitelist to be deprecated
     * @param newWhitelist new whitelist to be used
     * @param targets whitelist consumers to be triggered
     */
    function updateWhitelistConsumers(
        address whitelist,
        address newWhitelist,
        address[] memory targets
    ) external override {
        Whitelist(whitelist).triggerConsumers(newWhitelist, targets);
        emit OwnerFunctionCalled(17);
    }

    /**
     * @notice Updates a whitelist's entries
     * @dev user at position i will be assigned value i
     * @param whitelist whitelist to be updated
     * @param user users to be updated in the whitelist
     * @param val if user is or not allowed in the whitelist
     */
    function setWhitelistEntries(
        address whitelist,
        address[] memory user,
        bool[] memory val
    ) external override {
        require(user.length == val.length, "INVALID_INPUT");
        Whitelist(whitelist).setWhitelist(user, val);
        emit OwnerFunctionCalled(18);
    }

    /**
     * @notice Updates a sequencer address at the sequencer inbox
     * @param newSequencer new sequencer address to be used
     */
    function setIsSequencer(address newSequencer, bool isSequencer) external override {
        ISequencerInbox(sequencerBridge).setIsSequencer(newSequencer, isSequencer);
        emit OwnerFunctionCalled(19);
    }

    /**
     * @notice Upgrades the implementation of a beacon controlled by the rollup
     * @param beacon address of beacon to be upgraded
     * @param newImplementation new address of implementation
     */
    function upgradeBeacon(address beacon, address newImplementation) external override {
        UpgradeableBeacon(beacon).upgradeTo(newImplementation);
        emit OwnerFunctionCalled(20);
    }

    function forceResolveChallenge(address[] memory stakerA, address[] memory stakerB)
        external
        override
        whenPaused
    {
        require(stakerA.length == stakerB.length, "WRONG_LENGTH");
        for (uint256 i = 0; i < stakerA.length; i++) {
            address chall = inChallenge(stakerA[i], stakerB[i]);

            require(address(0) != chall, "NOT_IN_CHALL");
            clearChallenge(stakerA[i]);
            clearChallenge(stakerB[i]);

            IChallenge(chall).clearChallenge();
        }
        emit OwnerFunctionCalled(21);
    }

    function forceRefundStaker(address[] memory staker) external override whenPaused {
        for (uint256 i = 0; i < staker.length; i++) {
            reduceStakeTo(staker[i], 0);
            turnIntoZombie(staker[i]);
        }
        emit OwnerFunctionCalled(22);
    }

    function forceCreateNode(
        bytes32 expectedNodeHash,
        bytes32[3][2] calldata assertionBytes32Fields,
        uint256[4][2] calldata assertionIntFields,
        bytes calldata sequencerBatchProof,
        uint256 beforeProposedBlock,
        uint256 beforeInboxMaxCount,
        uint256 prevNode
    ) external override whenPaused {
        require(prevNode == latestConfirmed(), "ONLY_LATEST_CONFIRMED");

        RollupLib.Assertion memory assertion = RollupLib.decodeAssertion(
            assertionBytes32Fields,
            assertionIntFields,
            beforeProposedBlock,
            beforeInboxMaxCount,
            sequencerBridge.messageCount()
        );

        createNewNode(
            assertion,
            assertionBytes32Fields,
            assertionIntFields,
            sequencerBatchProof,
            CreateNodeDataFrame({
                avmGasSpeedLimitPerBlock: avmGasSpeedLimitPerBlock,
                confirmPeriodBlocks: confirmPeriodBlocks,
                prevNode: prevNode,
                sequencerInbox: sequencerBridge,
                rollupEventBridge: rollupEventBridge,
                nodeFactory: nodeFactory
            }),
            expectedNodeHash
        );

        emit OwnerFunctionCalled(23);
    }

    function forceConfirmNode(
        uint256 nodeNum,
        bytes32 beforeSendAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount
    ) external override whenPaused {
        // this skips deadline, staker and zombie validation
        confirmNode(
            nodeNum,
            beforeSendAcc,
            sendsData,
            sendLengths,
            afterSendCount,
            afterLogAcc,
            afterLogCount,
            outbox,
            rollupEventBridge
        );
        emit OwnerFunctionCalled(24);
    }

    /// @dev this function is intended to be called as part of the shutdown process of the classic contracts in favour of nitro
    /// It is expected that the rollup is not paused during the start of shutdown step, the shutdown procedure will pause the rollup.
    /// A final rollup node number is specified, then the rollup will only allow that node and its direct predecessors to be confirmed.
    /// All nodes that aren't directly previous to this are deleted, but in practice none are expected to be present (as this would mean an eventual challenge).
    /// Even though the rollup is paused, we use the `shutdownForNitroMode` var to allow validators to go through the sequence of final nodes confirming them so their send values are added to the outbox
    /// The deadline for the nodes marked as final are ignored to a lower value to allow for these faster confirmations, which will make L2 to L1 txs available for execution sooner
    function shutdownForNitro(
        uint256 finalNodeNum,
        bool destroyAlternatives,
        bool destroyChallenges
    ) external whenNotPaused {
        require(!shutdownForNitroMode(), "ALREADY_SHUTDOWN_MODE");
        // TODO: prove that final node num includes the last send by arbos

        // first we destroy all nodes that aren't in the correct chain
        uint256 latestConfirmedNodeNum = latestConfirmed();
        uint256 curr = latestNodeCreated();
        uint256 expectedPrev = finalNodeNum;
        // if finalNodeNum == latestConfirmed we don't need to delete any siblings
        while (curr != latestConfirmedNodeNum) {
            if (curr == expectedPrev) {
                INode currNode = getNode(curr);
                expectedPrev = currNode.prev();
            } else {
                require(destroyAlternatives, "ALTERNATIVES_NOT_EXPECTED");
                destroyNode(curr);
                emit NodeDestroyedInMigration(curr);
            }
            curr--;
        }

        uint256 stakerCount = stakerCount();
        address[] memory stakerAddresses = new address[](stakerCount);

        // we separate the loop that gets staker addresses to be different from the loop that withdraw stakers
        // since withdrawing stakers has side-effects on the array that is queried in `getStakerAddress`.
        for (uint64 i = 0; i < stakerCount; ++i) {
            stakerAddresses[i] = getStakerAddress(i);
        }

        for (uint64 i = 0; i < stakerCount; ++i) {
            address stakerAddr = stakerAddresses[i];
            address chall = currentChallenge(stakerAddr);

            if (chall != address(0)) {
                require(destroyChallenges, "CHALLENGE_NOT_EXPECTED");
                address asserter = IChallenge(chall).asserter();
                address challenger = IChallenge(chall).challenger();

                clearChallenge(asserter);
                clearChallenge(challenger);

                IChallenge(chall).clearChallenge();
                emit ChallengeDestroyedInMigration(chall);
            }

            if (getNode(latestStakedNode(stakerAddr)) == INode(0)) {
                // this node got destroyed, so we force refund the staker
                withdrawStaker(stakerAddr);
                emit StakerWithdrawnInMigration(stakerAddr);
            }
            // else the staker can unstake and withdraw regularly using `returnOldDeposit`
        }

        shutdownForNitroBlock = block.number;
        _pause();
        emit OwnerFunctionCalled(25);
    }

    /// @dev stops the rollup from shutdownForNitro mode in case something goes wrong during the migration process
    function undoShutdownForNitro() external whenPaused {
        require(shutdownForNitroMode(), "NOT_SHUTDOWN_MODE");
        shutdownForNitroBlock = type(uint256).max;
        _unpause();
        emit OwnerFunctionCalled(26);
    }

    /// @dev allows the admin to transfer the ownership of a contract controlled by the rollup
    function transferOwnership(Ownable target, address newOwner) external {
        target.transferOwnership(newOwner);
        emit OwnerFunctionCalled(27);
    }

    function isNitroReady() external pure returns (uint256) {
        return NitroReadyMagicNums.ROLLUP_ADMIN;
    }
}
