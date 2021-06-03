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
import "arb-bridge-eth/contracts/libraries/Whitelist.sol";
import "../../libraries/ITokenGateway.sol";

contract GatewayRouter is WhitelistConsumer, ITokenGateway {
    using Address for address;

    address internal constant ZERO_ADDR = address(0);
    address internal constant BLACKLISTED = address(1);

    mapping(address => address) public tokenToGateway;
    address public owner;
    address public defaultGateway;

    event TransferRouted(
        address indexed token,
        address indexed _userFrom,
        address indexed _userTo,
        address gateway
    );

    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    function initialize(
        address _owner,
        address _defaultGateway,
        address _whitelist
    ) public {
        require(_owner != address(0), "INVALID_OWNER");
        require(owner == address(0), "ALREADY_INIT");
        owner = _owner;
        // if defaultGateway is address(0), only tokens in mapping will not revert
        defaultGateway = _defaultGateway;
        WhitelistConsumer.whitelist = _whitelist;
    }

    function setDefaultGateway(address newDefaultGateway) external onlyOwner {
        defaultGateway = newDefaultGateway;
    }

    function setOwner(address newOwner) external onlyOwner {
        require(newOwner != address(0), "INVALID_OWNER");
        // set newOwner to address(1) to disable owner and keep `initialize` safe
        owner = newOwner;
    }

    function setGateways(address[] memory token, address[] memory gateway) external onlyOwner {
        require(token.length == gateway.length, "WRONG_LENGTH");

        for (uint256 i = 0; i < token.length; i++) {
            tokenToGateway[token[i]] = gateway[i];
        }
    }

    function getGateway(address _token) public view virtual returns (address gateway) {
        gateway = tokenToGateway[_token];
        require(gateway != BLACKLISTED, "BLACKLIST");

        if (gateway == ZERO_ADDR) {
            gateway = defaultGateway;
        }

        require(gateway.isContract(), "NO_GATEWAY_DEPLOYED");

        return gateway;
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable override onlyWhitelisted returns (bytes memory) {
        address gateway = getGateway(_token);
        bytes memory gatewayData = abi.encode(msg.sender, _data);

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
