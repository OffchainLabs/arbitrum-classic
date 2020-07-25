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
import "./Machine.sol";
import "../libraries/Keccak.sol";

// Sourced from https://github.com/leapdao/solEVM-enforcer/tree/master

library OneStepProof {
    using Machine for Machine.Data;
    using Hashing for Value.Data;
    using Value for Value.Data;

    uint256 private constant SEND_SIZE_LIMIT = 10000;

    uint256 private constant MAX_UINT256 = ((1 << 128) + 1) * ((1 << 128) - 1);

    struct ValidateProofData {
        bytes32 beforeHash;
        Value.Data beforeInbox;
        bool didInboxInsn;
        bytes32 firstMessage;
        bytes32 firstLog;
        uint64 gas;
        bytes proof;
    }

    struct AssertionContext {
        Machine.Data machine;
        Value.Data inbox;
        bool didInboxInsn;
        bytes32 messageAcc;
        bytes32 logAcc;
        uint64 gas;
    }

    function validateProof(
        bytes32 beforeHash,
        bytes32 beforeInbox,
        uint256 beforeInboxValueSize,
        bool didInboxInsn,
        bytes32 firstMessage,
        bytes32 firstLog,
        uint64 gas,
        bytes memory proof
    ) internal pure returns (AssertionContext memory) {
        return
            checkProof(
                ValidateProofData(
                    beforeHash,
                    Value.newTuplePreImage(beforeInbox, beforeInboxValueSize),
                    didInboxInsn,
                    firstMessage,
                    firstLog,
                    gas,
                    proof
                )
            );
    }

    /* solhint-disable no-inline-assembly */

    // Arithmetic

    function executeAddInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := add(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeMulInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := mul(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSubInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := sub(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeDivInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := div(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSdivInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := sdiv(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeModInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := mod(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSmodInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := smod(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeAddmodInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 m = vals[2].intVal;
        if (m == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := addmod(a, b, m)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeMulmodInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 m = vals[2].intVal;
        if (m == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := mulmod(a, b, m)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeExpInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := exp(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    // Comparison

    function executeLtInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := lt(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeGtInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := gt(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSltInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := slt(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSgtInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := sgt(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeEqInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newBoolean(vals[0].hash() == vals[1].hash())
        );
        return true;
    }

    function executeIszeroInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt()) {
            context.machine.addDataStackInt(0);
        } else {
            uint256 a = vals[0].intVal;
            uint256 c;
            assembly {
                c := iszero(a)
            }
            context.machine.addDataStackInt(c);
        }
        return true;
    }

    function executeAndInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := and(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeOrInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := or(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeXorInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        uint256 c;
        assembly {
            c := xor(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeNotInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 c;
        assembly {
            c := not(a)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeByteInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 x = vals[0].intVal;
        uint256 n = vals[1].intVal;
        uint256 c;
        assembly {
            c := byte(n, x)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    function executeSignextendInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 b = vals[0].intVal;
        uint256 a = vals[1].intVal;
        uint256 c;
        assembly {
            c := signextend(a, b)
        }
        context.machine.addDataStackInt(c);
        return true;
    }

    /* solhint-enable no-inline-assembly */

    // Hash

    function executeSha3Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackInt(uint256(vals[0].hash()));
        return true;
    }

    function executeTypeInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[0].typeCodeVal());
        return true;
    }

    function executeEthhash2Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isInt()) {
            return false;
        }
        uint256 a = vals[0].intVal;
        uint256 b = vals[1].intVal;
        bytes32 res = keccak256(abi.encodePacked(a, b));
        context.machine.addDataStackInt(uint256(res));
        return true;
    }

    function executeKeccakFInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isTuple()) {
            return false;
        }
        Value.Data[] memory values = vals[0].tupleVal;
        if (values.length != 7) {
            return false;
        }
        for (uint256 i = 0; i < 7; i++) {
            if (!values[i].isInt()) {
                return false;
            }
        }
        uint256[25] memory data;
        for (uint256 i = 0; i < 25; i++) {
            data[i] = uint256(uint64(values[i / 4].intVal));
            values[i / 4].intVal >>= 64;
        }

        data = Keccak.keccak_f(data);

        Value.Data[] memory outValues = new Value.Data[](7);
        for (uint256 i = 0; i < 7; i++) {
            outValues[i] = Value.newInt(0);
        }

        for (uint256 i = 0; i < 25; i++) {
            outValues[i / 4].intVal |= data[i] << ((i % 4) * 64);
        }

        context.machine.addDataStackValue(Value.newTuple(outValues));
        return true;
    }

    // Stack ops

    function executePopInsn(AssertionContext memory, Value.Data[] memory)
        internal
        pure
        returns (bool)
    {
        return true;
    }

    function executeSpushInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(context.machine.staticVal);
        return true;
    }

    function executeRpushInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(context.machine.registerVal);
        return true;
    }

    function executeRsetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.registerVal = vals[0];
        return true;
    }

    function executeJumpInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isCodePoint()) {
            return false;
        }
        context.machine.instructionStackHash = vals[0].hash();
        return true;
    }

    function executeCjumpInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isCodePoint()) {
            return false;
        }
        if (!vals[1].isInt()) {
            return false;
        }
        if (vals[1].intVal != 0) {
            context.machine.instructionStackHash = vals[0].hash();
        }
        return true;
    }

    function executeStackemptyInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newBoolean(
                context.machine.dataStack.hash() == Value.newEmptyTuple().hash()
            )
        );
        return true;
    }

    function executePcpushInsn(
        Machine.Data memory startMachine,
        AssertionContext memory context
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newHashedValue(startMachine.instructionStackHash, 1)
        );
        return true;
    }

    function executeAuxpushInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addAuxStackValue(vals[0]);
        return true;
    }

    function executeAuxstackemptyInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newBoolean(
                context.machine.auxStack.hash() == Value.newEmptyTuple().hash()
            )
        );
        return true;
    }

    function executeErrpushInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newHashedValue(context.machine.errHandlerHash, 1)
        );
        return true;
    }

    function executeErrsetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isCodePoint()) {
            return false;
        }
        context.machine.errHandlerHash = vals[0].hash();
        return true;
    }

    // Dup ops

    function executeDup0Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[0]);
        context.machine.addDataStackValue(vals[0]);
        return true;
    }

    function executeDup1Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[1]);
        context.machine.addDataStackValue(vals[0]);
        context.machine.addDataStackValue(vals[1]);
        return true;
    }

    function executeDup2Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[2]);
        context.machine.addDataStackValue(vals[1]);
        context.machine.addDataStackValue(vals[0]);
        context.machine.addDataStackValue(vals[2]);
        return true;
    }

    // Swap ops

    function executeSwap1Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[0]);
        context.machine.addDataStackValue(vals[1]);
        return true;
    }

    function executeSwap2Insn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(vals[0]);
        context.machine.addDataStackValue(vals[1]);
        context.machine.addDataStackValue(vals[2]);
        return true;
    }

    // Tuple ops

    function executeTgetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !vals[1].isTuple()) {
            return false;
        }

        if (vals[0].intVal >= vals[1].valLength()) {
            return false;
        }

        context.machine.addDataStackValue(vals[1].tupleVal[vals[0].intVal]);
        return true;
    }

    function executeTsetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[1].isTuple() || !vals[0].isInt()) {
            return false;
        }

        if (vals[0].intVal >= vals[1].valLength()) {
            return false;
        }
        Value.Data[] memory tupleVals = vals[1].tupleVal;
        tupleVals[vals[0].intVal] = vals[2];

        context.machine.addDataStackValue(Value.newTuple(tupleVals));
        return true;
    }

    function executeTlenInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isTuple()) {
            return false;
        }
        context.machine.addDataStackInt(vals[0].valLength());
        return true;
    }

    function executeXgetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals,
        Value.Data memory auxVal
    ) internal pure returns (bool) {
        if (!vals[0].isInt() || !auxVal.isTuple()) {
            return false;
        }

        if (vals[0].intVal >= auxVal.valLength()) {
            return false;
        }

        context.machine.addAuxStackValue(auxVal);
        context.machine.addDataStackValue(auxVal.tupleVal[vals[0].intVal]);
        return true;
    }

    function executeXsetInsn(
        AssertionContext memory context,
        Value.Data[] memory vals,
        Value.Data memory auxVal
    ) internal pure returns (bool) {
        if (!auxVal.isTuple() || !vals[0].isInt()) {
            return false;
        }

        if (vals[0].intVal >= auxVal.valLength()) {
            return false;
        }
        Value.Data[] memory tupleVals = auxVal.tupleVal;
        tupleVals[vals[0].intVal] = vals[1];

        context.machine.addAuxStackValue(Value.newTuple(tupleVals));
        return true;
    }

    // Logging

    function executeBreakpointInsn(AssertionContext memory, Value.Data[] memory)
        internal
        pure
        returns (bool)
    {
        return true;
    }

    function executeLogInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        context.logAcc = keccak256(
            abi.encodePacked(context.logAcc, vals[0].hash())
        );
        return true;
    }

    // System operations

    function executeSendInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (vals[0].size > SEND_SIZE_LIMIT) {
            return false;
        }
        if (!vals[0].isValidTypeForSend()) {
            return false;
        }
        context.messageAcc = keccak256(
            abi.encodePacked(context.messageAcc, vals[0].hash())
        );
        return true;
    }

    function executeInboxInsn(
        AssertionContext memory context,
        Value.Data memory beforeInbox
    ) internal pure returns (bool) {
        require(
            beforeInbox.hash() != Value.newEmptyTuple().hash(),
            "Inbox instruction was blocked"
        );
        context.machine.addDataStackValue(beforeInbox);
        return true;
    }

    function executeSetGasInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt()) {
            return false;
        }
        context.machine.arbGasRemaining = vals[0].intVal;
        return true;
    }

    function executePushGasInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackInt(context.machine.arbGasRemaining);
        return true;
    }

    function executeErrCodePointInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        context.machine.addDataStackValue(
            Value.newHashedValue(CODE_POINT_ERROR, 1)
        );
        return true;
    }

    function executePushInsnInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt()) {
            return false;
        }
        if (!vals[1].isCodePoint()) {
            return false;
        }
        context.machine.addDataStackValue(
            Value.newCodePoint(uint8(vals[0].intVal), vals[1].hash())
        );
        return true;
    }

    function executePushInsnImmInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (!vals[0].isInt()) {
            return false;
        }
        if (!vals[2].isCodePoint()) {
            return false;
        }
        context.machine.addDataStackValue(
            Value.newCodePoint(uint8(vals[0].intVal), vals[2].hash(), vals[1])
        );
        return true;
    }

    function executeSideloadInsn(
        AssertionContext memory context,
        Value.Data[] memory
    ) internal pure returns (bool) {
        Value.Data[] memory values = new Value.Data[](0);
        context.machine.addDataStackValue(Value.newTuple(values));
        return true;
    }

    function executeECRecoverInsn(
        AssertionContext memory context,
        Value.Data[] memory vals
    ) internal pure returns (bool) {
        if (
            !vals[0].isInt() ||
            !vals[1].isInt() ||
            !vals[2].isInt() ||
            !vals[3].isInt()
        ) {
            return false;
        }
        bytes32 r = bytes32(vals[0].intVal);
        bytes32 s = bytes32(vals[1].intVal);
        if (vals[2].intVal != 0 && vals[2].intVal != 1) {
            context.machine.addDataStackInt(0);
            return true;
        }
        uint8 v = uint8(vals[2].intVal) + 27;
        bytes32 message = bytes32(vals[3].intVal);
        address ret = ecrecover(message, v, r, s);
        context.machine.addDataStackInt(uint256(ret));
        return true;
    }

    // Stop and arithmetic ops
    uint8 internal constant OP_ADD = 0x01;
    uint8 internal constant OP_MUL = 0x02;
    uint8 internal constant OP_SUB = 0x03;
    uint8 internal constant OP_DIV = 0x04;
    uint8 internal constant OP_SDIV = 0x05;
    uint8 internal constant OP_MOD = 0x06;
    uint8 internal constant OP_SMOD = 0x07;
    uint8 internal constant OP_ADDMOD = 0x08;
    uint8 internal constant OP_MULMOD = 0x09;
    uint8 internal constant OP_EXP = 0x0a;

    // Comparison & bitwise logic
    uint8 internal constant OP_LT = 0x10;
    uint8 internal constant OP_GT = 0x11;
    uint8 internal constant OP_SLT = 0x12;
    uint8 internal constant OP_SGT = 0x13;
    uint8 internal constant OP_EQ = 0x14;
    uint8 internal constant OP_ISZERO = 0x15;
    uint8 internal constant OP_AND = 0x16;
    uint8 internal constant OP_OR = 0x17;
    uint8 internal constant OP_XOR = 0x18;
    uint8 internal constant OP_NOT = 0x19;
    uint8 internal constant OP_BYTE = 0x1a;
    uint8 internal constant OP_SIGNEXTEND = 0x1b;

    // SHA3
    uint8 internal constant OP_SHA3 = 0x20;
    uint8 internal constant OP_TYPE = 0x21;
    uint8 internal constant OP_ETHHASH2 = 0x22;
    uint8 internal constant OP_KECCAK_F = 0x23;

    // Stack, Memory, Storage and Flow Operations
    uint8 internal constant OP_POP = 0x30;
    uint8 internal constant OP_SPUSH = 0x31;
    uint8 internal constant OP_RPUSH = 0x32;
    uint8 internal constant OP_RSET = 0x33;
    uint8 internal constant OP_JUMP = 0x34;
    uint8 internal constant OP_CJUMP = 0x35;
    uint8 internal constant OP_STACKEMPTY = 0x36;
    uint8 internal constant OP_PCPUSH = 0x37;
    uint8 internal constant OP_AUXPUSH = 0x38;
    uint8 internal constant OP_AUXPOP = 0x39;
    uint8 internal constant OP_AUXSTACKEMPTY = 0x3a;
    uint8 internal constant OP_NOP = 0x3b;
    uint8 internal constant OP_ERRPUSH = 0x3c;
    uint8 internal constant OP_ERRSET = 0x3d;

    // Duplication and Exchange operations
    uint8 internal constant OP_DUP0 = 0x40;
    uint8 internal constant OP_DUP1 = 0x41;
    uint8 internal constant OP_DUP2 = 0x42;
    uint8 internal constant OP_SWAP1 = 0x43;
    uint8 internal constant OP_SWAP2 = 0x44;

    // Tuple opertations
    uint8 internal constant OP_TGET = 0x50;
    uint8 internal constant OP_TSET = 0x51;
    uint8 internal constant OP_TLEN = 0x52;
    uint8 internal constant OP_XGET = 0x53;
    uint8 internal constant OP_XSET = 0x54;

    // Logging opertations
    uint8 internal constant OP_BREAKPOINT = 0x60;
    uint8 internal constant OP_LOG = 0x61;

    // System operations
    uint8 internal constant OP_SEND = 0x70;
    uint8 internal constant OP_INBOX = 0x72;
    uint8 internal constant OP_ERROR = 0x73;
    uint8 internal constant OP_STOP = 0x74;
    uint8 internal constant OP_SETGAS = 0x75;
    uint8 internal constant OP_PUSHGAS = 0x76;
    uint8 internal constant OP_ERR_CODE_POINT = 0x77;
    uint8 internal constant OP_PUSH_INSN = 0x78;
    uint8 internal constant OP_PUSH_INSN_IMM = 0x79;
    // uint8 internal constant OP_OPEN_INSN = 0x7a;
    uint8 internal constant OP_SIDELOAD = 0x7b;

    uint8 internal constant OP_ECRECOVER = 0x80;

    // opInfo returns data stack pop count and gas used
    function opInfo(uint256 opCode) internal pure returns (uint256, uint256) {
        if (opCode == OP_ADD) {
            return (2, 3);
        } else if (opCode == OP_MUL) {
            return (2, 3);
        } else if (opCode == OP_SUB) {
            return (2, 3);
        } else if (opCode == OP_DIV) {
            return (2, 4);
        } else if (opCode == OP_SDIV) {
            return (2, 7);
        } else if (opCode == OP_MOD) {
            return (2, 4);
        } else if (opCode == OP_SMOD) {
            return (2, 7);
        } else if (opCode == OP_ADDMOD) {
            return (3, 4);
        } else if (opCode == OP_MULMOD) {
            return (3, 4);
        } else if (opCode == OP_EXP) {
            return (2, 25);
        } else if (opCode == OP_LT) {
            return (2, 2);
        } else if (opCode == OP_GT) {
            return (2, 2);
        } else if (opCode == OP_SLT) {
            return (2, 2);
        } else if (opCode == OP_SGT) {
            return (2, 2);
        } else if (opCode == OP_EQ) {
            return (2, 2);
        } else if (opCode == OP_ISZERO) {
            return (1, 1);
        } else if (opCode == OP_AND) {
            return (2, 2);
        } else if (opCode == OP_OR) {
            return (2, 2);
        } else if (opCode == OP_XOR) {
            return (2, 2);
        } else if (opCode == OP_NOT) {
            return (1, 1);
        } else if (opCode == OP_BYTE) {
            return (2, 4);
        } else if (opCode == OP_SIGNEXTEND) {
            return (2, 7);
        } else if (opCode == OP_SHA3) {
            return (1, 7);
        } else if (opCode == OP_TYPE) {
            return (1, 3);
        } else if (opCode == OP_ETHHASH2) {
            return (2, 8);
        } else if (opCode == OP_KECCAK_F) {
            return (1, 800);
        } else if (opCode == OP_POP) {
            return (1, 1);
        } else if (opCode == OP_SPUSH) {
            return (0, 1);
        } else if (opCode == OP_RPUSH) {
            return (0, 1);
        } else if (opCode == OP_RSET) {
            return (1, 2);
        } else if (opCode == OP_JUMP) {
            return (1, 4);
        } else if (opCode == OP_CJUMP) {
            return (2, 4);
        } else if (opCode == OP_STACKEMPTY) {
            return (0, 2);
        } else if (opCode == OP_PCPUSH) {
            return (0, 1);
        } else if (opCode == OP_AUXPUSH) {
            return (1, 1);
        } else if (opCode == OP_AUXPOP) {
            return (0, 1);
        } else if (opCode == OP_AUXSTACKEMPTY) {
            return (0, 2);
        } else if (opCode == OP_NOP) {
            return (0, 1);
        } else if (opCode == OP_ERRPUSH) {
            return (0, 1);
        } else if (opCode == OP_ERRSET) {
            return (1, 1);
        } else if (opCode == OP_DUP0) {
            return (1, 1);
        } else if (opCode == OP_DUP1) {
            return (2, 1);
        } else if (opCode == OP_DUP2) {
            return (3, 1);
        } else if (opCode == OP_SWAP1) {
            return (2, 1);
        } else if (opCode == OP_SWAP2) {
            return (3, 1);
        } else if (opCode == OP_TGET) {
            return (2, 2);
        } else if (opCode == OP_TSET) {
            return (3, 40);
        } else if (opCode == OP_TLEN) {
            return (1, 2);
        } else if (opCode == OP_XGET) {
            return (1, 3);
        } else if (opCode == OP_XSET) {
            return (2, 41);
        } else if (opCode == OP_BREAKPOINT) {
            return (0, 100);
        } else if (opCode == OP_LOG) {
            return (1, 100);
        } else if (opCode == OP_SEND) {
            return (1, 100);
        } else if (opCode == OP_INBOX) {
            return (0, 40);
        } else if (opCode == OP_ERROR) {
            return (0, 5);
        } else if (opCode == OP_STOP) {
            return (0, 10);
        } else if (opCode == OP_SETGAS) {
            return (1, 0);
        } else if (opCode == OP_PUSHGAS) {
            return (0, 1);
        } else if (opCode == OP_ERR_CODE_POINT) {
            return (0, 25);
        } else if (opCode == OP_PUSH_INSN) {
            return (2, 25);
        } else if (opCode == OP_PUSH_INSN_IMM) {
            return (3, 25);
        } else if (opCode == OP_SIDELOAD) {
            return (0, 10);
        } else if (opCode == OP_ECRECOVER) {
            return (4, 20000);
        } else {
            require(false, "Invalid opcode: opInfo()");
        }
    }

    function loadMachine(ValidateProofData memory _data)
        internal
        pure
        returns (
            uint8 opCode,
            uint256 gasCost,
            Value.Data[] memory stackVals,
            Machine.Data memory startMachine,
            AssertionContext memory context,
            uint256 offset
        )
    {
        startMachine.setExtensive();
        (offset, startMachine) = Machine.deserializeMachine(
            _data.proof,
            offset
        );

        context = AssertionContext(
            startMachine.clone(),
            _data.beforeInbox,
            _data.didInboxInsn,
            _data.firstMessage,
            _data.firstLog,
            _data.gas
        );

        uint8 immediate = uint8(_data.proof[offset]);
        opCode = uint8(_data.proof[offset + 1]);
        uint256 popCount;
        (popCount, gasCost) = opInfo(opCode);
        stackVals = new Value.Data[](popCount);
        offset += 2;

        require(
            immediate == 0 || immediate == 1,
            "Proof had bad operation type"
        );
        if (immediate == 0) {
            startMachine.instructionStackHash = Value
                .newCodePoint(uint8(opCode), startMachine.instructionStackHash)
                .hash();
        } else {
            Value.Data memory immediateVal;
            (offset, immediateVal) = Marshaling.deserialize(
                _data.proof,
                offset
            );
            if (popCount > 0) {
                stackVals[0] = immediateVal;
            } else {
                context.machine.addDataStackValue(immediateVal);
            }

            startMachine.instructionStackHash = Value
                .newCodePoint(
                uint8(opCode),
                startMachine
                    .instructionStackHash,
                immediateVal
            )
                .hash();
        }

        uint256 i = 0;
        for (i = immediate; i < popCount; i++) {
            (offset, stackVals[i]) = Marshaling.deserialize(
                _data.proof,
                offset
            );
        }
        if (stackVals.length > 0) {
            for (i = 0; i < stackVals.length - immediate; i++) {
                startMachine.addDataStackValue(
                    stackVals[stackVals.length - 1 - i]
                );
            }
        }
        return (opCode, gasCost, stackVals, startMachine, context, offset);
    }

    uint8 private constant CODE_POINT_TYPECODE = 1;
    bytes32 private constant CODE_POINT_ERROR = keccak256(
        abi.encodePacked(CODE_POINT_TYPECODE, uint8(0), bytes32(0))
    );

    function checkProof(ValidateProofData memory _data)
        internal
        pure
        returns (AssertionContext memory)
    {
        uint8 opCode;
        uint256 gasCost;
        uint256 offset;
        Value.Data[] memory stackVals;
        Machine.Data memory startMachine;
        AssertionContext memory context;
        (
            opCode,
            gasCost,
            stackVals,
            startMachine,
            context,
            offset
        ) = loadMachine(_data);

        bool correct = true;
        require(_data.gas == gasCost, "Invalid gas in proof");
        require(
            (_data.didInboxInsn && opCode == OP_INBOX) ||
                (!_data.didInboxInsn && opCode != OP_INBOX),
            "Invalid didInboxInsn claim"
        );
        // Update end machine gas remaining before running opcode
        // No need to overflow check since the check for whether we
        // have sufficient gas fixes the overflow case
        context.machine.arbGasRemaining =
            context.machine.arbGasRemaining -
            gasCost;

        if (startMachine.arbGasRemaining < gasCost) {
            context.machine.arbGasRemaining = MAX_UINT256;
            correct = false;
        } else if (opCode == OP_ADD) {
            correct = executeAddInsn(context, stackVals);
        } else if (opCode == OP_MUL) {
            correct = executeMulInsn(context, stackVals);
        } else if (opCode == OP_SUB) {
            correct = executeSubInsn(context, stackVals);
        } else if (opCode == OP_DIV) {
            correct = executeDivInsn(context, stackVals);
        } else if (opCode == OP_SDIV) {
            correct = executeSdivInsn(context, stackVals);
        } else if (opCode == OP_MOD) {
            correct = executeModInsn(context, stackVals);
        } else if (opCode == OP_SMOD) {
            correct = executeSmodInsn(context, stackVals);
        } else if (opCode == OP_ADDMOD) {
            correct = executeAddmodInsn(context, stackVals);
        } else if (opCode == OP_MULMOD) {
            correct = executeMulmodInsn(context, stackVals);
        } else if (opCode == OP_EXP) {
            correct = executeExpInsn(context, stackVals);
        } else if (opCode == OP_LT) {
            correct = executeLtInsn(context, stackVals);
        } else if (opCode == OP_GT) {
            correct = executeGtInsn(context, stackVals);
        } else if (opCode == OP_SLT) {
            correct = executeSltInsn(context, stackVals);
        } else if (opCode == OP_SGT) {
            correct = executeSgtInsn(context, stackVals);
        } else if (opCode == OP_EQ) {
            correct = executeEqInsn(context, stackVals);
        } else if (opCode == OP_ISZERO) {
            correct = executeIszeroInsn(context, stackVals);
        } else if (opCode == OP_AND) {
            correct = executeAndInsn(context, stackVals);
        } else if (opCode == OP_OR) {
            correct = executeOrInsn(context, stackVals);
        } else if (opCode == OP_XOR) {
            correct = executeXorInsn(context, stackVals);
        } else if (opCode == OP_NOT) {
            correct = executeNotInsn(context, stackVals);
        } else if (opCode == OP_BYTE) {
            correct = executeByteInsn(context, stackVals);
        } else if (opCode == OP_SIGNEXTEND) {
            correct = executeSignextendInsn(context, stackVals);
        } else if (opCode == OP_SHA3) {
            correct = executeSha3Insn(context, stackVals);
        } else if (opCode == OP_TYPE) {
            correct = executeTypeInsn(context, stackVals);
        } else if (opCode == OP_ETHHASH2) {
            correct = executeEthhash2Insn(context, stackVals);
        } else if (opCode == OP_KECCAK_F) {
            correct = executeKeccakFInsn(context, stackVals);
        } else if (opCode == OP_POP) {
            correct = executePopInsn(context, stackVals);
        } else if (opCode == OP_SPUSH) {
            correct = executeSpushInsn(context, stackVals);
        } else if (opCode == OP_RPUSH) {
            correct = executeRpushInsn(context, stackVals);
        } else if (opCode == OP_RSET) {
            correct = executeRsetInsn(context, stackVals);
        } else if (opCode == OP_JUMP) {
            correct = executeJumpInsn(context, stackVals);
        } else if (opCode == OP_CJUMP) {
            correct = executeCjumpInsn(context, stackVals);
        } else if (opCode == OP_STACKEMPTY) {
            correct = executeStackemptyInsn(context, stackVals);
        } else if (opCode == OP_PCPUSH) {
            correct = executePcpushInsn(startMachine, context);
        } else if (opCode == OP_AUXPUSH) {
            correct = executeAuxpushInsn(context, stackVals);
        } else if (opCode == OP_AUXPOP) {
            Value.Data memory auxVal;
            (offset, auxVal) = Marshaling.deserialize(_data.proof, offset);
            startMachine.addAuxStackValue(auxVal);
            context.machine.addDataStackValue(auxVal);
        } else if (opCode == OP_AUXSTACKEMPTY) {
            correct = executeAuxstackemptyInsn(context, stackVals);
        } else if (opCode == OP_NOP) {
            correct = true;
        } else if (opCode == OP_ERRPUSH) {
            correct = executeErrpushInsn(context, stackVals);
        } else if (opCode == OP_ERRSET) {
            correct = executeErrsetInsn(context, stackVals);
        } else if (opCode == OP_DUP0) {
            correct = executeDup0Insn(context, stackVals);
        } else if (opCode == OP_DUP1) {
            correct = executeDup1Insn(context, stackVals);
        } else if (opCode == OP_DUP2) {
            correct = executeDup2Insn(context, stackVals);
        } else if (opCode == OP_SWAP1) {
            correct = executeSwap1Insn(context, stackVals);
        } else if (opCode == OP_SWAP2) {
            correct = executeSwap2Insn(context, stackVals);
        } else if (opCode == OP_TGET) {
            correct = executeTgetInsn(context, stackVals);
        } else if (opCode == OP_TSET) {
            correct = executeTsetInsn(context, stackVals);
        } else if (opCode == OP_TLEN) {
            correct = executeTlenInsn(context, stackVals);
        } else if (opCode == OP_XGET) {
            Value.Data memory auxVal;
            (offset, auxVal) = Marshaling.deserialize(_data.proof, offset);
            startMachine.addAuxStackValue(auxVal);
            correct = executeXgetInsn(context, stackVals, auxVal);
        } else if (opCode == OP_XSET) {
            Value.Data memory auxVal;
            (offset, auxVal) = Marshaling.deserialize(_data.proof, offset);
            startMachine.addAuxStackValue(auxVal);
            correct = executeXsetInsn(context, stackVals, auxVal);
        } else if (opCode == OP_BREAKPOINT) {
            correct = executeBreakpointInsn(context, stackVals);
        } else if (opCode == OP_LOG) {
            correct = executeLogInsn(context, stackVals);
        } else if (opCode == OP_SEND) {
            correct = executeSendInsn(context, stackVals);
        } else if (opCode == OP_INBOX) {
            correct = executeInboxInsn(context, _data.beforeInbox);
        } else if (opCode == OP_ERROR) {
            correct = false;
        } else if (opCode == OP_STOP) {
            context.machine.setHalt();
        } else if (opCode == OP_SETGAS) {
            correct = executeSetGasInsn(context, stackVals);
        } else if (opCode == OP_PUSHGAS) {
            correct = executePushGasInsn(context, stackVals);
        } else if (opCode == OP_ERR_CODE_POINT) {
            correct = executeErrCodePointInsn(context, stackVals);
        } else if (opCode == OP_PUSH_INSN) {
            correct = executePushInsnInsn(context, stackVals);
        } else if (opCode == OP_PUSH_INSN_IMM) {
            correct = executePushInsnImmInsn(context, stackVals);
        } else if (opCode == OP_SIDELOAD) {
            correct = executeSideloadInsn(context, stackVals);
        } else if (opCode == OP_ECRECOVER) {
            correct = executeECRecoverInsn(context, stackVals);
        } else {
            correct = false;
        }

        if (!correct) {
            if (context.machine.errHandlerHash == CODE_POINT_ERROR) {
                context.machine.setErrorStop();
            } else {
                context.machine.instructionStackHash = context
                    .machine
                    .errHandlerHash;
            }
        }

        require(
            _data.beforeHash == startMachine.hash(),
            "Proof had non matching start state"
        );
        // require(
        //     _data.beforeHash == startMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching start state: ", startMachine.toString(),
        //     " beforeHash = ", DebugPrint.bytes32string(_data.beforeHash), "\nstartMachine = ", DebugPrint.bytes32string(startMachine.hash())))
        // );

        return context;
    }
}
