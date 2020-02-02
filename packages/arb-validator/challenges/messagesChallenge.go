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
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/message"
	errors2 "github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendMessagesClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockID *structures.BlockID,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	messageCount *big.Int,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockID, startLogIndex, contractWatcher)

	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}

	return defendMessages(
		reorgCtx,
		eventChan,
		contract,
		client,
		pendingInbox,
		beforePending,
		messageCount.Uint64(),
		bisectionCount,
	)
}

func ChallengeMessagesClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockID *structures.BlockID,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	messageCount *big.Int,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockID, startLogIndex, contractWatcher)

	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}

	return challengeMessages(
		reorgCtx,
		eventChan,
		contract,
		client,
		pendingInbox,
		beforePending,
		messageCount.Uint64(),
		challengeEverything,
	)
}

func defendMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	messageCount uint64,
	bisectionCount uint64,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, errNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge defender expected InitiateChallengeEvent but got %T", event)
	}

	inbox, err := pendingInbox.GenerateInbox(beforePending, messageCount)
	if err != nil {
		return 0, err
	}

	log.Println("Pending inbox", pendingInbox)
	log.Println("Full inbox", inbox)

	startPending := beforePending
	startMessages := value.NewEmptyTuple().Hash()
	inboxStartCount := uint64(0)

	for {
		log.Println(inboxStartCount, messageCount)
		if messageCount == 1 {
			timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
			if err != nil {
				return 0, err
			}
			if timedOut {
				msg, err := pendingInbox.GenerateOneStepProof(startPending)
				if err != nil {
					return 0, err
				}

				log.Println("OneStepProofEthMessage", startPending, startMessages)

				log.Println("pending after", hashing.SoliditySHA3(hashing.Bytes32(startPending), hashing.Bytes32(msg.CommitmentHash())))
				log.Println("inbox after", value.NewTuple2(value.NewHashOnlyValue(startMessages, 1), message.DeliveredValue(msg)).Hash())

				switch msg := msg.(type) {
				case message.DeliveredTransaction:
					err = contract.OneStepProofTransactionMessage(ctx, startPending, startMessages, msg)
				case message.DeliveredEth:
					err = contract.OneStepProofEthMessage(ctx, startPending, startMessages, msg)
				case message.DeliveredERC20:
					err = contract.OneStepProofERC20Message(ctx, startPending, startMessages, msg)
				case message.DeliveredERC721:
					err = contract.OneStepProofERC721Message(ctx, startPending, startMessages, msg)
				}
				if err != nil {
					return 0, errors2.Wrap(err, "failing making one step proof")
				}
				event, state, err = getNextEvent(eventChan)
				if err != nil {
					return 0, err
				}
			}

			if state != ChallengeContinuing {
				return state, err
			}
			_, ok = event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("MessagesChallenge defender expected OneStepProof but got %T", event)
			}
			return ChallengeAsserterWon, nil
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if err != nil {
			return 0, err
		}
		if timedOut {
			chainHashes, err := pendingInbox.GenerateBisection(startPending, bisectionCount, messageCount)
			if err != nil {
				return 0, err
			}
			inboxHashes, err := inbox.GenerateBisection(inboxStartCount, bisectionCount, messageCount)
			if err != nil {
				return 0, err
			}

			log.Println("chainHashes", chainHashes)
			log.Println("inboxHashes", inboxHashes)

			err = contract.Bisect(ctx, chainHashes, inboxHashes, new(big.Int).SetUint64(messageCount))
			if err != nil {
				return 0, errors2.Wrap(err, "failing making bisection")
			}

			event, state, err = getNextEvent(eventChan)
			if err != nil {
				return 0, err
			}
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected MessagesBisectionEvent but got %T", event)
		}

		event, state, err = getNextEventWithTimeout(
			ctx,
			eventChan,
			ev.Deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected ContinueChallengeEvent but got %T", event)
		}
		startPending = ev.ChainHashes[contEv.SegmentIndex.Uint64()]
		startMessages = ev.SegmentHashes[contEv.SegmentIndex.Uint64()]
		inboxStartCount += getSegmentStart(messageCount, uint64(len(ev.ChainHashes))-1, contEv.SegmentIndex.Uint64())
		log.Println("messageCount", messageCount, uint64(len(ev.ChainHashes))-1, contEv.SegmentIndex.Uint64())
		messageCount = getSegmentCount(messageCount, uint64(len(ev.ChainHashes))-1, contEv.SegmentIndex.Uint64())
	}
}

func challengeMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	messageCount uint64,
	challengeEverything bool,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, errNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

	inbox, err := pendingInbox.GenerateInbox(beforePending, messageCount)
	if err != nil {
		return 0, err
	}

	startInbox := uint64(0)

	deadline := ev.Deadline
	for {
		event, state, err := getNextEventWithTimeout(
			ctx,
			eventChan,
			deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}

		if _, ok := event.(arbbridge.OneStepProofEvent); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("messagesChallenge challenger expected MessagesBisectionEvent but got %T", event)
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if err != nil {
			return 0, err
		}
		if timedOut {
			pendingSegments, err := pendingInbox.GenerateBisection(ev.ChainHashes[0], uint64(len(ev.ChainHashes))-1, ev.TotalLength.Uint64())
			if err != nil {
				return 0, err
			}

			inboxSegments, err := inbox.GenerateBisection(startInbox, uint64(len(ev.SegmentHashes))-1, ev.TotalLength.Uint64())
			if err != nil {
				return 0, err
			}

			segmentToChallenge, found := func() (uint64, bool) {
				// If any pending inbox segment is wrong, we can easily win
				for i := uint64(1); i < uint64(len(pendingSegments)); i++ {
					if pendingSegments[i] != ev.ChainHashes[i] {
						return i - 1, true
					}
				}

				for i := uint64(1); i < uint64(len(inboxSegments)); i++ {
					if inboxSegments[i] != ev.SegmentHashes[i] {
						return i - 1, true
					}
				}
				return 0, false
			}()

			if !found {
				if challengeEverything {
					segmentToChallenge = uint64(rand.Int31n(int32(len(ev.ChainHashes) - 1)))
				} else {
					return 0, errors.New("nothing to challenge")
				}
			}
			log.Println("ChooseSegment", uint16(segmentToChallenge), ev.ChainHashes, ev.SegmentHashes, ev.TotalLength)
			err = contract.ChooseSegment(ctx, uint16(segmentToChallenge), ev.ChainHashes, ev.SegmentHashes, ev.TotalLength)
			if err != nil {
				return 0, err
			}
			event, state, err = getNextEvent(eventChan)
			if err != nil {
				return 0, err
			}
		}
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = contEv.Deadline
	}
}
