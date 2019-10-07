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

import "../libraries/ArbProtocol.sol";


library Challenge {
    enum State {
        NoChallenge,
        Challenged,
        Bisected
    }

    struct Data {
        address vmAddress;
        // After bisection this is an array of all sub-assertions
        // After a challenge, the first assertion is the main assertion
        bytes32 challengeState;

        uint128[2] escrows;

        address[2] players;

        uint64 deadline;

        // The current deadline at which the challenge timeouts and a winner is
        // declared. This deadline resets at each step in the challenge
        uint32 challengePeriod;

        State state;
    }

    struct BisectOtherData {
        bytes32 beforeHash;
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes21[] tokenTypes;
        uint256[] beforeBalances;
        bytes32 firstMessageHash;
        bytes32 firstLogHash;
        bytes32 a1AfterHash;
        uint32 a1NumSteps;
        bytes32 a1LastMessageHash;
        bytes32 a1LastLogHash;
        uint256[] a1OutputValues;
        bytes32 a2AfterHash;
        uint32 a2NumSteps;
        bytes32 a2LastMessageHash;
        bytes32 a2LastLogHash;
        uint256[] a2OutputValues;
    }

    function validateBisectionOther(
        Challenge.Data storage _challenge,
        Challenge.BisectOtherData memory _data
    )
        internal
        view
    {
        bytes32 oldPre = ArbProtocol.generatePreconditionHash(
            _data.beforeHash,
            _data.timeBounds,
            _data.beforeInbox,
            _data.tokenTypes,
            _data.beforeBalances
        );

        require(
            keccak256(
                abi.encodePacked(
                    oldPre,
                    ArbProtocol.generateAssertionHash(
                        _data.a1AfterHash,
                        _data.a1NumSteps,
                        _data.firstMessageHash,
                        _data.a1LastMessageHash,
                        _data.firstLogHash,
                        _data.a1LastLogHash,
                        _data.a1OutputValues
                    ),
                    ArbProtocol.generateAssertionHash(
                        _data.a2AfterHash,
                        _data.a2NumSteps,
                        _data.firstMessageHash,
                        _data.a2LastLogHash,
                        _data.firstLogHash,
                        _data.a2LastLogHash,
                        _data.a2OutputValues
                    ),
                    _data.a2NumSteps - _data.a1NumSteps
                )
            ) == _challenge.challengeState,
            "Bisector incorrectly revealed bisection segments"
        );

        uint tokenCount = _data.tokenTypes.length;

        for (uint i = 0; i < tokenCount; i++) {
            _data.beforeBalances[i] -= _data.a1OutputValues[i];
            _data.a2OutputValues[i] -= _data.a1OutputValues[i];
        }
    }
}
