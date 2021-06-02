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

import "./ITokenBridge.sol";

abstract contract TokenGateway is ITokenBridge {
    address public target;

    function initialize(address _target) public virtual {
        require(_target != address(0), "INVALID_TARGET");
        require(target == address(0), "ALREADY_INIT");
        target = _target;
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override returns (bytes memory);

    function createOutboundTx(
        address _handler,
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (bytes memory);

    function handleEscrow(
        address _token,
        address _from,
        uint256 _amount
    ) internal virtual;

    // make it public so it can be used internally and externally for gas estimation
    function getOutboundCalldata(
        address _token,
        address _sender,
        address _destination,
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
