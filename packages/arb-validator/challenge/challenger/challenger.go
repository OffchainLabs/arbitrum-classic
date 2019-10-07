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
	assertion [32]byte,
	numSteps uint32,
	machine machine.Machine,
	deadline uint64,
) challenge.State {
	return waitingContinuing{
		config,
		precondition,
		assertion,
		numSteps,
		machine,
		deadline,
	}
}

type waitingContinuing struct {
	*core.Config
	challengedPrecondition *protocol.Precondition
	challengedAssertion    [32]byte
	numSteps               uint32
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
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		startState := bot.startState.Clone()
		assertionNum, m, err := machine.ChooseAssertionToChallenge(bot.startState, ev.BisectionHashes, bot.challengedPrecondition, bot.numSteps)
		if err != nil && bot.ChallengeEverything {
			assertionNum, m = machine.ChooseRandomAssertionToChallenge(startState, ev.BisectionHashes, bot.challengedPrecondition, bot.numSteps)
		}
		if err != nil {
			return nil, &challenge.Error{Message: "ERROR: waitingContinuing: Critical bug: All segments in false Assertion are valid"}
		}
		_, err = bridge.ContinueChallenge(
			context.Background(),
			assertionNum,
			bot.challengedPrecondition,
			bot.numSteps,
			bot.challengedAssertion,
			ev.BisectionHashes,
		)
		if ev.SpentOutputValues != nil {
			bot.challengedPrecondition.BeforeBalance.RemoveAssertionValues(ev.SpentOutputValues)
		}
		return continuing{
			Config:          bot.Config,
			challengedState: m,
			deadline:        time + bot.VMConfig.GracePeriod,
			precondition:    bot.challengedPrecondition,
			bisections:      ev.BisectionHashes,
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
	bisections      [][32]byte
	numSteps        uint32
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
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		stepCounts := machine.BisectionStepCounts(uint32(len(bot.bisections)+1), bot.numSteps)
		return New(
			bot.Config,
			bot.precondition,
			bot.bisections[ev.ChallengedAssertion],
			stepCounts[ev.ChallengedAssertion],
			bot.challengedState,
			deadline,
		), nil
	default:
		return nil, &challenge.Error{Message: "ERROR: continuing: VM state got unsynchronized"}
	}
}
