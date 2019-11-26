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

    function generateMessageStubHash(
        bytes32 _data,
        bytes21 _tokenType,
        uint256 _value,
        address _destination
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

    function generatePreconditionHash(
        bytes32 _beforeHash,
        uint64[2] memory _timeBounds,
        bytes32 _beforeInbox
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
                _beforeInbox
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
        bytes21[] memory _tokenTypes,
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
                _tokenTypes,
                _totalMessageValueAmounts
            )
        );
    }

    function beforeBalancesValid(
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances
    )
        public
        pure
        returns(bool)
    {
        uint itemCount = _tokenTypes.length;
        if (itemCount == 0 || itemCount == 1) {
            return true;
        }
        for (uint i = 0; i < itemCount - 1; i++) {
            byte tokenType = _tokenTypes[i][20];
            if (tokenType == 0x00) {
                if (_tokenTypes[i + 1] <= _tokenTypes[i]) {
                    return false;
                }
            } else if (tokenType == 0x01) {
                if (
                    _tokenTypes[i + 1] < _tokenTypes[i] || (
                        _tokenTypes[i + 1] == _tokenTypes[i] &&
                        _beforeBalances[i + 1] <= _beforeBalances[i]
                    )
                ) {
                    return false;
                }

            } else {
                return false;
            }
        }

        if (_tokenTypes[itemCount - 1][20] > 0x01) {
            return false;
        }
        return true;
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
        uint messageCount = _messageTokenNums.length;
        uint256[] memory beforeBalances = new uint256[](_tokenTypes.length);

        for (uint i = 0; i < messageCount; i++) {
            uint16 tokenNum = _messageTokenNums[i];
            if (_tokenTypes[tokenNum][20] == 0x00) {
                beforeBalances[tokenNum] += _messageAmounts[i];
            } else {
                require(beforeBalances[tokenNum] == 0, "Can't include NFT token twice");
                require(_messageAmounts[i] != 0, "NFT token must have non-zero id");
                beforeBalances[tokenNum] = _messageAmounts[i];
            }
        }
        return beforeBalances;
    }
}
