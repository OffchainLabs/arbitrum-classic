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

import "@openzeppelin/contracts/proxy/Proxy.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "./RollupCore.sol";
import "./RollupEventBridge.sol";
import "./RollupLib.sol";
import "./INode.sol";
import "./INodeFactory.sol";

import "../libraries/ProxyUtil.sol";

contract Rollup is Proxy, RollupCore {
    using Address for address;

    constructor(uint256 _confirmPeriodBlocks) public Cloneable() Pausable() {
        // constructor is used so logic contract can't be init'ed
        confirmPeriodBlocks = _confirmPeriodBlocks;
        require(isInit(), "CONSTRUCTOR_NOT_INIT");
    }

    function isInit() internal view returns (bool) {
        return confirmPeriodBlocks != 0;
    }

    // _rollupParams = [ confirmPeriodBlocks, extraChallengeTimeBlocks, avmGasSpeedLimitPerBlock, baseStake ]
    // connectedContracts = [delayedBridge, sequencerInbox, outbox, rollupEventBridge, challengeFactory, nodeFactory]
    function initialize(
        bytes32 _machineHash,
        uint256[4] calldata _rollupParams,
        address _stakeToken,
        address _owner,
        bytes calldata _extraConfig,
        address[6] calldata connectedContracts,
        address[2] calldata _logicContracts,
        uint256[2] calldata sequencerInboxParams
    ) public {
        require(!isInit(), "ALREADY_INIT");

        // calls initialize method in user logic
        require(_logicContracts[0].isContract(), "LOGIC_0_NOT_CONTRACT");
        require(_logicContracts[1].isContract(), "LOGIC_1_NOT_CONTRACT");
        (bool success, ) = _logicContracts[1].delegatecall(
            abi.encodeWithSelector(IRollupUser.initialize.selector, _stakeToken)
        );
        adminLogic = IRollupAdmin(_logicContracts[0]);
        userLogic = IRollupUser(_logicContracts[1]);
        require(success, "FAIL_INIT_LOGIC");

        delayedBridge = IBridge(connectedContracts[0]);
        sequencerBridge = ISequencerInbox(connectedContracts[1]);
        outbox = IOutbox(connectedContracts[2]);
        delayedBridge.setOutbox(connectedContracts[2], true);
        rollupEventBridge = RollupEventBridge(connectedContracts[3]);
        delayedBridge.setInbox(connectedContracts[3], true);

        rollupEventBridge.rollupInitialized(
            _rollupParams[0],
            _rollupParams[2],
            _owner,
            _extraConfig
        );

        challengeFactory = IChallengeFactory(connectedContracts[4]);
        nodeFactory = INodeFactory(connectedContracts[5]);

        INode node = createInitialNode(_machineHash);
        initializeCore(node);

        confirmPeriodBlocks = _rollupParams[0];
        extraChallengeTimeBlocks = _rollupParams[1];
        avmGasSpeedLimitPerBlock = _rollupParams[2];
        baseStake = _rollupParams[3];
        owner = _owner;
        // A little over 15 minutes
        minimumAssertionPeriod = 75;
        challengeExecutionBisectionDegree = 400;

        sequencerBridge.setMaxDelay(sequencerInboxParams[0], sequencerInboxParams[1]);

        emit RollupCreated(_machineHash);
        require(isInit(), "INITIALIZE_NOT_INIT");
    }

    function postUpgradeInit() external {
        // it is assumed the rollup contract is behind a Proxy controlled by a proxy admin
        // this function can only be called by the proxy admin contract
        address proxyAdmin = ProxyUtil.getProxyAdmin();
        require(msg.sender == proxyAdmin, "NOT_FROM_ADMIN");
    }

    function createInitialNode(bytes32 _machineHash) private returns (INode) {
        bytes32 state = RollupLib.stateHash(
            RollupLib.ExecutionState(
                0, // total gas used
                _machineHash,
                0, // inbox count
                0, // send count
                0, // log count
                0, // send acc
                0, // log acc
                block.number, // block proposed
                1 // Initialization message already in inbox
            )
        );
        return
            INode(
                nodeFactory.createNode(
                    state,
                    0, // challenge hash (not challengeable)
                    0, // confirm data
                    0, // prev node
                    block.number // deadline block (not challengeable)
                )
            );
    }

    /**
     * @dev This is a virtual function that should be overriden so it returns the address to which the fallback function
     * and {_fallback} should delegate.
     */
    function _implementation() internal view virtual override returns (address) {
        require(msg.data.length >= 4, "NO_FUNC_SIG");
        address rollupOwner = owner;
        // if there is an owner and it is the sender, delegate to admin logic
        address target = rollupOwner != address(0) && rollupOwner == msg.sender
            ? address(adminLogic)
            : address(userLogic);
        require(target.isContract(), "TARGET_NOT_CONTRACT");
        return target;
    }
}
