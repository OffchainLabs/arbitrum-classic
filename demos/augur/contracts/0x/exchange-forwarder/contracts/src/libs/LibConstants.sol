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

import "../../../../exchange/contracts/src/interfaces/IExchange.sol";
import "../../../../erc20/contracts/src/interfaces/IEtherToken.sol";
import "../../../../exchange-forwarder/contracts/src/interfaces/IExchangeV2.sol";


contract LibConstants {

    uint256 constant internal MAX_UINT = uint256(-1);

    // The v2 order id is the first 4 bytes of the ExchangeV2 order schema hash.
    // bytes4(keccak256(abi.encodePacked(
    //     "Order(",
    //     "address makerAddress,",
    //     "address takerAddress,",
    //     "address feeRecipientAddress,",
    //     "address senderAddress,",
    //     "uint256 makerAssetAmount,",
    //     "uint256 takerAssetAmount,",
    //     "uint256 makerFee,",
    //     "uint256 takerFee,",
    //     "uint256 expirationTimeSeconds,",
    //     "uint256 salt,",
    //     "bytes makerAssetData,",
    //     "bytes takerAssetData",
    //     ")"
    // )));
    bytes4 constant public EXCHANGE_V2_ORDER_ID = 0x770501f8;

     // solhint-disable var-name-mixedcase
    IExchange internal EXCHANGE;
    IExchangeV2 internal EXCHANGE_V2;
    IEtherToken internal ETHER_TOKEN;
    // solhint-enable var-name-mixedcase

    constructor (
        address _exchange,
        address _exchangeV2,
        address _weth
    )
        public
    {
        EXCHANGE = IExchange(_exchange);
        EXCHANGE_V2 = IExchangeV2(_exchangeV2);
        ETHER_TOKEN = IEtherToken(_weth);
    }
}
