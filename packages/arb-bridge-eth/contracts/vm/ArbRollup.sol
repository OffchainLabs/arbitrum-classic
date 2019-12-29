/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./Leaves.sol";

import "./VM.sol";
import "./IArbRollup.sol";

import "../IGlobalPendingInbox.sol";

import "../arch/Value.sol";


contract ArbRollup is Leaves, IArbRollup {

    // invalid leaf
    string constant MAKE_LEAF = "MAKE_LEAF";
    // Can only disputable assert if machine is not errored or halted
    string constant MAKE_RUN = "MAKE_RUN";
    // Tried to execute too many steps
    string constant MAKE_STEP = "MAKE_STEP";
    // Precondition: not within time bounds
    string constant MAKE_TIME = "MAKE_TIME";
    // invalid prev leaf proof
    string constant MAKE_PREV_PROOF = "MAKE_PREV_PROOF";
    // invalid staker location proof
    string constant MAKE_STAKER_PROOF = "MAKE_STAKER_PROOF";

    // must include proof for all stakers
    string constant CONF_COUNT = "CONF_COUNT";
    // Stakers must be ordered
    string constant CONF_ORDER = "CONF_ORDER";
    // at least one active staker disagrees
    string constant CONF_STAKER_PROOF = "CONF_STAKER_PROOF";
    // Type is not invalid
    string constant CONF_INV_TYPE = "CONF_INV_TYPE";

    // Only callable by owner
    string constant ONLY_OWNER = "ONLY_OWNER";

    using SafeMath for uint256;

    IGlobalPendingInbox public globalInbox;

    address   owner;
    VM.Params vmParams;


    // Fields
    //   prevLeafHash
    //   afterPendingTop
    //   importedMesssagesSlice
    //   afterVMHash
    //   messagesAccHash
    //   logsAccHash

    event RollupAsserted(
        bytes32[6] fields,
        uint32 _importedMessageCount,
        uint64[2] _timeBounds,
        uint32 _numSteps,
        uint64 _numArbGas
    );

    event RollupConfirmed(bytes32 nodeHash);

    event ConfirmedAssertion(
        bytes32 logsAccHash
    );

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforeInboxHash;
        bytes32 beforePendingTop;
        bytes32 prevPrevLeafHash;
        bytes32 prevExtraDataHash;
        bytes32[] stakerProof;
        bytes32 afterPendingTop;
        bytes32 importedMessagesSlice;
        uint32 importedMessageCount;
        bytes32 afterVMHash;
        bytes32 afterInboxHash;
        bytes32 messagesAccHash;
        bytes32 logsAccHash;
        uint32 numSteps;
        uint64 numArbGas;
        uint64[2] timeBounds;
    }

    function init(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _stakeRequirement,
        address payable _owner,
        address _challengeFactoryAddress,
        address _globalInboxAddress
    )
        external
    {
        Leaves.init(_vmState, _stakeRequirement, _challengeFactoryAddress);

        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        globalInbox.registerForInbox();
        owner = _owner;

        // VM parameters
        vmParams.gracePeriod = _gracePeriod;
        vmParams.maxExecutionSteps = _maxExecutionSteps;
        vmParams.pendingInboxHash = Value.hashEmptyTuple();
    }

    // fields
    //  beforeVMHash
    //  beforeInboxHash
    //  beforePendingTop
    //  prevPrevLeafHash
    //  prevExtraDataHash
    //  afterPendingTop
    //  importedMessagesSlice
    //  afterVMHash
    //  afterInboxHash
    //  messagesAccHash
    //  logsAccHash

    function makeAssertion(
        bytes32[11] calldata _fields,
        bytes32[] calldata _stakerProof,
        uint32 _importedMessageCount,
        uint32 _numSteps,
        uint64 _numArbGas,
        uint64[2] calldata _timeBounds
    )
        external
    {
        return _makeAssertion(
            MakeAssertionData(
                _fields[0],
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _stakerProof,
                _fields[5],
                _fields[6],
                _importedMessageCount,
                _fields[7],
                _fields[8],
                _fields[9],
                _fields[10],
                _numSteps,
                _numArbGas,
                _timeBounds
            )
        );
    }

    function confirmValid(
        uint deadline,
        bytes calldata _messages,
        bytes32 logsAcc,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint[]  calldata stakerProofOffsets
    )
        external
    {
        _confirmNode(
            deadline,
            RollupUtils.validNodeHash(
                Protocol.generateLastMessageHash(_messages),
                logsAcc
            ),
            VALID_CHILD_TYPE,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );

        globalInbox.sendMessages(_messages);

        emit ConfirmedAssertion(
            logsAcc
        );
    }

    function confirmInvalid(
        uint    deadline,
        bytes32 challengeNodeData,
        uint    branch,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint[]  calldata stakerProofOffsets
    )
        external
    {
        require(branch < VALID_CHILD_TYPE, CONF_INV_TYPE);
        _confirmNode(
            deadline,
            challengeNodeData,
            branch,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );
    }

    modifier onlyOwner() {
        require(msg.sender == owner, ONLY_OWNER);
        _;
    }

/*    function activateVM() external onlyOwner {
        if (vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function ownerShutdown() external onlyOwner {
        _shutdown();
    }
    */

    function _makeAssertion(MakeAssertionData memory data) private {
        bytes32 vmProtoHashBefore = RollupUtils.protoStateHash(
            data.beforeVMHash,
            data.beforeInboxHash,
            data.beforePendingTop
        );
        bytes32 prevLeaf = RollupUtils.childNodeHash(
            data.prevPrevLeafHash,
            vmProtoHashBefore,
            data.prevExtraDataHash

        );
        require(isValidLeaf(prevLeaf), MAKE_LEAF);
        require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
        require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
        require(withinTimeBounds(data.timeBounds), MAKE_TIME);

        Staker storage staker = getValidStaker(msg.sender);
        require(RollupUtils.isPath(staker.location, prevLeaf, data.stakerProof), MAKE_STAKER_PROOF);

        uint deadline = block.number + vmParams.gracePeriod; //TODO: [Ed] compute this properly
        bytes32 afterInboxHash = Protocol.addMessagesToInbox(data.beforeInboxHash, data.importedMessagesSlice);
        bytes32[] memory leaves = new bytes32[](MAX_CHILD_TYPE);
        leaves[INVALID_PENDING_TOP_CHILD_TYPE] = RollupUtils.childNodeHash(
            prevLeaf,
            deadline,
            ChallengeUtils.pendingTopHash(
                globalInbox.getPendingMessages(),
                data.afterPendingTop,
                data.importedMessageCount
            ),
            INVALID_PENDING_TOP_CHILD_TYPE,
            vmProtoHashBefore
        );
        leaves[INVALID_MESSAGES_CHILD_TYPE] = RollupUtils.childNodeHash(
            prevLeaf,
            deadline,
            ChallengeUtils.messagesHash(
                data.beforePendingTop,
                data.afterPendingTop,
                0x00,
                data.importedMessagesSlice,
                data.importedMessageCount
            ),
            INVALID_MESSAGES_CHILD_TYPE,
            vmProtoHashBefore
        );
        bytes32 assertionHash = Protocol.generateAssertionHash(
            data.afterVMHash,
            data.numSteps,
            data.numArbGas,
            0x00,
            data.messagesAccHash,
            0x00,
            data.logsAccHash
        );
        leaves[INVALID_EXECUTION_CHILD_TYPE] = RollupUtils.childNodeHash(
            prevLeaf,
            deadline,
            ChallengeUtils.executionHash(
                keccak256(
                    abi.encodePacked(
                        data.timeBounds[0],
                        data.timeBounds[1],
                        afterInboxHash
                    )
                ),
                data.beforeVMHash,
                assertionHash
            ),
            INVALID_EXECUTION_CHILD_TYPE,
            vmProtoHashBefore
        );
        leaves[VALID_CHILD_TYPE] = RollupUtils.childNodeHash(
            prevLeaf,
            deadline,
            RollupUtils.validNodeHash(
                data.messagesAccHash,
                data.logsAccHash
            ),
            VALID_CHILD_TYPE,
            RollupUtils.protoStateHash(
                data.afterVMHash,
                afterInboxHash,
                data.afterPendingTop
            )
        );
        splitLeaf(prevLeaf, leaves);
        staker.location = leaves[VALID_CHILD_TYPE];

        emit RollupAsserted(
            [
                prevLeaf,
                data.afterPendingTop,
                data.importedMessagesSlice,
                data.afterVMHash,
                data.messagesAccHash,
                data.logsAccHash
            ],
            data.importedMessageCount,
            data.timeBounds,
            data.numSteps,
            data.numArbGas
        );
    }

    function _confirmNode(
        uint deadline,
        bytes32 nodeDataHash,
        uint branch,
        bytes32 vmProtoStateHash,
        address[] memory stakerAddresses,
        bytes32[] memory stakerProofs,
        uint[]  memory stakerProofOffsets
    ) private {
        uint _stakerCount = stakerAddresses.length;
        require(_stakerCount == getStakerCount(), CONF_COUNT);
        bytes32 to = RollupUtils.childNodeHash(
            latestConfirmed(),
            deadline,
            nodeDataHash,
            branch,
            vmProtoStateHash
        );
        bytes20 prevStaker = 0x00;
        for (uint i = 0; i < _stakerCount; i++) {
            address stakerAddress = stakerAddresses[i];
            require(bytes20(stakerAddress) > prevStaker, CONF_ORDER);
            Staker storage staker = getValidStaker(stakerAddress);
            if (staker.creationTime >= deadline) {
                require(
                    RollupUtils.isPathOffset(
                        to,
                        staker.location,
                        stakerProofs,
                        stakerProofOffsets[i],
                        stakerProofOffsets[i+1]
                    ),
                    CONF_STAKER_PROOF
                );
            }
            prevStaker = bytes20(stakerAddress);
        }

        updateLatestConfirmed(to);

        emit RollupConfirmed(to);
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) private view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }
}
