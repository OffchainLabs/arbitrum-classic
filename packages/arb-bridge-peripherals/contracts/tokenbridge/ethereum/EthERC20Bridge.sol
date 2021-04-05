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
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

import "../arbitrum/ArbTokenBridge.sol";

import "./IExitLiquidityProvider.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

import "../../buddybridge/ethereum/L1Buddy.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

enum StandardTokenType { ERC20, ERC777, Custom }

contract EthERC20Bridge {
    using SafeERC20 for IERC20;

    address internal constant USED_ADDRESS = address(0x01);

    // exitNum => exitDataHash => LP
    mapping(bytes32 => address) redirectedExits;

    mapping(address => address) public customL2Tokens;

    address private l2TemplateERC777;
    address private l2TemplateERC20;
    bytes32 private cloneableProxyHash;

    address public l2ArbTokenBridgeAddress;
    IInbox public inbox;

    // assumes only ERC20 tokens are deployed.
    // can only deposit after a deploy attempt
    mapping(address => bool) public deployAttempt;

    modifier onlyL2Address {
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        require(l2ArbTokenBridgeAddress == outbox.l2ToL1Sender(), "Not from l2 buddy");
        _;
    }

    event ActivateCustomToken(uint256 indexed seqNum, address indexed l1Address, address l2Address);

    event DeployToken(
        uint256 indexed seqNum,
        address indexed l1Address,
        StandardTokenType indexed tokenType
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
        address _l2ArbTokenBridgeAddress
    ) external payable {
        require(address(l2TemplateERC20) == address(0), "already initialized");
        l2TemplateERC777 = _l2TemplateERC777;
        l2TemplateERC20 = _l2TemplateERC20;

        // bytes memory deployCode =
        //     abi.encodePacked(
        //         type(ArbTokenBridge).creationCode,
        //         abi.encode(address(this), _l2TemplateERC777, _l2TemplateERC20)
        //     );
        l2ArbTokenBridgeAddress = _l2ArbTokenBridgeAddress;
        inbox = IInbox(_inbox);
        cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);
        // TODO: this stores the creation code in state, but we don't actually need that
        // L1Buddy.initiateBuddyDeploy(_maxSubmissionCost, _maxGas, _gasPrice, deployCode);
    }

    // function handleDeploySuccess() internal override {
    //     // this deletes the codehash from state!
    //     L1Buddy.handleDeploySuccess();
    // }

    // function handleDeployFail() internal override {}

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
    ) external payable returns (uint256) {
        address l1CustomTokenAddress = msg.sender;
        // TODO: what happens if users already bridged token to L2?
        // require(!deployAttempt[l1CustomTokenAddress], "Token already deployed in L2");
        require(
            customL2Tokens[l1CustomTokenAddress] == address(0),
            "Cannot re-register a custom token address"
        );
        customL2Tokens[l1CustomTokenAddress] = l2CustomTokenAddress;

        bytes memory data =
            abi.encodeWithSelector(
                ArbTokenBridge.customTokenRegistered.selector,
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
        uint256 exitNum
    ) public {
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        address msgSender = outbox.l2ToL1Sender();

        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, msgSender, erc20, amount));
        require(redirectedExits[withdrawData] == USED_ADDRESS, "ALREADY_EXITED");
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
    ) external onlyL2Address {
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
        bytes memory deployData,
        bytes memory callHookData
    ) internal returns (uint256) {
        require(tokenType != StandardTokenType.ERC777, "777 implementation disabled");
        // TODO: check deploy attempt
        IERC20(erc20).safeTransferFrom(msg.sender, l2ArbTokenBridgeAddress, amount);
        uint256 seqNum = 0;
        {
            bytes memory data =
                abi.encodeWithSelector(
                    ArbTokenBridge.mintFromL1.selector,
                    erc20,
                    sender,
                    tokenType,
                    destination,
                    amount,
                    deployData,
                    callHookData
                );

            seqNum = inbox.createRetryableTicket{ value: msg.value }(
                l2ArbTokenBridgeAddress,
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

    function deployAndDepositAsERC20(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256) {
        return
            deployAndDeposit(
                erc20,
                destination,
                amount,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                StandardTokenType.ERC20,
                callHookData
            );
    }

    function deployAndDepositAsERC777(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256) {
        return
            deployAndDeposit(
                erc20,
                destination,
                amount,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                StandardTokenType.ERC777,
                callHookData
            );
    }

    function deployAndDeposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        StandardTokenType tokenType,
        bytes calldata callHookData
    ) internal returns (uint256) {
        require(tokenType != StandardTokenType.Custom, "Custom token should already be deployed");
        require(tokenType != StandardTokenType.ERC777, "777 disabled");
        // TODO: write to mapping that deploy attempt

        bytes memory deployData =
            abi.encode(
                callStatic(erc20, ERC20.name.selector),
                callStatic(erc20, ERC20.symbol.selector),
                callStatic(erc20, ERC20.decimals.selector)
            );

        return
            depositToken(
                erc20,
                msg.sender,
                destination,
                amount,
                RetryableTxParams(maxSubmissionCost, maxGas, gasPriceBid),
                tokenType,
                deployData,
                callHookData
            );
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
                "",
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
                "",
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
                "",
                callHookData
            );
    }

    function calculateL2ERC777Address(address erc20) external view returns (address) {
        bytes32 salt = keccak256(abi.encodePacked(erc20, l2TemplateERC777));
        return Create2.computeAddress(salt, cloneableProxyHash, l2ArbTokenBridgeAddress);
    }

    function calculateL2ERC20Address(address erc20) external view returns (address) {
        bytes32 salt = keccak256(abi.encodePacked(erc20, l2TemplateERC20));
        return Create2.computeAddress(salt, cloneableProxyHash, l2ArbTokenBridgeAddress);
    }
}
