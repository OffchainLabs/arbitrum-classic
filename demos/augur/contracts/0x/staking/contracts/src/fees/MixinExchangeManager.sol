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


import "../libs/LibStakingRichErrors.sol";
import "../interfaces/IStakingEvents.sol";
import "../immutable/MixinStorage.sol";


contract MixinExchangeManager is
    IStakingEvents,
    MixinStorage
{
    /// @dev Asserts that the call is coming from a valid exchange.
    modifier onlyExchange() {
        if (!validExchanges[msg.sender]) {
            revert();
        }
        _;
    }

    /// @dev Adds a new exchange address
    /// @param addr Address of exchange contract to add
    function addExchangeAddress(address addr)
        external
        onlyAuthorized
    {
        if (validExchanges[addr]) {
            revert();
        }
        validExchanges[addr] = true;
        emit ExchangeAdded(addr);
    }

    /// @dev Removes an existing exchange address
    /// @param addr Address of exchange contract to remove
    function removeExchangeAddress(address addr)
        external
        onlyAuthorized
    {
        if (!validExchanges[addr]) {
            revert();
        }
        validExchanges[addr] = false;
        emit ExchangeRemoved(addr);
    }
}
