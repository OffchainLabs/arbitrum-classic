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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	errors2 "github.com/pkg/errors"
)

func DefendInboxTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	afterInboxTop common.Hash,
	messageCount *big.Int,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewInboxTopChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewInboxTopChallenge(address)
	if err != nil {
		return 0, err
	}
	log.Println("=======> defending inbox top claim")

	return defendInboxTop(
		reorgCtx,
		eventChan,
		contract,
		client,
		inbox,
		afterInboxTop,
		messageCount.Uint64(),
		bisectionCount,
	)
}

func ChallengeInboxTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewInboxTopChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewInboxTopChallenge(address)
	if err != nil {
		return 0, err
	}
	log.Println("=======> challenging inbox top claim")
	return challengeInboxTop(
		reorgCtx,
		eventChan,
		contract,
		client,
		inbox,
		challengeEverything,
	)
}

func defendInboxTop(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	afterInboxTop common.Hash,
	messageCount uint64,
	bisectionCount uint64,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("InboxTopChallenge defender expected InitiateChallengeEvent but got %T", event)
	}

	startState := afterInboxTop

	for {
		if messageCount == 1 {
			timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
			if timedOut {
				msg, err := inbox.GenerateOneStepProof(startState)
				if err != nil {
					return 0, err
				}
				err = contract.OneStepProof(ctx, startState, msg.CommitmentHash())
				if err != nil {
					return 0, errors2.Wrap(err, "Error making one step proof")
				}
				event, state, err = getNextEvent(eventChan)
			}

			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("InboxTopChallenge defender expected OneStepProof but got %T", event)
			}
			return ChallengeAsserterWon, nil
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			chainHashes, err := inbox.GenerateBisection(startState, bisectionCount, messageCount)
			if err != nil {
				return 0, err
			}
			err = contract.Bisect(ctx, chainHashes, new(big.Int).SetUint64(messageCount))
			if err != nil {
				return 0, errors2.Wrap(err, "Error bisecting")
			}
			event, state, err = getNextEvent(eventChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := event.(arbbridge.InboxTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("InboxTopChallenge defender expected InboxTopBisectionEvent but got %T", event)
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
			return 0, fmt.Errorf("InboxTopChallenge defender expected ContinueChallengeEvent but got %T", event)
		}
		startState = ev.ChainHashes[contEv.SegmentIndex.Uint64()]
		messageCount = getSegmentCount(messageCount, uint64(len(ev.ChainHashes))-1, contEv.SegmentIndex.Uint64())
	}
}

func challengeInboxTop(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	challengeEverything bool,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("InboxTopChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

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

		ev, ok := event.(arbbridge.InboxTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("InboxTopChallenge challenger expected InboxTopBisectionEvent but got %T", event)
		}

		// Wait to check if we've already chosen a segment
		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			err = nil
			segments, err := inbox.GenerateBisection(ev.ChainHashes[0], uint64(len(ev.ChainHashes))-1, ev.TotalLength.Uint64())
			segmentToChallenge, found := func() (uint64, bool) {
				for i := uint64(1); i < uint64(len(segments)); i++ {
					if segments[i] != ev.ChainHashes[i] {
						return i - 1, true
					}
				}
				return 0, false
			}()
			if !found {
				if challengeEverything {
					segmentToChallenge = uint64(rand.Int31n(int32(len(ev.ChainHashes) - 1)))
				} else {
					return 0, errors.New("can't find inbox segment to challenge")
				}
			}
			err = contract.ChooseSegment(ctx, uint16(segmentToChallenge), ev.ChainHashes, ev.TotalLength.Uint64())
			if err != nil {
				return 0, err
			}
			event, state, err = getNextEvent(eventChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("InboxTopChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = contEv.Deadline
	}
}
