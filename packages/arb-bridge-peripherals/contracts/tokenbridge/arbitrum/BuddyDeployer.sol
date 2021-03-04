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
    // L1 buddy bridge address
    BuddyBridge buddyBridge;

    constructor(address _buddyBridge) public {
        buddyBridge = BuddyBridge(_buddyBridge);
    }

    modifier onlyBuddyBridge {
        require(msg.sender == address(buddyBridge), "ONLY_BRIDGE");
        _;
    }

    function executeBuddyDeploy(address user, bytes memory deployCode)
        onlyBuddyBridge
        external
    {
        // user == L1 msg.sender
        // deployCode == type(ArbSymmetricTokenBridge).creationCode

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
        bytes memory calldataForL1 = abi.encodeWithSelector(BuddyContract.finalizeBuddyDeploy.selector, !deployFail);
        ArbSys(100).sendTxToL1(user, calldataForL1);
    }
}
