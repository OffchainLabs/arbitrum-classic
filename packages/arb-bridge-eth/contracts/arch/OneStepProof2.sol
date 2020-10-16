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

import "./IOneStepProof.sol";
import "./Value.sol";
import "./Machine.sol";
import "../inbox/Messages.sol";
import "../libraries/Precompiles.sol";

// Originally forked from https://github.com/leapdao/solEVM-enforcer/tree/master

contract OneStepProof2 is IOneStepProof2 {
    using Machine for Machine.Data;
    using Hashing for Value.Data;
    using Value for Value.Data;

    uint256 private constant SEND_SIZE_LIMIT = 10000;

    uint256 private constant MAX_UINT256 = ((1 << 128) + 1) * ((1 << 128) - 1);
    uint256 private constant MAX_PAIRING_COUNT = 30;

    string private constant BAD_IMM_TYP = "BAD_IMM_TYP";
    string private constant NO_IMM = "NO_IMM";
    string private constant STACK_MISSING = "STACK_MISSING";
    string private constant AUX_MISSING = "AUX_MISSING";
    string private constant STACK_MANY = "STACK_MANY";
    string private constant AUX_MANY = "AUX_MANY";
    string private constant INBOX_VAL = "INBOX_VAL";

    function executeStep(
        bytes32 inboxAcc,
        bytes32 messagesAcc,
        bytes32 logsAcc,
        bytes calldata proof,
        bytes calldata bproof
    ) external view returns (uint64 gas, bytes32[5] memory fields) {
        AssertionContext memory context = initializeExecutionContext(
            inboxAcc,
            messagesAcc,
            logsAcc,
            proof,
            bproof
        );

        executeOp(context);

        return returnContext(context);
    }

    // fields
    // startMachineHash,
    // endMachineHash,
    // afterInboxHash,
    // afterMessagesHash,
    // afterLogsHash

    function returnContext(AssertionContext memory context)
        private
        pure
        returns (uint64 gas, bytes32[5] memory fields)
    {
        return (
            context.gas,
            [
                Machine.hash(context.startMachine),
                Machine.hash(context.afterMachine),
                context.inboxAcc,
                context.messageAcc,
                context.logAcc
            ]
        );
    }

    struct ValueStack {
        uint256 length;
        Value.Data[] values;
    }

    function popVal(ValueStack memory stack) private pure returns (Value.Data memory) {
        Value.Data memory val = stack.values[stack.length - 1];
        stack.length--;
        return val;
    }

    function pushVal(ValueStack memory stack, Value.Data memory val) private pure {
        stack.values[stack.length] = val;
        stack.length++;
    }

    struct AssertionContext {
        Machine.Data startMachine;
        Machine.Data afterMachine;
        bytes32 inboxAcc;
        bytes32 messageAcc;
        bytes32 logAcc;
        uint64 gas;
        Value.Data inboxMessage;
        bytes32 inboxMessageHash;
        ValueStack stack;
        ValueStack auxstack;
        bool hadImmediate;
        uint8 opcode;
        bytes proof;
        uint256 offset;
        // merkle proofs for buffer
        bytes bufProof;
    }

    function handleError(AssertionContext memory context) private pure {
        if (context.afterMachine.errHandlerHash == CODE_POINT_ERROR) {
            context.afterMachine.setErrorStop();
        } else {
            context.afterMachine.instructionStackHash = context.afterMachine.errHandlerHash;
        }
    }

    function deductGas(AssertionContext memory context, uint64 amount) private pure returns (bool) {
        context.gas += amount;
        if (context.afterMachine.arbGasRemaining < amount) {
            context.afterMachine.arbGasRemaining = MAX_UINT256;
            handleError(context);
            return true;
        } else {
            context.afterMachine.arbGasRemaining -= amount;
            return false;
        }
    }

    function handleOpcodeError(AssertionContext memory context) private pure {
        handleError(context);
        // Also clear the stack and auxstack
        context.stack.length = 0;
        context.auxstack.length = 0;
    }

    function initializeExecutionContext(
        bytes32 inboxAcc,
        bytes32 messagesAcc,
        bytes32 logsAcc,
        bytes memory proof,
        bytes memory bproof
    ) internal pure returns (AssertionContext memory) {
        uint8 stackCount = uint8(proof[0]);
        uint8 auxstackCount = uint8(proof[1]);
        uint256 offset = 2;

        // Leave some extra space for values pushed on the stack in the proofs
        Value.Data[] memory stackVals = new Value.Data[](stackCount + 4);
        Value.Data[] memory auxstackVals = new Value.Data[](auxstackCount + 4);
        for (uint256 i = 0; i < stackCount; i++) {
            (offset, stackVals[i]) = Marshaling.deserialize(proof, offset);
        }
        for (uint256 i = 0; i < auxstackCount; i++) {
            (offset, auxstackVals[i]) = Marshaling.deserialize(proof, offset);
        }
        Machine.Data memory mach;
        (offset, mach) = Machine.deserializeMachine(proof, offset);

        AssertionContext memory context = AssertionContext(
            mach,
            mach.clone(),
            inboxAcc,
            messagesAcc,
            logsAcc,
            0,
            Value.newEmptyTuple(),
            0,
            ValueStack(stackCount, stackVals),
            ValueStack(auxstackCount, auxstackVals),
            uint8(proof[offset]) == 1,
            uint8(proof[offset + 1]),
            proof,
            offset+2,
            bproof
        );

        uint8 immediate = uint8(proof[offset]);
        uint8 opCode = uint8(proof[offset + 1]);
        offset += 2;

        require(immediate == 0 || immediate == 1, BAD_IMM_TYP);
        Value.Data memory cp;
        if (immediate == 0) {
            cp = Value.newCodePoint(uint8(opCode), context.startMachine.instructionStackHash);
        } else {
            // If we have an immediate, there must be at least one stack value
            require(stackVals.length > 0, NO_IMM);
            cp = Value.newCodePoint(
                uint8(opCode),
                context.startMachine.instructionStackHash,
                stackVals[stackCount - 1]
            );
        }
        context.startMachine.instructionStackHash = cp.hash();

        // Add the stack and auxstack values to the start machine
        uint256 i = 0;
        for (i = 0; i < stackCount - immediate; i++) {
            context.startMachine.addDataStackValue(stackVals[i]);
        }
        for (i = 0; i < auxstackCount; i++) {
            context.startMachine.addAuxStackValue(auxstackVals[i]);
        }

        return context;
    }

    function executeOp(AssertionContext memory context) internal view {
        (
            uint256 dataPopCount,
            uint256 auxPopCount,
            uint64 gasCost,
            function(AssertionContext memory) internal view impl
        ) = opInfo(context.opcode);

        // Update end machine gas remaining before running opcode
        if (deductGas(context, gasCost)) {
            return;
        }

        if (context.stack.length < dataPopCount) {
            // If we have insufficient values, reject the proof unless the stack has been fully exhausted
            require(
                context.afterMachine.dataStack.hash() == Value.newEmptyTuple().hash(),
                STACK_MISSING
            );
            // If the stack is empty, the instruction underflowed so we have hit an error
            handleError(context);
            return;
        }

        if (context.auxstack.length < auxPopCount) {
            // If we have insufficient values, reject the proof unless the auxstack has been fully exhausted
            require(
                context.afterMachine.auxStack.hash() == Value.newEmptyTuple().hash(),
                AUX_MISSING
            );
            // If the auxstack is empty, the instruction underflowed so we have hit an error
            handleError(context);
            return;
        }

        // Require the prover to submit the minimal number of stack items
        require(
            ((dataPopCount > 0 || !context.hadImmediate) && context.stack.length == dataPopCount) ||
                (context.hadImmediate && dataPopCount == 0 && context.stack.length == 1),
            STACK_MANY
        );
        require(context.auxstack.length == auxPopCount, AUX_MANY);

        impl(context);

        // Add the stack and auxstack values to the start machine
        uint256 i = 0;

        for (i = 0; i < context.stack.length; i++) {
            context.afterMachine.addDataStackValue(context.stack.values[i]);
        }

        for (i = 0; i < context.auxstack.length; i++) {
            context.afterMachine.addAuxStackValue(context.auxstack.values[i]);
        }
    }

    /* solhint-disable no-inline-assembly */

    function executeErrorInsn(AssertionContext memory context) internal pure {
        handleOpcodeError(context);
    }

    function executeNewBuffer(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newBuffer(keccak256(abi.encodePacked(bytes32(0)))));
    }

    
    function makeZeros() internal pure returns (bytes32[] memory) {
        bytes32[] memory zeros = new bytes32[](64);
        zeros[0] = keccak1(0);
        for (uint i = 1; i < 64; i++) {
            zeros[i] = keccak2(zeros[i-1], zeros[i-1]);
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
    function get(bytes32 buf, uint loc, bytes32[] memory proof) internal pure returns (bytes32) {
        // empty tree is full of zeros
        if (proof.length == 0) {
            require(buf == keccak1(bytes32(0)), "expected empty buffer");
            return 0;
        }
        bytes32 acc = keccak1(proof[0]);
        for (uint i = 1; i < proof.length; i++) {
            if (loc & 1 == 1) acc = keccak2(proof[i], acc);
            else acc = keccak2(acc, proof[i]);
            loc = loc >> 1;
        }
        require(acc == buf, "expected correct root");
        // maybe it is a zero outside the actual tree
        if (loc > 0) return 0;
        return proof[0];
    }

    function calcHeight(uint loc) internal pure returns (uint) {
        if (loc == 0) return 0;
        else return 1+calcHeight(loc>>1);
    }

    function set(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, uint nh, bytes32 normal1, bytes32 normal2) internal pure returns (bytes32) {
        // three possibilities, the tree depth stays same, it becomes lower or it's extended
        bytes32 acc = keccak1(v);
        // check that the proof matches original
        get(buf, loc, proof);
        bytes32[] memory zeros = makeZeros();
        // extended
        if (loc > (proof.length << 2)) {
            if (v == 0) return buf;
            uint height = calcHeight(loc);
            // build the left branch
            for (uint i = proof.length; i < height-1; i++) {
                buf = keccak2(buf, zeros[i]);
            }
            for (uint i = 1; i < height-1; i++) {
                if (loc & 1 == 1) acc = keccak2(acc, zeros[i]);
                else acc = keccak2(zeros[i], acc);
                loc = loc >> 1;
            }
            return keccak2(buf, acc);
        }
        for (uint i = 1; i < proof.length; i++) {
            bytes32 a = loc & 1 == 1 ? proof[i] : acc;
            bytes32 b = loc & 1 == 1 ? acc : proof[i];
            acc = keccak2(a, b);
            loc = loc >> 1;
        }
        if (v != bytes32(0)) return acc;
        require(normal2 != zeros[nh], "right subtree cannot be zero");
        bytes32 res = keccak2(normal1, normal2);
        bytes32 acc2 = res;
        for (uint i = nh; i < proof.length; i++) {
            acc2 = keccak2(res, zeros[i]);
        }
        require(acc2 == acc, "expected match");
        return res;
    }

    function getByte(bytes32 word, uint256 num) internal pure returns (uint256) {
        return (uint256(word) >> ((31-num)*8)) & 0xff;
    }

    function setByte(bytes32 word, uint256 num, uint256 b) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = bytes1(uint8(b));
        return bytes32(bytes32FromArray(arr));
    }

    function setByte(bytes32 word, uint256 num, bytes1 b) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = b;
        return bytes32(bytes32FromArray(arr));
    }

    function decode(bytes memory arr, bytes1 _start, bytes1 _end) internal pure returns (bytes32[] memory) {
        uint len = uint(uint8(_end-_start));
        uint start = uint(uint8(_start));
        bytes32[] memory res = new bytes32[](len);
        for (uint i = 0; i < len; i++) {
            res[i] = bytes32(bytes32FromArray(arr, (start+i)*32));
        }
        return res;
    }

    struct BufferProof {
        bytes32[] proof1;
        bytes32[] proof2;
        bytes32[] nproof1;
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
        for (uint i = 0; i < arr.length; i++) {
            res = res << 8;
            res = res | uint256(uint8(arr[arr.length-1-i]));
        }
        return res;
    }

    function bytes32FromArray(bytes memory arr, uint offset) internal pure returns (uint256) {
        uint256 res = 0;
        for (uint i = 0; i < 32; i++) {
            res = res << 8;
            res = res | uint256(uint8(arr[offset+32-1-i]));
        }
        return res;
    }

    function bytes32ToArray(bytes32 b) internal pure returns (bytes memory arr) {
        uint256 acc = uint256(b);
        bytes memory res = new bytes(32);
        for (uint i = 0; i < arr.length; i++) {
            res[31-i] = bytes1(uint8(acc));
            acc = acc >> 8;
        }
        return res;
    }

    function getBuffer8(bytes32 buf, uint256 offset, BufferProof memory proof) internal pure returns (uint256) {
        return getByte(get(buf, offset/32, proof.proof1), offset%32);
    }

    function getBuffer64(bytes32 buf, uint256 offset, BufferProof memory proof) internal pure returns (uint256) {
        bytes memory res = new bytes(8);
        bytes32 word = get(buf, offset/32, proof.proof1); 
        if (offset%32 + 8 >= 32) {
            bytes32 word2 = get(buf, offset/32 + 1, proof.proof2);
            for (uint i = 0; i < 8 - (offset%32 + 8 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
            for (uint i = 8 - (offset%32 + 8 - 32); i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint i = 0; i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function getBuffer256(bytes32 buf, uint256 offset, BufferProof memory proof) internal pure returns (uint256) {
        bytes memory res = new bytes(32);
        bytes32 word = get(buf, offset/32, proof.proof1); 
        if (offset%32 + 32 >= 32) {
            bytes32 word2 = get(buf, offset/32 + 1, proof.proof2);
            for (uint i = 0; i < 32 - (offset%32 + 32 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
            for (uint i = 8 - (offset%32 + 32 - 32); i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint i = 0; i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function set(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, bytes32[] memory nproof) internal pure returns (bytes32) {
        return set(buf, loc, v, proof, uint256(nproof[0]), nproof[1], nproof[2]);
    }

    function setBuffer8(bytes32 buf, uint256 offset, uint256 b, BufferProof memory proof) internal pure returns (bytes32) {
        bytes32 word = get(buf, offset/32, proof.proof1);
        bytes32 nword = setByte(word, offset%32, b);
        return set(buf, offset/32, nword, proof.proof1, proof.nproof1);
    }

    function setBuffer64(bytes32 buf, uint256 offset, uint256 val, BufferProof memory proof) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(bytes32(val));
        bytes32 nword = get(buf, offset/32, proof.proof1);
        if (offset%32 + 8 >= 32) {
            bytes32 nword2 = get(buf, offset/32 + 1, proof.proof2); 
            for (uint i = 0; i < 8 - (offset%32 + 8 - 32); i++) {
                nword = setByte(nword, offset%32 + i, arr[i+24]);
            }
            for (uint i = 8 - (offset%32 + 8 - 32); i < 8; i++) {
                nword2 = setByte(nword2, (offset+i)%32, arr[i+24]);
                buf = set(buf, offset/32, nword, proof.proof1, proof.nproof1);
                buf = set(buf, offset/32 + 1, nword2, proof.proof2, proof.nproof2);
            }
        } else {
            for (uint i = 0; i < 8; i++) {
                nword = setByte(nword, offset%32 + i, arr[i+24]);
            }
            buf = set(buf, offset/32, nword, proof.proof1, proof.nproof1);
        }
        return buf;
    }

    function setBuffer256(bytes32 buf, uint256 offset, uint256 val, BufferProof memory proof) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(bytes32(val));
        bytes32 nword = get(buf, offset/32, proof.proof1);
        if (offset%32 + 32 >= 32) {
            bytes32 nword2 = get(buf, offset/32 + 1, proof.proof2); 
            for (uint i = 0; i < 32 - (offset%32 + 32 - 32); i++) {
                nword = setByte(nword, offset%32 + i, arr[i]);
            }
            for (uint i = 32 - (offset%32 + 32 - 32); i < 32; i++) {
                nword2 = setByte(nword2, (offset+i)%32, arr[i]);
                buf = set(buf, offset/32, nword, proof.proof1, proof.nproof1);
                buf = set(buf, offset/32 + 1, nword2, proof.proof2, proof.nproof2);
            }
        } else {
            for (uint i = 0; i < 32; i++) {
                nword = setByte(nword, offset%32 + i, arr[i]);
            }
            buf = set(buf, offset/32, nword, proof.proof1, proof.nproof1);
        }
        return buf;
    }

    function executeGetBuffer8(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer8(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeGetBuffer64(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer64(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeGetBuffer256(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        uint256 res = getBuffer256(val1.bufferHash, val2.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newInt(res));
    }

    function executeSetBuffer8(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res = setBuffer8(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    function executeSetBuffer64(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res = setBuffer64(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    function executeSetBuffer256(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val2.isInt() || !val3.isInt() || !val1.isBuffer()) {
            handleOpcodeError(context);
            return;
        }
        bytes32 res = setBuffer256(val1.bufferHash, val2.intVal, val3.intVal, decodeProof(context.bufProof));
        pushVal(context.stack, Value.newBuffer(res));
    }

    // Stop and arithmetic ops
    uint8 private constant OP_ADD = 0x01;
    uint8 private constant OP_MUL = 0x02;
    uint8 private constant OP_SUB = 0x03;
    uint8 private constant OP_DIV = 0x04;
    uint8 private constant OP_SDIV = 0x05;
    uint8 private constant OP_MOD = 0x06;
    uint8 private constant OP_SMOD = 0x07;
    uint8 private constant OP_ADDMOD = 0x08;
    uint8 private constant OP_MULMOD = 0x09;
    uint8 private constant OP_EXP = 0x0a;
    uint8 private constant OP_SIGNEXTEND = 0x0b;

    // Comparison & bitwise logic
    uint8 private constant OP_LT = 0x10;
    uint8 private constant OP_GT = 0x11;
    uint8 private constant OP_SLT = 0x12;
    uint8 private constant OP_SGT = 0x13;
    uint8 private constant OP_EQ = 0x14;
    uint8 private constant OP_ISZERO = 0x15;
    uint8 private constant OP_AND = 0x16;
    uint8 private constant OP_OR = 0x17;
    uint8 private constant OP_XOR = 0x18;
    uint8 private constant OP_NOT = 0x19;
    uint8 private constant OP_BYTE = 0x1a;
    uint8 private constant OP_SHL = 0x1b;
    uint8 private constant OP_SHR = 0x1c;
    uint8 private constant OP_SAR = 0x1d;

    // SHA3
    uint8 private constant OP_HASH = 0x20;
    uint8 private constant OP_TYPE = 0x21;
    uint8 private constant OP_ETHHASH2 = 0x22;
    uint8 private constant OP_KECCAK_F = 0x23;
    uint8 private constant OP_SHA256_F = 0x24;

    // Stack, Memory, Storage and Flow Operations
    uint8 private constant OP_POP = 0x30;
    uint8 private constant OP_SPUSH = 0x31;
    uint8 private constant OP_RPUSH = 0x32;
    uint8 private constant OP_RSET = 0x33;
    uint8 private constant OP_JUMP = 0x34;
    uint8 private constant OP_CJUMP = 0x35;
    uint8 private constant OP_STACKEMPTY = 0x36;
    uint8 private constant OP_PCPUSH = 0x37;
    uint8 private constant OP_AUXPUSH = 0x38;
    uint8 private constant OP_AUXPOP = 0x39;
    uint8 private constant OP_AUXSTACKEMPTY = 0x3a;
    uint8 private constant OP_NOP = 0x3b;
    uint8 private constant OP_ERRPUSH = 0x3c;
    uint8 private constant OP_ERRSET = 0x3d;

    // Duplication and Exchange operations
    uint8 private constant OP_DUP0 = 0x40;
    uint8 private constant OP_DUP1 = 0x41;
    uint8 private constant OP_DUP2 = 0x42;
    uint8 private constant OP_SWAP1 = 0x43;
    uint8 private constant OP_SWAP2 = 0x44;

    // Tuple operations
    uint8 private constant OP_TGET = 0x50;
    uint8 private constant OP_TSET = 0x51;
    uint8 private constant OP_TLEN = 0x52;
    uint8 private constant OP_XGET = 0x53;
    uint8 private constant OP_XSET = 0x54;

    // Logging operations
    uint8 private constant OP_BREAKPOINT = 0x60;
    uint8 private constant OP_LOG = 0x61;

    // System operations
    uint8 private constant OP_SEND = 0x70;
    uint8 private constant OP_INBOX_PEEK = 0x71;
    uint8 private constant OP_INBOX = 0x72;
    uint8 private constant OP_ERROR = 0x73;
    uint8 private constant OP_STOP = 0x74;
    uint8 private constant OP_SETGAS = 0x75;
    uint8 private constant OP_PUSHGAS = 0x76;
    uint8 private constant OP_ERR_CODE_POINT = 0x77;
    uint8 private constant OP_PUSH_INSN = 0x78;
    uint8 private constant OP_PUSH_INSN_IMM = 0x79;
    // uint8 private constant OP_OPEN_INSN = 0x7a;
    uint8 private constant OP_SIDELOAD = 0x7b;

    uint8 private constant OP_ECRECOVER = 0x80;
    uint8 private constant OP_ECADD = 0x81;
    uint8 private constant OP_ECMUL = 0x82;
    uint8 private constant OP_ECPAIRING = 0x83;

    // Buffer operations
    uint8 private constant OP_NEWBUFFER = 0xa0;
    uint8 private constant OP_GETBUFFER8 = 0xa1;
    uint8 private constant OP_GETBUFFER64 = 0xa2;
    uint8 private constant OP_GETBUFFER256 = 0xa3;
    uint8 private constant OP_SETBUFFER8 = 0xa4;
    uint8 private constant OP_SETBUFFER64 = 0xa5;
    uint8 private constant OP_SETBUFFER256 = 0xa6;

    uint64 private constant EC_PAIRING_POINT_GAS_COST = 500000;

    uint8 private constant CODE_POINT_TYPECODE = 1;
    bytes32 private constant CODE_POINT_ERROR = keccak256(
        abi.encodePacked(CODE_POINT_TYPECODE, uint8(0), bytes32(0))
    );

    function opInfo(uint256 opCode)
        private
        pure
        returns (
            uint256, // stack pops
            uint256, // auxstack pops
            uint64, // gas used
            function(AssertionContext memory) internal view // impl
        )
    {
       
        if (opCode == OP_NEWBUFFER) {
            return (1, 0, 1, executeNewBuffer);
        } else if (opCode == OP_SETBUFFER8) {
            return (2, 0, 10, executeGetBuffer8);
        } else if (opCode == OP_GETBUFFER64) {
            return (2, 0, 10, executeGetBuffer64);
        } else if (opCode == OP_GETBUFFER8) {
            return (2, 0, 10, executeGetBuffer256);
        } else if (opCode == OP_SETBUFFER8) {
            return (3, 0, 100, executeSetBuffer8);
        } else if (opCode == OP_SETBUFFER64) {
            return (3, 0, 100, executeSetBuffer64);
        } else if (opCode == OP_SETBUFFER8) {
            return (3, 0, 100, executeSetBuffer256);
        } else {
            return (0, 0, 0, executeErrorInsn);
        }
    }
}
