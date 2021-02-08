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
import "./OneStepProofCommon.sol";

import "../bridge/Messages.sol";

import "../libraries/Precompiles.sol";

// Originally forked from https://github.com/leapdao/solEVM-enforcer/tree/master

contract OneStepProof is IOneStepProof, OneStepProofCommon {
    using Machine for Machine.Data;
    using Hashing for Value.Data;
    using Value for Value.Data;

    uint256 private constant MAX_PAIRING_COUNT = 30;

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    function executeStep(bytes32[3] calldata _machineFields, bytes calldata proof)
        external
        view
        override
        returns (uint64 gas, bytes32[5] memory fields)
    {
        bytes memory bproof = new bytes(0);
        AssertionContext memory context =
            initializeExecutionContext(
                _machineFields[0],
                _machineFields[1],
                _machineFields[2],
                proof,
                bproof
            );

        executeOp(context);

        return returnContext(context);
    }

    /* solhint-disable no-inline-assembly */

    // Arithmetic

    function binaryMathOp(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt()) {
            handleOpcodeError(context);
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

        pushVal(context.stack, Value.newInt(c));
    }

    function binaryMathOpZero(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || val2.intVal == 0) {
            handleOpcodeError(context);
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

        pushVal(context.stack, Value.newInt(c));
    }

    function executeMathModInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || !val3.isInt() || val3.intVal == 0) {
            handleOpcodeError(context);
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

        pushVal(context.stack, Value.newInt(c));
    }

    function executeEqInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        pushVal(context.stack, Value.newBoolean(val1.hash() == val2.hash()));
    }

    function executeIszeroInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        if (!val1.isInt()) {
            pushVal(context.stack, Value.newInt(0));
        } else {
            uint256 a = val1.intVal;
            uint256 c;
            assembly {
                c := iszero(a)
            }
            pushVal(context.stack, Value.newInt(c));
        }
    }

    function executeNotInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        if (!val1.isInt()) {
            handleOpcodeError(context);
            return;
        }
        uint256 a = val1.intVal;
        uint256 c;
        assembly {
            c := not(a)
        }
        pushVal(context.stack, Value.newInt(c));
    }

    /* solhint-enable no-inline-assembly */

    // Hash

    function executeHashInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        pushVal(context.stack, Value.newInt(uint256(val.hash())));
    }

    function executeTypeInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        pushVal(context.stack, val.typeCodeVal());
    }

    function executeKeccakFInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        if (!val.isTuple() || val.tupleVal.length != 7) {
            handleOpcodeError(context);
            return;
        }

        Value.Data[] memory values = val.tupleVal;
        for (uint256 i = 0; i < 7; i++) {
            if (!values[i].isInt()) {
                handleOpcodeError(context);
                return;
            }
        }
        uint256[25] memory data;
        for (uint256 i = 0; i < 25; i++) {
            data[5 * (i % 5) + i / 5] = uint256(uint64(values[i / 4].intVal >> ((i % 4) * 64)));
        }

        data = Precompiles.keccakF(data);

        Value.Data[] memory outValues = new Value.Data[](7);
        for (uint256 i = 0; i < 7; i++) {
            outValues[i] = Value.newInt(0);
        }

        for (uint256 i = 0; i < 25; i++) {
            outValues[i / 4].intVal |= data[5 * (i % 5) + i / 5] << ((i % 4) * 64);
        }

        pushVal(context.stack, Value.newTuple(outValues));
    }

    function executeSha256FInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || !val3.isInt()) {
            handleOpcodeError(context);
            return;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c = val3.intVal;

        pushVal(context.stack, Value.newInt(Precompiles.sha256Block([b, c], a)));
    }

    // Stack ops

    function executePopInsn(AssertionContext memory context) internal pure {
        popVal(context.stack);
    }

    function executeSpushInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, context.afterMachine.staticVal);
    }

    function executeRpushInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, context.afterMachine.registerVal);
    }

    function executeRsetInsn(AssertionContext memory context) internal pure {
        context.afterMachine.registerVal = popVal(context.stack);
    }

    function executeJumpInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        if (!val.isCodePoint()) {
            handleOpcodeError(context);
            return;
        }
        context.afterMachine.instructionStackHash = val.hash();
    }

    function executeCjumpInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isCodePoint() || !val2.isInt()) {
            handleOpcodeError(context);
            return;
        }
        if (val2.intVal != 0) {
            context.afterMachine.instructionStackHash = val1.hash();
        }
    }

    function executeStackemptyInsn(AssertionContext memory context) internal pure {
        bool empty =
            context.stack.length == 0 &&
                context.afterMachine.dataStack.hash() == Value.newEmptyTuple().hash();
        pushVal(context.stack, Value.newBoolean(empty));
    }

    function executePcpushInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newHashedValue(context.startMachine.instructionStackHash, 1));
    }

    function executeAuxpushInsn(AssertionContext memory context) internal pure {
        pushVal(context.auxstack, popVal(context.stack));
    }

    function executeAuxpopInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, popVal(context.auxstack));
    }

    function executeAuxstackemptyInsn(AssertionContext memory context) internal pure {
        bool empty =
            context.auxstack.length == 0 &&
                context.afterMachine.auxStack.hash() == Value.newEmptyTuple().hash();
        pushVal(context.stack, Value.newBoolean(empty));
    }

    /* solhint-disable-next-line no-empty-blocks */
    function executeNopInsn(AssertionContext memory) internal pure {}

    function executeErrpushInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newHashedValue(context.afterMachine.errHandlerHash, 1));
    }

    function executeErrsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        if (!val.isCodePoint()) {
            handleOpcodeError(context);
            return;
        }
        context.afterMachine.errHandlerHash = val.hash();
    }

    // Dup ops

    function executeDup0Insn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        pushVal(context.stack, val);
        pushVal(context.stack, val);
    }

    function executeDup1Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        pushVal(context.stack, val2);
        pushVal(context.stack, val1);
        pushVal(context.stack, val2);
    }

    function executeDup2Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        pushVal(context.stack, val3);
        pushVal(context.stack, val2);
        pushVal(context.stack, val1);
        pushVal(context.stack, val3);
    }

    // Swap ops

    function executeSwap1Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        pushVal(context.stack, val1);
        pushVal(context.stack, val2);
    }

    function executeSwap2Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        pushVal(context.stack, val1);
        pushVal(context.stack, val2);
        pushVal(context.stack, val3);
    }

    // Tuple ops

    function executeTgetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isInt() || !val2.isTuple() || val1.intVal >= val2.valLength()) {
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, val2.tupleVal[val1.intVal]);
    }

    function executeTsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val1.isInt() || !val2.isTuple() || val1.intVal >= val2.valLength()) {
            handleOpcodeError(context);
            return;
        }
        Value.Data[] memory tupleVals = val2.tupleVal;
        tupleVals[val1.intVal] = val3;
        pushVal(context.stack, Value.newTuple(tupleVals));
    }

    function executeTlenInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        if (!val1.isTuple()) {
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, Value.newInt(val1.valLength()));
    }

    function executeXgetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory auxVal = popVal(context.auxstack);
        if (!val1.isInt() || !auxVal.isTuple() || val1.intVal >= auxVal.valLength()) {
            handleOpcodeError(context);
            return;
        }
        pushVal(context.auxstack, auxVal);
        pushVal(context.stack, auxVal.tupleVal[val1.intVal]);
    }

    function executeXsetInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory auxVal = popVal(context.auxstack);
        if (!auxVal.isTuple() || !val1.isInt() || val1.intVal >= auxVal.valLength()) {
            handleOpcodeError(context);
            return;
        }
        Value.Data[] memory tupleVals = auxVal.tupleVal;
        tupleVals[val1.intVal] = val2;
        pushVal(context.auxstack, Value.newTuple(tupleVals));
    }

    // Logging

    function executeLogInsn(AssertionContext memory context) internal pure {
        context.logAcc = keccak256(abi.encodePacked(context.logAcc, popVal(context.stack).hash()));
    }

    // System operations

    function incrementInbox(AssertionContext memory context)
        private
        pure
        returns (Value.Data memory message)
    {
        uint256 newInboxDelta;
        (context.offset, message) = Marshaling.deserialize(context.proof, context.offset);
        (context.offset, newInboxDelta) = Marshaling.deserializeInt(context.proof, context.offset);
        require(
            context.inboxDelta == Messages.addMessageToInbox(bytes32(newInboxDelta), message.hash())
        );
        context.inboxDelta = bytes32(newInboxDelta);
        return message;
    }

    function executeInboxPeekInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        if (context.afterMachine.pendingMessage.hash() != Value.newEmptyTuple().hash()) {
            context.afterMachine.pendingMessage = incrementInbox(context);
        }
        // The pending message must be a tuple of size at least 2
        pushVal(
            context.stack,
            Value.newBoolean(context.afterMachine.pendingMessage.tupleVal[1].hash() == val.hash())
        );
    }

    function executeInboxInsn(AssertionContext memory context) internal pure {
        if (context.afterMachine.pendingMessage.hash() != Value.newEmptyTuple().hash()) {
            // The pending message field is already full
            pushVal(context.stack, context.afterMachine.pendingMessage);
            context.afterMachine.pendingMessage = Value.newEmptyTuple();
        } else {
            pushVal(context.stack, incrementInbox(context));
        }
    }

    function executeSetGasInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        if (!val1.isInt()) {
            handleOpcodeError(context);
            return;
        }
        context.afterMachine.arbGasRemaining = val1.intVal;
    }

    function executePushGasInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newInt(context.afterMachine.arbGasRemaining));
    }

    function executeErrCodePointInsn(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newHashedValue(CODE_POINT_ERROR, 1));
    }

    function executePushInsnInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isInt() || !val2.isCodePoint()) {
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, Value.newCodePoint(uint8(val1.intVal), val2.hash()));
    }

    function executePushInsnImmInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val1.isInt() || !val3.isCodePoint()) {
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, Value.newCodePoint(uint8(val1.intVal), val3.hash(), val2));
    }

    function executeSideloadInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        if (!val1.isInt()) {
            handleOpcodeError(context);
            return;
        }
        Value.Data[] memory values = new Value.Data[](0);
        pushVal(context.stack, Value.newTuple(values));
    }

    function executeECRecoverInsn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        Value.Data memory val4 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || !val3.isInt() || !val4.isInt()) {
            handleOpcodeError(context);
            return;
        }
        bytes32 r = bytes32(val1.intVal);
        bytes32 s = bytes32(val2.intVal);
        if (val3.intVal != 0 && val3.intVal != 1) {
            pushVal(context.stack, Value.newInt(0));
            return;
        }
        uint8 v = uint8(val3.intVal) + 27;
        bytes32 message = bytes32(val4.intVal);
        address ret = ecrecover(message, v, r, s);
        pushVal(context.stack, Value.newInt(uint256(ret)));
    }

    /* solhint-disable no-inline-assembly */

    function executeECAddInsn(AssertionContext memory context) internal view {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        Value.Data memory val4 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || !val3.isInt() || !val4.isInt()) {
            handleOpcodeError(context);
            return;
        }
        uint256[4] memory bnAddInput = [val1.intVal, val2.intVal, val3.intVal, val4.intVal];
        uint256[2] memory ret;
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 6, bnAddInput, 0x80, ret, 0x40)
        }
        if (!success) {
            // Must end on empty tuple
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, Value.newInt(uint256(ret[1])));
        pushVal(context.stack, Value.newInt(uint256(ret[0])));
    }

    function executeECMulInsn(AssertionContext memory context) internal view {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        Value.Data memory val3 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt() || !val3.isInt()) {
            handleOpcodeError(context);
            return;
        }
        uint256[3] memory bnAddInput = [val1.intVal, val2.intVal, val3.intVal];
        uint256[2] memory ret;
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 7, bnAddInput, 0x80, ret, 0x40)
        }
        if (!success) {
            // Must end on empty tuple
            handleOpcodeError(context);
            return;
        }
        pushVal(context.stack, Value.newInt(uint256(ret[1])));
        pushVal(context.stack, Value.newInt(uint256(ret[0])));
    }

    function executeECPairingInsn(AssertionContext memory context) internal view {
        Value.Data memory val = popVal(context.stack);

        Value.Data[MAX_PAIRING_COUNT] memory items;
        uint256 count;
        for (count = 0; count < MAX_PAIRING_COUNT; count++) {
            if (!val.isTuple()) {
                handleOpcodeError(context);
                return;
            }
            Value.Data[] memory stackTupleVals = val.tupleVal;
            if (stackTupleVals.length == 0) {
                // We reached the bottom of the stack
                break;
            }
            if (stackTupleVals.length != 2) {
                handleOpcodeError(context);
                return;
            }
            items[count] = stackTupleVals[0];
            val = stackTupleVals[1];
        }

        if (deductGas(context, uint64(EC_PAIRING_POINT_GAS_COST * count))) {
            return;
        }

        if (!val.isTuple() || val.tupleVal.length != 0) {
            // Must end on empty tuple
            handleOpcodeError(context);
            return;
        }

        // Allocate the maximum amount of space we might need
        uint256[MAX_PAIRING_COUNT * 6] memory input;
        for (uint256 i = 0; i < count; i++) {
            Value.Data memory pointVal = items[i];
            if (!pointVal.isTuple()) {
                handleOpcodeError(context);
                return;
            }

            Value.Data[] memory pointTupleVals = pointVal.tupleVal;
            if (pointTupleVals.length != 6) {
                handleOpcodeError(context);
                return;
            }

            for (uint256 j = 0; j < 6; j++) {
                if (!pointTupleVals[j].isInt()) {
                    handleOpcodeError(context);
                    return;
                }
            }
            input[i * 6] = pointTupleVals[0].intVal;
            input[i * 6 + 1] = pointTupleVals[1].intVal;
            input[i * 6 + 2] = pointTupleVals[3].intVal;
            input[i * 6 + 3] = pointTupleVals[2].intVal;
            input[i * 6 + 4] = pointTupleVals[5].intVal;
            input[i * 6 + 5] = pointTupleVals[4].intVal;
        }

        uint256 inputSize = count * 6 * 0x20;
        uint256[1] memory out;
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 8, input, inputSize, out, 0x20)
        }

        if (!success) {
            handleOpcodeError(context);
            return;
        }

        pushVal(context.stack, Value.newBoolean(out[0] != 0));
    }

    /* solhint-enable no-inline-assembly */

    function executeErrorInsn(AssertionContext memory context) internal pure {
        handleOpcodeError(context);
    }

    function executeStopInsn(AssertionContext memory context) internal pure {
        context.afterMachine.setHalt();
    }

    function executeNewBuffer(AssertionContext memory context) internal pure {
        pushVal(context.stack, Value.newBuffer(keccak256(abi.encodePacked(bytes32(0)))));
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
        if (opCode == OP_ADD || opCode == OP_MUL || opCode == OP_SUB) {
            return (2, 0, 3, binaryMathOp);
        } else if (opCode == OP_DIV || opCode == OP_MOD) {
            return (2, 0, 4, binaryMathOpZero);
        } else if (opCode == OP_SDIV || opCode == OP_SMOD) {
            return (2, 0, 7, binaryMathOpZero);
        } else if (opCode == OP_ADDMOD || opCode == OP_MULMOD) {
            return (3, 0, 4, executeMathModInsn);
        } else if (opCode == OP_EXP) {
            return (2, 0, 25, binaryMathOp);
        } else if (opCode == OP_SIGNEXTEND) {
            return (2, 0, 7, binaryMathOp);
        } else if (
            opCode == OP_LT ||
            opCode == OP_GT ||
            opCode == OP_SLT ||
            opCode == OP_SGT ||
            opCode == OP_AND ||
            opCode == OP_OR ||
            opCode == OP_XOR
        ) {
            return (2, 0, 2, binaryMathOp);
        } else if (opCode == OP_EQ) {
            return (2, 0, 2, executeEqInsn);
        } else if (opCode == OP_ISZERO) {
            return (1, 0, 1, executeIszeroInsn);
        } else if (opCode == OP_NOT) {
            return (1, 0, 1, executeNotInsn);
        } else if (opCode == OP_BYTE || opCode == OP_SHL || opCode == OP_SHR || opCode == OP_SAR) {
            return (2, 0, 4, binaryMathOp);
        } else if (opCode == OP_HASH) {
            return (1, 0, 7, executeHashInsn);
        } else if (opCode == OP_TYPE) {
            return (1, 0, 3, executeTypeInsn);
        } else if (opCode == OP_ETHHASH2) {
            return (2, 0, 8, binaryMathOp);
        } else if (opCode == OP_KECCAK_F) {
            return (1, 0, 600, executeKeccakFInsn);
        } else if (opCode == OP_SHA256_F) {
            return (3, 0, 250, executeSha256FInsn);
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
        } else if (opCode == OP_INBOX_PEEK) {
            return (1, 0, 40, executeInboxPeekInsn);
        } else if (opCode == OP_INBOX) {
            return (0, 0, 40, executeInboxInsn);
        } else if (opCode == OP_ERROR) {
            return (0, 0, 5, executeErrorInsn);
        } else if (opCode == OP_STOP) {
            return (0, 0, 10, executeStopInsn);
        } else if (opCode == OP_SETGAS) {
            return (1, 0, 1, executeSetGasInsn);
        } else if (opCode == OP_PUSHGAS) {
            return (0, 0, 1, executePushGasInsn);
        } else if (opCode == OP_ERR_CODE_POINT) {
            return (0, 0, 25, executeErrCodePointInsn);
        } else if (opCode == OP_PUSH_INSN) {
            return (2, 0, 25, executePushInsnInsn);
        } else if (opCode == OP_PUSH_INSN_IMM) {
            return (3, 0, 25, executePushInsnImmInsn);
        } else if (opCode == OP_SIDELOAD) {
            return (1, 0, 10, executeSideloadInsn);
        } else if (opCode == OP_ECRECOVER) {
            return (4, 0, 20000, executeECRecoverInsn);
        } else if (opCode == OP_ECADD) {
            return (4, 0, 3500, executeECAddInsn);
        } else if (opCode == OP_ECMUL) {
            return (3, 0, 82000, executeECMulInsn);
        } else if (opCode == OP_ECPAIRING) {
            return (1, 0, 1000, executeECPairingInsn);
        } else if (opCode == OP_DEBUGPRINT) {
            return (1, 0, 1, executePopInsn);
        } else if (opCode == OP_NEWBUFFER) {
            return (0, 0, 1, executeNewBuffer);
        } else if (opCode >= OP_GETBUFFER8 && opCode <= OP_SETBUFFER256) {
            revert("use another contract to handle buffer opcodes");
        } else {
            return (0, 0, 0, executeErrorInsn);
        }
    }
}
