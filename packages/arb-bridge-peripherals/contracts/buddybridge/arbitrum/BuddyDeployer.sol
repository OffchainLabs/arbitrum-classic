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

import "../ethereum/L1Buddy.sol";
import "arbos-contracts/arbos/builtin/ArbSys.sol";

contract BuddyDeployer {
    constructor() public {}

    event Deployed(address indexed _sender, address indexed _contract, bool _success);

    function executeBuddyDeploy(bytes memory contractInitCode) external payable {
        // we don't want nasty address clashes
        require(ArbSys(100).isTopLevelCall(), "Function must be called from L1");
        address user = msg.sender;
        uint256 salt = uint256(user);
        address addr;
        bool success;
        assembly {
            addr := create2(
                callvalue(), // wei sent in call
                add(contractInitCode, 0x20), // skip 32 bytes from rlp encoding length of bytearray
                mload(contractInitCode),
                salt
            )
            success := not(iszero(extcodesize(addr)))
        }

        // L1 callback to buddy
        /*
            If it calls back to an EOA the L1 call just won't execute as there is no function matching
            the selector, neither a fallback function to be executed.
        */
        bytes memory calldataForL1 =
            abi.encodeWithSelector(L1Buddy.finalizeBuddyDeploy.selector, success);
        ArbSys(100).sendTxToL1(user, calldataForL1);
        emit Deployed(user, addr, success);
    }
}
