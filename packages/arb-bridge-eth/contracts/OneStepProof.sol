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
import "./ArbProtocol.sol";
import "./ArbValue.sol";
import "./ArbMachine.sol";

// Sourced from https://github.com/leapdao/solEVM-enforcer/tree/master

library OneStepProof {

    using BytesLib for bytes;
    using ArbMachine for ArbMachine.Machine;
    using ArbValue for ArbValue.Value;

    // Arithmetic

    function execute_add_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := add(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_mul_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := mul(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_sub_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := sub(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_div_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := div(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_sdiv_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := sdiv(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_mod_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := mod(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_smod_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := smod(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_addmod_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2,
        ArbValue.Value memory val3
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint m = val3.intVal;
        uint c;
        assembly {
            c := addmod(a, b, m)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_mulmod_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2,
        ArbValue.Value memory val3
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint m = val3.intVal;
        uint c;
        assembly {
            c := mulmod(a, b, m)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_exp_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := exp(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    // Comparison

    function execute_lt_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := lt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_gt_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := gt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_slt_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := slt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_sgt_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := slt(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_eq_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(ArbValue.newBooleanValue(val1.hash().hash == val2.hash().hash));
        return true;
    }

    function execute_iszero_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        if (!val1.isInt()) {
            machine.addDataStackInt(0);
        } else {
            uint a = val1.intVal;
            uint c;
            assembly {
                c := iszero(a)
            }
            machine.addDataStackInt(c);
        }
        return true;
    }

    function execute_and_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := and(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_or_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := or(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_xor_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := xor(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_not_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        if (!val1.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint c;
        assembly {
            c := not(a)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_byte_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint n = val1.intVal;
        uint x = val2.intVal;
        uint c;
        assembly {
            c := byte(n, x)
        }
        machine.addDataStackInt(c);
        return true;
    }

    function execute_signextend_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isInt()) {
            return false;
        }
        uint a = val1.intVal;
        uint b = val2.intVal;
        uint c;
        assembly {
            c := signextend(a, b)
        }
        machine.addDataStackInt(c);
        return true;
    }

    // Hash

    function execute_sha3_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        machine.addDataStackInt(uint256(val1.hash().hash));
        return true;
    }

    function execute_type_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1.typeCodeVal());
        return true;
    }

    // Stack ops

    function execute_pop_insn(
        ArbMachine.Machine memory,
        ArbValue.Value memory
    ) internal pure returns (bool) {
        return true;
    }

    function execute_spush_insn(
        ArbMachine.Machine memory machine
    ) internal pure returns (bool) {
        machine.addDataStackHashValue(machine.staticHash);
        return true;
    }

    function execute_rpush_insn(
        ArbMachine.Machine memory machine
    ) internal pure returns (bool) {
        machine.addDataStackHashValue(machine.registerHash);
        return true;
    }

    function execute_rset_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        machine.registerHash = val1.hash();
        return true;
    }

    function execute_jump_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        machine.instructionStackHash = val1.hash();
        return true;
    }

    function execute_cjump_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val2.isInt()) {
            return false;
        }
        if (val1.intVal != 0) {
            machine.instructionStackHash = val1.hash();
        }
        return true;
    }

    function execute_stackempty_insn(
        ArbMachine.Machine memory machine
    ) internal pure returns (bool) {
        machine.addDataStackValue(
            ArbValue.newBooleanValue(machine.dataStackHash.hash == ArbValue.newNoneValue().hash().hash)
        );
        return true;
    }

    function execute_pcpush_insn(
        ArbMachine.Machine memory machine,
        ArbValue.HashOnlyValue memory pc
    ) internal pure returns (bool) {
        machine.addDataStackHashValue(pc);
        return true;
    }

    function execute_auxpush_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val
    ) internal pure returns (bool) {
        machine.addAuxStackValue(val);
        return true;
    }

    function execute_auxstackempty_insn(
        ArbMachine.Machine memory machine
    ) internal pure returns (bool) {
        machine.addDataStackValue(
            ArbValue.newBooleanValue(machine.auxStackHash.hash == ArbValue.newNoneValue().hash().hash)
        );
        return true;
    }

    function execute_errpush_insn(
        ArbMachine.Machine memory machine
    ) internal pure returns (bool) {
        machine.addDataStackHashValue(machine.errHandler);
        return true;
    }

    function execute_errset_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val
    ) internal pure returns (bool) {
        machine.errHandler = val.hash();
        return true;
    }


    // Dup ops

    function execute_dup0_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val1);
        return true;
    }

    function execute_dup1_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        return true;
    }

    function execute_dup2_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2,
        ArbValue.Value memory val3) internal pure returns (bool) {
        machine.addDataStackValue(val3);
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val3);
        return true;
    }

    // Swap ops

    function execute_swap1_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        return true;
    }

    function execute_swap2_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2,
        ArbValue.Value memory val3
    ) internal pure returns (bool) {
        machine.addDataStackValue(val1);
        machine.addDataStackValue(val2);
        machine.addDataStackValue(val3);
        return true;
    }

    // Tuple ops

    function execute_tget_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2
    ) internal pure returns (bool) {
        if (!val1.isInt() || !val2.isTuple()) {
            return false;
        }

        if (val1.intVal > val2.valLength()) {
            return false;
        }

        machine.addDataStackValue(val2.tupleVal[val1.intVal]);
        return true;
    }

    function execute_tset_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.Value memory val2,
        ArbValue.Value memory val3
    ) internal pure returns (bool) {
        if (!val2.isTuple() || !val1.isInt()) {
            return false;
        }

        if (val1.intVal > val2.valLength()) {
            return false;
        }
        val2.tupleVal[val1.intVal] = val3;
        machine.addDataStackValue(val2);
        return true;
    }

    function execute_tlen_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1
    ) internal pure returns (bool) {
        if (!val1.isTuple()) {
            return false;
        }
        machine.addDataStackInt(val1.valLength());
        return true;
    }

    // Logging

    function execute_breakpoint_insn(
        ArbMachine.Machine memory,
        ArbValue.Value memory
    ) internal pure returns (bool) {
        return true;
    }

    function execute_log_insn(
        ArbMachine.Machine memory,
        ArbValue.Value memory val1
    ) internal pure returns (bool, bytes32) {
        ArbValue.HashOnlyValue memory hashVal = val1.hash();
        return (true, hashVal.hash);
    }

    // System operations

    function send_insn_impl(
        ArbMachine.Machine memory,
        ArbValue.Value memory val1
    ) internal pure returns (
        bool, bytes32, bytes21, uint
    ) {
        bytes21 tokenType;
        uint amount;
        bytes32 messageHash;
        if (!val1.isTuple()) {
            return (false, messageHash, tokenType, amount);
        }
        if (!val1.tupleVal[1].isInt()) {
            return (false, messageHash, tokenType, amount);
        }
        if (!val1.tupleVal[2].isInt()) {
            return (false, messageHash, tokenType, amount);
        }
        if (!val1.tupleVal[3].isInt()) {
            return (false, messageHash, tokenType, amount);
        }
        tokenType = bytes21(bytes32(val1.tupleVal[1].intVal));
        amount = val1.tupleVal[2].intVal;
        messageHash = ArbProtocol.generateMessageStubHash(
            val1.hash().hash,
            tokenType,
            amount,
            bytes32(val1.tupleVal[3].intVal)
        );
        return (true, messageHash, tokenType, amount);
    }

    function execute_inbox_insn(
        ArbMachine.Machine memory machine,
        ArbValue.Value memory val1,
        ArbValue.HashOnlyValue memory beforeInbox
    ) internal pure returns (bool) {
        require(val1.hash().hash != beforeInbox.hash, "Inbox instruction was blocked");
        machine.addDataStackHashValue(beforeInbox);
        return true;
    }

    // Stop and arithmetic ops
    uint8 constant internal OP_ADD = 0x01;
    uint8 constant internal OP_MUL = 0x02;
    uint8 constant internal OP_SUB = 0x03;
    uint8 constant internal OP_DIV = 0x04;
    uint8 constant internal OP_SDIV = 0x05;
    uint8 constant internal OP_MOD = 0x06;
    uint8 constant internal OP_SMOD = 0x07;
    uint8 constant internal OP_ADDMOD = 0x08;
    uint8 constant internal OP_MULMOD = 0x09;
    uint8 constant internal OP_EXP = 0x0a;

    // Comparison & bitwise logic
    uint8 constant internal OP_LT = 0x10;
    uint8 constant internal OP_GT = 0x11;
    uint8 constant internal OP_SLT = 0x12;
    uint8 constant internal OP_SGT = 0x13;
    uint8 constant internal OP_EQ = 0x14;
    uint8 constant internal OP_ISZERO = 0x15;
    uint8 constant internal OP_AND = 0x16;
    uint8 constant internal OP_OR = 0x17;
    uint8 constant internal OP_XOR = 0x18;
    uint8 constant internal OP_NOT = 0x19;
    uint8 constant internal OP_BYTE = 0x1a;
    uint8 constant internal OP_SIGNEXTEND = 0x1b;

    // SHA3
    uint8 constant internal OP_SHA3 = 0x20;
    uint8 constant internal OP_TYPE = 0x21;

    // Stack, Memory, Storage and Flow Operations
    uint8 constant internal OP_POP = 0x30;
    uint8 constant internal OP_SPUSH = 0x31;
    uint8 constant internal OP_RPUSH = 0x32;
    uint8 constant internal OP_RSET = 0x33;
    uint8 constant internal OP_JUMP = 0x34;
    uint8 constant internal OP_CJUMP = 0x35;
    uint8 constant internal OP_STACKEMPTY = 0x36;
    uint8 constant internal OP_PCPUSH = 0x37;
    uint8 constant internal OP_AUXPUSH = 0x38;
    uint8 constant internal OP_AUXPOP = 0x39;
    uint8 constant internal OP_AUXSTACKEMPTY = 0x3a;
    uint8 constant internal OP_NOP = 0x3b;
    uint8 constant internal OP_ERRPUSH = 0x3c;
    uint8 constant internal OP_ERRSET = 0x3d;

    // Duplication and Exchange operations
    uint8 constant internal OP_DUP0 = 0x40;
    uint8 constant internal OP_DUP1 = 0x41;
    uint8 constant internal OP_DUP2 = 0x42;
    uint8 constant internal OP_SWAP1 = 0x43;
    uint8 constant internal OP_SWAP2 = 0x44;

    // Tuple opertations
    uint8 constant internal OP_TGET = 0x50;
    uint8 constant internal OP_TSET = 0x51;
    uint8 constant internal OP_TLEN = 0x52;

    // Logging opertations
    uint8 constant internal OP_BREAKPOINT = 0x60;
    uint8 constant internal OP_LOG = 0x61;

    // System operations
    uint8 constant internal OP_SEND = 0x70;
    uint8 constant internal OP_NBSEND = 0x71;
    uint8 constant internal OP_GETTIME = 0x72;
    uint8 constant internal OP_INBOX = 0x73;
    uint8 constant internal OP_ERROR = 0x74;
    uint8 constant internal OP_STOP = 0x75;

    function opInfo(uint opCode) internal pure returns (uint, uint) {
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
        } else if (opCode == OP_NBSEND) {
            return (1, 1);
        } else if (opCode == OP_GETTIME) {
            return (0, 1);
        } else if (opCode == OP_INBOX) {
            return (1, 1);
        } else if (opCode == OP_ERROR) {
            return (0, 0);
        } else if (opCode == OP_STOP) {
            return (0, 0);
        } else {
            require(false, "Invalid opcode");
        }
    }

    struct ValidateProofData {
        bytes32 beforeHash;
        uint64[2] timeBounds;
        bytes32 beforeInbox;
        bytes32 afterHash;
        bytes32 firstMessage;
        bytes32 lastMessage;
        bytes32 firstLog;
        bytes32 lastLog;
        bytes21 tokenType;
        uint amount;
        bool foundAmount;
        bytes proof;
    }

    // fields
    // _beforeHash
    // _beforeInbox
    // _afterHash
    // _firstMessageHash
    // _lastMessageHash
    // _firstLogHash
    // _lastLogHash

    event SawMachine(
        bytes32 instructionStack,
        bytes32 dataStack,
        bytes32 auxStack,
        bytes32 register,
        bytes32 staticHash,
        bytes32 errHandler
    );

    function validateProof(
        bytes32[7] memory fields,
        uint64[2] memory timeBounds,
        bytes21[] memory tokenTypes,
        uint256[] memory beforeValues,
        uint256[] memory messageValue,
        bytes memory proof
    ) public pure returns(uint) {
        // require(messageValue.length == 1 || messageValue.length == 0);
        bytes21 tokenType;
        uint amount;
        bool foundAmount;

        bool includesMessage = (fields[3] != fields[4]);
        int64 amountIndex = -1;
        if (includesMessage) {
            for (uint64 i = 0; i < messageValue.length; i++) {
                if (messageValue[i] != 0) {
                    require(amountIndex == -1, "multiple out messages");
                    amountIndex = int64(i);
                }
            }
            if (amountIndex != -1) {
                amount = messageValue[uint(amountIndex)];
                tokenType = tokenTypes[uint(amountIndex)];
                foundAmount = true;
                if (tokenTypes[uint(amountIndex)][20] == 0x01) {
                    require(beforeValues[uint(amountIndex)] == amount, "precondition must have nft");
                } else {
                    require(amount <= beforeValues[uint(amountIndex)], "precondition must have value");
                }
            }
        } else {
            for (uint64 i = 0; i < messageValue.length; i++) {
                require(messageValue[i] == 0, "Must have no message values");
            }
        }
        return checkProof(
            ValidateProofData(
                fields[0],
                timeBounds,
                fields[1],
                fields[2],
                fields[3],
                fields[4],
                fields[5],
                fields[6],
                tokenType,
                amount,
                foundAmount,
                proof
            )
        );
    }

    // Taken from https://github.com/oraclize/ethereum-api/blob/master/oraclizeAPI_0.5.sol
    function uint2str(uint _iParam) internal pure returns (string memory _uintAsString) {
        uint _i = _iParam;
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (_i != 0) {
            bstr[k--] = byte(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);
    }

    function opPopCount(uint8 opCode) internal pure returns(uint) {
        uint popCount;
        uint pushCount;
        (popCount, pushCount) = opInfo(opCode);
        return popCount;
    }

    function loadMachine(
        ValidateProofData memory _data
    ) internal pure returns (
        uint8,
        ArbValue.Value[] memory,
        ArbMachine.Machine memory,
        ArbMachine.Machine memory,
        uint
    ) {
        uint offset = 0;
        uint valid = 0;
        ArbMachine.Machine memory startMachine;
        startMachine.setExtensive();
        (valid, offset, startMachine) = ArbMachine.deserialize_machine(_data.proof, offset);
        ArbMachine.Machine memory endMachine = startMachine.clone();
        uint8 immediate = uint8(_data.proof[offset]);
        uint8 opCode = uint8(_data.proof[offset + 1]);
        uint popCount = opPopCount(opCode);
        ArbValue.Value[] memory stackVals = new ArbValue.Value[](popCount);
        offset += 2;

        require(immediate == 0 || immediate == 1, "Proof had bad operation type");
        if (immediate == 0) {
            startMachine.instructionStackHash = ArbValue.HashOnlyValue(ArbValue.hashCodePointBasicValue(
                uint8(opCode),
                startMachine.instructionStackHash.hash
            ));
        } else {
            ArbValue.Value memory immediateVal;
            (valid, offset, immediateVal) = ArbValue.deserialize_value(_data.proof, offset);
            // string(abi.encodePacked("Proof had bad immediate value ", uint2str(valid)))
            require(valid == 0, "Proof had bad immediate value");
            if (popCount > 0) {
                stackVals[0] = immediateVal;
            } else {
                endMachine.addDataStackValue(immediateVal);
            }

            startMachine.instructionStackHash = ArbValue.HashOnlyValue(ArbValue.hashCodePointImmediateValue(
                uint8(opCode),
                immediateVal.hash().hash,
                startMachine.instructionStackHash.hash
            ));
        }

        uint i = 0;
        for (i = immediate; i < popCount; i++) {
            (valid, offset, stackVals[i]) = ArbValue.deserialize_value(_data.proof, offset);
            require(valid == 0, "Proof had bad stack value");
        }
        if (stackVals.length > 0) {
            for (i = 0; i < stackVals.length - immediate; i++) {
                startMachine.addDataStackValue(stackVals[stackVals.length - 1 - i]);
            }
        }
        return (opCode, stackVals, startMachine, endMachine, offset);
    }

    uint8 constant CodePointCode = 1;
    bytes32 constant CODE_POINT_ERROR = keccak256(
        abi.encodePacked(
            CodePointCode,
            uint8(0),
            bytes32(0)
        )
    );

    function checkProof(ValidateProofData memory _data) internal pure returns(uint) {
        uint8 opCode;
        uint valid = 0;
        uint offset;
        ArbValue.Value[] memory stackVals;
        ArbMachine.Machine memory startMachine;
        ArbMachine.Machine memory endMachine;
        (opCode, stackVals, startMachine, endMachine, offset) = loadMachine(_data);
        bool correct = true;
        bytes32 messageHash;
        if (opCode == OP_ADD) {
            correct = execute_add_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_MUL) {
            correct = execute_mul_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SUB) {
            correct = execute_sub_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_DIV) {
            correct = execute_div_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SDIV) {
            correct = execute_sdiv_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_MOD) {
            correct = execute_mod_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SMOD) {
            correct = execute_smod_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_ADDMOD) {
            correct = execute_addmod_insn(endMachine, stackVals[0], stackVals[1], stackVals[2]);
        } else if (opCode == OP_MULMOD) {
            correct = execute_mulmod_insn(endMachine, stackVals[0], stackVals[1], stackVals[2]);
        } else if (opCode == OP_EXP) {
            correct = execute_exp_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_LT) {
            correct = execute_lt_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_GT) {
            correct = execute_gt_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SLT) {
            correct = execute_slt_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SGT) {
            correct = execute_sgt_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_EQ) {
            correct = execute_eq_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_ISZERO) {
            correct = execute_iszero_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_AND) {
            correct = execute_and_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_OR) {
            correct = execute_or_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_XOR) {
            correct = execute_xor_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_NOT) {
            correct = execute_not_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_BYTE) {
            correct = execute_byte_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SIGNEXTEND) {
            correct = execute_signextend_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SHA3) {
            correct = execute_sha3_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_TYPE) {
            correct = execute_type_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_POP) {
            correct = execute_pop_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_SPUSH) {
            correct = execute_spush_insn(endMachine);
        } else if (opCode == OP_RPUSH) {
            correct = execute_rpush_insn(endMachine);
        } else if (opCode == OP_RSET) {
            correct = execute_rset_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_JUMP) {
            correct = execute_jump_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_CJUMP) {
            correct = execute_cjump_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_STACKEMPTY) {
            correct = execute_stackempty_insn(endMachine);
        } else if (opCode == OP_PCPUSH) {
            correct = execute_pcpush_insn(endMachine, startMachine.instructionStackHash);
        } else if (opCode == OP_AUXPUSH) {
            correct = execute_auxpush_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_AUXPOP) {
            ArbValue.Value memory auxVal;
            (valid, offset, auxVal) = ArbValue.deserialize_value(_data.proof, offset);
            require(valid == 0, "Proof of auxpop had bad aux value");
            startMachine.addAuxStackValue(auxVal);
            endMachine.addDataStackValue(auxVal);
        } else if (opCode == OP_NOP) {

        } else if (opCode == OP_DUP0) {
            correct = execute_dup0_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_DUP1) {
            correct = execute_dup1_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_DUP2) {
            correct = execute_dup2_insn(endMachine, stackVals[0], stackVals[1], stackVals[2]);
        } else if (opCode == OP_SWAP1) {
            correct = execute_swap1_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_SWAP2) {
            correct = execute_swap2_insn(endMachine, stackVals[0], stackVals[1], stackVals[2]);
        } else if (opCode == OP_TGET) {
            correct = execute_tget_insn(endMachine, stackVals[0], stackVals[1]);
        } else if (opCode == OP_TSET) {
            correct = execute_tset_insn(endMachine, stackVals[0], stackVals[1], stackVals[2]);
        } else if (opCode == OP_TLEN) {
            correct = execute_tlen_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_BREAKPOINT) {
            correct = execute_breakpoint_insn(endMachine, stackVals[0]);
        } else if (opCode == OP_LOG) {
            (correct, messageHash) = execute_log_insn(endMachine, stackVals[0]);
            require(
                keccak256(
                    abi.encodePacked(
                        _data.firstLog,
                        messageHash
                    )
                ) == _data.lastLog,
                "Logged value doesn't match output log"
            );
            require(_data.firstMessage == _data.lastMessage, "Send not called, but message is nonzero");
        } else if (opCode == OP_SEND) {
            bytes21 tokenType;
            uint amount;
            (correct, messageHash, tokenType, amount) = send_insn_impl(endMachine, stackVals[0]);
            require(
                keccak256(
                    abi.encodePacked(
                        _data.firstMessage,
                        messageHash
                    )
                ) == _data.lastMessage,
                "sent message doesn't match output mesage"
            );
            require(_data.firstLog == _data.lastLog, "Log not called, but message is nonzero");
        } else if (opCode == OP_NBSEND) {
            bytes21 tokenType;
            uint amount;
            (correct, messageHash, tokenType, amount) = send_insn_impl(endMachine, stackVals[0]);
            require(
                keccak256(
                    abi.encodePacked(
                        _data.firstMessage,
                        messageHash
                    )
                ) == _data.lastMessage,
                "sent message doesn't match output mesage"
            );
            require(_data.firstLog == _data.lastLog, "Log not called, but message is nonzero");
        } else if (opCode == OP_GETTIME) {
            ArbValue.Value[] memory contents = new ArbValue.Value[](2);
            contents[0] = ArbValue.newIntValue(_data.timeBounds[0]);
            contents[1] = ArbValue.newIntValue(_data.timeBounds[1]);
            endMachine.addDataStackValue(ArbValue.newTupleValue(contents));
        } else if (opCode == OP_INBOX) {
            correct = execute_inbox_insn(endMachine, stackVals[0], ArbValue.HashOnlyValue(_data.beforeInbox));
        } else if (opCode == OP_ERROR) {
            correct = false;
        } else if (opCode == OP_STOP) {
            endMachine.setHalt();
        }

        if (messageHash == 0) {
            require(_data.firstMessage == _data.lastMessage, "Send not called, but message is nonzero");
            require(_data.firstLog == _data.lastLog, "Log not called, but message is nonzero");
        }

        if (!correct) {
            if (endMachine.errHandler.hash == CODE_POINT_ERROR) {
                endMachine.setErrorStop();
            } else {
                endMachine.instructionStackHash = endMachine.errHandler;
            }
        }

        // require(
        //     _data.beforeHash == startMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching start state: ", startMachine.toString()))
        // );
        require(_data.beforeHash == startMachine.hash(), "Proof had non matching start state");
        // require(
        //     _data.afterHash == endMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching end state: ", endMachine.toString()))
        // );
        require(_data.afterHash == endMachine.hash(), "Proof had non matching end state");

        return 0;
    }
}
