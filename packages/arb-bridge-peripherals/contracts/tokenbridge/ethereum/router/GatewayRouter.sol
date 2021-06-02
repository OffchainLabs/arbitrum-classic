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

import "../../libraries/ITokenGateway.sol";

contract GatewayRouter is ITokenGateway {
    using Address for address;

    address internal constant ZERO_ADDR = address(0);
    address internal constant BLACKLISTED = address(1);

    mapping(address => address) public tokenToConsumer;
    address public owner;
    address public inbox;
    // TODO: set defaultConsumer
    address public defaultConsumer;

    function initialize(address _owner, address _inbox) public {
        require(_inbox != address(0), "INVALID_INBOX");
        require(inbox == address(0), "ALREADY_INIT");
        owner = _owner;
        inbox = _inbox;
    }

    function getGateway(address _token) public view virtual returns (address gateway) {
        gateway = tokenToConsumer[_token];
        require(gateway != BLACKLISTED, "BLACKLIST");

        if (gateway == ZERO_ADDR) {
            gateway = defaultConsumer;
        }

        require(gateway.isContract(), "NO_CONSUMER_DEPLOYED");

        return gateway;
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable override returns (bytes memory) {
        // TODO: check whitelist
        address gateway = getGateway(_token);
        bytes memory gatewayData = abi.encode(inbox, msg.sender, _data);

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

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external override returns (bytes memory) {
        revert("ONLY_OUTBOUND_ROUTER");
    }
}
