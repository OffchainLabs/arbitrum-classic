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

interface ArbSys {
    function isTopLevelCall() external view returns (bool);
}

contract TopLevel {
    event TopLevelEvent(bool indexed top);

    function isTopLevel() external {
        bool top = ArbSys(100).isTopLevelCall();
        emit TopLevelEvent(top);
    }

    function nestedNotTop() external {
        TopLevel(address(this)).isTopLevel();
    }
}
