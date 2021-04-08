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
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
)

func NewValueFromOptional(val value.Value) (value.Value, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return nil, errors.New("optional must be a tuple")
	}
	if tup.Len() == 0 {
		return nil, errors.New("optional too short")
	}
	hasValue, _ := tup.GetByInt64(0)
	hasValueInt, ok := hasValue.(value.IntValue)
	if !ok {
		return nil, errors.New("hasValue must be an int")
	}
	if hasValueInt.BigInt().Uint64() == 0 {
		return nil, nil
	}
	if hasValueInt.BigInt().Uint64() != 1 {
		return nil, errors.New("optional had unknown code")
	}
	if tup.Len() != 2 {
		return nil, errors.New("optional with value wrong length")
	}
	nestedVal, _ := tup.GetByInt64(1)
	return nestedVal, nil
}

type AggregatorInfo struct {
	Aggregator    *common.Address
	CalldataBytes *big.Int
}

func NewAggregatorInfoFromOptionalValue(val value.Value) (*AggregatorInfo, error) {
	nestedVal, err := NewValueFromOptional(val)
	if err != nil {
		return nil, err
	}
	if nestedVal == nil {
		return nil, nil
	}
	nestedTup, ok := nestedVal.(*value.TupleValue)
	if !ok || nestedTup.Len() != 2 {
		return nil, errors.Errorf("expected tuple of length 2, but recieved %v", nestedVal)
	}

	aggregatorVal, _ := nestedTup.GetByInt64(0)
	calldataBytes, _ := nestedTup.GetByInt64(1)

	var aggAddress *common.Address

	// ArbOS version upgrade from https://github.com/OffchainLabs/arb-os/pull/429
	// Support aggregator field either as an address or and optional address
	if _, ok := aggregatorVal.(value.IntValue); !ok {
		aggregatorVal, err = NewValueFromOptional(aggregatorVal)
		if err != nil {
			return nil, err
		}
	}

	if aggregatorVal != nil {
		aggregatorInt, ok := aggregatorVal.(value.IntValue)
		if !ok {
			return nil, errors.New("aggregator must be an int")
		}
		rawAggregatorAddress := inbox.NewAddressFromInt(aggregatorInt)
		blankAddress := common.Address{}
		if rawAggregatorAddress != blankAddress {
			aggAddress = &rawAggregatorAddress
		}
	}
	calldataBytesInt, ok := calldataBytes.(value.IntValue)
	if !ok {
		return nil, errors.New("calldataBytes must be an int")
	}
	return &AggregatorInfo{
		Aggregator:    aggAddress,
		CalldataBytes: calldataBytesInt.BigInt(),
	}, nil
}

func newEmptyOptional() value.Value {
	tup, _ := value.NewTupleFromSlice([]value.Value{value.NewInt64Value(0)})
	return tup
}

func newOptional(val value.Value) value.Value {
	if val == nil {
		tup, _ := value.NewTupleFromSlice([]value.Value{value.NewInt64Value(0)})
		return tup
	}
	tup, _ := value.NewTupleFromSlice([]value.Value{value.NewInt64Value(1), val})
	return tup
}

func (a *AggregatorInfo) AsOptionalValue() value.Value {
	var val value.Value
	if a != nil {
		var aggVal value.Value
		if a.Aggregator != nil {
			aggVal = inbox.NewIntFromAddress(*a.Aggregator)
		}

		val := value.NewTuple2(
			newOptional(aggVal),
			value.NewIntValue(a.CalldataBytes),
		)
		tup, _ := value.NewTupleFromSlice([]value.Value{val})
		val = tup
	}
	return newOptional(val)
}

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
		return failRet, errors.New("provenance must be a tuple")
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

type IncomingRequest struct {
	Kind           inbox.Type
	Sender         common.Address
	MessageID      common.Hash
	Data           []byte
	L1BlockNumber  *big.Int
	L2BlockNumber  *big.Int
	L2Timestamp    *big.Int
	Provenance     Provenance
	AggregatorInfo *AggregatorInfo
	AdminMode      bool
}

func (r IncomingRequest) String() string {
	return fmt.Sprintf(
		"IncomingRequest(kind=%v, sender=%v, id=%v, data=%v, l1Block=%v, l2Block=%v, timestamp=%v, provenance=%v, aggregator=%v)",
		r.Kind,
		r.Sender,
		r.MessageID,
		hexutil.Encode(r.Data),
		r.L1BlockNumber,
		r.L2BlockNumber,
		r.L2Timestamp,
		r.Provenance,
		r.AggregatorInfo,
	)
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
	if req1.L1BlockNumber.Cmp(req2.L1BlockNumber) != 0 {
		differences = append(differences, fmt.Sprintf("different L1BlockNumber nums %v and %v", req1.L1BlockNumber, req2.L1BlockNumber))
	}
	if req1.L2BlockNumber.Cmp(req2.L2BlockNumber) != 0 {
		differences = append(differences, fmt.Sprintf("different L2BlockNumber nums %v and %v", req1.L2BlockNumber, req2.L2BlockNumber))
	}
	if req1.L2Timestamp.Cmp(req2.L2Timestamp) != 0 {
		differences = append(differences, fmt.Sprintf("different L2Timestamp nums %v and %v", req1.L2Timestamp, req2.L2Timestamp))
	}
	differences = append(differences, CompareProvenances(req1.Provenance, req2.Provenance)...)
	return differences
}

func NewIncomingRequestFromValue(val value.Value) (IncomingRequest, error) {
	failRet := IncomingRequest{}
	tup, ok := val.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("incoming request must be a tuple")
	}
	if tup.Len() != 8 {
		return failRet, errors.Errorf("expected incoming request to be tuple of length 8, but recieved tuple of length %v", tup.Len())
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	l2BlockNumber, _ := tup.GetByInt64(1)
	l1BlockNumber, _ := tup.GetByInt64(2)
	l2Timestamp, _ := tup.GetByInt64(3)
	sender, _ := tup.GetByInt64(4)
	inboxSeqNum, _ := tup.GetByInt64(5)
	messageData, _ := tup.GetByInt64(6)
	remVal, _ := tup.GetByInt64(7)

	remTup, ok := remVal.(*value.TupleValue)
	if !ok {
		return failRet, errors.New("remaining incoming request values must be a tuple")
	}
	if remTup.Len() != 3 {
		return failRet, errors.Errorf("expected incoming request remaining values to be tuple of length 3, but received tuple of length %v", remTup.Len())
	}
	provenanceVal, _ := remTup.GetByInt64(0)
	aggregatorInfoVal, _ := remTup.GetByInt64(1)
	adminModeVal, _ := remTup.GetByInt64(2)

	kindInt, ok := kind.(value.IntValue)
	if !ok {
		return failRet, errors.New("inbox message kind must be an int")
	}

	l1BlockNumberInt, ok := l1BlockNumber.(value.IntValue)
	if !ok {
		return failRet, errors.New("l1BlockNumber must be an int")
	}

	l2BlockNumberInt, ok := l2BlockNumber.(value.IntValue)
	if !ok {
		return failRet, errors.New("l2BlockNumber must be an int")
	}

	l2TimestampInt, ok := l2Timestamp.(value.IntValue)
	if !ok {
		return failRet, errors.New("l2Timestamp must be an int")
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

	data, err := inbox.ByteArrayToBytes(messageData)
	if err != nil {
		return failRet, errors.Wrap(err, "unmarshalling input data")
	}

	provenance, err := NewProvenanceFromValue(provenanceVal)
	if err != nil {
		return failRet, err
	}

	aggregatorInfo, err := NewAggregatorInfoFromOptionalValue(aggregatorInfoVal)
	if err != nil {
		return failRet, err
	}

	adminModeInt, ok := adminModeVal.(value.IntValue)
	if !ok {
		return failRet, errors.New("adminMode must be a boolean")
	}
	var adminMode bool
	if adminModeInt.Equal(value.NewInt64Value(0)) {
		adminMode = false
	} else if adminModeInt.Equal(value.NewInt64Value(1)) {
		adminMode = true
	} else {
		return failRet, errors.Errorf("expected adminMode to be an integer either 0 or 1, but received integer %v", adminModeInt)
	}

	return IncomingRequest{
		Kind:           inbox.Type(kindInt.BigInt().Uint64()),
		Sender:         inbox.NewAddressFromInt(senderInt),
		MessageID:      messageID,
		Data:           data,
		L1BlockNumber:  l1BlockNumberInt.BigInt(),
		L2BlockNumber:  l2BlockNumberInt.BigInt(),
		L2Timestamp:    l2TimestampInt.BigInt(),
		Provenance:     provenance,
		AggregatorInfo: aggregatorInfo,
		AdminMode:      adminMode,
	}, nil
}

func NewRandomIncomingRequest() IncomingRequest {
	return IncomingRequest{
		Kind:          inbox.Type(rand.Uint32()),
		Sender:        common.RandAddress(),
		MessageID:     common.RandHash(),
		Data:          common.RandBytes(200),
		L1BlockNumber: common.RandBigInt(),
		L2BlockNumber: common.RandBigInt(),
		L2Timestamp:   common.RandBigInt(),
	}
}
