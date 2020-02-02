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

package structures

import (
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-validator) -I. --go_out=paths=source_relative:. *.proto"

type ChildType uint

const (
	InvalidPendingChildType   ChildType = 0
	InvalidMessagesChildType  ChildType = 1
	InvalidExecutionChildType ChildType = 2
	ValidChildType            ChildType = 3

	MinChildType        ChildType = 0
	MaxInvalidChildType ChildType = 2
	MaxChildType        ChildType = 3
)

type VMProtoData struct {
	MachineHash  common.Hash
	PendingTop   common.Hash
	PendingCount *big.Int
}

func NewVMProtoData(
	machineHash common.Hash,
	pendingTop common.Hash,
	pendingCount *big.Int,
) *VMProtoData {
	return &VMProtoData{
		MachineHash:  machineHash,
		PendingTop:   pendingTop,
		PendingCount: pendingCount,
	}
}

func (d *VMProtoData) String() string {
	return fmt.Sprintf(
		"VMProtoData(MachineHash: %v, PendingTop: %v, PendingCount: %v)",
		d.MachineHash,
		d.PendingTop,
		d.PendingCount,
	)
}

func (d *VMProtoData) Equals(o *VMProtoData) bool {
	return d.MachineHash == o.MachineHash &&
		d.PendingTop == o.PendingTop &&
		d.PendingCount.Cmp(o.PendingCount) == 0
}

func (d *VMProtoData) Clone() *VMProtoData {
	return &VMProtoData{
		MachineHash:  d.MachineHash,
		PendingTop:   d.PendingTop,
		PendingCount: new(big.Int).Set(d.PendingCount),
	}
}

func (d *VMProtoData) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(d.MachineHash),
		hashing.Bytes32(d.PendingTop),
		hashing.Uint256(new(big.Int).Set(d.PendingCount)),
	)
}

func (d *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  d.MachineHash.MarshalToBuf(),
		PendingTop:   d.PendingTop.MarshalToBuf(),
		PendingCount: common.MarshalBigInt(d.PendingCount),
	}
}

func (m *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash:  m.MachineHash.Unmarshal(),
		PendingTop:   m.PendingTop.Unmarshal(),
		PendingCount: m.PendingCount.Unmarshal(),
	}
}

type AssertionParams struct {
	NumSteps             uint64
	TimeBounds           *protocol.TimeBoundsBlocks
	ImportedMessageCount *big.Int
}

func (ap *AssertionParams) String() string {
	return fmt.Sprintf(
		"AssertionParams(NumSteps: %v, TimeBounds: [%v, %v], ImportedCount: %v)",
		ap.NumSteps,
		ap.TimeBounds.Start.AsInt(),
		ap.TimeBounds.End.AsInt(),
		ap.ImportedMessageCount,
	)
}

func (ap *AssertionParams) Equals(o *AssertionParams) bool {
	return ap.NumSteps == o.NumSteps &&
		ap.TimeBounds.Equals(o.TimeBounds) &&
		ap.ImportedMessageCount.Cmp(o.ImportedMessageCount) == 0
}

func (ap *AssertionParams) Clone() *AssertionParams {
	return &AssertionParams{
		NumSteps:             ap.NumSteps,
		TimeBounds:           ap.TimeBounds.Clone(),
		ImportedMessageCount: new(big.Int).Set(ap.ImportedMessageCount),
	}
}

func (ap *AssertionParams) MarshalToBuf() *AssertionParamsBuf {
	return &AssertionParamsBuf{
		NumSteps:             ap.NumSteps,
		TimeBounds:           ap.TimeBounds.MarshalToBuf(),
		ImportedMessageCount: common.MarshalBigInt(ap.ImportedMessageCount),
	}
}

func (m *AssertionParamsBuf) Unmarshal() *AssertionParams {
	return &AssertionParams{
		NumSteps:             m.NumSteps,
		TimeBounds:           m.TimeBounds.Unmarshal(),
		ImportedMessageCount: m.ImportedMessageCount.Unmarshal(),
	}
}

type AssertionClaim struct {
	AfterPendingTop       common.Hash
	ImportedMessagesSlice common.Hash
	AssertionStub         *valprotocol.ExecutionAssertionStub
}

func (dn *AssertionClaim) String() string {
	return fmt.Sprintf(
		"AssertionClaim(AfterPendingTop: %v, ImportedMessagesSlice: %v, Assertion: %v)",
		dn.AfterPendingTop,
		dn.ImportedMessagesSlice,
		dn.AssertionStub,
	)
}

func (dn *AssertionClaim) Equals(o *AssertionClaim) bool {
	return dn.AfterPendingTop == o.AfterPendingTop &&
		dn.ImportedMessagesSlice == o.ImportedMessagesSlice &&
		dn.AssertionStub.Equals(o.AssertionStub)
}

func (dn *AssertionClaim) Clone() *AssertionClaim {
	return &AssertionClaim{
		AfterPendingTop:       dn.AfterPendingTop,
		ImportedMessagesSlice: dn.ImportedMessagesSlice,
		AssertionStub:         dn.AssertionStub.Clone(),
	}
}

func (dn *AssertionClaim) MarshalToBuf() *AssertionClaimBuf {
	return &AssertionClaimBuf{
		AfterPendingTop:       dn.AfterPendingTop.MarshalToBuf(),
		ImportedMessagesSlice: dn.ImportedMessagesSlice.MarshalToBuf(),
		AssertionStub:         dn.AssertionStub.MarshalToBuf(),
	}
}

func (m *AssertionClaimBuf) Unmarshal() *AssertionClaim {
	return &AssertionClaim{
		AfterPendingTop:       m.AfterPendingTop.Unmarshal(),
		ImportedMessagesSlice: m.ImportedMessagesSlice.Unmarshal(),
		AssertionStub:         m.AssertionStub.Unmarshal(),
	}
}

type DisputableNode struct {
	AssertionParams *AssertionParams
	AssertionClaim  *AssertionClaim
	MaxPendingTop   common.Hash
	MaxPendingCount *big.Int
}

func NewDisputableNode(
	assertionParams *AssertionParams,
	assertionClaim *AssertionClaim,
	maxPendingTop common.Hash,
	maxPendingCount *big.Int,
) *DisputableNode {
	return &DisputableNode{
		AssertionParams: assertionParams,
		AssertionClaim:  assertionClaim,
		MaxPendingTop:   maxPendingTop,
		MaxPendingCount: maxPendingCount,
	}
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		AssertionParams: dn.AssertionParams.MarshalToBuf(),
		AssertionClaim:  dn.AssertionClaim.MarshalToBuf(),
		MaxPendingTop:   dn.MaxPendingTop.MarshalToBuf(),
		MaxPendingCount: common.MarshalBigInt(dn.MaxPendingCount),
	}
}

func (m *DisputableNodeBuf) Unmarshal() *DisputableNode {
	return NewDisputableNode(
		m.AssertionParams.Unmarshal(),
		m.AssertionClaim.Unmarshal(),
		m.MaxPendingTop.Unmarshal(),
		m.MaxPendingCount.Unmarshal(),
	)
}

func (dn *DisputableNode) CheckTime(params ChainParams) common.TimeTicks {
	checkTimeRaw := dn.AssertionClaim.AssertionStub.NumGas / params.ArbGasSpeedLimitPerTick
	return common.TimeTicks{Val: new(big.Int).SetUint64(checkTimeRaw)}
}

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.AssertionClaim.AssertionStub.AfterHash,
		dn.AssertionClaim.AfterPendingTop,
		new(big.Int).Add(prevState.PendingCount, dn.AssertionParams.ImportedMessageCount),
	)
}
