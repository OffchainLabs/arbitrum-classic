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

import "./VM.sol";

import "../libraries/ArbProtocol.sol";

import "@openzeppelin/contracts/math/SafeMath.sol";


library Disputable {
    using SafeMath for uint256;

    // fields:
        // beforeHash
        // beforeInbox
        // afterHash
        // messagesAccHash
        // logsAccHash

    event PendingDisputableAssertion(
        bytes32[5] fields,
        address asserter,
        uint64[2] timeBounds,
        uint32 numSteps,
        uint64 deadline
    );

    event ConfirmedDisputableAssertion(
        bytes32 newState,
        bytes32 logsAccHash
    );

    function pendingDisputableAssert(
        VM.Data storage vm,
        bytes32 beforeHash,
        bytes32 beforeInbox,
        bytes32 afterHash,
        bytes32 messagesAccHash,
        bytes32 logsAccHash,
        uint32 numSteps,
        uint64[2] memory timeBounds
    )
        public
    {
        require(vm.state == VM.State.Waiting, "Can only disputable assert from waiting state");
        require(
            !VM.isErrored(vm) && !VM.isHalted(vm),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(vm.activeChallengeManager == address(0), "Can only disputable assert if not in challenge");
        require(numSteps <= vm.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(timeBounds), "Precondition: not within time bounds");
        require(beforeHash == vm.machineHash, "Precondition: state hash does not match");
        require(beforeInbox == vm.inbox, "Precondition: inbox does not match");

        VM.resetDeadline(vm);

        vm.pendingHash = keccak256(
            abi.encodePacked(
                ArbProtocol.generatePreconditionHash(
                    beforeHash,
                    timeBounds,
                    beforeInbox
                ),
                ArbProtocol.generateAssertionHash(
                    afterHash,
                    numSteps,
                    0x00,
                    messagesAccHash,
                    0x00,
                    logsAccHash
                )
            )
        );
        vm.asserter = msg.sender;
        vm.state = VM.State.PendingDisputable;

        emit PendingDisputableAssertion(
            [beforeHash, beforeInbox, afterHash, messagesAccHash, logsAccHash],
            msg.sender,
            timeBounds,
            numSteps,
            vm.deadline
        );
    }

    function confirmDisputableAsserted(
        VM.Data storage vm,
        bytes32 preconditionHash,
        bytes32 afterHash,
        uint32 numSteps,
        bytes memory messages,
        bytes32 logsAccHash
    )
        public
    {
        require(vm.state == VM.State.PendingDisputable, "VM does not have assertion pending");
        require(!VM.withinDeadline(vm), "Assertion is still pending challenge");
        require(
            keccak256(
                abi.encodePacked(
                    preconditionHash,
                    ArbProtocol.generateAssertionHash(
                        afterHash,
                        numSteps,
                        0x00,
                        ArbProtocol.generateLastMessageHash(messages),
                        0x00,
                        logsAccHash
                    )
                )
            ) == vm.pendingHash,
            "Confirm Disputable: Precondition and assertion do not match pending assertion"
        );
        VM.acceptAssertion(
            vm,
            afterHash
        );

        emit ConfirmedDisputableAssertion(
            afterHash,
            logsAccHash
        );
    }

    function initiateChallenge(
        VM.Data storage _vm,
        bytes32 beforeHash,
        bytes32 beforeInbox,
        uint64[2] memory timeBounds,
        bytes32 assertionHash
    ) public {
        require(msg.sender != _vm.asserter, "Challenge was created by asserter");
        require(VM.withinDeadline(_vm), "Challenge did not come before deadline");
        require(_vm.state == VM.State.PendingDisputable, "Assertion must be pending to initiate challenge");

        require(
            keccak256(
                abi.encodePacked(
                    ArbProtocol.generatePreconditionHash(
                        beforeHash,
                        timeBounds,
                        beforeInbox
                    ),
                    assertionHash
                )
            ) == _vm.pendingHash,
            "Initiate Challenge: Precondition and assertion do not match pending assertion"
        );

        _vm.pendingHash = 0;
        _vm.state = VM.State.Waiting;
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) public view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }
}
