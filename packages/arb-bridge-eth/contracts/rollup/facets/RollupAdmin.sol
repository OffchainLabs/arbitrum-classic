// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../Rollup.sol";

contract RollupAdminFacet is RollupBase {
    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    /**
     * @notice Add a contract authorized to put messages into this rollup's inbox
     * @param _outbox Outbox contract to add
     */
    function setOutbox(IOutbox _outbox) external onlyOwner {
        outbox = _outbox;
        delayedBridge.setOutbox(address(_outbox), true);
        emit OwnerFunctionCalled(0);
    }

    /**
     * @notice Disable an old outbox from interacting with the bridge
     * @param _outbox Outbox contract to remove
     */
    function removeOldOutbox(address _outbox) external onlyOwner {
        require(_outbox != address(outbox), "CUR_OUTBOX");
        delayedBridge.setOutbox(_outbox, false);
        emit OwnerFunctionCalled(1);
    }

    /**
     * @notice Enable or disable an inbox contract
     * @param _inbox Inbox contract to add or remove
     * @param _enabled New status of inbox
     */
    function setInbox(address _inbox, bool _enabled) external onlyOwner {
        delayedBridge.setInbox(address(_inbox), _enabled);
        emit OwnerFunctionCalled(2);
    }

    /**
     * @notice Pause interaction with the rollup contract
     */
    function pause() external onlyOwner {
        _pause();
        emit OwnerFunctionCalled(3);
    }

    /**
     * @notice Resume interaction with the rollup contract
     */
    function resume() external onlyOwner {
        _unpause();
        emit OwnerFunctionCalled(4);
    }

    /*
    function forceResolveChallenge(address[] memory stackerA, address[] memory stackerB) external onlyOwner whenPaused {
        require(stackerA.length == stackerB.length, "WRONG_LENGTH");
        for (uint256 i = 0; i < stackerA.length; i++) {
            address chall = inChallenge(stackerA[i], stackerB[i]);

            require(address(0) != chall, "NOT_IN_CHALL");
            clearChallenge(stackerA[i]);
            clearChallenge(stackerB[i]);

            IChallenge(chall).clearChallenge();
        }
    }

    function forceRefundStaker(address[] memory stacker) external onlyOwner whenPaused {
        for (uint256 i = 0; i < stacker.length; i++) {
            withdrawStaker(stacker[i]);
        }
    }

    function forceCreateNode(
        bytes32 expectedNodeHash,
        bytes32[3][2] calldata assertionBytes32Fields,
        uint256[4][2] calldata assertionIntFields,
        uint256 beforeProposedBlock,
        uint256 beforeInboxMaxCount,
        uint256 prevNode,
        uint256 deadlineBlock,
        uint256 sequencerBatchEnd,
        bytes32 sequencerBatchAcc
    ) external onlyOwner whenPaused {
        require(prevNode == latestConfirmed(), "ONLY_LATEST_CONFIRMED");

        RollupLib.Assertion memory assertion =
                RollupLib.decodeAssertion(
                    assertionBytes32Fields,
                    assertionIntFields,
                    beforeProposedBlock,
                    beforeInboxMaxCount,
                    sequencerBridge.messageCount()
                );

        bytes32 nodeHash =
            _newNode(
                assertion,
                deadlineBlock,
                sequencerBatchEnd,
                sequencerBatchAcc,
                prevNode,
                getNodeHash(prevNode),
                false
            );
        // TODO: should we add a stake?
        
        require(expectedNodeHash == nodeHash, "NOT_EXPECTED_HASH");
    }

    function forceConfirmNode(
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external onlyOwner whenPaused {
        outbox.processOutgoingMessages(sendsData, sendLengths);

        confirmLatestNode();

        rollupEventBridge.nodeConfirmed(latestConfirmed());

        // emit NodeConfirmed(
        //     firstUnresolved,
        //     afterSendAcc,
        //     afterSendCount,
        //     afterLogAcc,
        //     afterLogCount
        // );
    }
    */
}
