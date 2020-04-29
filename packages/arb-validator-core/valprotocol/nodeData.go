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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type ChildType uint

const (
	InvalidInboxTopChildType  ChildType = 0
	InvalidMessagesChildType  ChildType = 1
	InvalidExecutionChildType ChildType = 2
	ValidChildType            ChildType = 3

	MinChildType        ChildType = 0
	MaxInvalidChildType ChildType = 2
	MaxChildType        ChildType = 3
)

type VMProtoData struct {
	MachineHash common.Hash
	InboxTop    common.Hash
	InboxCount  *big.Int
}

func NewVMProtoData(
	machineHash common.Hash,
	inboxTop common.Hash,
	inboxCount *big.Int,
) *VMProtoData {
	return &VMProtoData{
		MachineHash: machineHash,
		InboxTop:    inboxTop,
		InboxCount:  inboxCount,
	}
}

func (d *VMProtoData) String() string {
	return fmt.Sprintf(
		"VMProtoData(MachineHash: %v, InboxTop: %v, InboxCount: %v)",
		d.MachineHash,
		d.InboxTop,
		d.InboxCount,
	)
}

func (d *VMProtoData) Equals(o *VMProtoData) bool {
	return d.MachineHash == o.MachineHash &&
		d.InboxTop == o.InboxTop &&
		d.InboxCount.Cmp(o.InboxCount) == 0
}

func (d *VMProtoData) Clone() *VMProtoData {
	return &VMProtoData{
		MachineHash: d.MachineHash,
		InboxTop:    d.InboxTop,
		InboxCount:  new(big.Int).Set(d.InboxCount),
	}
}

func (d *VMProtoData) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(d.MachineHash),
		hashing.Bytes32(d.InboxTop),
		hashing.Uint256(new(big.Int).Set(d.InboxCount)),
	)
}

func (node *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash: node.MachineHash.MarshalToBuf(),
		InboxTop:    node.InboxTop.MarshalToBuf(),
		InboxCount:  common.MarshalBigInt(node.InboxCount),
	}
}

func (buf *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash: buf.MachineHash.Unmarshal(),
		InboxTop:    buf.InboxTop.Unmarshal(),
		InboxCount:  buf.InboxCount.Unmarshal(),
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
		ap.TimeBounds.StartBlock.AsInt(),
		ap.TimeBounds.EndBlock.AsInt(),
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
	AfterInboxTop         common.Hash
	ImportedMessagesSlice common.Hash
	AssertionStub         *ExecutionAssertionStub
}

func (dn *AssertionClaim) String() string {
	return fmt.Sprintf(
		"AssertionClaim(AfterInboxTop: %v, ImportedMessagesSlice: %v, Assertion: %v)",
		dn.AfterInboxTop,
		dn.ImportedMessagesSlice,
		dn.AssertionStub,
	)
}

func (dn *AssertionClaim) Equals(o *AssertionClaim) bool {
	return dn.AfterInboxTop == o.AfterInboxTop &&
		dn.ImportedMessagesSlice == o.ImportedMessagesSlice &&
		dn.AssertionStub.Equals(o.AssertionStub)
}

func (dn *AssertionClaim) Clone() *AssertionClaim {
	return &AssertionClaim{
		AfterInboxTop:         dn.AfterInboxTop,
		ImportedMessagesSlice: dn.ImportedMessagesSlice,
		AssertionStub:         dn.AssertionStub.Clone(),
	}
}

func (dn *AssertionClaim) MarshalToBuf() *AssertionClaimBuf {
	return &AssertionClaimBuf{
		AfterInboxTop:         dn.AfterInboxTop.MarshalToBuf(),
		ImportedMessagesSlice: dn.ImportedMessagesSlice.MarshalToBuf(),
		AssertionStub:         dn.AssertionStub.MarshalToBuf(),
	}
}

func (m *AssertionClaimBuf) Unmarshal() *AssertionClaim {
	return &AssertionClaim{
		AfterInboxTop:         m.AfterInboxTop.Unmarshal(),
		ImportedMessagesSlice: m.ImportedMessagesSlice.Unmarshal(),
		AssertionStub:         m.AssertionStub.Unmarshal(),
	}
}

type DisputableNode struct {
	AssertionParams *AssertionParams
	AssertionClaim  *AssertionClaim
	MaxInboxTop     common.Hash
	MaxInboxCount   *big.Int
}

func NewDisputableNode(
	assertionParams *AssertionParams,
	assertionClaim *AssertionClaim,
	maxInboxTop common.Hash,
	maxInboxCount *big.Int,
) *DisputableNode {
	return &DisputableNode{
		AssertionParams: assertionParams,
		AssertionClaim:  assertionClaim,
		MaxInboxTop:     maxInboxTop,
		MaxInboxCount:   maxInboxCount,
	}
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		AssertionParams: dn.AssertionParams.MarshalToBuf(),
		AssertionClaim:  dn.AssertionClaim.MarshalToBuf(),
		MaxInboxTop:     dn.MaxInboxTop.MarshalToBuf(),
		MaxInboxCount:   common.MarshalBigInt(dn.MaxInboxCount),
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	return NewDisputableNode(
		buf.AssertionParams.Unmarshal(),
		buf.AssertionClaim.Unmarshal(),
		buf.MaxInboxTop.Unmarshal(),
		buf.MaxInboxCount.Unmarshal(),
	)
}

func (dn *DisputableNode) CheckTime(params ChainParams) common.TimeTicks {
	checkTimeRaw := dn.AssertionClaim.AssertionStub.NumGas / params.ArbGasSpeedLimitPerTick
	return common.TimeTicks{Val: new(big.Int).SetUint64(checkTimeRaw)}
}

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.AssertionClaim.AssertionStub.AfterHash,
		dn.AssertionClaim.AfterInboxTop,
		new(big.Int).Add(prevState.InboxCount, dn.AssertionParams.ImportedMessageCount),
	)
}
