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

import "../../../../erc20/contracts/src/interfaces/IEtherToken.sol";
import "../interfaces/IZrxVault.sol";


// solhint-disable separate-by-one-line-in-contract
contract MixinDeploymentConstants {

    // @TODO SET THESE VALUES FOR DEPLOYMENT

    // Mainnet WETH9 Address
    address constant private WETH_ADDRESS = address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);

    // Kovan WETH9 Address
    // address constant private WETH_ADDRESS = address(0xd0A1E359811322d97991E03f863a0C30C2cF029C);

    // Ropsten & Rinkeby WETH9 Address
    // address constant private WETH_ADDRESS = address(0xc778417E063141139Fce010982780140Aa0cD5Ab);

    // @TODO SET THESE VALUES FOR DEPLOYMENT

    // Mainnet ZrxVault address
    address constant private ZRX_VAULT_ADDRESS = address(0xBa7f8b5fB1b19c1211c5d49550fcD149177A5Eaf);

    // Kovan ZrxVault address
    // address constant private ZRX_VAULT_ADDRESS = address(0xf36eabdFE986B35b62c8FD5a98A7f2aEBB79B291);

    // Ropsten ZrxVault address
    // address constant private ZRX_VAULT_ADDRESS = address(0xffD161026865Ad8B4aB28a76840474935eEc4DfA);

    // Rinkeby ZrxVault address
    // address constant private ZRX_VAULT_ADDRESS = address(0xA5Bf6aC73bC40790FC6Ffc9DBbbCE76c9176e224);

    /// @dev An overridable way to access the deployed WETH contract.
    ///      Must be view to allow overrides to access state.
    /// @return wethContract The WETH contract instance.
    function getWethContract()
        public
        view
        returns (IEtherToken wethContract)
    {
        wethContract = IEtherToken(WETH_ADDRESS);
        return wethContract;
    }

    /// @dev An overridable way to access the deployed zrxVault.
    ///      Must be view to allow overrides to access state.
    /// @return zrxVault The zrxVault contract.
    function getZrxVault()
        public
        view
        returns (IZrxVault zrxVault)
    {
        zrxVault = IZrxVault(ZRX_VAULT_ADDRESS);
        return zrxVault;
    }
}
