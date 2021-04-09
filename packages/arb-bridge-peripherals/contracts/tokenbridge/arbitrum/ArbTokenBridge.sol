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

import "./IArbToken.sol";
import "./IArbCustomToken.sol";
import "./IArbTokenBridge.sol";
import "arbos-contracts/arbos/builtin/ArbSys.sol";

import "../libraries/IERC1363.sol";

contract ArbTokenBridge is ProxySetter, IArbTokenBridge, TokenAddressHandler {
    using Address for address;

    uint256 exitNum;

    bytes32 private cloneableProxyHash;
    address private deployBeacon;

    address public templateERC20;
    address public l1Pair;

    // amount of arbgas necessary to send user tokens in case
    // of the "onTokenTransfer" call consumes all available gas
    uint256 internal constant arbgasReserveIfCallRevert = 250000;

    modifier onlyEthPair {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == l1Pair, "ONLY_ETH_PAIR");
        _;
    }

    function initialize(address _l1Pair, address _templateERC20) external {
        require(address(l1Pair) == address(0), "already init");
        require(_l1Pair != address(0), "L1 pair can't be address 0");
        templateERC20 = _templateERC20;

        l1Pair = _l1Pair;

        cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);
    }

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

        // ~7 300 000 arbgas used to get here
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

    function handleCallHookData(
        address tokenAddress,
        uint256 amount,
        address sender,
        address dest,
        bytes memory callHookData
    ) internal {
        IArbToken token = IArbToken(tokenAddress);
        bool success;
        try ArbTokenBridge(this).mintAndCall(token, amount, sender, dest, callHookData) {
            success = true;
        } catch {
            // if reverted, then credit sender's account
            token.bridgeMint(sender, amount);
            success = false;
        }
        // if success tokens got minted to dest, else to sender
        emit TokenMinted(l1ERC20, expectedAddress, sender, success ? dest : sender, amount, true);
        emit MintAndCallTriggered(success, sender, dest, amount, callHookData);
    }

    function mintFromL1(
        address l1ERC20,
        address sender,
        address dest,
        uint256 amount,
        bytes calldata deployData,
        bytes calldata callHookData
    ) external override onlyEthPair {
        address expectedAddress =
            TokenAddressHandler.calculateL2TokenAddress(
                l1ERC20,
                templateERC20,
                address(this),
                cloneableProxyHash
            );

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

        if (callHookData.length > 0) {
            handleCallHookData(expectedAddress, amount, sender, dest, callHookData);
        } else {
            IArbToken(expectedAddress).bridgeMint(dest, amount);
            emit TokenMinted(l1ERC20, expectedAddress, sender, dest, amount, false);
        }
    }

    function deployToken(address l1ERC20, bytes memory deployData) internal returns (address) {
        address beacon = templateERC20;

        deployBeacon = beacon;
        bytes32 salt = keccak256(abi.encodePacked(l1ERC20, beacon));
        address createdContract = address(new ClonableBeaconProxy{ salt: salt }());
        deployBeacon = address(0);

        bool initSuccess = IArbToken(createdContract).bridgeInit(l1ERC20, deployData);
        assert(initSuccess);

        emit TokenCreated(l1ERC20, createdContract);
        return createdContract;
    }

    function customTokenRegistered(address l1Address, address l2Address)
        external
        override
        onlyEthPair
    {
        // This assumed token contract is initialized and ready to be used.
        TokenAddressHandler.customL2Token[l1Address] = l2Address;
        emit CustomTokenRegistered(l1Address, l2Address);
    }

    function withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) external override returns (uint256) {
        address expectedSender = calculateL2TokenAddress(l1ERC20);

        // if expected address is custom, but it got self destructed users can still withdraw
        require(
            msg.sender == expectedSender || msg.sender == calculateL2ERC20TokenAddress(l1ERC20),
            "Withdraw can only be triggered by expected sender"
        );
        return _withdraw(l1ERC20, destination, amount);
    }

    function _withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) internal returns (uint256) {
        uint256 id =
            ArbSys(100).sendTxToL1(
                l1Pair,
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

    // If a token is bridged before a custom implementation is set
    // users can call this method to migrate to the custom version
    function migrate(
        address l1ERC20,
        address target,
        address account,
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

        // this assumes the l2StandardToken has burnt the user funds
        IArbCustomToken(TokenAddressHandler.customL2Token[l1ERC20]).bridgeMint(account, amount);
        emit TokenMigrated(msg.sender, target, account, amount);
    }

    function calculateL2TokenAddress(address l1ERC20) public view override returns (address) {
        return
            TokenAddressHandler.calculateL2TokenAddress(
                l1ERC20,
                templateERC20,
                address(this),
                cloneableProxyHash
            );
    }

    function calculateL2ERC20TokenAddress(address l1ERC20) public view returns (address) {
        return
            TokenAddressHandler.calculateL2ERC20TokenAddress(
                l1ERC20,
                templateERC20,
                address(this),
                cloneableProxyHash
            );
    }

    function getBeacon() external view override returns (address) {
        return deployBeacon;
    }
}
