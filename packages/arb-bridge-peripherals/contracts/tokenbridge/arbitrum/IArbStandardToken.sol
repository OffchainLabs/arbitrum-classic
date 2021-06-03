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

import "./IArbToken.sol";

interface IArbStandardToken is IArbToken {
    /**
     * @notice initialize the token
     * @dev the L2 bridge assumes this does not fail or revert
     * @param _l1Address L1 address of ERC20
     * @param _data encoded symbol/name/decimal data for initial deploy
     */
    function bridgeInit(address _l1Address, bytes memory _data) external;
}
