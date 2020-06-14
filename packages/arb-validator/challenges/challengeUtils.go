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

package challenges

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
)

func challengeEnded(state ChallengeState, err error) bool {
	if err != nil || state != ChallengeContinuing {
		return true
	} else {
		return false
	}
}

func getVmInboxSegments(
	vmInbox *structures.VMInbox,
	bisectionEvent arbbridge.MessagesBisectionEvent,
	startInbox uint64,
) ([]value.HashPreImage, error) {
	bisectionLength := bisectionEvent.TotalLength.Uint64()

	vmInboxSegments, err := vmInbox.GenerateBisection(
		startInbox,
		uint64(len(bisectionEvent.SegmentHashes))-1,
		bisectionLength)
	if err != nil {
		return nil, err
	}

	return vmInboxSegments, nil
}

func getInboxSegments(
	inbox *structures.MessageStack,
	bisectionEvent arbbridge.MessagesBisectionEvent,
) ([]common.Hash, error) {
	bisectionLength := bisectionEvent.TotalLength.Uint64()

	inboxSegments, err := inbox.GenerateBisection(
		bisectionEvent.ChainHashes[0],
		uint64(len(bisectionEvent.ChainHashes))-1,
		bisectionLength)
	if err != nil {
		return nil, err
	} else {
		return inboxSegments, nil
	}
}

func getSegments(
	inbox *structures.MessageStack,
	bisectionEvent arbbridge.InboxTopBisectionEvent,
) ([]common.Hash, error) {
	bisectionLength := bisectionEvent.TotalLength.Uint64()

	inboxSegments, err := inbox.GenerateBisection(
		bisectionEvent.ChainHashes[0],
		uint64(len(bisectionEvent.ChainHashes))-1,
		bisectionLength)
	if err != nil {
		return nil, err
	} else {
		return inboxSegments, nil
	}
}

func findSegmentToChallenge(
	validatorHashes []common.Hash,
	chainHashes []common.Hash) (uint64, bool) {
	// If any inbox segment is wrong, we can easily win
	for i := uint64(1); i < uint64(len(validatorHashes)); i++ {
		if validatorHashes[i] != chainHashes[i] {
			return i - 1, true
		}
	}

	return 0, false
}

func getMsgStack1() *structures.MessageStack {
	msg1 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(532)),
			Timestamp: big.NewInt(5435254),
		},
	}
	msg2 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(532)),
			Timestamp: big.NewInt(5435254),
		},
	}
	msg3 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(532)),
			Timestamp: big.NewInt(5435254),
		},
	}
	msg4 := message.Received{
		Message: message.Eth{
			To:    common.Address{},
			From:  common.Address{},
			Value: big.NewInt(6745),
		},
		ChainTime: message.ChainTime{
			BlockNum:  common.NewTimeBlocks(big.NewInt(532)),
			Timestamp: big.NewInt(5435254),
		},
	}
	messageStack := structures.NewMessageStack()
	messageStack.DeliverMessage(msg1)
	messageStack.DeliverMessage(msg2)
	messageStack.DeliverMessage(msg3)
	messageStack.DeliverMessage(msg4)

	return messageStack
}

func getMsgStack2() *structures.MessageStack {
	messageStack := structures.NewMessageStack()
	for i := int64(0); i < 8; i++ {
		messageStack.DeliverMessage(message.Received{
			Message: message.Eth{
				To:    common.Address{},
				From:  common.Address{},
				Value: big.NewInt(6745),
			},
			ChainTime: message.ChainTime{
				BlockNum:  common.NewTimeBlocks(big.NewInt(532)),
				Timestamp: big.NewInt(5435254),
			},
		})
	}
	return messageStack
}
