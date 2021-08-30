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

import "./OneStepProofCommon.sol";

import "../libraries/Precompiles.sol";

// Originally forked from https://github.com/leapdao/solEVM-enforcer/tree/master

contract OneStepProofHash is OneStepProofCommon {

    uint64 internal constant BLAKE2BF_BASE_GAS_COST = 10;
    uint64 internal constant BLAKE2BF_ROUND_GAS_COST = 10;
    uint256 internal constant BLAKE2BF_DATA_LENGTH = 213;
    uint256 internal constant BLAKE2BF_RESULT_LENGTH = 64;

    function executeHashInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        pushVal(context.stack, Value.newInt(uint256(val.hash())));
    }

    function executeTypeInsn(AssertionContext memory context) internal pure {
        Value.Data memory val = popVal(context.stack);
        pushVal(context.stack, val.typeCodeVal());
    }

    function executeEthHash2Insn(AssertionContext memory context) internal pure {
        Value.Data memory val1 = popVal(context.stack);
        Value.Data memory val2 = popVal(context.stack);
        if (!val1.isInt() || !val2.isInt()) {
            handleOpcodeError(context);
            return;
        }
        uint256 a = val1.intVal;
        uint256 b = val2.intVal;
        uint256 c = uint256(keccak256(abi.encodePacked(a, b)));
        pushVal(context.stack, Value.newInt(c));
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

    function executeBlake2bFInsn(AssertionContext memory context) internal view {
        Value.Data memory val = popVal(context.stack);
        if (!val.isBuffer()) {
            handleOpcodeError(context);
            return;
        }

        // the full blake2f data, exactly 213 bytes, is appended to the proof
        if (context.proof.length == context.offset) {
            // we didn't pass data, which means it was invalid
            handleOpcodeError(context);
            return;
        }
        require(context.proof.length == context.offset + BLAKE2BF_DATA_LENGTH, "WRONG_BLAKE2F_BADDATA");
        bytes memory blake2fData = new bytes(BLAKE2BF_DATA_LENGTH);
        for (uint256 i = 0; i < BLAKE2BF_DATA_LENGTH; i++) {
            blake2fData[i] = context.proof[context.offset + i];
        }
        bytes32 bufferHash = Hashing.bytesToBufferHash(blake2fData, 0, BLAKE2BF_DATA_LENGTH);
        require(val.hash() == bufferHash, "WRONG_BLAKE2F_BADDATA");

        // rounds is a big-endian uint32_t we trim at 0xffff
        if (blake2fData[0] != 0 || blake2fData[1] != 0) {
            blake2fData[0] = 0x00;
            blake2fData[1] = 0x00;
            blake2fData[2] = 0xff;
            blake2fData[3] = 0xff;
        }
        uint rounds = (uint(uint8(blake2fData[2])) << 8) | uint(uint8(blake2fData[3]));

        //calculate gas
        if (deductGas(context, uint64(BLAKE2BF_ROUND_GAS_COST * rounds))) {
            // TODO: is that true for blake2F?
            // When we run out of gas, we only charge for an error + gas_set
            // That means we need to deduct the previously charged base cost here
            context.gas -= BLAKE2BF_BASE_GAS_COST;
            handleError(context);
            return;
        }

        //call ETH precompile
        bytes memory result = new bytes(BLAKE2BF_RESULT_LENGTH);
        bool success;
        assembly {
            success := staticcall(sub(gas(), 2000), 0x09, blake2fData, BLAKE2BF_DATA_LENGTH, result, BLAKE2BF_RESULT_LENGTH)
        }
        if (!success) {
                handleOpcodeError(context);
        }
        bytes32 resultHash = Hashing.bytesToBufferHash(result, 0, BLAKE2BF_RESULT_LENGTH);
        pushVal(context.stack, Value.newBuffer(resultHash));
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
        if (opCode == OP_HASH) {
            return (1, 0, 7, executeHashInsn);
        } else if (opCode == OP_TYPE) {
            return (1, 0, 3, executeTypeInsn);
        } else if (opCode == OP_ETHHASH2) {
            return (2, 0, 8, executeEthHash2Insn);
        } else if (opCode == OP_KECCAK_F) {
            return (1, 0, 600, executeKeccakFInsn);
        } else if (opCode == OP_SHA256_F) {
            return (3, 0, 250, executeSha256FInsn);
        } else if (opCode == OP_BLAKE2B_F) {
            return (1, 0, 10, executeBlake2bFInsn);
        } else {
            revert("use another contract to handle other opcodes");
        }
    }
}
