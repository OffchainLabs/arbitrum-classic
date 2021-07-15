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

import "./RollupCore.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "./RollupEventBridge.sol";

import "./INode.sol";
import "./INodeFactory.sol";
import "../challenge/IChallenge.sol";
import "../challenge/IChallengeFactory.sol";
import "../bridge/interfaces/IBridge.sol";
import "../bridge/interfaces/IOutbox.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../bridge/Messages.sol";
import "./RollupLib.sol";
import "../libraries/Cloneable.sol";
import "./facets/IRollupFacets.sol";

abstract contract RollupBase is Cloneable, RollupCore, Pausable {
    // Rollup Config
    uint256 public confirmPeriodBlocks;
    uint256 public extraChallengeTimeBlocks;
    uint256 public arbGasSpeedLimitPerBlock;
    uint256 public baseStake;

    // Bridge is an IInbox and IOutbox
    IBridge public delayedBridge;
    ISequencerInbox public sequencerBridge;
    IOutbox public outbox;
    RollupEventBridge public rollupEventBridge;
    IChallengeFactory public challengeFactory;
    INodeFactory public nodeFactory;
    address public owner;
    address public stakeToken;
    uint256 public minimumAssertionPeriod;

    uint256 public STORAGE_GAP_1;
    uint256 public STORAGE_GAP_2;
    uint256 public challengeExecutionBisectionDegree;

    address[] internal facets;

    mapping(address => bool) isValidator;

    event RollupCreated(bytes32 machineHash);

    event NodeCreated(
        uint256 indexed nodeNum,
        bytes32 indexed parentNodeHash,
        bytes32 nodeHash,
        bytes32 executionHash,
        uint256 inboxMaxCount,
        uint256 afterInboxBatchEndCount,
        bytes32 afterInboxBatchAcc,
        bytes32[3][2] assertionBytes32Fields,
        uint256[4][2] assertionIntFields
    );

    event NodeConfirmed(
        uint256 indexed nodeNum,
        bytes32 afterSendAcc,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount
    );

    event NodeRejected(uint256 indexed nodeNum);

    event RollupChallengeStarted(
        address indexed challengeContract,
        address asserter,
        address challenger,
        uint256 challengedNode
    );

    event StakerReassigned(address indexed staker, uint256 newNode);
    event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode);
    event OwnerFunctionCalled(uint256 indexed id);
}

contract Rollup is RollupBase {
    constructor() public Cloneable() Pausable() {
        // constructor is used so logic contract can't be init'ed
        confirmPeriodBlocks = 1;
        require(isInit(), "CONSTRUCTOR_NOT_INIT");
    }

    function isInit() internal view returns (bool) {
        return confirmPeriodBlocks != 0;
    }

    // _rollupParams = [ confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake ]
    // connectedContracts = [delayedBridge, sequencerInbox, outbox, rollupEventBridge, challengeFactory, nodeFactory]
    function initialize(
        bytes32 _machineHash,
        uint256[4] calldata _rollupParams,
        address _stakeToken,
        address _owner,
        bytes calldata _extraConfig,
        address[6] calldata connectedContracts,
        address[2] calldata _facets,
        uint256[2] calldata sequencerInboxParams
    ) public {
        require(!isInit(), "ALREADY_INIT");
        require(_rollupParams[0] != 0, "BAD_CONF_PERIOD");

        delayedBridge = IBridge(connectedContracts[0]);
        sequencerBridge = ISequencerInbox(connectedContracts[1]);
        outbox = IOutbox(connectedContracts[2]);
        delayedBridge.setOutbox(connectedContracts[2], true);
        rollupEventBridge = RollupEventBridge(connectedContracts[3]);
        delayedBridge.setInbox(connectedContracts[3], true);

        rollupEventBridge.rollupInitialized(
            _rollupParams[0],
            _rollupParams[2],
            _rollupParams[3],
            _stakeToken,
            _owner,
            _extraConfig
        );

        challengeFactory = IChallengeFactory(connectedContracts[4]);
        nodeFactory = INodeFactory(connectedContracts[5]);

        INode node = createInitialNode(_machineHash);
        initializeCore(node);

        confirmPeriodBlocks = _rollupParams[0];
        extraChallengeTimeBlocks = _rollupParams[1];
        arbGasSpeedLimitPerBlock = _rollupParams[2];
        baseStake = _rollupParams[3];
        owner = _owner;
        // A little over 15 minutes
        minimumAssertionPeriod = 75;
        challengeExecutionBisectionDegree = 400;

        sequencerBridge.setMaxDelayBlocks(sequencerInboxParams[0]);
        sequencerBridge.setMaxDelaySeconds(sequencerInboxParams[1]);

        // facets[0] == admin, facets[1] == user
        facets = _facets;

        (bool success, ) =
            _facets[1].delegatecall(
                abi.encodeWithSelector(IRollupUser.initialize.selector, _stakeToken)
            );
        require(success, "FAIL_INIT_FACET");

        emit RollupCreated(_machineHash);
        require(isInit(), "INITIALIZE_NOT_INIT");
    }

    function postUpgradeInit(address newAdminFacet) external {
        // this upgrade moves the delay blocks and seconds tracking to the sequencer inbox
        // because of that we need to update the admin facet logic to allow the owner to set
        // these values in the sequencer inbox

        uint256 oldDelayBlocks = STORAGE_GAP_1;
        uint256 oldDelaySeconds = STORAGE_GAP_2;

        require(oldDelayBlocks == 6545, "ALREADY_POST_INIT");
        require(oldDelaySeconds == 86400, "ALREADY_POST_INIT");

        sequencerBridge.setMaxDelayBlocks(oldDelayBlocks);
        sequencerBridge.setMaxDelaySeconds(oldDelaySeconds);

        STORAGE_GAP_1 = 0;
        STORAGE_GAP_2 = 0;

        facets[0] = newAdminFacet;
    }

    function createInitialNode(bytes32 _machineHash) private returns (INode) {
        bytes32 state =
            RollupLib.stateHash(
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
     * Fallback and delegate functions from OZ
     * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/proxy/TransparentUpgradeableProxy.sol
     * And dispatch pattern from EIP-2535: Diamonds
     */

    function getFacets() public view returns (address, address) {
        return (getAdminFacet(), getUserFacet());
    }

    function getAdminFacet() public view returns (address) {
        return facets[0];
    }

    function getUserFacet() public view returns (address) {
        return facets[1];
    }

    /**
     * @dev Fallback function that delegates calls to the address returned by `_implementation()`. Will run if no other
     * function in the contract matches the call data.
     */
    fallback() external payable {
        _fallback();
    }

    /**
     * @dev Fallback function that delegates calls to the address returned by `_implementation()`. Will run if call data
     * is empty.
     */
    receive() external payable {
        _fallback();
    }

    /**
     * @dev Delegates the current call to the address returned by `_implementation()`.
     *
     * This function does not return to its internall call site, it will return directly to the external caller.
     */
    function _fallback() internal virtual {
        require(msg.data.length >= 4, "NO_FUNC_SIG");
        address rollupOwner = owner;
        // if there is an owner and it is the sender, delegate to admin facet
        address target =
            rollupOwner != address(0) && rollupOwner == msg.sender
                ? getAdminFacet()
                : getUserFacet();
        _delegate(target);
    }

    /**
     * @dev Delegates the current call to `implementation`.
     *
     * This function does not return to its internall call site, it will return directly to the external caller.
     */
    function _delegate(address implementation) internal virtual {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code. We overwrite the
            // Solidity scratch pad at memory position 0.
            calldatacopy(0, 0, calldatasize())

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(gas(), implementation, 0, calldatasize(), 0, 0)

            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
                // delegatecall returns 0 on error.
                case 0 {
                    revert(0, returndatasize())
                }
                default {
                    return(0, returndatasize())
                }
        }
    }
}
