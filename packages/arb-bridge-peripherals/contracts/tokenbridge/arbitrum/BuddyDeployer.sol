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
    function calledFromL1() external view returns (bool);
}

contract L2Deployer {
    constructor() public {}

    event Deployed(address _sender, bool _success);

    function executeBuddyDeploy(bytes memory deployCode)
        external
        payable
    {
        // we don't want nasty address clashes
        require(ArbSys(100).calledFromL1(), "Function must be called from L1");
        address user = msg.sender;
        uint256 salt = uint256(user);
        address addr;
        bool success;
        assembly {
            addr := create2(
                callvalue(), // wei sent in call
                add(deployCode, 0x20), // do we need to actually skip this?
                mload(deployCode),
                salt
            )
            success := not(iszero(extcodesize(addr)))
        }

        // L1 callback to buddy
        bytes memory calldataForL1 = abi.encodeWithSelector(BuddyContract.finalizeBuddyDeploy.selector, success);
        ArbSys(100).sendTxToL1(user, calldataForL1);
        emit Deployed(user, success);
    }
}

abstract contract L2Buddy {
    constructor() public {
        bytes memory calldataForL1 = abi.encodeWithSelector(BuddyContract.finalizeBuddyDeploy.selector, true);
        ArbSys(100).sendTxToL1(msg.sender, calldataForL1);
    }
}
