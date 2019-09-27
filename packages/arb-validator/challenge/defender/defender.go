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

package defender

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
)

func New(core *core.Config, assDef machine.AssertionDefender, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	deadline := time + core.VMConfig.GracePeriod
	if assDef.GetAssertion().NumSteps == 1 {
		fmt.Println("Generating proof")
		proofData, err := assDef.SolidityOneStepProof()
		if err != nil {
			return nil, &challenge.Error{Err: err, Message: "AssertAndDefendBot: error generating one-step proof"}
		}
		_, err = bridge.OneStepProof(
			context.Background(),
			assDef.GetPrecondition(),
			assDef.GetAssertion().Stub(),
			proofData,
		)
		return oneStepChallenged{
			Config:       core,
			precondition: assDef.GetPrecondition(),
			assertion:    assDef.GetAssertion().Stub(),
			deadline:     deadline,
		}, err
	}

	defenders := assDef.NBisect(6)
	assertions := make([]*protocol.AssertionStub, 0, len(defenders))
	for _, defender := range defenders {
		assertions = append(assertions, defender.GetAssertion().Stub())
	}
	_, err := bridge.BisectAssertion(
		context.Background(),
		assDef.GetPrecondition(),
		assertions,
	)
	return bisectedAssert{
		Config:            core,
		wholePrecondition: assDef.GetPrecondition(),
		wholeAssertion:    assDef.GetAssertion().Stub(),
		splitDefenders:    defenders,
		deadline:          deadline,
	}, err
}

type bisectedAssert struct {
	*core.Config
	wholePrecondition *protocol.Precondition
	wholeAssertion    *protocol.AssertionStub
	splitDefenders    []machine.AssertionDefender
	deadline          uint64
}

func (bot bisectedAssert) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.wholePrecondition,
	//	bot.wholeAssertion,
	//}
	return challenge.TimedOutAsserter{Config: bot.Config}, nil
}

func (bot bisectedAssert) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	log.Printf("bisectedAssert UpdateState event %T", ev)
	switch ev.(type) {
	case ethbridge.BisectionEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingBisected{
			bot.Config,
			bot.splitDefenders,
			deadline,
		}, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: bisectedAssert: VM state got unsynchronized"}
	}
}

type waitingBisected struct {
	*core.Config
	defenders []machine.AssertionDefender
	deadline  uint64
}

func (bot waitingBisected) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	//if time <= bot.deadline {
	//	return bot, nil
	//}

	_, err := bridge.ChallengerTimedOut(
		context.Background(),
	)
	if err != nil {
		return challenge.TimedOutChallenger{Config: bot.Config}, &challenge.Error{err, "WARNING: UpdateTime: challenger timed out", true}
	}

	return challenge.TimedOutChallenger{Config: bot.Config}, err
}

func (bot waitingBisected) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	log.Printf("defender waitingBisected UpdateState event %T", ev)
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		if int(ev.ChallengedAssertion) >= len(bot.defenders) {
			return nil, errors.New("ChallengedAssertion number is out of bounds")
		}
		return New(bot.Config, bot.defenders[ev.ChallengedAssertion], time, bridge)
	default:
		return nil, &challenge.Error{Message: fmt.Sprintf("ERROR: waitingBisected: VM state got unsynchronized, %T", ev)}
	}
}

type oneStepChallenged struct {
	*core.Config
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot oneStepChallenged) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.Assertion,
	//}
	return challenge.TimedOutAsserter{Config: bot.Config}, nil
}

func (bot oneStepChallenged) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.ArbVMBridge) (challenge.State, error) {
	log.Printf("oneStepChallenged UpdateState event %T", ev)
	switch ev.(type) {
	case ethbridge.OneStepProofEvent:
		fmt.Println("oneStepChallenged: Proof was accepted")
		return nil, nil
	default:
		return nil, &challenge.Error{Message: fmt.Sprintf("ERROR: oneStepChallenged: VM state got unsynchronized, %T", ev)}
	}
}
