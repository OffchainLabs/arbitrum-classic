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

package validator

import (
	"errors"
	"github.com/offchainlabs/arb-validator/valmessage"

	"bytes"
	"fmt"
	"github.com/offchainlabs/arb-avm/protocol"
)

func NewDefendingValidator(core *validatorConfig, assDef protocol.AssertionDefender, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return defenseValidator(core, assDef, time)
}

func defenseValidator(core *validatorConfig, assDef protocol.AssertionDefender, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	deadline := time + core.config.GracePeriod
	if assDef.GetAssertion().NumSteps == 1 {
		fmt.Println("Generating proof")
		var buf bytes.Buffer
		if err := assDef.SolidityOneStepProof(&buf); err != nil {
			return nil, nil, &Error{err, "AssertAndDefendBot: error generating one-step proof"}
		}
		return OneStepChallengedAssertDefender{
				validatorConfig: core,
				precondition:    assDef.GetPrecondition(),
				assertion:       assDef.GetAssertion().Stub(),
				deadline:        deadline,
			}, []valmessage.OutgoingMessage{valmessage.SendOneStepProofMessage{
				Precondition: assDef.GetPrecondition(),
				Assertion:    assDef.GetAssertion(),
				Proof:        buf.Bytes(),
				Deadline:     deadline,
			}}, nil
	} else {
		defenders := assDef.NBisect(6)
		assertions := make([]*protocol.Assertion, 0, len(defenders))
		for _, defender := range defenders {
			assertions = append(assertions, defender.GetAssertion())
		}
		return BisectedAssertDefender{
				validatorConfig:   core,
				wholePrecondition: assDef.GetPrecondition(),
				wholeAssertion:    assDef.GetAssertion().Stub(),
				splitDefenders:    defenders,
				deadline:          deadline}, []valmessage.OutgoingMessage{valmessage.SendBisectionMessage{
				Deadline:     deadline,
				Precondition: assDef.GetPrecondition(),
				Assertions:   assertions,
			}}, nil
	}
}

type BisectedAssertDefender struct {
	*validatorConfig
	wholePrecondition *protocol.Precondition
	wholeAssertion    *protocol.AssertionStub
	splitDefenders    []protocol.AssertionDefender
	deadline          uint64
}

func (bot BisectedAssertDefender) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		//timeoutMsg := SendAsserterTimedOutChallengeMessage{
		//	bot.deadline,
		//	bot.wholePrecondition,
		//	bot.wholeAssertion,
		//}
		return TimedOutAsserterDefender{bot.validatorConfig}, nil, nil
	} else {
		return bot, nil, nil
	}
}

func (bot BisectedAssertDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.BisectMessage:
		deadline := time + bot.config.GracePeriod
		return WaitingBisectedDefender{
			bot.validatorConfig,
			bot.splitDefenders,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: BisectedAssertDefender: VM state got unsynchronized"}
	}
}

type WaitingBisectedDefender struct {
	*validatorConfig
	defenders []protocol.AssertionDefender
	deadline  uint64
}

func (bot WaitingBisectedDefender) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		preconditions := make([]*protocol.Precondition, 0, len(bot.defenders))
		assertions := make([]*protocol.AssertionStub, 0, len(bot.defenders))
		for _, defender := range bot.defenders {
			preconditions = append(preconditions, defender.GetPrecondition())
			assertions = append(assertions, defender.GetAssertion().Stub())
		}
		return TimedOutChallengerDefender{bot.validatorConfig},
			[]valmessage.OutgoingMessage{valmessage.SendChallengerTimedOutChallengeMessage{
				Deadline:      bot.deadline,
				Preconditions: preconditions,
				Assertions:    assertions,
			}}, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingBisectedDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case valmessage.ContinueChallengeMessage:
		if int(ev.ChallengedAssertion) >= len(bot.defenders) {
			return nil, nil, errors.New("ChallengedAssertion number is out of bounds")
		}
		return defenseValidator(bot.validatorConfig, bot.defenders[ev.ChallengedAssertion], time)
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: WaitingBisectedDefender: VM state got unsynchronized, %T", ev)}
	}
}

type OneStepChallengedAssertDefender struct {
	*validatorConfig
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot OneStepChallengedAssertDefender) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		//timeoutMsg := SendAsserterTimedOutChallengeMessage{
		//	bot.deadline,
		//	bot.precondition,
		//	bot.assertion,
		//}
		return TimedOutAsserterDefender{bot.validatorConfig}, nil, nil
	} else {
		return bot, nil, nil
	}
}

func (bot OneStepChallengedAssertDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.OneStepProofMessage:
		fmt.Println("OneStepChallengedAssertDefender: Proof was accepted")
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: OneStepChallengedAssertDefender: VM state got unsynchronized, %T", ev)}
	}
}

type TimedOutChallengerDefender struct {
	*validatorConfig
}

func (bot TimedOutChallengerDefender) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot TimedOutChallengerDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.ChallengerTimeoutMessage:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: TimedOutChallengerDefender: VM state got unsynchronized"}
	}
}

type TimedOutAsserterDefender struct {
	*validatorConfig
}

func (bot TimedOutAsserterDefender) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot TimedOutAsserterDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.AsserterTimeoutMessage:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: TimedOutAsserterDefender: VM state got unsynchronized"}
	}
}
