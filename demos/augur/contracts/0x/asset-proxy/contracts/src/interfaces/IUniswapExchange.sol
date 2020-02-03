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


interface IUniswapExchange {

    /// @dev Buys at least `minTokensBought` tokens with ETH and transfer them
    ///      to `recipient`.
    /// @param minTokensBought The minimum number of tokens to buy.
    /// @param deadline Time when this order expires.
    /// @param recipient Who to transfer the tokens to.
    /// @return tokensBought Amount of tokens bought.
    function ethToTokenTransferInput(
        uint256 minTokensBought,
        uint256 deadline,
        address recipient
    )
        external
        payable
        returns (uint256 tokensBought);

    /// @dev Buys at least `minEthBought` ETH with tokens.
    /// @param tokensSold Amount of tokens to sell.
    /// @param minEthBought The minimum amount of ETH to buy.
    /// @param deadline Time when this order expires.
    /// @return ethBought Amount of tokens bought.
    function tokenToEthSwapInput(
        uint256 tokensSold,
        uint256 minEthBought,
        uint256 deadline
    )
        external
        returns (uint256 ethBought);

    /// @dev Buys at least `minTokensBought` tokens with the exchange token
    ///      and transfer them to `recipient`.
    /// @param minTokensBought The minimum number of tokens to buy.
    /// @param minEthBought The minimum amount of intermediate ETH to buy.
    /// @param deadline Time when this order expires.
    /// @param recipient Who to transfer the tokens to.
    /// @param toTokenAddress The token being bought.
    /// @return tokensBought Amount of tokens bought.
    function tokenToTokenTransferInput(
        uint256 tokensSold,
        uint256 minTokensBought,
        uint256 minEthBought,
        uint256 deadline,
        address recipient,
        address toTokenAddress
    )
        external
        returns (uint256 tokensBought);
}
