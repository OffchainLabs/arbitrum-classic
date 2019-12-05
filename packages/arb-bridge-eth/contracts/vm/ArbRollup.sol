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

import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";
import "../arch/Value.sol";


contract ArbBase is IArbBase {
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

    IChallengeFactory public challengeFactory;
    IGlobalPendingInbox public globalInbox;

    struct Staker {
        bytes32 location;
        uint64 creationTime;
        address challenge;
    };

    uint stakeRequirement;
    bytes32 latestConfirmed;
    mapping(address => Staker) stakers;
    bytes32[] leaves;

    function validChild(bytes32 prevLeaf, bytes32 disputableHash) returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                prevLeaf,
                disputableHash,
                0
            )
        );
    }

    function invalidChild(bytes32 prevLeaf, bytes32 disputableHash) returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                prevLeaf,
                disputableHash,
                1
            )
        );
    }

    function isPath(bytes32 from, bytes32 to, bytes32[] proof) returns(bool) {

    }

    function isConflict(
        bytes32 from,
        bytes32 disputableHash,
        bytes32 toValid,
        bytes32 toInvalid,
        bytes32[] validProof,
        bytes32[] invalidProof
    )
        returns(bool)
    {
        bytes32 validChild = validChild(from, disputableHash);
        bytes32 invalidChild = invalidChild(from, disputableHash);
        return isPath(validChild, toValid, validProof) &&
            isPath(invalidChild, toInvalid, invalidProof);
    }

    function assert(
        bytes32 beforeHash,
        bytes32 beforeInbox,
        uint _prevLeafIndex,
        bytes32[] _prevLeafProof,
        bytes32[] _stakerProof,
        bytes32 _afterHash,
        bytes32 _messagesAccHash,
        bytes32 _logsAccHash,
        uint32 _numSteps,
        uint64[2] memory _timeBounds
    )
        public
    {
        Staker memory staker = stakers[msg.sender];
        require(staker.creationTime > 0, "must be called by a staker");
        require(_prevLeafIndex < leaves.length, "invalid leaf index");
        bytes32 prevLeaf = leaves[_prevLeafIndex];
        require(
            keccak256(abi.encodePacked(
                beforeHash,
                beforeInbox,
            )) == prevLeaf,
            "invalid prevLeaf reveal"
        )
        require(
            !VM.isErrored(vm) && !VM.isHalted(vm),
            "Can only disputable assert if machine is not errored or halted"
        );
        require(numSteps <= vm.maxExecutionSteps, "Tried to execute too many steps");
        require(withinTimeBounds(timeBounds), "Precondition: not within time bounds");
        require(isPath(latestConfirmed, prevLeaf, _prevLeafProof), "invalid prev leaf proof");
        require(isPath(staker.location, prevLeaf, _stakerProof), "invalid prev leaf proof");

        bytes32 disputableHash = keccak256(
            abi.encodePacked(
                block.number + deadline,
                Protocol.generatePreconditionHash(
                    beforeHash,
                    timeBounds,
                    beforeInbox
                ),
                Protocol.generateAssertionHash(
                    afterHash,
                    numSteps,
                    0x00,
                    messagesAccHash,
                    0x00,
                    logsAccHash
                )
            )
        );

        bytes32 validChild = validChild(prevLeaf, disputableHash);
        bytes32 invalidChild = invalidChild(prevLeaf, disputableHash);

        leaves[_prevLeafIndex] = leaves[leaves.length - 1];
        leaves[leaves.length - 1] = validChild;
        leaves.push(invalidChild);
        staker.location = validChild;
    }


    function pruneLeaf(
        uint _leafIndex,
        bytes32 from,
        bool isValidChild,
        bytes32 disputableHash,
        bytes32[] validProof,
        bytes32[] invalidProof
    )
        public
    {
        require(_leafIndex < leaves.length, "invalid leaf index");
        bytes32 leaf = leaves[_leafIndex];
        if (isValidChild) {
            require(
                isConflict(from, disputableHash, leaf, latestConfirmed, validProof, invalidProof),
                "Invalid conflict proof"
            );
        } else {
            require(
                isConflict(from, disputableHash, latestConfirmed, leaf, validProof, invalidProof),
                "Invalid conflict proof"
            );
        }
        leaves[_leafIndex] = leaves[leaves.length - 1];
        leaves.pop();
    }

    function createStake(
        bytes32 location,
        bytes32[] proof
    )
        public
        payable
    {
        require(isPath(latestConfirmed, location, proof), "invalid path proof");
        require(msg.amount == stakeRequirement, "must supply stake value");
        Staker memory staker = stakers[msg.sender];
        require(staker.creationTime == 0, "cannot be called by a staker");
        staker.location = location;
        staker.creationTime = block.number;
    }

    function recoverStakeA(
        bytes32[] proof
    )
        public
    {
        Staker memory staker = stakers[msg.sender];
        require(staker.challenge == address(0), "staker in challenge");
        require(isConflict(staker.location, latestConfirmed, proof), "invalid path proof");
        delete stakers[msg.sender];
        msg.sender.transfer(stakeRequirement);
    }

    function recoverStakeB(
        bytes32 node,
        bool isValidChild,
        bytes32 disputableHash,
        bytes32[] validProof,
        bytes32[] invalidProof
    )
        public
    {
        Staker memory staker = stakers[msg.sender];
        require(staker.challenge == address(0), "staker in challenge");
        if (isValidChild) {
            require(
                isConflict(staker.location, disputableHash, latestConfirmed, node, validProof, invalidProof),
                "Invalid conflict proof"
            );
        } else {
            require(
                isConflict(staker.location, disputableHash, node, latestConfirmed, validProof, invalidProof),
                "Invalid conflict proof"
            );
        }
        delete stakers[msg.sender];
        msg.sender.transfer(stakeRequirement);
    }

    function startChallenge(
        address staker1Address,
        address staker2Address,
        bytes32 node,
        uint64 disputableDeadline,
        bytes32 disputableHash,
        bytes32[] validProof,
        bytes32[] invalidProof,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _assertionHash,
    )
        public
    {
        Staker memory staker1 = stakers[staker1Address];
        Staker memory staker2 = stakers[staker2Address];

        require(keccak256(abi.encodePacked(
            disputableDeadline,
            Protocol.generatePreconditionHash(
                beforeHash,
                timeBounds,
                beforeInbox
            ),
            _assertionHash
        )) == disputableHash);
        require(staker1.creationTime < disputableDeadline);
        require(staker2.creationTime < disputableDeadline);
        require(staker1.challenge == address(0));
        require(staker2.challenge == address(0));
        require(
            isConflict(node, disputableHash, staker1.location, staker2.location, validProof, invalidProof),
            "Invalid conflict proof"
        );
        validatorBalances[msg.sender] -= vm.escrowRequired;

        Disputable.initiateChallenge(
            vm,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );

        vm.activeChallengeManager = challengeFactory.createChallenge(
            [vm.asserter, msg.sender],
            [vm.escrowRequired, vm.escrowRequired],
            vm.gracePeriod,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );

        emit ChallengeLaunched(vm.activeChallengeManager, msg.sender);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only callable by owner");
        _;
    }

    function initialize(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _escrowRequired,
        address payable _owner,
        address _challengeFactoryAddress,
        address _globalInboxAddress
    )
        public
    {
        require(address(challengeFactory) == address(0), "VM already initialized");
        require(_challengeFactoryAddress != address(0), "Challenge factory address not set");

        globalInbox = IGlobalPendingInbox(_globalInboxAddress);
        challengeFactory = IChallengeFactory(_challengeFactoryAddress);

        globalInbox.registerForInbox();
        owner = _owner;

        // Machine state
        vm.machineHash = _vmState;
        vm.state = VM.State.Uninitialized;
        vm.inbox = Value.hashEmptyTuple();

        // Validator options
        vm.escrowRequired = _escrowRequired;
        vm.gracePeriod = _gracePeriod;
        vm.maxExecutionSteps = _maxExecutionSteps;
    }

    function currentDeposit(address validator) external view returns(uint256) {
        return validatorBalances[validator];
    }

    function escrowRequired() external view returns(uint256) {
        return vm.escrowRequired;
    }

    function getState() external view returns(VM.State) {
        return vm.state;
    }

    function activateVM() external onlyOwner {
        if (vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function ownerShutdown() external onlyOwner {
        _shutdown();
    }

    function completeChallenge(address[2] calldata _players, uint128[2] calldata _rewards) external {
        require(
            msg.sender == address(vm.activeChallengeManager),
            "Only challenge manager can complete challenge"
        );

        vm.activeChallengeManager = address(0);
        validatorBalances[_players[0]] = validatorBalances[_players[0]].add(_rewards[0]);
        validatorBalances[_players[1]] = validatorBalances[_players[1]].add(_rewards[1]);
    }

    function assert(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        bytes32 _afterHash,
        bytes32 _messagesAccHash,
        bytes32 _logsAccHash,
        uint32 _numSteps,
        uint64[2] memory _timeBounds
    )
        public
    {
        require(
            vm.escrowRequired <= validatorBalances[msg.sender],
            "Validator does not have required escrow to assert"
        );
        validatorBalances[msg.sender] -= vm.escrowRequired;

        Disputable.pendingDisputableAssert(
            vm,
            _beforeHash,
            _beforeInbox,
            _afterHash,
            _messagesAccHash,
            _logsAccHash,
            _numSteps,
            _timeBounds
        );
    }

    function confirmDisputableAsserted(
        bytes32 _preconditionHash,
        bytes32 _afterHash,
        uint32 _numSteps,
        bytes memory _messages,
        bytes32 _logsAccHash
    )
        public
    {
        Disputable.confirmDisputableAsserted(
            vm,
            _preconditionHash,
            _afterHash,
            _numSteps,
            _messages,
            _logsAccHash
        );

        validatorBalances[vm.asserter] = validatorBalances[vm.asserter].add(vm.escrowRequired);

        _completeAssertion(_messages);
    }

    function initiateChallenge(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _assertionHash
    )
        public
    {
        require(
            vm.escrowRequired <= validatorBalances[msg.sender],
            "Challenger did not have enough escrowed"
        );
        validatorBalances[msg.sender] -= vm.escrowRequired;

        Disputable.initiateChallenge(
            vm,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );

        vm.activeChallengeManager = challengeFactory.createChallenge(
            [vm.asserter, msg.sender],
            [vm.escrowRequired, vm.escrowRequired],
            vm.gracePeriod,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );

        emit ChallengeLaunched(vm.activeChallengeManager, msg.sender);
    }

    function _completeAssertion(bytes memory _messages) internal {
        bytes32 pending = globalInbox.pullPendingMessages();
        if (pending != Value.hashEmptyTuple()) {
            vm.inbox = Value.hashTuple([
                Value.newInt(1),
                Value.newHashOnly(vm.inbox),
                Value.newHashOnly(pending)
            ]);
        }

        globalInbox.sendMessages(_messages);
    }

    function _shutdown() private {
        // TODO: transfer all owned funds to halt address
        selfdestruct(owner);
    }
}
