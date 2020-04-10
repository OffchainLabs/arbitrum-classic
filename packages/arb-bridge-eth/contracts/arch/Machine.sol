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
        vals[0] = stackValHash;
        vals[1] = valHash;
        return Value.newTuple(vals);

        // return Value.HashOValue.hashTuple(tuple));
    }

    struct Data {
        Value.Data instructionStackHash;
        Value.Data dataStackHash;
        Value.Data auxStackHash;
        Value.Data registerHash;
        Value.Data staticHash;
        Value.Data errHandler;
        uint256 status;
    }

    function toString(Data memory machine) internal pure returns (string memory) {
        return string(
            abi.encodePacked(
                "Machine(",
                DebugPrint.bytes32string(Value.hash(machine.instructionStackHash).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.dataStackHash).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.auxStackHash).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.registerHash).hash),
                ", \n",
                DebugPrint.bytes32string(Value.hash(machine.staticHash).hash),
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

    function addDataStackHashValue(Data memory machine, Value.Data memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val);
    }

    function addAuxStackHashValue(Data memory machine, Value.Data memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val);
    }

    function addDataStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val);
    }

    function addAuxStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val);
    }

    function addDataStackInt(Data memory machine, uint256 val) internal pure {
        machine.dataStackHash = addStackVal(
            machine.dataStackHash,
            Value.newInt(val)
        );
    }

    function machineHash(
        Value.Data memory instructionStackHash,
        Value.Data memory dataStackHash,
        Value.Data memory auxStackHash,
        Value.Data memory registerHash,
        Value.Data memory staticHash,
        Value.Data memory errHandlerHash
    )
        internal
        pure
        returns (bytes32)
    {
        return hash(
            Data(
                instructionStackHash,
                dataStackHash,
                auxStackHash,
                registerHash,
                staticHash,
                errHandlerHash,
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
                    Value.hash(machine.instructionStackHash).hash,
                    Value.hash(machine.dataStackHash).hash,
                    Value.hash(machine.auxStackHash).hash,
                    Value.hash(machine.registerHash).hash,
                    Value.hash(machine.staticHash).hash,
                    Value.hash(machine.errHandler).hash
                )
            );
        }

    }

    function clone(Data memory machine) internal pure returns (Data memory) {
        return Data(
            machine.instructionStackHash,
            machine.dataStackHash,
            machine.auxStackHash,
            machine.registerHash,
            machine.staticHash,
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
        bool valid;
        (valid, offset, m.instructionStackHash) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.dataStackHash) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.auxStackHash) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.registerHash) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.staticHash) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        (valid, offset, m.errHandler) = Value.deserializeHashValue(data, offset);
        if (!valid) {
            return (false, offset, m);
        }
        return (true, offset, m);
    }
}
