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

pragma solidity ^0.5.11;

import "arbos-contracts/contracts/ArbSys.sol";
import "../inbox/IGlobalInbox.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";

contract ArbBaseERC20 is ERC20, ERC20Detailed {
    constructor() public ERC20Detailed("Token Buddy", "TB", 18) {}
}

// contract PairedErc20 is ArbBaseERC20 {
//     function mint(address account, uint256 amount) external;
//     function burn(address account, uint256 amount) external;
// }

contract ArbERC20 is ArbBaseERC20 {
    function adminMint(address account, uint256 amount) public {
        _mint(account, amount);
    }

    function withdraw(address account, uint256 amount) public {
        _burn(msg.sender, amount);
        ArbSys(100).withdrawERC20(account, amount);
    }
}

contract BuddyERC20 is ArbBaseERC20 {
    address public inbox;

    constructor() public {}

    function initialize(address _rollupChain, address _inbox) public {
        inbox = _inbox;
        IGlobalInbox(_inbox).deployL2ContractPair(
            _rollupChain,
            type(ArbERC20).creationCode
        );
    }

    function mint(address account, uint256 amount) public {
        require(inbox == msg.sender, "must be authorized rollup-chain");
        _mint(account, amount);
    }

    function burn(address account, uint256 amount) public {
        require(inbox == msg.sender, "must be authorized rollup-chain");
        _burn(account, amount);
    }
}
