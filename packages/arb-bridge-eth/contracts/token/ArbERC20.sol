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
import "../rollup/IInbox.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";

contract ArbERC20 is ERC20, ERC20Detailed {
    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals
    ) public ERC20Detailed(name, symbol, decimals) {}

    function mintFromL1(address account, uint256 amount) public {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == address(this));
        _mint(account, amount);
    }

    function withdraw(address destination, uint256 amount) public {
        _burn(msg.sender, amount);
        ArbSys(100).sendTxToL1(
            address(this),
            abi.encodeWithSignature("withdrawFromL2(address,uint256)", destination, amount)
        );
    }
}

contract EthERC20Escrow is Ownable {
    address public inbox;
    address public wrappedToken;
    mapping(address => bool) pairedRollups;

    constructor(address _wrappedToken) public {
        wrappedToken = _wrappedToken;
    }

    function connectToChain(
        address rollupChain,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyOwner {
        ERC20Detailed t = ERC20Detailed(wrappedToken);
        // Pay for gas
        IInbox(rollupChain).depositEthMessage.value(msg.value)(address(this));
        IInbox(rollupChain).deployL2ContractPair(
            maxGas, // max gas
            gasPriceBid, // gas price
            0, // payment
            abi.encodePacked(
                type(ArbERC20).creationCode,
                abi.encode(t.name(), t.symbol(), t.decimals())
            )
        );
    }

    function deposit(
        address rollupChain,
        address destination,
        uint256 amount,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable {
        require(pairedRollups[rollupChain]);
        require(IERC20(wrappedToken).transferFrom(msg.sender, address(this), amount));
        // Pay for gas
        IInbox(rollupChain).depositEthMessage.value(msg.value)(address(this));
        IInbox(rollupChain).sendL2Message(
            abi.encodePacked(
                maxGas,
                gasPriceBid,
                uint256(address(this)),
                uint256(0),
                abi.encodeWithSignature("mintFromL1(address,uint256)", destination, amount)
            )
        );
    }

    function withdrawFromL2(address destination, uint256 amount) external {
        require(pairedRollups[msg.sender]);
        require(IERC20(wrappedToken).transfer(destination, amount));
    }
}
