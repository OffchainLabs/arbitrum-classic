// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

import "../bridge/interfaces/IBridge.sol";

interface IOneStepProof {
    function executeStep(
        IBridge bridge,
        uint256 initialMessagesRead,
        bytes32 initialSendAcc,
        bytes32 initialLogAcc,
        bytes calldata proof
    )
        external
        view
        returns (
            uint64 gas,
            uint256 totalMessagesRead,
            bytes32[4] memory fields
        );
}

interface IOneStepProof2 {
    function executeStep(
        uint256 initialMessagesRead,
        bytes32 initialSendAcc,
        bytes32 initialLogAcc,
        bytes calldata proof,
        bytes calldata bproof
    )
        external
        view
        returns (
            uint64 gas,
            uint256 totalMessagesRead,
            bytes32[4] memory fields
        );
}
