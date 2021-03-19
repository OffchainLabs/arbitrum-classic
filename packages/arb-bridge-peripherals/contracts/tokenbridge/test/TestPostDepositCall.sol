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

// import "../ethereum/EthERC20Bridge.sol";
// import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

// contract TestPostDepositCall {
//     EthERC20Bridge tokenBridge;
//     IInbox inbox;
    
//     constructor(address _tokenBridge, address erc20Template, address erc777Template, address _inbox) {
//         tokenBridge = EthErc20Bridge(_tokenBridge, erc20Template, erc777Template);
//         inbox = IInbox(_inbox);
//     }

//     function depositAndCall(
//         address erc20,
//         address destination,
//         uint256 amount,
//         uint256 maxGas,
//         uint256 gasPriceBid,
//         address l2CallDestination,
//         bytes memory data
//     ) external payable {
//         tokenBridge.depositAsERC20(erc20, destination, amount, maxGas, gasPriceBid);
//         inbox.sendContractTransaction(maxGas, gasPriceBid, l2CallDestination, 0, data);
//     }
// }

contract L2Called {
    event Called(uint256 num);
    
    constructor() public {}

    // This function can be anything
    function postDepositHook(uint256 num) public {
        emit Called(num);
    }
}
