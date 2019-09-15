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
import "./ArbProtocol.sol";

import "@openzeppelin/contracts/math/SafeMath.sol";


library Disputable {
    using SafeMath for uint256;

    // fields
    //    beforeHash
    //    beforeInbox
    //    afterHash

    event PendingDisputableAssertion (
        bytes32[3] fields,
        address asserter,
        uint64[2] timeBounds,
        bytes21[] tokenTypes,
        uint32 numSteps,
        bytes32 lastMessageHash,
        bytes32 logsAccHash,
        uint256[] amounts
    );

    event ConfirmedDisputableAssertion(
        bytes32 newState,
        bytes32 logsAccHash
    );

    event InitiatedChallenge(
        address challenger
    );

    struct PendingDisputableAssertData {
        bytes32 beforeHash;
        bytes32 beforeInbox;
        bytes32 afterHash;
        uint32 numSteps;
        uint64[2] timeBounds;
        bytes21[] tokenTypes;
        bytes32[] messageDataHash;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        address[] messageDestinations;
        bytes32 logsAccHash;
    }

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
    // _logsAccHash

    function pendingDisputableAssert(
        VM.Data storage _vm,
        bytes32[4] memory _fields,
        uint32 _numSteps,
        uint64[2] memory timeBounds,
        bytes21[] memory _tokenTypes,
        bytes32[] memory _messageDataHash,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations
    )
        public
    {
        return _pendingDisputableAssert(
            _vm,
            PendingDisputableAssertData(
                _fields[0],
                _fields[1],
                _fields[2],
                _numSteps,
                timeBounds,
                _tokenTypes,
                _messageDataHash,
                _messageTokenNums,
                _messageAmounts,
                _messageDestinations,
                _fields[3]
            )
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
        require(_vm.escrowRequired <= _vm.validators[msg.sender].balance, "Challenger did not have enough escrowed");

        require(
            _assertPreHash == _vm.pendingHash,
            "Precondition and assertion do not match pending assertion"
        );

        _vm.validators[msg.sender].balance -= _vm.escrowRequired;
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

    function _pendingDisputableAssert(
        VM.Data storage _vm,
        PendingDisputableAssertData memory _data
    )
        internal
    {
        require(_vm.state == VM.State.Waiting, "Can only disputable assert from waiting state");
        require(
            !VM.isErrored(_vm) && !VM.isHalted(_vm),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(!_vm.inChallenge, "Can only disputable assert if not in challenge");
        require(_vm.escrowRequired <= _vm.validators[msg.sender].balance, "Validator does not have required escrow");
        require(_data.numSteps <= _vm.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(_data.timeBounds), "Precondition: not within time bounds");
        require(_data.beforeHash == _vm.machineHash, "Precondition: state hash does not match");
        require(_vm.inbox == _data.beforeInbox, "Precondition: inbox does not match");

        uint256[] memory beforeBalances = ArbProtocol.calculateBeforeValues(
            _data.tokenTypes,
            _data.messageTokenNums,
            _data.messageAmounts
        );

        VM.resetDeadline(_vm);

        bytes32 lastMessageHash = ArbProtocol.generateLastMessageHashStub(
            _data.tokenTypes,
            _data.messageDataHash,
            _data.messageTokenNums,
            _data.messageAmounts,
            _data.messageDestinations
        );

        _vm.pendingHash = keccak256(
            abi.encodePacked(
                ArbProtocol.generatePreconditionHash(
                    _data.beforeHash,
                    _data.timeBounds,
                    _data.beforeInbox,
                    _data.tokenTypes,
                    beforeBalances
                ),
                ArbProtocol.generateAssertionHash(
                    _data.afterHash,
                    _data.numSteps,
                    0x00,
                    lastMessageHash,
                    0x00,
                    _data.logsAccHash,
                    beforeBalances
                )
            )
        );
        _vm.validators[msg.sender].balance -= _vm.escrowRequired;
        _vm.asserter = msg.sender;
        _vm.state = VM.State.PendingDisputable;

        emit PendingDisputableAssertion(
            [_data.beforeHash, _data.beforeInbox, _data.afterHash],
            msg.sender,
            _data.timeBounds,
            _data.tokenTypes,
            _data.numSteps,
            lastMessageHash,
            _data.logsAccHash,
            beforeBalances
        );
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
                        _data.assertion.logsAccHash,
                        ArbProtocol.calculateBeforeValues(
                            _data.tokenTypes,
                            _data.assertion.messageTokenNums,
                            _data.assertion.messageAmounts
                        )
                    )
                )
            ) == _vm.pendingHash,
            "Precondition and assertion do not match pending assertion"
        );
        _vm.validators[_vm.asserter].balance = _vm.validators[_vm.asserter].balance.add(_vm.escrowRequired);
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
