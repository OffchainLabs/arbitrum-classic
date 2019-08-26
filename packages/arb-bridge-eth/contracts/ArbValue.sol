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

import "bytes/BytesLib.sol";

library ArbValue {
    using BytesLib for bytes;
    using ArbValue for Value;

    uint8 constant IntTypeCode = 0;
    uint8 constant CodePointCode = 1;
    uint8 constant HashOnlyTypeCode = 2;
    uint8 constant TupleTypeCode = 3;
    uint8 constant ValueTypeCount = TupleTypeCode + 9;

    function isTupleType(uint8 typeCode) private pure returns (bool) {
        return typeCode < ValueTypeCount && typeCode >= TupleTypeCode;
    }

    function typeLength(uint8 typeCode) private pure returns (uint8) {
        if (isTupleType(typeCode)) {
            return typeCode - TupleTypeCode;
        } else {
            return 1;
        }
    }

    function hashIntValue(uint256 val) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(val));
    }

    function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                CodePointCode,
                opcode,
                nextCodePoint
            )
        );
    }

    function hashCodePointImmediateValue(
        uint8 opcode,
        bytes32 immediateVal,
        bytes32 nextCodePoint
    ) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                CodePointCode,
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
                uint8(TupleTypeCode),
                hashes
            )
        );
    }

    function hashTupleValue(Value[1] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        for (uint i = 0; i < vals.length; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }

    function hashTupleValue(Value[2] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        for (uint i = 0; i < vals.length; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }

    function hashTupleValue(Value[3] memory val) internal pure returns (bytes32) {
        Value[] memory vals = new Value[](val.length);
        for (uint i = 0; i < vals.length; i++) {
            vals[i] = val[i];
        }
        return hashTupleValue(vals);
    }


    function hashTupleValue(Value[] memory val) private pure returns (bytes32) {
        require(val.length <= 8, "Invalid tuple length");
        bytes32[] memory hashes = new bytes32[](val.length);
        for (uint i = 0; i < hashes.length; i++) {
            HashOnlyValue memory hashVal = val[i].hash();
            hashes[i] = hashVal.hash;
        }
        return keccak256(
            abi.encodePacked(
                uint8(TupleTypeCode + val.length),
                hashes
            )
        );
    }

    struct HashOnlyValue {
        bytes32 hash;
    }

    function deserialize_hash_only_value(
        bytes memory data,
        uint startOffset
    ) internal pure returns(uint retCode, uint, HashOnlyValue memory) {
        uint offset = startOffset;
        bytes32 valHash = data.toBytes32(offset);
        offset += 32;
        return (0, offset, HashOnlyValue(valHash));
    }

    struct Value {
        uint256 intVal;
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
        return val.typeCode == IntTypeCode;
    }

    function isTuple(Value memory val) internal pure returns (bool) {
        return isTupleType(val.typeCode);
    }

    function hash(Value memory val) internal pure returns (HashOnlyValue memory) {
        require(val.typeCode < ValueTypeCount, "Invalid type code");
        if (val.typeCode == IntTypeCode) {
            return HashOnlyValue(hashIntValue(val.intVal));
        } else if (val.typeCode == HashOnlyTypeCode) {
            return HashOnlyValue(bytes32(val.intVal));
        } else if (val.typeCode >= TupleTypeCode && val.typeCode < ValueTypeCount) {
            return HashOnlyValue(hashTupleValue(val.tupleVal));
        } else {
            assert(false);
        }
    }

    function newNoneValue() internal pure returns (Value memory) {
        return Value(0, new Value[](0), TupleTypeCode);
    }

    function newBooleanValue(bool val) internal pure returns (Value memory) {
        if (val) {
            return newIntValue(1);
        } else {
            return newIntValue(0);
        }
    }

    function newIntValue(uint256 _val) internal pure returns (Value memory) {
        return Value(_val, new Value[](0), IntTypeCode);
    }

    function isValidTupleSize(uint size) public pure returns (bool) {
        return size <= 8;
    }

    function newTupleValue(Value[] memory _val) internal pure returns (Value memory) {
        require(isValidTupleSize(_val.length), "Tuple must have valid size");
        return Value(0, _val, uint8(TupleTypeCode + _val.length));
    }

    function newTupleHashValues(HashOnlyValue[] memory _val) internal pure returns (Value memory) {
        Value[] memory values = new Value[](_val.length);
        for (uint i = 0; i < _val.length; i++) {
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
        return Value(uint256(_val), new Value[](0), HashOnlyTypeCode);
    }

    function deserialize_int(bytes memory data, uint startOffset) internal pure returns (uint, uint256) {
        uint offset = startOffset;
        uint256 intVal = data.toUint(offset);
        offset += 32;
        return (offset, intVal);
    }

    function deserialize_tuple(
        uint8 memberCount,
        bytes memory data,
        uint startOffset
    ) internal pure returns (uint, uint, Value[] memory) {
        uint offset = startOffset;
        uint retVal;
        Value[] memory members = new Value[](memberCount);
        for (uint8 i = 0; i < memberCount; i++) {
            (retVal, offset, members[i]) = deserialize_value(data, offset);
            if (retVal != 0) {
                return (retVal, offset, members);
            }
        }
        return (0, offset, members);
    }

    function deserialize_value(
        bytes memory data,
        uint startOffset
    ) internal pure returns(uint retCode, uint, Value memory) {
        uint offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        if (valType == IntTypeCode) {
            (offset, intVal) = deserialize_int(data, offset);
            return (0, offset, newIntValue(intVal));
        } else if (valType == HashOnlyTypeCode) {
            (offset, intVal) = deserialize_int(data, offset);
            return (0, offset, newHashOnlyValue(bytes32(intVal)));
        } else if (valType >= TupleTypeCode && valType < ValueTypeCount) {
            uint8 tupLength = uint8(valType - TupleTypeCode);
            Value[] memory tupleVal;
            uint valid;
            (valid, offset, tupleVal) = deserialize_tuple(tupLength, data, offset);
            return (valid, offset, newTupleValue(tupleVal));
        }
        return (10000 + uint(valType), 0, newIntValue(0));
    }

    function deserialize_valid_value_hash(bytes memory data, uint offset) public pure returns(uint, bytes32) {
        uint valid;
        uint newOffset;
        Value memory value;
        (valid, newOffset, value) = deserialize_value(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (newOffset, value.hash().hash);
    }

    function get_next_valid_value(bytes memory data, uint offset) public pure returns(uint, bytes memory) {
        uint valid;
        uint nextOffset;
        Value memory value;
        (valid, nextOffset, value) = deserialize_value(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (nextOffset, data.slice(offset, nextOffset - offset));
    }

    function deserialize_value_hash(bytes memory data) public pure returns (bytes32) {
        uint valid;
        uint offset = 0;
        Value memory value;
        (valid, offset, value) = deserialize_value(data, 0);
        require(valid == 0, "Marshalled value must be valid");
        return value.hash().hash;
    }
}
