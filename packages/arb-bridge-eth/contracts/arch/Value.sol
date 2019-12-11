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
    uint8 internal constant HASH_ONLY_TYPECODE = 2;
    uint8 internal constant TUPLE_TYPECODE = 3;
    uint8 internal constant VALUE_TYPE_COUNT = TUPLE_TYPECODE + 9;

    struct CodePoint {
        uint8 opcode;
        bytes32 nextCodePoint;
        bool immediate;
        bytes32 immediateVal;
    }

    struct HashOnly {
        bytes32 hash;
    }

    struct Data {
        uint256 intVal;
        CodePoint cpVal;
        Data[] tupleVal;
        uint8 typeCode;
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

    function hashInt(uint256 val) public pure returns (bytes32) {
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

    function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) public pure returns (bytes32) {
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

    function hashTuple(Data[1] memory val) internal pure returns (bytes32) {
        Data[] memory vals = new Data[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTuple(vals);
    }

    function hashTuple(Data[2] memory val) internal pure returns (bytes32) {
        Data[] memory vals = new Data[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTuple(vals);
    }

    function hashTuple(Data[3] memory val) internal pure returns (bytes32) {
        Data[] memory vals = new Data[](val.length);
        uint valCount = vals.length;
        for (uint i = 0; i < valCount; i++) {
            vals[i] = val[i];
        }
        return hashTuple(vals);
    }

    function hashTuple(Data[] memory val) private pure returns (bytes32) {
        require(val.length <= 8, "Invalid tuple length");
        bytes32[] memory hashes = new bytes32[](val.length);
        uint hashCount = hashes.length;
        for (uint i = 0; i < hashCount; i++) {
            HashOnly memory hashVal = val[i].hash();
            hashes[i] = hashVal.hash;
        }
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE + val.length),
                hashes
            )
        );
    }

    function hashTuple(bytes32[] memory hashes) private pure returns (bytes32) {
        require(hashes.length <= 8, "Invalid tuple length");
        return keccak256(
            abi.encodePacked(
                uint8(TUPLE_TYPECODE + hashes.length),
                hashes
            )
        );
    }

    function deserializeHashOnly(
        bytes memory data,
        uint startOffset
    )
        internal
        pure
        returns(uint retCode, uint, HashOnly memory)
    {
        uint offset = startOffset;
        bytes32 valHash = data.toBytes32(offset);
        offset += 32;
        return (0, offset, HashOnly(valHash));
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

    function hash(Data memory val) internal pure returns (HashOnly memory) {
        require(val.typeCode < VALUE_TYPE_COUNT, "Invalid type code");
        if (val.typeCode == INT_TYPECODE) {
            return HashOnly(hashInt(val.intVal));
        } else if (val.typeCode == CODE_POINT_TYPECODE) {
            return HashOnly(hashCodePoint(val.cpVal.opcode, val.cpVal.immediate, val.cpVal.immediateVal, val.cpVal.nextCodePoint));
        } else if (val.typeCode == HASH_ONLY_TYPECODE) {
            return HashOnly(bytes32(val.intVal));
        } else if (val.typeCode >= TUPLE_TYPECODE && val.typeCode < VALUE_TYPE_COUNT) {
            return HashOnly(hashTuple(val.tupleVal));
        } else {
            assert(false);
        }
    }

    function newNone() internal pure returns (Data memory) {
        return Data(0, CodePoint(0, 0, false, 0), new Data[](0), TUPLE_TYPECODE);
    }

    function newBoolean(bool val) internal pure returns (Data memory) {
        if (val) {
            return newInt(1);
        } else {
            return newInt(0);
        }
    }

    function newInt(uint256 _val) internal pure returns (Data memory) {
        return Data(_val, CodePoint(0, 0, false, 0), new Data[](0), INT_TYPECODE);
    }

    function newCodePoint(CodePoint memory _val) internal pure returns (Data memory) {
        return Data(0, _val, new Data[](0), CODE_POINT_TYPECODE);
    }

    function isValidTupleSize(uint size) public pure returns (bool) {
        return size <= 8;
    }

    function newTuple(Data[] memory _val) internal pure returns (Data memory) {
        require(isValidTupleSize(_val.length), "Tuple must have valid size");
        return Data(0, CodePoint(0, 0, false, 0), _val, uint8(TUPLE_TYPECODE + _val.length));
    }

    function newTupleHashValues(HashOnly[] memory _val) internal pure returns (Data memory) {
        Data[] memory values = new Data[](_val.length);
        uint valCount = _val.length;
        for (uint i = 0; i < valCount; i++) {
            values[i] = newHashOnly(_val[i].hash);
        }
        return newTuple(values);
    }

    function newRepeatedTuple(Data memory _val, uint8 _count) internal pure returns (Data memory) {
        Data[] memory values = new Data[](_count);
        for (uint i = 0; i < _count; i++) {
            values[i] = _val;
        }
        return newTuple(values);
    }

    function newHashOnly(bytes32 _val) internal pure returns (Data memory) {
        return Data(uint256(_val), CodePoint(0, 0, false, 0), new Data[](0), HASH_ONLY_TYPECODE);
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
            Data memory value;
            (valid, offset, value) = deserialize(data, offset);
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
        returns (uint, uint, Data[] memory)
    {
        uint offset = startOffset;
        uint retVal;
        Data[] memory members = new Data[](memberCount);
        for (uint8 i = 0; i < memberCount; i++) {
            (retVal, offset, members[i]) = deserialize(data, offset);
            if (retVal != 0) {
                return (retVal, offset, members);
            }
        }
        return (0, offset, members);
    }

    function deserialize(
        bytes memory data,
        uint startOffset
    )
        internal
        pure
        returns(uint retCode, uint, Data memory)
    {
        require(startOffset < data.length, "Data offset out of bounds");
        uint offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;
        uint256 intVal;
        CodePoint memory cpVal;
        if (valType == INT_TYPECODE) {
            (offset, intVal) = deserializeInt(data, offset);
            return (0, offset, newInt(intVal));
        } else if (valType == CODE_POINT_TYPECODE) {
            (offset, cpVal) = deserializeCodePoint(data, offset);
            return (0, offset, newCodePoint(cpVal));
        } else if (valType == HASH_ONLY_TYPECODE) {
            (offset, intVal) = deserializeInt(data, offset);
            return (0, offset, newHashOnly(bytes32(intVal)));
        } else if (valType >= TUPLE_TYPECODE && valType < VALUE_TYPE_COUNT) {
            uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
            Data[] memory tupleVal;
            uint valid;
            (valid, offset, tupleVal) = deserializeTuple(tupLength, data, offset);
            return (valid, offset, newTuple(tupleVal));
        }
        return (10000 + uint(valType), 0, newInt(0));
    }

    function deserializeValidHashed(bytes memory data, uint offset) public pure returns(uint, bytes32) {
        uint valid;
        uint newOffset;
        Data memory value;
        (valid, newOffset, value) = deserialize(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (newOffset, value.hash().hash);
    }

    function getNextValid(bytes memory data, uint offset) public pure returns(uint, bytes memory) {
        uint valid;
        uint nextOffset;
        Data memory value;
        (valid, nextOffset, value) = deserialize(data, offset);
        require(valid == 0, "Marshalled value must be valid");
        return (nextOffset, data.slice(offset, nextOffset - offset));
    }

    function deserializeHashed(bytes memory data) public pure returns (bytes32) {
        uint valid;
        uint offset = 0;
        Data memory value;
        (valid, offset, value) = deserialize(data, 0);
        require(valid == 0, "Marshalled value must be valid");
        return value.hash().hash;
    }

    function deserializeMessage(
        bytes memory data,
        uint startOffset)
        public
        pure
        returns(
            bool valid,
            uint offset,
            bytes32 messageHash,
            uint256 msg_type,
            uint256 sender,
            Data[] memory messageData
        )
    {
        offset = startOffset;
        uint8 valType = uint8(data[offset]);
        offset++;

        // typle_type + what?
        if(valType != TUPLE_TYPECODE){ 
            // return error?
            return (valid, offset, messageHash, msg_type, sender, messageData); 
        }

        (msg_type, offset) = deserializeInt(data, offset);
        valType = uint8(data[offset]);
        offset++;

        (sender, offset) = deserializeInt(data, offset);
        valType = uint8(data[offset]);
        offset++;

        uint8 tupLength = uint8(valType - TUPLE_TYPECODE);
        Data[] memory tupleVal;
        uint tupleValid;
        (tupleValid, offset, tupleVal) = deserializeTuple(tupLength, data, offset);

        valid = true;

        bytes32[] memory hashes = new bytes32[](3);
        hashes[0] = hashInt(msg_type);
        hashes[1] = hashInt(sender);
        hashes[2] = hashTuple(tupleVal);
        messageHash = hashTuple(hashes);

        return (valid, offset, messageHash, msg_type, sender, tupleVal);
    }

    function getERCTokenMsgData(
        Data[] memory data)
        public
        pure
        returns(
            bool valid,
            uint256 tokenAddress,
            uint256 destination,
            uint256 value)
    {
        return (tokenAddress, destination, value);
    }

    function getTransactionMsgData(
        Data[] memory data)
        public
        pure
        returns(
            bool valid,
            uint offset,
            uint256 destination,
            uint256 seqNumber,
            uint256 value,
            bytes memory messageData
        )
    {
        return (destination, seqNumber, value, messageData);
    }

    function getEthMsgData(
        Data[] memory data)
        public
        pure
        returns(
            bool valid,
            uint256 destination,
            uint256 value)
    {
        return (destination, value);
    }
}
