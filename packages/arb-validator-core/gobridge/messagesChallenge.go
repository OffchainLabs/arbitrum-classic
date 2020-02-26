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
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type messagesChallenge struct {
	*bisectionChallenge
}

func newMessagesChallenge(address common.Address, client *GoArbAuthClient) (*messagesChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client) //, auth??
	if err != nil {
		return nil, err
	}
	vm := &messagesChallenge{bisectionChallenge: bisectionChallenge}
	return vm, err
}

func (c *messagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {

	bisectionCount := len(chainHashes) - 1
	if bisectionCount+1 != len(segmentHashes) {
		return errors.New("Bisect Incorrect previous state - bisection count")
	}

	msgHash := valprotocol.MessageChallengeDataHash(chainHashes[0], chainHashes[bisectionCount], segmentHashes[0], segmentHashes[bisectionCount], chainLength)
	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(msgHash) {
		return errors.New("Bisect Incorrect previous state msgHash")
	}

	hashes := make([][32]byte, 0, bisectionCount)
	hashes = append(hashes, valprotocol.MessageChallengeDataHash(
		chainHashes[0],
		chainHashes[1],
		segmentHashes[0],
		segmentHashes[1],
		new(big.Int).Add(new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount))), new(big.Int).Mod(chainLength, big.NewInt(int64(bisectionCount)))),
	))
	for i := 1; i < bisectionCount; i++ {
		hashes = append(hashes, valprotocol.MessageChallengeDataHash(
			chainHashes[i],
			chainHashes[i+1],
			segmentHashes[i],
			segmentHashes[i+1],
			new(big.Int).Div(chainLength, big.NewInt(int64(bisectionCount)))))
	}

	c.commitToSegment(hashes)
	c.asserterResponded()

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
	messageHash := msg.CommitmentHash()
	arbMessageHash := message.DeliveredValue(msg).Hash()

	// oneStepProof
	if !c.challenge.challengeData.challengerDataHash.Equals(hashing.SoliditySHA3(
		lowerHashA,
		lowerHashB,
		messageHash,
		arbMessageHash,
	)) {
		return errors.New("OneStepProofTransactionMessage Incorrect previous state")
	}
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	// TODO: challenge resolution
	// resolveChallengeAsserterWon

	return nil
}

func (c *messagesChallenge) OneStepProofEthMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredEth,
) error {
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
	matchHash := valprotocol.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("OneStepProofEthMessage Incorrect previous state")
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

func (c *messagesChallenge) OneStepProofERC20Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredERC20,
) error {
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
	matchHash := valprotocol.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("OneStepProofERC20Message Incorrect previous state")
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
	matchHash := valprotocol.MessageChallengeDataHash(
		lowerHashA,
		hashing.SoliditySHA3(
			hashing.Bytes32(lowerHashA),
			hashing.Bytes32(messageHash)),
		lowerHashB,
		msgs.Hash(),
		big.NewInt(1),
	)

	if !c.client.GoEthClient.challenges[c.contractAddress].challengerDataHash.Equals(matchHash) {
		return errors.New("OneStepProofERC721Message Incorrect previous state")
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

func (c *messagesChallenge) OneStepProofContractTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB common.Hash,
	msg message.DeliveredContractTransaction,
) error {
	messageHash := msg.CommitmentHash()
	txHash := msg.ReceiptHash()
	msgType := msg.AsValue()
	arbMessageHash, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(new(big.Int).Set(msg.BlockNum.AsInt())),
		value.NewIntValue(new(big.Int).SetBytes(txHash[:])),
		msgType,
	})

	if !c.challenge.challengeData.challengerDataHash.Equals(hashing.SoliditySHA3(
		lowerHashA,
		lowerHashB,
		messageHash,
		arbMessageHash,
	)) {
		return errors.New("OneStepProofContractTransactionMessage Incorrect previous state")
	}
	c.client.GoEthClient.pubMsg(c.challengeData, arbbridge.MaybeEvent{
		Event: arbbridge.OneStepProofEvent{
			ChainInfo: arbbridge.ChainInfo{
				BlockId: c.client.GoEthClient.getCurrentBlock(),
			},
		},
	})
	// TODO: challenge resolution
	// resolveChallengeAsserterWon

	return nil
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	bisectionCount := uint64(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint64(0); i < bisectionCount; i++ {
		stepCount := valprotocol.CalculateBisectionStepCount(i, bisectionCount, chainLength.Uint64())
		bisectionHashes = append(
			bisectionHashes,
			valprotocol.MessageChallengeDataHash(
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
