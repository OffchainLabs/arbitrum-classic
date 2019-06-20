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

package challenge

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/core"
	"github.com/offchainlabs/arb-validator/ethbridge"

	"github.com/offchainlabs/arb-avm/protocol"
)

func NewDefender(core *core.Config, assDef protocol.AssertionDefender, time uint64, bridge bridge.Bridge) (State, error) {
	deadline := time + core.VMConfig.GracePeriod
	if assDef.GetAssertion().NumSteps == 1 {
		fmt.Println("Generating proof")
		var buf bytes.Buffer
		if err := assDef.SolidityOneStepProof(&buf); err != nil {
			return nil, &Error{err, "AssertAndDefendBot: error generating one-step proof"}
		}
		bridge.OneStepProof(
			assDef.GetPrecondition(),
			assDef.GetAssertion(),
			buf.Bytes(),
			deadline,
		)
		return oneStepChallengedAssertDefender{
				Config:       core,
				precondition: assDef.GetPrecondition(),
				assertion:    assDef.GetAssertion().Stub(),
				deadline:     deadline,
			}, nil
	}

	defenders := assDef.NBisect(6)
	assertions := make([]*protocol.Assertion, 0, len(defenders))
	for _, defender := range defenders {
		assertions = append(assertions, defender.GetAssertion())
	}
	bridge.BisectAssertion(
		assDef.GetPrecondition(),
		assertions,
		deadline,
	)
	return bisectedAssertDefender{
			Config:            core,
			wholePrecondition: assDef.GetPrecondition(),
			wholeAssertion:    assDef.GetAssertion().Stub(),
			splitDefenders:    defenders,
			deadline:          deadline,
		}, nil
}

type bisectedAssertDefender struct {
	*core.Config
	wholePrecondition *protocol.Precondition
	wholeAssertion    *protocol.AssertionStub
	splitDefenders    []protocol.AssertionDefender
	deadline          uint64
}

func (bot bisectedAssertDefender) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.wholePrecondition,
	//	bot.wholeAssertion,
	//}
	return timedOutAsserterDefender{bot.Config}, nil
}

func (bot bisectedAssertDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.BisectionEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingBisectedDefender{
			bot.Config,
			bot.splitDefenders,
			deadline,
		}, nil
	default:
		return nil, &Error{nil, "ERROR: bisectedAssertDefender: VM state got unsynchronized"}
	}
}

type waitingBisectedDefender struct {
	*core.Config
	defenders []protocol.AssertionDefender
	deadline  uint64
}

func (bot waitingBisectedDefender) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	preconditions := make([]*protocol.Precondition, 0, len(bot.defenders))
	assertions := make([]*protocol.AssertionStub, 0, len(bot.defenders))
	for _, defender := range bot.defenders {
		preconditions = append(preconditions, defender.GetPrecondition())
		assertions = append(assertions, defender.GetAssertion().Stub())
	}
	bridge.TimeoutChallenger(
		preconditions,
		assertions,
		bot.deadline,
	)
	return timedOutChallengerDefender{bot.Config}, nil
}

func (bot waitingBisectedDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		if int(ev.ChallengedAssertion) >= len(bot.defenders) {
			return nil, errors.New("ChallengedAssertion number is out of bounds")
		}
		return NewDefender(bot.Config, bot.defenders[ev.ChallengedAssertion], time, bridge)
	default:
		return nil, &Error{nil, fmt.Sprintf("ERROR: waitingBisectedDefender: VM state got unsynchronized, %T", ev)}
	}
}

type oneStepChallengedAssertDefender struct {
	*core.Config
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot oneStepChallengedAssertDefender) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.Assertion,
	//}
	return timedOutAsserterDefender{bot.Config}, nil
}

func (bot oneStepChallengedAssertDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.OneStepProofEvent:
		fmt.Println("oneStepChallengedAssertDefender: Proof was accepted")
		return nil, nil
	default:
		return nil, &Error{nil, fmt.Sprintf("ERROR: oneStepChallengedAssertDefender: VM state got unsynchronized, %T", ev)}
	}
}

type timedOutChallengerDefender struct {
	*core.Config
}

func (bot timedOutChallengerDefender) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot timedOutChallengerDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: timedOutChallengerDefender: VM state got unsynchronized"}
	}
}

type timedOutAsserterDefender struct {
	*core.Config
}

func (bot timedOutAsserterDefender) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot timedOutAsserterDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: timedOutAsserterDefender: VM state got unsynchronized"}
	}
}
