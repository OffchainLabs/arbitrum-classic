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

import "../../libraries/ITokenBridge.sol";
import "../../libraries/RouterConsumer.sol";

abstract contract L1ArbitrumConsumer is RouterConsumer {
    function createOutboundTx(
        address _handler,
        address _target,
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override returns (bytes memory) {
        // msg.value is sent, but 0 is set to the L2 call value
        // the eth sent is used to pay for the tx's gas
        uint256 seqNum =
            IInbox(_handler).createRetryableTicket{ value: msg.value }(
                _target,
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

contract ERC20Bridge is L1ArbitrumConsumer {
    using SafeERC20 for IERC20;

    function initialize(address _l2Target) public virtual override {
        super.initialize(_l2Target);
        // TODO: add onlyRouter to `outboundTransfer`
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
        address _l2Target,
        address _sender,
        address _destination,
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
            ITokenBridge.finalizeInboundTransfer.selector,
            // _token,
            // sender,
            // destination,
            // amount,
            deployData,
            _data
        );

        return outboundCalldata;
    }

    function triggerEscrow(
        address _token,
        address _from,
        uint256 _amount
    ) internal virtual override {
        IERC20(_token).safeTransferFrom(_from, address(this), _amount);
    }

    function finalizeInboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override returns (bytes memory) {
        revert("NOT_IMPLEMENTED");
    }
}
