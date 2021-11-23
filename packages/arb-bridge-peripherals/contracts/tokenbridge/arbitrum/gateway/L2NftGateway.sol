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

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "arb-bridge-eth/contracts/libraries/AddressAliasHelper.sol";

import "../../libraries/Escrow721.sol";
import "../../libraries/ClonableBeaconProxy.sol";
import "../../ethereum/gateway/L1NftGateway.sol";
import "../L2ArbitrumMessenger.sol";
import "../StandardArbERC721.sol";

/**
 * @title Common interface for L1 and L2 Gateway Routers
 */
contract L2NftGateway is L2ArbitrumMessenger, IERC721Receiver {
    address public counterpartGateway;
    address public router;
    address public beaconProxyFactory;
    // TODO: be careful, these break through upgrades
    bytes constant bytecode = type(Escrow721).creationCode;
    bytes32 constant bytecodeHash = keccak256(type(Escrow721).creationCode);

    function initialize(
        address _counterpartGateway,
        address _router,
        address _beaconProxyFactory
    ) public {
        require(counterpartGateway == address(0), "ALREADY_INIT");
        require(_counterpartGateway != address(0), "BAD_COUNTERPART");
        require(_router != address(0), "BAD_ROUTER");
        counterpartGateway = _counterpartGateway;
        router = _router;
        beaconProxyFactory = _beaconProxyFactory;
    }

    function withdraw(
        address l1Token,
        uint256 tokenId,
        address from,
        address to,
        bool shouldCreate2Escrow,
        bytes calldata data
    ) external returns (uint256) {
        address expectedAddress = calculateL2TokenAddress(l1Token);

        require(StandardArbERC721(expectedAddress).l1Address() == l1Token, "INVALID_TOKEN");

        require(!shouldCreate2Escrow, "EXTERNAL_ESCROW_DISABLED");
        if (shouldCreate2Escrow) {
            // TODO: can we just check if the token is escrowed there, and allow this to be deployed during an eventual withdrawal?
            // TODO: should we allow the user to deploy their own escrow logic?
            address escrow = Escrow721Handler.create2Deploy(l1Token, tokenId);
            require(Escrow721(escrow).requestEscrow(from, l1Token, tokenId), "NO_ESCROW");
        } else {
            StandardArbERC721(l1Token).burn(tokenId);
        }

        bytes memory outboundCalldata = getOutboundCalldata(l1Token, from, to, tokenId, data);
        // TODO: add exitId for tradeable exits?
        return sendTxToL1(0, from, to, outboundCalldata);
    }

    function finalizeDeposit(
        address l1Token,
        uint256 tokenId,
        address from,
        address to,
        bytes calldata name,
        bytes calldata symbol,
        bytes calldata tokenURI,
        bytes calldata /* data */
    ) external {
        require(msg.sender == AddressAliasHelper.applyL1ToL2Alias(counterpartGateway));
        address expectedAddress = calculateL2TokenAddress(l1Token);

        if (!Address.isContract(expectedAddress)) {
            address createdContract = BeaconProxyFactory(beaconProxyFactory).createProxy(
                keccak256(abi.encodePacked(l1Token))
            );

            if (createdContract == expectedAddress) {
                StandardArbERC721(createdContract).bridgeInit(l1Token, name, symbol);
            } else {
                // trigger withdrawal then halt
                // this codepath should only be hit if the system is setup incorrectly
                // this withdrawal is for error recovery, not composing with L2 dapps, so we ignore the return value
                // TODO: trigger withdrawal
                revert("NOT_IMPLEMENTED");
                // triggerWithdrawal(l1ERC20, address(this), _from, _amount, "");
                // return;
            }
        }
        StandardArbERC721(expectedAddress).mint(to, tokenId, tokenURI);
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    function getCreate2EscrowAddress(address l1Token, uint256 tokenId)
        external
        view
        returns (address)
    {
        return Escrow721Handler.getCreate2EscrowAddress(l1Token, tokenId);
    }

    function updateTokenUriToL2() external {
        // TODO: take in batch
    }

    function getOutboundCalldata(
        address l1Token,
        address from,
        address to,
        uint256 tokenId,
        bytes calldata data
    ) public view returns (bytes memory) {
        // TODO: what if this is a L2 native erc20?
        return
            abi.encodeWithSelector(
                L1NftGateway.finalizeWithdraw.selector,
                l1Token,
                tokenId,
                to,
                data
            );
    }

    /**
     * @notice Calculate the address used when bridging an ERC721 token
     * @dev the L1 and L2 address oracles may not always be in sync.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1Token address of L1 token
     * @return L2 address of a bridged ERC721 token
     */
    function calculateL2TokenAddress(address l1Token) public view returns (address) {
        return
            BeaconProxyFactory(beaconProxyFactory).calculateExpectedAddress(
                address(this),
                keccak256(abi.encodePacked(l1Token))
            );
    }
}
