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
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/libraries/ProxyUtil.sol";

import "../L1ArbitrumMessenger.sol";
import "../../libraries/gateway/GatewayMessageHandler.sol";
import "../../libraries/gateway/TokenGateway.sol";
import "../../libraries/ITransferAndCall.sol";

/**
 * @title Common interface for gatways on L1 messaging to Arbitrum.
 */
abstract contract L1ArbitrumGateway is L1ArbitrumMessenger, TokenGateway {
    using SafeERC20 for IERC20;
    using Address for address;

    address public inbox;

    event DepositInitiated(
        address l1Token,
        address indexed _from,
        address indexed _to,
        uint256 indexed _sequenceNumber,
        uint256 _amount
    );

    event WithdrawalFinalized(
        address l1Token,
        address indexed _from,
        address indexed _to,
        uint256 indexed _exitNum,
        uint256 _amount
    );

    modifier onlyCounterpartGateway() override {
        address _inbox = inbox;

        // a message coming from the counterpart gateway was executed by the bridge
        address bridge = address(super.getBridge(_inbox));
        require(msg.sender == bridge, "NOT_FROM_BRIDGE");

        // and the outbox reports that the L2 address of the sender is the counterpart gateway
        address l2ToL1Sender = super.getL2ToL1Sender(_inbox);
        require(l2ToL1Sender == counterpartGateway, "ONLY_COUNTERPART_GATEWAY");
        _;
    }

    function postUpgradeInit() external {
        // it is assumed the L1 Arbitrum Gateway contract is behind a Proxy controlled by a proxy admin
        // this function can only be called by the proxy admin contract
        address proxyAdmin = ProxyUtil.getProxyAdmin();
        require(msg.sender == proxyAdmin, "NOT_FROM_ADMIN");
        // this has no other logic since the current upgrade doesn't require this logic
    }

    function _initialize(
        address _l2Counterpart,
        address _router,
        address _inbox
    ) internal {
        TokenGateway._initialize(_l2Counterpart, _router);
        // L1 gateway must have a router
        require(_router != address(0), "BAD_ROUTER");
        require(_inbox != address(0), "BAD_INBOX");
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
    ) public payable virtual override onlyCounterpartGateway {
        // this function is marked as virtual so superclasses can override it to add modifiers
        (uint256 exitNum, bytes memory callHookData) = GatewayMessageHandler.parseToL1GatewayMsg(
            _data
        );

        if (callHookData.length != 0) {
            // callHookData should always be 0 since inboundEscrowAndCall is disabled
            callHookData = bytes("");
        }

        // we ignore the returned data since the callHook feature is now disabled
        (_to, ) = getExternalCall(exitNum, _to, callHookData);
        inboundEscrowTransfer(_token, _to, _amount);

        emit WithdrawalFinalized(_token, _from, _to, exitNum, _amount);
    }

    function getExternalCall(
        uint256, /* _exitNum */
        address _initialDestination,
        bytes memory _initialData
    ) public view virtual returns (address target, bytes memory data) {
        // this method is virtual so the destination of a call can be changed
        // using tradeable exits in a subclass (L1ArbitrumExtendedGateway)
        target = _initialDestination;
        data = _initialData;
    }

    function inboundEscrowTransfer(
        address _l1Token,
        address _dest,
        uint256 _amount
    ) internal virtual {
        // this method is virtual since different subclasses can handle escrow differently
        IERC20(_l1Token).safeTransfer(_dest, _amount);
    }

    /**
     * @dev Only excess gas is refunded to the _refundTo account, l2 call value is always returned to the _to account
     */
    function createOutboundTxCustomRefund(
        address _refundTo,
        address _from,
        uint256, /* _tokenAmount */
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        bytes memory _outboundCalldata
    ) internal virtual returns (uint256) {
        // We make this function virtual since outboundTransfer logic is the same for many gateways
        // but sometimes (ie weth) you construct the outgoing message differently.

        // msg.value is sent, but 0 is set to the L2 call value
        // the eth sent is used to pay for the tx's gas
        return
            sendTxToL2CustomRefund(
                inbox,
                counterpartGateway,
                _refundTo,
                _from,
                msg.value, // we forward the L1 call value to the inbox
                0, // l2 call value 0 by default
                L2GasParams({
                    _maxSubmissionCost: _maxSubmissionCost,
                    _maxGas: _maxGas,
                    _gasPriceBid: _gasPriceBid
                }),
                _outboundCalldata
            );
    }

    /**
     * @notice DEPRECATED - look at createOutboundTxCustomRefund instead
     */
    function createOutboundTx(
        address _from,
        uint256 _tokenAmount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        bytes memory _outboundCalldata
    ) internal returns (uint256) {
        return
            createOutboundTxCustomRefund(
                _from,
                _from,
                _tokenAmount,
                _maxGas,
                _gasPriceBid,
                _maxSubmissionCost,
                _outboundCalldata
            );
    }

    /**
     * @notice DEPRECATED - look at outboundTransferCustomRefund instead
     */
    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable override returns (bytes memory res) {
        return
            outboundTransferCustomRefund(_l1Token, _to, _to, _amount, _maxGas, _gasPriceBid, _data);
    }

    /**
     * @notice Deposit ERC20 token from Ethereum into Arbitrum. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy. Initiate by GatewayRouter.
     * @dev L2 address alias will not be applied to the following types of addresses on L1:
     *      - an externally-owned account
     *      - a contract in construction
     *      - an address where a contract will be created
     *      - an address where a contract lived, but was destroyed
     * @param _l1Token L1 address of ERC20
     * @param _refundTo Account, or its L2 alias if it have code in L1, to be credited with excess gas refund in L2
     * @param _to Account to be credited with the tokens in the L2 (can be the user's L2 account or a contract), not subject to L2 aliasing
                  This account, or its L2 alias if it have code in L1, will also be able to cancel the retryable ticket and receive callvalue refund
     * @param _amount Token Amount
     * @param _maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param _gasPriceBid Gas price for L2 execution
     * @param _data encoded data from router and user
     * @return res abi encoded inbox sequence number
     */
    //  * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
    function outboundTransferCustomRefund(
        address _l1Token,
        address _refundTo,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable virtual override returns (bytes memory res) {
        require(isRouter(msg.sender), "NOT_FROM_ROUTER");
        // This function is set as public and virtual so that subclasses can override
        // it and add custom validation for callers (ie only whitelisted users)
        address _from;
        uint256 seqNum;
        bytes memory extraData;
        {
            uint256 _maxSubmissionCost;
            if (super.isRouter(msg.sender)) {
                // router encoded
                (_from, extraData) = GatewayMessageHandler.parseFromRouterToGateway(_data);
            } else {
                _from = msg.sender;
                extraData = _data;
            }
            // user encoded
            (_maxSubmissionCost, extraData) = abi.decode(extraData, (uint256, bytes));
            // the inboundEscrowAndCall functionality has been disabled, so no data is allowed
            require(extraData.length == 0, "EXTRA_DATA_DISABLED");

            require(_l1Token.isContract(), "L1_NOT_CONTRACT");
            address l2Token = calculateL2TokenAddress(_l1Token);
            require(l2Token != address(0), "NO_L2_TOKEN_SET");

            _amount = outboundEscrowTransfer(_l1Token, _from, _amount);

            // we override the res field to save on the stack
            res = getOutboundCalldata(_l1Token, _from, _to, _amount, extraData);

            seqNum = createOutboundTxCustomRefund(
                _refundTo,
                _from,
                _amount,
                _maxGas,
                _gasPriceBid,
                _maxSubmissionCost,
                res
            );
        }
        emit DepositInitiated(_l1Token, _from, _to, seqNum, _amount);
        return abi.encode(seqNum);
    }

    function outboundEscrowTransfer(
        address _l1Token,
        address _from,
        uint256 _amount
    ) internal virtual returns (uint256 amountReceived) {
        // this method is virtual since different subclasses can handle escrow differently
        // user funds are escrowed on the gateway using this function
        uint256 prevBalance = IERC20(_l1Token).balanceOf(address(this));
        IERC20(_l1Token).safeTransferFrom(_from, address(this), _amount);
        uint256 postBalance = IERC20(_l1Token).balanceOf(address(this));
        return SafeMath.sub(postBalance, prevBalance);
    }

    function getOutboundCalldata(
        address _l1Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata) {
        // this function is public so users can query how much calldata will be sent to the L2
        // before execution
        // it is virtual since different gateway subclasses can build this calldata differently
        // ( ie the standard ERC20 gateway queries for a tokens name/symbol/decimals )
        bytes memory emptyBytes = "";

        outboundCalldata = abi.encodeWithSelector(
            ITokenGateway.finalizeInboundTransfer.selector,
            _l1Token,
            _from,
            _to,
            _amount,
            GatewayMessageHandler.encodeToL2GatewayMsg(emptyBytes, _data)
        );

        return outboundCalldata;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC165)
        returns (bool)
    {
        // registering interfaces that is added after arb-bridge-peripherals >1.0.11
        // using function selector instead of single function interfaces to reduce bloat
        return
            interfaceId == this.outboundTransferCustomRefund.selector ||
            ERC165.supportsInterface(interfaceId);
    }
}
