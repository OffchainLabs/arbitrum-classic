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

pragma solidity ^0.6.11;

import "../arch/Value.sol";
import "../arch/Marshaling.sol";

contract ValueTester {
    using Hashing for Value.Data;

    function deserializeHash(bytes memory data, uint256 startOffset)
        public
        pure
        returns (
            uint256, // offset
            bytes32 // valHash
        )
    {
        (uint256 offset, Value.Data memory value) = Marshaling.deserialize(data, startOffset);
        return (offset, value.hash());
    }

    function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) public pure returns (bytes32) {
        return Hashing.hashTuplePreImage(innerHash, valueSize);
    }

    function hashTestTuple() public pure returns (bytes32) {
        Value.Data[] memory tupVals = new Value.Data[](2);
        tupVals[0] = Value.newInt(uint256(111));
        tupVals[1] = Value.newTuple(new Value.Data[](0));
        return Value.newTuple(tupVals).hash();
    }
}
