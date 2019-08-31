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
import "./MerkleLib.sol";
import "./ArbProtocol.sol";


library Unanimous {

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

    struct PendingUnanimousAssertData {
        bytes32 unanRest;
        bytes21[] tokenTypes;
        uint16[] messageTokenNums;
        uint256[] messageAmounts;
        uint64 sequenceNum;
        bytes32 logsAccHash;
        bytes signatures;
    }

    struct FinalizedUnanimousAssertData {
        bytes32 afterHash;
        bytes32 newInbox;
        bytes21[] tokenTypes;
        VM.FullAssertion assertion;
        bytes signatures;
    }

    // fields
    //   _afterHash
    //   _newInbox
    //   _logsAccHash

    function finalizedUnanimousAssert(
        VM.Data storage _vm,
        bytes32[3] memory _fields,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts,
        bytes32[] memory _messageDestinations,
        bytes memory _signatures
    )
        public
    {
        _finalizedUnanimousAssert(
            _vm,
            FinalizedUnanimousAssertData(
                _fields[0],
                _fields[1],
                _tokenTypes,
                VM.FullAssertion(
                    _messageData,
                    _messageTokenNums,
                    _messageAmounts,
                    _messageDestinations,
                    _fields[2]
                ),
                _signatures
            )
        );
    }

    function pendingUnanimousAssert(
        VM.Data storage _vm,
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
        return _pendingUnanimousAssert(
            _vm,
            PendingUnanimousAssertData(
                _unanRest,
                _tokenTypes,
                _messageTokenNums,
                _messageAmounts,
                _sequenceNum,
                _logsAccHash,
                _signatures
            )
        );
    }

    function confirmUnanimousAsserted(
        VM.Data storage _vm,
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
        require(_vm.state == VM.State.PendingUnanimous, "Can only confirm if there is a pending assertion");
        require(!VM.withinDeadline(_vm), "Can only confirm assertion whose challenge deadline has passed");
        require(
            keccak256(
                abi.encodePacked(
                    _tokenTypes,
                    _messageTokenNums,
                    _messageAmounts,
                    keccak256(
                        abi.encodePacked(
                            _newInbox,
                            _afterHash,
                            _messageData,
                            _messageDestinations
                        )
                    )
                )
            ) == _vm.pendingHash,
            "Can only confirm assertion that is currently pending"
        );

        _vm.inboxHash = _newInbox;
        VM.acceptAssertion(
            _vm,
            _afterHash
        );

        emit ConfirmedUnanimousAssertion(
            _vm.id,
            _vm.sequenceNum
        );
    }

    function _finalizedUnanimousAssert(
        VM.Data storage vm,
        FinalizedUnanimousAssertData memory data
    )
        private
    {
        require(!VM.isHalted(vm), "Can't assert halted machine");
        bytes32 unanHash = keccak256(
            abi.encodePacked(
                vm.id,
                keccak256(
                    abi.encodePacked(
                        keccak256(
                            abi.encodePacked(
                                data.newInbox,
                                data.afterHash,
                                data.assertion.messageData,
                                data.assertion.messageDestinations
                            )
                        ),
                        vm.machineHash,
                        vm.inboxHash,
                        data.tokenTypes,
                        data.assertion.messageTokenNums,
                        data.assertion.messageAmounts
                    )
                ),
                data.assertion.logsAccHash
            )
        );
        require(
            MerkleLib.generateAddressRoot(
                ArbProtocol.recoverAddresses(unanHash, data.signatures)
            ) == vm.validatorRoot,
            "Validator signatures don't match"
        );

        VM.cancelCurrentState(vm);
        vm.state = VM.State.Waiting;
        vm.inboxHash = data.newInbox;
        VM.acceptAssertion(
            vm,
            data.afterHash
        );

        emit FinalizedUnanimousAssertion(
            vm.id,
            unanHash
        );
    }

    function _pendingUnanimousAssert(
        VM.Data storage vm,
        PendingUnanimousAssertData memory data
    )
        private
    {
        require(!VM.isHalted(vm), "Can't assert halted machine");
        bytes32 unanHash = keccak256(
            abi.encodePacked(
                vm.id,
                keccak256(
                    abi.encodePacked(
                        data.unanRest,
                        vm.machineHash,
                        vm.inboxHash,
                        data.tokenTypes,
                        data.messageTokenNums,
                        data.messageAmounts,
                        data.sequenceNum
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

        if (vm.state == VM.State.PendingUnanimous) {
            require(
                data.sequenceNum > vm.sequenceNum,
                "Can only supersede previous assertion with greater sequence number"
            );
        }

        VM.cancelCurrentState(vm);
        VM.resetDeadline(vm);

        vm.state = VM.State.PendingUnanimous;
        vm.sequenceNum = data.sequenceNum;
        vm.pendingHash = keccak256(
            abi.encodePacked(
                data.tokenTypes,
                data.messageTokenNums,
                data.messageAmounts,
                data.unanRest
            )
        );

        emit PendingUnanimousAssertion(
            vm.id,
            unanHash,
            data.sequenceNum
        );
    }
}
