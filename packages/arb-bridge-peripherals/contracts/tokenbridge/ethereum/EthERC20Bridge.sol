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


import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../arbitrum/ArbTokenBridge.sol";
import "../libraries/SafeERC20Namer.sol";
import "../libraries/TransferHelper.sol";
import "./IExitLiquidityProvider.sol";
import "./L1Buddy.sol";
import "./TicketFactory.sol";

contract EthERC20Bridge is L1Buddy {
    address internal constant USED_ADDRESS = address(0x01);

    TicketFactory public immutable ticketFactory;

    mapping(address => address) customL2Tokens;

    constructor(IInbox _inbox, TicketFactory _ticketFactory) public L1Buddy(_inbox) {
        ticketFactory = _ticketFactory;
    }

    function connectToChain(uint256 maxGas, uint256 gasPriceBid) external payable {
        // Pay for gas
        if (msg.value > 0) {
            inbox.depositEth{ value: msg.value }(address(this));
        }
        inbox.deployL2ContractPair(
            maxGas, // max gas
            gasPriceBid, // gas price
            0, // payment
            type(ArbTokenBridge).creationCode
        );
    }

    /**
     * @notice Notify the L2 side of the bridge that a given token has opted into a custom implementation
     * @dev Anyone can call this method repeatedly in case the L2 call fails for some reason. There's no harm in
     * allowing this to be called multiple times
     */
    function notifyCustomToken(
        address l1Address,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable {
        address l2Address = customL2Tokens[l1Address];
        require(l2Address != address(0), "NOT_REGISTERED");
        sendPairedContractTransaction(
            maxGas,
            gasPriceBid,
            abi.encodeWithSignature("customTokenRegistered(address,address)", l1Address, l2Address)
        );
    }

    function registerCustomL2Token(address l2Address) external {
        customL2Tokens[msg.sender] = l2Address;
    }

    function fastWithdrawalFromL2(
        address liquidityProvider,
        bytes memory liquidityProof,
        address erc20,
        uint256 amount,
        uint256 exitNum
    ) public onlyIfConnected {
        ticketFactory.mint(erc20, msg.sender, exitNum, liquidityProvider);

        IExitLiquidityProvider(liquidityProvider).requestLiquidity(
            msg.sender,
            erc20,
            amount,
            exitNum,
            liquidityProof
        );
    }

    function withdrawFromL2(
        uint256 exitNum,
        address erc20,
        address destination,
        uint256 amount
    ) external onlyIfConnected onlyL2 {
        // This method is only callable by this contract's buddy contract on L2
        require(l2Sender() == address(this), "L2_SENDER");

        // Unsafe external calls must occur below checks and effects
        address recipient;
        if (ticketFactory.exists(address(this), erc20, destination, exitNum)) {
            recipient = ticketFactory.burn(erc20, destination, exitNum);
        } else {
            recipient = destination;
        }
        TransferHelper.safeTransfer(erc20, recipient, amount);
    }

    function updateTokenInfo(
        address erc20,
        uint256 maxGas,
        uint256 gasPriceBid,
        bool isERC20
    ) external payable onlyIfConnected {
        string memory name = SafeERC20Namer.tokenName(erc20);
        string memory symbol = SafeERC20Namer.tokenSymbol(erc20);
        uint8 decimals = ERC20(erc20).decimals();

        bytes4 _selector = isERC20
            ? ArbTokenBridge.updateERC777TokenInfo.selector
            : ArbTokenBridge.updateERC20TokenInfo.selector;

        sendPairedContractTransaction(
            maxGas,
            gasPriceBid,
            abi.encodeWithSelector(_selector, erc20, name, symbol, decimals)
        );
    }

    function depositAsERC777(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        require(IERC20(erc20).transferFrom(msg.sender, address(this), amount));
        uint8 decimals = ERC20(erc20).decimals();
        // This transfers along any ETH sent for to pay for gas in L2
        sendPairedContractTransaction(
            maxGas,
            gasPriceBid,
            abi.encodeWithSelector(
                ArbTokenBridge.mintERC777FromL1.selector,
                erc20,
                destination,
                amount,
                decimals
            )
        );
    }

    function depositAsERC20(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        require(IERC20(erc20).transferFrom(msg.sender, address(this), amount));
        uint8 decimals = ERC20(erc20).decimals();
        // This transfers along any ETH sent for to pay for gas in L2
        sendPairedContractTransaction(
            maxGas,
            gasPriceBid,
            abi.encodeWithSelector(
                ArbTokenBridge.mintERC20FromL1.selector,
                erc20,
                destination,
                amount,
                decimals
            )
        );
    }

    function sendPairedContractTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes memory data
    ) private {
        inbox.depositEth{ value: msg.value }(address(this));
        inbox.sendContractTransaction(maxGas, gasPriceBid, address(this), 0, data);
    }
}
