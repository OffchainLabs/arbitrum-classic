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


contract EthBalanceChecker {

    /// @dev Batch fetches ETH balances
    /// @param addresses Array of addresses.
    /// @return Array of ETH balances.
    function getEthBalances(address[] memory addresses)
        public
        view
        returns (uint256[] memory)
    {
        uint256[] memory balances = new uint256[](addresses.length);
        for (uint256 i = 0; i != addresses.length; i++) {
            balances[i] = addresses[i].balance;
        }
        return balances;
    }

}
