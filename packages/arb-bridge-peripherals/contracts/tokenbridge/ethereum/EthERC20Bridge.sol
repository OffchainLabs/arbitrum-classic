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

import "@openzeppelin/contracts/utils/Create2.sol";
import "../libraries/ClonableBeaconProxy.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

import "../arbitrum/IArbTokenBridge.sol";

import "./IExitLiquidityProvider.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

import "../../buddybridge/ethereum/L1Buddy.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

import "./IEthERC20Bridge.sol";

contract EthERC20Bridge is IEthERC20Bridge {
    using SafeERC20 for IERC20;

    address internal constant USED_ADDRESS = address(0x01);

    // exitNum => exitDataHash => LP
    mapping(bytes32 => address) public redirectedExits;

    mapping(address => address) public customL2Tokens;

    // TODO: delete __placeholder__
    // Can't delete now as it will break the storage layout of the proxy contract
    address private __placeholder__;
    address private l2TemplateERC20;
    bytes32 private cloneableProxyHash;

    address public l2ArbTokenBridgeAddress;
    IInbox public inbox;

    // assumes only ERC20 tokens are deployed.
    // can only deposit after a deploy attempt
    mapping(address => bool) public override hasTriedDeploy;

    modifier onlyL2Address {
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        require(l2ArbTokenBridgeAddress == outbox.l2ToL1Sender(), "Not from l2 buddy");
        _;
    }

    function initialize(
        address _inbox,
        address _l2Deployer,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPrice,
        address _l2TemplateERC20,
        address _l2ArbTokenBridgeAddress
    ) external payable {
        require(address(l2TemplateERC20) == address(0), "already initialized");
        l2TemplateERC20 = _l2TemplateERC20;

        l2ArbTokenBridgeAddress = _l2ArbTokenBridgeAddress;
        inbox = IInbox(_inbox);
        cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);
    }

    /**
     * @notice Update the L1 custom token registry directly and L2 side via  a retryable ticket
     * @dev
     */
    function registerCustomL2Token(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) external payable override returns (uint256) {
        address l1CustomTokenAddress = msg.sender;
        // Token must be registering for the first time, or retrying the same address
        require(
            customL2Tokens[l1CustomTokenAddress] == address(0) ||
                customL2Tokens[l1CustomTokenAddress] == l2CustomTokenAddress,
            "Cannot register a different custom token address"
        );
        customL2Tokens[l1CustomTokenAddress] = l2CustomTokenAddress;

        bytes memory data =
            abi.encodeWithSelector(
                IArbTokenBridge.customTokenRegistered.selector,
                l1CustomTokenAddress,
                l2CustomTokenAddress
            );
        uint256 seqNum =
            inbox.createRetryableTicket{ value: msg.value }(
                l2ArbTokenBridgeAddress,
                0,
                maxSubmissionCost,
                refundAddress,
                refundAddress,
                maxGas,
                gasPriceBid,
                data
            );
        emit ActivateCustomToken(seqNum, l1CustomTokenAddress, l2CustomTokenAddress);
        return seqNum;
    }

    function fastWithdrawalFromL2(
        address liquidityProvider,
        bytes memory liquidityProof,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        uint256 maxFee
    ) external override {
        // TODO: this only allows withdrawal if you were the withdrawal initiator
        bytes32 withdrawData = encodeWithdrawal(exitNum, msg.sender, erc20, amount);

        require(redirectedExits[withdrawData] != USED_ADDRESS, "ALREADY_EXITED");
        redirectedExits[withdrawData] = liquidityProvider;

        uint256 balancePrior = IERC20(erc20).balanceOf(msg.sender);

        // Liquidity provider is responsible for validating if this is a valid exit
        IExitLiquidityProvider(liquidityProvider).requestLiquidity(
            msg.sender,
            erc20,
            amount,
            exitNum,
            liquidityProof
        );

        uint256 balancePost = IERC20(erc20).balanceOf(msg.sender);

        // User must be sent at least (amount - maxFee) or execution reverts
        require(
            SafeMath.sub(balancePost, balancePrior) >= SafeMath.sub(amount, maxFee),
            "User did not get credited with enough tokens"
        );

        emit WithdrawRedirected(msg.sender, liquidityProvider, erc20, amount, exitNum);
    }

    function withdrawFromL2(
        uint256 exitNum,
        address erc20,
        address withdrawInitiator,
        uint256 amount
    ) external override onlyL2Address {
        bytes32 withdrawData = encodeWithdrawal(exitNum, withdrawInitiator, erc20, amount);
        address exitAddress = redirectedExits[withdrawData];
        redirectedExits[withdrawData] = USED_ADDRESS;
        address dest = exitAddress != address(0) ? exitAddress : withdrawInitiator;
        // Unsafe external calls must occur below checks and effects
        IERC20(erc20).safeTransfer(dest, amount);

        emit WithdrawExecuted(withdrawInitiator, dest, erc20, amount, exitNum);
    }

    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        returns (bytes memory)
    {
        (bool success, bytes memory res) =
            targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
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
        bytes memory deployData,
        bytes memory callHookData
    ) internal returns (uint256) {
        IERC20(erc20).safeTransferFrom(sender, l2ArbTokenBridgeAddress, amount);
        uint256 seqNum = 0;
        {
            bytes memory data =
                abi.encodeWithSelector(
                    IArbTokenBridge.mintFromL1.selector,
                    erc20,
                    sender,
                    destination,
                    amount,
                    deployData,
                    callHookData
                );

            seqNum = inbox.createRetryableTicket{ value: msg.value }(
                l2ArbTokenBridgeAddress,
                0,
                retryableParams.maxSubmissionCost,
                sender,
                sender,
                retryableParams.maxGas,
                retryableParams.gasPriceBid,
                data
            );
        }

        emit DepositToken(destination, sender, seqNum, amount, erc20);
        return seqNum;
    }

    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable override returns (uint256) {
        bytes memory deployData = "";

        // if no deploy done and no custom L2 token set
        if (!hasTriedDeploy[erc20] && customL2Tokens[erc20] == address(0)) {
            // TODO: use OZ's ERC20Metadata once available
            // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/IERC20Metadata.sol
            deployData = abi.encode(
                callStatic(erc20, ERC20.name.selector),
                callStatic(erc20, ERC20.symbol.selector),
                callStatic(erc20, ERC20.decimals.selector)
            );
            hasTriedDeploy[erc20] = true;
        }

        return
            depositToken(
                erc20,
                msg.sender,
                destination,
                amount,
                RetryableTxParams(maxSubmissionCost, maxGas, gasPriceBid),
                deployData,
                callHookData
            );
    }

    function encodeWithdrawal(
        uint256 exitNum,
        address user,
        address erc20,
        uint256 amount
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(exitNum, user, erc20, amount));
    }

    function calculateL2TokenAddress(address erc20) public view override returns (address) {
        address customTokenAddr = customL2Tokens[erc20];
        if (customTokenAddr != address(0)) {
            return customTokenAddr;
        } else {
            bytes32 salt = keccak256(abi.encodePacked(erc20, l2TemplateERC20));
            return Create2.computeAddress(salt, cloneableProxyHash, l2ArbTokenBridgeAddress);
        }
    }
}
