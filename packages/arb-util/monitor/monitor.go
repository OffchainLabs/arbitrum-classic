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

package monitor

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var logger = log.With().Str("component", "monitor").Logger()

var GlobalMonitor = NewMonitor()

type txInfo struct {
	gotFromUserTime *time.Time
	batch           *common.Hash
}

func (tx *txInfo) LoggedAllValues() bool {
	return tx.gotFromUserTime != nil && tx.batch != nil
}

type batchInfo struct {
	submittedTime       *time.Time
	includedInBlockTime *time.Time
	readBatch           *time.Time
}

func (b *batchInfo) LoggedAllValues() bool {
	return b.submittedTime != nil &&
		b.includedInBlockTime != nil &&
		b.readBatch != nil
}

type Monitor struct {
	sync.Mutex
	txInfo    map[common.Hash]*txInfo
	batchInfo map[common.Hash]*batchInfo
	lastClear time.Time
}

func NewMonitor() *Monitor {
	return &Monitor{
		txInfo:    make(map[common.Hash]*txInfo),
		batchInfo: make(map[common.Hash]*batchInfo),
		lastClear: time.Now(),
	}
}

func (m *Monitor) maybeClear() {
	// Clear periodically to avoid too much memory bloat
	if time.Since(m.lastClear) > time.Hour {
		m.txInfo = make(map[common.Hash]*txInfo)
		m.batchInfo = make(map[common.Hash]*batchInfo)
		m.lastClear = time.Now()
	}
}

func (m *Monitor) GotTransactionFromUser(txHash common.Hash) {
	m.Lock()
	defer m.Unlock()
	m.maybeClear()
	_, ok := m.txInfo[txHash]
	if ok {
		// Already got from user
		return
	}
	currentTime := time.Now()
	m.txInfo[txHash] = &txInfo{
		gotFromUserTime: &currentTime,
	}
}

func (m *Monitor) IncludedInBatch(txHash common.Hash, batchHash common.Hash) {
	m.Lock()
	defer m.Unlock()
	m.maybeClear()
	tx, ok := m.txInfo[txHash]
	if !ok {
		return
	}
	tx.batch = &batchHash
}

func (m *Monitor) SubmittedBatch(txHash common.Hash) {
	m.Lock()
	defer m.Unlock()
	m.maybeClear()
	_, ok := m.batchInfo[txHash]
	if ok {
		return
	}
	currentTime := time.Now()
	m.batchInfo[txHash] = &batchInfo{
		submittedTime: &currentTime,
	}
}

func (m *Monitor) BatchAccepted(txHash common.Hash) {
	m.Lock()
	defer m.Unlock()
	m.maybeClear()
	batch, ok := m.batchInfo[txHash]
	if !ok {
		return
	}
	currentTime := time.Now()
	batch.includedInBlockTime = &currentTime
}

func (m *Monitor) ReaderGotBatch(txHash common.Hash) {
	m.Lock()
	defer m.Unlock()
	m.maybeClear()
	batch, ok := m.batchInfo[txHash]
	if !ok {
		return
	}
	currentTime := time.Now()
	batch.readBatch = &currentTime
}

func (m *Monitor) GotLog(txHash common.Hash) {
	m.Lock()
	defer m.Unlock()

	tx, ok := m.txInfo[txHash]
	if !ok {
		return
	}

	if !tx.LoggedAllValues() {
		return
	}

	batch, ok := m.batchInfo[*tx.batch]
	if !ok {
		return
	}

	if !batch.LoggedAllValues() {
		return
	}

	timeToSubmit := batch.submittedTime.Sub(*tx.gotFromUserTime)
	timeFromSubmissionToInclusion := batch.includedInBlockTime.Sub(*batch.submittedTime)
	timeFromSubmissionToReading := batch.readBatch.Sub(*batch.submittedTime)
	timeFromReadingToLog := time.Now().Sub(*batch.readBatch)

	delete(m.txInfo, txHash)

	logger.Info().
		Dur("timeToSubmit", timeToSubmit).
		Dur("timeFromSubmissionToInclusion", timeFromSubmissionToInclusion).
		Dur("timeFromSubmissionToReading", timeFromSubmissionToReading).
		Dur("timeFromReadingToLog", timeFromReadingToLog).
		Msg("transaction timing")

	m.maybeClear()
}
