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
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ArbOutputLookup interface {
	GetLogCount() (*big.Int, error)
	GetLogs(startIndex, count *big.Int) ([]ValueAndInbox, error)

	GetSendCount() (*big.Int, error)
	GetSends(startIndex, count *big.Int) ([][]byte, error)

	GetMessageCount() (*big.Int, error)
	GetMessages(startIndex, count *big.Int) ([]inbox.InboxMessage, error)

	GetSequencerBatchItems(startIndex *big.Int) ([]inbox.SequencerBatchItem, error)

	GetDelayedMessageCount() (*big.Int, error)
	GetTotalDelayedMessagesSequenced() (*big.Int, error)

	GetMachineForSideload(uint64, bool) (machine.Machine, error)
}

type InMemoryOutputLookup struct {
	sync.Mutex
	messages []inbox.InboxMessage
	logs     []value.Value
	sends    [][]byte
}

func (as *InMemoryOutputLookup) GetLogCount() (*big.Int, error) {
	return big.NewInt(int64(len(as.logs))), nil
}

func (as *InMemoryOutputLookup) GetLogs(startIndex, count *big.Int) ([]value.Value, error) {
	start := startIndex.Uint64()
	num := count.Uint64()
	totalCount := uint64(len(as.logs))

	if start > totalCount {
		return nil, errors.New("message index to low")
	}

	if start+num > totalCount {
		num = totalCount - start
	}

	return as.logs[start : start+num], nil
}

func (as *InMemoryOutputLookup) GetSendCount() (*big.Int, error) {
	return big.NewInt(int64(len(as.sends))), nil
}

func (as *InMemoryOutputLookup) GetSends(startIndex, count *big.Int) ([][]byte, error) {
	panic("implement me")
}

func (as *InMemoryOutputLookup) GetMessageCount() (*big.Int, error) {
	panic("implement me")
}

func (as *InMemoryOutputLookup) GetMessages(startIndex, count *big.Int) ([]inbox.InboxMessage, error) {
	panic("implement me")
}

func NewInMemoryOutputLookup() *InMemoryOutputLookup {
	return &InMemoryOutputLookup{}
}

func (as *InMemoryOutputLookup) GetMessage(index uint64) (value.Value, error) {
	as.Lock()
	defer as.Unlock()
	if index >= uint64(len(as.messages)) {
		return nil, errors.New("failed to get l2message")
	}
	panic("UNSUPPORTED")
	//return as.messages[index], nil
}

func (as *InMemoryOutputLookup) GetLog(index uint64) (value.Value, error) {
	as.Lock()
	defer as.Unlock()
	if index >= uint64(len(as.logs)) {
		return nil, errors.New("failed to get log")
	}
	return as.logs[index], nil
}

func (as *InMemoryOutputLookup) SaveLog(val value.Value) error {
	as.Lock()
	defer as.Unlock()
	as.logs = append(as.logs, val)
	return nil
}

func (as *InMemoryOutputLookup) LogCount() (uint64, error) {
	as.Lock()
	defer as.Unlock()
	return uint64(len(as.logs)), nil
}

func (as *InMemoryOutputLookup) MessageCount() (uint64, error) {
	as.Lock()
	defer as.Unlock()
	return uint64(len(as.messages)), nil
}
