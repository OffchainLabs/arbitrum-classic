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

import "arb-bridge-eth/contracts/libraries/ProxyUtil.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "./TokenGateway.sol";
import "./GatewayMessageHandler.sol";

/**
 * @title Common interface for L1 and L2 Gateway Routers
 */
abstract contract GatewayRouter is TokenGateway {
    using Address for address;

    address internal constant ZERO_ADDR = address(0);
    address internal constant DISABLED = address(1);

    mapping(address => address) public l1TokenToGateway;
    address public defaultGateway;

    event TransferRouted(
        address indexed token,
        address indexed _userFrom,
        address indexed _userTo,
        address gateway
    );

    event GatewaySet(address indexed l1Token, address indexed gateway);
    event DefaultGatewayUpdated(address newDefaultGateway);

    function postUpgradeInit() external {
        // it is assumed the L2 Arbitrum Gateway contract is behind a Proxy controlled by a proxy admin
        // this function can only be called by the proxy admin contract
        address proxyAdmin = ProxyUtil.getProxyAdmin();
        require(msg.sender == proxyAdmin, "NOT_FROM_ADMIN");
        // this has no other logic since the current upgrade doesn't require this logic
    }

    function _initialize(
        address _counterpartGateway,
        address _router,
        address _defaultGateway
    ) internal {
        // if you are a router, you can't have a router
        require(_router == address(0), "BAD_ROUTER");
        TokenGateway._initialize(_counterpartGateway, _router);
        // default gateway can have 0 address
        defaultGateway = _defaultGateway;
    }

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external payable virtual override returns (bytes memory) {
        revert("ONLY_OUTBOUND_ROUTER");
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable virtual override returns (bytes memory) {
        address gateway = getGateway(_token);
        bytes memory gatewayData = GatewayMessageHandler.encodeFromRouterToGateway(
            msg.sender,
            _data
        );

        emit TransferRouted(_token, msg.sender, _to, gateway);
        return
            ITokenGateway(gateway).outboundTransfer{ value: msg.value }(
                _token,
                _to,
                _amount,
                _maxGas,
                _gasPriceBid,
                gatewayData
            );
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory) {
        address gateway = getGateway(_token);
        return TokenGateway(gateway).getOutboundCalldata(_token, _from, _to, _amount, _data);
    }

    function getGateway(address _token) public view virtual returns (address gateway) {
        gateway = l1TokenToGateway[_token];

        if (gateway == ZERO_ADDR) {
            // if no gateway value set, use default gateway
            gateway = defaultGateway;
        }

        if (gateway == DISABLED || !gateway.isContract()) {
            // not a valid gateway
            return ZERO_ADDR;
        }

        return gateway;
    }

    function calculateL2TokenAddress(address l1ERC20)
        public
        view
        virtual
        override
        returns (address)
    {
        address gateway = getGateway(l1ERC20);
        if (gateway == ZERO_ADDR) {
            return ZERO_ADDR;
        }
        return TokenGateway(gateway).calculateL2TokenAddress(l1ERC20);
    }
}
