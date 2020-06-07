/*
 * Copyright 2020, Offchain Labs, Inc.
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

package evm

import (
	"math/rand"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func newRandomFullLog(topicCount int32) FullLog {
	return FullLog{
		Log:        NewRandomLog(topicCount),
		TxIndex:    rand.Uint64(),
		TxHash:     common.RandHash(),
		NodeHeight: rand.Uint64(),
		NodeHash:   common.RandHash(),
		Index:      rand.Uint64(),
		Removed:    false,
	}
}

func TestFullLogMarshal(t *testing.T) {
	rand.Seed(43242)
	l := newRandomFullLog(3)
	lBuf := l.Marshal()

	l2, err := lBuf.Unmarshal()
	if err != nil {
		t.Fatal(err)
	}

	if !l.Equals(l2) {
		t.Fatal("not equal after unmarshal")
	}
}

func TestFullLogToEVMLog(t *testing.T) {
	rand.Seed(43242)
	l := newRandomFullLog(3)
	evmLog := l.ToEVMLog()
	if evmLog == nil {
		t.Fatal("log was nil")
	}
}
