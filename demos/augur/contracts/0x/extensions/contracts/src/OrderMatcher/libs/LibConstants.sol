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

import "../../../../../exchange/contracts/src/interfaces/IExchange.sol";


contract LibConstants {

    // bytes4(keccak256("transfer(address,uint256)"))
    bytes4 constant internal ERC20_TRANSFER_SELECTOR = 0xa9059cbb;
    // bytes4(keccak256("ERC20Token(address)"))
    bytes4 constant internal ERC20_DATA_ID = 0xf47261b0;
    // bytes4(keccak256("ERC721Token(address,uint256)"))
    bytes4 constant internal ERC721_DATA_ID = 0x02571792;
 
     // solhint-disable var-name-mixedcase
    IExchange internal EXCHANGE;
    address internal ERC20_PROXY_ADDRESS;
    address internal ERC721_PROXY_ADDRESS;
    // solhint-enable var-name-mixedcase

    constructor (address _exchange)
        public
    {
        EXCHANGE = IExchange(_exchange);

        ERC20_PROXY_ADDRESS = EXCHANGE.getAssetProxy(ERC20_DATA_ID);
        require(
            ERC20_PROXY_ADDRESS != address(0),
            "UNREGISTERED_ASSET_PROXY"
        );

        ERC721_PROXY_ADDRESS = EXCHANGE.getAssetProxy(ERC721_DATA_ID);
        require(
            ERC721_PROXY_ADDRESS != address(0),
            "UNREGISTERED_ASSET_PROXY"
        );
    }
}
