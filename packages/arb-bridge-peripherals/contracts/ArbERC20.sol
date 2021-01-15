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

pragma solidity ^0.6.11;

import "arbos-contracts/contracts/ArbSys.sol";
import "arb-bridge-eth/contracts/bridge/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/IOutbox.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ArbERC20 is ERC20 {
    address public l1Address;

    constructor(
        address _l1Address,
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) public ERC20(_name, _symbol) {
        _setupDecimals(_decimals);
        l1Address = _l1Address;
    }

    function mintFromL1(address account, uint256 amount) external {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == address(this));
        _mint(account, amount);
    }

    function withdraw(address destination, uint256 amount) external {
        _burn(msg.sender, amount);
        ArbSys(100).sendTxToL1(
            address(this),
            abi.encodeWithSignature("withdrawFromL2(address,uint256)", destination, amount)
        );
    }
}

contract EthERC20Escrow {
    address public wrappedToken;
    string name;
    string symbol;
    uint8 decimals = 18;
    mapping(address => uint256) balances;

    constructor(address _wrappedToken) public {
        wrappedToken = _wrappedToken;

        try ERC20(_wrappedToken).name() returns (string memory _name) {
            name = _name;
        } catch {}

        try ERC20(_wrappedToken).symbol() returns (string memory _symbol) {
            symbol = _symbol;
        } catch {}

        try ERC20(_wrappedToken).decimals() returns (uint8 _decimals) {
            decimals = _decimals;
        } catch {}
    }

    function connectToChain(
        address rollupChain,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable {
        // Pay for gas
        if (msg.value > 0) {
            IInbox(rollupChain).depositEthMessage{ value: msg.value }(address(this));
        }
        IInbox(rollupChain).deployL2ContractPair(
            maxGas, // max gas
            gasPriceBid, // gas price
            0, // payment
            abi.encodePacked(
                type(ArbERC20).creationCode,
                abi.encode(wrappedToken, name, symbol, decimals)
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
        require(IERC20(wrappedToken).transferFrom(msg.sender, address(this), amount));
        balances[rollupChain] += amount;
        // Pay for gas
        if (msg.value > 0) {
            IInbox(rollupChain).depositEthMessage{ value: msg.value }(address(this));
        }
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
        require(balances[msg.sender] >= amount, "LOW_BALANCE");
        require(IOutbox(msg.sender).l2ToL1Sender() == address(this), "L2_SENDER");

        balances[msg.sender] -= amount;

        // Unsafe external call must occur below checks and effects
        require(IERC20(wrappedToken).transfer(destination, amount));
    }
}
