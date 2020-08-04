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
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
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
	MachineHash  common.Hash
	InboxTop     common.Hash
	InboxCount   *big.Int
	MessageCount *big.Int
	LogCount     *big.Int
}

func NewVMProtoData(
	machineHash common.Hash,
	inboxTop common.Hash,
	inboxCount *big.Int,
	messageCount *big.Int,
	logCount *big.Int,
) *VMProtoData {
	return &VMProtoData{
		MachineHash:  machineHash,
		InboxTop:     inboxTop,
		InboxCount:   inboxCount,
		MessageCount: messageCount,
		LogCount:     logCount,
	}
}

func (d *VMProtoData) String() string {
	return fmt.Sprintf(
		"VMProtoData(MachineHash: %v, InboxTop: %v, InboxCount: %v, MessageCount: %v, LogCount: %v)",
		d.MachineHash,
		d.InboxTop,
		d.InboxCount,
		d.MessageCount,
		d.LogCount,
	)
}

func (d *VMProtoData) Equals(o *VMProtoData) bool {
	return d.MachineHash == o.MachineHash &&
		d.InboxTop == o.InboxTop &&
		d.InboxCount.Cmp(o.InboxCount) == 0 &&
		d.MessageCount.Cmp(o.MessageCount) == 0 &&
		d.LogCount.Cmp(o.LogCount) == 0
}

func (d *VMProtoData) Clone() *VMProtoData {
	return &VMProtoData{
		MachineHash:  d.MachineHash,
		InboxTop:     d.InboxTop,
		InboxCount:   new(big.Int).Set(d.InboxCount),
		MessageCount: new(big.Int).Set(d.MessageCount),
		LogCount:     new(big.Int).Set(d.LogCount),
	}
}

func (d *VMProtoData) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(d.MachineHash),
		hashing.Bytes32(d.InboxTop),
		hashing.Uint256(new(big.Int).Set(d.InboxCount)),
		hashing.Uint256(new(big.Int).Set(d.MessageCount)),
		hashing.Uint256(new(big.Int).Set(d.LogCount)),
	)
}

func (node *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  node.MachineHash.MarshalToBuf(),
		InboxTop:     node.InboxTop.MarshalToBuf(),
		InboxCount:   common.MarshalBigInt(node.InboxCount),
		MessageCount: common.MarshalBigInt(node.MessageCount),
		LogCount:     common.MarshalBigInt(node.LogCount),
	}
}

func (buf *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash:  buf.MachineHash.Unmarshal(),
		InboxTop:     buf.InboxTop.Unmarshal(),
		InboxCount:   buf.InboxCount.Unmarshal(),
		MessageCount: buf.MessageCount.Unmarshal(),
		LogCount:     buf.LogCount.Unmarshal(),
	}
}

type AssertionParams struct {
	NumSteps             uint64
	ImportedMessageCount *big.Int
}

func NewRandomAssertionParams() *AssertionParams {
	return &AssertionParams{
		NumSteps:             rand.Uint64(),
		ImportedMessageCount: common.RandBigInt(),
	}
}

func (ap *AssertionParams) String() string {
	return fmt.Sprintf(
		"AssertionParams(NumSteps: %v, ImportedCount: %v)",
		ap.NumSteps,
		ap.ImportedMessageCount,
	)
}

func (ap *AssertionParams) Equals(o *AssertionParams) bool {
	return ap.NumSteps == o.NumSteps &&
		ap.ImportedMessageCount.Cmp(o.ImportedMessageCount) == 0
}

func (ap *AssertionParams) Clone() *AssertionParams {
	return &AssertionParams{
		NumSteps:             ap.NumSteps,
		ImportedMessageCount: new(big.Int).Set(ap.ImportedMessageCount),
	}
}

func (ap *AssertionParams) MarshalToBuf() *AssertionParamsBuf {
	return &AssertionParamsBuf{
		NumSteps:             ap.NumSteps,
		ImportedMessageCount: common.MarshalBigInt(ap.ImportedMessageCount),
	}
}

func (m *AssertionParamsBuf) Unmarshal() *AssertionParams {
	return &AssertionParams{
		NumSteps:             m.NumSteps,
		ImportedMessageCount: m.ImportedMessageCount.Unmarshal(),
	}
}

type AssertionClaim struct {
	AfterInboxTop common.Hash
	AssertionStub *ExecutionAssertionStub
}

func NewRandomAssertionClaim(assertion *ExecutionAssertionStub) *AssertionClaim {
	return &AssertionClaim{
		AfterInboxTop: common.RandHash(),
		AssertionStub: assertion,
	}
}

func (dn *AssertionClaim) String() string {
	return fmt.Sprintf(
		"AssertionClaim(AfterInboxTop: %v, Assertion: %v)",
		dn.AfterInboxTop,
		dn.AssertionStub,
	)
}

func (dn *AssertionClaim) Equals(o *AssertionClaim) bool {
	return dn.AfterInboxTop == o.AfterInboxTop &&
		dn.AssertionStub.Equals(o.AssertionStub)
}

func (dn *AssertionClaim) Clone() *AssertionClaim {
	return &AssertionClaim{
		AfterInboxTop: dn.AfterInboxTop,
		AssertionStub: dn.AssertionStub.Clone(),
	}
}

func (dn *AssertionClaim) MarshalToBuf() *AssertionClaimBuf {
	return &AssertionClaimBuf{
		AfterInboxTop: dn.AfterInboxTop.MarshalToBuf(),
		AssertionStub: dn.AssertionStub.MarshalToBuf(),
	}
}

func (m *AssertionClaimBuf) Unmarshal() *AssertionClaim {
	return &AssertionClaim{
		AfterInboxTop: m.AfterInboxTop.Unmarshal(),
		AssertionStub: m.AssertionStub.Unmarshal(),
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

func NewRandomDisputableNode(assertion *ExecutionAssertionStub) *DisputableNode {
	return &DisputableNode{
		AssertionParams: NewRandomAssertionParams(),
		AssertionClaim:  NewRandomAssertionClaim(assertion),
		MaxInboxTop:     common.RandHash(),
		MaxInboxCount:   common.RandBigInt(),
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

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.AssertionClaim.AssertionStub.AfterMachineHash,
		dn.AssertionClaim.AfterInboxTop,
		new(big.Int).Add(prevState.InboxCount, dn.AssertionParams.ImportedMessageCount),
		new(big.Int).Add(prevState.MessageCount, new(big.Int).SetUint64(dn.AssertionClaim.AssertionStub.MessageCount)),
		new(big.Int).Add(prevState.LogCount, new(big.Int).SetUint64(dn.AssertionClaim.AssertionStub.LogCount)),
	)
}
