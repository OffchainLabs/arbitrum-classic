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

import "../../../utils/contracts/src/LibSafeMath.sol";
import "./UnlimitedAllowanceERC20Token.sol";


contract MintableERC20Token is
    UnlimitedAllowanceERC20Token
{
    using LibSafeMath for uint256;

    /// @dev Mints new tokens
    /// @param _to Address of the beneficiary that will own the minted token
    /// @param _value Amount of tokens to mint
    function _mint(address _to, uint256 _value)
        internal
    {
        balances[_to] = _value.safeAdd(balances[_to]);
        _totalSupply = _totalSupply.safeAdd(_value);

        emit Transfer(
            address(0),
            _to,
            _value
        );
    }

    /// @dev Mints new tokens
    /// @param _owner Owner of tokens that will be burned
    /// @param _value Amount of tokens to burn
    function _burn(address _owner, uint256 _value)
        internal
    {
        balances[_owner] = balances[_owner].safeSub(_value);
        _totalSupply = _totalSupply.safeSub(_value);

        emit Transfer(
            _owner,
            address(0),
            _value
        );
    }
}
