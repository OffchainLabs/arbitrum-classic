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

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arbos-contracts/arbos/builtin/ArbSys.sol";

abstract contract L1ArbitrumMessenger {
    event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data);

    function sendTxToL2(
        address _inbox,
        address _to,
        address _user,
        uint256 _l2CallValue,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual returns (uint256) {
        uint256 seqNum =
            IInbox(_inbox).createRetryableTicket{ value: msg.value }(
                _to,
                _l2CallValue,
                _maxSubmissionCost,
                _user,
                _user,
                _maxGas,
                _gasPriceBid,
                _data
            );
        emit TxToL2(_user, _to, seqNum, _data);
        return seqNum;
    }
}

abstract contract L2ArbitrumMessenger {
    address internal constant arbsysAddr = address(100);

    event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data);

    function sendTxToL1(
        uint256 _l1CallValue,
        address _from,
        address _to,
        bytes memory _data
    ) internal virtual returns (uint256) {
        uint256 _id = ArbSys(arbsysAddr).sendTxToL1{ value: _l1CallValue }(_to, _data);
        emit TxToL1(_from, _to, _id, _data);
        return _id;
    }
}
