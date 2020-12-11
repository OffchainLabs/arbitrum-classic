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

package evm

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
)

type Provenance struct {
	L1SeqNum        *big.Int
	ParentRequestId common.Hash
	IndexInParent   *big.Int
}

func CompareProvenances(prov1 Provenance, prov2 Provenance) []string {
	var differences []string
	if prov1.L1SeqNum.Cmp(prov2.L1SeqNum) != 0 {
		differences = append(differences, fmt.Sprintf("different seq nums %v and %v", prov1.L1SeqNum, prov2.L1SeqNum))
	}
	if prov1.ParentRequestId != prov2.ParentRequestId {
		differences = append(differences, fmt.Sprintf("different parents %v and %v", prov1.ParentRequestId, prov2.ParentRequestId))
	}
	if prov1.IndexInParent.Cmp(prov2.IndexInParent) != 0 {
		differences = append(differences, fmt.Sprintf("different indexes %v and %v", prov1.IndexInParent, prov2.IndexInParent))
	}
	return differences
}

func NewProvenanceFromValue(val value.Value) (Provenance, error) {
	failRet := Provenance{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("val must be a tuple")
	}
	if tup.Len() != 3 {
		return failRet, errors.Errorf("expected tuple of length 3, but recieved tuple of length %v", tup.Len())
	}

	// Tuple size already verified above, so error can be ignored
	l1SeqNumVal, _ := tup.GetByInt64(0)
	parentRequestIdVal, _ := tup.GetByInt64(1)
	indexInParentVal, _ := tup.GetByInt64(2)

	l1SeqNumInt, ok := l1SeqNumVal.(value.IntValue)
	if !ok {
		return failRet, errors.New("provenance l1SeqNum must be an int")
	}

	parentRequestIdInt, ok := parentRequestIdVal.(value.IntValue)
	if !ok {
		return failRet, errors.New("provenance parentRequestId must be an int")
	}

	indexInParentInt, ok := indexInParentVal.(value.IntValue)
	if !ok {
		return failRet, errors.New("provenance indexInParent must be an int")
	}

	indexInParent := indexInParentInt.BigInt()
	if indexInParent.Cmp(math.MaxBig256) == 0 {
		indexInParent = nil
	}

	var parentRequestId common.Hash
	copy(parentRequestId[:], math.U256Bytes(parentRequestIdInt.BigInt()))

	return Provenance{
		L1SeqNum:        l1SeqNumInt.BigInt(),
		ParentRequestId: parentRequestId,
		IndexInParent:   indexInParent,
	}, nil
}

func (p Provenance) AsValue() value.Value {
	parent := p.IndexInParent
	if parent == nil {
		parent = math.MaxBig256
	}
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(p.L1SeqNum),
		value.NewIntValue(new(big.Int).SetBytes(p.ParentRequestId[:])),
		value.NewIntValue(parent),
	})
	return tup
}

type IncomingRequest struct {
	Kind       inbox.Type
	Sender     common.Address
	MessageID  common.Hash
	Data       []byte
	ChainTime  inbox.ChainTime
	Provenance Provenance
}

func CompareIncomingRequests(req1 IncomingRequest, req2 IncomingRequest) []string {
	var differences []string
	if req1.Kind != req2.Kind {
		differences = append(differences, fmt.Sprintf("different kinds %v and %v", req1.Kind, req2.Kind))
	}
	if req1.Sender != req2.Sender {
		differences = append(differences, fmt.Sprintf("different senders %v and %v", req1.Sender, req2.Sender))
	}
	if req1.MessageID != req2.MessageID {
		differences = append(differences, fmt.Sprintf("different message ids %v and %v", req1.MessageID, req2.MessageID))
	}
	if !bytes.Equal(req1.Data, req2.Data) {
		differences = append(differences, fmt.Sprintf("different data 0x%X and 0x%X", req1.Data, req2.Data))
	}
	if req1.ChainTime.BlockNum.Cmp(req2.ChainTime.BlockNum) != 0 {
		differences = append(differences, fmt.Sprintf("different block nums %v and %v", req1.ChainTime.BlockNum, req2.ChainTime.BlockNum))
	}
	if req1.ChainTime.Timestamp.Cmp(req2.ChainTime.Timestamp) != 0 {
		differences = append(differences, fmt.Sprintf("different timestamps %v and %v", req1.ChainTime.Timestamp, req2.ChainTime.Timestamp))
	}
	differences = append(differences, CompareProvenances(req1.Provenance, req2.Provenance)...)
	return differences
}

func NewIncomingRequestFromValue(val value.Value) (IncomingRequest, error) {
	failRet := IncomingRequest{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("val must be a tuple")
	}
	if tup.Len() != 7 {
		return failRet, errors.Errorf("expected tuple of length 7, but recieved tuple of length %v", tup.Len())
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	blockNumber, _ := tup.GetByInt64(1)
	timestamp, _ := tup.GetByInt64(2)
	sender, _ := tup.GetByInt64(3)
	inboxSeqNum, _ := tup.GetByInt64(4)
	messageData, _ := tup.GetByInt64(5)
	provenanceVal, _ := tup.GetByInt64(6)

	kindInt, ok := kind.(value.IntValue)
	if !ok {
		return failRet, errors.New("inbox message kind must be an int")
	}

	blockNumberInt, ok := blockNumber.(value.IntValue)
	if !ok {
		return failRet, errors.New("blockNumber must be an int")
	}

	timestampInt, ok := timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("timestamp must be an int")
	}

	senderInt, ok := sender.(value.IntValue)
	if !ok {
		return failRet, errors.New("sender must be an int")
	}

	messageIDInt, ok := inboxSeqNum.(value.IntValue)
	if !ok {
		return failRet, errors.New("inboxSeqNum must be an int")
	}
	var messageID common.Hash
	copy(messageID[:], math.U256Bytes(messageIDInt.BigInt()))

	data, err := inbox.ByteStackToHex(messageData)
	if err != nil {
		return failRet, errors.Wrap(err, "unmarshalling input data")
	}

	provenance, err := NewProvenanceFromValue(provenanceVal)
	if err != nil {
		return failRet, err
	}

	return IncomingRequest{
		Kind:      inbox.Type(kindInt.BigInt().Uint64()),
		Sender:    inbox.NewAddressFromInt(senderInt),
		MessageID: messageID,
		Data:      data,
		ChainTime: inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(blockNumberInt.BigInt()),
			Timestamp: timestampInt.BigInt(),
		},
		Provenance: provenance,
	}, nil
}

func NewRandomIncomingRequest() IncomingRequest {
	return IncomingRequest{
		Kind:      inbox.Type(rand.Uint32()),
		Sender:    common.RandAddress(),
		MessageID: common.RandHash(),
		Data:      common.RandBytes(200),
		ChainTime: inbox.NewRandomChainTime(),
	}
}

func (ir IncomingRequest) AsValue() value.Value {
	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(int64(ir.Kind)),
		value.NewIntValue(ir.ChainTime.BlockNum.AsInt()),
		value.NewIntValue(ir.ChainTime.Timestamp),
		inbox.NewIntFromAddress(ir.Sender),
		value.NewIntValue(new(big.Int).SetBytes(ir.MessageID[:])),
		inbox.BytesToByteStack(ir.Data),
		ir.Provenance.AsValue(),
	})
	return tup
}
