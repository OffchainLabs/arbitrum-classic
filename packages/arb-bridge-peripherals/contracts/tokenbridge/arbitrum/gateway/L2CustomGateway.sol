// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.6.11;

import "./L2ArbitrumGateway.sol";
import "../../libraries/gateway/ICustomGateway.sol";

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract L2CustomGateway is L2ArbitrumGateway, ICustomGateway {
    // stores addresses of L2 tokens to be used
    mapping(address => address) public override l1ToL2Token;

    function initialize(address _l1Counterpart, address _router) public {
        L2ArbitrumGateway._initialize(_l1Counterpart, _router);
    }

    /**
     * @notice internal utility function used to handle when no contract is deployed at expected address
     */
    function handleNoContract(
        address _l1Token,
        address, /* expectedL2Address */
        address _from,
        address, /* _to */
        uint256 _amount,
        bytes memory /* gatewayData */
    ) internal override returns (bool shouldHalt) {
        // it is assumed that the custom token is deployed in the L2 before deposits are made
        // trigger withdrawal
        // we don't need the return value from triggerWithdrawal since this is forcing a withdrawal back to the L1
        // instead of composing with a L2 dapp
        triggerWithdrawal(_l1Token, address(this), _from, _amount, "");
        return true;
    }

    function outboundEscrowTransfer(
        address _l2Token,
        address _from,
        uint256 _amount
    ) internal override returns (uint256 amountBurnt) {
        uint256 prevBalance = IERC20(_l2Token).balanceOf(_from);

        // in the custom gateway, we do the same behaviour as the superclass, but actually check
        // for the balances of tokens to ensure that inflationary / deflationary changes in the amount
        // are taken into account
        // we ignore the return value since we actually query the token before and after to calculate
        // the amount of tokens that were burnt
        super.outboundEscrowTransfer(_l2Token, _from, _amount);

        uint256 postBalance = IERC20(_l2Token).balanceOf(_from);
        return SafeMath.sub(prevBalance, postBalance);
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev the L1 and L2 address oracles may not always be in sync.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20) public view override returns (address) {
        return l1ToL2Token[l1ERC20];
    }

    function registerTokenFromL1(address[] calldata l1Address, address[] calldata l2Address)
        external
        onlyCounterpartGateway
    {
        // we assume both arrays are the same length, safe since its encoded by the L1
        for (uint256 i = 0; i < l1Address.length; i++) {
            // here we don't check if l2Address is a contract and instead deal with that behaviour
            // in `handleNoContract` this way we keep the l1 and l2 address oracles in sync
            l1ToL2Token[l1Address[i]] = l2Address[i];
            emit TokenSet(l1Address[i], l2Address[i]);
        }
    }
}
