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
}
