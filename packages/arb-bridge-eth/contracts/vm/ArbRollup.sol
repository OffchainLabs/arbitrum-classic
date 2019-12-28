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
    // Previous leaf incorrectly unwrapped
    string constant MAKE_PREV = "MAKE_PREV";
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

    // invalid leaf
    string constant CONF_LEAF = "CONF_LEAF";
    // Invalid child type
    string constant CONF_TYPE = "CONF_TYPE";
    // must include proof for all stakers
    string constant CONF_COUNT = "CONF_COUNT";
    // node does not exist
    string constant CONF_LEAF_PROOF = "CONF_LEAF_PROOF";
    // Stakers must be ordered
    string constant CONF_ORDER = "CONF_ORDER";
    // at least one active staker disagrees
    string constant CONF_STAKER_PROOF = "CONF_STAKER_PROOF";

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

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforeInboxHash;
        bytes32 beforePendingTop;
        bytes32 prevPrevLeafHash;
        uint prevDeadline;
        bytes32 prevDisputableHash;
        uint prevChildType;
        bytes32 prevLeaf;
        bytes32[] prevLeafProof;
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

    struct ConfirmData {
        address[] stakerAddresses;
        bytes32[] stakerProofs;
        uint[] stakerProofOffsets;
        uint branch;
        uint deadline;
        bytes32 preconditionHash;
        bytes32 pendingAssertion;
        bytes32 importedAssertion;
        bytes32 executionAssertion;
        bytes32 vmProtoStateHash;
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
    //  prevDisputableHash
    //  afterPendingTop
    //  importedMessagesSlice
    //  afterVMHash
    //  afterInboxHash
    //  messagesAccHash
    //  logsAccHash

    function makeAssertion(
        bytes32[11] calldata _fields,
        uint _prevDeadline,
        uint    _prevChildType,
        bytes32 _prevLeaf,
        bytes32[] calldata _prevLeafProof,
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
                _prevDeadline,
                _fields[4],
                _prevChildType,
                _prevLeaf,
                _prevLeafProof,
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

    // fields
    //   preconditionHash
    //   pendingAssertion
    //   importedAssertion
    //   executionAssertion
    //   vmProtoStateHash

    function confirm(
        bytes32[5] calldata fields,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint[]  calldata stakerProofOffsets,
        uint    branch,
        uint    deadline
    )
        external
    {
        return _confirm(ConfirmData(
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets,
            branch,
            deadline,
            fields[0],
            fields[1],
            fields[2],
            fields[3],
            fields[4]
        ));
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
        Staker storage staker = getValidStaker(msg.sender);
        require(isValidLeaf(data.prevLeaf), MAKE_LEAF);
        bytes32 vmProtoHashBefore = RollupUtils.protoStateHash(
            data.beforeVMHash,
            data.beforeInboxHash,
            data.beforePendingTop
        );
        require(
            RollupUtils.childNodeHash(
                data.prevPrevLeafHash,
                data.prevDeadline,
                data.prevDisputableHash,
                data.prevChildType,
                vmProtoHashBefore
            ) == data.prevLeaf,
            MAKE_PREV
        );

        require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
        require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
        require(withinTimeBounds(data.timeBounds), MAKE_TIME);
        require(RollupUtils.isPath(latestConfirmed(), data.prevLeaf, data.prevLeafProof), MAKE_PREV_PROOF);
        require(RollupUtils.isPath(staker.location, data.prevLeaf, data.stakerProof), MAKE_STAKER_PROOF);

        uint deadline = block.number + vmParams.gracePeriod; //TODO: [Ed] compute this properly
        bytes32 assertionHash = Protocol.generateAssertionHash(
            data.afterVMHash,
            data.numSteps,
            data.numArbGas,
            0x00,
            data.messagesAccHash,
            0x00,
            data.logsAccHash
        );
        bytes32 disputableHash = RollupUtils.disputableNodeHash(
            Protocol.generatePreconditionHash(
                data.beforeVMHash,
                data.timeBounds,
                data.beforeInboxHash
            ),
            RollupUtils.pendingAssertionHash(
                data.afterPendingTop,
                vmParams.pendingInboxHash
            ),
            RollupUtils.importedAssertionHash(
                data.beforePendingTop,
                data.importedMessageCount,
                data.importedMessagesSlice
            ),
            assertionHash
        );

        bytes32 validKid = RollupUtils.childNodeHash(
            data.prevLeaf,
            deadline,
            disputableHash,
            VALID_CHILD_TYPE,
            RollupUtils.protoStateHash(
                data.afterVMHash,
                data.afterInboxHash,
                data.afterPendingTop
            )
        );

        bytes32[] memory leaves = new bytes32[](MAX_CHILD_TYPE);
        leaves[0] = validKid;
        for (uint i = 1; i<=MAX_CHILD_TYPE; i++) {
            leaves[i] = RollupUtils.childNodeHash(
                data.prevLeaf,
                deadline,
                disputableHash,
                i,
                vmProtoHashBefore
            );
        }
        splitLeaf(data.prevLeaf, leaves);
        staker.location = validKid;

        emit RollupAsserted(
            [
                data.prevLeaf,
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

    function _confirm(ConfirmData memory data) private {
        uint _stakerCount = data.stakerAddresses.length;
        require(_stakerCount == getStakerCount(), CONF_COUNT);
        bytes32 to = RollupUtils.childNodeHash(
            latestConfirmed(),
            data.deadline,
            RollupUtils.disputableNodeHash(
                data.preconditionHash,
                data.pendingAssertion,
                data.importedAssertion,
                data.executionAssertion
            ),
            data.branch,
            data.vmProtoStateHash
        );
        bytes20 prevStaker = 0x00;
        for (uint i = 0; i < _stakerCount; i++) {
            address stakerAddress = data.stakerAddresses[i];
            require(bytes20(stakerAddress) > prevStaker, CONF_ORDER);
            Staker storage staker = getValidStaker(stakerAddress);
            if (staker.creationTime >= data.deadline) {
                require(
                    RollupUtils.isPathOffset(
                        to,
                        staker.location,
                        data.stakerProofs,
                        data.stakerProofOffsets[i],
                        data.stakerProofOffsets[i+1]
                    ),
                    CONF_STAKER_PROOF
                );
            }
            prevStaker = bytes20(stakerAddress);
        }

        updateLatestConfirmed(to);
        if (data.branch == 0) {
            //TODO: execute actions from the DA before the confirmed assertion (to)
        }

        emit RollupConfirmed(to);
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) private view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }
}
