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

import "../ethereum/gateway/L1GatewayRouter.sol";
import "../ethereum/gateway/L1WethGateway.sol";
import "../ethereum/gateway/L1CustomGateway.sol";
import "../ethereum/gateway/L1ERC20Gateway.sol";
import "../ethereum/L1ArbitrumMessenger.sol";

import "../arbitrum/gateway/L2GatewayRouter.sol";
import "../arbitrum/gateway/L2WethGateway.sol";
import "../arbitrum/gateway/L2CustomGateway.sol";
import "../arbitrum/gateway/L2ERC20Gateway.sol";
import "../arbitrum/L2ArbitrumMessenger.sol";
import "arb-bridge-eth/contracts/libraries/AddressAliasHelper.sol";

contract AddressMappingTest is L2ArbitrumMessenger {
    function getL1AddressTest(address sender) external pure returns (address l1Address) {
        return AddressAliasHelper.undoL1ToL2Alias(sender);
    }
}

// these contracts are used to "flatten" out communication between contracts
// this way the token bridge can be tested fully in the base layer
// assembly code from OZ's proxy is used to surface revert messages correctly
abstract contract L1ArbitrumTestMessenger is L1ArbitrumMessenger {
    bool shouldUseInbox;

    function setInboxUse(bool _shouldUseInbox) public {
        shouldUseInbox = _shouldUseInbox;
    }

    function sendTxToL2(
        address, /* _inbox */
        address _to,
        address, /* _user */
        uint256, /* _l1CallValue */
        uint256 _l2CallValue,
        uint256, /* _maxSubmissionCost */
        uint256, /* _maxGas */
        uint256, /* _gasPriceBid */
        bytes memory _data
    ) internal virtual override returns (uint256) {
        (bool success, bytes memory retdata) = _to.call{ value: _l2CallValue }(_data);
        assembly {
            switch success
            case 0 {
                revert(add(retdata, 32), mload(retdata))
            }
        }
        return 1337;
    }

    function getBridge(address _inboxMock) internal view virtual override returns (IBridge) {
        if (shouldUseInbox) {
            // the inbox mock covers the role of bridge/inbox/outbox
            return IBridge(_inboxMock);
        } else {
            return IBridge(msg.sender);
        }
    }

    function getL2ToL1Sender(address _inbox) internal view virtual override returns (address) {
        if (shouldUseInbox) {
            return super.getL2ToL1Sender(_inbox);
        } else {
            return msg.sender;
        }
    }
}

abstract contract L2ArbitrumTestMessenger is L2ArbitrumMessenger {
    struct PendingCall {
        address _to;
        uint256 _l1CallValue;
        bytes _data;
    }
    PendingCall[] pending;

    function sendTxToL1(
        uint256 _l1CallValue,
        address, /* _from */
        address _to,
        bytes memory _data
    ) internal virtual override returns (uint256) {
        pending.push(PendingCall({ _to: _to, _l1CallValue: _l1CallValue, _data: _data }));
    }

    function triggerTxToL1() external {
        PendingCall storage currCall = pending[pending.length - 1];

        address _to = currCall._to;
        uint256 _l1CallValue = currCall._l1CallValue;
        bytes memory _data = currCall._data;

        pending.pop();

        (bool success, bytes memory retdata) = _to.call{ value: _l1CallValue }(_data);
        assembly {
            switch success
            case 0 {
                revert(add(retdata, 32), mload(retdata))
            }
        }
    }
}

contract L1GatewayTester is L1ArbitrumTestMessenger, L1ERC20Gateway {
    function sendTxToL2(
        address _inbox,
        address _to,
        address _user,
        uint256 _l1CallValue,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override(L1ArbitrumMessenger, L1ArbitrumTestMessenger) returns (uint256) {
        return
            L1ArbitrumTestMessenger.sendTxToL2(
                _inbox,
                _to,
                _user,
                _l1CallValue,
                _l2CallValue,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                _data
            );
    }

    function getL2ToL1Sender(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (address)
    {
        return L1ArbitrumTestMessenger.getL2ToL1Sender(_inbox);
    }

    function getBridge(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (IBridge)
    {
        return L1ArbitrumTestMessenger.getBridge(_inbox);
    }
}

contract L2GatewayTester is L2ArbitrumTestMessenger, L2ERC20Gateway {
    function sendTxToL1(
        uint256 _l1CallValue,
        address _from,
        address _to,
        bytes memory _data
    ) internal virtual override(L2ArbitrumMessenger, L2ArbitrumTestMessenger) returns (uint256) {
        return L2ArbitrumTestMessenger.sendTxToL1(_l1CallValue, _from, _to, _data);
    }

    address public stubAddressOracleReturn;

    function setStubAddressOracleReturn(address _stubValue) external {
        stubAddressOracleReturn = _stubValue;
    }

    function calculateL2TokenAddress(address l1ERC20)
        public
        view
        virtual
        override
        returns (address)
    {
        // only return stub address if it is set
        // we use this to test the _withdraws initiated by the bridge
        // in case something goes wrong
        if (stubAddressOracleReturn != address(0)) {
            return stubAddressOracleReturn;
        }
        return super.calculateL2TokenAddress(l1ERC20);
    }
}

contract L1CustomGatewayTester is L1ArbitrumTestMessenger, L1CustomGateway {
    function sendTxToL2(
        address _inbox,
        address _to,
        address _user,
        uint256 _l1CallValue,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override(L1ArbitrumMessenger, L1ArbitrumTestMessenger) returns (uint256) {
        return
            L1ArbitrumTestMessenger.sendTxToL2(
                _inbox,
                _to,
                _user,
                _l1CallValue,
                _l2CallValue,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                _data
            );
    }

    function getL2ToL1Sender(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (address)
    {
        return L1ArbitrumTestMessenger.getL2ToL1Sender(_inbox);
    }

    function getBridge(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (IBridge)
    {
        return L1ArbitrumTestMessenger.getBridge(_inbox);
    }
}

contract L2CustomGatewayTester is L2ArbitrumTestMessenger, L2CustomGateway {
    function sendTxToL1(
        uint256 _l1CallValue,
        address _from,
        address _to,
        bytes memory _data
    ) internal virtual override(L2ArbitrumMessenger, L2ArbitrumTestMessenger) returns (uint256) {
        return L2ArbitrumTestMessenger.sendTxToL1(_l1CallValue, _from, _to, _data);
    }
}

contract L1WethGatewayTester is L1ArbitrumTestMessenger, L1WethGateway {
    function sendTxToL2(
        address _inbox,
        address _to,
        address _user,
        uint256 _l1CallValue,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override(L1ArbitrumMessenger, L1ArbitrumTestMessenger) returns (uint256) {
        return
            L1ArbitrumTestMessenger.sendTxToL2(
                _inbox,
                _to,
                _user,
                _l1CallValue,
                _l2CallValue,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                _data
            );
    }

    function getL2ToL1Sender(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (address)
    {
        return L1ArbitrumTestMessenger.getL2ToL1Sender(_inbox);
    }

    function getBridge(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (IBridge)
    {
        return L1ArbitrumTestMessenger.getBridge(_inbox);
    }
}

contract L2WethGatewayTester is L2ArbitrumTestMessenger, L2WethGateway {
    function sendTxToL1(
        uint256 _l1CallValue,
        address _from,
        address _to,
        bytes memory _data
    ) internal virtual override(L2ArbitrumMessenger, L2ArbitrumTestMessenger) returns (uint256) {
        return L2ArbitrumTestMessenger.sendTxToL1(_l1CallValue, _from, _to, _data);
    }

    function setL2WethAddress(address _l2Weth) external {
        L2WethGateway.l2Weth = _l2Weth;
    }
}

contract L1GatewayRouterTester is L1ArbitrumTestMessenger, L1GatewayRouter {
    function sendTxToL2(
        address _inbox,
        address _to,
        address _user,
        uint256 _l1CallValue,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override(L1ArbitrumMessenger, L1ArbitrumTestMessenger) returns (uint256) {
        return
            L1ArbitrumTestMessenger.sendTxToL2(
                _inbox,
                _to,
                _user,
                _l1CallValue,
                _l2CallValue,
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                _data
            );
    }

    function getL2ToL1Sender(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (address)
    {
        return L1ArbitrumTestMessenger.getL2ToL1Sender(_inbox);
    }

    function getBridge(address _inbox)
        internal
        view
        virtual
        override(L1ArbitrumMessenger, L1ArbitrumTestMessenger)
        returns (IBridge)
    {
        return L1ArbitrumTestMessenger.getBridge(_inbox);
    }
}

contract L2GatewayRouterTester is L2ArbitrumTestMessenger, L2GatewayRouter {
    function sendTxToL1(
        uint256 _l1CallValue,
        address _from,
        address _to,
        bytes memory _data
    ) internal virtual override(L2ArbitrumMessenger, L2ArbitrumTestMessenger) returns (uint256) {
        return L2ArbitrumTestMessenger.sendTxToL1(_l1CallValue, _from, _to, _data);
    }
}
