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
import "./IArbChannel.sol";

import "../libraries/Protocol.sol";
import "../libraries/SigUtils.sol";


library Unanimous {

    event PendingUnanimousAssertion(
        bytes32 unanHash,
        uint64 sequenceNum,
        uint64 deadline
    );

    event ConfirmedUnanimousAssertion(
        uint64 sequenceNum
    );

    event FinalizedUnanimousAssertion(
        bytes32 unanHash
    );

    function finalizedUnanimousAssert(
        VM.Data storage vm,
        IArbChannel channel,
        bytes32 afterHash,
        bytes32 newInbox,
        bytes memory messages,
        bytes32 logsAccHash,
        bytes memory signatures
    )
        public
    {
        require(!VM.isHalted(vm), "Can't assert halted machine");
        require(
            vm.state == VM.State.Waiting ||
            vm.state == VM.State.PendingDisputable ||
            vm.state == VM.State.PendingUnanimous,
            "Tried to finalize unanimous from invalid state"
        );
        if (vm.state != VM.State.Waiting) {
            require(block.number <= vm.deadline, "Can't cancel finalized state");
        }
        bool allSigned;
        bytes32 unanHash;
        (allSigned, unanHash) = _checkAllSignedAssertion(
            vm,
            channel,
            keccak256(
                abi.encodePacked(
                    newInbox,
                    afterHash
                )
            ),
            ~uint64(0),
            Protocol.generateLastMessageHash(messages),
            logsAccHash,
            signatures
        );

        require(allSigned, "Invalid signature list");

        vm.inbox = newInbox;

        emit FinalizedUnanimousAssertion(
            unanHash
        );
    }

    function pendingUnanimousAssert(
        VM.Data storage vm,
        IArbChannel channel,
        bytes32 unanRest,
        uint64 sequenceNum,
        bytes32 messagesAccHash,
        bytes32 logsAccHash,
        bytes memory signatures
    )
        public
    {
        require(!VM.isHalted(vm), "Can't assert halted machine");
        require(
            vm.state == VM.State.Waiting ||
            vm.state == VM.State.PendingDisputable ||
            vm.state == VM.State.PendingUnanimous,
            "Tried to pending unanimous from invalid state"
        );
        if (vm.state != VM.State.Waiting) {
            require(block.number <= vm.deadline, "Can't cancel finalized state");
        }

        bool allSigned;
        bytes32 unanHash;
        (allSigned, unanHash) = _checkAllSignedAssertion(
            vm,
            channel,
            unanRest,
            sequenceNum,
            messagesAccHash,
            logsAccHash,
            signatures
        );

        require(allSigned, "Invalid signature list");

        if (vm.state == VM.State.PendingUnanimous) {
            require(
                sequenceNum > vm.sequenceNum,
                "Can only supersede previous assertion with greater sequence number"
            );
        }

        VM.resetDeadline(vm);
        vm.sequenceNum = sequenceNum;
        vm.pendingHash = keccak256(
            abi.encodePacked(
                messagesAccHash,
                unanRest
            )
        );

        emit PendingUnanimousAssertion(
            unanHash,
            sequenceNum,
            vm.deadline
        );
    }

    function confirmUnanimousAsserted(
        VM.Data storage _vm,
        bytes32 _afterHash,
        bytes32 _newInbox,
        bytes memory _messages
    )
        public
    {
        require(_vm.state == VM.State.PendingUnanimous, "Can only confirm if there is a pending unanimous assertion");
        require(!VM.withinDeadline(_vm), "Can only confirm assertion whose challenge deadline has passed");
        require(
            keccak256(
                abi.encodePacked(
                    Protocol.generateLastMessageHash(_messages),
                    keccak256(
                        abi.encodePacked(
                            _newInbox,
                            _afterHash
                        )
                    )
                )
            ) == _vm.pendingHash,
            "Can only confirm assertion that is currently pending"
        );

        _vm.inbox = _newInbox;
        VM.acceptAssertion(
            _vm,
            _afterHash
        );

        emit ConfirmedUnanimousAssertion(
            _vm.sequenceNum
        );
    }

    function _checkAllSignedAssertion(
        VM.Data storage vm,
        IArbChannel channel,
        bytes32 _unanRest,
        uint64 _sequenceNum,
        bytes32 _messagesAccHash,
        bytes32 _logsAccHash,
        bytes memory _signatures
    )
        private
        view
        returns(bool, bytes32)
    {
        bytes32 partialHash = keccak256(
            abi.encodePacked(
                _unanRest,
                vm.machineHash,
                vm.inbox,
                _sequenceNum,
                _messagesAccHash

            )
        );
        bytes32 unanHash = keccak256(
            abi.encodePacked(
                address(this),
                partialHash,
                _logsAccHash
            )
        );
        bool allSigned = channel.isValidatorList(SigUtils.recoverAddresses(unanHash, _signatures));
        return (allSigned, unanHash);
    }
}
