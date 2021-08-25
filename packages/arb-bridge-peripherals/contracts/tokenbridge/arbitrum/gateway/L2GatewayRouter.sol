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

import "../../libraries/gateway/GatewayRouter.sol";
import "../../ethereum/gateway/L1GatewayRouter.sol";
import "../L2ArbitrumMessenger.sol";

/**
 * @title Handles withdrawals from Ethereum into Arbitrum. Tokens are routered to their appropriate L2 gateway (Router itself also conforms to the Gateway interface).
 * @notice Router also serves as an L2-L1 token address oracle.
 */
contract L2GatewayRouter is GatewayRouter, L2ArbitrumMessenger {
    modifier onlyCounterpartGateway() override {
        require(
            msg.sender == counterpartGateway ||
                L2ArbitrumMessenger.getL1Address(msg.sender) == counterpartGateway,
            "ONLY_COUNTERPART_GATEWAY"
        );
        _;
    }

    function initialize(address _counterpartGateway, address _defaultGateway) public {
        GatewayRouter._initialize(_counterpartGateway, address(0), _defaultGateway);
    }

    function setGateway(address[] memory _l1Token, address[] memory _gateway)
        external
        onlyCounterpartGateway
    {
        // counterpart gateway (L1 router) should never allow wrong lengths
        assert(_l1Token.length == _gateway.length);

        for (uint256 i = 0; i < _l1Token.length; i++) {
            l1TokenToGateway[_l1Token[i]] = _gateway[i];
            emit GatewaySet(_l1Token[i], _gateway[i]);
        }
    }

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) public payable returns (bytes memory) {
        return outboundTransfer(_l1Token, _to, _amount, 0, 0, _data);
    }

    function setDefaultGateway(address newL2DefaultGateway) external onlyCounterpartGateway {
        defaultGateway = newL2DefaultGateway;
        emit DefaultGatewayUpdated(newL2DefaultGateway);
    }
}
