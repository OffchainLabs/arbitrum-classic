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

import "@openzeppelin/contracts/math/SafeMath.sol";


library VM {
    using SafeMath for uint256;

    bytes32 private constant MACHINE_HALT_HASH = bytes32(0);
    bytes32 private constant MACHINE_ERROR_HASH = bytes32(uint(1));

    struct Params {  // these are defined just once for each vM
        uint128 stakeRequirement;
        uint32  gracePeriod;
        uint32  maxExecutionSteps;
        bytes32 pendingInboxHash;
    }

    struct ProtocolState {
        bytes32 machineHash;
        bytes32 inboxHash;
    }

    struct FullAssertion {
        bytes messageData;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        address[] messageDestinations;
        bytes32 logsAccHash;
    }

    function protoStateHash(bytes32 machineHash, bytes32 inboxHash) external pure returns(bytes32) {
        return keccak256(abi.encodePacked(
            machineHash,
            inboxHash
        ));
    }

    function isErrored(bytes32 vmStateHash) external pure returns(bool) {
        return vmStateHash == MACHINE_ERROR_HASH;
    }

    function isHalted(bytes32 vmStateHash) external pure returns(bool) {
        return vmStateHash == MACHINE_HALT_HASH;
    }
}
