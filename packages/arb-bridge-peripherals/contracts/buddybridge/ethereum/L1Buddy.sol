// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";

import "../arbitrum/BuddyDeployer.sol";
import "../util/BuddyUtil.sol";

// contracts that want to have buddies should inherit from this
abstract contract L1Buddy {
    enum L2Connection {
        Null, // 0
        Initiated, // 1
        Complete // 2
    }

    L2Connection public l2Connection;
    BuddyDeployer public l2Deployer;
    IInbox public inbox;
    address public l2Buddy;
    bytes32 public codeHash;

    modifier onlyIfConnected {
        require(l2Connection == L2Connection.Complete, "Not connected");
        _;
    }

    modifier onlyL2Buddy {
        require(l2Buddy != address(0), "l2 buddy not set");
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        require(l2Buddy == outbox.l2ToL1Sender(), "Not from l2 buddy");
        _;
    }

    constructor(address _inbox, address _l2Deployer) public {
        l2Connection = L2Connection.Null;
        inbox = IInbox(_inbox);
        l2Deployer = BuddyDeployer(_l2Deployer);
    }

    function initiateBuddyDeploy(
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes memory contractInitCode
    ) public payable {
        require(l2Connection != L2Connection.Complete, "already connected");
        require(
            codeHash == bytes32(0) || codeHash == keccak256(contractInitCode),
            "Only retry if same deploy code"
        );
        bytes memory data = abi.encodeWithSelector(BuddyDeployer.executeBuddyDeploy.selector, contractInitCode);

        if(msg.value > 0) {
            // gas paid in L1
            inbox.sendL1FundedContractTransaction{value: msg.value}(maxGas, gasPriceBid, address(l2Deployer), data);
        } else {
            // gas paid in L2
            inbox.sendContractTransaction(maxGas, gasPriceBid, address(l2Deployer), 0, data);
        }
        codeHash = keccak256(contractInitCode);
        l2Buddy = BuddyUtil.calculateL2Address(
            address(l2Deployer),
            address(this),
            codeHash
        );
        l2Connection = L2Connection.Initiated;
    }

    function finalizeBuddyDeploy(
        bool success
    ) external {
        require(l2Connection == L2Connection.Initiated, "Connection not in initiated state");
        // get sender from outbox
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());
        require(outbox.l2ToL1Sender() == address(l2Deployer), "Wrong L2 address triggering outbox");
        /*
            The callback from L2 can come from buddy's constructor if
            you don't want to rely on the L2Deployer's correctness.
        */

        if(success) {
            handleDeploySuccess();
        } else {
            handleDeployFail();
        }
    }
    
    function handleDeploySuccess() internal virtual {
        l2Connection = L2Connection.Complete;
    }

    function handleDeployFail() internal virtual;
}
