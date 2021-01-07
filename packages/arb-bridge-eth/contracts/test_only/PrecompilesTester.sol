// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "../libraries/Precompiles.sol";

library PrecompilesTester {
    function keccakF(uint256[25] memory input) public pure returns (uint256[25] memory) {
        return Precompiles.keccakF(input);
    }

    function sha256Block(bytes32[2] memory inputChunk, bytes32 hashState)
        public
        pure
        returns (bytes32)
    {
        return
            bytes32(
                Precompiles.sha256Block(
                    [uint256(inputChunk[0]), uint256(inputChunk[1])],
                    uint256(hashState)
                )
            );
    }
}
