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

import "../ethereum/gateway/L1WethGateway.sol";
import "../ethereum/gateway/L1CustomGateway.sol";
import "../ethereum/gateway/L1ERC20Gateway.sol";

import "../arbitrum/gateway/L2WethGateway.sol";
import "../arbitrum/gateway/L2CustomGateway.sol";
import "../arbitrum/gateway/L2ERC20Gateway.sol";

contract L1GatewayTester is L1ERC20Gateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL2(
        address _user,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) = counterpartGateway.call(_data);
        assembly {
            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch success
                case 0 {
                    revert(0, retdata)
                }
        }
        return 1337;
    }
}

contract L2GatewayTester is L2ERC20Gateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL1(
        address _from,
        uint256 _l1CallValue,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) =
            counterpartGateway.call{ value: _l1CallValue }(_data);
        require(success, "OUTBOUND_REVERT");
        return 1337;
    }

    function arbgasReserveIfCallRevert() internal pure virtual override returns (uint256) {
        return 50000;
    }
}

contract L1CustomGatewayTester is L1CustomGateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL2(
        address _user,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) = counterpartGateway.call(_data);
        require(success, "OUTBOUND_REVERT");
        return 1337;
    }
}

contract L2CustomGatewayTester is L2CustomGateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL1(
        address _from,
        uint256 _l1CallValue,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) =
            counterpartGateway.call{ value: _l1CallValue }(_data);
        require(success, "OUTBOUND_REVERT");
        return 1337;
    }

    function arbgasReserveIfCallRevert() internal pure virtual override returns (uint256) {
        return 50000;
    }
}

contract L1WethGatewayTester is L1WethGateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL2(
        address _user,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) =
            counterpartGateway.call{ value: _l2CallValue }(_data);
        require(success, "OUTBOUND_REVERT");
        return 1337;
    }
}

contract L2WethGatewayTester is L2WethGateway {
    function isSenderCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function sendTxToL1(
        address _from,
        uint256 _l1CallValue,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) =
            counterpartGateway.call{ value: _l1CallValue }(_data);
        require(success, "OUTBOUND_REVERT");
        return 1337;
    }

    function arbgasReserveIfCallRevert() internal pure virtual override returns (uint256) {
        return 50000;
    }
}
