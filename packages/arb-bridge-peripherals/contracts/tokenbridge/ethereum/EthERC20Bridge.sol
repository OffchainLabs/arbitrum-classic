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
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

import "../arbitrum/ArbTokenBridge.sol";

import "./IExitLiquidityProvider.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

import "../../buddybridge/ethereum/L1Buddy.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

enum StandardTokenType { ERC20, ERC777, Custom }

contract EthERC20Bridge {
    using SafeERC20 for IERC20;

    address internal constant USED_ADDRESS = address(0x01);

    // exitNum => exitDataHash => LP
    mapping(bytes32 => address) redirectedExits;

    mapping(address => address) public customL2Tokens;

    address private l2TemplateERC777;
    address private l2TemplateERC20;

    address owner;

    function updateOwner(address newOwner) external {
        require(msg.sender == owner, "Only owner");
        owner = newOwner;
    }

    function updateTemplates(address erc20, address erc777) external {
        require(msg.sender == owner, "Only owner");
        l2TemplateERC777 = erc777;
        l2TemplateERC20 = erc20;
    }

    function updateL2Address(address newL2Address) external {
        require(msg.sender == owner, "Only owner");
        l2Address = newL2Address;
    }

    address public l2Address;
    IInbox public inbox;

    event ActivateCustomToken(uint256 indexed seqNum, address indexed l1Address, address l2Address);

    event UpdateTokenInfo(
        uint256 indexed seqNum,
        address indexed l1Address,
        bytes name,
        bytes symbol,
        bytes decimals
    );

    event DepositToken(
        address indexed destination,
        address sender,
        uint256 indexed seqNum,
        StandardTokenType indexed tokenType,
        uint256 value,
        address tokenAddress
    );

    function initialize(
        address _inbox,
        address _l2Deployer,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPrice,
        address _l2TemplateERC777,
        address _l2TemplateERC20,
        address _l2Address
    ) external payable {
        require(address(l2TemplateERC20) == address(0), "already initialized");
        require(owner == address(0), "owner already set");
        owner = msg.sender;
        l2TemplateERC777 = _l2TemplateERC777;
        l2TemplateERC20 = _l2TemplateERC20;

        // bytes memory deployCode =
        //     abi.encodePacked(
        //         type(ArbTokenBridge).creationCode,
        //         abi.encode(address(this), _l2TemplateERC777, _l2TemplateERC20)
        //     );
        l2Address = _l2Address;
        inbox = IInbox(_inbox);
        // TODO: this stores the creation code in state, but we don't actually need that
        // L1Buddy.initiateBuddyDeploy(_maxSubmissionCost, _maxGas, _gasPrice, deployCode);
    }

    // function handleDeploySuccess() internal override {
    //     // this deletes the codehash from state!
    //     L1Buddy.handleDeploySuccess();
    // }

    // function handleDeployFail() internal override {}

    /**
     * @notice Notify the L2 side of the bridge that a given token has opted into a custom implementation
     * @dev Anyone can call this method repeatedly in case the L2 call fails for some reason. There's no harm in
     * allowing this to be called multiple times
     */
    function notifyCustomToken(
        address l1Address,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable returns (uint256) {
        address l2Address = customL2Tokens[l1Address];
        require(l2Address != address(0), "NOT_REGISTERED");
        bytes memory data =
            abi.encodeWithSignature("customTokenRegistered(address,address)", l1Address, l2Address);
        uint256 seqNum =
            inbox.createRetryableTicket{ value: msg.value }(
                l2Address,
                0,
                maxSubmissionCost,
                msg.sender,
                msg.sender,
                maxGas,
                gasPriceBid,
                data
            );
        emit ActivateCustomToken(seqNum, l1Address, l2Address);
        return seqNum;
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
    ) public {
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
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
    ) external {
        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, destination, erc20, amount));
        address exitAddress = redirectedExits[withdrawData];
        redirectedExits[withdrawData] = USED_ADDRESS;
        // Unsafe external calls must occur below checks and effects
        if (exitAddress != address(0)) {
            IERC20(erc20).safeTransfer(exitAddress, amount);
        } else {
            IERC20(erc20).safeTransfer(destination, amount);
        }
    }

    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        returns (bytes memory)
    {
        (bool success, bytes memory res) =
            targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
    }

    function updateTokenInfo(
        address erc20,
        StandardTokenType tokenType,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable returns (uint256) {
        bytes memory name = callStatic(erc20, ERC20.name.selector);
        bytes memory symbol = callStatic(erc20, ERC20.symbol.selector);
        bytes memory decimals = callStatic(erc20, ERC20.decimals.selector);

        bytes memory data =
            abi.encodeWithSelector(
                ArbTokenBridge.updateTokenInfo.selector,
                erc20,
                tokenType,
                name,
                symbol,
                decimals
            );
        uint256 seqNum =
            inbox.createRetryableTicket{ value: msg.value }(
                l2Address,
                0,
                maxSubmissionCost,
                msg.sender,
                msg.sender,
                maxGas,
                gasPriceBid,
                data
            );
        emit UpdateTokenInfo(seqNum, erc20, name, symbol, decimals);
        return seqNum;
    }

    // hacky struct to avoid stack size limit
    struct RetryableTxParams {
        uint256 maxSubmissionCost;
        uint256 maxGas;
        uint256 gasPriceBid;
    }

    function depositToken(
        address erc20,
        address sender,
        address destination,
        uint256 amount,
        RetryableTxParams memory retryableParams,
        StandardTokenType tokenType,
        bytes memory callHookData
    ) private returns (uint256) {
        require(tokenType != StandardTokenType.ERC777, "777 implementation disabled");
        IERC20(erc20).safeTransferFrom(msg.sender, l2Address, amount);
        uint256 seqNum = 0;
        {
            bytes memory decimals = callStatic(erc20, ERC20.decimals.selector);
            bytes memory data =
                abi.encodeWithSelector(
                    ArbTokenBridge.mintFromL1.selector,
                    erc20,
                    sender,
                    tokenType,
                    destination,
                    amount,
                    decimals,
                    callHookData
                );

            seqNum = inbox.createRetryableTicket{ value: msg.value }(
                l2Address,
                0,
                retryableParams.maxSubmissionCost,
                msg.sender,
                msg.sender,
                retryableParams.maxGas,
                retryableParams.gasPriceBid,
                data
            );
        }

        emit DepositToken(destination, sender, seqNum, tokenType, amount, erc20);
        return seqNum;
    }

    function depositAsERC777(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256) {
        return
            depositToken(
                erc20,
                msg.sender,
                destination,
                amount,
                RetryableTxParams(maxSubmissionCost, maxGas, gasPriceBid),
                StandardTokenType.ERC777,
                callHookData
            );
    }

    function depositAsERC20(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256) {
        return
            depositToken(
                erc20,
                msg.sender,
                destination,
                amount,
                RetryableTxParams(maxSubmissionCost, maxGas, gasPriceBid),
                StandardTokenType.ERC20,
                callHookData
            );
    }

    function depositAsCustomToken(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256) {
        // TODO: should this not be checked in the L2?
        require(customL2Tokens[erc20] != address(0), "Custom token not deployed");
        return
            depositToken(
                erc20,
                msg.sender,
                destination,
                amount,
                RetryableTxParams(maxSubmissionCost, maxGas, gasPriceBid),
                StandardTokenType.Custom,
                callHookData
            );
    }

    function calculateL2ERC777Address(address erc20) external view returns (address) {
        bytes32 salt = bytes32(uint256(erc20));
        return Clones.predictDeterministicAddress(l2TemplateERC777, salt, l2Address);
    }

    function calculateL2ERC20Address(address erc20) external view returns (address) {
        bytes32 salt = bytes32(uint256(erc20));
        return Clones.predictDeterministicAddress(l2TemplateERC20, salt, l2Address);
    }
}
