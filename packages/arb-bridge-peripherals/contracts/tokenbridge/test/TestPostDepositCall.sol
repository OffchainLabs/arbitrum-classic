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

import "../libraries/IERC677.sol";

contract L2Called is IERC677Receiver {
    event Called(uint256 num);

    constructor() public {}

    // This function can be anything
    function postDepositHook(uint256 num) public {
        emit Called(num);
    }

    function onTokenTransfer(
        address sender,
        uint256 amount,
        bytes calldata data
    ) external override {
        uint256 num = abi.decode(data, (uint256));

        if (num == 5) {
            postDepositHook(num);
        } else if (num == 7) {
            revert("should fail because 7");
        } else if (num == 9) {
            // this should use all gas
            while (gasleft() > 0) {}
        } else {
            revert("should fail");
        }
    }
}
