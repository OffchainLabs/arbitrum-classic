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

pragma solidity ^0.6.11;

import "./IOneStepProof.sol";
import "./Value.sol";
import "./Machine.sol";
import "../bridge/interfaces/IBridge.sol";

abstract contract OneStepProofCommon {
    using Machine for Machine.Data;
    using Hashing for Value.Data;
    using Value for Value.Data;

    uint256 internal constant MAX_UINT256 = ((1 << 128) + 1) * ((1 << 128) - 1);

    uint64 internal constant ERROR_GAS_COST = 5;

    string internal constant BAD_IMM_TYP = "BAD_IMM_TYP";
    string internal constant NO_IMM = "NO_IMM";
    string internal constant STACK_MISSING = "STACK_MISSING";
    string internal constant AUX_MISSING = "AUX_MISSING";
    string internal constant STACK_MANY = "STACK_MANY";
    string internal constant AUX_MANY = "AUX_MANY";
    string internal constant INBOX_VAL = "INBOX_VAL";

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
    uint8 internal constant OP_SIGNEXTEND = 0x0b;

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
    uint8 internal constant OP_SHL = 0x1b;
    uint8 internal constant OP_SHR = 0x1c;
    uint8 internal constant OP_SAR = 0x1d;

    // SHA3
    uint8 internal constant OP_HASH = 0x20;
    uint8 internal constant OP_TYPE = 0x21;
    uint8 internal constant OP_ETHHASH2 = 0x22;
    uint8 internal constant OP_KECCAK_F = 0x23;
    uint8 internal constant OP_SHA256_F = 0x24;

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

    // Tuple operations
    uint8 internal constant OP_TGET = 0x50;
    uint8 internal constant OP_TSET = 0x51;
    uint8 internal constant OP_TLEN = 0x52;
    uint8 internal constant OP_XGET = 0x53;
    uint8 internal constant OP_XSET = 0x54;

    // Logging operations
    uint8 internal constant OP_BREAKPOINT = 0x60;
    uint8 internal constant OP_LOG = 0x61;

    // System operations
    uint8 internal constant OP_SEND = 0x70;
    uint8 internal constant OP_INBOX_PEEK = 0x71;
    uint8 internal constant OP_INBOX = 0x72;
    uint8 internal constant OP_ERROR = 0x73;
    uint8 internal constant OP_STOP = 0x74;
    uint8 internal constant OP_SETGAS = 0x75;
    uint8 internal constant OP_PUSHGAS = 0x76;
    uint8 internal constant OP_ERR_CODE_POINT = 0x77;
    uint8 internal constant OP_PUSH_INSN = 0x78;
    uint8 internal constant OP_PUSH_INSN_IMM = 0x79;
    // uint8 private constant OP_OPEN_INSN = 0x7a;
    uint8 internal constant OP_SIDELOAD = 0x7b;

    uint8 internal constant OP_ECRECOVER = 0x80;
    uint8 internal constant OP_ECADD = 0x81;
    uint8 internal constant OP_ECMUL = 0x82;
    uint8 internal constant OP_ECPAIRING = 0x83;

    uint8 internal constant OP_DEBUGPRINT = 0x90;

    // Buffer operations
    uint8 internal constant OP_NEWBUFFER = 0xa0;
    uint8 internal constant OP_GETBUFFER8 = 0xa1;
    uint8 internal constant OP_GETBUFFER64 = 0xa2;
    uint8 internal constant OP_GETBUFFER256 = 0xa3;
    uint8 internal constant OP_SETBUFFER8 = 0xa4;
    uint8 internal constant OP_SETBUFFER64 = 0xa5;
    uint8 internal constant OP_SETBUFFER256 = 0xa6;

    uint64 internal constant EC_PAIRING_POINT_GAS_COST = 500000;

    uint8 internal constant CODE_POINT_TYPECODE = 1;
    bytes32 internal constant CODE_POINT_ERROR =
        keccak256(abi.encodePacked(CODE_POINT_TYPECODE, uint8(0), bytes32(0)));

    uint256 internal constant SEND_SIZE_LIMIT = 10000;

    // fields
    // startMachineHash,
    // endMachineHash,
    // afterInboxHash,
    // afterMessagesHash,
    // afterLogsHash

    function returnContext(AssertionContext memory context)
        internal
        pure
        returns (
            uint64 gas,
            uint256 totalMessagesRead,
            bytes32[4] memory fields
        )
    {
        return (
            context.gas,
            context.totalMessagesRead,
            [
                Machine.hash(context.startMachine),
                Machine.hash(context.afterMachine),
                context.sendAcc,
                context.logAcc
            ]
        );
    }

    struct ValueStack {
        uint256 length;
        Value.Data[] values;
    }

    function popVal(ValueStack memory stack) internal pure returns (Value.Data memory) {
        Value.Data memory val = stack.values[stack.length - 1];
        stack.length--;
        return val;
    }

    function pushVal(ValueStack memory stack, Value.Data memory val) internal pure {
        stack.values[stack.length] = val;
        stack.length++;
    }

    struct AssertionContext {
        IBridge bridge;
        Machine.Data startMachine;
        Machine.Data afterMachine;
        uint256 totalMessagesRead;
        bytes32 sendAcc;
        bytes32 logAcc;
        uint64 gas;
        ValueStack stack;
        ValueStack auxstack;
        bool hadImmediate;
        uint8 opcode;
        bytes proof;
        uint256 offset;
        // merkle proofs for buffer
        bytes bufProof;
        bool errorOccurred;
    }

    function handleError(AssertionContext memory context) internal pure {
        context.errorOccurred = true;
    }

    function deductGas(AssertionContext memory context, uint64 amount)
        internal
        pure
        returns (bool)
    {
        if (context.afterMachine.arbGasRemaining < amount) {
            // ERROR + GAS_SET
            context.gas += ERROR_GAS_COST;
            context.afterMachine.arbGasRemaining = MAX_UINT256;
            return true;
        } else {
            context.gas += amount;
            context.afterMachine.arbGasRemaining -= amount;
            return false;
        }
    }

    function handleOpcodeError(AssertionContext memory context) internal pure {
        handleError(context);
    }

    function initializeExecutionContext(
        uint256 initialMessagesRead,
        bytes32 initialSendAcc,
        bytes32 initialLogAcc,
        bytes memory proof,
        bytes memory bproof,
        IBridge bridge
    ) internal pure returns (AssertionContext memory) {
        uint8 opCode = uint8(proof[0]);
        uint8 stackCount = uint8(proof[1]);
        uint8 auxstackCount = uint8(proof[2]);
        uint256 offset = 3;

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

        uint8 immediate = uint8(proof[offset]);
        offset += 1;

        AssertionContext memory context;
        context.bridge = bridge;
        context.startMachine = mach;
        context.afterMachine = mach.clone();
        context.totalMessagesRead = initialMessagesRead;
        context.sendAcc = initialSendAcc;
        context.logAcc = initialLogAcc;
        context.gas = 0;
        context.stack = ValueStack(stackCount, stackVals);
        context.auxstack = ValueStack(auxstackCount, auxstackVals);
        context.hadImmediate = uint8(proof[offset]) == 1;
        context.opcode = opCode;
        context.proof = proof;
        context.bufProof = bproof;
        context.errorOccurred = false;
        context.offset = offset;

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

        // Require the prover to submit the minimal number of stack items
        require(
            ((dataPopCount > 0 || !context.hadImmediate) && context.stack.length <= dataPopCount) ||
                (context.hadImmediate && dataPopCount == 0 && context.stack.length == 1),
            STACK_MANY
        );
        require(context.auxstack.length <= auxPopCount, AUX_MANY);

        // Update end machine gas remaining before running opcode
        if (context.stack.length < dataPopCount) {
            // If we have insufficient values, reject the proof unless the stack has been fully exhausted
            require(
                context.afterMachine.dataStack.hash() == Value.newEmptyTuple().hash(),
                STACK_MISSING
            );
            deductGas(context, ERROR_GAS_COST);
            // If the stack is empty, the instruction underflowed so we have hit an error
            handleError(context);
        } else if (context.auxstack.length < auxPopCount) {
            // If we have insufficient values, reject the proof unless the auxstack has been fully exhausted
            require(
                context.afterMachine.auxStack.hash() == Value.newEmptyTuple().hash(),
                AUX_MISSING
            );
            deductGas(context, ERROR_GAS_COST);
            // If the auxstack is empty, the instruction underflowed so we have hit an error
            handleError(context);
        } else if (deductGas(context, gasCost)) {
            handleError(context);
        } else {
            impl(context);
        }

        if (context.errorOccurred) {
            if (context.afterMachine.errHandlerHash == CODE_POINT_ERROR) {
                context.afterMachine.setErrorStop();
            } else {
                // Clear error
                context.errorOccurred = false;
                context.afterMachine.instructionStackHash = context.afterMachine.errHandlerHash;

                if (!(context.hadImmediate && dataPopCount == 0)) {
                    context.stack.length = 0;
                }
                context.auxstack.length = 0;
            }
        }

        // Add the stack and auxstack values to the start machine
        uint256 i = 0;

        for (i = 0; i < context.stack.length; i++) {
            context.afterMachine.addDataStackValue(context.stack.values[i]);
        }

        for (i = 0; i < context.auxstack.length; i++) {
            context.afterMachine.addAuxStackValue(context.auxstack.values[i]);
        }
    }

    function opInfo(uint256 opCode)
        internal
        pure
        virtual
        returns (
            uint256, // stack pops
            uint256, // auxstack pops
            uint64, // gas used
            function(AssertionContext memory) internal view // impl
        );
}
