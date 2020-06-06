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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/rand"
	"testing"
)

func generateSampleFullLog() (FullLog, error) {
	l, err := generateSampleLog()
	if err != nil {
		return FullLog{}, err
	}
	var txHash common.Hash
	var nodeHash common.Hash

	slicesToFill := [][]byte{txHash[:], nodeHash[:]}
	for _, sl := range slicesToFill {
		_, err := rand.Read(sl)
		if err != nil {
			return FullLog{}, err
		}
	}
	return FullLog{
		Log:        l,
		TxIndex:    rand.Uint64(),
		TxHash:     txHash,
		NodeHeight: rand.Uint64(),
		NodeHash:   nodeHash,
		Index:      rand.Uint64(),
		Removed:    false,
	}, nil
}

func TestFullLogMarshal(t *testing.T) {
	rand.Seed(43242)
	l, err := generateSampleFullLog()
	if err != nil {
		t.Fatal(err)
	}

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
	l, err := generateSampleFullLog()
	if err != nil {
		t.Fatal(err)
	}

	evmLog := l.ToEVMLog()
	if evmLog == nil {
		t.Fatal("log was nil")
	}
}
