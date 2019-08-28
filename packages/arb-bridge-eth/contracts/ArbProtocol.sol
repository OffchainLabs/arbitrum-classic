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

import "./ArbValue.sol";


library ArbProtocol {
    using ArbValue for ArbValue.Value;

    function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) public pure returns (bytes32) {
        return ArbValue.hashTupleValue([
            ArbValue.newIntValue(1),
            ArbValue.newHashOnlyValue(_inboxHash),
            ArbValue.newHashOnlyValue(_pendingMessages)
        ]);
    }

    function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) public pure returns (bytes32) {
        return ArbValue.hashTupleValue([
            ArbValue.newIntValue(0),
            ArbValue.newHashOnlyValue(_pendingMessages),
            ArbValue.newHashOnlyValue(_newMessage)
        ]);
    }

    function generateMessageStubHash(
        bytes32 _data,
        bytes21 _tokenType,
        uint256 _value,
        bytes32 _destination
    )
        public
        pure
        returns (bytes32)
    {
        ArbValue.Value[] memory values = new ArbValue.Value[](4);
        values[0] = ArbValue.newHashOnlyValue(_data);
        values[1] = ArbValue.newIntValue(uint256(_destination));
        values[2] = ArbValue.newIntValue(_value);
        values[3] = ArbValue.newIntValue(uint256(bytes32(_tokenType)));
        return ArbValue.newTupleValue(values).hash().hash;
    }

    function generateSentMessageHash(
        bytes32 _dest,
        bytes32 _data,
        bytes21 _tokenType,
        uint256 _value,
        bytes32 _sender
    )
        public
        view
        returns (bytes32)
    {
        bytes32 txHash = keccak256(
            abi.encodePacked(
                _dest,
                _data,
                _value,
                _tokenType
            )
        );
        ArbValue.Value[] memory dataValues = new ArbValue.Value[](4);
        dataValues[0] = ArbValue.newHashOnlyValue(_data);
        dataValues[1] = ArbValue.newIntValue(block.timestamp);
        dataValues[2] = ArbValue.newIntValue(block.number);
        dataValues[3] = ArbValue.newIntValue(uint(txHash));

        ArbValue.Value[] memory values = new ArbValue.Value[](4);
        values[0] = ArbValue.newTupleValue(dataValues);
        values[1] = ArbValue.newIntValue(uint256(_sender));
        values[2] = ArbValue.newIntValue(_value);
        values[3] = ArbValue.newIntValue(uint256(bytes32(_tokenType)));
        return ArbValue.newTupleValue(values).hash().hash;
    }

    function generatePreconditionHash(
        bytes32 _beforeHash,
        uint64[2] memory _timeBounds,
        bytes32 _beforeInbox,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances
    )
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _beforeHash,
                _timeBounds[0],
                _timeBounds[1],
                _beforeInbox,
                _tokenTypes,
                _beforeBalances
            )
        );
    }

    function generateAssertionHash(
        bytes32 _afterHash,
        uint32 _numSteps,
        bytes32 _firstMessageHash,
        bytes32 _lastMessageHash,
        bytes32 _firstLogHash,
        bytes32 _lastLogHash,
        uint256[] memory _totalMessageValueAmounts
    )
        public
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _afterHash,
                _numSteps,
                _firstMessageHash,
                _lastMessageHash,
                _firstLogHash,
                _lastLogHash,
                _totalMessageValueAmounts
            )
        );
    }

    // fields:
    // vmId
    // beforeHash
    // beforeInbox
    // afterHash
    // newInbox

    function unanimousAssertHash(
        bytes32[5] memory _fields,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        bytes memory _messageData,
        uint16[] memory _messageTokenNum,
        uint256[] memory _messageAmount,
        bytes32[] memory _messageDestination
    )
        public
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _fields,
                _timeBounds,
                _tokenTypes,
                _messageData,
                _messageTokenNum,
                _messageAmount,
                _messageDestination
            )
        );
    }

    function calculateBeforeValues(
        bytes21[] memory _tokenTypes,
        uint16[] memory _messageTokenNums,
        uint256[] memory _messageAmounts
    )
        public
        pure
        returns(uint256[] memory)
    {
        uint256[] memory beforeBalances = new uint256[](_tokenTypes.length);
        uint tokenNumCount = _messageTokenNums.length;
        for (uint i = 0; i < tokenNumCount; i++) {
            if (_tokenTypes[_messageTokenNums[i]][20] == 0) {
                beforeBalances[_messageTokenNums[i]] += _messageAmounts[i];
            } else {
                require(beforeBalances[_messageTokenNums[i]] == 0, "Can't include NFT token twice");
                require(_messageAmounts[i] != 0, "NFT token must have non-zero id");
                beforeBalances[_messageTokenNums[i]] = _messageAmounts[i];
            }
        }
        return beforeBalances;
    }

    function generateLastMessageHash(
        bytes21[] memory _tokenTypes,
        bytes memory _data,
        uint16[] memory _tokenNums,
        uint256[] memory _amounts,
        bytes32[] memory _destinations
    )
        public
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
            msgHash = generateMessageStubHash(
                msgHash,
                _tokenTypes[_tokenNums[i]],
                _amounts[i],
                _destinations[i]
            );
            hashVal = keccak256(abi.encodePacked(hashVal, msgHash));
        }
    }

    function generateLastMessageHashStub(
        bytes21[] memory _tokenTypes,
        bytes32[] memory _dataHashes,
        uint16[] memory _tokenNums,
        uint256[] memory _amounts,
        bytes32[] memory _destinations
    )
        public
        pure
        returns (bytes32)
    {
        require(_dataHashes.length == _tokenNums.length, "Input size mismatch");
        require(_dataHashes.length == _amounts.length, "Input size mismatch");
        require(_dataHashes.length == _destinations.length, "Input size mismatch");
        bytes32 hashVal = 0x00;
        bytes32 msgHash;
        uint dataHashCount = _dataHashes.length;
        for (uint i = 0; i < dataHashCount; i++) {
            msgHash = generateMessageStubHash(
                _dataHashes[i],
                _tokenTypes[_tokenNums[i]],
                _amounts[i],
                _destinations[i]
            );
            hashVal = keccak256(abi.encodePacked(hashVal, msgHash));
        }
        return hashVal;
    }

    function parseSignature(
        bytes memory _signatures,
        uint _pos
    )
        public
        pure
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        uint offset = _pos * 65;
        // The signature format is a compact form of:
        //   {bytes32 r}{bytes32 s}{uint8 v}
        // Compact means, uint8 is not padded to 32 bytes.
        assembly { // solium-disable-line security/no-inline-assembly
            r := mload(add(_signatures, add(32, offset)))
            s := mload(add(_signatures, add(64, offset)))
            // Here we are loading the last 32 bytes, including 31 bytes
            // of 's'. There is no 'mload8' to do this.
            //
            // 'byte' is not working due to the Solidity parser, so lets
            // use the second best option, 'and'
            v := and(mload(add(_signatures, add(65, offset))), 0xff)
        }

        if (v < 27) {
            v += 27;
        }

        require(v == 27 || v == 28, "Incorrect v value");
    }

    /// @notice Counts the number of signatures in a signatures bytes array. Returns 0 if the length is invalid.
    /// @param _signatures The signatures bytes array
    /// @dev Signatures are 65 bytes long and are densely packed.
    function countSignatures(bytes memory _signatures) public pure returns (uint) {
        return _signatures.length % 65 == 0 ? _signatures.length / 65 : 0;
    }

    /// @notice Recovers an array of addresses using a message hash and a signatures bytes array.
    /// @param _messageHash The signed message hash
    /// @param _signatures The signatures bytes array
    function recoverAddresses(
        bytes32 _messageHash,
        bytes memory _signatures
    )
        public
        pure
        returns (address[] memory)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;
        uint count = countSignatures(_signatures);
        address[] memory addresses = new address[](count);
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, _messageHash));
        for (uint i = 0; i < count; i++) {
            (v, r, s) = parseSignature(_signatures, i);
            addresses[i] = ecrecover(
                prefixedHash,
                v,
                r,
                s
            );
        }
        return addresses;
    }

    /// @notice Recovers an array of addresses using a message hash and a signatures bytes array.
    /// @param _messageHash The signed message hash
    /// @param _signature The signature bytes array
    function recoverAddress(
        bytes32 _messageHash,
        bytes memory _signature
    )
        public
        pure
        returns (address)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, _messageHash));
        (v, r, s) = parseSignature(_signature, 0);
        return ecrecover(
            prefixedHash,
            v,
            r,
            s
        );
    }
}
