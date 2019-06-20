package validator

import (
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type callRequest struct {
	Message    protocol.Message
	ResultChan chan<- value.Value
	ErrorChan  chan<- error
}

type pendingMessageCheck struct {
	ResultChan chan<- bool
}

type vmStateRequest struct {
	ResultChan chan<- valmessage.VMStateData
}

type disputableDefenderRequest struct {
	Length                 uint64
	IncludePendingMessages bool
	ResultChan             chan<- bool
}

type initiateUnanimousRequest struct {
	TimeLength  uint64
	NewMessages []protocol.Message
	Final       bool
	RequestChan chan<- valmessage.UnanimousRequest
	ResultChan  chan<- valmessage.UnanimousUpdateResults
	ErrChan     chan<- error
}

type followUnanimousRequest struct {
	valmessage.UnanimousRequestData
	NewMessages []protocol.Message

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

type closeUnanimousAssertionRequest struct {
	ResultChan chan<- bool
	ErrChan    chan<- error
}

type unanimousConfirmRequest struct {
	valmessage.UnanimousRequestData
	Signatures [][]byte

	ResultChan chan<- bool
	ErrChan    chan<- error
}