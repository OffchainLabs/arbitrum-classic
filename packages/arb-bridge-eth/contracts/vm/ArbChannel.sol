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

import "./ArbitrumVM.sol";
import "./Unanimous.sol";


contract ArbChannel is ArbitrumVM {
    using SafeMath for uint256;

    event PendingUnanimousAssertion (
        bytes32 unanHash,
        uint64 sequenceNum
    );

    event ConfirmedUnanimousAssertion (
        uint64 sequenceNum
    );

    event FinalizedUnanimousAssertion(
        bytes32 unanHash
    );

    constructor(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _escrowRequired,
        address payable _owner,
        address _challengeManagerAddress,
        address _globalInboxAddress,
        address[] memory _validatorKeys
    )
        ArbitrumVM(
            _vmState,
            _gracePeriod,
            _maxExecutionSteps,
            _escrowRequired,
            _owner,
            _challengeManagerAddress,
            _globalInboxAddress
        )
        public
    {
        uint16 validatorCount = uint16(_validatorKeys.length);
        vm.validatorCount = validatorCount;
        for (uint16 i = 0; i < validatorCount; i++) {
            vm.validators[_validatorKeys[i]] = VM.Validator(0, true);
        }
    }

    function increaseDeposit() external payable validatorOnly {
        VM.Validator storage validator = vm.validators[msg.sender];
        bool wasInactive = validator.balance < uint256(vm.escrowRequired);
        vm.validators[msg.sender].balance += msg.value;
        if (wasInactive && validator.balance >= uint256(vm.escrowRequired)) {
            activatedValidators++;
        }
        if (activatedValidators == vm.validatorCount && vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function isValidatorList(address[] memory _validators) public view returns(bool) {
        return VM.isValidatorList(vm, _validators);
    }

    function finalizedUnanimousAssert(
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        public
    {
        Unanimous.finalizedUnanimousAssert(
            vm,
            [_afterHash, _newInbox, _logsAccHash],
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations,
            _signatures
        );

        _completeAssertion(
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function pendingUnanimousAssert(
        bytes32 _unanRest,
        bytes21[] memory _tokenTypes,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        uint64 _sequenceNum,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        public
    {
        uint256[] memory beforeBalances = ArbProtocol.calculateBeforeValues(
            _tokenTypes,
            _messageTokenNums,
            _messageAmounts
        );
        require(ArbProtocol.beforeBalancesValid(_tokenTypes, beforeBalances), "Token types must be valid and sorted");
        require(
            globalInbox.hasFunds(
                address(this),
                _tokenTypes,
                beforeBalances
            ),
            "VM has insufficient balance"
        );
        Unanimous.pendingUnanimousAssert(
            vm,
            _unanRest,
            _tokenTypes,
            _messageTokenNums,
            _messageAmounts,
            _sequenceNum,
            _logsAccHash,
            _signatures
        );
    }

    function confirmUnanimousAsserted(
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations
    )
        public
    {
        Unanimous.confirmUnanimousAsserted(
            vm,
            _afterHash,
            _newInbox,
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );

        _completeAssertion(
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }
}
