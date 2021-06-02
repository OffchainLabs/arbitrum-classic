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

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "./ITokenBridge.sol";

abstract contract RouterConsumer is ITokenBridge {
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
    ) external payable virtual override returns (bytes memory res) {
        (
            address _handler, // inbox
            address _from,
            uint256 _maxSubmissionCost,
            bytes memory extraData
        ) = parseArbitrumData(_data);

        triggerEscrow(_token, _from, _amount);

        bytes memory outboundCalldata =
            getOutboundCalldata(_token, target, _from, _to, _amount, extraData);

        res = createOutboundTx(
            _handler,
            target,
            _from,
            _maxSubmissionCost,
            _maxGas,
            _gasPriceBid,
            outboundCalldata
        );

        emit OutboundTransferInitiated(_token, _from, _to, _amount, _data);

        return res;
    }

    function parseArbitrumData(bytes memory _data)
        internal
        pure
        virtual
        returns (
            address inbox,
            address from,
            uint256 maxSubmissionCost,
            bytes memory _extraData
        )
    {
        // router encoded
        (inbox, from, _extraData) = abi.decode(_data, (address, address, bytes));
        // user encoded
        (maxSubmissionCost, _extraData) = abi.decode(_extraData, (uint256, bytes));
    }

    function createOutboundTx(
        address _handler,
        address _target,
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (bytes memory);

    function triggerEscrow(
        address _token,
        address _from,
        uint256 _amount
    ) internal virtual;

    // make it public so it can be used internally and externally for gas estimation
    function getOutboundCalldata(
        address _token,
        address _target,
        address _sender,
        address _destination,
        uint256 _amount,
        bytes memory _data
    ) public view virtual returns (bytes memory);

    function finalizeInboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override returns (bytes memory);
}
