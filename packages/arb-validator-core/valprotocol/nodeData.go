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
	InvalidExecutionChildType ChildType = 1
	ValidChildType            ChildType = 2

	MinChildType        ChildType = 0
	MaxInvalidChildType ChildType = 1
	MaxChildType        ChildType = 2
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

func (d *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  d.MachineHash.MarshalToBuf(),
		InboxTop:     d.InboxTop.MarshalToBuf(),
		InboxCount:   common.MarshalBigInt(d.InboxCount),
		MessageCount: common.MarshalBigInt(d.MessageCount),
		LogCount:     common.MarshalBigInt(d.LogCount),
	}
}

func (x *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash:  x.MachineHash.Unmarshal(),
		InboxTop:     x.InboxTop.Unmarshal(),
		InboxCount:   x.InboxCount.Unmarshal(),
		MessageCount: x.MessageCount.Unmarshal(),
		LogCount:     x.LogCount.Unmarshal(),
	}
}

type AssertionParams struct {
	NumSteps             uint64
	ImportedMessageCount *big.Int
}

func (ct ChildType) String() string {
	switch ct {
	case InvalidInboxTopChildType:
		return "InvalidInboxTopChildType"
	case InvalidExecutionChildType:
		return "InvalidExecutionChildType"
	case ValidChildType:
		return "ValidChildType"
	}

	return "UnknownChildType"
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

func (x *AssertionParamsBuf) Unmarshal() *AssertionParams {
	return &AssertionParams{
		NumSteps:             x.NumSteps,
		ImportedMessageCount: x.ImportedMessageCount.Unmarshal(),
	}
}

type DisputableNode struct {
	AssertionParams *AssertionParams
	Assertion       *ExecutionAssertionStub
	MaxInboxTop     common.Hash
	MaxInboxCount   *big.Int
}

func NewDisputableNode(
	assertionParams *AssertionParams,
	assertion *ExecutionAssertionStub,
	maxInboxTop common.Hash,
	maxInboxCount *big.Int,
) *DisputableNode {
	return &DisputableNode{
		AssertionParams: assertionParams,
		Assertion:       assertion,
		MaxInboxTop:     maxInboxTop,
		MaxInboxCount:   maxInboxCount,
	}
}

func NewRandomDisputableNode(assertion *ExecutionAssertionStub) *DisputableNode {
	return &DisputableNode{
		AssertionParams: NewRandomAssertionParams(),
		Assertion:       assertion,
		MaxInboxTop:     common.RandHash(),
		MaxInboxCount:   common.RandBigInt(),
	}
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		AssertionParams: dn.AssertionParams.MarshalToBuf(),
		Assertion:       dn.Assertion.MarshalToBuf(),
		MaxInboxTop:     dn.MaxInboxTop.MarshalToBuf(),
		MaxInboxCount:   common.MarshalBigInt(dn.MaxInboxCount),
	}
}

func (x *DisputableNodeBuf) Unmarshal() *DisputableNode {
	return NewDisputableNode(
		x.AssertionParams.Unmarshal(),
		x.Assertion.Unmarshal(),
		x.MaxInboxTop.Unmarshal(),
		x.MaxInboxCount.Unmarshal(),
	)
}

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.Assertion.AfterMachineHash,
		dn.Assertion.AfterInboxAcc,
		new(big.Int).Add(prevState.InboxCount, dn.AssertionParams.ImportedMessageCount),
		new(big.Int).Add(prevState.MessageCount, new(big.Int).SetUint64(dn.Assertion.MessageCount)),
		new(big.Int).Add(prevState.LogCount, new(big.Int).SetUint64(dn.Assertion.LogCount)),
	)
}
