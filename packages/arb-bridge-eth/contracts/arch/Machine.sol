// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

import "./Marshaling.sol";

import "../libraries/DebugPrint.sol";

library Machine {
    using Hashing for Value.Data;

    // Make sure these don't conflict with Challenge.MACHINE_UNREACHABLE (currently 100)
    uint256 internal constant MACHINE_EXTENSIVE = 0;
    uint256 internal constant MACHINE_ERRORSTOP = 1;
    uint256 internal constant MACHINE_HALT = 2;

    function addStackVal(Value.Data memory stackValHash, Value.Data memory valHash)
        internal
        pure
        returns (Value.Data memory)
    {
        Value.Data[] memory vals = new Value.Data[](2);
        vals[0] = valHash;
        vals[1] = stackValHash;

        return Hashing.getTuplePreImage(vals);
    }

    struct Data {
        bytes32 instructionStackHash;
        Value.Data dataStack;
        Value.Data auxStack;
        Value.Data registerVal;
        Value.Data staticVal;
        uint256 avmGasRemaining;
        bytes32 errHandlerHash;
        uint256 status;
    }

    function toString(Data memory machine) internal pure returns (string memory) {
        return
            string(
                abi.encodePacked(
                    "Machine(",
                    DebugPrint.bytes32string(machine.instructionStackHash),
                    ", \n",
                    DebugPrint.bytes32string(machine.dataStack.hash()),
                    ", \n",
                    DebugPrint.bytes32string(machine.auxStack.hash()),
                    ", \n",
                    DebugPrint.bytes32string(machine.registerVal.hash()),
                    ", \n",
                    DebugPrint.bytes32string(machine.staticVal.hash()),
                    ", \n",
                    DebugPrint.uint2str(machine.avmGasRemaining),
                    ", \n",
                    DebugPrint.bytes32string(machine.errHandlerHash),
                    ")\n"
                )
            );
    }

    function setErrorStop(Data memory machine) internal pure {
        machine.status = MACHINE_ERRORSTOP;
    }

    function setHalt(Data memory machine) internal pure {
        machine.status = MACHINE_HALT;
    }

    function addDataStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.dataStack = addStackVal(machine.dataStack, val);
    }

    function addAuxStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.auxStack = addStackVal(machine.auxStack, val);
    }

    function hash(Data memory machine) internal pure returns (bytes32) {
        if (machine.status == MACHINE_HALT) {
            return bytes32(uint256(0));
        } else if (machine.status == MACHINE_ERRORSTOP) {
            return bytes32(uint256(1));
        } else {
            return
                keccak256(
                    abi.encodePacked(
                        machine.instructionStackHash,
                        machine.dataStack.hash(),
                        machine.auxStack.hash(),
                        machine.registerVal.hash(),
                        machine.staticVal.hash(),
                        machine.avmGasRemaining,
                        machine.errHandlerHash
                    )
                );
        }
    }

    function clone(Data memory machine) internal pure returns (Data memory) {
        return
            Data(
                machine.instructionStackHash,
                machine.dataStack,
                machine.auxStack,
                machine.registerVal,
                machine.staticVal,
                machine.avmGasRemaining,
                machine.errHandlerHash,
                machine.status
            );
    }

    function deserializeMachine(bytes memory data, uint256 offset)
        internal
        pure
        returns (
            uint256, // offset
            Data memory // machine
        )
    {
        Data memory m;
        m.status = MACHINE_EXTENSIVE;
        uint256 instructionStack;
        uint256 errHandler;
        (offset, instructionStack) = Marshaling.deserializeInt(data, offset);

        (offset, m.dataStack) = Marshaling.deserializeHashPreImage(data, offset);
        (offset, m.auxStack) = Marshaling.deserializeHashPreImage(data, offset);
        (offset, m.registerVal) = Marshaling.deserialize(data, offset);
        (offset, m.staticVal) = Marshaling.deserialize(data, offset);
        (offset, m.avmGasRemaining) = Marshaling.deserializeInt(data, offset);
        (offset, errHandler) = Marshaling.deserializeInt(data, offset);

        m.instructionStackHash = bytes32(instructionStack);
        m.errHandlerHash = bytes32(errHandler);
        return (offset, m);
    }
}
