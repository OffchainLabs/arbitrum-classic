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

import "@openzeppelin/contracts/utils/Address.sol";
import "./TokenGateway.sol";

/**
 * @title Handles routing tokens to their appropriate gateways (Router itself also conforms to the Gateway interface).
 * @notice Router also serves as an L1-L2 token address oracle.
 */
abstract contract GatewayRouter is TokenGateway {
    using Address for address;

    mapping(address => address) public l1TokenToGateway;

    event TransferRouted(
        address indexed token,
        address indexed _userFrom,
        address indexed _userTo,
        address gateway
    );

    event GatewaySet(address indexed l1Token, address indexed gateway);

    function _initialize(address _counterpartGateway) internal virtual {
        TokenGateway._initialize(_counterpartGateway, address(0));
    }

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override returns (bytes memory) {
        revert("ONLY_OUTBOUND_ROUTER");
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override returns (bytes memory) {
        preTransferHook();
        address gateway = getGateway(_token);
        bytes memory gatewayData = getOutboundCalldata(_token, msg.sender, _to, _amount, _data);

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
        return abi.encode(_from, _data);
    }

    function isRouter() internal view virtual override returns (bool) {
        // nothing routes to gateway router
        return false;
    }

    function getGateway(address _token) public view virtual returns (address gateway);

    function preTransferHook() internal virtual;
}
