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

package protocol

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ExecutionAssertionStub struct {
	AfterHash        [32]byte
	DidInboxInsn     bool
	NumGas           uint64
	FirstMessageHash [32]byte
	LastMessageHash  [32]byte
	FirstLogHash     [32]byte
	LastLogHash      [32]byte
}

func (a *ExecutionAssertionStub) MarshalToBuf() *ExecutionAssertionStubBuf {
	return &ExecutionAssertionStubBuf{
		AfterHash:        utils.MarshalHash(a.AfterHash),
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: utils.MarshalHash(a.FirstMessageHash),
		LastMessageHash:  utils.MarshalHash(a.LastMessageHash),
		FirstLogHash:     utils.MarshalHash(a.FirstLogHash),
		LastLogHash:      utils.MarshalHash(a.LastLogHash),
	}
}

func (a *ExecutionAssertionStubBuf) Unmarshal() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		AfterHash:        utils.UnmarshalHash(a.AfterHash),
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: utils.UnmarshalHash(a.FirstMessageHash),
		LastMessageHash:  utils.UnmarshalHash(a.LastMessageHash),
		FirstLogHash:     utils.UnmarshalHash(a.FirstLogHash),
		LastLogHash:      utils.UnmarshalHash(a.LastLogHash),
	}
}

func (a *ExecutionAssertionStub) Clone() *ExecutionAssertionStub {
	return &ExecutionAssertionStub{
		AfterHash:        a.AfterHash,
		DidInboxInsn:     a.DidInboxInsn,
		NumGas:           a.NumGas,
		FirstMessageHash: a.FirstMessageHash,
		LastMessageHash:  a.LastMessageHash,
		FirstLogHash:     a.FirstLogHash,
		LastLogHash:      a.LastLogHash,
	}
}

func (a *ExecutionAssertionStub) String() string {
	return fmt.Sprintf(
		"Assertion(AfterHash: %v, DidInboxInsn: %v, NumGas: %v, "+
			"FirstMessageHash: %v, LastMessageHash: %v, FirstLogHash: %v LastLogHash: %v)",
		hexutil.Encode(a.AfterHash[:]),
		a.DidInboxInsn,
		a.NumGas,
		hexutil.Encode(a.FirstMessageHash[:]),
		hexutil.Encode(a.LastMessageHash[:]),
		hexutil.Encode(a.FirstLogHash[:]),
		hexutil.Encode(a.LastLogHash[:]),
	)
}

func (a *ExecutionAssertionStub) Equals(b *ExecutionAssertionStub) bool {
	return a.AfterHash == b.AfterHash &&
		a.NumGas == b.NumGas &&
		a.FirstMessageHash == b.FirstMessageHash &&
		a.LastMessageHash == b.LastMessageHash &&
		a.FirstLogHash == b.FirstLogHash &&
		a.LastLogHash == b.LastLogHash
}

func (a *ExecutionAssertionStub) Hash() [32]byte {
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(a.AfterHash),
		solsha3.Bool(a.DidInboxInsn),
		solsha3.Uint64(a.NumGas),
		solsha3.Bytes32(a.FirstMessageHash),
		solsha3.Bytes32(a.LastMessageHash),
		solsha3.Bytes32(a.FirstLogHash),
		solsha3.Bytes32(a.LastLogHash),
	)
	copy(ret[:], hashVal)
	return ret
}

func (a *ExecutionAssertionStub) GeneratePostcondition(pre *Precondition) *Precondition {
	nextBeforeInbox := pre.BeforeInbox
	if a.DidInboxInsn {
		nextBeforeInbox = value.NewEmptyTuple()
	}
	return &Precondition{
		BeforeHash:  a.AfterHash,
		TimeBounds:  pre.TimeBounds,
		BeforeInbox: nextBeforeInbox,
	}
}

func GeneratePreconditions(pre *Precondition, assertions []*ExecutionAssertionStub) []*Precondition {
	preconditions := make([]*Precondition, 0, len(assertions))
	for _, assertion := range assertions {
		preconditions = append(preconditions, pre)
		pre = assertion.GeneratePostcondition(pre)
	}
	return preconditions
}
