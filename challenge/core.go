package challenge

import (
	"fmt"

	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/core"
	"github.com/offchainlabs/arb-validator/ethbridge"
)

type State interface {
	UpdateTime(uint64, bridge.Bridge) (State, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (State, error)
}

type Error struct {
	Err     error
	Message string
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%v: %v", e.Message, e.Err)
	}
	return e.Message
}

type TimedOutChallenger struct {
	*core.Config
}

func (bot TimedOutChallenger) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot TimedOutChallenger) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: TimedOutChallenger: VM state got unsynchronized"}
	}
}

type TimedOutAsserter struct {
	*core.Config
}

func (bot TimedOutAsserter) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot TimedOutAsserter) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: TimedOutAsserter: VM state got unsynchronized"}
	}
}
