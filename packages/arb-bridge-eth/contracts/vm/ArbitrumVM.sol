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
import "./Disputable.sol";

import "../IGlobalPendingInbox.sol";

import "../challenge/IChallengeManager.sol";

import "../libraries/ArbProtocol.sol";


contract ArbitrumVM {
    using SafeMath for uint256;

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
        address payable _owner,
        address _challengeManagerAddress,
        address _globalInboxAddress
    )
        public
    {
        globalInbox = IGlobalPendingInbox(_globalInboxAddress);
        challengeManager = IChallengeManager(_challengeManagerAddress);

        globalInbox.registerForInbox();
        owner = _owner;
        activatedValidators = 0;

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

        _completeAssertion(
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

    function _completeAssertion(
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        address[] memory _messageDestinations
    )
        internal
    {
        bytes32 pending = globalInbox.pullPendingMessages(address(this));
        if (pending != ArbValue.hashEmptyTuple()) {
            vm.inbox = ArbProtocol.appendInboxMessages(vm.inbox, pending);
        }

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
}
