// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2012, Offchain Labs, Inc.
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

import "../arch/Machine.sol";
import "../arch/Value.sol";

contract MachineTester {
    using Machine for Machine.Data;

    function deserializeMachine(bytes memory data)
        public
        pure
        returns (bytes32)
    {
        uint256 offset;
        Machine.Data memory machine;
        (offset, machine) = Machine.deserializeMachine(data, 0);
        return Machine.hash(machine);
    }

    function addStackVal(bytes memory data1, bytes memory data2)
        public
        pure
        returns (bytes32)
    {
        uint256 offset;
        Value.Data memory val1;
        Value.Data memory val2;

        (offset, val1) = Value.deserialize(data1, 0);

        (offset, val2) = Value.deserialize(data2, 0);

        return Value.hash(Machine.addStackVal(val1, val2));
    }
}
