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
import "@openzeppelin/contracts/utils/Address.sol";
import "../libraries/ClonableBeaconProxy.sol";
import "../libraries/TokenAddressHandler.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

import "../arbitrum/IArbTokenBridge.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

import "../../buddybridge/ethereum/L1Buddy.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

import "./IEthERC20Bridge.sol";

/**
 * @title Layer 1 contract for bridging ERC20s and custom fungible tokens
 * @notice This contract handles token deposits, holds the escrowed tokens on layer 1, and (ulimately) finalizes withdrawals.
 * @dev Custom tokens that are sufficiently "weird," (i.e., dynamic supply adjustment, say) should use their own, custom bridge.
 * All messages to layer 2 use the inbox's createRetryableTicket method.
 */
contract EthERC20Bridge is IEthERC20Bridge, TokenAddressHandler {
    using SafeERC20 for IERC20;
    using Address for address;

    address internal constant USED_ADDRESS = address(0x01);

    mapping(bytes32 => address) public redirectedExits;

    address private l2TemplateERC20;
    bytes32 private cloneableProxyHash;

    address public l2ArbTokenBridgeAddress;
    IInbox public inbox;
    address public owner;

    // can only deposit after a deploy attempt
    mapping(address => bool) public override hasTriedDeploy;

    /**
     * @notice This ensures that a method can only be called from the L2 pair of this contract
     */
    modifier onlyL2Address {
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        require(l2ArbTokenBridgeAddress == outbox.l2ToL1Sender(), "Not from l2 buddy");
        _;
    }

    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    function setOwner(address _owner) external onlyOwner {
        owner = _owner;
    }

    /**
     * @notice Initialize L1 bridge
     * @param _inbox Address of Arbitrum chain's L1 Inbox.sol contract used to submit transactions to the L2
     * @param _l2TemplateERC20 Address of template ERC20 (i.e, StandardArbERC20.sol). Used for salt in computing L2 address.
     * @param _l2ArbTokenBridgeAddress Address of L2 side of token bridge (ArbTokenBridge.sol)
     * @param _owner Bridge owner that is able to force set token pairs
     */
    function initialize(
        address _inbox,
        address _l2TemplateERC20,
        address _l2ArbTokenBridgeAddress,
        address _owner
    ) external payable {
        require(address(l2TemplateERC20) == address(0), "already initialized");
        l2TemplateERC20 = _l2TemplateERC20;

        l2ArbTokenBridgeAddress = _l2ArbTokenBridgeAddress;
        inbox = IInbox(_inbox);
        cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);
        owner = _owner;
    }

    function _registerCustomL2Token(
        address l1CustomTokenAddress,
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress,
        uint256 value
    ) internal returns (uint256) {
        // Token must be registering for the first time, or retrying the same address
        require(
            !TokenAddressHandler.isCustomToken(l1CustomTokenAddress) ||
                TokenAddressHandler.customL2Token[l1CustomTokenAddress] == l2CustomTokenAddress,
            "Cannot register a different custom token address"
        );
        TokenAddressHandler.customL2Token[l1CustomTokenAddress] = l2CustomTokenAddress;

        bytes memory data =
            abi.encodeWithSelector(
                IArbTokenBridge.customTokenRegistered.selector,
                l1CustomTokenAddress,
                l2CustomTokenAddress
            );
        uint256 seqNum =
            inbox.createRetryableTicket{ value: value }(
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

    /**
     * @notice Called by a custom token on L1 to register with a previously deployed custom token on L2.
     * The L1 contract should conform to ICustomToken.sol; L2 contract should conform to IArbCustomToken.sol.
     * @dev If the L2 side hasn't yet been deployed, a safe, temporary fallback scenario will take place
     * (see ArbTokenBridge.customTokenRegistered). But please, save yourself and the trouble, and just deploy the L2 contract first.
     * @param l2CustomTokenAddress  L2 address of previously deployed custom token contract
     * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
     * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param gasPriceBid Gas price for L2 execution
     * @param refundAddress  Address to refund overbid for maxSubmissionCost and/or maxGas*gasPriceBid execution
     */
    function registerCustomL2Token(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) external payable override returns (uint256) {
        return
            _registerCustomL2Token(
                msg.sender,
                l2CustomTokenAddress,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                refundAddress,
                msg.value
            );
    }

    function forceRegisterCustomL2Token(
        address l1CustomTokenAddress,
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) external payable onlyOwner returns (uint256) {
        return
            _registerCustomL2Token(
                l1CustomTokenAddress,
                l2CustomTokenAddress,
                maxSubmissionCost,
                maxGas,
                gasPriceBid,
                refundAddress,
                msg.value
            );
    }

    /**
     * @notice Allows a user to redirect their right to claim a withdrawal to another address
     * @dev This method also allows you to make an arbitrary call after the transfer, similar to ERC677
     * @param initialDestination address the L2 withdrawal call initially set as the destination.
     * @param erc20 L1 token address
     * @param amount token amount (should match amount in previously-initiated withdrawal)
     * @param exitNum Sequentially increasing exit counter determined by the L2 bridge
     * @param data optional data for external call upon transfering the exit
     */
    function transferExitAndCall(
        address initialDestination,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        address to,
        bytes calldata data
    ) external override {
        bytes32 withdrawData = encodeWithdrawal(exitNum, initialDestination, erc20, amount);
        address redirectedAddress = redirectedExits[withdrawData];
        require(redirectedAddress != USED_ADDRESS, "ALREADY_EXITED");

        address expectedSender =
            redirectedAddress == address(0) ? initialDestination : redirectedAddress;
        require(msg.sender == expectedSender, "EXPECTED_SENDER");

        redirectedExits[withdrawData] = to;

        if (data.length > 0) {
            require(to.isContract(), "TO_NOT_CONTRACT");
            bytes4 res =
                IExitTransferCallReceiver(to).onExitTransfered(expectedSender, amount, erc20, data);
            require(
                res == IExitTransferCallReceiver.onExitTransfered.selector,
                "EXTERNAL_CALL_FAIL"
            );
        }

        emit WithdrawRedirected(expectedSender, to, erc20, amount, exitNum, data.length > 0);
    }

    /**
     * @notice Finalizes a withdraw via Outbox message; callable only by ArbTokenBridge._withdraw
     * @param exitNum Sequentially increasing exit counter determined by the L2 bridge
     * @param erc20 L1 address of token being withdrawn from
     * @param initialDestination address the L2 withdrawal call initially set as the destination.
     * @param amount Token amount being withdrawn
     */
    function withdrawFromL2(
        uint256 exitNum,
        address erc20,
        address initialDestination,
        uint256 amount
    ) external override onlyL2Address {
        bytes32 withdrawData = encodeWithdrawal(exitNum, initialDestination, erc20, amount);
        address exitAddress = redirectedExits[withdrawData];
        redirectedExits[withdrawData] = USED_ADDRESS;
        address dest = exitAddress != address(0) ? exitAddress : initialDestination;
        // Unsafe external calls must occur below checks and effects
        IERC20(erc20).safeTransfer(dest, amount);

        emit WithdrawExecuted(initialDestination, dest, erc20, amount, exitNum);
    }

    /**
     * @notice utility function used to perform external read-only calls.
     * @dev the result is returned even if the call failed, the L2 is expected to
     * identify and deal with this.
     * @return result bytes, even if the call failed.
     */
    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        view
        returns (bytes memory)
    {
        (bool success, bytes memory res) =
            targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
    }

    /**
     * @notice Utility method that allows you to get the calldata to be submitted to the L2 for a token deposit
     * @param erc20 L1 address of ERC20
     * @param sender account initiating the L1 deposit
     * @param destination account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param amount Token Amount
     * @param callHookData optional data for external call upon minting
     * @return isDeployed if token has already been deployed to the L2
     * @return depositCalldata calldata submitted to the L2
     */
    function getDepositCalldata(
        address erc20,
        address sender,
        address destination,
        uint256 amount,
        bytes calldata callHookData
    ) public view override returns (bool isDeployed, bytes memory depositCalldata) {
        isDeployed = hasTriedDeploy[erc20];

        bytes memory deployData = "";
        if (!isDeployed && !TokenAddressHandler.isCustomToken(erc20)) {
            // TODO: use OZ's ERC20Metadata once available
            // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/IERC20Metadata.sol
            deployData = abi.encode(
                callStatic(erc20, ERC20.name.selector),
                callStatic(erc20, ERC20.symbol.selector),
                callStatic(erc20, ERC20.decimals.selector)
            );
        }

        depositCalldata = abi.encodeWithSelector(
            IArbTokenBridge.mintFromL1.selector,
            erc20,
            sender,
            destination,
            amount,
            deployData,
            callHookData
        );

        return (isDeployed, depositCalldata);
    }

    /**
     * @notice Deposit standard or custom ERC20 token. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy.
     * @param erc20 L1 address of ERC20
     * @param destination account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param amount Token Amount
     * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
     * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param gasPriceBid Gas price for L2 execution
     * @param callHookData optional data for external call upon minting
     * @return seqNum ticket ID used to redeem the retryable transaction in the L2
     * @return depositCalldataLength length of calldata submitted to the L2
     */
    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable override returns (uint256 seqNum, uint256 depositCalldataLength) {
        IERC20(erc20).safeTransferFrom(msg.sender, address(this), amount);

        bytes memory depositCalldata;
        {
            bool isDeployed;
            (isDeployed, depositCalldata) = getDepositCalldata(
                erc20,
                msg.sender,
                destination,
                amount,
                callHookData
            );

            if (!isDeployed) {
                hasTriedDeploy[erc20] = true;
            }
        }
        seqNum = inbox.createRetryableTicket{ value: msg.value }(
            l2ArbTokenBridgeAddress,
            0,
            maxSubmissionCost,
            msg.sender,
            msg.sender,
            maxGas,
            gasPriceBid,
            depositCalldata
        );

        emit DepositToken(destination, msg.sender, seqNum, amount, erc20);
        return (seqNum, depositCalldata.length);
    }

    /**
     * @notice Output unique identifier for a token withdrawal. Used for tracking fast exits.
     * @param exitNum Sequentially increasing exit counter
     * @param initialDestination address for tokens before/unless otherwise redirected  (via, i.e., a fast-withdrawal)
     * @param erc20 L1 address of token being withdrawn
     * @param amount amount of token being withdrawn
     * @return bytes hash uniquely identifying withdrawal
     */
    function encodeWithdrawal(
        uint256 exitNum,
        address initialDestination,
        address erc20,
        uint256 amount
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(exitNum, initialDestination, erc20, amount));
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L@ oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param erc20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address erc20) public view override returns (address) {
        return
            TokenAddressHandler.calculateL2TokenAddress(
                erc20,
                l2TemplateERC20,
                l2ArbTokenBridgeAddress,
                cloneableProxyHash
            );
    }
}
