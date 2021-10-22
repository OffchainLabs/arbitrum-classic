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
}

type ArbCoreLookup interface {
	ArbOutputLookup

	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetDelayedInboxAcc(index *big.Int) (common.Hash, error)
	GetInboxAccPair(index1 *big.Int, index2 *big.Int) (common.Hash, common.Hash, error)
	CountMatchingBatchAccs(lastSeqNums []*big.Int, accs []common.Hash) (ret int, err error)
	GetDelayedMessagesToSequence(maxBlock *big.Int) (*big.Int, error)
	GetSequencerBlockNumberAt(index *big.Int) (*big.Int, error)
	GenInboxProof(seqNum *big.Int, batchIndex *big.Int, batchEndCount *big.Int) ([]byte, error)

	MachineMessagesRead() *big.Int

	// GetLastMachine gets a copy of the machine from the last time machinethread stopped or the last reorg
	GetLastMachine() (machine.Machine, error)

	GetLastMachineTotalGas() (*big.Int, error)

	// GetExecutionCursor returns a cursor containing the machine after executing totalGasUsed
	// from the original machine
	GetExecutionCursor(totalGasUsed *big.Int, allowSlowLookup bool) (ExecutionCursor, error)

	// AdvanceExecutionCursor executes as much as it can without going over maxGas or
	// optionally until it reaches or goes over maxGas
	AdvanceExecutionCursor(executionCursor ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool) error

	// AdvanceExecutionCursorWithTracing executes as much as it can without going over maxGas or
	// optionally until it reaches or goes over maxGas and returns all debug prints created
	AdvanceExecutionCursorWithTracing(executionCursor ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool) ([]value.Value, error)

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine(executionCursor ExecutionCursor) (machine.Machine, error)
}

type ArbCoreInbox interface {
	DeliverMessages(previousMessageCount *big.Int, previousSeqBatchAcc common.Hash, seqBatchItems []inbox.SequencerBatchItem, delayedMessages []inbox.DelayedMessage, reorgSeqBatchItemCount *big.Int) bool
	MessagesStatus() (MessageStatus, error)
	PrintCoreThreadBacktrace()
}

func DeliverMessagesAndWait(db ArbCoreInbox, previousMessageCount *big.Int, previousSeqBatchAcc common.Hash, seqBatchItems []inbox.SequencerBatchItem, delayedMessages []inbox.DelayedMessage, reorgSeqBatchItemCount *big.Int) error {
	if !db.DeliverMessages(previousMessageCount, previousSeqBatchAcc, seqBatchItems, delayedMessages, reorgSeqBatchItemCount) {
		return errors.New("unable to deliver messages")
	}
	status, err := waitForMessages(db)
	if err != nil {
		return err
	}
	if status != MessagesSuccess {
		return errors.New("Unexpected status")
	}
	return nil
}

func ReorgAndWait(db ArbCoreInbox, reorgMessageCount *big.Int) error {
	if !db.DeliverMessages(big.NewInt(0), common.Hash{}, nil, nil, reorgMessageCount) {
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

func WaitForMachineIdle(db ArbCore) {
	for {
		idle := db.MachineIdle()
		if idle {
			break
		}
		time.Sleep(time.Millisecond * 20)
	}
}

func waitForMessages(db ArbCoreInbox) (MessageStatus, error) {
	start := time.Now()
	nextLog := time.Second * 30
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
		if duration > nextLog {
			logger.Warn().Dur("elapsed", duration).Msg("Message delivery taking too long")
			db.PrintCoreThreadBacktrace()
			nextLog += time.Second * 30
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

func GetZeroOrOneLog(lookup ArbOutputLookup, index *big.Int) (ValueAndInbox, error) {
	logs, err := lookup.GetLogs(index, big.NewInt(1))
	if err != nil {
		return ValueAndInbox{}, err
	}
	if len(logs) == 0 {
		return ValueAndInbox{}, nil
	}
	if len(logs) > 1 {
		return ValueAndInbox{}, errors.New("too many logs")
	}
	return logs[0], nil
}

type ExecutionState struct {
	MachineHash       common.Hash `json:"machineHash"`
	InboxAcc          common.Hash `json:"inboxAcc"`
	TotalMessagesRead *big.Int    `json:"inboxCount"`
	TotalGasConsumed  *big.Int    `json:"gasUsed"`
	TotalSendCount    *big.Int    `json:"sendCount"`
	TotalLogCount     *big.Int    `json:"logCount"`
	SendAcc           common.Hash `json:"sendAcc"`
	LogAcc            common.Hash `json:"logAcc"`
}

func NewExecutionState(c ExecutionCursor) (*ExecutionState, error) {
	return &ExecutionState{
		MachineHash:       c.MachineHash(),
		InboxAcc:          c.InboxAcc(),
		TotalMessagesRead: c.TotalMessagesRead(),
		TotalGasConsumed:  c.TotalGasConsumed(),
		TotalSendCount:    c.TotalSendCount(),
		TotalLogCount:     c.TotalLogCount(),
		SendAcc:           c.SendAcc(),
		LogAcc:            c.LogAcc(),
	}, nil
}

func (e *ExecutionState) IsPermanentlyBlocked() bool {
	var haltedHash common.Hash = [32]byte{}
	var erroredHash common.Hash = [32]byte{}
	erroredHash[31] = 1
	return e.MachineHash == haltedHash || e.MachineHash == erroredHash
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

func (e *ExecutionState) CutHash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(e.TotalGasConsumed),
		hashing.Bytes32(e.RestHash()),
	)
}

type InboxState struct {
	Count       *big.Int
	Accumulator common.Hash
}

type ValueAndInbox struct {
	Value value.Value
	Inbox InboxState
}

type LogConsumer interface {
	AddLogs(initialIndex *big.Int, avmLogs []ValueAndInbox) error
	DeleteLogs(avmLogs []ValueAndInbox) error
}

type LogsCursor interface {
	LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error
	LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []ValueAndInbox, []ValueAndInbox, error)
	LogsCursorCheckError(cursorIndex *big.Int) error
	LogsCursorConfirmReceived(cursorIndex *big.Int) (bool, error)
	LogsCursorPosition(cursorIndex *big.Int) (*big.Int, error)
}
