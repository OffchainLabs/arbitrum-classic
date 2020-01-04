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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/utils"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

type VMProtoData struct {
	MachineHash  [32]byte
	InboxHash    [32]byte
	PendingTop   [32]byte
	PendingCount *big.Int
}

func NewVMProtoData(
	machineHash [32]byte,
	inboxHash [32]byte,
	pendingTop [32]byte,
	pendingCount *big.Int,
) *VMProtoData {
	return &VMProtoData{
		MachineHash:  machineHash,
		InboxHash:    inboxHash,
		PendingTop:   pendingTop,
		PendingCount: pendingCount,
	}
}

func (d *VMProtoData) Equals(o *VMProtoData) bool {
	return d.MachineHash == o.MachineHash &&
		d.InboxHash == o.InboxHash &&
		d.PendingTop == o.PendingTop &&
		d.PendingCount.Cmp(o.PendingCount) == 0
}

func (d *VMProtoData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(d.MachineHash),
		solsha3.Bytes32(d.InboxHash),
		solsha3.Bytes32(d.PendingTop),
		solsha3.Uint256(d.PendingCount),
	))
	return ret
}

func (node *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  utils.MarshalHash(node.MachineHash),
		InboxHash:    utils.MarshalHash(node.InboxHash),
		PendingTop:   utils.MarshalHash(node.PendingTop),
		PendingCount: utils.MarshalBigInt(node.PendingCount),
	}
}

func (buf *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		MachineHash:  utils.UnmarshalHash(buf.MachineHash),
		InboxHash:    utils.UnmarshalHash(buf.InboxHash),
		PendingTop:   utils.UnmarshalHash(buf.PendingTop),
		PendingCount: utils.UnmarshalBigInt(buf.PendingCount),
	}
}

type AssertionParams struct {
	NumSteps             uint32
	TimeBoundsBlocks     [2]RollupTime
	ImportedMessageCount *big.Int
}

func (dn *AssertionParams) MarshalToBuf() *AssertionParamsBuf {
	return &AssertionParamsBuf{
		NumSteps:             dn.NumSteps,
		TimeLowerBound:       dn.TimeBoundsBlocks[0].MarshalToBuf(),
		TimeUpperBound:       dn.TimeBoundsBlocks[1].MarshalToBuf(),
		ImportedMessageCount: utils.MarshalBigInt(dn.ImportedMessageCount),
	}
}

func (m *AssertionParamsBuf) Unmarshal() *AssertionParams {
	return &AssertionParams{
		NumSteps:             m.NumSteps,
		TimeBoundsBlocks:     [2]RollupTime{m.TimeLowerBound.Unmarshal(), m.TimeUpperBound.Unmarshal()},
		ImportedMessageCount: utils.UnmarshalBigInt(m.ImportedMessageCount),
	}
}

type AssertionClaim struct {
	AfterPendingTop       [32]byte
	ImportedMessagesSlice [32]byte
	AssertionStub         *protocol.ExecutionAssertionStub
}

func (dn *AssertionClaim) MarshalToBuf() *AssertionClaimBuf {
	return &AssertionClaimBuf{
		AfterPendingTop:       utils.MarshalHash(dn.AfterPendingTop),
		ImportedMessagesSlice: utils.MarshalHash(dn.ImportedMessagesSlice),
		AssertionStub:         dn.AssertionStub,
	}
}

func (m *AssertionClaimBuf) Unmarshal() *AssertionClaim {
	return &AssertionClaim{
		AfterPendingTop:       utils.UnmarshalHash(m.AfterPendingTop),
		ImportedMessagesSlice: utils.UnmarshalHash(m.ImportedMessagesSlice),
		AssertionStub:         m.AssertionStub,
	}
}

type DisputableNode struct {
	assertionParams *AssertionParams
	assertionClaim  *AssertionClaim
	hash            [32]byte
}

func NewDisputableNode(
	assertionParams *AssertionParams,
	assertionClaim *AssertionClaim,
) *DisputableNode {
	ret := &DisputableNode{
		assertionParams: assertionParams,
		assertionClaim:  assertionClaim,
	}
	ret.hash = ret._hash()
	return ret
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		AssertionParams: dn.assertionParams.MarshalToBuf(),
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	return NewDisputableNode(
		buf.AssertionParams.Unmarshal(),
		buf.AssertionClaim.Unmarshal(),
	)
}

func (dn *DisputableNode) Hash() [32]byte {
	return dn.Hash()
}

func (dn *DisputableNode) _hash() [32]byte {
	// Hash calculation is incorrect
	var ret [32]byte
	retSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(utils.UnmarshalHash(dn.assertionClaim.AssertionStub.AfterHash)),
		solsha3.Bool(dn.assertionClaim.AssertionStub.DidInboxInsn),
		//solsha3.Uint32(dn.numSteps),
		solsha3.Uint64(dn.assertionClaim.AssertionStub.NumGas),
		solsha3.Bytes32(utils.UnmarshalHash(dn.assertionClaim.AssertionStub.FirstMessageHash)),
		solsha3.Bytes32(utils.UnmarshalHash(dn.assertionClaim.AssertionStub.LastMessageHash)),
		solsha3.Bytes32(utils.UnmarshalHash(dn.assertionClaim.AssertionStub.FirstLogHash)),
		solsha3.Bytes32(utils.UnmarshalHash(dn.assertionClaim.AssertionStub.LastLogHash)),
	)
	copy(ret[:], retSlice)
	return ret
}

func (dn *DisputableNode) ValidAfterVMProtoData(prevState *VMProtoData) *VMProtoData {
	return NewVMProtoData(
		dn.assertionClaim.AssertionStub.AfterHashValue(),
		value.NewEmptyTuple().Hash(),
		dn.assertionClaim.AfterPendingTop,
		new(big.Int).Add(prevState.PendingCount, dn.assertionParams.ImportedMessageCount),
	)
}
