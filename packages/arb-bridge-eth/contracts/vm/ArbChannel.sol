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

import "./IArbChannel.sol";
import "./ArbitrumVM.sol";
import "./Unanimous.sol";


contract ArbChannel is ArbitrumVM, IArbChannel {
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

    mapping (address => bool) validators;
    uint16 public validatorCount;
    uint16 public activatedValidators;

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
        activatedValidators = 0;
        validatorCount = uint16(_validatorKeys.length);
        for (uint16 i = 0; i < validatorCount; i++) {
            validators[_validatorKeys[i]] = true;
        }
    }

    function increaseDeposit() external payable {
        require(validators[msg.sender], "Caller must be validator");
        uint256 balance = validatorBalances[msg.sender];
        bool wasInactive = balance < uint256(vm.escrowRequired);
        balance += msg.value;
        validatorBalances[msg.sender] = balance;
        if (wasInactive && balance >= uint256(vm.escrowRequired)) {
            activatedValidators++;
        }
        if (activatedValidators == validatorCount && vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function isListedValidator(address validator) external view returns(bool) {
        return validators[validator];
    }

    function isValidatorList(address[] calldata _validators) external view returns(bool) {
        uint _validatorCount = _validators.length;
        if (_validatorCount != validatorCount) {
            return false;
        }
        for (uint i = 0; i < validatorCount; i++) {
            if (!validators[_validators[i]]) {
                return false;
            }
        }
        return true;
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
            this,
            [_afterHash, _newInbox, _logsAccHash],
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations,
            _signatures
        );

        if (vm.state == VM.State.PendingDisputable) {
            validatorBalances[vm.asserter] = validatorBalances[vm.asserter].add(vm.escrowRequired);
        }

        VM.acceptAssertion(
            vm,
            _afterHash
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
        Unanimous.pendingUnanimousAssert(
            vm,
            this,
            _unanRest,
            _tokenTypes,
            _messageTokenNums,
            _messageAmounts,
            _sequenceNum,
            _logsAccHash,
            _signatures
        );

        if (vm.state == VM.State.PendingDisputable) {
            validatorBalances[vm.asserter] = validatorBalances[vm.asserter].add(vm.escrowRequired);
        }
        vm.state = VM.State.PendingUnanimous;
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
