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
import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibZeroExTransaction.sol";
import "../../../utils/contracts/src/LibEIP712.sol";
import "../../../utils/contracts/src/LibBytes.sol";
import "./OrderValidationUtils.sol";
import "./OrderTransferSimulationUtils.sol";
import "./EthBalanceChecker.sol";


// solhint-disable no-empty-blocks
contract DevUtils is
    OrderValidationUtils,
    LibEIP712ExchangeDomain,
    EthBalanceChecker
{
    constructor (
        address _exchange,
        address _chaiBridge
    )
        public
        OrderValidationUtils(
            _exchange,
            _chaiBridge
        )
        OrderTransferSimulationUtils(_exchange)
        LibEIP712ExchangeDomain(uint256(0), address(0)) // null args because because we only use constants
    {}

    function getOrderHash(
        LibOrder.Order memory order,
        uint256 chainId,
        address exchange
    )
        public
        pure
        returns (bytes32 orderHash)
    {
        return LibOrder.getTypedDataHash(
            order,
            LibEIP712.hashEIP712Domain(_EIP712_EXCHANGE_DOMAIN_NAME, _EIP712_EXCHANGE_DOMAIN_VERSION, chainId, exchange)
        );
    }

    function getTransactionHash(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        uint256 chainId,
        address exchange
    )
        public
        pure
        returns (bytes32 transactionHash)
    {
        return LibZeroExTransaction.getTypedDataHash(
            transaction,
            LibEIP712.hashEIP712Domain(_EIP712_EXCHANGE_DOMAIN_NAME, _EIP712_EXCHANGE_DOMAIN_VERSION, chainId, exchange)
        );
    }
}
