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

import "./Value.sol";

import "../libraries/DebugPrint.sol";


library Machine {
    using Value for Value.Data;

    uint256 internal constant MACHINE_EXTENSIVE = 0;
    uint256 internal constant MACHINE_ERRORSTOP = 1;
    uint256 internal constant MACHINE_HALT = 2;

    function addStackVal(
        Value.Data memory stackValHash,
        Value.Data memory valHash
    )
        internal
        pure
        returns (Value.Data memory)
    {
        Value.Data[] memory vals = new Value.Data[](2);
        vals[0] = valHash;
        vals[1] = stackValHash;
        Value.Data memory tuple = Value.newTuple(vals);
        return Value.newHashOnly(Value.hash(tuple).hash, tuple.size);
    }

    struct Data {
        Value.Data instructionStack;
        Value.Data dataStack;
        Value.Data auxStack;
        Value.Data registerVal;
        Value.Data staticVal;
        Value.Data errHandler;
        uint256 status;
    }

    function toString(Data memory machine) internal pure returns (string memory) {
        return string(
            abi.encodePacked(
                "Machine(",
                DebugPrint.bytes32string(Value.hash(machine.instructionStack).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.dataStack).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.auxStack).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.registerVal).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.staticVal).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.errHandler).hash),
                ")\n"
            )
        );
    }

    function setExtensive(Data memory machine) internal pure {
        machine.status = MACHINE_EXTENSIVE;
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

    function addDataStackInt(Data memory machine, uint256 val) internal pure {
        machine.dataStack = addStackVal(
            machine.dataStack,
            Value.newInt(val)
        );
    }

    function machineHash(
        Value.Data memory instructionStack,
        Value.Data memory dataStack,
        Value.Data memory auxStack,
        Value.Data memory registerVal,
        Value.Data memory staticVal,
        Value.Data memory errHandler
    )
        internal
        pure
        returns (bytes32)
    {
        return hash(
            Data(
                instructionStack,
                dataStack,
                auxStack,
                registerVal,
                staticVal,
                errHandler,
                MACHINE_EXTENSIVE
            )
        );
    }

    function hash(Data memory machine) internal pure returns (bytes32) {
        if (machine.status == MACHINE_HALT) {
            return bytes32(uint(0));
        } else if (machine.status == MACHINE_ERRORSTOP) {
            return bytes32(uint(1));
        } else {
            return keccak256(
                abi.encodePacked(
                    Value.hash(machine.instructionStack).hash,
                    Value.hash(machine.dataStack).hash,
                    Value.hash(machine.auxStack).hash,
                    Value.hash(machine.registerVal).hash,
                    Value.hash(machine.staticVal).hash,
                    Value.hash(machine.errHandler).hash
                )
            );
        }

    }

    function clone(Data memory machine) internal pure returns (Data memory) {
        return Data(
            machine.instructionStack,
            machine.dataStack,
            machine.auxStack,
            machine.registerVal,
            machine.staticVal,
            machine.errHandler,
            machine.status
        );
    }

    function deserializeMachine(
        bytes memory data,
        uint256 offset
    )
        internal
        pure
        returns(
            bool, // valid
            uint256, // offset
            Data memory // machine
        )
    {
        Data memory m;
        m.status = MACHINE_EXTENSIVE;
        bytes32 hashVal;
        bool valid;
        (valid, offset, hashVal) = Value.deserializeHashed(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        m.instructionStack = Value.newHashOnly(hashVal, 1);

        (valid, offset, m.dataStack) = Value.deserializeHashPreImage(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.auxStack) = Value.deserializeHashPreImage(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, hashVal) = Value.deserializeHashed(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        m.registerVal = Value.newHashOnly(hashVal, 1);

        (valid, offset, hashVal) = Value.deserializeHashed(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        m.staticVal = Value.newHashOnly(hashVal, 1);

        (valid, offset, hashVal) = Value.deserializeHashed(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        m.errHandler = Value.newHashOnly(hashVal, 1);

        return (true, offset, m);
    }
}
