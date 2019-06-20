package challenge

import (
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"
)

type State interface {
	UpdateTime(uint64, bridge.Bridge) (State, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (State, error)
}