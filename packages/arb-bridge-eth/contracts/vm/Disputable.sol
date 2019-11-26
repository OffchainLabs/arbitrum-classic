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

import "../libraries/ArbValue.sol";
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

    event PendingDisputableAssertion (
        bytes32[5] fields,
        address asserter,
        uint64[2] timeBounds,
        uint32 numSteps
    );

    event ConfirmedDisputableAssertion(
        bytes32 newState,
        bytes32 logsAccHash
    );

    event InitiatedChallenge(
        address challenger
    );

    struct ConfirmDisputableAssertedData {
        bytes32 preconditionHash;
        bytes32 afterHash;
        uint32  numSteps;
        bytes21[] tokenTypes;
        VM.FullAssertion assertion;
    }

    // fields:
    // _beforeHash
    // _beforeInbox
    // _afterHash
    // _messagesAccHash
    // _logsAccHash

    function pendingDisputableAssert(
        VM.Data storage _vm,
        bytes32[5] memory _fields,
        uint32 _numSteps,
        uint64[2] memory _timeBounds
    )
        public
    {
        require(_vm.state == VM.State.Waiting, "Can only disputable assert from waiting state");
        require(
            !VM.isErrored(_vm) && !VM.isHalted(_vm),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(!_vm.inChallenge, "Can only disputable assert if not in challenge");
        require(_numSteps <= _vm.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(_timeBounds), "Precondition: not within time bounds");
        require(_fields[0] == _vm.machineHash, "Precondition: state hash does not match");
        require(_fields[1] == _vm.inbox, "Precondition: inbox does not match");

        VM.resetDeadline(_vm);

        _vm.pendingHash = keccak256(
            abi.encodePacked(
                ArbProtocol.generatePreconditionHash(
                    _fields[0],
                    _timeBounds,
                    _fields[1]
                ),
                ArbProtocol.generateAssertionHash(
                    _fields[2],
                    _numSteps,
                    0x00,
                    _fields[3],
                    0x00,
                    _fields[4]
                )
            )
        );
        _vm.asserter = msg.sender;
        _vm.state = VM.State.PendingDisputable;

        emit PendingDisputableAssertion(
            _fields,
            msg.sender,
            _timeBounds,
            _numSteps
        );
    }

    function confirmDisputableAsserted(
        VM.Data storage _vm,
        bytes32 _preconditionHash,
        bytes32 _afterHash,
        uint32 _numSteps,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations,
        bytes32 _logsAccHash
    )
        public
    {
        return _confirmDisputableAsserted(
            _vm,
            ConfirmDisputableAssertedData(
                _preconditionHash,
                _afterHash,
                _numSteps,
                _tokenTypes,
                VM.FullAssertion(
                    _messageData,
                    _messageTokenNums,
                    _messageAmounts,
                    _messageDestinations,
                    _logsAccHash
                )
            )
        );
    }

    function initiateChallenge(VM.Data storage _vm, bytes32 _assertPreHash) public {
        require(msg.sender != _vm.asserter, "Challenge was created by asserter");
        require(VM.withinDeadline(_vm), "Challenge did not come before deadline");
        require(_vm.state == VM.State.PendingDisputable, "Assertion must be pending to initiate challenge");

        require(
            _assertPreHash == _vm.pendingHash,
            "Initiate Challenge: Precondition and assertion do not match pending assertion"
        );

        _vm.pendingHash = 0;
        _vm.state = VM.State.Waiting;
        _vm.inChallenge = true;

        emit InitiatedChallenge(
            msg.sender
        );
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) public view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }

    function _confirmDisputableAsserted(
        VM.Data storage _vm,
        ConfirmDisputableAssertedData memory _data
    )
        internal
    {
        require(_vm.state == VM.State.PendingDisputable, "VM does not have assertion pending");
        require(!VM.withinDeadline(_vm), "Assertion is still pending challenge");
        require(
            keccak256(
                abi.encodePacked(
                    _data.preconditionHash,
                    ArbProtocol.generateAssertionHash(
                        _data.afterHash,
                        _data.numSteps,
                        0x00,
                        ArbProtocol.generateLastMessageHash(
                            _data.tokenTypes,
                            _data.assertion.messageData,
                            _data.assertion.messageTokenNums,
                            _data.assertion.messageAmounts,
                            _data.assertion.messageDestinations
                        ),
                        0x00,
                        _data.assertion.logsAccHash
                    )
                )
            ) == _vm.pendingHash,
            "Confirm Disputable: Precondition and assertion do not match pending assertion"
        );
        VM.acceptAssertion(
            _vm,
            _data.afterHash
        );

        emit ConfirmedDisputableAssertion(
            _data.afterHash,
            _data.assertion.logsAccHash
        );
    }
}
