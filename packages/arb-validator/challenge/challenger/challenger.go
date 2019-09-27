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

package challenger

import (
	"context"
	"log"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
)

func New(
	config *core.Config,
	precondition *protocol.Precondition,
	machine machine.Machine,
	deadline uint64,
) challenge.State {
	return waitingContinuing{
		config,
		precondition,
		machine,
		deadline,
	}
}

type waitingContinuing struct {
	*core.Config
	challengedPrecondition *protocol.Precondition
	startState             machine.Machine
	deadline               uint64
}

func (bot waitingContinuing) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	_, err := bridge.AsserterTimedOut(
		context.Background(),
	)
	return challenge.TimedOutAsserter{Config: bot.Config}, err
}

func (bot waitingContinuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	log.Printf("waitingContinuing UpdateState event %T", ev)
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		assertionNum, m, err := machine.ChooseAssertionToChallenge(bot.startState, ev.Assertions, bot.challengedPrecondition)
		if err != nil && bot.ChallengeEverything {
			assertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
			m = bot.startState
			for i := uint16(0); i < assertionNum; i++ {
				m.ExecuteAssertion(
					int32(ev.Assertions[i].NumSteps),
					bot.challengedPrecondition.TimeBounds,
				)
			}
			err = nil
		}
		if err != nil {
			return nil, &challenge.Error{Message: "ERROR: waitingContinuing: Critical bug: All segments in false Assertion are valid"}
		}
		_, err = bridge.ContinueChallenge(
			context.Background(),
			assertionNum,
			bot.challengedPrecondition,
			ev.Assertions,
		)
		return continuing{
			Config:          bot.Config,
			challengedState: m,
			deadline:        time + bot.VMConfig.GracePeriod,
			precondition:    bot.challengedPrecondition,
			assertions:      ev.Assertions,
		}, err
	case ethbridge.OneStepProofEvent:
		return nil, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: waitingContinuing: VM state got unsynchronized"}
	}
}

type continuing struct {
	*core.Config
	challengedState machine.Machine
	deadline        uint64
	precondition    *protocol.Precondition
	assertions      []*protocol.AssertionStub
}

func (bot continuing) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	// Currently not sending a timeout valmessage if this validator timed out
	return challenge.TimedOutChallenger{Config: bot.Config}, nil
}

func (bot continuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	log.Printf("continuing UpdateState event %T", ev)
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		preconditions := protocol.GeneratePreconditions(bot.precondition, bot.assertions)
		return waitingContinuing{
			bot.Config,
			preconditions[ev.ChallengedAssertion],
			bot.challengedState,
			deadline,
		}, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: continuing: VM state got unsynchronized"}
	}
}
