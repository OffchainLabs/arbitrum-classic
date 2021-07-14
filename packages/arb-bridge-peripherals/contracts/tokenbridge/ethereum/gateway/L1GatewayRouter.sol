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

import "arb-bridge-eth/contracts/libraries/Whitelist.sol";

import { L1ArbitrumMessenger } from "../../libraries/gateway/ArbitrumMessenger.sol";
import "../../libraries/gateway/GatewayRouter.sol";
import "../../arbitrum/gateway/L2GatewayRouter.sol";

/**
 * @title Handles deposits from Erhereum into Arbitrum. Tokens are routered to their appropriate L1 gateway (Router itself also conforms to the Gateway itnerface).
 * @notice Router also serves as an L1-L2 token address oracle.
 */
contract L1GatewayRouter is WhitelistConsumer, L1ArbitrumMessenger, GatewayRouter {
    address public owner;
    address public inbox;

    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    function initialize(
        address _owner,
        address _defaultGateway,
        address _whitelist,
        address _counterpartGateway,
        address _inbox
    ) public virtual {
        GatewayRouter._initialize(_counterpartGateway, _defaultGateway);
        owner = _owner;
        WhitelistConsumer.whitelist = _whitelist;
        inbox = _inbox;
    }

    function sendTxToL2(
        address _user,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (uint256) {
        return
            L1ArbitrumMessenger.sendTxToL2(
                inbox,
                counterpartGateway,
                _user,
                _l2CallValue,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                _data
            );
    }

    function setDefaultGateway(
        address newL1DefaultGateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost
    ) external payable virtual onlyOwner returns (uint256) {
        defaultGateway = newL1DefaultGateway;

        emit DefaultGatewayUpdated(newL1DefaultGateway);

        address l2NewDefaultGateway;

        if (newL1DefaultGateway != address(0)) {
            l2NewDefaultGateway = TokenGateway(newL1DefaultGateway).counterpartGateway();
        }

        bytes memory data =
            abi.encodeWithSelector(L2GatewayRouter.setDefaultGateway.selector, l2NewDefaultGateway);

        return sendTxToL2(msg.sender, 0, _maxSubmissionCost, _maxGas, _gasPriceBid, data);
    }

    function setOwner(address newOwner) external onlyOwner {
        require(newOwner != address(0), "INVALID_OWNER");
        // set newOwner to address(1) to disable owner and keep `initialize` safe
        owner = newOwner;
    }

    function _setGateways(
        address[] memory _token,
        address[] memory _gateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        address _creditBackAddress
    ) internal returns (uint256) {
        require(_token.length == _gateway.length, "WRONG_LENGTH");

        for (uint256 i = 0; i < _token.length; i++) {
            l1TokenToGateway[_token[i]] = _gateway[i];
            emit GatewaySet(_token[i], _gateway[i]);
            // overwrite memory so the L2 router receives the L2 address of each gateway
            if (_gateway[i] != address(0)) {
                _gateway[i] = TokenGateway(_gateway[i]).counterpartGateway();
            }
        }

        bytes memory data =
            abi.encodeWithSelector(L2GatewayRouter.setGateway.selector, _token, _gateway);

        return sendTxToL2(_creditBackAddress, 0, _maxSubmissionCost, _maxGas, _gasPriceBid, data);
    }

    /**
     * @notice Allows L1 Token contract to trustlessly register its gateway. (other setGateway method allows excess eth recovery from _maxSubmissionCost and is recommended)

     * @param _gateway l1 gateway address
     * @param _maxGas max gas for L2 retryable exrecution 
     * @param _gasPriceBid gas price for L2 retryable ticket 
     * @param  _maxSubmissionCost base submission cost  L2 retryable tick3et 
     * @return Retryable ticket ID
     */
    function setGateway(
        address _gateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost
    ) external payable returns (uint256) {
        return setGateway(_gateway, _maxGas, _gasPriceBid, _maxSubmissionCost, msg.sender);
    }

    /**
     * @notice Allows L1 Token contract to trustlessly register its gateway.

     * @param _gateway l1 gateway address
     * @param _maxGas max gas for L2 retryable exrecution 
     * @param _gasPriceBid gas price for L2 retryable ticket 
     * @param  _maxSubmissionCost base submission cost  L2 retryable tick3et 
     * @param _creditBackAddress address for crediting back overpayment of _maxSubmissionCost
     * @return Retryable ticket ID
     */
    function setGateway(
        address _gateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        address _creditBackAddress
    ) public payable returns (uint256) {
        require(address(msg.sender).isContract(), "NOT_FROM_CONTRACT");
        require(_gateway.isContract(), "NOT_TO_CONTRACT");

        address[] memory _tokenArr = new address[](1);
        _tokenArr[0] = address(msg.sender);

        address[] memory _gatewayArr = new address[](1);
        _gatewayArr[0] = _gateway;

        return
            _setGateways(
                _tokenArr,
                _gatewayArr,
                _maxGas,
                _gasPriceBid,
                _maxSubmissionCost,
                _creditBackAddress
            );
    }

    function setGateways(
        address[] memory _token,
        address[] memory _gateway,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost
    ) external payable onlyOwner returns (uint256) {
        // it is assumed that token and gateway are both contracts
        // require(_token[i].isContract() && _gateway[i].isContract(), "NOT_CONTRACT");
        return
            _setGateways(_token, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost, msg.sender);
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable virtual override onlyWhitelisted returns (bytes memory) {
        // will revert if msg.sender is not whitelisted
        super.outboundTransfer(_token, _to, _amount, _maxGas, _gasPriceBid, _data);
    }

    function isCounterpartGateway(address _target) internal view virtual override returns (bool) {
        // don't expect messages from L2 router
        return false;
    }
}
