package valmessage

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/protocol"
)

type UnanimousRequestData struct {
	BeforeHash  [32]byte
	BeforeInbox [32]byte
	SequenceNum uint64
	TimeBounds  protocol.TimeBounds
}

func (r UnanimousRequestData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(r.BeforeHash),
		solsha3.Bytes32(r.BeforeInbox),
		solsha3.Uint64(r.SequenceNum),
		solsha3.Uint64(r.TimeBounds[0]),
		solsha3.Uint64(r.TimeBounds[1]),
	))
	return ret
}

type UnanimousRequest struct {
	UnanimousRequestData
	NewMessages []protocol.Message
}

type UnanimousUpdateResults struct {
	SequenceNum       uint64
	BeforeHash        [32]byte
	TimeBounds        protocol.TimeBounds
	NewInboxHash      [32]byte
	OriginalInboxHash [32]byte
	Assertion         *protocol.Assertion
}

type VMStateData struct {
	MachineState [32]byte
	Config       VMConfiguration
}