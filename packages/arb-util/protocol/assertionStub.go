/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type AssertionStub struct {
	AfterHash        [32]byte
	NumSteps         uint32
	FirstMessageHash [32]byte
	LastMessageHash  [32]byte
	FirstLogHash     [32]byte
	LastLogHash      [32]byte
	TotalVals        []*big.Int
}

func (a *AssertionStub) String() string {
	return fmt.Sprintf("AssertionStub(%x, %v, %x, %x, %v)", a.AfterHash, a.NumSteps, a.FirstMessageHash, a.LastMessageHash, a.TotalVals)
}

func NewAssertionStubFromReader(rd io.Reader) (AssertionStub, error) {
	var afterHash [32]byte
	_, err := io.ReadFull(rd, afterHash[:])
	if err != nil {
		return AssertionStub{}, err
	}
	var numSteps uint32
	err = binary.Read(rd, binary.LittleEndian, &numSteps)
	if err != nil {
		return AssertionStub{}, err
	}
	var firstMessageHash [32]byte
	_, err = io.ReadFull(rd, firstMessageHash[:])
	if err != nil {
		return AssertionStub{}, err
	}
	var lastMessageHash [32]byte
	_, err = io.ReadFull(rd, lastMessageHash[:])
	if err != nil {
		return AssertionStub{}, err
	}

	var firstLogHash [32]byte
	_, err = io.ReadFull(rd, firstMessageHash[:])
	if err != nil {
		return AssertionStub{}, err
	}
	var lastLogHash [32]byte
	_, err = io.ReadFull(rd, lastMessageHash[:])
	if err != nil {
		return AssertionStub{}, err
	}

	var valCount int32
	err = binary.Read(rd, binary.LittleEndian, &valCount)
	if err != nil {
		return AssertionStub{}, err
	}
	totalVals := make([]*big.Int, valCount)
	for i := range totalVals {
		intVal, err := value.NewIntValueFromReader(rd)
		if err != nil {
			return AssertionStub{}, err
		}
		totalVals[i] = intVal.BigInt()
	}
	return AssertionStub{afterHash, numSteps, firstMessageHash, lastMessageHash, firstLogHash, lastLogHash, totalVals}, nil
}

func (a *AssertionStub) Marshal(wr io.Writer) error {
	if wr == nil {
		return nil
	}
	_, err := wr.Write(a.AfterHash[:])
	if err != nil {
		return err
	}
	err = binary.Write(wr, binary.LittleEndian, &a.NumSteps)
	if err != nil {
		return err
	}
	_, err = wr.Write(a.FirstMessageHash[:])
	if err != nil {
		return err
	}
	_, err = wr.Write(a.LastMessageHash[:])
	if err != nil {
		return err
	}
	numMsgs := int32(len(a.TotalVals))
	err = binary.Write(wr, binary.LittleEndian, &numMsgs)
	if err != nil {
		return err
	}
	for _, val := range a.TotalVals {
		err := value.NewIntValue(val).Marshal(wr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AssertionStub) Equals(b *AssertionStub) bool {
	if a.AfterHash != b.AfterHash ||
		a.NumSteps != b.NumSteps ||
		a.FirstMessageHash != b.FirstMessageHash ||
		a.LastMessageHash != b.LastMessageHash ||
		(len(a.TotalVals) != len(b.TotalVals)) {
		return false
	}
	for i, ao := range a.TotalVals {
		if ao.Cmp(b.TotalVals[i]) != 0 {
			return false
		}
	}
	return true
}

func (a *AssertionStub) Hash() [32]byte {
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(a.AfterHash),
		solsha3.Uint32(a.NumSteps),
		solsha3.Bytes32(a.FirstMessageHash),
		solsha3.Bytes32(a.LastMessageHash),
		solsha3.Bytes32(a.FirstLogHash),
		solsha3.Bytes32(a.LastLogHash),
		solsha3.Uint256Array(a.TotalVals),
	)
	copy(ret[:], hashVal)
	return ret
}

func (a *AssertionStub) GeneratePostcondition(pre *Precondition) *Precondition {
	bt := pre.BeforeBalance.Clone()
	for i, val := range a.TotalVals {
		bt.TokenAmounts[i].Sub(bt.TokenAmounts[i], val)
	}
	return NewPrecondition(a.AfterHash, pre.TimeBounds, bt, pre.BeforeInbox)
}

func GeneratePreconditions(pre *Precondition, assertions []*AssertionStub) []*Precondition {
	preconditions := make([]*Precondition, 0, len(assertions))
	for _, assertion := range assertions {
		preconditions = append(preconditions, pre)
		pre = assertion.GeneratePostcondition(pre)
	}
	return preconditions
}
