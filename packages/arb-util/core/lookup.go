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

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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
	MachineHash() common.Hash
	TotalMessagesRead() *big.Int
	InboxAcc() common.Hash
	SendAcc() common.Hash
	LogAcc() common.Hash
	TotalGasConsumed() *big.Int
	TotalSteps() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine() (machine.Machine, error)
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
		if time.Since(start) > time.Second*30 {
			return 0, errors.New("timed out adding messages")
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
}

func NewExecutionState(c ExecutionCursor) *ExecutionState {
	return &ExecutionState{
		MachineHash:       c.MachineHash(),
		InboxAcc:          c.InboxAcc(),
		TotalMessagesRead: c.TotalMessagesRead(),
		TotalGasConsumed:  c.TotalGasConsumed(),
		TotalSendCount:    c.TotalSendCount(),
		TotalLogCount:     c.TotalLogCount(),
	}
}

func (e *ExecutionState) Equals(o *ExecutionState) bool {
	// We don't check InboxAcc here intentionally.
	// We don't assert InboxAcc, it's more of a side product.
	// Any relevant inbox changes will be reflected in other fields.
	return e.MachineHash == o.MachineHash &&
		e.TotalMessagesRead.Cmp(o.TotalMessagesRead) == 0 &&
		e.TotalGasConsumed.Cmp(o.TotalGasConsumed) == 0 &&
		e.TotalSendCount.Cmp(o.TotalSendCount) == 0 &&
		e.TotalLogCount.Cmp(o.TotalLogCount) == 0
}

func (e *ExecutionState) IsPermanentlyBlocked() bool {
	var haltedHash common.Hash = [32]byte{}
	var erroredHash common.Hash = [32]byte{}
	erroredHash[31] = 1
	return e.MachineHash == haltedHash || e.MachineHash == erroredHash
}

type ExecutionInfo struct {
	Before  *ExecutionState
	After   *ExecutionState
	SendAcc common.Hash
	LogAcc  common.Hash
}

func (e *ExecutionInfo) Equals(o *ExecutionInfo) bool {
	return e.Before.Equals(o.Before) &&
		e.After.Equals(o.After) &&
		e.SendAcc == o.SendAcc &&
		e.LogAcc == o.LogAcc
}

func (e *ExecutionInfo) GasUsed() *big.Int {
	return new(big.Int).Sub(e.After.TotalGasConsumed, e.Before.TotalGasConsumed)
}

func (e *ExecutionInfo) SendCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalSendCount, e.Before.TotalSendCount)
}

func (e *ExecutionInfo) LogCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalLogCount, e.Before.TotalLogCount)
}

func (e *ExecutionInfo) InboxMessagesRead() *big.Int {
	return new(big.Int).Sub(e.After.TotalMessagesRead, e.Before.TotalMessagesRead)
}

type LogConsumer interface {
	AddLogs(initialIndex *big.Int, avmLogs []value.Value) error
	DeleteLogs(avmLogs []value.Value) error
	CurrentLogCount() (*big.Int, error)
	UpdateCurrentLogCount(count *big.Int) error
}

type LogsCursor interface {
	LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error
	LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []value.Value, error)
	LogsCursorGetDeletedLogs(cursorIndex *big.Int) (*big.Int, []value.Value, error)
	LogsCursorCheckError(cursorIndex *big.Int) error
	LogsCursorConfirmReceived(cursorIndex *big.Int) (bool, error)
}
