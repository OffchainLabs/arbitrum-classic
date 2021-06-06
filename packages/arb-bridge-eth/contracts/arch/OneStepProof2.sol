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

/*

Structure of the extra proofs passed to operations accessing buffers (d, array of words):
 * d_0: 32 bytes header, includes the locations of other proofs as the first 5 bytes b_0 ... b_1
 * The words d[b_0..b_1]: merkle proof for first access, first element is the leaf that is accessed
 * The words d[b_1..b_2]: normalization proof for the case the buffer shrinks
 * The words d[b_2..b_3]: merkle proof for second access
 * The words d[b_4..b_5]: normalization proof for second access

Structure of merkle proofs:
 * first element is the leaf
 * other elements are the adjacent subtrees
 * the location in the tree is known from the argument passed to opcodes
 * if the access is outside the tree, the merkle proof is needed to confirm the size of the tree, and is the accessed location mod the original size of the tree

Structure of normalization proof:
 * needed if the tree shrinks
 * has three words
 * height of the tree (minus one)
 * left subtree hash
 * right subtree hash
 * if the height of the tree is 0, the left subtree hash is the single leaf of the tree instead
 * right subtree hash is checked that it's not zero, this ensures that the resulting tree is of minimal height

*/

pragma solidity ^0.6.11;

import "./IOneStepProof.sol";
import "./OneStepProofCommon.sol";
import "./Value.sol";
import "./Machine.sol";

// Originally forked from https://github.com/leapdao/solEVM-enforcer/tree/master

contract OneStepProof2 is OneStepProofCommon {
    /* solhint-disable no-inline-assembly */

    function makeZeros() internal pure returns (bytes32[] memory) {
        bytes32[] memory zeros = new bytes32[](64);
        zeros[0] = keccak1(0);
        for (uint256 i = 1; i < 64; i++) {
            zeros[i] = keccak2(zeros[i - 1], zeros[i - 1]);
        }
        return zeros;
    }

    function keccak1(bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(b));
    }

    function keccak2(bytes32 a, bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(a, b));
    }

    // hashes are normalized
    function get(
        bytes32 buf,
        uint256 loc,
        bytes32[] memory proof
    ) internal pure returns (bytes32) {
        // empty tree is full of zeros
        if (proof.length == 0) {
            require(buf == keccak1(bytes32(0)), "expected empty buffer");
            return 0;
        }
        bytes32 acc = keccak1(proof[0]);
        for (uint256 i = 1; i < proof.length; i++) {
            if (loc & 1 == 1) acc = keccak2(proof[i], acc);
            else acc = keccak2(acc, proof[i]);
            loc = loc >> 1;
        }
        require(acc == buf, "expected correct root");
        // maybe it is a zero outside the actual tree
        if (loc > 0) return 0;
        return proof[0];
    }

    function checkSize(
        bytes32 buf,
        uint256 loc,
        bytes32[] memory proof
    ) internal pure returns (bool) {
        // empty tree is full of zeros
        if (proof.length == 0) {
            require(buf == keccak1(bytes32(0)), "expected empty buffer");
            return true;
        }
        bytes32 acc = keccak1(proof[0]);
        bool check = true;
        bytes32[] memory zeros = makeZeros();
        for (uint256 i = 1; i < proof.length; i++) {
            if (loc & 1 == 1) acc = keccak2(proof[i], acc);
            else {
                acc = keccak2(acc, proof[i]);
                check = check && proof[i] == zeros[i - 1];
            }
            loc = loc >> 1;
        }
        require(acc == buf, "expected correct root");
        // maybe it is a zero outside the actual tree
        if (loc > 0) return true;
        return check;
    }

    function calcHeight(uint256 loc) internal pure returns (uint256) {
        if (loc == 0) return 1;
        else return 1 + calcHeight(loc >> 1);
    }

    function set(
        bytes32 buf,
        uint256 loc,
        bytes32 v,
        bytes32[] memory proof,
        uint256 nh,
        bytes32 normal1,
        bytes32 normal2
    ) internal pure returns (bytes32) {
        // three possibilities, the tree depth stays same, it becomes lower or it's extended
        bytes32 acc = keccak1(v);
        // check that the proof matches original
        get(buf, loc, proof);
        bytes32[] memory zeros = makeZeros();
        // extended
        if (loc >= (1 << (proof.length - 1))) {
            if (v == 0) return buf;
            uint256 height = calcHeight(loc);
            // build the left branch
            for (uint256 i = proof.length; i < height - 1; i++) {
                buf = keccak2(buf, zeros[i - 1]);
            }
            for (uint256 i = 1; i < height - 1; i++) {
                if (loc & 1 == 1) acc = keccak2(zeros[i - 1], acc);
                else acc = keccak2(acc, zeros[i - 1]);
                loc = loc >> 1;
            }
            return keccak2(buf, acc);
        }
        for (uint256 i = 1; i < proof.length; i++) {
            bytes32 a = loc & 1 == 1 ? proof[i] : acc;
            bytes32 b = loc & 1 == 1 ? acc : proof[i];
            acc = keccak2(a, b);
            loc = loc >> 1;
        }
        if (v != bytes32(0)) return acc;
        bytes32 res;
        if (nh == 0) {
            // Here we specify the leaf hash directly, since we're at height 0
            // There's no need for the leaf to be non-zero
            res = normal1;
        } else {
            // Since this is a branch, prove that its right side isn't 0,
            // as that wouldn't be normalized
            require(normal2 != zeros[nh], "right subtree cannot be zero");
            res = keccak2(normal1, normal2);
        }
        bytes32 acc2 = res;
        for (uint256 i = nh; i < proof.length - 1; i++) {
            acc2 = keccak2(acc2, zeros[i]);
        }
        require(acc2 == acc, "expected match");
        return res;
    }

    function getByte(bytes32 word, uint256 num) internal pure returns (uint256) {
        return (uint256(word) >> ((31 - num) * 8)) & 0xff;
    }

    function setByte(
        bytes32 word,
        uint256 num,
        uint256 b
    ) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = bytes1(uint8(b));
        return bytes32(bytes32FromArray(arr));
    }

    function setByte(
        bytes32 word,
        uint256 num,
        bytes1 b
    ) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = b;
        return bytes32(bytes32FromArray(arr));
    }

    function decode(
        bytes memory arr,
        bytes1 _start,
        bytes1 _end
    ) internal pure returns (bytes32[] memory) {
        uint256 len = uint256(uint8(_end) - uint8(_start));
        uint256 start = uint256(uint8(_start));
        bytes32[] memory res = new bytes32[](len);
        for (uint256 i = 0; i < len; i++) {
            res[i] = bytes32(bytes32FromArray(arr, (start + i) * 32));
        }
        return res;
    }

    struct BufferProof {
        bytes32[] proof1;
        bytes32[] nproof1;
        bytes32[] proof2;
        bytes32[] nproof2;
    }

    function decodeProof(bytes memory proof) internal pure returns (BufferProof memory) {
        bytes32[] memory proof1 = decode(proof, proof[0], proof[1]);
        bytes32[] memory nproof1 = decode(proof, proof[1], proof[2]);
        bytes32[] memory proof2 = decode(proof, proof[2], proof[3]);
        bytes32[] memory nproof2 = decode(proof, proof[3], proof[4]);
        return BufferProof(proof1, nproof1, proof2, nproof2);
    }

    function bytes32FromArray(bytes memory arr) internal pure returns (uint256) {
        uint256 res = 0;
        for (uint256 i = 0; i < arr.length; i++) {
            res = res << 8;
            res = res | uint256(uint8(arr[i]));
        }
        return res;
    }

    function bytes32FromArray(bytes memory arr, uint256 offset) internal pure returns (uint256) {
        uint256 res = 0;
        for (uint256 i = 0; i < 32; i++) {
            res = res << 8;
            res = res | uint256(uint8(arr[offset + i]));
        }
        return res;
    }

    function bytes32ToArray(bytes32 b) internal pure returns (bytes memory) {
        uint256 acc = uint256(b);
        bytes memory res = new bytes(32);
        for (uint256 i = 0; i < 32; i++) {
            res[31 - i] = bytes1(uint8(acc));
            acc = acc >> 8;
        }
        return res;
    }

    function getBuffer8(
        bytes32 buf,
        uint256 offset,
        BufferProof memory proof
    ) internal pure returns (uint256) {
        return getByte(get(buf, offset / 32, proof.proof1), offset % 32);
    }

    function checkBufferSize(
        bytes32 buf,
        uint256 offset,
        BufferProof memory proof
    ) internal pure returns (bool) {
        bytes32 w = get(buf, offset / 32, proof.proof1);
        for (uint256 i = offset % 32; i < 32; i++) {
            if (getByte(w, i) != 0) return false;
        }
        return checkSize(buf, offset / 32, proof.proof1);
    }

    function getBuffer64(
        bytes32 buf,
        uint256 offset,
        BufferProof memory proof
    ) internal pure returns (uint256) {
        bytes memory res = new bytes(8);
        bytes32 word = get(buf, offset / 32, proof.proof1);
        if ((offset % 32) + 8 > 32) {
            bytes32 word2 = get(buf, offset / 32 + 1, proof.proof2);
            for (uint256 i = 0; i < 8 - ((offset % 32) + 8 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, (offset % 32) + i)));
            }
            for (uint256 i = 8 - ((offset % 32) + 8 - 32); i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint256 i = 0; i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word, (offset % 32) + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function getBuffer256(
        bytes32 buf,
        uint256 offset,
        BufferProof memory proof
    ) internal pure returns (uint256) {
        bytes memory res = new bytes(32);
        bytes32 word = get(buf, offset / 32, proof.proof1);
        if ((offset % 32) + 32 > 32) {
            bytes32 word2 = get(buf, offset / 32 + 1, proof.proof2);
            for (uint256 i = 0; i < 32 - ((offset % 32) + 32 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, (offset % 32) + i)));
            }
            for (uint256 i = 32 - ((offset % 32) + 32 - 32); i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint256 i = 0; i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word, (offset % 32) + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function set(
        bytes32 buf,
        uint256 loc,
        bytes32 v,
        bytes32[] memory proof,
        bytes32[] memory nproof
    ) internal pure returns (bytes32) {
        require(nproof.length == 3, "BAD_NORMALIZATION_PROOF");
        return set(buf, loc, v, proof, uint256(nproof[0]), nproof[1], nproof[2]);
    }

    function setBuffer8(
        bytes32 buf,
        uint256 offset,
        uint256 b,
        BufferProof memory proof
    ) internal pure returns (bytes32) {
        bytes32 word = get(buf, offset / 32, proof.proof1);
        bytes32 nword = setByte(word, offset % 32, b);
        bytes32 res = set(buf, offset / 32, nword, proof.proof1, proof.nproof1);
        return res;
    }

    function setBuffer64(
        bytes32 buf,
        uint256 offset,
        uint256 val,
        BufferProof memory proof
    ) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(bytes32(val));
        bytes32 nword = get(buf, offset / 32, proof.proof1);
        if ((offset % 32) + 8 > 32) {
            for (uint256 i = 0; i < 8 - ((offset % 32) + 8 - 32); i++) {
                nword = setByte(nword, (offset + i) % 32, arr[i + 24]);
            }
            buf = set(buf, offset / 32, nword, proof.proof1, proof.nproof1);
            bytes32 nword2 = get(buf, offset / 32 + 1, proof.proof2);
            for (uint256 i = 8 - ((offset % 32) + 8 - 32); i < 8; i++) {
                nword2 = setByte(nword2, (offset + i) % 32, arr[i + 24]);
            }
            buf = set(buf, offset / 32 + 1, nword2, proof.proof2, proof.nproof2);
        } else {
            for (uint256 i = 0; i < 8; i++) {
                nword = setByte(nword, (offset % 32) + i, arr[i + 24]);
            }
            buf = set(buf, offset / 32, nword, proof.proof1, proof.nproof1);
        }
        return buf;
    }

    function parseProof(bytes memory proof)
        public
        pure
        returns (
            bytes32[] memory,
            bytes32[] memory,
            bytes32[] memory,
            bytes32[] memory
        )
    {
        BufferProof memory p = decodeProof(proof);
        return (p.proof1, p.nproof1, p.proof2, p.nproof2);
    }

    function setBuffer256(
        bytes32 buf,
        uint256 offset,
        uint256 val,
        BufferProof memory proof
    ) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(bytes32(val));
        bytes32 nword = get(buf, offset / 32, proof.proof1);
        if ((offset % 32) + 32 > 32) {
            for (uint256 i = 0; i < 32 - ((offset % 32) + 32 - 32); i++) {
                nword = setByte(nword, (offset % 32) + i, arr[i]);
            }
            buf = set(buf, offset / 32, nword, proof.proof1, proof.nproof1);
            bytes32 nword2 = get(buf, offset / 32 + 1, proof.proof2);
            for (uint256 i = 32 - ((offset % 32) + 32 - 32); i < 32; i++) {
                nword2 = setByte(nword2, (offset + i) % 32, arr[i]);
            }
            buf = set(buf, offset / 32 + 1, nword2, proof.proof2, proof.nproof2);
        } else {
            for (uint256 i = 0; i < 32; i++) {
                nword = setByte(nword, (offset % 32) + i, arr[i]);
            }
            buf = set(buf, offset / 32, nword, proof.proof1, proof.nproof1);
        }
        return buf;
    }

    function executeSendInsn(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal > SEND_SIZE_LIMIT || val2.intVal == 0) {
            handleOpcodeError(context);
            return;
        }

        if (context.offset == context.proof.length) {
            // If we didn't pass the message data, the buffer must have been longer than the length param passed
            require(
                !checkBufferSize(val1.bufferHash, val2.intVal, decodeProof(context.bufProof)),
                "BUF_LENGTH"
            );
            handleOpcodeError(context);
            return;
        }

        // We've passed more data in the proof which is the data of the send because it isn't too long
        uint256 dataStart = context.offset;
        uint256 dataLength = val2.intVal;
        bytes memory proof = context.proof;
        bytes32 bufferHash = Hashing.bytesToBufferHash(proof, dataStart, dataLength);
        require(val1.hash() == bufferHash, "WRONG_SEND");

        bytes32 dataHash;
        assembly {
            dataHash := keccak256(add(add(proof, 32), dataStart), dataLength)
        }

        context.sendAcc = keccak256(abi.encodePacked(context.sendAcc, dataHash));
    }

    function executeGetBuffer8(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= 1 << 64) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer8(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeGetBuffer64(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= (1 << 64) - 7) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer64(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeGetBuffer256(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= (1 << 64) - 31) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer256(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeSetBuffer8(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= 1 << 64 || val3.intVal >= 1 << 8) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res =
            setBuffer8(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    function executeSetBuffer64(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= (1 << 64) - 7 || val3.intVal >= 1 << 64) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res =
            setBuffer64(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    function executeSetBuffer256(AssertionContext memory context) internal pure {
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        Value.Data memory val1 = popVal(context.stack);
        if (!val2.isInt64() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal >= (1 << 64) - 31) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res =
            setBuffer256(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    function opInfo(uint256 opCode)
        internal
        pure
        override
        returns (
            uint256, // stack pops
            uint256, // auxstack pops
            uint64, // gas used
            function(AssertionContext memory) internal view // impl
        )
    {
        if (opCode == OP_GETBUFFER8) {
            return (2, 0, 10, executeGetBuffer8);
        } else if (opCode == OP_GETBUFFER64) {
            return (2, 0, 10, executeGetBuffer64);
        } else if (opCode == OP_GETBUFFER256) {
            return (2, 0, 10, executeGetBuffer256);
        } else if (opCode == OP_SETBUFFER8) {
            return (3, 0, 100, executeSetBuffer8);
        } else if (opCode == OP_SETBUFFER64) {
            return (3, 0, 100, executeSetBuffer64);
        } else if (opCode == OP_SETBUFFER256) {
            return (3, 0, 100, executeSetBuffer256);
        } else if (opCode == OP_SEND) {
            return (2, 0, 100, executeSendInsn);
        } else {
            revert("use another contract to handle other opcodes");
        }
    }
}
