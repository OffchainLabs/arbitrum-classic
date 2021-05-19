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

import "./StandardArbERC20.sol";
import "../libraries/ClonableBeaconProxy.sol";
import "../libraries/TokenAddressHandler.sol";
import "../ethereum/IEthERC20Bridge.sol";

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "../libraries/BytesParser.sol";

import "./IArbStandardToken.sol";
import "./IArbToken.sol";
import "./IArbTokenBridge.sol";
import "arbos-contracts/arbos/builtin/ArbSys.sol";

import "../libraries/IERC1363.sol";

contract ArbTokenBridge is ProxySetter, IArbTokenBridge, TokenAddressHandler {
    using Address for address;

    uint256 exitNum;

    bytes32 private cloneableProxyHash;

    address public templateERC20;
    address public ethTokenBridge;

    // amount of arbgas necessary to send user tokens in case
    // of the "onTokenTransfer" call consumes all available gas
    uint256 internal constant arbgasReserveIfCallRevert = 2500;

    /**
     * @notice This ensures that a method can only be called from the L1 pair of this contract
     */
    modifier onlyEthPair {
        require(msg.sender == ethTokenBridge, "ONLY_ETH_PAIR");
        _;
    }

    /**
     * @notice Initialize L2 bridge
     * @param _ethTokenBridge Address of L1 side of token bridge (EthERC20Bridge.sol)
     * @param _templateERC20 Address of template ERC20 (i.e, StandardArbERC20.sol). Used for salt in computing L2 address.
     */
    function initialize(address _ethTokenBridge, address _templateERC20) external {
        require(_ethTokenBridge != address(0), "L1 pair can't be address 0");
        require(address(ethTokenBridge) == address(0), "already init");
        templateERC20 = _templateERC20;

        ethTokenBridge = _ethTokenBridge;

        cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);
    }

    /**
     * @notice this function can only be callable by the bridge itself
     * @dev This method is inspired by EIP 677/1363 for calls to be executed after minting.
     * A reserve amount of gas is always kept in case this call reverts or uses up all gas.
     * The reserve is the amount of gas needed to catch the revert and do the necessary alternative logic.
     */
    function mintAndCall(
        IArbToken token,
        uint256 amount,
        address sender,
        address dest,
        bytes memory data
    ) external {
        require(msg.sender == address(this), "Mint can only be called by self");
        require(dest.isContract(), "Destination must be a contract");

        token.bridgeMint(dest, amount);

        // ~73 000 arbgas used to get here
        uint256 gasAvailable = gasleft() - arbgasReserveIfCallRevert;
        require(gasleft() > gasAvailable, "Mint and call gas left calculation undeflow");

        // TODO: should the operator be L1 or L2 bridge instead of the user?
        bytes4 retval =
            IERC1363Receiver(dest).onTransferReceived{ gas: gasAvailable }(
                sender,
                sender,
                amount,
                data
            );

        require(
            retval == IERC1363Receiver.onTransferReceived.selector,
            "external logic on call fail"
        );
    }

    /**
     * @notice Mint on L2 upon L1 deposit.
     * If token not yet deployed and symbol/name/decimal data is included, deploys StandardArbERC20
     * If minting a custom token whose L2 counterpart hasn't yet been deployed/registered (!) deploys a temporary StandardArbERC20 that can later be migrated to custom token.
     * @dev Callable only by the EthERC20Bridge.depositToken function. For initial deployments of a token the L1 EthERC20Bridge
     * is expected to include the deployData. If not a L1 withdrawal is automatically triggered for the user
     * @param l1ERC20 L1 address of ERC20
     * @param sender account that initiated the deposit in the L1
     * @param dest account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param amount token amount to be minted to the user
     * @param deployData encoded symbol/name/decimal data for initial deploy
     * @param callHookData optional data for external call upon minting
     */
    function mintFromL1(
        address l1ERC20,
        address sender,
        address dest,
        uint256 amount,
        bytes calldata deployData,
        bytes calldata callHookData
    ) external override onlyEthPair {
        address expectedAddress = calculateL2TokenAddress(l1ERC20);

        if (!expectedAddress.isContract()) {
            if (deployData.length > 0) {
                address deployedToken = deployToken(l1ERC20, deployData);
                assert(deployedToken == expectedAddress);
            } else {
                if (TokenAddressHandler.isCustomToken(l1ERC20)) {
                    // address handler expects a custom, but nothing deployed
                    // no custom token deployed, expectedAddress is a temporary erc20
                    expectedAddress = calculateL2ERC20TokenAddress(l1ERC20);
                    if (!expectedAddress.isContract()) {
                        // deploy erc20 temporarily, but users can migrate to custom implementation once deployed
                        bytes memory deployData =
                            abi.encode(
                                abi.encode("Temporary Migrateable Token"),
                                abi.encode("TMT"),
                                abi.encode(uint8(18))
                            );
                        address deployedAddress = deployToken(l1ERC20, deployData);
                        assert(deployedAddress == expectedAddress);
                    }
                } else {
                    // withdraw funds to user as no deployData and no contract deployed
                    // The L1 contract shouldn't let this happen!
                    // if it does happen, withdraw to sender
                    _withdraw(l1ERC20, sender, amount);
                    return;
                }
            }
        }
        // ignores deployData if token already deployed

        IArbToken token = IArbToken(expectedAddress);
        if (callHookData.length > 0) {
            bool success;
            try ArbTokenBridge(this).mintAndCall(token, amount, sender, dest, callHookData) {
                success = true;
            } catch {
                // if reverted, then credit sender's account
                try token.bridgeMint(sender, amount) {} catch {
                    // if external bridgeMint fails, withdraw user funds and return
                    _withdraw(l1ERC20, sender, amount);
                    return;
                }
                success = false;
            }
            // if success tokens got minted to dest, else to sender
            emit TokenMinted(
                l1ERC20,
                expectedAddress,
                sender,
                success ? dest : sender,
                amount,
                true
            );
            emit MintAndCallTriggered(success, sender, dest, amount, callHookData);
        } else {
            try token.bridgeMint(dest, amount) {} catch {
                // if external bridgeMint fails, withdraw user funds and return
                _withdraw(l1ERC20, sender, amount);
                return;
            }
            emit TokenMinted(l1ERC20, expectedAddress, sender, dest, amount, false);
        }
    }

    /**
     * @notice internal utility function used to deploy ERC20 tokens with the beacon proxy pattern.
     * @dev the transparent proxy implementation by OpenZeppelin can't be used if we want to be able to
     * upgrade the token logic.
     * @param l1ERC20 L1 address of ERC20
     * @param deployData encoded symbol/name/decimal data for initial deploy
     */
    function deployToken(address l1ERC20, bytes memory deployData) internal returns (address) {
        bytes32 salt = TokenAddressHandler.getCreate2Salt(l1ERC20, templateERC20);
        address createdContract = address(new ClonableBeaconProxy{ salt: salt }());

        IArbStandardToken(createdContract).bridgeInit(l1ERC20, deployData);

        emit TokenCreated(l1ERC20, createdContract);
        return createdContract;
    }

    /**
     * @notice Sets the L1 / L2 custom token pairing; called from the L1 via EthErc20Bridge.registerCustomL2Token
     * @dev this doesn't check if the L2 token is actually deployed - this way the L1 and L2 address oracles are
     * always consistent. The necessary existence checks are done before interacting with the tokens.
     * @param l1Address Address of L1 custom token implementation
     * @param l2Address Address of L2 custom token implementation
     */
    function customTokenRegistered(address l1Address, address l2Address)
        external
        override
        onlyEthPair
    {
        // This assumed token contract is initialized and ready to be used.
        TokenAddressHandler.customL2Token[l1Address] = l2Address;
        emit CustomTokenRegistered(l1Address, l2Address);
    }

    /**
     * @notice send a withdraw message to the L1 outbox
     * @dev this call is initiated by the token, ie StandardArbERC20.withdraw or WhateverCustomToken.whateverWithdrawMethod
     * @param l1ERC20 L1 address of ERC20
     * @param destination the account to be credited with the tokens
     * @param amount token amount to be withdrawn
     */
    function withdraw(
        address l1ERC20,
        address sender,
        address destination,
        uint256 amount
    ) external override returns (uint256) {
        address expectedSender = calculateL2TokenAddress(l1ERC20);

        if (!expectedSender.isContract()) {
            // if this is a TMT or a standard token deployed by the bridge before the custom token got registered
            expectedSender = calculateL2ERC20TokenAddress(l1ERC20);
            assert(expectedSender.isContract());
        }

        require(msg.sender == expectedSender, "Withdraw can only be triggered by expected sender");

        IArbToken(expectedSender).bridgeBurn(sender, amount);
        return _withdraw(l1ERC20, destination, amount);
    }

    /**
     * @notice internal utility function that encodes and calls a L2 to L1 withdrawal transaction
     * @dev this executes the withdrawal without validating the inputs.
     * It is expected that the input is validated before calling this function.
     * @param l1ERC20 L1 address of ERC20
     * @param destination the account to be credited with the tokens
     * @param amount token amount to be withdrawn
     * @return identifier used to trigger the transaction in the L1.
     */
    function _withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) internal returns (uint256) {
        uint256 id =
            ArbSys(100).sendTxToL1(
                ethTokenBridge,
                abi.encodeWithSelector(
                    IEthERC20Bridge.withdrawFromL2.selector,
                    exitNum,
                    l1ERC20,
                    destination,
                    amount
                )
            );
        exitNum++;
        emit WithdrawToken(id, l1ERC20, amount, destination, exitNum);
        return id;
    }

    /**
     * @notice If a token is bridged as a StandardArbERC20 before a custom implementation is set,
     * users can call this method via StandardArbERC20.migrate to migrate to the custom version
     * @param l1ERC20 L1 address of ERC20
     * @param sender the account that called the migration
     * @param destination the account to be credited with the tokens
     * @param amount token amount to be migrated
     */
    function migrate(
        address l1ERC20,
        address sender,
        address destination,
        uint256 amount
    ) external override {
        require(
            TokenAddressHandler.isCustomToken(l1ERC20),
            "Needs to have custom token implementation"
        );
        require(
            msg.sender == calculateL2ERC20TokenAddress(l1ERC20),
            "Migration should be called by erc20 token contract"
        );
        // burn the tokens sent from the standard implementation
        IArbToken(msg.sender).bridgeBurn(sender, amount);

        // validate custom contract is correctly setup
        address l2CustomTokenAddress = calculateL2TokenAddress(l1ERC20);
        require(l2CustomTokenAddress.isContract(), "L2 custom token must already be deployed");

        // mint tokens of custom implementation
        IArbToken(l2CustomTokenAddress).bridgeMint(destination, amount);
        emit TokenMigrated(l1ERC20, destination, amount);
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20) public view override returns (address) {
        return
            TokenAddressHandler.calculateL2TokenAddress(
                l1ERC20,
                templateERC20,
                address(this),
                cloneableProxyHash
            );
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev If there is a custom token registered with the bridge, this address won't be used.
     * @param l1ERC20 address of L1 token
     * @return L2 address of ERC20 tokens deployed by this bridge
     */
    function calculateL2ERC20TokenAddress(address l1ERC20) public view returns (address) {
        return
            TokenAddressHandler.calculateL2ERC20TokenAddress(
                l1ERC20,
                templateERC20,
                address(this),
                cloneableProxyHash
            );
    }

    /**
     * @notice utility function used in ClonableBeaconProxy.
     * @dev this method makes it possible to use ClonableBeaconProxy.creationCode without encoding constructor parameters
     * @return the token logic to be used in a proxy contract.
     */
    function getBeacon() external view override returns (address) {
        return templateERC20;
    }
}
