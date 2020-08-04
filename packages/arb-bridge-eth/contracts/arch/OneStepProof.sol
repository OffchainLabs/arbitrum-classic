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
import "../inbox/Messages.sol";
import "../libraries/Keccak.sol";

// Sourced from https://github.com/leapdao/solEVM-enforcer/tree/master

library OneStepProof {
    using Machine for Machine.Data;
    using Hashing for Value.Data;
    using Value for Value.Data;
    using OneStepProof for ValueStack;
    using OneStepProof for AssertionContext;

    uint256 private constant SEND_SIZE_LIMIT = 10000;

    uint256 private constant MAX_UINT256 = ((1 << 128) + 1) * ((1 << 128) - 1);

    struct ValueStack {
        uint256 length;
        Value.Data[] values;
    }

    function popVal(ValueStack memory stack)
        internal
        pure
        returns (Value.Data memory)
    {
        Value.Data memory val = stack.values[stack.length - 1];
        stack.length--;
        return val;
    }

    function pushVal(ValueStack memory stack, Value.Data memory val)
        internal
        pure
    {
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
    }

    function handleError(AssertionContext memory context) internal pure {
        if (context.afterMachine.errHandlerHash == CODE_POINT_ERROR) {
            context.afterMachine.setErrorStop();
        } else {
            context.afterMachine.instructionStackHash = context
                .afterMachine
                .errHandlerHash;
        }
    }

    function handleOpcodeError(AssertionContext memory context) internal pure {
        context.handleError();
        // Also clear the stack and auxstack
        context.stack.length = 0;
        context.auxstack.length = 0;
    }

    function initializeInboxExecutionContext(
        bytes32 inboxAcc,
        bytes32 messagesAcc,
        bytes32 logsAcc,
        bytes memory proof,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _msgData
    ) internal pure returns (AssertionContext memory) {
        AssertionContext memory context = initializeExecutionContext(
            inboxAcc,
            messagesAcc,
            logsAcc,
            proof
        );

        context.inboxMessageHash = Messages.messageHash(
            _kind,
            _sender,
            _blockNumber,
            _timestamp,
            _inboxSeqNum,
            keccak256(_msgData)
        );

        context.inboxMessage = Messages.messageValue(
            _kind,
            _blockNumber,
            _timestamp,
            _sender,
            _inboxSeqNum,
            _msgData
        );
    }

    function initializeExecutionContext(
        bytes32 inboxAcc,
        bytes32 messagesAcc,
        bytes32 logsAcc,
        bytes memory proof
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

        uint8 immediate = uint8(proof[offset]);
        uint8 opCode = uint8(proof[offset + 1]);
        offset += 2;
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
            immediate == 1,
            opCode,
            proof,
            offset
        );

        require(
            immediate == 0 || immediate == 1,
            "Proof had bad operation type"
        );
        Value.Data memory cp;
        if (immediate == 0) {
            cp = Value.newCodePoint(
                uint8(opCode),
                context.startMachine.instructionStackHash
            );
        } else {
            // If we have an immediate, there must be at least one stack value
            require(stackVals.length > 0, "no immediate value");
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

    function executeOp(AssertionContext memory context) internal pure {
        (
            uint256 dataPopCount,
            uint256 auxPopCount,
            uint64 gasCost,
            function(AssertionContext memory) internal pure impl
        ) = opInfo(context.opcode);
        context.gas = gasCost;

        // Update end machine gas remaining before running opcode
        // No need to overflow check since the check for whether we
        // have sufficient gas fixes the overflow case
        context.afterMachine.arbGasRemaining =
            context.afterMachine.arbGasRemaining -
            gasCost;

        if (context.startMachine.arbGasRemaining < gasCost) {
            context.afterMachine.arbGasRemaining = MAX_UINT256;
            context.handleError();
            return;
        }

        if (context.stack.length < dataPopCount) {
            // If we have insufficient values, reject the proof unless the stack has been fully exhausted
            require(
                context.afterMachine.dataStack.hash() ==
                    Value.newEmptyTuple().hash(),
                "stack item missing from proof"
            );
            // If the stack is empty, the instruction underflowed so we have hit an error
            context.handleError();
            return;
        }

        if (context.auxstack.length < auxPopCount) {
            // If we have insufficient values, reject the proof unless the auxstack has been fully exhausted
            require(
                context.afterMachine.auxStack.hash() ==
                    Value.newEmptyTuple().hash(),
                "auxstack item missing from proof"
            );
            // If the auxstack is empty, the instruction underflowed so we have hit an error
            context.handleError();
            return;
        }

        // Require the prover to submit the minimal number of stack items
        require(
            ((dataPopCount > 0 || !context.hadImmediate) &&
                context.stack.length == dataPopCount) ||
                (context.hadImmediate &&
                    dataPopCount == 0 &&
                    context.stack.length == 1),
            "too many stack items"
        );
        require(
            context.auxstack.length == auxPopCount,
            "too many auxstack items"
        );

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

    // Arithmetic

    function binaryMathOp(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        if (!val1.isInt() || !val2.isInt()) {
            context.handleOpcodeError();
            return;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;

        uint256 c;
        if (context.opcode == OP_ADD) {
            assembly {
                c := add(a, b)
            }
        } else if (context.opcode == OP_MUL) {
            assembly {
                c := mul(a, b)
            }
        } else if (context.opcode == OP_SUB) {
            assembly {
                c := sub(a, b)
            }
        } else if (context.opcode == OP_EXP) {
            assembly {
                c := exp(a, b)
            }
        } else if (context.opcode == OP_SIGNEXTEND) {
            assembly {
                c := signextend(a, b)
            }
        } else if (context.opcode == OP_LT) {
            assembly {
                c := lt(a, b)
            }
        } else if (context.opcode == OP_GT) {
            assembly {
                c := gt(a, b)
            }
        } else if (context.opcode == OP_SLT) {
            assembly {
                c := slt(a, b)
            }
        } else if (context.opcode == OP_SGT) {
            assembly {
                c := sgt(a, b)
            }
        } else if (context.opcode == OP_AND) {
            assembly {
                c := and(a, b)
            }
        } else if (context.opcode == OP_OR) {
            assembly {
                c := or(a, b)
            }
        } else if (context.opcode == OP_XOR) {
            assembly {
                c := xor(a, b)
            }
        } else if (context.opcode == OP_BYTE) {
            assembly {
                c := byte(a, b)
            }
        } else if (context.opcode == OP_SHL) {
            assembly {
                c := shl(a, b)
            }
        } else if (context.opcode == OP_SHR) {
            assembly {
                c := shr(a, b)
            }
        } else if (context.opcode == OP_SAR) {
            assembly {
                c := sar(a, b)
            }
        } else if (context.opcode == OP_ETHHASH2) {
            c = uint256(keccak256(abi.encodePacked(a, b)));
        } else {
            assert(false);
        }

        context.stack.pushVal(Value.newInt(c));
    }

    function binaryMathOpZero(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        if (!val1.isInt() || !val2.isInt() || val2.intVal == 0) {
            context.handleOpcodeError();
            return;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;

        uint256 c;
        if (context.opcode == OP_DIV) {
            assembly {
                c := div(a, b)
            }
        } else if (context.opcode == OP_SDIV) {
            assembly {
                c := sdiv(a, b)
            }
        } else if (context.opcode == OP_MOD) {
            assembly {
                c := mod(a, b)
            }
        } else if (context.opcode == OP_SMOD) {
            assembly {
                c := smod(a, b)
            }
        } else {
            assert(false);
        }

        context.stack.pushVal(Value.newInt(c));
    }

    function executeMathModInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        if (
            !val1.isInt() || !val2.isInt() || !val3.isInt() || val3.intVal == 0
        ) {
            context.handleOpcodeError();
            return;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 m = val3.intVal;

        uint256 c;

        if (context.opcode == OP_ADDMOD) {
            assembly {
                c := addmod(a, b, m)
            }
        } else if (context.opcode == OP_MULMOD) {
            assembly {
                c := mulmod(a, b, m)
            }
        } else {
            assert(false);
        }

        context.stack.pushVal(Value.newInt(c));
    }

    function executeEqInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        context.stack.pushVal(Value.newBoolean(val1.hash() == val2.hash()));
    }

    function executeIszeroInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        if (!val1.isInt()) {
            context.stack.pushVal(Value.newInt(0));
        } else {
            uint256 a = val1.intVal;
            uint256 c;
            assembly {
                c := iszero(a)
            }
            context.stack.pushVal(Value.newInt(c));
        }
    }

    function executeNotInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        if (!val1.isInt()) {
            context.handleOpcodeError();
            return;
        }
        uint256 a = val1.intVal;
        uint256 c;
        assembly {
            c := not(a)
        }
        context.stack.pushVal(Value.newInt(c));
    }

    /* solhint-enable no-inline-assembly */

    // Hash

    function executeHashInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        context.stack.pushVal(Value.newInt(uint256(val.hash())));
    }

    function executeTypeInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        context.stack.pushVal(val.typeCodeVal());
    }

    function executeKeccakFInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        if (!val.isTuple() || val.tupleVal.length != 7) {
            context.handleOpcodeError();
            return;
        }

        Value.Data[] memory values = val.tupleVal;
        for (uint256 i = 0; i < 7; i++) {
            if (!values[i].isInt()) {
                context.handleOpcodeError();
                return;
            }
        }
        uint256[25] memory data;
        for (uint256 i = 0; i < 25; i++) {
            data[5 * (i % 5) + i / 5] = uint256(
                uint64(values[i / 4].intVal >> ((i % 4) * 64))
            );
        }

        data = Keccak.keccak_f(data);

        Value.Data[] memory outValues = new Value.Data[](7);
        for (uint256 i = 0; i < 7; i++) {
            outValues[i] = Value.newInt(0);
        }

        for (uint256 i = 0; i < 25; i++) {
            outValues[i / 4].intVal |=
                data[5 * (i % 5) + i / 5] <<
                ((i % 4) * 64);
        }

        context.stack.pushVal(Value.newTuple(outValues));
    }

    // Stack ops

    function executePopInsn(AssertionContext memory context) internal pure {
        context.stack.popVal();
    }

    function executeSpushInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(context.afterMachine.staticVal);
    }

    function executeRpushInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(context.afterMachine.registerVal);
    }

    function executeRsetInsn(AssertionContext memory context) internal pure {
        context.afterMachine.registerVal = context.stack.popVal();
    }

    function executeJumpInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        if (!val.isCodePoint()) {
            context.handleOpcodeError();
            return;
        }
        context.afterMachine.instructionStackHash = val.hash();
    }

    function executeCjumpInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        if (!val1.isCodePoint() || !val2.isInt()) {
            context.handleOpcodeError();
            return;
        }
        if (val2.intVal != 0) {
            context.afterMachine.instructionStackHash = val1.hash();
        }
    }

    function executeStackemptyInsn(AssertionContext memory context)
        internal
        pure
    {
        bool empty = context.stack.length == 0 &&
            context.afterMachine.dataStack.hash() ==
            Value.newEmptyTuple().hash();
        context.stack.pushVal(Value.newBoolean(empty));
    }

    function executePcpushInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(
            Value.newHashedValue(context.startMachine.instructionStackHash, 1)
        );
    }

    function executeAuxpushInsn(AssertionContext memory context) internal pure {
        context.auxstack.pushVal(context.stack.popVal());
    }

    function executeAuxpopInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(context.auxstack.popVal());
    }

    function executeAuxstackemptyInsn(AssertionContext memory context)
        internal
        pure
    {
        bool empty = context.auxstack.length == 0 &&
            context.afterMachine.auxStack.hash() ==
            Value.newEmptyTuple().hash();
        context.stack.pushVal(Value.newBoolean(empty));
    }

    function executeNopInsn(AssertionContext memory) internal pure {}

    function executeErrpushInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(
            Value.newHashedValue(context.afterMachine.errHandlerHash, 1)
        );
    }

    function executeErrsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        if (!val.isCodePoint()) {
            context.handleOpcodeError();
            return;
        }
        context.afterMachine.errHandlerHash = val.hash();
    }

    // Dup ops

    function executeDup0Insn(AssertionContext memory context) internal pure {
        Value.Data memory val = context.stack.popVal();
        context.stack.pushVal(val);
        context.stack.pushVal(val);
    }

    function executeDup1Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        context.stack.pushVal(val2);
        context.stack.pushVal(val1);
        context.stack.pushVal(val2);
    }

    function executeDup2Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        context.stack.pushVal(val3);
        context.stack.pushVal(val2);
        context.stack.pushVal(val1);
        context.stack.pushVal(val3);
    }

    // Swap ops

    function executeSwap1Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        context.stack.pushVal(val1);
        context.stack.pushVal(val2);
    }

    function executeSwap2Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        context.stack.pushVal(val1);
        context.stack.pushVal(val2);
        context.stack.pushVal(val3);
    }

    // Tuple ops

    function executeTgetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        if (
            !val1.isInt() || !val2.isTuple() || val1.intVal >= val2.valLength()
        ) {
            context.handleOpcodeError();
            return;
        }
        context.stack.pushVal(val2.tupleVal[val1.intVal]);
    }

    function executeTsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        if (
            !val1.isInt() || !val2.isTuple() || val1.intVal >= val2.valLength()
        ) {
            context.handleOpcodeError();
            return;
        }
        Value.Data[] memory tupleVals = val2.tupleVal;
        tupleVals[val1.intVal] = val3;
        context.stack.pushVal(Value.newTuple(tupleVals));
    }

    function executeTlenInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        if (!val1.isTuple()) {
            context.handleOpcodeError();
            return;
        }
        context.stack.pushVal(Value.newInt(val1.valLength()));
    }

    function executeXgetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory auxVal = context.auxstack.popVal();
        if (
            !val1.isInt() ||
            !auxVal.isTuple() ||
            val1.intVal >= auxVal.valLength()
        ) {
            context.handleOpcodeError();
            return;
        }
        context.auxstack.pushVal(auxVal);
        context.stack.pushVal(auxVal.tupleVal[val1.intVal]);
    }

    function executeXsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory auxVal = context.auxstack.popVal();
        if (
            !auxVal.isTuple() ||
            !val1.isInt() ||
            val1.intVal >= auxVal.valLength()
        ) {
            context.handleOpcodeError();
            return;
        }
        Value.Data[] memory tupleVals = auxVal.tupleVal;
        tupleVals[val1.intVal] = val2;
        context.auxstack.pushVal(Value.newTuple(tupleVals));
    }

    // Logging

    function executeLogInsn(AssertionContext memory context) internal pure {
        context.logAcc = keccak256(
            abi.encodePacked(context.logAcc, context.stack.popVal().hash())
        );
    }

    // System operations

    function executeSendInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        if (val1.size > SEND_SIZE_LIMIT || !val1.isValidTypeForSend()) {
            context.handleOpcodeError();
            return;
        }
        context.messageAcc = keccak256(
            abi.encodePacked(context.messageAcc, val1.hash())
        );
    }

    function executeInboxInsn(AssertionContext memory context) internal pure {
        require(context.inboxMessageHash != 0, "must supply message");
        context.stack.pushVal(context.inboxMessage);
        context.inboxAcc = Messages.addMessageToInbox(
            context.inboxAcc,
            context.inboxMessageHash
        );
    }

    function executeSetGasInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = context.stack.popVal();
        if (!val1.isInt()) {
            context.handleOpcodeError();
            return;
        }
        context.afterMachine.arbGasRemaining = val1.intVal;
    }

    function executePushGasInsn(AssertionContext memory context) internal pure {
        context.stack.pushVal(
            Value.newInt(context.afterMachine.arbGasRemaining)
        );
    }

    function executeErrCodePointInsn(AssertionContext memory context)
        internal
        pure
    {
        context.stack.pushVal(Value.newHashedValue(CODE_POINT_ERROR, 1));
    }

    function executePushInsnInsn(AssertionContext memory context)
        internal
        pure
    {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        if (!val1.isInt() || !val2.isCodePoint()) {
            context.handleOpcodeError();
            return;
        }
        context.stack.pushVal(
            Value.newCodePoint(uint8(val1.intVal), val2.hash())
        );
    }

    function executePushInsnImmInsn(AssertionContext memory context)
        internal
        pure
    {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        if (!val1.isInt() || !val3.isCodePoint()) {
            context.handleOpcodeError();
            return;
        }
        context.stack.pushVal(
            Value.newCodePoint(uint8(val1.intVal), val3.hash(), val2)
        );
    }

    function executeSideloadInsn(AssertionContext memory context)
        internal
        pure
    {
        Value.Data[] memory values = new Value.Data[](0);
        context.stack.pushVal(Value.newTuple(values));
    }

    function executeECRecoverInsn(AssertionContext memory context)
        internal
        pure
    {
        Value.Data memory val1 = context.stack.popVal();
        Value.Data memory val2 = context.stack.popVal();
        Value.Data memory val3 = context.stack.popVal();
        Value.Data memory val4 = context.stack.popVal();
        if (!val1.isInt() || !val2.isInt() || !val3.isInt() || !val4.isInt()) {
            context.handleOpcodeError();
            return;
        }
        bytes32 r = bytes32(val1.intVal);
        bytes32 s = bytes32(val2.intVal);
        if (val3.intVal != 0 && val3.intVal != 1) {
            context.stack.pushVal(Value.newInt(0));
            return;
        }
        uint8 v = uint8(val3.intVal) + 27;
        bytes32 message = bytes32(val4.intVal);
        address ret = ecrecover(message, v, r, s);
        context.stack.pushVal(Value.newInt(uint256(ret)));
    }

    function executeErrorInsn(AssertionContext memory context) internal pure {
        context.handleOpcodeError();
    }

    function executeStopInsn(AssertionContext memory context) internal pure {
        context.afterMachine.setHalt();
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

    uint8 private constant CODE_POINT_TYPECODE = 1;
    bytes32 private constant CODE_POINT_ERROR = keccak256(
        abi.encodePacked(CODE_POINT_TYPECODE, uint8(0), bytes32(0))
    );

    function opInfo(uint256 opCode)
        internal
        pure
        returns (
            uint256, // stack pops
            uint256, // auxstack pops
            uint64, // gas used
            function(AssertionContext memory) internal pure // impl
        )
    {
        if (opCode == OP_ADD) {
            return (2, 0, 3, binaryMathOp);
        } else if (opCode == OP_MUL) {
            return (2, 0, 3, binaryMathOp);
        } else if (opCode == OP_SUB) {
            return (2, 0, 3, binaryMathOp);
        } else if (opCode == OP_DIV) {
            return (2, 0, 4, binaryMathOpZero);
        } else if (opCode == OP_SDIV) {
            return (2, 0, 7, binaryMathOpZero);
        } else if (opCode == OP_MOD) {
            return (2, 0, 4, binaryMathOpZero);
        } else if (opCode == OP_SMOD) {
            return (2, 0, 7, binaryMathOpZero);
        } else if (opCode == OP_ADDMOD) {
            return (3, 0, 4, executeMathModInsn);
        } else if (opCode == OP_MULMOD) {
            return (3, 0, 4, executeMathModInsn);
        } else if (opCode == OP_EXP) {
            return (2, 0, 25, binaryMathOp);
        } else if (opCode == OP_SIGNEXTEND) {
            return (2, 0, 7, binaryMathOp);
        } else if (opCode == OP_LT) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_GT) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_SLT) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_SGT) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_EQ) {
            return (2, 0, 2, executeEqInsn);
        } else if (opCode == OP_ISZERO) {
            return (1, 0, 1, executeIszeroInsn);
        } else if (opCode == OP_AND) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_OR) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_XOR) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_NOT) {
            return (1, 0, 1, executeNotInsn);
        } else if (opCode == OP_BYTE) {
            return (2, 0, 4, binaryMathOp);
        } else if (opCode == OP_SHL) {
            return (2, 0, 4, binaryMathOp);
        } else if (opCode == OP_SHR) {
            return (2, 0, 4, binaryMathOp);
        } else if (opCode == OP_SAR) {
            return (2, 0, 4, binaryMathOp);
        } else if (opCode == OP_HASH) {
            return (1, 0, 7, executeHashInsn);
        } else if (opCode == OP_TYPE) {
            return (1, 0, 3, executeTypeInsn);
        } else if (opCode == OP_ETHHASH2) {
            return (2, 0, 8, binaryMathOp);
        } else if (opCode == OP_KECCAK_F) {
            return (1, 0, 600, executeKeccakFInsn);
        } else if (opCode == OP_POP) {
            return (1, 0, 1, executePopInsn);
        } else if (opCode == OP_SPUSH) {
            return (0, 0, 1, executeSpushInsn);
        } else if (opCode == OP_RPUSH) {
            return (0, 0, 1, executeRpushInsn);
        } else if (opCode == OP_RSET) {
            return (1, 0, 2, executeRsetInsn);
        } else if (opCode == OP_JUMP) {
            return (1, 0, 4, executeJumpInsn);
        } else if (opCode == OP_CJUMP) {
            return (2, 0, 4, executeCjumpInsn);
        } else if (opCode == OP_STACKEMPTY) {
            return (0, 0, 2, executeStackemptyInsn);
        } else if (opCode == OP_PCPUSH) {
            return (0, 0, 1, executePcpushInsn);
        } else if (opCode == OP_AUXPUSH) {
            return (1, 0, 1, executeAuxpushInsn);
        } else if (opCode == OP_AUXPOP) {
            return (0, 1, 1, executeAuxpopInsn);
        } else if (opCode == OP_AUXSTACKEMPTY) {
            return (0, 0, 2, executeAuxstackemptyInsn);
        } else if (opCode == OP_NOP) {
            return (0, 0, 1, executeNopInsn);
        } else if (opCode == OP_ERRPUSH) {
            return (0, 0, 1, executeErrpushInsn);
        } else if (opCode == OP_ERRSET) {
            return (1, 0, 1, executeErrsetInsn);
        } else if (opCode == OP_DUP0) {
            return (1, 0, 1, executeDup0Insn);
        } else if (opCode == OP_DUP1) {
            return (2, 0, 1, executeDup1Insn);
        } else if (opCode == OP_DUP2) {
            return (3, 0, 1, executeDup2Insn);
        } else if (opCode == OP_SWAP1) {
            return (2, 0, 1, executeSwap1Insn);
        } else if (opCode == OP_SWAP2) {
            return (3, 0, 1, executeSwap2Insn);
        } else if (opCode == OP_TGET) {
            return (2, 0, 2, executeTgetInsn);
        } else if (opCode == OP_TSET) {
            return (3, 0, 40, executeTsetInsn);
        } else if (opCode == OP_TLEN) {
            return (1, 0, 2, executeTlenInsn);
        } else if (opCode == OP_XGET) {
            return (1, 1, 3, executeXgetInsn);
        } else if (opCode == OP_XSET) {
            return (2, 1, 41, executeXsetInsn);
        } else if (opCode == OP_BREAKPOINT) {
            return (0, 0, 100, executeNopInsn);
        } else if (opCode == OP_LOG) {
            return (1, 0, 100, executeLogInsn);
        } else if (opCode == OP_SEND) {
            return (1, 0, 100, executeSendInsn);
        } else if (opCode == OP_INBOX) {
            return (0, 0, 40, executeInboxInsn);
        } else if (opCode == OP_ERROR) {
            return (0, 0, 5, executeErrorInsn);
        } else if (opCode == OP_STOP) {
            return (0, 0, 10, executeStopInsn);
        } else if (opCode == OP_SETGAS) {
            return (1, 0, 0, executeSetGasInsn);
        } else if (opCode == OP_PUSHGAS) {
            return (0, 0, 1, executePushGasInsn);
        } else if (opCode == OP_ERR_CODE_POINT) {
            return (0, 0, 25, executeErrCodePointInsn);
        } else if (opCode == OP_PUSH_INSN) {
            return (2, 0, 25, executePushInsnInsn);
        } else if (opCode == OP_PUSH_INSN_IMM) {
            return (3, 0, 25, executePushInsnImmInsn);
        } else if (opCode == OP_SIDELOAD) {
            return (0, 0, 10, executeSideloadInsn);
        } else if (opCode == OP_ECRECOVER) {
            return (4, 0, 20000, executeECRecoverInsn);
        } else {
            return (0, 0, 0, executeErrorInsn);
        }
    }
}
