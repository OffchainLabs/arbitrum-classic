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
        address addr;
        bytes32 location;
        uint    creationTime;
        address challenge;
    }

    address   owner;
    VM.Params vmParams;
    bytes32   latestConfirmed;
    mapping(address => uint) stakerIndex;
    Staker[]  stakers;
    bytes32[] leaves;


    // Fields
    //   prevLeafHash
    //   afterPendingTop
    //   importedMesssagesSlice
    //   afterVMHash
    //   messagesAccHash
    //   logsAccHash

    event RollupAsserted(
        bytes32[6] fields,
        uint _importedMessageCount,
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

    modifier nonStaker() {
        require(stakerIndex[msg.sender] == 0, "Sender cannot be a staker");
        _;
    }

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
        uint prevLeafIndex;
        bytes32[] prevLeafProof;
        bytes32[] stakerProof;
        bytes32 afterPendingTop;
        bytes32 importedMessagesSlice;
        uint importedMessageCount;
        bytes32 afterVMHash;
        bytes32 afterInboxHash;
        bytes32 messagesAccHash;
        bytes32 logsAccHash;
        uint32 numSteps;
        uint64 numArbGas;
        uint64[2] timeBounds;
    }

    struct StartExecutionChallengeData {
        address[2] stakerAddresses;
        bytes32 node;
        uint64 disputableDeadline;
        uint[2] stakerPositions;
        bytes32[2] vmProtoHashes;
        bytes32[] proof1;
        bytes32[] proof2;
        bytes32 beforeHash;
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes32 afterPendingTop;
        bytes32 importedMessageSlice;
        uint importedMessageCount;
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
        leaves.push(latestConfirmed);
    }

    function resolveChallenge(address winner, address loser) external {
        Staker storage winningStaker = getStaker(winner);
        Staker storage loserStaker = getStaker(loser);
        require(winningStaker.challenge==msg.sender, "verdict can only be declared by challenge contract");
        require(loserStaker.challenge==msg.sender, "verdict can only be declared by challenge contract");
        //TODO: slash the loser, deliver half to the winner
        winningStaker.challenge = address(0);
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, winner, loser);
    }

    // fields
    //  beforeVMHash
    //  beforeInboxHash
    //  beforePendingTop
    //  prevPrevLeafHash
    //  prevDisputableHash
    //  afterPendingTop
    //  importedMessagesSlice

    function makeAssertion(
        bytes32[7] memory _fields,
        uint    _prevChildType,
        uint _prevLeafIndex,
        bytes32[] memory _prevLeafProof,
        bytes32[] memory _stakerProof,
        uint _importedMessageCount,
        bytes32 _afterVMHash,
        bytes32 _afterInboxHash,
        bytes32 _messagesAccHash,
        bytes32 _logsAccHash,
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
                _prevLeafIndex,
                _prevLeafProof,
                _stakerProof,
                _fields[5],
                _fields[6],
                _importedMessageCount,
                _afterVMHash,
                _afterInboxHash,
                _messagesAccHash,
                _logsAccHash,
                _numSteps,
                _numArbGas,
                _timeBounds
            )
        );
    }

    function confirm(
        uint    _leafIndex,
        bytes32[] memory proof1,
        bytes32[] memory stakerProofs,
        uint[]  memory stakerProofOffsets,
        bytes32 prev,
        uint    branch,
        uint    deadline,
        bytes32 _preconditionHash,
        bytes32 _afterPendingTop,
        bytes32 _importedAssertionHash,
        bytes32 _executionAssertionHash,
        bytes32 _vmProtoStateHash
    )
        public
    {
        require(_leafIndex < leaves.length, "invalid leaf index");
        bytes32 to = childNodeHash(
            prev,
            disputableNodeHash(
                deadline,
                _preconditionHash,
                _afterPendingTop,
                _importedAssertionHash,
                _executionAssertionHash
            ),
            branch,
            _vmProtoStateHash
        );
        require(isPath(to, leaves[_leafIndex], proof1), "node does not exist");
        for (uint i = 0; i<stakers.length; i++) {
            require(
                (stakers[i].creationTime >= deadline) ||
                isPathOffset(
                    to,
                    stakers[i].location,
                    stakerProofs,
                    stakerProofOffsets[i],
                    stakerProofOffsets[i+1]
                ),
                "at least one active staker disagrees");
        }

        latestConfirmed = to;
        if (branch == 0) {
            //TODO: execute actions from the DA before the confirmed assertion (to)
        }

        emit RollupConfirmed(latestConfirmed);
    }

    function pruneLeaf(
        uint _leafIndex,
        bytes32 from,
        bytes32[] memory leafProof,
        bytes32[] memory latestConfirmedProof
    )
        public
    {
        require(_leafIndex < leaves.length, "invalid leaf index");
        bytes32 leaf = leaves[_leafIndex];
        require(
            isConflict(
                from,
                leaf,
                latestConfirmed,
                leafProof,
                latestConfirmedProof
            ),
            "Invalid conflict proof"
        );
        leaves[_leafIndex] = leaves[leaves.length - 1];
        leaves.pop();

        emit RollupPruned(leaf);
    }

    function createStake(
        bytes32 location,
        bytes32[] memory proof
    )
        public
        payable
        nonStaker
    {
        require(isPath(latestConfirmed, location, proof), "invalid path proof");
        require(msg.value == vmParams.stakeRequirement, "must supply stake value");
        Staker memory staker;
        staker.addr = msg.sender;
        staker.location = location;
        staker.creationTime = block.number;
        stakerIndex[msg.sender] = stakers.length + 1;
        stakers.push(staker);

        emit RollupStakeCreated(msg.sender, location, block.number);
    }

    function moveStake(
        bytes32 newLocation,
        uint    _leafIndex,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        public
    {
        require(_leafIndex < leaves.length, "invalid leaf index");
        bytes32 leaf = leaves[_leafIndex];
        Staker storage staker = getStaker(msg.sender);
        require(isPath(staker.location, newLocation, proof1), "stake must move forward");
        require(isPath(newLocation, leaf, proof2), "node does not exist");

        staker.location = newLocation;

        emit RollupStakeMoved(msg.sender, newLocation);
    }

    function recoverStakeConfirmed(
        bytes32[] memory proof
    )
        public
    {
        Staker storage staker = getStaker(msg.sender);
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
        Staker storage staker = getStaker(msg.sender);
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
    //  afterPendingTop
    //  importedMessageSlice
    //  assertionHash

    function startExecutionChallenge(
        bytes32[6] memory _fields,
        address[2] memory stakerAddresses,
        uint64 disputableDeadline,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        uint64[2] memory _timeBounds,
        uint _importedMessageCount
    )
        public
    {
        return _startExecutionChallenge(
            StartExecutionChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2,
                _fields[1],
                _fields[2],
                _timeBounds,
                _fields[3],
                _fields[4],
                _importedMessageCount,
                _fields[5]
            )
        );
    }

    // function startChallenge(
    //     address staker1Address,
    //     address staker2Address,
    //     bytes32 node,
    //     uint64 disputableDeadline,
    //     bytes32 disputableHash,
    //     uint    staker1position,
    //     uint    staker2position,
    //     bytes32 vmProtoHash1,
    //     bytes32 vmProtoHash2,
    //     bytes32[] memory proof1,
    //     bytes32[] memory proof2,
    //     bytes32 _beforeHash,
    //     bytes32 _beforeInbox,
    //     uint64[2] memory _timeBounds,
    //     bytes32 _assertionHash
    // )
    //     public
    // {
    //     Staker memory staker1 = stakers[getStakerIndex(staker1Address)];
    //     Staker memory staker2 = stakers[getStakerIndex(staker2Address)];

    //     require(keccak256(abi.encodePacked(
    //         disputableDeadline,
    //         Protocol.generatePreconditionHash(
    //             _beforeHash,
    //             _timeBounds,
    //             _beforeInbox
    //         ),
    //         _assertionHash
    //     )) == disputableHash);
    //     require(staker1.creationTime < disputableDeadline, "staker1 staked after deadline");
    //     require(staker2.creationTime < disputableDeadline, "staker2 staked after deadline");
    //     require(staker1.challenge == address(0), "staker1 already in a challenge");
    //     require(staker2.challenge == address(0), "staker2 already in a challenge");
    //     require(
    //         isSpecifiedConflict(
    //             node, disputableHash,
    //             staker1position, vmProtoHash1, staker1.location, proof1,
    //             staker2position, vmProtoHash2, staker2.location, proof2
    //         ),
    //         "Invalid conflict proof"
    //     );

    //     address newChallengeAddr;
    //     if (staker2position==INVALID_PENDING_TOP_CHILD_TYPE) {
    //         newChallengeAddr = challengeFactory.createPendingTopChallenge(
    //             staker1Address,
    //             staker2Address,
    //             disputableHash
    //         );
    //     } else if (staker2position==INVALID_MESSAGES_CHILD_TYPE) {
    //         newChallengeAddr = challengeFactory.createMessagesChallenge(
    //             staker1Address,
    //             staker2Address,
    //             disputableHash
    //         );
    //     } else {
    //         newChallengeAddr = challengeFactory.createExecutionChallenge(
    //             staker1Address,
    //             staker2Address,
    //             disputableHash
    //         );
    //     }
    //     staker1.challenge = newChallengeAddr;
    //     staker2.challenge = newChallengeAddr;

    //     emit rollupChallengeStarted(staker1.addr, staker2.addr, staker2position, newChallengeAddr);
    // }

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
        Staker storage staker = getStaker(msg.sender);
        require(data.prevLeafIndex < leaves.length, "invalid leaf index");
        bytes32 prevLeaf = leaves[data.prevLeafIndex];
        bytes32 vmProtoHashBefore = VM.protoStateHash(data.beforeVMHash, data.beforeInboxHash, data.beforePendingTop);
        require(
            childNodeHash(
                data.prevPrevLeafHash,
                data.prevDisputableHash,
                data.prevChildType,
                vmProtoHashBefore
            ) == prevLeaf,
            "Previous leaf incorrectly unwrapped"
        );

        require(
            !VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(data.numSteps <= vmParams.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(data.timeBounds), "Precondition: not within time bounds");
        require(isPath(latestConfirmed, prevLeaf, data.prevLeafProof), "invalid prev leaf proof");
        require(isPath(staker.location, prevLeaf, data.stakerProof), "invalid staker location proof");

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
            data.afterPendingTop,
            importedAssertionHash(
                data.importedMessageCount,
                data.importedMessagesSlice
            ),
            assertionHash
        );

        bytes32 validKid = childNodeHash(
            prevLeaf,
            disputableHash,
            VALID_CHILD_TYPE,
            VM.protoStateHash(
                data.afterVMHash,
                data.afterInboxHash,
                data.afterPendingTop
            )
        );

        leaves[data.prevLeafIndex] = leaves[leaves.length - 1];
        leaves[leaves.length - 1] = validKid;
        for (uint i = 1; i<=MAX_CHILD_TYPE; i++) {
            leaves.push(childNodeHash(
                prevLeaf,
                disputableHash,
                i,
                vmProtoHashBefore
            ));
        }
        staker.location = validKid;

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

    function _startExecutionChallenge(StartExecutionChallengeData memory data) internal {
        Staker storage staker1 = getStaker(data.stakerAddresses[0]);
        Staker storage staker2 = getStaker(data.stakerAddresses[1]);
        require(data.stakerPositions[1] == INVALID_EXECUTION_CHILD_TYPE, "Stakers must have a conflict over execution");

        verifyConflict(
            staker1,
            staker2,
            data.node,
            data.disputableDeadline,
            disputableNodeHash(
                data.disputableDeadline,
                Protocol.generatePreconditionHash(
                    data.beforeHash,
                    data.timeBounds,
                    data.beforeInbox
                ),
                data.afterPendingTop,
                importedAssertionHash(
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            ),
            data.stakerPositions,
            data.vmProtoHashes,
            data.proof1,
            data.proof2
        );

        address newChallengeAddr = challengeFactory.createExecutionChallenge(
            data.stakerAddresses[0],
            data.stakerAddresses[1],
            0, // Challenge period
            data.beforeHash,
            Protocol.addMessagesToInbox(data.beforeInbox, data.importedMessageSlice),
            data.timeBounds,
            data.assertionHash
        );
        staker1.challenge = newChallengeAddr;
        staker2.challenge = newChallengeAddr;

        emit RollupChallengeStarted(
            data.stakerAddresses[0],
            data.stakerAddresses[1],
            data.stakerPositions[1],
            newChallengeAddr
        );
    }

    // function _startPendingTopChallenge(StartExecutionChallengeData memory data) internal {
    //     Staker storage staker1 = getStaker(data.stakerAddresses[0]);
    //     Staker storage staker2 = getStaker(data.stakerAddresses[1]);
    //     require(data.stakerPositions[1] == INVALID_PENDING_TOP_CHILD_TYPE, "Stakers must have a conflict over pending top");

    //     verifyConflict(
    //         staker1,
    //         staker2,
    //         data.node,
    //         data.disputableDeadline,
    //         disputableNodeHash(
    //             data.disputableDeadline,
    //             Protocol.generatePreconditionHash(
    //                 data.beforeHash,
    //                 data.timeBounds,
    //                 data.beforeInbox
    //             ),
    //             data.afterPendingTop,
    //             importedAssertionHash(
    //                 data.importedMessageCount,
    //                 data.importedMessageSlice
    //             ),
    //             data.assertionHash
    //         ),
    //         data.stakerPositions,
    //         data.vmProtoHashes,
    //         data.proof1,
    //         data.proof2
    //     );

    //     address newChallengeAddr = challengeFactory.createPendingChallenge(
    //         data.stakerAddresses[0],
    //         data.stakerAddresses[1],
    //         0, // Challenge period
    //         data.beforeHash,
    //         Protocol.addMessagesToInbox(data.beforeInbox, data.importedMessageSlice),
    //         data.timeBounds,
    //         data.assertionHash
    //     );
    //     staker1.challenge = newChallengeAddr;
    //     staker2.challenge = newChallengeAddr;

    //     emit RollupChallengeStarted(
    //         data.stakerAddresses[0],
    //         data.stakerAddresses[1],
    //         data.stakerPositions[1],
    //         newChallengeAddr
    //     );
    // }

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
        require(staker1.challenge == address(0), "staker1 already in a challenge");
        require(staker2.challenge == address(0), "staker2 already in a challenge");
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

    function getStaker(address addr) internal view returns(Staker storage) {
        uint index = stakerIndex[addr];
        require(index != 0, "not a staker");
        return stakers[index - 1];
    }

    function deleteStaker(address staker) internal {
        uint index = stakerIndex[staker];
        delete stakerIndex[staker];
        if (index < stakers.length-1) {
            stakers[index] = stakers[stakers.length-1];
            stakerIndex[stakers[index].addr] = index + 1;
        }
        stakers.pop();
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

    function importedAssertionHash(uint messageCount, bytes32 messagesSlice) internal pure returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                messageCount,
                messagesSlice
            )
        );
    }

    function disputableNodeHash(
        uint deadline,
        bytes32 preconditionHash,
        bytes32 afterPendingTop,
        bytes32 importedHash,
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
                afterPendingTop,
                importedHash,
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
