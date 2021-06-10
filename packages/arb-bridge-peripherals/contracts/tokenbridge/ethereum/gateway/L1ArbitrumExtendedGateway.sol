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

import "../../libraries/IERC677.sol";

import "./L1ArbitrumGateway.sol";

interface ITradeableExitReceiver {
    function onExitTransfer(
        address sender,
        uint256 amount,
        address erc20,
        uint256 exitNum,
        bytes calldata data
    ) external returns (bool);
}

abstract contract L1ArbitrumExtendedGateway is L1ArbitrumGateway {
    address internal constant USED_ADDRESS = address(0x01);
    // exitNum to redirected address
    mapping(uint256 => address) public redirectedExits;

    event WithdrawRedirected(
        address indexed from,
        address indexed to,
        address erc20,
        uint256 amount,
        uint256 indexed exitNum,
        bool madeExternalCall
    );

    /**
     * @notice Allows a user to redirect their right to claim a withdrawal to another address.
     * @dev This method also allows you to make an arbitrary call after the transfer, similar to ERC677.
     * This does not change the original data that will be triggered with the withdrawal's external call.
     * @param _from address the L2 withdrawal call initially set as the destination.
     * @param _token L1 token address
     * @param _amount token amount (should match amount in previously-initiated withdrawal)
     * @param _exitNum Sequentially increasing exit counter determined by the L2 bridge
     * @param _initialDestination address the L2 withdrawal call initially set as the destination.
     * @param _newDestination address the L1 will now call instead of the previously set destination
     * @param _data optional data for external call upon transfering the exit
     */
    function transferExitAndCall(
        address _from,
        address _token,
        uint256 _amount,
        uint256 _exitNum,
        address _initialDestination,
        address _newDestination,
        bytes calldata _data
    ) external virtual {
        // if you want to transfer your exit, you must be the current destination
        address expectedSender = getCurrentDestination(_exitNum, _initialDestination);
        require(msg.sender == expectedSender, "NOT_EXPECTED_SENDER");

        updateDestination(_exitNum, _newDestination);

        if (_data.length > 0) {
            require(_newDestination.isContract(), "TO_NOT_CONTRACT");
            bool success =
                ITradeableExitReceiver(_newDestination).onExitTransfer(
                    expectedSender,
                    _amount,
                    _token,
                    _exitNum,
                    _data
                );
            require(success, "TRANSFER_HOOK_FAIL");
        }

        emit WithdrawRedirected(
            expectedSender,
            _newDestination,
            _token,
            _amount,
            _exitNum,
            _data.length > 0
        );
    }

    function getCurrentDestination(uint256 _exitNum, address _initialDestination)
        public
        view
        virtual
        override
        returns (address)
    {
        // here we assume the L2 bridge gives a unique exitNum to each exit
        address redirectedAddress = redirectedExits[_exitNum];
        require(redirectedAddress != USED_ADDRESS, "ALREADY_EXITED");
        return redirectedAddress == address(0) ? _initialDestination : redirectedAddress;
    }

    function updateDestination(uint256 _exitNum, address _newDestination) internal virtual {
        redirectedExits[_exitNum] = _newDestination;
    }
}
