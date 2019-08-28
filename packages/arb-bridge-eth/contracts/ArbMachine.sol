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

import "solidity-bytes-utils/contracts/BytesLib.sol";
import "./ArbValue.sol";


library ArbMachine {

    using BytesLib for bytes;
    using ArbValue for ArbValue.Value;

    uint internal constant MACHINE_EXTENSIVE = uint(0);
    uint internal constant MACHINE_ERRORSTOP = uint(1);
    uint internal constant MACHINE_HALT = uint(2);

    function addStackVal(
        ArbValue.HashOnlyValue memory stackVal,
        ArbValue.HashOnlyValue memory valHash
    )
        internal
        pure
        returns (ArbValue.HashOnlyValue memory)
    {
        ArbValue.HashOnlyValue[] memory values = new ArbValue.HashOnlyValue[](2);
        values[0] = valHash;
        values[1] = stackVal;
        return ArbValue.HashOnlyValue(ArbValue.hashTupleValue([
            ArbValue.newHashOnlyValue(valHash.hash),
            ArbValue.newHashOnlyValue(stackVal.hash)
        ]));
    }

    struct Machine {
        ArbValue.HashOnlyValue instructionStackHash;
        ArbValue.HashOnlyValue dataStackHash;
        ArbValue.HashOnlyValue auxStackHash;
        ArbValue.HashOnlyValue registerHash;
        ArbValue.HashOnlyValue staticHash;
        ArbValue.HashOnlyValue errHandler;
        uint status;
    }

    function char(byte b) internal pure returns (byte c) {
        if (uint8(b) < 10) {
            return byte(uint8(b) + 0x30);
        } else {
            return byte(uint8(b) + 0x57);
        }
    }

    function bytes32string(bytes32 b32) internal pure returns (string memory out) {
        bytes memory s = new bytes(64);

        for (uint i = 0; i < 32; i++) {
            byte b = byte(b32[i]);
            byte hi = byte(uint8(b) / 16);
            byte lo = byte(uint8(b) - 16 * uint8(hi));
            s[i*2] = char(hi);
            s[i*2+1] = char(lo);
        }

        out = string(s);
    }

    function toString(Machine memory machine) internal pure returns (string memory) {
        return string(
            abi.encodePacked(
                "Machine(",
                bytes32string(machine.instructionStackHash.hash),
                ", ",
                bytes32string(machine.dataStackHash.hash),
                ", ",
                bytes32string(machine.auxStackHash.hash),
                ", ",
                bytes32string(machine.registerHash.hash),
                ", ",
                bytes32string(machine.staticHash.hash),
                ", ",
                bytes32string(machine.errHandler.hash),
                ")"
            )
        );
    }

    function setExtensive(Machine memory machine) internal pure {
        machine.status = MACHINE_EXTENSIVE;
    }

    function setErrorStop(Machine memory machine) internal pure {
        machine.status = MACHINE_ERRORSTOP;
    }

    function setHalt(Machine memory machine) internal pure {
        machine.status = MACHINE_HALT;
    }

    function addDataStackHashValue(Machine memory machine, ArbValue.HashOnlyValue memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val);
    }

    function addAuxStackHashValue(Machine memory machine, ArbValue.HashOnlyValue memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val);
    }

    function addDataStackValue(Machine memory machine, ArbValue.Value memory val) internal pure {
        machine.dataStackHash = addStackVal(machine.dataStackHash, val.hash());
    }

    function addAuxStackValue(Machine memory machine, ArbValue.Value memory val) internal pure {
        machine.auxStackHash = addStackVal(machine.auxStackHash, val.hash());
    }

    function addDataStackInt(Machine memory machine, uint val) internal pure {
        machine.dataStackHash = addStackVal(
            machine.dataStackHash,
            ArbValue.newIntValue(val).hash()
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
            Machine(
                ArbValue.HashOnlyValue(instructionStackHash),
                ArbValue.HashOnlyValue(dataStackHash),
                ArbValue.HashOnlyValue(auxStackHash),
                ArbValue.HashOnlyValue(registerHash),
                ArbValue.HashOnlyValue(staticHash),
                ArbValue.HashOnlyValue(errHandlerHash),
                MACHINE_EXTENSIVE
            )
        );
    }

    function hash(Machine memory machine) internal pure returns (bytes32) {
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

    function clone(Machine memory machine) internal pure returns (Machine memory) {
        return Machine(
            machine.instructionStackHash,
            machine.dataStackHash,
            machine.auxStackHash,
            machine.registerHash,
            machine.staticHash,
            machine.errHandler,
            machine.status
        );
    }

    function deserializeMachine(bytes memory data, uint offset) internal pure returns (uint, uint, Machine memory) {
        Machine memory m;
        m.status = MACHINE_EXTENSIVE;
        uint retVal;
        (retVal, offset, m.instructionStackHash) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.dataStackHash) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.auxStackHash) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.registerHash) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.staticHash) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        (retVal, offset, m.errHandler) = ArbValue.deserializeHashOnlyValue(data, offset);
        if (retVal != 0) {
            return (retVal, offset, m);
        }
        return (0, offset, m);
    }
}
