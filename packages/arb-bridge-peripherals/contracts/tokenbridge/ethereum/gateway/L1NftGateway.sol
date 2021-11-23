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

import "@openzeppelin/contracts/introspection/IERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

import "../../arbitrum/gateway/L2NftGateway.sol";
import "../../libraries/gateway/GatewayMessageHandler.sol";
import "../L1ArbitrumMessenger.sol";
import "./L1NftRouter.sol";

contract Escrow721 is IERC721Receiver {
    address owner;

    constructor() public {}

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    function requestEscrow(
        address user,
        address tokenAddr,
        uint256 tokenId
    ) external returns (bool) {
        require(IERC165(tokenAddr).supportsInterface(0x80ac58cd), "165_INTERFACE_NOT_DETECTED");
        owner = msg.sender;
        // TODO: should we check if was transfered here before getting created?
        IERC721(tokenAddr).safeTransferFrom(user, address(this), tokenId);
    }

    function releaseEscrow(
        address to,
        address tokenAddr,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bool) {
        require(msg.sender == owner, "NOT_OWNER");
        IERC721(tokenAddr).safeTransferFrom(address(this), to, tokenId, data);
        return true;
    }
}

/**
 * @title Common interface for L1 and L2 Gateway Routers
 */
contract L1NftGateway is L1ArbitrumMessenger, IERC721Receiver {
    address public counterpartGateway;
    address public inbox;
    address public router;

    function initialize(
        address _counterpartGateway,
        address _router,
        address _inbox
    ) public {
        require(counterpartGateway == address(0), "ALREADY_INIT");
        require(_counterpartGateway != address(0), "BAD_COUNTERPART");
        require(_router != address(0), "BAD_ROUTER");
        require(_inbox != address(0), "BAD_INBOX");
        counterpartGateway = _counterpartGateway;
        router = _router;
        inbox = _inbox;
    }

    function finalizeWithdraw(
        address l1Token,
        uint256 tokenId,
        address to,
        bytes calldata data
    ) external {
        address _inbox = inbox;
        IOutbox outbox = IOutbox(getBridge(_inbox).activeOutbox());
        require(msg.sender == address(outbox), "NOT_OUTBOX");
        L1ArbitrumMessenger.getL2ToL1Sender(_inbox);
        // TODO: implement tradeable exits?
        // TODO: what if NFT is L2 native?

        address escrow = getCreate2EscrowAddress(l1Token, tokenId);
        if (Address.isContract(escrow)) {
            Escrow721(escrow).releaseEscrow(to, l1Token, tokenId, data);
        } else {
            IERC721(l1Token).safeTransferFrom(address(this), to, tokenId, data);
        }
    }

    function depositFromRouter(
        address l1Token,
        uint256 tokenId,
        address from,
        address to,
        bool shouldUseNewEscrowAddress,
        uint256 maxGas,
        uint256 gasPrice,
        uint256 maxSubmissionCost,
        address creditBackAddress,
        bytes calldata data
    ) external payable returns (uint256) {
        require(msg.sender == router, "ONLY_ROUTER");
        return
            depositImpl(
                l1Token,
                tokenId,
                from,
                to,
                shouldUseNewEscrowAddress,
                L1ArbitrumMessenger.L2GasParams({
                    _maxSubmissionCost: maxSubmissionCost,
                    _gasPriceBid: gasPrice,
                    _maxGas: maxGas
                }),
                creditBackAddress,
                data
            );
    }

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
        address expectedGateway = L1NftRouter(router).getGateway(l1Token, tokenId);
        require(expectedGateway == address(this), "INVALID_GATEWAY");

        return
            depositImpl(
                l1Token,
                tokenId,
                msg.sender,
                to,
                shouldUseNewEscrowAddress,
                L1ArbitrumMessenger.L2GasParams({
                    _maxSubmissionCost: maxSubmissionCost,
                    _gasPriceBid: gasPrice,
                    _maxGas: maxGas
                }),
                creditBackAddress,
                data
            );
    }

    function depositImpl(
        address l1Token,
        uint256 tokenId,
        address from,
        address to,
        bool shouldCreate2Escrow,
        L1ArbitrumMessenger.L2GasParams memory l2GasParams,
        address creditBackAddress,
        bytes calldata data
    ) internal returns (uint256) {
        // when sending a L1 to L2 transaction, we expect the user to send
        // eth in flight in order to pay for L2 gas costs
        // this check prevents users from misconfiguring the msg.value
        require(
            msg.value ==
                (l2GasParams._maxGas * l2GasParams._gasPriceBid) + l2GasParams._maxSubmissionCost,
            "WRONG_MSG_VAL"
        );
        require(!shouldCreate2Escrow, "EXTERNAL_ESCROW_DISABLED");

        if (shouldCreate2Escrow) {
            address escrow = create2Deploy(l1Token, tokenId);
            require(Escrow721(escrow).requestEscrow(from, l1Token, tokenId), "NO_ESCROW");
        } else {
            IERC721(l1Token).safeTransferFrom(from, address(this), tokenId);
        }

        bytes memory outboundCalldata = getOutboundCalldata(l1Token, from, to, tokenId, data);
        return
            sendTxToL2(
                inbox,
                counterpartGateway,
                creditBackAddress,
                msg.value, // l1CallValue
                0, // l2CallValue
                l2GasParams,
                outboundCalldata
            );
    }

    function create2Deploy(address l1Token, uint256 tokenId) internal returns (address) {
        // "The pair (contract address, uint256 tokenId) [...] globally unique"
        // ~ https://eips.ethereum.org/EIPS/eip-721#rationale
        bytes32 salt = keccak256(abi.encodePacked(l1Token, tokenId));
        bytes memory bytecode = type(Escrow721).creationCode;
        return Create2.deploy(0, salt, bytecode);
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        // TODO: should we check with the router that this is a valid gateway?
        // should we SSTORE here instead of approval/transferfrom?
        return this.onERC721Received.selector;
    }

    function getCreate2EscrowAddress(address l1Token, uint256 tokenId)
        public
        view
        returns (address)
    {
        // TODO: do the create2 in a factory so that the address oracle doesn't break on upgrades
        bytes32 salt = keccak256(abi.encodePacked(l1Token, tokenId));
        // TODO: hash this during compiletime and inline
        bytes32 bytecodeHash = keccak256(type(Escrow721).creationCode);
        return Create2.computeAddress(salt, bytecodeHash);
    }

    function updateBaseUriToL2() external {
        // TODO: should we also allow update name/symbol?
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
        // TODO: query the 721 for "name" / "symbol" / "uri"
        return
            abi.encodeWithSelector(
                L2NftGateway.finalizeDeposit.selector,
                l1Token,
                tokenId,
                from,
                to,
                data
            );
    }

    function calculateL2TokenAddress(address l1ERC721, uint256 tokenId)
        public
        view
        returns (address)
    {
        // TODO: implement address oracle
        revert("NOT_IMPLEMENTED");
    }
}
