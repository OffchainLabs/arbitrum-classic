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

import "@openzeppelin/contracts/proxy/Clones.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "../arbitrum/ArbTokenBridge.sol";

import "./IExitLiquidityProvider.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "../libraries/SafeERC20Namer.sol";

import "../../buddybridge/ethereum/L1Buddy.sol";

enum StandardTokenType { ERC20, ERC777 }

contract EthERC20Bridge is L1Buddy {
    address internal constant USED_ADDRESS = address(0x01);

    // exitNum => exitDataHash => LP
    mapping(bytes32 => address) redirectedExits;

    mapping(address => address) public customL2Tokens;

    address private immutable l2TemplateERC777;
    address private immutable l2TemplateERC20;

    event DepositERC20(
        address indexed destination,
        address sender,
        uint256 indexed seqNum,
        uint256 amount,
        address tokenAddress
    );

    event DepositERC777(
        address indexed destination,
        address sender,
        uint256 indexed seqNum,
        uint256 amount,
        address tokenAddress
    );

    event DepositCustomTokem(
        address indexed destination,
        address sender,
        uint256 indexed seqNum,
        uint256 value,
        address tokenAddress
    );

    constructor(
        address _inbox,
        address _l2Deployer,
        uint256 _maxGas,
        uint256 _gasPrice,
        address _l2TemplateERC777,
        address _l2TemplateERC20
    ) public payable L1Buddy(_inbox, _l2Deployer) {
        l2TemplateERC777 = _l2TemplateERC777;
        l2TemplateERC20 = _l2TemplateERC20;

        bytes memory deployCode =
            abi.encodePacked(
                type(ArbTokenBridge).creationCode,
                abi.encode(address(this), _l2TemplateERC20, _l2TemplateERC777)
            );

        // TODO: this stores the creation code in state, but we don't actually need that
        L1Buddy.initiateBuddyDeploy(_maxGas, _gasPrice, deployCode);
    }

    function handleDeploySuccess() internal override {
        // this deletes the codehash from state!
        L1Buddy.handleDeploySuccess();
    }

    function handleDeployFail() internal override {}

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
        require(
            customL2Tokens[msg.sender] == address(0),
            "Cannot re-register a custom token address"
        );
        customL2Tokens[msg.sender] = l2Address;
    }

    function fastWithdrawalFromL2(
        address liquidityProvider,
        bytes memory liquidityProof,
        address erc20,
        uint256 amount,
        uint256 exitNum
    ) public onlyIfConnected {
        IOutbox outbox = IOutbox(L1Buddy.inbox.bridge().activeOutbox());
        address msgSender = outbox.l2ToL1Sender();

        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, msgSender, erc20, amount));
        require(redirectedExits[withdrawData] == address(0), "ALREADY_EXITED");
        redirectedExits[withdrawData] = liquidityProvider;

        IExitLiquidityProvider(liquidityProvider).requestLiquidity(
            msgSender,
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
    ) external onlyIfConnected onlyL2Buddy {
        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, destination, erc20, amount));
        address exitAddress = redirectedExits[withdrawData];
        redirectedExits[withdrawData] = USED_ADDRESS;
        // Unsafe external calls must occur below checks and effects
        if (exitAddress != address(0)) {
            require(IERC20(erc20).transfer(exitAddress, amount));
        } else {
            require(IERC20(erc20).transfer(destination, amount));
        }
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

        bytes4 _selector =
            isERC20
                ? ArbTokenBridge.updateERC777TokenInfo.selector
                : ArbTokenBridge.updateERC20TokenInfo.selector;

        sendPairedContractTransaction(
            maxGas,
            gasPriceBid,
            abi.encodeWithSelector(_selector, erc20, name, symbol, decimals)
        );
    }

    function depositToken(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        StandardTokenType tokenType
    ) private returns (uint256) {
        require(IERC20(erc20).transferFrom(msg.sender, l2Buddy, amount));
        uint8 decimals = ERC20(erc20).decimals();
        bytes4 selector;
        if (tokenType == StandardTokenType.ERC20) {
            selector = ArbTokenBridge.mintERC20FromL1.selector;
        } else if (tokenType == StandardTokenType.ERC777) {
            selector = ArbTokenBridge.mintERC777FromL1.selector;
        }

        bytes memory data = abi.encodeWithSelector(selector, erc20, destination, amount, decimals);
        return
            inbox.createRetryableTicket(
                destination,
                0,
                maxSubmissionCost,
                msg.sender,
                msg.sender,
                maxGas,
                gasPriceBid,
                data
            );
    }

    function depositAsERC777(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable returns (uint256) {
        uint256 seqNum =
            depositToken(
                erc20,
                destination,
                amount,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                StandardTokenType.ERC777
            );
        emit DepositERC777(destination, msg.sender, seqNum, amount, erc20);
        return seqNum;
    }

    function depositAsERC20(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable returns (uint256) {
        uint256 seqNum =
            depositToken(
                erc20,
                destination,
                amount,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                StandardTokenType.ERC20
            );
        emit DepositERC20(destination, msg.sender, seqNum, amount, erc20);
        return seqNum;
    }

    function depositAsCustomToken(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable returns (uint256) {
        require(customL2Tokens[erc20] != address(0), "Custom token not deployed");
        require(IERC20(erc20).transferFrom(msg.sender, l2Buddy, amount));
        bytes memory data =
            abi.encodeWithSelector(
                ArbTokenBridge.mintCustomtokenFromL1.selector,
                erc20,
                destination,
                amount
            );
        uint256 seqNum =
            inbox.createRetryableTicket(
                destination,
                0,
                maxSubmissionCost,
                msg.sender,
                msg.sender,
                maxGas,
                gasPriceBid,
                data
            );
        emit DepositCustomTokem(destination, msg.sender, seqNum, amount, erc20);
        return seqNum;
    }

    function calculateL2ERC777Address(address erc20) external view returns (address) {
        bytes32 salt = bytes32(uint256(erc20));
        return Clones.predictDeterministicAddress(l2TemplateERC777, salt, L1Buddy.l2Buddy);
    }

    function calculateL2ERC20Address(address erc20) external view returns (address) {
        bytes32 salt = bytes32(uint256(erc20));
        return Clones.predictDeterministicAddress(l2TemplateERC20, salt, L1Buddy.l2Buddy);
    }

    // TODO: does this carry over the msg.value of the internal call implicitly?
    function sendPairedContractTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes memory data
    ) private {
        inbox.depositEth{ value: msg.value }(L1Buddy.l2Buddy);
        inbox.sendContractTransaction(maxGas, gasPriceBid, L1Buddy.l2Buddy, 0, data);
    }
}
