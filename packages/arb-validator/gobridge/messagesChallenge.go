/*
 * Copyright 2020, Offchain Labs, Inc.
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

package gobridge

import (
	"context"
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type messagesChallenge struct {
	*bisectionChallenge
}

func newMessagesChallenge(address common.Address, client *MockArbAuthClient) (*messagesChallenge, error) {
	fmt.Println("in messagesChallenge newMessagesChallenge")
	bisectionChallenge, err := newBisectionChallenge(address, client) //, auth??
	if err != nil {
		return nil, err
	}
	vm := &messagesChallenge{bisectionChallenge: bisectionChallenge}
	return vm, err
	//messagesContract, err := messageschallenge.NewMessagesChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to messagesChallenge")
	//}
	//return &messagesChallenge{bisectionChallenge: bisectionChallenge}, nil
}

func (vm *messagesChallenge) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	fmt.Println("in messagesChallenge GetEvents")
	return nil, nil
}

func (c *messagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	fmt.Println("in messagesChallenge Bisect")
	//c.auth.Context = ctx
	//tx, err := c.challenge.Bisect(
	//	c.auth,
	//	chainHashes,
	//	segmentHashes,
	//	chainLength,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "Bisect")
	c.client.MockEthClient.pubMsg(arbbridge.MaybeEvent{
		Event: arbbridge.MessagesBisectionEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.MockEthClient.getCurrentBlock(),
			},
			ChainHashes:   chainHashes,
			SegmentHashes: segmentHashes,
			TotalLength:   chainLength,
			Deadline:      c.client.MockEthClient.challenges[c.contractAddress].deadline,
		},
	})

	return nil
}

func (c *messagesChallenge) OneStepProofTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredTransaction,
) error {
	fmt.Println("in messagesChallenge OneStepProofTransactionMessage")
	return nil
}

func (c *messagesChallenge) OneStepProofEthMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredEth,
) error {
	fmt.Println("in messagesChallenge OneStepProofEthMessage")
	//bytes32 messageHash = Messages.ethHash(
	//	_to,
	//	_from,
	//	_value,
	//	_blockNumber,
	//	_messageNum
	//);
	//		return keccak256(
	//			abi.encodePacked(
	//				ETH_DEPOSIT,
	//				to,
	//				from,
	//				value,
	//				blockNumber,
	//				messageNum
	//		)
	//	);
	//bytes32 arbMessageHash = Messages.ethMessageHash(
	//	_to,
	//	_from,
	//	_value,
	//	_blockNumber,
	//	_messageNum
	//);
	//		Value.Data[] memory msgValues = new Value.Data[](2);
	//		msgValues[0] = Value.newInt(uint256(to));
	//		msgValues[1] = Value.newInt(value);
	//
	//		Value.Data[] memory msgType = new Value.Data[](3);
	//		msgType[0] = Value.newInt(ETH_DEPOSIT);
	//		msgType[1] = Value.newInt(uint256(from));
	//		msgType[2] = Value.newTuple(msgValues);
	//
	//		Value.Data[] memory ethMsg = new Value.Data[](3);
	//		ethMsg[0] = Value.newInt(blockNumber);
	//		ethMsg[1] = Value.newInt(messageNum);
	//		ethMsg[2] = Value.newTuple(msgType);
	//
	//		return Value.newTuple(ethMsg).hash().hash;
	//
	//oneStepProof(
	//	_lowerHashA,
	//	_lowerHashB,
	//	messageHash,
	//	arbMessageHash
	//);
	//	requireMatchesPrevState(
	//		ChallengeUtils.messagesHash(
	//			_lowerHashA,
	//			Protocol.addMessageToPending(_lowerHashA, _valueHashA),
	//			_lowerHashB,
	//			Protocol.addMessageToInbox(_lowerHashB, _valueHashB),
	//			1
	//	)
	//);
	//
	//	emit OneStepProofCompleted();
	c.client.MockEthClient.pubMsg(arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.MockEthClient.getCurrentBlock(),
			},
		},
	})
	//	_asserterWin();
	//		resolveChallengeAsserterWon();
	//			require(challenges[msg.sender], RES_CHAL_SENDER);
	//			delete challenges[msg.sender];
	//
	//			Staker storage winningStaker = getValidStaker(address(winner));
	//			winner.transfer(stakeRequirement / 2);
	//			winningStaker.inChallenge = false;
	//			deleteStaker(loser);
	//
	//			emit RollupChallengeCompleted(msg.sender, address(winner), loser);
	c.client.MockEthClient.pubMsg(arbbridge.MaybeEvent{
		Event: arbbridge.ChallengeCompletedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.MockEthClient.getCurrentBlock(),
			},
			Winner:            msg.From,
			Loser:             msg.To,
			ChallengeContract: c.contractAddress,
		},
	})

	//		selfdestruct(msg.sender);

	return nil
}

func (c *messagesChallenge) OneStepProofERC20Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC20,
) error {
	fmt.Println("in messagesChallenge OneStepProofERC20Message")
	return nil
}

func (c *messagesChallenge) OneStepProofERC721Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC721,
) error {
	fmt.Println("in messagesChallenge OneStepProofERC721Message")
	return nil
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	fmt.Println("in messagesChallenge ChooseSegment")
	bisectionCount := uint64(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint64(0); i < bisectionCount; i++ {
		stepCount := structures.CalculateBisectionStepCount(i, bisectionCount, chainLength.Uint64())
		bisectionHashes = append(
			bisectionHashes,
			structures.MessageChallengeDataHash(
				chainHashes[i],
				chainHashes[i+1],
				segmentHashes[i],
				segmentHashes[i+1],
				new(big.Int).SetUint64(stepCount),
			),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}
