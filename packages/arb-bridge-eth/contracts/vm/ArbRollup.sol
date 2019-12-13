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
import "./Disputable.sol";
import "./IArbBase.sol";

import "../IGlobalPendingInbox.sol";

import "../challenge/ChallengeLauncher.sol";

import "../arch/Protocol.sol";
import "../arch/Value.sol";


//TODO: emit events to announce everything interesting that happens

contract ArbRollup is IArbBase {
    using SafeMath for uint256;

    // fields:
        // beforeHash
        // beforeInbox
        // afterHash
        // messagesAccHash
        // logsAccHash

    event PendingDisputableAssertion (
        bytes32[5] fields,
        address asserter,
        uint64[2] timeBounds,
        uint32 numSteps,
        uint64 deadline
    );

    event ConfirmedDisputableAssertion(
        bytes32 newState,
        bytes32 logsAccHash
    );

    event ChallengeLaunched(
        address challengeContract,
        address challenger
    );

    address internal constant ETH_ADDRESS = address(0);

    IGlobalPendingInbox public globalInbox;

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
    address[] activeChallenges;

    function myStakerIndex() internal view returns(uint) {
        uint index = stakerIndex[msg.sender];
        require(stakers[index].addr == msg.sender, "must be called by a staker");
        return index;
    }

    function getStakerIndex(address addr) internal view returns(uint) {
        uint index = stakerIndex[addr];
        require(stakers[index].addr == addr, "not a staker");
        return index;
    }

    function deleteStaker(uint index) internal {
        delete stakerIndex[stakers[index].addr];
        if (index < stakers.length-1) {
            stakers[index] = stakers[stakers.length-1];
            stakerIndex[stakers[index].addr] = index;
        }
        stakers.pop();
    }

    uint constant ValidChildType = 0;
    uint constant InvalidPendingTopChildType = 1;
    uint constant InvalidMessagesChildType = 2;
    uint constant InvalidExecutionChildType = 3;
    uint constant MaxChildType = 3;

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
        require((childType>=ValidChildType) && (childType<=MaxChildType), "Invalid child type");
        return keccak256(abi.encodePacked(
            prevNodeHash,
            keccak256(abi.encodePacked(
                disputableNodeHash,
                childType,
                vmProtoStateHash
            ))
        ));
    }

    function isPath(bytes32 from, bytes32 to, bytes32[] memory proof) internal pure returns(bool) {
        bytes32 node = from;
        for (uint i=0; i<proof.length; i++) {
            node = keccak256(abi.encodePacked(node, proof[i]));
        }
        return (node==to);
    }

    function isPath_offset(bytes32 from, bytes32 to, bytes32[] memory proof, uint start, uint end) internal pure returns(bool) {
        bytes32 node = from;
        for (uint i=start; i<end; i++) {
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
        bytes32 disputableNodeHash,
        uint    childType1,
        bytes32 vmProtoHash1,
        bytes32 to1,
        bytes32[] memory proof1,
        uint    childType2,
        bytes32 vmProtoHash2,
        bytes32 to2,
        bytes32[] memory proof2
    )
        internal
        pure
        returns(bool)
    {
        require(childType1 < childType2);
        return isPath(childNodeHash(from, disputableNodeHash, childType1, vmProtoHash1), to1, proof1) &&
            isPath(childNodeHash(from, disputableNodeHash, childType2, vmProtoHash2), to2, proof2);
    }

    function disputableNodeHash(
        uint deadline, 
        bytes32 preconditionHash, 
        bytes32 assertionHash
    ) 
        internal
        pure 
        returns(bytes32) 
    {
        return keccak256(abi.encodePacked(
            deadline,
            preconditionHash,
            assertionHash
        ));
    }

    function withinTimeBounds(uint64[2] memory _timeBounds) public view returns (bool) {
        return block.number >= _timeBounds[0] && block.number <= _timeBounds[1];
    }

    function makeAssertion(
        bytes32 beforeVMHash,
        bytes32 beforeInboxHash,
        bytes32 _prevPrevLeafHash,
        bytes32 _prevDisputableHash,
        uint    _prevChildType,
        bytes32 _prevVMprotoHash,
        uint _prevLeafIndex,
        bytes32[] memory _prevLeafProof,
        bytes32[] memory _stakerProof,
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
        Staker memory staker = stakers[myStakerIndex()];
        require(_prevLeafIndex < leaves.length, "invalid leaf index");
        bytes32 prevLeaf = leaves[_prevLeafIndex];
        require(
            childNodeHash(
                _prevPrevLeafHash, 
                _prevDisputableHash,
                _prevChildType,
                _prevVMprotoHash
            ) == VM.protoStateHash(beforeVMHash, beforeInboxHash),
            "Precondition does not match prior state"
        );
        require(
            !VM.isErrored(beforeVMHash) && !VM.isHalted(beforeVMHash),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(_numSteps <= vmParams.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(_timeBounds), "Precondition: not within time bounds");
        require(isPath(latestConfirmed, prevLeaf, _prevLeafProof), "invalid prev leaf proof");
        require(isPath(staker.location, prevLeaf, _stakerProof), "invalid staker location proof");

        uint deadline = block.number + vmParams.gracePeriod; //TODO: [Ed] compute this properly
        bytes32 vmProtoHashBefore = VM.protoStateHash(beforeVMHash, beforeInboxHash);
        bytes32 vmProtoHashAfter = VM.protoStateHash(_afterVMHash, _afterInboxHash);
        bytes32 disputableHash = disputableNodeHash(
            deadline,
            Protocol.generatePreconditionHash(
                beforeVMHash,
                _timeBounds,
                beforeInboxHash
            ),
            Protocol.generateAssertionHash(
                _afterVMHash,
                _numSteps,
                _numArbGas,
                0x00,
                _messagesAccHash,
                0x00,
                _logsAccHash
            )
        );

        bytes32 validKid = childNodeHash(prevLeaf, disputableHash, ValidChildType, vmProtoHashAfter);
        leaves[_prevLeafIndex] = leaves[leaves.length - 1];
        leaves[leaves.length - 1] = validKid;
        for (uint i=1; i<=MaxChildType; i++) {
            leaves.push(childNodeHash(prevLeaf, disputableHash, i, vmProtoHashBefore));
        }
        stakers[myStakerIndex()].location = validKid;
    }

    function confirm(
        bytes32 to,
        uint    _leafIndex,
        bytes32[] memory proof1,
        bytes32[] memory stakerProofs,
        uint[]  memory stakerProofOffsets,
        bytes32 prev,
        uint    branch,
        uint    deadline,
        bytes32 _preconditionHash,
        bytes32 _assertionHash
    )
        public
    {
        require(_leafIndex < leaves.length, "invalid leaf index");
        bytes32 leaf = leaves[_leafIndex];
        require(isPath(to, leaf, proof1), "node does not exist");
        require(keccak256(abi.encodePacked(
            prev,
            keccak256(abi.encodePacked(
                keccak256(abi.encodePacked(
                    deadline,
                    _preconditionHash,
                    _assertionHash
                )),
                branch
            ))
        )) == to, "invalid parameters for prev node");

        for (uint i=0; i<stakers.length; i++) {
            require((stakers[i].creationTime >= deadline) || isPath_offset(to, stakers[i].location, stakerProofs, stakerProofOffsets[i], stakerProofOffsets[i+1]),
                "at least one active staker disagrees");
        }

        latestConfirmed = to;
        //TODO: execute actions from the confirmed assertion (to)
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
            isConflict(from, leaf, latestConfirmed, leafProof, latestConfirmedProof),
            "Invalid conflict proof"
        );
        leaves[_leafIndex] = leaves[leaves.length - 1];
        leaves.pop();
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
        require(stakers[stakerIndex[msg.sender]].addr != msg.sender, "cannot be called by a staker");
        Staker memory staker;
        staker.addr = msg.sender;
        staker.location = location;
        staker.creationTime = block.number;
        stakerIndex[msg.sender] = stakers.length;
        stakers.push(staker);
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
        Staker memory staker = stakers[myStakerIndex()];
        require(isPath(staker.location, newLocation, proof1), "stake must move forward");
        require(isPath(newLocation, leaf, proof2), "node does not exist");

        stakers[myStakerIndex()].location = newLocation;
    }

    function recoverStakeConfirmed(
        bytes32[] memory proof
    )
        public
    {
        uint index = myStakerIndex();
        Staker memory staker = stakers[index];
        require(isPath(staker.location, latestConfirmed, proof), "invalid path proof");
        delete stakerIndex[msg.sender];
        if (index < stakers.length-1) {
            stakers[index] = stakers[stakers.length-1];
            stakerIndex[stakers[index].addr] = index;
        }
        stakers.pop();
        msg.sender.transfer(vmParams.stakeRequirement);
    }

    function recoverStakeMooted(
        bytes32 disputableHash,
        bytes32[] memory latestConfirmedProof,
        bytes32[] memory nodeProof
    )
        public
    {
        uint index = myStakerIndex();
        Staker memory staker = stakers[index];
        require(
            isConflict(staker.location, disputableHash, latestConfirmed, latestConfirmedProof, nodeProof),
            "Invalid conflict proof"
        );
        delete stakerIndex[msg.sender];
        if (index < stakers.length-1) {
            stakers[index] = stakers[stakers.length-1];
            stakerIndex[stakers[index].addr] = index;
        }
        stakers.pop();
        msg.sender.transfer(vmParams.stakeRequirement);
    }

    function startChallenge(
        address staker1Address,
        address staker2Address,
        bytes32 node,
        uint64 disputableDeadline,
        bytes32 disputableHash,
        uint    staker1position,
        uint    staker2position,
        bytes32 vmProtoHash1,
        bytes32 vmProtoHash2,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _assertionHash
    )
        public
    {
        Staker memory staker1 = stakers[getStakerIndex(staker1Address)];
        Staker memory staker2 = stakers[getStakerIndex(staker2Address)];

        require(keccak256(abi.encodePacked(
            disputableDeadline,
            Protocol.generatePreconditionHash(
                _beforeHash,
                _timeBounds,
                _beforeInbox
            ),
            _assertionHash
        )) == disputableHash);
        require(staker1.creationTime < disputableDeadline, "staker1 staked after deadline");
        require(staker2.creationTime < disputableDeadline, "staker2 staked after deadline");
        require(staker1.challenge == address(0), "staker1 already in a challenge");
        require(staker2.challenge == address(0), "staker2 already in a challenge");
        require(
            isSpecifiedConflict(
                node, disputableHash,
                staker1position, vmProtoHash1, staker1.location, proof1, 
                staker2position, vmProtoHash2, staker2.location, proof2
            ),
            "Invalid conflict proof"
        );
   
        address newChallengeAddr;
        if (staker2position==InvalidPendingTopChildType) {
            newChallengeAddr = ChallengeLauncher.startInvalidPendingTopChallenge(
                staker1Address, 
                staker2Address,
                disputableHash
            );
        } else if (staker2position==InvalidMessagesChildType) {
            newChallengeAddr = ChallengeLauncher.startInvalidMessagesChallenge(
                staker1Address,
                staker2Address,
                disputableHash
            );
        } else {
            newChallengeAddr = ChallengeLauncher.startExecutionChallenge(
                staker1Address,
                staker2Address,
                disputableHash
            );
        }
        staker1.challenge = newChallengeAddr;
        staker2.challenge = newChallengeAddr;
    }

    function resolveChallenge(address winner, address loser) public {
        uint winnerIndex = getStakerIndex(winner);
        uint loserIndex = getStakerIndex(loser);
        require(stakers[winnerIndex].challenge==msg.sender, "verdict can only be declared by challenge");
        require(stakers[loserIndex].challenge==msg.sender, "verdict can only be declared by challenge");
        //TODO: slash the loser, deliver half to the winner
        stakers[winnerIndex].challenge = address(0);
        deleteStaker(loserIndex);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only callable by owner");
        _;
    }

    function initialize(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _stakeRequirement,
        address payable _owner,
        address _globalInboxAddress
    )
        public
    {
        require(latestConfirmed == bytes32(0), "VM already initialized");

        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        globalInbox.registerForInbox();
        owner = _owner;

        // VM parameters
        vmParams.stakeRequirement = _stakeRequirement;
        vmParams.gracePeriod = _gracePeriod;
        vmParams.maxExecutionSteps = _maxExecutionSteps;
        vmParams.pendingHash = Value.hashEmptyTuple();

        // VM protocol state
        bytes32 vmProtoStateHash = VM.protoStateHash(_vmState, Value.hashEmptyTuple());
        latestConfirmed = childNodeHash(0, 0, 0, vmProtoStateHash);
        leaves.push(latestConfirmed);
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
}