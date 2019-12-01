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

import "./Challenge.sol";
import "./OneStepProof.sol";

import "../libraries/MerkleLib.sol";
import "../libraries/ArbProtocol.sol";


library ChallengeImpl {

    event InitiatedChallenge(
        uint64 deadline
    );

    event ContinuedChallenge (
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        bytes32[] machineHashes,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint32 totalSteps,
        uint64 deadline
    );

    // Can online initialize once
    string constant CHAL_INIT_STATE = "CHAL_INIT_STATE";
    // Can only continue challenge in response to bisection
    string constant CON_STATE = "CON_STATE";
    // Incorrect previous state
    string constant CON_PREV = "CON_PREV";
    // deadline expired
    string constant CON_DEADLINE = "CON_DEADLINE";
    // Only original challenger can continue challenge
    string constant CON_SENDER = "CON_SENDER";
    // Invalid assertion selected
    string constant CON_PROOF = "CON_PROOF";
    // Can only bisect assertion in response to a challenge
    string constant BIS_STATE = "BIS_STATE";
    // Incorrect previous state
    string constant BIS_INPLEN = "BIS_INPLEN";
    // Incorrect previous state
    string constant BIS_PREV = "BIS_PREV";
    // deadline expired
    string constant BIS_DEADLINE = "BIS_DEADLINE";
    // Only original asserter can continue bisect
    string constant BIS_SENDER = "BIS_SENDER";

    // Can only one step proof following a single step challenge
    string constant OSP_STATE = "OSP_STATE";
    // One step proof with invalid prev state
    string constant OSP_PREV = "OSP_PREV";
    // deadline expired
    string constant OSP_DEADLINE = "OSP_DEADLINE";
    // Proof was incorrect
    string constant OSP_PROOF = "OSP_PROOF";


    function initializeChallenge(
        Challenge.Data storage challenge,
        address vmAddress,
        address[2] memory players,
        uint128[2] memory escrows,
        uint32 challengePeriod,
        bytes32 beforeHash,
        bytes32 beforeInbox,
        uint64[2] memory timeBounds,
        bytes32 assertionHash
    )
        public
    {
        require(challenge.state == Challenge.State.NoChallenge, CHAL_INIT_STATE);

        uint64 deadline = uint64(block.number) + uint64(challengePeriod);
        challenge.vmAddress = vmAddress;
        challenge.challengeState = keccak256(
            abi.encodePacked(
                keccak256(
                    abi.encodePacked(
                        timeBounds[0],
                        timeBounds[1],
                        beforeInbox
                    )
                ),
                beforeHash,
                assertionHash
            )
        );
        challenge.escrows = escrows;
        challenge.players = players;
        challenge.deadline = deadline;
        challenge.challengePeriod = challengePeriod;
        challenge.state = Challenge.State.Challenged;

        emit InitiatedChallenge(
            deadline
        );
    }

    function continueChallenge(
        Challenge.Data storage challenge,
        uint assertionToChallenge,
        bytes memory proof,
        bytes32 bisectionRoot,
        bytes32 bisectionHash
    )
        public
    {
        require(challenge.state == Challenge.State.Bisected, CON_STATE);
        require(bisectionRoot == challenge.challengeState, CON_PREV);
        require(block.number <= challenge.deadline, CON_DEADLINE);
        require(msg.sender == challenge.players[1], CON_SENDER);
        require(
            MerkleLib.verifyProof(
                proof,
                bisectionRoot,
                bisectionHash,
                assertionToChallenge + 1
            ),
            CON_PROOF
        );

        challenge.state = Challenge.State.Challenged;
        uint64 deadline = uint64(block.number) + uint64(challenge.challengePeriod);
        challenge.deadline = deadline;
        challenge.challengeState = bisectionHash;
        emit ContinuedChallenge(assertionToChallenge, deadline);
    }

    function bisectAssertion(
        Challenge.Data storage challenge,
        bytes32 preData,
        bytes32[] memory machineHashes,
        bytes32[] memory messageAccs,
        bytes32[] memory logAccs,
        uint32 totalSteps
    )
        public
    {
        uint bisectionCount = machineHashes.length - 1;
        require(bisectionCount + 1 == messageAccs.length, BIS_INPLEN);
        require(bisectionCount + 1 == logAccs.length, BIS_INPLEN);
        require(Challenge.State.Challenged == challenge.state, BIS_STATE);
        require(block.number <= challenge.deadline, BIS_DEADLINE);
        require(msg.sender == challenge.players[0], BIS_SENDER);



        require(
            keccak256(
                abi.encodePacked(
                    preData,
                    machineHashes[0],
                    ArbProtocol.generateAssertionHash(
                        machineHashes[bisectionCount],
                        totalSteps,
                        messageAccs[0],
                        messageAccs[bisectionCount],
                        logAccs[0],
                        logAccs[bisectionCount]
                    )
                )
            ) == challenge.challengeState,
            BIS_PREV
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = keccak256(
            abi.encodePacked(
                preData,
                machineHashes[0],
                ArbProtocol.generateAssertionHash(
                    machineHashes[1],
                    totalSteps / uint32(bisectionCount) + totalSteps%uint32(bisectionCount),
                    messageAccs[0],
                    messageAccs[1],
                    logAccs[0],
                    logAccs[1]
                )
            )
        );
        for (uint i = 1; i < bisectionCount; i++) {
            hashes[i] = keccak256(
                abi.encodePacked(
                    preData,
                    machineHashes[i],
                    ArbProtocol.generateAssertionHash(
                        machineHashes[i + 1],
                        totalSteps / uint32(bisectionCount),
                        messageAccs[i],
                        messageAccs[i + 1],
                        logAccs[i],
                        logAccs[i + 1]
                    )
                )
            );
        }
        challenge.challengeState = MerkleLib.generateRoot(hashes);
        challenge.state = Challenge.State.Bisected;
        uint64 deadline = uint64(block.number) + uint64(challenge.challengePeriod);
        challenge.deadline = deadline;

        emit BisectedAssertion(
            machineHashes,
            messageAccs,
            logAccs,
            totalSteps,
            deadline
        );
    }

    function oneStepProof(
        Challenge.Data storage challenge,
        bytes32 beforeHash,
        bytes32 beforeInbox,
        uint64[2] memory timeBounds,
        bytes32 afterHash,
        bytes32 firstMessage,
        bytes32 lastMessage,
        bytes32 firstLog,
        bytes32 lastLog,
        bytes memory proof
    )
        public view
    {
        require(challenge.state == Challenge.State.Challenged, OSP_STATE);
        require(block.number <= challenge.deadline, OSP_DEADLINE);
        require(
            keccak256(
                abi.encodePacked(
                    keccak256(
                        abi.encodePacked(
                            timeBounds[0],
                            timeBounds[1],
                            beforeInbox
                        )
                    ),
                    beforeHash,
                    ArbProtocol.generateAssertionHash(
                        afterHash,
                        1,
                        firstMessage,
                        lastMessage,
                        firstLog,
                        lastLog
                    )
                )
            ) == challenge.challengeState,
            OSP_PREV
        );

        uint correctProof = OneStepProof.validateProof(
            beforeHash,
            timeBounds,
            beforeInbox,
            afterHash,
            firstMessage,
            lastMessage,
            firstLog,
            lastLog,
            proof
        );

        require(correctProof == 0, OSP_PROOF);
    }
}
