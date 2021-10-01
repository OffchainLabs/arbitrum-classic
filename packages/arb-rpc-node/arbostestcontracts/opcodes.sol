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

pragma solidity >=0.4.21 <0.7.0;

contract OpCodes {
    function getBlockHash() external returns (bytes32) {
        return blockhash(block.number - 1);
    }

    function getNestedOrigin(address conn) external returns (address) {
        return OpCodes(conn).getOrigin();
    }

    function getNestedSend(address conn) external returns (address) {
        return OpCodes(conn).getSender();
    }

    function getSender() external returns (address) {
        return msg.sender;
    }

    function getOrigin() external returns (address) {
        return tx.origin;
    }

    function getGasLeft() external returns (uint256) {
        return gasleft();
    }
}
