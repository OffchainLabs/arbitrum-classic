/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;

import "../../../exchange-libs/contracts/src/LibEIP712ExchangeDomain.sol";
import "../libs/LibConstants.sol";
import "../libs/LibEIP712CoordinatorDomain.sol";
import "./MixinSignatureValidator.sol";
import "./MixinCoordinatorApprovalVerifier.sol";
import "./MixinCoordinatorCore.sol";


// solhint-disable no-empty-blocks
contract Coordinator is
    LibConstants,
    MixinSignatureValidator,
    MixinCoordinatorApprovalVerifier,
    MixinCoordinatorCore
{
    /// @param exchange Address of the 0x Exchange contract.
    /// @param chainId Chain ID of the network this contract is deployed on.
    constructor (address exchange, uint256 chainId)
        public
        LibConstants(exchange)
        LibEIP712CoordinatorDomain(chainId, address(0))
        LibEIP712ExchangeDomain(chainId, exchange)
    {}
}
