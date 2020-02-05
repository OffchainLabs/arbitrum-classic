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

import "../../../../erc20/contracts/src/interfaces/IERC20Token.sol";


contract PotLike {
    function chi() external returns (uint256);
    function rho() external returns (uint256);
    function drip() external returns (uint256);
    function join(uint256) external;
    function exit(uint256) external;
}


// The actual Chai contract can be found here: https://github.com/dapphub/chai
contract IChai is
    IERC20Token
{
    /// @dev Withdraws Dai owned by `src`
    /// @param src Address that owns Dai.
    /// @param wad Amount of Dai to withdraw.
    function draw(
        address src,
        uint256 wad
    )
        external;

    /// @dev Queries Dai balance of Chai holder.
    /// @param usr Address of Chai holder.
    /// @return Dai balance.
    function dai(address usr)
        external
        returns (uint256);

    /// @dev Queries the Pot contract used by the Chai contract.
    function pot()
        external
        returns (PotLike);

    /// @dev Deposits Dai in exchange for Chai
    /// @param dst Address to receive Chai.
    /// @param wad Amount of Dai to deposit.
    function join(
        address dst,
        uint256 wad
    )
        external;
}
