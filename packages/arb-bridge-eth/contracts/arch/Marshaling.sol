// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.11;

import "./Value.sol";
import "./Hashing.sol";

import "../libraries/BytesLib.sol";

library Marshaling {
    using BytesLib for bytes;
    using Value for Value.Data;

    function deserializeHashPreImage(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            uint256, // offset
            Value.Data memory
        )
    {
        uint256 size;
        require(
            data.length >= startOffset && data.length - startOffset >= 64,
            "to short"
        );
        bytes32 hashData = data.toBytes32(startOffset);
        startOffset += 32;
        (startOffset, size) = deserializeInt(data, startOffset);
        return (startOffset, Value.newTuplePreImage(hashData, size));
    }

    function deserializeInt(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            uint256, // offset
            uint256 // val
        )
    {
        require(
            data.length >= startOffset && data.length - startOffset >= 32,
            "too short"
        );
        return (startOffset + 32, data.toUint(startOffset));
    }

    function deserializeCheckedInt(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            uint256 // val
        )
    {
        uint256 totalLength = data.length;
        if (
            totalLength < startOffset ||
            totalLength - startOffset < 33 ||
            uint8(data[startOffset]) != Value.intTypeCode()
        ) {
            return (false, startOffset, 0);
        }
        return (true, startOffset + 33, data.toUint(startOffset + 1));
    }

    function deserializeCodePoint(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            uint256, // offset
            Value.Data memory // val
        )
    {
        uint256 offset = startOffset;
        uint8 immediateType = uint8(data[offset]);
        offset++;
        uint8 opCode = uint8(data[offset]);
        offset++;
        Value.Data memory immediate;
        if (immediateType == 1) {
            (offset, immediate) = deserialize(data, offset);
        }
        bytes32 nextHash = data.toBytes32(offset);
        offset += 32;
        if (immediateType == 1) {
            return (offset, Value.newCodePoint(opCode, nextHash, immediate));
        }
        return (offset, Value.newCodePoint(opCode, nextHash));
    }

    function deserializeTuple(
        uint8 memberCount,
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns (
            uint256, // offset
            Value.Data[] memory // val
        )
    {
        uint256 offset = startOffset;
        Value.Data[] memory members = new Value.Data[](memberCount);
        for (uint8 i = 0; i < memberCount; i++) {
            (offset, members[i]) = deserialize(data, offset);
        }
        return (offset, members);
    }

    function deserialize(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            uint256, // offset
            Value.Data memory // val
        )
    {
        require(startOffset < data.length, "invalid offset");
        uint256 offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        if (valType == Value.intTypeCode()) {
            (offset, intVal) = deserializeInt(data, offset);
            return (offset, Value.newInt(intVal));
        } else if (valType == Value.codePointTypeCode()) {
            return deserializeCodePoint(data, offset);
        } else if (valType == Value.tuplePreImageTypeCode()) {
            return deserializeHashPreImage(data, offset);
        } else if (
            valType >= Value.tupleTypeCode() && valType < Value.valueTypeCode()
        ) {
            uint8 tupLength = uint8(valType - Value.tupleTypeCode());
            Value.Data[] memory tupleVal;
            (offset, tupleVal) = deserializeTuple(tupLength, data, offset);
            return (offset, Value.newTuple(tupleVal));
        }
        require(false, "invalid typecode");
    }

    function bytesToBytestack(
        bytes memory data,
        uint256 startOffset,
        uint256 dataLength
    ) internal pure returns (Value.Data memory) {
        uint256 wholeChunkCount = dataLength / 32;

        // tuple code + size + (for each chunk tuple code + chunk val) + empty tuple code
        Value.Data memory stack = Value.newNone();
        Value.Data[] memory vals = new Value.Data[](2);

        for (uint256 i = 0; i < wholeChunkCount; i++) {
            vals[0] = Value.newInt(data.toUint(startOffset + i * 32));
            vals[1] = stack;
            stack = Hashing.getTuplePreImage(vals);
        }

        if (dataLength % 32 != 0) {
            uint256 lastVal = data.toUint(startOffset + dataLength - 32);
            lastVal <<= (32 - (dataLength % 32)) * 8;
            vals[0] = Value.newInt(lastVal);
            vals[1] = stack;
            stack = Hashing.getTuplePreImage(vals);
        }

        vals[0] = Value.newInt(dataLength);
        vals[1] = stack;

        return Hashing.getTuplePreImage(vals);
    }

    function bytestackToBytes(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool valid,
            uint256 offset,
            bytes memory byteData
        )
    {
        offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        if (valType != Value.tupleTypeCode() + 2) {
            return (false, offset, byteData);
        }

        uint256 byteCount;
        (valid, offset, byteCount) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, offset, byteData);
        }
        uint256 fullChunkCount = byteCount / 32;
        uint256 partialChunkSize = byteCount % 32;
        uint256 totalChunkCount = fullChunkCount +
            (partialChunkSize > 0 ? 1 : 0);

        bytes32[] memory fullChunks = new bytes32[](fullChunkCount);
        bytes memory partialChunk = new bytes(partialChunkSize);

        uint256 fullChunkIndex = 0;

        for (uint256 i = 0; i < totalChunkCount; i++) {
            valType = uint8(data[offset]);
            offset++;
            if (valType != Value.tupleTypeCode() + 2) {
                return (false, offset, byteData);
            }

            uint256 nextChunk;
            (valid, offset, nextChunk) = deserializeCheckedInt(data, offset);
            if (!valid) {
                return (false, offset, byteData);
            }

            if (i == 0 && partialChunkSize > 0) {
                bytes32 chunkBytes = bytes32(nextChunk);
                for (uint256 j = 0; j < partialChunkSize; j++) {
                    partialChunk[j] = chunkBytes[j];
                }
            } else {
                fullChunks[fullChunkCount - 1 - fullChunkIndex] = bytes32(
                    nextChunk
                );
                fullChunkIndex++;
            }
        }

        valType = uint8(data[offset]);
        offset++;
        if (valType != Value.tupleTypeCode()) {
            return (false, offset, byteData);
        }
        return (true, offset, abi.encodePacked(fullChunks, partialChunk));
    }
}
