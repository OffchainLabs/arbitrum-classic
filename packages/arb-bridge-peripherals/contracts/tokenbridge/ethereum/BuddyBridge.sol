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

contract BuddyBridge {
    IInbox public inbox;
    L2Deployer public l2Deployer;

    constructor(address _inbox, address _l2Deployer) public {
        inbox = IInbox(_inbox);
        l2Deployer = L2Deployer(_l2Deployer);
    }

    function initiateBuddyDeploy(
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata deployCode
    ) external payable {
        // TODO: check if called by contract
        // TODO: check if contract adheres to interface?
        address user = msg.sender;

        // deployCode == type(ArbSymmetricTokenBridge).creationCode
        bytes memory data = abi.encodeWithSelector(L2Deployer.executeBuddyDeploy.selector, user, deployCode);

        if(msg.value > 0) {
            // gas paid in L1
            inbox.sendL1FundedContractTransaction(maxGas, gasPriceBid, address(l2Deployer), data);
        } else {
            // gas paid in L2
            inbox.sendContractTransaction(maxGas, gasPriceBid, address(l2Deployer), 0, data);
        }
    }
}

// contracts that want to have buddies should inherit from this
abstract contract BuddyContract {
    bool public connected;
    bytes32 public codeHash;
    BuddyBridge public buddyBridge;

    constructor(
        address _buddyBridge,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes memory l2ContractCode
    ) public payable {
        connected = false;
        buddyBridge = BuddyBridge(_buddyBridge);
        // l2ContractCode == type(ArbSymmetricTokenBridge).creationCode
        buddyBridge.initiateBuddyDeploy(maxGas, gasPriceBid, l2ContractCode);
        codeHash = keccak256(l2ContractCode);
    }

    function finalizeBuddyDeploy(
        bool success
    ) external {
        // get sender from outbox
        IInbox inbox = buddyBridge.inbox();
        IOutbox outbox = IOutbox(inbox.bridge().activeOutbox());

        // this logic only needed if call comes from deployed contract
        /*
        address expectedL2Address = calculateL2Address(
            address(buddyBridge.l2Deployer()),
            address(this),
            codeHash
        );
        require(outbox.l2ToL1Sender() == expectedL2Address, "Wrong L2 address triggering outbox");
        */

        // instead the L2 bridge can tell us if the deploy worked
        require(outbox.l2ToL1Sender() == address(buddyBridge.l2Deployer()), "Only L2 buddy bridge");

        if(success) {
            handleDeploySuccess();
        } else {
            handleDeployFail();
        }
    }
    
    function handleDeploySuccess() internal virtual { connected = true; }
    function handleDeployFail() internal virtual;

    function calculateL2Address(
        address _deployer,
        address _salt,
        bytes32 _codeHash
    )
        internal
        pure
        returns (address)
    {
        // bytes32 temp = keccak256( 0xff ++ address ++ salt ++ keccak256(init_code))[12:]
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                _deployer,
                _salt,
                _codeHash
            )
        );
        return address(uint160(uint256(hash)));
    }
}
