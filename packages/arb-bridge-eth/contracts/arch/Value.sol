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
    using Value for CodePoint;

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
        Data[] immediate;
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

    function hashInt(uint256 val) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(val));
    }

    function hashCodePoint(CodePoint memory cp)
        internal
        pure
        returns (bytes32)
    {
        assert(cp.immediate.length < 2);
        if (cp.immediate.length == 0) {
            return
                keccak256(
                    abi.encodePacked(
                        CODE_POINT_TYPECODE,
                        cp.opcode,
                        cp.nextCodePoint
                    )
                );
        }
        return
            keccak256(
                abi.encodePacked(
                    CODE_POINT_TYPECODE,
                    cp.opcode,
                    cp.immediate[0].hash(),
                    cp.nextCodePoint
                )
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

    function hashEmptyTuple() internal pure returns (bytes32) {
        return newNone().hash();
    }

    function hash(Data memory val) internal pure returns (bytes32) {
        if (val.typeCode == INT_TYPECODE) {
            return hashInt(val.intVal);
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return hashCodePoint(val.cpVal);
        } else if (val.typeCode == HASH_PRE_IMAGE_TYPECODE) {
            return hashTuplePreImage(bytes32(val.intVal), val.size);
        } else if (
            val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT
        ) {
            Data memory preImage = getTuplePreImage(val.tupleVal);
            return preImage.hash();
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
                CodePoint(0, 0, new Data[](0)),
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
                CodePoint(0, 0, new Data[](0)),
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
        return newCodePoint(CodePoint(opCode, nextHash, new Data[](0)));
    }

    function newCodePoint(
        uint8 opCode,
        bytes32 nextHash,
        Data memory immediate
    ) internal pure returns (Data memory) {
        Data[] memory imm = new Data[](1);
        imm[0] = immediate;
        return newCodePoint(CodePoint(opCode, nextHash, imm));
    }

    function newHashedValue(bytes32 valueHash, uint256 valueSize)
        internal
        pure
        returns (Data memory)
    {
        return
            Data(
                uint256(valueHash),
                CodePoint(0, 0, new Data[](0)),
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
                CodePoint(0, 0, new Data[](0)),
                _val,
                uint8(TUPLE_TYPECODE + _val.length),
                size
            );
    }

    function getTuplePreImage(Data[] memory vals)
        internal
        pure
        returns (Data memory)
    {
        require(vals.length <= 8, "Invalid tuple length");
        bytes32[] memory hashes = new bytes32[](vals.length);
        uint256 hashCount = hashes.length;
        uint256 size = 1;
        for (uint256 i = 0; i < hashCount; i++) {
            hashes[i] = vals[i].hash();
            size += vals[i].size;
        }
        bytes32 firstHash = keccak256(
            abi.encodePacked(uint8(hashes.length), hashes)
        );
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
                CodePoint(0, 0, new Data[](0)),
                new Data[](0),
                HASH_PRE_IMAGE_TYPECODE,
                size
            );
    }

    function deserializeHashPreImage(bytes memory data, uint256 startOffset)
        internal
        pure
        returns (
            uint256, // offset
            Data memory
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
        Data memory hashValue = newTuplePreImage(hashData, size);
        return (startOffset, hashValue);
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
            (offset, immediate) = deserialize(data, offset);
        }
        bytes32 nextHash = data.toBytes32(offset);
        offset += 32;
        if (immediateType == 1) {
            return (offset, newCodePoint(opCode, nextHash, immediate));
        }
        return (offset, newCodePoint(opCode, nextHash));
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
            Data[] memory // val
        )
    {
        uint256 offset = startOffset;
        Data[] memory members = new Data[](memberCount);
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
            Data memory // val
        )
    {
        require(startOffset < data.length, "invalid offset");
        uint256 offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        if (valType == INT_TYPECODE) {
            (offset, intVal) = deserializeInt(data, offset);
            return (offset, newInt(intVal));
        } else if (valType == CODE_POINT_TYPECODE) {
            return deserializeCodePoint(data, offset);
        } else if (valType == HASH_PRE_IMAGE_TYPECODE) {
            return deserializeHashPreImage(data, offset);
        } else if (valType >= TUPLE_TYPECODE && valType < VALUE_TYPE_COUNT) {
            uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
            Data[] memory tupleVal;
            (offset, tupleVal) = deserializeTuple(tupLength, data, offset);
            return (offset, newTuple(tupleVal));
        }
        require(false, "invalid typecode");
    }

    function bytesToBytestack(
        bytes memory data,
        uint256 startOffset,
        uint256 dataLength
    ) internal pure returns (Data memory) {
        uint256 wholeChunkCount = dataLength / 32;

        // tuple code + size + (for each chunk tuple code + chunk val) + empty tuple code
        Value.Data memory stack = newNone();
        Value.Data[] memory vals = new Value.Data[](2);

        for (uint256 i = 0; i < wholeChunkCount; i++) {
            vals[0] = newInt(data.toUint(startOffset + i * 32));
            vals[1] = stack;
            stack = getTuplePreImage(vals);
        }

        if (dataLength % 32 != 0) {
            uint256 lastVal = data.toUint(startOffset + dataLength - 32);
            lastVal <<= (32 - (dataLength % 32)) * 8;
            vals[0] = newInt(lastVal);
            vals[1] = stack;
            stack = getTuplePreImage(vals);
        }

        vals[0] = newInt(dataLength);
        vals[1] = stack;

        return getTuplePreImage(vals);
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
