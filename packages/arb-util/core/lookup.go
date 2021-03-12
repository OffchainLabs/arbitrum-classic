/*
 * Copyright 2021, Offchain Labs, Inc.
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

package core

import (
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type MessageStatus uint8

const (
	MessagesEmpty MessageStatus = iota
	MessagesReady
	MessagesSuccess
	MessagesNeedOlder
	MessagesError
)

type ExecutionCursor interface {
	Clone() ExecutionCursor
	MachineHash() (common.Hash, error)
	TotalMessagesRead() *big.Int
	InboxAcc() common.Hash
	SendAcc() common.Hash
	LogAcc() common.Hash
	TotalGasConsumed() *big.Int
	TotalSteps() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int
}

type ArbCoreLookup interface {
	ArbOutputLookup

	GetSendAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)
	GetLogAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)
	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetInboxAccPair(index1 *big.Int, index2 *big.Int) (common.Hash, common.Hash, error)

	// GetExecutionCursor returns a cursor containing the machine after executing totalGasUsed
	// from the original machine
	GetExecutionCursor(totalGasUsed *big.Int) (ExecutionCursor, error)

	// Advance executes as much as it can without going over maxGas or
	// optionally until it reaches or goes over maxGas
	AdvanceExecutionCursor(executionCursor ExecutionCursor, maxGas *big.Int, goOverGas bool) error

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine(executionCursor ExecutionCursor) (machine.Machine, error)
}

type ArbCoreInbox interface {
	DeliverMessages(messages []inbox.InboxMessage, previousInboxAcc common.Hash, lastBlockComplete bool, reorgHeight *big.Int) bool
	MessagesStatus() (MessageStatus, error)
}

func DeliverMessagesAndWait(db ArbCoreInbox, messages []inbox.InboxMessage, previousInboxAcc common.Hash, lastBlockComplete bool) (bool, error) {
	if !db.DeliverMessages(messages, previousInboxAcc, lastBlockComplete, nil) {
		return false, errors.New("unable to deliver messages")
	}
	status, err := waitForMessages(db)
	if err != nil {
		return false, err
	}
	if status == MessagesSuccess {
		return true, nil
	}
	if status == MessagesNeedOlder {
		return false, nil
	}
	return false, errors.New("Unexpected status")
}

func ReorgAndWait(db ArbCoreInbox, reorgMessageCount *big.Int) error {
	if !db.DeliverMessages(nil, common.Hash{}, false, reorgMessageCount) {
		return errors.New("unable to deliver messages")
	}
	status, err := waitForMessages(db)
	if err != nil {
		return err
	}
	if status == MessagesSuccess {
		return nil
	}
	return errors.New("Unexpected status")
}

func waitForMessages(db ArbCoreInbox) (MessageStatus, error) {
	start := time.Now()
	var status MessageStatus
	var err error
	for {
		status, err = db.MessagesStatus()
		if err != nil {
			return 0, err
		}

		if status == MessagesEmpty {
			return 0, errors.New("should have messages")
		}
		if status != MessagesReady {
			break
		}
		duration := time.Since(start)
		if duration > time.Second*30 {
			logger.Warn().Dur("elapsed", duration).Msg("Message delivery taking too long")
			start = time.Now()
		}
		<-time.After(time.Millisecond * 50)
	}
	return status, nil
}

type ArbCore interface {
	ArbCoreLookup
	ArbCoreInbox
	LogsCursor
	StartThread() bool
	StopThread()
	MachineIdle() bool
}

func GetSingleMessage(lookup ArbOutputLookup, index *big.Int) (inbox.InboxMessage, error) {
	messages, err := lookup.GetMessages(index, big.NewInt(1))
	if err != nil {
		return inbox.InboxMessage{}, err
	}
	if len(messages) == 0 {
		return inbox.InboxMessage{}, errors.New("no send found")
	}
	if len(messages) > 1 {
		return inbox.InboxMessage{}, errors.New("too many sends")
	}
	return messages[0], nil
}

func GetSingleSend(lookup ArbOutputLookup, index *big.Int) ([]byte, error) {
	sends, err := lookup.GetSends(index, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	if len(sends) == 0 {
		return nil, errors.New("no send found")
	}
	if len(sends) > 1 {
		return nil, errors.New("too many sends")
	}
	return sends[0], nil
}

func GetSingleLog(lookup ArbOutputLookup, index *big.Int) (value.Value, error) {
	logs, err := lookup.GetLogs(index, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("no log found")
	}
	if len(logs) > 1 {
		return nil, errors.New("too many logs")
	}
	return logs[0], nil
}

type ExecutionState struct {
	MachineHash       common.Hash
	InboxAcc          common.Hash
	TotalMessagesRead *big.Int
	TotalGasConsumed  *big.Int
	TotalSendCount    *big.Int
	TotalLogCount     *big.Int
	SendAcc           common.Hash
	LogAcc            common.Hash
}

func NewExecutionState(c ExecutionCursor) *ExecutionState {
	hash, err := c.MachineHash()
	if err != nil {
		panic("Unable to compute hash for execution state")
	}
	return &ExecutionState{
		MachineHash:       hash,
		InboxAcc:          c.InboxAcc(),
		TotalMessagesRead: c.TotalMessagesRead(),
		TotalGasConsumed:  c.TotalGasConsumed(),
		TotalSendCount:    c.TotalSendCount(),
		TotalLogCount:     c.TotalLogCount(),
		SendAcc:           c.SendAcc(),
		LogAcc:            c.LogAcc(),
	}
}

func (e *ExecutionState) IsPermanentlyBlocked() bool {
	var haltedHash common.Hash = [32]byte{}
	var erroredHash common.Hash = [32]byte{}
	erroredHash[31] = 1
	return e.MachineHash == haltedHash || e.MachineHash == erroredHash
}

func (e *ExecutionState) Equals(other Cut) bool {
	return e.CutHash() == other.CutHash()
}

func (e *ExecutionState) RestHash() [32]byte {
	return hashing.SoliditySHA3(
		hashing.Uint256(e.TotalMessagesRead),
		hashing.Bytes32(e.MachineHash),
		hashing.Bytes32(e.SendAcc),
		hashing.Uint256(e.TotalSendCount),
		hashing.Bytes32(e.LogAcc),
		hashing.Uint256(e.TotalLogCount),
	)
}

func (e *ExecutionState) CutHash() [32]byte {
	return hashing.SoliditySHA3(
		hashing.Uint256(e.TotalGasConsumed),
		hashing.Bytes32(e.RestHash()),
	)
}

type LogConsumer interface {
	AddLogs(initialIndex *big.Int, avmLogs []value.Value) error
	DeleteLogs(avmLogs []value.Value) error
}

type LogsCursor interface {
	LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error
	LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []value.Value, []value.Value, error)
	LogsCursorCheckError(cursorIndex *big.Int) error
	LogsCursorConfirmReceived(cursorIndex *big.Int) (bool, error)
	LogsCursorPosition(cursorIndex *big.Int) (*big.Int, error)
}
