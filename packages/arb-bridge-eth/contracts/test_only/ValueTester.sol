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

import "../arch/Value.sol";

contract ValueTester {
    function deserializeHash(
        bytes memory data,
        uint256 startOffset
    )
        public
        pure
        returns(
            bool, // valid
            uint256, // offset
            bytes32 // valHash
        )
    {
        (bool valid, uint256 offset, Value.Data memory value) = Value.deserialize(data, startOffset);
        return (valid, offset, Value.hash(value));
    }

    function deserializeMessageData(
        bytes memory data,
        uint256 startOffset
    )
        public
        pure
        returns(
            bool, // valid
            uint256, // offset
            uint256, // msgType
            address // sender
        )
    {
        return Value.deserializeMessageData(data, startOffset);
    }

    function getERCTokenMsgData(
        bytes memory data,
        uint256 startOffset
    )
        public
        pure
        returns(
            bool, // valid
            uint256, // offset
            address, // tokenAddress
            address, // destination
            uint256 // value
        )
    {
        return Value.getERCTokenMsgData(data, startOffset);
    }

    function getEthMsgData(
        bytes memory data,
        uint256 startOffset
    )
        public
        pure
        returns(
            bool, // valid
            uint256, // offset
            address, // destination
            uint256 // value
        )
    {
        return Value.getEthMsgData(data, startOffset);
    }

    function bytesToBytestackHash(bytes memory data)
        public
        pure
        returns(bytes32){
        return Value.hash(Value.bytesToBytestackHash(data, 0, data.length));
    }


    function bytesToBytestackHash(
        bytes memory data,
        uint256 startOffset,
        uint256 dataLength
    )
        public
        pure
        returns(bytes32)
    {
        return Value.hash(Value.bytesToBytestackHash(data, startOffset, dataLength));
    }

    function bytestackToBytes(bytes memory data) public pure returns (bytes memory) {
        return Value.bytestackToBytes(data);
    }

    function hashTuplePreImage(
        bytes32 innerHash,
        uint256 valueSize
    )
        public pure returns (bytes32)
    {
        return Value.hashTuplePreImage(innerHash, valueSize);
    }

    function hashEmptyTuple() public pure returns (bytes32)
    {
        return Value.hashEmptyTuple();
    }

    function hashTestTuple() public pure returns (bytes32)
    {
        Value.Data[] memory tupVals = new Value.Data[](2);
        tupVals[0] = Value.newInt(uint256(111));
        tupVals[1] = Value.newTuple(new Value.Data[](0));
        Value.Data memory tuple = Value.newTuple(tupVals);

        return Value.hash(tuple);
    }

}
