// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

pragma experimental ABIEncoderV2;

import "../rollup/IRollup.sol";

contract Validator {
    address owner;

    constructor() public {
        owner = msg.sender;
    }

    function executeTransactions(
        bytes[] calldata data,
        address[] calldata destination,
        uint256[] calldata amount
    ) external payable {
        require(msg.sender == owner, "ONLY_OWNER");
        uint256 numTxes = data.length;
        for (uint256 i = 0; i < numTxes; i++) {
            (bool success, ) = address(destination[i]).call{ value: amount[i] }(data[i]);
            if (!success) {
                assembly {
                    let ptr := mload(0x40)
                    let size := returndatasize()
                    returndatacopy(ptr, 0, size)
                    revert(ptr, size)
                }
            }
        }
    }

    function executeTransaction(
        bytes calldata data,
        address destination,
        uint256 amount
    ) external payable {
        require(msg.sender == owner, "ONLY_OWNER");
        (bool success, ) = destination.call{ value: amount }(data);
        if (!success) {
            assembly {
                let ptr := mload(0x40)
                let size := returndatasize()
                returndatacopy(ptr, 0, size)
                revert(ptr, size)
            }
        }
    }

    function returnOldDeposits(IRollup rollup, address payable[] calldata stakers) external {
        uint256 stakerCount = stakers.length;
        for (uint256 i = 0; i < stakerCount; i++) {
            try rollup.returnOldDeposit(stakers[i]) {} catch {}
        }
    }
}
