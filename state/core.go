package state

import (
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/challenge"
	"github.com/offchainlabs/arb-validator/core"
	"github.com/offchainlabs/arb-validator/ethbridge"
)

type State interface {
	UpdateTime(uint64, bridge.Bridge) (State, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (State, challenge.State, error)

	SendMessageToVM(msg protocol.Message)
	GetCore() *core.Core
	GetConfig() *core.Config
}
