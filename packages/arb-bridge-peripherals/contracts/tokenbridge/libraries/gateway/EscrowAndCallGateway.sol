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

import "../ITransferAndCall.sol";
import "./IEscrowAndCallGateway.sol";
import "./TokenGateway.sol";

abstract contract EscrowAndCallGateway is IEscrowAndCallGateway, TokenGateway {
    using Address for address;

    event TransferAndCallTriggered(
        bool success,
        address indexed _from,
        address indexed _to,
        uint256 _amount,
        bytes _data
    );

    function inboundEscrowAndCall(
        address _l2Address,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) external virtual override {
        require(msg.sender == address(this), "Mint can only be called by self");
        require(_to.isContract(), "Destination must be a contract");

        inboundEscrowTransfer(_l2Address, _to, _amount);

        ITransferAndCallReceiver(_to).onTokenTransfer(_from, _amount, _data);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IEscrowAndCallGateway).interfaceId;
    }
}
