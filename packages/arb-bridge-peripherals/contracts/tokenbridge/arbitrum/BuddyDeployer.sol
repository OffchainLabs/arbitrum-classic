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

import "../ethereum/BuddyBridge.sol";

// TODO: get from arb-os submodule
interface ArbSys {
    function sendTxToL1(address destAddr, bytes calldata calldataForL1) external payable;
}

contract L2Deployer {

    constructor() public {}
    event DeployedSuccess(address _sender, address _contract);
    event DeployedFail(address _sender);

    function executeBuddyDeploy(bytes memory deployCode)
        external
    {
        address user = msg.sender;
        // we don't want nasty address clashes
        require(user == tx.origin, "Can't be called by L2 contract");
        uint256 salt = uint256(user);
        address addr;
        bool deployFail;
        // TODO: is there any time we don't want callvalue() to be 0?
        assembly {
            addr := create2(
                callvalue(), // wei sent in call
                add(deployCode, 0x20),
                mload(deployCode),
                salt
            )
            deployFail := iszero(extcodesize(addr))
        }

        // L1 callback to buddy
        if(deployFail) {
            bytes memory calldataForL1 = abi.encodeWithSelector(BuddyContract.finalizeBuddyDeploy.selector, false);
            ArbSys(100).sendTxToL1(user, calldataForL1);
            emit DeployedFail(user);
        } else {
            emit DeployedSuccess(user, addr);
        }
    }
}

abstract contract L2Buddy {
    constructor() public {
        bytes memory calldataForL1 = abi.encodeWithSelector(BuddyContract.finalizeBuddyDeploy.selector, true);
        ArbSys(100).sendTxToL1(msg.sender, calldataForL1);
    }
}
