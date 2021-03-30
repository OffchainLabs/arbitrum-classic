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

interface Sys {
    // Send given Eth to dest.
    function withdrawEth(address dest) external payable;
}

interface IFailedSend {
    function withdrawFunds(address payable dest) external;
}

contract FailedSend is IFailedSend {
    constructor() public {}

    function send(address payable dest) external payable {
        IFailedSend(address(this)).withdrawFunds(dest);
        require(false, "force failure");
    }

    function withdrawFunds(address payable dest) external override {
        Sys(100).withdrawEth.value(address(this).balance)(dest);
    }
}
