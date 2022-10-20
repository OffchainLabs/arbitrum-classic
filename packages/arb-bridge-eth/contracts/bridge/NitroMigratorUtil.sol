// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

pragma experimental ABIEncoderV2;

import "./interfaces/IInbox.sol";
import "@arbitrum/nitro-contracts/src/bridge/IBridge.sol" as INitroBridge;
import "@arbitrum/nitro-contracts/src/bridge/IInbox.sol" as INitroInbox;

/**
 * @notice DEPRECATED - only for classic version, see new repo (https://github.com/OffchainLabs/nitro/tree/master/contracts) 
 * for new updates
 */
interface INitroRollup {
    struct GlobalState {
        bytes32[2] bytes32Vals;
        uint64[2] u64Vals;
    }

    enum MachineStatus {
        RUNNING,
        FINISHED,
        ERRORED,
        TOO_FAR
    }

    struct ExecutionState {
        GlobalState globalState;
        MachineStatus machineStatus;
    }
    struct NitroRollupAssertion {
        ExecutionState beforeState;
        ExecutionState afterState;
        uint64 numBlocks;
    }

    function bridge() external view returns (INitroBridge.IBridge);

    function inbox() external view returns (INitroInbox.IInbox);

    function setInbox(IInbox newInbox) external;

    function setOwner(address newOwner) external;

    function paused() external view returns (bool);

    function pause() external;

    function resume() external;

    function latestNodeCreated() external returns (uint64);

    function createNitroMigrationGenesis(NitroRollupAssertion calldata assertion) external;
}

interface IArbOwner {
    function addChainOwner(address newOwner) external;
}

/// @dev lib used since file level consts aren't available in this solc version
library NitroReadyMagicNums {
    uint256 constant ROLLUP_USER = 0xa4b1;
    uint256 constant ROLLUP_ADMIN = 0xa4b2;
    uint256 constant NODE_BEACON = 0xa4b3;
    uint256 constant OUTBOX = 0xa4b4;
    uint256 constant BRIDGE = 0xa4b5;
    uint256 constant DELAYED_INBOX = 0xa4b6;
    uint256 constant SEQ_INBOX = 0xa4b7;
}
