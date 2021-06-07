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

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";

import "../../libraries/gateway/ITokenGateway.sol";
import "../../libraries/gateway/TokenGateway.sol";
import "../../libraries/ClonableBeaconProxy.sol";

abstract contract L1ArbitrumGateway is TokenGateway {
    using SafeERC20 for IERC20;

    address public inbox;

    function isCounterpartGateway() internal view virtual override returns (bool) {
        IOutbox outbox = IOutbox(IInbox(inbox).bridge().activeOutbox());
        return counterpartGateway == outbox.l2ToL1Sender();
    }

    function _initialize(
        address _l2Counterpart,
        address _router,
        address _inbox
    ) internal virtual {
        // L1 gateway must have a router
        require(_router != address(0), "BAD_ROUTER");
        TokenGateway._initialize(_l2Counterpart, _router);
        require(_inbox != address(0), "BAD_INBOX");
        router = _router;
        inbox = _inbox;
    }

    /**
     * @notice Finalizes a withdrawal via Outbox message; callable only by L2Gateway.outboundTransfer
     * @param _token L1 address of token being withdrawn from
     * @param _from initiator of withdrawal
     * @param _to address the L2 withdrawal call set as the destination.
     * @param _amount Token amount being withdrawn
     * @param _data encoded exitNum (Sequentially increasing exit counter determined by the L2Gateway) and additinal hook data
     */
    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override onlyCounterpartGateway returns (bytes memory) {
        (uint256 exitNum, bytes memory extraData) = abi.decode(_data, (uint256, bytes));

        // TODO: add withdraw and call
        // TODO: add transferExit
        IERC20(_token).safeTransfer(_to, _amount);
        emit InboundTransferFinalized(_token, _from, _to, exitNum, _amount, _data);

        return bytes("");
    }

    function createOutboundTx(
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (uint256) {
        // msg.value is sent, but 0 is set to the L2 call value
        // the eth sent is used to pay for the tx's gas
        uint256 seqNum =
            IInbox(inbox).createRetryableTicket{ value: msg.value }(
                counterpartGateway,
                0,
                _maxSubmissionCost,
                _user,
                _user,
                _maxGas,
                _gasPriceBid,
                _data
            );
        return seqNum;
    }

    /**
     * @notice Deposit ERC20 token from Ethereum into Arbitrum. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy. Initiate by GatewayRouter.
     * @param _token L1 address of ERC20
     * @param _to account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param _amount Token Amount
     * @param _maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param _gasPriceBid Gas price for L2 execution
     * @param _data encoded data from router and user
     * @return res abi encoded inbox sequence number
     */
    //  * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override onlyRouter returns (bytes memory res) {
        address _from;
        uint256 seqNum;
        {
            uint256 _maxSubmissionCost;
            bytes memory extraData;
            (_from, _maxSubmissionCost, extraData) = parseOutboundData(_data);

            // escrow funds in gateway
            IERC20(_token).safeTransferFrom(_from, address(this), _amount);

            bytes memory outboundCalldata =
                getOutboundCalldata(_token, _from, _to, _amount, extraData);

            seqNum = createOutboundTx(
                _from,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                outboundCalldata
            );
        }

        emit OutboundTransferInitiated(_token, _from, _to, seqNum, _amount, _data);
        return abi.encode(seqNum);
    }

    function parseOutboundData(bytes memory _data)
        internal
        view
        virtual
        returns (
            address _from,
            uint256 _maxSubmissionCost,
            bytes memory _extraData
        )
    {
        if (isRouter()) {
            // router encoded
            (_from, _extraData) = abi.decode(_data, (address, bytes));
        } else {
            _from = msg.sender;
            _extraData = _data;
        }
        // user encoded
        (_maxSubmissionCost, _extraData) = abi.decode(_extraData, (uint256, bytes));
    }

    function isRouter() internal view virtual override returns (bool) {
        return msg.sender == router;
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata);
}
