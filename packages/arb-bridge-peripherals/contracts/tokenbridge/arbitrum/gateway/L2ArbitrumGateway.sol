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

import "@openzeppelin/contracts/utils/Address.sol";
import "arb-bridge-eth/contracts/libraries/BytesLib.sol";

import "../IArbToken.sol";

import "../L2ArbitrumMessenger.sol";
import "../../libraries/gateway/GatewayMessageHandler.sol";
import "../../libraries/gateway/EscrowAndCallGateway.sol";
import "../../libraries/gateway/TokenGateway.sol";

/**
 * @title Common interface for gatways on Arbitrum messaging to L1.
 */
abstract contract L2ArbitrumGateway is L2ArbitrumMessenger, TokenGateway, EscrowAndCallGateway {
    using Address for address;

    uint256 public exitNum;

    function _initialize(address _l1Counterpart, address _router) internal virtual override {
        TokenGateway._initialize(_l1Counterpart, _router);
        // L1 gateway must have a router
        require(_router != address(0), "BAD_ROUTER");
    }

    function gasReserveIfCallRevert() public pure virtual override returns (uint256) {
        // amount of arbgas necessary to send user tokens in case
        // of the "onTokenTransfer" call consumes all available gas
        return 2500;
    }

    function createOutboundTx(
        address _from,
        uint256 _tokenAmount,
        bytes memory _outboundCalldata
    ) internal virtual returns (uint256) {
        // We make this function virtual since outboundTransfer logic is the same for many gateways
        // but sometimes (ie weth) you construct the outgoing message differently.

        // exitNum incremented after being included in _outboundCalldata
        exitNum++;
        return
            sendTxToL1(
                // default to sending no callvalue to the L1
                0,
                _from,
                counterpartGateway,
                _outboundCalldata
            );
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view override returns (bytes memory outboundCalldata) {
        outboundCalldata = abi.encodeWithSelector(
            TokenGateway.finalizeInboundTransfer.selector,
            _token,
            _from,
            _to,
            _amount,
            GatewayMessageHandler.encodeFromL2GatewayMsgV1(exitNum, _data)
        );

        return outboundCalldata;
    }

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) public payable virtual returns (bytes memory) {
        return outboundTransfer(_l1Token, _to, _amount, 0, 0, _data);
    }

    /**
     * @notice Initiates a token withdrawal from Arbitrum to Ethereum
     * @param _l1Token l1 address of token
     * @param _to destination address
     * @param _amount amount of tokens withdrawn
     * @param _maxGas max gas provided for outbox execution market (todo)
     * @param _gasPriceBid provided for outbox execution market (todo)
     @ @return encoded unique identifier for withdrawal
     */

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable virtual override returns (bytes memory res) {
        // This function is set as public and virtual so that subclasses can override
        // it and add custom validation for callers (ie only whitelisted users)

        // the function is marked as payable to conform to the inheritance setup
        // this particular code path shouldn't have a msg.value > 0
        // TODO: remove this invariant for execution markets
        require(msg.value == 0, "NO_VALUE");

        address _from;
        bytes memory _extraData;
        {
            if (isRouter(msg.sender)) {
                (_from, _extraData) = GatewayMessageHandler.parseFromRouterToGateway(_data);
            } else {
                _from = msg.sender;
                _extraData = _data;
            }
        }

        uint256 id;
        {
            address l2Token = calculateL2TokenAddress(_l1Token);
            require(l2Token.isContract(), "TOKEN_NOT_DEPLOYED");

            outboundEscrowTransfer(l2Token, _from, _amount);

            // we override the res field to save on the stack
            res = getOutboundCalldata(_l1Token, _from, _to, _amount, _extraData);
            id = createOutboundTx(_from, _amount, res);
        }
        emit OutboundTransferInitiated(_l1Token, _from, _to, id, _amount, _extraData);
        return abi.encode(id);
    }

    function outboundEscrowTransfer(
        address _l2Token,
        address _from,
        uint256 _amount
    ) internal virtual {
        // this method is virtual since different subclasses can handle escrow differently
        // user funds are escrowed on the gateway using this function
        // burns L2 tokens in order to release escrowed L1 tokens
        IArbToken(_l2Token).bridgeBurn(_from, _amount);
    }

    function inboundEscrowTransfer(
        address _l2Address,
        address _dest,
        uint256 _amount
    ) internal virtual override {
        // this method is virtual since different subclasses can handle escrow differently
        IArbToken(_l2Address).bridgeMint(_dest, _amount);
    }

    /**
     * @notice Mint on L2 upon L1 deposit.
     * If token not yet deployed and symbol/name/decimal data is included, deploys StandardArbERC20
     * @dev Callable only by the L1ERC20Gateway.outboundTransfer method. For initial deployments of a token the L1 L1ERC20Gateway
     * is expected to include the deployData. If not a L1 withdrawal is automatically triggered for the user
     * @param _token L1 address of ERC20
     * @param _from account that initiated the deposit in the L1
     * @param _to account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param _amount token amount to be minted to the user
     * @param _data encoded symbol/name/decimal data for deploy, in addition to any additional callhook data
     */
    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external payable override onlyCounterpartGateway returns (bytes memory) {
        bytes memory gatewayData;
        bytes memory callHookData;

        {
            uint8 version = GatewayMessageHandler.getGatewayMessageVersion(_data);
            if (version == 0) {
                // something went wrong
                revert("PANIC! 0");
            } else if (version == 1) {
                (gatewayData, callHookData) = GatewayMessageHandler.parseFromL1GatewayMsgV1(_data);
            } else {
                // something went wrong, we don't have these yet
                revert("PANIC! > 1");
            }
        }

        address expectedAddress = calculateL2TokenAddress(_token);

        if (!expectedAddress.isContract()) {
            bool shouldHalt = handleNoContract(
                _token,
                expectedAddress,
                _from,
                _to,
                _amount,
                gatewayData
            );
            if (shouldHalt) return bytes("");
        }
        // ignores gatewayData if token already deployed

        {
            // validate if L1 address supplied matches that of the expected L2 address
            (bool success, bytes memory _l1AddressData) = expectedAddress.staticcall(
                abi.encodeWithSelector(IArbToken.l1Address.selector)
            );

            bool shouldWithdraw;
            if (!success || _l1AddressData.length < 32) {
                shouldWithdraw = true;
            } else {
                // we do this in the else branch since we want to avoid reverts
                // and `toAddress` reverts if _l1AddressData has a short length
                // `_l1AddressData` should be 12 bytes of padding then 20 bytes for the address
                address expectedL1Address = BytesLib.toAddress(_l1AddressData, 12);
                if (expectedL1Address != _token) {
                    shouldWithdraw = true;
                }
            }

            if (shouldWithdraw) {
                // we override the gatewayData field to save on the stack
                gatewayData = getOutboundCalldata(_token, address(this), _from, _amount, "");

                createOutboundTx(address(this), _amount, gatewayData);
                return bytes("");
            }
        }

        if (callHookData.length > 0) {
            bool success;
            try this.inboundEscrowAndCall(expectedAddress, _amount, _from, _to, callHookData) {
                success = true;
            } catch {
                // if reverted, then credit _from's account
                inboundEscrowTransfer(expectedAddress, _from, _amount);
                // success default value is false
            }
            emit TransferAndCallTriggered(success, _from, _to, _amount, callHookData);
        } else {
            inboundEscrowTransfer(expectedAddress, _to, _amount);
        }

        emit InboundTransferFinalized(
            _token,
            _from,
            _to,
            uint256(uint160(expectedAddress)),
            _amount,
            _data
        );

        return bytes("");
    }

    // returns if function should halt after
    function handleNoContract(
        address _l1Token,
        address expectedL2Address,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory gatewayData
    ) internal virtual returns (bool shouldHalt);
}
