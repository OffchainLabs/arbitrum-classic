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

import "./ArbValue.sol";
import "./ArbProtocol.sol";

import "@openzeppelin/contracts/math/SafeMath.sol";


library VM {
    using SafeMath for uint256;

    bytes32 private constant MACHINE_HALT_HASH = bytes32(0);
    bytes32 private constant MACHINE_ERROR_HASH = bytes32(uint(1));

    enum State {
        Waiting,
        PendingAssertion,
        PendingUnanimous
    }

    struct Data {
        bytes32 machineHash;
        bytes32 pendingHash; // Lock pending and confirm asserts together
        bytes32 inboxHash;
        bytes32 pendingMessages;
        bytes32 validatorRoot;
        bytes32 exitAddress;
        bytes32 terminateAddress;
        bytes32 id;
        address owner;
        address asserter;
        address currencyType;
        uint128 escrowRequired;
        uint64 deadline;
        uint64 sequenceNum;
        uint32 gracePeriod;
        uint32 maxExecutionSteps;
        uint16 validatorCount;
        uint16 challengeManagersNum;
        State state;
        bool inChallenge;
        mapping(address => uint256) validatorBalances;
    }

    struct FullAssertion {
        bytes messageData;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        bytes32[] messageDestinations;
        bytes32 logsAccHash;
    }

    function acceptAssertion(Data storage _vm, bytes32 _afterHash) external {
        _vm.machineHash = _afterHash;
        _vm.state = VM.State.Waiting;

        if (_vm.pendingMessages != ArbValue.hashEmptyTuple()) {
            _vm.inboxHash = ArbProtocol.appendInboxMessages(_vm.inboxHash, _vm.pendingMessages);
            _vm.pendingMessages = ArbValue.hashEmptyTuple();
        }
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

    function cancelCurrentState(Data storage vm) external {
        if (vm.state != VM.State.Waiting) {
            require(block.number <= vm.deadline, "Can't cancel finalized state");
        }

        if (vm.state == VM.State.PendingAssertion) {
            // If there is a pending disputable assertion, cancel it
            vm.validatorBalances[vm.asserter] = vm.validatorBalances[vm.asserter].add(vm.escrowRequired);
        }
    }

    function deliverMessage(
        Data storage _vm,
        bytes21 _tokenType,
        uint256 _value,
        bytes32 _sender,
        bytes calldata _data
    )
        external
    {
        bytes32 messageHash = ArbProtocol.generateSentMessageHash(
            _vm.id,
            ArbValue.deserializeValueHash(_data),
            _tokenType,
            _value,
            _sender
        );
        _vm.pendingMessages = ArbProtocol.appendInboxPendingMessage(
            _vm.pendingMessages,
            messageHash
        );
    }

}
