/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package valprotocol

import (
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type ExecutionAssertionStub struct {
	NumGas            uint64
	BeforeMachineHash common.Hash
	AfterMachineHash  common.Hash
	BeforeInboxHash   common.Hash
	AfterInboxHash    common.Hash
	FirstMessageHash  common.Hash
	LastMessageHash   common.Hash
	MessageCount      uint64
	FirstLogHash      common.Hash
	LastLogHash       common.Hash
	LogCount          uint64
}

var zeroBufferHash = hashing.SoliditySHA3(hashing.Uint256(big.NewInt(0)))
var tagHash = hashing.SoliditySHA3(hashing.Uint256(big.NewInt(123)))

func roundUpToPow2(len int) int {
	if len <= 1 {
		return 1
	}
	return 2 * roundUpToPow2((len+1)/2)
}

func hashBufferAux(buf []byte, pack bool) (common.Hash, bool) {
	if len(buf) == 0 {
		return zeroBufferHash, true
	}
	if len(buf) == 32 {
		var arr [32]byte
		copy(arr[:], buf)
		res := hashing.SoliditySHA3(hashing.Bytes32(arr))
		return res, res == zeroBufferHash
	}
	len := len(buf)
	h2, zero2 := hashBufferAux(buf[len/2:len], false)
	if zero2 && pack {
		return hashBufferAux(buf[0:len/2], pack)
	}
	h1, zero1 := hashBufferAux(buf[0:len/2], false)
	return hashing.SoliditySHA3(hashing.Bytes32(h1), hashing.Bytes32(h2)), zero1 && zero2
}

func hashBuffer(buf []byte) common.Hash {
	// convert to pow2
	arr := make([]byte, roundUpToPow2(len(buf)))
	copy(arr, buf)
	res, _ := hashBufferAux(arr, true)
	return res
}

/*
func BytesArrayAccumHash(initialHash common.Hash, data []byte, valCount uint64) common.Hash {
	lastMsgHash := initialHash
	rd := bytes.NewReader(data)
	for i := uint64(0); i < valCount; i++ {
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			panic(err)
		}
		lastMsgHash = hashing.SoliditySHA3(
			hashing.Bytes32(lastMsgHash),
			hashing.Bytes32(val.Hash()))
	}
	return lastMsgHash
}
*/

func BufferAccumHash(initialHash common.Hash, data [][]byte) common.Hash {
	lastMsgHash := initialHash
	for msg := range data {
		bufHash := hashing.SoliditySHA3(
			hashing.Bytes32(tagHash),
			hashing.Bytes32(hashBuffer(data[msg])))
		msgHash := hashing.SoliditySHA3(
			hashing.Uint256(big.NewInt(int64(len(data[msg])))),
			hashing.Bytes32(bufHash))
		lastMsgHash = hashing.SoliditySHA3(
			hashing.Bytes32(lastMsgHash),
			hashing.Bytes32(msgHash))
	}
	return lastMsgHash
}

func (a *ExecutionAssertionStub) MarshalToBuf() *ExecutionAssertionStubBuf {
	return &ExecutionAssertionStubBuf{
		NumGas:            a.NumGas,
		BeforeMachineHash: a.BeforeMachineHash.MarshalToBuf(),
		AfterMachineHash:  a.AfterMachineHash.MarshalToBuf(),
		BeforeInboxHash:   a.BeforeInboxHash.MarshalToBuf(),
		AfterInboxHash:    a.AfterInboxHash.MarshalToBuf(),
		FirstMessageHash:  a.FirstMessageHash.MarshalToBuf(),
		LastMessageHash:   a.LastMessageHash.MarshalToBuf(),
		MessageCount:      a.MessageCount,
		FirstLogHash:      a.FirstLogHash.MarshalToBuf(),
		LastLogHash:       a.LastLogHash.MarshalToBuf(),
		LogCount:          a.LogCount,
	}
}

func (a *ExecutionAssertionStubBuf) Unmarshal() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		NumGas:            a.NumGas,
		BeforeMachineHash: a.BeforeMachineHash.Unmarshal(),
		AfterMachineHash:  a.AfterMachineHash.Unmarshal(),
		BeforeInboxHash:   a.BeforeInboxHash.Unmarshal(),
		AfterInboxHash:    a.AfterInboxHash.Unmarshal(),
		FirstMessageHash:  a.FirstMessageHash.Unmarshal(),
		LastMessageHash:   a.LastMessageHash.Unmarshal(),
		MessageCount:      a.MessageCount,
		FirstLogHash:      a.FirstLogHash.Unmarshal(),
		LastLogHash:       a.LastLogHash.Unmarshal(),
		LogCount:          a.LogCount,
	}
}

func (a *ExecutionAssertionStub) Clone() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		NumGas:            a.NumGas,
		BeforeMachineHash: a.BeforeMachineHash,
		AfterMachineHash:  a.AfterMachineHash,
		BeforeInboxHash:   a.BeforeInboxHash,
		AfterInboxHash:    a.AfterInboxHash,
		FirstMessageHash:  a.FirstMessageHash,
		LastMessageHash:   a.LastMessageHash,
		MessageCount:      a.MessageCount,
		FirstLogHash:      a.FirstLogHash,
		LastLogHash:       a.LastLogHash,
		LogCount:          a.LogCount,
	}
}

func (a *ExecutionAssertionStub) String() string {
	return fmt.Sprintf(
		"Assertion(NumGas: %v, BeforeMachineHash: %v, AfterMachineHash: %v, BeforeInboxHash: %v, AfterInboxHash: %v "+
			"FirstMessageHash: %v, LastMessageHash: %v, MessageCount %v, FirstLogHash: %v LastLogHash: %v, LogCount %v)",
		a.NumGas,
		a.BeforeMachineHash,
		a.AfterMachineHash,
		a.BeforeInboxHash,
		a.AfterInboxHash,
		a.FirstMessageHash,
		a.LastMessageHash,
		a.MessageCount,
		a.FirstLogHash,
		a.LastLogHash,
		a.LogCount,
	)
}

func (a *ExecutionAssertionStub) Equals(b *ExecutionAssertionStub) bool {
	return a.NumGas == b.NumGas &&
		a.BeforeMachineHash == b.BeforeMachineHash &&
		a.AfterMachineHash == b.AfterMachineHash &&
		a.BeforeInboxHash == b.BeforeInboxHash &&
		a.AfterInboxHash == b.AfterInboxHash &&
		a.FirstMessageHash == b.FirstMessageHash &&
		a.LastMessageHash == b.LastMessageHash &&
		a.MessageCount == b.MessageCount &&
		a.FirstLogHash == b.FirstLogHash &&
		a.LastLogHash == b.LastLogHash &&
		a.LogCount == b.LogCount
}

func (dn *ExecutionAssertionStub) CheckTime(params ChainParams) common.TimeTicks {
	checkTimeRaw := dn.NumGas / params.ArbGasSpeedLimitPerTick
	return common.TimeTicks{Val: new(big.Int).SetUint64(checkTimeRaw)}
}

func CalculateNodeDeadline(
	assertion *ExecutionAssertionStub,
	params ChainParams,
	prevDeadline common.TimeTicks,
	assertionTime common.TimeTicks,
) common.TimeTicks {
	checkTime := assertion.CheckTime(params)
	deadlineTicks := assertionTime.Add(params.GracePeriod)
	if deadlineTicks.Cmp(prevDeadline) >= 0 {
		return deadlineTicks.Add(checkTime)
	} else {
		return prevDeadline.Add(checkTime)
	}
}
