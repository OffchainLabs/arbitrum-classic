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

import "../../../../exchange/contracts/src/interfaces/ITransactions.sol";


// solhint-disable var-name-mixedcase
contract LibConstants {

    // The 0x Exchange contract.
    ITransactions internal EXCHANGE;

    /// @param exchange Address of the 0x Exchange contract.
    constructor (address exchange)
        public
    {
        EXCHANGE = ITransactions(exchange);
    }
}
