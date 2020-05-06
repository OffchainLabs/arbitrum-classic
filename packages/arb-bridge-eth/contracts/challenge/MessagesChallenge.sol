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

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/Protocol.sol";
import "../Messages.sol";


contract MessagesChallenge is BisectionChallenge {

    event Bisected(
        bytes32[] chainHashes,
        bytes32[] segmentHashes,
        uint256 totalLength,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string constant HS_BIS_INPLEN = "HS_BIS_INPLEN";

    function bisect(
        bytes32[] memory _chainHashes,
        bytes32[] memory _segmentHashes,
        uint256 _chainLength
    )
        public
        asserterAction
    {
        uint256 bisectionCount = _chainHashes.length - 1;
        require(bisectionCount + 1 == _segmentHashes.length, HS_BIS_INPLEN);

        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                _chainHashes[0],
                _chainHashes[bisectionCount],
                _segmentHashes[0],
                _segmentHashes[bisectionCount],
                _chainLength
            )
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.messagesHash(
            _chainHashes[0],
            _chainHashes[1],
            _segmentHashes[0],
            _segmentHashes[1],
            firstSegmentSize(_chainLength, bisectionCount)
        );
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeUtils.messagesHash(
                _chainHashes[i],
                _chainHashes[i + 1],
                _segmentHashes[i],
                _segmentHashes[i + 1],
                otherSegmentSize(_chainLength, bisectionCount)
            );
        }

        commitToSegment(hashes);
        asserterResponded();
        emit Bisected(
            _chainHashes,
            _segmentHashes,
            _chainLength,
            deadlineTicks
        );
    }

    function oneStepProofTransactionMessage(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        address _chain,
        address _to,
        address _from,
        uint256 _seqNumber,
        uint256 _value,
        bytes memory _data,
        uint256 _blockNumber,
        uint256 _timestamp
    )
        public
        asserterAction
    {

        bytes32 messageHash = Messages.transactionHash(
            _chain,
            _to,
            _from,
            _seqNumber,
            _value,
            keccak256(_data),
            _blockNumber,
            _timestamp
        );
        bytes32 arbMessageHash = Messages.transactionMessageHash(
            _chain,
            _to,
            _from,
            _seqNumber,
            _value,
            _data,
            _blockNumber,
            _timestamp
        );

        oneStepProof(
            _lowerHashA,
            _lowerHashB,
            messageHash,
            arbMessageHash
        );
    }

    function oneStepProofEthMessage(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        address _to,
        address _from,
        uint256 _value,
        uint256 _blockNumber,
        uint256 _timestamp,
        uint256 _messageNum
    )
        public
        asserterAction
    {

        bytes32 messageHash = Messages.ethHash(
            _to,
            _from,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );
        bytes32 arbMessageHash = Messages.ethMessageHash(
            _to,
            _from,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );

        oneStepProof(
            _lowerHashA,
            _lowerHashB,
            messageHash,
            arbMessageHash
        );
    }

    function oneStepProofERC20Message(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        address _to,
        address _from,
        address _erc20,
        uint256 _value,
        uint256 _blockNumber,
        uint256 _timestamp,
        uint256 _messageNum
    )
        public
        asserterAction
    {

        bytes32 messageHash = Messages.erc20Hash(
            _to,
            _from,
            _erc20,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );
        bytes32 arbMessageHash = Messages.erc20MessageHash(
            _to,
            _from,
            _erc20,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );

        oneStepProof(
            _lowerHashA,
            _lowerHashB,
            messageHash,
            arbMessageHash
        );
    }

    function oneStepProofERC721Message(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        address _to,
        address _from,
        address _erc721,
        uint256 _value,
        uint256 _blockNumber,
        uint256 _timestamp,
        uint256 _messageNum
    )
        public
        asserterAction
    {

        bytes32 messageHash = Messages.erc721Hash(
            _to,
            _from,
            _erc721,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );
        bytes32 arbMessageHash = Messages.erc721MessageHash(
            _to,
            _from,
            _erc721,
            _value,
            _blockNumber,
            _timestamp,
            _messageNum
        );

        oneStepProof(
            _lowerHashA,
            _lowerHashB,
            messageHash,
            arbMessageHash
        );
    }

    function oneStepProofContractTransactionMessage(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        address _to,
        address _from,
        uint256 _value,
        bytes memory _data,
        uint256 _blockNumber,
        uint256 _timestamp,
        uint256 _messageNum
    )
        public
        asserterAction
    {

        bytes32 messageHash = Messages.contractTransactionHash(
            _to,
            _from,
            _value,
            _data,
            _blockNumber,
            _timestamp,
            _messageNum
        );
        bytes32 arbMessageHash = Messages.contractTransactionMessageHash(
            _to,
            _from,
            _value,
            _data,
            _blockNumber,
            _timestamp,
            _messageNum
        );

        oneStepProof(
            _lowerHashA,
            _lowerHashB,
            messageHash,
            arbMessageHash
        );
    }

    function oneStepProofTransactionBatchMessage(
        bytes32 lowerHashA,
        bytes32 lowerHashB,
        address chain,
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    )
        public
        asserterAction
    {
        bytes32 messageHash = Messages.transactionBatchHash(
            transactions,
            blockNum,
            blockTimestamp
        );

        bytes32 afterInboxHash = Messages.transactionMessageBatchHash(
            lowerHashB,
            chain,
            transactions,
            blockNum,
            blockTimestamp
        );

        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                lowerHashA,
                Protocol.addMessageToInbox(lowerHashA, messageHash),
                lowerHashB,
                afterInboxHash,
                1
            )
        );
        finishOneStepProof();
    }

    function oneStepProof(
        bytes32 _lowerHashA,
        bytes32 _lowerHashB,
        bytes32 _valueHashA,
        bytes32 _valueHashB
    )
        private
    {
        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                _lowerHashA,
                Protocol.addMessageToInbox(_lowerHashA, _valueHashA),
                _lowerHashB,
                Protocol.addMessageToVMInbox(_lowerHashB, _valueHashB),
                1
            )
        );

        finishOneStepProof();
    }

    function finishOneStepProof() private {
        emit OneStepProofCompleted();
        _asserterWin();
    }

    function resolveChallengeAsserterWon() internal {
        IStaking(vmAddress).resolveChallenge(asserter, challenger, INVALID_MESSAGES_TYPE);
    }

    function resolveChallengeChallengerWon() internal {
        IStaking(vmAddress).resolveChallenge(challenger, asserter, INVALID_MESSAGES_TYPE);
    }
}
