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
import "./StandardArbERC777.sol";
import "../libraries/ClonableBeaconProxy.sol";
import "../ethereum/EthERC20Bridge.sol";

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "../libraries/BytesParser.sol";

import "./IArbToken.sol";
import "./IArbTokenBridge.sol";
import "arbos-contracts/arbos/builtin/ArbSys.sol";

interface ITransferReceiver {
    function onTokenTransfer(
        address,
        uint256,
        bytes calldata
    ) external returns (bool);
}

contract ArbTokenBridge is ProxySetter, IArbTokenBridge {
    using Address for address;

    /// @notice This mapping is from L1 address to L2 address
    mapping(address => address) public customToken;

    uint256 exitNum;

    bytes32 private cloneableProxyHash;
    address private deployBeacon;

    address public templateERC20;
    address public templateERC777;
    address public l1Pair;

    // amount of arbgas necessary to send user tokens in case
    // of the "onTokenTransfer" call consumes all available gas
    uint256 internal immutable arbgasReserveIfCallRevert = 250000;

    modifier onlyEthPair {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == l1Pair, "ONLY_ETH_PAIR");
        _;
    }

    modifier onlyFromStandardL2Token(address l1ERC20) {
        // I.e., can't be called by a custom token
        require(
            msg.sender == calculateBridgedERC777Address(l1ERC20) ||
                msg.sender == calculateBridgedERC20Address(l1ERC20),
            "NOT_FROM_STANDARD_TOKEN"
        );
        _;
    }

    modifier onlyFromL2Token(address l1ERC20) {
        // This ensures that this method can only be called by the L2 token
        require(
            msg.sender == calculateBridgedERC777Address(l1ERC20) ||
                msg.sender == calculateBridgedERC20Address(l1ERC20) ||
                msg.sender == customToken[l1ERC20],
            "NOT_FROM_TOKEN"
        );
        _;
    }
    modifier onlyToL2Token(address l1ERC20, address to) {
        // This ensures that this method can only be called by the L2 token
        require(
            to == calculateBridgedERC777Address(l1ERC20) ||
                to == calculateBridgedERC20Address(l1ERC20) ||
                to == customToken[l1ERC20],
            "NOT_TO_TOKEN"
        );
        _;
    }
    modifier noCustomToken(address l1ERC20) {
        require(customToken[l1ERC20] == address(0), "No_CUSTOM_TOKEN");
        _;
    }
    modifier ifCustomSelectedRequireCustom(address l1ERC20, StandardTokenType tokenType) {
        if (tokenType == StandardTokenType.Custom) {
            require(customToken[l1ERC20] != address(0), "No_CUSTOM_TOKEN");
        }
        _;
    }

    function initialize(
        address _l1Pair,
        address _templateERC777,
        address _templateERC20
    ) external {
        require(address(l1Pair) == address(0), "already init");
        require(_l1Pair != address(0), "L1 pair can't be address 0");
        templateERC20 = _templateERC20;
        templateERC777 = _templateERC777;

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

        // the token's transfer hook does not get triggered here
        // since the bridge already triggers a hook
        token.bridgeMint(dest, amount, "");

        // ~7 300 000 arbgas used to get here
        uint256 gasAvailable = gasleft() - arbgasReserveIfCallRevert;
        require(gasleft() > gasAvailable, "Mint and call gas left calculation undeflow");

        bool success =
            ITransferReceiver(dest).onTokenTransfer{ gas: gasAvailable }(sender, amount, data);

        require(success, "External onTokenTransfer reverted");
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
            token.bridgeMint(sender, amount, "");
            // TODO: should try to submit callHookData for the hook?
            success = false;
        }
        emit MintAndCallTriggered(success, sender, dest, amount, callHookData);
    }

    function mintFromL1(
        address l1ERC20,
        address sender,
        StandardTokenType tokenType,
        address dest,
        uint256 amount,
        bytes calldata deployData,
        bytes calldata callHookData
    ) external override onlyEthPair ifCustomSelectedRequireCustom(l1ERC20, tokenType) {
        address expectedAddress = calculateBridgeTokenAddress(l1ERC20, tokenType);

        if (!expectedAddress.isContract()) {
            if (deployData.length > 0) {
                address deployedToken = deployToken(l1ERC20, tokenType, deployData);
                require(deployedToken == expectedAddress, "Token not deployed to expected address");
            } else {
                // withdraw funds to user as no deployData and no contract deployed
                // The L1 contract shouldn't let this happen!
                // if it does happen, withdraw to sender
                _withdraw(l1ERC20, sender, amount);
            }
        }
        // ignores deployData if token already deployed
        // IArbToken token = IArbToken(expectedAddress);

        if (callHookData.length > 0) {
            handleCallHookData(expectedAddress, amount, sender, dest, callHookData);
        } else {
            IArbToken(expectedAddress).bridgeMint(dest, amount, "");
        }

        emit TokenMinted(
            l1ERC20,
            expectedAddress,
            tokenType,
            sender,
            dest,
            amount,
            callHookData.length > 0
        );
    }

    function deployToken(
        address l1ERC20,
        StandardTokenType tokenType,
        bytes memory deployData
    ) internal returns (address) {
        address beacon = tokenType == StandardTokenType.ERC20 ? templateERC20 : templateERC777;

        deployBeacon = beacon;
        bytes32 salt = keccak256(abi.encodePacked(l1ERC20, beacon));
        address createdContract = address(new ClonableBeaconProxy{ salt: salt }());
        deployBeacon = address(0);

        bool initSuccess = IArbToken(createdContract).bridgeInit(l1ERC20, deployData);
        require(initSuccess, "Bridge init on token failed");

        emit TokenCreated(l1ERC20, createdContract, tokenType);
        return createdContract;
    }

    function customTokenRegistered(address l1Address, address l2Address)
        external
        override
        onlyEthPair
    {
        // TODO: what happens if users already bridged tokens?
        customToken[l1Address] = l2Address;
        emit TokenCreated(l1Address, l2Address, StandardTokenType.Custom);
    }

    function withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) external override onlyFromL2Token(l1ERC20) returns (uint256) {
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
                    EthERC20Bridge.withdrawFromL2.selector,
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

    // A token can be bridged into different L2 implementations (ie 777 and 20)
    // this method allows you to migrate your balance between them.
    function migrate(
        address l1ERC20,
        address target,
        address account,
        uint256 amount,
        bytes memory data
    )
        external
        override
        onlyFromStandardL2Token(l1ERC20)
        onlyToL2Token(l1ERC20, target)
        noCustomToken(l1ERC20)
    {
        require(false, "Method disabled");
        // TODO: ensureTokenExists(l1ERC20, decimals, tokenType);
        IArbToken(target).bridgeMint(account, amount, data);
        emit TokenMigrated(msg.sender, target, account, amount, data);
    }

    function calculateBridgeTokenAddress(address l1ERC20, StandardTokenType tokenType)
        public
        view
        override
        returns (address)
    {
        if (tokenType == StandardTokenType.ERC20) {
            return calculateBridgedERC20Address(l1ERC20);
        } else if (tokenType == StandardTokenType.ERC777) {
            return calculateBridgedERC777Address(l1ERC20);
        } else if (tokenType == StandardTokenType.Custom) {
            address l2Addr = customToken[l1ERC20];
            require(l2Addr != address(0), "No custom address set");
            return l2Addr;
        } else {
            revert("Token type not recognized");
        }
    }

    function calculateBridgedERC777Address(address l1ERC20) public view override returns (address) {
        return
            Create2.computeAddress(
                keccak256(abi.encodePacked(l1ERC20, templateERC777)),
                cloneableProxyHash
            );
    }

    function calculateBridgedERC20Address(address l1ERC20) public view override returns (address) {
        return
            Create2.computeAddress(
                keccak256(abi.encodePacked(l1ERC20, templateERC20)),
                cloneableProxyHash
            );
    }

    function getBeacon() external view override returns (address) {
        return deployBeacon;
    }
}
