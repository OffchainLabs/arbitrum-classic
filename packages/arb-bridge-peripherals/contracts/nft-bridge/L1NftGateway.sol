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
import "@openzeppelin/contracts/token/ERC721/IERC721Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

import "./L2NftGateway.sol";
import "../tokenbridge/libraries/gateway/GatewayMessageHandler.sol";
import "../tokenbridge/libraries/Escrow721.sol";
import "../tokenbridge/ethereum/L1ArbitrumMessenger.sol";
import "./L1NftRouter.sol";

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

    modifier onlyCounterpartGateway() {
        address _inbox = inbox;

        // a message coming from the counterpart gateway was executed by the bridge
        address bridge = address(super.getBridge(_inbox));
        require(msg.sender == bridge, "NOT_FROM_BRIDGE");

        // and the outbox reports that the L2 address of the sender is the counterpart gateway
        address l2ToL1Sender = super.getL2ToL1Sender(_inbox);
        require(l2ToL1Sender == counterpartGateway, "ONLY_COUNTERPART_GATEWAY");
        _;
    }

    function finalizeWithdraw(
        address l1Token,
        uint256 tokenId,
        address to,
        bytes calldata data
    ) external onlyCounterpartGateway {
        // TODO: implement tradeable exits?
        // TODO: what if NFT is L2 native?

        address escrow = Escrow721Handler.getCreate2EscrowAddress(l1Token, tokenId);
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
            // TODO: can we just check if the token is escrowed there, and allow this to be deployed during an eventual withdrawal?
            // TODO: should we allow the user to deploy their own escrow logic?
            address escrow = Escrow721Handler.create2Deploy(l1Token, tokenId);
            require(Escrow721(escrow).requestEscrow(from, l1Token, tokenId), "NO_ESCROW");
        } else {
            IERC721(l1Token).transferFrom(from, address(this), tokenId);
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

    function getCreate2EscrowAddress(address l1Token, uint256 tokenId)
        external
        view
        returns (address)
    {
        return Escrow721Handler.getCreate2EscrowAddress(l1Token, tokenId);
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        // this shouldn't be triggered since we don't do a safe `transferFrom`
        revert("INVALID_DEPOSIT");
        // return this.onERC721Received.selector;
    }

    function updateTokenUriToL2() external {
        // TODO: allow permissionless updating of URI in L2
    }

    function getOutboundCalldata(
        address l1Token,
        address from,
        address to,
        uint256 tokenId,
        bytes calldata data
    ) public view returns (bytes memory) {
        (, bytes memory tokenURI) = l1Token.staticcall(
            abi.encodeWithSelector(IERC721Metadata.tokenURI.selector, tokenId)
        );
        // TODO: is it cheaper to only send these once?
        (, bytes memory name) = l1Token.staticcall(
            abi.encodeWithSelector(IERC721Metadata.name.selector)
        );
        (, bytes memory symbol) = l1Token.staticcall(
            abi.encodeWithSelector(IERC721Metadata.symbol.selector)
        );
        // TODO: should we send baseUri? not part of the standard, but OZ implements it
        // l1Token.staticcall(abi.encodeWithSelector(IERC721Metadata.baseURI.selector)),

        return
            abi.encodeWithSelector(
                L2NftGateway.finalizeDeposit.selector,
                l1Token,
                tokenId,
                from,
                to,
                name,
                symbol,
                tokenURI,
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
