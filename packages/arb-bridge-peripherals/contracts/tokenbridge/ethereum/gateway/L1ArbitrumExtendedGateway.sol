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

import "../../libraries/ITransferAndCall.sol";

import "./L1ArbitrumGateway.sol";

interface ITradeableExitReceiver {
    function onExitTransfer(
        address sender,
        uint256 exitNum,
        bytes calldata data
    ) external returns (bool);
}

abstract contract L1ArbitrumExtendedGateway is L1ArbitrumGateway {
    struct ExitData {
        bool isExit;
        address _newTo;
        bytes _newData;
    }

    mapping(bytes32 => ExitData) public redirectedExits;

    event WithdrawRedirected(
        address indexed from,
        address indexed to,
        uint256 indexed exitNum,
        bytes newData,
        bytes data,
        bool madeExternalCall
    );

    /**
     * @notice Allows a user to redirect their right to claim a withdrawal to another address.
     * @dev This method also allows you to make an arbitrary call after the transfer.
     * This does not validate if the exit was already triggered. It is assumed the `_exitNum` is
     * validated off-chain to ensure this was not yet triggered.
     * @param _exitNum Sequentially increasing exit counter determined by the L2 bridge
     * @param _initialDestination address the L2 withdrawal call initially set as the destination.
     * @param _newDestination address the L1 will now call instead of the previously set destination
     * @param _newData data to be used in inboundEscrowAndCall
     * @param _data optional data for external call upon transfering the exit
     */
    function transferExitAndCall(
        uint256 _exitNum,
        address _initialDestination,
        address _newDestination,
        bytes calldata _newData,
        bytes calldata _data
    ) external {
        // the initial data doesn't make a difference when transfering you exit
        // since the L2 bridge gives a unique exit ID to each exit
        (address expectedSender, ) = getExternalCall(_exitNum, _initialDestination, "");

        // if you want to transfer your exit, you must be the current destination
        require(msg.sender == expectedSender, "NOT_EXPECTED_SENDER");
        // the inboundEscrowAndCall functionality has been disabled, so no data is allowed
        require(_newData.length == 0, "NO_DATA_ALLOWED");

        setRedirectedExit(_exitNum, _initialDestination, _newDestination, _newData);

        if (_data.length > 0) {
            require(_newDestination.isContract(), "TO_NOT_CONTRACT");
            bool success = ITradeableExitReceiver(_newDestination).onExitTransfer(
                expectedSender,
                _exitNum,
                _data
            );
            require(success, "TRANSFER_HOOK_FAIL");
        }

        emit WithdrawRedirected(
            expectedSender,
            _newDestination,
            _exitNum,
            _newData,
            _data,
            _data.length > 0
        );
    }

    /// @notice this does not verify if the external call was already done
    function getExternalCall(
        uint256 _exitNum,
        address _initialDestination,
        bytes memory _initialData
    ) public view virtual override returns (address target, bytes memory data) {
        // this function is virtual so that subclasses can override it with custom logic where necessary
        bytes32 withdrawData = encodeWithdrawal(_exitNum, _initialDestination);
        ExitData storage exit = redirectedExits[withdrawData];

        // here we don't authenticate `_initialData`. we could hash it into `withdrawData` but would increase gas costs
        // this is safe because if the exit isn't overriden, the _initialData coming from L2 is trusted
        // but if the exit is traded, all we care about is the latest user calldata
        if (exit.isExit) {
            return (exit._newTo, exit._newData);
        } else {
            return (_initialDestination, _initialData);
        }
    }

    function setRedirectedExit(
        uint256 _exitNum,
        address _initialDestination,
        address _newDestination,
        bytes memory _newData
    ) internal {
        bytes32 withdrawData = encodeWithdrawal(_exitNum, _initialDestination);
        redirectedExits[withdrawData] = ExitData(true, _newDestination, _newData);
    }

    function encodeWithdrawal(uint256 _exitNum, address _initialDestination)
        public
        pure
        returns (bytes32)
    {
        // here we assume the L2 bridge gives a unique exitNum to each exit
        return keccak256(abi.encode(_exitNum, _initialDestination));
    }
}
