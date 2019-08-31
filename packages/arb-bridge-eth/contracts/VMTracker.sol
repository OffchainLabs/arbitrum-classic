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
import "./ArbValue.sol";
import "./IChallengeManager.sol";
import "./ArbBalanceTracker.sol";
import "./MerkleLib.sol";
import "./VM.sol";
import "./Disputable.sol";
import "./Unanimous.sol";

import "@openzeppelin/contracts/ownership/Ownable.sol";


contract VMTracker is Ownable {
    using SafeMath for uint256;
    using BytesLib for bytes;

    event MessageDelivered(
        bytes32 indexed vmId,
        bytes32 destination,
        bytes21 tokenType,
        uint256 value,
        bytes data
    );

    event VMCreated(
        bytes32 indexed vmId,
        uint32 _gracePeriod,
        uint128 _escrowRequired,
        address _escrowCurrency,
        uint32 _maxExecutionSteps,
        bytes32 _vmState,
        uint16 _challengeManagerNum,
        address _owner,
        address[] validators
    );

    event PendingUnanimousAssertion (
        bytes32 indexed vmId,
        bytes32 unanHash,
        uint64 sequenceNum
    );

    event ConfirmedUnanimousAssertion (
        bytes32 indexed vmId,
        uint64 sequenceNum
    );

    event FinalizedUnanimousAssertion(
        bytes32 indexed vmId,
        bytes32 unanHash
    );

    // fields:
    // beforeHash
    // beforeInbox
    // afterHash

    event PendingDisputableAssertion (
        bytes32 indexed vmId,
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
        bytes32 indexed vmId,
        bytes32 newState,
        bytes32 logsAccHash
    );

    event InitiatedChallenge(
        bytes32 indexed vmId,
        address challenger
    );

    address internal constant ETH_ADDRESS = address(0);

    IChallengeManager[] challengeManagers;
    ArbBalanceTracker arbBalanceTracker;
    mapping(bytes32 => VM.Data) vms;
    mapping(address => bool) acceptedCurrencies;

    constructor(address _balanceTrackerAddress) public {
        arbBalanceTracker = ArbBalanceTracker(_balanceTrackerAddress);

        acceptedCurrencies[ETH_ADDRESS] = true;
    }

    function ownerShutdown(bytes32 _vmId) external {
        VM.Data storage vm = vms[_vmId];
        require(msg.sender == vm.owner, "Only owner can shutdown the VM");
        _shutdown(vm);
    }

    function completeChallenge(bytes32 _vmId, address[2] calldata _players, uint128[2] calldata _rewards) external {
        VM.Data storage vm = vms[_vmId];
        require(
            msg.sender == address(challengeManagers[vm.challengeManagersNum]),
            "Only challenge manager can complete challenge"
        );
        require(vm.inChallenge, "VM must be in challenge to complete it");

        vm.inChallenge = false;
        vm.validatorBalances[_players[0]] = vm.validatorBalances[_players[0]].add(_rewards[0]);
        vm.validatorBalances[_players[1]] = vm.validatorBalances[_players[1]].add(_rewards[1]);
    }

    function addChallengeManager(IChallengeManager _challengeManager) public onlyOwner {
        challengeManagers.push(_challengeManager);
    }

    // fields
    // _vmId
    // _vmState
    // _createHash

    function createVm(
        bytes32[3] memory _fields,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint16 _challengeManagerNum,
        uint128 _escrowRequired,
        address _escrowCurrency,
        address _owner,
        bytes memory _signatures
    )
        public
    {
        require(_signatures.length / 65 < 2 ** 16, "Too many validators");
        require(bytes32(bytes20(_fields[0])) != _fields[0], "Invalid vmId");
        require(_challengeManagerNum < challengeManagers.length, "Invalid challenge manager num");
        require(_escrowRequired > 0, "VM must require non-zero deposit");
        require(acceptedCurrencies[_escrowCurrency], "Selected currency is not an accepted type");

        address[] memory assertKeys = ArbProtocol.recoverAddresses(_fields[2], _signatures);
        require(
            keccak256(
                abi.encodePacked(
                    _gracePeriod,
                    _escrowRequired,
                    _escrowCurrency,
                    _maxExecutionSteps,
                    _fields[1],
                    _challengeManagerNum,
                    _owner,
                    assertKeys
                )
            ) == _fields[2],
            "Create data incorrect"
        );

        for (uint i = 0; i < assertKeys.length; i++) {
            arbBalanceTracker.ownerRemoveToken(
                bytes32(bytes20(assertKeys[i])),
                _escrowCurrency,
                _escrowRequired
            );
        }

        VM.Data storage vm = vms[_fields[0]];

        // Machine state
        vm.machineHash = _fields[1];
        vm.inboxHash = ArbValue.hashEmptyTuple();
        vm.pendingMessages = ArbValue.hashEmptyTuple();
        vm.challengeManagersNum = _challengeManagerNum;
        vm.state = VM.State.Waiting;
        vm.pendingHash = 0;

        vm.id = _fields[0];

        // Validator options
        vm.validatorRoot = MerkleLib.generateAddressRoot(assertKeys);
        vm.validatorCount = uint16(assertKeys.length);
        vm.escrowRequired = _escrowRequired;
        vm.currencyType = _escrowCurrency;
        vm.owner = _owner;
        vm.gracePeriod = _gracePeriod;
        vm.maxExecutionSteps = _maxExecutionSteps;

        for (uint i = 0; i < assertKeys.length; i++) {
            vm.validatorBalances[assertKeys[i]] = _escrowRequired;
        }

        emit VMCreated(
            _fields[0],
            _gracePeriod,
            _escrowRequired,
            _escrowCurrency,
            _maxExecutionSteps,
            _fields[1],
            _challengeManagerNum,
            _owner,
            assertKeys
        );
    }

    function sendMessage(
        bytes32 _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes memory _data
    )
        public
    {
        _sendUnpaidMessage(
            _destination,
            _tokenType,
            _amount,
            bytes32(uint256(uint160(msg.sender))),
            _data
        );
    }

    function forwardMessage(
        bytes32 _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes memory _data,
        bytes memory _signature
    )
        public
    {
        address sender = ArbProtocol.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _destination,
                    ArbValue.deserializeValueHash(_data),
                    _amount,
                    _tokenType
                )
            ),
            _signature
        );

        _sendUnpaidMessage(
            _destination,
            _tokenType,
            _amount,
            bytes32(uint256(uint160(sender))),
            _data
        );
    }

    function sendEthMessage(bytes32 _destination, bytes memory _data) public payable {
        arbBalanceTracker.depositEth.value(msg.value)(_destination);
        _deliverMessage(
            _destination,
            bytes21(0),
            msg.value,
            bytes32(uint256(uint160(msg.sender))),
            _data
        );
    }

    // fields:
    // _beforeHash
    // _beforeInbox
    // _afterHash
    // _logsAccHash

    function pendingDisputableAssert(
        bytes32 _vmId,
        bytes32[4] memory _fields,
        uint32 _numSteps,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        bytes32[] memory _messageDataHash,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        bytes32[] memory _messageDestinations
    )
        public
    {
        VM.Data storage vm = vms[_vmId];
        uint256[] memory beforeBalances = ArbProtocol.calculateBeforeValues(
            _tokenTypes,
            _messageTokenNums,
            _messageAmounts
        );

        require(
            arbBalanceTracker.hasFunds(
                _vmId,
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
        bytes32 _vmId,
        bytes32 _preconditionHash,
        bytes32 _afterHash,
        uint32 _numSteps,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        bytes32[] memory _messageDestinations,
        bytes32 _logsAccHash
    )
        public
    {
        VM.Data storage vm = vms[_vmId];
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

        _deliverSentMessages(
            _vmId,
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function initiateChallenge(bytes32 _vmId, bytes32 _assertPreHash) public {
        VM.Data storage vm = vms[_vmId];
        Disputable.initiateChallenge(
            vm,
            _assertPreHash
        );

        challengeManagers[vm.challengeManagersNum].initiateChallenge(
            _vmId,
            [vm.asserter, msg.sender],
            [vm.escrowRequired, vm.escrowRequired],
            vm.gracePeriod,
            _assertPreHash
        );
    }

    // fields
    //   _afterHash
    //   _newInbox
    //   _logsAccHash

    function finalizedUnanimousAssert(
        bytes32 _vmId,
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        bytes32[] memory _messageDestinations,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        public
    {
        Unanimous.finalizedUnanimousAssert(
            vms[_vmId],
            [_afterHash, _newInbox, _logsAccHash],
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations,
            _signatures
        );

        _deliverSentMessages(
            _vmId,
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function pendingUnanimousAssert(
        bytes32 _vmId,
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
        require(
            arbBalanceTracker.hasFunds(
                _vmId,
                _tokenTypes,
                ArbProtocol.calculateBeforeValues(
                    _tokenTypes,
                    _messageTokenNums,
                    _messageAmounts
                )
            ),
            "VM has insufficient balance"
        );
        VM.Data storage vm = vms[_vmId];
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
        bytes32 _vmId,
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        bytes32[] memory _messageDestinations
    )
        public
    {
        VM.Data storage vm = vms[_vmId];
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
        _deliverSentMessages(
            vm.id,
            _tokenTypes,
            _messageData,
            _messageTokenNums,
            _messageAmounts,
            _messageDestinations
        );
    }

    function _sendUnpaidMessage(
        bytes32 _destination,
        bytes21 _tokenType,
        uint256 _value,
        bytes32 _sender,
        bytes memory _data
    )
        internal
    {
        if (_tokenType[20] == 0x01) {
            arbBalanceTracker.transferNFT(
                _sender,
                _destination,
                address(bytes20(_tokenType)),
                _value
            );
        } else {
            arbBalanceTracker.transferToken(
                _sender,
                _destination,
                address(bytes20(_tokenType)),
                _value
            );
        }
        _deliverMessage(
            _destination,
            _tokenType,
            _value,
            _sender,
            _data
        );
    }

    function _deliverMessage(
        bytes32 _destination,
        bytes21 _tokenType,
        uint256 _value,
        bytes32 _sender,
        bytes memory _data
    )
        internal
    {
        if (bytes32(bytes20(_destination)) != _destination) {
            VM.Data storage vm = vms[_destination];
            VM.deliverMessage(
                vm,
                _tokenType,
                _value,
                _sender,
                _data
            );
        }

        emit MessageDelivered(
            _destination,
            _sender,
            _tokenType,
            _value,
            _data
        );
    }

    function _deliverSentMessages(
        bytes32 _vmId,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _tokenTypeNum,
        uint256[] memory _amounts,
        bytes32[] memory _destinations
    )
        internal
    {
        uint offset = 0;
        bytes memory msgData;
        uint amountCount = _amounts.length;
        for (uint i = 0; i < amountCount; i++) {
            (offset, msgData) = ArbValue.getNextValidValue(_messageData, offset);
            _sendUnpaidMessage(
                _destinations[i],
                _tokenTypes[_tokenTypeNum[i]],
                _amounts[i],
                _vmId,
                msgData
            );
        }
    }

    function _shutdown(VM.Data storage _vm) private {
        // TODO: transfer all owned funds to halt address
        delete vms[_vm.id];
    }
}
