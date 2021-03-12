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

import "arb-bridge-eth/contracts/rollup/Rollup.sol";
import "arb-bridge-eth/contracts/libraries/BytesLib.sol";

contract ConfirmRoots {
    using BytesLib for bytes;

    bytes1 internal constant MSG_ROOT = 0;

    Rollup public rollup;

    // ConfirmRoot => nodeNum => valid
    mapping(bytes32 => mapping(uint256 => bool)) public confirmRoots;

    constructor(Rollup _rollup) public {
        rollup = _rollup;
    }

    function setupConfirmData(
        uint256 nodeNum,
        bytes32 beforeSendAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount
    ) external {
        INode node = rollup.getNode(nodeNum);
        bytes32 afterSendAcc = RollupLib.feedAccumulator(sendsData, sendLengths, beforeSendAcc);
        require(
            node.confirmData() ==
                RollupLib.confirmHash(
                    beforeSendAcc,
                    afterSendAcc,
                    afterLogAcc,
                    afterSendCount,
                    afterLogCount
                ),
            "CONFIRM_DATA"
        );
        uint256 messageCount = sendLengths.length;
        uint256 offset = 0;
        for (uint256 i = 0; i < messageCount; i++) {
            handleOutgoingMessage(nodeNum, bytes(sendsData[offset:sendLengths[i]]));
            offset += sendLengths[i];
        }
    }

    function handleOutgoingMessage(uint256 nodeNum, bytes memory data) private {
        // Otherwise we have an unsupported message type and we skip the message
        if (data[0] == MSG_ROOT) {
            bytes32 outputRoot = data.toBytes32(65);
            confirmRoots[outputRoot][nodeNum] = true;
        }
    }
}
