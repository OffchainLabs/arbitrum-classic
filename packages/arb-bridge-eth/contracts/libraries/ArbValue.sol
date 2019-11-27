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

import "solidity-bytes-utils/contracts/BytesLib.sol";


library ArbValue {
    using BytesLib for bytes;
    using ArbValue for Value;

    uint8 internal constant INT_TYPECODE = 0;
    uint8 internal constant CODE_POINT_TYPECODE = 1;
    uint8 internal constant HASH_ONLY_TYPECODE = 2;
    uint8 internal constant TUPLE_TYPECODE = 3;
    uint8 internal constant VALUE_TYPE_COUNT = TUPLE_TYPECODE + 9;

    struct CodePoint {
        uint8 opcode;
        bytes32 nextCodePoint;
        bool immediate;
        bytes32 immediateVal;
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

    function hashIntValue(uint256 val) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(val));
    }

    function hashCodePoint(
            uint8 opcode,
            bool immediate,
            bytes32 immediateVal,
            bytes32 nextCodePoint
    ) public pure returns (bytes32) {
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

    function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                CODE_POINT_TYPECODE,
                opcode,
                nextCodePoint
            )
        );
    }

    function hashCodePointImmediateValue(
        uint8 opcode,
        bytes32 immediateVal,
        bytes32 nextCodePoint
    )
        public
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

    function hashEmptyTuple() public pure returns (bytes32) {
        bytes32[] memory hashes = new bytes32[](0);
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE),
                hashes
            )
        );
    }

    function hashTupleValue(Value[1] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }

    function hashTupleValue(Value[2] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }

    function hashTupleValue(Value[3] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }

    function hashTupleValue(Value[] memory val) private pure returns (bytes32) {
        require(val.length <= 8, "Invalid tuple length");
        bytes32[] memory hashes = new bytes32[](val.length);
        uint hashCount = hashes.length;
        for (uint i = 0; i < hashCount; i++) {
            HashOnlyValue memory hashVal = val[i].hash();
            hashes[i] = hashVal.hash;
        }
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE + val.length),
                hashes
            )
        );
    }

    function hashTupleValue(bytes32[] memory hashes) private pure returns (bytes32) {
        require(hashes.length <= 8, "Invalid tuple length");
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE + hashes.length),
                hashes
            )
        );
    }

    struct HashOnlyValue {
        bytes32 hash;
    }

    function deserializeHashOnlyValue(
        bytes memory data,
        uint startOffset
    )
        internal
        pure
        returns(uint retCode, uint, HashOnlyValue memory)
    {
        uint offset = startOffset;
        bytes32 valHash = data.toBytes32(offset);
        offset += 32;
        return (0, offset, HashOnlyValue(valHash));
    }

    struct Value {
        uint256 intVal;
        CodePoint cpVal;
        Value[] tupleVal;
        uint8 typeCode;
    }

    function typeCodeVal(Value memory val) internal pure returns (Value memory) {
        require(val.typeCode != 2, "Value must have a valid type code");
        if (val.typeCode == 0) {
            return newIntValue(0);
        } else if (val.typeCode == 1) {
            return newIntValue(1);
        } else {
            return newIntValue(3);
        }
    }

    function valLength(Value memory val) internal pure returns (uint8) {
        return typeLength(val.typeCode);
    }

    function isInt(Value memory val) internal pure returns (bool) {
        return val.typeCode == INT_TYPECODE;
    }

    function isCodePoint(Value memory val) internal pure returns (bool) {
        return val.typeCode == CODE_POINT_TYPECODE;
    }

    function isTuple(Value memory val) internal pure returns (bool) {
        return isTupleType(val.typeCode);
    }

    function hash(Value memory val) internal pure returns (HashOnlyValue memory) {
        require(val.typeCode < VALUE_TYPE_COUNT, "Invalid type code");
        if (val.typeCode == INT_TYPECODE) {
            return HashOnlyValue(hashIntValue(val.intVal));
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return HashOnlyValue(hashCodePoint(val.cpVal.opcode, val.cpVal.immediate, val.cpVal.immediateVal, val.cpVal.nextCodePoint));
        } else if (val.typeCode == HASH_ONLY_TYPECODE) {
            return HashOnlyValue(bytes32(val.intVal));
        } else if (val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT) {
            return HashOnlyValue(hashTupleValue(val.tupleVal));
        } else {
            assert(false);
        }
    }

    function newNoneValue() internal pure returns (Value memory) {
        return Value(0, CodePoint(0, 0, false, 0), new Value[](0), TUPLE_TYPECODE);
    }

    function newBooleanValue(bool val) internal pure returns (Value memory) {
        if (val) {
            return newIntValue(1);
        } else {
            return newIntValue(0);
        }
    }

    function newIntValue(uint256 _val) internal pure returns (Value memory) {
        return Value(_val, CodePoint(0, 0, false, 0), new Value[](0), INT_TYPECODE);
    }

    function newCodePointValue(CodePoint memory _val) internal pure returns (Value memory) {
        return Value(0, _val, new Value[](0), CODE_POINT_TYPECODE);
    }

    function isValidTupleSize(uint size) public pure returns (bool) {
        return size <= 8;
    }

    function newTupleValue(Value[] memory _val) internal pure returns (Value memory) {
        require(isValidTupleSize(_val.length), "Tuple must have valid size");
        return Value(0, CodePoint(0, 0, false, 0), _val, uint8(TUPLE_TYPECODE + _val.length));
    }

    function newTupleHashValues(HashOnlyValue[] memory _val) internal pure returns (Value memory) {
        Value[] memory values = new Value[](_val.length);
        uint valCount = _val.length;
        for (uint i = 0; i < valCount; i++) {
            values[i] = newHashOnlyValue(_val[i].hash);
        }
        return newTupleValue(values);
    }

    function newRepeatedTuple(Value memory _val, uint8 _count) internal pure returns (Value memory) {
        Value[] memory values = new Value[](_count);
        for (uint i = 0; i < _count; i++) {
            values[i] = _val;
        }
        return newTupleValue(values);
    }

    function newHashOnlyValue(bytes32 _val) internal pure returns (Value memory) {
        return Value(uint256(_val), CodePoint(0, 0, false, 0), new Value[](0), HASH_ONLY_TYPECODE);
    }

    function deserializeInt(bytes memory data, uint startOffset) internal pure returns (uint, uint256) {
        uint offset = startOffset;
        uint256 intVal = data.toUint(offset);
        offset += 32;
        return (offset, intVal);
    }

    function deserializeCodePoint(bytes memory data, uint startOffset) internal pure returns (uint, CodePoint memory) {
        uint offset = startOffset;
        uint8 immediateType = uint8(data[offset]);
        offset ++;
        uint8 opCode = uint8(data[offset]);
        offset++;
        bytes32 immediateVal;
        if (immediateType == 1) {
            uint valid;
            Value memory value;
            (valid, offset, value) = deserializeValue(data, offset);
            require(valid == 0, "Marshalled value must be valid");
            immediateVal = value.hash().hash;
        }
        bytes32 nextHash = data.toBytes32(offset);
        offset += 32;
        if (immediateType == 1) {
            return (offset, CodePoint(opCode, nextHash, true, immediateVal));
        }
        return (offset, CodePoint(opCode, nextHash, false, 0));
    }

    function deserializeTuple(
        uint8 memberCount,
        bytes memory data,
        uint startOffset
    )
        internal
        pure
        returns (uint, uint, Value[] memory)
    {
        uint offset = startOffset;
        uint retVal;
        Value[] memory members = new Value[](memberCount);
        for (uint8 i = 0; i < memberCount; i++) {
            (retVal, offset, members[i]) = deserializeValue(data, offset);
            if (retVal != 0) {
                return (retVal, offset, members);
            }
        }
        return (0, offset, members);
    }

    function deserializeValue(
        bytes memory data,
        uint startOffset
    )
        internal
        pure
        returns(uint retCode, uint, Value memory)
    {
        require(startOffset < data.length, "Data offset out of bounds");
        uint offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        CodePoint memory cpVal;
        if (valType == INT_TYPECODE) {
            (offset, intVal) = deserializeInt(data, offset);
            return (0, offset, newIntValue(intVal));
        } else if (valType == CODE_POINT_TYPECODE) {
            (offset, cpVal) = deserializeCodePoint(data, offset);
            return (0, offset, newCodePointValue(cpVal));
        } else if (valType == HASH_ONLY_TYPECODE) {
            (offset, intVal) = deserializeInt(data, offset);
            return (0, offset, newHashOnlyValue(bytes32(intVal)));
        } else if (valType >= TUPLE_TYPECODE && valType < VALUE_TYPE_COUNT) {
            uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
            Value[] memory tupleVal;
            uint valid;
            (valid, offset, tupleVal) = deserializeTuple(tupLength, data, offset);
            return (valid, offset, newTupleValue(tupleVal));
        }
        return (10000 + uint(valType), 0, newIntValue(0));
    }

    function deserializeValidValueHash(bytes memory data, uint offset) public pure returns(uint, bytes32) {
        uint valid;
        uint newOffset;
        Value memory value;
        (valid, newOffset, value) = deserializeValue(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (newOffset, value.hash().hash);
    }

    function getNextValidValue(bytes memory data, uint offset) public pure returns(uint, bytes memory) {
        uint valid;
        uint nextOffset;
        Value memory value;
        (valid, nextOffset, value) = deserializeValue(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (nextOffset, data.slice(offset, nextOffset - offset));
    }

    function deserializeValueHash(bytes memory data) public pure returns (bytes32) {
        uint valid;
        uint offset = 0;
        Value memory value;
        (valid, offset, value) = deserializeValue(data, 0);
        require(valid == 0, "Marshalled value must be valid");
        return value.hash().hash;
    }

    function deserializeMessage(
        bytes memory data,
        uint startOffset
    )
        public
        pure
        returns(
            bool valid,
            uint offset,
            bytes32 messageHash,
            uint256 destination,
            uint256 value,
            uint256 tokenType,
            bytes memory messageData
        )
    {
        bytes32 messageDataHash;
        offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        if (valType != TUPLE_TYPECODE + 4) {
            (offset, messageHash) = deserializeValidValueHash(data, offset - 1);
            return (valid, offset, messageHash, destination, value, tokenType, messageData);
        }

        (offset, messageDataHash) = deserializeValidValueHash(data, offset);
        messageData = data.slice(startOffset + 1, offset - startOffset - 1);

        valType = uint8(data[offset]);
        offset++;
        if (valType != INT_TYPECODE) {
            (offset, messageHash) = deserializeValidValueHash(data, offset - 1);
            return (valid, offset, messageHash, destination, value, tokenType, messageData);
        }
        (destination, offset) = deserializeInt(data, offset);

        valType = uint8(data[offset]);
        offset++;
        if (valType != INT_TYPECODE) {
            (offset, messageHash) = deserializeValidValueHash(data, offset - 1);
            return (valid, offset, messageHash, destination, value, tokenType, messageData);
        }
        (value, offset) = deserializeInt(data, offset);

        valType = uint8(data[offset]);
        offset++;
        if (valType != INT_TYPECODE) {
            (offset, messageHash) = deserializeValidValueHash(data, offset - 1);
            return (valid, offset, messageHash, destination, value, tokenType, messageData);
        }
        (tokenType, offset) = deserializeInt(data, offset);

        valid = true;

        bytes32[] memory hashes = new bytes32[](4);
        hashes[0] = messageDataHash;
        hashes[1] = hashIntValue(destination);
        hashes[2] = hashIntValue(value);
        hashes[3] = hashIntValue(tokenType);
        messageHash = hashTupleValue(hashes);

        return (valid, offset, messageHash, destination, value, tokenType, messageData);
    }
}
