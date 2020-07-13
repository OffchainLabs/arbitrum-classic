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

library Value {
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

    function tuplePreImageTypeCode() internal pure returns (uint8) {
        return HASH_PRE_IMAGE_TYPECODE;
    }

    function intTypeCode() internal pure returns (uint8) {
        return INT_TYPECODE;
    }

    function codePointTypeCode() internal pure returns (uint8) {
        return CODE_POINT_TYPECODE;
    }

    function valueTypeCode() internal pure returns (uint8) {
        return VALUE_TYPE_COUNT;
    }

    function isTupleType(uint8 typeCode) private pure returns (bool) {
        return typeCode < VALUE_TYPE_COUNT && typeCode >= TUPLE_TYPECODE;
    }

    function isValidTupleSize(uint256 size) internal pure returns (bool) {
        return size <= 8;
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

    function newNone() internal pure returns (Data memory) {
        return newTuple(new Data[](0));
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

    function newCodePoint(CodePoint memory _val)
        private
        pure
        returns (Data memory)
    {
        return Data(0, _val, new Data[](0), CODE_POINT_TYPECODE, uint256(1));
    }
}
