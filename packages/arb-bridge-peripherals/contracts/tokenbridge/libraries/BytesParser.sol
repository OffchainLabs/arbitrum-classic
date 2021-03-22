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

library BytesParserWithDefault {

    function toUint8(bytes memory input, uint8 defaultValue) internal pure returns (uint8) {
        if(input.length == 0) {
            return defaultValue;
        } else {
            // TODO: try catch to handle error
            return abi.decode(input, (uint8));
        }
    }

    function toString(bytes memory input, string memory defaultValue) internal pure returns (string memory) {
        if(input.length == 0) {
            return defaultValue;
        } else if (input.length == 32) {
            // TODO: remove padding and parse ascii
            return string(input);
        } else {
            // TODO: try catch to handle error
            return abi.decode(input, (string));
        }
    }
}
