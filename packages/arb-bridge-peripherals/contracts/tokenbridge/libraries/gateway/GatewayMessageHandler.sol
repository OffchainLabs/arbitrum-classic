// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

/// @notice this library manages encoding and decoding of gateway communication
library GatewayMessageHandler {
    // TODO: set actual fork block numbers. for now this should always default to v1
    uint256 private constant L1_MAINNET_FORK_BLOCK = 999999999999999999999999999;
    uint256 private constant L2_MAINNET_FORK_BLOCK = 999999999999999999999999999;
    uint256 private constant L1_RINKEBY_FORK_BLOCK = 999999999999999999999999999;
    uint256 private constant L2_RINKEBY_FORK_BLOCK = 999999999999999999999999999;

    uint256 private constant HARDHAT_FORK_BLOCK = 999999999999999999999999999;

    function getForkBlock() private view returns (uint256 forkBlock) {
        assembly {
            switch chainid()
            case 1 {
                forkBlock := L1_MAINNET_FORK_BLOCK
            }
            case 42161 {
                forkBlock := L2_MAINNET_FORK_BLOCK
            }
            case 4 {
                forkBlock := L1_RINKEBY_FORK_BLOCK
            }
            case 421611 {
                forkBlock := L2_RINKEBY_FORK_BLOCK
            }
            // TODO: should we set hardhat network to use mainnet/rinkeby chainid instead of special case here?
            case 1337 {
                forkBlock := HARDHAT_FORK_BLOCK
            }
            default {
                // forkBlock is set to 0 by default
            }
        }
        if (forkBlock == 0) revert("INVALID_CHAIN_ID");
    }

    function getGatewayMessageVersion(bytes calldata _data) internal view returns (uint8) {
        // It is easy to version bump from L2 to L1, but L1 to L2 you need to ensure all
        // retryables are consumed before setting a fork block number

        // TODO: handle L1 fork block numbers too. this includes querying the outbox for the L2 block num
        if (block.number < getForkBlock()) {
            // we start at v1 to avoid errors from uninitialized variables
            return 1;
        } else {
            // TODO: should we concat instead of rlp encode?
            (uint8 version, ) = abi.decode(_data, (uint8, bytes));
            // TODO: should we assert version isn't 1? can do this on the consumer
            return version;
        }
    }

    // TODO: should we split L1 and L2 sides? maybe not so its easier to see encode/decode are inverse of each other

    // these are for communication from L1 to L2 gateway

    /// @notice message v1 type will be deprecated
    /// @dev this assumes the message version was previously validated
    function encodeToL2GatewayMsgV1(bytes memory gatewayData, bytes memory callHookData)
        internal
        pure
        returns (bytes memory res)
    {
        res = abi.encode(gatewayData, callHookData);
    }

    /// @notice message v1 type will be deprecated
    /// @dev this assumes the message version was previously validated
    function parseFromL1GatewayMsgV1(bytes calldata _data)
        internal
        pure
        returns (bytes memory gatewayData, bytes memory callHookData)
    {
        // abi decode may revert, but the encoding is done by L1 gateway, so we trust it
        (gatewayData, callHookData) = abi.decode(_data, (bytes, bytes));
    }

    // these are for communication from L2 to L1 gateway

    /// @notice message v1 type will be deprecated
    /// @dev this assumes the message version was previously validated
    function encodeFromL2GatewayMsgV1(uint256 exitNum, bytes memory callHookData)
        internal
        pure
        returns (bytes memory res)
    {
        res = abi.encode(exitNum, callHookData);
    }

    /// @notice message v1 type will be deprecated
    /// @dev this assumes the message version was previously validated
    function parseToL1GatewayMsgV1(bytes calldata _data)
        internal
        pure
        returns (uint256 exitNum, bytes memory callHookData)
    {
        // abi decode may revert, but the encoding is done by L1 gateway, so we trust it
        (exitNum, callHookData) = abi.decode(_data, (uint256, bytes));
    }

    function encodeFromRouterToGateway(address _from, bytes calldata _data)
        internal
        view
        returns (bytes memory res)
    {
        // abi decode may revert, but the encoding is done by L1 gateway, so we trust it
        return abi.encode(_from, _data);
    }

    function parseFromRouterToGateway(bytes calldata _data)
        internal
        view
        returns (address, bytes memory res)
    {
        // abi decode may revert, but the encoding is done by L1 gateway, so we trust it
        return abi.decode(_data, (address, bytes));
    }
}
