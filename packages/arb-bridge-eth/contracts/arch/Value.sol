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

import "bytes/contracts/BytesLib.sol";


library Value {
    using BytesLib for bytes;
    using Value for Data;

    uint8 internal constant INT_TYPECODE = 0;
    uint8 internal constant CODE_POINT_TYPECODE = 1;
    uint8 internal constant HASH_PRE_IMAGE_TYPECODE = 2;
    uint8 internal constant TUPLE_TYPECODE = 3;
    uint8 internal constant VALUE_TYPE_COUNT = TUPLE_TYPECODE + 9;

    struct CodePoint {
        uint8 opcode;
        bytes32 nextCodePoint;
        bool immediate;
        bytes32 immediateVal;
    }

    struct Data {
        uint256 intVal;
        CodePoint cpVal;
        Data[] tupleVal;
        uint8 typeCode;
        uint256 size;
    }

    function isTupleType(uint8 typeCode) private pure returns (bool) {
        return typeCode < VALUE_TYPE_COUNT && typeCode >= TUPLE_TYPECODE;
    }

    function typeLength(uint8 typeCode) private pure returns (uint8) {
        if (isTupleType(typeCode)) {
            return typeCode - TUPLE_TYPECODE;
        } else {
            return 1;
        }
    }

    function hashInt(uint256 val) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(val));
    }

    function hashCodePoint(
        uint8 opcode,
        bool immediate,
        bytes32 immediateVal,
        bytes32 nextCodePoint
    ) internal pure returns (bytes32) {
        if (immediate) {
            return keccak256(
                abi.encodePacked(
                    CODE_POINT_TYPECODE,
                    opcode,
                    immediateVal,
                    nextCodePoint
                )
            );
        }
        return keccak256(
            abi.encodePacked(
                CODE_POINT_TYPECODE,
                opcode,
                nextCodePoint
            )
        );
    }

    function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                CODE_POINT_TYPECODE,
                opcode,
                nextCodePoint
            )
        );
    }

    function hashCodePointImmediate(
        uint8 opcode,
        bytes32 immediateVal,
        bytes32 nextCodePoint
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                CODE_POINT_TYPECODE,
                opcode,
                immediateVal,
                nextCodePoint
            )
        );
    }

    function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE),
                innerHash,
                valueSize
            )
        );
    }

    function hashTupleInner(Data[] memory vals) private pure returns (bytes32[] memory){
        require(vals.length <= 8, "Invalid tuple length");

        bytes32[] memory hashes = new bytes32[](vals.length);

        uint256 hashCount = hashes.length;
        for (uint256 i = 0; i < hashCount; i++) {
            bytes32 hashVal = vals[i].hash();
            hashes[i] = hashVal;
        }

        return hashes;
    }

    function hashTuplePreImage(Data memory preImage) private pure returns (bytes32) {
        require(preImage.typeCode == HASH_PRE_IMAGE_TYPECODE, "Must be PreImageHsh");
        return hashTuplePreImage(bytes32(preImage.intVal), preImage.size);
    }

    function tuplePreImage(bytes32[] memory hashes) private pure returns (bytes32) {
        require(hashes.length <= 8, "Invalid tuple length");

        bytes32 firstHash = keccak256(
            abi.encodePacked(
                uint8(hashes.length),
                hashes
            )
        );

        return (firstHash);
    }

    function hashTuple(bytes32[] memory hashes, uint256 size) private pure returns (bytes32) {
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
        return typeLength(val.typeCode);
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
        require(val.typeCode < VALUE_TYPE_COUNT, "Invalid type code");
        if (val.typeCode == INT_TYPECODE) {
            return hashInt(val.intVal);
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return hashCodePoint(val.cpVal.opcode, val.cpVal.immediate, val.cpVal.immediateVal, val.cpVal.nextCodePoint);
        } else if (val.typeCode == HASH_PRE_IMAGE_TYPECODE) {
            if(val.cpVal.nextCodePoint == bytes32(uint(1))){
                return bytes32(val.intVal);
            }else{
                return hashTuplePreImage(val);
            }
        } else if (val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT) {
            return hashTuple(val);
        } else {
            assert(false);
        }
    }

    function newNone() internal pure returns (Data memory) {
        return Data(0, CodePoint(0, 0, false, 0), new Data[](0), TUPLE_TYPECODE, uint256(1));
    }

    function newBoolean(bool val) internal pure returns (Data memory) {
        if (val) {
            return newInt(1);
        } else {
            return newInt(0);
        }
    }

    function newInt(uint256 _val) internal pure returns (Data memory) {
        return Data(_val, CodePoint(0, 0, false, 0), new Data[](0), INT_TYPECODE, uint256(1));
    }

    function newCodePoint(CodePoint memory _val) internal pure returns (Data memory) {
        return Data(0, _val, new Data[](0), CODE_POINT_TYPECODE, uint256(1));
    }

    function newCodePoint(uint8 opCode, bytes32 nextHash) internal pure returns(Data memory){
        return newCodePoint(CodePoint(opCode, nextHash, false, 0));
    }

    function newCodePoint(uint8 opCode, bytes32 nextHash, bytes32 immediateVal) internal pure returns(Data memory){
        return newCodePoint(CodePoint(opCode, nextHash, true, immediateVal));
    }

    function isValidTupleSize(uint256 size) internal pure returns (bool) {
        return size <= 8;
    }

    function newTuple(Data[] memory _val) internal pure returns (Data memory) {
        require(isValidTupleSize(_val.length), "Tuple must have valid size");
        uint256 size = 1;

        for(uint256 i = 0; i < _val.length; i++){
            size += _val[i].size;
        }

        return Data(0, CodePoint(0, 0, false, 0), _val, uint8(TUPLE_TYPECODE + _val.length), size);
    }

    function newRepeatedTuple(Data memory _val, uint8 _count) internal pure returns (Data memory) {
        Data[] memory values = new Data[](_count);
        for (uint256 i = 0; i < _count; i++) {
            values[i] = _val;
        }
        return newTuple(values);
    }

    function getTuplePreImage(Data memory tuple) internal pure returns (Data memory) {
        require(isTuple(tuple), "Must be Tuple type");
        bytes32[] memory hashes = hashTupleInner(tuple.tupleVal);
        return newTuplePreImage(hashes, tuple.size);
    }

    function newTuplePreImage(bytes32[] memory preImageHashes, uint256 size) private pure returns (Data memory) {
        bytes32 firstHash = tuplePreImage(preImageHashes);
        return newTuplePreImage(firstHash, size);
    }

    function newTuplePreImage(bytes32 preImageHash, uint256 size) internal pure returns (Data memory){
        return Data(uint256(preImageHash), CodePoint(0, 0, false, 0), new Data[](0), HASH_PRE_IMAGE_TYPECODE, size);
    }

    function deserializeHashed(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            bytes32)
    {
        bytes32 hashData;
        uint256 totalLength = data.length;

        if (totalLength < startOffset || totalLength - startOffset < 32) {
            return (false, startOffset, hashData);
        }else{
            hashData = data.toBytes32(startOffset);
            return (true, startOffset + 32, hashData);
        }
    }

    function deserializeHashPreImage(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
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

        if(valid){
            hashValue = newTuplePreImage(hashData, size);

            return (true, startOffset, hashValue);
        }else{
            return (false, startOffset, hashValue);
        }
    }

    function deserializeInt(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
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

    function deserializeCheckedInt(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
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

    function deserializeCodePoint(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            CodePoint memory // val
        ) {
        uint256 offset = startOffset;
        uint8 immediateType = uint8(data[offset]);
        offset++;
        uint8 opCode = uint8(data[offset]);
        offset++;
        bytes32 immediateVal;
        if (immediateType == 1) {
            bool valid;
            Data memory value;
            (valid, offset, value) = deserialize(data, offset);
            if (!valid) {
                return (false, startOffset, CodePoint(0, 0, false, 0));
            }
            immediateVal = value.hash();
        }
        bytes32 nextHash = data.toBytes32(offset);
        offset += 32;
        if (immediateType == 1) {
            return (true, offset, CodePoint(opCode, nextHash, true, immediateVal));
        }
        return (true, offset, CodePoint(opCode, nextHash, false, 0));
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
            uint, // offset
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

    function deserialize(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint, // offset
            Data memory // val
        )
    {
        if(startOffset >= data.length) {
            return (false, startOffset, newInt(0));
        }
        bool valid;
        uint256 offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        CodePoint memory cpVal;
        if (valType == INT_TYPECODE) {
            (valid, offset, intVal) = deserializeInt(data, offset);
            return (valid, offset, newInt(intVal));
        } else if (valType == CODE_POINT_TYPECODE) {
            (valid, offset, cpVal) = deserializeCodePoint(data, offset);
            return (valid, offset, newCodePoint(cpVal));
        } else if (valType == HASH_PRE_IMAGE_TYPECODE) {
            return deserializeHashPreImage(data, offset);
        } else if (valType >= TUPLE_TYPECODE && valType < VALUE_TYPE_COUNT) {
            uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
            Data[] memory tupleVal;
            (valid, offset, tupleVal) = deserializeTuple(tupLength, data, offset);
            return (valid, offset, newTuple(tupleVal));
        }
        return (false, 0, newInt(0));
    }

    function getNextValid(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset,
            bytes memory // dataSlice
        )
    {
        (bool valid, uint256 offset,) = deserialize(data, startOffset);
        if (!valid) {
            return (false, startOffset, new bytes(0));
        }
        return (true, offset, data.slice(startOffset, offset - startOffset));
    }

    function deserializeMessageData(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            uint256, // msgType
            address // sender
        )
    {
        bool valid;
        uint256 msgType;
        uint256 senderRaw;
        uint256 offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;

        if(valType != TUPLE_TYPECODE + 3) {
            return (false, startOffset, 0, address(0));
        }

        (valid, offset, msgType) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, 0, address(0));
        }

        (valid, offset, senderRaw) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, 0, address(0));
        }

        return (
            true,
            offset,
            msgType,
            address(uint160((senderRaw)))
        );
    }

    function getTransactionMsgData(
        bytes memory data
    )
        internal
        pure
        returns(
            bool valid,
            uint256 vmAddress,
            uint256 destination,
            uint256 seqNumber,
            uint256 value,
            bytes memory messageData
        )
    {
        uint offset = 0;
        uint8 valType = uint8(data[offset]);
        offset++;

        if(valType == TUPLE_TYPECODE + 4){

            (valid, offset, destination) = deserializeCheckedInt(data, offset);

            (valid, offset, seqNumber) = deserializeCheckedInt(data, offset);

            (valid, offset, value) = deserializeCheckedInt(data, offset);

            // fix incorrect
            bytes32 messageDataHash;
            (valid, offset, messageDataHash) = deserializeHashed(data, offset);
            messageData = data.slice(1, offset - 1);// fix incorrect

            valid = true;
        }

        return (valid,vmAddress,destination, seqNumber, value, messageData);
    }

    function getEthMsgData(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            address, // destination
            uint256 // value
        )
    {
        bool valid;
        uint256 destRaw;
        uint256 value;
        uint offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;

        if(valType != TUPLE_TYPECODE + 2) {
            return (false, startOffset, address(0), 0);
        }

        (valid, offset, destRaw) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, address(0), 0);
        }

        (valid, offset, value) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, address(0), 0);
        }

        return (
            true,
            offset,
            address(uint160((destRaw))),
            value
        );
    }

    function getERCTokenMsgData(
        bytes memory data,
        uint256 startOffset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            address, // tokenAddress
            address, // destination
            uint256 // value
        )
    {
        bool valid;
        uint256 tokenAddressRaw;
        uint256 destRaw;
        uint256 value;
        uint offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;

        if(valType != TUPLE_TYPECODE + 3) {
            return (false, startOffset, address(0), address(0), 0);
        }

        (valid, offset, tokenAddressRaw) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, address(0), address(0), 0);
        }

        (valid, offset, destRaw) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, address(0), address(0), 0);
        }

        (valid, offset, value) = deserializeCheckedInt(data, offset);
        if (!valid) {
            return (false, startOffset, address(0), address(0), 0);
        }

        return (
            true,
            offset,
            address(uint160((tokenAddressRaw))),
            address(uint160((destRaw))),
            value
        );
    }

    function bytesToBytestackHash(bytes memory data, uint256 startOffset, uint256 dataLength) internal pure returns (Data memory) {
        uint wholeChunkCount = dataLength / 32;
        uint chunkCount = (dataLength + 31) / 32;

        // tuple code + size + (for each chunk tuple code + chunk val) + empty tuple code
        bytes32 stackHash = hashEmptyTuple();
        uint256 size = 1;
        bytes32[] memory vals = new bytes32[](2);

        for (uint i = 0; i < wholeChunkCount; i++) {
            vals[0] = stackHash;
            vals[1] = newInt(data.toUint(startOffset + i * 32)).hash();
            size += 2;

            stackHash = hashTuple(vals, size);
        }

        if (wholeChunkCount < chunkCount) {
            uint lastVal = data.toUint(startOffset + dataLength - 32);
            lastVal <<= (32 - dataLength - wholeChunkCount * 32) * 8;

            vals[0] = stackHash;
            vals[1] = newInt(lastVal).hash();
            size += 2;

            stackHash = hashTuple(vals, size);
        }

        vals[0] = newInt(dataLength).hash();
        vals[1] = stackHash;
        size += 2;

        return newTuplePreImage(vals, size);
    }

    function bytesToBytestackHash(bytes memory data) internal pure returns (Data memory) {
        uint dataLength = data.length;
        uint256 startOffset = 0;

        return bytesToBytestackHash(data, dataLength, startOffset);
    }

    function bytestackToBytes(bytes memory data) internal pure returns (bytes memory) {
        uint byteCount = data.toUint(2);
        uint chunkCount = (byteCount + 31) / 32;

        bytes32[] memory chunks = new bytes32[](chunkCount);
        uint offset = 35;
        for (uint i = 0; i < chunkCount; i++) {
            chunks[i] = data.toBytes32(offset + 2);
            offset += 34;
        }
        return abi.encodePacked(chunks).slice(0, byteCount);
    }
}
