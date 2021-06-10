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

import "./L2ArbitrumMessenger.sol";
import "../IArbToken.sol";

import "../../libraries/gateway/ITokenGateway.sol";
import "../../libraries/gateway/TokenGateway.sol";
import "../../libraries/IERC677.sol";

/**
 * @title Common interface for gatways on Arbitrum messaging to L1.
 */
abstract contract L2ArbitrumGateway is L2ArbitrumMessenger, TokenGateway {
    using Address for address;

    uint256 public exitNum;

    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function _initialize(address _l1Counterpart, address _router) internal virtual override {
        // L2 gateway may have a router address(0)
        TokenGateway._initialize(_l1Counterpart, _router);
    }

    function arbgasReserveIfCallRevert() internal pure virtual returns (uint256) {
        // amount of arbgas necessary to send user tokens in case
        // of the "onTokenTransfer" call consumes all available gas
        return 2500;
    }

    function createOutboundTx(
        address _l1Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _extraData
    ) internal virtual returns (uint256) {
        return sendTxToL1(_from, 0, getOutboundCalldata(_l1Token, _from, _to, _amount, _extraData));
    }

    function sendTxToL1(
        address _from,
        uint256 _l1CallValue,
        bytes memory _data
    ) internal virtual returns (uint256) {
        return sendTxToL1(_l1CallValue, _from, counterpartGateway, _data);
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata) {
        outboundCalldata = abi.encodeWithSelector(
            ITokenGateway.finalizeInboundTransfer.selector,
            _token,
            _from,
            _to,
            _amount,
            abi.encode(exitNum, _data)
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
    ) public payable virtual override returns (bytes memory) {
        // can be triggered directly or by router
        require(msg.value == 0, "NO_VALUE");

        (address _from, bytes memory _extraData) = parseOutboundData(_data);

        uint256 id;
        {
            address l2Token = _calculateL2TokenAddress(_l1Token);
            require(l2Token.isContract(), "TOKEN_NOT_DEPLOYED");

            outboundEscrowTransfer(l2Token, _from, _amount);

            id = createOutboundTx(_l1Token, _from, _to, _amount, _extraData);
        }
        // exitNum incremented after being used in createOutboundTx
        exitNum++;
        emit OutboundTransferInitiated(_l1Token, _from, _to, id, _amount, _extraData);
        return abi.encode(id);
    }

    function outboundEscrowTransfer(
        address _l2Token,
        address _from,
        uint256 _amount
    ) internal virtual {
        // burns L2 tokens in order to release escrowed L1 tokens
        IArbToken(_l2Token).bridgeBurn(_from, _amount);
    }

    function parseOutboundData(bytes memory _data)
        internal
        view
        virtual
        returns (address _from, bytes memory _extraData)
    {
        if (isSenderRouter()) {
            (_from, _extraData) = abi.decode(_data, (address, bytes));
        } else {
            _from = msg.sender;
            _extraData = _data;
        }
    }

    /**
     * @notice this function can only be callable by the bridge itself
     * @dev This method is inspired by EIP 677/1363 for calls to be executed after minting.
     * A reserve amount of gas is always kept in case this call reverts or uses up all gas.
     * The reserve is the amount of gas needed to catch the revert and do the necessary alternative logic.
     */
    function mintAndCall(
        address _l2Address,
        uint256 _amount,
        address _sender,
        address _dest,
        bytes memory _data
    ) external virtual {
        require(msg.sender == address(this), "Mint can only be called by self");
        require(_dest.isContract(), "Destination must be a contract");

        inboundEscrowTransfer(_l2Address, _dest, _amount);

        // ~73 000 arbgas used to get here
        uint256 gasAvailable = gasleft() - arbgasReserveIfCallRevert();
        require(gasleft() > gasAvailable, "Mint and call gas left calculation undeflow");

        IERC677Receiver(_dest).onTokenTransfer{ gas: gasAvailable }(_sender, _amount, _data);
    }

    function inboundEscrowTransfer(
        address _l2Address,
        address _dest,
        uint256 _amount
    ) internal virtual {
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
    ) external payable virtual override onlyCounterpartGateway returns (bytes memory) {
        (bytes memory gatewayData, bytes memory callHookData) = abi.decode(_data, (bytes, bytes));

        address expectedAddress = _calculateL2TokenAddress(_token);

        if (!expectedAddress.isContract()) {
            bool shouldHalt =
                handleNoContract(_token, expectedAddress, _from, _to, _amount, gatewayData);
            if (shouldHalt) return bytes("");
        }
        // ignores gatewayData if token already deployed

        if (callHookData.length > 0) {
            bool success;
            try this.mintAndCall(expectedAddress, _amount, _from, _to, callHookData) {
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

    function isSenderRouter() internal view virtual override returns (bool) {
        return msg.sender == router;
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
