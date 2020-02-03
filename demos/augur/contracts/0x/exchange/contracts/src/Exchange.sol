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
import "./MixinMatchOrders.sol";
import "./MixinWrapperFunctions.sol";
import "./MixinTransferSimulator.sol";


// solhint-disable no-empty-blocks
// MixinAssetProxyDispatcher, MixinExchangeCore, MixinSignatureValidator,
// and MixinTransactions are all inherited via the other Mixins that are
// used.
/// @dev The 0x Exchange contract.
contract Exchange is
    LibEIP712ExchangeDomain,
    MixinMatchOrders,
    MixinWrapperFunctions,
    MixinTransferSimulator
{
    /// @dev Mixins are instantiated in the order they are inherited
    /// @param chainId Chain ID of the network this contract is deployed on.
    constructor (uint256 chainId)
        public
        LibEIP712ExchangeDomain(chainId, address(0))
    {}
    
    // For testing purposes
    function isValidSignature(LibOrder.Order memory order, bytes32 orderHash, bytes memory signature) public view returns (bool) {
        return _isValidOrderWithHashSignature(order, orderHash, signature);
    }
}
