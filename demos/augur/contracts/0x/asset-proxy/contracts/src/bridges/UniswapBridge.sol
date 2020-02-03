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

import "../../../../erc20/contracts/src/interfaces/IERC20Token.sol";
import "../../../../erc20/contracts/src/interfaces/IEtherToken.sol";
import "../../../../erc20/contracts/src/LibERC20Token.sol";
import "../../../../exchange-libs/contracts/src/IWallet.sol";
import "../../../../utils/contracts/src/DeploymentConstants.sol";
import "../interfaces/IUniswapExchangeFactory.sol";
import "../interfaces/IUniswapExchange.sol";
import "../interfaces/IERC20Bridge.sol";


// solhint-disable space-after-comma
// solhint-disable not-rely-on-time
contract UniswapBridge is
    IERC20Bridge,
    IWallet,
    DeploymentConstants
{
    // Struct to hold `bridgeTransferFrom()` local variables in memory and to avoid
    // stack overflows.
    struct WithdrawToState {
        IUniswapExchange exchange;
        uint256 fromTokenBalance;
        IEtherToken weth;
    }

    // solhint-disable no-empty-blocks
    /// @dev Payable fallback to receive ETH from uniswap.
    function ()
        external
        payable
    {}

    /// @dev Callback for `IERC20Bridge`. Tries to buy `amount` of
    ///      `toTokenAddress` tokens by selling the entirety of the `fromTokenAddress`
    ///      token encoded in the bridge data.
    /// @param toTokenAddress The token to buy and transfer to `to`.
    /// @param to The recipient of the bought tokens.
    /// @param amount Minimum amount of `toTokenAddress` tokens to buy.
    /// @param bridgeData The abi-encoded "from" token address.
    /// @return success The magic bytes if successful.
    function bridgeTransferFrom(
        address toTokenAddress,
        address /* from */,
        address to,
        uint256 amount,
        bytes calldata bridgeData
    )
        external
        returns (bytes4 success)
    {
        // State memory object to avoid stack overflows.
        WithdrawToState memory state;
        // Decode the bridge data to get the `fromTokenAddress`.
        (address fromTokenAddress) = abi.decode(bridgeData, (address));

        // Just transfer the tokens if they're the same.
        if (fromTokenAddress == toTokenAddress) {
            LibERC20Token.transfer(fromTokenAddress, to, amount);
            return BRIDGE_SUCCESS;
        }

        // Get the exchange for the token pair.
        state.exchange = _getUniswapExchangeForTokenPair(
            fromTokenAddress,
            toTokenAddress
        );
        // Get our balance of `fromTokenAddress` token.
        state.fromTokenBalance = IERC20Token(fromTokenAddress).balanceOf(address(this));
        // Get the weth contract.
        state.weth = IEtherToken(_getWethAddress());

        // Convert from WETH to a token.
        if (fromTokenAddress == address(state.weth)) {
            // Unwrap the WETH.
            state.weth.withdraw(state.fromTokenBalance);
            // Buy as much of `toTokenAddress` token with ETH as possible and
            // transfer it to `to`.
            state.exchange.ethToTokenTransferInput.value(state.fromTokenBalance)(
                // Minimum buy amount.
                amount,
                // Expires after this block.
                block.timestamp,
                // Recipient is `to`.
                to
            );

        // Convert from a token to WETH.
        } else if (toTokenAddress == address(state.weth)) {
            // Grant the exchange an allowance.
            _grantExchangeAllowance(state.exchange, fromTokenAddress);
            // Buy as much ETH with `fromTokenAddress` token as possible.
            uint256 ethBought = state.exchange.tokenToEthSwapInput(
                // Sell all tokens we hold.
                state.fromTokenBalance,
                // Minimum buy amount.
                amount,
                // Expires after this block.
                block.timestamp
            );
            // Wrap the ETH.
            state.weth.deposit.value(ethBought)();
            // Transfer the WETH to `to`.
            IEtherToken(toTokenAddress).transfer(to, ethBought);

        // Convert from one token to another.
        } else {
            // Grant the exchange an allowance.
            _grantExchangeAllowance(state.exchange, fromTokenAddress);
            // Buy as much `toTokenAddress` token with `fromTokenAddress` token
            // and transfer it to `to`.
            state.exchange.tokenToTokenTransferInput(
                // Sell all tokens we hold.
                state.fromTokenBalance,
                // Minimum buy amount.
                amount,
                // Must buy at least 1 intermediate ETH.
                1,
                // Expires after this block.
                block.timestamp,
                // Recipient is `to`.
                to,
                // Convert to `toTokenAddress`.
                toTokenAddress
            );
        }
        return BRIDGE_SUCCESS;
    }

    /// @dev `SignatureType.Wallet` callback, so that this bridge can be the maker
    ///      and sign for itself in orders. Always succeeds.
    /// @return magicValue Success bytes, always.
    function isValidSignature(
        bytes32,
        bytes calldata
    )
        external
        view
        returns (bytes4 magicValue)
    {
        return LEGACY_WALLET_MAGIC_VALUE;
    }

    /// @dev Grants an unlimited allowance to the exchange for its token
    ///      on behalf of this contract.
    /// @param exchange The Uniswap token exchange.
    /// @param tokenAddress The token address for the exchange.
    function _grantExchangeAllowance(IUniswapExchange exchange, address tokenAddress)
        private
    {
        LibERC20Token.approve(tokenAddress, address(exchange), uint256(-1));
    }

    /// @dev Retrieves the uniswap exchange for a given token pair.
    ///      In the case of a WETH-token exchange, this will be the non-WETH token.
    ///      In th ecase of a token-token exchange, this will be the first token.
    /// @param fromTokenAddress The address of the token we are converting from.
    /// @param toTokenAddress The address of the token we are converting to.
    /// @return exchange The uniswap exchange.
    function _getUniswapExchangeForTokenPair(
        address fromTokenAddress,
        address toTokenAddress
    )
        private
        view
        returns (IUniswapExchange exchange)
    {
        address exchangeTokenAddress = fromTokenAddress;
        // Whichever isn't WETH is the exchange token.
        if (fromTokenAddress == _getWethAddress()) {
            exchangeTokenAddress = toTokenAddress;
        }
        exchange = IUniswapExchange(
            IUniswapExchangeFactory(_getUniswapExchangeFactoryAddress())
            .getExchange(exchangeTokenAddress)
        );
        require(address(exchange) != address(0), "NO_UNISWAP_EXCHANGE_FOR_TOKEN");
        return exchange;
    }
}
