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

import "./Value.sol";


library Protocol {
    using Value for Value.Data;

    function generateMessageStubHash(
        Value.Data memory _dataHashValue,
        bytes21 _tokenType,
        uint256 _value,
        address _destination
    )
        internal
        pure
        returns (bytes32)
    {
        Value.Data[] memory values = new Value.Data[](4);
        values[0] = _dataHashValue;
        values[1] = Value.newInt(uint256(_destination));
        values[2] = Value.newInt(_value);
        values[3] = Value.newInt(uint256(bytes32(_tokenType)));
        return Value.newTuple(values).hash().hash;
    }

    function generatePreconditionHash(
        bytes32 _beforeHash,
        uint128[2] memory _timeBounds,
        bytes32 _beforeInbox
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _beforeHash,
                _timeBounds[0],
                _timeBounds[1],
                _beforeInbox
            )
        );
    }

    function generateAssertionHash(
        bytes32 _afterHash,
        bool    _didInboxInsn,
        uint64  _numGas,
        bytes32 _firstMessageHash,
        bytes32 _lastMessageHash,
        bytes32 _firstLogHash,
        bytes32 _lastLogHash
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _afterHash,
                _didInboxInsn,
                _numGas,
                _firstMessageHash,
                _lastMessageHash,
                _firstLogHash,
                _lastLogHash
            )
        );
    }

    function generateLastMessageHash(bytes memory messages, uint256 startOffset, uint256 length) internal pure returns (bytes32) {
        bool valid;
        bytes32 hashVal = 0x00;
        bytes32 msgHash;
        uint256 endOffset = startOffset + length;
        require(endOffset <= messages.length, "invalid length");
        uint256 offset;
        for (offset = startOffset; offset < endOffset;) {
            (valid, offset, msgHash) = Value.deserializeHashed(messages, offset);
            require(valid, "Invalid output message");
            hashVal = keccak256(abi.encodePacked(hashVal, msgHash));
        }
        require(offset == startOffset + length, "value extended past length");
        return hashVal;
    }

    function addMessageToVMInboxHash(Value.Data memory vmInboxHashValue, Value.Data memory messageHashValue) internal pure returns (bytes32) {
        Value.Data[] memory vals = new Value.Data[](2);
        vals[0] = vmInboxHashValue;
        vals[1] = messageHashValue;
        Value.Data memory tuple = Value.newTuple(vals);

        return Value.hashTuple(tuple);
    }

    function addMessageToInbox(bytes32 inbox, bytes32 message) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                inbox,
                message
            )
        );
    }
}
