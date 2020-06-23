// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "../arch/Protocol.sol";

contract ProtocolTester{

    function generateAssertionHash(
    	bytes32 _afterHash,
        bool    _didInboxInsn,
        uint64  _numGas,
        bytes32 _firstMessageHash,
        bytes32 _lastMessageHash,
        bytes32 _firstLogHash,
        bytes32 _lastLogHash
    )
    	public
        pure
        returns (bytes32)
    {
       return Protocol.generateAssertionHash(
       	_afterHash,
        _didInboxInsn,
        _numGas,
        _firstMessageHash,
        _lastMessageHash,
        _firstLogHash,
        _lastLogHash);

    }

    function generatePreconditionHash(
        bytes32 _beforeHash,
        uint128[4] memory _timeBounds,
        bytes32 _beforeInboxHash
    )
        public
        pure
        returns(bytes32)
    {
        return Protocol.generatePreconditionHash(
        	_beforeHash, 
        	_timeBounds, 
        	_beforeInboxHash);
    }
}