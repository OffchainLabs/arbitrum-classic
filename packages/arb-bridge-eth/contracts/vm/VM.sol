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

    enum State {
        Uninitialized,
        Waiting,
        PendingDisputable,
        PendingUnanimous
    }

    struct Data {
        bytes32 machineHash;
        bytes32 pendingHash; // Lock pending and confirm asserts together
        bytes32 inbox;
        address asserter;
        uint128 escrowRequired;
        uint64 deadline;
        uint64 sequenceNum;
        uint32 gracePeriod;
        uint32 maxExecutionSteps;
        State state;
        address activeChallengeManager;
    }

    struct FullAssertion {
        bytes messageData;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        address[] messageDestinations;
        bytes32 logsAccHash;
    }

    function acceptAssertion(Data storage _vm, bytes32 _afterHash) external {
        _vm.machineHash = _afterHash;
        _vm.state = VM.State.Waiting;
    }

    function withinDeadline(Data storage _vm) external view returns(bool) {
        return block.number <= _vm.deadline;
    }

    function resetDeadline(Data storage _vm) external {
        _vm.deadline = uint64(block.number) + _vm.gracePeriod;
    }

    function isErrored(Data storage _vm) external view returns(bool) {
        return _vm.machineHash == MACHINE_ERROR_HASH;
    }

    function isHalted(Data storage _vm) external view returns(bool) {
        return _vm.machineHash == MACHINE_HALT_HASH;
    }
}
