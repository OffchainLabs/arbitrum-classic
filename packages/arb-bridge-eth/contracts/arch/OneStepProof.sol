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

import "./Value.sol";
import "./Machine.sol";

// Sourced from https://github.com/leapdao/solEVM-enforcer/tree/master

library OneStepProof {
    using Machine for Machine.Data;
    using Value for Value.Data;

    uint256 private constant SEND_SIZE_LIMIT = 10000;

    struct ValidateProofData {
        bytes32 beforeHash;
        uint128[4] timeBounds;
        Value.Data beforeInbox;
        bytes32 afterHash;
        bool didInboxInsn;
        bytes32 firstMessage;
        bytes32 lastMessage;
        bytes32 firstLog;
        bytes32 lastLog;
        uint64 gas;
        bytes proof;
    }

    function validateProof(
        bytes32 beforeHash,
        uint128[4] memory timeBounds,
        bytes32 beforeInbox,
        uint256 beforeInboxValueSize,
        bytes32 afterHash,
        bool didInboxInsn,
        bytes32 firstMessage,
        bytes32 lastMessage,
        bytes32 firstLog,
        bytes32 lastLog,
        uint64 gas,
        bytes memory proof
    ) internal pure returns (uint256) {
        return
            checkProof(
                ValidateProofData(
                    beforeHash,
                    timeBounds,
                    Value.newTuplePreImage(beforeInbox, beforeInboxValueSize),
                    afterHash,
                    didInboxInsn,
                    firstMessage,
                    lastMessage,
                    firstLog,
                    lastLog,
                    gas,
                    proof
                )
            );
    }

    /* solhint-disable no-inline-assembly */

    // Arithmetic

    function executeAddInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := add(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeMulInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := mul(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSubInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := sub(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeDivInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := div(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSdivInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := sdiv(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeModInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := mod(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSmodInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        if (b == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := smod(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeAddmodInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 m = val3.intVal;
        if (m == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := addmod(a, b, m)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeMulmodInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 m = val3.intVal;
        if (m == 0) {
            return false;
        }
        uint256 c;
        assembly {
            c := mulmod(a, b, m)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeExpInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := exp(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    // Comparison

    function executeLtInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := lt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeGtInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := gt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSltInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := slt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSgtInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := sgt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeEqInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(Value.newBoolean(val1.hash() == val2.hash()));
        return true;
    }

    function executeIszeroInsn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        if (!val1.isInt()) {
            machine.addDataStackInt(0);
        } else {
            uint256 a = val1.intVal;
            uint256 c;
            assembly {
                c := iszero(a)
            }
            machine.addDataStackInt(c);
        }
        return true;
    }

    function executeAndInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := and(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeOrInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := or(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeXorInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c;
        assembly {
            c := xor(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeNotInsn(Machine.Data memory machine, Value.Data memory val1)
        internal
        pure
        returns (bool)
    {
        if (!val1.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 c;
        assembly {
            c := not(a)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeByteInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 x = val1.intVal;
        uint256 n = val2.intVal;
        uint256 c;
        assembly {
            c := byte(n, x)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function executeSignextendInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 b = val1.intVal;
        uint256 a = val2.intVal;
        uint256 c;
        assembly {
            c := signextend(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    /* solhint-enable no-inline-assembly */

    // Hash

    function executeSha3Insn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        machine.addDataStackInt(uint256(val1.hash()));
        return true;
    }

    function executeTypeInsn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1.typeCodeVal());
        return true;
    }

    function executeEthhash2Insn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        bytes32 res = keccak256(abi.encodePacked(a, b));
        machine.addDataStackInt(uint256(res));
        return true;
    }

    // Stack ops

    function executePopInsn(Machine.Data memory, Value.Data memory)
        internal
        pure
        returns (bool)
    {
        return true;
    }

    function executeSpushInsn(Machine.Data memory machine)
        internal
        pure
        returns (bool)
    {
        machine.addDataStackValue(machine.staticVal);
        return true;
    }

    function executeRpushInsn(Machine.Data memory machine)
        internal
        pure
        returns (bool)
    {
        machine.addDataStackValue(machine.registerVal);
        return true;
    }

    function executeRsetInsn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        machine.registerVal = val1;
        return true;
    }

    function executeJumpInsn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        if (!val1.isCodePoint()) {
            return false;
        }
        machine.instructionStackHash = val1.hash();
        return true;
    }

    function executeCjumpInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isCodePoint()) {
            return false;
        }
        if (!val2.isInt()) {
            return false;
        }
        if (val2.intVal != 0) {
            machine.instructionStackHash = val1.hash();
        }
        return true;
    }

    function executeStackemptyInsn(Machine.Data memory machine)
        internal
        pure
        returns (bool)
    {
        machine.addDataStackValue(
            Value.newBoolean(
                Value.hash(machine.dataStack) == Value.hashEmptyTuple()
            )
        );
        return true;
    }

    function executePcpushInsn(
        Machine.Data memory machine,
        bytes32 codepointHash
    ) internal pure returns (bool) {
        machine.addDataStackValue(Value.newCodepointHash(codepointHash));
        return true;
    }

    function executeAuxpushInsn(
        Machine.Data memory machine,
        Value.Data memory val
    ) internal pure returns (bool) {
        machine.addAuxStackValue(val);
        return true;
    }

    function executeAuxstackemptyInsn(Machine.Data memory machine)
        internal
        pure
        returns (bool)
    {
        machine.addDataStackValue(
            Value.newBoolean(
                Value.hash(machine.auxStack) == Value.hashEmptyTuple()
            )
        );
        return true;
    }

    function executeErrpushInsn(Machine.Data memory machine)
        internal
        pure
        returns (bool)
    {
        machine.addDataStackValue(
            Value.newCodepointHash(machine.errHandlerHash)
        );
        return true;
    }

    function executeErrsetInsn(
        Machine.Data memory machine,
        Value.Data memory val
    ) internal pure returns (bool) {
        if (!val.isCodePoint()) {
            return false;
        }
        machine.errHandlerHash = val.hash();
        return true;
    }

    // Dup ops

    function executeDup0Insn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val1);
        return true;
    }

    function executeDup1Insn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        return true;
    }

    function executeDup2Insn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3
    ) internal pure returns (bool) {
        machine.addDataStackValue(val3);
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val3);
        return true;
    }

    // Swap ops

    function executeSwap1Insn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        return true;
    }

    function executeSwap2Insn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val3);
        return true;
    }

    // Tuple ops

    function executeTgetInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isTuple()) {
            return false;
        }

        if (val1.intVal >= val2.valLength()) {
            return false;
        }

        machine.addDataStackValue(val2.tupleVal[val1.intVal]);
        return true;
    }

    function executeTsetInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3
    ) internal pure returns (bool) {
        if (!val2.isTuple() || !val1.isInt()) {
            return false;
        }

        if (val1.intVal >= val2.valLength()) {
            return false;
        }
        Value.Data[] memory tupleVals = val2.tupleVal;
        tupleVals[val1.intVal] = val3;

        machine.addDataStackValue(Value.newTuple(tupleVals));
        return true;
    }

    function executeTlenInsn(
        Machine.Data memory machine,
        Value.Data memory val1
    ) internal pure returns (bool) {
        if (!val1.isTuple()) {
            return false;
        }
        machine.addDataStackInt(val1.valLength());
        return true;
    }

    // Logging

    function executeBreakpointInsn(Machine.Data memory)
        internal
        pure
        returns (bool)
    {
        return true;
    }

    function executeLogInsn(Machine.Data memory, Value.Data memory val1)
        internal
        pure
        returns (bool, bytes32)
    {
        return (true, val1.hash());
    }

    // System operations

    function executeSendInsn(Machine.Data memory, Value.Data memory val1)
        internal
        pure
        returns (bool, bytes32)
    {
        if (val1.size > SEND_SIZE_LIMIT) {
            return (false, 0);
        }
        if (!val1.isValidForSend()) {
            return (false, 0);
        }
        return (true, val1.hash());
    }

    function executeInboxInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory beforeInbox,
        uint256 lowerTimeBound
    ) internal pure returns (bool) {
        if (!val1.isInt()) {
            return false;
        }
        require(
            val1.intVal >= lowerTimeBound ||
                Value.hash(beforeInbox) != Value.hashEmptyTuple(),
            "Inbox instruction was blocked"
        );
        machine.addDataStackValue(beforeInbox);
        return true;
    }

    function executeECRecoverInsn(
        Machine.Data memory machine,
        Value.Data memory val1,
        Value.Data memory val2,
        Value.Data memory val3,
        Value.Data memory val4
    )
        internal
        pure
        returns (bool)
    {
        if (!val1.isInt() || !val2.isInt() || !val3.isInt() || !val4.isInt()) {
            return false;
        }
        bytes32 r = bytes32(val1.intVal);
        bytes32 s = bytes32(val2.intVal);
        if (val3.intVal != 0 && val3.intVal != 1) {
            machine.addDataStackInt(0);
            return true;
        }
        uint8 v = uint8(val3.intVal) + 27;
        bytes32 message = bytes32(val4.intVal);
        address ret = ecrecover(
            message,
            v,
            r,
            s
        );
        machine.addDataStackInt(uint256(ret));
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

    // Logging opertations
    uint8 internal constant OP_BREAKPOINT = 0x60;
    uint8 internal constant OP_LOG = 0x61;

    // System operations
    uint8 internal constant OP_SEND = 0x70;
    uint8 internal constant OP_GETTIME = 0x71;
    uint8 internal constant OP_INBOX = 0x72;
    uint8 internal constant OP_ERROR = 0x73;
    uint8 internal constant OP_STOP = 0x74;

    uint8 constant internal OP_ECRECOVER = 0x80;

    function opInfo(uint256 opCode) internal pure returns (uint256, uint256) {
        if (opCode == OP_ADD) {
            return (2, 1);
        } else if (opCode == OP_MUL) {
            return (2, 1);
        } else if (opCode == OP_SUB) {
            return (2, 1);
        } else if (opCode == OP_DIV) {
            return (2, 1);
        } else if (opCode == OP_SDIV) {
            return (2, 1);
        } else if (opCode == OP_MOD) {
            return (2, 1);
        } else if (opCode == OP_SMOD) {
            return (2, 1);
        } else if (opCode == OP_ADDMOD) {
            return (3, 1);
        } else if (opCode == OP_MULMOD) {
            return (3, 1);
        } else if (opCode == OP_EXP) {
            return (2, 1);
        } else if (opCode == OP_LT) {
            return (2, 1);
        } else if (opCode == OP_GT) {
            return (2, 1);
        } else if (opCode == OP_SLT) {
            return (2, 1);
        } else if (opCode == OP_SGT) {
            return (2, 1);
        } else if (opCode == OP_EQ) {
            return (2, 1);
        } else if (opCode == OP_ISZERO) {
            return (1, 1);
        } else if (opCode == OP_AND) {
            return (2, 1);
        } else if (opCode == OP_OR) {
            return (2, 1);
        } else if (opCode == OP_XOR) {
            return (2, 1);
        } else if (opCode == OP_NOT) {
            return (1, 1);
        } else if (opCode == OP_BYTE) {
            return (2, 1);
        } else if (opCode == OP_SIGNEXTEND) {
            return (2, 1);
        } else if (opCode == OP_SHA3) {
            return (1, 1);
        } else if (opCode == OP_TYPE) {
            return (1, 1);
        } else if (opCode == OP_ETHHASH2) {
            return (2, 1);
        } else if (opCode == OP_POP) {
            return (1, 0);
        } else if (opCode == OP_SPUSH) {
            return (0, 1);
        } else if (opCode == OP_RPUSH) {
            return (0, 1);
        } else if (opCode == OP_RSET) {
            return (1, 0);
        } else if (opCode == OP_JUMP) {
            return (1, 0);
        } else if (opCode == OP_CJUMP) {
            return (2, 0);
        } else if (opCode == OP_STACKEMPTY) {
            return (0, 1);
        } else if (opCode == OP_PCPUSH) {
            return (0, 1);
        } else if (opCode == OP_AUXPUSH) {
            return (1, 0);
        } else if (opCode == OP_AUXPOP) {
            return (0, 1);
        } else if (opCode == OP_AUXSTACKEMPTY) {
            return (0, 1);
        } else if (opCode == OP_NOP) {
            return (0, 0);
        } else if (opCode == OP_ERRPUSH) {
            return (0, 1);
        } else if (opCode == OP_ERRSET) {
            return (1, 0);
        } else if (opCode == OP_DUP0) {
            return (1, 2);
        } else if (opCode == OP_DUP1) {
            return (2, 3);
        } else if (opCode == OP_DUP2) {
            return (3, 4);
        } else if (opCode == OP_SWAP1) {
            return (2, 2);
        } else if (opCode == OP_SWAP2) {
            return (3, 3);
        } else if (opCode == OP_TGET) {
            return (2, 1);
        } else if (opCode == OP_TSET) {
            return (3, 1);
        } else if (opCode == OP_TLEN) {
            return (1, 1);
        } else if (opCode == OP_BREAKPOINT) {
            return (0, 0);
        } else if (opCode == OP_LOG) {
            return (1, 0);
        } else if (opCode == OP_SEND) {
            return (1, 0);
        } else if (opCode == OP_GETTIME) {
            return (0, 1);
        } else if (opCode == OP_INBOX) {
            return (1, 1);
        } else if (opCode == OP_ERROR) {
            return (0, 0);
        } else if (opCode == OP_STOP) {
            return (0, 0);
        } else if (opCode == OP_ECRECOVER) {
            return (4, 1);
        } else {
            require(false, "Invalid opcode: opInfo()");
        }
    }

    function opPopCount(uint8 opCode) internal pure returns (uint256) {
        uint256 popCount;
        uint256 pushCount;
        (popCount, pushCount) = opInfo(opCode);
        return popCount;
    }

    function opGasCost(uint8 opCode) internal pure returns (uint64) {
        if (opCode == OP_ADD) {
            return 3;
        } else if (opCode == OP_MUL) {
            return 3;
        } else if (opCode == OP_SUB) {
            return 3;
        } else if (opCode == OP_DIV) {
            return 4;
        } else if (opCode == OP_SDIV) {
            return 7;
        } else if (opCode == OP_MOD) {
            return 4;
        } else if (opCode == OP_SMOD) {
            return 7;
        } else if (opCode == OP_ADDMOD) {
            return 4;
        } else if (opCode == OP_MULMOD) {
            return 4;
        } else if (opCode == OP_EXP) {
            return 25;
        } else if (opCode == OP_LT) {
            return 2;
        } else if (opCode == OP_GT) {
            return 2;
        } else if (opCode == OP_SLT) {
            return 2;
        } else if (opCode == OP_SGT) {
            return 2;
        } else if (opCode == OP_EQ) {
            return 2;
        } else if (opCode == OP_ISZERO) {
            return 1;
        } else if (opCode == OP_AND) {
            return 2;
        } else if (opCode == OP_OR) {
            return 2;
        } else if (opCode == OP_XOR) {
            return 2;
        } else if (opCode == OP_NOT) {
            return 1;
        } else if (opCode == OP_BYTE) {
            return 4;
        } else if (opCode == OP_SIGNEXTEND) {
            return 7;
        } else if (opCode == OP_SHA3) {
            return 7;
        } else if (opCode == OP_TYPE) {
            return 3;
        } else if (opCode == OP_ETHHASH2) {
            return 8;
        } else if (opCode == OP_POP) {
            return 1;
        } else if (opCode == OP_SPUSH) {
            return 1;
        } else if (opCode == OP_RPUSH) {
            return 1;
        } else if (opCode == OP_RSET) {
            return 2;
        } else if (opCode == OP_JUMP) {
            return 4;
        } else if (opCode == OP_CJUMP) {
            return 4;
        } else if (opCode == OP_STACKEMPTY) {
            return 2;
        } else if (opCode == OP_PCPUSH) {
            return 1;
        } else if (opCode == OP_AUXPUSH) {
            return 1;
        } else if (opCode == OP_AUXPOP) {
            return 1;
        } else if (opCode == OP_AUXSTACKEMPTY) {
            return 2;
        } else if (opCode == OP_NOP) {
            return 1;
        } else if (opCode == OP_ERRPUSH) {
            return 1;
        } else if (opCode == OP_ERRSET) {
            return 1;
        } else if (opCode == OP_DUP0) {
            return 1;
        } else if (opCode == OP_DUP1) {
            return 1;
        } else if (opCode == OP_DUP2) {
            return 1;
        } else if (opCode == OP_SWAP1) {
            return 1;
        } else if (opCode == OP_SWAP2) {
            return 1;
        } else if (opCode == OP_TGET) {
            return 2;
        } else if (opCode == OP_TSET) {
            return 40;
        } else if (opCode == OP_TLEN) {
            return 2;
        } else if (opCode == OP_BREAKPOINT) {
            return 100;
        } else if (opCode == OP_LOG) {
            return 100;
        } else if (opCode == OP_SEND) {
            return 100;
        } else if (opCode == OP_GETTIME) {
            return 40;
        } else if (opCode == OP_INBOX) {
            return 40;
        } else if (opCode == OP_ERROR) {
            return 5;
        } else if (opCode == OP_STOP) {
            return 10;
        } else if (opCode == OP_ECRECOVER) {
            return 20000;
        } else {
            require(false, "Invalid opcode: opGasCost()");
        }
    }

    function loadMachine(ValidateProofData memory _data)
        internal
        pure
        returns (
            uint8,
            Value.Data[] memory,
            Machine.Data memory,
            Machine.Data memory,
            uint256
        )
    {
        uint256 offset = 0;
        bool valid;
        Machine.Data memory startMachine;
        startMachine.setExtensive();
        (valid, offset, startMachine) = Machine.deserializeMachine(
            _data.proof,
            offset
        );

        require(valid, "loadMachine(): invalid machine");

        Machine.Data memory endMachine = startMachine.clone();
        uint8 immediate = uint8(_data.proof[offset]);
        uint8 opCode = uint8(_data.proof[offset + 1]);
        uint256 popCount = opPopCount(opCode);
        Value.Data[] memory stackVals = new Value.Data[](popCount);
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
            (valid, offset, immediateVal) = Value.deserialize(
                _data.proof,
                offset
            );
            // string(abi.encodePacked("Proof had bad immediate value ", uint2str(valid)))
            require(valid, "Proof had bad immediate value");
            if (popCount > 0) {
                stackVals[0] = immediateVal;
            } else {
                endMachine.addDataStackValue(immediateVal);
            }

            startMachine.instructionStackHash = Value
                .newCodePoint(
                uint8(opCode),
                startMachine
                    .instructionStackHash,
                Value.hash(immediateVal)
            )
                .hash();
        }

        uint256 i = 0;
        for (i = immediate; i < popCount; i++) {
            (valid, offset, stackVals[i]) = Value.deserialize(
                _data.proof,
                offset
            );
            require(valid, "Proof had bad stack value");
        }
        if (stackVals.length > 0) {
            for (i = 0; i < stackVals.length - immediate; i++) {
                startMachine.addDataStackValue(
                    stackVals[stackVals.length - 1 - i]
                );
            }
        }
        return (opCode, stackVals, startMachine, endMachine, offset);
    }

    uint8 private constant CODE_POINT_TYPECODE = 1;
    bytes32 private constant CODE_POINT_ERROR = keccak256(
        abi.encodePacked(CODE_POINT_TYPECODE, uint8(0), bytes32(0))
    );

    function checkProof(ValidateProofData memory _data)
        internal
        pure
        returns (uint256)
    {
        uint8 opCode;
        bool valid;
        uint256 offset;
        Value.Data[] memory stackVals;
        Machine.Data memory startMachine;
        Machine.Data memory endMachine;
        (opCode, stackVals, startMachine, endMachine, offset) = loadMachine(
            _data
        );

        bool correct = true;
        bytes32 messageHash;
        require(_data.gas == opGasCost(opCode), "Invalid gas in proof");
        require(
            (_data.didInboxInsn && opCode == OP_INBOX) ||
                (!_data.didInboxInsn && opCode != OP_INBOX),
            "Invalid didInboxInsn claim"
        );
        if (opCode == OP_ADD) {
            correct = executeAddInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_MUL) {
            correct = executeMulInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SUB) {
            correct = executeSubInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_DIV) {
            correct = executeDivInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SDIV) {
            correct = executeSdivInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_MOD) {
            correct = executeModInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SMOD) {
            correct = executeSmodInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_ADDMOD) {
            correct = executeAddmodInsn(
                endMachine,
                stackVals[0],
                stackVals[1],
                stackVals[2]
            );
        } else if (opCode == OP_MULMOD) {
            correct = executeMulmodInsn(
                endMachine,
                stackVals[0],
                stackVals[1],
                stackVals[2]
            );
        } else if (opCode == OP_EXP) {
            correct = executeExpInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_LT) {
            correct = executeLtInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_GT) {
            correct = executeGtInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SLT) {
            correct = executeSltInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SGT) {
            correct = executeSgtInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_EQ) {
            correct = executeEqInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_ISZERO) {
            correct = executeIszeroInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_AND) {
            correct = executeAndInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_OR) {
            correct = executeOrInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_XOR) {
            correct = executeXorInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_NOT) {
            correct = executeNotInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_BYTE) {
            correct = executeByteInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SIGNEXTEND) {
            correct = executeSignextendInsn(
                endMachine,
                stackVals[0],
                stackVals[1]
            );
        } else if (opCode == OP_SHA3) {
            correct = executeSha3Insn(endMachine, stackVals[0]);
        } else if (opCode == OP_TYPE) {
            correct = executeTypeInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_ETHHASH2) {
            correct = executeEthhash2Insn(
                endMachine,
                stackVals[0],
                stackVals[1]
            );
        } else if (opCode == OP_POP) {
            correct = executePopInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_SPUSH) {
            correct = executeSpushInsn(endMachine);
        } else if (opCode == OP_RPUSH) {
            correct = executeRpushInsn(endMachine);
        } else if (opCode == OP_RSET) {
            correct = executeRsetInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_JUMP) {
            correct = executeJumpInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_CJUMP) {
            correct = executeCjumpInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_STACKEMPTY) {
            correct = executeStackemptyInsn(endMachine);
        } else if (opCode == OP_PCPUSH) {
            correct = executePcpushInsn(
                endMachine,
                startMachine.instructionStackHash
            );
        } else if (opCode == OP_AUXPUSH) {
            correct = executeAuxpushInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_AUXPOP) {
            Value.Data memory auxVal;
            (valid, offset, auxVal) = Value.deserialize(_data.proof, offset);
            require(valid, "Proof of auxpop had bad aux value");
            startMachine.addAuxStackValue(auxVal);
            endMachine.addDataStackValue(auxVal);
        } else if (opCode == OP_AUXSTACKEMPTY) {
            correct = executeAuxstackemptyInsn(endMachine);
        } else if (opCode == OP_NOP) {
            correct = true;
        } else if (opCode == OP_ERRPUSH) {
            correct = executeErrpushInsn(endMachine);
        } else if (opCode == OP_ERRSET) {
            correct = executeErrsetInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_DUP0) {
            correct = executeDup0Insn(endMachine, stackVals[0]);
        } else if (opCode == OP_DUP1) {
            correct = executeDup1Insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_DUP2) {
            correct = executeDup2Insn(
                endMachine,
                stackVals[0],
                stackVals[1],
                stackVals[2]
            );
        } else if (opCode == OP_SWAP1) {
            correct = executeSwap1Insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SWAP2) {
            correct = executeSwap2Insn(
                endMachine,
                stackVals[0],
                stackVals[1],
                stackVals[2]
            );
        } else if (opCode == OP_TGET) {
            correct = executeTgetInsn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_TSET) {
            correct = executeTsetInsn(
                endMachine,
                stackVals[0],
                stackVals[1],
                stackVals[2]
            );
        } else if (opCode == OP_TLEN) {
            correct = executeTlenInsn(endMachine, stackVals[0]);
        } else if (opCode == OP_BREAKPOINT) {
            correct = executeBreakpointInsn(endMachine);
        } else if (opCode == OP_LOG) {
            (correct, messageHash) = executeLogInsn(endMachine, stackVals[0]);
            if (correct) {
                require(
                    keccak256(abi.encodePacked(_data.firstLog, messageHash)) ==
                        _data.lastLog,
                    "Logged value doesn't match output log"
                );
                require(
                    _data.firstMessage == _data.lastMessage,
                    "Send not called, but message is nonzero"
                );
            } else {
                messageHash = 0;
            }
        } else if (opCode == OP_SEND) {
            (correct, messageHash) = executeSendInsn(endMachine, stackVals[0]);
            if (correct) {
                if (messageHash == 0) {
                    require(
                        _data.firstMessage == _data.lastMessage,
                        "Send value exceeds size limit, no message should be sent"
                    );
                } else {
                    require(
                        keccak256(
                            abi.encodePacked(_data.firstMessage, messageHash)
                        ) == _data.lastMessage,
                        "sent message doesn't match output message"
                    );

                    require(
                        _data.firstLog == _data.lastLog,
                        "Log not called, but message is nonzero"
                    );
                }
            } else {
                messageHash = 0;
            }
        } else if (opCode == OP_GETTIME) {
            Value.Data[] memory contents = new Value.Data[](4);
            contents[0] = Value.newInt(_data.timeBounds[0]);
            contents[1] = Value.newInt(_data.timeBounds[1]);
            contents[2] = Value.newInt(_data.timeBounds[2]);
            contents[3] = Value.newInt(_data.timeBounds[3]);
            endMachine.addDataStackValue(Value.newTuple(contents));
        } else if (opCode == OP_INBOX) {
            correct = executeInboxInsn(
                endMachine,
                stackVals[0],
                _data.beforeInbox,
                _data.timeBounds[0]
            );
        } else if (opCode == OP_ERROR) {
            correct = false;
        } else if (opCode == OP_STOP) {
            endMachine.setHalt();
        } else if (opCode == OP_ECRECOVER) {
            correct = executeECRecoverInsn(endMachine, stackVals[0], stackVals[1], stackVals[2], stackVals[3]);
        } else {
            correct = false;
        }

        if (messageHash == 0) {
            require(
                _data.firstMessage == _data.lastMessage,
                "Send not called, but message is nonzero"
            );
            require(
                _data.firstLog == _data.lastLog,
                "Log not called, but message is nonzero"
            );
        }

        if (!correct) {
            if (endMachine.errHandlerHash == CODE_POINT_ERROR) {
                endMachine.setErrorStop();
            } else {
                endMachine.instructionStackHash = endMachine.errHandlerHash;
            }
        }

        require(
            _data.beforeHash == startMachine.hash(),
            "Proof had non matching start state"
        );
        require(
            _data.afterHash == endMachine.hash(),
            "Proof had non matching end state"
        );

        // require(
        //     _data.beforeHash == startMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching start state: ", startMachine.toString(),
        //     " beforeHash = ", DebugPrint.bytes32string(_data.beforeHash), "\nstartMachine = ", DebugPrint.bytes32string(startMachine.hash())))
        // );
        // require(
        //     _data.afterHash == endMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching end state: ", endMachine.toString(),
        //     " afterHash = ", DebugPrint.bytes32string(_data.afterHash), "\nendMachine = ", DebugPrint.bytes32string(endMachine.hash())))
        // );

        return 0;
    }
}
