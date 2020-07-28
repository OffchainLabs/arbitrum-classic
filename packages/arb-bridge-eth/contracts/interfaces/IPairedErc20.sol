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

// import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";
import "./IERC20.sol";

// contract ArbBaseERC20 is ERC20, ERC20Detailed {
//     constructor() public ERC20Detailed("Token Buddy", "TB", 18) {}
// }

contract PairedErc20 is IERC20 {
    function mint(address account, uint256 amount) external;

    function burn(address account, uint256 amount) external;
}

// contract ArbERC20 is ArbBaseERC20 {
//     function adminMint(address account, uint256 amount) public {
//         _mint(account, amount);
//     }

//     function withdraw(address account, uint256 amount) public {
//         _burn(msg.sender, amount);
//         ArbSys(100).withdrawERC20(account, amount);
//     }
// }
