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

import "./ArbERC20Bridge.sol";
import "arbos-contracts/contracts/ArbSys.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";

contract EthERC20Bridge {
    IInbox inbox;
    IBridge bridge;
    bool connectedToChain;

    modifier onlyIfConnected {
        require(connectedToChain, "NOT_CONNECTED");
        _;
    }

    modifier onlyL2 {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    constructor(IInbox _inbox) public {
        inbox = _inbox;
        bridge = IBridge(_inbox.bridge());
    }

    function connectToChain(uint256 maxGas, uint256 gasPriceBid) external payable {
        // Pay for gas
        if (msg.value > 0) {
            inbox.depositEthMessage{ value: msg.value }(address(this));
        }
        inbox.deployL2ContractPair(
            maxGas, // max gas
            gasPriceBid, // gas price
            0, // payment
            type(ArbERC20Bridge).creationCode
        );
    }

    function buddyCreated(bool successful) external onlyL2 {
        // This method must be called by the l2 system rather than a contract
        require(l2Sender() == address(0), "ONLY_SYSTEM");
        if (successful) {
            connectedToChain = true;
        }
    }

    function withdrawFromL2(
        address erc20,
        address destination,
        uint256 amount
    ) external onlyIfConnected onlyL2 {
        // This method is only callable by this contract's buddy contract on L2
        require(l2Sender() == address(this), "L2_SENDER");
        // Unsafe external call must occur below checks and effects
        require(IERC20(erc20).transfer(destination, amount));
    }

    function updateTokenName(
        address erc20,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        string memory name = ERC20(erc20).name();
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature("updateTokenName(address,string)", erc20, name)
        );
    }

    function updateTokenSymbol(
        address erc20,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        string memory symbol = ERC20(erc20).symbol();
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature("updateTokenSymbol(address,string)", erc20, symbol)
        );
    }

    function updateTokenDecimals(
        address erc20,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        uint8 decimals = ERC20(erc20).decimals();
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature("updateTokenDecimals(address,uint8)", erc20, decimals)
        );
    }

    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        require(IERC20(erc20).transferFrom(msg.sender, address(this), amount));
        // This transfers along any ETH sent for to pay for gas in L2
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature(
                "mintFromL1(address,address,uint256)",
                erc20,
                destination,
                amount
            )
        );
    }

    function l2Sender() private view returns (address) {
        return IOutbox(bridge.activeOutbox()).l2ToL1Sender();
    }
}
