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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
)

func New(assDef machine.AssertionDefender, deadline uint64, brdg bridge.Challenge) (challenge.State, error) {
	if assDef.NumSteps() == 1 {
		fmt.Println("Generating proof")
		proofData, err := assDef.SolidityOneStepProof()
		if err != nil {
			return nil, &bridge.Error{Err: err, Message: "AssertAndDefendBot: error generating solidity one-step proof"}
		}
		_, err = brdg.OneStepProof(
			context.Background(),
			assDef.GetPrecondition(),
			assDef.GetAssertion().Stub(),
			proofData,
		)
		if err != nil {
			err = &bridge.Error{Err: err, Message: "AssertAndDefendBot: error generating one-step proof"}
		}
		return oneStepChallenged{
			precondition: assDef.GetPrecondition(),
			assertion:    assDef.GetAssertion().Stub(),
			deadline:     deadline,
		}, err
	}

	defenders := assDef.NBisect(50)
	assertions := make([]*protocol.ExecutionAssertionStub, 0, len(defenders))
	for _, defender := range defenders {
		assertions = append(assertions, defender.GetAssertion().Stub())
	}
	_, err := brdg.BisectAssertion(
		context.Background(),
		assDef.GetPrecondition(),
		assertions,
	)
	return bisectedAssert{
		wholePrecondition: assDef.GetPrecondition(),
		wholeAssertion:    assDef.GetAssertion().Stub(),
		splitDefenders:    defenders,
		deadline:          deadline,
	}, err
}

type bisectedAssert struct {
	wholePrecondition *protocol.Precondition
	wholeAssertion    *protocol.ExecutionAssertionStub
	splitDefenders    []machine.AssertionDefender
	deadline          uint64
}

func (bot bisectedAssert) UpdateTime(time uint64, bridge bridge.Challenge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.wholePrecondition,
	//	bot.wholeAssertion,
	//}
	return challenge.TimedOutAsserter{}, nil
}

func (bot bisectedAssert) UpdateState(ev ethbridge.Event, time uint64, brdg bridge.Challenge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ExecutionBisectionEvent:
		return waitingBisected{
			bot.splitDefenders,
			ev.Deadline,
		}, nil
	default:
		return nil, &bridge.Error{Message: "ERROR: bisectedAssert: VM state got unsynchronized"}
	}
}

type waitingBisected struct {
	defenders []machine.AssertionDefender
	deadline  uint64
}

func (bot waitingBisected) UpdateTime(time uint64, brdg bridge.Challenge) (challenge.State, error) {
	// comment out to force challenge
	if time <= bot.deadline {
		return bot, nil
	}

	_, err := brdg.ChallengerTimedOut(
		context.Background(),
	)
	if err != nil {
		return challenge.TimedOutChallenger{}, &bridge.Error{err, "WARNING: UpdateTime: challenger timed out", true}
	}

	return challenge.TimedOutChallenger{}, err
}

func (bot waitingBisected) UpdateState(ev ethbridge.Event, time uint64, brdg bridge.Challenge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		if int(ev.ChallengedAssertion) >= len(bot.defenders) {
			return nil, errors.New("ChallengedAssertion number is out of bounds")
		}
		return New(bot.defenders[ev.ChallengedAssertion], ev.Deadline, brdg)
	default:
		return nil, &bridge.Error{Message: fmt.Sprintf("ERROR: waitingBisected: VM state got unsynchronized, %T", ev)}
	}
}

type oneStepChallenged struct {
	precondition *protocol.Precondition
	assertion    *protocol.ExecutionAssertionStub
	deadline     uint64
}

func (bot oneStepChallenged) UpdateTime(time uint64, bridge bridge.Challenge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.ExecutionAssertion,
	//}
	return challenge.TimedOutAsserter{}, nil
}

func (bot oneStepChallenged) UpdateState(ev ethbridge.Event, time uint64, brdg bridge.Challenge) (challenge.State, error) {
	switch ev.(type) {
	case ethbridge.OneStepProofEvent:
		fmt.Println("oneStepChallenged: Proof was accepted")
		brdg.SendMonitorMsg(bridge.ProofAccepted)
		return nil, nil
	default:
		return nil, &bridge.Error{Message: fmt.Sprintf("ERROR: oneStepChallenged: VM state got unsynchronized, %T", ev)}
	}
}
