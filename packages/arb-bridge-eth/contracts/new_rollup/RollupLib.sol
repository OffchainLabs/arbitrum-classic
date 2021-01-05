// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

import "../new_challenge/ChallengeLib.sol";

library RollupLib {
    function nodeStateHash(
        uint256 stepsRun,
        bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount,
        uint256 messageCount,
        uint256 logCount,
        uint256 inboxMaxCount
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    stepsRun,
                    machineHash,
                    inboxTop,
                    inboxCount,
                    messageCount,
                    logCount,
                    inboxMaxCount
                )
            );
    }

    struct Assertion {
        uint256 beforeStepsRun;
        bytes32 beforeMachineHash;
        bytes32 beforeInboxHash;
        uint256 beforeInboxCount;
        uint256 beforeSendCount;
        uint256 beforeLogCount;
        uint256 beforeInboxMaxCount;
        uint256 stepsExecuted;
        bytes32 inboxDelta;
        uint256 inboxMessagesRead;
        uint256 gasUsed;
        bytes32 sendAcc;
        uint256 sendCount;
        bytes32 logAcc;
        uint256 logCount;
        bytes32 afterInboxHash;
        bytes32 afterMachineHash;
    }

    function decodeAssertion(bytes32[7] memory bytes32Fields, uint256[10] memory intFields)
        internal
        pure
        returns (Assertion memory)
    {
        return
            Assertion(
                intFields[0],
                bytes32Fields[0],
                bytes32Fields[1],
                intFields[1],
                intFields[2],
                intFields[3],
                intFields[4],
                intFields[5],
                bytes32Fields[2],
                intFields[6],
                intFields[7],
                bytes32Fields[3],
                intFields[8],
                bytes32Fields[4],
                intFields[9],
                bytes32Fields[5],
                bytes32Fields[6]
            );
    }

    function beforeNodeStateHash(Assertion memory assertion) internal pure returns (bytes32) {
        return
            nodeStateHash(
                assertion.beforeStepsRun,
                assertion.beforeMachineHash,
                assertion.beforeInboxHash,
                assertion.beforeInboxCount,
                assertion.beforeSendCount,
                assertion.beforeLogCount,
                assertion.beforeInboxMaxCount
            );
    }

    function afterNodeStateHash(Assertion memory assertion, uint256 inboxMaxCount)
        internal
        pure
        returns (bytes32)
    {
        return
            nodeStateHash(
                assertion.beforeStepsRun + assertion.stepsExecuted,
                assertion.afterMachineHash,
                assertion.afterInboxHash,
                assertion.beforeInboxCount + assertion.inboxMessagesRead,
                assertion.beforeSendCount + assertion.sendCount,
                assertion.beforeLogCount + assertion.logCount,
                inboxMaxCount
            );
    }

    function challengeRoot(
        Assertion memory assertion,
        uint256 inboxTopCount,
        bytes32 inboxTopHash,
        uint256 executionCheckTime
    ) internal pure returns (bytes32) {
        bytes32 executionHash = ChallengeLib.bisectionChunkHash(
            assertion.stepsExecuted,
            ChallengeLib.assertionHash(
                assertion.inboxDelta,
                0,
                ChallengeLib.outputAccHash(0, 0, 0, 0),
                assertion.beforeMachineHash
            ),
            ChallengeLib.assertionHash(
                0,
                assertion.gasUsed,
                ChallengeLib.outputAccHash(
                    assertion.sendAcc,
                    assertion.sendCount,
                    assertion.logAcc,
                    assertion.logCount
                ),
                assertion.afterMachineHash
            )
        );

        bytes32 inboxConsistencyHash = ChallengeLib.bisectionChunkHash(
            inboxTopCount - assertion.beforeInboxCount - assertion.inboxMessagesRead,
            inboxTopHash,
            assertion.afterInboxHash
        );

        bytes32 inboxDeltaHash = ChallengeLib.bisectionChunkHash(
            assertion.inboxMessagesRead,
            ChallengeLib.inboxDeltaHash(assertion.afterInboxHash, 0),
            ChallengeLib.inboxDeltaHash(assertion.beforeInboxHash, assertion.inboxDelta)
        );

        return
            ChallengeLib.challengeRootHash(
                inboxConsistencyHash,
                inboxDeltaHash,
                executionHash,
                executionCheckTime
            );
    }
}
