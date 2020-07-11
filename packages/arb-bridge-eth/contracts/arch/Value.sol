// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

import "../libraries/BytesLib.sol";

library Value {
    using BytesLib for bytes;
    using Value for Data;

    uint8 internal constant INT_TYPECODE = 0;
    uint8 internal constant CODE_POINT_TYPECODE = 1;
    uint8 internal constant HASH_PRE_IMAGE_TYPECODE = 2;
    uint8 internal constant TUPLE_TYPECODE = 3;
    // All values received from clients will have type codes less than the VALUE_TYPE_COUNT
    uint8 internal constant VALUE_TYPE_COUNT = TUPLE_TYPECODE + 9;

    // The following types do not show up in the marshalled format and is
    // only used for internal tracking purposes
    uint8 internal constant HASH_ONLY = 100;

    struct CodePoint {
        uint8 opcode;
        bytes32 nextCodePoint;
        bool immediate;
        bytes32 immediateHash;
        uint256 immediateSize;
    }

    struct Data {
        uint256 intVal;
        CodePoint cpVal;
        Data[] tupleVal;
        uint8 typeCode;
        uint256 size;
    }

    function tupleTypeCode() internal pure returns (uint8) {
        return TUPLE_TYPECODE;
    }

    function isTupleType(uint8 typeCode) private pure returns (bool) {
        return typeCode < VALUE_TYPE_COUNT && typeCode >= TUPLE_TYPECODE;
    }

    function hashInt(uint256 val) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(val));
    }

    function hashCodePoint(
        uint8 opcode,
        bool immediate,
        bytes32 immediateHash,
        bytes32 nextCodePoint
    ) internal pure returns (bytes32) {
        if (immediate) {
            return
                keccak256(
                    abi.encodePacked(
                        CODE_POINT_TYPECODE,
                        opcode,
                        immediateHash,
                        nextCodePoint
                    )
                );
        }
        return
            keccak256(
                abi.encodePacked(CODE_POINT_TYPECODE, opcode, nextCodePoint)
            );
    }

    function hashTuplePreImage(bytes32 innerHash, uint256 valueSize)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked(uint8(TUPLE_TYPECODE), innerHash, valueSize)
            );
    }

    function hashTupleInner(Data[] memory vals)
        private
        pure
        returns (bytes32[] memory)
    {
        require(vals.length <= 8, "Invalid tuple length");

        bytes32[] memory hashes = new bytes32[](vals.length);

        uint256 hashCount = hashes.length;
        for (uint256 i = 0; i < hashCount; i++) {
            bytes32 hashVal = vals[i].hash();
            hashes[i] = hashVal;
        }

        return hashes;
    }

    function hashTuplePreImage(Data memory preImage)
        internal
        pure
        returns (bytes32)
    {
        require(
            preImage.typeCode == HASH_PRE_IMAGE_TYPECODE,
            "Must be PreImageHsh"
        );
        return hashTuplePreImage(bytes32(preImage.intVal), preImage.size);
    }

    function tuplePreImage(bytes32[] memory hashes)
        private
        pure
        returns (bytes32)
    {
        require(hashes.length <= 8, "Invalid tuple length");

        bytes32 firstHash = keccak256(
            abi.encodePacked(uint8(hashes.length), hashes)
        );

        return (firstHash);
    }

    function hashTuple(bytes32[] memory hashes, uint256 size)
        private
        pure
        returns (bytes32)
    {
        Data memory preImage = newTuplePreImage(hashes, size);
        return hashTuplePreImage(preImage);
    }

    function hashTuple(Data memory val) internal pure returns (bytes32) {
        Data memory preImage = getTuplePreImage(val);
        return hashTuplePreImage(preImage);
    }

    function hashEmptyTuple() internal pure returns (bytes32) {
        bytes32[] memory hashes = new bytes32[](0);
        return hashTuple(hashes, uint256(1));
    }

    function typeCodeVal(Data memory val) internal pure returns (Data memory) {
        require(val.typeCode != 2, "Value must have a valid type code");
        if (val.typeCode == 0) {
            return newInt(0);
        } else if (val.typeCode == 1) {
            return newInt(1);
        } else {
            return newInt(3);
        }
    }

    function valLength(Data memory val) internal pure returns (uint8) {
        if (isTupleType(val.typeCode)) {
            return val.typeCode - TUPLE_TYPECODE;
        } else {
            return 1;
        }
    }

    function isInt(Data memory val) internal pure returns (bool) {
        return val.typeCode == INT_TYPECODE;
    }

    function isCodePoint(Data memory val) internal pure returns (bool) {
        return val.typeCode == CODE_POINT_TYPECODE;
    }

    function isTuple(Data memory val) internal pure returns (bool) {
        return isTupleType(val.typeCode);
    }

    function hash(Data memory val) internal pure returns (bytes32) {
        if (val.typeCode == INT_TYPECODE) {
            return hashInt(val.intVal);
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return
                hashCodePoint(
                    val.cpVal.opcode,
                    val.cpVal.immediate,
                    val.cpVal.immediateHash,
                    val.cpVal.nextCodePoint
                );
        } else if (val.typeCode == HASH_PRE_IMAGE_TYPECODE) {
            return hashTuplePreImage(val);
        } else if (
            val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT
        ) {
            return hashTuple(val);
        } else if (val.typeCode == HASH_ONLY) {
            return bytes32(val.intVal);
        } else {
            require(false, "Invalid type code");
        }
    }

    function isValidTypeForSend(Data memory val) internal pure returns (bool) {
        if (val.typeCode == INT_TYPECODE) {
            return true;
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return false;
        } else if (val.typeCode == HASH_PRE_IMAGE_TYPECODE) {
            require(false, "must have full value");
        } else if (
            val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT
        ) {
            uint256 valueCount = val.tupleVal.length;
            for (uint256 i = 0; i < valueCount; i++) {
                if (!isValidTypeForSend(val.tupleVal[i])) {
                    return false;
                }
            }
            return true;
        } else if (val.typeCode == HASH_ONLY) {
            return false;
        } else {
            require(false, "Invalid type code");
        }
    }

    function newNone() internal pure returns (Data memory) {
        return
            Data(
                0,
                CodePoint(0, 0, false, 0, 0),
                new Data[](0),
                TUPLE_TYPECODE,
                uint256(1)
            );
    }

    function newBoolean(bool val) internal pure returns (Data memory) {
        if (val) {
            return newInt(1);
        } else {
            return newInt(0);
        }
    }

    function newInt(uint256 _val) internal pure returns (Data memory) {
        return
            Data(
                _val,
                CodePoint(0, 0, false, 0, 0),
                new Data[](0),
                INT_TYPECODE,
                uint256(1)
            );
    }

    function newCodePoint(uint8 opCode, bytes32 nextHash)
        internal
        pure
        returns (Data memory)
    {
        return newCodePoint(CodePoint(opCode, nextHash, false, 0, 0));
    }

    function newCodePoint(
        uint8 opCode,
        bytes32 nextHash,
        Data memory immediate
    ) internal pure returns (Data memory) {
        return
            newCodePoint(
                CodePoint(
                    opCode,
                    nextHash,
                    true,
                    immediate.hash(),
                    immediate.size
                )
            );
    }

    function newHashedValue(bytes32 valueHash, uint256 valueSize)
        internal
        pure
        returns (Data memory)
    {
        return
            Data(
                uint256(valueHash),
                CodePoint(0, 0, false, 0, 0),
                new Data[](0),
                HASH_ONLY,
                valueSize
            );
    }

    function isValidTupleSize(uint256 size) internal pure returns (bool) {
        return size <= 8;
    }

    function newTuple(Data[] memory _val) internal pure returns (Data memory) {
        require(isValidTupleSize(_val.length), "Tuple must have valid size");
        uint256 size = 1;

        for (uint256 i = 0; i < _val.length; i++) {
            size += _val[i].size;
        }

        return
            Data(
                0,
                CodePoint(0, 0, false, 0, 0),
                _val,
                uint8(TUPLE_TYPECODE + _val.length),
                size
            );
    }

    function getTuplePreImage(Data memory tuple)
        internal
        pure
        returns (Data memory)
    {
        require(isTuple(tuple), "Must be Tuple type");
        bytes32[] memory hashes = hashTupleInner(tuple.tupleVal);
        return newTuplePreImage(hashes, tuple.size);
    }

    function newTuplePreImage(bytes32[] memory preImageHashes, uint256 size)
        private
        pure
        returns (Data memory)
    {
        bytes32 firstHash = tuplePreImage(preImageHashes);
        return newTuplePreImage(firstHash, size);
    }

    function newTuplePreImage(bytes32 preImageHash, uint256 size)
        internal
        pure
        returns (Data memory)
    {
        return
            Data(
                uint256(preImageHash),
                CodePoint(0, 0, false, 0, 0),
                new Data[](0),
                HASH_PRE_IMAGE_TYPECODE,
                size
            );
    }

    function deserializeHashed(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            bytes32
        )
    {
        bytes32 hashData;
        uint256 totalLength = data.length;

        if (totalLength < startOffset || totalLength - startOffset < 32) {
            return (false, startOffset, hashData);
        } else {
            hashData = data.toBytes32(startOffset);
            return (true, startOffset + 32, hashData);
        }
    }

    function deserializeHashPreImage(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            Data memory
        )
    {
        Data memory hashValue;
        uint256 size;
        bool valid;

        uint256 totalLength = data.length;
        if (totalLength < startOffset || totalLength - startOffset < 64) {
            return (false, startOffset, hashValue);
        }

        bytes32 hashData = data.toBytes32(startOffset);
        startOffset += 32;
        (valid, startOffset, size) = deserializeInt(data, startOffset);

        if (valid) {
            hashValue = newTuplePreImage(hashData, size);

            return (true, startOffset, hashValue);
        } else {
            return (false, startOffset, hashValue);
        }
    }

    function deserializeInt(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            uint256 // val
        )
    {
        uint256 totalLength = data.length;
        if (totalLength < startOffset || totalLength - startOffset < 32) {
            return (false, startOffset, 0);
        }
        return (true, startOffset + 32, data.toUint(startOffset));
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
            uint8(data[startOffset]) != INT_TYPECODE
        ) {
            return (false, startOffset, 0);
        }
        return (true, startOffset + 33, data.toUint(startOffset + 1));
    }

    function deserializeCodePoint(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            Data memory // val
        )
    {
        uint256 offset = startOffset;
        uint8 immediateType = uint8(data[offset]);
        offset++;
        uint8 opCode = uint8(data[offset]);
        offset++;
        Data memory immediate;
        if (immediateType == 1) {
            bool valid;
            (valid, offset, immediate) = deserialize(data, offset);
            if (!valid) {
                return (false, startOffset, newNone());
            }
        }
        bytes32 nextHash = data.toBytes32(offset);
        offset += 32;
        if (immediateType == 1) {
            return (true, offset, newCodePoint(opCode, nextHash, immediate));
        }
        return (true, offset, newCodePoint(opCode, nextHash));
    }

    function deserializeTuple(
        uint8 memberCount,
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            Data[] memory // val
        )
    {
        uint256 offset = startOffset;
        bool valid;
        Data[] memory members = new Data[](memberCount);
        for (uint8 i = 0; i < memberCount; i++) {
            (valid, offset, members[i]) = deserialize(data, offset);
            if (!valid) {
                return (false, startOffset, members);
            }
        }
        return (true, offset, members);
    }

    function deserialize(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            bool, // valid
            uint256, // offset
            Data memory // val
        )
    {
        if (startOffset >= data.length) {
            return (false, startOffset, newInt(0));
        }
        bool valid;
        uint256 offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        if (valType == INT_TYPECODE) {
            (valid, offset, intVal) = deserializeInt(data, offset);
            return (valid, offset, newInt(intVal));
        } else if (valType == CODE_POINT_TYPECODE) {
            return deserializeCodePoint(data, offset);
        } else if (valType == HASH_PRE_IMAGE_TYPECODE) {
            return deserializeHashPreImage(data, offset);
        } else if (valType >= TUPLE_TYPECODE && valType < VALUE_TYPE_COUNT) {
            uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
            Data[] memory tupleVal;
            (valid, offset, tupleVal) = deserializeTuple(
                tupLength,
                data,
                offset
            );
            return (valid, offset, newTuple(tupleVal));
        }
        return (false, 0, newInt(0));
    }

    function bytesToBytestack(
        bytes memory data,
        uint256 startOffset,
        uint256 dataLength
    ) internal pure returns (Data memory) {
        uint256 wholeChunkCount = dataLength / 32;

        // tuple code + size + (for each chunk tuple code + chunk val) + empty tuple code
        bytes32 stackHash = hashEmptyTuple();
        uint256 size = 1;
        bytes32[] memory vals = new bytes32[](2);

        for (uint256 i = 0; i < wholeChunkCount; i++) {
            vals[0] = newInt(data.toUint(startOffset + i * 32)).hash();
            vals[1] = stackHash;
            size += 2;

            stackHash = hashTuple(vals, size);
        }

        if (dataLength % 32 != 0) {
            uint256 lastVal = data.toUint(startOffset + dataLength - 32);
            lastVal <<= (32 - (dataLength % 32)) * 8;

            vals[0] = newInt(lastVal).hash();
            vals[1] = stackHash;
            size += 2;

            stackHash = hashTuple(vals, size);
        }

        vals[0] = newInt(dataLength).hash();
        vals[1] = stackHash;
        size += 2;

        return newTuplePreImage(vals, size);
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
        if (valType != TUPLE_TYPECODE + 2) {
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
            if (valType != TUPLE_TYPECODE + 2) {
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
        if (valType != TUPLE_TYPECODE) {
            return (false, offset, byteData);
        }
        return (true, offset, abi.encodePacked(fullChunks, partialChunk));
    }

    function newCodePoint(CodePoint memory _val)
        private
        pure
        returns (Data memory)
    {
        return Data(0, _val, new Data[](0), CODE_POINT_TYPECODE, uint256(1));
    }
}
