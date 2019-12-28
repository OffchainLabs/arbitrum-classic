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

import "./VM.sol";
import "./IArbRollup.sol";

import "../IGlobalPendingInbox.sol";

import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";
import "../arch/Value.sol";


contract ArbRollup is IArbRollup {
    using SafeMath for uint256;

    address internal constant ETH_ADDRESS = address(0);

    IGlobalPendingInbox public globalInbox;
    IChallengeFactory public challengeFactory;

    struct Staker {
        bytes32 location;
        uint    creationTime;
        bool    inChallenge;
    }

    address   owner;
    VM.Params vmParams;
    bytes32   latestConfirmed;
    mapping(address => Staker) stakers;
    uint stakerCount;
    mapping (bytes32 => bool) leaves;


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

    event RollupPruned(bytes32 nodeHash);

    event RollupStakeCreated(
        address staker,
        bytes32 nodeHash,
        uint    blockNumber
    );

    event RollupStakeMoved(
        address staker,
        bytes32 toNodeHash
    );

    event RollupStakeRefunded(address staker);

    event RollupChallengeStarted(
        address asserter,
        address challenger,
        uint    challengeType,
        address challengeContract
    );

    event RollupChallengeCompleted(
        address challengeContract,
        address winner,
        address loser
    );

    uint constant VALID_CHILD_TYPE = 0;
    uint constant INVALID_PENDING_TOP_CHILD_TYPE = 1;
    uint constant INVALID_MESSAGES_CHILD_TYPE = 2;
    uint constant INVALID_EXECUTION_CHILD_TYPE = 3;
    uint constant MAX_CHILD_TYPE = 3;

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforeInboxHash;
        bytes32 beforePendingTop;
        bytes32 prevPrevLeafHash;
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

    struct ChallengeData {
        address payable[2] stakerAddresses;
        bytes32 node;
        uint64 disputableDeadline;
        uint[2] stakerPositions;
        bytes32[2] vmProtoHashes;
        bytes32[] proof1;
        bytes32[] proof2;
    }

    struct StartExecutionChallengeData {
        bytes32 beforeHash;
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes32 pendingAssertion;
        bytes32 beforePendingTop;
        bytes32 importedMessageSlice;
        uint32 importedMessageCount;
        bytes32 assertionHash;
    }

    struct StartPendingTopChallengeData {
        bytes32 preconditionHash;
        bytes32 afterPendingTop;
        bytes32 currentPending;
        bytes32 importedAssertion;
        bytes32 assertionHash;
    }

    struct StartMessagesChallengeData {
        bytes32 preconditionHash;
        bytes32 afterPendingTop;
        bytes32 currentPending;
        bytes32 beforePendingTop;
        bytes32 importedMessageSlice;
        uint32 importedMessageCount;
        bytes32 assertionHash;
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
        require(latestConfirmed == bytes32(0), "VM already initialized");

        challengeFactory = IChallengeFactory(_challengeFactoryAddress);
        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        globalInbox.registerForInbox();
        owner = _owner;

        // VM parameters
        vmParams.stakeRequirement = _stakeRequirement;
        vmParams.gracePeriod = _gracePeriod;
        vmParams.maxExecutionSteps = _maxExecutionSteps;
        vmParams.pendingInboxHash = Value.hashEmptyTuple();

        // VM protocol state
        bytes32 vmProtoStateHash = VM.protoStateHash(_vmState, Value.hashEmptyTuple(), Value.hashEmptyTuple());
        latestConfirmed = childNodeHash(
            0,
            0,
            0,
            vmProtoStateHash
        );
        leaves[latestConfirmed] = true;
    }

    function getValidStaker(address _stakerAddress) internal view returns (Staker storage) {
        Staker storage staker = stakers[_stakerAddress];
        require(staker.location != 0x00, "Invalid staker");
        return staker;
    }

    function deleteStaker(address _stakerAddress) internal {
        delete stakers[_stakerAddress];
        stakerCount--;
    }

    function resolveChallenge(address payable winner, address loser) external {
        address sender = msg.sender;
        bytes32 codehash;
        assembly { codehash := extcodehash(sender) }
        address challengeContract = challengeFactory.generateCloneAddress(address(winner), loser, codehash);
        require(challengeContract == msg.sender, "Challenge can only be resolved by spawned contract");
        Staker storage winningStaker = getValidStaker(address(winner));
        winner.transfer(vmParams.stakeRequirement / 2);
        winningStaker.inChallenge = false;
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, address(winner), loser);
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
        bytes32[11] memory _fields,
        uint    _prevChildType,
        bytes32 _prevLeaf,
        bytes32[] memory _prevLeafProof,
        bytes32[] memory _stakerProof,
        uint32 _importedMessageCount,
        uint32 _numSteps,
        uint64 _numArbGas,
        uint64[2] memory _timeBounds
    )
        public
    {
        return _makeAssertion(
            MakeAssertionData(
                _fields[0],
                _fields[1],
                _fields[2],
                _fields[3],
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

    struct ConfirmData {
        bytes32 leaf;
        address[] stakerAddresses;
        bytes32[] proof1;
        bytes32[] stakerProofs;
        uint[] stakerProofOffsets;
        bytes32 prev;
        uint branch;
        uint deadline;
        bytes32 preconditionHash;
        bytes32 pendingAssertion;
        bytes32 importedAssertion;
        bytes32 executionAssertion;
        bytes32 vmProtoStateHash;
    }

    function confirm(
        bytes32    _leaf,
        address[] memory stakerAddresses,
        bytes32[] memory proof1,
        bytes32[] memory stakerProofs,
        uint[]  memory stakerProofOffsets,
        bytes32 prev,
        uint    branch,
        uint    deadline,
        bytes32 _preconditionHash,
        bytes32 _pendingAssertion,
        bytes32 _importedAssertion,
        bytes32 _executionAssertion,
        bytes32 _vmProtoStateHash
    )
        public
    {
        return _confirm(ConfirmData(
            _leaf,
            stakerAddresses,
            proof1,
            stakerProofs,
            stakerProofOffsets,
            prev,
            branch,
            deadline,
            _preconditionHash,
            _pendingAssertion,
            _importedAssertion,
            _executionAssertion,
            _vmProtoStateHash
        ));
    }

    function _confirm(ConfirmData memory data) internal {
        require(leaves[data.leaf], "invalid leaf");
        uint _stakerCount = data.stakerAddresses.length;
        require(_stakerCount == stakerCount, "must include proof for all stakers");
        bytes32 to = childNodeHash(
            data.prev,
            disputableNodeHash(
                data.deadline,
                data.preconditionHash,
                data.pendingAssertion,
                data.importedAssertion,
                data.executionAssertion
            ),
            data.branch,
            data.vmProtoStateHash
        );
        require(isPath(to, data.leaf, data.proof1), "node does not exist");
        bytes20 prevStaker = 0x00;
        for (uint i = 0; i < _stakerCount; i++) {
            address stakerAddress = data.stakerAddresses[i];
            require(bytes20(stakerAddress) > prevStaker, "Stakers must be ordered");
            Staker storage staker = getValidStaker(stakerAddress);
            if (staker.creationTime >= data.deadline) {
                require(
                    isPathOffset(
                        to,
                        staker.location,
                        data.stakerProofs,
                        data.stakerProofOffsets[i],
                        data.stakerProofOffsets[i+1]
                    ),
                    "at least one active staker disagrees"
                );
            }
            prevStaker = bytes20(stakerAddress);
        }

        latestConfirmed = to;
        if (data.branch == 0) {
            //TODO: execute actions from the DA before the confirmed assertion (to)
        }

        emit RollupConfirmed(to);
    }

    function pruneLeaf(
        bytes32 _leaf,
        bytes32 from,
        bytes32[] memory leafProof,
        bytes32[] memory latestConfirmedProof
    )
        public
    {
        require(leaves[_leaf], "invalid leaf");
        require(
            isConflict(
                from,
                _leaf,
                latestConfirmed,
                leafProof,
                latestConfirmedProof
            ),
            "Invalid conflict proof"
        );
        delete leaves[_leaf];

        emit RollupPruned(_leaf);
    }

    function createStake(
        bytes32 location,
        bytes32[] memory proof
    )
        public
        payable
    {
        require(isPath(latestConfirmed, location, proof), "invalid path proof");
        require(msg.value == vmParams.stakeRequirement, "must supply stake value");
        require(stakers[msg.sender].location != 0x00, "Staker already exists");
        stakers[msg.sender] = Staker(
            location,
            block.number,
            false
        );
        stakerCount++;

        emit RollupStakeCreated(msg.sender, location, block.number);
    }

    function moveStake(
        bytes32 newLocation,
        bytes32    _leaf,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        public
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(leaves[_leaf], "invalid leaf");
        require(isPath(staker.location, newLocation, proof1), "stake must move forward");
        require(isPath(newLocation, _leaf, proof2), "node does not exist");

        staker.location = newLocation;

        emit RollupStakeMoved(msg.sender, newLocation);
    }

    function recoverStakeConfirmed(
        bytes32[] memory proof
    )
        public
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(isPath(staker.location, latestConfirmed, proof), "invalid path proof");
        deleteStaker(msg.sender);
        msg.sender.transfer(vmParams.stakeRequirement);

        emit RollupStakeRefunded(msg.sender);
    }

    function recoverStakeMooted(
        bytes32 disputableHash,
        bytes32[] memory latestConfirmedProof,
        bytes32[] memory nodeProof
    )
        public
    {
        Staker storage staker = getValidStaker(msg.sender);
        require(
            isConflict(
                staker.location,
                disputableHash,
                latestConfirmed,
                latestConfirmedProof,
                nodeProof
            ),
            "Invalid conflict proof"
        );
        deleteStaker(msg.sender);
        msg.sender.transfer(vmParams.stakeRequirement);

        emit RollupStakeRefunded(msg.sender);
    }

    // fields
    //  node
    //  beforeHash
    //  beforeInbox
    //  pendingAssertionHash
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startExecutionChallenge(
        bytes32[7] memory _fields,
        address payable[2] memory stakerAddresses,
        uint64 disputableDeadline,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        uint64[2] memory _timeBounds,
        uint32 _importedMessageCount
    )
        public
    {
        return _startExecutionChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartExecutionChallengeData(
                _fields[1],
                _fields[2],
                _timeBounds,
                _fields[3],
                _fields[4],
                _fields[5],
                _importedMessageCount,
                _fields[6]
            )
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  importedAssertion
    //  assertionHash

    function startPendingTopChallenge(
        bytes32[6] memory _fields,
        address payable[2] memory stakerAddresses,
        uint64 disputableDeadline,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        public
    {
        return _startPendingTopChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartPendingTopChallengeData(
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _fields[5]
            )
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startMessagesChallenge(
        bytes32[7] memory _fields,
        address payable[2] memory stakerAddresses,
        uint64 disputableDeadline,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        uint32 _importedMessageCount
    )
        public
    {
        return _startMessagesChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartMessagesChallengeData(
                _fields[1],
                _fields[2],
                _fields[3],
                _fields[4],
                _fields[5],
                _importedMessageCount,
                _fields[6]
            )
        );
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only callable by owner");
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

    function _makeAssertion(MakeAssertionData memory data) internal {
        Staker storage staker = getValidStaker(msg.sender);
        require(leaves[data.prevLeaf], "invalid leaf");
        bytes32 vmProtoHashBefore = VM.protoStateHash(data.beforeVMHash, data.beforeInboxHash, data.beforePendingTop);
        require(
            childNodeHash(
                data.prevPrevLeafHash,
                data.prevDisputableHash,
                data.prevChildType,
                vmProtoHashBefore
            ) == data.prevLeaf,
            "Previous leaf incorrectly unwrapped"
        );

        require(
            !VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(data.numSteps <= vmParams.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(data.timeBounds), "Precondition: not within time bounds");
        require(isPath(latestConfirmed, data.prevLeaf, data.prevLeafProof), "invalid prev leaf proof");
        require(isPath(staker.location, data.prevLeaf, data.stakerProof), "invalid staker location proof");

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
        bytes32 disputableHash = disputableNodeHash(
            deadline,
            Protocol.generatePreconditionHash(
                data.beforeVMHash,
                data.timeBounds,
                data.beforeInboxHash
            ),
            pendingAssertionHash(
                data.afterPendingTop,
                vmParams.pendingInboxHash
            ),
            importedAssertionHash(
                data.beforePendingTop,
                data.importedMessageCount,
                data.importedMessagesSlice
            ),
            assertionHash
        );

        bytes32 validKid = childNodeHash(
            data.prevLeaf,
            disputableHash,
            VALID_CHILD_TYPE,
            VM.protoStateHash(
                data.afterVMHash,
                data.afterInboxHash,
                data.afterPendingTop
            )
        );

        delete leaves[data.prevLeaf];
        leaves[validKid] = true;
        for (uint i = 1; i<=MAX_CHILD_TYPE; i++) {
            leaves[childNodeHash(
                data.prevLeaf,
                disputableHash,
                i,
                vmProtoHashBefore
            )] = true;
        }
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

    function _startPendingTopChallenge(
        ChallengeData memory _challenge,
        StartPendingTopChallengeData memory data
    )
        internal
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_PENDING_TOP_CHILD_TYPE, "Stakers must have a conflict over pending top");

        verifyConflict(
            staker1,
            staker2,
            _challenge.node,
            _challenge.disputableDeadline,
            disputableNodeHash(
                _challenge.disputableDeadline,
                data.preconditionHash,
                pendingAssertionHash(
                    data.afterPendingTop,
                    data.currentPending
                ),
                data.importedAssertion,
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createPendingTopChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.currentPending,
            data.afterPendingTop
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startMessagesChallenge(
        ChallengeData memory _challenge,
        StartMessagesChallengeData memory data
    )
        internal
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_MESSAGES_CHILD_TYPE, "Stakers must have a conflict over pending top");

        verifyConflict(
            staker1,
            staker2,
            _challenge.node,
            _challenge.disputableDeadline,
            disputableNodeHash(
                _challenge.disputableDeadline,
                data.preconditionHash,
                pendingAssertionHash(
                    data.afterPendingTop,
                    data.currentPending
                ),
                importedAssertionHash(
                    data.beforePendingTop,
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createMessagesChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.beforePendingTop,
            data.afterPendingTop,
            data.importedMessageSlice,
            data.importedMessageCount
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startExecutionChallenge(
        ChallengeData memory _challenge,
        StartExecutionChallengeData memory data
    )
        internal
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);
        require(_challenge.stakerPositions[1] == INVALID_EXECUTION_CHILD_TYPE, "Stakers must have a conflict over execution");

        verifyConflict(
            staker1,
            staker2,
            _challenge.node,
            _challenge.disputableDeadline,
            disputableNodeHash(
                _challenge.disputableDeadline,
                Protocol.generatePreconditionHash(
                    data.beforeHash,
                    data.timeBounds,
                    data.beforeInbox
                ),
                data.pendingAssertion,
                importedAssertionHash(
                    data.beforePendingTop,
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            ),
            _challenge.stakerPositions,
            _challenge.vmProtoHashes,
            _challenge.proof1,
            _challenge.proof2
        );

        address newChallengeAddr = challengeFactory.createExecutionChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.beforeHash,
            Protocol.addMessagesToInbox(data.beforeInbox, data.importedMessageSlice),
            data.timeBounds,
            data.assertionHash
        );
        staker1.inChallenge = true;
        staker2.inChallenge = true;

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function verifyConflict(
        Staker storage staker1,
        Staker storage staker2,
        bytes32 node,
        uint64 disputableDeadline,
        bytes32 disputableNodeHashVal,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2

    )
        internal
        view
    {
        require(staker1.creationTime < disputableDeadline, "staker1 staked after deadline");
        require(staker2.creationTime < disputableDeadline, "staker2 staked after deadline");
        require(!staker1.inChallenge, "staker1 already in a challenge");
        require(!staker2.inChallenge, "staker2 already in a challenge");
        require(
            isSpecifiedConflict(
                node,
                disputableNodeHashVal,
                stakerPositions, vmProtoHashes,
                staker1.location, proof1,
                staker2.location, proof2
            ),
            "Invalid conflict proof"
        );
    }

    function childNodeHash(
        bytes32 prevNodeHash,
        bytes32 disputableNodeHash,
        uint    childType,
        bytes32 vmProtoStateHash
    )
        internal
        pure
        returns(bytes32)
    {
        require((childType>=VALID_CHILD_TYPE) && (childType<=MAX_CHILD_TYPE), "Invalid child type");
        return keccak256(
            abi.encodePacked(
                prevNodeHash,
                keccak256(
                    abi.encodePacked(
                        disputableNodeHash,
                        childType,
                        vmProtoStateHash
                    )
                )
            )
        );
    }

    function importedAssertionHash(bytes32 beforePendingTop, uint32 messageCount, bytes32 messagesSlice) internal pure returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                beforePendingTop,
                messageCount,
                messagesSlice
            )
        );
    }

    function pendingAssertionHash(bytes32 afterPendingTop, bytes32 currentPending) internal pure returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                afterPendingTop,
                currentPending
            )
        );
    }

    function disputableNodeHash(
        uint deadline,
        bytes32 preconditionHash,
        bytes32 pendingAssertion,
        bytes32 importedAssertion,
        bytes32 assertionHash
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                deadline,
                preconditionHash,
                pendingAssertion,
                importedAssertion,
                assertionHash
            )
        );
    }

    function isPath(bytes32 from, bytes32 to, bytes32[] memory proof) internal pure returns(bool) {
        return isPathOffset(
            from,
            to,
            proof,
            0,
            proof.length
        );
    }

    function isPathOffset(
        bytes32 from,
        bytes32 to,
        bytes32[] memory proof,
        uint start,
        uint end
    )
        internal
        pure
        returns(bool)
    {
        bytes32 node = from;
        for (uint i = start; i<end; i++) {
            node = keccak256(abi.encodePacked(node, proof[i]));
        }
        return (node==to);
    }

    function isConflict(
        bytes32 from,
        bytes32 to1,
        bytes32 to2,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        internal
        pure
        returns(bool)
    {
        return (proof1[0] != proof2[0]) &&
            isPath(from, to1, proof1) &&
            isPath(from, to2, proof2);
    }

    function isSpecifiedConflict(
        bytes32 from,
        bytes32 disputableNode,
        uint[2] memory childTypes,
        bytes32[2] memory vmProtoHashes,
        bytes32 to1,
        bytes32[] memory proof1,
        bytes32 to2,
        bytes32[] memory proof2
    )
        internal
        pure
        returns(bool)
    {
        require(childTypes[0] < childTypes[1], "Child types must be ordered");
        return isPath(
            childNodeHash(
                from,
                disputableNode,
                childTypes[0],
                vmProtoHashes[0]
            ),
            to1,
            proof1
        ) && isPath(
            childNodeHash(
                from,
                disputableNode,
                childTypes[1],
                vmProtoHashes[1]
            ),
            to2,
            proof2
        );
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) internal view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }
}
