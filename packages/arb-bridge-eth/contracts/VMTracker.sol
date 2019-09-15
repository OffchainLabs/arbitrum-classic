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

import "./ArbProtocol.sol";
import "./IChallengeManager.sol";
import "./IGlobalPendingInbox.sol";
import "./VM.sol";
import "./Disputable.sol";
import "./Unanimous.sol";
import "./SigUtils.sol";


contract VMTracker {
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

    // fields:
    // beforeHash
    // beforeInbox
    // afterHash

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

    address internal constant ETH_ADDRESS = address(0);

    IChallengeManager public challengeManager;
    IGlobalPendingInbox public globalInbox;

    VM.Data public vm;
    address public escrowCurrency;

    uint16 public activatedValidators;
    address payable public exitAddress;
    address payable public terminateAddress;
    address payable public owner;

    modifier validatorOnly() {
        require(vm.validators[msg.sender].valid, "Caller must be validator");
        _;
    }

    constructor(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _escrowRequired,
        address _escrowCurrency,
        address payable _owner,
        address _challengeManagerAddress,
        address _globalInboxAddress,
        address[] memory _validatorKeys
    )
        public
    {
        require(
            _escrowCurrency == ETH_ADDRESS,
            "Validator deposits must be in ETH"
        );
        globalInbox = IGlobalPendingInbox(_globalInboxAddress);
        challengeManager = IChallengeManager(_challengeManagerAddress);

        globalInbox.registerForInbox();
        owner = _owner;
        activatedValidators = 0;
        escrowCurrency = _escrowCurrency;

        uint16 validatorCount = uint16(_validatorKeys.length);
        vm.validatorCount = validatorCount;
        for (uint16 i = 0; i < validatorCount; i++) {
            vm.validators[_validatorKeys[i]] = VM.Validator(0, true);
        }

        // Machine state
        vm.machineHash = _vmState;
        vm.state = VM.State.Uninitialized;
        vm.inbox = ArbValue.hashEmptyTuple();

        // Validator options
        vm.escrowRequired = _escrowRequired;
        vm.gracePeriod = _gracePeriod;
        vm.maxExecutionSteps = _maxExecutionSteps;
    }

    function isListedValidator(address validator) external view returns(bool) {
        return vm.validators[validator].valid;
    }

    function currentDeposit(address validator) external view returns(uint256) {
        return vm.validators[validator].balance;
    }

    function escrowRequired() external view returns(uint256) {
        return vm.escrowRequired;
    }

    function getState() external view returns(VM.State) {
        return vm.state;
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

    function ownerShutdown() external {
        require(msg.sender == owner, "Only owner can shutdown the VM");
        _shutdown();
    }

    function completeChallenge(address[2] calldata _players, uint128[2] calldata _rewards) external {
        require(
            msg.sender == address(challengeManager),
            "Only challenge manager can complete challenge"
        );
        require(vm.inChallenge, "VM must be in challenge to complete it");

        vm.inChallenge = false;
        vm.validators[_players[0]].balance = vm.validators[_players[0]].balance.add(_rewards[0]);
        vm.validators[_players[1]].balance = vm.validators[_players[1]].balance.add(_rewards[1]);
    }

    function isValidatorList(address[] memory _validators) public view returns(bool) {
        uint validatorCount = _validators.length;
        if (validatorCount != vm.validatorCount) {
            return false;
        }
        for (uint i = 0; i < validatorCount; i++) {
            if (!vm.validators[_validators[i]].valid) {
                return false;
            }
        }
        return true;
    }

    // fields:
    // _beforeHash
    // _beforeInbox
    // _afterHash
    // _logsAccHash

    function pendingDisputableAssert(
        bytes32[4] memory _fields,
        uint32 _numSteps,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        bytes32[] memory _messageDataHash,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations
    )
        public
        validatorOnly
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
        Disputable.pendingDisputableAssert(
            vm,
            _fields,
            _numSteps,
            _timeBounds,
            _tokenTypes,
            _messageDataHash,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function confirmDisputableAsserted(
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
        Disputable.confirmDisputableAsserted(
            vm,
            _preconditionHash,
            _afterHash,
            _numSteps,
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations,
            _logsAccHash
        );

        _pushPendingToInbox();

        globalInbox.sendMessages(
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function initiateChallenge(bytes32 _assertPreHash) public validatorOnly {
        Disputable.initiateChallenge(
            vm,
            _assertPreHash
        );

        challengeManager.initiateChallenge(
            [vm.asserter, msg.sender],
            [vm.escrowRequired, vm.escrowRequired],
            vm.gracePeriod,
            _assertPreHash
        );
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

        _pushPendingToInbox();

        globalInbox.sendMessages(
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

        _pushPendingToInbox();

        globalInbox.sendMessages(
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function _shutdown() private {
        // TODO: transfer all owned funds to halt address
        selfdestruct(owner);
    }

    function _pushPendingToInbox() private {
        bytes32 pending = globalInbox.pullPendingMessages(address(this));
        if (pending != ArbValue.hashEmptyTuple()) {
            vm.inbox = ArbProtocol.appendInboxMessages(vm.inbox, pending);
        }
    }
}
