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

import "./ITokenGateway.sol";

abstract contract TokenGateway is ITokenGateway {
    address public counterpartGateway;

    modifier onlyCounterpartGateway {
        require(isCounterpartGateway(), "ONLY_COUNTERPART_GATEWAY");
        _;
    }

    function initialize(address _counterpartGateway) public virtual {
        require(_counterpartGateway != address(0), "INVALID_COUNTERPART");
        require(counterpartGateway == address(0), "ALREADY_INIT");
        counterpartGateway = _counterpartGateway;
    }

    function isCounterpartGateway() internal view virtual returns (bool);

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override returns (bytes memory);

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual returns (bytes memory);

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override returns (bytes memory);
}
