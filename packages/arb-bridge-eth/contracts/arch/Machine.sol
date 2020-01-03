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
        Value.HashOnly memory stackVal,
        Value.HashOnly memory valHash
    )
        internal
        pure
        returns (Value.HashOnly memory)
    {
        Value.HashOnly[] memory values = new Value.HashOnly[](2);
        values[0] = valHash;
        values[1] = stackVal;
        return Value.HashOnly(Value.hashTuple([
            Value.newHashOnly(valHash.hash),
            Value.newHashOnly(stackVal.hash)
        ]));
    }

    struct Data {
        Value.HashOnly instructionStackHash;
        Value.HashOnly dataStackHash;
        Value.HashOnly auxStackHash;
        Value.HashOnly registerHash;
        Value.HashOnly staticHash;
        Value.HashOnly errHandler;
        uint256 status;
    }

    function toString(Data memory machine) internal pure returns (string memory) {
        return string(
            abi.encodePacked(
                "Machine(",
                DebugPrint.bytes32string(machine.instructionStackHash.hash),
                ", \n",
                DebugPrint.bytes32string(machine.dataStackHash.hash),
                ", \n",
                DebugPrint.bytes32string(machine.auxStackHash.hash),
                ", \n",
                DebugPrint.bytes32string(machine.registerHash.hash),
                ", \n",
                DebugPrint.bytes32string(machine.staticHash.hash),
                ", \n",
                DebugPrint.bytes32string(machine.errHandler.hash),
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

    function addDataStackHashValue(Data memory machine, Value.HashOnly memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val);
    }

    function addAuxStackHashValue(Data memory machine, Value.HashOnly memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val);
    }

    function addDataStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val.hash());
    }

    function addAuxStackValue(Data memory machine, Value.Data memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val.hash());
    }

    function addDataStackInt(Data memory machine, uint256 val) internal pure {
        machine.dataStackHash = addStackVal(
            machine.dataStackHash,
            Value.newInt(val).hash()
        );
    }

    function machineHash(
        bytes32 instructionStackHash,
        bytes32 dataStackHash,
        bytes32 auxStackHash,
        bytes32 registerHash,
        bytes32 staticHash,
        bytes32 errHandlerHash
    )
        public
        pure
        returns (bytes32)
    {
        return hash(
            Data(
                Value.HashOnly(instructionStackHash),
                Value.HashOnly(dataStackHash),
                Value.HashOnly(auxStackHash),
                Value.HashOnly(registerHash),
                Value.HashOnly(staticHash),
                Value.HashOnly(errHandlerHash),
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
                    machine.instructionStackHash.hash,
                    machine.dataStackHash.hash,
                    machine.auxStackHash.hash,
                    machine.registerHash.hash,
                    machine.staticHash.hash,
                    machine.errHandler.hash
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

    function deserializeMachine(bytes memory data, uint256 offset) internal pure returns (uint, uint, Data memory) {
        Data memory m;
        m.status = MACHINE_EXTENSIVE;
        uint256 retVal;
        (retVal, offset, m.instructionStackHash) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.dataStackHash) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.auxStackHash) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.registerHash) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.staticHash) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.errHandler) = Value.deserializeHashOnly(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        return (0, offset, m);
    }
}
