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

	"github.com/ethereum/go-ethereum/common/hexutil"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

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
	MachineHash  [32]byte
	PendingTop   [32]byte
	PendingCount *big.Int
}

func NewVMProtoData(
	machineHash [32]byte,
	pendingTop [32]byte,
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
		hexutil.Encode(d.MachineHash[:]),
		hexutil.Encode(d.PendingTop[:]),
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

func (d *VMProtoData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(d.MachineHash),
		solsha3.Bytes32(d.PendingTop),
		solsha3.Uint256(d.PendingCount),
	))
	return ret
}

func (node *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  utils.MarshalHash(node.MachineHash),
		PendingTop:   utils.MarshalHash(node.PendingTop),
		PendingCount: utils.MarshalBigInt(node.PendingCount),
	}
}

func (buf *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash:  utils.UnmarshalHash(buf.MachineHash),
		PendingTop:   utils.UnmarshalHash(buf.PendingTop),
		PendingCount: utils.UnmarshalBigInt(buf.PendingCount),
	}
}

type AssertionParams struct {
	NumSteps             uint32
	TimeBounds           *protocol.TimeBoundsBlocks
	ImportedMessageCount *big.Int
}

func (ap *AssertionParams) String() string {
	return fmt.Sprintf(
		"AssertionParams(NumSteps: %v, TimeBounds: [%v, %v], ImportedCount: %v)",
		ap.NumSteps,
		ap.TimeBounds.Start.Unmarshal().AsInt(),
		ap.TimeBounds.End.Unmarshal().AsInt(),
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
		TimeBounds:           ap.TimeBounds,
		ImportedMessageCount: utils.MarshalBigInt(ap.ImportedMessageCount),
	}
}

func (m *AssertionParamsBuf) Unmarshal() *AssertionParams {
	return &AssertionParams{
		NumSteps:             m.NumSteps,
		TimeBounds:           m.TimeBounds,
		ImportedMessageCount: utils.UnmarshalBigInt(m.ImportedMessageCount),
	}
}

type AssertionClaim struct {
	AfterPendingTop       [32]byte
	ImportedMessagesSlice [32]byte
	AssertionStub         *protocol.ExecutionAssertionStub
}

func (dn *AssertionClaim) String() string {
	return fmt.Sprintf(
		"AssertionClaim(AfterPendingTop: %v, ImportedMessagesSlice: %v, Assertion: %v)",
		hexutil.Encode(dn.AfterPendingTop[:]),
		hexutil.Encode(dn.ImportedMessagesSlice[:]),
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
		AfterPendingTop:       utils.MarshalHash(dn.AfterPendingTop),
		ImportedMessagesSlice: utils.MarshalHash(dn.ImportedMessagesSlice),
		AssertionStub:         dn.AssertionStub.MarshalToBuf(),
	}
}

func (m *AssertionClaimBuf) Unmarshal() *AssertionClaim {
	return &AssertionClaim{
		AfterPendingTop:       utils.UnmarshalHash(m.AfterPendingTop),
		ImportedMessagesSlice: utils.UnmarshalHash(m.ImportedMessagesSlice),
		AssertionStub:         m.AssertionStub.Unmarshal(),
	}
}

type DisputableNode struct {
	AssertionParams *AssertionParams
	AssertionClaim  *AssertionClaim
	MaxPendingTop   [32]byte
	MaxPendingCount *big.Int
}

func NewDisputableNode(
	assertionParams *AssertionParams,
	assertionClaim *AssertionClaim,
	maxPendingTop [32]byte,
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
		MaxPendingTop:   utils.MarshalHash(dn.MaxPendingTop),
		MaxPendingCount: utils.MarshalBigInt(dn.MaxPendingCount),
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	return NewDisputableNode(
		buf.AssertionParams.Unmarshal(),
		buf.AssertionClaim.Unmarshal(),
		utils.UnmarshalHash(buf.MaxPendingTop),
		utils.UnmarshalBigInt(buf.MaxPendingCount),
	)
}

func (dn *DisputableNode) CheckTime(params ChainParams) TimeTicks {
	checkTimeRaw := dn.AssertionClaim.AssertionStub.NumGas / params.ArbGasSpeedLimitPerTick
	return TimeTicks{Val: new(big.Int).SetUint64(checkTimeRaw)}
}

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.AssertionClaim.AssertionStub.AfterHash,
		dn.AssertionClaim.AfterPendingTop,
		new(big.Int).Add(prevState.PendingCount, dn.AssertionParams.ImportedMessageCount),
	)
}
