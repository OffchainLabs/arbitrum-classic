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
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type messagesChallenge struct {
	*bisectionChallenge
}

func newMessagesChallenge(address common.Address, client *GoArbAuthClient) (*messagesChallenge, error) {
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

	bisectionCount := len(chainHashes) - 1
	fmt.Println("bisectionCount", bisectionCount)
	fmt.Println("len(segmentHashes)", len(segmentHashes))
	if bisectionCount+1 != len(segmentHashes) {
		return errors.New("Incorrect previous state")
	}
	//	uint256 bisectionCount = _chainHashes.length - 1;
	//	require(bisectionCount + 1 == _segmentHashes.length, HS_BIS_INPLEN);
	//
	//	requireMatchesPrevState(
	//		ChallengeUtils.messagesHash(
	//			_chainHashes[0],
	//			_chainHashes[bisectionCount],
	//			_segmentHashes[0],
	//			_segmentHashes[bisectionCount],
	//			_chainLength
	//	)
	//);

	fmt.Println("chainHashes[0]", chainHashes[0])
	fmt.Println("chainHashes[bisectionCount]", chainHashes[bisectionCount])
	fmt.Println("segmentHashes[0]", segmentHashes[0])
	fmt.Println("segmentHashes[bisectionCount]", segmentHashes[bisectionCount])
	fmt.Println("chainLength", chainLength)
	msgHash := structures.MessageChallengeDataHash(chainHashes[0], chainHashes[bisectionCount], segmentHashes[0], segmentHashes[bisectionCount], chainLength)
	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(msgHash) {
		return errors.New("Incorrect previous state")
	}

	//
	//	bytes32[] memory hashes = new bytes32[](bisectionCount);
	//	hashes[0] = ChallengeUtils.messagesHash(
	//		_chainHashes[0],
	//		_chainHashes[1],
	//		_segmentHashes[0],
	//		_segmentHashes[1],
	//		firstSegmentSize(_chainLength, bisectionCount)
	//	);
	//	for (uint256 i = 1; i < bisectionCount; i++) {
	//		hashes[i] = ChallengeUtils.messagesHash(
	//			_chainHashes[i],
	//			_chainHashes[i + 1],
	//			_segmentHashes[i],
	//			_segmentHashes[i + 1],
	//			otherSegmentSize(_chainLength, bisectionCount)
	//		);
	//	}

	hashes := make([][32]byte, 0, bisectionCount)
	hashes = append(hashes, structures.MessageChallengeDataHash(
		chainHashes[0],
		chainHashes[1],
		segmentHashes[0],
		segmentHashes[1],
		new(big.Int).Add(new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount))), new(big.Int).Mod(chainLength, big.NewInt(int64(bisectionCount)))),
	))
	for i := 1; i < bisectionCount; i++ {
		hashes = append(hashes, structures.MessageChallengeDataHash(
			chainHashes[i],
			chainHashes[i+1],
			segmentHashes[i],
			segmentHashes[i+1],
			new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount)))))
	}

	c.commitToSegment(hashes)
	c.asserterResponded()
	//	emit Bisected(
	//		_chainHashes,
	//		_segmentHashes,
	//		_chainLength,
	//		deadlineTicks
	//	);

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.MessagesBisectionEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
			ChainHashes:   chainHashes,
			SegmentHashes: segmentHashes,
			TotalLength:   chainLength,
			Deadline:      c.client.GoEthClient.challenges[c.contractAddress].deadline,
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

	messageHash := msg.CommitmentHash()

	msgType := msg.AsValue()
	ethMsg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(msg.BlockNum.AsInt()),
		value.NewIntValue(msg.MessageNum),
		msgType,
	})
	ethMsgHash := ethMsg.Hash()

	msgs, _ := value.NewTupleFromSlice([]value.Value{
		value.NewHashOnlyValue(lowerHashB, 32),
		value.NewHashOnlyValue(ethMsgHash, 32),
	})
	matchHash := structures.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("Incorrect previous state")
	}

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	// TODO: handle stake distribution
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
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.ChallengeCompletedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
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
	messageHash := msg.CommitmentHash()

	msgType := msg.AsValue()
	ethMsg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(msg.BlockNum.AsInt()),
		value.NewIntValue(msg.MessageNum),
		msgType,
	})
	ethMsgHash := ethMsg.Hash()

	msgs, _ := value.NewTupleFromSlice([]value.Value{
		value.NewHashOnlyValue(lowerHashB, 32),
		value.NewHashOnlyValue(ethMsgHash, 32),
	})
	matchHash := structures.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("Incorrect previous state")
	}

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	// TODO: handle stake distribution
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
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.ChallengeCompletedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
			Winner:            msg.From,
			Loser:             msg.To,
			ChallengeContract: c.contractAddress,
		},
	})
	return nil
}

func (c *messagesChallenge) OneStepProofERC721Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC721,
) error {
	fmt.Println("in messagesChallenge OneStepProofERC721Message")
	messageHash := msg.CommitmentHash()

	msgType := msg.AsValue()
	ethMsg, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(msg.BlockNum.AsInt()),
		value.NewIntValue(msg.MessageNum),
		msgType,
	})
	ethMsgHash := ethMsg.Hash()

	msgs, _ := value.NewTupleFromSlice([]value.Value{
		value.NewHashOnlyValue(lowerHashB, 32),
		value.NewHashOnlyValue(ethMsgHash, 32),
	})
	matchHash := structures.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("Incorrect previous state")
	}

	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	// TODO: handle stake distribution
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
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.ChallengeCompletedEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
			Winner:            msg.From,
			Loser:             msg.To,
			ChallengeContract: c.contractAddress,
		},
	})
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
