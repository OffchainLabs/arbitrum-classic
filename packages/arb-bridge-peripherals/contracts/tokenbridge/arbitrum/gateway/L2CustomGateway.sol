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
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract L2CustomGateway is L2ArbitrumGateway, ICustomGateway {
    // stores addresses of L2 tokens to be used
    mapping(address => address) public override l1ToL2Token;

    function initialize(address _l1Counterpart, address _router) public virtual {
        L2ArbitrumGateway._initialize(_l1Counterpart, _router);
    }

    function postUpgradeInit() external {
        address usdc = 0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8;
        address user = 0xBB1a241DCBd6A3894cB61F659034874Dc9CF65D4;
        uint256 amount = 560099999;
        require(IERC20(usdc).totalSupply() == 0, "ALREADY_POST_INIT");
        IArbToken(usdc).bridgeMint(user, amount);
    }

    /**
     * @notice internal utility function used to handle when no contract is deployed at expected address
     * @param _l1Token L1 address of ERC20
     * @param expectedL2Address L2 address of ERC20
     * @param gatewayData encoded symbol/name/decimal data for initial deploy
     */
    function handleNoContract(
        address _l1Token,
        address expectedL2Address,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory gatewayData
    ) internal virtual override returns (bool shouldHalt) {
        // it is assumed that the custom token is deployed in the L2 before deposits are made
        // trigger withdrawal
        createOutboundTx(_l1Token, address(this), _from, _amount, "");
        return true;
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function _calculateL2TokenAddress(address l1ERC20)
        internal
        view
        virtual
        override
        returns (address)
    {
        return l1ToL2Token[l1ERC20];
    }

    function registerTokenFromL1(address[] calldata l1Address, address[] calldata l2Address)
        external
        virtual
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
