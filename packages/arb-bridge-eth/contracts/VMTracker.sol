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

    enum VMState {
        Waiting,
        PendingAssertion,
        PendingUnanimous
    }

    struct VM {
        bytes32 machineHash;
        bytes32 pendingHash; // Lock pending and confirm asserts together
        bytes32 inboxHash;
        bytes32 pendingMessages;
        bytes32 validatorRoot;
        bytes32 exitAddress;
        bytes32 terminateAddress;
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
        VMState state;
        bool inChallenge;
        mapping(address => uint256) validatorBalances;
    }

    bytes32 internal constant MACHINE_HALT_HASH = bytes32(0);
    bytes32 internal constant MACHINE_ERROR_HASH = bytes32(uint(1));
    address internal constant ETH_ADDRESS = address(0);

    IChallengeManager[] challengeManagers;
    ArbBalanceTracker arbBalanceTracker;
    mapping(bytes32 => VM) vms;
    mapping(address => bool) acceptedCurrencies;

    constructor(address _balanceTrackerAddress) public {
        arbBalanceTracker = ArbBalanceTracker(_balanceTrackerAddress);

        acceptedCurrencies[ETH_ADDRESS] = true;
    }

    function ownerShutdown(bytes32 _vmId) external {
        VM storage vm = vms[_vmId];
        require(msg.sender == vm.owner, "Only owner can shutdown the VM");
        _shutdownVM(_vmId);
    }

    function completeChallenge(bytes32 _vmId, address[2] calldata _players, uint128[2] calldata _rewards) external {
        VM storage vm = vms[_vmId];
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

    function withinTimeBounds(uint64[2] memory _timeBounds) public view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
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

        VM storage vm = vms[_fields[0]];

        // Machine state
        vm.machineHash = _fields[1];
        vm.inboxHash = ArbValue.hashEmptyTuple();
        vm.pendingMessages = ArbValue.hashEmptyTuple();
        vm.challengeManagersNum = _challengeManagerNum;
        vm.state = VMState.Waiting;
        vm.pendingHash = 0;

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

    struct FinalizedUnanimousAssertData {
        bytes32 vmId;
        bytes32 afterHash;
        bytes32 newInbox;
        bytes21[] tokenTypes;
        bytes messageData;
        uint16[] messageTokenNum;
        uint256[] messageAmount;
        bytes32[] messageDestination;
        bytes32 logsAccHash;
        bytes signatures;
    }

    function finalizedUnanimousAssert(
        bytes32 _vmId,
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNum,
        uint256[] memory _messageAmount,
        bytes32[] memory _messageDestination,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        public
    {
        _finalizedUnanimousAssert(
            FinalizedUnanimousAssertData(
                _vmId,
                _afterHash,
                _newInbox,
                _tokenTypes,
                _messageData,
                _messageTokenNum,
                _messageAmount,
                _messageDestination,
                _logsAccHash,
                _signatures
            )
        );
    }

    function pendingUnanimousAssert(
        bytes32 _vmId,
        bytes32 _unanRest,
        bytes21[] memory _tokenTypes,
        uint16[] memory _messageTokenNum,
        uint256[] memory _messageAmount,
        uint64 _sequenceNum,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        public
    {
        VM storage vm = vms[_vmId];
        require(vm.machineHash != MACHINE_HALT_HASH, "Can't assert halted machine");
        bytes32 unanHash = keccak256(
            abi.encodePacked(
                _vmId,
                keccak256(
                    abi.encodePacked(
                        _unanRest,
                        vm.machineHash,
                        vm.inboxHash,
                        _tokenTypes,
                        _messageTokenNum,
                        _messageAmount,
                        _sequenceNum
                    )
                ),
                _logsAccHash
            )
        );
        require(
            MerkleLib.generateAddressRoot(
                ArbProtocol.recoverAddresses(unanHash, _signatures)
            ) == vm.validatorRoot,
            "Validator signatures don't match"
        );

        if (vm.state == VMState.PendingUnanimous) {
            require(
                _sequenceNum > vm.sequenceNum,
                "Can only supersede previous assertion with greater sequence number"
            );
        }

        require(
            arbBalanceTracker.hasFunds(
                _vmId,
                _tokenTypes,
                ArbProtocol.calculateBeforeValues(
                    _tokenTypes,
                    _messageTokenNum,
                    _messageAmount
                )
            ),
            "VM has insufficient balance"
        );

        _cancelCurrentState(vm);
        _resetDeadline(vm);

        vm.state = VMState.PendingUnanimous;
        vm.sequenceNum = _sequenceNum;
        vm.pendingHash = keccak256(
            abi.encodePacked(
                _tokenTypes,
                _messageTokenNum,
                _messageAmount,
                _unanRest
            )
        );

        emit PendingUnanimousAssertion(
            _vmId,
            unanHash,
            _sequenceNum
        );
    }

    function confirmUnanimousAsserted(
        bytes32 _vmId,
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNum,
        uint256[] memory _messageAmount,
        bytes32[] memory _messageDestination
    )
        public
    {
        VM storage vm = vms[_vmId];
        require(vm.state == VMState.PendingUnanimous, "Can only confirm if there is a pending assertion");
        require(block.number > vm.deadline, "Can only confirm assertion whose challenge deadline has passed");
        require(
            keccak256(
                abi.encodePacked(
                    _tokenTypes,
                    _messageTokenNum,
                    _messageAmount,
                    keccak256(
                        abi.encodePacked(
                            _newInbox,
                            _afterHash,
                            _messageData,
                            _messageDestination
                        )
                    )
                )
            ) == vm.pendingHash,
            "Can only confirm assertion that is currently pending"
        );

        vm.inboxHash = _newInbox;
        _acceptAssertion(
            _vmId,
            _afterHash,
            _tokenTypes,
            _messageData,
            _messageTokenNum,
            _messageAmount,
            _messageDestination
        );

        emit ConfirmedUnanimousAssertion(
            _vmId,
            vm.sequenceNum
        );
    }

    struct PendingDisputableAssertData {
        bytes32 vmId;
        bytes32 beforeHash;
        bytes32 beforeInbox;
        bytes32 afterHash;
        bytes32 logsAccHash;
        uint32 numSteps;
        uint64[2] timeBounds;
        bytes21[] tokenTypes;
        bytes32[] messageDataHash;
        uint16[] messageTokenNum;
        uint256[] msgAmount;
        bytes32[] msgDestination;
    }

    // fields:
    // _vmId
    // _beforeHash
    // _beforeInbox
    // _afterHash
    // _logsAccHash

    function pendingDisputableAssert(
        bytes32[5] memory _fields,
        uint32 _numSteps,
        uint64[2] memory timeBounds,
        bytes21[] memory _tokenTypes,
        bytes32[] memory _messageDataHash,
        uint16[] memory _messageTokenNum,
        uint256[] memory _msgAmount,
        bytes32[] memory _msgDestination
    )
        public
    {
        return _pendingDisputableAssert(
            PendingDisputableAssertData(
                _fields[0],
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _numSteps,
                timeBounds,
                _tokenTypes,
                _messageDataHash,
                _messageTokenNum,
                _msgAmount,
                _msgDestination
            )
        );
    }

    // fields:
    // _vmId
    // _preconditionHash
    // _afterHash
    // _logsAccHash

    struct ConfirmDisputableAssertedData {
        bytes32 vmId;
        bytes32 preconditionHash;
        bytes32 afterHash;
        uint32  numSteps;
        bytes21[] tokenTypes;
        bytes messageData;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        bytes32[] messageDestination;
        bytes32 logsAccHash;
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
        bytes32[] memory _messageDestination,
        bytes32 _logsAccHash
    )
        public
    {
        return _confirmDisputableAsserted(
            ConfirmDisputableAssertedData(
                _vmId,
                _preconditionHash,
                _afterHash,
                _numSteps,
                _tokenTypes,
                _messageData,
                _messageTokenNums,
                _messageAmounts,
                _messageDestination,
                _logsAccHash
            )
        );
    }

    event InitiatedChallenge(
        bytes32 indexed vmId,
        address challenger
    );

    // fields
    // _vmId
    // _assertionHash

    function initiateChallenge(bytes32 _vmId, bytes32 _assertPreHash) public {
        VM storage vm = vms[_vmId];
        require(msg.sender != vm.asserter, "Challenge was created by asserter");
        require(block.number <= vm.deadline, "Challenge did not come before deadline");
        require(vm.state == VMState.PendingAssertion, "Assertion must be pending to initiate challenge");
        require(vm.escrowRequired <= vm.validatorBalances[msg.sender], "Challenger did not have enough escrowed");

        require(
            _assertPreHash == vm.pendingHash,
            "Precondition and assertion do not match pending assertion"
        );

        vm.validatorBalances[msg.sender] = vm.validatorBalances[msg.sender].sub(vm.escrowRequired);
        vm.pendingHash = 0;
        vm.state = VMState.Waiting;
        vm.inChallenge = true;

        challengeManagers[vm.challengeManagersNum].initiateChallenge(
            _vmId,
            [vm.asserter, msg.sender],
            [vm.escrowRequired, vm.escrowRequired],
            vm.gracePeriod,
            _assertPreHash
        );
        emit InitiatedChallenge(
            _vmId,
            msg.sender
        );
    }

    function _finalizedUnanimousAssert(FinalizedUnanimousAssertData memory data) internal {
        VM storage vm = vms[data.vmId];
        require(vm.machineHash != MACHINE_HALT_HASH, "Can't assert halted machine");
        bytes32 unanHash = keccak256(
            abi.encodePacked(
                data.vmId,
                keccak256(
                    abi.encodePacked(
                        keccak256(
                            abi.encodePacked(
                                data.newInbox,
                                data.afterHash,
                                data.messageData,
                                data.messageDestination
                            )
                        ),
                        vm.machineHash,
                        vm.inboxHash,
                        data.tokenTypes,
                        data.messageTokenNum,
                        data.messageAmount
                    )
                ),
                data.logsAccHash
            )
        );
        require(
            MerkleLib.generateAddressRoot(
                ArbProtocol.recoverAddresses(unanHash, data.signatures)
            ) == vm.validatorRoot,
            "Validator signatures don't match"
        );

        _cancelCurrentState(vm);
        vm.state = VMState.Waiting;
        vm.inboxHash = data.newInbox;
        _acceptAssertion(
            data.vmId,
            data.afterHash,
            data.tokenTypes,
            data.messageData,
            data.messageTokenNum,
            data.messageAmount,
            data.messageDestination
        );

        emit FinalizedUnanimousAssertion(
            data.vmId,
            unanHash
        );
    }

    function _pendingDisputableAssert(PendingDisputableAssertData memory _data) internal {
        VM storage vm = vms[_data.vmId];
        require(vm.state == VMState.Waiting, "Can only disputable assert from waiting state");
        require(
            vm.machineHash != MACHINE_HALT_HASH && vm.machineHash != MACHINE_ERROR_HASH,
            "Can only disputable assert if machine is not errored or halted"
        );
        require(!vm.inChallenge, "Can only disputable assert if not in challenge");
        require(vm.escrowRequired <= vm.validatorBalances[msg.sender], "Validator does not have required escrow");
        require(_data.numSteps <= vm.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(_data.timeBounds), "Precondition: not within time bounds");
        require(_data.beforeHash == vm.machineHash, "Precondition: state hash does not match");
        require(
            _data.beforeInbox == vm.inboxHash ||
            _data.beforeInbox == ArbProtocol.appendInboxMessages(vm.inboxHash, vm.pendingMessages),
            "Precondition: inbox does not match"
        );

        uint256[] memory beforeBalances = ArbProtocol.calculateBeforeValues(
            _data.tokenTypes,
            _data.messageTokenNum,
            _data.msgAmount
        );

        require(
            arbBalanceTracker.hasFunds(
                _data.vmId,
                _data.tokenTypes,
                beforeBalances
            ),
            "VM has insufficient balance"
        );
        _resetDeadline(vm);

        bytes32 lastMessageHash = ArbProtocol.generateLastMessageHashStub(
            _data.tokenTypes,
            _data.messageDataHash,
            _data.messageTokenNum,
            _data.msgAmount,
            _data.msgDestination
        );

        vm.pendingHash = keccak256(
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
        vm.validatorBalances[msg.sender] = vm.validatorBalances[msg.sender].sub(vm.escrowRequired);
        vm.asserter = msg.sender;
        vm.state = VMState.PendingAssertion;

        emit PendingDisputableAssertion(
            _data.vmId,
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

    function _confirmDisputableAsserted(ConfirmDisputableAssertedData memory _data) internal {
        VM storage vm = vms[_data.vmId];
        require(vm.state == VMState.PendingAssertion, "VM does not have assertion pending");
        require(block.number > vm.deadline, "Assertion is still pending challenge");
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
                            _data.messageData,
                            _data.messageTokenNums,
                            _data.messageAmounts,
                            _data.messageDestination
                        ),
                        0x00,
                        _data.logsAccHash,
                        ArbProtocol.calculateBeforeValues(
                            _data.tokenTypes,
                            _data.messageTokenNums,
                            _data.messageAmounts
                        )
                    )
                )
            ) == vm.pendingHash,
            "Precondition and assertion do not match pending assertion"
        );
        vm.validatorBalances[vm.asserter] = vm.validatorBalances[vm.asserter].add(vm.escrowRequired);
        _acceptAssertion(
            _data.vmId,
            _data.afterHash,
            _data.tokenTypes,
            _data.messageData,
            _data.messageTokenNums,
            _data.messageAmounts,
            _data.messageDestination
        );

        emit ConfirmedDisputableAssertion(
            _data.vmId,
            _data.afterHash,
            _data.logsAccHash
        );
    }

    function _resetDeadline(VM storage _vm) internal {
        _vm.deadline = uint64(block.number) + _vm.gracePeriod;
    }

    function _acceptAssertion(
        bytes32 _vmId,
        bytes32 _afterHash,
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
        VM storage vm = vms[_vmId];
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

        vm.machineHash = _afterHash;
        vm.state = VMState.Waiting;

        if (vm.pendingMessages != ArbValue.hashEmptyTuple()) {
            vm.inboxHash = ArbProtocol.appendInboxMessages(vm.inboxHash, vm.pendingMessages);
            vm.pendingMessages = ArbValue.hashEmptyTuple();
        }

        if (_afterHash == MACHINE_HALT_HASH) {
            _shutdownVM(_vmId);
        }
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
            VM storage vm = vms[_destination];
            bytes32 messageHash = ArbProtocol.generateSentMessageHash(
                _destination,
                ArbValue.deserializeValueHash(_data),
                _tokenType,
                _value,
                _sender
            );
            vm.pendingMessages = ArbProtocol.appendInboxPendingMessage(
                vm.pendingMessages,
                messageHash
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

    function _cancelCurrentState(VM storage vm) internal {
        if (vm.state != VMState.Waiting) {
            require(block.number <= vm.deadline, "Can't cancel finalized state");
        }

        if (vm.state == VMState.PendingAssertion) {
            // If there is a pending disputable assertion, cancel it
            vm.validatorBalances[vm.asserter] = vm.validatorBalances[vm.asserter].add(vm.escrowRequired);
        }
    }

    function _shutdownVM(bytes32 _vmId) private {
        // TODO: transfer all owned funds to halt address
        delete vms[_vmId];
    }
}
