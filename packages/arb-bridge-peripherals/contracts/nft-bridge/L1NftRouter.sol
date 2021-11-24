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

import "../tokenbridge/libraries/gateway/GatewayMessageHandler.sol";
import "../tokenbridge/ethereum/L1ArbitrumMessenger.sol";
import "./L1NftGateway.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

/**
 * how to bridge NFT:
 *  router keeps mapping of bridge (escrow) contract.
 *  allows self-registration or owner to register another escrow contract
 */
contract L1NftRouter is L1ArbitrumMessenger, IERC721Receiver {
    using Address for address;

    address internal constant ZERO_ADDR = address(0);
    address internal constant DISABLED = address(1);

    mapping(bytes32 => address) public l1TokenToGateway;
    address public counterpartGateway;
    address public defaultGateway;
    address public inbox;

    function initialize(
        address _counterpartGateway,
        address _defaultGateway,
        address _inbox
    ) public {
        require(_counterpartGateway != address(0), "INVALID_COUNTERPART");
        require(counterpartGateway == address(0), "ALREADY_INIT");
        counterpartGateway = _counterpartGateway;
        // default gateway can have 0 address
        defaultGateway = _defaultGateway;
        inbox = _inbox;
    }

    // TODO: compare this event topics to erc20. maybe shared topic positions can be interesting
    event TransferRouted(
        uint256 indexed tokenId,
        address indexed tokenAddr,
        address _userFrom,
        address indexed _userTo,
        address gateway
    );

    function deposit(
        address l1Token,
        uint256 tokenId,
        address to,
        bool shouldUseNewEscrowAddress,
        uint256 maxGas,
        uint256 gasPrice,
        uint256 maxSubmissionCost,
        address creditBackAddress,
        bytes calldata data
    ) external payable returns (uint256) {
        address gateway = getGateway(l1Token, tokenId);
        require(gateway != address(0), "INVALID_GATEWAY");
        emit TransferRouted(tokenId, l1Token, msg.sender, to, gateway);
        return
            L1NftGateway(gateway).depositFromRouter{ value: msg.value }(
                l1Token,
                tokenId,
                msg.sender,
                to,
                shouldUseNewEscrowAddress,
                maxGas,
                gasPrice,
                maxSubmissionCost,
                creditBackAddress,
                data
            );
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        revert("NO_COLLATERAL_ROUTER");
        // // we could instead just transfer the collateral to the correct contract
        // IERC721(_token).safeTransferFrom(address(this), gateway, tokenId, gatewayData);
        // return this.onERC721Received.selector;
    }

    function getGateway(address _tokenAddress, uint256 tokenId)
        public
        view
        returns (address gateway)
    {
        bytes32 key = keccak256(abi.encodePacked(_tokenAddress, tokenId));
        gateway = l1TokenToGateway[key];

        if (gateway == ZERO_ADDR) {
            // if no gateway value set, use default gateway
            gateway = defaultGateway;
            // TODO: add default 721 / 1155 gateways. can infer based on 165
            // or have the same gateway handle both?
        }

        if (gateway == DISABLED || !gateway.isContract()) {
            // not a valid gateway
            return ZERO_ADDR;
        }

        return gateway;
    }

    function setDefaultGateway(
        address newL1DefaultGateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost
    ) external {
        // TODO: implement
        revert("NOT_IMPLEMENTED");
    }

    function _setGateways(
        address[] memory _token,
        address[] memory _gateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        address _creditBackAddress
    ) external {
        // TODO: implement
        revert("NOT_IMPLEMENTED");
    }

    function getOutboundCalldata(
        address l1Token,
        address _from,
        address _to,
        uint256 tokenId,
        bytes memory _data
    ) public view returns (bytes memory) {
        address gateway = getGateway(l1Token, tokenId);
        return L1NftGateway(gateway).getOutboundCalldata(l1Token, _from, _to, tokenId, _data);
    }

    function calculateL2TokenAddress(address l1ERC721, uint256 tokenId)
        public
        view
        returns (address)
    {
        address gateway = getGateway(l1ERC721, tokenId);
        if (gateway == ZERO_ADDR) {
            return ZERO_ADDR;
        }
        return L1NftGateway(gateway).calculateL2TokenAddress(l1ERC721, tokenId);
    }
}
