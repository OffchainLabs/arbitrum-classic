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

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

import "../../libraries/ITokenGateway.sol";
import "../../libraries/TokenGateway.sol";

abstract contract L1ArbitrumGateway is TokenGateway {
    address router;

    function initialize(address _l2Counterpart, address _router) public virtual {
        super.initialize(_l2Counterpart);
        require(_router != address(0), "BAD_ROUTER");
        require(router == address(0), "ALREADY_INIT");
        router = _router;
    }

    modifier onlyRouter {
        require(msg.sender == router, "ONLY_ROUTER");
        _;
    }

    function createOutboundTx(
        address _inbox,
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (bytes memory) {
        // msg.value is sent, but 0 is set to the L2 call value
        // the eth sent is used to pay for the tx's gas
        uint256 seqNum =
            IInbox(_inbox).createRetryableTicket{ value: msg.value }(
                counterpartGateway,
                0,
                _maxSubmissionCost,
                _user,
                _user,
                _maxGas,
                _gasPriceBid,
                _data
            );
        return abi.encode(seqNum);
    }
}

contract L1ERC20Gateway is L1ArbitrumGateway {
    using SafeERC20 for IERC20;

    function initialize(address _l2Counterpart, address _router) public virtual override {
        super.initialize(_l2Counterpart, _router);
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override onlyRouter returns (bytes memory res) {
        (address _inbox, address _from, uint256 _maxSubmissionCost, bytes memory extraData) =
            parseArbitrumData(_data);

        // escrow funds in gateway
        IERC20(_token).safeTransferFrom(_from, address(this), _amount);

        bytes memory outboundCalldata = getOutboundCalldata(_token, _from, _to, _amount, extraData);

        res = createOutboundTx(
            _inbox,
            _from,
            _maxSubmissionCost,
            _maxGas,
            _gasPriceBid,
            outboundCalldata
        );

        emit OutboundTransferInitiated(_token, _from, _to, _amount, _data);

        return res;
    }

    function parseArbitrumData(bytes memory _data)
        internal
        pure
        virtual
        returns (
            address _inbox,
            address _from,
            uint256 _maxSubmissionCost,
            bytes memory _extraData
        )
    {
        // router encoded
        (_inbox, _from, _extraData) = abi.decode(_data, (address, address, bytes));
        // user encoded
        (_maxSubmissionCost, _extraData) = abi.decode(_extraData, (uint256, bytes));
    }

    /**
     * @notice utility function used to perform external read-only calls.
     * @dev the result is returned even if the call failed, the L2 is expected to
     * identify and deal with this.
     * @return result bytes, even if the call failed.
     */
    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        view
        returns (bytes memory)
    {
        (bool success, bytes memory res) =
            targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata) {
        // TODO: cheaper to make static calls or save isDeployed to storage?
        bytes memory deployData =
            abi.encode(
                callStatic(_token, ERC20.name.selector),
                callStatic(_token, ERC20.symbol.selector),
                callStatic(_token, ERC20.decimals.selector)
            );

        outboundCalldata = abi.encodeWithSelector(
            ITokenGateway.finalizeInboundTransfer.selector,
            _token,
            _from,
            _to,
            _amount,
            abi.encode(deployData, _data)
        );

        return outboundCalldata;
    }

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override onlyCounterpartGateway returns (bytes memory) {
        // TODO: implement withdrawal
        revert("NOT_IMPLEMENTED");
    }
}
