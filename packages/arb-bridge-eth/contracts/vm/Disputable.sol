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

    event PendingDisputableAssertion (
        bytes32 beforeHash,
        uint64[2] timeBounds,
        bytes32 beforeInbox,
        bytes21[] tokenTypes,
        uint256[] beforeBalances,
        bytes32 assertionHash,
        address asserter,
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

    function pendingDisputableAssert(
        VM.Data storage _vm,
        bytes32 _beforeHash,
        uint64[2] memory _timeBounds,
        bytes32 _beforeInbox,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances,
        uint32 _numSteps,
        bytes32 _assertionHash
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
        require(_beforeHash == _vm.machineHash, "Precondition: state hash does not match");
        require(_beforeInbox == _vm.inbox, "Precondition: inbox does not match");

        VM.resetDeadline(_vm);

        _vm.pendingHash = keccak256(
            abi.encodePacked(
                ArbProtocol.generatePreconditionHash(
                    _beforeHash,
                    _timeBounds,
                    _beforeInbox,
                    _tokenTypes,
                    _beforeBalances
                ),
                _assertionHash,
                _numSteps
            )
        );
        _vm.asserter = msg.sender;
        _vm.state = VM.State.PendingDisputable;

        emit PendingDisputableAssertion(
            _beforeHash,
            _timeBounds,
            _beforeInbox,
            _tokenTypes,
            _beforeBalances,
            _assertionHash,
            msg.sender,
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

    function initiateChallenge(
        VM.Data storage _vm,
        bytes32 _preconditionHash,
        bytes32 _assertionHash,
        uint32 _numSteps
    )
        public
    {
        require(msg.sender != _vm.asserter, "Challenge was created by asserter");
        require(VM.withinDeadline(_vm), "Challenge did not come before deadline");
        require(_vm.state == VM.State.PendingDisputable, "Assertion must be pending to initiate challenge");

        require(
            keccak256(
                abi.encodePacked(
                    _preconditionHash,
                    _assertionHash,
                    _numSteps
                )
            ) == _vm.pendingHash,
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
                        generateLastMessageHash(
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
                    ),
                    _data.numSteps
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

    function generateLastMessageHash(
        bytes21[] memory _tokenTypes,
        bytes memory _data,
        uint16[] memory _tokenNums,
        uint256[] memory _amounts,
        address[] memory _destinations
    )
        private
        pure
        returns (bytes32)
    {
        require(_amounts.length == _destinations.length, "Input size mismatch");
        require(_amounts.length == _tokenNums.length, "Input size mismatch");
        bytes32 hashVal = 0x00;
        uint256 offset = 0;
        bytes32 msgHash;
        uint amountCount = _amounts.length;
        for (uint i = 0; i < amountCount; i++) {
            (offset, msgHash) = ArbValue.deserializeValidValueHash(_data, offset);
            msgHash = ArbProtocol.generateMessageStubHash(
                msgHash,
                _tokenTypes[_tokenNums[i]],
                _amounts[i],
                _destinations[i]
            );
            hashVal = keccak256(abi.encodePacked(hashVal, msgHash));
        }
        return hashVal;
    }
}
